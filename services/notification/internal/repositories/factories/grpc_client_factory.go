package factories

import (
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/domain/repositories"
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/internal/delivery/grpc_client"
	repository_grpc "github.com/febrihidayan/go-architecture-monorepo/services/notification/internal/repositories/grpc"
)

type GrpcClientFactory struct {
	UserRepo repositories.UserRepository
}

func NewGrpcFactory(client *grpc_client.ServerClient) *GrpcClientFactory {
	var (
		userRepoGrpc repository_grpc.UserRepository
	)

	userRepoGrpc = repository_grpc.NewUserRepository(client.UserClient)

	return &GrpcClientFactory{
		UserRepo: &userRepoGrpc,
	}
}
