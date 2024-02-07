-- name: CreateWishlist :one
INSERT INTO wishlists (id, user_id, name, visibility, created_at, last_updated)
VALUES (UUID_GENERATE_V4(), $1,'My Wishlist', 'private', NOW(), NOW())
RETURNING id, user_id, name, visibility, created_at, last_updated;

-- name: AddItemToWishlist :one
INSERT INTO wishlist_items (id, wishlist_id, product_id, priority, created_at, last_updated)
VALUES (UUID_GENERATE_V4(), $1, $2, 'medium', NOW(), NOW())
RETURNING id, wishlist_id, product_id, priority, created_at, last_updated;

-- name: RemoveItemFromWishlist :exec
DELETE FROM wishlist_items
WHERE wishlist_id = $1 AND product_id = $2;

-- name: ListAllItemsInUserWishlist :many
SELECT p.id, p.name, p.description, p.image_url, p.category_id
FROM products p
JOIN wishlist_items wi ON p.id = wi.product_id
JOIN wishlists w ON wi.wishlist_id = w.id
WHERE w.user_id = $1
ORDER BY wi.created_at DESC
LIMIT $2 OFFSET $3;

-- name: UpdateWishlist :one
UPDATE wishlists SET
    name = $2,
    visibility = $3,
    last_updated = NOW()
WHERE id = $1
RETURNING id, user_id, name, visibility, created_at, last_updated;

-- name: TrackInterestInWishlistItem :many
SELECT product_id, COUNT(*) AS interest_count
FROM wishlist_items
GROUP BY product_id;

-- name: FindCommonWishlistLists :many
SELECT product_id, COUNT(DISTINCT user_id) AS user_count
FROM wishlist_items wi
JOIN wishlists w ON wi.wishlist_id = w.id
GROUP BY product_id
HAVING user_count > 1;

-- name: DeleteWishlist :exec
DELETE FROM wishlists
WHERE id = $1;

-- name: CopyWishlistsIntoAnotherWishlist :exec
INSERT INTO wishlist_items (id, wishlist_id, product_id, priority, created_at, last_updated)
SELECT UUID_GENERATE_V4(), $2, product_id, 'medium', NOW(), NOW()
FROM wishlist_items
WHERE wishlist_id = $1
  AND NOT EXISTS (
    SELECT 1
    FROM wishlist_items wi2
    WHERE wi2.wishlist_id = $2
      AND wi2.product_id = wishlist_items.product_id
);
