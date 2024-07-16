-- migrate:up
CREATE TYPE currency_name AS ENUM ('EUR', 'USD');

CREATE TABLE currencies (
    id BIGSERIAL PRIMARY KEY,
    name currency_name NOT NULL UNIQUE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NULL,
    deleted_at TIMESTAMP WITH TIME ZONE DEFAULT NULL
);

-- migrate:down
DROP TABLE IF EXISTS currencies;
DROP TYPE IF EXISTS currency_name;