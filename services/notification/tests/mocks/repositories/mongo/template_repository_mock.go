package mongo_repositories

import (
	"context"

	"github.com/febrihidayan/go-architecture-monorepo/services/notification/domain/entities"
	"github.com/stretchr/testify/mock"
)

type TemplateRepositoryMock struct {
	mock.Mock
}

func (x *TemplateRepositoryMock) Create(ctx context.Context, payload *entities.Template) (err error) {
	args := x.Called(payload)

	if n, ok := args.Get(0).(error); ok {
		err = n
	}

	return
}

func (x *TemplateRepositoryMock) Find(ctx context.Context, id string) (result *entities.Template, err error) {
	args := x.Called(id)

	if n, ok := args.Get(0).(*entities.Template); ok {
		result = n
	}

	if n, ok := args.Get(1).(error); ok {
		err = n
	}

	return
}

func (x *TemplateRepositoryMock) FindByName(ctx context.Context, name string) (result *entities.Template, err error) {
	args := x.Called(name)

	if n, ok := args.Get(0).(*entities.Template); ok {
		result = n
	}

	if n, ok := args.Get(1).(error); ok {
		err = n
	}

	return
}

func (x *TemplateRepositoryMock) Update(ctx context.Context, payload *entities.Template) (err error) {
	args := x.Called(payload)

	if n, ok := args.Get(0).(error); ok {
		err = n
	}

	return
}

func (x *TemplateRepositoryMock) Delete(ctx context.Context, id string) (err error) {
	args := x.Called(id)

	if n, ok := args.Get(0).(error); ok {
		err = n
	}

	return
}
