package router

import (
	"moonlogs/api/server/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterLogSchemaRouter(r *mux.Router) {
	logSchemaRouter := r.PathPrefix("/api/schemas").Subrouter()

	logSchemaRouter.HandleFunc("", controllers.LogSchemaGetAll).Methods(http.MethodGet, http.MethodOptions)
	logSchemaRouter.HandleFunc("", controllers.LogSchemaCreate).Methods(http.MethodPost, http.MethodOptions)
	logSchemaRouter.HandleFunc("/{id}", controllers.LogSchemaGetById).Methods(http.MethodGet, http.MethodOptions)
	logSchemaRouter.HandleFunc("/search", controllers.LogSchemaGetByQuery).Methods(http.MethodPost, http.MethodOptions)
}

func RegisterLogRecordRouter(r *mux.Router) {
	logRecordRouter := r.PathPrefix("/api/logs").Subrouter()

	logRecordRouter.HandleFunc("", controllers.LogRecordGetAll).Methods(http.MethodGet, http.MethodOptions)
	logRecordRouter.HandleFunc("", controllers.LogRecordCreate).Methods(http.MethodPost, http.MethodOptions)
	logRecordRouter.HandleFunc("/{id}", controllers.LogRecordGetById).Methods(http.MethodGet, http.MethodOptions)
	logRecordRouter.HandleFunc("/group/{schemaName}/{hash}", controllers.LogRecordsByGroupHash).Methods(http.MethodGet, http.MethodOptions)
	logRecordRouter.HandleFunc("/search", controllers.LogRecordGetByQuery).Methods(http.MethodPost, http.MethodOptions)
}

func RegisterWebRouter(r *mux.Router) {
	r.PathPrefix("/").HandlerFunc(controllers.Web)
}
