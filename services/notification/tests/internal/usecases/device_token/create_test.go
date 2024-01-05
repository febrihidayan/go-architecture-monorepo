package device_token

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

func (x *DeviceTokenUsecaseSuite) TestCreate() {
	id := common.NewID()
	var (
		template *entities.DeviceToken
	)

	payloadDto := entities.DeviceTokenDto{
		ID:     &id,
		UserId: id.String(),
		Token:  "sdfghjkvftyu",
	}

	payloadEmptyDto := entities.DeviceTokenDto{
		ID:     &id,
		UserId: id.String(),
		Token:  "",
	}

	template = &entities.DeviceToken{
		ID:        id,
		UserId:    id.String(),
		Token:     "sdfghjkvftyu",
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
				x.deviceTokenRepo.Mock.On("Create", template).Return(nil)

				result, err := x.deviceTokenUsecase.Create(context.Background(), payloadDto)
				x.Nil(err)
				x.Equal(template, result)
			},
		},
		{
			name: "Failed Domain Negatif Case",
			tests: func(arg Any) {
				_, err := x.deviceTokenUsecase.Create(context.Background(), payloadEmptyDto)
				e := &exceptions.CustomError{
					Status: exceptions.ERRDOMAIN,
					Errors: multierror.Append(lang.Trans("filled", "Token")),
				}

				x.Equal(err, e)
			},
		},
		{
			name: "Failed Repository Negatif Case",
			tests: func(arg Any) {
				x.deviceTokenRepo.Mock.On("Create", template).Return(errors.New(mock.Anything))

				_, err := x.deviceTokenUsecase.Create(context.Background(), payloadDto)
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
