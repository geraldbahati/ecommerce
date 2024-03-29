-- +goose Up
ALTER TABLE products
    ADD COLUMN sub_category_id UUID NULL,
ADD CONSTRAINT fk_products_sub_categories FOREIGN KEY (sub_category_id) REFERENCES sub_categories(id) ON DELETE SET NULL;

-- +goose Down
ALTER TABLE products
    DROP COLUMN sub_category_id;