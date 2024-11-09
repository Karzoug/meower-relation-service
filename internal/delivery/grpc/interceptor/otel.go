package interceptor

import (
	"context"

	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"

	"github.com/Karzoug/meower-relation-service/pkg/trace/otlp"
)

func Otel(tracer trace.Tracer) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		return handler(otlp.InjectTracing(ctx, tracer), req)
	}
}
