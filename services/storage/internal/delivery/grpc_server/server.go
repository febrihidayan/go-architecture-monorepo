package grpc_server

import (
	"github.com/febrihidayan/go-architecture-monorepo/services/storage/domain/usecases"
	"github.com/febrihidayan/go-architecture-monorepo/services/storage/internal/config"
)

type server struct {
	cfg          config.StorageConfig
	cloudUsecase usecases.CloudUsecase
}
