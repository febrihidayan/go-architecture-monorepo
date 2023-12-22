package usecases

import (
	"context"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/exceptions"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/domain/entities"
)

type AuthUsecase interface {
	Login(ctx context.Context, payload entities.AuthDto) (*entities.AuthTokenMeta, *exceptions.CustomError)
	Register(ctx context.Context, payload entities.RegisterDto) (*entities.Auth, *exceptions.CustomError)
	CreateOrUpdate(ctx context.Context, payload entities.AuthDto) (*entities.Auth, *exceptions.CustomError)
}
