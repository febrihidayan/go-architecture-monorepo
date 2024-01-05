package device_token

import (
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/domain/repositories"
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/internal/config"
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/internal/repositories/factories"
)

type deviceTokenInteractor struct {
	cfg             *config.NotificationConfig
	deviceTokenRepo repositories.DeviceTokenRepository
}

func NewDeviceTokenInteractor(
	config *config.NotificationConfig,
	mongoFactory *factories.MongoFactory,
) *deviceTokenInteractor {

	return &deviceTokenInteractor{
		cfg:             config,
		deviceTokenRepo: mongoFactory.DeviceTokenRepo,
	}
}
