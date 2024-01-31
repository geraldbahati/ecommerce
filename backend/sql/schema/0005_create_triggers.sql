-- +goose Up
CREATE TRIGGER update_categories_last_updated
    BEFORE UPDATE ON categories
    FOR EACH ROW
    EXECUTE FUNCTION update_last_updated_column();

CREATE TRIGGER update_products_last_updated
    BEFORE UPDATE ON products
    FOR EACH ROW
    EXECUTE FUNCTION update_last_updated_column();

CREATE TRIGGER update_orders_last_updated
    BEFORE UPDATE ON orders
    FOR EACH ROW
    EXECUTE FUNCTION update_last_updated_column();

CREATE TRIGGER update_shopping_carts_last_updated
    BEFORE UPDATE ON shopping_carts
    FOR EACH ROW
    EXECUTE FUNCTION update_last_updated_column();

CREATE TRIGGER update_order_items_last_updated
    BEFORE UPDATE ON order_items
    FOR EACH ROW
    EXECUTE FUNCTION update_last_updated_column();

CREATE TRIGGER update_cart_items_last_updated
    BEFORE UPDATE ON cart_items
    FOR EACH ROW
    EXECUTE FUNCTION update_last_updated_column();

CREATE TRIGGER update_reviews_last_updated
    BEFORE UPDATE ON reviews
    FOR EACH ROW
    EXECUTE FUNCTION update_last_updated_column();

CREATE TRIGGER update_wishlists_last_updated
    BEFORE UPDATE ON wishlists
    FOR EACH ROW
    EXECUTE FUNCTION update_last_updated_column();

CREATE TRIGGER update_wishlist_items_last_updated
    BEFORE UPDATE ON wishlist_items
    FOR EACH ROW
    EXECUTE FUNCTION update_last_updated_column();

-- +goose Down
DROP TRIGGER update_categories_last_updated ON categories;
DROP TRIGGER update_products_last_updated ON products;
DROP TRIGGER update_orders_last_updated ON orders;
DROP TRIGGER update_shopping_carts_last_updated ON shopping_carts;
DROP TRIGGER update_order_items_last_updated ON order_items;
DROP TRIGGER update_cart_items_last_updated ON cart_items;
DROP TRIGGER update_reviews_last_updated ON reviews;
DROP TRIGGER update_wishlists_last_updated ON wishlists;
DROP TRIGGER update_wishlist_items_last_updated ON wishlist_items;