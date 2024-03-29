-- name: UpdateProductMaterial :one
INSERT INTO product_materials (id, product_id, material_id, created_at, last_updated)
VALUES ($1, $2, $3, NOW(), NULL)
RETURNING *;

-- name: DeleteProductMaterial :exec
DELETE FROM product_materials
WHERE id = $1;

-- name: GetProductsByMaterial :many
SELECT p.* FROM products p
    INNER JOIN product_materials pm ON p.id = pm.product_id
WHERE pm.material_id = $1
LIMIT $2 OFFSET $3;