-- ============================
-- Database initialization file
-- Compatible with GORM models
-- ============================

-- Create database if it does not exist.
DO
$$
BEGIN
    IF NOT EXISTS (
        SELECT FROM pg_database WHERE datname = 'erp'
    ) THEN
        PERFORM dblink_exec('dbname=postgres', 'CREATE DATABASE erp');
    END IF;
END
$$;

-- Switch to the target database
\c erp;

-- ============================
-- Create "users" table
-- ============================

CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) UNIQUE NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,

    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- ============================
-- Create "products" table
-- ============================

CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    category VARCHAR(255) NOT NULL,
    image_url TEXT NOT NULL,
    description TEXT,

    price INTEGER NOT NULL,
    stock INTEGER NOT NULL,

    colors TEXT[] NOT NULL,
    sizes  TEXT[] NOT NULL,

    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- ============================
-- Auto-update updated_at column on modifications
-- ============================

CREATE OR REPLACE FUNCTION update_timestamp()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- --- Users trigger ---
DROP TRIGGER IF EXISTS update_users_timestamp ON users;

CREATE TRIGGER update_users_timestamp
BEFORE UPDATE ON users
FOR EACH ROW
EXECUTE FUNCTION update_timestamp();

-- --- Products trigger ---
DROP TRIGGER IF EXISTS update_products_timestamp ON products;

CREATE TRIGGER update_products_timestamp
BEFORE UPDATE ON products
FOR EACH ROW
EXECUTE FUNCTION update_timestamp();
