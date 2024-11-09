package entity

import "time"

type User struct {
	ID          string
	FollowStart time.Time
	Hidden      bool
}
