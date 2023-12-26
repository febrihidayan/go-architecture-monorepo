package grpc_server

import (
	storagePb "github.com/febrihidayan/go-architecture-monorepo/proto/_generated/storage"

	"github.com/febrihidayan/go-architecture-monorepo/services/storage/internal/config"
	"github.com/febrihidayan/go-architecture-monorepo/services/storage/internal/repositories/factories"
	"github.com/febrihidayan/go-architecture-monorepo/services/storage/internal/usecases/cloud"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

func HandlerStorageServices(s *grpc.Server, db *mongo.Database, cfg config.StorageConfig) {

	mongoFactory := factories.NewMongoFactory(db)

	storagePb.RegisterStorageServicesServer(s, &server{
		cloudUsecase: cloud.NewCloudInteractor(&cfg, mongoFactory, nil),
	})
}
