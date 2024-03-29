-- name: CreateSubCategory :one
INSERT INTO sub_categories (id, category_id, name, description, image_url, SEO_keywords, is_active, created_at, last_updated)
VALUES ($1, $2, $3, $4, $5, $6, true, NOW(), NOW())
RETURNING *;

-- name: GetSubCategory :one
SELECT id, category_id, name, description, image_url, SEO_keywords, is_active, created_at, last_updated
FROM sub_categories
WHERE id = $1;

-- name: ListAllSubCategories :many
SELECT id, category_id, name, description, image_url, SEO_keywords, is_active, created_at, last_updated
FROM sub_categories;

-- name: UpdateSubCategory :one
UPDATE sub_categories SET
    category_id = $2,
    name = $3,
    description = $4,
    image_url = $5,
    SEO_keywords = $6,
    last_updated = NOW()
WHERE id = $1
RETURNING *;

-- name: DeleteSubCategory :exec
DELETE FROM sub_categories
WHERE id = $1;

-- name: GetSubCategoryByCategory :many
SELECT id, category_id, name, description, image_url, SEO_keywords, is_active, created_at, last_updated
FROM sub_categories
WHERE category_id = $1
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;

-- name: GetProductBySubCategory :many
SELECT *
FROM products
WHERE sub_category_id = $1
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;


-- name: GetProductCountBySubCategory :one
SELECT COUNT(*)
FROM products
WHERE sub_category_id = $1;

-- name: GetSubCategoryCountByCategory :one
SELECT COUNT(*)
FROM sub_categories
WHERE category_id = $1;