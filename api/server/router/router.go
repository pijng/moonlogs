package router

import (
	"moonlogs/api/server/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterLogSchemaRouter(r *mux.Router) {
	logSchemaRouter := r.PathPrefix("/api/schemas").Subrouter()

	logSchemaRouter.HandleFunc("", controllers.LogSchemaGetAll).Methods(http.MethodGet)
	logSchemaRouter.HandleFunc("", controllers.LogSchemaCreate).Methods(http.MethodPost)
	logSchemaRouter.HandleFunc("/{id}", controllers.LogSchemaGetById).Methods(http.MethodGet)
}

func RegisterLogRecordRouter(r *mux.Router) {
	logRecordRouter := r.PathPrefix("/api/records").Subrouter()

	logRecordRouter.HandleFunc("", controllers.LogRecordGetAll).Methods(http.MethodGet)
	logRecordRouter.HandleFunc("", controllers.LogRecordCreate).Methods(http.MethodPost)
	logRecordRouter.HandleFunc("/{id}", controllers.LogRecordGetById).Methods(http.MethodGet)
	logRecordRouter.HandleFunc("/search", controllers.LogRecordGetByMeta).Methods(http.MethodPost)
}

func RegisterWebRouter(r *mux.Router) {
	r.PathPrefix("/").HandlerFunc(controllers.Web)
}
