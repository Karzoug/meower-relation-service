package relation

import (
	"context"
	"net/http"

	"github.com/rs/zerolog"
	"go.opentelemetry.io/otel/trace"

	gen "github.com/Karzoug/meower-relation-service/internal/delivery/http/gen/relation/v1"
	"github.com/Karzoug/meower-relation-service/internal/delivery/http/handler/errfunc"
	zerologHook "github.com/Karzoug/meower-relation-service/internal/delivery/http/handler/zerolog"
	"github.com/Karzoug/meower-relation-service/internal/delivery/http/middleware"
	"github.com/Karzoug/meower-relation-service/internal/relation/service"
	"github.com/Karzoug/meower-relation-service/pkg/auth"
)

const baseURL = "/api/v1"

var _ gen.StrictServerInterface = handlers{}

func RoutesFunc(rs service.RelationService, tracer trace.Tracer, logger zerolog.Logger) func(mux *http.ServeMux) {
	logger = logger.With().
		Str("component", "http server: relation handlers").
		Logger().
		Hook(zerologHook.TraceIDHook())

	hdl := handlers{
		relationService: rs,
		logger:          logger,
	}

	return func(mux *http.ServeMux) {
		gen.HandlerWithOptions(
			gen.NewStrictHandlerWithOptions(hdl,
				[]gen.StrictMiddlewareFunc{
					middleware.Recover,
					middleware.AuthN,
					middleware.Error(logger),
					middleware.Logger(logger),
					middleware.Otel(tracer),
				},
				gen.StrictHTTPServerOptions{
					RequestErrorHandlerFunc:  errfunc.JSONRequest(logger, tracer),
					ResponseErrorHandlerFunc: errfunc.JSONResponse(logger),
				}),
			gen.StdHTTPServerOptions{
				BaseURL:    baseURL,
				BaseRouter: mux,
			})
	}
}

type handlers struct {
	relationService service.RelationService
	logger          zerolog.Logger
}

// DeleteRelationHiddenUserID implements v1.StrictServerInterface.
func (h handlers) DeleteRelationHiddenUserID(ctx context.Context, request gen.DeleteRelationHiddenUserIDRequestObject) (gen.DeleteRelationHiddenUserIDResponseObject, error) {
	if err := h.relationService.Unhide(ctx,
		auth.UserIDFromContext(ctx),
		request.UserID,
	); err != nil {
		return nil, err
	}

	return gen.DeleteRelationHiddenUserID200Response{}, nil
}

// DeleteRelationFollowingsUserID implements v1.StrictServerInterface.
func (h handlers) DeleteRelationFollowingsUserID(ctx context.Context, request gen.DeleteRelationFollowingsUserIDRequestObject) (gen.DeleteRelationFollowingsUserIDResponseObject, error) {
	if err := h.relationService.Unfollow(ctx,
		auth.UserIDFromContext(ctx),
		request.UserID,
	); err != nil {
		return nil, err
	}

	return gen.DeleteRelationFollowingsUserID200Response{}, nil
}

// PostRelationFollowings implements v1.StrictServerInterface.
func (h handlers) PostRelationFollowings(ctx context.Context, request gen.PostRelationFollowingsRequestObject) (gen.PostRelationFollowingsResponseObject, error) {
	if err := h.relationService.Follow(ctx,
		auth.UserIDFromContext(ctx),
		request.Body.Id,
	); err != nil {
		return nil, err
	}

	return gen.PostRelationFollowings201Response{}, nil
}

// PostRelationHidden implements v1.StrictServerInterface.
func (h handlers) PostRelationHidden(ctx context.Context, request gen.PostRelationHiddenRequestObject) (gen.PostRelationHiddenResponseObject, error) {
	if err := h.relationService.Hide(ctx,
		auth.UserIDFromContext(ctx),
		request.Body.Id,
	); err != nil {
		return nil, err
	}

	return gen.PostRelationHidden201Response{}, nil
}
