package sqlite_adapter

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"moonlogs/internal/entities"
	"moonlogs/internal/lib/qrx"
	"moonlogs/internal/storage"
)

type NotificationProfileStorage struct {
	readDB  *sql.DB
	writeDB *sql.DB
}

func NewNotificationProfileStorage(readDB *sql.DB, writeDB *sql.DB) *NotificationProfileStorage {
	return &NotificationProfileStorage{
		writeDB: readDB,
		readDB:  writeDB,
	}
}

func (s *NotificationProfileStorage) CreateNotificationProfile(ctx context.Context, profile entities.NotificationProfile) (*entities.NotificationProfile, error) {
	tx, err := qrx.BeginImmediate(s.writeDB)
	if err != nil {
		return nil, fmt.Errorf("failed to start transaction: %w", err)
	}

	query := `INSERT INTO notification_profiles
		(name, description, rule_ids, enabled, silence_for,
		url, method, headers, payload
		VALUES (?, ?, ?, ?, ?, ?, ?, ?);`

	result, err := tx.ExecContext(ctx, query,
		profile.Name, profile.Description, profile.RuleIDs, profile.Enabled, profile.SilenceFor,
		profile.URL, profile.Method, profile.Headers, profile.Payload)
	if err != nil {
		return nil, fmt.Errorf("failed inserting notification profile: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("failed retrieving notification profile last insert id: %w", err)
	}

	np, err := s.GetNotificationProfileByID(ctx, int(id))
	if err != nil {
		return nil, fmt.Errorf("failed querying notification profile: %w", err)
	}

	return np, nil
}

func (s *NotificationProfileStorage) GetNotificationProfileByID(ctx context.Context, id int) (*entities.NotificationProfile, error) {
	query := `SELECT id, name, description, rule_ids, enabled, silence_for,
		url, method, headers, payload FROM notification_profiles WHERE id=? LIMIT 1;`
	stmt, err := s.readDB.PrepareContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed preparing statement: %w", err)
	}
	defer stmt.Close()

	row := stmt.QueryRowContext(ctx, id)

	var np entities.NotificationProfile
	err = row.Scan(&np.ID, &np.Name, &np.Description, &np.RuleIDs, &np.Enabled, &np.SilenceFor, &np.URL, &np.Method, &np.Headers, &np.Payload)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = storage.ErrNotFound
		}

		return nil, fmt.Errorf("failed scanning notification profile: %w", err)
	}

	return &np, nil
}

func (s *NotificationProfileStorage) DeleteNotificationProfileByID(ctx context.Context, id int) error {
	tx, err := qrx.BeginImmediate(s.writeDB)
	if err != nil {
		return fmt.Errorf("failed to start transaction: %w", err)
	}

	query := "DELETE FROM notification_profiles WHERE id=?;"

	_, err = tx.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("failed deleting notification profile: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

func (s *NotificationProfileStorage) UpdateNotificationProfileByID(ctx context.Context, id int, profile entities.NotificationProfile) (*entities.NotificationProfile, error) {
	tx, err := qrx.BeginImmediate(s.writeDB)
	if err != nil {
		return nil, fmt.Errorf("failed to start transaction: %w", err)
	}

	query := "UPDATE notification_profiles SET name=?, description=?, rule_ids=?, enabled=?, silence_for=?, url=?, method=?, headers=?, payload=? WHERE id=?;"

	_, err = tx.ExecContext(ctx, query, profile.Name, profile.Description, profile.RuleIDs, profile.Enabled, profile.SilenceFor, profile.URL, profile.Method, profile.Headers, profile.Payload, id)
	if err != nil {
		return nil, fmt.Errorf("failed updating notification profile: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return s.GetNotificationProfileByID(ctx, id)
}

func (s *NotificationProfileStorage) GetAllNotificationProfiles(ctx context.Context) ([]*entities.NotificationProfile, error) {
	query := `SELECT id, name, description, rule_ids, enabled, silence_for,
		url, method, headers, payload FROM notification_profiles ORDER BY id DESC;`
	stmt, err := s.readDB.PrepareContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("failed preparing statement: %w", err)
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying notification profiles: %w", err)
	}
	defer rows.Close()

	notificationProfiles := make([]*entities.NotificationProfile, 0)

	for rows.Next() {
		var np entities.NotificationProfile

		err = rows.Scan(
			&np.ID, &np.Name, &np.Description, &np.RuleIDs, &np.Enabled, &np.SilenceFor, &np.URL, &np.Method, &np.Headers, &np.Payload)
		if err != nil {
			return nil, fmt.Errorf("failed scanning notification profile: %w", err)
		}

		notificationProfiles = append(notificationProfiles, &np)
	}

	return notificationProfiles, nil
}
