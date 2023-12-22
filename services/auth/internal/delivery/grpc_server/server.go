package grpc_server

import (
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/domain/usecases"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/config"
)

type server struct {
	authUsecase usecases.AuthUsecase
	cfg         config.AuthConfig
}
