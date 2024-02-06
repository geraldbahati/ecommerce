package usecases

import (
	"context"
	"database/sql"
	"time"

	"github.com/geraldbahati/ecommerce/internal/database"
	"github.com/geraldbahati/ecommerce/pkg/repository"
	"github.com/google/uuid"
)

type ProductService struct{
	productRepo repository.ProductRepository
}

func NewProductService(productRepo repository.ProductRepository) *ProductService{
	return &ProductService{
		productRepo: productRepo,
	}
}

// Get All Products
func (s *ProductService) GetProducts(ctx context.Context)([]database.Product, error){
	return s.productRepo.GetProducts(ctx)
}

// Get a specific product details
func (s *ProductService) GetProductDetails(ctx context.Context, productID uuid.UUID)(database.Product, error){
	return s.productRepo.GetProductById(ctx, productID)
}

// Adds a Product to the database
func(s *ProductService) AddProduct(
	ctx context.Context,
	ID uuid.UUID,
	Name         string,
	Description  sql.NullString,
	ImageUrl     sql.NullString,
	Price        string,
	Stock        int32,
	CategoryID   uuid.UUID,
	Brand        sql.NullString,
	Rating       string,
	ReviewCount  int32,
	DiscountRate string,
	Keywords     sql.NullString,
	IsActive     bool,
	CreatedAt    time.Time,
	LastUpdated  sql.NullTime)(database.Product, error){
	return s.productRepo.AddProduct(ctx, database.Product{
		ID: ID,
		Name: Name,
		Description: Description,
		ImageUrl: ImageUrl,
		Price: Price,
		Stock: Stock,
		CategoryID: CategoryID,
		Brand: Brand,
		Rating: Rating,
		ReviewCount: ReviewCount,
		DiscountRate: DiscountRate,
		Keywords: Keywords,
		IsActive: IsActive,
		CreatedAt: CreatedAt,
		LastUpdated: LastUpdated,
	})
}

// Update an existing product
func (s *ProductService) UpdateProduct(ctx context.Context,
	ID           uuid.UUID,
    Name         string,
    Description  sql.NullString,
    ImageUrl     sql.NullString,
    Price        string,
    Stock        int32,
    CategoryID   uuid.UUID,
    Brand        sql.NullString,
    Rating       string,
    ReviewCount  int32,
    DiscountRate string,
    Keywords     sql.NullString,
    IsActive     bool)(database.Product, error){
	return s.productRepo.UpdateProduct(ctx, database.UpdateProductParams{
		ID: ID,
		Name: Name,
		Description: Description,
		ImageUrl: ImageUrl,
		Price: Price,
		Stock: Stock,
		CategoryID: CategoryID,
		Brand: Brand,
		Rating: Rating,
		ReviewCount: ReviewCount,
		DiscountRate: DiscountRate,
		Keywords: Keywords,
		IsActive: IsActive,
	})
}

// Delete and existing product
func (s *ProductService) DeleteProduct(ctx context.Context, productID uuid.UUID) error{
	return s.productRepo.DeleteProduct(ctx, productID)
}

// Fetches a particular product
func (s *ProductService) GetProductById(ctx context.Context, id uuid.UUID)(database.Product, error){
	return s.productRepo.GetProductById(ctx, id)
}

// Fetches available products
func(s *ProductService) GetAvailableProducts(ctx context.Context)([]database.Product, error){
	return s.productRepo.GetAvailableProducts(ctx)
}

// Fetches products based on filters
func (s *ProductService) GetFilteredProducts(ctx context.Context, CategoryID uuid.UUID, Price string)([]database.Product, error){
		return s.productRepo.GetFilteredProducts(ctx, database.GetFilteredProductsParams{
			CategoryID: CategoryID,
			Price: Price,
		})
}

// Paginates products fetched from database
func (s *ProductService) GetPaginatedProducts(ctx context.Context, Offset int32, Limit  int32)([]database.Product, error){
	return s.productRepo.GetPaginatedProducts(ctx, database.GetPaginatedProductsParams{
		Offset: Offset,
		Limit: Limit,
	})
}

// Filters products based on calculated recommendations
func (s *ProductService) GetProductWithRecommendations(ctx context.Context, id uuid.UUID)(database.GetProductWithRecommendationsRow, error){
	return s.productRepo.GetProductWithRecommendations(ctx, id)
}

// Filters products based by category
func(s *ProductService) GetProductsByCategory(ctx context.Context, categoryID uuid.UUID)([]database.Product, error){
	return s.productRepo.GetProductsByCategory(ctx, categoryID)
}

// Searches for a particular product
func (s *ProductService) SearchProducts(ctx context.Context, query sql.NullString)([]database.Product, error){
	return s.productRepo.SearchProducts(ctx ,query)
}

// Returns a sales trend for the current month
func (s *ProductService) GetSalesTrends(ctx context.Context)([]database.GetSalesTrendsRow, error){
	return s.productRepo.GetSalesTrends(ctx)
}