package rabbitmq

import (
	"fmt"

	"github.com/rabbitmq/amqp091-go"
)

type RabbitMQ struct {
	Conn    *amqp091.Connection
	Channel *amqp091.Channel
}

func NewRabbitMQ(url string) (*RabbitMQ, error) {
	conn, err := amqp091.Dial(url)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return &RabbitMQ{
		Conn:    conn,
		Channel: ch,
	}, nil
}

func (x *RabbitMQ) Close() {
	x.Channel.Close()
	x.Conn.Close()
}

// DeclareExchange for deklarasi exchange
func (x *RabbitMQ) DeclareExchange(exchangeName, exchangeType string) error {
	return x.Channel.ExchangeDeclare(
		exchangeName, exchangeType, true, false, false, false, nil,
	)
}

// Publish for publish message
func (x *RabbitMQ) Publish(exchange, routingKey string, body []byte) error {
	return x.Channel.Publish(
		exchange,
		routingKey,
		false,
		false,
		amqp091.Publishing{
			ContentType: "application/protobuf",
			Body:        body,
		},
	)
}

// SetupQueue binds multiple routing keys to a single queue using a slice
func (x *RabbitMQ) SetupQueue(queueName, exchangeName string, routingKeys []string) error {
	// Declare queue
	_, err := x.Channel.QueueDeclare(
		queueName, // Nama queue
		true,      // Durable
		false,     // Delete when unused
		false,     // Exclusive
		false,     // No-wait
		nil,       // Arguments
	)
	if err != nil {
		return fmt.Errorf("failed to declare queue: %w", err)
	}

	// Looping routing keys dan binding ke queue
	for _, routingKey := range routingKeys {
		err = x.Channel.QueueBind(
			queueName,    // Nama queue
			routingKey,   // Routing key
			exchangeName, // Nama exchange
			false,        // No-wait
			nil,          // Arguments
		)
		if err != nil {
			return fmt.Errorf("failed to bind queue: %w", err)
		}
	}

	return nil
}

// Consume for start consumer
func (x *RabbitMQ) Consume(queueName string) (<-chan amqp091.Delivery, error) {
	return x.Channel.Consume(queueName, "", true, false, false, false, nil)
}
