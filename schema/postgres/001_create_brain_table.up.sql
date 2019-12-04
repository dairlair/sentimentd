CREATE TABLE brain (
    brain_id BIGSERIAL PRIMARY KEY
    , name TEXT NOT NULL
    , description TEXT NOT NULL
    , created_at   TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);