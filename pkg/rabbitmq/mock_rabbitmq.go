package rabbitmq

import (
	"github.com/rabbitmq/amqp091-go"
	"github.com/stretchr/testify/mock"
)

type MockRabbitMQ struct {
	mock.Mock
}

func (m *MockRabbitMQ) Close() {
	m.Called()
}

func (m *MockRabbitMQ) DeclareExchange(exchangeName, exchangeType string) error {
	args := m.Called(exchangeName, exchangeType)
	return args.Error(0)
}

func (m *MockRabbitMQ) Publish(exchange, routingKey string, body []byte) error {
	args := m.Called(exchange, routingKey, body)
	return args.Error(0)
}

func (m *MockRabbitMQ) SetupQueue(queueName, exchangeName string, routingKeys []string) error {
	args := m.Called(queueName, exchangeName, routingKeys)
	return args.Error(0)
}

func (m *MockRabbitMQ) Consume(queueName string) (<-chan amqp091.Delivery, error) {
	args := m.Called(queueName)
	return args.Get(0).(<-chan amqp091.Delivery), args.Error(1)
}
