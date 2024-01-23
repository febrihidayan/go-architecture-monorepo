package device_token_handler

import (
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/domain/usecases"
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/internal/config"
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/internal/repositories/factories"
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/internal/usecases/device_token"

	"github.com/gorilla/mux"
)

type DeviceTokenHttpHandler struct {
	Cfg                *config.NotificationConfig
	DeviceTokenUsecase usecases.DeviceTokenUsecase
}

func NewDeviceTokenHttpHandler(
	r *mux.Router,
	config *config.NotificationConfig,
	mongoFactory *factories.MongoFactory,
) {
	handler := &DeviceTokenHttpHandler{
		Cfg: config,
		DeviceTokenUsecase: device_token.NewDeviceTokenInteractor(
			config,
			mongoFactory,
		),
	}

	r.HandleFunc("/v1/notification/device-token", handler.Create).Methods("POST")
}
