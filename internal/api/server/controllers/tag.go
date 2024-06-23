package controllers

import (
	"fmt"
	"moonlogs/internal/api/server/response"
	"moonlogs/internal/entities"
	"moonlogs/internal/lib/serialize"
	"moonlogs/internal/services"
	"moonlogs/internal/usecases"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type TagController struct {
	tagUseCase    *usecases.TagUseCase
	userUseCase   *usecases.UserUseCase
	schemaUseCase *usecases.SchemaUseCase
}

func NewTagController(tagUseCase *usecases.TagUseCase, userUseCase *usecases.UserUseCase, schemaUseCase *usecases.SchemaUseCase) *TagController {
	return &TagController{
		tagUseCase:    tagUseCase,
		userUseCase:   userUseCase,
		schemaUseCase: schemaUseCase,
	}
}

func (tc *TagController) CreateTag(w http.ResponseWriter, r *http.Request) {
	var newTag entities.Tag

	err := serialize.NewJSONDecoder(r.Body).Decode(&newTag)
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, err, nil, response.Meta{})
		return
	}

	tag, err := tc.tagUseCase.CreateTag(r.Context(), newTag)
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, err, nil, response.Meta{})
		return
	}

	response.Return(w, true, http.StatusOK, nil, tag, response.Meta{})
}

func (tc *TagController) DeleteTagByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, fmt.Errorf("`id` path parameter is invalid: %w", err), nil, response.Meta{})
		return
	}

	tagService := services.NewTagService(tc.tagUseCase, tc.schemaUseCase, tc.userUseCase)

	err = tagService.DestroyTagByID(r.Context(), id)
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, err, nil, response.Meta{})
		return
	}

	response.Return(w, true, http.StatusOK, nil, id, response.Meta{})
}

func (tc *TagController) GetTagByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, fmt.Errorf("`id` path parameter is invalid: %w", err), nil, response.Meta{})
		return
	}

	tag, err := tc.tagUseCase.GetTagByID(r.Context(), id)
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, err, nil, response.Meta{})
		return
	}

	response.Return(w, true, http.StatusOK, nil, tag, response.Meta{})
}

func (tc *TagController) UpdateTagByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, fmt.Errorf("`id` path parameter is invalid: %w", err), nil, response.Meta{})
		return
	}

	var tagToUpdate entities.Tag

	err = serialize.NewJSONDecoder(r.Body).Decode(&tagToUpdate)
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, err, nil, response.Meta{})
		return
	}

	tag, err := tc.tagUseCase.UpdateTagByID(r.Context(), id, tagToUpdate)
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, err, nil, response.Meta{})
		return
	}

	if tag == nil {
		response.Return(w, false, http.StatusNotFound, err, nil, response.Meta{})
		return
	}

	response.Return(w, true, http.StatusOK, nil, tag, response.Meta{})
}

func (tc *TagController) GetAllTags(w http.ResponseWriter, r *http.Request) {
	tags, err := tc.tagUseCase.GetAllTags(r.Context())
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, err, nil, response.Meta{})
		return
	}

	response.Return(w, true, http.StatusOK, nil, tags, response.Meta{})
}
