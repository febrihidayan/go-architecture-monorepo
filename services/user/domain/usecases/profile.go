package usecases

import (
	"context"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/exceptions"
	"github.com/febrihidayan/go-architecture-monorepo/services/user/domain/entities"
)

type ProfileUsecase interface {
	Find(ctx context.Context, id string) (*entities.User, *exceptions.CustomError)
	Update(ctx context.Context, payload entities.UserDto) (*entities.User, *exceptions.CustomError)
}
