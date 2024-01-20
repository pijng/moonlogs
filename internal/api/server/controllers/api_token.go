package controllers

import (
	"encoding/json"
	"fmt"
	"moonlogs/internal/api/server/util"
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
		util.Return(w, false, http.StatusBadRequest, err, nil, util.Meta{})
		return
	}

	apiTokenRepository := repositories.NewApiTokenRepository(r.Context())
	apiToken, err := usecases.NewApiTokenUseCase(apiTokenRepository).CreateApiToken(newApiToken.Name)
	if err != nil {
		util.Return(w, false, http.StatusBadRequest, err, nil, util.Meta{})
		return
	}

	util.Return(w, true, http.StatusOK, nil, apiToken, util.Meta{})
}

func DestroyApiTokenByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		util.Return(w, false, http.StatusBadRequest, fmt.Errorf("`id` path parameter is invalid: %w", err), nil, util.Meta{})
		return
	}

	apiTokenRepository := repositories.NewApiTokenRepository(r.Context())
	err = usecases.NewApiTokenUseCase(apiTokenRepository).DestroyApiTokenByID(id)
	if err != nil {
		util.Return(w, false, http.StatusInternalServerError, err, nil, util.Meta{})
		return
	}

	util.Return(w, true, http.StatusOK, nil, id, util.Meta{})
}

func GetApiTokenByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		util.Return(w, false, http.StatusBadRequest, fmt.Errorf("`id` path parameter is invalid: %w", err), nil, util.Meta{})
		return
	}

	apiTokenRepository := repositories.NewApiTokenRepository(r.Context())
	apiToken, err := usecases.NewApiTokenUseCase(apiTokenRepository).GetApiTokenByID(id)
	if err != nil {
		util.Return(w, false, http.StatusBadRequest, err, nil, util.Meta{})
		return
	}

	if apiToken.ID == 0 {
		util.Return(w, false, http.StatusNotFound, err, nil, util.Meta{})
		return
	}

	util.Return(w, true, http.StatusOK, nil, apiToken, util.Meta{})
}

func UpdateApiTokenByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		util.Return(w, false, http.StatusBadRequest, fmt.Errorf("`id` path parameter is invalid: %w", err), nil, util.Meta{})
		return
	}

	var apiTokenToUpdate entities.ApiToken

	err = json.NewDecoder(r.Body).Decode(&apiTokenToUpdate)
	if err != nil {
		util.Return(w, false, http.StatusBadRequest, err, nil, util.Meta{})
		return
	}

	apiTokenRepository := repositories.NewApiTokenRepository(r.Context())
	apiToken, err := usecases.NewApiTokenUseCase(apiTokenRepository).UpdateApiTokenByID(id, apiTokenToUpdate)
	if err != nil {
		util.Return(w, false, http.StatusBadRequest, err, nil, util.Meta{})
		return
	}

	if apiToken == nil {
		util.Return(w, false, http.StatusNotFound, err, nil, util.Meta{})
		return
	}

	util.Return(w, true, http.StatusOK, nil, apiToken, util.Meta{})
}

func GetAllApiTokens(w http.ResponseWriter, r *http.Request) {
	apiTokenRepository := repositories.NewApiTokenRepository(r.Context())
	apiTokens, err := usecases.NewApiTokenUseCase(apiTokenRepository).GetAllApiTokens()
	if err != nil {
		util.Return(w, false, http.StatusBadRequest, err, nil, util.Meta{})
		return
	}

	util.Return(w, true, http.StatusOK, nil, apiTokens, util.Meta{})
}
