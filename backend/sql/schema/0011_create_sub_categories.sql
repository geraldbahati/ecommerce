-- +goose Up
CREATE TABLE sub_categories (
    id UUID PRIMARY KEY,
    category_id UUID NOT NULL,
    name VARCHAR(50) NOT NULL,
    description VARCHAR(100) NULL,
    image_url VARCHAR(100) NULL,
    SEO_keywords VARCHAR(100) NULL,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    last_updated TIMESTAMP NULL,
    FOREIGN KEY (category_id) REFERENCES categories (id) ON DELETE CASCADE,
    UNIQUE (name)
);

-- +goose Down
DROP TABLE sub_categories;