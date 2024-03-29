-- name: SearchProducts :many
SELECT * FROM products
WHERE name ILIKE '%' || $1 || '%' OR keywords ILIKE '%' || $1 || '%';

-- name: GetSalesTrends :many
SELECT DATE_TRUNC('month', created_at) AS month, SUM(price) AS total_sales
FROM orders
GROUP BY month
ORDER BY month;

-- name: CheckProductStock :many
SELECT id FROM products
WHERE stock > 0
AND (last_updated > NOW() - INTERVAL '1 DAY');

-- Product structure changes
--CREATE TABLE products (
-- id UUID PRIMARY KEY,
--     name VARCHAR(50) NOT NULL,
--     description VARCHAR(255) NULL,
--     image_url VARCHAR(100) NULL,
--     price DECIMAL(10, 2) NOT NULL,
--     stock INT NOT NULL,
--     sub_category_id UUID NOT NULL,
--     brand VARCHAR(50) NULL,
--     rating DECIMAL(2, 1) NOT NULL DEFAULT 0.0,
--     review_count INT NOT NULL DEFAULT 0,
--     discount_rate DECIMAL(2, 1) NOT NULL DEFAULT 0.0,
--     keywords VARCHAR(100) NULL,
--     is_active BOOLEAN NOT NULL DEFAULT TRUE,
--     created_at TIMESTAMP NOT NULL DEFAULT NOW(),
--     last_updated TIMESTAMP NULL,
--     FOREIGN KEY (sub_category_id) REFERENCES categories (id) ON DELETE CASCADE,
--     UNIQUE (name)
-- );

-- name: CreateProduct :one
INSERT INTO products (id, name, description, image_url, price, stock, sub_category_id, brand, rating, review_count, discount_rate, keywords, is_active, created_at, last_updated)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, 0.0, 0, 0.0, $9, TRUE, NOW(), NULL)
RETURNING *;

-- name: UpdateProduct :one
UPDATE products SET
    name = $2,
    description = $3,
    image_url = $4,
    price = $5,
    stock = $6,
    sub_category_id = $7,
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
WHERE id = $1;

-- name: GetProducts :many
SELECT * FROM products
ORDER BY created_at DESC
LIMIT $1 OFFSET $2;

-- name: GetProductById :one
SELECT * FROM products
WHERE id = $1;

-- name: GetProductsByCategory :many
SELECT p.*, sc.name AS sub_category_name, c.name AS category_name
FROM products p
    INNER JOIN sub_categories sc ON p.sub_category_id = sc.id
    INNER JOIN categories c ON sc.category_id = c.id
WHERE c.id = $1
ORDER BY p.created_at DESC
LIMIT $2 OFFSET $3;

-- name: GetAvailableProducts :many
SELECT * FROM products
WHERE stock > 0 AND is_active = TRUE;

-- name: GetProductCount :one
SELECT COUNT(*) FROM products;

-- name: GetProductCountByCategory :one
SELECT COUNT(*)
FROM products p
    INNER JOIN sub_categories sc ON p.sub_category_id = sc.id
    INNER JOIN categories c ON sc.category_id = c.id
WHERE c.id = $1;

-- name: GetTrendingProducts :many
WITH TrendingProducts AS (
    SELECT
        sc.category_id,
        p.id AS product_id,
        SUM(oi.quantity) AS sales_volume
    FROM
        order_items oi
            JOIN orders o ON oi.order_id = o.id
            JOIN products p ON oi.product_id = p.id
            JOIN sub_categories sc ON p.sub_category_id = sc.id
    WHERE
        o.created_at > NOW() - INTERVAL '1 month'
GROUP BY
    sc.category_id, p.id
    )
SELECT
    tp.product_id,
    p.name AS product_name,
    p.price,
    sc.id AS sub_category_id,
    sc.name AS sub_category_name,
    c.id AS category_id,
    c.name AS category_name,
    tp.sales_volume
FROM
    TrendingProducts tp
        JOIN products p ON tp.product_id = p.id
        JOIN sub_categories sc ON p.sub_category_id = sc.id
        JOIN categories c ON sc.category_id = c.id
ORDER BY
    c.name, tp.sales_volume DESC;


-- name: GetTrendingProductsByCategory :many
WITH TrendingProducts AS (
    SELECT
        sc.category_id,
        p.id AS product_id,
        SUM(oi.quantity) AS sales_volume
    FROM
        order_items oi
            JOIN orders o ON oi.order_id = o.id
            JOIN products p ON oi.product_id = p.id
            JOIN sub_categories sc ON p.sub_category_id = sc.id
    WHERE
        o.created_at > NOW() - INTERVAL '1 month'
    AND sc.category_id = $1
GROUP BY
    sc.category_id, p.id
    )
SELECT
    tp.product_id,
    p.name AS product_name,
    p.price,
    sc.id AS sub_category_id,
    sc.name AS sub_category_name,
    c.id AS category_id,
    c.name AS category_name,
    tp.sales_volume
FROM
    TrendingProducts tp
        JOIN products p ON tp.product_id = p.id
        JOIN sub_categories sc ON p.sub_category_id = sc.id
        JOIN categories c ON sc.category_id = c.id
ORDER BY
    tp.sales_volume DESC;