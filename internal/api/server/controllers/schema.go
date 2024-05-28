package controllers

import (
	"fmt"
	"moonlogs/internal/api/server/access"
	"moonlogs/internal/api/server/response"
	"moonlogs/internal/api/server/session"
	"moonlogs/internal/entities"
	"moonlogs/internal/lib/serialize"
	"moonlogs/internal/usecases"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type SchemaController struct {
	schemaUseCase *usecases.SchemaUseCase
}

func NewSchemaController(schemaUseCase *usecases.SchemaUseCase) *SchemaController {
	return &SchemaController{
		schemaUseCase: schemaUseCase,
	}
}

func (c *SchemaController) CreateSchema(w http.ResponseWriter, r *http.Request) {
	var newSchema entities.Schema

	err := serialize.NewJSONDecoder(r.Body).Decode(&newSchema)
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, err, nil, response.Meta{})
		return
	}

	schema, err := c.schemaUseCase.CreateSchema(r.Context(), newSchema)
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, err, nil, response.Meta{})
		return
	}

	response.Return(w, true, http.StatusOK, nil, schema, response.Meta{})
}

func (c *SchemaController) UpdateSchemaByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		error := fmt.Errorf("`id` path parameter is invalid")
		response.Return(w, false, http.StatusBadRequest, error, nil, response.Meta{})
		return
	}

	var schemaToUpdate entities.Schema

	err = serialize.NewJSONDecoder(r.Body).Decode(&schemaToUpdate)
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, err, nil, response.Meta{})
		return
	}

	schema, err := c.schemaUseCase.UpdateSchemaByID(r.Context(), id, schemaToUpdate)
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

func (c *SchemaController) GetAllSchemas(w http.ResponseWriter, r *http.Request) {
	user := session.GetUserFromContext(r)

	schemas, err := c.schemaUseCase.GetAllSchemas(r.Context(), user)
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, err, nil, response.Meta{})
		return
	}

	response.Return(w, true, http.StatusOK, nil, schemas, response.Meta{})
}

func (c *SchemaController) GetSchemaByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		error := fmt.Errorf("`id` path parameter is invalid")
		response.Return(w, false, http.StatusBadRequest, error, nil, response.Meta{})
		return
	}

	schema, err := c.schemaUseCase.GetSchemaByID(r.Context(), id)
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, err, nil, response.Meta{})
		return
	}

	if schema.ID == 0 {
		response.Return(w, false, http.StatusNotFound, err, nil, response.Meta{})
		return
	}

	if access.IsSchemaForbiddenForUser(c.schemaUseCase, schema.Name, r) {
		response.Return(w, false, http.StatusForbidden, nil, nil, response.Meta{})
		return
	}

	response.Return(w, true, http.StatusOK, nil, schema, response.Meta{})
}

func (c *SchemaController) DeleteSchemaByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, fmt.Errorf("`id` path parameter is invalid: %w", err), nil, response.Meta{})
		return
	}

	err = c.schemaUseCase.DeleteSchemaByID(r.Context(), id)
	if err != nil {
		response.Return(w, false, http.StatusInternalServerError, err, nil, response.Meta{})
		return
	}

	response.Return(w, true, http.StatusOK, nil, id, response.Meta{})
}
