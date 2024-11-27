package auth_handler

import (
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/domain/usecases"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/config"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/repositories/factories"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/usecases/auth"

	"github.com/gorilla/mux"
)

type AuthHttpHandler struct {
	Cfg         *config.AuthConfig
	AuthUsecase usecases.AuthUsecase
}

func NewAuthHttpHandler(
	r *mux.Router,
	deps *factories.Dependencies,
) {
	handler := &AuthHttpHandler{
		Cfg: deps.Config,
		AuthUsecase: auth.NewAuthInteractor(deps),
	}

	r.HandleFunc("/v1/auth/login", handler.Login).Methods("POST")
	r.HandleFunc("/v1/auth/register", handler.Register).Methods("POST")
	r.HandleFunc("/v1/auth/email/verified", handler.SendEmailVerified).Methods("POST")
	r.HandleFunc("/v1/auth/email/{token}", handler.EmailVerified).Methods("GET")
	r.HandleFunc("/v1/auth/password/email", handler.PasswordEmail).Methods("POST")
	r.HandleFunc("/v1/auth/password/reset", handler.PasswordReset).Methods("POST")
}
