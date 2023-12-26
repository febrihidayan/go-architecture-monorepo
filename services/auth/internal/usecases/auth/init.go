package auth

import (
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/domain/repositories"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/config"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/repositories/factories"
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
	mongoFactory *factories.MongoFactory,
	grpcClientFactory *factories.GrpcClientFactory,
) *authInteractor {

	return &authInteractor{
		cfg:          config,
		authRepo:     mongoFactory.AuthRepo,
		roleUserRepo: mongoFactory.RoleUserRepo,
		roleRepo:     mongoFactory.RoleRepo,
		userRepo:     grpcClientFactory.UserRepo,
	}
}
