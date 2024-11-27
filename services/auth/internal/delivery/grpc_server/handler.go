package grpc_server

import (
	authPb "github.com/febrihidayan/go-architecture-monorepo/proto/_generated/auth"

	"github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/factories"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/usecases/auth"
	"google.golang.org/grpc"
)

func HandlerAuthServices(s *grpc.Server, deps *factories.Dependencies) {
	authPb.RegisterAuthServicesServer(s, &server{
		authUsecase: auth.NewAuthInteractor(deps),
	})
}
