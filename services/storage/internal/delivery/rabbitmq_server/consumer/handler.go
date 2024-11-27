package customer

import (
	"github.com/febrihidayan/go-architecture-monorepo/pkg/rabbitmq"
	"github.com/febrihidayan/go-architecture-monorepo/services/storage/domain/usecases"
	"github.com/febrihidayan/go-architecture-monorepo/services/storage/internal/config"
	"github.com/febrihidayan/go-architecture-monorepo/services/storage/internal/factories"
	"github.com/febrihidayan/go-architecture-monorepo/services/storage/internal/usecases/cloud"
)

type CustomerRabbitMQ struct {
	cfg          *config.StorageConfig
	rmq          *rabbitmq.RabbitMQ
	cloudUsecase usecases.CloudUsecase
}

func NewCustomerRabbitMQ(deps *factories.Dependencies) *CustomerRabbitMQ {
	return &CustomerRabbitMQ{
		rmq:          deps.RabbitMQConn,
		cloudUsecase: cloud.NewCloudInteractor(deps),
	}
}
