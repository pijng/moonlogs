package sqlite_adapter

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"moonlogs/internal/entities"
	"moonlogs/internal/lib/qrx"
	"moonlogs/internal/persistence"
)

type ApiTokenStorage struct {
	ctx    context.Context
	tokens *qrx.TableQuerier[entities.ApiToken]
	db     *sql.DB
}

func NewApiTokenStorage(ctx context.Context) *ApiTokenStorage {
	return &ApiTokenStorage{
		ctx:    ctx,
		tokens: qrx.Scan(entities.ApiToken{}).With(persistence.DB()).From("api_tokens"),
		db:     persistence.DB(),
	}
}

func (s *ApiTokenStorage) CreateApiToken(apiToken entities.ApiToken) (*entities.ApiToken, error) {
	query := "INSERT INTO api_tokens (name, token, token_digest, is_revoked) VALUES (?,?,?,?);"
	stmt, err := s.db.PrepareContext(s.ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed preparing statement: %w", err)
	}
	defer stmt.Close()

	result, err := stmt.ExecContext(s.ctx, apiToken.Name, "", apiToken.TokenDigest, apiToken.IsRevoked)
	if err != nil {
		return nil, fmt.Errorf("failed inserting api_token: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("failed retrieving api_token last insert id: %w", err)
	}

	query = "SELECT * FROM api_tokens WHERE id=? LIMIT 1;"
	stmt, err = s.db.PrepareContext(s.ctx, query)
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

func (s *ApiTokenStorage) UpdateApiTokenByID(id int, apiToken entities.ApiToken) (*entities.ApiToken, error) {
	t, err := s.tokens.Where("id = ?", id).UpdateOne(s.ctx, map[string]interface{}{
		"name":       apiToken.Name,
		"is_revoked": apiToken.IsRevoked,
	})

	if err != nil {
		return nil, fmt.Errorf("failed updating api token: %w", err)
	}

	return t, nil
}

func (s *ApiTokenStorage) GetApiTokenByID(id int) (*entities.ApiToken, error) {
	t, err := s.tokens.Where("id = ?", id).First(s.ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying api token: %w", err)
	}

	return t, nil
}

func (s *ApiTokenStorage) GetApiTokenByDigest(digest string) (*entities.ApiToken, error) {
	query := "SELECT * FROM api_tokens WHERE token_digest=? LIMIT 1;"
	stmt, err := s.db.PrepareContext(s.ctx, query)
	if err != nil {
		return &entities.ApiToken{}, fmt.Errorf("failed preparing statement: %w", err)
	}
	defer stmt.Close()

	row := stmt.QueryRowContext(s.ctx, digest)

	var t entities.ApiToken
	err = row.Scan(&t.ID, &t.Token, &t.TokenDigest, &t.Name, &t.IsRevoked)
	if errors.Is(err, sql.ErrNoRows) {
		return &entities.ApiToken{}, nil
	}

	if err != nil {
		return &entities.ApiToken{}, fmt.Errorf("failed scanning api_token: %w", err)
	}

	return &t, nil
}

func (s *ApiTokenStorage) GetAllApiTokens() ([]*entities.ApiToken, error) {
	t, err := s.tokens.All(s.ctx, "ORDER BY id DESC")
	if err != nil {
		return nil, fmt.Errorf("failed querying api tokens: %w", err)
	}

	return t, nil
}

func (s *ApiTokenStorage) DeleteApiTokenByID(id int) error {
	_, err := s.tokens.DeleteOne(s.ctx, "id = ?", id)

	return err
}
