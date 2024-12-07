package service

import (
	"context"

	"github.com/rs/xid"

	"github.com/Karzoug/meower-relation-service/internal/relation/entity"
)

type relationRepository interface {
	Follow(ctx context.Context, sourceUserID, targetUserID xid.ID) error
	Unfollow(ctx context.Context, sourceUserID, targetUserID xid.ID) error
	Mute(ctx context.Context, sourceUserID, targetUserID xid.ID) error
	Unmute(ctx context.Context, sourceUserID, targetUserID xid.ID) error
	ListFollowings(ctx context.Context, userID, afterID xid.ID, limit int) ([]entity.User, xid.ID, error)
	ListFollowers(ctx context.Context, userID, afterID xid.ID, limit int) ([]entity.User, xid.ID, error)
}
