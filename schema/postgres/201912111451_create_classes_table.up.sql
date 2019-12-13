CREATE TABLE classes
(
    id         BIGSERIAL PRIMARY KEY,
    brain_id   BIGINT      NOT NULL,
    label      TEXT        NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMPTZ          DEFAULT NULL
);

ALTER TABLE classes ADD CONSTRAINT fk_classes_brains FOREIGN KEY (brain_id) REFERENCES brains (id);

CREATE INDEX non_deleted_classes ON classes USING BTREE (deleted_at);

CREATE UNIQUE INDEX unique_class_label_in_brain ON classes USING BTREE (brain_id, lower(label));

CREATE INDEX find_class_by_id ON classes USING HASH (id);