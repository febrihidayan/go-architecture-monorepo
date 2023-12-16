package user

import (
	"testing"

	"github.com/febrihidayan/go-architecture-monorepo/services/user/domain/usecases"
	"github.com/febrihidayan/go-architecture-monorepo/services/user/internal/config"
	"github.com/febrihidayan/go-architecture-monorepo/services/user/internal/usecases/user"
	mongo_repositories "github.com/febrihidayan/go-architecture-monorepo/services/user/tests/mocks/repositories/mongo"
	"github.com/stretchr/testify/suite"
)

type UserUsecaseSuite struct {
	suite.Suite
	cfg         *config.UserConfig
	userRepo    *mongo_repositories.UserRepositoryMock
	userUsecase usecases.UserUsecase
}

func (x *UserUsecaseSuite) SetupTest() {
	x.cfg = &config.UserConfig{}

	x.userRepo = new(mongo_repositories.UserRepositoryMock)

	x.userUsecase = user.NewUserInteractor(x.cfg, x.userRepo)
}

func TestUserUsecase(t *testing.T) {
	suite.Run(t, new(UserUsecaseSuite))
}

type Any []interface{}

func (a Any) Get(i int) interface{} {
	return a[i]
}
