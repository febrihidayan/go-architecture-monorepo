package factories

import (
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/domain/repositories"
	mongo_repositories "github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/repositories/mongo"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoFactory struct {
	AuthRepo           repositories.AuthRepository
	PermissionRoleRepo repositories.PermissionRoleRepository
	PermissionUserRepo repositories.PermissionUserRepository
	PermissionRepo     repositories.PermissionRepository
	RoleRepo           repositories.RoleRepository
	RoleUserRepo       repositories.RoleUserRepository
}

func NewMongoFactory(db *mongo.Database) *MongoFactory {
	var (
		AuthRepo           = mongo_repositories.NewAuthRepository(db)
		PermissionRoleRepo = mongo_repositories.NewPermissionRoleRepository(db)
		PermissionUserRepo = mongo_repositories.NewPermissionUserRepository(db)
		PermissionRepo     = mongo_repositories.NewPermissionRepository(db)
		RoleRepo           = mongo_repositories.NewRoleRepository(db)
		RoleUserRepo       = mongo_repositories.NewRoleUserRepository(db)
	)

	return &MongoFactory{
		AuthRepo:           &AuthRepo,
		PermissionRoleRepo: &PermissionRoleRepo,
		PermissionUserRepo: &PermissionUserRepo,
		PermissionRepo:     &PermissionRepo,
		RoleRepo:           &RoleRepo,
		RoleUserRepo:       &RoleUserRepo,
	}
}
