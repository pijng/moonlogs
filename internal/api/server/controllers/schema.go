package controllers

import (
	"encoding/json"
	"fmt"
	"moonlogs/internal/api/server/access"
	"moonlogs/internal/api/server/response"
	"moonlogs/internal/api/server/session"
	"moonlogs/internal/config"
	"moonlogs/internal/entities"
	"moonlogs/internal/storage"
	"moonlogs/internal/usecases"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CreateSchema(w http.ResponseWriter, r *http.Request) {
	var newSchema entities.Schema

	err := json.NewDecoder(r.Body).Decode(&newSchema)
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, err, nil, response.Meta{})
		return
	}

	schemaStorage := storage.NewSchemaStorage(r.Context(), config.Get().DBAdapter)
	schema, err := usecases.NewSchemaUseCase(schemaStorage).CreateSchema(newSchema)
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, err, nil, response.Meta{})
		return
	}

	response.Return(w, true, http.StatusOK, nil, schema, response.Meta{})
}

func UpdateSchemaByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		error := fmt.Errorf("`id` path parameter is invalid")
		response.Return(w, false, http.StatusBadRequest, error, nil, response.Meta{})
		return
	}

	var schemaToUpdate entities.Schema

	err = json.NewDecoder(r.Body).Decode(&schemaToUpdate)
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, err, nil, response.Meta{})
		return
	}

	schemaStorage := storage.NewSchemaStorage(r.Context(), config.Get().DBAdapter)
	schema, err := usecases.NewSchemaUseCase(schemaStorage).UpdateSchemaByID(id, schemaToUpdate)
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, err, nil, response.Meta{})
		return
	}

	if schema == nil {
		response.Return(w, false, http.StatusNotFound, err, nil, response.Meta{})
		return
	}

	response.Return(w, true, http.StatusOK, nil, schema, response.Meta{})
}

func GetAllSchemas(w http.ResponseWriter, r *http.Request) {
	user := session.GetUserFromContext(r)

	schemaStorage := storage.NewSchemaStorage(r.Context(), config.Get().DBAdapter)
	schemas, err := usecases.NewSchemaUseCase(schemaStorage).GetAllSchemas(user)
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, err, nil, response.Meta{})
		return
	}

	response.Return(w, true, http.StatusOK, nil, schemas, response.Meta{})
}

func GetSchemaByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		error := fmt.Errorf("`id` path parameter is invalid")
		response.Return(w, false, http.StatusBadRequest, error, nil, response.Meta{})
		return
	}

	schemaStorage := storage.NewSchemaStorage(r.Context(), config.Get().DBAdapter)
	schema, err := usecases.NewSchemaUseCase(schemaStorage).GetSchemaByID(id)
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, err, nil, response.Meta{})
		return
	}

	if schema.ID == 0 {
		response.Return(w, false, http.StatusNotFound, err, nil, response.Meta{})
		return
	}

	if access.IsSchemaForbiddenForUser(schema.Name, r) {
		response.Return(w, false, http.StatusForbidden, nil, nil, response.Meta{})
		return
	}

	response.Return(w, true, http.StatusOK, nil, schema, response.Meta{})
}

func DeleteSchemaByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, fmt.Errorf("`id` path parameter is invalid: %w", err), nil, response.Meta{})
		return
	}

	schemaStorage := storage.NewSchemaStorage(r.Context(), config.Get().DBAdapter)
	err = usecases.NewSchemaUseCase(schemaStorage).DeleteSchemaByID(id)
	if err != nil {
		response.Return(w, false, http.StatusInternalServerError, err, nil, response.Meta{})
		return
	}

	response.Return(w, true, http.StatusOK, nil, id, response.Meta{})
}
