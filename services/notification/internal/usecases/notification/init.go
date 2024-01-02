package notification

import (
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/domain/repositories"
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/internal/config"
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/internal/repositories/factories"
)

type notificationInteractor struct {
	cfg              *config.NotificationConfig
	notificationRepo repositories.NotificationRepository
	templateRepo     repositories.TemplateRepository
}

func NewTemplateInteractor(
	config *config.NotificationConfig,
	mongoFactory *factories.MongoFactory,
) *notificationInteractor {

	return &notificationInteractor{
		cfg:              config,
		templateRepo:     mongoFactory.TemplateRepo,
		notificationRepo: mongoFactory.NotificationRepo,
	}
}
