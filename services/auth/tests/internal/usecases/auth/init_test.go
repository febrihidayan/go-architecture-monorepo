package auth

import (
	"testing"
	"time"

	"bou.ke/monkey"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/domain/usecases"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/config"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/usecases/auth"
	grpc_repositories "github.com/febrihidayan/go-architecture-monorepo/services/auth/tests/mocks/repositories/grpc"
	mongo_repositories "github.com/febrihidayan/go-architecture-monorepo/services/auth/tests/mocks/repositories/mongo"
	"github.com/stretchr/testify/suite"
)

type AuthUsecaseSuite struct {
	suite.Suite
	cfg         *config.AuthConfig
	authRepo    *mongo_repositories.AuthRepositoryMock
	userRepo    *grpc_repositories.UserRepositoryMock
	authUsecase usecases.AuthUsecase
}

func (x *AuthUsecaseSuite) SetupTest() {
	x.cfg = &config.AuthConfig{}

	x.authRepo = new(mongo_repositories.AuthRepositoryMock)
	x.userRepo = new(grpc_repositories.UserRepositoryMock)

	x.authUsecase = auth.NewAuthInteractor(x.cfg, x.authRepo, x.userRepo)

	// fake time now for testing
	monkey.Patch(time.Now, func() time.Time {
		return time.Date(2020, 1, 1, 1, 1, 1, 1, time.UTC)
	})
}

func TestAuthUsecase(t *testing.T) {
	suite.Run(t, new(AuthUsecaseSuite))
}

type Any []interface{}

func (a Any) Get(i int) interface{} {
	return a[i]
}
