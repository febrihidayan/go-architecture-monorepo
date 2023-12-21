package profile

import (
	"github.com/febrihidayan/go-architecture-monorepo/services/user/domain/repositories"
	"github.com/febrihidayan/go-architecture-monorepo/services/user/internal/config"
)

type profileInteractor struct {
	cfg      *config.UserConfig
	userRepo repositories.UserRepository
}

func NewProfileInteractor(
	config *config.UserConfig,
	userRepo repositories.UserRepository,
) *profileInteractor {

	return &profileInteractor{
		cfg:      config,
		userRepo: userRepo,
	}
}
