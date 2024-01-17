package usecases

import (
	"fmt"
	"moonlogs/api/server/session"
	"moonlogs/internal/entities"
	"moonlogs/internal/repositories"
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
	userRepository *repositories.UserRepository
}

func NewUserUseCase(userRepository *repositories.UserRepository) *UserUseCase {
	return &UserUseCase{userRepository: userRepository}
}

func (uc *UserUseCase) CreateUser(user entities.User) (*entities.User, error) {
	userWithIdenticalEmail, err := uc.userRepository.GetUserByEmail(user.Email)
	if userWithIdenticalEmail.ID != 0 && err == nil {
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

	user.PasswordDigest = passwordDigest

	return uc.userRepository.CreateUser(user)
}

func (uc *UserUseCase) DestroyUserByID(id int) error {
	return uc.userRepository.DestroyUserByID(id)
}

func (uc *UserUseCase) GetUserByID(id int) (*entities.User, error) {
	return uc.userRepository.GetUserByID(id)
}

func (uc *UserUseCase) UpdateUserByID(id int, user entities.User) (*entities.User, error) {
	if len(user.Password) > 0 {
		passwordDigest, err := hashPassword(user.Password)
		if err != nil {
			return nil, fmt.Errorf("failed hashing user password: %w", err)
		}

		user.PasswordDigest = passwordDigest
	}

	if len(user.PasswordDigest) > 0 {
		token, err := session.GenerateAuthToken()
		if err != nil {
			return nil, fmt.Errorf("failed generating auth token for user: %w", err)
		}

		user.Token = token
	}

	return uc.userRepository.UpdateUserByID(id, user)
}

func (uc *UserUseCase) GetAllUsers() ([]*entities.User, error) {
	return uc.userRepository.GetAllUsers()
}

func (uc *UserUseCase) GetUserByToken(token string) (*entities.User, error) {
	return uc.userRepository.GetUserByToken(token)
}

func (uc *UserUseCase) GetUserByEmail(email string) (*entities.User, error) {
	return uc.userRepository.GetUserByEmail(email)
}

func (uc *UserUseCase) UpdateUserTokenByID(id int, token string) error {
	return uc.userRepository.UpdateUserTokenByID(id, token)
}

func (uc *UserUseCase) ShouldCreateInitialAdmin() (bool, error) {
	users, err := uc.userRepository.GetAllUsers()
	if err != nil {
		return false, fmt.Errorf("failed querying system user: %w", err)
	}

	if len(users) > 0 {
		return false, nil
	}

	return true, nil
}

func (uc *UserUseCase) CreateInitialAdmin(admin entities.User) (*entities.User, error) {
	users, err := uc.userRepository.GetAllUsers()
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

	return uc.userRepository.CreateInitialAdmin(admin)
}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}
