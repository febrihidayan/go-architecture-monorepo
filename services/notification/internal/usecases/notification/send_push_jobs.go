package notification

import (
	"context"
	"log"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/exceptions"
	"github.com/febrihidayan/go-architecture-monorepo/pkg/utils"
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/domain/entities"
	"github.com/hashicorp/go-multierror"
)

func (x *notificationInteractor) SendPushJobs(ctx context.Context, params entities.NotificationSends) *exceptions.CustomError {
	var (
		multilerr *multierror.Error
	)

	user, err := x.userGrpcRepo.FindUser(ctx, params.UserId)
	if err != nil {
		log.Println("SendPushJobs::error#1:", err)
		multilerr = multierror.Append(multilerr, err)
		return &exceptions.CustomError{
			Status: exceptions.ERRREPOSITORY,
			Errors: multilerr,
		}
	}

	template, err := x.templateRepo.FindByName(ctx, params.TemplateName)
	if err != nil {
		log.Println("SendPushJobs::error#2:", err)
		multilerr = multierror.Append(multilerr, err)
		return &exceptions.CustomError{
			Status: exceptions.ERRREPOSITORY,
			Errors: multilerr,
		}
	}

	content := template.GetTemplateData(params.GetData(), user.LangCode)

	// insert notification
	if utils.ContainsString(params.Services, entities.NotificationTypeApp) {
		go func(param entities.NotificationSends) {
			notification := entities.NewNotification(entities.NotificationDto{
				UserId: param.UserId,
				Type:   param.TemplateName,
				Data:   param.Data,
			})

			if err := x.notificationRepo.Create(context.Background(), notification); err != nil {
				log.Println("SendPushJobs::error#3:", err)
			}
		}(params)
	}

	// send push fcm
	if utils.ContainsString(params.Services, entities.NotificationTypeFCM) {
		go func(param entities.NotificationSends, data entities.TemplateData) {
			ctx := context.Background()

			devices, err := x.deviceTokenRepo.All(ctx, &entities.DeviceTokenQueryParams{
				UserId: params.UserId,
			})
			if err != nil {
				log.Println("SendPushJobs::error#4:", err)
			}

			for _, item := range devices {
				if err := x.firebaseGoogleService.SendPushMessage(item.OsName, item.Token, data.Title, data.Body); err != nil {
					log.Println("SendPushJobs::error#5:", err)
				}
			}
		}(params, content)
	}

	// send email
	if utils.ContainsString(params.Services, entities.NotificationTypeEmail) {
		// TODO
	}

	return nil
}
