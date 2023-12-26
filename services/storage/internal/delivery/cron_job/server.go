package cron_job

import "github.com/febrihidayan/go-architecture-monorepo/services/storage/domain/usecases"

type CronJob struct {
	cloudUsecase usecases.CloudUsecase
}
