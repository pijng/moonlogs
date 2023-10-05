package controllers

import (
	"encoding/json"
	"fmt"
	"moonlogs/api/server/util"
	"moonlogs/ent"
	"moonlogs/internal/repository"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func LogSchemaCreate(w http.ResponseWriter, r *http.Request) {
	var newLogSchema ent.LogSchema

	err := json.NewDecoder(r.Body).Decode(&newLogSchema)
	if err != nil {
		util.Return(w, false, http.StatusBadRequest, err, nil, util.Meta{})
		return
	}

	if len(newLogSchema.Fields) == 0 {
		error := fmt.Errorf("`fields` field is required")
		util.Return(w, false, http.StatusBadRequest, error, nil, util.Meta{})
		return
	}

	formattedSchemaName := strings.ReplaceAll(strings.ToLower(newLogSchema.Name), " ", "_")
	createdLogSchema, err := repository.NewLogSchemaRepository(r.Context()).Create(newLogSchema, formattedSchemaName)
	if err != nil {
		util.Return(w, false, http.StatusInternalServerError, err, nil, util.Meta{})
		return
	}

	util.Return(w, true, http.StatusOK, nil, createdLogSchema, util.Meta{})
}

func LogSchemaUpdateById(w http.ResponseWriter, r *http.Request) {
	var logSchemaToUpdate ent.LogSchema

	err := json.NewDecoder(r.Body).Decode(&logSchemaToUpdate)
	if err != nil {
		util.Return(w, false, http.StatusBadRequest, err, nil, util.Meta{})
		return
	}

	logSchema, err := repository.NewLogSchemaRepository(r.Context()).UpdateById(logSchemaToUpdate)
	if err != nil {
		util.Return(w, false, http.StatusInternalServerError, err, nil, util.Meta{})
		return
	}

	util.Return(w, true, http.StatusOK, nil, logSchema, util.Meta{})
}

func LogSchemaGetAll(w http.ResponseWriter, r *http.Request) {
	logSchemas, err := repository.NewLogSchemaRepository(r.Context()).GetAll()
	if err != nil {
		util.Return(w, false, http.StatusNotFound, err, nil, util.Meta{})
		return
	}

	util.Return(w, true, http.StatusOK, nil, logSchemas, util.Meta{})
}

func LogSchemaGetById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		error := fmt.Errorf("`id` path parameter is invalid")
		util.Return(w, false, http.StatusBadRequest, error, nil, util.Meta{})
		return
	}

	logSchema, err := repository.NewLogSchemaRepository(r.Context()).GetById(id)
	if err != nil {
		util.Return(w, false, http.StatusNotFound, err, nil, util.Meta{})
		return
	}

	util.Return(w, true, http.StatusOK, nil, logSchema, util.Meta{})
}

func LogSchemaGetByQuery(w http.ResponseWriter, r *http.Request) {
	var newLogSchema ent.LogSchema

	err := json.NewDecoder(r.Body).Decode(&newLogSchema)
	if err != nil {
		util.Return(w, false, http.StatusBadRequest, err, nil, util.Meta{})
		return
	}

	logSchemas, err := repository.
		NewLogSchemaRepository(r.Context()).
		GetByTitleOrDescriptionAll(newLogSchema.Title, newLogSchema.Description)

	if err != nil {
		util.Return(w, false, http.StatusNotFound, err, nil, util.Meta{})
		return
	}

	util.Return(w, true, http.StatusOK, nil, logSchemas, util.Meta{})
}
