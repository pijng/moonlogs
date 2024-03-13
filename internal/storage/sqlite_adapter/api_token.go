package sqlite_adapter

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"moonlogs/internal/entities"
	"moonlogs/internal/lib/qrx"
	"moonlogs/internal/persistence"

	"github.com/newrelic/go-agent/v3/newrelic"
)

type ApiTokenStorage struct {
	ctx     context.Context
	readDB  *sql.DB
	writeDB *sql.DB
}

func NewApiTokenStorage(ctx context.Context) *ApiTokenStorage {
	return &ApiTokenStorage{
		ctx:     ctx,
		readDB:  persistence.SqliteReadDB(),
		writeDB: persistence.SqliteWriteDB(),
	}
}

func (s *ApiTokenStorage) CreateApiToken(apiToken entities.ApiToken) (*entities.ApiToken, error) {
	tx, err := qrx.BeginImmediate(s.writeDB)
	if err != nil {
		return nil, fmt.Errorf("failed to start transaction: %w", err)
	}

	query := "INSERT INTO api_tokens (name, token, token_digest, is_revoked) VALUES (?,?,?,?);"

	result, err := tx.ExecContext(s.ctx, query, apiToken.Name, "", apiToken.TokenDigest, apiToken.IsRevoked)
	if err != nil {
		return nil, fmt.Errorf("failed inserting api_token: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("failed retrieving api_token last insert id: %w", err)
	}

	return s.GetApiTokenByID(int(id))
}

func (s *ApiTokenStorage) UpdateApiTokenByID(id int, apiToken entities.ApiToken) (*entities.ApiToken, error) {
	tx, err := qrx.BeginImmediate(s.writeDB)
	if err != nil {
		return nil, fmt.Errorf("failed to start transaction: %w", err)
	}
	defer func(tx *sql.Tx) {
		_ = tx.Commit()
	}(tx)

	query := "UPDATE api_tokens SET name=?, is_revoked=? WHERE id=?;"

	_, err = tx.ExecContext(s.ctx, query, apiToken.Name, apiToken.IsRevoked, id)
	if err != nil {
		return nil, fmt.Errorf("failed updating api token: %w", err)
	}

	return s.GetApiTokenByID(id)
}

func (s *ApiTokenStorage) GetApiTokenByID(id int) (*entities.ApiToken, error) {
	query := "SELECT * FROM api_tokens WHERE id=? LIMIT 1;"
	stmt, err := s.readDB.PrepareContext(s.ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed preparing statement: %w", err)
	}
	defer stmt.Close()

	row := stmt.QueryRowContext(s.ctx, id)

	var t entities.ApiToken
	err = row.Scan(&t.ID, &t.Token, &t.TokenDigest, &t.Name, &t.IsRevoked)
	if err != nil {
		return nil, fmt.Errorf("failed scanning api_token: %w", err)
	}

	return &t, nil
}

func (s *ApiTokenStorage) GetApiTokenByDigest(digest string) (*entities.ApiToken, error) {
	txn := newrelic.FromContext(s.ctx)
	defer txn.StartSegment("storage.sqlite_adapter.GetApiTokenByDigest").End()

	query := "SELECT * FROM api_tokens WHERE token_digest=? LIMIT 1;"
	stmt, err := s.readDB.PrepareContext(s.ctx, query)
	if err != nil {
		return &entities.ApiToken{}, fmt.Errorf("failed preparing statement: %w", err)
	}
	defer stmt.Close()

	txnQueryApiToken := txn.StartSegment("storage.sqlite_adapter.GetApiTokenByDigest#QueryRowContext")
	row := stmt.QueryRowContext(s.ctx, digest)
	txnQueryApiToken.End()

	txnScanApiToken := txn.StartSegment("storage.sqlite_adapter.GetApiTokenByDigest#ScanApiToken")
	var t entities.ApiToken
	err = row.Scan(&t.ID, &t.Token, &t.TokenDigest, &t.Name, &t.IsRevoked)
	if errors.Is(err, sql.ErrNoRows) {
		return &entities.ApiToken{}, nil
	}
	txnScanApiToken.End()

	if err != nil {
		return &entities.ApiToken{}, fmt.Errorf("failed scanning api_token: %w", err)
	}

	return &t, nil
}

func (s *ApiTokenStorage) GetAllApiTokens() ([]*entities.ApiToken, error) {
	query := "SELECT * FROM api_tokens ORDER BY id DESC;"
	stmt, err := s.readDB.PrepareContext(s.ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed preparing statement: %w", err)
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(s.ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying api_tokens: %w", err)
	}
	defer rows.Close()

	tokens := make([]*entities.ApiToken, 0)

	for rows.Next() {
		var dest entities.ApiToken

		err := rows.Scan(&dest.ID, &dest.Token, &dest.TokenDigest, &dest.Name, &dest.IsRevoked)
		if err != nil {
			return nil, fmt.Errorf("failed querying api_token: %w", err)
		}

		tokens = append(tokens, &dest)
	}

	return tokens, nil
}

func (s *ApiTokenStorage) DeleteApiTokenByID(id int) error {
	tx, err := qrx.BeginImmediate(s.writeDB)
	if err != nil {
		return fmt.Errorf("failed to start transaction: %w", err)
	}
	defer func(tx *sql.Tx) {
		_ = tx.Commit()
	}(tx)

	query := "DELETE FROM api_tokens WHERE id=?;"

	_, err = tx.ExecContext(s.ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed deleting api token: %w", err)
	}

	return err
}
