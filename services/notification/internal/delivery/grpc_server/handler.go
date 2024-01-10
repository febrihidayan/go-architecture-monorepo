package grpc_server

import (
	notificationPb "github.com/febrihidayan/go-architecture-monorepo/proto/_generated/notification"

	"github.com/febrihidayan/go-architecture-monorepo/services/notification/internal/config"
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/internal/delivery/grpc_client"
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/internal/repositories/factories"
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/internal/services"
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/internal/usecases/device_token"
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/internal/usecases/notification"
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/internal/usecases/template"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

func HandlerNotificationServices(
	s *grpc.Server,
	db *mongo.Database,
	grpcClient *grpc_client.ServerClient,
	firebaseGoogleService *services.FirebaseGoogleService,
	cfg config.NotificationConfig) {
	mongoFactory := factories.NewMongoFactory(db)
	grpcClientFactory := factories.NewGrpcFactory(grpcClient)

	notificationPb.RegisterNotificationServicesServer(s, &server{
		templateUsecase:     template.NewTemplateInteractor(&cfg, mongoFactory),
		deviceTokenUsecase:  device_token.NewDeviceTokenInteractor(&cfg, mongoFactory),
		notificationUsecase: notification.NewNotificationInteractor(&cfg, mongoFactory, grpcClientFactory, firebaseGoogleService),
	})
}
