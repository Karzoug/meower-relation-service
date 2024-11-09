package service

import (
	"context"

	"github.com/Karzoug/meower-relation-service/internal/relation/entity"
)

type relationRepository interface {
	Follow(ctx context.Context, reqUserID, targetUserID string) error
	Unfollow(ctx context.Context, reqUserID, targetUserID string) error
	Hide(ctx context.Context, reqUserID, targetUserID string) error
	Unhide(ctx context.Context, reqUserID, targetUserID string) error
	ListFollowings(ctx context.Context, reqUserID, targetUserID, afterID string, limit int) ([]entity.User, *string, error)
	ListFollowers(ctx context.Context, reqUserID, targetUserID, afterID string, limit int) ([]entity.User, *string, error)
}
