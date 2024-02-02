package usecases

import (
	"context"

	"github.com/geraldbahati/ecommerce/pkg/model"
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

// Retrieves a list of products
func (s *ProductService) GetProductList(ctx context.Context) ([]model.ProductListing, error){
	// Logic to retrieve and return a list of products
	// Fetch only important details

	return s.productRepo.GetProductList(ctx)
}

// Retrieves detailed information for a specified product
func (s *ProductService) GetProductDetails(ctx context.Context, productID uuid.UUID) (model.ProductDetails, error){
	// Logic to retrieve and return detailed product information

	return s.productRepo.GetProductDetails(ctx, productID)
}

// Adds the product to the user's wishlist
func (s *ProductService) AddToWishlist(ctx context.Context,userID uuid.UUID, productID uuid.UUID) error{
	// Logic to add the specified product to the user's wishlist
	// TODO: Ensure proper validation and handling of duplicates if necessary
	return s.productRepo.AddToWishlist(ctx, userID,productID)
}

// TODO: Implement search, cart management, checkout process etc

// Helper function to extract the user ID from the context
func getUserIDFromContext(ctx context.Context) uuid.UUID{
	return ctx.Value("userId").(uuid.UUID)
}