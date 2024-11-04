package service

import (
	"context"

	"github.com/Karzoug/meower-relation-service/internal/relation/entity"
	"github.com/Karzoug/meower-relation-service/internal/relation/service/options"
)

type RelationService struct{}

func NewRelationService() RelationService {
	return RelationService{}
}

func (s RelationService) ListFollowings(ctx context.Context,
	reqUserID, targetUserID string,
	opts ...options.Lister[options.ListFollowingsOptions],
) ([]entity.User, error) {
	panic("not implemented")
}

func (s RelationService) ListFollowers(ctx context.Context,
	reqUserID, targetUserID string,
	opts ...options.Lister[options.ListFollowersOptions],
) ([]entity.User, error) {
	panic("not implemented")
}

func (s RelationService) Follow(ctx context.Context, reqUserID, targetUserID string) error {
	panic("not implemented")
}

func (s RelationService) Unfollow(ctx context.Context, reqUserID, targetUserID string) error {
	panic("not implemented")
}

func (s RelationService) Hide(ctx context.Context, reqUserID, targetUserID string) error {
	panic("not implemented")
}

func (s RelationService) Unhide(ctx context.Context, reqUserID, targetUserID string) error {
	panic("not implemented")
}
