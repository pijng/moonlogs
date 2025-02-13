-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS alerting_rules (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL,
	description TEXT NOT NULL,
	enabled INTEGER DEFAULT 1,
	severity TEXT NOT NULL,
	interval INTEGER NOT NULL,
	threshold INTEGER NOT NULL,
	condition TEXT NOT NULL,
	filter_level TEXT NOT NULL,
	filter_schema_ids TEXT,
	filter_schema_fields TEXT,
	filter_schema_kinds TEXT,
	aggregation_type TEXT NOT NULL,
	aggregation_group_by TEXT,
	aggregation_time_window INTEGER NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE alerting_rules;
-- +goose StatementEnd
