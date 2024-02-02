package model

import (
	"time"

	"github.com/google/uuid"
)

type Product struct{
	ID uuid.UUID					`json:"id"`
	Name string						`json:"name"`
	Description string				`json:"description"`
	Price float32					`json:"price"`
	StockQuantity int				`json:"stock_quantity"`
	CategoryID uuid.UUID			`json:"category_id"`
	ImageURLs []string				`json:"image_urls"`
	Brand string					`json:"brand"`
	Color string					`json:"color"`
	Rating float32					`json:"rating"`
	ReviewCount int					`json:"review_count"`
	DiscountRate float32			`json:"discount_rate"`
	Tags []string					`json:"tags"`
	WarrantyPeriod string			`json:"warranty_period"`
	RelatedProducts []uuid.UUID		`json:"related_products"`
	CreatedAt time.Time				`json:"created_at"`
	UpdatedAt time.Time				`json:"updated_at"`
}

type ProductListing struct{
	ID uuid.UUID		`json:"id"`
	Name string			`json:"name"`
	Price float32		`json:"price"`
	ImageURL string		`json:"image_url"`
	Rating float32		`json:"rating"`
	ReviewCount int 	`json:"review_count"`
}

type ProductDetails struct{
	Product
	CategoryName string		`json:"category_name"`
}

type ProductReview struct{
	ID uuid.UUID			`json:"id"`
	UserID uuid.UUID		`json:"user_id"`
	ProductID uuid.UUID		`json:"product_id"`
	Rating float32			`json:"rating"`
	Comment string			`json:"comment"`
	CreatedAt time.Time		`json:"created_at"`
}