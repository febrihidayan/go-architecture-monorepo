package role

import (
	"testing"
	"time"

	"bou.ke/monkey"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/domain/usecases"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/config"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/factories"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/usecases/role"
	mongo_repositories "github.com/febrihidayan/go-architecture-monorepo/services/auth/tests/mocks/repositories/mongo"
	"github.com/stretchr/testify/suite"
)

type RoleUsecaseSuite struct {
	suite.Suite
	cfg          *config.AuthConfig
	deps         *factories.Dependencies
	mongoFactory *factories.MongoFactory
	roleRepo     *mongo_repositories.RoleRepositoryMock
	roleUsecase  usecases.RoleUsecase
}

func (x *RoleUsecaseSuite) SetupTest() {
	x.cfg = &config.AuthConfig{}

	x.roleRepo = new(mongo_repositories.RoleRepositoryMock)

	x.mongoFactory = &factories.MongoFactory{
		RoleRepo: x.roleRepo,
	}

	// Initialize dependencies
	x.deps = &factories.Dependencies{
		Config:       x.cfg,
		MongoFactory: x.mongoFactory,
	}

	x.roleUsecase = role.NewRoleInteractor(x.deps)

	// fake time now for testing
	monkey.Patch(time.Now, func() time.Time {
		return time.Date(2020, 1, 1, 1, 1, 1, 1, time.UTC)
	})
}

func TestRoleUsecase(t *testing.T) {
	suite.Run(t, new(RoleUsecaseSuite))
}

type Any []interface{}

func (a Any) Get(i int) interface{} {
	return a[i]
}
