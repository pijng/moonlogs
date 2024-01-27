package sqlite_adapter

import (
	"context"
	"fmt"
	"moonlogs/internal/entities"
	"moonlogs/internal/lib/qrx"
	"moonlogs/internal/persistence"
)

type UserStorage struct {
	ctx   context.Context
	users *qrx.TableQuerier[entities.User]
}

func NewUserStorage(ctx context.Context) *UserStorage {
	return &UserStorage{
		ctx:   ctx,
		users: qrx.Scan(entities.User{}).With(persistence.DB()).From("users"),
	}
}

func (s *UserStorage) CreateUser(user entities.User) (*entities.User, error) {
	u, err := s.users.Create(s.ctx, map[string]interface{}{
		"email":           user.Email,
		"password":        "",
		"password_digest": user.PasswordDigest,
		"name":            user.Name,
		"role":            user.Role,
		"tag_ids":         user.Tags,
		"token":           "",
	})

	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}

	return u, nil
}

func (r *UserStorage) GetUserByID(id int) (*entities.User, error) {
	u, err := r.users.Where("id = ?", id).First(r.ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}

	return u, nil
}

func (r *UserStorage) GetUsersByTagID(tagID int) ([]*entities.User, error) {
	u, err := r.users.Where("tag_ids LIKE ?", qrx.Contains(tagID)).All(r.ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}

	return u, nil
}

func (r *UserStorage) GetUserByEmail(email string) (*entities.User, error) {
	u, err := r.users.Where("email = ?", email).First(r.ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}

	return u, nil
}

func (r *UserStorage) GetUserByToken(token string) (*entities.User, error) {
	u, err := r.users.Where("token = ?", token).First(r.ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}

	return u, nil
}

func (r *UserStorage) DeleteUserByID(id int) error {
	_, err := r.users.DeleteOne(r.ctx, "id=?", id)

	return err
}

func (r *UserStorage) UpdateUserByID(id int, user entities.User) (*entities.User, error) {
	data := map[string]interface{}{
		"email":   user.Email,
		"name":    user.Name,
		"role":    user.Role,
		"tag_ids": user.Tags,
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

func (r *UserStorage) UpdateUserTokenByID(id int, token string) error {
	_, err := r.users.Where("id = ?", id).UpdateOne(r.ctx, map[string]interface{}{
		"token": token,
	})

	if err != nil {
		return fmt.Errorf("failed updating user: %w", err)
	}

	return nil
}

func (r *UserStorage) GetAllUsers() ([]*entities.User, error) {
	u, err := r.users.All(r.ctx, "")
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}

	return u, nil
}

func (r *UserStorage) GetSystemUser() (*entities.User, error) {
	return r.users.Where("role = 'System'").First(r.ctx)
}

func (r *UserStorage) CreateInitialAdmin(admin entities.User) (*entities.User, error) {
	return r.users.Create(r.ctx, map[string]interface{}{
		"name":            admin.Name,
		"email":           admin.Email,
		"password":        "",
		"password_digest": admin.PasswordDigest,
		"role":            "Admin",
		"token":           "",
	})
}
