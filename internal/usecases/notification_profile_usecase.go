package usecases

import (
	"context"
	"fmt"
	"moonlogs/internal/entities"
	"moonlogs/internal/storage"
)

type NotificationProfileUseCase struct {
	notificationProfileStorage storage.NotificationProfileStorage
}

func NewNotificationProfileUseCase(notificationProfileStorage storage.NotificationProfileStorage) *NotificationProfileUseCase {
	return &NotificationProfileUseCase{notificationProfileStorage: notificationProfileStorage}
}

func (nc *NotificationProfileUseCase) CreateNotificationProfile(ctx context.Context, profile entities.NotificationProfile) (*entities.NotificationProfile, error) {
	if profile.Name == "" {
		return nil, fmt.Errorf("failed creating notification profile: `name` attribute is required")
	}

	if profile.Description == "" {
		return nil, fmt.Errorf("failed creating notification profile: `description` attribute is required")
	}

	if profile.Method == "" {
		return nil, fmt.Errorf("failed creating notification profile: `method` attribute is required")
	}

	if profile.URL == "" {
		return nil, fmt.Errorf("failed creating notification profile: `url` attribute is required")
	}

	return nc.notificationProfileStorage.CreateNotificationProfile(ctx, profile)
}

func (nc *NotificationProfileUseCase) GetAllNotificationProfiles(ctx context.Context) ([]*entities.NotificationProfile, error) {
	return nc.notificationProfileStorage.GetAllNotificationProfiles(ctx)
}

func (nc *NotificationProfileUseCase) DeleteNotificationProfileByID(ctx context.Context, id int) error {
	return nc.notificationProfileStorage.DeleteNotificationProfileByID(ctx, id)
}

func (nc *NotificationProfileUseCase) GetNotificationProfileByID(ctx context.Context, id int) (*entities.NotificationProfile, error) {
	return nc.notificationProfileStorage.GetNotificationProfileByID(ctx, id)
}

func (nc *NotificationProfileUseCase) UpdateNotificationProfileByID(ctx context.Context, id int, profile entities.NotificationProfile) (*entities.NotificationProfile, error) {
	return nc.notificationProfileStorage.UpdateNotificationProfileByID(ctx, id, profile)
}
