package model

import (
	"database/sql"
	"github.com/google/uuid"
	"time"
)

type Wishlist struct {
	ID          uuid.UUID    `json:"id"`
	UserID      uuid.UUID    `json:"user_id"`
	Name        string       `json:"name"`
	Visibility  string       `json:"visibility"`
	CreatedAt   time.Time    `json:"created_at"`
	LastUpdated sql.NullTime `json:"last_updated"`
}

type WishlistItem struct {
	ID          uuid.UUID    `json:"id"`
	WishlistID  uuid.UUID    `json:"wishlist_id"`
	ProductID   uuid.UUID    `json:"product_id"`
	Priority    string       `json:"priority"`
	CreatedAt   time.Time    `json:"created_at"`
	LastUpdated sql.NullTime `json:"last_updated"`
}

type InterestCount struct {
	ProductID uuid.UUID `json:"product_id"`
	Count     int64     `json:"count"`
}

type UserCount struct {
	ProductID uuid.UUID `json:"product_id"`
	Count     int64     `json:"count"`
}
