package user

import (
	"github.com/febrihidayan/go-architecture-monorepo/services/user/domain/repositories"
	"github.com/febrihidayan/go-architecture-monorepo/services/user/internal/config"
)

type userInteractor struct {
	cfg      *config.UserConfig
	userRepo repositories.UserRepository
	authRepo repositories.AuthRepository
}

func NewUserInteractor(
	config *config.UserConfig,
	userRepo repositories.UserRepository,
	authRepo repositories.AuthRepository,
) *userInteractor {

	return &userInteractor{
		cfg:      config,
		userRepo: userRepo,
		authRepo: authRepo,
	}
}
