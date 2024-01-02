package notification

import (
	"testing"
	"time"

	"bou.ke/monkey"
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/domain/usecases"
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/internal/config"
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/internal/repositories/factories"
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/internal/usecases/notification"
	mongo_repositories "github.com/febrihidayan/go-architecture-monorepo/services/notification/tests/mocks/repositories/mongo"
	"github.com/stretchr/testify/suite"
)

type NotificationUsecaseSuite struct {
	suite.Suite
	cfg              *config.NotificationConfig
	mongoFactory     *factories.MongoFactory
	templateRepo     *mongo_repositories.TemplateRepositoryMock
	notificationRepo *mongo_repositories.NotificationRepositoryMock
	notificationUsecase  usecases.NotificationUsecase
}

func (x *NotificationUsecaseSuite) SetupTest() {
	x.cfg = &config.NotificationConfig{}

	x.templateRepo = new(mongo_repositories.TemplateRepositoryMock)
	x.notificationRepo = new(mongo_repositories.NotificationRepositoryMock)

	x.mongoFactory = &factories.MongoFactory{
		NotificationRepo: x.notificationRepo,
		TemplateRepo:     x.templateRepo,
	}

	x.notificationUsecase = notification.NewNotificationInteractor(x.cfg, x.mongoFactory)

	// fake time now for testing
	monkey.Patch(time.Now, func() time.Time {
		return time.Date(2020, 1, 1, 1, 1, 1, 1, time.UTC)
	})
}

func TestNotificationUsecase(t *testing.T) {
	suite.Run(t, new(NotificationUsecaseSuite))
}

type Any []interface{}

func (a Any) Get(i int) interface{} {
	return a[i]
}
