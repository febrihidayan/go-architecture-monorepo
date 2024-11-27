package rabbitmq

import (
	"fmt"
	"log"

	"github.com/rabbitmq/amqp091-go"
)

type RabbitMQ struct {
	Conn    *amqp091.Connection
	Channel *amqp091.Channel
}

// NewRabbitMQ initializes a RabbitMQ connection and channel
func NewRabbitMQ(url string) (*RabbitMQ, error) {
	conn, err := amqp091.Dial(url)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to RabbitMQ: %w", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, fmt.Errorf("failed to open a channel: %w", err)
	}

	return &RabbitMQ{
		Conn:    conn,
		Channel: ch,
	}, nil
}

// Close closes the RabbitMQ connection and channel
func (x *RabbitMQ) Close() {
	if x.Channel != nil {
		if err := x.Channel.Close(); err != nil {
			log.Printf("failed to close channel: %v", err)
		}
	}
	if x.Conn != nil {
		if err := x.Conn.Close(); err != nil {
			log.Printf("failed to close connection: %v", err)
		}
	}
}

// DeclareExchange declares an exchange if it doesn't exist
func (x *RabbitMQ) DeclareExchange(exchangeName, exchangeType string) error {
	return x.Channel.ExchangeDeclare(
		exchangeName, // Exchange name
		exchangeType, // Exchange type (e.g., "direct", "topic", "fanout")
		true,      // Durable
		false,        // Auto-deleted
		false,        // Internal
		false,        // No-wait
		nil,          // Arguments
	)
}

// Publish sends a message to the specified exchange and routing key
func (x *RabbitMQ) Publish(exchange, routingKey string, body []byte) error {
	return x.Channel.Publish(
		exchange,   // Exchange name
		routingKey, // Routing key
		false,      // Mandatory
		false,      // Immediate
		amqp091.Publishing{
			ContentType: "application/protobuf", // You can change this if needed
			Body:        body,
		},
	)
}

// SetupQueue declares a queue and binds it to multiple routing keys
func (x *RabbitMQ) SetupQueue(queueName, exchangeName string, routingKeys []string) error {
	// Declare the queue
	_, err := x.Channel.QueueDeclare(
		queueName, // Queue name
		true,   // Durable
		false,     // Auto-deleted
		false,     // Exclusive
		false,     // No-wait
		nil,       // Arguments
	)
	if err != nil {
		return fmt.Errorf("failed to declare queue: %w", err)
	}

	// Bind each routing key to the queue
	for _, routingKey := range routingKeys {
		err := x.Channel.QueueBind(
			queueName,    // Queue name
			routingKey,   // Routing key
			exchangeName, // Exchange name
			false,        // No-wait
			nil,          // Arguments
		)
		if err != nil {
			return fmt.Errorf("failed to bind queue to exchange: %w", err)
		}
	}

	return nil
}

// Consume starts consuming messages from the specified queue
func (x *RabbitMQ) Consume(queueName string) (<-chan amqp091.Delivery, error) {
	return x.Channel.Consume(
		queueName, // Queue name
		"",        // Consumer tag
		true,   // Auto-acknowledge
		false,     // Exclusive
		false,     // No-local
		false,     // No-wait
		nil,       // Arguments
	)
}
