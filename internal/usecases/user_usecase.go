package usecases

import (
	"context"
	"errors"
	"fmt"
	"moonlogs/internal/entities"
	"moonlogs/internal/shared"
	"moonlogs/internal/storage"
	"slices"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

var appropriateRoles = []string{
	string(entities.MemberRole),
	string(entities.AdminRole),
}

var appropriateRolesInfo = strings.Join(appropriateRoles, ", ")

type UserUseCase struct {
	userStorage storage.UserStorage
}

func NewUserUseCase(userStorage storage.UserStorage) *UserUseCase {
	return &UserUseCase{userStorage: userStorage}
}

func (uc *UserUseCase) CreateUser(ctx context.Context, user entities.User) (*entities.User, error) {
	userWithIdenticalEmail, err := uc.userStorage.GetUserByEmail(ctx, user.Email)
	if err != nil && !errors.Is(err, storage.ErrNotFound) {
		return nil, fmt.Errorf("failed getting user: %w", err)
	}

	if userWithIdenticalEmail != nil {
		return nil, fmt.Errorf("user with email %s already exists", user.Email)
	}

	if len(user.Role) == 0 {
		return nil, fmt.Errorf("user role is empty")
	}

	isValidRole := slices.Contains(appropriateRoles, string(user.Role))
	if !isValidRole {
		return nil, fmt.Errorf("role attribute should be one of: %v", appropriateRolesInfo)
	}

	passwordDigest, err := hashPassword(user.Password)
	if err != nil {
		return nil, fmt.Errorf("failed hashing user password: %w", err)
	}

	if user.Role == entities.AdminRole {
		user.Tags = []int{}
	}

	user.PasswordDigest = passwordDigest

	return uc.userStorage.CreateUser(ctx, user)
}

func (uc *UserUseCase) DeleteUserByID(ctx context.Context, id int) error {
	return uc.userStorage.DeleteUserByID(ctx, id)
}

func (uc *UserUseCase) GetUserByID(ctx context.Context, id int) (*entities.User, error) {
	return uc.userStorage.GetUserByID(ctx, id)
}

func (uc *UserUseCase) GetUsersByTagID(ctx context.Context, tagID int) ([]*entities.User, error) {
	return uc.userStorage.GetUsersByTagID(ctx, tagID)
}

func (uc *UserUseCase) UpdateUserByID(ctx context.Context, id int, user entities.User) (*entities.User, error) {
	if len(user.Password) > 0 {
		passwordDigest, err := hashPassword(user.Password)
		if err != nil {
			return nil, fmt.Errorf("failed hashing user password: %w", err)
		}

		user.PasswordDigest = passwordDigest
	}

	if len(user.PasswordDigest) > 0 {
		token, err := shared.GenerateRandomToken(16)
		if err != nil {
			return nil, fmt.Errorf("failed generating auth token for user: %w", err)
		}

		user.Token = token
	}

	if user.Role == entities.AdminRole {
		user.Tags = []int{}
	}

	if user.IsRevoked {
		user.Token = ""
	}

	return uc.userStorage.UpdateUserByID(ctx, id, user)
}

func (uc *UserUseCase) GetAllUsers(ctx context.Context) ([]*entities.User, error) {
	return uc.userStorage.GetAllUsers(ctx)
}

func (uc *UserUseCase) GetUserByToken(ctx context.Context, token string) (*entities.User, error) {
	return uc.userStorage.GetUserByToken(ctx, token)
}

func (uc *UserUseCase) GetUserByEmail(ctx context.Context, email string) (*entities.User, error) {
	return uc.userStorage.GetUserByEmail(ctx, email)
}

func (uc *UserUseCase) UpdateUserTokenByID(ctx context.Context, id int, token string) error {
	return uc.userStorage.UpdateUserTokenByID(ctx, id, token)
}

func (uc *UserUseCase) ShouldCreateInitialAdmin(ctx context.Context) (bool, error) {
	users, err := uc.userStorage.GetAllUsers(ctx)
	if err != nil {
		return false, fmt.Errorf("failed querying system user: %w", err)
	}

	if len(users) > 0 {
		return false, nil
	}

	return true, nil
}

func (uc *UserUseCase) CreateInitialAdmin(ctx context.Context, admin entities.User) (*entities.User, error) {
	users, err := uc.userStorage.GetAllUsers(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying users: %w", err)
	}

	if len(users) > 0 {
		return nil, fmt.Errorf("initial admin already exist: %w", err)
	}

	passwordDigest, err := hashPassword(admin.Password)
	if err != nil {
		return nil, fmt.Errorf("failed hashing admin user password: %w", err)
	}

	admin.PasswordDigest = passwordDigest

	return uc.userStorage.CreateInitialAdmin(ctx, admin)
}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}
