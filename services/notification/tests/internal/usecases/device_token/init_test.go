package device_token

import (
	"testing"
	"time"

	"bou.ke/monkey"
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/domain/usecases"
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/internal/config"
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/internal/factories"
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/internal/usecases/device_token"
	mongo_repositories "github.com/febrihidayan/go-architecture-monorepo/services/notification/tests/mocks/repositories/mongo"
	"github.com/stretchr/testify/suite"
)

type DeviceTokenUsecaseSuite struct {
	suite.Suite
	cfg                *config.NotificationConfig
	deps               *factories.Dependencies
	mongoFactory       *factories.MongoFactory
	deviceTokenRepo    *mongo_repositories.DeviceTokenRepositoryMock
	deviceTokenUsecase usecases.DeviceTokenUsecase
}

func (x *DeviceTokenUsecaseSuite) SetupTest() {
	x.cfg = &config.NotificationConfig{}

	x.deviceTokenRepo = new(mongo_repositories.DeviceTokenRepositoryMock)

	x.mongoFactory = &factories.MongoFactory{
		DeviceTokenRepo: x.deviceTokenRepo,
	}

	x.deps = &factories.Dependencies{
		Config:       x.cfg,
		MongoFactory: x.mongoFactory,
	}

	x.deviceTokenUsecase = device_token.NewDeviceTokenInteractor(x.deps)

	// fake time now for testing
	monkey.Patch(time.Now, func() time.Time {
		return time.Date(2020, 1, 1, 1, 1, 1, 1, time.UTC)
	})
}

func TestDeviceTokenUsecase(t *testing.T) {
	suite.Run(t, new(DeviceTokenUsecaseSuite))
}

type Any []interface{}

func (a Any) Get(i int) interface{} {
	return a[i]
}
