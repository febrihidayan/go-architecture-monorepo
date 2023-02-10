package usecases

import (
	"context"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/exceptions"
	"github.com/febrihidayan/go-architecture-monorepo/services/user/domain/entities"
)

type UserUsecase interface {
	Create(ctx context.Context, payload entities.UserDto) (*entities.User, *exceptions.CustomError)
	Profile(ctx context.Context, id string) (*entities.User, *exceptions.CustomError)
}
