package grpc_server

import (
	notificationPb "github.com/febrihidayan/go-architecture-monorepo/proto/_generated/notification"

	"github.com/febrihidayan/go-architecture-monorepo/services/notification/internal/config"
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/internal/repositories/factories"
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/internal/usecases/template"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

func HandlerNotificationServices(s *grpc.Server, db *mongo.Database, cfg config.NotificationConfig) {
	mongoFactory := factories.NewMongoFactory(db)

	notificationPb.RegisterNotificationServicesServer(s, &server{
		templateUsecase: template.NewTemplateInteractor(&cfg, mongoFactory),
	})
}
