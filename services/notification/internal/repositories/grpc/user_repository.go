package grpc_repositories

import (
	"context"

	userPb "github.com/febrihidayan/go-architecture-monorepo/proto/_generated/user"
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/domain/entities"
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/internal/repositories/grpc/mappers"

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

func (u *UserRepository) FindUser(ctx context.Context, id string) (*entities.User, error) {
	user, err := u.svc.FindUser(ctx, &userPb.FindUserRequest{
		Id: id,
	})
	if err != nil {
		return nil, err
	}

	return mappers.ToDomainUser(user.GetData()), nil
}
