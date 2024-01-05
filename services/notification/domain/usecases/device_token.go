package usecases

import (
	"context"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/exceptions"
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/domain/entities"
)

type DeviceTokenUsecase interface {
	Create(ctx context.Context, payload entities.DeviceTokenDto) (*entities.DeviceToken, *exceptions.CustomError)
}
