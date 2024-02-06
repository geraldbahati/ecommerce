package usecases

import (
	"context"
	"database/sql"
	"github.com/geraldbahati/ecommerce/pkg/model"
	"github.com/geraldbahati/ecommerce/pkg/repository"
	"github.com/geraldbahati/ecommerce/pkg/utils"
	"github.com/google/uuid"
	"strings"
	"time"
)

type CategoryService struct {
	categoryRepo repository.CategoryRepository
}

func NewCategoryService(categoryRepo repository.CategoryRepository) *CategoryService {
	return &CategoryService{
		categoryRepo: categoryRepo,
	}
}

// CreateCategory creates a new category
func (s *CategoryService) CreateCategory(
	ctx context.Context,
	name string,
	description string,
	imageUrl string,
	seoKeywords string,

) (model.Category, error) {
	// create category model
	descriptionValue := sql.NullString{}
	if description != "" {
		descriptionValue.String = description
		descriptionValue.Valid = true
	}

	imageUrlValue := sql.NullString{}
	if imageUrl != "" {
		imageUrlValue.String = imageUrl
		imageUrlValue.Valid = true
	}

	seoKeywordsValue := sql.NullString{}
	if seoKeywords != "" {
		seoKeywordsValue.String = strings.ToLower(seoKeywords)
		seoKeywordsValue.Valid = true
	}

	category := model.Category{
		ID:          uuid.New(),
		Name:        strings.ToLower(name),
		Description: descriptionValue,
		ImageUrl:    imageUrlValue,
		SeoKeywords: seoKeywordsValue,
		IsActive:    true,
		CreatedAt:   time.Now().UTC(),
	}

	// create category
	newCategory, err := s.categoryRepo.CreateCategory(ctx, category)
	if err != nil {
		return model.Category{}, err
	}

	// return created category
	return newCategory, nil
}

// UpdateCategory updates a category
func (s *CategoryService) UpdateCategory(
	ctx context.Context,
	id uuid.UUID,
	name string,
	description string,
	imageUrl string,
	seoKeywords string,
	isActive bool,
) (model.Category, error) {
	// get category by id
	category, err := s.categoryRepo.GetCategoryById(ctx, id)
	if err != nil {
		return model.Category{}, err
	}

	// update category
	if name != "" {
		category.Name = name
	}

	if description != "" {
		descriptionValue := sql.NullString{String: description, Valid: description != ""}
		category.Description = descriptionValue
	}

	if imageUrl != "" {
		imageUrlValue := sql.NullString{String: imageUrl, Valid: imageUrl != ""}
		category.ImageUrl = imageUrlValue
	}

	if seoKeywords != "" {
		seoKeywordsValue := sql.NullString{String: seoKeywords, Valid: seoKeywords != ""}
		category.SeoKeywords = seoKeywordsValue
	}

	category.IsActive = isActive

	lastUpdated := sql.NullTime{Time: time.Now().UTC(), Valid: true}
	category.LastUpdated = lastUpdated

	// update category
	updatedCategory, err := s.categoryRepo.UpdateCategory(ctx, category)
	if err != nil {
		return model.Category{}, err
	}

	// return updated category
	return updatedCategory, nil
}

// DeleteCategory deletes a category
func (s *CategoryService) DeleteCategory(ctx context.Context, categoryId uuid.UUID) error {
	// delete category
	err := s.categoryRepo.DeleteCategory(ctx, categoryId)
	if err != nil {
		return err
	}

	// return nil
	return nil
}

// GetCategoryById gets a category by id
func (s *CategoryService) GetCategoryById(ctx context.Context, categoryId uuid.UUID) (model.Category, error) {
	// get category by id
	category, err := s.categoryRepo.GetCategoryById(ctx, categoryId)
	if err != nil {
		return model.Category{}, err
	}

	// return category
	return category, nil
}

// GetAllCategories gets all categories
func (s *CategoryService) GetAllCategories(ctx context.Context, pageSize int32, page int32) (model.PaginationResult, error) {
	// get category count
	totalCount, err := s.categoryRepo.GetCategoryCount(ctx)
	if err != nil {
		return model.PaginationResult{}, err
	}

	// get all categories
	paginatedCategories, err := utils.Paginate(ctx, totalCount, page, pageSize, func(offset int32, limit int32) (interface{}, error) {
		return s.categoryRepo.GetAllCategories(ctx, offset, limit)
	})
	if err != nil {
		return model.PaginationResult{}, err
	}

	// return categories
	return *paginatedCategories, nil
}

// SearchCategoriesByName searches categories by name
func (s *CategoryService) SearchCategoriesByName(ctx context.Context, name string, pageSize int32, page int32) (model.PaginationResult, error) {
	// get category count
	totalCount, err := s.categoryRepo.GetCategoryCount(ctx)
	if err != nil {
		return model.PaginationResult{}, err
	}

	// add wildcard to name
	name = "%" + name + "%"

	// search categories by name
	paginatedCategories, err := utils.Paginate(ctx, totalCount, page, pageSize, func(offset int32, limit int32) (interface{}, error) {
		return s.categoryRepo.SoftSearchCategoriesByName(ctx, name, offset, limit)
	})
	if err != nil {
		return model.PaginationResult{}, err
	}

	// return categories
	return *paginatedCategories, nil
}
