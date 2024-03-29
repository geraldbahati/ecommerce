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
	CreateProductColour(ctx context.Context, colourHex string) (model.Colour, error)
	CreateProductMaterial(ctx context.Context, materialName string) (model.Material, error)

	// Update product
	UpdateProduct(ctx context.Context, product model.UpdateProductParams) (model.Product, error)
	UpdateProductColour(ctx context.Context, productId uuid.UUID, colourId uuid.UUID) (model.ProductColour, error)
	UpdateProductMaterial(ctx context.Context, productId uuid.UUID, materialId uuid.UUID) (model.ProductMaterial, error)

	// Delete Product
	DeleteProduct(ctx context.Context, productID uuid.UUID) error

	// Get Product methods
	GetProducts(ctx context.Context, offset int32, limit int32) (interface{}, error)

	GetProductColours(ctx context.Context, productId uuid.UUID, offset int32, limit int32) ([]model.Colour, error)
	GetAllMaterials(ctx context.Context, offset int32, limit int32) ([]model.Material, error)
	GetAllColours(ctx context.Context, offset int32, limit int32) ([]model.Colour, error)
	GetProductMaterials(ctx context.Context, productId uuid.UUID, offset int32, limit int32) ([]model.Material, error)
	GetColourByHex(ctx context.Context, hex string) (model.Colour, error)
	GetColourCount(ctx context.Context) (int64, error)
	GetMaterialCount(ctx context.Context) (int64, error)
	GetMaterialByName(ctx context.Context, materialName string) (model.Material, error)

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
