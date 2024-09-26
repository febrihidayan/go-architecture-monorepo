package profile

import (
	"github.com/febrihidayan/go-architecture-monorepo/services/user/domain/repositories"
	"github.com/febrihidayan/go-architecture-monorepo/services/user/internal/config"
	"github.com/febrihidayan/go-architecture-monorepo/services/user/internal/delivery/rabbitmq_server/publisher"
	"github.com/febrihidayan/go-architecture-monorepo/services/user/internal/repositories/factories"
)

type profileInteractor struct {
	cfg          *config.UserConfig
	userRepo     repositories.UserRepository
	rabbitmqRepo repositories.RabbitMQRepository
}

func NewProfileInteractor(
	config *config.UserConfig,
	mongoFactory *factories.MongoFactory,
	rabbitmq *publisher.PublisherRabbitMQ,
) *profileInteractor {

	return &profileInteractor{
		cfg:          config,
		userRepo:     mongoFactory.UserRepo,
		rabbitmqRepo: rabbitmq,
	}
}
