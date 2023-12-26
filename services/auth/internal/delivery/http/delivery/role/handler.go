package role_handler

import (
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/domain/usecases"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/config"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/repositories/factories"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/usecases/role"

	"github.com/gorilla/mux"
)

type roleHttpHandler struct {
	cfg         *config.AuthConfig
	roleUsecase usecases.RoleUsecase
}

func RoleHttpHandler(
	r *mux.Router,
	config *config.AuthConfig,
	mongoFactory *factories.MongoFactory,
) {
	handler := &roleHttpHandler{
		cfg: config,
		roleUsecase: role.NewRoleInteractor(
			config,
			mongoFactory,
		),
	}

	r.HandleFunc("/v1/auth/roles", handler.GetAll).Methods("GET")
	r.HandleFunc("/v1/auth/role", handler.Create).Methods("POST")
	r.HandleFunc("/v1/auth/role/{id}", handler.Find).Methods("GET")
	r.HandleFunc("/v1/auth/role/{id}", handler.Update).Methods("PUT")
}
