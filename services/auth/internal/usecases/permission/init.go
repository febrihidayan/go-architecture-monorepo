package permission

import (
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/domain/repositories"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/config"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/repositories/factories"
)

type permissionInteractor struct {
	cfg            *config.AuthConfig
	permissionRepo repositories.PermissionRepository
}

func NewPermissionInteractor(
	config *config.AuthConfig,
	mongoFactory *factories.MongoFactory,
) *permissionInteractor {

	return &permissionInteractor{
		cfg:            config,
		permissionRepo: mongoFactory.PermissionRepo,
	}
}
