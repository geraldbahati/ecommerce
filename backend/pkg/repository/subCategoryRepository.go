package repository

import (
	"context"
	"github.com/geraldbahati/ecommerce/pkg/model"
	"github.com/google/uuid"
)

type SubCategoryRepository interface {
	// create
	CreateSubCategory(ctx context.Context, subCategory model.AddSubCategoryParams) (model.SubCategory, error)

	// update
	UpdateSubCategory(ctx context.Context, subCategory model.SubCategory) (model.SubCategory, error)

	// delete
	DeleteSubCategory(ctx context.Context, subCategoryId uuid.UUID) error

	// get
	GetProductBySubCategory(ctx context.Context, subCategoryId uuid.NullUUID, offset int32, limit int32) (interface{}, error)
	ListSubCategories(ctx context.Context) ([]model.SubCategory, error)
	GetProductCountBySubCategory(ctx context.Context, subCategoryId uuid.NullUUID) (int64, error)
	GetSubCategoryByCategory(ctx context.Context, categoryId uuid.UUID, offset int32, limit int32) (interface{}, error)
	GetSubCategoryCountByCategory(ctx context.Context, categoryId uuid.UUID) (int64, error)
}
