package sqlc

import (
	"context"
	"github.com/geraldbahati/ecommerce/internal/database"
	"github.com/geraldbahati/ecommerce/pkg/model"
	"github.com/google/uuid"
)

type SQLCategoryRepository struct {
	DB *database.Queries
}

func NewSQLCategoryRepository(db *database.Queries) *SQLCategoryRepository {
	return &SQLCategoryRepository{
		DB: db,
	}
}

// CreateCategory creates a new category
func (r *SQLCategoryRepository) CreateCategory(ctx context.Context, category model.Category) (model.Category, error) {
	// insert category into database
	createdCategory, err := r.DB.CreateCategory(ctx, database.CreateCategoryParams{
		ID:          category.ID,
		Name:        category.Name,
		Description: category.Description,
		ImageUrl:    category.ImageUrl,
		SeoKeywords: category.SeoKeywords,
		IsActive:    category.IsActive,
		CreatedAt:   category.CreatedAt,
	})
	if err != nil {
		return model.Category{}, err
	}

	// return created category
	return model.Category{
		ID:          createdCategory.ID,
		Name:        createdCategory.Name,
		Description: createdCategory.Description,
		ImageUrl:    createdCategory.ImageUrl,
		SeoKeywords: createdCategory.SeoKeywords,
		IsActive:    createdCategory.IsActive,
		CreatedAt:   createdCategory.CreatedAt,
		LastUpdated: createdCategory.LastUpdated,
	}, nil
}

// UpdateCategory updates a category
func (r *SQLCategoryRepository) UpdateCategory(ctx context.Context, category model.Category) (model.Category, error) {
	// update category in database
	updatedCategory, err := r.DB.UpdateCategory(ctx, database.UpdateCategoryParams{
		ID:          category.ID,
		Name:        category.Name,
		Description: category.Description,
		ImageUrl:    category.ImageUrl,
		SeoKeywords: category.SeoKeywords,
		IsActive:    category.IsActive,
	})
	if err != nil {
		return model.Category{}, err
	}

	// return updated category
	return model.Category{
		ID:          updatedCategory.ID,
		Name:        updatedCategory.Name,
		Description: updatedCategory.Description,
		ImageUrl:    updatedCategory.ImageUrl,
		SeoKeywords: updatedCategory.SeoKeywords,
		IsActive:    updatedCategory.IsActive,
		CreatedAt:   updatedCategory.CreatedAt,
		LastUpdated: updatedCategory.LastUpdated,
	}, nil
}

// DeleteCategory deletes a category
func (r *SQLCategoryRepository) DeleteCategory(ctx context.Context, categoryId uuid.UUID) error {
	// delete category from database
	err := r.DB.DeleteCategory(ctx, categoryId)
	if err != nil {
		return err
	}

	// return nil
	return nil
}

// GetCategoryById gets a category by id
func (r *SQLCategoryRepository) GetCategoryById(ctx context.Context, categoryId uuid.UUID) (model.Category, error) {
	// get category from database
	category, err := r.DB.FindCategoryByID(ctx, categoryId)
	if err != nil {
		return model.Category{}, err
	}

	// return category
	return model.Category{
		ID:          category.ID,
		Name:        category.Name,
		Description: category.Description,
		ImageUrl:    category.ImageUrl,
		SeoKeywords: category.SeoKeywords,
		IsActive:    category.IsActive,
		CreatedAt:   category.CreatedAt,
		LastUpdated: category.LastUpdated,
	}, nil
}

// GetAllCategories gets all categories
func (r *SQLCategoryRepository) GetAllCategories(ctx context.Context, offset int32, limit int32) (interface{}, error) {
	// get all categories from database
	categories, err := r.DB.GetAllCategories(ctx, database.GetAllCategoriesParams{
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return nil, err
	}

	// return categories
	var modelCategories []model.Category
	for _, category := range categories {
		modelCategories = append(modelCategories, model.Category{
			ID:          category.ID,
			Name:        category.Name,
			Description: category.Description,
			ImageUrl:    category.ImageUrl,
			SeoKeywords: category.SeoKeywords,
			IsActive:    category.IsActive,
			CreatedAt:   category.CreatedAt,
			LastUpdated: category.LastUpdated,
		})
	}
	return modelCategories, nil
}

// SoftSearchCategoriesByName soft searches categories by name
func (r *SQLCategoryRepository) SoftSearchCategoriesByName(ctx context.Context, categoryName string, offset int32, limit int32) (interface{}, error) {
	// soft search categories from database
	categories, err := r.DB.FindCategoriesBySoftName(ctx, database.FindCategoriesBySoftNameParams{
		Name:   categoryName,
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return nil, err
	}

	// return categories
	var modelCategories []model.Category
	for _, category := range categories {
		modelCategories = append(modelCategories, model.Category{
			ID:          category.ID,
			Name:        category.Name,
			Description: category.Description,
			ImageUrl:    category.ImageUrl,
			SeoKeywords: category.SeoKeywords,
			IsActive:    category.IsActive,
			CreatedAt:   category.CreatedAt,
			LastUpdated: category.LastUpdated,
		})
	}
	return modelCategories, nil
}

// GetActiveCategories gets active categories
func (r *SQLCategoryRepository) GetActiveCategories(ctx context.Context, offset int32, limit int32) (interface{}, error) {
	// get active categories from database
	categories, err := r.DB.GetActiveCategories(ctx, database.GetActiveCategoriesParams{
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return nil, err
	}

	// return categories
	var modelCategories []model.Category
	for _, category := range categories {
		modelCategories = append(modelCategories, model.Category{
			ID:          category.ID,
			Name:        category.Name,
			Description: category.Description,
			ImageUrl:    category.ImageUrl,
			SeoKeywords: category.SeoKeywords,
			IsActive:    category.IsActive,
			CreatedAt:   category.CreatedAt,
			LastUpdated: category.LastUpdated,
		})
	}
	return modelCategories, nil
}

// GetInactiveCategories gets inactive categories
func (r *SQLCategoryRepository) GetInactiveCategories(ctx context.Context, offset int32, limit int32) (interface{}, error) {
	// get inactive categories from database
	categories, err := r.DB.GetInactiveCategories(ctx, database.GetInactiveCategoriesParams{
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return nil, err
	}

	// return categories
	var modelCategories []model.Category
	for _, category := range categories {
		modelCategories = append(modelCategories, model.Category{
			ID:          category.ID,
			Name:        category.Name,
			Description: category.Description,
			ImageUrl:    category.ImageUrl,
			SeoKeywords: category.SeoKeywords,
			IsActive:    category.IsActive,
			CreatedAt:   category.CreatedAt,
			LastUpdated: category.LastUpdated,
		})
	}
	return modelCategories, nil
}

// GetCategoryCount gets the count of categories
func (r *SQLCategoryRepository) GetCategoryCount(ctx context.Context) (int64, error) {
	// get category count from database
	count, err := r.DB.GetCategoryCount(ctx)
	if err != nil {
		return 0, err
	}

	// return category count
	return count, nil
}
