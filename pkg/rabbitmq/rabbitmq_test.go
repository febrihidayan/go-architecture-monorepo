package rabbitmq

import (
	"testing"

	"github.com/rabbitmq/amqp091-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestPublishMessage(t *testing.T) {
	mockRabbitMQ := new(MockRabbitMQ)

	// Define expected behavior
	mockRabbitMQ.On("Publish", "my-exchange", "my-routing-key", mock.Anything).Return(nil)

	// Call the Publish method
	err := mockRabbitMQ.Publish("my-exchange", "my-routing-key", []byte("test message"))

	// Assert that the mock was called with expected parameters
	mockRabbitMQ.AssertCalled(t, "Publish", "my-exchange", "my-routing-key", []byte("test message"))

	// Check that no error was returned
	assert.NoError(t, err)
}

func TestDeclareExchange(t *testing.T) {
	mockRabbitMQ := new(MockRabbitMQ)

	// Setup mock to expect call to DeclareExchange
	mockRabbitMQ.On("DeclareExchange", "my-exchange", "direct").Return(nil)

	// Call the DeclareExchange method
	err := mockRabbitMQ.DeclareExchange("my-exchange", "direct")

	// Assert that the method was called correctly
	mockRabbitMQ.AssertCalled(t, "DeclareExchange", "my-exchange", "direct")

	// Assert no error returned
	assert.NoError(t, err)
}

func TestSetupQueue(t *testing.T) {
	mockRabbitMQ := new(MockRabbitMQ)

	// Setup mock to expect call to SetupQueue
	mockRabbitMQ.On("SetupQueue", "my-queue", "my-exchange", []string{"routing-key"}).Return(nil)

	// Call the SetupQueue method
	err := mockRabbitMQ.SetupQueue("my-queue", "my-exchange", []string{"routing-key"})

	// Assert that the method was called correctly
	mockRabbitMQ.AssertCalled(t, "SetupQueue", "my-queue", "my-exchange", []string{"routing-key"})

	// Assert no error returned
	assert.NoError(t, err)
}

func TestConsume(t *testing.T) {
	mockRabbitMQ := new(MockRabbitMQ)

	// Setup mock to expect call to Consume
	mockRabbitMQ.On("Consume", "my-queue").Return(make(<-chan amqp091.Delivery), nil)

	// Call the Consume method
	deliveryChannel, err := mockRabbitMQ.Consume("my-queue")

	// Assert that the method was called correctly
	mockRabbitMQ.AssertCalled(t, "Consume", "my-queue")

	// Assert no error returned and deliveryChannel is not nil
	assert.NoError(t, err)
	assert.NotNil(t, deliveryChannel)
}
