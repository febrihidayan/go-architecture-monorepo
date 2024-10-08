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

	"github.com/febrihidayan/go-architecture-monorepo/pkg/rabbitmq"
	"github.com/febrihidayan/go-architecture-monorepo/services/storage/internal/config"
	"github.com/febrihidayan/go-architecture-monorepo/services/storage/internal/delivery/cron_job"
	"github.com/febrihidayan/go-architecture-monorepo/services/storage/internal/delivery/grpc_server"
	cloud_handler "github.com/febrihidayan/go-architecture-monorepo/services/storage/internal/delivery/http/delivery/cloud"
	"github.com/febrihidayan/go-architecture-monorepo/services/storage/internal/delivery/rabbitmq_server"
	"github.com/febrihidayan/go-architecture-monorepo/services/storage/internal/repositories/factories"
	"github.com/febrihidayan/go-architecture-monorepo/services/storage/internal/services"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

var (
	cfg          = config.Storage()
	ctx, cancel  = context.WithCancel(context.Background())
	db           = config.InitDatabaseMongodb()
	mongoFactory = factories.NewMongoFactory(db)
	awsService   = services.NewAwsService(&cfg.Aws)
)

func main() {
	defer func() {
		db.Client().Disconnect(ctx)
	}()

	// run Grpc Server
	go RunGrpcServer()
	// end run Grpc Server

	// run cron job
	go cron_job.HandlerJobService(cfg, db)

	// Run RabbitMQ
	go RunRabbitMQServer()
	// end Run RabbitMQ

	router := mux.NewRouter()
	initHandler(router, cfg)
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
	grpc_server.HandlerStorageServices(grpcServer, db, *cfg)

	lis, err := net.Listen("tcp", cfg.RpcPort)
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	go func() {
		log.Println(fmt.Sprintf("Grpc Server listen to: %s", cfg.RpcPort))
		log.Fatal(grpcServer.Serve(lis))
	}()
}

func RunRabbitMQServer() {
	dns := fmt.Sprintf(
		"amqp://%s:%s@%s:%s/",
		cfg.RabbitMQ.User,
		cfg.RabbitMQ.Password,
		cfg.RabbitMQ.Host,
		cfg.RabbitMQ.Port,
	)

	rmq, err := rabbitmq.NewRabbitMQ(dns)
	if err != nil {
		log.Fatalln("Failed to connect to RabbitMQ:", err)
	}

	defer rmq.Close()

	server := rabbitmq_server.HandlerRabbitMQServices(cfg, rmq, db)

	server.Worker()
}

func initHandler(
	router *mux.Router,
	cfg *config.StorageConfig) {

	cloud_handler.NewCloudHttpHandler(router, cfg, mongoFactory, *awsService)
}
