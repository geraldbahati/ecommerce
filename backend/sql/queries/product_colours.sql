-- name: UpdateProductColour :one
INSERT INTO product_colours (id, product_id, colour_id, created_at, last_updated)
VALUES ($1, $2, $3, NOW(), NULL)
RETURNING *;

-- name: DeleteProductColour :exec
DELETE FROM product_colours
WHERE id = $1;

-- name: GetProductsByColour :many
SELECT p.* FROM products p
    INNER JOIN product_colours pc ON p.id = pc.product_id
WHERE pc.colour_id = $1
LIMIT $2 OFFSET $3;