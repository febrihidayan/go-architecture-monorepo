package notification

import (
	"context"
	"errors"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/common"
	"github.com/febrihidayan/go-architecture-monorepo/pkg/exceptions"
	"github.com/febrihidayan/go-architecture-monorepo/pkg/utils"
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/domain/entities"
	"github.com/hashicorp/go-multierror"
	"github.com/stretchr/testify/mock"
)

func (x *NotificationUsecaseSuite) TestGetAll() {
	id := common.NewID()
	var (
		template      *entities.Template
		notification  *entities.Notification
		notifications []*entities.Notification
		results       *entities.NotificationMeta
		user          *entities.User
	)

	template = &entities.Template{
		ID:        id,
		Name:      "welcome",
		Data:      `{"title":{"en":"Welcome","id":"Selamat Datang {{.name}}"}}`,
		CreatedAt: utils.TimeUTC(),
		UpdatedAt: utils.TimeUTC(),
	}

	notification = &entities.Notification{
		ID:        id,
		Type:      "welcome",
		Data:      `{"name":"Febri Hidayan"}`,
		CreatedAt: utils.TimeUTC(),
		UpdatedAt: utils.TimeUTC(),
	}

	notifications = append(notifications, notification)

	params := entities.NotificationQueryParams{
		UserId:  id.String(),
		Page:    1,
		PerPage: 10,
	}

	results = &entities.NotificationMeta{
		Data:    notifications,
		Total:   len(notifications),
		Page:    1,
		PerPage: 10,
	}

	user = &entities.User{
		ID:       id,
		LangCode: entities.TemplateLangEN,
	}

	args := []struct {
		name  string
		args  Any
		tests func(arg Any)
	}{
		{
			name: "Success Positive Case",
			tests: func(arg Any) {
				x.userGrpcRepo.Mock.On("FindUser", params.UserId).Return(user, nil)

				x.notificationRepo.Mock.On("GetAll", &params).Return(notifications, len(notifications), nil)

				x.templateRepo.Mock.On("FindByName", notification.Type).Return(template, nil)

				result, err := x.notificationUsecase.GetAll(context.Background(), params)
				x.Nil(err)
				x.Equal(result, results)
			},
		},
		{
			name: "Failed Negatif Case",
			tests: func(arg Any) {
				x.userGrpcRepo.Mock.On("FindUser", params.UserId).Return(user, nil)

				x.notificationRepo.Mock.On("GetAll", &params).Return(notifications, len(notifications), errors.New(mock.Anything))

				_, err := x.notificationUsecase.GetAll(context.Background(), params)
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
