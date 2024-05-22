package controllers

import (
	"fmt"
	"moonlogs/internal/api/server/response"
	"moonlogs/internal/config"
	"moonlogs/internal/entities"
	"moonlogs/internal/lib/serialize"
	"moonlogs/internal/storage"
	"moonlogs/internal/usecases"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CreateAction(w http.ResponseWriter, r *http.Request) {
	var newAction entities.Action

	err := serialize.NewJSONDecoder(r.Body).Decode(&newAction)
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, err, nil, response.Meta{})
		return
	}

	schemaStorage := storage.NewSchemaStorage(r.Context(), config.Get().DBAdapter)
	schema, err := usecases.NewSchemaUseCase(schemaStorage).GetSchemaByID(newAction.SchemaID)
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, err, nil, response.Meta{})
		return
	}

	if schema.ID == 0 {
		err = fmt.Errorf("provided schema is not found: %s", newAction.SchemaName)
		response.Return(w, false, http.StatusBadRequest, err, nil, response.Meta{})
		return
	}

	newAction.SchemaName = schema.Name

	actionStorage := storage.NewActionStorage(r.Context(), config.Get().DBAdapter)
	action, err := usecases.NewActionUseCase(actionStorage).CreateAction(newAction)
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, err, nil, response.Meta{})
		return
	}

	response.Return(w, true, http.StatusOK, nil, action, response.Meta{})
}

func DeleteActionByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, fmt.Errorf("`id` path parameter is invalid: %w", err), nil, response.Meta{})
		return
	}

	actionStorage := storage.NewActionStorage(r.Context(), config.Get().DBAdapter)
	err = usecases.NewActionUseCase(actionStorage).DeleteActionByID(id)
	if err != nil {
		response.Return(w, false, http.StatusInternalServerError, err, nil, response.Meta{})
		return
	}

	response.Return(w, true, http.StatusOK, nil, id, response.Meta{})
}

func GetActionByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, fmt.Errorf("`id` path parameter is invalid: %w", err), nil, response.Meta{})
		return
	}

	actionStorage := storage.NewActionStorage(r.Context(), config.Get().DBAdapter)
	action, err := usecases.NewActionUseCase(actionStorage).GetActionByID(id)
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, err, nil, response.Meta{})
		return
	}

	if action.ID == 0 {
		response.Return(w, false, http.StatusNotFound, err, nil, response.Meta{})
		return
	}

	response.Return(w, true, http.StatusOK, nil, action, response.Meta{})
}

func UpdateActionByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, fmt.Errorf("`id` path parameter is invalid: %w", err), nil, response.Meta{})
		return
	}

	var actionToUpdate entities.Action

	err = serialize.NewJSONDecoder(r.Body).Decode(&actionToUpdate)
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, err, nil, response.Meta{})
		return
	}

	actionStorage := storage.NewActionStorage(r.Context(), config.Get().DBAdapter)
	action, err := usecases.NewActionUseCase(actionStorage).UpdateActionByID(id, actionToUpdate)
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, err, nil, response.Meta{})
		return
	}

	if action == nil {
		response.Return(w, false, http.StatusNotFound, err, nil, response.Meta{})
		return
	}

	response.Return(w, true, http.StatusOK, nil, action, response.Meta{})
}

func GetAllActions(w http.ResponseWriter, r *http.Request) {
	actionStorage := storage.NewActionStorage(r.Context(), config.Get().DBAdapter)
	actions, err := usecases.NewActionUseCase(actionStorage).GetAllActions()
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, err, nil, response.Meta{})
		return
	}

	response.Return(w, true, http.StatusOK, nil, actions, response.Meta{})
}
