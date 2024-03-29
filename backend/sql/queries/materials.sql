-- name: CreateMaterial :one
INSERT INTO materials (id, name, created_at, last_updated)
VALUES ($1, $2, NOW(), NULL)
RETURNING *;

-- name: UpdateMaterial :one
UPDATE materials SET
    name = $2,
    last_updated = NOW()
WHERE id = $1
RETURNING *;

-- name: DeleteMaterial :exec
DELETE FROM materials
WHERE id = $1;

-- name: FindMaterialByID :one
SELECT * FROM materials
WHERE id = $1;

-- name: FindMaterialsBySoftName :many
SELECT * FROM materials
WHERE name ILIKE '%' || $1 || '%'
LIMIT $2 OFFSET $3;

-- name: GetMaterials :many
SELECT * FROM materials
LIMIT $1 OFFSET $2;

-- name: GetMaterialCount :one
SELECT COUNT(*) FROM materials;

-- name: GetProductMaterials :many
SELECT m.* FROM materials m
    INNER JOIN product_materials pm ON m.id = pm.material_id
WHERE pm.product_id = $1
LIMIT $2 OFFSET $3;

-- name: GetMaterialByName :one
SELECT * FROM materials
WHERE name = $1;