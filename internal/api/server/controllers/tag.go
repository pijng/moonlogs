package controllers

import (
	"encoding/json"
	"fmt"
	"moonlogs/internal/api/server/response"
	"moonlogs/internal/entities"
	"moonlogs/internal/repositories"
	"moonlogs/internal/usecases"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CreateTag(w http.ResponseWriter, r *http.Request) {
	var newTag entities.Tag

	err := json.NewDecoder(r.Body).Decode(&newTag)
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, err, nil, response.Meta{})
		return
	}

	tagRepository := repositories.NewTagRepository(r.Context())
	tag, err := usecases.NewTagUseCase(tagRepository).CreateTag(newTag.Name)
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, err, nil, response.Meta{})
		return
	}

	response.Return(w, true, http.StatusOK, nil, tag, response.Meta{})
}

func DestroyTagByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, fmt.Errorf("`id` path parameter is invalid: %w", err), nil, response.Meta{})
		return
	}

	tagRepository := repositories.NewTagRepository(r.Context())
	err = usecases.NewTagUseCase(tagRepository).DestroyTagByID(id)
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, err, nil, response.Meta{})
		return
	}

	response.Return(w, true, http.StatusOK, nil, id, response.Meta{})
}

func GetTagByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, fmt.Errorf("`id` path parameter is invalid: %w", err), nil, response.Meta{})
		return
	}

	tagRepository := repositories.NewTagRepository(r.Context())
	tag, err := usecases.NewTagUseCase(tagRepository).GetTagByID(id)
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, err, nil, response.Meta{})
		return
	}

	if tag.ID == 0 {
		response.Return(w, false, http.StatusNotFound, err, nil, response.Meta{})
		return
	}

	response.Return(w, true, http.StatusOK, nil, tag, response.Meta{})
}

func UpdateTagByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, fmt.Errorf("`id` path parameter is invalid: %w", err), nil, response.Meta{})
		return
	}

	var tagToUpdate entities.Tag

	err = json.NewDecoder(r.Body).Decode(&tagToUpdate)
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, err, nil, response.Meta{})
		return
	}

	tagRepository := repositories.NewTagRepository(r.Context())
	tag, err := usecases.NewTagUseCase(tagRepository).UpdateTagByID(id, tagToUpdate)
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

func GetAllTags(w http.ResponseWriter, r *http.Request) {
	tagRepository := repositories.NewTagRepository(r.Context())
	tags, err := usecases.NewTagUseCase(tagRepository).GetAllTags()
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, err, nil, response.Meta{})
		return
	}

	response.Return(w, true, http.StatusOK, nil, tags, response.Meta{})
}
