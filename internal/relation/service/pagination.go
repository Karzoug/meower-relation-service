package service

import "github.com/rs/xid"

type ListUsersPagination struct {
	Token xid.ID
	Size  int
}
