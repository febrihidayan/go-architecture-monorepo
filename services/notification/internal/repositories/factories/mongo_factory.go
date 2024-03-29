package factories

import (
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/domain/repositories"
	mongo_repositories "github.com/febrihidayan/go-architecture-monorepo/services/notification/internal/repositories/mongo"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoFactory struct {
	TemplateRepo     repositories.TemplateRepository
	NotificationRepo repositories.NotificationRepository
	DeviceTokenRepo  repositories.DeviceTokenRepository
}

func NewMongoFactory(db *mongo.Database) *MongoFactory {
	var (
		TemplateRepo        = mongo_repositories.NewTemplateRepository(db)
		NotificationRepo    = mongo_repositories.NewNotificationRepository(db)
		DeviceTokenRepoRepo = mongo_repositories.NewDeviceTokenRepository(db)
	)

	return &MongoFactory{
		TemplateRepo:     &TemplateRepo,
		NotificationRepo: &NotificationRepo,
		DeviceTokenRepo:  &DeviceTokenRepoRepo,
	}
}
