package controllers

import (
	"fmt"
	"moonlogs/internal/api/server/response"
	"moonlogs/internal/entities"
	"moonlogs/internal/lib/serialize"
	"moonlogs/internal/usecases"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type ActionController struct {
	actionUseCase *usecases.ActionUseCase
	schemaUseCase *usecases.SchemaUseCase
}

func NewActionController(actionUseCase *usecases.ActionUseCase, schemaUseCase *usecases.SchemaUseCase) *ActionController {
	return &ActionController{
		actionUseCase: actionUseCase,
		schemaUseCase: schemaUseCase,
	}
}

func (c *ActionController) CreateAction(w http.ResponseWriter, r *http.Request) {
	var newAction entities.Action

	err := serialize.NewJSONDecoder(r.Body).Decode(&newAction)
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, err, nil, response.Meta{})
		return
	}

	for _, id := range newAction.SchemaIDs {
		schema, err := c.schemaUseCase.GetSchemaByID(r.Context(), id)
		if err != nil {
			response.Return(w, false, http.StatusBadRequest, err, nil, response.Meta{})
			return
		}

		if schema.ID == 0 {
			err = fmt.Errorf("provided schema is not found by id: %d", id)
			response.Return(w, false, http.StatusBadRequest, err, nil, response.Meta{})
			return
		}
	}

	action, err := c.actionUseCase.CreateAction(r.Context(), newAction)
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, err, nil, response.Meta{})
		return
	}

	response.Return(w, true, http.StatusOK, nil, action, response.Meta{})
}

func (c *ActionController) DeleteActionByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, fmt.Errorf("`id` path parameter is invalid: %w", err), nil, response.Meta{})
		return
	}

	err = c.actionUseCase.DeleteActionByID(r.Context(), id)
	if err != nil {
		response.Return(w, false, http.StatusInternalServerError, err, nil, response.Meta{})
		return
	}

	response.Return(w, true, http.StatusOK, nil, id, response.Meta{})
}

func (c *ActionController) GetActionByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, fmt.Errorf("`id` path parameter is invalid: %w", err), nil, response.Meta{})
		return
	}

	action, err := c.actionUseCase.GetActionByID(r.Context(), id)
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

func (c *ActionController) UpdateActionByID(w http.ResponseWriter, r *http.Request) {
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

	action, err := c.actionUseCase.UpdateActionByID(r.Context(), id, actionToUpdate)
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

func (c *ActionController) GetAllActions(w http.ResponseWriter, r *http.Request) {
	actions, err := c.actionUseCase.GetAllActions(r.Context())
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, err, nil, response.Meta{})
		return
	}

	response.Return(w, true, http.StatusOK, nil, actions, response.Meta{})
}
