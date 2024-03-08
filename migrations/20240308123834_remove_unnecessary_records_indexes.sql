-- +goose Up
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_kind;
DROP INDEX IF EXISTS idx_level;
DROP INDEX IF EXISTS idx_group_hash;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
CREATE INDEX IF NOT EXISTS idx_kind ON records(kind);
CREATE INDEX IF NOT EXISTS idx_level ON records(level);
CREATE INDEX IF NOT EXISTS idx_group_hash ON records(group_hash);
-- +goose StatementEnd
