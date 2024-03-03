package sqlite_adapter

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"moonlogs/internal/entities"
	"moonlogs/internal/lib/qrx"
	"moonlogs/internal/persistence"
	"strings"
)

type UserStorage struct {
	ctx context.Context
	db  *sql.DB
}

func NewUserStorage(ctx context.Context) *UserStorage {
	return &UserStorage{
		ctx: ctx,
		db:  persistence.DB(),
	}
}

func (s *UserStorage) CreateUser(user entities.User) (*entities.User, error) {
	query := "INSERT INTO users (email, password, password_digest, name, role, tag_ids, token, is_revoked) VALUES (?,?,?,?,?,?,?,?);"
	stmt, err := s.db.PrepareContext(s.ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed preparing statement: %w", err)
	}
	defer stmt.Close()

	result, err := stmt.ExecContext(s.ctx, user.Email, "", user.PasswordDigest, user.Name, user.Role, user.Tags, "", user.IsRevoked)
	if err != nil {
		return nil, fmt.Errorf("failed inserting user: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("failed retrieving user last insert id: %w", err)
	}

	return s.GetUserByID(int(id))
}

func (s *UserStorage) GetUserByID(id int) (*entities.User, error) {
	query := "SELECT * FROM users WHERE id=? LIMIT 1;"
	stmt, err := s.db.PrepareContext(s.ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed preparing statement: %w", err)
	}
	defer stmt.Close()

	row := stmt.QueryRowContext(s.ctx, id)

	var u entities.User
	err = row.Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.PasswordDigest, &u.Role, &u.Tags, &u.Token, &u.IsRevoked)
	if err != nil {
		return nil, fmt.Errorf("failed scanning user: %w", err)
	}

	return &u, nil

}

func (s *UserStorage) GetUsersByTagID(tagID int) ([]*entities.User, error) {
	query := "SELECT * FROM users WHERE tag_ids LIKE ? ORDER BY id desc;"
	stmt, err := s.db.PrepareContext(s.ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed preparing statement: %w", err)
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(s.ctx, qrx.Contains(tagID))
	if err != nil {
		return nil, fmt.Errorf("failed querying users: %w", err)
	}
	defer rows.Close()

	users := make([]*entities.User, 0)

	for rows.Next() {
		var u entities.User

		err = rows.Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.PasswordDigest, &u.Role, &u.Tags, &u.Token, &u.IsRevoked)
		if err != nil {
			return nil, fmt.Errorf("failed scanning user: %w", err)
		}

		users = append(users, &u)
	}

	return users, nil
}

func (s *UserStorage) GetUserByEmail(email string) (*entities.User, error) {
	query := "SELECT * FROM users WHERE email=? LIMIT 1;"
	stmt, err := s.db.PrepareContext(s.ctx, query)
	if err != nil {
		return &entities.User{}, fmt.Errorf("failed preparing statement: %w", err)
	}
	defer stmt.Close()

	row := stmt.QueryRowContext(s.ctx, email)

	var u entities.User
	err = row.Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.PasswordDigest, &u.Role, &u.Tags, &u.Token, &u.IsRevoked)
	if errors.Is(err, sql.ErrNoRows) {
		return &entities.User{}, nil
	}

	if err != nil {
		return &entities.User{}, fmt.Errorf("failed scanning user: %w", err)
	}

	return &u, nil
}

func (s *UserStorage) GetUserByToken(token string) (*entities.User, error) {
	query := "SELECT * FROM users WHERE token=? LIMIT 1;"
	stmt, err := s.db.PrepareContext(s.ctx, query)
	if err != nil {
		return &entities.User{}, fmt.Errorf("failed preparing statement: %w", err)
	}
	defer stmt.Close()

	row := stmt.QueryRowContext(s.ctx, token)

	var u entities.User
	err = row.Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.PasswordDigest, &u.Role, &u.Tags, &u.Token, &u.IsRevoked)
	if errors.Is(err, sql.ErrNoRows) {
		return &entities.User{}, nil
	}

	if err != nil {
		return &entities.User{}, fmt.Errorf("failed scanning user: %w", err)
	}

	return &u, nil
}

func (s *UserStorage) DeleteUserByID(id int) error {
	query := "DELETE FROM users WHERE id=?;"
	stmt, err := s.db.PrepareContext(s.ctx, query)
	if err != nil {
		return fmt.Errorf("failed preparing statement: %w", err)
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(s.ctx, id)
	if err != nil {
		return fmt.Errorf("failed deleting user: %w", err)
	}

	return err
}

func (s *UserStorage) UpdateUserByID(id int, user entities.User) (*entities.User, error) {
	var queryBuilder strings.Builder
	args := make([]interface{}, 0)

	queryBuilder.WriteString("UPDATE users SET email=?, name=?, role=?, tag_ids=?, is_revoked=?")
	args = append(args, user.Email, user.Name, user.Role, user.Tags, user.IsRevoked)

	if len(user.PasswordDigest) > 0 {
		queryBuilder.WriteString(", password_digest=?, token=?")
		args = append(args, user.PasswordDigest, user.Token)
	}

	queryBuilder.WriteString(" WHERE id=?;")
	args = append(args, id)

	stmt, err := s.db.PrepareContext(s.ctx, queryBuilder.String())
	if err != nil {
		return nil, fmt.Errorf("failed preparing statement: %w", err)
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(s.ctx, args...)
	fmt.Println(err)
	if err != nil {
		return nil, fmt.Errorf("failed updating user: %w", err)
	}

	return s.GetUserByID(id)
}

func (s *UserStorage) UpdateUserTokenByID(id int, token string) error {
	query := "UPDATE users SET token=? WHERE id=?;"
	stmt, err := s.db.PrepareContext(s.ctx, query)
	if err != nil {
		return fmt.Errorf("failed preparing statement: %w", err)
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(s.ctx, token, id)
	if err != nil {
		return fmt.Errorf("failed updating user: %w", err)
	}

	return err
}

func (s *UserStorage) GetAllUsers() ([]*entities.User, error) {
	query := "SELECT * FROM users ORDER BY id DESC;"
	stmt, err := s.db.PrepareContext(s.ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed preparing statement: %w", err)
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(s.ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %w", err)
	}
	defer rows.Close()

	users := make([]*entities.User, 0)

	for rows.Next() {
		var u entities.User

		err = rows.Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.PasswordDigest, &u.Role, &u.Tags, &u.Token, &u.IsRevoked)
		if err != nil {
			return nil, fmt.Errorf("failed scanning user: %w", err)
		}

		users = append(users, &u)
	}

	return users, nil
}

func (s *UserStorage) CreateInitialAdmin(admin entities.User) (*entities.User, error) {
	query := "INSERT INTO users (name, email, password, password_digest, role, token, is_revoked) VALUES (?,?,?,?,?,?,?);"
	stmt, err := s.db.PrepareContext(s.ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed preparing statement: %w", err)
	}
	defer stmt.Close()

	result, err := stmt.ExecContext(s.ctx, admin.Name, admin.Email, "", admin.PasswordDigest, "Admin", "", admin.IsRevoked)
	if err != nil {
		return nil, fmt.Errorf("failed inserting user: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("failed retrieving user last insert id: %w", err)
	}

	return s.GetUserByID(int(id))
}
