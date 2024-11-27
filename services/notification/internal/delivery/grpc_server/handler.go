package grpc_server

import (
	notificationPb "github.com/febrihidayan/go-architecture-monorepo/proto/_generated/notification"

	"github.com/febrihidayan/go-architecture-monorepo/services/notification/internal/factories"
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/internal/usecases/device_token"
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/internal/usecases/notification"
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/internal/usecases/template"
	"google.golang.org/grpc"
)

func HandlerNotificationServices(s *grpc.Server, deps *factories.Dependencies) {
	notificationPb.RegisterNotificationServicesServer(s, &server{
		templateUsecase:     template.NewTemplateInteractor(deps),
		deviceTokenUsecase:  device_token.NewDeviceTokenInteractor(deps),
		notificationUsecase: notification.NewNotificationInteractor(deps),
	})
}
