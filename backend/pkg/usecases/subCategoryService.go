package usecases

import (
	"context"
	"database/sql"
	"github.com/geraldbahati/ecommerce/pkg/model"
	"github.com/geraldbahati/ecommerce/pkg/repository"
	"github.com/geraldbahati/ecommerce/pkg/utils"
	"github.com/google/uuid"
	"log"
)

type SubCategoryService struct {
	subCategoryRepo repository.SubCategoryRepository
}

func NewSubCategoryService(subCategoryRepo repository.SubCategoryRepository) *SubCategoryService {
	return &SubCategoryService{
		subCategoryRepo: subCategoryRepo,
	}
}

// CreateSubCategory creates a new sub category
func (s *SubCategoryService) CreateSubCategory(
	ctx context.Context,
	categoryId string,
	name string,
	description string,
	imageUrl string,
	seoKeywords string,
) (model.SubCategory, error) {
	// create sub category model
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
		seoKeywordsValue.String = seoKeywords
		seoKeywordsValue.Valid = true
	}

	addSubCategoryParams := model.AddSubCategoryParams{
		ID:          uuid.New(),
		CategoryID:  uuid.MustParse(categoryId),
		Name:        name,
		Description: descriptionValue,
		ImageUrl:    imageUrlValue,
		SeoKeywords: seoKeywordsValue,
	}

	// log the sub category creation
	log.Printf("Creating sub category: %v", addSubCategoryParams)

	// create sub category
	return s.subCategoryRepo.CreateSubCategory(ctx, addSubCategoryParams)
}

// GetProductsBySubCategory returns a list of products by sub category
func (s *SubCategoryService) GetProductsBySubCategory(ctx context.Context, subCategoryId string, pageSize int32, page int32) (model.PaginationResult, error) {
	log.Printf("Getting products by sub category: %v", subCategoryId)

	// convert sub category id to uuid
	subCategoryIdUUID, err := uuid.Parse(subCategoryId)
	subCategoryIdUUIDValue := uuid.NullUUID{
		UUID:  subCategoryIdUUID,
		Valid: true,
	}
	if err != nil {
		return model.PaginationResult{}, err
	}
	log.Printf("Sub category id: %v", subCategoryIdUUIDValue)

	// get product count by sub category
	productCount, err := s.subCategoryRepo.GetProductCountBySubCategory(ctx, subCategoryIdUUIDValue)
	if err != nil {
		return model.PaginationResult{}, err
	}

	log.Printf("Product count: %v", productCount)
	// get products by sub category
	paginatedProducts, err := utils.Paginate(
		ctx,
		productCount,
		page,
		pageSize,
		func(offset, limit int32) (interface{}, error) {
			return s.subCategoryRepo.GetProductBySubCategory(ctx, subCategoryIdUUIDValue, offset, limit)
		},
	)
	if err != nil {
		return model.PaginationResult{}, err
	}

	return *paginatedProducts, nil
}

// ListSubCategories returns a list of sub categories
func (s *SubCategoryService) ListSubCategories(ctx context.Context) ([]model.SubCategory, error) {
	// get sub categories
	subCategories, err := s.subCategoryRepo.ListSubCategories(ctx)
	if err != nil {
		return nil, err
	}

	return subCategories, nil
}

// ListSubCategoriesByCategory returns a list of sub categories by category
func (s *SubCategoryService) ListSubCategoriesByCategory(ctx context.Context, categoryId string, pageSize int32, page int32) (model.PaginationResult, error) {
	// convert category id to uuid
	categoryIdUUID, err := uuid.Parse(categoryId)
	if err != nil {
		return model.PaginationResult{}, err
	}

	// get sub category count by category
	subCategoryCount, err := s.subCategoryRepo.GetSubCategoryCountByCategory(ctx, categoryIdUUID)
	if err != nil {
		return model.PaginationResult{}, err
	}

	// get sub categories by category
	paginatedSubCategories, err := utils.Paginate(
		ctx,
		subCategoryCount,
		page,
		pageSize,
		func(offset, limit int32) (interface{}, error) {
			return s.subCategoryRepo.GetSubCategoryByCategory(ctx, categoryIdUUID, offset, limit)
		},
	)
	if err != nil {
		return model.PaginationResult{}, err
	}

	return *paginatedSubCategories, nil
}
