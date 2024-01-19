package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/config"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/delivery/grpc_client"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/delivery/grpc_server"
	acl_handler "github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/delivery/http/delivery/acl"
	auth_handler "github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/delivery/http/delivery/auth"
	permision_handler "github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/delivery/http/delivery/permission"
	role_handler "github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/delivery/http/delivery/role"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/repositories/factories"
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
	roleUserRepo   = repository_mongo.NewRoleUserRepository(db)
	mongoFactory   = factories.NewMongoFactory(db)
)

func main() {
	defer func() {
		db.Client().Disconnect(ctx)
	}()

	// run rpc client
	grpcClient, errs := grpc_client.NewGrpcClient(&cfg.GrpcClient)
	if len(errs) > 0 {
		cancel()
		log.Fatalf("did not connect grpc client: %v", errs)
	}

	// run Grpc Server
	go RunGrpcServer()
	// end run Grpc Server

	router := mux.NewRouter()
	initHandler(router, cfg, grpcClient)
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

func RunGrpcServer() {

	grpcServer := grpc.NewServer()
	grpc_server.HandlerAuthServices(grpcServer, db, *cfg)

	lis, err := net.Listen("tcp", cfg.RpcPort)
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	go func() {
		log.Println(fmt.Sprintf("Grpc Server listen to: %v", cfg.RpcPort))
		log.Fatal(grpcServer.Serve(lis))
	}()
}

func initHandler(
	router *mux.Router,
	cfg *config.AuthConfig,
	grpcClient *grpc_client.ServerClient) {

	grpcClientFactory := factories.NewGrpcFactory(grpcClient)

	auth_handler.AuthHttpHandler(router, cfg, mongoFactory, grpcClientFactory)
	permision_handler.PermissionHttpHandler(router, cfg, mongoFactory)
	role_handler.RoleHttpHandler(router, cfg, mongoFactory)
	acl_handler.AclHttpHandler(router, cfg, mongoFactory)
}
