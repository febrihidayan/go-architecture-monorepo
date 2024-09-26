package rabbitmq

import (
	"github.com/rabbitmq/amqp091-go"
)

type RabbitMQInterface interface {
	Close()
	DeclareExchange(exchangeName, exchangeType string) error
	Publish(exchange, routingKey string, body []byte) error
	SetupQueue(queueName, exchangeName string, routingKeys []string) error
	Consume(queueName string) (<-chan amqp091.Delivery, error)
}
