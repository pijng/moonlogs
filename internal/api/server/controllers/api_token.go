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

func CreateApiToken(w http.ResponseWriter, r *http.Request) {
	var newApiToken entities.ApiToken

	err := json.NewDecoder(r.Body).Decode(&newApiToken)
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, err, nil, response.Meta{})
		return
	}

	apiTokenRepository := repositories.NewApiTokenRepository(r.Context())
	apiToken, err := usecases.NewApiTokenUseCase(apiTokenRepository).CreateApiToken(newApiToken.Name)
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, err, nil, response.Meta{})
		return
	}

	response.Return(w, true, http.StatusOK, nil, apiToken, response.Meta{})
}

func DestroyApiTokenByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, fmt.Errorf("`id` path parameter is invalid: %w", err), nil, response.Meta{})
		return
	}

	apiTokenRepository := repositories.NewApiTokenRepository(r.Context())
	err = usecases.NewApiTokenUseCase(apiTokenRepository).DestroyApiTokenByID(id)
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

	apiTokenRepository := repositories.NewApiTokenRepository(r.Context())
	apiToken, err := usecases.NewApiTokenUseCase(apiTokenRepository).GetApiTokenByID(id)
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

	apiTokenRepository := repositories.NewApiTokenRepository(r.Context())
	apiToken, err := usecases.NewApiTokenUseCase(apiTokenRepository).UpdateApiTokenByID(id, apiTokenToUpdate)
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
	apiTokenRepository := repositories.NewApiTokenRepository(r.Context())
	apiTokens, err := usecases.NewApiTokenUseCase(apiTokenRepository).GetAllApiTokens()
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, err, nil, response.Meta{})
		return
	}

	response.Return(w, true, http.StatusOK, nil, apiTokens, response.Meta{})
}
