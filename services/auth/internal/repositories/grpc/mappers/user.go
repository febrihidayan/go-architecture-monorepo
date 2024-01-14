package mappers

import (
	userPb "github.com/febrihidayan/go-architecture-monorepo/proto/_generated/user"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/domain/entities"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToProtoUser(x entities.User) *userPb.User {
	return &userPb.User{
		Id:              x.ID.String(),
		Name:            x.Name,
		Email:           x.Email,
		EmailVerifiedAt: timestamppb.New(x.EmailVerifiedAt),
		CreatedAt:       timestamppb.New(x.CreatedAt),
		UpdatedAt:       timestamppb.New(x.UpdatedAt),
	}
}
