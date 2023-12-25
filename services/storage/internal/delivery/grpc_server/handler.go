package grpc_server

import (
	storagePb "github.com/febrihidayan/go-architecture-monorepo/proto/_generated/storage"

	"github.com/febrihidayan/go-architecture-monorepo/services/storage/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

func HandlerStorageServices(s *grpc.Server, db *mongo.Database, cfg config.StorageConfig) {

	storagePb.RegisterStorageServicesServer(s, &server{})
}
