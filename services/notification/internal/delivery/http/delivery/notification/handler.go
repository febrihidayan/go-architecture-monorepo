package notification_handler

import (
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/domain/usecases"
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/internal/config"
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/internal/repositories/factories"
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/internal/usecases/notification"

	"github.com/gorilla/mux"
)

type notificationHttpHandler struct {
	cfg                 *config.NotificationConfig
	notificationUsecase usecases.NotificationUsecase
}

func TemplateHttpHandler(
	r *mux.Router,
	config *config.NotificationConfig,
	mongoFactory *factories.MongoFactory,
	grpcClientFactory *factories.GrpcClientFactory,
) {
	handler := &notificationHttpHandler{
		cfg: config,
		notificationUsecase: notification.NewNotificationInteractor(
			config,
			mongoFactory,
			grpcClientFactory,
		),
	}

	r.HandleFunc("/v1/notifications", handler.GetAll).Methods("GET")
}
