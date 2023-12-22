package grpc_repositories

import (
	"context"

	authPb "github.com/febrihidayan/go-architecture-monorepo/proto/_generated/auth"
	"github.com/febrihidayan/go-architecture-monorepo/services/user/domain/entities"

	"google.golang.org/grpc"
)

type AuthRepository struct {
	svc authPb.AuthServicesClient
	ctx context.Context
}

func NewAuthRepository(con *grpc.ClientConn) AuthRepository {
	client := authPb.NewAuthServicesClient(con)

	return AuthRepository{
		svc: client,
		ctx: nil,
	}
}

func (x *AuthRepository) CreateOrUpdate(ctx context.Context, payload *entities.Auth) error {
	_, err := x.svc.CreateOrUpdateAuth(ctx, &authPb.CreateOrUpdateAuthRequest{
		Data: &authPb.Auth{
			Email:    payload.Email,
			UserId:   payload.UserId,
			Password: payload.Password,
		},
	})

	return err
}
