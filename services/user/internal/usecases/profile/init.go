package profile

import (
	"github.com/febrihidayan/go-architecture-monorepo/services/user/domain/repositories"
	"github.com/febrihidayan/go-architecture-monorepo/services/user/internal/config"
	"github.com/febrihidayan/go-architecture-monorepo/services/user/internal/repositories/factories"
)

type profileInteractor struct {
	cfg      *config.UserConfig
	userRepo repositories.UserRepository
}

func NewProfileInteractor(
	config *config.UserConfig,
	mongoFactory *factories.MongoFactory,
) *profileInteractor {

	return &profileInteractor{
		cfg:      config,
		userRepo: mongoFactory.UserRepo,
	}
}
