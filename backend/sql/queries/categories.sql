-- name: CreateCategory :one
INSERT INTO categories (id, name, description, image_url, SEO_keywords, is_active, created_at, last_updated)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING id, name, description, image_url, SEO_keywords, is_active, created_at, last_updated;

-- name: UpdateCategory :one
UPDATE categories SET
    name = $2,
    description = $3,
    image_url = $4,
    SEO_keywords = $5,
    is_active = $6,
    last_updated = $7
WHERE id = $1
RETURNING id, name, description, image_url, SEO_keywords, is_active, created_at, last_updated;

-- name: DeleteCategory :exec
DELETE FROM categories
WHERE id = $1;

-- name: FindCategoryByID :one
SELECT * FROM categories
WHERE id = $1;

-- name: FindCategoriesBySoftName :many
SELECT * FROM categories
WHERE name ILIKE '%' || $1 || '%'
LIMIT $2 OFFSET $3;

-- name: GetActiveCategories :many
SELECT * FROM categories
WHERE is_active = TRUE
LIMIT $1 OFFSET $2;

-- name: GetInactiveCategories :many
SELECT * FROM categories
WHERE is_active = FALSE
LIMIT $1 OFFSET $2;

-- name: GetAllCategories :many
SELECT * FROM categories
LIMIT $1 OFFSET $2;

-- name: GetCategoryCount :one
SELECT COUNT(*) FROM categories;