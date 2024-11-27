package acl

import (
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/domain/repositories"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/config"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/factories"
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

func NewAclInteractor(deps *factories.Dependencies) *aclInteractor {

	return &aclInteractor{
		cfg:                deps.Config,
		authRepo:           deps.MongoFactory.AuthRepo,
		roleRepo:           deps.MongoFactory.RoleRepo,
		roleUserRepo:       deps.MongoFactory.RoleUserRepo,
		permissionRepo:     deps.MongoFactory.PermissionRepo,
		permissionRoleRepo: deps.MongoFactory.PermissionRoleRepo,
		permissionUserRepo: deps.MongoFactory.PermissionUserRepo,
	}
}
