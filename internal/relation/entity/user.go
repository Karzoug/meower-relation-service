package entity

import (
	"time"

	"github.com/rs/xid"
)

type User struct {
	ID          xid.ID
	FollowStart time.Time
	Muted       bool
}
