// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: materials.sql

package database

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

const createMaterial = `-- name: CreateMaterial :one
INSERT INTO materials (id, name, created_at, last_updated)
VALUES ($1, $2, NOW(), NULL)
RETURNING id, name, created_at, last_updated
`

type CreateMaterialParams struct {
	ID   uuid.UUID
	Name string
}

func (q *Queries) CreateMaterial(ctx context.Context, arg CreateMaterialParams) (Material, error) {
	row := q.db.QueryRowContext(ctx, createMaterial, arg.ID, arg.Name)
	var i Material
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatedAt,
		&i.LastUpdated,
	)
	return i, err
}

const deleteMaterial = `-- name: DeleteMaterial :exec
DELETE FROM materials
WHERE id = $1
`

func (q *Queries) DeleteMaterial(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteMaterial, id)
	return err
}

const findMaterialByID = `-- name: FindMaterialByID :one
SELECT id, name, created_at, last_updated FROM materials
WHERE id = $1
`

func (q *Queries) FindMaterialByID(ctx context.Context, id uuid.UUID) (Material, error) {
	row := q.db.QueryRowContext(ctx, findMaterialByID, id)
	var i Material
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatedAt,
		&i.LastUpdated,
	)
	return i, err
}

const findMaterialsBySoftName = `-- name: FindMaterialsBySoftName :many
SELECT id, name, created_at, last_updated FROM materials
WHERE name ILIKE '%' || $1 || '%'
LIMIT $2 OFFSET $3
`

type FindMaterialsBySoftNameParams struct {
	Column1 sql.NullString
	Limit   int32
	Offset  int32
}

func (q *Queries) FindMaterialsBySoftName(ctx context.Context, arg FindMaterialsBySoftNameParams) ([]Material, error) {
	rows, err := q.db.QueryContext(ctx, findMaterialsBySoftName, arg.Column1, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Material
	for rows.Next() {
		var i Material
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.CreatedAt,
			&i.LastUpdated,
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

const getMaterialByName = `-- name: GetMaterialByName :one
SELECT id, name, created_at, last_updated FROM materials
WHERE name = $1
`

func (q *Queries) GetMaterialByName(ctx context.Context, name string) (Material, error) {
	row := q.db.QueryRowContext(ctx, getMaterialByName, name)
	var i Material
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatedAt,
		&i.LastUpdated,
	)
	return i, err
}

const getMaterialCount = `-- name: GetMaterialCount :one
SELECT COUNT(*) FROM materials
`

func (q *Queries) GetMaterialCount(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, getMaterialCount)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const getMaterials = `-- name: GetMaterials :many
SELECT id, name, created_at, last_updated FROM materials
LIMIT $1 OFFSET $2
`

type GetMaterialsParams struct {
	Limit  int32
	Offset int32
}

func (q *Queries) GetMaterials(ctx context.Context, arg GetMaterialsParams) ([]Material, error) {
	rows, err := q.db.QueryContext(ctx, getMaterials, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Material
	for rows.Next() {
		var i Material
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.CreatedAt,
			&i.LastUpdated,
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

const getProductMaterials = `-- name: GetProductMaterials :many
SELECT m.id, m.name, m.created_at, m.last_updated FROM materials m
    INNER JOIN product_materials pm ON m.id = pm.material_id
WHERE pm.product_id = $1
LIMIT $2 OFFSET $3
`

type GetProductMaterialsParams struct {
	ProductID uuid.UUID
	Limit     int32
	Offset    int32
}

func (q *Queries) GetProductMaterials(ctx context.Context, arg GetProductMaterialsParams) ([]Material, error) {
	rows, err := q.db.QueryContext(ctx, getProductMaterials, arg.ProductID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Material
	for rows.Next() {
		var i Material
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.CreatedAt,
			&i.LastUpdated,
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

const updateMaterial = `-- name: UpdateMaterial :one
UPDATE materials SET
    name = $2,
    last_updated = NOW()
WHERE id = $1
RETURNING id, name, created_at, last_updated
`

type UpdateMaterialParams struct {
	ID   uuid.UUID
	Name string
}

func (q *Queries) UpdateMaterial(ctx context.Context, arg UpdateMaterialParams) (Material, error) {
	row := q.db.QueryRowContext(ctx, updateMaterial, arg.ID, arg.Name)
	var i Material
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.CreatedAt,
		&i.LastUpdated,
	)
	return i, err
}
