package notification_handler

import (
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/domain/usecases"
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/internal/config"
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/internal/factories"
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/internal/usecases/notification"

	"github.com/gorilla/mux"
)

type NotificationHttpHandler struct {
	Cfg                 *config.NotificationConfig
	NotificationUsecase usecases.NotificationUsecase
}

func NewNotificationHttpHandler(
	r *mux.Router,
	deps *factories.Dependencies,
) {
	handler := &NotificationHttpHandler{
		Cfg:                 deps.Config,
		NotificationUsecase: notification.NewNotificationInteractor(deps),
	}

	r.HandleFunc("/v1/notifications", handler.GetAll).Methods("GET")
}
