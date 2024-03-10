-- +goose Up
-- +goose StatementBegin
CREATE INDEX IF NOT EXISTS idx_token_digest ON api_tokens (token_digest);
REINDEX idx_token_digest;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_token_digest;
-- +goose StatementEnd
