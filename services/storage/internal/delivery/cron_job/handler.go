package cron_job

import (
	"context"
	"log"
	"time"

	"github.com/febrihidayan/go-architecture-monorepo/services/storage/internal/config"
	"github.com/febrihidayan/go-architecture-monorepo/services/storage/internal/repositories/factories"
	"github.com/febrihidayan/go-architecture-monorepo/services/storage/internal/usecases/cloud"
	"github.com/go-co-op/gocron/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func HandlerJobService(cfg *config.StorageConfig, db *mongo.Database) {
	// create a scheduler
	service, err := gocron.NewScheduler()
	if err != nil {
		log.Panic("start cron job:", err)
	}

	mongoFactory := factories.NewMongoFactory(db)

	jobs := CronJob{
		cloudUsecase: cloud.NewCloudInteractor(cfg, mongoFactory, nil),
	}

	service.NewJob(gocron.DurationJob(24*time.Hour), gocron.NewTask(func() {
		jobs.cloudUsecase.DeleteAllStatusJob(context.Background())
	}))

	// start the scheduler
	service.Start()
}
