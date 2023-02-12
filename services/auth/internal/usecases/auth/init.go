package auth

import (
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/domain/repositories"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/config"
)

type authInteractor struct {
	cfg      *config.AuthConfig
	authRepo repositories.AuthRepository
	userRepo repositories.UserRepository
}

func NewAuthInteractor(
	config *config.AuthConfig,
	authRepo repositories.AuthRepository,
	userRepo repositories.UserRepository,
) *authInteractor {

	return &authInteractor{
		cfg:      config,
		authRepo: authRepo,
		userRepo: userRepo,
	}
}
