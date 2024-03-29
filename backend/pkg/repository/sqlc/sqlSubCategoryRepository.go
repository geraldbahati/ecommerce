package sqlc

import (
	"context"
	"github.com/geraldbahati/ecommerce/internal/database"
	"github.com/geraldbahati/ecommerce/pkg/model"
	"github.com/google/uuid"
	"log"
)

type SQLSubCategoryRepository struct {
	DB *database.Queries
}

func NewSQLSubCategoryRepository(db *database.Queries) *SQLSubCategoryRepository {
	return &SQLSubCategoryRepository{
		DB: db,
	}
}

// CreateSubCategory creates a new subcategory in the database
func (r *SQLSubCategoryRepository) CreateSubCategory(ctx context.Context, subCategory model.AddSubCategoryParams) (model.SubCategory, error) {
	// Add subcategory into database
	addSubCategory, err := r.DB.CreateSubCategory(ctx, database.CreateSubCategoryParams{
		ID:          subCategory.ID,
		CategoryID:  subCategory.CategoryID,
		Name:        subCategory.Name,
		Description: subCategory.Description,
		ImageUrl:    subCategory.ImageUrl,
		SeoKeywords: subCategory.SeoKeywords,
	})
	if err != nil {
		log.Printf("Error adding subcategory: %v", err)
		return model.SubCategory{}, err
	}

	// Return newly added subcategory
	return model.SubCategory{
		ID:          addSubCategory.ID,
		CategoryID:  addSubCategory.CategoryID,
		Name:        addSubCategory.Name,
		Description: addSubCategory.Description,
		ImageUrl:    addSubCategory.ImageUrl,
		SeoKeywords: addSubCategory.SeoKeywords,
		IsActive:    addSubCategory.IsActive,
		CreatedAt:   addSubCategory.CreatedAt,
		LastUpdated: addSubCategory.LastUpdated,
	}, err
}

// UpdateSubCategory updates a subcategory in the database
func (r *SQLSubCategoryRepository) UpdateSubCategory(ctx context.Context, subCategory model.SubCategory) (model.SubCategory, error) {
	// Update subcategory in the database
	updateSubCategory, err := r.DB.UpdateSubCategory(ctx, database.UpdateSubCategoryParams{
		ID:          subCategory.ID,
		CategoryID:  subCategory.CategoryID,
		Name:        subCategory.Name,
		Description: subCategory.Description,
		ImageUrl:    subCategory.ImageUrl,
		SeoKeywords: subCategory.SeoKeywords,
	})
	if err != nil {
		return model.SubCategory{}, err
	}

	// Return newly updated subcategory
	return model.SubCategory{
		ID:          updateSubCategory.ID,
		CategoryID:  updateSubCategory.CategoryID,
		Name:        updateSubCategory.Name,
		Description: updateSubCategory.Description,
		ImageUrl:    updateSubCategory.ImageUrl,
		SeoKeywords: updateSubCategory.SeoKeywords,
		IsActive:    updateSubCategory.IsActive,
		CreatedAt:   updateSubCategory.CreatedAt,
		LastUpdated: updateSubCategory.LastUpdated,
	}, err
}

// DeleteSubCategory deletes a subcategory from the database
func (r *SQLSubCategoryRepository) DeleteSubCategory(ctx context.Context, subCategoryId uuid.UUID) error {
	// Delete subcategory from the database
	err := r.DB.DeleteSubCategory(ctx, subCategoryId)
	if err != nil {
		return err
	}

	return nil
}

// GetProductBySubCategory returns a list of products in a subcategory
func (r *SQLSubCategoryRepository) GetProductBySubCategory(ctx context.Context, subCategoryId uuid.NullUUID, offset int32, limit int32) (interface{}, error) {
	// Get products in a subcategory from the database
	products, err := r.DB.GetProductBySubCategory(ctx, database.GetProductBySubCategoryParams{
		SubCategoryID: subCategoryId,
		Offset:        offset,
		Limit:         limit,
	})
	if err != nil {
		log.Fatalf("Error fetching products: %v", err)
		return nil, err
	}

	// Return products
	productList := make([]model.Product, len(products))

	for i, product := range products {
		productList[i] = model.Product{
			ID:            product.ID,
			Name:          product.Name,
			Description:   product.Description,
			ImageUrl:      product.ImageUrl,
			Price:         product.Price,
			Stock:         product.Stock,
			SubCategoryID: product.SubCategoryID,
			Brand:         product.Brand,
			Rating:        product.Rating,
			ReviewCount:   product.ReviewCount,
			DiscountRate:  product.DiscountRate,
			Keywords:      product.Keywords,
			IsActive:      product.IsActive,
			CreatedAt:     product.CreatedAt,
			LastUpdated:   product.LastUpdated,
		}
	}

	return productList, nil
}

// ListSubCategories returns a list of subcategories in a category
func (r *SQLSubCategoryRepository) ListSubCategories(ctx context.Context) ([]model.SubCategory, error) {
	// Get subcategories in a category from the database
	subCategories, err := r.DB.ListAllSubCategories(ctx)
	if err != nil {
		log.Fatalf("Error fetching subcategories: %v", err)
		return nil, err
	}

	// Return subcategories
	subCategoryList := make([]model.SubCategory, len(subCategories))

	for i, subCategory := range subCategories {
		subCategoryList[i] = model.SubCategory{
			ID:          subCategory.ID,
			CategoryID:  subCategory.CategoryID,
			Name:        subCategory.Name,
			Description: subCategory.Description,
			ImageUrl:    subCategory.ImageUrl,
			SeoKeywords: subCategory.SeoKeywords,
			IsActive:    subCategory.IsActive,
			CreatedAt:   subCategory.CreatedAt,
			LastUpdated: subCategory.LastUpdated,
		}
	}

	return subCategoryList, nil
}

// GetProductCountBySubCategory returns the number of products in a subcategory
func (r *SQLSubCategoryRepository) GetProductCountBySubCategory(ctx context.Context, subCategoryId uuid.NullUUID) (int64, error) {
	// Get product count in a subcategory from the database
	productCount, err := r.DB.GetProductCountBySubCategory(ctx, subCategoryId)
	if err != nil {
		log.Printf("Error fetching product count: %v", err)
		return 0, err
	}

	// Return product count
	return productCount, nil
}

// GetSubCategoryByCategory returns a list of subcategories in a category
func (r *SQLSubCategoryRepository) GetSubCategoryByCategory(ctx context.Context, categoryId uuid.UUID, offset int32, limit int32) (interface{}, error) {
	// Get subcategories in a category from the database
	subCategories, err := r.DB.GetSubCategoryByCategory(ctx, database.GetSubCategoryByCategoryParams{
		CategoryID: categoryId,
		Offset:     offset,
		Limit:      limit,
	})
	if err != nil {
		log.Fatalf("Error fetching subcategories: %v", err)
		return nil, err
	}

	// Return subcategories
	subCategoryList := make([]model.SubCategory, len(subCategories))

	for i, subCategory := range subCategories {
		subCategoryList[i] = model.SubCategory{
			ID:          subCategory.ID,
			CategoryID:  subCategory.CategoryID,
			Name:        subCategory.Name,
			Description: subCategory.Description,
			ImageUrl:    subCategory.ImageUrl,
			SeoKeywords: subCategory.SeoKeywords,
			IsActive:    subCategory.IsActive,
			CreatedAt:   subCategory.CreatedAt,
			LastUpdated: subCategory.LastUpdated,
		}
	}

	return subCategoryList, nil
}

// GetSubCategoryCountByCategory returns the number of subcategories in a category
func (r *SQLSubCategoryRepository) GetSubCategoryCountByCategory(ctx context.Context, categoryId uuid.UUID) (int64, error) {
	// Get subcategory count in a category from the database
	subCategoryCount, err := r.DB.GetSubCategoryCountByCategory(ctx, categoryId)
	if err != nil {
		log.Printf("Error fetching subcategory count: %v", err)
		return 0, err
	}

	// Return subcategory count
	return subCategoryCount, nil
}
