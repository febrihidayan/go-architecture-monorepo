package repositories

import (
	"context"

	"github.com/febrihidayan/go-architecture-monorepo/services/notification/domain/entities"
)

type NotificationRepository interface {
	Create(ctx context.Context, payload *entities.Notification) error
	GetAll(ctx context.Context, params *entities.NotificationQueryParams) ([]*entities.Notification, int, error)
	Delete(ctx context.Context, id string) error
}
