-- +goose up

CREATE TABLE IF NOT EXISTS tasks (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    status VARCHAR(50) NOT NULL,
    type VARCHAR(50) NOT NULL,
    retries INT NOT NULL DEFAULT 0,
    error JSONB,
    result JSONB,
    payload JSONB,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);

CREATE TABLE IF NOT EXISTS task_deps (
    task_id UUID NOT NULL,
    dep_id UUID NOT NULL,
    PRIMARY KEY (task_id, dep_id),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ
);

CREATE INDEX idx_task_deps_task_id ON task_deps(task_id);
CREATE INDEX idx_task_deps_dep_id ON task_deps(dep_id);


-- +goose down

DROP INDEX IF EXISTS idx_task_deps_task_id;
DROP INDEX IF EXISTS idx_task_deps_dep_id;

DROP TABLE IF EXISTS task_deps;
DROP TABLE IF EXISTS tasks;
