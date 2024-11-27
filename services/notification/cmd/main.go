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

	"github.com/febrihidayan/go-architecture-monorepo/services/notification/internal/delivery/grpc_server"
	device_token_handler "github.com/febrihidayan/go-architecture-monorepo/services/notification/internal/delivery/http/delivery/device_token"
	notification_handler "github.com/febrihidayan/go-architecture-monorepo/services/notification/internal/delivery/http/delivery/notification"
	template_handler "github.com/febrihidayan/go-architecture-monorepo/services/notification/internal/delivery/http/delivery/template"
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/internal/factories"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

var (
	ctx, cancel = context.WithCancel(context.Background())
)

func main() {
	defer cancel()

	// Initialize dependencies
	deps := factories.InitializeDependencies()
	defer deps.Close()

	// Run HTTP Server
	go RunHTTPServer(deps)

	// Run gRPC Server
	go RunGrpcServer(deps)

	// Handle graceful shutdown
	HandleGracefulShutdown()
}

func RunGrpcServer(deps *factories.Dependencies) {
	grpcServer := grpc.NewServer()
	grpc_server.HandlerNotificationServices(grpcServer, deps)

	lis, err := net.Listen("tcp", deps.Config.RpcPort)
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	go func() {
		log.Println(fmt.Sprintf("Grpc Server listen to: %v", deps.Config.RpcPort))
		log.Fatal(grpcServer.Serve(lis))
	}()
}

func RunHTTPServer(deps *factories.Dependencies) {
	router := mux.NewRouter()

	notification_handler.NewNotificationHttpHandler(router, deps)
	template_handler.NewTemplateHttpHandler(router, deps)
	device_token_handler.NewDeviceTokenHttpHandler(router, deps)

	log.Printf("HTTP Server running on %s", deps.Config.HttpPort)
	if err := http.ListenAndServe(deps.Config.HttpPort, router); err != nil {
		log.Fatalf("HTTP Server stopped: %v", err)
	}
}

func HandleGracefulShutdown() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit
	log.Println("Shutting down gracefully...")
	cancel() // Cancel the context for all components
}
