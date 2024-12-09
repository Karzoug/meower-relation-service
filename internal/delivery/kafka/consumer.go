package kafka

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/rs/zerolog"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/protobuf/proto"

	ck "github.com/Karzoug/meower-common-go/kafka"
	"github.com/Karzoug/meower-common-go/trace/otlp"
	"github.com/Karzoug/meower-relation-service/internal/relation/service"
	gen "github.com/Karzoug/meower-relation-service/pkg/proto/kafka/user/v1"
)

const userTopic = "users"

type consumer struct {
	c               *kafka.Consumer
	relationService service.RelationService
	tracer          trace.Tracer
	logger          zerolog.Logger
}

func NewConsumer(ctx context.Context, cfg Config, service service.RelationService, tracer trace.Tracer, logger zerolog.Logger) (consumer, error) {
	const op = "create kafka consumer"

	logger = logger.With().
		Str("component", "kafka consumer").
		Logger()

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":        cfg.Brokers,
		"group.id":                 cfg.GroupID,
		"auto.offset.reset":        "earliest",
		"auto.commit.interval.ms":  cfg.CommitIntervalMilliseconds,
		"enable.auto.offset.store": false,
	})
	if err != nil {
		return consumer{}, fmt.Errorf("%s: %w", op, err)
	}

	var (
		topic   = userTopic
		timeout int
	)
	if t, ok := ctx.Deadline(); ok {
		timeout = int(time.Until(t).Milliseconds())
	} else {
		timeout = 500
	}

	// analog PING here
	_, err = c.GetMetadata(&topic, false, timeout)
	if err != nil {
		return consumer{}, fmt.Errorf("%s: failed to get metadata: %w", op, err)
	}

	return consumer{
		c:               c,
		relationService: service,
		tracer:          tracer,
		logger:          logger,
	}, nil
}

func (c consumer) Run(ctx context.Context) (err error) {
	userChangedEventFngpnt := ck.MessageTypeHeaderValue(&gen.ChangedEvent{})

	defer func() {
		if defErr := c.c.Close(); defErr != nil {
			err = errors.Join(err,
				fmt.Errorf("failed to close consumer: %w", defErr))
		}
	}()

	if err := c.c.Subscribe(userTopic, nil); err != nil {
		return err
	}

	run := true
	for run {
		select {
		case <-ctx.Done():
			run = false
		default:
			msg, err := c.c.ReadMessage(100 * time.Millisecond)
			if err != nil {
				var kafkaErr kafka.Error
				if errors.As(err, &kafkaErr) {
					if kafkaErr.IsFatal() {
						return fmt.Errorf("fatal error while read message: %w", err)
					}
					if !kafkaErr.IsTimeout() {
						c.logger.Warn().
							Err(err).
							Msg("failed to read message")
					}
				}
				continue
			}

			if len(msg.Headers) == 0 {
				c.storeOffset(msg)
				continue
			}
			eventType, ok := lookupHeaderValue(msg.Headers, ck.MessageTypeHeaderKey)
			if !ok {
				c.storeOffset(msg)
				continue
			}

			eventTypeFngpnt := string(eventType)

			ctx = otlp.InjectTracing(ctx, c.tracer)
			hlogger := c.logger.With().
				Str("topic", *msg.TopicPartition.Topic).
				Str("key", string(msg.Key)).
				Str("event fingerprint", eventTypeFngpnt).
				Ctx(ctx).
				Logger()

			if eventTypeFngpnt == userChangedEventFngpnt {

				event := &gen.ChangedEvent{}
				if err := proto.Unmarshal(msg.Value, event); err != nil {
					return fmt.Errorf("failed to deserialize payload: %w", err)
				}

				switch event.ChangeType {
				case gen.ChangeType_CHANGE_TYPE_CREATED:
					err = c.userCreatedHandler(ctx, event, hlogger)
				case gen.ChangeType_CHANGE_TYPE_DELETED:
					err = c.userDeletedHandler(ctx, event, hlogger)
				}
			}

			if err != nil {
				// log outside, not store offset, return from consumer with error
				return err
			}

			c.storeOffset(msg)
		}
	}

	return nil
}

func (c consumer) storeOffset(msg *kafka.Message) {
	_, err := c.c.StoreMessage(msg)
	if err != nil {
		c.logger.Error().
			Err(err).
			Str("topic", *msg.TopicPartition.Topic).
			Str("key", string(msg.Key)).
			Msg("failed to store offset after message")
	}
}

func lookupHeaderValue(headers []kafka.Header, key string) ([]byte, bool) {
	for _, header := range headers {
		if header.Key == key {
			return header.Value, true
		}
	}
	return nil, false
}
