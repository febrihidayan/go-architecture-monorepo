package usecases

import (
	"context"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/exceptions"
	"github.com/febrihidayan/go-architecture-monorepo/services/user/domain/entities"
)

type UserUsecase interface {
	Create(ctx context.Context, payload entities.UserDto) (*entities.User, *exceptions.CustomError)
	CreateAuth(ctx context.Context, payload entities.UserDto) (*entities.User, *exceptions.CustomError)
	GetAll(ctx context.Context, params entities.UserQueryParams) (*entities.UserMeta, *exceptions.CustomError)
	Update(ctx context.Context, payload entities.UserDto) (*entities.User, *exceptions.CustomError)
	UpdateEmailVerified(ctx context.Context, payload entities.UserDto) *exceptions.CustomError
}
