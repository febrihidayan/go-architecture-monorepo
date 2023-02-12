package grpc_server

import (
	"context"
	"errors"

	"github.com/febrihidayan/go-architecture-monorepo/proto/_generated/user"
	userPb "github.com/febrihidayan/go-architecture-monorepo/proto/_generated/user"
	"github.com/febrihidayan/go-architecture-monorepo/services/user/domain/entities"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (x *server) CreateUser(ctx context.Context, req *userPb.CreateUserRequest) (*user.CreateUserResponse, error) {
	user, err := x.userUsecase.Create(ctx, entities.UserDto{
		Name:  req.Data.GetName(),
		Email: req.Data.GetEmail(),
		Role:  req.Data.GetRole(),
	})
	if err != nil {
		return nil, status.Error(codes.Canceled, errors.New(err.Errors.GoString()).Error())
	}

	return &userPb.CreateUserResponse{
		Data: &userPb.User{
			Id:        user.ID.String(),
			Name:      user.Name,
			Role:      user.Role,
			CreatedAt: timestamppb.New(user.CreatedAt),
			UpdatedAt: timestamppb.New(user.UpdatedAt),
		},
	}, nil
}
