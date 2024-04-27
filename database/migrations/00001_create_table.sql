-- +goose Up
-- +goose StatementBegin
CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "username" varchar unique,
  "password" varchar,
  "role" varchar,
  "created_at" timestamp
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

Drop table if exists users;

-- +goose StatementEnd
