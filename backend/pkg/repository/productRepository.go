package repository

import (
	"context"
	"database/sql"
	"github.com/geraldbahati/ecommerce/pkg/model"

	"github.com/geraldbahati/ecommerce/internal/database"
	"github.com/google/uuid"
)

type ProductRepository interface {
	// Create product
	AddProduct(ctx context.Context, product model.AddProductParams) (model.Product, error)

	// Update product
	UpdateProduct(ctx context.Context, product database.UpdateProductParams) (database.Product, error)

	// Delete Product
	DeleteProduct(ctx context.Context, productID uuid.UUID) error

	// Get Product methods
	GetProducts(ctx context.Context, offset int32, limit int32) (interface{}, error)

	GetAvailableProducts(ctx context.Context) ([]database.Product, error)

	GetProductById(ctx context.Context, id uuid.UUID) (database.Product, error)

	GetProductsByCategory(ctx context.Context, categoryID uuid.UUID) (interface{}, error)

	GetProductCountByCategory(ctx context.Context, categoryID uuid.UUID) (int64, error)

	GetProductCount(ctx context.Context) (int64, error)

	GetTrendingProducts(ctx context.Context) ([]model.TrendingProduct, error)

	// Search Products
	SearchProducts(ctx context.Context, query sql.NullString) ([]database.Product, error)

	// Additional methods ...
	GetSalesTrends(ctx context.Context) ([]database.GetSalesTrendsRow, error)
}
