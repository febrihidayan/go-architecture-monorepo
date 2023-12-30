package repositories

import (
	"context"

	"github.com/febrihidayan/go-architecture-monorepo/services/notification/domain/entities"
)

type TemplateRepository interface {
	Create(ctx context.Context, payload *entities.Template) error
	Find(ctx context.Context, id string) (*entities.Template, error)
	FindByName(ctx context.Context, name string) (*entities.Template, error)
	Update(ctx context.Context, payload *entities.Template) error
	Delete(ctx context.Context, id string) error
}
