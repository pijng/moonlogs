package usecases

import "moonlogs/internal/persistence"

type ProxyCfg struct {
	ProxyHost string
	ProxyPort string
	ProxyUser string
	ProxyPass string
}

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
	InsightsUseCase            *InsightsUseCase
}

func InitUsecases(storages persistence.Storages, insightsAdapter InsightsAdapter, proxyCfg ProxyCfg) *UseCases {
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
		InsightsUseCase:            NewInsightsUseCase(insightsAdapter, proxyCfg.ProxyUser, proxyCfg.ProxyPass, proxyCfg.ProxyHost, proxyCfg.ProxyPort),
	}
}
