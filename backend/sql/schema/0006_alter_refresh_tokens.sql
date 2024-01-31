-- +goose Up
ALTER TABLE refresh_tokens
    ALTER COLUMN created_at SET DEFAULT CURRENT_TIMESTAMP;

-- +goose Down
ALTER TABLE refresh_tokens
    ALTER COLUMN created_at DROP DEFAULT;
