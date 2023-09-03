// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// LogRecordsColumns holds the columns for the "log_records" table.
	LogRecordsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "text", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "schema_name", Type: field.TypeString},
		{Name: "schema_id", Type: field.TypeInt},
		{Name: "query", Type: field.TypeJSON},
		{Name: "group_hash", Type: field.TypeString, Nullable: true},
		{Name: "level", Type: field.TypeString, Default: "Info"},
	}
	// LogRecordsTable holds the schema information for the "log_records" table.
	LogRecordsTable = &schema.Table{
		Name:       "log_records",
		Columns:    LogRecordsColumns,
		PrimaryKey: []*schema.Column{LogRecordsColumns[0]},
	}
	// LogSchemasColumns holds the columns for the "log_schemas" table.
	LogSchemasColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "title", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "fields", Type: field.TypeJSON},
	}
	// LogSchemasTable holds the schema information for the "log_schemas" table.
	LogSchemasTable = &schema.Table{
		Name:       "log_schemas",
		Columns:    LogSchemasColumns,
		PrimaryKey: []*schema.Column{LogSchemasColumns[0]},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "email", Type: field.TypeString},
		{Name: "password_digest", Type: field.TypeString},
		{Name: "role", Type: field.TypeString, Default: "Member"},
		{Name: "token", Type: field.TypeString, Nullable: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		LogRecordsTable,
		LogSchemasTable,
		UsersTable,
	}
)

func init() {
}
