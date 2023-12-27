package user

import (
	"github.com/febrihidayan/go-architecture-monorepo/services/user/domain/repositories"
	"github.com/febrihidayan/go-architecture-monorepo/services/user/internal/config"
	"github.com/febrihidayan/go-architecture-monorepo/services/user/internal/repositories/factories"
)

type userInteractor struct {
	cfg             *config.UserConfig
	userRepo        repositories.UserRepository
	authGrpcRepo    repositories.AuthRepository
	storageGrpcRepo repositories.StorageRepository
}

func NewUserInteractor(
	config *config.UserConfig,
	mongoFactory *factories.MongoFactory,
	grpcFactory *factories.GrpcClientFactory,
) *userInteractor {

	return &userInteractor{
		cfg:             config,
		userRepo:        mongoFactory.UserRepo,
		authGrpcRepo:    grpcFactory.AuthRepo,
		storageGrpcRepo: grpcFactory.StorageRepo,
	}
}
