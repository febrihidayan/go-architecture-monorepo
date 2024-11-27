package rabbitmq_server

import (
	"context"
	"log"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/rabbitmq"
	"github.com/febrihidayan/go-architecture-monorepo/services/storage/domain/entities"
	"github.com/febrihidayan/go-architecture-monorepo/services/storage/internal/config"
	customer "github.com/febrihidayan/go-architecture-monorepo/services/storage/internal/delivery/rabbitmq_server/consumer"
	"github.com/febrihidayan/go-architecture-monorepo/services/storage/internal/factories"
)

type RabbitMQServer struct {
	ctx      context.Context
	cfg      *config.StorageConfig
	rmq      *rabbitmq.RabbitMQ
	customer *customer.CustomerRabbitMQ
}

func HandlerRabbitMQServices(deps *factories.Dependencies) *RabbitMQServer {
	return &RabbitMQServer{
		ctx:      context.Background(),
		cfg:      deps.Config,
		rmq:      deps.RabbitMQConn,
		customer: customer.NewCustomerRabbitMQ(deps),
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
		entities.RABBITMQ_STORAGE_CLOUD_UPDATE,
		entities.RABBITMQ_STORAGE_CLOUD_DELETE,
	}

	if err := x.rmq.SetupQueue("storage.queue", x.cfg.RabbitMQ.Exchange, routingKeys); err != nil {
		log.Fatalf("Failed to setup storage queue: %v", err)
	}

	msgs, err := x.rmq.Consume("storage.queue")
	if err != nil {
		log.Fatal(err)
	}

	forever := make(chan bool)

	go func() {
		for msg := range msgs {
			log.Println("CustomerListen:", msg.RoutingKey)

			switch msg.RoutingKey {
			case entities.RABBITMQ_STORAGE_CLOUD_UPDATE:
				if err := x.customer.UpdateCloudApprove(x.ctx, msg.Body); err != nil {
					log.Println("CustomerListen#UpdateCloudApprove", err)
				}
			case entities.RABBITMQ_STORAGE_CLOUD_DELETE:
				if err := x.customer.DeleteCloudApprove(x.ctx, msg.Body); err != nil {
					log.Println("CustomerListen#DeleteCloudApprove", err)
				}
			default:
				log.Printf("Unknown event: %s", msg.RoutingKey)
			}
		}
	}()

	log.Printf("RabbitMQ#Waiting for messages")
	<-forever
}
