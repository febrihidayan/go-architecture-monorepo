package profile_handler

import (
	"github.com/febrihidayan/go-architecture-monorepo/services/user/domain/usecases"
	"github.com/febrihidayan/go-architecture-monorepo/services/user/internal/config"
	"github.com/febrihidayan/go-architecture-monorepo/services/user/internal/repositories/factories"
	"github.com/febrihidayan/go-architecture-monorepo/services/user/internal/usecases/profile"

	"github.com/gorilla/mux"
)

type ProfileHttpHandler struct {
	Cfg            *config.UserConfig
	ProfileUsecase usecases.ProfileUsecase
}

func NewProfileHttpHandler(
	r *mux.Router,
	config *config.UserConfig,
	mongoFactory *factories.MongoFactory,
	grpcFactory *factories.GrpcClientFactory,

	) {
	handler := &ProfileHttpHandler{
		Cfg: config,
		ProfileUsecase: profile.NewProfileInteractor(
			config,
			mongoFactory,
			grpcFactory,
		),
	}

	r.HandleFunc("/v1/user/profile", handler.Find).Methods("GET")
	r.HandleFunc("/v1/user/profile", handler.Update).Methods("PUT")
}
