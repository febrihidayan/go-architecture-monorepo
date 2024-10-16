package repositories

import (
	"context"
	"errors"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/mongoqb"
	"github.com/febrihidayan/go-architecture-monorepo/services/storage/domain/entities"
	"github.com/febrihidayan/go-architecture-monorepo/services/storage/internal/repositories/mongo/mappers"
	"github.com/febrihidayan/go-architecture-monorepo/services/storage/internal/repositories/mongo/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CloudRepository struct {
	db *mongoqb.MongoQueryBuilder
}

func NewCloudRepository(db *mongo.Database) CloudRepository {
	return CloudRepository{
		db: mongoqb.NewMongoQueryBuilder(db.Collection(models.Cloud{}.TableName())),
	}
}

func (x *CloudRepository) Create(ctx context.Context, payload *entities.Cloud) error {
	_, err := x.db.InsertOne(ctx, mappers.ToModelCloud(payload))

	if err != nil {
		return err
	}

	return nil
}

func (x *CloudRepository) All(ctx context.Context, params *entities.CloudQueryParams) ([]*entities.Cloud, error) {
	var (
		results []*entities.Cloud
		query   = x.db.NewPipeline()
	)

	query.
		AddConditions(func(builder *mongoqb.MongoQueryBuilder) {
			builder.Match(builder.WhereGroup(func(condition *bson.D) {
				if params.Status != "" {
					builder.Where(condition, "status", "=", params.Status)
				}

				if !params.CreatedAt.IsZero() {
					builder.Where(condition, "status", ">=", primitive.NewDateTimeFromTime(params.CreatedAt.Local()))
				}
			}))
		})

	cursor, err := query.Execute(ctx)
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

	err := x.db.FindOne(ctx, bson.M{"_id": id}).Decode(&item)

	if err != nil {
		return nil, errors.New("cloud not found")
	}

	return mappers.ToDomainCloud(&item), nil
}

func (x *CloudRepository) FindByUrl(ctx context.Context, url string) (*entities.Cloud, error) {
	var item models.Cloud

	err := x.db.FindOne(ctx, bson.M{"url": url}).Decode(&item)
	if err != nil {
		return nil, errors.New("cloud not found")
	}

	return mappers.ToDomainCloud(&item), nil
}

func (x *CloudRepository) Update(ctx context.Context, payload *entities.Cloud) error {
	_, err := x.db.ReplaceOne(ctx, bson.M{
		"_id": payload.ID.String(),
	}, mappers.ToModelCloud(payload))

	if err != nil {
		return err
	}

	return nil
}

func (x *CloudRepository) Delete(ctx context.Context, id string) error {
	_, err := x.db.DeleteOne(ctx, bson.M{
		"_id": id,
	})

	return err
}
