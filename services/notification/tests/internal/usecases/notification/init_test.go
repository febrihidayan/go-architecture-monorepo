package notification

import (
	"testing"
	"time"

	"bou.ke/monkey"
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/domain/usecases"
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/internal/config"
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/internal/factories"
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/internal/usecases/notification"
	grpc_repositories "github.com/febrihidayan/go-architecture-monorepo/services/notification/tests/mocks/repositories/grpc"
	mongo_repositories "github.com/febrihidayan/go-architecture-monorepo/services/notification/tests/mocks/repositories/mongo"
	"github.com/stretchr/testify/suite"
)

type NotificationUsecaseSuite struct {
	suite.Suite
	cfg                 *config.NotificationConfig
	deps                *factories.Dependencies
	mongoFactory        *factories.MongoFactory
	grpcClientFactory   *factories.GrpcClientFactory
	templateRepo        *mongo_repositories.TemplateRepositoryMock
	notificationRepo    *mongo_repositories.NotificationRepositoryMock
	userGrpcRepo        *grpc_repositories.UserRepositoryMock
	notificationUsecase usecases.NotificationUsecase
}

func (x *NotificationUsecaseSuite) SetupTest() {
	x.cfg = &config.NotificationConfig{}

	x.templateRepo = new(mongo_repositories.TemplateRepositoryMock)
	x.notificationRepo = new(mongo_repositories.NotificationRepositoryMock)
	x.userGrpcRepo = new(grpc_repositories.UserRepositoryMock)

	x.mongoFactory = &factories.MongoFactory{
		NotificationRepo: x.notificationRepo,
		TemplateRepo:     x.templateRepo,
	}

	x.grpcClientFactory = &factories.GrpcClientFactory{
		UserRepo: x.userGrpcRepo,
	}

	x.deps = &factories.Dependencies{
		Config:            x.cfg,
		MongoFactory:      x.mongoFactory,
		GrpcClientFactory: x.grpcClientFactory,
	}

	x.notificationUsecase = notification.NewNotificationInteractor(x.deps)

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
