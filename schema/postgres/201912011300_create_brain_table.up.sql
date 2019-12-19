CREATE TABLE brains
(
    id          BIGSERIAL PRIMARY KEY,
    name        TEXT        NOT NULL,
    description TEXT        NOT NULL,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at  TIMESTAMPTZ          DEFAULT NULL
);

CREATE INDEX non_deleted_brains ON brains USING BTREE (deleted_at);

CREATE INDEX find_brain_by_id ON brains USING HASH (id);

CREATE UNIQUE INDEX unique_brain_name ON brains USING BTREE (name);