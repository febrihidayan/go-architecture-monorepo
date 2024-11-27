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

	"github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/delivery/grpc_server"
	acl_handler "github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/delivery/http/delivery/acl"
	auth_handler "github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/delivery/http/delivery/auth"
	permision_handler "github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/delivery/http/delivery/permission"
	role_handler "github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/delivery/http/delivery/role"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/delivery/rabbitmq_server"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/factories"
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

	// Run RabbitMQ Worker
	go RunRabbitMQWorker(deps)

	// Handle graceful shutdown
	HandleGracefulShutdown()
}

func RunGrpcServer(deps *factories.Dependencies) {
	grpcServer := grpc.NewServer()
	grpc_server.HandlerAuthServices(grpcServer, deps)

	lis, err := net.Listen("tcp", deps.Config.RpcPort)
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	go func() {
		log.Println(fmt.Sprintf("Grpc Server listen to: %v", deps.Config.RpcPort))
		log.Fatal(grpcServer.Serve(lis))
	}()
}

func RunRabbitMQWorker(deps *factories.Dependencies) {
	server := rabbitmq_server.HandlerRabbitMQServices(deps, deps.RabbitMQConn)
	server.Worker()
	log.Println("RabbitMQ Worker started")
}

func RunHTTPServer(deps *factories.Dependencies) {
	router := mux.NewRouter()

	auth_handler.NewAuthHttpHandler(router, deps)
	permision_handler.NewPermissionHttpHandler(router, deps)
	role_handler.NewRoleHttpHandler(router, deps)
	acl_handler.NewAclHttpHandler(router, deps)

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
