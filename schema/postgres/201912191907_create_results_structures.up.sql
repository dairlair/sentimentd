CREATE TABLE trainings
(
    id            BIGSERIAL PRIMARY KEY,
    brain_id      BIGINT      NOT NULL,
    comment       TEXT                 DEFAULT NULL,
    samples_count BIGINT      NOT NULL,
    created_at    TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at    TIMESTAMPTZ          DEFAULT NULL
);

ALTER TABLE trainings
    ADD CONSTRAINT fk_trainings_brains FOREIGN KEY (brain_id) REFERENCES brains (id);

CREATE TABLE training_classes
(
    id            BIGSERIAL PRIMARY KEY,
    brain_id      BIGINT NOT NULL,
    training_id   BIGINT NOT NULL,
    class_id      BIGINT NOT NULL,
    samples_count BIGINT NOT NULL,
    created_at    TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at    TIMESTAMPTZ          DEFAULT NULL
);

ALTER TABLE training_classes
    ADD CONSTRAINT fk_training_classes_brain FOREIGN KEY (brain_id) REFERENCES brains (id);

ALTER TABLE training_classes
    ADD CONSTRAINT fk_training_classes_training FOREIGN KEY (training_id) REFERENCES trainings (id);

ALTER TABLE training_classes
    ADD CONSTRAINT fk_training_classes_class FOREIGN KEY (class_id) REFERENCES classes (id);

CREATE TABLE training_tokens
(
    id            BIGSERIAL PRIMARY KEY,
    brain_id      BIGINT NOT NULL,
    training_id   BIGINT NOT NULL,
    class_id      BIGINT NOT NULL,
    token_id      BIGINT NOT NULL,
    samples_count BIGINT NOT NULL,
    created_at    TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at    TIMESTAMPTZ          DEFAULT NULL
);

ALTER TABLE training_tokens
    ADD CONSTRAINT fk_training_tokens_brain FOREIGN KEY (brain_id) REFERENCES brains (id);

ALTER TABLE training_tokens
    ADD CONSTRAINT fk_training_tokens_training FOREIGN KEY (training_id) REFERENCES trainings (id);

ALTER TABLE training_tokens
    ADD CONSTRAINT fk_training_tokens_class FOREIGN KEY (class_id) REFERENCES classes (id);

ALTER TABLE training_tokens
    ADD CONSTRAINT fk_training_tokens_token FOREIGN KEY (token_id) REFERENCES tokens (id);