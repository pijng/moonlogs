package mongodb_adapter

import (
	"context"
	"log"
	"moonlogs/internal/entities"
	"moonlogs/internal/testutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApiTokenStorage(t *testing.T) {
	ctx := context.Background()

	mongoC, client, err := testutil.SetupMongoContainer(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := testutil.TeardownMongoContainer(ctx, mongoC); err != nil {
			log.Fatal(err)
		}
	}()

	apiTokenStorage := NewApiTokenStorage(client.Database("test_moonlogs"))

	t.Run("CreateApiToken", func(t *testing.T) {
		apiToken := entities.ApiToken{
			Token:       "testtoken",
			TokenDigest: "digest123",
			Name:        "Test Token",
			IsRevoked:   false,
		}
		createdToken, err := apiTokenStorage.CreateApiToken(ctx, apiToken)
		assert.NoError(t, err)
		assert.NotNil(t, createdToken)
		assert.Equal(t, apiToken.Name, createdToken.Name)
	})

	t.Run("GetApiTokenByID", func(t *testing.T) {
		apiToken := entities.ApiToken{
			Token:       "testtokenByID",
			TokenDigest: "digest456",
			Name:        "Test Token By ID",
			IsRevoked:   false,
		}
		createdToken, err := apiTokenStorage.CreateApiToken(ctx, apiToken)
		assert.NoError(t, err)

		fetchedToken, err := apiTokenStorage.GetApiTokenByID(ctx, createdToken.ID)
		assert.NoError(t, err)
		assert.NotNil(t, fetchedToken)
		assert.Equal(t, createdToken.Name, fetchedToken.Name)
	})

	t.Run("UpdateApiTokenByID", func(t *testing.T) {
		apiToken := entities.ApiToken{
			Token:       "testtokenToUpdate",
			TokenDigest: "digest789",
			Name:        "Test Token To Update",
			IsRevoked:   false,
		}
		createdToken, err := apiTokenStorage.CreateApiToken(ctx, apiToken)
		assert.NoError(t, err)

		updatedData := entities.ApiToken{
			Name:      "Updated Token Name",
			IsRevoked: true,
		}
		updatedToken, err := apiTokenStorage.UpdateApiTokenByID(ctx, createdToken.ID, updatedData)
		assert.NoError(t, err)
		assert.NotNil(t, updatedToken)
		assert.Equal(t, updatedData.Name, updatedToken.Name)
		assert.Equal(t, updatedData.IsRevoked, updatedToken.IsRevoked)
	})

	t.Run("GetApiTokenByDigest", func(t *testing.T) {
		apiToken := entities.ApiToken{
			Token:       "testtokenByDigest",
			TokenDigest: "digest101112",
			Name:        "Test Token By Digest",
			IsRevoked:   false,
		}
		createdToken, err := apiTokenStorage.CreateApiToken(ctx, apiToken)
		assert.NoError(t, err)

		fetchedToken, err := apiTokenStorage.GetApiTokenByDigest(ctx, createdToken.TokenDigest)
		assert.NoError(t, err)
		assert.NotNil(t, fetchedToken)
		assert.Equal(t, createdToken.Name, fetchedToken.Name)
	})

	t.Run("GetAllApiTokens", func(t *testing.T) {
		apiTokens, err := apiTokenStorage.GetAllApiTokens(ctx)
		assert.NoError(t, err)
		assert.True(t, len(apiTokens) > 0)
	})

	t.Run("DeleteApiTokenByID", func(t *testing.T) {
		apiToken := entities.ApiToken{
			Token:       "testtokenToDelete",
			TokenDigest: "digest131415",
			Name:        "Test Token To Delete",
			IsRevoked:   false,
		}
		createdToken, err := apiTokenStorage.CreateApiToken(ctx, apiToken)
		assert.NoError(t, err)

		err = apiTokenStorage.DeleteApiTokenByID(ctx, createdToken.ID)
		assert.NoError(t, err)

		deletedToken, err := apiTokenStorage.GetApiTokenByID(ctx, createdToken.ID)
		assert.NoError(t, err)
		assert.Nil(t, deletedToken)
	})
}
