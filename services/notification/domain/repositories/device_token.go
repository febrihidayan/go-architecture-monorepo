package repositories

import (
	"context"

	"github.com/febrihidayan/go-architecture-monorepo/services/notification/domain/entities"
)

type DeviceTokenRepository interface {
	Create(ctx context.Context, payload *entities.DeviceToken) error
	All(ctx context.Context, params *entities.DeviceTokenQueryParams) ([]*entities.DeviceToken, error)
	Delete(ctx context.Context, id string) error
}
