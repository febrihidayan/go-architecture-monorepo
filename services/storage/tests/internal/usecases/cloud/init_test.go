package cloud

import (
	"testing"
	"time"

	"bou.ke/monkey"
	"github.com/febrihidayan/go-architecture-monorepo/services/storage/domain/usecases"
	"github.com/febrihidayan/go-architecture-monorepo/services/storage/internal/config"
	"github.com/febrihidayan/go-architecture-monorepo/services/storage/internal/repositories/factories"
	"github.com/febrihidayan/go-architecture-monorepo/services/storage/internal/usecases/cloud"
	mongo_repositories "github.com/febrihidayan/go-architecture-monorepo/services/storage/tests/mocks/repositories/mongo"
	services "github.com/febrihidayan/go-architecture-monorepo/services/storage/tests/mocks/services"
	"github.com/stretchr/testify/suite"
)

type CloudUsecaseSuite struct {
	suite.Suite
	cfg          *config.StorageConfig
	mongoFactory *factories.MongoFactory
	cloudRepo    *mongo_repositories.CloudRepositoryMock
	awsService   *services.AwsServiceMock
	cloudUsecase usecases.CloudUsecase
}

func (x *CloudUsecaseSuite) SetupTest() {
	x.cfg = &config.StorageConfig{}

	x.cloudRepo = new(mongo_repositories.CloudRepositoryMock)
	x.awsService = new(services.AwsServiceMock)

	x.mongoFactory = &factories.MongoFactory{
		CloudRepo: x.cloudRepo,
	}

	x.cloudUsecase = cloud.NewCloudInteractor(x.cfg, x.mongoFactory, x.awsService)

	// fake time now for testing
	monkey.Patch(time.Now, func() time.Time {
		return time.Date(2020, 1, 1, 1, 1, 1, 1, time.UTC)
	})
}

func TestCloudUsecase(t *testing.T) {
	suite.Run(t, new(CloudUsecaseSuite))
}

type Any []interface{}

func (a Any) Get(i int) interface{} {
	return a[i]
}
