// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: product_colours.sql

package database

import (
	"context"

	"github.com/google/uuid"
)

const deleteProductColour = `-- name: DeleteProductColour :exec
DELETE FROM product_colours
WHERE id = $1
`

func (q *Queries) DeleteProductColour(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteProductColour, id)
	return err
}

const getProductsByColour = `-- name: GetProductsByColour :many
SELECT p.id, p.name, p.description, p.image_url, p.price, p.stock, p.brand, p.rating, p.review_count, p.discount_rate, p.keywords, p.is_active, p.created_at, p.last_updated, p.sub_category_id FROM products p
    INNER JOIN product_colours pc ON p.id = pc.product_id
WHERE pc.colour_id = $1
LIMIT $2 OFFSET $3
`

type GetProductsByColourParams struct {
	ColourID uuid.UUID
	Limit    int32
	Offset   int32
}

func (q *Queries) GetProductsByColour(ctx context.Context, arg GetProductsByColourParams) ([]Product, error) {
	rows, err := q.db.QueryContext(ctx, getProductsByColour, arg.ColourID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Product
	for rows.Next() {
		var i Product
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.ImageUrl,
			&i.Price,
			&i.Stock,
			&i.Brand,
			&i.Rating,
			&i.ReviewCount,
			&i.DiscountRate,
			&i.Keywords,
			&i.IsActive,
			&i.CreatedAt,
			&i.LastUpdated,
			&i.SubCategoryID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateProductColour = `-- name: UpdateProductColour :one
INSERT INTO product_colours (id, product_id, colour_id, created_at, last_updated)
VALUES ($1, $2, $3, NOW(), NULL)
RETURNING id, product_id, colour_id, created_at, last_updated
`

type UpdateProductColourParams struct {
	ID        uuid.UUID
	ProductID uuid.UUID
	ColourID  uuid.UUID
}

func (q *Queries) UpdateProductColour(ctx context.Context, arg UpdateProductColourParams) (ProductColour, error) {
	row := q.db.QueryRowContext(ctx, updateProductColour, arg.ID, arg.ProductID, arg.ColourID)
	var i ProductColour
	err := row.Scan(
		&i.ID,
		&i.ProductID,
		&i.ColourID,
		&i.CreatedAt,
		&i.LastUpdated,
	)
	return i, err
}
