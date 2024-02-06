package sqlc

import (
	"context"
	"database/sql"
	"log"

	"github.com/geraldbahati/ecommerce/internal/database"
	"github.com/google/uuid"
)

type SQLProductRepository struct {
	DB *database.Queries
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


// DeleteProduct implements repository.ProductRepository.
func (r *SQLProductRepository) DeleteProduct(ctx context.Context, productID uuid.UUID) error {
	deletedProduct, err := r.DB.GetProductById(ctx, productID)
	if err != nil{
		log.Printf("Error fetching product with id %s: %s", productID.String(), err.Error())
		return err
	}

	err = r.DB.DeleteProduct(ctx, deletedProduct.ID)
	if err != nil {
		log.Printf("Error deleting product with id %s: %s", productID.String(), err.Error())
		return err
	}

	return nil
}

// GetAvailableProducts implements repository.ProductRepository.
func (r *SQLProductRepository) GetAvailableProducts(ctx context.Context) ([]database.Product, error) {
	availableProducts, err := r.DB.GetAvailableProducts(ctx)
	if err != nil{
		log.Printf("Error fetching available products : %s", err.Error())
		return []database.Product{}, err
	}

	return availableProducts, nil
}

// GetFilteredProducts implements repository.ProductRepository.
func (r *SQLProductRepository) GetFilteredProducts(ctx context.Context, arg database.GetFilteredProductsParams) ([]database.Product, error) {
	filteredProducts, err := r.DB.GetFilteredProducts(ctx, arg)
	if err != nil {
		log.Printf("Error fetching filtered products with args%s, %s: %s", arg.CategoryID.String(),arg.Price, err.Error())
		return []database.Product{}, err
	}
	return filteredProducts, nil
}

// GetPaginatedProducts implements repository.ProductRepository.
func (r *SQLProductRepository) GetPaginatedProducts(ctx context.Context, arg database.GetPaginatedProductsParams) ([]database.Product, error) {
	paginatedProducts, err := r.DB.GetPaginatedProducts(ctx, arg)
	if err != nil {
		log.Printf("Error fetching paginated products with args %d, %d: %s", arg.Limit,arg.Offset, err.Error())
		return []database.Product{}, err
	}
	return paginatedProducts, nil
}

// GetProductById implements repository.ProductRepository.
func (r *SQLProductRepository) GetProductById(ctx context.Context, id uuid.UUID) (database.Product, error) {
	product, err := r.DB.GetProductById(ctx, id)
	if err != nil{
		log.Printf("Error fetching product with id %s: %s", id.String(), err.Error())
		return database.Product{}, err
	}
	return product, nil
}


// GetProductWithRecommendations implements repository.ProductRepository.
func (r *SQLProductRepository) GetProductWithRecommendations(ctx context.Context, id uuid.UUID) (database.GetProductWithRecommendationsRow, error) {
	recommendedProduct, err := r.DB.GetProductWithRecommendations(ctx, id)
	if err != nil{
		log.Printf("Error fetching recommended products with reference product id %s: %s", id.String(), err.Error())
		return database.GetProductWithRecommendationsRow{}, err
	}
	return recommendedProduct, nil
}

// GetProducts implements repository.ProductRepository.
func (r *SQLProductRepository) GetProducts(ctx context.Context) ([]database.Product, error) {
	products, err := r.DB.GetProducts(ctx)
	if err != nil{
		log.Printf("Error fetching all products in the database : %s", err.Error())
		return []database.Product{}, err
	}
	return products, nil
}

// GetProductsByCategory implements repository.ProductRepository.
func (r *SQLProductRepository) GetProductsByCategory(ctx context.Context, categoryID uuid.UUID) ([]database.Product, error) {
	categorizedProducts, err := r.DB.GetProductsByCategory(ctx, categoryID)
	if err != nil {
		log.Printf("Error fetching categorized products with category id %s: %s", categoryID.String(), err.Error())
		return []database.Product{}, err
	}
	return categorizedProducts, nil
}

// GetSalesTrends implements repository.ProductRepository.
func (r *SQLProductRepository) GetSalesTrends(ctx context.Context) ([]database.GetSalesTrendsRow, error) {
	salesTrendRow, err := r.DB.GetSalesTrends(ctx)
	if err != nil{
		log.Printf("Error fetching the row of sales trends : %s",  err.Error())
		return []database.GetSalesTrendsRow{}, err
	}
	return salesTrendRow, nil
}

// SearchProducts implements repository.ProductRepository.
func (r *SQLProductRepository) SearchProducts(ctx context.Context, query sql.NullString) ([]database.Product, error) {
	queryResults, err := r.DB.SearchProducts(ctx, query)
	if err != nil{
		log.Printf("Error fetching products with query  %s: %s", query.String ,err.Error())
		return []database.Product{}, err
	}
	return queryResults, nil
}
