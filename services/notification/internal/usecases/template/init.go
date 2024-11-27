package template

import (
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/domain/repositories"
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/internal/config"
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/internal/factories"
)

type templateInteractor struct {
	cfg          *config.NotificationConfig
	templateRepo repositories.TemplateRepository
}

func NewTemplateInteractor(deps *factories.Dependencies) *templateInteractor {

	return &templateInteractor{
		cfg:          deps.Config,
		templateRepo: deps.MongoFactory.TemplateRepo,
	}
}
