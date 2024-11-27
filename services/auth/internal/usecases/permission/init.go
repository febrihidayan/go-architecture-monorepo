package permission

import (
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/domain/repositories"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/config"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/factories"
)

type permissionInteractor struct {
	cfg            *config.AuthConfig
	permissionRepo repositories.PermissionRepository
}

func NewPermissionInteractor(deps *factories.Dependencies) *permissionInteractor {
	return &permissionInteractor{
		cfg:            deps.Config,
		permissionRepo: deps.MongoFactory.PermissionRepo,
	}
}
