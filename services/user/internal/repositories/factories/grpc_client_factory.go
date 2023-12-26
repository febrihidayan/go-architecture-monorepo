package factories

import (
	"github.com/febrihidayan/go-architecture-monorepo/services/user/domain/repositories"
	"github.com/febrihidayan/go-architecture-monorepo/services/user/internal/delivery/grpc_client"
	repository_grpc "github.com/febrihidayan/go-architecture-monorepo/services/user/internal/repositories/grpc"
)

type GrpcClientFactory struct {
	AuthRepo repositories.AuthRepository
}

func NewGrpcFactory(client *grpc_client.ServerClient) *GrpcClientFactory {
	var (
		userRepoGrpc repository_grpc.AuthRepository
	)

	userRepoGrpc = repository_grpc.NewAuthRepository(client.AuthClient)

	return &GrpcClientFactory{
		AuthRepo: &userRepoGrpc,
	}
}
