package template

import (
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/domain/repositories"
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/internal/config"
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/internal/repositories/factories"
)

type templateInteractor struct {
	cfg          *config.NotificationConfig
	templateRepo repositories.TemplateRepository
}

func NewTemplateInteractor(
	config *config.NotificationConfig,
	mongoFactory *factories.MongoFactory,
) *templateInteractor {

	return &templateInteractor{
		cfg:          config,
		templateRepo: mongoFactory.TemplateRepo,
	}
}
