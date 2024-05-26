package mongodb_adapter

import (
	"context"
	"log"
	"testing"

	"moonlogs/internal/entities"
	"moonlogs/internal/testutil"

	"github.com/stretchr/testify/assert"
)

func TestUserStorage(t *testing.T) {
	ctx := context.Background()
	mongoC, client, err := testutil.SetupMongoContainer(ctx)
	if err != nil {
		t.Fatalf("Failed to set up MongoDB container: %v", err)
	}
	defer func() {
		if err := testutil.TeardownMongoContainer(ctx, mongoC); err != nil {
			log.Fatalf("Failed to tear down MongoDB container: %v", err)
		}
	}()

	userStorage := &UserStorage{
		ctx:        ctx,
		client:     client,
		collection: client.Database("test_moonlogs").Collection("users"),
	}

	admin := entities.User{
		Name:           "Admin",
		Email:          "admin@example.com",
		PasswordDigest: "hashed_password",
		Role:           "Admin",
		IsRevoked:      false,
	}

	t.Run("CreateUser", func(t *testing.T) {
		user := entities.User{
			Name:           "Test User",
			Email:          "test@example.com",
			PasswordDigest: "hashed_password",
			Role:           "User",
			IsRevoked:      false,
		}
		createdUser, err := userStorage.CreateUser(user)
		assert.NoError(t, err)
		assert.NotNil(t, createdUser)
		assert.Equal(t, "Test User", createdUser.Name)
	})

	t.Run("GetUserByID", func(t *testing.T) {
		testUser, err := userStorage.GetUserByEmail("test@example.com")
		assert.NoError(t, err)
		assert.NotNil(t, testUser)

		foundUser, err := userStorage.GetUserByID(testUser.ID)
		assert.NoError(t, err)
		assert.NotNil(t, foundUser)
		assert.Equal(t, "Test User", foundUser.Name)
	})

	t.Run("GetUsersByTagID", func(t *testing.T) {
		users, err := userStorage.GetUsersByTagID(1)
		assert.NoError(t, err)
		assert.NotNil(t, users)
	})

	t.Run("GetUserByEmail", func(t *testing.T) {
		foundUser, err := userStorage.GetUserByEmail("test@example.com")
		assert.NoError(t, err)
		assert.NotNil(t, foundUser)
		assert.Equal(t, "Test User", foundUser.Name)
	})

	t.Run("GetUserByToken", func(t *testing.T) {
		foundUser, err := userStorage.GetUserByToken("valid_token")
		assert.NoError(t, err)
		assert.NotNil(t, foundUser)
	})

	t.Run("DeleteUserByID", func(t *testing.T) {
		err := userStorage.DeleteUserByID(1)
		assert.NoError(t, err)
	})

	t.Run("UpdateUserByID", func(t *testing.T) {
		user := entities.User{
			Name:           "Test User",
			Email:          "test@example.com",
			PasswordDigest: "hashed_password",
			Role:           "User",
			IsRevoked:      false,
		}
		createdUser, err := userStorage.CreateUser(user)
		assert.NoError(t, err)
		assert.NotNil(t, createdUser)

		updatedUser := entities.User{
			Name: "Updated Test User",
		}
		updated, err := userStorage.UpdateUserByID(createdUser.ID, updatedUser)
		assert.NoError(t, err)
		assert.NotNil(t, updated)
		assert.Equal(t, "Updated Test User", updated.Name)
	})

	t.Run("GetAllUsers", func(t *testing.T) {
		users, err := userStorage.GetAllUsers()
		assert.NoError(t, err)
		assert.NotNil(t, users)
	})

	t.Run("CreateInitialAdmin", func(t *testing.T) {
		createdAdmin, err := userStorage.CreateInitialAdmin(admin)
		assert.NoError(t, err)
		assert.NotNil(t, createdAdmin)
		assert.Equal(t, "Admin", createdAdmin.Name)
	})
}
