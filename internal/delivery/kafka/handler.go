package kafka

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/cenkalti/backoff/v4"
	"github.com/rs/xid"
	"github.com/rs/zerolog"
	ocodes "go.opentelemetry.io/otel/codes"
	"google.golang.org/grpc/codes"

	"github.com/Karzoug/meower-common-go/ucerr"

	gen "github.com/Karzoug/meower-relation-service/pkg/proto/kafka/user/v1"
)

const (
	defaultOperationTimeout   = 5 * time.Second
	maxRetryTimeoutBeforeExit = 60 * time.Second
	preffixSpanName           = "RelationService.KafkaConsumer/"
)

func (c consumer) userCreatedHandler(ctx context.Context, event *gen.ChangedEvent, logger zerolog.Logger) error {
	logger.Info().Msg("received message")

	id, err := xid.FromString(event.Id)
	if err != nil {
		return fmt.Errorf("bug: invalid user id format: %w", err)
	}

	ctx, span := c.tracer.Start(ctx, preffixSpanName+"userCreate")
	defer span.End()

	operation := func() error {
		ctx, cancel := context.WithTimeout(ctx, defaultOperationTimeout)
		defer cancel()

		if err := c.relationService.CreateUser(ctx, id); err != nil {
			var serr ucerr.Error
			if errors.As(err, &serr) {
				if serr.Code() == codes.AlreadyExists {
					return nil
				}
				logger.Error().
					Str("user_id", id.String()).
					Err(serr.Unwrap()).
					Msg("create user failed")
			} else {
				logger.Error().
					Str("user_id", id.String()).
					Err(err).
					Msg("create user failed")
			}

			return err
		}

		return nil
	}
	if err := backoff.Retry(operation,
		backoff.NewExponentialBackOff(
			backoff.WithMaxElapsedTime(maxRetryTimeoutBeforeExit),
		),
	); err != nil {
		span.RecordError(err)
		span.SetStatus(ocodes.Error, "all operation retries failed")
		return err
	}

	logger.Info().
		Ctx(ctx).
		Str("created_user_id", id.String()).
		Msg("processed message")

	return nil
}

func (c consumer) userDeletedHandler(ctx context.Context, event *gen.ChangedEvent, logger zerolog.Logger) error {
	logger.Info().Msg("received message")

	id, err := xid.FromString(event.Id)
	if err != nil {
		return fmt.Errorf("bug: invalid user id format: %w", err)
	}

	ctx, span := c.tracer.Start(ctx, preffixSpanName+"userDelete")
	defer span.End()

	operation := func() error {
		ctx, cancel := context.WithTimeout(ctx, defaultOperationTimeout)
		defer cancel()

		if err := c.relationService.DeleteUser(ctx, id); err != nil {
			var serr ucerr.Error
			if errors.As(err, &serr) {
				if serr.Code() == codes.AlreadyExists {
					return nil
				}
				logger.Error().
					Str("user_id", id.String()).
					Err(serr.Unwrap()).
					Msg("delete user failed")
			} else {
				logger.Error().
					Str("user_id", id.String()).
					Err(err).
					Msg("delete user failed")
			}

			return err
		}

		return nil
	}
	if err := backoff.Retry(operation,
		backoff.NewExponentialBackOff(
			backoff.WithMaxElapsedTime(maxRetryTimeoutBeforeExit),
		),
	); err != nil {
		span.RecordError(err)
		span.SetStatus(ocodes.Error, "all operation retries failed")
		return err
	}

	logger.Info().
		Ctx(ctx).
		Str("deleted_user_id", id.String()).
		Msg("processed message")

	return nil
}
