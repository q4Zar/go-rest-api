-- migrate:up
CREATE TABLE assets (
    id BIGSERIAL PRIMARY KEY,
    amount DOUBLE PRECISION NOT NULL,
    user_id BIGINT REFERENCES users (id),
    currency_id BIGINT REFERENCES currencies (id),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NULL,
    deleted_at TIMESTAMP WITH TIME ZONE DEFAULT NULL,
    Unique(user_id, currency_id)
);

-- migrate:down
DROP TABLE IF EXISTS assets;