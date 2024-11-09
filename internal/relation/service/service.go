package service

import (
	"context"

	"github.com/Karzoug/meower-relation-service/internal/relation/entity"
	"github.com/Karzoug/meower-relation-service/pkg/ucerr"
)

type RelationService struct{}

func NewRelationService() RelationService {
	return RelationService{}
}

func (rs RelationService) ListFollowings(ctx context.Context,
	reqUserID, targetUserID string,
	options ...PaginationOption,
) (users []entity.User, token *string, err error) {
	return
}

func (rs RelationService) ListFollowers(ctx context.Context,
	reqUserID, targetUserID string,
	options ...PaginationOption,
) (users []entity.User, token *string, err error) {
	return
}
}

func (rs RelationService) Follow(ctx context.Context, reqUserID, targetUserID string) error {
	return nil
}
func (rs RelationService) Unfollow(ctx context.Context, reqUserID, targetUserID string) error {
	return nil
}
func (rs RelationService) Hide(ctx context.Context, reqUserID, targetUserID string) error {

	return nil
}
func (rs RelationService) Unhide(ctx context.Context, reqUserID, targetUserID string) error {

	return nil
}
