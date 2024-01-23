package usecases

import (
	"context"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/exceptions"
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/domain/entities"
	"github.com/stretchr/testify/mock"
)

type TemplateUsecaseMock struct {
	mock.Mock
}

func (x *TemplateUsecaseMock) Create(ctx context.Context, payload entities.TemplateDto) (result *entities.Template, err *exceptions.CustomError) {
	args := x.Called(payload)

	if n, ok := args.Get(0).(*entities.Template); ok {
		result = n
	}

	if n, ok := args.Get(1).(*exceptions.CustomError); ok {
		err = n
	}

	return
}
