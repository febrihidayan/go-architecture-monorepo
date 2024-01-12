package repositories

import (
	"context"

	"github.com/febrihidayan/go-architecture-monorepo/services/auth/domain/entities"
)

type NotificationRepository interface {
	SendNotification(ctx context.Context, payload entities.NotificationSends) error
}
