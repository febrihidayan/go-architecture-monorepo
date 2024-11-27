package device_token

import (
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/domain/repositories"
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/internal/config"
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/internal/factories"
)

type deviceTokenInteractor struct {
	cfg             *config.NotificationConfig
	deviceTokenRepo repositories.DeviceTokenRepository
}

func NewDeviceTokenInteractor(deps *factories.Dependencies) *deviceTokenInteractor {
	return &deviceTokenInteractor{
		cfg:             deps.Config,
		deviceTokenRepo: deps.MongoFactory.DeviceTokenRepo,
	}
}
