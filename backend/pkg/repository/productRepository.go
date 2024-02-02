package repository

import (
	"context"

	"github.com/geraldbahati/ecommerce/pkg/model"
	"github.com/google/uuid"
)

type ProductRepository interface{
	// Create product
	CreateProduct(ctx context.Context, product model.Product) (model.Product, error)

	// Update product
	UpdateProduct(ctx context.Context, product model.Product) (model.Product, error)

	// Add to wishlist
	AddToWishlist(ctx context.Context, userID uuid.UUID,productID uuid.UUID) error

	// Delete Product
	DeleteProduct(ctx context.Context, productID uuid.UUID) error

	// Get Product
	GetProductList(ctx context.Context) ([]model.ProductListing, error)
	GetProductDetails(ctx context.Context, productID uuid.UUID) (model.ProductDetails, error)

	// Additional methods ...
}