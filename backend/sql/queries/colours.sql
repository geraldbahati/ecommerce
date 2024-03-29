-- name: CreateColour :one
INSERT INTO colours (id, colour_hex, created_at, last_updated)
VALUES ($1, $2, NOW(), NULL)
RETURNING *;

-- name: UpdateColour :one
UPDATE colours SET
    colour_hex = $2,
    last_updated = NOW()
WHERE id = $1
RETURNING *;

-- name: DeleteColour :exec
DELETE FROM colours
WHERE id = $1;

-- name: GetColourByID :one
SELECT * FROM colours
WHERE id = $1;

-- name: GetColourByHex :one
SELECT * FROM colours
WHERE colour_hex = $1;

-- name: GetColours :many
SELECT * FROM colours
LIMIT $1 OFFSET $2;

-- name: GetColourCount :one
SELECT COUNT(*) FROM colours;

-- name: GetProductColours :many
SELECT c.* FROM colours c
    INNER JOIN product_colours pc ON c.id = pc.colour_id
WHERE pc.product_id = $1
LIMIT $2 OFFSET $3;