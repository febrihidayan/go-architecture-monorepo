package repositories

import (
	"context"
	"errors"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/mongoqb"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/domain/entities"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/repositories/mongo/mappers"
	"github.com/febrihidayan/go-architecture-monorepo/services/auth/internal/repositories/mongo/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type RoleRepository struct {
	db *mongoqb.MongoQueryBuilder
}

func NewRoleRepository(db *mongo.Database) RoleRepository {
	return RoleRepository{
		db: mongoqb.NewMongoQueryBuilder(db.Collection(models.Role{}.TableName())),
	}
}

func (x *RoleRepository) Create(ctx context.Context, payload *entities.Role) error {
	_, err := x.db.InsertOne(ctx, mappers.ToModelRole(payload))

	if err != nil {
		return err
	}

	return nil
}

func (x *RoleRepository) Find(ctx context.Context, id string) (*entities.Role, error) {
	var role models.Role

	err := x.db.FindOne(ctx, bson.M{"_id": id}).Decode(&role)

	if err != nil {
		return nil, errors.New("role not found")
	}

	return mappers.ToDomainRole(&role), nil
}

func (x *RoleRepository) FindByName(ctx context.Context, name string) (*entities.Role, error) {
	var role models.Role

	err := x.db.FindOne(ctx, bson.M{"name": name}).Decode(&role)

	if err != nil {
		return nil, errors.New("role not found")
	}

	return mappers.ToDomainRole(&role), nil
}

func (x *RoleRepository) All(ctx context.Context) ([]*entities.Role, error) {
	var roles []*models.Role

	cursor, err := x.db.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	if err := cursor.All(ctx, &roles); err != nil {
		return nil, errors.New("roles not found")
	}

	return mappers.ToListDomainRole(roles), nil
}

func (x *RoleRepository) GetAll(ctx context.Context, params *entities.RoleQueryParams) ([]*entities.Role, int, error) {
	query := x.db.NewPipeline()

	if params.Search != "" {
		query.SearchSingleField("name", params.Search)
	}

	query.
		Sort("created_at", false).
		CountFacet().
		Unwind("$total").
		Paginate(params.Page, params.PerPage)

	cursor, err := query.Execute(ctx)
	if err != nil {
		return nil, 0, err
	}

	defer cursor.Close(ctx)

	var result models.RoleMeta
	for cursor.Next(ctx) {
		if err := cursor.Decode(&result); err != nil {
			return nil, result.Total, err
		}
	}

	return mappers.ToListDomainRole(result.Data), result.Total, nil
}

func (x *RoleRepository) AllByUserId(ctx context.Context, userId string) ([]*entities.Role, error) {
	var (
		query   = x.db.NewPipeline()
		results []*models.Role
	)

	query.
		Lookup("role_user", "_id", "role_id", "role_user").
		Match(bson.D{{
			"role_user.user_id", userId,
		}})

	cursor, err := query.Execute(ctx)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	if err := cursor.All(ctx, &results); err != nil {
		return nil, errors.New("roles not found")
	}

	return mappers.ToListDomainRole(results), nil
}

func (x *RoleRepository) Update(ctx context.Context, payload *entities.Role) error {
	_, err := x.db.ReplaceOne(ctx, bson.M{
		"_id": payload.ID.String(),
	}, mappers.ToModelRole(payload))

	if err != nil {
		return err
	}

	return nil
}
