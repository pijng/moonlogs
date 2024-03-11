package usecases

import (
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

var cachedTokens = make(map[string]bool)

type ApiTokenUseCase struct {
	apiTokenStorage storage.ApiTokenStorage
}

func NewApiTokenUseCase(apiTokenStorage storage.ApiTokenStorage) *ApiTokenUseCase {
	return &ApiTokenUseCase{apiTokenStorage: apiTokenStorage}
}

func (uc *ApiTokenUseCase) CreateApiToken(name string) (*entities.ApiToken, error) {
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

	apiToken, err := uc.apiTokenStorage.CreateApiToken(entities.ApiToken{Name: name, TokenDigest: tokenHash, IsRevoked: false})
	if err != nil {
		return nil, fmt.Errorf("failed creating api token: %w", err)
	}

	// Assign the real token to the ApiToken entity so that we can show it in the UI
	// right after creation, but keep it secure afterwards.
	apiToken.Token = token

	return apiToken, nil
}

func (uc *ApiTokenUseCase) IsTokenValid(token string) (bool, error) {
	valid, ok := cachedTokens[token]
	if ok && valid {
		return true, nil
	}

	tokenHash, err := hashToken(token)
	if err != nil {
		return false, fmt.Errorf("failed hashing token: %w", err)
	}

	apiToken, err := uc.apiTokenStorage.GetApiTokenByDigest(tokenHash)
	if err != nil {
		return false, fmt.Errorf("failed querying token by digest: %w", err)
	}

	if apiToken.IsRevoked {
		return false, nil
	}

	apiTokenExist := apiToken.ID > 0

	cachedTokens[token] = true

	return apiTokenExist, nil
}

func (uc *ApiTokenUseCase) GetAllApiTokens() ([]*entities.ApiToken, error) {
	return uc.apiTokenStorage.GetAllApiTokens()
}

func (uc *ApiTokenUseCase) DeleteApiTokenByID(id int) error {
	return uc.apiTokenStorage.DeleteApiTokenByID(id)
}

func (uc *ApiTokenUseCase) GetApiTokenByID(id int) (*entities.ApiToken, error) {
	return uc.apiTokenStorage.GetApiTokenByID(id)
}

func (uc *ApiTokenUseCase) UpdateApiTokenByID(id int, apiToken entities.ApiToken) (*entities.ApiToken, error) {
	return uc.apiTokenStorage.UpdateApiTokenByID(id, apiToken)
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
