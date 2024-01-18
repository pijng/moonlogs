package router

import (
	"context"
	"moonlogs/api/server/controllers"
	"moonlogs/api/server/session"
	"moonlogs/api/server/util"
	"moonlogs/internal/entities"
	"moonlogs/internal/repositories"
	"moonlogs/internal/usecases"
	"net/http"
	"slices"
	"strings"

	"github.com/gorilla/mux"
)

func RegisterSchemaRouter(r *mux.Router) {
	schemaRouter := r.PathPrefix("/api/schemas").Subrouter()
	schemaRouter.Use(SessionMiddleware)

	schemaRouter.HandleFunc("", controllers.GetAllSchemas).Methods(http.MethodGet)
	schemaRouter.HandleFunc("", roleMiddleware(controllers.CreateSchema, entities.AdminRole, entities.TokenRole)).Methods(http.MethodPost)
	schemaRouter.HandleFunc("/{id}", controllers.GetSchemaByID).Methods(http.MethodGet)
	schemaRouter.HandleFunc("/{id}", roleMiddleware(controllers.UpdateSchemaByID, entities.AdminRole, entities.TokenRole)).Methods(http.MethodPut)
	schemaRouter.HandleFunc("/{id}", roleMiddleware(controllers.DestroySchemaByID, entities.AdminRole)).Methods(http.MethodDelete)
	schemaRouter.HandleFunc("/search", controllers.GetSchemasByTitleOrDescription).Methods(http.MethodPost)
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
	userRouter.HandleFunc("/{id}", roleMiddleware(controllers.DestroyUserByID, entities.AdminRole)).Methods(http.MethodDelete)
}

func RegisterApiTokenRouter(r *mux.Router) {
	apiTokenRouter := r.PathPrefix("/api/api_tokens").Subrouter()
	apiTokenRouter.Use(SessionMiddleware)

	apiTokenRouter.HandleFunc("", roleMiddleware(controllers.GetAllApiTokens, entities.AdminRole)).Methods(http.MethodGet)
	apiTokenRouter.HandleFunc("", roleMiddleware(controllers.CreateApiToken, entities.AdminRole)).Methods(http.MethodPost)
	apiTokenRouter.HandleFunc("/{id}", roleMiddleware(controllers.GetApiTokenByID, entities.AdminRole)).Methods(http.MethodGet)
	apiTokenRouter.HandleFunc("/{id}", roleMiddleware(controllers.UpdateApiTokenByID, entities.AdminRole)).Methods(http.MethodPut)
	apiTokenRouter.HandleFunc("/{id}", roleMiddleware(controllers.DestroyApiTokenByID, entities.AdminRole)).Methods(http.MethodDelete)
}

func RegisterTagRouter(r *mux.Router) {
	apiTokenRouter := r.PathPrefix("/api/tags").Subrouter()
	apiTokenRouter.Use(SessionMiddleware)

	apiTokenRouter.HandleFunc("", controllers.GetAllTags).Methods(http.MethodGet)
	apiTokenRouter.HandleFunc("", roleMiddleware(controllers.CreateTag, entities.AdminRole)).Methods(http.MethodPost)
	apiTokenRouter.HandleFunc("/{id}", controllers.GetTagByID).Methods(http.MethodGet)
	apiTokenRouter.HandleFunc("/{id}", roleMiddleware(controllers.UpdateTagByID, entities.AdminRole)).Methods(http.MethodPut)
	apiTokenRouter.HandleFunc("/{id}", roleMiddleware(controllers.DestroyTagByID, entities.AdminRole)).Methods(http.MethodDelete)
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
		var bearerToken string

		reqAuth := r.Header.Get("Authorization")
		splitToken := strings.Split(reqAuth, "Bearer ")
		if len(splitToken) == 2 {
			bearerToken = splitToken[1]
		}

		apiTokenRepository := repositories.NewApiTokenRepository(r.Context())
		ok, err := usecases.NewApiTokenUseCase(apiTokenRepository).IsTokenValid(bearerToken)
		if err != nil {
			util.Return(w, false, http.StatusInternalServerError, nil, nil, util.Meta{})
			return
		}

		if ok {
			next.ServeHTTP(w, r)
			return
		}

		sessionCookie, err := session.GetSessionStore().Get(r, session.NAME)
		if err != nil {
			util.Return(w, false, http.StatusInternalServerError, nil, nil, util.Meta{})
			return
		}

		userRepository := repositories.NewUserRepository(r.Context())
		user, _ := usecases.NewUserUseCase(userRepository).GetUserByToken(bearerToken)

		token, ok := sessionCookie.Values["token"].(string)
		if !ok || user.ID == 0 {
			util.Return(w, false, http.StatusUnauthorized, nil, nil, util.Meta{})
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
		var bearerToken string

		reqAuth := r.Header.Get("Authorization")
		splitToken := strings.Split(reqAuth, "Bearer ")
		if len(splitToken) == 2 {
			bearerToken = splitToken[1]
		}

		apiTokenRepository := repositories.NewApiTokenRepository(r.Context())
		validAPIToken, err := usecases.NewApiTokenUseCase(apiTokenRepository).IsTokenValid(bearerToken)
		if err != nil {
			util.Return(w, false, http.StatusInternalServerError, nil, nil, util.Meta{})
			return
		}

		// allow access by API token for certain actions only
		if validAPIToken {
			if !slices.Contains(requiredRoles, entities.TokenRole) {
				util.Return(w, false, http.StatusForbidden, nil, nil, util.Meta{})
				return
			}

			next.ServeHTTP(w, r)
			return
		}

		userRepository := repositories.NewUserRepository(r.Context())
		user, _ := usecases.NewUserUseCase(userRepository).GetUserByToken(bearerToken)

		if user.ID == 0 {
			util.Return(w, false, http.StatusForbidden, nil, nil, util.Meta{})
			return
		}

		if !slices.Contains(requiredRoles, user.Role) {
			util.Return(w, false, http.StatusForbidden, nil, nil, util.Meta{})
			return
		}

		next(w, r)
	}
}
