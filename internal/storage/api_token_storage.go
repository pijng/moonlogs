package storage

import (
	"context"
	"moonlogs/internal/entities"
)

type ApiTokenStorage interface {
	CreateApiToken(ctx context.Context, apiToken entities.ApiToken) (*entities.ApiToken, error)
	DeleteApiTokenByID(ctx context.Context, id int) error
	GetAllApiTokens(ctx context.Context) ([]*entities.ApiToken, error)
	GetApiTokenByDigest(ctx context.Context, digest string) (*entities.ApiToken, error)
	GetApiTokenByID(ctx context.Context, id int) (*entities.ApiToken, error)
	UpdateApiTokenByID(ctx context.Context, id int, apiToken entities.ApiToken) (*entities.ApiToken, error)
}
