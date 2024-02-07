-- name: AddProduct :one
INSERT INTO products (id, name, description, image_url, price, stock, category_id, brand, rating, review_count, discount_rate, keywords, is_active, created_at, last_updated)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, NOW(), NULL)
RETURNING *;

-- name: UpdateProduct :one
UPDATE products SET
    name = $2,
    description = $3,
    image_url = $4,
    price = $5,
    stock = $6,
    category_id = $7,
    brand = $8,
    rating = $9,
    review_count = $10,
    discount_rate = $11,
    keywords = $12,
    is_active = $13,
    last_updated = NOW()
WHERE id = $1
RETURNING *;

-- name: DeleteProduct :exec
UPDATE products SET
    is_active = FALSE
WHERE id = $1
RETURNING *;

-- name: GetProducts :many
SELECT * FROM products;

-- name: GetProductById :one
SELECT * FROM products
WHERE id = $1;

-- name: GetProductsByCategory :many
SELECT * FROM products
WHERE category_id = $1;

-- name: GetAvailableProducts :many
SELECT * FROM products
WHERE stock > 0 AND is_active = TRUE;

-- name: GetProductWithRecommendations :one
WITH current_product AS (
    SELECT * FROM products WHERE products.id = $1
)
SELECT cp.*,
    (
        SELECT jsonb_agg(jsonb_build_object('id',  rec.id, 'name',rec.name, 'price',rec.price))
        FROM products rec
        WHERE rec.category_id = cp.category_id AND rec.id != cp.id
        ORDER BY random()
        LIMIT 5
    ) AS recommendations
FROM current_product cp;

-- name: GetPaginatedProducts :many
SELECT * FROM products
ORDER BY created_at DESC
OFFSET $1 LIMIT $2;

-- name: GetFilteredProducts :many
SELECT * FROM products
WHERE category_id = $1 AND price <= $2;

-- name: SearchProducts :many
SELECT * FROM products
WHERE name ILIKE '%' || $1 || '%' OR keywords ILIKE '%' || $1 || '%';

-- name: GetSalesTrends :many
SELECT DATE_TRUNC('month', created_at) AS month, SUM(price) AS total_sales
FROM orders
GROUP BY month
ORDER BY month;