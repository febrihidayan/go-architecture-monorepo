package device_token_handler

import (
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/domain/usecases"
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/internal/config"
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/internal/repositories/factories"
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/internal/usecases/device_token"

	"github.com/gorilla/mux"
)

type deviceTokenHttpHandler struct {
	cfg                *config.NotificationConfig
	deviceTokenUsecase usecases.DeviceTokenUsecase
}

func TemplateHttpHandler(
	r *mux.Router,
	config *config.NotificationConfig,
	mongoFactory *factories.MongoFactory,
) {
	handler := &deviceTokenHttpHandler{
		cfg: config,
		deviceTokenUsecase: device_token.NewDeviceTokenInteractor(
			config,
			mongoFactory,
		),
	}

	r.HandleFunc("/v1/notification/device-token", handler.Create).Methods("POST")
}
