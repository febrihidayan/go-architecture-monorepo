package rabbitmq_repositories

import "github.com/stretchr/testify/mock"

type RabbitMQRepositoryMock struct {
	mock.Mock
}

func (m *RabbitMQRepositoryMock) CloudApprove(url []string, _type string) error {
	arg := m.Called(url, _type)
	var err error

	if n, ok := arg.Get(0).(error); ok {
		err = n
	}

	return err
}

func (m *RabbitMQRepositoryMock) AuthDelete(id string) error {
	arg := m.Called(id)
	var err error

	if n, ok := arg.Get(0).(error); ok {
		err = n
	}

	return err
}
