package usecases

import (
	"context"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/exceptions"
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/domain/entities"
	"github.com/stretchr/testify/mock"
)

type NotificationUsecaseMock struct {
	mock.Mock
}

func (x *NotificationUsecaseMock) GetAll(ctx context.Context, payload entities.NotificationQueryParams) (result *entities.NotificationMeta, err *exceptions.CustomError) {
	args := x.Called(payload)

	if n, ok := args.Get(0).(*entities.NotificationMeta); ok {
		result = n
	}

	if n, ok := args.Get(1).(*exceptions.CustomError); ok {
		err = n
	}

	return
}

func (x *NotificationUsecaseMock) SendPushJobs(ctx context.Context, params entities.NotificationSends) (err *exceptions.CustomError) {
	args := x.Called(params)

	if n, ok := args.Get(0).(*exceptions.CustomError); ok {
		err = n
	}

	return
}
