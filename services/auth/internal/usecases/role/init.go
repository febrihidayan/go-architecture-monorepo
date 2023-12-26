package role

import (
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/domain/repositories"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/config"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/repositories/factories"
)

type roleInteractor struct {
	cfg      *config.AuthConfig
	roleRepo repositories.RoleRepository
}

func NewRoleInteractor(
	config *config.AuthConfig,
	mongoFactory *factories.MongoFactory,
) *roleInteractor {

	return &roleInteractor{
		cfg:      config,
		roleRepo: mongoFactory.RoleRepo,
	}
}
