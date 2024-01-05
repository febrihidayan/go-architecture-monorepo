package grpc_server

import (
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/domain/usecases"
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/internal/config"
)

type server struct {
	cfg                config.NotificationConfig
	templateUsecase    usecases.TemplateUsecase
	deviceTokenUsecase usecases.DeviceTokenUsecase
}
