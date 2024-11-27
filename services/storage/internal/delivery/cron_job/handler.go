package cron_job

import (
	"context"
	"log"
	"time"

	"github.com/febrihidayan/go-architecture-monorepo/services/storage/internal/factories"
	"github.com/febrihidayan/go-architecture-monorepo/services/storage/internal/usecases/cloud"
	"github.com/go-co-op/gocron/v2"
)

func HandlerJobService(deps *factories.Dependencies) {
	// create a scheduler
	service, err := gocron.NewScheduler()
	if err != nil {
		log.Panic("start cron job:", err)
	}

	jobs := CronJob{
		cloudUsecase: cloud.NewCloudInteractor(deps),
	}

	service.NewJob(gocron.DurationJob(24*time.Hour), gocron.NewTask(func() {
		jobs.cloudUsecase.DeleteAllStatusJob(context.Background())
	}))

	// start the scheduler
	service.Start()
}
