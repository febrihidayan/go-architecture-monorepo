package usecases

import (
	"context"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/exceptions"
	"github.com/febrihidayan/go-architecture-monorepo/services/storage/domain/entities"
)

type CloudUsecase interface {
	Create(ctx context.Context, payload entities.CloudDto) (*entities.Cloud, *exceptions.CustomError)
	UpdateStatus(ctx context.Context, payloads []*entities.Cloud) *exceptions.CustomError
	DeleteAllStatusJob(ctx context.Context) *exceptions.CustomError
}
