package repositories

import (
	"context"

	"github.com/febrihidayan/go-architecture-monorepo/services/auth/domain/entities"
)

type AuthRepository interface {
	Create(ctx context.Context, payload *entities.Auth) error
	Find(ctx context.Context, id string) (*entities.Auth, error)
	FindByEmail(ctx context.Context, email string) (*entities.Auth, error)
	FindByUserId(ctx context.Context, userId string) (*entities.Auth, error)
	Update(ctx context.Context, payload *entities.Auth) error
	DeleteByUserID(ctx context.Context, userId string) error
}
