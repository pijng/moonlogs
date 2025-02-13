package router

import (
	"moonlogs/internal/api/server/controllers"
	"moonlogs/internal/entities"
	"moonlogs/internal/usecases"
	"net/http"

	"github.com/gorilla/mux"
)

type SubRouterConfig struct {
	R  *mux.Router
	MW *Middlewares
	UC *usecases.UseCases
}

func RegisterSchemaRouter(cfg *SubRouterConfig) {
	schemaRouter := cfg.R.PathPrefix("/api/schemas").Subrouter()
	schemaRouter.Use(cfg.MW.SessionMiddleware)

	schemaController := controllers.NewSchemaController(cfg.UC.SchemaUseCase)

	schemaRouter.HandleFunc("", schemaController.GetAllSchemas).Methods(http.MethodGet)
	schemaRouter.HandleFunc("", cfg.MW.RoleMiddleware(schemaController.CreateSchema, entities.AdminRole, entities.TokenRole)).Methods(http.MethodPost)
	schemaRouter.HandleFunc("/{id}", schemaController.GetSchemaByID).Methods(http.MethodGet)
	schemaRouter.HandleFunc("/{id}", cfg.MW.RoleMiddleware(schemaController.UpdateSchemaByID, entities.AdminRole, entities.TokenRole)).Methods(http.MethodPut)
	schemaRouter.HandleFunc("/{id}", cfg.MW.RoleMiddleware(schemaController.DeleteSchemaByID, entities.AdminRole)).Methods(http.MethodDelete)
}

func RegisterRecordRouter(cfg *SubRouterConfig) {
	recordRouter := cfg.R.PathPrefix("/api/logs").Subrouter()
	recordRouter.Use(cfg.MW.SessionMiddleware)

	recordController := controllers.NewRecordController(cfg.UC.RecordUseCase, cfg.UC.SchemaUseCase)

	recordRouter.HandleFunc("", recordController.GetAllRecords).Methods(http.MethodGet)
	recordRouter.HandleFunc("", cfg.MW.RoleMiddleware(recordController.CreateRecord, entities.AdminRole, entities.TokenRole)).Methods(http.MethodPost)
	recordRouter.HandleFunc("/async", cfg.MW.RoleMiddleware(recordController.CreateRecordAsync, entities.AdminRole, entities.TokenRole)).Methods(http.MethodPost)
	recordRouter.HandleFunc("/{id}", recordController.GetRecordByID).Methods(http.MethodGet)
	recordRouter.HandleFunc("/{id}/request", recordController.GetRecordRequestByID).Methods(http.MethodGet)
	recordRouter.HandleFunc("/{id}/response", recordController.GetRecordResponseByID).Methods(http.MethodGet)
	recordRouter.HandleFunc("/group/{schemaName}/{hash}", recordController.GetRecordsByGroupHash).Methods(http.MethodGet)
	recordRouter.HandleFunc("/search", recordController.GetRecordsByQuery).Methods(http.MethodPost)
}

func RegisterUserRouter(cfg *SubRouterConfig) {
	userRouter := cfg.R.PathPrefix("/api/users").Subrouter()
	userRouter.Use(cfg.MW.SessionMiddleware)

	userController := controllers.NewUserController(cfg.UC.UserUseCase)

	userRouter.HandleFunc("", userController.GetAllUsers).Methods(http.MethodGet)
	userRouter.HandleFunc("", cfg.MW.RoleMiddleware(userController.CreateUser, entities.AdminRole)).Methods(http.MethodPost)
	userRouter.HandleFunc("/{id}", userController.GetUserByID).Methods(http.MethodGet)
	userRouter.HandleFunc("/{id}", cfg.MW.RoleMiddleware(userController.UpdateUserByID, entities.AdminRole)).Methods(http.MethodPut)
	userRouter.HandleFunc("/{id}", cfg.MW.RoleMiddleware(userController.DeleteUserByID, entities.AdminRole)).Methods(http.MethodDelete)
}

func RegisterApiTokenRouter(cfg *SubRouterConfig) {
	apiTokenRouter := cfg.R.PathPrefix("/api/api_tokens").Subrouter()
	apiTokenRouter.Use(cfg.MW.SessionMiddleware)

	apiTokenController := controllers.NewApiTokenController(cfg.UC.ApiTokenUseCase)

	apiTokenRouter.HandleFunc("", cfg.MW.RoleMiddleware(apiTokenController.GetAllApiTokens, entities.AdminRole)).Methods(http.MethodGet)
	apiTokenRouter.HandleFunc("", cfg.MW.RoleMiddleware(apiTokenController.CreateApiToken, entities.AdminRole)).Methods(http.MethodPost)
	apiTokenRouter.HandleFunc("/{id}", cfg.MW.RoleMiddleware(apiTokenController.GetApiTokenByID, entities.AdminRole)).Methods(http.MethodGet)
	apiTokenRouter.HandleFunc("/{id}", cfg.MW.RoleMiddleware(apiTokenController.UpdateApiTokenByID, entities.AdminRole)).Methods(http.MethodPut)
	apiTokenRouter.HandleFunc("/{id}", cfg.MW.RoleMiddleware(apiTokenController.DeleteApiTokenByID, entities.AdminRole)).Methods(http.MethodDelete)
}

func RegisterTagRouter(cfg *SubRouterConfig) {
	tagRouter := cfg.R.PathPrefix("/api/tags").Subrouter()
	tagRouter.Use(cfg.MW.SessionMiddleware)

	tagController := controllers.NewTagController(cfg.UC.TagUseCase, cfg.UC.UserUseCase, cfg.UC.SchemaUseCase)

	tagRouter.HandleFunc("", tagController.GetAllTags).Methods(http.MethodGet)
	tagRouter.HandleFunc("", cfg.MW.RoleMiddleware(tagController.CreateTag, entities.AdminRole)).Methods(http.MethodPost)
	tagRouter.HandleFunc("/{id}", tagController.GetTagByID).Methods(http.MethodGet)
	tagRouter.HandleFunc("/{id}", cfg.MW.RoleMiddleware(tagController.UpdateTagByID, entities.AdminRole)).Methods(http.MethodPut)
	tagRouter.HandleFunc("/{id}", cfg.MW.RoleMiddleware(tagController.DeleteTagByID, entities.AdminRole)).Methods(http.MethodDelete)
}

func RegisterAlertingRuleRouter(cfg *SubRouterConfig) {
	ruleRouter := cfg.R.PathPrefix("/api/alerting_rules").Subrouter()
	ruleRouter.Use(cfg.MW.SessionMiddleware)

	ruleController := controllers.NewAlertingRuleController(cfg.UC.AlertingRuleUseCase)

	ruleRouter.HandleFunc("", cfg.MW.RoleMiddleware(ruleController.GetAllRules, entities.AdminRole)).Methods(http.MethodGet)
	ruleRouter.HandleFunc("", cfg.MW.RoleMiddleware(ruleController.CreateRule, entities.AdminRole)).Methods(http.MethodPost)
	ruleRouter.HandleFunc("/{id}", cfg.MW.RoleMiddleware(ruleController.GetRuleByID, entities.AdminRole)).Methods(http.MethodGet)
	ruleRouter.HandleFunc("/{id}", cfg.MW.RoleMiddleware(ruleController.UpdateRuleByID, entities.AdminRole)).Methods(http.MethodPut)
	ruleRouter.HandleFunc("/{id}", cfg.MW.RoleMiddleware(ruleController.DeleteRuleByID, entities.AdminRole)).Methods(http.MethodDelete)
}

func RegisterIncidentsRouter(cfg *SubRouterConfig) {
	incidentRouter := cfg.R.PathPrefix("/api/incidents").Subrouter()
	incidentRouter.Use(cfg.MW.SessionMiddleware)

	incidentController := controllers.NewIncidentController(cfg.UC.IncidentUseCase)

	incidentRouter.HandleFunc("", cfg.MW.RoleMiddleware(incidentController.GetAllIncidents, entities.AdminRole, entities.TokenRole)).Methods(http.MethodGet)
	incidentRouter.HandleFunc("/search", cfg.MW.RoleMiddleware(incidentController.GetIncidentsByKeys, entities.AdminRole, entities.TokenRole)).Methods(http.MethodPost)
}

func RegisterNotificationProfileRouter(cfg *SubRouterConfig) {
	profileRouter := cfg.R.PathPrefix("/api/notification_profiles").Subrouter()
	profileRouter.Use(cfg.MW.SessionMiddleware)

	profileController := controllers.NewNotificationProfileController(cfg.UC.NotificationProfileUseCase)

	profileRouter.HandleFunc("", cfg.MW.RoleMiddleware(profileController.GetAllNotificationProfiles, entities.AdminRole)).Methods(http.MethodGet)
	profileRouter.HandleFunc("", cfg.MW.RoleMiddleware(profileController.CreateNotificationProfile, entities.AdminRole)).Methods(http.MethodPost)
	profileRouter.HandleFunc("/{id}", cfg.MW.RoleMiddleware(profileController.GetNotificationProfileByID, entities.AdminRole)).Methods(http.MethodGet)
	profileRouter.HandleFunc("/{id}", cfg.MW.RoleMiddleware(profileController.UpdateNotificationProfileByID, entities.AdminRole)).Methods(http.MethodPut)
	profileRouter.HandleFunc("/{id}", cfg.MW.RoleMiddleware(profileController.DeleteNotificationProfileByID, entities.AdminRole)).Methods(http.MethodDelete)
}

func RegisterActionRouter(cfg *SubRouterConfig) {
	actionRouter := cfg.R.PathPrefix("/api/actions").Subrouter()
	actionRouter.Use(cfg.MW.SessionMiddleware)

	actionController := controllers.NewActionController(cfg.UC.ActionUseCase, cfg.UC.SchemaUseCase)

	actionRouter.HandleFunc("", actionController.GetAllActions).Methods(http.MethodGet)
	actionRouter.HandleFunc("", cfg.MW.RoleMiddleware(actionController.CreateAction, entities.AdminRole)).Methods(http.MethodPost)
	actionRouter.HandleFunc("/{id}", actionController.GetActionByID).Methods(http.MethodGet)
	actionRouter.HandleFunc("/{id}", cfg.MW.RoleMiddleware(actionController.UpdateActionByID, entities.AdminRole)).Methods(http.MethodPut)
	actionRouter.HandleFunc("/{id}", cfg.MW.RoleMiddleware(actionController.DeleteActionByID, entities.AdminRole)).Methods(http.MethodDelete)
}

func RegisterSessionRouter(cfg *SubRouterConfig) {
	sessionRouter := cfg.R.PathPrefix("/api/session").Subrouter()

	sessionController := controllers.NewSessionController(cfg.UC.UserUseCase)

	sessionRouter.HandleFunc("", sessionController.Login).Methods(http.MethodPost)
	sessionRouter.HandleFunc("", sessionController.GetSession).Methods(http.MethodGet)
}

func RegisterSetupRouter(cfg *SubRouterConfig) {
	setupRouter := cfg.R.PathPrefix("/api/setup").Subrouter()

	userController := controllers.NewUserController(cfg.UC.UserUseCase)

	setupRouter.HandleFunc("/register_admin", userController.CreateInitialAdmin).Methods(http.MethodPost)
}

func RegisterWebRouter(r *mux.Router) {
	r.PathPrefix("/").HandlerFunc(controllers.Web)
}
