package sqlite_adapter

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"moonlogs/internal/entities"
	"moonlogs/internal/lib/qrx"
	"moonlogs/internal/storage"
)

type ApiTokenStorage struct {
	readDB  *sql.DB
	writeDB *sql.DB
}

func NewApiTokenStorage(readDB *sql.DB, writeDB *sql.DB) *ApiTokenStorage {
	return &ApiTokenStorage{
		readDB:  readDB,
		writeDB: writeDB,
	}
}

func (s *ApiTokenStorage) CreateApiToken(ctx context.Context, apiToken entities.ApiToken) (*entities.ApiToken, error) {
	tx, err := qrx.BeginImmediate(s.writeDB)
	if err != nil {
		return nil, fmt.Errorf("failed to start transaction: %w", err)
	}

	query := "INSERT INTO api_tokens (name, token, token_digest, is_revoked) VALUES (?,?,?,?);"

	result, err := tx.ExecContext(ctx, query, apiToken.Name, "", apiToken.TokenDigest, apiToken.IsRevoked)
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

	return s.GetApiTokenByID(ctx, int(id))
}

func (s *ApiTokenStorage) UpdateApiTokenByID(ctx context.Context, id int, apiToken entities.ApiToken) (*entities.ApiToken, error) {
	tx, err := qrx.BeginImmediate(s.writeDB)
	if err != nil {
		return nil, fmt.Errorf("failed to start transaction: %w", err)
	}

	query := "UPDATE api_tokens SET name=?, is_revoked=? WHERE id=?;"

	_, err = tx.ExecContext(ctx, query, apiToken.Name, apiToken.IsRevoked, id)
	if err != nil {
		return nil, fmt.Errorf("failed updating api token: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return s.GetApiTokenByID(ctx, id)
}

func (s *ApiTokenStorage) GetApiTokenByID(ctx context.Context, id int) (*entities.ApiToken, error) {
	query := "SELECT * FROM api_tokens WHERE id=? LIMIT 1;"
	stmt, err := s.readDB.PrepareContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed preparing statement: %w", err)
	}
	defer stmt.Close()

	row := stmt.QueryRowContext(ctx, id)

	var t entities.ApiToken
	err = row.Scan(&t.ID, &t.Token, &t.TokenDigest, &t.Name, &t.IsRevoked)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = storage.ErrNotFound
		}

		return nil, fmt.Errorf("failed scanning api_token: %w", err)
	}

	return &t, nil
}

func (s *ApiTokenStorage) GetApiTokenByDigest(ctx context.Context, digest string) (*entities.ApiToken, error) {
	query := "SELECT * FROM api_tokens WHERE token_digest=? LIMIT 1;"
	stmt, err := s.readDB.PrepareContext(ctx, query)
	if err != nil {
		return &entities.ApiToken{}, fmt.Errorf("failed preparing statement: %w", err)
	}
	defer stmt.Close()

	row := stmt.QueryRowContext(ctx, digest)

	var t entities.ApiToken
	err = row.Scan(&t.ID, &t.Token, &t.TokenDigest, &t.Name, &t.IsRevoked)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = storage.ErrNotFound
		}

		return nil, fmt.Errorf("failed scanning api token: %w", err)
	}
	return &t, nil
}

func (s *ApiTokenStorage) GetAllApiTokens(ctx context.Context) ([]*entities.ApiToken, error) {
	query := "SELECT * FROM api_tokens ORDER BY id DESC;"
	stmt, err := s.readDB.PrepareContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed preparing statement: %w", err)
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx)
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

func (s *ApiTokenStorage) DeleteApiTokenByID(ctx context.Context, id int) error {
	tx, err := qrx.BeginImmediate(s.writeDB)
	if err != nil {
		return fmt.Errorf("failed to start transaction: %w", err)
	}

	query := "DELETE FROM api_tokens WHERE id=?;"

	_, err = tx.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed deleting api token: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return err
}
