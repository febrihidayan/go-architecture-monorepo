package repositories

import (
	"context"

	"github.com/febrihidayan/go-architecture-monorepo/services/notification/domain/entities"
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/internal/repositories/mongo/mappers"
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/internal/repositories/mongo/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type DeviceTokenRepository struct {
	db *mongo.Database
}

func NewDeviceTokenRepository(db *mongo.Database) DeviceTokenRepository {
	return DeviceTokenRepository{db: db}
}

func (x *DeviceTokenRepository) Create(ctx context.Context, payload *entities.DeviceToken) error {
	_, err := x.db.Collection(models.DeviceToken{}.TableName()).InsertOne(ctx, mappers.ToModelDeviceToken(payload))

	if err != nil {
		return err
	}

	return nil
}

func (x *DeviceTokenRepository) All(ctx context.Context, params *entities.DeviceTokenQueryParams) ([]*entities.DeviceToken, error) {
	var (
		results []*entities.DeviceToken
		filter  = mongo.Pipeline{}
		match   bson.D
	)

	if params.UserId != "" {
		match = append(match, bson.D{{"user_id", params.UserId}}...)
	}

	if len(match) > 0 {
		filter = append(filter, mongo.Pipeline{
			bson.D{{
				"$match", match,
			}},
		}...)
	}

	cursor, err := x.db.Collection(models.DeviceToken{}.TableName()).Aggregate(ctx, filter)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var item models.DeviceToken
		if err := cursor.Decode(&item); err != nil {
			return nil, err
		}

		results = append(results, mappers.ToDomainDeviceToken(&item))
	}

	return results, nil
}

func (x *DeviceTokenRepository) Delete(ctx context.Context, id string) error {
	_, err := x.db.Collection(models.DeviceToken{}.TableName()).DeleteOne(ctx, bson.M{
		"_id": id,
	})

	return err
}
