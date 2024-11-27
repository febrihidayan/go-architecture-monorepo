package publisher

import (
	"log"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/rabbitmq"
	authPb "github.com/febrihidayan/go-architecture-monorepo/proto/_generated/auth"
	"github.com/febrihidayan/go-architecture-monorepo/services/user/domain/entities"
)

func (x *PublisherRabbitMQ) AuthDelete(id string) error {
	body, err := rabbitmq.MarshalProto(&authPb.FindByIdRequest{
		Id: id,
	})
	if err != nil {
		log.Fatalf("AuthDelete#1Failed to marshal message: %v", err)
		return err
	}

	err = x.rmq.Publish(x.cfg.RabbitMQ.Exchange, entities.RABBITMQ_AUTH_AUTH_DELETE, body)
	if err != nil {
		log.Fatalf("AuthDelete#2Failed to publish message: %v", err)
	}

	return err
}
