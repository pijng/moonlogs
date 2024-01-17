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
	"strings"

	"github.com/gorilla/mux"
)

func RegisterSchemaRouter(r *mux.Router) {
	schemaRouter := r.PathPrefix("/api/schemas").Subrouter()
	schemaRouter.Use(SessionMiddleware)

	schemaRouter.HandleFunc("", controllers.GetAllSchemas).Methods(http.MethodGet)
	schemaRouter.HandleFunc("", RoleMiddleware(entities.AdminRole, controllers.CreateSchema)).Methods(http.MethodPost)
	schemaRouter.HandleFunc("/{id}", controllers.GetSchemaByID).Methods(http.MethodGet)
	schemaRouter.HandleFunc("/{id}", RoleMiddleware(entities.AdminRole, controllers.UpdateSchemaByID)).Methods(http.MethodPut)
	schemaRouter.HandleFunc("/{id}", RoleMiddleware(entities.AdminRole, controllers.DestroySchemaByID)).Methods(http.MethodDelete)
	schemaRouter.HandleFunc("/search", controllers.GetSchemasByTitleOrDescription).Methods(http.MethodPost)
}

func RegisterRecordRouter(r *mux.Router) {
	recordRouter := r.PathPrefix("/api/logs").Subrouter()
	// TODO: uncomment this
	// recordRouter.Use(SessionMiddleware)

	recordRouter.HandleFunc("", controllers.GetAllRecords).Methods(http.MethodGet)
	recordRouter.HandleFunc("", controllers.CreateRecord).Methods(http.MethodPost)
	recordRouter.HandleFunc("/{id}", controllers.GetRecordByID).Methods(http.MethodGet)
	recordRouter.HandleFunc("/group/{schemaName}/{hash}", controllers.GetRecordsByGroupHash).Methods(http.MethodGet)
	recordRouter.HandleFunc("/search", controllers.GetRecordsByQuery).Methods(http.MethodPost)
}

func RegisterUserRouter(r *mux.Router) {
	userRouter := r.PathPrefix("/api/users").Subrouter()
	userRouter.Use(SessionMiddleware)

	userRouter.HandleFunc("", controllers.GetAllUsers).Methods(http.MethodGet)
	userRouter.HandleFunc("", RoleMiddleware(entities.AdminRole, controllers.CreateUser)).Methods(http.MethodPost)
	userRouter.HandleFunc("/{id}", controllers.GetUserByID).Methods(http.MethodGet)
	userRouter.HandleFunc("/{id}", RoleMiddleware(entities.AdminRole, controllers.UpdateUserByID)).Methods(http.MethodPut)
	userRouter.HandleFunc("/{id}", RoleMiddleware(entities.AdminRole, controllers.DestroyUserByID)).Methods(http.MethodDelete)
}

func RegisterApiTokenRouter(r *mux.Router) {
	apiTokenRouter := r.PathPrefix("/api/api_tokens").Subrouter()
	apiTokenRouter.Use(SessionMiddleware)

	apiTokenRouter.HandleFunc("", RoleMiddleware(entities.AdminRole, controllers.GetAllApiTokens)).Methods(http.MethodGet)
	apiTokenRouter.HandleFunc("", RoleMiddleware(entities.AdminRole, controllers.CreateApiToken)).Methods(http.MethodPost)
	apiTokenRouter.HandleFunc("/{id}", RoleMiddleware(entities.AdminRole, controllers.GetApiTokenByID)).Methods(http.MethodGet)
	apiTokenRouter.HandleFunc("/{id}", RoleMiddleware(entities.AdminRole, controllers.UpdateApiTokenByID)).Methods(http.MethodPut)
	apiTokenRouter.HandleFunc("/{id}", RoleMiddleware(entities.AdminRole, controllers.DestroyApiTokenByID)).Methods(http.MethodDelete)
}

func RegisterTagRouter(r *mux.Router) {
	apiTokenRouter := r.PathPrefix("/api/tags").Subrouter()
	apiTokenRouter.Use(SessionMiddleware)

	apiTokenRouter.HandleFunc("", controllers.GetAllTags).Methods(http.MethodGet)
	apiTokenRouter.HandleFunc("", RoleMiddleware(entities.AdminRole, controllers.CreateTag)).Methods(http.MethodPost)
	apiTokenRouter.HandleFunc("/{id}", controllers.GetTagByID).Methods(http.MethodGet)
	apiTokenRouter.HandleFunc("/{id}", RoleMiddleware(entities.AdminRole, controllers.UpdateTagByID)).Methods(http.MethodPut)
	apiTokenRouter.HandleFunc("/{id}", RoleMiddleware(entities.AdminRole, controllers.DestroyTagByID)).Methods(http.MethodDelete)
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

func RoleMiddleware(requiredRole entities.Role, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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

		userRepository := repositories.NewUserRepository(r.Context())
		user, _ := usecases.NewUserUseCase(userRepository).GetUserByToken(bearerToken)

		if user.ID == 0 {
			util.Return(w, false, http.StatusForbidden, nil, nil, util.Meta{})
			return
		}

		if !strings.EqualFold(string(user.Role), string(requiredRole)) {
			util.Return(w, false, http.StatusForbidden, nil, nil, util.Meta{})
			return
		}

		next(w, r)
	}
}
