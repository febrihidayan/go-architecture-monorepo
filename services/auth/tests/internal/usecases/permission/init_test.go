package permission

import (
	"testing"
	"time"

	"bou.ke/monkey"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/domain/usecases"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/config"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/repositories/factories"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/usecases/permission"
	mongo_repositories "github.com/febrihidayan/go-architecture-monorepo/services/auth/tests/mocks/repositories/mongo"
	"github.com/stretchr/testify/suite"
)

type PermissionUsecaseSuite struct {
	suite.Suite
	cfg               *config.AuthConfig
	deps              *factories.Dependencies
	mongoFactory      *factories.MongoFactory
	permissionRepo    *mongo_repositories.PermissionRepositoryMock
	permissionUsecase usecases.PermissionUsecase
}

func (x *PermissionUsecaseSuite) SetupTest() {
	x.cfg = &config.AuthConfig{}

	x.permissionRepo = new(mongo_repositories.PermissionRepositoryMock)

	x.mongoFactory = &factories.MongoFactory{
		PermissionRepo: x.permissionRepo,
	}

	// Initialize dependencies
	x.deps = &factories.Dependencies{
		Config:       x.cfg,
		MongoFactory: x.mongoFactory,
	}

	x.permissionUsecase = permission.NewPermissionInteractor(x.deps)

	// fake time now for testing
	monkey.Patch(time.Now, func() time.Time {
		return time.Date(2020, 1, 1, 1, 1, 1, 1, time.UTC)
	})
}

func TestPermissionUsecase(t *testing.T) {
	suite.Run(t, new(PermissionUsecaseSuite))
}

type Any []interface{}

func (a Any) Get(i int) interface{} {
	return a[i]
}
