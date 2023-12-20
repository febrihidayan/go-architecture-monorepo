package grpc_repositories

import (
	"context"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/common"
	userPb "github.com/febrihidayan/go-architecture-monorepo/proto/_generated/user"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/domain/entities"

	"google.golang.org/grpc"
)

type UserRepository struct {
	svc userPb.UserServicesClient
	ctx context.Context
}

func NewUserRepository(con *grpc.ClientConn) UserRepository {
	client := userPb.NewUserServicesClient(con)

	return UserRepository{
		svc: client,
		ctx: nil,
	}
}

func (u *UserRepository) CreateUser(ctx context.Context, payload entities.User) (*entities.User, error) {
	user, err := u.svc.CreateUser(ctx, &userPb.CreateUserRequest{
		Data: &userPb.User{
			Name:  payload.Name,
			Email: payload.Email,
		},
	})
	if err != nil {
		return nil, err
	}
	userId, errUserId := common.StringToID(user.Data.GetId())
	if errUserId != nil {
		return nil, errUserId
	}
	return &entities.User{
		ID:        userId,
		Name:      user.Data.GetName(),
		Email:     user.Data.GetEmail(),
		CreatedAt: user.Data.GetCreatedAt().AsTime(),
		UpdatedAt: user.Data.GetUpdatedAt().AsTime(),
	}, nil
}
