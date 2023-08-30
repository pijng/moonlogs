// Code generated by ent, DO NOT EDIT.

package ent

import (
	"moonlogs/ent/logrecord"
	"moonlogs/ent/logschema"
	"moonlogs/ent/schema"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	logrecordFields := schema.LogRecord{}.Fields()
	_ = logrecordFields
	// logrecordDescText is the schema descriptor for text field.
	logrecordDescText := logrecordFields[0].Descriptor()
	// logrecord.TextValidator is a validator for the "text" field. It is called by the builders before save.
	logrecord.TextValidator = logrecordDescText.Validators[0].(func(string) error)
	// logrecordDescCreatedAt is the schema descriptor for created_at field.
	logrecordDescCreatedAt := logrecordFields[1].Descriptor()
	// logrecord.DefaultCreatedAt holds the default value on creation for the created_at field.
	logrecord.DefaultCreatedAt = logrecordDescCreatedAt.Default.(func() time.Time)
	// logrecordDescSchemaName is the schema descriptor for schema_name field.
	logrecordDescSchemaName := logrecordFields[2].Descriptor()
	// logrecord.SchemaNameValidator is a validator for the "schema_name" field. It is called by the builders before save.
	logrecord.SchemaNameValidator = logrecordDescSchemaName.Validators[0].(func(string) error)
	// logrecordDescSchemaID is the schema descriptor for schema_id field.
	logrecordDescSchemaID := logrecordFields[3].Descriptor()
	// logrecord.SchemaIDValidator is a validator for the "schema_id" field. It is called by the builders before save.
	logrecord.SchemaIDValidator = logrecordDescSchemaID.Validators[0].(func(int) error)
	// logrecordDescLevel is the schema descriptor for level field.
	logrecordDescLevel := logrecordFields[6].Descriptor()
	// logrecord.DefaultLevel holds the default value on creation for the level field.
	logrecord.DefaultLevel = logrecordDescLevel.Default.(string)
	logschemaFields := schema.LogSchema{}.Fields()
	_ = logschemaFields
	// logschemaDescTitle is the schema descriptor for title field.
	logschemaDescTitle := logschemaFields[0].Descriptor()
	// logschema.TitleValidator is a validator for the "title" field. It is called by the builders before save.
	logschema.TitleValidator = logschemaDescTitle.Validators[0].(func(string) error)
	// logschemaDescName is the schema descriptor for name field.
	logschemaDescName := logschemaFields[2].Descriptor()
	// logschema.NameValidator is a validator for the "name" field. It is called by the builders before save.
	logschema.NameValidator = logschemaDescName.Validators[0].(func(string) error)
}
