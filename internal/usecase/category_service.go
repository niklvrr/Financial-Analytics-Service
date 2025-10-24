package usecase

import (
	"context"
	"github.com/niklvrr/Financial-Analytics-Service/internal/domain/model"
)

type CategoryRepo interface {
	CreateCategory(ctx context.Context, c *model.Category) (*model.Category, error)
	GetCategory(ctx context.Context, id int64) (*model.Category, error)
}
