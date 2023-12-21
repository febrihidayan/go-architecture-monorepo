package profile

import (
	"testing"
	"time"

	"bou.ke/monkey"
	"github.com/febrihidayan/go-architecture-monorepo/services/user/domain/usecases"
	"github.com/febrihidayan/go-architecture-monorepo/services/user/internal/config"
	"github.com/febrihidayan/go-architecture-monorepo/services/user/internal/usecases/profile"
	mongo_repositories "github.com/febrihidayan/go-architecture-monorepo/services/user/tests/mocks/repositories/mongo"
	"github.com/stretchr/testify/suite"
)

type ProfileUsecaseSuite struct {
	suite.Suite
	cfg            *config.UserConfig
	userRepo       *mongo_repositories.UserRepositoryMock
	profileUsecase usecases.ProfileUsecase
}

func (x *ProfileUsecaseSuite) SetupTest() {
	x.cfg = &config.UserConfig{}

	x.userRepo = new(mongo_repositories.UserRepositoryMock)

	x.profileUsecase = profile.NewProfileInteractor(x.cfg, x.userRepo)

	// fake time now for testing
	monkey.Patch(time.Now, func() time.Time {
		return time.Date(2020, 1, 1, 1, 1, 1, 1, time.UTC)
	})
}

func TestUserUsecase(t *testing.T) {
	suite.Run(t, new(ProfileUsecaseSuite))
}

type Any []interface{}

func (a Any) Get(i int) interface{} {
	return a[i]
}
