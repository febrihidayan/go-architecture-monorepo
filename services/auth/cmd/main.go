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
	repository_mongo "github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/repositories/mongo"
	"github.com/gorilla/mux"
)

var (
	cfg         = config.Auth()
	ctx, cancel = context.WithCancel(context.Background())
	db          = config.InitDatabaseMongodb()
	authRepo    = repository_mongo.NewAuthRepository(db)
)

func main() {
	defer func() {
		db.Client().Disconnect(ctx)
	}()

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

func initHandler(
	router *mux.Router,
	cfg *config.AuthConfig) {
	auth_handler.AuthHttpHandler(router, cfg, authRepo)
}
