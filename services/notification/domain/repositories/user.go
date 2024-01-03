package repositories

import (
	"context"

	"github.com/febrihidayan/go-architecture-monorepo/services/notification/domain/entities"
)

type UserRepository interface {
	FindUser(ctx context.Context, id string) (*entities.User, error)
}
