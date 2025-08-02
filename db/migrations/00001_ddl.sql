-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS pg_uuidv7;

CREATE TYPE category_type AS ENUM ('Smartphones', 'Laptops', 'Accessories', 'TV');
CREATE TYPE brand_type AS ENUM ('Apple', 'Samsung', 'Sony', 'Xiaomi', 'Other');

CREATE TABLE products (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v7(),
    name VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    available BOOLEAN DEFAULT TRUE,
    category category_type NOT NULL,
    brand brand_type NOT NULL,
    rating DECIMAL(2, 1) NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS products;

DROP TYPE IF EXISTS category_type;
DROP TYPE IF EXISTS brand_type;

DROP EXTENSION IF EXISTS pg_uuidv7;
-- +goose StatementEnd
