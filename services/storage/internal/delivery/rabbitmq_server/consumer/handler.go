package customer

import (
	"github.com/febrihidayan/go-architecture-monorepo/pkg/rabbitmq"
	"github.com/febrihidayan/go-architecture-monorepo/services/storage/domain/usecases"
	"github.com/febrihidayan/go-architecture-monorepo/services/storage/internal/config"
	"github.com/febrihidayan/go-architecture-monorepo/services/storage/internal/repositories/factories"
	"github.com/febrihidayan/go-architecture-monorepo/services/storage/internal/usecases/cloud"
)

type CustomerRabbitMQ struct {
	cfg          *config.StorageConfig
	rmq          *rabbitmq.RabbitMQ
	cloudUsecase usecases.CloudUsecase
}

func NewCustomerRabbitMQ(cfg *config.StorageConfig, rmq *rabbitmq.RabbitMQ, mongoFactory *factories.MongoFactory) *CustomerRabbitMQ {
	return &CustomerRabbitMQ{
		rmq:          rmq,
		cloudUsecase: cloud.NewCloudInteractor(cfg, mongoFactory, nil),
	}
}
