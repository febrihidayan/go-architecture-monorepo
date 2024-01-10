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

func (x server) SendNotification(ctx context.Context, req *notificationPb.SendNotificationRequest) (*emptypb.Empty, error) {
	err := x.notificationUsecase.SendPushJobs(ctx, mappers.ToDomainNotificationParamsDto(req.GetData()))
	if err != nil {
		return nil, status.Error(codes.Code(exceptions.MapToHttpStatusCode(err.Status)), err.Errors.Error())
	}

	return &emptypb.Empty{}, nil
}
