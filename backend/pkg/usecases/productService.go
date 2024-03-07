package usecases

import (
	"context"
	"database/sql"
	"github.com/geraldbahati/ecommerce/internal/database"
	"github.com/geraldbahati/ecommerce/pkg/model"
	"github.com/geraldbahati/ecommerce/pkg/repository"
	"github.com/geraldbahati/ecommerce/pkg/utils"
	"github.com/google/uuid"
	"log"
)

type ProductService struct {
	productRepo repository.ProductRepository
}

func NewProductService(productRepo repository.ProductRepository) *ProductService {
	return &ProductService{
		productRepo: productRepo,
	}
}

// Get All Products
func (s *ProductService) GetProducts(ctx context.Context, pageSize int32, page int32) (model.PaginationResult, error) {
	// get product count
	productCount, err := s.productRepo.GetProductCount(ctx)
	if err != nil {
		return model.PaginationResult{}, err
	}

	paginatedProducts, err := utils.Paginate(
		ctx,
		productCount,
		page,
		pageSize,
		func(offset, limit int32) (interface{}, error) {
			return s.productRepo.GetProducts(ctx, pageSize, page)
		},
	)
	if err != nil {
		return model.PaginationResult{}, err
	}

	return *paginatedProducts, nil
}

// Get a specific product details
func (s *ProductService) GetProductDetails(ctx context.Context, productID uuid.UUID) (database.Product, error) {
	return s.productRepo.GetProductById(ctx, productID)
}

// AddProduct creates a new product
func (s *ProductService) AddProduct(
	ctx context.Context,
	Name string,
	Description string,
	ImageUrl string,
	Price string,
	Stock int32,
	SubCategoryID string,
	Brand string,
	Keywords string,
) (model.Product, error) {

	descriptionValue := sql.NullString{}
	if Description != "" {
		descriptionValue.String = Description
		descriptionValue.Valid = true
	}

	imageUrlValue := sql.NullString{}
	if ImageUrl != "" {
		imageUrlValue.String = ImageUrl
		imageUrlValue.Valid = true
	}

	brandValue := sql.NullString{}
	if Brand != "" {
		brandValue.String = Brand
		brandValue.Valid = true
	}

	keywordsValue := sql.NullString{}
	if Keywords != "" {
		keywordsValue.String = Keywords
		keywordsValue.Valid = true
	}

	// parse category id to uuid
	categoryUUID, err := uuid.Parse(SubCategoryID)
	categoryIDValue := uuid.NullUUID{
		UUID:  categoryUUID,
		Valid: true,
	}
	if err != nil {
		return model.Product{}, err
	}

	createProuduct := model.AddProductParams{
		Name:          Name,
		Description:   descriptionValue,
		ImageUrl:      imageUrlValue,
		Price:         Price,
		Stock:         Stock,
		SubCategoryID: categoryIDValue,
		Brand:         brandValue,
		Keywords:      keywordsValue,
	}

	// create product
	newProduct, err := s.productRepo.AddProduct(ctx, createProuduct)
	if err != nil {
		return model.Product{}, err
	}

	// return created product
	return newProduct, nil
}

// Update an existing product
func (s *ProductService) UpdateProduct(ctx context.Context,
	ID uuid.UUID,
	Name string,
	Description sql.NullString,
	ImageUrl sql.NullString,
	Price string,
	Stock int32,
	SubCategoryID uuid.NullUUID,
	Brand sql.NullString,
	Rating string,
	ReviewCount int32,
	DiscountRate string,
	Keywords sql.NullString,
	IsActive bool) (database.Product, error) {
	return s.productRepo.UpdateProduct(ctx, database.UpdateProductParams{
		ID:            ID,
		Name:          Name,
		Description:   Description,
		ImageUrl:      ImageUrl,
		Price:         Price,
		Stock:         Stock,
		SubCategoryID: SubCategoryID,
		Brand:         Brand,
		Rating:        Rating,
		ReviewCount:   ReviewCount,
		DiscountRate:  DiscountRate,
		Keywords:      Keywords,
		IsActive:      IsActive,
	})
}

// Delete and existing product
func (s *ProductService) DeleteProduct(ctx context.Context, productID uuid.UUID) error {
	return s.productRepo.DeleteProduct(ctx, productID)
}

// Fetches a particular product
func (s *ProductService) GetProductById(ctx context.Context, id uuid.UUID) (database.Product, error) {
	return s.productRepo.GetProductById(ctx, id)
}

// Fetches available products
func (s *ProductService) GetAvailableProducts(ctx context.Context) ([]database.Product, error) {
	return s.productRepo.GetAvailableProducts(ctx)
}

// Filters products based by category
func (s *ProductService) GetProductsByCategory(ctx context.Context, categoryIdStr string, pageSize int32, page int32) (model.PaginationResult, error) {
	// parse category id to uuid
	categoryID, err := uuid.Parse(categoryIdStr)
	if err != nil {
		log.Fatalf("Failed to parse category id: %v", err)
		return model.PaginationResult{}, err
	}

	// get product count by category
	productCount, err := s.productRepo.GetProductCountByCategory(ctx, categoryID)
	if err != nil {
		return model.PaginationResult{}, err
	}

	paginatedProducts, err := utils.Paginate(
		ctx,
		productCount,
		page,
		pageSize,
		func(offset, limit int32) (interface{}, error) {
			return s.productRepo.GetProductsByCategory(ctx, categoryID)
		},
	)
	if err != nil {
		return model.PaginationResult{}, err
	}

	return *paginatedProducts, nil
}

// Searches for a particular product
func (s *ProductService) SearchProducts(ctx context.Context, query sql.NullString) ([]database.Product, error) {
	return s.productRepo.SearchProducts(ctx, query)
}

// Returns a sales trend for the current month
func (s *ProductService) GetSalesTrends(ctx context.Context) ([]database.GetSalesTrendsRow, error) {
	return s.productRepo.GetSalesTrends(ctx)
}

// GetTrendingProducts Returns trending products
func (s *ProductService) GetTrendingProducts(ctx context.Context) ([]model.TrendingProduct, error) {
	trendingProducts, err := s.productRepo.GetTrendingProducts(ctx)
	if err != nil {
		return nil, err
	}
	return trendingProducts, nil
}
