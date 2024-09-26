package publisher

import (
	"log"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/rabbitmq"
	storagePb "github.com/febrihidayan/go-architecture-monorepo/proto/_generated/storage"
)

func (x *PublisherRabbitMQ) CloudApprove(url []string, _type string) error {
	body, err := rabbitmq.MarshalProto(&storagePb.CloudApproveRequest{
		Url: url,
	})
	if err != nil {
		log.Fatalf("CloudApprove#Failed to marshal message: %v", err)
		return err
	}

	err = x.rmq.Publish(x.cfg.RabbitMQ.Exchange, _type, body)
	if err != nil {
		log.Fatalf("Failed to publish message: %v", err)
	}

	return err
}
