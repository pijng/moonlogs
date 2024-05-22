-- +goose Up
-- +goose StatementBegin
CREATE TABLE actions (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL,
  pattern TEXT NOT NULL,
  method TEXT NOT NULL,
  conditions JSON,
  schema_name TEXT NOT NULL,
  schema_id INTEGER NOT NULL,
  disabled INTEGER DEFAULT 0
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE actions;
-- +goose StatementEnd
