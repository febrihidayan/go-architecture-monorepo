package grpc_server

import (
	"context"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/exceptions"
	notificationPb "github.com/febrihidayan/go-architecture-monorepo/proto/_generated/notification"
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/internal/delivery/grpc_server/mappers"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (x server) CreateDeviceToken(ctx context.Context, req *notificationPb.CreateDeviceTokenRequest) (*emptypb.Empty, error) {
	_, err := x.deviceTokenUsecase.Create(ctx, mappers.ToDomainDeviceTokenDto(req.GetData()))
	if err != nil {
		return nil, status.Error(codes.Code(exceptions.MapToHttpStatusCode(err.Status)), err.Errors.Error())
	}

	return &emptypb.Empty{}, nil
}
