package grpc_server

import (
	"context"
	"errors"

	"github.com/febrihidayan/go-architecture-monorepo/proto/_generated/auth"
	authPb "github.com/febrihidayan/go-architecture-monorepo/proto/_generated/auth"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/domain/entities"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (x *server) CreateOrUpdateAuth(ctx context.Context, req *authPb.CreateOrUpdateAuthRequest) (*auth.CreateOrUpdateAuthResponse, error) {
	_, err := x.authUsecase.CreateOrUpdate(ctx, entities.AuthDto{
		UserId:   req.Data.GetUserId(),
		Email:    req.Data.GetEmail(),
		Password: req.Data.GetPassword(),
	})
	if err != nil {
		return nil, status.Error(codes.Canceled, errors.New(err.Errors.GoString()).Error())
	}

	return &authPb.CreateOrUpdateAuthResponse{
		Response: &authPb.RequestResponse{
			Status:  int32(codes.OK),
			Message: "success",
		},
	}, nil
}
