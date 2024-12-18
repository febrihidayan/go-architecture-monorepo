package user

import (
	"github.com/febrihidayan/go-architecture-monorepo/services/user/domain/repositories"
	"github.com/febrihidayan/go-architecture-monorepo/services/user/internal/config"
	"github.com/febrihidayan/go-architecture-monorepo/services/user/internal/delivery/rabbitmq_server/publisher"
	"github.com/febrihidayan/go-architecture-monorepo/services/user/internal/factories"
)

type userInteractor struct {
	cfg             *config.UserConfig
	userRepo        repositories.UserRepository
	authGrpcRepo    repositories.AuthRepository
	storageGrpcRepo repositories.StorageRepository
	rabbitmqRepo    repositories.RabbitMQRepository
}

func NewUserInteractor(
	config *config.UserConfig,
	mongoFactory *factories.MongoFactory,
	grpcFactory *factories.GrpcClientFactory,
	rabbitmq *publisher.PublisherRabbitMQ,
) *userInteractor {

	return &userInteractor{
		cfg:             config,
		userRepo:        mongoFactory.UserRepo,
		authGrpcRepo:    grpcFactory.AuthRepo,
		storageGrpcRepo: grpcFactory.StorageRepo,
		rabbitmqRepo:    rabbitmq,
	}
}
