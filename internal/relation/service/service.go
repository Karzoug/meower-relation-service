package service

import (
	"context"
	"errors"

	"github.com/Karzoug/meower-relation-service/internal/relation/entity"
	rerr "github.com/Karzoug/meower-relation-service/internal/relation/repo"
	"github.com/Karzoug/meower-relation-service/pkg/ucerr"
	"github.com/Karzoug/meower-relation-service/pkg/ucerr/codes"
)

type RelationService struct {
	repo relationRepository
}

func NewRelationService(rr relationRepository) RelationService {
	return RelationService{
		repo: rr,
	}
}

func (rs RelationService) ListFollowings(ctx context.Context,
	reqUserID, targetUserID string,
	options ...PaginationOption,
) (users []entity.User, token *string, err error) {
	pagination := defaultPagination()
	for _, opt := range options {
		opt(&pagination)
	}
	users, token, err = rs.repo.ListFollowings(ctx, reqUserID, targetUserID, pagination.pageToken, pagination.maxPageSize)
	if err != nil {
		return nil, nil, ucerr.NewInternalError(err)
	}

	return
}

func (rs RelationService) ListFollowers(ctx context.Context,
	reqUserID, targetUserID string,
	options ...PaginationOption,
) (users []entity.User, token *string, err error) {
	pagination := defaultPagination()
	for _, opt := range options {
		opt(&pagination)
	}
	users, token, err = rs.repo.ListFollowers(ctx, reqUserID, targetUserID, pagination.pageToken, pagination.maxPageSize)
	if err != nil {
		return nil, nil, ucerr.NewInternalError(err)
	}

	return
}

func (rs RelationService) Follow(ctx context.Context, reqUserID, targetUserID string) error {
	if err := rs.repo.Follow(ctx, reqUserID, targetUserID); err != nil {
		if errors.Is(err, rerr.ErrNoAffected) {
			return ucerr.NewError(err,
				"follow operation failed: follow relation already exists or users not found",
				codes.FailedPrecondition)
		}
		return ucerr.NewInternalError(err)
	}

	return nil
}

func (rs RelationService) Unfollow(ctx context.Context, reqUserID, targetUserID string) error {
	if err := rs.repo.Unfollow(ctx, reqUserID, targetUserID); err != nil {
		if errors.Is(err, rerr.ErrNoAffected) {
			return ucerr.NewError(err,
				"unfollow operation failed: follow relation not exists or users not found",
				codes.FailedPrecondition)
		}
	}

	return nil
}

func (rs RelationService) Hide(ctx context.Context, reqUserID, targetUserID string) error {
	if err := rs.repo.Hide(ctx, reqUserID, targetUserID); err != nil {
		if errors.Is(err, rerr.ErrNoAffected) {
			return ucerr.NewError(err,
				"hide operation failed: hide relation already exists or users not found",
				codes.FailedPrecondition)
		}
	}

	return nil
}

func (rs RelationService) Unhide(ctx context.Context, reqUserID, targetUserID string) error {
	if err := rs.repo.Unhide(ctx, reqUserID, targetUserID); err != nil {
		if errors.Is(err, rerr.ErrNoAffected) {
			return ucerr.NewError(err,
				"unhide operation failed: hide relation not exists or users not found",
				codes.FailedPrecondition)
		}
	}

	return nil
}
