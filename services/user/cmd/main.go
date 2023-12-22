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

	"github.com/febrihidayan/go-architecture-monorepo/services/user/internal/config"
	grpc_server "github.com/febrihidayan/go-architecture-monorepo/services/user/internal/delivery/grpc_server"
	profile_handler "github.com/febrihidayan/go-architecture-monorepo/services/user/internal/delivery/http/delivery/profile"
	user_handler "github.com/febrihidayan/go-architecture-monorepo/services/user/internal/delivery/http/delivery/user"
	repository_grpc "github.com/febrihidayan/go-architecture-monorepo/services/user/internal/repositories/grpc"
	repository_mongo "github.com/febrihidayan/go-architecture-monorepo/services/user/internal/repositories/mongo"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

var (
	cfg         = config.User()
	ctx, cancel = context.WithCancel(context.Background())
	db          = config.InitDatabaseMongodb()
	userRepo    = repository_mongo.NewUserRepository(db)
)

func main() {
	defer func() {
		db.Client().Disconnect(ctx)
	}()

	// run Grpc Server
	go RunGrpcServer()
	// end run Grpc Server

	// run grpc client
	conGrpcAuthClient := RunGrpcAuthClient()

	router := mux.NewRouter()
	initHandler(router, cfg, conGrpcAuthClient)
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
	grpc_server.HandlerUserServices(grpcServer, db, *cfg)

	lis, err := net.Listen("tcp", cfg.RpcPort)
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	go func() {
		log.Println(fmt.Sprintf("Grpc Server listen to: %v", cfg.RpcPort))
		log.Fatal(grpcServer.Serve(lis))
	}()
}

func RunGrpcAuthClient() *grpc.ClientConn {
	conGrpcClient, err := grpc.Dial(cfg.GrpcClient.Auth, grpc.WithInsecure())
	if err != nil {
		cancel()
		log.Fatalf("did not connect: %v", err)
	}
	log.Println("rpc user started on", cfg.GrpcClient.Auth)

	return conGrpcClient
}

func initHandler(
	router *mux.Router,
	cfg *config.UserConfig,
	grpcConnAuth *grpc.ClientConn) {

	authRepo := repository_grpc.NewAuthRepository(grpcConnAuth)

	profile_handler.ProfileHttpHandler(router, cfg, userRepo)
	user_handler.UserHttpHandler(router, cfg, userRepo, authRepo)
}
