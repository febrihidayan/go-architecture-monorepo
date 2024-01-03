package notification

import (
	"context"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/exceptions"
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/domain/entities"
	"github.com/hashicorp/go-multierror"
)

func (x *notificationInteractor) GetAll(ctx context.Context, params entities.NotificationQueryParams) (*entities.NotificationMeta, *exceptions.CustomError) {
	var (
		multilerr *multierror.Error
		results   = make([]*entities.Notification, 0)
	)

	user, err := x.userGrpcRepo.FindUser(ctx, params.UserId)
	if err != nil {
		multilerr = multierror.Append(multilerr, err)
		return nil, &exceptions.CustomError{
			Status: exceptions.ERRREPOSITORY,
			Errors: multilerr,
		}
	}

	getAll, total, err := x.notificationRepo.GetAll(ctx, &params)
	if err != nil {
		multilerr = multierror.Append(multilerr, err)
		return nil, &exceptions.CustomError{
			Status: exceptions.ERRREPOSITORY,
			Errors: multilerr,
		}
	}

	for _, item := range getAll {
		temp, _ := x.templateRepo.FindByName(ctx, item.Type)

		content := temp.GetTemplateData(item.GetData(), user.LangCode)

		item.SetData(content.Body)

		results = append(results, item)
	}

	return &entities.NotificationMeta{
		Data:    results,
		Total:   total,
		Page:    params.Page,
		PerPage: params.PerPage,
	}, nil
}
