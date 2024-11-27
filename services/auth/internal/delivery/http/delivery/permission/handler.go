package permission_handler

import (
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/domain/usecases"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/config"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/factories"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/usecases/permission"

	"github.com/gorilla/mux"
)

type PermissionHttpHandler struct {
	Cfg               *config.AuthConfig
	PermissionUsecase usecases.PermissionUsecase
}

func NewPermissionHttpHandler(
	r *mux.Router,
	deps *factories.Dependencies,
) {
	handler := &PermissionHttpHandler{
		Cfg:               deps.Config,
		PermissionUsecase: permission.NewPermissionInteractor(deps),
	}

	r.HandleFunc("/v1/auth/permissions", handler.GetAll).Methods("GET")
	r.HandleFunc("/v1/auth/permission", handler.Create).Methods("POST")
	r.HandleFunc("/v1/auth/permission/{id}", handler.Find).Methods("GET")
	r.HandleFunc("/v1/auth/permission/{id}", handler.Update).Methods("PUT")
}
