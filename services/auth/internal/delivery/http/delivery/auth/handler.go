package auth_handler

import (
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/domain/usecases"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/config"
	repository_mongo "github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/repositories/mongo"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/usecases/auth"

	"github.com/gorilla/mux"
)

type authHttpHandler struct {
	cfg         *config.AuthConfig
	authUsecase usecases.AuthUsecase
}

func AuthHttpHandler(
	r *mux.Router,
	config *config.AuthConfig,
	authRepo repository_mongo.AuthRepository,
) {
	handler := &authHttpHandler{
		cfg: config,
		authUsecase: auth.NewAuthInteractor(
			config,
			&authRepo,
		),
	}

	r.HandleFunc("/v1/auth/login", handler.Login).Methods("POST")
	r.HandleFunc("/v1/auth/register", handler.Register).Methods("POST")
}
