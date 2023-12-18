package permission

import (
	"testing"

	"github.com/febrihidayan/go-architecture-monorepo/services/auth/domain/usecases"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/config"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/usecases/permission"
	mongo_repositories "github.com/febrihidayan/go-architecture-monorepo/services/auth/tests/mocks/repositories/mongo"
	"github.com/stretchr/testify/suite"
)

type PermissionUsecaseSuite struct {
	suite.Suite
	cfg               *config.AuthConfig
	permissionRepo    *mongo_repositories.PermissionRepositoryMock
	permissionUsecase usecases.PermissionUsecase
}

func (x *PermissionUsecaseSuite) SetupTest() {
	x.cfg = &config.AuthConfig{}

	x.permissionRepo = new(mongo_repositories.PermissionRepositoryMock)
	x.permissionUsecase = permission.NewPermissionInteractor(x.cfg, x.permissionRepo)
}

func TestPermissionUsecase(t *testing.T) {
	suite.Run(t, new(PermissionUsecaseSuite))
}

type Any []interface{}

func (a Any) Get(i int) interface{} {
	return a[i]
}
