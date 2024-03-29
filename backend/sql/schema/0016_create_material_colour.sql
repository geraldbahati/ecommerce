-- +goose Up
CREATE TABLE materials (
    id UUID PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    last_updated TIMESTAMP NULL
);

CREATE TABLE product_materials (
    id UUID PRIMARY KEY,
    product_id UUID NOT NULL,
    material_id UUID NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    last_updated TIMESTAMP NULL,
    FOREIGN KEY (product_id) REFERENCES products (id) ON DELETE CASCADE,
    FOREIGN KEY (material_id) REFERENCES materials (id) ON DELETE CASCADE,
    UNIQUE (product_id, material_id)
);

CREATE TABLE colours (
    id UUID PRIMARY KEY,
    colour_hex VARCHAR(7) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    last_updated TIMESTAMP NULL
);

CREATE TABLE product_colours (
    id UUID PRIMARY KEY,
    product_id UUID NOT NULL,
    colour_id UUID NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    last_updated TIMESTAMP NULL,
    FOREIGN KEY (product_id) REFERENCES products (id) ON DELETE CASCADE,
    FOREIGN KEY (colour_id) REFERENCES colours (id) ON DELETE CASCADE,
    UNIQUE (product_id, colour_id)
);

-- Indexes
CREATE INDEX idx_product_materials_product_id ON product_materials (product_id);
CREATE INDEX idx_product_materials_material_id ON product_materials (material_id);
CREATE INDEX idx_product_colours_product_id ON product_colours (product_id);
CREATE INDEX idx_product_colours_colour_id ON product_colours (colour_id);

-- +goose Down
DROP TABLE product_colours;
DROP TABLE colours;
DROP TABLE product_materials;
DROP TABLE materials;
