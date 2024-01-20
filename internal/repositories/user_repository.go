package repositories

import (
	"context"
	"fmt"
	"moonlogs/internal/entities"
	"moonlogs/internal/lib/qrx"
	"moonlogs/internal/persistence"
)

type UserRepository struct {
	ctx   context.Context
	users *qrx.TableQuerier[entities.User]
}

func NewUserRepository(ctx context.Context) *UserRepository {
	return &UserRepository{
		ctx:   ctx,
		users: qrx.Scan(entities.User{}).With(persistence.DB()).From("users"),
	}
}

func (r *UserRepository) CreateUser(user entities.User) (*entities.User, error) {
	u, err := r.users.Create(r.ctx, map[string]interface{}{
		"email":           user.Email,
		"password":        "",
		"password_digest": user.PasswordDigest,
		"name":            user.Name,
		"role":            user.Role,
		"token":           "",
	})

	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}

	return u, nil
}

func (r *UserRepository) GetUserByID(id int) (*entities.User, error) {
	u, err := r.users.Where("id = ?", id).First(r.ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}

	return u, nil
}

func (r *UserRepository) GetUserByEmail(email string) (*entities.User, error) {
	u, err := r.users.Where("email = ?", email).First(r.ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}

	return u, nil
}

func (r *UserRepository) GetUserByToken(token string) (*entities.User, error) {
	u, err := r.users.Where("token = ?", token).First(r.ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}

	return u, nil
}

func (r *UserRepository) DestroyUserByID(id int) error {
	_, err := r.users.DeleteOne(r.ctx, "id=?", id)

	return err
}

func (r *UserRepository) UpdateUserByID(id int, user entities.User) (*entities.User, error) {
	data := map[string]interface{}{
		"email": user.Email,
		"name":  user.Name,
		"role":  user.Role,
	}

	if len(user.PasswordDigest) > 0 {
		data["password_digest"] = user.PasswordDigest
		data["token"] = user.Token
	}

	u, err := r.users.Where("id = ?", id).UpdateOne(r.ctx, data)
	if err != nil {
		return nil, fmt.Errorf("failed updating user: %w", err)
	}

	return u, nil
}

func (r *UserRepository) UpdateUserTokenByID(id int, token string) error {
	_, err := r.users.Where("id = ?", id).UpdateOne(r.ctx, map[string]interface{}{
		"token": token,
	})

	if err != nil {
		return fmt.Errorf("failed updating user: %w", err)
	}

	return nil
}

func (r *UserRepository) GetAllUsers() ([]*entities.User, error) {
	u, err := r.users.All(r.ctx, "")
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}

	return u, nil
}

func (r *UserRepository) GetSystemUser() (*entities.User, error) {
	return r.users.Where("role = 'System'").First(r.ctx)
}

func (r *UserRepository) CreateInitialAdmin(admin entities.User) (*entities.User, error) {
	return r.users.Create(r.ctx, map[string]interface{}{
		"name":            admin.Name,
		"email":           admin.Email,
		"password":        "",
		"password_digest": admin.PasswordDigest,
		"role":            "Admin",
		"token":           "",
	})
}
