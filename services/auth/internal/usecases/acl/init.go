package acl

import (
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/domain/repositories"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/config"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/repositories/factories"
)

type aclInteractor struct {
	cfg                *config.AuthConfig
	authRepo           repositories.AuthRepository
	roleRepo           repositories.RoleRepository
	roleUserRepo       repositories.RoleUserRepository
	permissionRepo     repositories.PermissionRepository
	permissionRoleRepo repositories.PermissionRoleRepository
	permissionUserRepo repositories.PermissionUserRepository
}

func NewAclInteractor(
	config *config.AuthConfig,
	mongoFactory *factories.MongoFactory,
) *aclInteractor {

	return &aclInteractor{
		cfg:                config,
		authRepo:           mongoFactory.AuthRepo,
		roleRepo:           mongoFactory.RoleRepo,
		roleUserRepo:       mongoFactory.RoleUserRepo,
		permissionRepo:     mongoFactory.PermissionRepo,
		permissionRoleRepo: mongoFactory.PermissionRoleRepo,
		permissionUserRepo: mongoFactory.PermissionUserRepo,
	}
}
