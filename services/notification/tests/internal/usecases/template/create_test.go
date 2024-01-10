package template

import (
	"context"
	"errors"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/common"
	"github.com/febrihidayan/go-architecture-monorepo/pkg/exceptions"
	"github.com/febrihidayan/go-architecture-monorepo/pkg/lang"
	"github.com/febrihidayan/go-architecture-monorepo/pkg/utils"
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/domain/entities"
	"github.com/hashicorp/go-multierror"
	"github.com/stretchr/testify/mock"
)

func (x *TemplateUsecaseSuite) TestCreate() {
	id := common.NewID()
	var (
		template *entities.Template
	)

	payloadDto := entities.TemplateDto{
		ID:   &id,
		Name: "register_user",
		Type: "in-app",
		Data: `{"title":{"en":"Welcome","id":"Selamat Datang"}}`,
	}

	payloadEmptyDto := entities.TemplateDto{
		ID:   &id,
		Name: "register_user",
		Type: "in-app",
		Data: "",
	}

	template = &entities.Template{
		ID:        id,
		Name:      "register_user",
		Type:      "in-app",
		Data:      `{"title":{"en":"Welcome","id":"Selamat Datang"}}`,
		CreatedAt: utils.TimeUTC(),
		UpdatedAt: utils.TimeUTC(),
	}

	args := []struct {
		name  string
		args  Any
		tests func(arg Any)
	}{
		{
			name: "Success Positive Case",
			tests: func(arg Any) {
				x.templateRepo.Mock.On("FindByName", payloadDto.Name).Return(nil, errors.New(mock.Anything))

				x.templateRepo.Mock.On("Create", template).Return(nil)

				result, err := x.templateUsecase.Create(context.Background(), payloadDto)
				x.Nil(err)
				x.Equal(template, result)
			},
		},
		{
			name: "Failed Name Ready Negatif Case",
			tests: func(arg Any) {
				x.templateRepo.Mock.On("FindByName", payloadDto.Name).Return(template, nil)

				_, err := x.templateUsecase.Create(context.Background(), payloadDto)
				e := &exceptions.CustomError{
					Status: exceptions.ERRREPOSITORY,
					Errors: multierror.Append(lang.ErrTemplateAlready),
				}

				x.Equal(err, e)
			},
		},
		{
			name: "Failed Domain Negatif Case",
			tests: func(arg Any) {
				x.templateRepo.Mock.On("FindByName", payloadEmptyDto.Name).Return(nil, errors.New(mock.Anything))

				_, err := x.templateUsecase.Create(context.Background(), payloadEmptyDto)
				e := &exceptions.CustomError{
					Status: exceptions.ERRDOMAIN,
					Errors: multierror.Append(lang.Trans("filled", "Data")),
				}

				x.Equal(err, e)
			},
		},
		{
			name: "Failed Repository Negatif Case",
			tests: func(arg Any) {
				x.templateRepo.Mock.On("FindByName", payloadDto.Name).Return(nil, errors.New(mock.Anything))

				x.templateRepo.Mock.On("Create", template).Return(errors.New(mock.Anything))

				_, err := x.templateUsecase.Create(context.Background(), payloadDto)
				e := &exceptions.CustomError{
					Status: exceptions.ERRREPOSITORY,
					Errors: multierror.Append(errors.New(mock.Anything)),
				}

				x.Equal(err, e)
			},
		},
	}

	for _, arg := range args {
		x.Run(arg.name, func() {
			x.SetupTest()
			arg.tests(arg.args)
		})
	}
}
