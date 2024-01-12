package grpc_repositories

import (
	"context"

	notificationPb "github.com/febrihidayan/go-architecture-monorepo/proto/_generated/notification"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/domain/entities"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/repositories/grpc/mappers"

	"google.golang.org/grpc"
)

type NotificationRepository struct {
	svc notificationPb.NotificationServicesClient
	ctx context.Context
}

func NewNotificationRepository(con *grpc.ClientConn) NotificationRepository {
	client := notificationPb.NewNotificationServicesClient(con)

	return NotificationRepository{
		svc: client,
		ctx: nil,
	}
}

func (u *NotificationRepository) SendNotification(ctx context.Context, payload entities.NotificationSends) error {
	_, err := u.svc.SendNotification(ctx, &notificationPb.SendNotificationRequest{
		Data: mappers.ToProtoNotificationSendParams(payload),
	})
	if err != nil {
		return err
	}

	return nil
}
