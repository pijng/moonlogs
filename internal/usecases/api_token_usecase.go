package usecases

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"moonlogs/internal/entities"
	"moonlogs/internal/repositories"
)

var tokenHasher = sha256.New()

type ApiTokenUseCase struct {
	apiTokenRepository *repositories.ApiTokenRepository
}

func NewApiTokenUseCase(apiTokenRepository *repositories.ApiTokenRepository) *ApiTokenUseCase {
	return &ApiTokenUseCase{apiTokenRepository: apiTokenRepository}
}

func (uc *ApiTokenUseCase) CreateApiToken(name string) (*entities.ApiToken, error) {
	if name == "" {
		return nil, fmt.Errorf("failed creating api token: `name` attribute is required")
	}

	token, err := generateToken(32)
	if err != nil {
		return nil, fmt.Errorf("failed generating token string: %w", err)
	}

	tokenHash, err := hashToken(token)
	if err != nil {
		return nil, fmt.Errorf("failed hashing token: %w", err)
	}

	apiToken, err := uc.apiTokenRepository.CreateApiToken(entities.ApiToken{Name: name, TokenDigest: tokenHash, IsRevoked: false})
	if err != nil {
		return nil, fmt.Errorf("failed creating api token: %w", err)
	}

	// Assign the real token to the ApiToken entity so that we can show it in the UI
	// right after creation, but keep it secure afterwards.
	apiToken.Token = token

	return apiToken, nil
}

func (uc *ApiTokenUseCase) IsTokenValid(token string) (bool, error) {
	tokenHash, err := hashToken(token)
	if err != nil {
		return false, fmt.Errorf("failed hashing token: %w", err)
	}

	apiToken, err := uc.apiTokenRepository.GetApiTokenByDigest(tokenHash)
	if err != nil {
		return false, fmt.Errorf("failed querying token by digest: %w", err)
	}

	if apiToken.IsRevoked {
		return false, nil
	}

	apiTokenExist := apiToken.ID > 0

	return apiTokenExist, nil
}

func (uc *ApiTokenUseCase) GetAllApiTokens() ([]*entities.ApiToken, error) {
	return uc.apiTokenRepository.GetAllApiTokens()
}

func (uc *ApiTokenUseCase) DestroyApiTokenByID(id int) error {
	return uc.apiTokenRepository.DestroyApiTokenByID(id)
}

func (uc *ApiTokenUseCase) GetApiTokenByID(id int) (*entities.ApiToken, error) {
	return uc.apiTokenRepository.GetApiTokenByID(id)
}

func (uc *ApiTokenUseCase) UpdateApiTokenByID(id int, apiToken entities.ApiToken) (*entities.ApiToken, error) {
	return uc.apiTokenRepository.UpdateApiTokenByID(id, apiToken)
}

func generateToken(length int) (string, error) {
	numBytes := (length * 6) / 8

	randomBytes := make([]byte, numBytes)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}

	token := base64.URLEncoding.EncodeToString(randomBytes)

	token = token[:length]

	return token, nil
}

func hashToken(token string) (string, error) {
	_, err := tokenHasher.Write([]byte(token))
	if err != nil {
		return "", fmt.Errorf("failed hashing token: %w", err)
	}

	hashBytes := tokenHasher.Sum(nil)
	hashString := hex.EncodeToString(hashBytes)
	tokenHasher.Reset()

	return hashString, nil
}
