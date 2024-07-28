-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp" SCHEMA public;

CREATE TABLE IF NOT EXISTS mock(
    id uuid DEFAULT public.uuid_generate_v4(),
    content json,
    created_at timestamp,
    updated_at timestamp,
    CONSTRAINT mock_pk PRIMARY KEY (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS mock;
-- +goose StatementEnd