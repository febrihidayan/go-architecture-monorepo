package role

import (
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/domain/repositories"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/config"
)

type roleInteractor struct {
	cfg      *config.AuthConfig
	roleRepo repositories.RoleRepository
}

func NewRoleInteractor(
	config *config.AuthConfig,
	roleRepo repositories.RoleRepository,
) *roleInteractor {

	return &roleInteractor{
		cfg:      config,
		roleRepo: roleRepo,
	}
}
