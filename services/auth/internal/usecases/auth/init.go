package auth

import (
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/domain/repositories"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/config"
)

type authInteractor struct {
	cfg      *config.AuthConfig
	authRepo repositories.AuthRepository
}

func NewAuthInteractor(
	config *config.AuthConfig,
	authRepo repositories.AuthRepository,
) *authInteractor {

	return &authInteractor{
		cfg:      config,
		authRepo: authRepo,
	}
}
