package repositories

import (
	"context"
	"fmt"
	"moonlogs/internal/entities"
	"moonlogs/internal/persistence"
	"moonlogs/lib/qrx"
)

type ApiTokenRepository struct {
	ctx    context.Context
	tokens *qrx.TableQuerier[entities.ApiToken]
}

func NewApiTokenRepository(ctx context.Context) *ApiTokenRepository {
	return &ApiTokenRepository{
		ctx:    ctx,
		tokens: qrx.Scan(entities.ApiToken{}).With(persistence.DB()).From("api_tokens"),
	}
}

func (r *ApiTokenRepository) CreateApiToken(apiToken entities.ApiToken) (*entities.ApiToken, error) {
	t, err := r.tokens.Create(r.ctx, map[string]interface{}{
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

func (r *ApiTokenRepository) UpdateApiTokenByID(id int, apiToken entities.ApiToken) (*entities.ApiToken, error) {
	t, err := r.tokens.Where("id = ?", id).UpdateOne(r.ctx, map[string]interface{}{
		"name":       apiToken.Name,
		"is_revoked": apiToken.IsRevoked,
	})

	if err != nil {
		return nil, fmt.Errorf("failed updating api token: %w", err)
	}

	return t, nil
}

func (r *ApiTokenRepository) GetApiTokenByID(id int) (*entities.ApiToken, error) {
	t, err := r.tokens.Where("id = ?", id).First(r.ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying api token: %w", err)
	}

	return t, nil
}

func (r *ApiTokenRepository) GetApiTokenByDigest(digest string) (*entities.ApiToken, error) {
	t, err := r.tokens.Where("token_digest = ?", digest).First(r.ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying api token: %w", err)
	}

	return t, nil
}

func (r *ApiTokenRepository) GetAllApiTokens() ([]*entities.ApiToken, error) {
	t, err := r.tokens.All(r.ctx, "ORDER BY id DESC")
	if err != nil {
		return nil, fmt.Errorf("failed querying api tokens: %w", err)
	}

	return t, nil
}

func (r *ApiTokenRepository) DestroyApiTokenByID(id int) error {
	_, err := r.tokens.DeleteOne(r.ctx, "id = ?", id)

	return err
}
