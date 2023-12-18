package role

import (
	"testing"
	"time"

	"bou.ke/monkey"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/domain/usecases"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/config"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/usecases/role"
	mongo_repositories "github.com/febrihidayan/go-architecture-monorepo/services/auth/tests/mocks/repositories/mongo"
	"github.com/stretchr/testify/suite"
)

type RoleUsecaseSuite struct {
	suite.Suite
	cfg         *config.AuthConfig
	roleRepo    *mongo_repositories.RoleRepositoryMock
	roleUsecase usecases.RoleUsecase
}

func (x *RoleUsecaseSuite) SetupTest() {
	x.cfg = &config.AuthConfig{}

	x.roleRepo = new(mongo_repositories.RoleRepositoryMock)
	x.roleUsecase = role.NewRoleInteractor(x.cfg, x.roleRepo)

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
