package repositories

import (
	"context"

	"github.com/febrihidayan/go-architecture-monorepo/services/storage/domain/entities"
)

type CloudRepository interface {
	Create(ctx context.Context, payload *entities.Cloud) error
	Find(ctx context.Context, id string) (*entities.Cloud, error)
	Delete(ctx context.Context, id string) error
}
