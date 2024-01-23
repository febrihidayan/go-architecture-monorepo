package usecases

import (
	"context"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/exceptions"
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/domain/entities"
	"github.com/stretchr/testify/mock"
)

type DeviceTokenUsecaseMock struct {
	mock.Mock
}

func (x *DeviceTokenUsecaseMock) Create(ctx context.Context, payload entities.DeviceTokenDto) (result *entities.DeviceToken, err *exceptions.CustomError) {
	args := x.Called(payload)

	if n, ok := args.Get(0).(*entities.DeviceToken); ok {
		result = n
	}

	if n, ok := args.Get(1).(*exceptions.CustomError); ok {
		err = n
	}

	return
}
