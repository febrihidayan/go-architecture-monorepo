package cloud

import (
	"github.com/febrihidayan/go-architecture-monorepo/services/storage/domain/repositories"
	"github.com/febrihidayan/go-architecture-monorepo/services/storage/domain/services"
	"github.com/febrihidayan/go-architecture-monorepo/services/storage/internal/config"
	"github.com/febrihidayan/go-architecture-monorepo/services/storage/internal/repositories/factories"
)

type cloudInteractor struct {
	cfg        *config.StorageConfig
	cloudRepo  repositories.CloudRepository
	awsService services.AwsService
}

func NewCloudInteractor(
	config *config.StorageConfig,
	mongoFactory *factories.MongoFactory,
	awsService services.AwsService,
) *cloudInteractor {

	return &cloudInteractor{
		cfg:        config,
		cloudRepo:  mongoFactory.CloudRepo,
		awsService: awsService,
	}
}
