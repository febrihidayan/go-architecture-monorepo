package permission

import (
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/domain/repositories"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/config"
)

type permissionInteractor struct {
	cfg            *config.AuthConfig
	permissionRepo repositories.PermissionRepository
}

func NewPermissionInteractor(
	config *config.AuthConfig,
	permissionRepo repositories.PermissionRepository,
) *permissionInteractor {

	return &permissionInteractor{
		cfg:            config,
		permissionRepo: permissionRepo,
	}
}
