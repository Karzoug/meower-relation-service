package relation

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/Karzoug/meower-common-go/auth"
	"github.com/rs/xid"

	"github.com/Karzoug/meower-relation-service/internal/delivery/grpc/converter"
	gen "github.com/Karzoug/meower-relation-service/internal/delivery/grpc/gen/relation/v1"
	"github.com/Karzoug/meower-relation-service/internal/relation/service"
)

func RegisterService(rs service.RelationService) func(grpcServer *grpc.Server) {
	hdl := handlers{
		relationService: rs,
	}
	return func(grpcServer *grpc.Server) {
		gen.RegisterRelationServiceServer(grpcServer, hdl)
	}
}

type handlers struct {
	gen.UnimplementedRelationServiceServer
	relationService service.RelationService
}

func (h handlers) ListFollowings(ctx context.Context, req *gen.ListFollowingsRequest) (*gen.ListFollowingsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	id, err := xid.FromString(req.Parent)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid id: "+req.Parent)
	}

	token := xid.NilID()
	if len(req.PageToken) != 0 {
		tkn, err := xid.FromString(req.PageToken)
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, "invalid page token: "+req.PageToken)
		}
		token = tkn
	}

	users, nextID, err := h.relationService.ListFollowings(ctx, id, service.ListUsersPagination{
		Token: token,
		Size:  int(req.PageSize),
	})
	if err != nil {
		return nil, err
	}

	var nextToken string
	if nextID.IsNil() {
		nextToken = ""
	} else {
		nextToken = nextID.String()
	}

	return &gen.ListFollowingsResponse{
		Followings:    converter.ToProtoUsers(users),
		NextPageToken: nextToken,
	}, nil
}

func (h handlers) ListFollowers(ctx context.Context, req *gen.ListFollowersRequest) (*gen.ListFollowersResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	id, err := xid.FromString(req.Parent)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid id: "+req.Parent)
	}

	token := xid.NilID()
	if len(req.PageToken) != 0 {
		tkn, err := xid.FromString(req.PageToken)
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, "invalid page token: "+req.PageToken)
		}
		token = tkn
	}

	users, nextID, err := h.relationService.ListFollowers(ctx, id, service.ListUsersPagination{
		Token: token,
		Size:  int(req.PageSize),
	})
	if err != nil {
		return nil, err
	}

	var nextToken string
	if nextID.IsNil() {
		nextToken = ""
	} else {
		nextToken = nextID.String()
	}

	return &gen.ListFollowersResponse{
		Followers:     converter.ToProtoUsers(users),
		NextPageToken: nextToken,
	}, nil
}

func (h handlers) CreateRelation(ctx context.Context, req *gen.CreateRelationRequest) (*emptypb.Empty, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	userID, err := xid.FromString(req.UserId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid id: "+req.UserId)
	}

	switch req.RelationType {
	case gen.RelationType_RELATION_TYPE_FOLLOW:
		err = h.relationService.Follow(ctx, auth.UserIDFromContext(ctx), userID)
	case gen.RelationType_RELATION_TYPE_MUTE:
		err = h.relationService.Mute(ctx, auth.UserIDFromContext(ctx), userID)
	default:
		return nil, status.Error(codes.InvalidArgument, "unknown relation type: "+req.RelationType.String())
	}

	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (h handlers) DeleteRelation(ctx context.Context, req *gen.DeleteRelationRequest) (*emptypb.Empty, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "empty request")
	}

	userID, err := xid.FromString(req.UserId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid id: "+req.UserId)
	}

	switch req.RelationType {
	case gen.RelationType_RELATION_TYPE_FOLLOW:
		err = h.relationService.Unfollow(ctx, auth.UserIDFromContext(ctx), userID)
	case gen.RelationType_RELATION_TYPE_MUTE:
		err = h.relationService.Unmute(ctx, auth.UserIDFromContext(ctx), userID)
	default:
		return nil, status.Error(codes.InvalidArgument, "unknown relation type: "+req.RelationType.String())
	}

	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
