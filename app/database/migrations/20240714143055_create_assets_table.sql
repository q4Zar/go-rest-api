-- migrate:up
CREATE TYPE asset_type AS ENUM ('EUR', 'USD');

CREATE TABLE assets (
    id BIGSERIAL PRIMARY KEY,
    balance DOUBLE PRECISION NOT NULL,
    user_id BIGINT REFERENCES users (id),
    asset_type VARCHAR(10) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NULL,
    deleted_at TIMESTAMP WITH TIME ZONE DEFAULT NULL,
    Unique(user_id, asset_type)
);

-- migrate:down
DROP TABLE IF EXISTS assets;
DROP TYPE IF EXISTS asset_type;