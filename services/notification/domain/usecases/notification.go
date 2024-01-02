package usecases

import (
	"context"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/exceptions"
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/domain/entities"
)

type NotificationUsecase interface {
	GetAll(ctx context.Context, params entities.NotificationQueryParams) (*entities.NotificationMeta, *exceptions.CustomError)
}
