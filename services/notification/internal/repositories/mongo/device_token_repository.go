package repositories

import (
	"context"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/mongoqb"
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/domain/entities"
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/internal/repositories/mongo/mappers"
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/internal/repositories/mongo/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type DeviceTokenRepository struct {
	db *mongoqb.MongoQueryBuilder
}

func NewDeviceTokenRepository(db *mongo.Database) DeviceTokenRepository {
	return DeviceTokenRepository{
		db: mongoqb.NewMongoQueryBuilder(db.Collection(models.DeviceToken{}.TableName())),
	}
}

func (x *DeviceTokenRepository) Create(ctx context.Context, payload *entities.DeviceToken) error {
	_, err := x.db.InsertOne(ctx, mappers.ToModelDeviceToken(payload))

	if err != nil {
		return err
	}

	return nil
}

func (x *DeviceTokenRepository) All(ctx context.Context, params *entities.DeviceTokenQueryParams) ([]*entities.DeviceToken, error) {
	var (
		results []*entities.DeviceToken
		query   = x.db.NewPipeline()
	)

	query.AddConditions(func(builder *mongoqb.MongoQueryBuilder) {
		builder.Match(builder.WhereGroup(func(condition *bson.D) {
			if params.UserId != "" {
				builder.Where(condition, "user_id", "=", params.UserId)
			}
		}))
	})

	cursor, err := query.Execute(ctx)
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
	_, err := x.db.DeleteOne(ctx, bson.M{
		"_id": id,
	})

	return err
}
