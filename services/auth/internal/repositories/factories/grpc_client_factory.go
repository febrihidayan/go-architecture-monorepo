package factories

import (
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/domain/repositories"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/delivery/grpc_client"
	repository_grpc "github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/repositories/grpc"
)

type GrpcClientFactory struct {
	UserRepo         repositories.UserRepository
	NotificationRepo repositories.NotificationRepository
}

func NewGrpcFactory(client *grpc_client.ServerClient) *GrpcClientFactory {
	var (
		userRepoGrpc         repository_grpc.UserRepository
		notificationRepoGrpc repository_grpc.NotificationRepository
	)

	userRepoGrpc = repository_grpc.NewUserRepository(client.UserClient)
	notificationRepoGrpc = repository_grpc.NewNotificationRepository(client.NotificationClient)

	return &GrpcClientFactory{
		UserRepo:         &userRepoGrpc,
		NotificationRepo: &notificationRepoGrpc,
	}
}
