package sqlc

import (
	"context"
	"database/sql"
	"log"

	"github.com/geraldbahati/ecommerce/internal/database"
	"github.com/geraldbahati/ecommerce/pkg/model"
	"github.com/google/uuid"
)

type SQLProductRepository struct {
	DB *database.Queries
}

// DeleteProduct implements repository.ProductRepository.
func (*SQLProductRepository) DeleteProduct(ctx context.Context, productID uuid.UUID) error {
	panic("unimplemented")
}

// GetAvailableProducts implements repository.ProductRepository.
func (*SQLProductRepository) GetAvailableProducts(ctx context.Context) ([]database.Product, error) {
	panic("unimplemented")
}

// GetFilteredProducts implements repository.ProductRepository.
func (*SQLProductRepository) GetFilteredProducts(ctx context.Context, arg database.GetFilteredProductsParams) ([]database.Product, error) {
	panic("unimplemented")
}

// GetPaginatedProducts implements repository.ProductRepository.
func (*SQLProductRepository) GetPaginatedProducts(ctx context.Context, arg database.GetPaginatedProductsParams) ([]database.Product, error) {
	panic("unimplemented")
}

// GetProductById implements repository.ProductRepository.
func (*SQLProductRepository) GetProductById(ctx context.Context, id uuid.UUID) (database.Product, error) {
	panic("unimplemented")
}

// GetProductDetails implements repository.ProductRepository.
func (*SQLProductRepository) GetProductDetails(ctx context.Context, productID uuid.UUID) (model.ProductDetails, error) {
	panic("unimplemented")
}

// GetProductWithRecommendations implements repository.ProductRepository.
func (*SQLProductRepository) GetProductWithRecommendations(ctx context.Context, id uuid.UUID) (database.GetProductWithRecommendationsRow, error) {
	panic("unimplemented")
}

// GetProducts implements repository.ProductRepository.
func (*SQLProductRepository) GetProducts(ctx context.Context) ([]database.Product, error) {
	panic("unimplemented")
}

// GetProductsByCategory implements repository.ProductRepository.
func (*SQLProductRepository) GetProductsByCategory(ctx context.Context, categoryID uuid.UUID) ([]database.Product, error) {
	panic("unimplemented")
}

// GetSalesTrends implements repository.ProductRepository.
func (*SQLProductRepository) GetSalesTrends(ctx context.Context) ([]database.GetSalesTrendsRow, error) {
	panic("unimplemented")
}

// SearchProducts implements repository.ProductRepository.
func (*SQLProductRepository) SearchProducts(ctx context.Context, query sql.NullString) ([]database.Product, error) {
	panic("unimplemented")
}

func NewSQLProductRepository(db *database.Queries) *SQLProductRepository {
	return &SQLProductRepository{
		DB: db,
	}
}

// Adds a new product to the database
func (r *SQLProductRepository) AddProduct(ctx context.Context, product database.Product) (database.Product, error) {
	// Add product into database
	addProduct, err := r.DB.AddProduct(ctx, database.AddProductParams{
		ID:           product.ID,
		Name:         product.Name,
		Description:  product.Description,
		ImageUrl:     product.ImageUrl,
		Price:        product.Price,
		Stock:        product.Stock,
		CategoryID:   product.CategoryID,
		Brand:        product.Brand,
		Rating:       product.Rating,
		ReviewCount:  product.ReviewCount,
		DiscountRate: product.DiscountRate,
		Keywords:     product.Keywords,
		IsActive:     product.IsActive,
	})

	if err != nil {
		return database.Product{}, err
	}

	// Return newely added product
	return database.Product{
		ID:           addProduct.ID,
		Name:         addProduct.Name,
		Description:  addProduct.Description,
		ImageUrl:     addProduct.ImageUrl,
		Price:        addProduct.Price,
		Stock:        addProduct.Stock,
		CategoryID:   addProduct.CategoryID,
		Brand:        addProduct.Brand,
		Rating:       addProduct.Rating,
		ReviewCount:  addProduct.ReviewCount,
		DiscountRate: addProduct.DiscountRate,
		Keywords:     addProduct.Keywords,
		IsActive:     addProduct.IsActive,
		CreatedAt:    addProduct.CreatedAt,
		LastUpdated:  addProduct.LastUpdated,
	}, err
}

// Updates an already existing product in the database
func (r *SQLProductRepository) UpdateProduct(ctx context.Context, product database.UpdateProductParams) (database.Product, error) {
	updatedProduct, err := r.DB.UpdateProduct(ctx, database.UpdateProductParams{
		ID:           product.ID,
		Name:         product.Name,
		Description:  product.Description,
		ImageUrl:     product.ImageUrl,
		Price:        product.Price,
		Stock:        product.Stock,
		CategoryID:   product.CategoryID,
		Brand:        product.Brand,
		Rating:       product.Rating,
		ReviewCount:  product.ReviewCount,
		DiscountRate: product.DiscountRate,
		Keywords:     product.Keywords,
		IsActive:     product.IsActive,
	})

	if err != nil {
		log.Printf("Error updating product with id %s: %s", product.ID.String(), err.Error())
		return database.Product{}, err
	}

	// Return updated Product
	return database.Product{
		ID:           updatedProduct.ID,
		Name:         updatedProduct.Name,
		Description:  updatedProduct.Description,
		ImageUrl:     updatedProduct.ImageUrl,
		Price:        updatedProduct.Price,
		Stock:        updatedProduct.Stock,
		CategoryID:   updatedProduct.CategoryID,
		Brand:        updatedProduct.Brand,
		Rating:       updatedProduct.Rating,
		ReviewCount:  updatedProduct.ReviewCount,
		DiscountRate: updatedProduct.DiscountRate,
		Keywords:     updatedProduct.Keywords,
		IsActive:     updatedProduct.IsActive,
		CreatedAt:    updatedProduct.CreatedAt,
		LastUpdated:  updatedProduct.LastUpdated,
	}, err
}
