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
	"moonlogs/internal/usecases"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

const (
	AsyncProcessingMessage = "Logs are being queued for asynchronous processing"
)

type RecordController struct {
	recordUseCase *usecases.RecordUseCase
	schemaUseCase *usecases.SchemaUseCase
}

func NewRecordController(recordUseCase *usecases.RecordUseCase, schemaUseCase *usecases.SchemaUseCase) *RecordController {
	return &RecordController{
		recordUseCase: recordUseCase,
		schemaUseCase: schemaUseCase,
	}
}

func (c *RecordController) CreateRecord(w http.ResponseWriter, r *http.Request) {
	var newRecord entities.Record

	err := serialize.NewJSONDecoder(r.Body).Decode(&newRecord)
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, err, nil, response.Meta{})
		return
	}

	fmt.Println(newRecord)

	if config.Get().AsyncRecordCreation {
		go c.createRecordAsync(newRecord)
		response.Return(w, true, http.StatusOK, nil, AsyncProcessingMessage, response.Meta{})
		return
	}

	c.createRecord(w, r, newRecord)
}

func (c *RecordController) CreateRecordAsync(w http.ResponseWriter, r *http.Request) {
	var newRecord entities.Record

	err := serialize.NewJSONDecoder(r.Body).Decode(&newRecord)
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, err, nil, response.Meta{})
		return
	}

	go c.createRecordAsync(newRecord)

	response.Return(w, true, http.StatusOK, nil, AsyncProcessingMessage, response.Meta{})
}

func (c *RecordController) createRecord(w http.ResponseWriter, r *http.Request, newRecord entities.Record) {
	schema, err := c.schemaUseCase.GetSchemaByName(r.Context(), newRecord.SchemaName)
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, err, nil, response.Meta{})
		return
	}

	record, err := c.recordUseCase.CreateRecord(r.Context(), newRecord, schema.ID)
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, err, nil, response.Meta{})
		return
	}

	response.Return(w, true, http.StatusOK, nil, record, response.Meta{})
}

func (c *RecordController) createRecordAsync(newRecord entities.Record) {
	ctx := context.Background()

	schema, err := c.schemaUseCase.GetSchemaByName(ctx, newRecord.SchemaName)
	if err != nil {
		return
	}

	_, _ = c.recordUseCase.CreateRecord(ctx, newRecord, schema.ID)
}

func (c *RecordController) GetAllRecords(w http.ResponseWriter, r *http.Request) {
	user := session.GetUserFromContext(r)
	// Deny access to all logs if user has any tags
	if len(user.Tags) > 0 {
		response.Return(w, false, http.StatusForbidden, nil, nil, response.Meta{})
		return
	}

	limit, offset, page := pagination.Paginate(r)

	records, err := c.recordUseCase.GetAllRecords(r.Context(), limit, offset)
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, err, nil, response.Meta{})
		return
	}

	count, err := c.recordUseCase.GetAllRecordsCount(r.Context())
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, err, nil, response.Meta{})
		return
	}

	pages := int(math.Ceil(float64(count) / float64(limit)))

	response.Return(w, true, http.StatusOK, nil, records, response.Meta{Page: page, Count: count, Pages: pages})
}

func (c *RecordController) GetRecordByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		error := fmt.Errorf("`id` path parameter is invalid")
		response.Return(w, false, http.StatusBadRequest, error, nil, response.Meta{})
		return
	}

	record, err := c.recordUseCase.GetRecordByID(r.Context(), id)
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, err, nil, response.Meta{})
		return
	}

	if access.IsSchemaForbiddenForUser(c.schemaUseCase, record.SchemaName, r) {
		response.Return(w, false, http.StatusForbidden, nil, nil, response.Meta{})
		return
	}

	response.Return(w, true, http.StatusOK, nil, record, response.Meta{})
}

func (c *RecordController) GetRecordRequestByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		error := fmt.Errorf("`id` path parameter is invalid")
		response.Return(w, false, http.StatusBadRequest, error, nil, response.Meta{})
		return
	}

	record, err := c.recordUseCase.GetRecordByID(r.Context(), id)
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, err, nil, response.Meta{})
		return
	}

	if access.IsSchemaForbiddenForUser(c.schemaUseCase, record.SchemaName, r) {
		response.Return(w, false, http.StatusForbidden, nil, nil, response.Meta{})
		return
	}

	response.ReturnPlain(w, http.StatusOK, record.Request)
}

func (c *RecordController) GetRecordResponseByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		error := fmt.Errorf("`id` path parameter is invalid")
		response.Return(w, false, http.StatusBadRequest, error, nil, response.Meta{})
		return
	}

	record, err := c.recordUseCase.GetRecordByID(r.Context(), id)
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, err, nil, response.Meta{})
		return
	}

	if access.IsSchemaForbiddenForUser(c.schemaUseCase, record.SchemaName, r) {
		response.Return(w, false, http.StatusForbidden, nil, nil, response.Meta{})
		return
	}

	response.ReturnPlain(w, http.StatusOK, record.Response)
}

func (c *RecordController) GetRecordsByQuery(w http.ResponseWriter, r *http.Request) {
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

	if access.IsSchemaForbiddenForUser(c.schemaUseCase, recordsToGet.SchemaName, r) {
		response.Return(w, false, http.StatusForbidden, nil, nil, response.Meta{})
		return
	}

	records, count, err := c.recordUseCase.GetRecordsByQuery(r.Context(), recordsToGet, from, to, limit, offset)
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, err, nil, response.Meta{})
		return
	}

	pages := int(math.Ceil(float64(count) / float64(limit)))

	response.Return(w, true, http.StatusOK, nil, records, response.Meta{Page: page, Count: count, Pages: pages})
}

func (c *RecordController) GetRecordsByGroupHash(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	groupHash := vars["hash"]
	schemaName := vars["schemaName"]

	if access.IsSchemaForbiddenForUser(c.schemaUseCase, schemaName, r) {
		response.Return(w, false, http.StatusForbidden, nil, nil, response.Meta{})
		return
	}

	_, err := c.schemaUseCase.GetSchemaByName(r.Context(), schemaName)
	if err != nil {
		response.Return(w, false, http.StatusNotFound, err, nil, response.Meta{})
		return
	}

	records, err := c.recordUseCase.GetRecordsByGroupHash(r.Context(), schemaName, groupHash)
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
