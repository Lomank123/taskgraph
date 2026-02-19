-- +goose up

CREATE TABLE IF NOT EXISTS tasks (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    status VARCHAR(50) NOT NULL,
    type VARCHAR(50) NOT NULL,
    retries INT NOT NULL DEFAULT 0,
    error TEXT,
    result TEXT,
    payload TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);


-- +goose down

DROP TABLE IF EXISTS tasks;
