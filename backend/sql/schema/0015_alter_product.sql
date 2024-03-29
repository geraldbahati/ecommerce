-- +goose Up
ALTER TABLE products
ALTER COLUMN description TYPE VARCHAR(255);

-- +goose Down
ALTER TABLE products
ALTER COLUMN description TYPE VARCHAR(100);