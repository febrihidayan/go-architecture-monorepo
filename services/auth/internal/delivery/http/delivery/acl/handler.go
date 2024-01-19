package acl_handler

import (
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/domain/usecases"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/config"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/repositories/factories"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/usecases/acl"

	"github.com/gorilla/mux"
)

type aclHttpHandler struct {
	cfg        *config.AuthConfig
	aclUsecase usecases.AclUsecase
}

func AclHttpHandler(
	r *mux.Router,
	config *config.AuthConfig,
	mongoFactory *factories.MongoFactory,
) {
	handler := &aclHttpHandler{
		cfg: config,
		aclUsecase: acl.NewAclInteractor(
			config,
			mongoFactory,
		),
	}

	r.HandleFunc("/v1/auth/acl/roles", handler.GetAllRole).Methods("GET")
	r.HandleFunc("/v1/auth/acl/permissions", handler.GetAllPermission).Methods("GET")
	r.HandleFunc("/v1/auth/acl/access", handler.Access).Methods("GET")
	r.HandleFunc("/v1/auth/acl/user/{id}", handler.GetAllUser).Methods("GET")
	r.HandleFunc("/v1/auth/acl/user/{id}", handler.UpdateUser).Methods("PUT")
}
