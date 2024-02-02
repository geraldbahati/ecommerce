-- +goose Up

-- Remove existing address columns
ALTER TABLE users
DROP COLUMN shipping_address,
DROP COLUMN billing_address;

-- Create a new table for shipping addresses
CREATE TABLE shipping_addresses (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL,
    street VARCHAR(255) NOT NULL,
    city VARCHAR(100) NOT NULL,
    state VARCHAR(100) NOT NULL,
    country VARCHAR(100) NOT NULL,
    postal_code VARCHAR(20) NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- Create a new table for billing addresses
CREATE TABLE billing_addresses (
   id UUID PRIMARY KEY,
   user_id UUID NOT NULL,
   street VARCHAR(255) NOT NULL,
   city VARCHAR(100) NOT NULL,
   state VARCHAR(100) NOT NULL,
   country VARCHAR(100) NOT NULL,
   postal_code VARCHAR(20) NOT NULL,
   FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- +goose Down

-- Add back the original address columns to the users table
ALTER TABLE users
ADD COLUMN shipping_address VARCHAR(100),
ADD COLUMN billing_address VARCHAR(100);

-- Drop the new address tables
DROP TABLE shipping_addresses;
DROP TABLE billing_addresses;
