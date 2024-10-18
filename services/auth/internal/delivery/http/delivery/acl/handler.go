package acl_handler

import (
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/domain/usecases"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/config"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/repositories/factories"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/usecases/acl"

	"github.com/gorilla/mux"
)

type AclHttpHandler struct {
	Cfg        *config.AuthConfig
	AclUsecase usecases.AclUsecase
}

func NewAclHttpHandler(
	r *mux.Router,
	config *config.AuthConfig,
	mongoFactory *factories.MongoFactory,
) {
	handler := &AclHttpHandler{
		Cfg: config,
		AclUsecase: acl.NewAclInteractor(
			config,
			mongoFactory,
		),
	}

	r.HandleFunc("/v1/auth/acl/roles", handler.GetAllRole).Methods("GET")
	r.HandleFunc("/v1/auth/acl/permissions", handler.GetAllPermission).Methods("GET")
	r.HandleFunc("/v1/auth/acl/permission/role/{id}", handler.GetAllPermissionByRole).Methods("GET")
	r.HandleFunc("/v1/auth/acl/permission/role/{id}", handler.UpdatePermissionByRole).Methods("PUT")
	r.HandleFunc("/v1/auth/acl/access", handler.AccessUserLogin).Methods("GET")
	r.HandleFunc("/v1/auth/acl/user/{id}", handler.GetAllUser).Methods("GET")
	r.HandleFunc("/v1/auth/acl/user/{id}", handler.UpdateUser).Methods("PUT")
}
