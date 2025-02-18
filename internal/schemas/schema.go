package schemas

const SQLITE_SCHEMA = `
CREATE TABLE IF NOT EXISTS records (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	text TEXT,
	created_at INTEGER NOT NULL,
	schema_name TEXT NOT NULL,
	schema_id INTEGER NOT NULL,
	query JSON,
	kind string,
	group_hash TEXT NOT NULL,
	level TEXT NOT NULL
);
CREATE TABLE IF NOT EXISTS schemas (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	title TEXT NOT NULL,
	description TEXT,
	name TEXT NOT NULL,
	fields JSON,
	kinds JSON,
	tag_id INTEGER,
	retention_days INTEGER
);
CREATE TABLE IF NOT EXISTS users (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL,
	email TEXT NOT NULL,
	password TEXT,
	password_digest TEXT NOT NULL,
	role TEXT NOT NULL,
	tag_ids TEXT,
	token TEXT,
	is_revoked INTEGER DEFAULT 0
);
CREATE TABLE IF NOT EXISTS api_tokens (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	token TEXT,
	token_digest TEXT NOT NULL,
	name TEXT NOT NULL,
	is_revoked INTEGER DEFAULT 0
);
CREATE TABLE IF NOT EXISTS tags (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL
);
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
CREATE TABLE IF NOT EXISTS incidents (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	rule_id INTEGER NOT NULL,
	keys JSON,
	count INTEGER NOT NULL,
	ttl INTEGER NOT NULL
);
CREATE TABLE IF NOT EXISTS notification_profiles (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT NOT NULL,
	description TEXT NOT NULL,
  rule_ids TEXT,
	enabled INTEGER DEFAULT 1,
	silence_for INTEGER,
	url TEXT,
  method TEXT,
  headers TEXT,
  payload text
);
`
