-- +goose Up
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_schema_id;
DROP INDEX IF EXISTS idx_schema_name;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
CREATE INDEX IF NOT EXISTS idx_schema_id ON records(schema_id);
CREATE INDEX IF NOT EXISTS idx_schema_name ON records(schema_name);
-- +goose StatementEnd
