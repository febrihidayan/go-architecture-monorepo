package grpc_server

import (
	"context"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/exceptions"
	"github.com/febrihidayan/go-architecture-monorepo/proto/_generated/user"
	userPb "github.com/febrihidayan/go-architecture-monorepo/proto/_generated/user"
	"github.com/febrihidayan/go-architecture-monorepo/services/user/internal/delivery/grpc_server/mappers"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (x *server) CreateUser(ctx context.Context, req *userPb.CreateUserRequest) (*user.CreateUserResponse, error) {
	user, err := x.userUsecase.CreateAuth(ctx, mappers.ToDomainUserDto(req.GetData()))
	if err != nil {
		return nil, status.Error(codes.Code(exceptions.MapToHttpStatusCode(err.Status)), err.Errors.Error())
	}

	return &userPb.CreateUserResponse{
		Data: mappers.ToProtoUser(user),
	}, nil
}

func (x *server) FindUser(ctx context.Context, req *userPb.FindUserRequest) (*user.FindUserResponse, error) {
	user, err := x.profileUsecase.Find(ctx, req.GetId())
	if err != nil {
		return nil, status.Error(codes.Code(exceptions.MapToHttpStatusCode(err.Status)), err.Errors.Error())
	}

	return &userPb.FindUserResponse{
		Data: mappers.ToProtoUser(user),
	}, nil
}

func (x *server) UpdateEmailVerifiedUser(ctx context.Context, req *userPb.User) (*emptypb.Empty, error) {
	if err := x.userUsecase.UpdateEmailVerified(ctx, mappers.ToDomainUserDto(req)); err != nil {
		return nil, status.Error(codes.Code(exceptions.MapToHttpStatusCode(err.Status)), err.Errors.Error())
	}

	return &emptypb.Empty{}, nil
}
