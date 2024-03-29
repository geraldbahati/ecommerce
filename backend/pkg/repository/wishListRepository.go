package repository

import (
	"context"
	"github.com/geraldbahati/ecommerce/pkg/model"
	"github.com/google/uuid"
)

type WishListRepository interface {
	// create
	CreateWishList(ctx context.Context, userId uuid.UUID) (model.Wishlist, error)
	CopyWishlistsIntoAnotherWishlist(ctx context.Context, sourceWishListId uuid.UUID, targetWishListId uuid.UUID) error

	// update
	UpdateWishList(ctx context.Context, id uuid.UUID, name string, visibility string) (model.Wishlist, error)
	AddItemToWishlist(ctx context.Context, wishListId uuid.UUID, productId uuid.UUID) (model.WishlistItem, error)

	// delete
	DeleteWishList(ctx context.Context, wishListId uuid.UUID) error
	RemoveItemFromWishlist(ctx context.Context, wishListId uuid.UUID, productId uuid.UUID) error

	// get
	ListAllItemsInUserWishlist(ctx context.Context, userId uuid.UUID, offset int32, limit int32) (interface{}, error)

	// track
	TrackInterestInWishlistItem(ctx context.Context) ([]model.InterestCount, error)
	FindCommonWishlistLists(ctx context.Context) ([]model.UserCount, error)
}
