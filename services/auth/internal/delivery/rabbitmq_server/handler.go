package rabbitmq_server

import (
	"context"
	"log"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/rabbitmq"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/domain/entities"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/config"
	customer "github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/delivery/rabbitmq_server/consumer"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/repositories/factories"
)

type RabbitMQServer struct {
	ctx      context.Context
	cfg      *config.AuthConfig
	rmq      *rabbitmq.RabbitMQ
	customer *customer.CustomerRabbitMQ
}

func HandlerRabbitMQServices(deps *factories.Dependencies, rmq *rabbitmq.RabbitMQ) *RabbitMQServer {

	return &RabbitMQServer{
		ctx:      context.Background(),
		cfg:      deps.Config,
		rmq:      rmq,
		customer: customer.NewCustomerRabbitMQ(deps, rmq),
	}
}

func (x *RabbitMQServer) Worker() {
	x.CustomerListen()
}

func (x *RabbitMQServer) CustomerListen() {
	err := x.rmq.DeclareExchange(x.cfg.RabbitMQ.Exchange, "direct")
	if err != nil {
		log.Fatalf("Failed to declare exchange: %v", err)
	}

	routingKeys := []string{
		entities.RABBITMQ_AUTH_AUTH_DELETE,
	}

	if err := x.rmq.SetupQueue("auth.queue", x.cfg.RabbitMQ.Exchange, routingKeys); err != nil {
		log.Fatalf("Failed to setup auth queue: %v", err)
	}

	msgs, err := x.rmq.Consume("auth.queue")
	if err != nil {
		log.Fatal(err)
	}

	forever := make(chan bool)

	go func() {
		for msg := range msgs {
			log.Println("CustomerListen:", msg.RoutingKey)

			switch msg.RoutingKey {
			case entities.RABBITMQ_AUTH_AUTH_DELETE:
				if err := x.customer.AuthDelete(x.ctx, msg.Body); err != nil {
					log.Println("CustomerListen#AuthDelete", err)
				}
			default:
				log.Printf("Unknown event: %s", msg.RoutingKey)
			}
		}
	}()

	log.Printf("RabbitMQ#Waiting for messages")
	<-forever
}
