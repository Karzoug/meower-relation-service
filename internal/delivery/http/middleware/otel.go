package middleware

import (
	"context"
	"net/http"

	"github.com/oapi-codegen/runtime/strictmiddleware/nethttp"
	"go.opentelemetry.io/otel/trace"

	gen "github.com/Karzoug/meower-relation-service/internal/delivery/http/gen/relation/v1"
	"github.com/Karzoug/meower-relation-service/pkg/trace/otlp"
)

// Otel starts the otel tracing and stores the trace id in the context.
func Otel(tracer trace.Tracer) gen.StrictMiddlewareFunc {
	return func(f nethttp.StrictHTTPHandlerFunc, operationID string) nethttp.StrictHTTPHandlerFunc {
		return func(ctx context.Context, w http.ResponseWriter, r *http.Request, request any) (response any, err error) {
			return f(otlp.InjectTracing(ctx, tracer), w, r, request)
		}
	}
}
