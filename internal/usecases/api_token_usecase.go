package usecases

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"hash"
	"moonlogs/internal/entities"
	"moonlogs/internal/shared"
	"moonlogs/internal/storage"
	"sync"
)

var sha256HasherPool = sync.Pool{
	New: func() interface{} {
		return sha256.New()
	},
}

type ApiTokenUseCase struct {
	apiTokenStorage storage.ApiTokenStorage
}

func NewApiTokenUseCase(apiTokenStorage storage.ApiTokenStorage) *ApiTokenUseCase {
	return &ApiTokenUseCase{apiTokenStorage: apiTokenStorage}
}

func (uc *ApiTokenUseCase) CreateApiToken(ctx context.Context, name string) (*entities.ApiToken, error) {
	if name == "" {
		return nil, fmt.Errorf("failed creating api token: `name` attribute is required")
	}

	token, err := shared.GenerateRandomToken(32)
	if err != nil {
		return nil, fmt.Errorf("failed generating token string: %w", err)
	}

	tokenHash, err := hashToken(token)
	if err != nil {
		return nil, fmt.Errorf("failed hashing token: %w", err)
	}

	apiToken, err := uc.apiTokenStorage.CreateApiToken(ctx, entities.ApiToken{Name: name, TokenDigest: tokenHash, IsRevoked: false})
	if err != nil {
		return nil, fmt.Errorf("failed creating api token: %w", err)
	}

	// Assign the real token to the ApiToken entity so that we can show it in the UI
	// right after creation, but keep it secure afterwards.
	apiToken.Token = token

	return apiToken, nil
}

func (uc *ApiTokenUseCase) IsTokenValid(ctx context.Context, token string) (bool, error) {
	tokenHash, err := hashToken(token)
	if err != nil {
		return false, fmt.Errorf("failed hashing token: %w", err)
	}

	apiToken, err := uc.apiTokenStorage.GetApiTokenByDigest(ctx, tokenHash)
	if err != nil {
		return false, fmt.Errorf("failed querying token by digest: %w", err)
	}

	if apiToken.IsRevoked {
		return false, nil
	}

	apiTokenExist := apiToken.ID > 0

	return apiTokenExist, nil
}

func (uc *ApiTokenUseCase) GetAllApiTokens(ctx context.Context) ([]*entities.ApiToken, error) {
	return uc.apiTokenStorage.GetAllApiTokens(ctx)
}

func (uc *ApiTokenUseCase) DeleteApiTokenByID(ctx context.Context, id int) error {
	return uc.apiTokenStorage.DeleteApiTokenByID(ctx, id)
}

func (uc *ApiTokenUseCase) GetApiTokenByID(ctx context.Context, id int) (*entities.ApiToken, error) {
	return uc.apiTokenStorage.GetApiTokenByID(ctx, id)
}

func (uc *ApiTokenUseCase) UpdateApiTokenByID(ctx context.Context, id int, apiToken entities.ApiToken) (*entities.ApiToken, error) {
	return uc.apiTokenStorage.UpdateApiTokenByID(ctx, id, apiToken)
}

func hashToken(token string) (string, error) {
	tokenHasher := sha256HasherPool.Get().(hash.Hash)
	defer sha256HasherPool.Put(tokenHasher)

	_, err := tokenHasher.Write([]byte(token))
	if err != nil {
		return "", fmt.Errorf("failed hashing token: %w", err)
	}

	hashBytes := tokenHasher.Sum(nil)
	hashString := hex.EncodeToString(hashBytes)
	tokenHasher.Reset()

	return hashString, nil
}
