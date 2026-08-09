package cloud

import (
	"context"
	"reflect"
	"testing"
	"time"

	"bou.ke/monkey"
	"github.com/febrihidayan/go-architecture-monorepo/services/storage/domain/entities"
	"github.com/febrihidayan/go-architecture-monorepo/services/storage/domain/usecases"
	"github.com/febrihidayan/go-architecture-monorepo/services/storage/internal/config"
	"github.com/febrihidayan/go-architecture-monorepo/services/storage/internal/factories"
	awsService "github.com/febrihidayan/go-architecture-monorepo/services/storage/internal/services"
	"github.com/febrihidayan/go-architecture-monorepo/services/storage/internal/usecases/cloud"
	mongo_repositories "github.com/febrihidayan/go-architecture-monorepo/services/storage/tests/mocks/repositories/mongo"
	services "github.com/febrihidayan/go-architecture-monorepo/services/storage/tests/mocks/services"
	"github.com/stretchr/testify/suite"
)

type CloudUsecaseSuite struct {
	suite.Suite
	cfg          *config.StorageConfig
	deps         *factories.Dependencies
	mongoFactory *factories.MongoFactory
	cloudRepo    *mongo_repositories.CloudRepositoryMock
	awsService   *services.AwsServiceMock
	cloudUsecase usecases.CloudUsecase
}

func (x *CloudUsecaseSuite) SetupTest() {
	x.cfg = &config.StorageConfig{
		Aws: config.AwsConfig{
			AccessKey:    "FAKE_ACCESS_KEY_123",
			AccessSecret: "FAKE_ACCESS_SECRET_123",
			Region:       "ap-southeast-1",
			Bucket:       "storage_test",
		},
	}

	x.cloudRepo = new(mongo_repositories.CloudRepositoryMock)
	x.awsService = new(services.AwsServiceMock)

	x.mongoFactory = &factories.MongoFactory{
		CloudRepo: x.cloudRepo,
	}

	awsRealService := awsService.NewAwsService(&x.cfg.Aws)

	x.deps = &factories.Dependencies{
		Config:       x.cfg,
		MongoFactory: x.mongoFactory,
		AwsService:   awsRealService,
	}

	x.cloudUsecase = cloud.NewCloudInteractor(x.deps)

	var s *awsService.AwsService
	monkey.PatchInstanceMethod(reflect.TypeOf(s), "UploadFile", func(s *awsService.AwsService, ctx context.Context, payload *entities.File) (string, error) {
		return x.awsService.UploadFile(ctx, payload)
	})

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
