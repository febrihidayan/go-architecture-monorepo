package profile

import (
	"github.com/febrihidayan/go-architecture-monorepo/services/user/domain/repositories"
	"github.com/febrihidayan/go-architecture-monorepo/services/user/internal/config"
	"github.com/febrihidayan/go-architecture-monorepo/services/user/internal/repositories/factories"
)

type profileInteractor struct {
	cfg             *config.UserConfig
	userRepo        repositories.UserRepository
	storageGrpcRepo repositories.StorageRepository
}

func NewProfileInteractor(
	config *config.UserConfig,
	mongoFactory *factories.MongoFactory,
	grpcFactory *factories.GrpcClientFactory,
) *profileInteractor {

	return &profileInteractor{
		cfg:             config,
		userRepo:        mongoFactory.UserRepo,
		storageGrpcRepo: grpcFactory.StorageRepo,
	}
}
