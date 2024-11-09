package relation

import (
	"context"

	"google.golang.org/grpc"

	"github.com/Karzoug/meower-relation-service/internal/delivery/grpc/converter"
	desc "github.com/Karzoug/meower-relation-service/internal/delivery/grpc/gen/relation/v1"
	"github.com/Karzoug/meower-relation-service/internal/relation/service"
	"github.com/Karzoug/meower-relation-service/pkg/auth"
)

func RegisterService(rs service.RelationService) func(grpcServer *grpc.Server) {
	hdl := handlers{
		relationService: rs,
	}
	return func(grpcServer *grpc.Server) {
		desc.RegisterRelationServiceServer(grpcServer, hdl)
	}
}

type handlers struct {
	desc.UnimplementedRelationServiceServer
	relationService service.RelationService
}

func (h handlers) ListFollowings(ctx context.Context, req *desc.ListFollowingsRequest) (*desc.ListFollowingsResponse, error) {
	users, token, err := h.relationService.ListFollowings(ctx,
		auth.UserIDFromContext(ctx),
		req.UserId,
		service.WithPaginationToken(req.Pagination.PageToken),
		service.WithPaginationMaxPageSize(int(req.Pagination.MaxPageSize)),
	)
	if err != nil {
		return nil, err
	}

	var respToken *desc.PaginationResponse
	if token != nil {
		respToken = &desc.PaginationResponse{NextPageToken: *token}
	}
	return &desc.ListFollowingsResponse{
		Followings: converter.ToProtoUsers(users),
		Pagination: respToken,
	}, nil
}

func (h handlers) ListFollowers(ctx context.Context, req *desc.ListFollowersRequest) (*desc.ListFollowersResponse, error) {
	users, token, err := h.relationService.ListFollowers(ctx,
		auth.UserIDFromContext(ctx),
		req.UserId,
		service.WithPaginationToken(req.Pagination.PageToken),
		service.WithPaginationMaxPageSize(int(req.Pagination.MaxPageSize)),
	)
	if err != nil {
		return nil, err
	}

	var respToken *desc.PaginationResponse
	if token != nil {
		respToken = &desc.PaginationResponse{NextPageToken: *token}
	}
	return &desc.ListFollowersResponse{
		Followers:  converter.ToProtoUsers(users),
		Pagination: respToken,
	}, nil
}
