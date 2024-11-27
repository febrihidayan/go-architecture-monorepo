package customer

import (
	"github.com/febrihidayan/go-architecture-monorepo/pkg/rabbitmq"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/domain/usecases"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/repositories/factories"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/usecases/auth"
)

type CustomerRabbitMQ struct {
	rmq         *rabbitmq.RabbitMQ
	authUsecase usecases.AuthUsecase
}

func NewCustomerRabbitMQ(deps *factories.Dependencies, rmq *rabbitmq.RabbitMQ) *CustomerRabbitMQ {
	return &CustomerRabbitMQ{
		rmq:         rmq,
		authUsecase: auth.NewAuthInteractor(deps),
	}
}
