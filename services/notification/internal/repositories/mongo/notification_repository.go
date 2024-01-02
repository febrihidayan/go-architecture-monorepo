package repositories

import (
	"context"

	"github.com/febrihidayan/go-architecture-monorepo/services/notification/domain/entities"
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/internal/repositories/mongo/mappers"
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/internal/repositories/mongo/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type NotificationRepository struct {
	db *mongo.Database
}

func NewNotificationRepository(db *mongo.Database) NotificationRepository {
	return NotificationRepository{db: db}
}

func (x *NotificationRepository) Create(ctx context.Context, payload *entities.Notification) error {
	_, err := x.db.Collection(models.Notification{}.TableName()).InsertOne(ctx, mappers.ToModelNotification(payload))

	if err != nil {
		return err
	}

	return nil
}

func (x *NotificationRepository) GetAll(ctx context.Context, params *entities.NotificationQueryParams) ([]*entities.Notification, int, error) {
	var (
		filter = mongo.Pipeline{}
		match  bson.D
		skip   = (params.Page - 1) * params.PerPage
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

	filter = append(filter, mongo.Pipeline{
		bson.D{{
			"$sort", bson.D{
				{"created_at", -1},
			},
		}},
		bson.D{{
			"$facet", bson.D{
				{"total", bson.A{
					bson.D{{
						"$count", "count",
					}},
				}},
				{"data", bson.A{
					bson.D{{
						"$addFields", bson.D{
							{"_id", "$_id"},
						},
					}},
				}},
			},
		}},
		bson.D{{
			"$unwind", "$total",
		}},
		bson.D{{
			"$project", bson.D{
				{"data", bson.D{
					{"$slice", bson.A{
						"$data", skip, bson.D{
							{"$ifNull", bson.A{
								params.PerPage, "$total.count",
							}},
						},
					}},
				}},
				{"page", bson.D{
					{"$literal", skip/params.PerPage + 1},
				}},
				{"per_page", bson.D{
					{"$literal", params.PerPage},
				}},
				{"total", "$total.count"},
			},
		}},
	}...)

	cursor, err := x.db.Collection(models.Notification{}.TableName()).Aggregate(ctx, filter)
	if err != nil {
		return nil, 0, err
	}

	defer cursor.Close(ctx)

	var result models.NotificationMeta
	for cursor.Next(ctx) {
		if err := cursor.Decode(&result); err != nil {
			return nil, result.Total, err
		}
	}

	return mappers.ToListDomainNotification(result.Data), result.Total, nil
}

func (x *NotificationRepository) Delete(ctx context.Context, id string) error {
	_, err := x.db.Collection(models.Notification{}.TableName()).DeleteOne(ctx, bson.M{
		"_id": id,
	})

	return err
}
