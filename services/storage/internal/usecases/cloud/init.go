package cloud

import (
	"github.com/febrihidayan/go-architecture-monorepo/services/storage/domain/repositories"
	"github.com/febrihidayan/go-architecture-monorepo/services/storage/domain/services"
	"github.com/febrihidayan/go-architecture-monorepo/services/storage/internal/config"
	"github.com/febrihidayan/go-architecture-monorepo/services/storage/internal/factories"
)

type cloudInteractor struct {
	cfg        *config.StorageConfig
	cloudRepo  repositories.CloudRepository
	awsService services.AwsService
}

func NewCloudInteractor(deps *factories.Dependencies) *cloudInteractor {

	return &cloudInteractor{
		cfg:        deps.Config,
		cloudRepo:  deps.MongoFactory.CloudRepo,
		awsService: deps.AwsService,
	}
}
