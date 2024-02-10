package controllers

import (
	"fmt"
	"moonlogs/internal/api/server/response"
	"moonlogs/internal/config"
	"moonlogs/internal/entities"
	"moonlogs/internal/lib/serialize"
	"moonlogs/internal/mediators"
	"moonlogs/internal/storage"
	"moonlogs/internal/usecases"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CreateTag(w http.ResponseWriter, r *http.Request) {
	var newTag entities.Tag

	err := serialize.NewJSONDecoder(r.Body).Decode(&newTag)
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, err, nil, response.Meta{})
		return
	}

	tagStorage := storage.NewTagStorage(r.Context(), config.Get().DBAdapter)
	tag, err := usecases.NewTagUseCase(tagStorage).CreateTag(newTag.Name)
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, err, nil, response.Meta{})
		return
	}

	response.Return(w, true, http.StatusOK, nil, tag, response.Meta{})
}

func DeleteTagByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, fmt.Errorf("`id` path parameter is invalid: %w", err), nil, response.Meta{})
		return
	}

	tagStorage := storage.NewTagStorage(r.Context(), config.Get().DBAdapter)
	schemaStorage := storage.NewSchemaStorage(r.Context(), config.Get().DBAdapter)
	userStorage := storage.NewUserStorage(r.Context(), config.Get().DBAdapter)
	tagUsecase := usecases.NewTagUseCase(tagStorage)
	schemaUsecase := usecases.NewSchemaUseCase(schemaStorage)
	userUsecase := usecases.NewUserUseCase(userStorage)

	tagMediator := mediators.NewTagMediator(tagUsecase, schemaUsecase, userUsecase)

	err = tagMediator.DestroyTagByID(id)
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

	tagStorage := storage.NewTagStorage(r.Context(), config.Get().DBAdapter)
	tag, err := usecases.NewTagUseCase(tagStorage).GetTagByID(id)
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

	err = serialize.NewJSONDecoder(r.Body).Decode(&tagToUpdate)
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, err, nil, response.Meta{})
		return
	}

	tagStorage := storage.NewTagStorage(r.Context(), config.Get().DBAdapter)
	tag, err := usecases.NewTagUseCase(tagStorage).UpdateTagByID(id, tagToUpdate)
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
	tagStorage := storage.NewTagStorage(r.Context(), config.Get().DBAdapter)
	tags, err := usecases.NewTagUseCase(tagStorage).GetAllTags()
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, err, nil, response.Meta{})
		return
	}

	response.Return(w, true, http.StatusOK, nil, tags, response.Meta{})
}
