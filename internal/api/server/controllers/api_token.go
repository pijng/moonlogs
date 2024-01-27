package controllers

import (
	"encoding/json"
	"fmt"
	"moonlogs/internal/api/server/response"
	"moonlogs/internal/config"
	"moonlogs/internal/entities"
	"moonlogs/internal/storage"
	"moonlogs/internal/usecases"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CreateApiToken(w http.ResponseWriter, r *http.Request) {
	var newApiToken entities.ApiToken

	err := json.NewDecoder(r.Body).Decode(&newApiToken)
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, err, nil, response.Meta{})
		return
	}

	apiTokenStorage := storage.NewApiTokenStorage(r.Context(), config.Get().DBAdapter)
	apiToken, err := usecases.NewApiTokenUseCase(apiTokenStorage).CreateApiToken(newApiToken.Name)
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, err, nil, response.Meta{})
		return
	}

	response.Return(w, true, http.StatusOK, nil, apiToken, response.Meta{})
}

func DeleteApiTokenByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, fmt.Errorf("`id` path parameter is invalid: %w", err), nil, response.Meta{})
		return
	}

	apiTokenStorage := storage.NewApiTokenStorage(r.Context(), config.Get().DBAdapter)
	err = usecases.NewApiTokenUseCase(apiTokenStorage).DeleteApiTokenByID(id)
	if err != nil {
		response.Return(w, false, http.StatusInternalServerError, err, nil, response.Meta{})
		return
	}

	response.Return(w, true, http.StatusOK, nil, id, response.Meta{})
}

func GetApiTokenByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, fmt.Errorf("`id` path parameter is invalid: %w", err), nil, response.Meta{})
		return
	}

	apiTokenStorage := storage.NewApiTokenStorage(r.Context(), config.Get().DBAdapter)
	apiToken, err := usecases.NewApiTokenUseCase(apiTokenStorage).GetApiTokenByID(id)
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, err, nil, response.Meta{})
		return
	}

	if apiToken.ID == 0 {
		response.Return(w, false, http.StatusNotFound, err, nil, response.Meta{})
		return
	}

	response.Return(w, true, http.StatusOK, nil, apiToken, response.Meta{})
}

func UpdateApiTokenByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, fmt.Errorf("`id` path parameter is invalid: %w", err), nil, response.Meta{})
		return
	}

	var apiTokenToUpdate entities.ApiToken

	err = json.NewDecoder(r.Body).Decode(&apiTokenToUpdate)
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, err, nil, response.Meta{})
		return
	}

	apiTokenStorage := storage.NewApiTokenStorage(r.Context(), config.Get().DBAdapter)
	apiToken, err := usecases.NewApiTokenUseCase(apiTokenStorage).UpdateApiTokenByID(id, apiTokenToUpdate)
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, err, nil, response.Meta{})
		return
	}

	if apiToken == nil {
		response.Return(w, false, http.StatusNotFound, err, nil, response.Meta{})
		return
	}

	response.Return(w, true, http.StatusOK, nil, apiToken, response.Meta{})
}

func GetAllApiTokens(w http.ResponseWriter, r *http.Request) {
	apiTokenStorage := storage.NewApiTokenStorage(r.Context(), config.Get().DBAdapter)
	apiTokens, err := usecases.NewApiTokenUseCase(apiTokenStorage).GetAllApiTokens()
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, err, nil, response.Meta{})
		return
	}

	response.Return(w, true, http.StatusOK, nil, apiTokens, response.Meta{})
}
