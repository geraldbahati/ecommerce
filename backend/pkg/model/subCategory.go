package model

import (
	"database/sql"
	"github.com/google/uuid"
	"time"
)

type SubCategory struct {
	ID          uuid.UUID      `json:"id"`
	CategoryID  uuid.UUID      `json:"category_id"`
	Name        string         `json:"name"`
	Description sql.NullString `json:"description"`
	ImageUrl    sql.NullString `json:"image_url"`
	SeoKeywords sql.NullString `json:"seo_keywords"`
	IsActive    bool           `json:"is_active"`
	CreatedAt   time.Time      `json:"created_at"`
	LastUpdated sql.NullTime   `json:"last_updated"`
}

type AddSubCategoryParams struct {
	ID          uuid.UUID      `json:"id"`
	CategoryID  uuid.UUID      `json:"category_id"`
	Name        string         `json:"name"`
	Description sql.NullString `json:"description"`
	ImageUrl    sql.NullString `json:"image_url"`
	SeoKeywords sql.NullString `json:"seo_keywords"`
}
