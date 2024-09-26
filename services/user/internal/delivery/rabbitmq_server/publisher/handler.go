package publisher

import (
	"github.com/febrihidayan/go-architecture-monorepo/pkg/rabbitmq"
	"github.com/febrihidayan/go-architecture-monorepo/services/user/internal/config"
)

type PublisherRabbitMQ struct {
	cfg *config.UserConfig
	rmq *rabbitmq.RabbitMQ
}

func NewPublisherRabbitMQ(cfg *config.UserConfig, rmq *rabbitmq.RabbitMQ) *PublisherRabbitMQ {
	return &PublisherRabbitMQ{
		cfg: cfg,
		rmq: rmq,
	}
}
