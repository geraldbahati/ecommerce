package repository

import (
	"context"
	"github.com/geraldbahati/ecommerce/pkg/model"
	"github.com/google/uuid"
)

type CategoryRepository interface {
	// create
	CreateCategory(ctx context.Context, category model.Category) (model.Category, error)

	// update
	UpdateCategory(ctx context.Context, category model.Category) (model.Category, error)

	// delete
	DeleteCategory(ctx context.Context, categoryId uuid.UUID) error

	// get
	GetCategoryById(ctx context.Context, categoryId uuid.UUID) (model.Category, error)
	GetAllCategories(ctx context.Context, offset int32, limit int32) (interface{}, error)
	SoftSearchCategoriesByName(ctx context.Context, categoryName string, offset int32, limit int32) (interface{}, error)
	GetActiveCategories(ctx context.Context, offset int32, limit int32) (interface{}, error)
	GetInactiveCategories(ctx context.Context, offset int32, limit int32) (interface{}, error)
	GetCategoryCount(ctx context.Context) (int64, error)
}
