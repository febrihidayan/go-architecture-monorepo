package grpc_server

import (
	"github.com/febrihidayan/go-architecture-monorepo/services/user/domain/usecases"
	"github.com/febrihidayan/go-architecture-monorepo/services/user/internal/config"
)

type server struct {
	userUsecase usecases.UserUsecase
	cfg         config.UserConfig
}
