package profile_handler

import (
	"github.com/febrihidayan/go-architecture-monorepo/services/user/domain/usecases"
	"github.com/febrihidayan/go-architecture-monorepo/services/user/internal/config"
	repository_mongo "github.com/febrihidayan/go-architecture-monorepo/services/user/internal/repositories/mongo"
	"github.com/febrihidayan/go-architecture-monorepo/services/user/internal/usecases/profile"

	"github.com/gorilla/mux"
)

type profileHttpHandler struct {
	cfg            *config.UserConfig
	profileUsecase usecases.ProfileUsecase
}

func ProfileHttpHandler(
	r *mux.Router,
	config *config.UserConfig,
	userRepo repository_mongo.UserRepository,
) {
	handler := &profileHttpHandler{
		cfg: config,
		profileUsecase: profile.NewProfileInteractor(
			config,
			&userRepo,
		),
	}

	r.HandleFunc("/v1/user/profile", handler.Find).Methods("GET")
	r.HandleFunc("/v1/user/profile", handler.Update).Methods("PUT")
}
