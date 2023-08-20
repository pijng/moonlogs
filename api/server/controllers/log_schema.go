package controllers

import (
	"encoding/json"
	"fmt"
	"moonlogs/api/server/util"
	"moonlogs/ent"
	"moonlogs/internal/repository"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func LogSchemaCreate(w http.ResponseWriter, r *http.Request) {
	var newLogSchema ent.LogSchema

	err := json.NewDecoder(r.Body).Decode(&newLogSchema)
	if err != nil {
		util.Return(w, false, http.StatusBadRequest, err, nil)
		return
	}

	if len(newLogSchema.Fields) == 0 {
		error := fmt.Errorf("`fields` field is required")
		util.Return(w, false, http.StatusBadRequest, error, nil)
		return
	}

	createdLogSchema, err := repository.NewLogSchemaRepository(r.Context()).Create(newLogSchema)
	if err != nil {
		util.Return(w, false, http.StatusInternalServerError, err, nil)
		return
	}

	util.Return(w, true, http.StatusOK, nil, createdLogSchema)
}

func LogSchemaGetAll(w http.ResponseWriter, r *http.Request) {
	logSchemas, err := repository.NewLogSchemaRepository(r.Context()).GetAll()
	if err != nil {
		util.Return(w, false, http.StatusNotFound, err, nil)
		return
	}

	util.Return(w, true, http.StatusOK, nil, logSchemas)
}

func LogSchemaGetById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		error := fmt.Errorf("`id` path parameter is invalid")
		util.Return(w, false, http.StatusBadRequest, error, nil)
		return
	}

	logSchema, err := repository.NewLogSchemaRepository(r.Context()).GetById(id)
	if err != nil {
		util.Return(w, false, http.StatusNotFound, err, nil)
		return
	}

	util.Return(w, true, http.StatusOK, nil, logSchema)
}
