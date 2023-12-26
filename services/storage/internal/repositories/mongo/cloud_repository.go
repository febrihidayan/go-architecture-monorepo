package repositories

import (
	"context"
	"errors"

	"github.com/febrihidayan/go-architecture-monorepo/services/storage/domain/entities"
	"github.com/febrihidayan/go-architecture-monorepo/services/storage/internal/repositories/mongo/mappers"
	"github.com/febrihidayan/go-architecture-monorepo/services/storage/internal/repositories/mongo/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type CloudRepository struct {
	db *mongo.Database
}

func NewCloudRepository(db *mongo.Database) CloudRepository {
	return CloudRepository{db: db}
}

func (x *CloudRepository) Create(ctx context.Context, payload *entities.Cloud) error {
	_, err := x.db.Collection(models.Cloud{}.TableName()).InsertOne(ctx, mappers.ToModelCloud(payload))

	if err != nil {
		return err
	}

	return nil
}

func (x *CloudRepository) All(ctx context.Context, params *entities.CloudQueryParams) ([]*entities.Cloud, error) {
	var (
		results []*entities.Cloud
		filter  bson.M
	)

	if params.Status != "" {
		filter = bson.M{
			"status": params.Status,
		}
	}

	cursor, err := x.db.Collection(models.Cloud{}.TableName()).Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var item models.Cloud
		if err := cursor.Decode(&item); err != nil {
			return nil, err
		}

		results = append(results, mappers.ToDomainCloud(&item))
	}

	return results, nil
}

func (x *CloudRepository) Find(ctx context.Context, id string) (*entities.Cloud, error) {
	var item models.Cloud

	err := x.db.Collection(models.Cloud{}.TableName()).FindOne(ctx, bson.M{"_id": id}).Decode(&item)

	if err != nil {
		return nil, errors.New("cloud not found")
	}

	return mappers.ToDomainCloud(&item), nil
}

func (x *CloudRepository) Update(ctx context.Context, payload *entities.Cloud) error {
	_, err := x.db.Collection(models.Cloud{}.TableName()).ReplaceOne(ctx, bson.M{
		"_id": payload.ID.String(),
	}, mappers.ToModelCloud(payload))

	if err != nil {
		return err
	}

	return nil
}

func (x *CloudRepository) Delete(ctx context.Context, id string) error {
	_, err := x.db.Collection(models.Cloud{}.TableName()).DeleteOne(ctx, bson.M{
		"_id": id,
	})

	return err
}
