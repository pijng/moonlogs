package sqlite_adapter

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

	writeDB, readDB, err := testutil.SetupSqlite()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := testutil.TeardownSqlite(); err != nil {
			log.Fatal(err)
		}
	}()

	userStorage := NewUserStorage(readDB, writeDB)

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
		createdUser, err := userStorage.CreateUser(ctx, user)
		assert.NoError(t, err)
		assert.NotNil(t, createdUser)
		assert.Equal(t, "Test User", createdUser.Name)
	})

	t.Run("GetUserByID", func(t *testing.T) {
		testUser, err := userStorage.GetUserByEmail(ctx, "test@example.com")
		assert.NoError(t, err)
		assert.NotNil(t, testUser)

		foundUser, err := userStorage.GetUserByID(ctx, testUser.ID)
		assert.NoError(t, err)
		assert.NotNil(t, foundUser)
		assert.Equal(t, "Test User", foundUser.Name)
	})

	t.Run("GetUsersByTagID", func(t *testing.T) {
		users, err := userStorage.GetUsersByTagID(ctx, 1)
		assert.NoError(t, err)
		assert.NotNil(t, users)
	})

	t.Run("GetUserByEmail", func(t *testing.T) {
		foundUser, err := userStorage.GetUserByEmail(ctx, "test@example.com")
		assert.NoError(t, err)
		assert.NotNil(t, foundUser)
		assert.Equal(t, "Test User", foundUser.Name)
	})

	t.Run("GetUserByToken", func(t *testing.T) {
		user := entities.User{
			Name:           "Test User",
			Email:          "test@example.com",
			PasswordDigest: "hashed_password",
			Role:           "User",
			IsRevoked:      false,
		}
		createdUser, err := userStorage.CreateUser(ctx, user)
		assert.NoError(t, err)
		assert.NotNil(t, createdUser)

		token := "secure_token"
		err = userStorage.UpdateUserTokenByID(ctx, createdUser.ID, token)
		assert.NoError(t, err)

		foundUser, err := userStorage.GetUserByToken(ctx, token)
		assert.NoError(t, err)
		assert.NotNil(t, foundUser)
	})

	t.Run("DeleteUserByID", func(t *testing.T) {
		err := userStorage.DeleteUserByID(ctx, 1)
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
		createdUser, err := userStorage.CreateUser(ctx, user)
		assert.NoError(t, err)
		assert.NotNil(t, createdUser)

		updatedUser := entities.User{
			Name: "Updated Test User",
		}
		updated, err := userStorage.UpdateUserByID(ctx, createdUser.ID, updatedUser)
		assert.NoError(t, err)
		assert.NotNil(t, updated)
		assert.Equal(t, "Updated Test User", updated.Name)
	})

	t.Run("GetAllUsers", func(t *testing.T) {
		users, err := userStorage.GetAllUsers(ctx)
		assert.NoError(t, err)
		assert.NotNil(t, users)
	})

	t.Run("CreateInitialAdmin", func(t *testing.T) {
		createdAdmin, err := userStorage.CreateInitialAdmin(ctx, admin)
		assert.NoError(t, err)
		assert.NotNil(t, createdAdmin)
		assert.Equal(t, "Admin", createdAdmin.Name)
	})
}
