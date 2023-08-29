package controllers

import (
	"encoding/json"
	"fmt"
	"hash/fnv"
	"math"
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
		util.Return(w, false, http.StatusBadRequest, err, nil, util.Meta{})
		return
	}

	if len(newLogRecord.Query) == 0 {
		error := fmt.Errorf("`query` field is required")
		util.Return(w, false, http.StatusBadRequest, error, nil, util.Meta{})
		return
	}

	logSchema, err := repository.NewLogSchemaRepository(r.Context()).GetByName(newLogRecord.SchemaName)
	if err != nil {
		util.Return(w, false, http.StatusInternalServerError, err, nil, util.Meta{})
		return
	}

	bytes, err := json.Marshal(newLogRecord.Query)
	if err != nil {
		util.Return(w, false, http.StatusInternalServerError, err, nil, util.Meta{})
		return
	}

	hasher := fnv.New64a()
	hasher.Write(bytes)
	hashSum := hasher.Sum64()

	groupHash := fmt.Sprint(hashSum)

	createdLogRecord, err := repository.NewLogRecordRepository(r.Context()).Create(newLogRecord, logSchema.ID, groupHash)
	if err != nil {
		util.Return(w, false, http.StatusInternalServerError, err, nil, util.Meta{})
		return
	}

	util.Return(w, true, http.StatusOK, nil, createdLogRecord, util.Meta{})
}

func LogRecordGetAll(w http.ResponseWriter, r *http.Request) {
	limit, offset, page := util.Pagination(r)

	logRecords, err := repository.NewLogRecordRepository(r.Context()).GetAll(limit, offset)
	if err != nil {
		util.Return(w, false, http.StatusNotFound, err, nil, util.Meta{})
		return
	}

	count, err := repository.NewLogRecordRepository(r.Context()).GetCountAll()
	if err != nil {
		util.Return(w, false, http.StatusInternalServerError, err, nil, util.Meta{})
		return
	}

	pages := int(math.Ceil(float64(count) / float64(limit)))

	util.Return(w, true, http.StatusOK, nil, logRecords, util.Meta{Page: page, Count: count, Pages: pages})
}

func LogRecordGetById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		error := fmt.Errorf("`id` path parameter is invalid")
		util.Return(w, false, http.StatusBadRequest, error, nil, util.Meta{})
		return
	}

	logRecord, err := repository.NewLogRecordRepository(r.Context()).GetById(id)
	if err != nil {
		util.Return(w, false, http.StatusNotFound, err, nil, util.Meta{})
		return
	}

	util.Return(w, true, http.StatusOK, nil, logRecord, util.Meta{})
}

func LogRecordGetByQuery(w http.ResponseWriter, r *http.Request) {
	limit, offset, page := util.Pagination(r)

	var newLogRecord ent.LogRecord

	err := json.NewDecoder(r.Body).Decode(&newLogRecord)
	if err != nil {
		util.Return(w, false, http.StatusBadRequest, err, nil, util.Meta{})
		return
	}

	logRecord, err := repository.
		NewLogRecordRepository(r.Context()).
		GetBySchemaAndQuery(newLogRecord.SchemaID, newLogRecord.SchemaName, newLogRecord.Text, newLogRecord.Query, limit, offset)

	if err != nil {
		util.Return(w, false, http.StatusNotFound, err, nil, util.Meta{})
		return
	}

	count, err := repository.NewLogRecordRepository(r.Context()).GetCountBySchemaAndQuery(newLogRecord.SchemaName, newLogRecord.Text, newLogRecord.Query)
	if err != nil {
		util.Return(w, false, http.StatusInternalServerError, err, nil, util.Meta{})
		return
	}

	pages := int(math.Ceil(float64(count) / float64(limit)))

	util.Return(w, true, http.StatusOK, nil, logRecord, util.Meta{Page: page, Count: count, Pages: pages})
}

func LogRecordsByGroupHash(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	groupHash := vars["hash"]
	schemaName := vars["schemaName"]

	logRecords, err := repository.NewLogRecordRepository(r.Context()).GetByGroupHash(schemaName, groupHash)
	if err != nil {
		util.Return(w, false, http.StatusNotFound, err, nil, util.Meta{})
		return
	}

	util.Return(w, true, http.StatusOK, nil, logRecords, util.Meta{})
}
