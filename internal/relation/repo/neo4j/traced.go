package neo4j

import (
	"context"

	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"

	"github.com/Karzoug/meower-relation-service/internal/relation/entity"
)

type tracedRepo struct {
	repo
	tracer trace.Tracer
}

func (tr tracedRepo) Follow(ctx context.Context, reqUserID, targetUserID string) (err error) {
	const operationID = "neo4j.Follow"
	ctx, span := tr.tracer.Start(ctx, operationID, trace.WithSpanKind(trace.SpanKindClient))
	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, operationID+": operation failed")
			span.RecordError(err)
		}
		span.End()
	}()

	return tr.repo.Follow(ctx, reqUserID, targetUserID)
}

func (tr tracedRepo) Unfollow(ctx context.Context, reqUserID, targetUserID string) (err error) {
	const operationID = "neo4j.Unfollow"
	ctx, span := tr.tracer.Start(ctx, operationID, trace.WithSpanKind(trace.SpanKindClient))
	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, operationID+": operation failed")
			span.RecordError(err)
		}
		span.End()
	}()

	return tr.repo.Unfollow(ctx, reqUserID, targetUserID)
}

func (tr tracedRepo) Hide(ctx context.Context, reqUserID, targetUserID string) (err error) {
	const operationID = "neo4j.Hide"
	ctx, span := tr.tracer.Start(ctx, operationID, trace.WithSpanKind(trace.SpanKindClient))
	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, operationID+": operation failed")
			span.RecordError(err)
		}
		span.End()
	}()

	return tr.repo.Hide(ctx, reqUserID, targetUserID)
}

func (tr tracedRepo) Unhide(ctx context.Context, reqUserID, targetUserID string) (err error) {
	const operationID = "neo4j.Unhide"
	ctx, span := tr.tracer.Start(ctx, operationID, trace.WithSpanKind(trace.SpanKindClient))
	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, operationID+": operation failed")
			span.RecordError(err)
		}
		span.End()
	}()

	return tr.repo.Unhide(ctx, reqUserID, targetUserID)
}

func (tr tracedRepo) ListFollowers(ctx context.Context,
	reqUserID, targetUserID, afterID string, limit int,
) (users []entity.User, token *string, err error) {
	const operationID = "neo4j.ListFollowers"
	ctx, span := tr.tracer.Start(ctx, operationID, trace.WithSpanKind(trace.SpanKindClient))
	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, operationID+": operation failed")
			span.RecordError(err)
		}
		span.End()
	}()

	return tr.repo.ListFollowers(ctx, reqUserID, targetUserID, afterID, limit)
}

func (tr tracedRepo) ListFollowings(ctx context.Context,
	reqUserID, targetUserID, afterID string, limit int,
) (users []entity.User, token *string, err error) {
	const operationID = "neo4j.ListFollowings"
	ctx, span := tr.tracer.Start(ctx, operationID, trace.WithSpanKind(trace.SpanKindClient))
	defer func() {
		if err != nil {
			span.SetStatus(codes.Error, operationID+": operation failed")
			span.RecordError(err)
		}
		span.End()
	}()

	return tr.repo.ListFollowings(ctx, reqUserID, targetUserID, afterID, limit)
}
