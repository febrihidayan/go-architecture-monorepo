package permission_handler

import (
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/domain/usecases"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/config"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/repositories/factories"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/usecases/permission"

	"github.com/gorilla/mux"
)

type permissionHttpHandler struct {
	cfg               *config.AuthConfig
	permissionUsecase usecases.PermissionUsecase
}

func PermissionHttpHandler(
	r *mux.Router,
	config *config.AuthConfig,
	mongoFactory *factories.MongoFactory,
) {
	handler := &permissionHttpHandler{
		cfg: config,
		permissionUsecase: permission.NewPermissionInteractor(
			config,
			mongoFactory,
		),
	}

	r.HandleFunc("/v1/auth/permissions", handler.GetAll).Methods("GET")
	r.HandleFunc("/v1/auth/permission", handler.Create).Methods("POST")
	r.HandleFunc("/v1/auth/permission/{id}", handler.Find).Methods("GET")
	r.HandleFunc("/v1/auth/permission/{id}", handler.Update).Methods("PUT")
}
