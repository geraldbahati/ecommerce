package usecases

import (
	"context"
	"github.com/geraldbahati/ecommerce/pkg/model"
	"github.com/geraldbahati/ecommerce/pkg/repository"
	"github.com/google/uuid"
)

type WishlistService struct {
	wishlistRepo repository.WishListRepository
}

func NewWishlistService(wishlistRepo repository.WishListRepository) *WishlistService {
	return &WishlistService{
		wishlistRepo: wishlistRepo,
	}
}

// CreateWishlist creates a new wishlist
func (s *WishlistService) CreateWishlist(ctx context.Context) (model.Wishlist, error) {
	// get user id from context
	userId := ctx.Value("userId").(uuid.UUID)

	// create wishlist
	return s.wishlistRepo.CreateWishList(ctx, userId)
}

// AddProductToWishlist adds a product to a wishlist
func (s *WishlistService) AddProductToWishlist(ctx context.Context, productId string, wishlistId string) (model.WishlistItem, error) {
	// convert wishlist id to uuid
	wishlistUUID, err := uuid.Parse(wishlistId)
	if err != nil {
		return model.WishlistItem{}, err
	}

	// convert product id to uuid
	productIdUUID, err := uuid.Parse(productId)
	if err != nil {
		return model.WishlistItem{}, err
	}

	// add product to wishlist
	return s.wishlistRepo.AddItemToWishlist(ctx, wishlistUUID, productIdUUID)
}

// RemoveProductFromWishlist removes a product from a wishlist
func (s *WishlistService) RemoveProductFromWishlist(ctx context.Context, productId string, wishlistId string) error {
	// convert wishlist id to uuid
	wishlistUUID, err := uuid.Parse(wishlistId)
	if err != nil {
		return err
	}

	// convert product id to uuid
	productIdUUID, err := uuid.Parse(productId)
	if err != nil {
		return err
	}

	// remove product from wishlist
	return s.wishlistRepo.RemoveItemFromWishlist(ctx, wishlistUUID, productIdUUID)
}

// ListAllItemsInUserWishlist lists all products in a wishlist
func (s *WishlistService) ListAllItemsInUserWishlist(ctx context.Context, offset int32, limit int32) (interface{}, error) {
	// get user id from context
	userId := ctx.Value("userId").(uuid.UUID)

	// list all products in wishlist
	return s.wishlistRepo.ListAllItemsInUserWishlist(ctx, userId, offset, limit)
}

// UpdateWishlist updates a wishlist
func (s *WishlistService) UpdateWishlist(ctx context.Context, wishlistId string, name string, visibility string) (model.Wishlist, error) {
	// convert wishlist id to uuid
	wishlistUUID, err := uuid.Parse(wishlistId)
	if err != nil {
		return model.Wishlist{}, err
	}

	// update wishlist
	return s.wishlistRepo.UpdateWishList(ctx, wishlistUUID, name, visibility)
}

// DeleteWishlist deletes a wishlist
func (s *WishlistService) DeleteWishlist(ctx context.Context, wishlistId string) error {
	// convert wishlist id to uuid
	wishlistUUID, err := uuid.Parse(wishlistId)
	if err != nil {
		return err
	}

	// delete wishlist
	return s.wishlistRepo.DeleteWishList(ctx, wishlistUUID)
}

// WishlistNotification sends a notification to a user when a product in their wishlist is on sale

// WishlistCleanup automatically removes products from a wishlist that are no longer available
//func (s *WishlistService) WishlistCleanup(ctx context.Context) error {
//	// get user id from context
//	userId := ctx.Value("userId").(uuid.UUID)
//
//	// cleanup wishlist
//	return s.wishlistRepo.WishlistCleanup(ctx, userId)
//}

// WishlistStats returns statistics about a wishlist

// WishlistItemInterestTracks the interest of a user in a product in a wishlist
