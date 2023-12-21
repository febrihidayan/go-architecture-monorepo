package repositories

import (
	"context"

	"github.com/febrihidayan/go-architecture-monorepo/services/user/domain/entities"
)

type UserRepository interface {
	Create(ctx context.Context, payload *entities.User) error
	Find(ctx context.Context, id string) (*entities.User, error)
	FindByEmail(ctx context.Context, email string) (*entities.User, error)
	GetAll(ctx context.Context, params *entities.UserQueryParams) ([]*entities.User, int, error)
	Update(ctx context.Context, payload *entities.User) error
}
