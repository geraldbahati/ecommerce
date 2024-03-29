-- +goose Up
CREATE INDEX idx_wishlist_items_wishlist_product ON wishlist_items(wishlist_id, product_id);

-- +goose Down
DROP INDEX idx_wishlist_items_wishlist_product;