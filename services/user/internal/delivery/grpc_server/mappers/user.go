package mappers

import (
	"github.com/febrihidayan/go-architecture-monorepo/pkg/common"
	userPb "github.com/febrihidayan/go-architecture-monorepo/proto/_generated/user"
	"github.com/febrihidayan/go-architecture-monorepo/services/user/domain/entities"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToProtoUser(x *entities.User) *userPb.User {
	return &userPb.User{
		Id:        x.ID.String(),
		Name:      x.Name,
		Email:     x.Email,
		Avatar:    x.Avatar,
		LangCode:  x.LangCode,
		CreatedAt: timestamppb.New(x.CreatedAt),
		UpdatedAt: timestamppb.New(x.UpdatedAt),
	}
}

func ToDomainUserDto(x *userPb.User) entities.UserDto {
	id, _ := common.StringToID(x.GetId())
	return entities.UserDto{
		ID:       &id,
		Name:     x.GetName(),
		Email:    x.GetEmail(),
		Avatar:   x.GetAvatar(),
		LangCode: x.GetLangCode(),
	}
}
