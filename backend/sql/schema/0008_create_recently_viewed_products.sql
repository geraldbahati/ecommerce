-- +goose Up
CREATE TABLE recently_viewed_products (
    user_id UUID NOT NULL,
    product_id UUID NOT NULL,
    last_viewed_at TIMESTAMP WITH TIME ZONE NOT NULL,
    PRIMARY KEY (user_id, product_id),
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
    FOREIGN KEY (product_id) REFERENCES products (id) ON DELETE CASCADE
);

CREATE INDEX recently_viewed_products_user_id_idx ON recently_viewed_products (user_id);
CREATE INDEX recently_viewed_products_product_id_idx ON recently_viewed_products (product_id);
CREATE INDEX recently_viewed_products_last_viewed_at_idx ON recently_viewed_products (last_viewed_at);

-- +goose Down
DROP TABLE recently_viewed_products;