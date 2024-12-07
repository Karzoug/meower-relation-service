package service

import (
	"context"
	"errors"

	"github.com/Karzoug/meower-common-go/ucerr"
	"github.com/rs/xid"
	"google.golang.org/grpc/codes"

	"github.com/Karzoug/meower-relation-service/internal/relation/entity"
	rerr "github.com/Karzoug/meower-relation-service/internal/relation/repo"
)

type RelationService struct {
	repo relationRepository
}

func NewRelationService(rr relationRepository) RelationService {
	return RelationService{
		repo: rr,
	}
}

func (rs RelationService) ListFollowings(ctx context.Context, userID xid.ID, pgn ListUsersPagination) (users []entity.User, token xid.ID, err error) {
	if pgn.Size < -1 {
		return nil, xid.NilID(), ucerr.NewError(
			nil,
			"invalid pagination parameter: negative size & not -1",
			codes.InvalidArgument,
		)
	}
	if pgn.Size == 0 {
		pgn.Size = 100
	} else if pgn.Size > 100 {
		pgn.Size = 100
	}

	users, token, err = rs.repo.ListFollowings(ctx, userID, pgn.Token, pgn.Size)
	if err != nil {
		return nil, xid.NilID(), ucerr.NewInternalError(err)
	}

	return
}

func (rs RelationService) ListFollowers(ctx context.Context, userID xid.ID, pgn ListUsersPagination) (users []entity.User, token xid.ID, err error) {
	if pgn.Size < -1 {
		return nil, xid.NilID(), ucerr.NewError(
			nil,
			"invalid pagination parameter: negative size & not -1",
			codes.InvalidArgument,
		)
	}
	if pgn.Size == 0 {
		pgn.Size = 100
	} else if pgn.Size > 100 {
		pgn.Size = 100
	}

	users, token, err = rs.repo.ListFollowers(ctx, userID, pgn.Token, pgn.Size)
	if err != nil {
		return nil, xid.NilID(), ucerr.NewInternalError(err)
	}

	return
}

func (rs RelationService) Follow(ctx context.Context, sourceUserID, targetUserID xid.ID) error {
	if err := rs.repo.Follow(ctx, sourceUserID, targetUserID); err != nil {
		if errors.Is(err, rerr.ErrNoAffected) {
			return ucerr.NewError(err,
				"follow operation failed: follow relation already exists or users not found",
				codes.FailedPrecondition)
		}
		return ucerr.NewInternalError(err)
	}

	return nil
}

func (rs RelationService) Unfollow(ctx context.Context, sourceUserID, targetUserID xid.ID) error {
	if err := rs.repo.Unfollow(ctx, sourceUserID, targetUserID); err != nil {
		if errors.Is(err, rerr.ErrNoAffected) {
			return ucerr.NewError(err,
				"unfollow operation failed: follow relation not exists or users not found",
				codes.FailedPrecondition)
		}
	}

	return nil
}

func (rs RelationService) Mute(ctx context.Context, sourceUserID, targetUserID xid.ID) error {
	if err := rs.repo.Mute(ctx, sourceUserID, targetUserID); err != nil {
		if errors.Is(err, rerr.ErrNoAffected) {
			return ucerr.NewError(err,
				"mute operation failed: mute relation already exists or users not found",
				codes.FailedPrecondition)
		}
	}

	return nil
}

func (rs RelationService) Unmute(ctx context.Context, sourceUserID, targetUserID xid.ID) error {
	if err := rs.repo.Unmute(ctx, sourceUserID, targetUserID); err != nil {
		if errors.Is(err, rerr.ErrNoAffected) {
			return ucerr.NewError(err,
				"unmute operation failed: mute relation not exists or users not found",
				codes.FailedPrecondition)
		}
	}

	return nil
}
