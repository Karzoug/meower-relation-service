package converter

import (
	"github.com/Karzoug/meower-relation-service/internal/relation/entity"
	desc "github.com/Karzoug/meower-relation-service/pkg/proto/grpc/relation/v1"
)

func ToProtoUsers(users []entity.User) []*desc.User {
	pUsers := make([]*desc.User, len(users))
	for i := range users {
		pUsers[i] = &desc.User{
			Id:    users[i].ID.String(),
			Muted: users[i].Muted,
		}
	}
	return pUsers
}
