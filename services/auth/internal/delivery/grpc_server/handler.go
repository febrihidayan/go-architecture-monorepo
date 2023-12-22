package grpc_server

import (
	authPb "github.com/febrihidayan/go-architecture-monorepo/proto/_generated/auth"

	"github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/config"
	repository_mongo "github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/repositories/mongo"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/usecases/auth"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

func HandlerAuthServices(s *grpc.Server, db *mongo.Database, cfg config.AuthConfig) {
	var (
		authRepo     = repository_mongo.NewAuthRepository(db)
		roleRepo     = repository_mongo.NewRoleRepository(db)
		roleUserRepo = repository_mongo.NewRoleUserRepository(db)
	)

	authPb.RegisterAuthServicesServer(s, &server{
		authUsecase: auth.NewAuthInteractor(&cfg, &authRepo, nil, &roleUserRepo, &roleRepo),
	})
}
