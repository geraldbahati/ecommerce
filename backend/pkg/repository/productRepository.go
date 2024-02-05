package repository

import (
	"context"
	"database/sql"

	"github.com/geraldbahati/ecommerce/internal/database"
	"github.com/google/uuid"
)

type ProductRepository interface{
	// Create product
	AddProduct(ctx context.Context, product database.Product) (database.Product, error)

	// Update product
	UpdateProduct(ctx context.Context, arg database.UpdateProductParams) (database.Product, error)

	// Delete Product
	DeleteProduct(ctx context.Context, productID uuid.UUID) error

	// Get Product methods
	GetProducts(ctx context.Context)([]database.Product, error)

	GetAvailableProducts(ctx context.Context)([]database.Product, error)

	GetFilteredProducts(ctx context.Context, arg database.GetFilteredProductsParams)([]database.Product, error)

	GetPaginatedProducts(ctx context.Context, arg database.GetPaginatedProductsParams)([]database.Product, error)

	GetProductWithRecommendations(ctx context.Context, id uuid.UUID)(database.GetProductWithRecommendationsRow, error)

	GetProductById(ctx context.Context, id uuid.UUID)(database.Product, error)

	GetProductsByCategory(ctx context.Context, categoryID uuid.UUID)([]database.Product, error)

	// Search Products
	SearchProducts(ctx context.Context, query sql.NullString)([]database.Product, error)

	// Additional methods ...
	GetSalesTrends(ctx context.Context)([]database.GetSalesTrendsRow, error)


}