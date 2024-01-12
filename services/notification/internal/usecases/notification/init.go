package notification

import (
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/domain/repositories"
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/internal/config"
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/internal/repositories/factories"
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

func NewNotificationInteractor(
	config *config.NotificationConfig,
	mongoFactory *factories.MongoFactory,
	grpcClientFactory *factories.GrpcClientFactory,
	firebaseGoogleService *services.FirebaseGoogleService,
	mailgunService *services.MailgunService,
) *notificationInteractor {

	return &notificationInteractor{
		cfg:                   config,
		templateRepo:          mongoFactory.TemplateRepo,
		notificationRepo:      mongoFactory.NotificationRepo,
		deviceTokenRepo:       mongoFactory.DeviceTokenRepo,
		userGrpcRepo:          grpcClientFactory.UserRepo,
		firebaseGoogleService: firebaseGoogleService,
		mailgunService:        mailgunService,
	}
}
