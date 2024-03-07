-- +goose Up
ALTER TABLE sub_categories
ALTER COLUMN description TYPE VARCHAR(255);

-- +goose Down
ALTER TABLE sub_categories
ALTER COLUMN description TYPE VARCHAR(100);