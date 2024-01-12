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

	"github.com/febrihidayan/go-architecture-monorepo/services/notification/internal/config"
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/internal/delivery/grpc_client"
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/internal/delivery/grpc_server"
	notification_handler "github.com/febrihidayan/go-architecture-monorepo/services/notification/internal/delivery/http/delivery/notification"
	template_handler "github.com/febrihidayan/go-architecture-monorepo/services/notification/internal/delivery/http/delivery/template"
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/internal/repositories/factories"
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/internal/services"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

var (
	cfg          = config.Notification()
	ctx, cancel  = context.WithCancel(context.Background())
	db           = config.InitDatabaseMongodb()
	mongoFactory = factories.NewMongoFactory(db)
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

	// run firebase google service
	clientService := services.NewServiceHandler(cfg)

	// run Grpc Server
	go RunGrpcServer(grpcClient, clientService)
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

func RunGrpcServer(
	grpcClient *grpc_client.ServerClient,
	clientService *services.ClientService) {

	grpcServer := grpc.NewServer()
	grpc_server.HandlerNotificationServices(grpcServer, db, grpcClient, clientService, *cfg)

	lis, err := net.Listen("tcp", cfg.RpcPort)
	if err != nil {
		cancel()
		log.Fatalln("Failed to listen:", err)
	}

	go func() {
		log.Println(fmt.Sprintf("Grpc Server listen to: %s", cfg.RpcPort))
		log.Fatal(grpcServer.Serve(lis))
	}()
}

func initHandler(
	router *mux.Router,
	cfg *config.NotificationConfig,
	grpcClient *grpc_client.ServerClient) {

	grpcClientFactory := factories.NewGrpcFactory(grpcClient)

	notification_handler.TemplateHttpHandler(router, cfg, mongoFactory, grpcClientFactory)
	template_handler.TemplateHttpHandler(router, cfg, mongoFactory)
}
