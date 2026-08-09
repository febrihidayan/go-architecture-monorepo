package user

import (
	"reflect"
	"testing"
	"time"

	"bou.ke/monkey"
	"github.com/febrihidayan/go-architecture-monorepo/services/user/domain/usecases"
	"github.com/febrihidayan/go-architecture-monorepo/services/user/internal/config"
	"github.com/febrihidayan/go-architecture-monorepo/services/user/internal/delivery/rabbitmq_server/publisher"
	"github.com/febrihidayan/go-architecture-monorepo/services/user/internal/factories"
	"github.com/febrihidayan/go-architecture-monorepo/services/user/internal/usecases/user"
	grpc_repositories "github.com/febrihidayan/go-architecture-monorepo/services/user/tests/mocks/repositories/grpc"
	mongo_repositories "github.com/febrihidayan/go-architecture-monorepo/services/user/tests/mocks/repositories/mongo"
	rabbitmq_repositories "github.com/febrihidayan/go-architecture-monorepo/services/user/tests/mocks/repositories/rabbitmq"
	"github.com/stretchr/testify/suite"
)

type UserUsecaseSuite struct {
	suite.Suite
	cfg             *config.UserConfig
	mongoFactory    *factories.MongoFactory
	grpcFactory     *factories.GrpcClientFactory
	userRepo        *mongo_repositories.UserRepositoryMock
	authRepo        *grpc_repositories.AuthRepositoryMock
	storageGrpcRepo *grpc_repositories.StorageRepositoryMock
	rabbitmqRepo    *rabbitmq_repositories.RabbitMQRepositoryMock
	userUsecase     usecases.UserUsecase
}

func (x *UserUsecaseSuite) SetupTest() {
	x.cfg = &config.UserConfig{}

	x.userRepo = new(mongo_repositories.UserRepositoryMock)
	x.authRepo = new(grpc_repositories.AuthRepositoryMock)
	x.storageGrpcRepo = new(grpc_repositories.StorageRepositoryMock)

	x.mongoFactory = &factories.MongoFactory{
		UserRepo: x.userRepo,
	}

	x.grpcFactory = &factories.GrpcClientFactory{
		AuthRepo:    x.authRepo,
		StorageRepo: x.storageGrpcRepo,
	}

	x.rabbitmqRepo = new(rabbitmq_repositories.RabbitMQRepositoryMock)

	rmqRealPublisher := &publisher.PublisherRabbitMQ{}

	x.userUsecase = user.NewUserInteractor(x.cfg, x.mongoFactory, x.grpcFactory, rmqRealPublisher)

	var p *publisher.PublisherRabbitMQ
	monkey.PatchInstanceMethod(reflect.TypeOf(p), "AuthDelete", func(p *publisher.PublisherRabbitMQ, id string) error {
		return x.rabbitmqRepo.AuthDelete(id)
	})

	// fake time now for testing
	monkey.Patch(time.Now, func() time.Time {
		return time.Date(2020, 1, 1, 1, 1, 1, 1, time.UTC)
	})
}

func TestUserUsecase(t *testing.T) {
	suite.Run(t, new(UserUsecaseSuite))
}

type Any []interface{}

func (a Any) Get(i int) interface{} {
	return a[i]
}
