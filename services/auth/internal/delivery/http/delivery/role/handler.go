package role_handler

import (
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/domain/usecases"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/config"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/factories"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/usecases/role"

	"github.com/gorilla/mux"
)

type RoleHttpHandler struct {
	Cfg         *config.AuthConfig
	RoleUsecase usecases.RoleUsecase
}

func NewRoleHttpHandler(
	r *mux.Router,
	deps *factories.Dependencies,
) {
	handler := &RoleHttpHandler{
		Cfg:         deps.Config,
		RoleUsecase: role.NewRoleInteractor(deps),
	}

	r.HandleFunc("/v1/auth/roles", handler.GetAll).Methods("GET")
	r.HandleFunc("/v1/auth/role", handler.Create).Methods("POST")
	r.HandleFunc("/v1/auth/role/{id}", handler.Find).Methods("GET")
	r.HandleFunc("/v1/auth/role/{id}", handler.Update).Methods("PUT")
}
