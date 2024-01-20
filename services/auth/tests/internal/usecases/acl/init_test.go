package acl

import (
	"testing"
	"time"

	"bou.ke/monkey"
	"github.com/febrihidayan/go-architecture-monorepo/pkg/common"
	"github.com/febrihidayan/go-architecture-monorepo/pkg/middleware"
	"github.com/febrihidayan/go-architecture-monorepo/pkg/utils"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/domain/entities"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/domain/usecases"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/config"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/repositories/factories"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/usecases/acl"
	mongo_repositories "github.com/febrihidayan/go-architecture-monorepo/services/auth/tests/mocks/repositories/mongo"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
)

type AclUsecaseSuite struct {
	suite.Suite
	cfg                *config.AuthConfig
	mongoFactory       *factories.MongoFactory
	authRepo           *mongo_repositories.AuthRepositoryMock
	permissionRepo     *mongo_repositories.PermissionRepositoryMock
	permissionRoleRepo *mongo_repositories.PermissionRoleRepositoryMock
	permissionUserRepo *mongo_repositories.PermissionUserRepositoryMock
	roleRepo           *mongo_repositories.RoleRepositoryMock
	roleUserRepo       *mongo_repositories.RoleUserRepositoryMock
	aclUsecase         usecases.AclUsecase

	role           *entities.Role
	roleUser       *entities.RoleUser
	permissionRole *entities.PermissionRole
	permissionUser *entities.PermissionUser
	permission     *entities.Permission
	auth           *entities.Auth
	id             uuid.UUID
}

func (x *AclUsecaseSuite) SetupTest() {
	x.cfg = &config.AuthConfig{}

	x.authRepo = new(mongo_repositories.AuthRepositoryMock)
	x.permissionRepo = new(mongo_repositories.PermissionRepositoryMock)
	x.permissionRoleRepo = new(mongo_repositories.PermissionRoleRepositoryMock)
	x.permissionUserRepo = new(mongo_repositories.PermissionUserRepositoryMock)
	x.roleUserRepo = new(mongo_repositories.RoleUserRepositoryMock)
	x.roleRepo = new(mongo_repositories.RoleRepositoryMock)

	x.mongoFactory = &factories.MongoFactory{
		AuthRepo:           x.authRepo,
		PermissionRepo:     x.permissionRepo,
		PermissionRoleRepo: x.permissionRoleRepo,
		PermissionUserRepo: x.permissionUserRepo,
		RoleUserRepo:       x.roleUserRepo,
		RoleRepo:           x.roleRepo,
	}

	x.aclUsecase = acl.NewAclInteractor(x.cfg, x.mongoFactory)

	// fake time now for testing
	monkey.Patch(time.Now, func() time.Time {
		return time.Date(2020, 1, 1, 1, 1, 1, 1, time.UTC)
	})

	// storage data
	x.id, _ = common.StringToID("83f8619c-608c-4e5d-8941-3c394551085d")

	x.role = &entities.Role{
		ID:          x.id,
		Name:        middleware.ROLE_SUPERADMINISTRATOR,
		DisplayName: middleware.ROLE_SUPERADMINISTRATOR,
		Description: middleware.ROLE_SUPERADMINISTRATOR,
		CreatedAt:   utils.TimeUTC(),
		UpdatedAt:   utils.TimeUTC(),
	}

	x.roleUser = &entities.RoleUser{
		ID:     x.id,
		RoleId: x.id.String(),
		UserId: x.id.String(),
	}

	x.permission = &entities.Permission{
		ID:          x.id,
		Name:        "users_create",
		DisplayName: "users create",
		Description: "users create",
		CreatedAt:   utils.TimeUTC(),
		UpdatedAt:   utils.TimeUTC(),
	}

	x.permissionRole = &entities.PermissionRole{
		ID:           x.id,
		RoleId:       x.id.String(),
		PermissionId: x.id.String(),
	}

	x.permissionUser = &entities.PermissionUser{
		ID:           x.id,
		UserId:       x.id.String(),
		PermissionId: x.id.String(),
	}

	x.auth = &entities.Auth{
		ID:     x.id,
		UserId: x.id.String(),
	}
}

func TestAclUsecase(t *testing.T) {
	suite.Run(t, new(AclUsecaseSuite))
}

type Any []interface{}

func (a Any) Get(i int) interface{} {
	return a[i]
}
