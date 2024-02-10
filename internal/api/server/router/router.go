package router

import (
	"context"
	"moonlogs/internal/api/server/controllers"
	"moonlogs/internal/api/server/response"
	"moonlogs/internal/api/server/session"
	"moonlogs/internal/config"
	"moonlogs/internal/entities"
	"moonlogs/internal/shared"
	"moonlogs/internal/storage"
	"moonlogs/internal/usecases"
	"net/http"
	"slices"

	"github.com/gorilla/mux"
)

func RegisterSchemaRouter(r *mux.Router) {
	schemaRouter := r.PathPrefix("/api/schemas").Subrouter()
	schemaRouter.Use(SessionMiddleware)

	schemaRouter.HandleFunc("", controllers.GetAllSchemas).Methods(http.MethodGet)
	schemaRouter.HandleFunc("", roleMiddleware(controllers.CreateSchema, entities.AdminRole, entities.TokenRole)).Methods(http.MethodPost)
	schemaRouter.HandleFunc("/{id}", controllers.GetSchemaByID).Methods(http.MethodGet)
	schemaRouter.HandleFunc("/{id}", roleMiddleware(controllers.UpdateSchemaByID, entities.AdminRole, entities.TokenRole)).Methods(http.MethodPut)
	schemaRouter.HandleFunc("/{id}", roleMiddleware(controllers.DeleteSchemaByID, entities.AdminRole)).Methods(http.MethodDelete)
}

func RegisterRecordRouter(r *mux.Router) {
	recordRouter := r.PathPrefix("/api/logs").Subrouter()
	recordRouter.Use(SessionMiddleware)

	recordRouter.HandleFunc("", controllers.GetAllRecords).Methods(http.MethodGet)
	recordRouter.HandleFunc("", roleMiddleware(controllers.CreateRecord, entities.AdminRole, entities.TokenRole)).Methods(http.MethodPost)
	recordRouter.HandleFunc("/{id}", controllers.GetRecordByID).Methods(http.MethodGet)
	recordRouter.HandleFunc("/group/{schemaName}/{hash}", controllers.GetRecordsByGroupHash).Methods(http.MethodGet)
	recordRouter.HandleFunc("/search", controllers.GetRecordsByQuery).Methods(http.MethodPost)
}

func RegisterUserRouter(r *mux.Router) {
	userRouter := r.PathPrefix("/api/users").Subrouter()
	userRouter.Use(SessionMiddleware)

	userRouter.HandleFunc("", controllers.GetAllUsers).Methods(http.MethodGet)
	userRouter.HandleFunc("", roleMiddleware(controllers.CreateUser, entities.AdminRole)).Methods(http.MethodPost)
	userRouter.HandleFunc("/{id}", controllers.GetUserByID).Methods(http.MethodGet)
	userRouter.HandleFunc("/{id}", roleMiddleware(controllers.UpdateUserByID, entities.AdminRole)).Methods(http.MethodPut)
	userRouter.HandleFunc("/{id}", roleMiddleware(controllers.DeleteUserByID, entities.AdminRole)).Methods(http.MethodDelete)
}

func RegisterApiTokenRouter(r *mux.Router) {
	apiTokenRouter := r.PathPrefix("/api/api_tokens").Subrouter()
	apiTokenRouter.Use(SessionMiddleware)

	apiTokenRouter.HandleFunc("", roleMiddleware(controllers.GetAllApiTokens, entities.AdminRole)).Methods(http.MethodGet)
	apiTokenRouter.HandleFunc("", roleMiddleware(controllers.CreateApiToken, entities.AdminRole)).Methods(http.MethodPost)
	apiTokenRouter.HandleFunc("/{id}", roleMiddleware(controllers.GetApiTokenByID, entities.AdminRole)).Methods(http.MethodGet)
	apiTokenRouter.HandleFunc("/{id}", roleMiddleware(controllers.UpdateApiTokenByID, entities.AdminRole)).Methods(http.MethodPut)
	apiTokenRouter.HandleFunc("/{id}", roleMiddleware(controllers.DeleteApiTokenByID, entities.AdminRole)).Methods(http.MethodDelete)
}

func RegisterTagRouter(r *mux.Router) {
	apiTokenRouter := r.PathPrefix("/api/tags").Subrouter()
	apiTokenRouter.Use(SessionMiddleware)

	apiTokenRouter.HandleFunc("", controllers.GetAllTags).Methods(http.MethodGet)
	apiTokenRouter.HandleFunc("", roleMiddleware(controllers.CreateTag, entities.AdminRole)).Methods(http.MethodPost)
	apiTokenRouter.HandleFunc("/{id}", controllers.GetTagByID).Methods(http.MethodGet)
	apiTokenRouter.HandleFunc("/{id}", roleMiddleware(controllers.UpdateTagByID, entities.AdminRole)).Methods(http.MethodPut)
	apiTokenRouter.HandleFunc("/{id}", roleMiddleware(controllers.DeleteTagByID, entities.AdminRole)).Methods(http.MethodDelete)
}

func RegisterSessionRouter(r *mux.Router) {
	sessionRouter := r.PathPrefix("/api/session").Subrouter()

	sessionRouter.HandleFunc("", controllers.Login).Methods(http.MethodPost)
	sessionRouter.HandleFunc("", controllers.GetSession).Methods(http.MethodGet)
}

func RegisterSetupRouter(r *mux.Router) {
	setupRouter := r.PathPrefix("/api/setup").Subrouter()

	setupRouter.HandleFunc("/register_admin", controllers.CreateInitialAdmin).Methods(http.MethodPost)
}

func RegisterWebRouter(r *mux.Router) {
	r.PathPrefix("/").HandlerFunc(controllers.Web)
}

func SessionMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bearerToken := shared.ExtractBearerToken(r)

		apiTokenStorage := storage.NewApiTokenStorage(r.Context(), config.Get().DBAdapter)
		ok, err := usecases.NewApiTokenUseCase(apiTokenStorage).IsTokenValid(bearerToken)
		if err != nil {
			response.Return(w, false, http.StatusInternalServerError, nil, nil, response.Meta{})
			return
		}

		if ok {
			next.ServeHTTP(w, r)
			return
		}

		sessionCookie, err := session.GetSessionStore().Get(r, session.NAME)
		if err != nil {
			response.Return(w, false, http.StatusInternalServerError, nil, nil, response.Meta{})
			return
		}

		userStorage := storage.NewUserStorage(r.Context(), config.Get().DBAdapter)
		user, _ := usecases.NewUserUseCase(userStorage).GetUserByToken(bearerToken)

		token, ok := sessionCookie.Values["token"].(string)
		if !ok || user.ID == 0 || bool(user.IsRevoked) {
			response.Return(w, false, http.StatusUnauthorized, nil, nil, response.Meta{})
			return
		}

		if token == "" {
			token = user.Token
			sessionCookie.Values["token"] = token
			sessionCookie.Values["userID"] = user.ID
		}

		ctx := context.WithValue(r.Context(), session.KEY, sessionCookie)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func roleMiddleware(next http.HandlerFunc, requiredRoles ...entities.Role) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		bearerToken := shared.ExtractBearerToken(r)

		apiTokenStorage := storage.NewApiTokenStorage(r.Context(), config.Get().DBAdapter)
		validAPIToken, err := usecases.NewApiTokenUseCase(apiTokenStorage).IsTokenValid(bearerToken)
		if err != nil {
			response.Return(w, false, http.StatusInternalServerError, nil, nil, response.Meta{})
			return
		}

		// allow access by API token for certain actions only
		if validAPIToken {
			if !slices.Contains(requiredRoles, entities.TokenRole) {
				response.Return(w, false, http.StatusForbidden, nil, nil, response.Meta{})
				return
			}

			next.ServeHTTP(w, r)
			return
		}

		userStorage := storage.NewUserStorage(r.Context(), config.Get().DBAdapter)
		user, _ := usecases.NewUserUseCase(userStorage).GetUserByToken(bearerToken)

		if user.ID == 0 {
			response.Return(w, false, http.StatusForbidden, nil, nil, response.Meta{})
			return
		}

		if !slices.Contains(requiredRoles, user.Role) {
			response.Return(w, false, http.StatusForbidden, nil, nil, response.Meta{})
			return
		}

		next(w, r)
	}
}
