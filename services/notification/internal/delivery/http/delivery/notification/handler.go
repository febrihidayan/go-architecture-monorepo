package notification_handler

import (
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/domain/usecases"
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/internal/config"
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/internal/repositories/factories"
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/internal/usecases/notification"

	"github.com/gorilla/mux"
)

type NotificationHttpHandler struct {
	Cfg                 *config.NotificationConfig
	NotificationUsecase usecases.NotificationUsecase
}

func NewNotificationHttpHandler(
	r *mux.Router,
	config *config.NotificationConfig,
	mongoFactory *factories.MongoFactory,
	grpcClientFactory *factories.GrpcClientFactory,
) {
	handler := &NotificationHttpHandler{
		Cfg: config,
		NotificationUsecase: notification.NewNotificationInteractor(
			config,
			mongoFactory,
			grpcClientFactory,
			nil,
			nil,
		),
	}

	r.HandleFunc("/v1/notifications", handler.GetAll).Methods("GET")
}
