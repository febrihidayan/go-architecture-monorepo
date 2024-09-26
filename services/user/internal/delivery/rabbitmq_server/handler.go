package rabbitmq_server

import (
	"github.com/febrihidayan/go-architecture-monorepo/pkg/rabbitmq"
	"github.com/febrihidayan/go-architecture-monorepo/services/user/internal/config"
)

type RabbitMQServer struct {
	cfg *config.UserConfig
	rmq *rabbitmq.RabbitMQ
}

func HandlerRabbitMQServices(cfg *config.UserConfig, rmq *rabbitmq.RabbitMQ) *RabbitMQServer {
	return &RabbitMQServer{
		cfg: cfg,
		rmq: rmq,
	}
}
