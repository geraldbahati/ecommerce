package model

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Product struct {
	ID            uuid.UUID      `json:"id"`
	Name          string         `json:"name"`
	Description   sql.NullString `json:"description"`
	ImageUrl      sql.NullString `json:"image_url"`
	Price         string         `json:"price"`
	Stock         int32          `json:"stock"`
	SubCategoryID uuid.NullUUID  `json:"sub_category_id"`
	Brand         sql.NullString `json:"brand"`
	Rating        string         `json:"rating"`
	ReviewCount   int32          `json:"review_count"`
	DiscountRate  string         `json:"discount_rate"`
	Keywords      sql.NullString `json:"keywords"`
	IsActive      bool           `json:"is_active"`
	CreatedAt     time.Time      `json:"created_at"`
	LastUpdated   sql.NullTime   `json:"last_updated"`
}

type ProductListing struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Price       float32   `json:"price"`
	ImageURL    string    `json:"image_url"`
	Rating      float32   `json:"rating"`
	ReviewCount int       `json:"review_count"`
}

type ProductDetails struct {
	Product
	CategoryName string `json:"category_name"`
}

type ProductReview struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"user_id"`
	ProductID uuid.UUID `json:"product_id"`
	Rating    float32   `json:"rating"`
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"created_at"`
}

type TrendingProduct struct {
	ProductID    uuid.UUID `json:"product_id"`
	ProductName  string    `json:"product_name"`
	Price        string    `json:"price"`
	CategoryID   uuid.UUID `json:"category_id"`
	CategoryName string    `json:"category_name"`
	SalesVolume  int64     `json:"sales_volume"`
}

type AddProductParams struct {
	Name          string         `json:"name"`
	Description   sql.NullString `json:"description"`
	ImageUrl      sql.NullString `json:"image_url"`
	Price         string         `json:"price"`
	Stock         int32          `json:"stock"`
	SubCategoryID uuid.NullUUID  `json:"sub_category_id"`
	Brand         sql.NullString `json:"brand"`
	Keywords      sql.NullString `json:"keywords"`
}

type GetProductsByCategoryRow struct {
	ID              uuid.UUID      `json:"id"`
	Name            string         `json:"name"`
	Description     sql.NullString `json:"description"`
	ImageUrl        sql.NullString `json:"image_url"`
	Price           string         `json:"price"`
	Stock           int32          `json:"stock"`
	Brand           sql.NullString `json:"brand"`
	Rating          string         `json:"rating"`
	ReviewCount     int32          `json:"review_count"`
	DiscountRate    string         `json:"discount_rate"`
	Keywords        sql.NullString `json:"keywords"`
	IsActive        bool           `json:"is_active"`
	CreatedAt       time.Time      `json:"created_at"`
	LastUpdated     sql.NullTime   `json:"last_updated"`
	SubCategoryID   uuid.NullUUID  `json:"sub_category_id"`
	SubCategoryName string         `json:"sub_category_name"`
	CategoryName    string         `json:"category_name"`
}

type ProductMaterial struct {
	ID          uuid.UUID    `json:"id"`
	ProductID   uuid.UUID    `json:"product_id"`
	MaterialID  uuid.UUID    `json:"material_id"`
	CreatedAt   time.Time    `json:"created_at"`
	LastUpdated sql.NullTime `json:"last_updated"`
}

type ProductColour struct {
	ID          uuid.UUID    `json:"id"`
	ProductID   uuid.UUID    `json:"product_id"`
	ColourID    uuid.UUID    `json:"colour_id"`
	CreatedAt   time.Time    `json:"created_at"`
	LastUpdated sql.NullTime `json:"last_updated"`
}

type UpdateProductParams struct {
	ID            uuid.UUID      `json:"id"`
	Name          string         `json:"name"`
	Description   sql.NullString `json:"description"`
	ImageUrl      sql.NullString `json:"image_url"`
	Price         string         `json:"price"`
	Stock         int32          `json:"stock"`
	SubCategoryID uuid.NullUUID  `json:"sub_category_id"`
	Brand         sql.NullString `json:"brand"`
	Rating        string         `json:"rating"`
	ReviewCount   int32          `json:"review_count"`
	DiscountRate  string         `json:"discount_rate"`
	Keywords      sql.NullString `json:"keywords"`
	IsActive      bool           `json:"is_active"`
}

type UpdateProductColourParams struct {
	ID        uuid.UUID `json:"id"`
	ProductID uuid.UUID `json:"product_id"`
	ColourID  uuid.UUID `json:"colour_id"`
}

type UpdateProductMaterialParams struct {
	ID         uuid.UUID `json:"id"`
	ProductID  uuid.UUID `json:"product_id"`
	MaterialID uuid.UUID `json:"material_id"`
}
