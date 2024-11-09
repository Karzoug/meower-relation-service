package converter

import (
	desc "github.com/Karzoug/meower-relation-service/internal/delivery/grpc/gen/relation/v1"
	"github.com/Karzoug/meower-relation-service/internal/relation/entity"
)

func ToProtoUsers(users []entity.User) []*desc.User {
	pUsers := make([]*desc.User, len(users))
	for i := range users {
		pUsers[i] = &desc.User{
			Id:     users[i].ID,
			Hidden: users[i].Hidden,
		}
	}
	return pUsers
}
