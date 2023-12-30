package usecases

import (
	"context"

	"github.com/febrihidayan/go-architecture-monorepo/pkg/exceptions"
	"github.com/febrihidayan/go-architecture-monorepo/services/notification/domain/entities"
)

type TemplateUsecase interface {
	Create(ctx context.Context, payload entities.TemplateDto) (*entities.Template, *exceptions.CustomError)
}
