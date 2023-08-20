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

func LogRecordCreate(w http.ResponseWriter, r *http.Request) {
	var newLogRecord ent.LogRecord

	err := json.NewDecoder(r.Body).Decode(&newLogRecord)
	if err != nil {
		util.Return(w, false, http.StatusBadRequest, err, nil)
		return
	}

	if len(newLogRecord.Meta) == 0 {
		error := fmt.Errorf("`meta` field is required")
		util.Return(w, false, http.StatusBadRequest, error, nil)
		return
	}

	fmt.Println(newLogRecord.SchemaName)
	logSchema, err := repository.NewLogSchemaRepository(r.Context()).GetByName(newLogRecord.SchemaName)
	if err != nil {
		util.Return(w, false, http.StatusInternalServerError, err, nil)
		return
	}

	createdLogRecord, err := repository.NewLogRecordRepository(r.Context()).Create(newLogRecord, logSchema.ID)
	if err != nil {
		util.Return(w, false, http.StatusInternalServerError, err, nil)
		return
	}

	util.Return(w, true, http.StatusOK, nil, createdLogRecord)
}

func LogRecordGetAll(w http.ResponseWriter, r *http.Request) {
	limit, offset := util.Pagination(r)

	logRecords, err := repository.NewLogRecordRepository(r.Context()).GetAll(limit, offset)
	if err != nil {
		util.Return(w, false, http.StatusNotFound, err, nil)
		return
	}

	util.Return(w, true, http.StatusOK, nil, logRecords)
}

func LogRecordGetById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		error := fmt.Errorf("`id` path parameter is invalid")
		util.Return(w, false, http.StatusBadRequest, error, nil)
		return
	}

	logRecord, err := repository.NewLogRecordRepository(r.Context()).GetById(id)
	if err != nil {
		util.Return(w, false, http.StatusNotFound, err, nil)
		return
	}

	util.Return(w, true, http.StatusOK, nil, logRecord)
}

func LogRecordGetByMeta(w http.ResponseWriter, r *http.Request) {
	limit, offset := util.Pagination(r)

	var newLogRecord ent.LogRecord

	err := json.NewDecoder(r.Body).Decode(&newLogRecord)
	if err != nil {
		util.Return(w, false, http.StatusBadRequest, err, nil)
		return
	}

	logRecord, err := repository.
		NewLogRecordRepository(r.Context()).
		GetBySchemaAndMeta(newLogRecord.SchemaID, newLogRecord.SchemaName, newLogRecord.Meta, limit, offset)
	if err != nil {
		util.Return(w, false, http.StatusNotFound, err, nil)
		return
	}

	util.Return(w, true, http.StatusOK, nil, logRecord)
}
