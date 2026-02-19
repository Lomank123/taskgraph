-- +goose up

CREATE TABLE IF NOT EXISTS tasks (
    id UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    status VARCHAR(50) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);


-- +goose down

DROP TABLE IF EXISTS tasks;
