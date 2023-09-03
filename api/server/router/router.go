package router

import (
	"moonlogs/api/server/controllers"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

func RegisterLogSchemaRouter(r *mux.Router, store *sessions.CookieStore) {
	logSchemaRouter := r.PathPrefix("/api/schemas").Subrouter()

	logSchemaRouter.HandleFunc("", SessionMiddleware(store, controllers.LogSchemaGetAll)).Methods(http.MethodGet, http.MethodOptions)
	logSchemaRouter.HandleFunc("", SessionMiddleware(store, controllers.LogSchemaCreate)).Methods(http.MethodPost, http.MethodOptions)
	logSchemaRouter.HandleFunc("/{id}", SessionMiddleware(store, controllers.LogSchemaGetById)).Methods(http.MethodGet, http.MethodOptions)
	logSchemaRouter.HandleFunc("/search", SessionMiddleware(store, controllers.LogSchemaGetByQuery)).Methods(http.MethodPost, http.MethodOptions)
}

func RegisterLogRecordRouter(r *mux.Router, store *sessions.CookieStore) {
	logRecordRouter := r.PathPrefix("/api/logs").Subrouter()

	logRecordRouter.HandleFunc("", SessionMiddleware(store, controllers.LogRecordGetAll)).Methods(http.MethodGet, http.MethodOptions)
	logRecordRouter.HandleFunc("", SessionMiddleware(store, controllers.LogRecordCreate)).Methods(http.MethodPost, http.MethodOptions)
	logRecordRouter.HandleFunc("/{id}", SessionMiddleware(store, controllers.LogRecordGetById)).Methods(http.MethodGet, http.MethodOptions)
	logRecordRouter.HandleFunc("/group/{schemaName}/{hash}", SessionMiddleware(store, controllers.LogRecordsByGroupHash)).Methods(http.MethodGet, http.MethodOptions)
	logRecordRouter.HandleFunc("/search", SessionMiddleware(store, controllers.LogRecordGetByQuery)).Methods(http.MethodPost, http.MethodOptions)
}

func RegisterUserRouter(r *mux.Router, store *sessions.CookieStore) {
	logRecordRouter := r.PathPrefix("/api/users").Subrouter()

	logRecordRouter.HandleFunc("", SessionMiddleware(store, controllers.UserGetAll)).Methods(http.MethodGet, http.MethodOptions)
	logRecordRouter.HandleFunc("", SessionMiddleware(store, controllers.UserCreate)).Methods(http.MethodPost, http.MethodOptions)
	logRecordRouter.HandleFunc("/{id}", SessionMiddleware(store, controllers.UserGetById)).Methods(http.MethodGet, http.MethodOptions)
	logRecordRouter.HandleFunc("/{id}", SessionMiddleware(store, controllers.UserDestroyById)).Methods(http.MethodDelete, http.MethodOptions)
}

func RegisterSessionRouter(r *mux.Router) {
	sessionRouter := r.PathPrefix("/api/session").Subrouter()

	sessionRouter.HandleFunc("", controllers.Login).Methods(http.MethodPost, http.MethodOptions)
	sessionRouter.HandleFunc("", controllers.GetSession).Methods(http.MethodGet, http.MethodOptions)
}

func RegisterWebRouter(r *mux.Router) {
	r.PathPrefix("/").HandlerFunc(controllers.Web)
}

func SessionMiddleware(store *sessions.CookieStore, next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Remove when auth is done
		next.ServeHTTP(w, r)

		// session, err := store.Get(r, sessionName)
		// if err != nil {
		// 	util.Return(w, false, http.StatusInternalServerError, fmt.Errorf("Session error: %w", err), nil, util.Meta{})
		// 	return
		// }

		// if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		// 	util.Return(w, false, http.StatusUnauthorized, fmt.Errorf("Unauthorized"), nil, util.Meta{})
		// 	return
		// }

		// ctx := context.WithValue(r.Context(), sessionKey, session)
		// next.ServeHTTP(w, r.WithContext(ctx))
	})
}
