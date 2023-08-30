// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"moonlogs/ent/logrecord"
	"moonlogs/ent/logschema"
	"moonlogs/ent/predicate"
	"moonlogs/ent/schema"
	"sync"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeLogRecord = "LogRecord"
	TypeLogSchema = "LogSchema"
)

// LogRecordMutation represents an operation that mutates the LogRecord nodes in the graph.
type LogRecordMutation struct {
	config
	op            Op
	typ           string
	id            *int
	text          *string
	created_at    *time.Time
	schema_name   *string
	schema_id     *int
	addschema_id  *int
	query         *schema.Query
	group_hash    *string
	level         *string
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*LogRecord, error)
	predicates    []predicate.LogRecord
}

var _ ent.Mutation = (*LogRecordMutation)(nil)

// logrecordOption allows management of the mutation configuration using functional options.
type logrecordOption func(*LogRecordMutation)

// newLogRecordMutation creates new mutation for the LogRecord entity.
func newLogRecordMutation(c config, op Op, opts ...logrecordOption) *LogRecordMutation {
	m := &LogRecordMutation{
		config:        c,
		op:            op,
		typ:           TypeLogRecord,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withLogRecordID sets the ID field of the mutation.
func withLogRecordID(id int) logrecordOption {
	return func(m *LogRecordMutation) {
		var (
			err   error
			once  sync.Once
			value *LogRecord
		)
		m.oldValue = func(ctx context.Context) (*LogRecord, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().LogRecord.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withLogRecord sets the old LogRecord of the mutation.
func withLogRecord(node *LogRecord) logrecordOption {
	return func(m *LogRecordMutation) {
		m.oldValue = func(context.Context) (*LogRecord, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m LogRecordMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m LogRecordMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *LogRecordMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *LogRecordMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().LogRecord.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetText sets the "text" field.
func (m *LogRecordMutation) SetText(s string) {
	m.text = &s
}

// Text returns the value of the "text" field in the mutation.
func (m *LogRecordMutation) Text() (r string, exists bool) {
	v := m.text
	if v == nil {
		return
	}
	return *v, true
}

// OldText returns the old "text" field's value of the LogRecord entity.
// If the LogRecord object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *LogRecordMutation) OldText(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldText is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldText requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldText: %w", err)
	}
	return oldValue.Text, nil
}

// ResetText resets all changes to the "text" field.
func (m *LogRecordMutation) ResetText() {
	m.text = nil
}

// SetCreatedAt sets the "created_at" field.
func (m *LogRecordMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *LogRecordMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the LogRecord entity.
// If the LogRecord object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *LogRecordMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *LogRecordMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetSchemaName sets the "schema_name" field.
func (m *LogRecordMutation) SetSchemaName(s string) {
	m.schema_name = &s
}

// SchemaName returns the value of the "schema_name" field in the mutation.
func (m *LogRecordMutation) SchemaName() (r string, exists bool) {
	v := m.schema_name
	if v == nil {
		return
	}
	return *v, true
}

// OldSchemaName returns the old "schema_name" field's value of the LogRecord entity.
// If the LogRecord object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *LogRecordMutation) OldSchemaName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldSchemaName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldSchemaName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldSchemaName: %w", err)
	}
	return oldValue.SchemaName, nil
}

// ResetSchemaName resets all changes to the "schema_name" field.
func (m *LogRecordMutation) ResetSchemaName() {
	m.schema_name = nil
}

// SetSchemaID sets the "schema_id" field.
func (m *LogRecordMutation) SetSchemaID(i int) {
	m.schema_id = &i
	m.addschema_id = nil
}

// SchemaID returns the value of the "schema_id" field in the mutation.
func (m *LogRecordMutation) SchemaID() (r int, exists bool) {
	v := m.schema_id
	if v == nil {
		return
	}
	return *v, true
}

// OldSchemaID returns the old "schema_id" field's value of the LogRecord entity.
// If the LogRecord object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *LogRecordMutation) OldSchemaID(ctx context.Context) (v int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldSchemaID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldSchemaID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldSchemaID: %w", err)
	}
	return oldValue.SchemaID, nil
}

// AddSchemaID adds i to the "schema_id" field.
func (m *LogRecordMutation) AddSchemaID(i int) {
	if m.addschema_id != nil {
		*m.addschema_id += i
	} else {
		m.addschema_id = &i
	}
}

// AddedSchemaID returns the value that was added to the "schema_id" field in this mutation.
func (m *LogRecordMutation) AddedSchemaID() (r int, exists bool) {
	v := m.addschema_id
	if v == nil {
		return
	}
	return *v, true
}

// ResetSchemaID resets all changes to the "schema_id" field.
func (m *LogRecordMutation) ResetSchemaID() {
	m.schema_id = nil
	m.addschema_id = nil
}

// SetQuery sets the "query" field.
func (m *LogRecordMutation) SetQuery(s schema.Query) {
	m.query = &s
}

// Query returns the value of the "query" field in the mutation.
func (m *LogRecordMutation) Query() (r schema.Query, exists bool) {
	v := m.query
	if v == nil {
		return
	}
	return *v, true
}

// OldQuery returns the old "query" field's value of the LogRecord entity.
// If the LogRecord object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *LogRecordMutation) OldQuery(ctx context.Context) (v schema.Query, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldQuery is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldQuery requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldQuery: %w", err)
	}
	return oldValue.Query, nil
}

// ResetQuery resets all changes to the "query" field.
func (m *LogRecordMutation) ResetQuery() {
	m.query = nil
}

// SetGroupHash sets the "group_hash" field.
func (m *LogRecordMutation) SetGroupHash(s string) {
	m.group_hash = &s
}

// GroupHash returns the value of the "group_hash" field in the mutation.
func (m *LogRecordMutation) GroupHash() (r string, exists bool) {
	v := m.group_hash
	if v == nil {
		return
	}
	return *v, true
}

// OldGroupHash returns the old "group_hash" field's value of the LogRecord entity.
// If the LogRecord object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *LogRecordMutation) OldGroupHash(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldGroupHash is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldGroupHash requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldGroupHash: %w", err)
	}
	return oldValue.GroupHash, nil
}

// ClearGroupHash clears the value of the "group_hash" field.
func (m *LogRecordMutation) ClearGroupHash() {
	m.group_hash = nil
	m.clearedFields[logrecord.FieldGroupHash] = struct{}{}
}

// GroupHashCleared returns if the "group_hash" field was cleared in this mutation.
func (m *LogRecordMutation) GroupHashCleared() bool {
	_, ok := m.clearedFields[logrecord.FieldGroupHash]
	return ok
}

// ResetGroupHash resets all changes to the "group_hash" field.
func (m *LogRecordMutation) ResetGroupHash() {
	m.group_hash = nil
	delete(m.clearedFields, logrecord.FieldGroupHash)
}

// SetLevel sets the "level" field.
func (m *LogRecordMutation) SetLevel(s string) {
	m.level = &s
}

// Level returns the value of the "level" field in the mutation.
func (m *LogRecordMutation) Level() (r string, exists bool) {
	v := m.level
	if v == nil {
		return
	}
	return *v, true
}

// OldLevel returns the old "level" field's value of the LogRecord entity.
// If the LogRecord object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *LogRecordMutation) OldLevel(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldLevel is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldLevel requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldLevel: %w", err)
	}
	return oldValue.Level, nil
}

// ResetLevel resets all changes to the "level" field.
func (m *LogRecordMutation) ResetLevel() {
	m.level = nil
}

// Where appends a list predicates to the LogRecordMutation builder.
func (m *LogRecordMutation) Where(ps ...predicate.LogRecord) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the LogRecordMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *LogRecordMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.LogRecord, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *LogRecordMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *LogRecordMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (LogRecord).
func (m *LogRecordMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *LogRecordMutation) Fields() []string {
	fields := make([]string, 0, 7)
	if m.text != nil {
		fields = append(fields, logrecord.FieldText)
	}
	if m.created_at != nil {
		fields = append(fields, logrecord.FieldCreatedAt)
	}
	if m.schema_name != nil {
		fields = append(fields, logrecord.FieldSchemaName)
	}
	if m.schema_id != nil {
		fields = append(fields, logrecord.FieldSchemaID)
	}
	if m.query != nil {
		fields = append(fields, logrecord.FieldQuery)
	}
	if m.group_hash != nil {
		fields = append(fields, logrecord.FieldGroupHash)
	}
	if m.level != nil {
		fields = append(fields, logrecord.FieldLevel)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *LogRecordMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case logrecord.FieldText:
		return m.Text()
	case logrecord.FieldCreatedAt:
		return m.CreatedAt()
	case logrecord.FieldSchemaName:
		return m.SchemaName()
	case logrecord.FieldSchemaID:
		return m.SchemaID()
	case logrecord.FieldQuery:
		return m.Query()
	case logrecord.FieldGroupHash:
		return m.GroupHash()
	case logrecord.FieldLevel:
		return m.Level()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *LogRecordMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case logrecord.FieldText:
		return m.OldText(ctx)
	case logrecord.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case logrecord.FieldSchemaName:
		return m.OldSchemaName(ctx)
	case logrecord.FieldSchemaID:
		return m.OldSchemaID(ctx)
	case logrecord.FieldQuery:
		return m.OldQuery(ctx)
	case logrecord.FieldGroupHash:
		return m.OldGroupHash(ctx)
	case logrecord.FieldLevel:
		return m.OldLevel(ctx)
	}
	return nil, fmt.Errorf("unknown LogRecord field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *LogRecordMutation) SetField(name string, value ent.Value) error {
	switch name {
	case logrecord.FieldText:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetText(v)
		return nil
	case logrecord.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case logrecord.FieldSchemaName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetSchemaName(v)
		return nil
	case logrecord.FieldSchemaID:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetSchemaID(v)
		return nil
	case logrecord.FieldQuery:
		v, ok := value.(schema.Query)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetQuery(v)
		return nil
	case logrecord.FieldGroupHash:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetGroupHash(v)
		return nil
	case logrecord.FieldLevel:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetLevel(v)
		return nil
	}
	return fmt.Errorf("unknown LogRecord field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *LogRecordMutation) AddedFields() []string {
	var fields []string
	if m.addschema_id != nil {
		fields = append(fields, logrecord.FieldSchemaID)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *LogRecordMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case logrecord.FieldSchemaID:
		return m.AddedSchemaID()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *LogRecordMutation) AddField(name string, value ent.Value) error {
	switch name {
	case logrecord.FieldSchemaID:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddSchemaID(v)
		return nil
	}
	return fmt.Errorf("unknown LogRecord numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *LogRecordMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(logrecord.FieldGroupHash) {
		fields = append(fields, logrecord.FieldGroupHash)
	}
	return fields
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *LogRecordMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *LogRecordMutation) ClearField(name string) error {
	switch name {
	case logrecord.FieldGroupHash:
		m.ClearGroupHash()
		return nil
	}
	return fmt.Errorf("unknown LogRecord nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *LogRecordMutation) ResetField(name string) error {
	switch name {
	case logrecord.FieldText:
		m.ResetText()
		return nil
	case logrecord.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case logrecord.FieldSchemaName:
		m.ResetSchemaName()
		return nil
	case logrecord.FieldSchemaID:
		m.ResetSchemaID()
		return nil
	case logrecord.FieldQuery:
		m.ResetQuery()
		return nil
	case logrecord.FieldGroupHash:
		m.ResetGroupHash()
		return nil
	case logrecord.FieldLevel:
		m.ResetLevel()
		return nil
	}
	return fmt.Errorf("unknown LogRecord field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *LogRecordMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *LogRecordMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *LogRecordMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *LogRecordMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *LogRecordMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *LogRecordMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *LogRecordMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown LogRecord unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *LogRecordMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown LogRecord edge %s", name)
}

// LogSchemaMutation represents an operation that mutates the LogSchema nodes in the graph.
type LogSchemaMutation struct {
	config
	op            Op
	typ           string
	id            *int
	title         *string
	description   *string
	name          *string
	fields        *[]schema.Field
	appendfields  []schema.Field
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*LogSchema, error)
	predicates    []predicate.LogSchema
}

var _ ent.Mutation = (*LogSchemaMutation)(nil)

// logschemaOption allows management of the mutation configuration using functional options.
type logschemaOption func(*LogSchemaMutation)

// newLogSchemaMutation creates new mutation for the LogSchema entity.
func newLogSchemaMutation(c config, op Op, opts ...logschemaOption) *LogSchemaMutation {
	m := &LogSchemaMutation{
		config:        c,
		op:            op,
		typ:           TypeLogSchema,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withLogSchemaID sets the ID field of the mutation.
func withLogSchemaID(id int) logschemaOption {
	return func(m *LogSchemaMutation) {
		var (
			err   error
			once  sync.Once
			value *LogSchema
		)
		m.oldValue = func(ctx context.Context) (*LogSchema, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().LogSchema.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withLogSchema sets the old LogSchema of the mutation.
func withLogSchema(node *LogSchema) logschemaOption {
	return func(m *LogSchemaMutation) {
		m.oldValue = func(context.Context) (*LogSchema, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m LogSchemaMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m LogSchemaMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *LogSchemaMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *LogSchemaMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().LogSchema.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetTitle sets the "title" field.
func (m *LogSchemaMutation) SetTitle(s string) {
	m.title = &s
}

// Title returns the value of the "title" field in the mutation.
func (m *LogSchemaMutation) Title() (r string, exists bool) {
	v := m.title
	if v == nil {
		return
	}
	return *v, true
}

// OldTitle returns the old "title" field's value of the LogSchema entity.
// If the LogSchema object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *LogSchemaMutation) OldTitle(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldTitle is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldTitle requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldTitle: %w", err)
	}
	return oldValue.Title, nil
}

// ResetTitle resets all changes to the "title" field.
func (m *LogSchemaMutation) ResetTitle() {
	m.title = nil
}

// SetDescription sets the "description" field.
func (m *LogSchemaMutation) SetDescription(s string) {
	m.description = &s
}

// Description returns the value of the "description" field in the mutation.
func (m *LogSchemaMutation) Description() (r string, exists bool) {
	v := m.description
	if v == nil {
		return
	}
	return *v, true
}

// OldDescription returns the old "description" field's value of the LogSchema entity.
// If the LogSchema object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *LogSchemaMutation) OldDescription(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldDescription is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldDescription requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldDescription: %w", err)
	}
	return oldValue.Description, nil
}

// ResetDescription resets all changes to the "description" field.
func (m *LogSchemaMutation) ResetDescription() {
	m.description = nil
}

// SetName sets the "name" field.
func (m *LogSchemaMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *LogSchemaMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the LogSchema entity.
// If the LogSchema object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *LogSchemaMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *LogSchemaMutation) ResetName() {
	m.name = nil
}

// SetFields sets the "fields" field.
func (m *LogSchemaMutation) SetFields(s []schema.Field) {
	m.fields = &s
	m.appendfields = nil
}

// GetFields returns the value of the "fields" field in the mutation.
func (m *LogSchemaMutation) GetFields() (r []schema.Field, exists bool) {
	v := m.fields
	if v == nil {
		return
	}
	return *v, true
}

// OldFields returns the old "fields" field's value of the LogSchema entity.
// If the LogSchema object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *LogSchemaMutation) OldFields(ctx context.Context) (v []schema.Field, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldFields is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldFields requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldFields: %w", err)
	}
	return oldValue.Fields, nil
}

// AppendFields adds s to the "fields" field.
func (m *LogSchemaMutation) AppendFields(s []schema.Field) {
	m.appendfields = append(m.appendfields, s...)
}

// AppendedFields returns the list of values that were appended to the "fields" field in this mutation.
func (m *LogSchemaMutation) AppendedFields() ([]schema.Field, bool) {
	if len(m.appendfields) == 0 {
		return nil, false
	}
	return m.appendfields, true
}

// ResetFields resets all changes to the "fields" field.
func (m *LogSchemaMutation) ResetFields() {
	m.fields = nil
	m.appendfields = nil
}

// Where appends a list predicates to the LogSchemaMutation builder.
func (m *LogSchemaMutation) Where(ps ...predicate.LogSchema) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the LogSchemaMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *LogSchemaMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.LogSchema, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *LogSchemaMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *LogSchemaMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (LogSchema).
func (m *LogSchemaMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *LogSchemaMutation) Fields() []string {
	fields := make([]string, 0, 4)
	if m.title != nil {
		fields = append(fields, logschema.FieldTitle)
	}
	if m.description != nil {
		fields = append(fields, logschema.FieldDescription)
	}
	if m.name != nil {
		fields = append(fields, logschema.FieldName)
	}
	if m.fields != nil {
		fields = append(fields, logschema.FieldFields)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *LogSchemaMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case logschema.FieldTitle:
		return m.Title()
	case logschema.FieldDescription:
		return m.Description()
	case logschema.FieldName:
		return m.Name()
	case logschema.FieldFields:
		return m.GetFields()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *LogSchemaMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case logschema.FieldTitle:
		return m.OldTitle(ctx)
	case logschema.FieldDescription:
		return m.OldDescription(ctx)
	case logschema.FieldName:
		return m.OldName(ctx)
	case logschema.FieldFields:
		return m.OldFields(ctx)
	}
	return nil, fmt.Errorf("unknown LogSchema field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *LogSchemaMutation) SetField(name string, value ent.Value) error {
	switch name {
	case logschema.FieldTitle:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetTitle(v)
		return nil
	case logschema.FieldDescription:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetDescription(v)
		return nil
	case logschema.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	case logschema.FieldFields:
		v, ok := value.([]schema.Field)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetFields(v)
		return nil
	}
	return fmt.Errorf("unknown LogSchema field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *LogSchemaMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *LogSchemaMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *LogSchemaMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown LogSchema numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *LogSchemaMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *LogSchemaMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *LogSchemaMutation) ClearField(name string) error {
	return fmt.Errorf("unknown LogSchema nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *LogSchemaMutation) ResetField(name string) error {
	switch name {
	case logschema.FieldTitle:
		m.ResetTitle()
		return nil
	case logschema.FieldDescription:
		m.ResetDescription()
		return nil
	case logschema.FieldName:
		m.ResetName()
		return nil
	case logschema.FieldFields:
		m.ResetFields()
		return nil
	}
	return fmt.Errorf("unknown LogSchema field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *LogSchemaMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *LogSchemaMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *LogSchemaMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *LogSchemaMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *LogSchemaMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *LogSchemaMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *LogSchemaMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown LogSchema unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *LogSchemaMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown LogSchema edge %s", name)
}
