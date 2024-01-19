package usecases

import (
	"context"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/exceptions"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/domain/entities"
)

type AclUsecase interface {
	GetAll(ctx context.Context) (*entities.AclMeta, *exceptions.CustomError)
	GetAllUser(ctx context.Context, userId string) (*entities.AclMeta, *exceptions.CustomError)
	UpdateUser(ctx context.Context, payload entities.AclUserDto) *exceptions.CustomError
}
