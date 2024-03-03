package controllers

import (
	"context"
	"fmt"
	"math"
	"moonlogs/internal/api/server/access"
	"moonlogs/internal/api/server/pagination"
	"moonlogs/internal/api/server/response"
	"moonlogs/internal/api/server/session"
	"moonlogs/internal/api/server/timerange"
	"moonlogs/internal/config"
	"moonlogs/internal/entities"
	"moonlogs/internal/lib/serialize"
	"moonlogs/internal/storage"
	"moonlogs/internal/usecases"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

const (
	AsyncProcessingMessage = "Logs are being queued for asynchronous processing"
)

func CreateRecord(w http.ResponseWriter, r *http.Request) {
	var newRecord entities.Record

	err := serialize.NewJSONDecoder(r.Body).Decode(&newRecord)
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, err, nil, response.Meta{})
		return
	}

	if config.Get().AsyncRecordCreation {
		go createRecordAsync(newRecord)
		response.Return(w, true, http.StatusOK, nil, AsyncProcessingMessage, response.Meta{})
		return
	}

	createRecord(w, r, newRecord)
}

func CreateRecordAsync(w http.ResponseWriter, r *http.Request) {
	var newRecord entities.Record

	err := serialize.NewJSONDecoder(r.Body).Decode(&newRecord)
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, err, nil, response.Meta{})
		return
	}

	go createRecordAsync(newRecord)

	response.Return(w, true, http.StatusOK, nil, AsyncProcessingMessage, response.Meta{})
}

func createRecord(w http.ResponseWriter, r *http.Request, newRecord entities.Record) {
	schemaStorage := storage.NewSchemaStorage(r.Context(), config.Get().DBAdapter)
	schema, err := usecases.NewSchemaUseCase(schemaStorage).GetSchemaByName(newRecord.SchemaName)
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, err, nil, response.Meta{})
		return
	}

	if schema.ID == 0 {
		response.Return(w, false, http.StatusBadRequest, nil, nil, response.Meta{})
		return
	}

	recordStorage := storage.NewRecordStorage(r.Context(), config.Get().DBAdapter)
	record, err := usecases.NewRecordUseCase(recordStorage).CreateRecord(newRecord, schema.ID)
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, err, nil, response.Meta{})
		return
	}

	response.Return(w, true, http.StatusOK, nil, record, response.Meta{})
}

func createRecordAsync(newRecord entities.Record) {
	ctx := context.Background()

	schemaStorage := storage.NewSchemaStorage(ctx, config.Get().DBAdapter)
	schema, err := usecases.NewSchemaUseCase(schemaStorage).GetSchemaByName(newRecord.SchemaName)
	if err != nil {
		return
	}

	if schema.ID == 0 {
		return
	}

	recordStorage := storage.NewRecordStorage(ctx, config.Get().DBAdapter)
	_, _ = usecases.NewRecordUseCase(recordStorage).CreateRecord(newRecord, schema.ID)
}

func GetAllRecords(w http.ResponseWriter, r *http.Request) {
	user := session.GetUserFromContext(r)
	// Deny access to all logs if user has any tags
	if len(user.Tags) > 0 {
		response.Return(w, false, http.StatusForbidden, nil, nil, response.Meta{})
		return
	}

	limit, offset, page := pagination.Paginate(r)

	recordStorage := storage.NewRecordStorage(r.Context(), config.Get().DBAdapter)
	recordUseCase := usecases.NewRecordUseCase(recordStorage)

	records, err := recordUseCase.GetAllRecords(limit, offset)
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, err, nil, response.Meta{})
		return
	}

	count, err := recordUseCase.GetAllRecordsCount()
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, err, nil, response.Meta{})
		return
	}

	pages := int(math.Ceil(float64(count) / float64(limit)))

	response.Return(w, true, http.StatusOK, nil, records, response.Meta{Page: page, Count: count, Pages: pages})
}

func GetRecordByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		error := fmt.Errorf("`id` path parameter is invalid")
		response.Return(w, false, http.StatusBadRequest, error, nil, response.Meta{})
		return
	}

	recordStorage := storage.NewRecordStorage(r.Context(), config.Get().DBAdapter)
	record, err := usecases.NewRecordUseCase(recordStorage).GetRecordByID(id)
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, err, nil, response.Meta{})
		return
	}

	if record.ID == 0 {
		response.Return(w, false, http.StatusNotFound, err, nil, response.Meta{})
		return
	}

	if access.IsSchemaForbiddenForUser(record.SchemaName, r) {
		response.Return(w, false, http.StatusForbidden, nil, nil, response.Meta{})
		return
	}

	response.Return(w, true, http.StatusOK, nil, record, response.Meta{})
}

func GetRecordRequestByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		error := fmt.Errorf("`id` path parameter is invalid")
		response.Return(w, false, http.StatusBadRequest, error, nil, response.Meta{})
		return
	}

	recordStorage := storage.NewRecordStorage(r.Context(), config.Get().DBAdapter)
	record, err := usecases.NewRecordUseCase(recordStorage).GetRecordByID(id)
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, err, nil, response.Meta{})
		return
	}

	if record.ID == 0 {
		response.Return(w, false, http.StatusNotFound, err, nil, response.Meta{})
		return
	}

	if access.IsSchemaForbiddenForUser(record.SchemaName, r) {
		response.Return(w, false, http.StatusForbidden, nil, nil, response.Meta{})
		return
	}

	response.ReturnPlain(w, http.StatusOK, record.Request)
}

func GetRecordResponseByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		error := fmt.Errorf("`id` path parameter is invalid")
		response.Return(w, false, http.StatusBadRequest, error, nil, response.Meta{})
		return
	}

	recordStorage := storage.NewRecordStorage(r.Context(), config.Get().DBAdapter)
	record, err := usecases.NewRecordUseCase(recordStorage).GetRecordByID(id)
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, err, nil, response.Meta{})
		return
	}

	if record.ID == 0 {
		response.Return(w, false, http.StatusNotFound, err, nil, response.Meta{})
		return
	}

	if access.IsSchemaForbiddenForUser(record.SchemaName, r) {
		response.Return(w, false, http.StatusForbidden, nil, nil, response.Meta{})
		return
	}

	response.ReturnPlain(w, http.StatusOK, record.Response)
}

func GetRecordsByQuery(w http.ResponseWriter, r *http.Request) {
	limit, offset, page := pagination.Paginate(r)
	from, to, err := timerange.Parse(r)
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, err, nil, response.Meta{})
		return
	}

	var recordsToGet entities.Record

	err = serialize.NewJSONDecoder(r.Body).Decode(&recordsToGet)
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, err, nil, response.Meta{})
		return
	}

	if access.IsSchemaForbiddenForUser(recordsToGet.SchemaName, r) {
		response.Return(w, false, http.StatusForbidden, nil, nil, response.Meta{})
		return
	}

	recordStorage := storage.NewRecordStorage(r.Context(), config.Get().DBAdapter)
	recordUseCase := usecases.NewRecordUseCase(recordStorage)

	records, count, err := recordUseCase.GetRecordsByQuery(recordsToGet, from, to, limit, offset)
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, err, nil, response.Meta{})
		return
	}

	pages := int(math.Ceil(float64(count) / float64(limit)))

	response.Return(w, true, http.StatusOK, nil, records, response.Meta{Page: page, Count: count, Pages: pages})
}

func GetRecordsByGroupHash(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	groupHash := vars["hash"]
	schemaName := vars["schemaName"]

	if access.IsSchemaForbiddenForUser(schemaName, r) {
		response.Return(w, false, http.StatusForbidden, nil, nil, response.Meta{})
		return
	}

	schemaStorage := storage.NewSchemaStorage(r.Context(), config.Get().DBAdapter)
	schema, err := usecases.NewSchemaUseCase(schemaStorage).GetSchemaByName(schemaName)
	if err != nil || schema.ID == 0 {
		response.Return(w, false, http.StatusNotFound, err, nil, response.Meta{})
		return
	}

	recordStorage := storage.NewRecordStorage(r.Context(), config.Get().DBAdapter)
	records, err := usecases.NewRecordUseCase(recordStorage).GetRecordsByGroupHash(schemaName, groupHash)
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, err, nil, response.Meta{})
		return
	}

	if len(records) == 0 {
		response.Return(w, false, http.StatusNotFound, err, nil, response.Meta{})
		return
	}

	response.Return(w, true, http.StatusOK, nil, records, response.Meta{})
}
