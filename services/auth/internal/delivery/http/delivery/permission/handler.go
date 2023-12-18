package permission_handler

import (
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/domain/usecases"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/config"
	repository_mongo "github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/repositories/mongo"
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
	permissionRepo repository_mongo.PermissionRepository,
) {
	handler := &permissionHttpHandler{
		cfg: config,
		permissionUsecase: permission.NewPermissionInteractor(
			config,
			&permissionRepo,
		),
	}

	r.HandleFunc("/v1/auth/permission", handler.Create).Methods("POST")
}
