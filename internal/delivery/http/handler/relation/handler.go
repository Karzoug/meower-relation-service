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

// PostRelationUserIDFollow implements v1.StrictServerInterface.
func (h handlers) PostRelationUserIDFollow(ctx context.Context, request gen.PostRelationUserIDFollowRequestObject) (gen.PostRelationUserIDFollowResponseObject, error) {
	if err := h.relationService.Follow(ctx,
		auth.UserIDFromContext(ctx),
		request.UserID,
	); err != nil {
		return nil, err
	}

	return gen.PostRelationUserIDFollow200Response{}, nil
}

// PostRelationUserIDUnfollow implements v1.StrictServerInterface.
func (h handlers) PostRelationUserIDUnfollow(ctx context.Context, request gen.PostRelationUserIDUnfollowRequestObject) (gen.PostRelationUserIDUnfollowResponseObject, error) {
	if err := h.relationService.Unfollow(ctx,
		auth.UserIDFromContext(ctx),
		request.UserID,
	); err != nil {
		return nil, err
	}

	return gen.PostRelationUserIDUnfollow200Response{}, nil
}

// PostRelationUserIDHide implements v1.StrictServerInterface.
func (h handlers) PostRelationUserIDHide(ctx context.Context, request gen.PostRelationUserIDHideRequestObject) (gen.PostRelationUserIDHideResponseObject, error) {
	if err := h.relationService.Hide(ctx,
		auth.UserIDFromContext(ctx),
		request.UserID,
	); err != nil {
		return nil, err
	}

	return gen.PostRelationUserIDHide200Response{}, nil
}

// PostRelationUserIDUnhide implements v1.StrictServerInterface.
func (h handlers) PostRelationUserIDUnhide(ctx context.Context, request gen.PostRelationUserIDUnhideRequestObject) (gen.PostRelationUserIDUnhideResponseObject, error) {
	if err := h.relationService.Unhide(ctx,
		auth.UserIDFromContext(ctx),
		request.UserID,
	); err != nil {
		return nil, err
	}

	return gen.PostRelationUserIDUnhide200Response{}, nil
}
