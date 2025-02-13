-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS notification_profiles (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL,
	description TEXT NOT NULL,
  rule_ids TEXT,
	enabled INTEGER DEFAULT 1,
	url TEXT,
  method TEXT,
  headers TEXT,
  payload text
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE notification_profiles;
-- +goose StatementEnd
