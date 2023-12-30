package repositories

import (
	"context"
	"errors"

	"github.com/febrihidayan/go-architecture-monorepo/services/notification/domain/entities"
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/internal/repositories/mongo/mappers"
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/internal/repositories/mongo/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type TemplateRepository struct {
	db *mongo.Database
}

func NewTemplateRepository(db *mongo.Database) TemplateRepository {
	return TemplateRepository{db: db}
}

func (x *TemplateRepository) Create(ctx context.Context, payload *entities.Template) error {
	_, err := x.db.Collection(models.Template{}.TableName()).InsertOne(ctx, mappers.ToModelTemplate(payload))

	if err != nil {
		return err
	}

	return nil
}

func (x *TemplateRepository) Find(ctx context.Context, id string) (*entities.Template, error) {
	var item models.Template

	err := x.db.Collection(models.Template{}.TableName()).FindOne(ctx, bson.M{"_id": id}).Decode(&item)

	if err != nil {
		return nil, errors.New("template not found")
	}

	return mappers.ToDomainTemplate(&item), nil
}

func (x *TemplateRepository) FindByName(ctx context.Context, name string) (*entities.Template, error) {
	var item models.Template

	err := x.db.Collection(models.Template{}.TableName()).FindOne(ctx, bson.M{"name": name}).Decode(&item)

	if err != nil {
		return nil, errors.New("template not found")
	}

	return mappers.ToDomainTemplate(&item), nil
}

func (x *TemplateRepository) Update(ctx context.Context, payload *entities.Template) error {
	_, err := x.db.Collection(models.Template{}.TableName()).ReplaceOne(ctx, bson.M{
		"_id": payload.ID.String(),
	}, mappers.ToModelTemplate(payload))

	if err != nil {
		return err
	}

	return nil
}

func (x *TemplateRepository) Delete(ctx context.Context, id string) error {
	_, err := x.db.Collection(models.Template{}.TableName()).DeleteOne(ctx, bson.M{
		"_id": id,
	})

	return err
}
