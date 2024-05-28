package storage

import (
	"context"
	"moonlogs/internal/entities"
)

type UserStorage interface {
	CreateInitialAdmin(ctx context.Context, admin entities.User) (*entities.User, error)
	CreateUser(ctx context.Context, user entities.User) (*entities.User, error)
	DeleteUserByID(ctx context.Context, id int) error
	GetAllUsers(ctx context.Context) ([]*entities.User, error)
	GetUserByEmail(ctx context.Context, email string) (*entities.User, error)
	GetUserByID(ctx context.Context, id int) (*entities.User, error)
	GetUsersByTagID(ctx context.Context, id int) ([]*entities.User, error)
	GetUserByToken(ctx context.Context, token string) (*entities.User, error)
	UpdateUserByID(ctx context.Context, id int, user entities.User) (*entities.User, error)
	UpdateUserTokenByID(ctx context.Context, id int, token string) error
}
