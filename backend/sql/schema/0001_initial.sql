-- +goose Up
CREATE TABLE users (
   id UUID PRIMARY KEY,
   username VARCHAR(50) NOT NULL,
   email VARCHAR(50) NOT NULL,
   hashed_password VARCHAR(60) NOT NULL,
   first_name VARCHAR(50) NOT NULL,
   last_name VARCHAR(50) NOT NULL,
   phone_number VARCHAR(20) NULL,
   date_of_birth DATE NULL,
   gender VARCHAR(10) NULL,
   shipping_address VARCHAR(100) NULL,
   billing_address VARCHAR(100) NULL,
   created_at TIMESTAMP NOT NULL DEFAULT NOW(),
   last_login TIMESTAMP NULL,
   account_status VARCHAR(10) NOT NULL DEFAULT 'active',
   user_role VARCHAR(10) NOT NULL DEFAULT 'customer',
   profile_picture VARCHAR(100) NULL,
   two_factor_auth BOOLEAN NOT NULL DEFAULT FALSE,
   UNIQUE (username),
   UNIQUE (email),
   CHECK (account_status IN ('active', 'inactive', 'suspended', 'deleted')),
   CHECK (user_role IN ('customer', 'admin', 'superadmin'))
);



CREATE TABLE categories (
    id UUID PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    description VARCHAR(100) NULL,
    image_url VARCHAR(100) NULL,
    SEO_keywords VARCHAR(100) NULL,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    last_updated TIMESTAMP NULL,
    UNIQUE (name)
);

CREATE TABLE products (
    id UUID PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    description VARCHAR(100) NULL,
    image_url VARCHAR(100) NULL,
    price DECIMAL(10, 2) NOT NULL,
    stock INT NOT NULL,
    category_id UUID NOT NULL,
    brand VARCHAR(50) NULL,
    rating DECIMAL(2, 1) NOT NULL DEFAULT 0.0,
    review_count INT NOT NULL DEFAULT 0,
    discount_rate DECIMAL(2, 1) NOT NULL DEFAULT 0.0,
    keywords VARCHAR(100) NULL,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    last_updated TIMESTAMP NULL,
    FOREIGN KEY (category_id) REFERENCES categories (id) ON DELETE CASCADE,
    UNIQUE (name)
);

CREATE TABLE orders (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL,
    order_status VARCHAR(10) NOT NULL DEFAULT 'pending',
    payment_status VARCHAR(10) NOT NULL DEFAULT 'pending',
    payment_method VARCHAR(10) NOT NULL DEFAULT 'cash',
    shipping_address VARCHAR(100) NOT NULL,
    billing_address VARCHAR(100) NOT NULL,
    total_price DECIMAL(10, 2) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    last_updated TIMESTAMP NULL,
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
    CHECK (order_status IN ('pending', 'processing', 'shipped', 'delivered', 'cancelled')),
    CHECK (payment_status IN ('pending', 'paid', 'failed')),
    CHECK (payment_method IN ('cash', 'credit_card', 'debit_card', 'paypal'))
);

CREATE TABLE shopping_carts (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    last_updated TIMESTAMP NULL,
    total_items INT NOT NULL DEFAULT 0,
    total_price DECIMAL(10, 2) NOT NULL DEFAULT 0.0,
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
    UNIQUE (user_id)
);

CREATE TABLE cart_items (
    id UUID PRIMARY KEY,
    shopping_cart_id UUID NOT NULL,
    product_id UUID NOT NULL,
    quantity INT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    last_updated TIMESTAMP NULL,
    FOREIGN KEY (shopping_cart_id) REFERENCES shopping_carts (id) ON DELETE CASCADE,
    FOREIGN KEY (product_id) REFERENCES products (id) ON DELETE CASCADE,
    UNIQUE (shopping_cart_id, product_id)
);

CREATE TABLE order_items (
    id UUID PRIMARY KEY,
    order_id UUID NOT NULL,
    product_id UUID NOT NULL,
    quantity INT NOT NULL,
    taxed_price DECIMAL(10, 2) NOT NULL DEFAULT 0.0,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    last_updated TIMESTAMP NULL,
    FOREIGN KEY (order_id) REFERENCES orders (id) ON DELETE CASCADE,
    FOREIGN KEY (product_id) REFERENCES products (id) ON DELETE CASCADE,
    UNIQUE (order_id, product_id)
);

CREATE TABLE reviews (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL,
    product_id UUID NOT NULL,
    rating DECIMAL(2, 1) NOT NULL,
    comment VARCHAR(100) NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    last_updated TIMESTAMP NULL,
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
    FOREIGN KEY (product_id) REFERENCES products (id) ON DELETE CASCADE,
    UNIQUE (user_id, product_id)
);

CREATE TABLE wishlists (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL,
    name VARCHAR(50) NOT NULL,
    visibility VARCHAR(10) NOT NULL DEFAULT 'private',
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    last_updated TIMESTAMP NULL,
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
    CHECK (visibility IN ('private', 'public')),
    UNIQUE (user_id, name)
);

CREATE TABLE wishlist_items (
    id UUID PRIMARY KEY,
    wishlist_id UUID NOT NULL,
    product_id UUID NOT NULL,
    priority VARCHAR(10) NOT NULL DEFAULT 'medium',
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    last_updated TIMESTAMP NULL,
    FOREIGN KEY (wishlist_id) REFERENCES wishlists (id) ON DELETE CASCADE,
    FOREIGN KEY (product_id) REFERENCES products (id) ON DELETE CASCADE,
    CHECK (priority IN ('low', 'medium', 'high')),
    UNIQUE (wishlist_id, product_id)
);

-- +goose Down
DROP TABLE users;
DROP TABLE categories;
DROP TABLE products;
DROP TABLE orders;
DROP TABLE shopping_carts;
DROP TABLE cart_items;
DROP TABLE order_items;
DROP TABLE reviews;
DROP TABLE wishlists;
DROP TABLE wishlist_items;