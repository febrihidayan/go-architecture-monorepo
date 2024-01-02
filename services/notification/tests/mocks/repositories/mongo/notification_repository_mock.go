package mongo_repositories

import (
	"context"

	"github.com/febrihidayan/go-architecture-monorepo/services/notification/domain/entities"
	"github.com/stretchr/testify/mock"
)

type NotificationRepositoryMock struct {
	mock.Mock
}

func (x *NotificationRepositoryMock) Create(ctx context.Context, payload *entities.Notification) (err error) {
	args := x.Called(payload)

	if n, ok := args.Get(0).(error); ok {
		err = n
	}

	return
}

func (x *NotificationRepositoryMock) GetAll(ctx context.Context, params *entities.NotificationQueryParams) (result []*entities.Notification, total int, err error) {
	args := x.Called(params)

	if n, ok := args.Get(0).([]*entities.Notification); ok {
		result = n
	}

	if n, ok := args.Get(1).(int); ok {
		total = n
	}

	if n, ok := args.Get(2).(error); ok {
		err = n
	}

	return
}

func (x *NotificationRepositoryMock) Delete(ctx context.Context, id string) (err error) {
	args := x.Called(id)

	if n, ok := args.Get(0).(error); ok {
		err = n
	}

	return
}
