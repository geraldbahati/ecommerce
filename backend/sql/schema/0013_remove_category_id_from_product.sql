-- +goose Up
ALTER TABLE products
    DROP COLUMN category_id;

-- +goose Down
ALTER TABLE products
    ADD COLUMN category_id UUID NOT NULL,