-- migrate:up
CREATE TYPE currency_name AS ENUM ('EUR', 'USD');

CREATE TABLE currencies (
    id BIGSERIAL PRIMARY KEY,
    name currency_name NOT NULL,
    amount BIGINT,
    owner_id BIGINT NOT NULL REFERENCES users (id),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NULL,
    deleted_at TIMESTAMP WITH TIME ZONE DEFAULT NULL,
    UNIQUE (name, owner_id)
);

-- migrate:down
DROP TABLE IF EXISTS currencies;
DROP TYPE IF EXISTS currency_name;