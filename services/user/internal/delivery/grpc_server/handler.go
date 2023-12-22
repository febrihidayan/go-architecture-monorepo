package grpc_server

import (
	userPb "github.com/febrihidayan/go-architecture-monorepo/proto/_generated/user"

	"github.com/febrihidayan/go-architecture-monorepo/services/user/internal/config"
	repositories "github.com/febrihidayan/go-architecture-monorepo/services/user/internal/repositories/mongo"
	"github.com/febrihidayan/go-architecture-monorepo/services/user/internal/usecases/user"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

func HandlerUserServices(s *grpc.Server, db *mongo.Database, cfg config.UserConfig) {
	userRepo := repositories.NewUserRepository(db)

	userPb.RegisterUserServicesServer(s, &server{
		userUsecase: user.NewUserInteractor(&cfg, &userRepo, nil),
	})
}
