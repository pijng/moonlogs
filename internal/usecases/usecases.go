package usecases

import "moonlogs/internal/persistence"

type UseCases struct {
	ActionUseCase              *ActionUseCase
	ApiTokenUseCase            *ApiTokenUseCase
	RecordUseCase              *RecordUseCase
	SchemaUseCase              *SchemaUseCase
	TagUseCase                 *TagUseCase
	UserUseCase                *UserUseCase
	AlertingRuleUseCase        *AlertingRuleUseCase
	IncidentUseCase            *IncidentUseCase
	NotificationProfileUseCase *NotificationProfileUseCase
}

func InitUsecases(storages persistence.Storages) *UseCases {
	return &UseCases{
		ActionUseCase:              NewActionUseCase(storages.ActionStorage),
		ApiTokenUseCase:            NewApiTokenUseCase(storages.ApiTokenStorage),
		RecordUseCase:              NewRecordUseCase(storages.RecordStorage),
		SchemaUseCase:              NewSchemaUseCase(storages.SchemaStorage),
		TagUseCase:                 NewTagUseCase(storages.TagStorage),
		UserUseCase:                NewUserUseCase(storages.UserStorage),
		AlertingRuleUseCase:        NewAlertingRuleUseCase(storages.AlertingRuleStorage),
		IncidentUseCase:            NewIncidentUseCase(storages.IncidentStorage),
		NotificationProfileUseCase: NewNotificationProfileUseCase(storages.NotificationProfileStorage),
	}
}
