package controllers

import (
	"encoding/json"
	"fmt"
	"math"
	"moonlogs/api/server/util"
	"moonlogs/internal/entities"
	"moonlogs/internal/repositories"
	"moonlogs/internal/usecases"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CreateRecord(w http.ResponseWriter, r *http.Request) {
	var newRecord entities.Record

	err := json.NewDecoder(r.Body).Decode(&newRecord)
	if err != nil {
		util.Return(w, false, http.StatusBadRequest, err, nil, util.Meta{})
		return
	}

	schema, err := repositories.NewSchemaRepository(r.Context()).GetByName(newRecord.SchemaName)
	if err != nil {
		util.Return(w, false, http.StatusBadRequest, err, nil, util.Meta{})
		return
	}

	recordRepository := repositories.NewRecordRepository(r.Context())
	record, err := usecases.NewRecordUseCase(recordRepository).CreateRecord(newRecord, schema.ID)
	if err != nil {
		util.Return(w, false, http.StatusBadRequest, err, nil, util.Meta{})
		return
	}

	util.Return(w, true, http.StatusOK, nil, record, util.Meta{})
}

func GetAllRecords(w http.ResponseWriter, r *http.Request) {
	limit, offset, page := util.Pagination(r)

	recordRepository := repositories.NewRecordRepository(r.Context())
	recordUseCase := usecases.NewRecordUseCase(recordRepository)

	records, err := recordUseCase.GetAllRecords(limit, offset)
	if err != nil {
		util.Return(w, false, http.StatusBadRequest, err, nil, util.Meta{})
		return
	}

	count, err := recordUseCase.GetAllRecordsCount()
	if err != nil {
		util.Return(w, false, http.StatusBadRequest, err, nil, util.Meta{})
		return
	}

	pages := int(math.Ceil(float64(count) / float64(limit)))

	util.Return(w, true, http.StatusOK, nil, records, util.Meta{Page: page, Count: count, Pages: pages})
}

func GetRecordByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		error := fmt.Errorf("`id` path parameter is invalid")
		util.Return(w, false, http.StatusBadRequest, error, nil, util.Meta{})
		return
	}

	recordRepository := repositories.NewRecordRepository(r.Context())
	record, err := usecases.NewRecordUseCase(recordRepository).GetRecordByID(id)
	if err != nil {
		util.Return(w, false, http.StatusBadRequest, err, nil, util.Meta{})
		return
	}

	if record.ID == 0 {
		util.Return(w, false, http.StatusNotFound, err, nil, util.Meta{})
		return
	}

	util.Return(w, true, http.StatusOK, nil, record, util.Meta{})
}

func GetRecordsByQuery(w http.ResponseWriter, r *http.Request) {
	limit, offset, page := util.Pagination(r)

	var recordsToGet entities.Record

	err := json.NewDecoder(r.Body).Decode(&recordsToGet)
	if err != nil {
		util.Return(w, false, http.StatusBadRequest, err, nil, util.Meta{})
		return
	}

	recordRepository := repositories.NewRecordRepository(r.Context())
	recordUseCase := usecases.NewRecordUseCase(recordRepository)

	records, err := recordUseCase.GetRecordsByQuery(recordsToGet, limit, offset)
	if err != nil {
		util.Return(w, false, http.StatusBadRequest, err, nil, util.Meta{})
		return
	}

	count, err := recordUseCase.GetRecordsCountByQuery(recordsToGet)
	if err != nil {
		util.Return(w, false, http.StatusBadRequest, err, nil, util.Meta{})
		return
	}

	pages := int(math.Ceil(float64(count) / float64(limit)))

	util.Return(w, true, http.StatusOK, nil, records, util.Meta{Page: page, Count: count, Pages: pages})
}

func GetRecordsByGroupHash(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	groupHash := vars["hash"]
	schemaName := vars["schemaName"]

	schemaRepository := repositories.NewSchemaRepository(r.Context())
	schema, err := usecases.NewSchemaUseCase(schemaRepository).GetSchemaByName(schemaName)
	if err != nil || schema.ID == 0 {
		util.Return(w, false, http.StatusNotFound, err, nil, util.Meta{})
		return
	}

	recordRepository := repositories.NewRecordRepository(r.Context())
	records, err := usecases.NewRecordUseCase(recordRepository).GetRecordsByGroupHash(schemaName, groupHash)
	if err != nil {
		util.Return(w, false, http.StatusBadRequest, err, nil, util.Meta{})
		return
	}

	if len(records) == 0 {
		util.Return(w, false, http.StatusNotFound, err, nil, util.Meta{})
		return
	}

	util.Return(w, true, http.StatusOK, nil, records, util.Meta{})
}
