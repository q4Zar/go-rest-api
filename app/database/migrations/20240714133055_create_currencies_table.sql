-- migrate:up
CREATE TABLE currencies (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(200) NOT NULL,
    amount BIGINT,
    owner_id BIGINT REFERENCES users (id),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NULL,
    deleted_at TIMESTAMP WITH TIME ZONE DEFAULT NULL
);

-- migrate:down
DROP TABLE IF EXISTS currencies;