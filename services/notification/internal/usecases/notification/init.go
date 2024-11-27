package notification

import (
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/domain/repositories"
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/internal/config"
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/internal/factories"
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/internal/services"
)

type notificationInteractor struct {
	cfg                   *config.NotificationConfig
	notificationRepo      repositories.NotificationRepository
	templateRepo          repositories.TemplateRepository
	deviceTokenRepo       repositories.DeviceTokenRepository
	userGrpcRepo          repositories.UserRepository
	firebaseGoogleService *services.FirebaseGoogleService
	mailgunService        *services.MailgunService
}

func NewNotificationInteractor(deps *factories.Dependencies) *notificationInteractor {

	return &notificationInteractor{
		cfg:                   deps.Config,
		templateRepo:          deps.MongoFactory.TemplateRepo,
		notificationRepo:      deps.MongoFactory.NotificationRepo,
		deviceTokenRepo:       deps.MongoFactory.DeviceTokenRepo,
		userGrpcRepo:          deps.GrpcClientFactory.UserRepo,
		firebaseGoogleService: deps.FirebaseGoogle,
		mailgunService:        deps.Mailgun,
	}
}
