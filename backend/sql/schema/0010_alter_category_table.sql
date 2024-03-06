-- +goose Up
ALTER TABLE categories
ALTER COLUMN description TYPE VARCHAR(255);

-- +goose Down
ALTER TABLE categories
ALTER COLUMN description TYPE VARCHAR(100);