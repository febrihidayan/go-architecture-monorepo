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

	"github.com/febrihidayan/go-architecture-monorepo/services/storage/internal/delivery/cron_job"
	"github.com/febrihidayan/go-architecture-monorepo/services/storage/internal/delivery/grpc_server"
	cloud_handler "github.com/febrihidayan/go-architecture-monorepo/services/storage/internal/delivery/http/delivery/cloud"
	"github.com/febrihidayan/go-architecture-monorepo/services/storage/internal/factories"
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

	// run cron job
	go cron_job.HandlerJobService(deps)

	// Run HTTP Server
	go RunHTTPServer(deps)

	// Run gRPC Server
	go RunGrpcServer(deps)

	// Handle graceful shutdown
	HandleGracefulShutdown()
}

func RunGrpcServer(deps *factories.Dependencies) {
	grpcServer := grpc.NewServer()
	grpc_server.HandlerStorageServices(grpcServer, deps)

	lis, err := net.Listen("tcp", deps.Config.RpcPort)
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	go func() {
		log.Println(fmt.Sprintf("Grpc Server listen to: %s", deps.Config.RpcPort))
		log.Fatal(grpcServer.Serve(lis))
	}()
}

func RunHTTPServer(deps *factories.Dependencies) {
	router := mux.NewRouter()

	cloud_handler.NewCloudHttpHandler(router, deps)

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
