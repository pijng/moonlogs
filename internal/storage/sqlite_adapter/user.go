package sqlite_adapter

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"moonlogs/internal/entities"
	"moonlogs/internal/lib/qrx"
	"moonlogs/internal/storage"
	"strings"
)

type UserStorage struct {
	readDB  *sql.DB
	writeDB *sql.DB
}

func NewUserStorage(readDB *sql.DB, writeDB *sql.DB) *UserStorage {
	return &UserStorage{
		writeDB: readDB,
		readDB:  writeDB,
	}
}

func (s *UserStorage) CreateUser(ctx context.Context, user entities.User) (*entities.User, error) {
	tx, err := qrx.BeginImmediate(s.writeDB)
	if err != nil {
		return nil, fmt.Errorf("failed to start transaction: %w", err)
	}

	query := "INSERT INTO users (email, password, password_digest, name, role, tag_ids, token, is_revoked) VALUES (?,?,?,?,?,?,?,?);"

	result, err := tx.ExecContext(ctx, query, user.Email, "", user.PasswordDigest, user.Name, user.Role, user.Tags, "", user.IsRevoked)
	if err != nil {
		return nil, fmt.Errorf("failed inserting user: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("failed retrieving user last insert id: %w", err)
	}

	return s.GetUserByID(ctx, int(id))
}

func (s *UserStorage) GetUserByID(ctx context.Context, id int) (*entities.User, error) {
	query := "SELECT * FROM users WHERE id=? LIMIT 1;"
	stmt, err := s.readDB.PrepareContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed preparing statement: %w", err)
	}
	defer stmt.Close()

	row := stmt.QueryRowContext(ctx, id)

	var u entities.User
	err = row.Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.PasswordDigest, &u.Role, &u.Tags, &u.Token, &u.IsRevoked)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = storage.ErrNotFound
		}

		return nil, fmt.Errorf("failed scanning user: %w", err)
	}

	return &u, nil

}

func (s *UserStorage) GetUsersByTagID(ctx context.Context, tagID int) ([]*entities.User, error) {
	query := "SELECT * FROM users WHERE tag_ids LIKE ? ORDER BY id desc;"
	stmt, err := s.readDB.PrepareContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed preparing statement: %w", err)
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, qrx.Contains(tagID))
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

func (s *UserStorage) GetUserByEmail(ctx context.Context, email string) (*entities.User, error) {
	query := "SELECT * FROM users WHERE email=? LIMIT 1;"
	stmt, err := s.readDB.PrepareContext(ctx, query)
	if err != nil {
		return &entities.User{}, fmt.Errorf("failed preparing statement: %w", err)
	}
	defer stmt.Close()

	row := stmt.QueryRowContext(ctx, email)

	var u entities.User
	err = row.Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.PasswordDigest, &u.Role, &u.Tags, &u.Token, &u.IsRevoked)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = storage.ErrNotFound
		}

		return nil, fmt.Errorf("failed scanning user: %w", err)
	}

	return &u, nil
}

func (s *UserStorage) GetUserByToken(ctx context.Context, token string) (*entities.User, error) {
	query := "SELECT * FROM users WHERE token=? LIMIT 1;"
	stmt, err := s.readDB.PrepareContext(ctx, query)
	if err != nil {
		return &entities.User{}, fmt.Errorf("failed preparing statement: %w", err)
	}
	defer stmt.Close()

	row := stmt.QueryRowContext(ctx, token)

	var u entities.User
	err = row.Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.PasswordDigest, &u.Role, &u.Tags, &u.Token, &u.IsRevoked)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}

	if err != nil {
		return &entities.User{}, fmt.Errorf("failed scanning user: %w", err)
	}

	return &u, nil
}

func (s *UserStorage) DeleteUserByID(ctx context.Context, id int) error {
	tx, err := qrx.BeginImmediate(s.writeDB)
	if err != nil {
		return fmt.Errorf("failed to start transaction: %w", err)
	}

	query := "DELETE FROM users WHERE id=?;"

	_, err = tx.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed deleting user: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func (s *UserStorage) UpdateUserByID(ctx context.Context, id int, user entities.User) (*entities.User, error) {
	tx, err := qrx.BeginImmediate(s.writeDB)
	if err != nil {
		return nil, fmt.Errorf("failed to start transaction: %w", err)
	}

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

	_, err = tx.ExecContext(ctx, queryBuilder.String(), args...)
	if err != nil {
		return nil, fmt.Errorf("failed updating user: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return s.GetUserByID(ctx, id)
}

func (s *UserStorage) UpdateUserTokenByID(ctx context.Context, id int, token string) error {
	tx, err := qrx.BeginImmediate(s.writeDB)
	if err != nil {
		return fmt.Errorf("failed to start transaction: %w", err)
	}

	query := "UPDATE users SET token=? WHERE id=?;"

	_, err = tx.ExecContext(ctx, query, token, id)
	if err != nil {
		return fmt.Errorf("failed updating user: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func (s *UserStorage) GetAllUsers(ctx context.Context) ([]*entities.User, error) {
	query := "SELECT * FROM users ORDER BY id DESC;"
	stmt, err := s.readDB.PrepareContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed preparing statement: %w", err)
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx)
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

func (s *UserStorage) CreateInitialAdmin(ctx context.Context, admin entities.User) (*entities.User, error) {
	tx, err := qrx.BeginImmediate(s.writeDB)
	if err != nil {
		return nil, fmt.Errorf("failed to start transaction: %w", err)
	}

	query := "INSERT INTO users (name, email, password, password_digest, role, token, is_revoked) VALUES (?,?,?,?,?,?,?);"

	result, err := tx.ExecContext(ctx, query, admin.Name, admin.Email, "", admin.PasswordDigest, "Admin", "", admin.IsRevoked)
	if err != nil {
		return nil, fmt.Errorf("failed inserting user: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("failed retrieving user last insert id: %w", err)
	}

	return s.GetUserByID(ctx, int(id))
}
