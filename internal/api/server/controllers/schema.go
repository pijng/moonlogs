package controllers

import (
	"encoding/json"
	"fmt"
	"moonlogs/internal/api/server/util"
	"moonlogs/internal/entities"
	"moonlogs/internal/repositories"
	"moonlogs/internal/usecases"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CreateSchema(w http.ResponseWriter, r *http.Request) {
	var newSchema entities.Schema

	err := json.NewDecoder(r.Body).Decode(&newSchema)
	if err != nil {
		util.Return(w, false, http.StatusBadRequest, err, nil, util.Meta{})
		return
	}

	schemaRepository := repositories.NewSchemaRepository(r.Context())
	schema, err := usecases.NewSchemaUseCase(schemaRepository).CreateSchema(newSchema)
	if err != nil {
		util.Return(w, false, http.StatusBadRequest, err, nil, util.Meta{})
		return
	}

	util.Return(w, true, http.StatusOK, nil, schema, util.Meta{})
}

func UpdateSchemaByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		error := fmt.Errorf("`id` path parameter is invalid")
		util.Return(w, false, http.StatusBadRequest, error, nil, util.Meta{})
		return
	}

	var schemaToUpdate entities.Schema

	err = json.NewDecoder(r.Body).Decode(&schemaToUpdate)
	if err != nil {
		util.Return(w, false, http.StatusBadRequest, err, nil, util.Meta{})
		return
	}

	schemaRepository := repositories.NewSchemaRepository(r.Context())
	schema, err := usecases.NewSchemaUseCase(schemaRepository).UpdateSchemaByID(id, schemaToUpdate)
	if err != nil {
		util.Return(w, false, http.StatusBadRequest, err, nil, util.Meta{})
		return
	}

	if schema == nil {
		util.Return(w, false, http.StatusNotFound, err, nil, util.Meta{})
		return
	}

	util.Return(w, true, http.StatusOK, nil, schema, util.Meta{})
}

func GetAllSchemas(w http.ResponseWriter, r *http.Request) {
	schemaRepository := repositories.NewSchemaRepository(r.Context())
	schemas, err := usecases.NewSchemaUseCase(schemaRepository).GetAllSchemas()
	if err != nil {
		util.Return(w, false, http.StatusBadRequest, err, nil, util.Meta{})
		return
	}

	util.Return(w, true, http.StatusOK, nil, schemas, util.Meta{})
}

func GetSchemaByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		error := fmt.Errorf("`id` path parameter is invalid")
		util.Return(w, false, http.StatusBadRequest, error, nil, util.Meta{})
		return
	}

	schemaRepository := repositories.NewSchemaRepository(r.Context())
	schema, err := usecases.NewSchemaUseCase(schemaRepository).GetSchemaByID(id)
	if err != nil {
		util.Return(w, false, http.StatusBadRequest, err, nil, util.Meta{})
		return
	}

	if schema.ID == 0 {
		util.Return(w, false, http.StatusNotFound, err, nil, util.Meta{})
		return
	}

	util.Return(w, true, http.StatusOK, nil, schema, util.Meta{})
}

func GetSchemasByTitleOrDescription(w http.ResponseWriter, r *http.Request) {
	var schemaToGet entities.Schema

	err := json.NewDecoder(r.Body).Decode(&schemaToGet)
	if err != nil {
		util.Return(w, false, http.StatusBadRequest, err, nil, util.Meta{})
		return
	}

	schemaRepository := repositories.NewSchemaRepository(r.Context())
	schemas, err := usecases.NewSchemaUseCase(schemaRepository).GetSchemasByTitleOrDescription(schemaToGet.Title, schemaToGet.Description)

	if err != nil {
		util.Return(w, false, http.StatusNotFound, err, nil, util.Meta{})
		return
	}

	util.Return(w, true, http.StatusOK, nil, schemas, util.Meta{})
}

func DestroySchemaByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		util.Return(w, false, http.StatusBadRequest, fmt.Errorf("`id` path parameter is invalid: %w", err), nil, util.Meta{})
		return
	}

	schemaRepository := repositories.NewSchemaRepository(r.Context())
	err = usecases.NewSchemaUseCase(schemaRepository).DestroySchemaByID(id)
	if err != nil {
		util.Return(w, false, http.StatusInternalServerError, err, nil, util.Meta{})
		return
	}

	util.Return(w, true, http.StatusOK, nil, id, util.Meta{})
}
