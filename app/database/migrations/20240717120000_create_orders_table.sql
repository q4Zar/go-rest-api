-- 20240717120000_create_orders_table.sql
-- Enum types for PostgreSQL
-- migrate:up
-- CREATE TYPE order_side_name AS ENUM ('BUY', 'SELL');
-- CREATE TYPE order_pair_name AS ENUM ('EUR-USD', 'USD-EUR');
-- CREATE TYPE order_status_name AS ENUM ('Pending', 'Filled');

-- Table Definition
CREATE TABLE orders (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT REFERENCES users (id),
    side VARCHAR(10) NOT NULL,
    amount FLOAT NOT NULL CHECK (amount > 0),
    price FLOAT NOT NULL CHECK (price > 0),
    asset_pair VARCHAR(10) NOT NULL,
    status VARCHAR(10) DEFAULT 'Pending' NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE NULL,
    deleted_at TIMESTAMP WITH TIME ZONE NULL
);


-- migrate:down
DROP TABLE IF EXISTS orders;
-- DROP TYPE IF EXISTS order_side_name;
-- DROP TYPE IF EXISTS order_pair_name;
-- DROP TYPE IF EXISTS order_status_name;