package repository

import (
	"context"
	"fmt"
	"moonlogs/ent"
	"moonlogs/ent/user"
	"moonlogs/internal/config"
)

type UserRepository struct {
	ctx    context.Context
	client *ent.Client
}

func NewUserRepository(ctx context.Context) *UserRepository {
	return &UserRepository{
		ctx:    ctx,
		client: config.GetClient(),
	}
}

func (r *UserRepository) Create(user ent.User) (*ent.User, error) {
	u, err := r.client.User.
		Create().
		SetEmail(user.Email).
		SetPasswordDigest(user.PasswordDigest).
		SetName(user.Name).
		SetRole(user.Role).
		Save(r.ctx)

	if err != nil {
		return nil, fmt.Errorf("failed creating user: %w", err)
	}

	return u, nil
}

func (r *UserRepository) GetById(id int) (*ent.User, error) {
	u, err := r.client.User.
		Query().
		Where(user.ID(id)).First(r.ctx)

	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}

	return u, nil
}

func (r *UserRepository) GetByEmail(email string) (*ent.User, error) {
	u, err := r.client.User.
		Query().
		Where(user.Email(email)).First(r.ctx)

	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}

	return u, nil
}

func (r *UserRepository) GetByToken(token string) (*ent.User, error) {
	u, err := r.client.User.
		Query().
		Where(user.Token(token)).First(r.ctx)

	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}

	return u, nil
}

func (r *UserRepository) DestroyById(id int) error {
	return r.client.User.DeleteOneID(id).Exec(r.ctx)
}

func (r *UserRepository) UpdateById(userToUpdate ent.User) (*ent.User, error) {
	transaction := r.client.User.UpdateOneID(userToUpdate.ID).
		SetEmail(userToUpdate.Email).
		SetName(userToUpdate.Name).
		SetRole(userToUpdate.Role)

	if len(userToUpdate.PasswordDigest) > 0 {
		transaction = transaction.SetPasswordDigest(userToUpdate.PasswordDigest).SetToken(userToUpdate.Token)
	}

	u, err := transaction.Save(r.ctx)
	if err != nil {
		return nil, fmt.Errorf("failed updating user: %w", err)
	}

	return u, nil
}

func (r *UserRepository) UpdateTokenById(id int, token string) error {
	_, err := r.client.User.UpdateOneID(id).
		SetToken(token).
		Save(r.ctx)

	if err != nil {
		return fmt.Errorf("failed updating user: %w", err)
	}

	return nil
}

func (r *UserRepository) GetAll() ([]*ent.User, error) {
	u, err := r.client.User.
		Query().All(r.ctx)

	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}

	return u, nil
}
