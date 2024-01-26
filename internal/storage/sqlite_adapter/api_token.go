package sqlite_adapter

import (
	"context"
	"fmt"
	"moonlogs/internal/entities"
	"moonlogs/internal/lib/qrx"
	"moonlogs/internal/persistence"
)

type ApiTokenStorage struct {
	ctx    context.Context
	tokens *qrx.TableQuerier[entities.ApiToken]
}

func NewApiTokenStorage(ctx context.Context) *ApiTokenStorage {
	return &ApiTokenStorage{
		ctx:    ctx,
		tokens: qrx.Scan(entities.ApiToken{}).With(persistence.DB()).From("api_tokens"),
	}
}

func (s *ApiTokenStorage) CreateApiToken(apiToken entities.ApiToken) (*entities.ApiToken, error) {
	t, err := s.tokens.Create(s.ctx, map[string]interface{}{
		"name":         apiToken.Name,
		"token":        "",
		"token_digest": apiToken.TokenDigest,
		"is_revoked":   apiToken.IsRevoked,
	})

	if err != nil {
		return nil, fmt.Errorf("failed creating api token: %w", err)
	}

	return t, nil
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
	t, err := s.tokens.Where("token_digest = ?", digest).First(s.ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying api token: %w", err)
	}

	return t, nil
}

func (s *ApiTokenStorage) GetAllApiTokens() ([]*entities.ApiToken, error) {
	t, err := s.tokens.All(s.ctx, "ORDER BY id DESC")
	if err != nil {
		return nil, fmt.Errorf("failed querying api tokens: %w", err)
	}

	return t, nil
}

func (s *ApiTokenStorage) DestroyApiTokenByID(id int) error {
	_, err := s.tokens.DeleteOne(s.ctx, "id = ?", id)

	return err
}
