-- +goose Up
CREATE INDEX idx_categories_name ON categories(name);

CREATE INDEX idx_products_category_id ON products(category_id);
CREATE INDEX idx_products_name ON products(name);

CREATE INDEX idx_orders_user_id ON orders(user_id);
CREATE INDEX idx_orders_order_status ON orders(order_status);
CREATE INDEX idx_orders_payment_status ON orders(payment_status);

CREATE INDEX idx_shopping_carts_user_id ON shopping_carts(user_id);

CREATE INDEX idx_cart_items_shopping_cart_id ON cart_items(shopping_cart_id);
CREATE INDEX idx_cart_items_product_id ON cart_items(product_id);

CREATE INDEX idx_order_items_order_id ON order_items(order_id);
CREATE INDEX idx_order_items_product_id ON order_items(product_id);

CREATE INDEX idx_reviews_user_id ON reviews(user_id);
CREATE INDEX idx_reviews_product_id ON reviews(product_id);

CREATE INDEX idx_wishlists_user_id ON wishlists(user_id);

CREATE INDEX idx_wishlist_items_wishlist_id ON wishlist_items(wishlist_id);
CREATE INDEX idx_wishlist_items_product_id ON wishlist_items(product_id);

-- +goose Down
DROP INDEX idx_categories_name;

DROP INDEX idx_products_category_id;
DROP INDEX idx_products_name;

DROP INDEX idx_orders_user_id;
DROP INDEX idx_orders_order_status;
DROP INDEX idx_orders_payment_status;

DROP INDEX idx_shopping_carts_user_id;

DROP INDEX idx_cart_items_shopping_cart_id;
DROP INDEX idx_cart_items_product_id;

DROP INDEX idx_order_items_order_id;
DROP INDEX idx_order_items_product_id;

DROP INDEX idx_reviews_user_id;
DROP INDEX idx_reviews_product_id;

DROP INDEX idx_wishlists_user_id;

DROP INDEX idx_wishlist_items_wishlist_id;
DROP INDEX idx_wishlist_items_product_id;