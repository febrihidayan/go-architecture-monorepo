package mappers

import (
	"github.com/febrihidayan/go-architecture-monorepo/pkg/common"
	userPb "github.com/febrihidayan/go-architecture-monorepo/proto/_generated/user"
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/domain/entities"
)

func ToDomainUser(x *userPb.User) *entities.User {
	id, _ := common.StringToID(x.GetId())
	return &entities.User{
		ID:        id,
		Name:      x.GetName(),
		Email:     x.GetEmail(),
		Avatar:    x.GetAvatar(),
		LangCode:  x.GetLangCode(),
		CreatedAt: x.CreatedAt.AsTime(),
		UpdatedAt: x.UpdatedAt.AsTime(),
	}
}
