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

type NotificationRepository struct {
	db *mongoqb.MongoQueryBuilder
}

func NewNotificationRepository(db *mongo.Database) NotificationRepository {
	return NotificationRepository{
		db: mongoqb.NewMongoQueryBuilder(db.Collection(models.Notification{}.TableName())),
	}
}

func (x *NotificationRepository) Create(ctx context.Context, payload *entities.Notification) error {
	_, err := x.db.InsertOne(ctx, mappers.ToModelNotification(payload))

	if err != nil {
		return err
	}

	return nil
}

func (x *NotificationRepository) GetAll(ctx context.Context, params *entities.NotificationQueryParams) ([]*entities.Notification, int, error) {
	query := x.db.NewPipeline()

	query.
		AddConditions(func(builder *mongoqb.MongoQueryBuilder) {
			builder.Match(builder.WhereGroup(func(condition *bson.D) {
				if params.UserId != "" {
					builder.Where(condition, "user_id", "=", params.UserId)
				}
			}))
		}).
		Sort("created_at", false).
		CountFacet().
		Unwind("$total").
		Paginate(params.Page, params.PerPage)

	cursor, err := query.Execute(ctx)
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
	_, err := x.db.DeleteOne(ctx, bson.M{
		"_id": id,
	})

	return err
}
