package sqlc

import (
	"context"
	"database/sql"
	"github.com/geraldbahati/ecommerce/pkg/model"
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

// AddProduct creates a new product in the database
func (r *SQLProductRepository) AddProduct(ctx context.Context, product model.AddProductParams) (model.Product, error) {
	// Add product into database
	addProduct, err := r.DB.CreateProduct(ctx, database.CreateProductParams{
		ID:            uuid.New(),
		Name:          product.Name,
		Description:   product.Description,
		ImageUrl:      product.ImageUrl,
		Price:         product.Price,
		Stock:         product.Stock,
		SubCategoryID: product.SubCategoryID,
		Brand:         product.Brand,
		Keywords:      product.Keywords,
	})
	if err != nil {
		return model.Product{}, err
	}

	// Return newly added product
	return model.Product{
		ID:            addProduct.ID,
		Name:          addProduct.Name,
		Description:   addProduct.Description,
		ImageUrl:      addProduct.ImageUrl,
		Price:         addProduct.Price,
		Stock:         addProduct.Stock,
		SubCategoryID: addProduct.SubCategoryID,
		Brand:         addProduct.Brand,
		Rating:        addProduct.Rating,
		ReviewCount:   addProduct.ReviewCount,
		DiscountRate:  addProduct.DiscountRate,
		Keywords:      addProduct.Keywords,
		IsActive:      addProduct.IsActive,
		CreatedAt:     addProduct.CreatedAt,
		LastUpdated:   addProduct.LastUpdated,
	}, err
}

// Updates an already existing product in the database
func (r *SQLProductRepository) UpdateProduct(ctx context.Context, product database.UpdateProductParams) (database.Product, error) {
	updatedProduct, err := r.DB.UpdateProduct(ctx, database.UpdateProductParams{
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
	})

	if err != nil {
		log.Printf("Error updating product with id %s: %s", product.ID.String(), err.Error())
		return database.Product{}, err
	}

	// Return updated Product
	return database.Product{
		ID:            updatedProduct.ID,
		Name:          updatedProduct.Name,
		Description:   updatedProduct.Description,
		ImageUrl:      updatedProduct.ImageUrl,
		Price:         updatedProduct.Price,
		Stock:         updatedProduct.Stock,
		SubCategoryID: updatedProduct.SubCategoryID,
		Brand:         updatedProduct.Brand,
		Rating:        updatedProduct.Rating,
		ReviewCount:   updatedProduct.ReviewCount,
		DiscountRate:  updatedProduct.DiscountRate,
		Keywords:      updatedProduct.Keywords,
		IsActive:      updatedProduct.IsActive,
		CreatedAt:     updatedProduct.CreatedAt,
		LastUpdated:   updatedProduct.LastUpdated,
	}, err
}

// DeleteProduct implements repository.ProductRepository.
func (r *SQLProductRepository) DeleteProduct(ctx context.Context, productID uuid.UUID) error {
	deletedProduct, err := r.DB.GetProductById(ctx, productID)
	if err != nil {
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
	if err != nil {
		log.Printf("Error fetching available products : %s", err.Error())
		return []database.Product{}, err
	}

	return availableProducts, nil
}

// GetProductById implements repository.ProductRepository.
func (r *SQLProductRepository) GetProductById(ctx context.Context, id uuid.UUID) (database.Product, error) {
	product, err := r.DB.GetProductById(ctx, id)
	if err != nil {
		log.Printf("Error fetching product with id %s: %s", id.String(), err.Error())
		return database.Product{}, err
	}
	return product, nil
}

// GetProducts implements repository.ProductRepository.
func (r *SQLProductRepository) GetProducts(ctx context.Context, offset int32, limit int32) (interface{}, error) {
	products, err := r.DB.GetProducts(ctx, database.GetProductsParams{
		Offset: offset,
		Limit:  limit,
	})
	if err != nil {
		log.Printf("Error fetching all products in the database : %s", err.Error())
		return []database.Product{}, err
	}
	return products, nil
}

// GetProductsByCategory implements repository.ProductRepository.
func (r *SQLProductRepository) GetProductsByCategory(ctx context.Context, categoryID uuid.UUID) (interface{}, error) {
	categorizedProducts, err := r.DB.GetProductsByCategory(ctx, database.GetProductsByCategoryParams{
		ID: categoryID,
	})
	if err != nil {
		log.Printf("Error fetching categorized products with category id %s: %s", categoryID.String(), err.Error())
		return []database.Product{}, err
	}
	return categorizedProducts, nil
}

// GetSalesTrends implements repository.ProductRepository.
func (r *SQLProductRepository) GetSalesTrends(ctx context.Context) ([]database.GetSalesTrendsRow, error) {
	salesTrendRow, err := r.DB.GetSalesTrends(ctx)
	if err != nil {
		log.Printf("Error fetching the row of sales trends : %s", err.Error())
		return []database.GetSalesTrendsRow{}, err
	}
	return salesTrendRow, nil
}

// SearchProducts implements repository.ProductRepository.
func (r *SQLProductRepository) SearchProducts(ctx context.Context, query sql.NullString) ([]database.Product, error) {
	queryResults, err := r.DB.SearchProducts(ctx, query)
	if err != nil {
		log.Printf("Error fetching products with query  %s: %s", query.String, err.Error())
		return []database.Product{}, err
	}
	return queryResults, nil
}

// GetTrendingProducts implements repository.ProductRepository.
func (r *SQLProductRepository) GetTrendingProducts(ctx context.Context) ([]model.TrendingProduct, error) {
	trendingProducts, err := r.DB.GetTrendingProducts(ctx)
	if err != nil {
		log.Printf("Error fetching trending products : %s", err.Error())
		return []model.TrendingProduct{}, err
	}

	// Return trending products
	var modelTrendingProducts []model.TrendingProduct
	for _, product := range trendingProducts {
		modelTrendingProducts = append(modelTrendingProducts, model.TrendingProduct{
			ProductID:    product.ProductID,
			ProductName:  product.ProductName,
			Price:        product.Price,
			CategoryID:   product.CategoryID,
			CategoryName: product.CategoryName,
			SalesVolume:  product.SalesVolume,
		})
	}

	return modelTrendingProducts, nil
}

// GetProductCountByCategory implements repository.ProductRepository.
func (r *SQLProductRepository) GetProductCountByCategory(ctx context.Context, categoryID uuid.UUID) (int64, error) {
	productCount, err := r.DB.GetProductCountByCategory(ctx, categoryID)
	if err != nil {
		log.Printf("Error fetching product count by category id %s: %s", categoryID.String(), err.Error())
		return 0, err
	}
	return productCount, nil
}

// GetProductCount implements repository.ProductRepository.
func (r *SQLProductRepository) GetProductCount(ctx context.Context) (int64, error) {
	productCount, err := r.DB.GetProductCount(ctx)
	if err != nil {
		log.Printf("Error fetching product count : %s", err.Error())
		return 0, err
	}
	return productCount, nil
}
