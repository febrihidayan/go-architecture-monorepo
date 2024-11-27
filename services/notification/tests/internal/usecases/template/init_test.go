package template

import (
	"testing"
	"time"

	"bou.ke/monkey"
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/domain/usecases"
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/internal/config"
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/internal/factories"
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/internal/usecases/template"
	mongo_repositories "github.com/febrihidayan/go-architecture-monorepo/services/notification/tests/mocks/repositories/mongo"
	"github.com/stretchr/testify/suite"
)

type TemplateUsecaseSuite struct {
	suite.Suite
	cfg             *config.NotificationConfig
	deps            *factories.Dependencies
	mongoFactory    *factories.MongoFactory
	templateRepo    *mongo_repositories.TemplateRepositoryMock
	templateUsecase usecases.TemplateUsecase
}

func (x *TemplateUsecaseSuite) SetupTest() {
	x.cfg = &config.NotificationConfig{}

	x.templateRepo = new(mongo_repositories.TemplateRepositoryMock)

	x.mongoFactory = &factories.MongoFactory{
		TemplateRepo: x.templateRepo,
	}

	x.deps = &factories.Dependencies{
		Config:       x.cfg,
		MongoFactory: x.mongoFactory,
	}

	x.templateUsecase = template.NewTemplateInteractor(x.deps)

	// fake time now for testing
	monkey.Patch(time.Now, func() time.Time {
		return time.Date(2020, 1, 1, 1, 1, 1, 1, time.UTC)
	})
}

func TestTemplateUsecase(t *testing.T) {
	suite.Run(t, new(TemplateUsecaseSuite))
}

type Any []interface{}

func (a Any) Get(i int) interface{} {
	return a[i]
}
