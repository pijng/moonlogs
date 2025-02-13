package storage

import (
	"context"
	"moonlogs/internal/entities"
)

type NotificationProfileStorage interface {
	CreateNotificationProfile(ctx context.Context, profile entities.NotificationProfile) (*entities.NotificationProfile, error)
	DeleteNotificationProfileByID(ctx context.Context, id int) error
	GetAllNotificationProfiles(ctx context.Context) ([]*entities.NotificationProfile, error)
	GetNotificationProfileByID(ctx context.Context, id int) (*entities.NotificationProfile, error)
	UpdateNotificationProfileByID(ctx context.Context, id int, profile entities.NotificationProfile) (*entities.NotificationProfile, error)
}
