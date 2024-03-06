package sqlc

import (
	"context"
	"github.com/geraldbahati/ecommerce/internal/database"
	"github.com/geraldbahati/ecommerce/pkg/model"
	"github.com/google/uuid"
)

type SQLWishlistRepository struct {
	DB *database.Queries
}

func NewSQLWishlistRepository(db *database.Queries) *SQLWishlistRepository {
	return &SQLWishlistRepository{
		DB: db,
	}
}

// CreateWishList creates a new wishlist
func (r *SQLWishlistRepository) CreateWishList(ctx context.Context, userId uuid.UUID) (model.Wishlist, error) {
	// insert wishlist into database
	createdWishList, err := r.DB.CreateWishlist(ctx, userId)
	if err != nil {
		return model.Wishlist{}, err
	}

	// return created wishlist
	return model.Wishlist{
		ID:          createdWishList.ID,
		UserID:      createdWishList.UserID,
		Name:        createdWishList.Name,
		Visibility:  createdWishList.Visibility,
		CreatedAt:   createdWishList.CreatedAt,
		LastUpdated: createdWishList.LastUpdated,
	}, nil
}

// CopyWishlistsIntoAnotherWishlist copies all items from one wishlist into another
//func (r *SQLWishlistRepository) CopyWishlistsIntoAnotherWishlist(ctx context.Context, sourceWishListId uuid.UUID, targetWishListId uuid.UUID) error {
//	_, err := r.DB.CopyWishlistsIntoAnotherWishlist(ctx, database.CopyWishlistsIntoAnotherWishlistParams{
//		WishlistID:   sourceWishListId,
//		WishlistID_2: targetWishListId,
//	})
//	return err
//}

// UpdateWishList updates a wishlist
func (r *SQLWishlistRepository) UpdateWishList(ctx context.Context, id uuid.UUID, name string, visibility string) (model.Wishlist, error) {
	// update wishlist in database
	updatedWishList, err := r.DB.UpdateWishlist(ctx, database.UpdateWishlistParams{
		ID:         id,
		Name:       name,
		Visibility: visibility,
	})
	if err != nil {
		return model.Wishlist{}, err
	}

	// return updated wishlist
	return model.Wishlist{
		ID:          updatedWishList.ID,
		UserID:      updatedWishList.UserID,
		Name:        updatedWishList.Name,
		Visibility:  updatedWishList.Visibility,
		CreatedAt:   updatedWishList.CreatedAt,
		LastUpdated: updatedWishList.LastUpdated,
	}, nil
}

// AddItemToWishlist adds an item to a wishlist
func (r *SQLWishlistRepository) AddItemToWishlist(ctx context.Context, wishListId uuid.UUID, productId uuid.UUID) (model.WishlistItem, error) {
	// add item to wishlist in database
	addedItem, err := r.DB.AddItemToWishlist(ctx, database.AddItemToWishlistParams{
		WishlistID: wishListId,
		ProductID:  productId,
	})
	if err != nil {
		return model.WishlistItem{}, err
	}

	// return added item
	return model.WishlistItem{
		ID:          addedItem.ID,
		WishlistID:  addedItem.WishlistID,
		ProductID:   addedItem.ProductID,
		Priority:    addedItem.Priority,
		CreatedAt:   addedItem.CreatedAt,
		LastUpdated: addedItem.LastUpdated,
	}, nil
}

// DeleteWishList deletes a wishlist
func (r *SQLWishlistRepository) DeleteWishList(ctx context.Context, wishListId uuid.UUID) error {
	//tx, err := db.BeginTx(ctx, nil)
	return nil
}

// RemoveItemFromWishlist removes an item from a wishlist
func (r *SQLWishlistRepository) RemoveItemFromWishlist(ctx context.Context, wishListId uuid.UUID, productId uuid.UUID) error {
	// remove item from wishlist in database
	err := r.DB.RemoveItemFromWishlist(ctx, database.RemoveItemFromWishlistParams{
		WishlistID: wishListId,
		ProductID:  productId,
	})
	return err
}

// ListAllItemsInUserWishlist lists all items in a user's wishlist
func (r *SQLWishlistRepository) ListAllItemsInUserWishlist(ctx context.Context, userId uuid.UUID, offset int32, limit int32) (interface{}, error) {
	// list all items in user's wishlist from database
	items, err := r.DB.ListAllItemsInUserWishlist(ctx, database.ListAllItemsInUserWishlistParams{
		UserID: userId,
		Offset: offset,
		Limit:  limit,
	})
	if err != nil {
		return nil, err
	}

	// return items
	return items, nil
}

// TrackInterestInWishlistItem tracks interest in a wishlist item
func (r *SQLWishlistRepository) TrackInterestInWishlistItem(ctx context.Context) ([]model.InterestCount, error) {
	// track interest in wishlist item in database
	interest, err := r.DB.TrackInterestInWishlistItem(ctx)
	if err != nil {
		return nil, err
	}

	// return interest
	modelInterest := make([]model.InterestCount, len(interest))
	for i, v := range interest {
		modelInterest[i] = model.InterestCount{
			ProductID: v.ProductID,
			Count:     v.InterestCount,
		}
	}
	return modelInterest, nil
}

// FindCommonWishlistLists finds common wishlist lists
func (r *SQLWishlistRepository) FindCommonWishlistLists(ctx context.Context) ([]model.UserCount, error) {
	// find common wishlist lists in database
	users, err := r.DB.FindCommonWishlistLists(ctx)
	if err != nil {
		return nil, err
	}

	// return users
	modelUsers := make([]model.UserCount, len(users))
	for i, v := range users {
		modelUsers[i] = model.UserCount{
			ProductID: v.ProductID,
			Count:     v.UserCount,
		}
	}
	return modelUsers, nil
}
