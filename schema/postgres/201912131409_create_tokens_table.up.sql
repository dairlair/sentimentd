CREATE TABLE tokens
(
    id         BIGSERIAL PRIMARY KEY,
    brain_id   BIGINT      NOT NULL,
    text       TEXT        NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ          DEFAULT NULL
);

ALTER TABLE tokens ADD CONSTRAINT fk_tokens_brains FOREIGN KEY (brain_id) REFERENCES brains (id);

CREATE INDEX non_deleted_tokens ON tokens USING BTREE (deleted_at);

CREATE UNIQUE INDEX unique_tokens_text_in_brain ON tokens USING BTREE (brain_id, text);