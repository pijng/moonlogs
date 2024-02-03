// Package qrx provides a lightweight query builder and executor for SQLite databases in Go.
//
// Usage:
// 1. Create an `Executor` or `Querier` based on your needs with the `With` method.
// 2. Use the query builder methods to build SQL queries easily.
//
// Example:
// ```go
// db, err := sql.Open("sqlite3", "example.db")
// if err != nil {
//     log.Fatal(err)
// }
//
// executor := qrx.With(db)
//
// // Example 1: Executing a simple SQL statement
// result, err := executor.Exec("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, name TEXT);")
// if err != nil {
//     log.Fatal(err)
// }
//
// // Example 2: Creating a Querier for a specific data structure
// type User struct {
//     ID   int
//     Name string
// }
//
// querier := qrx.Scan(User{}).With(db).From("users").Where("name = ?", "John")
//
// // Example 3: Executing a query and retrieving the first result
// user, err := querier.First()
// if err != nil {
//     log.Fatal(err)
// }
//
// fmt.Println(user)

package qrx

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"math"
	"reflect"
	"strings"
	"sync"
	"time"
)

const (
	LOWEST_TIME = math.MinInt
	MAX_TIME    = math.MaxInt
)

// Executor is a struct for executing SQL queries.
type Executor struct {
	db *sql.DB
}

// With creates an Executor with the given database connection.
func With(db *sql.DB) *Executor {
	return &Executor{db: db}
}

// Exec executes an SQL statement with context and returns the result.
func (eq *Executor) Exec(ctx context.Context, stmt string, args ...any) (sql.Result, error) {
	query := putSemicolon(stmt)
	return eq.db.ExecContext(ctx, query, args...)
}

// Querier is a struct for building queries on a specific data structure.
type Querier[T any] struct {
	structure T
}

// DBQuerier is a struct for building queries with a specific data structure and database connection.
type DBQuerier[T any] struct {
	db        *sql.DB
	structure T
}

// TableQuerier is a struct for building queries with a specific data structure, database connection, and table name.
type TableQuerier[T any] struct {
	db        *sql.DB
	structure T
	tableName string
}

var stmtsCache struct {
	sync.RWMutex
	stmts map[string]*sql.Stmt
}

// StmtQuerier is a struct for building queries with a specific data structure, database connection, query, and arguments.
type StmtQuerier[T any] struct {
	db        *sql.DB
	structure T
	tableName string
	query     string
	stmt      string
	args      []any
}

// Scan creates a Querier for building queries on a specific data structure.
func Scan[T any](structure T) *Querier[T] {
	return &Querier[T]{structure: structure}
}

// With creates a DBQuerier with the given database connection.
func (q *Querier[T]) With(db *sql.DB) *DBQuerier[T] {
	return &DBQuerier[T]{db: db, structure: q.structure}
}

// From creates a TableQuerier with the given table name.
func (dq *DBQuerier[T]) From(tableName string) *TableQuerier[T] {
	return &TableQuerier[T]{db: dq.db, structure: dq.structure, tableName: tableName}
}

// Exec creates a StmtQuerier for executing a custom SQL statement.
func (dq *DBQuerier[T]) Exec(stmt string, args ...any) *StmtQuerier[T] {
	return &StmtQuerier[T]{db: dq.db, structure: dq.structure, query: stmt, args: args}
}

// Create inserts data into the table with context and returns the inserted data.
func (tq *TableQuerier[T]) Create(ctx context.Context, data map[string]interface{}) (*T, error) {
	return tq.create(ctx, data)
}

// create is a helper function for inserting data into the table.
func (tq *TableQuerier[T]) create(ctx context.Context, data map[string]interface{}) (*T, error) {
	var columns []string
	var placeholders []string
	var values []interface{}

	for column, value := range data {
		columns = append(columns, column)
		placeholders = append(placeholders, "?")
		values = append(values, value)
	}

	columnsText := strings.Join(columns, ", ")
	placeholdersText := strings.Join(placeholders, ", ")

	query := putSemicolon(fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s)", tq.tableName, columnsText, placeholdersText))

	result, err := tq.db.ExecContext(ctx, query, values...)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	structure, err := tq.Where("id = ?", id).First(ctx)
	if err != nil {
		return nil, err
	}

	return structure, nil
}

// Where creates a StmtQuerier for querying with a WHERE clause.
func (tq *TableQuerier[T]) Where(stmt string, args ...any) *StmtQuerier[T] {
	query := fmt.Sprintf("SELECT * FROM %s WHERE %s", tq.tableName, stmt)
	return &StmtQuerier[T]{db: tq.db, structure: tq.structure, tableName: tq.tableName, query: query, stmt: stmt, args: args}
}

// All retrieves all records from the table with context.
func (tq *TableQuerier[T]) All(ctx context.Context, stmt string, args ...any) ([]*T, error) {
	query := fmt.Sprintf("SELECT * FROM %s %s", tq.tableName, stmt)
	querier := &StmtQuerier[T]{db: tq.db, structure: tq.structure, tableName: tq.tableName, query: query, stmt: stmt, args: args}

	return querier.all(ctx)
}

// Count retrieves the count of records in the table with context.
func (tq *TableQuerier[T]) Count(ctx context.Context) (int, error) {
	query := fmt.Sprintf("SELECT COUNT(*) FROM %s", tq.tableName)
	querier := &StmtQuerier[int]{db: tq.db, structure: 0, tableName: tq.tableName, query: query}

	return querier.count(ctx)
}

// DeleteAll deletes all the records from the table based on a WHERE clause with context.
func (tq *TableQuerier[T]) DeleteAll(ctx context.Context, stmt string, args ...any) (sql.Result, error) {
	query := fmt.Sprintf("DELETE FROM %s WHERE %s", tq.tableName, stmt)
	querier := &StmtQuerier[int]{db: tq.db, structure: 0, tableName: tq.tableName, query: query, stmt: stmt, args: args}

	return querier.delete(ctx)
}

// DeleteOne deletes one record from the table based on a WHERE clause with context.
func (tq *TableQuerier[T]) DeleteOne(ctx context.Context, stmt string, args ...any) (sql.Result, error) {
	query := fmt.Sprintf("DELETE FROM %s WHERE %s", tq.tableName, stmt)
	querier := &StmtQuerier[int]{db: tq.db, structure: 0, tableName: tq.tableName, query: query, stmt: stmt, args: args}

	return querier.delete(ctx)
}

// CountWhere retrieves the count of records from the table based on the provided conditions with context.
func (tq *TableQuerier[T]) CountWhere(ctx context.Context, stmt string, args ...any) (int, error) {
	query := fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE %s", tq.tableName, stmt)
	querier := &StmtQuerier[int]{db: tq.db, structure: 0, tableName: tq.tableName, query: query, stmt: stmt, args: args}

	return querier.count(ctx)
}

// count retrieves the count of records from the database.
func (sq *StmtQuerier[T]) count(ctx context.Context) (int, error) {
	row := sq.db.QueryRowContext(ctx, sq.query, sq.args...)

	var count int
	err := row.Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

// delete deletes records from the database.
func (sq *StmtQuerier[T]) delete(ctx context.Context) (sql.Result, error) {
	query := putSemicolon(sq.query)

	return sq.db.ExecContext(ctx, query, sq.args...)
}

// First retrieves the first record from the query result with context.
func (sq StmtQuerier[T]) First(ctx context.Context) (*T, error) {
	return sq.first(ctx)
}

// first retrieves the first record from the query result.
func (sq StmtQuerier[T]) first(ctx context.Context) (*T, error) {
	limitQuery := fmt.Sprintf("%s LIMIT 1;", trimSemicolon(sq.query))

	rows, err := sq.db.QueryContext(ctx, limitQuery, sq.args...)
	if err != nil {
		return nil, err
	}

	dest := new(T)
	for rows.Next() {
		err := scanRow(rows, dest)
		if err != nil {
			return nil, fmt.Errorf("failed to Scan row: %w", err)
		}
	}

	return dest, nil
}

// UpdateAll updates all the records by WHERE clause and return the updated data.
func (sq *StmtQuerier[T]) UpdateAll(ctx context.Context, data map[string]interface{}) ([]*T, error) {
	return sq.update(ctx, data)
}

// UpdateOne update one record by WHERE clause and return the updated data.
func (sq *StmtQuerier[T]) UpdateOne(ctx context.Context, data map[string]interface{}) (*T, error) {
	structures, err := sq.update(ctx, data)
	if len(structures) == 0 || err != nil {
		return nil, err
	}

	return structures[0], nil
}

func (sq *StmtQuerier[T]) update(ctx context.Context, data map[string]interface{}) ([]*T, error) {
	var setColumns []string
	var values []any

	for column, value := range data {
		setColumns = append(setColumns, fmt.Sprintf("%s = ?", column))
		values = append(values, value)
	}

	values = append(values, sq.args...)

	setClause := strings.Join(setColumns, ", ")
	query := putSemicolon(fmt.Sprintf("UPDATE %s SET %s WHERE %s", sq.tableName, setClause, sq.stmt))

	_, err := sq.db.ExecContext(ctx, query, values...)
	if err != nil {
		return nil, err
	}

	tq := &TableQuerier[T]{db: sq.db, structure: sq.structure, tableName: sq.tableName}
	structures, err := tq.Where(sq.stmt, sq.args...).All(ctx)
	if err != nil {
		return nil, err
	}

	return structures, nil
}

// All retrieves all records from the query result with context.
func (sq *StmtQuerier[T]) All(ctx context.Context) ([]*T, error) {
	return sq.all(ctx)
}

// all retrieves all records from the database.
func (sq *StmtQuerier[T]) all(ctx context.Context) ([]*T, error) {
	query := putSemicolon(sq.query)

	stmt, err := sq.cachedStmt(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed calling all: %w", err)
	}

	rows, err := stmt.QueryContext(ctx, sq.args...)
	if err != nil {
		return nil, fmt.Errorf("failed to QueryContext: %w", err)
	}
	defer rows.Close()

	resultSlice := make([]*T, 0)

	for rows.Next() {
		dest := new(T)

		err := scanRow(rows, dest)
		if err != nil {
			return nil, fmt.Errorf("failed to Scan row: %w", err)
		}

		resultSlice = append(resultSlice, dest)
	}

	return resultSlice, nil
}

func (sq *StmtQuerier[T]) cachedStmt(ctx context.Context, query string) (*sql.Stmt, error) {
	if stmtsCache.stmts == nil {
		stmtsCache.stmts = make(map[string]*sql.Stmt)
	}

	stmtsCache.RLock()
	stmt, ok := stmtsCache.stmts[query]
	stmtsCache.RUnlock()

	if ok {
		return stmt, nil
	}

	stmt, err := sq.db.PrepareContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed preparing statement: %w", err)
	}

	stmtsCache.Lock()
	defer stmtsCache.Unlock()
	stmtsCache.stmts[query] = stmt

	return stmt, err
}

// trimSemicolon removes the trailing semicolon from a query.
func trimSemicolon(query string) string {
	if len(query) > 0 && query[len(query)-1] == ';' {
		return query[:len(query)-1]
	}

	return query
}

// putSemicolon adds a trailing semicolon to a query.
func putSemicolon(query string) string {
	if len(query) > 0 && query[len(query)-1] == ';' {
		return query
	}

	return fmt.Sprintf("%s;", query)
}

// Contains returns a formatted string for SQL LIKE queries with wildcard characters.
func Contains(value any) string {
	return fmt.Sprintf("%%%s%%", value)
}

func Between(from *time.Time, to *time.Time) string {
	fromPart := LOWEST_TIME
	if from != nil {
		fromPart = int(from.Unix())
	}

	toPart := MAX_TIME
	if to != nil {
		toPart = int(to.Unix())
	}

	return fmt.Sprintf("%d and %d", fromPart, toPart)
}

func In[T any](values []T) ([]string, []any) {
	placeholders := make([]string, len(values))
	args := make([]any, len(values))
	for i, value := range values {
		placeholders[i] = "?"
		args[i] = value
	}

	return placeholders, args
}

// MapLike generates a WHERE clause for querying based on a map of conditions.
func MapLike(query map[string]interface{}) string {
	if len(query) == 0 {
		return "1=1"
	}

	var placeholders []string

	for key, value := range query {
		placeholders = append(placeholders, fmt.Sprintf("json_extract(query, '$.%s') like '%s'", key, Contains(value)))
	}

	return strings.Join(placeholders, " AND ")
}

func scanRow(rows *sql.Rows, dest any) error {
	rv := reflect.ValueOf(dest)
	if rv.Kind() != reflect.Pointer || rv.IsNil() {
		return errors.New("dest must be a non-nil pointer")
	}

	elem := rv.Elem()
	if elem.Kind() != reflect.Struct {
		return errors.New("dest must point to a struct")
	}

	indexes := cachedFieldIndexes(reflect.TypeOf(dest).Elem())

	columns, err := rows.Columns()
	if err != nil {
		return fmt.Errorf("cannot fetch columns: %w", err)
	}

	var scanArgs []any
	for _, column := range columns {
		index, ok := indexes[column]
		if ok {
			// We have a column to field mapping, scan the value.
			field := elem.Field(index)
			scanArgs = append(scanArgs, field.Addr().Interface())
		} else {
			// Unassigned column, throw away the scanned value.
			var throwAway any
			scanArgs = append(scanArgs, &throwAway)
		}
	}

	return rows.Scan(scanArgs...)
}

// fieldIndexes returns a map of database column name to struct field index.
func fieldIndexes(structType reflect.Type) map[string]int {
	indexes := make(map[string]int)

	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)
		tag := field.Tag.Get("sql")
		if tag != "" {
			indexes[tag] = i
		} else {
			indexes[field.Name] = i
		}
	}

	return indexes
}

var fieldIndexesCache sync.Map // map[reflect.Type]map[string]int

// cachedFieldIndexes is like fieldIndexes, but cached per struct type.
func cachedFieldIndexes(structType reflect.Type) map[string]int {
	if f, ok := fieldIndexesCache.Load(structType); ok {
		return f.(map[string]int)
	}

	indexes := fieldIndexes(structType)
	fieldIndexesCache.Store(structType, indexes)

	return indexes
}

func CleanCachedStatements() {
	stmtsCache.RWMutex.Lock()
	defer stmtsCache.Unlock()

	clear(stmtsCache.stmts)
}
