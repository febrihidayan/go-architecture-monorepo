package grpc_server

import (
	storagePb "github.com/febrihidayan/go-architecture-monorepo/proto/_generated/storage"

	"github.com/febrihidayan/go-architecture-monorepo/services/storage/internal/factories"
	"github.com/febrihidayan/go-architecture-monorepo/services/storage/internal/usecases/cloud"
	"google.golang.org/grpc"
)

func HandlerStorageServices(s *grpc.Server, deps *factories.Dependencies) {
	storagePb.RegisterStorageServicesServer(s, &server{
		cloudUsecase: cloud.NewCloudInteractor(deps),
	})
}
