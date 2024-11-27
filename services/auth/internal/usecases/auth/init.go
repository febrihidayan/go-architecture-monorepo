package auth

import (
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/domain/repositories"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/config"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/repositories/factories"
)

type authInteractor struct {
	cfg                  *config.AuthConfig
	authRepo             repositories.AuthRepository
	userRepo             repositories.UserRepository
	roleUserRepo         repositories.RoleUserRepository
	roleRepo             repositories.RoleRepository
	permissionUserRepo   repositories.PermissionUserRepository
	notificationGrpcRepo repositories.NotificationRepository
}

func NewAuthInteractor(deps *factories.Dependencies) *authInteractor {

	return &authInteractor{
		cfg:                  deps.Config,
		authRepo:             deps.MongoFactory.AuthRepo,
		roleUserRepo:         deps.MongoFactory.RoleUserRepo,
		roleRepo:             deps.MongoFactory.RoleRepo,
		permissionUserRepo:   deps.MongoFactory.PermissionUserRepo,
		userRepo:             deps.GrpcClientFactory.UserRepo,
		notificationGrpcRepo: deps.GrpcClientFactory.NotificationRepo,
	}
}
