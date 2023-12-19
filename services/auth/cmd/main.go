package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/config"
	auth_handler "github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/delivery/http/delivery/auth"
	permision_handler "github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/delivery/http/delivery/permission"
	role_handler "github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/delivery/http/delivery/role"
	repository_grpc "github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/repositories/grpc"
	repository_mongo "github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/repositories/mongo"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

var (
	cfg            = config.Auth()
	ctx, cancel    = context.WithCancel(context.Background())
	db             = config.InitDatabaseMongodb()
	authRepo       = repository_mongo.NewAuthRepository(db)
	permissionRepo = repository_mongo.NewPermissionRepository(db)
	roleRepo       = repository_mongo.NewRoleRepository(db)
)

func main() {
	defer func() {
		db.Client().Disconnect(ctx)
	}()

	// run rpc user client
	conGrpcUser, errGUser := grpc.Dial(cfg.GrpcClient.User, grpc.WithInsecure())
	if errGUser != nil {
		cancel()
		log.Fatalf("did not connect: %v", errGUser)
	}
	log.Println("rpc user started on", cfg.GrpcClient.User)
	// end run rpc user client

	router := mux.NewRouter()
	initHandler(router, cfg, conGrpcUser)
	http.Handle("/", router)

	log.Println("Http Run on", cfg.HttpPort)
	err := http.ListenAndServe(cfg.HttpPort, router)
	if err != nil {
		log.Fatal(err)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	select {
	case v := <-quit:
		log.Fatal(fmt.Sprintf("signal.Notify: %v", v))
	case done := <-ctx.Done():
		log.Fatal(fmt.Sprintf("ctx.Done: %v", done))
	}

	log.Println("Server Exited Properly")
}

func initHandler(
	router *mux.Router,
	cfg *config.AuthConfig,
	grpcConnUser *grpc.ClientConn) {
	userRepo := repository_grpc.NewUserRepository(grpcConnUser)

	auth_handler.AuthHttpHandler(router, cfg, authRepo, userRepo)
	permision_handler.PermissionHttpHandler(router, cfg, permissionRepo)
	role_handler.RoleHttpHandler(router, cfg, roleRepo)
}
