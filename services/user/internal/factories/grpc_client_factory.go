package factories

import (
	"github.com/febrihidayan/go-architecture-monorepo/services/user/domain/repositories"
	"github.com/febrihidayan/go-architecture-monorepo/services/user/internal/delivery/grpc_client"
	repository_grpc "github.com/febrihidayan/go-architecture-monorepo/services/user/internal/repositories/grpc"
)

type GrpcClientFactory struct {
	AuthRepo    repositories.AuthRepository
	StorageRepo repositories.StorageRepository
}

func NewGrpcFactory(client *grpc_client.ServerClient) *GrpcClientFactory {
	var (
		authRepoGrpc    repository_grpc.AuthRepository
		storageRepoGrpc repository_grpc.StorageRepository
	)

	authRepoGrpc = repository_grpc.NewAuthRepository(client.AuthClient)
	storageRepoGrpc = repository_grpc.NewStorageRepository(client.StorageClient)

	return &GrpcClientFactory{
		AuthRepo:    &authRepoGrpc,
		StorageRepo: &storageRepoGrpc,
	}
}
