package auth

import (
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/domain/repositories"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/config"
)

type authInteractor struct {
	cfg          *config.AuthConfig
	authRepo     repositories.AuthRepository
	userRepo     repositories.UserRepository
	roleUserRepo repositories.RoleUserRepository
	roleRepo     repositories.RoleRepository
}

func NewAuthInteractor(
	config *config.AuthConfig,
	authRepo repositories.AuthRepository,
	userRepo repositories.UserRepository,
	roleUserRepo repositories.RoleUserRepository,
	roleRepo repositories.RoleRepository,
) *authInteractor {

	return &authInteractor{
		cfg:          config,
		authRepo:     authRepo,
		userRepo:     userRepo,
		roleUserRepo: roleUserRepo,
		roleRepo:     roleRepo,
	}
}
