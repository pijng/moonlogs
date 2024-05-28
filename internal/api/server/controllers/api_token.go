package controllers

import (
	"fmt"
	"moonlogs/internal/api/server/response"
	"moonlogs/internal/entities"
	"moonlogs/internal/lib/serialize"
	"moonlogs/internal/usecases"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type ApiTokenController struct {
	apiTokenUseCase *usecases.ApiTokenUseCase
}

func NewApiTokenController(apiTokenUseCase *usecases.ApiTokenUseCase) *ApiTokenController {
	return &ApiTokenController{
		apiTokenUseCase: apiTokenUseCase,
	}
}

func (c *ApiTokenController) CreateApiToken(w http.ResponseWriter, r *http.Request) {
	var newApiToken entities.ApiToken

	err := serialize.NewJSONDecoder(r.Body).Decode(&newApiToken)
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, err, nil, response.Meta{})
		return
	}

	apiToken, err := c.apiTokenUseCase.CreateApiToken(r.Context(), newApiToken.Name)
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, err, nil, response.Meta{})
		return
	}

	response.Return(w, true, http.StatusOK, nil, apiToken, response.Meta{})
}

func (c *ApiTokenController) DeleteApiTokenByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, fmt.Errorf("`id` path parameter is invalid: %w", err), nil, response.Meta{})
		return
	}

	err = c.apiTokenUseCase.DeleteApiTokenByID(r.Context(), id)
	if err != nil {
		response.Return(w, false, http.StatusInternalServerError, err, nil, response.Meta{})
		return
	}

	response.Return(w, true, http.StatusOK, nil, id, response.Meta{})
}

func (c *ApiTokenController) GetApiTokenByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, fmt.Errorf("`id` path parameter is invalid: %w", err), nil, response.Meta{})
		return
	}

	apiToken, err := c.apiTokenUseCase.GetApiTokenByID(r.Context(), id)
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

func (c *ApiTokenController) UpdateApiTokenByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, fmt.Errorf("`id` path parameter is invalid: %w", err), nil, response.Meta{})
		return
	}

	var apiTokenToUpdate entities.ApiToken

	err = serialize.NewJSONDecoder(r.Body).Decode(&apiTokenToUpdate)
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, err, nil, response.Meta{})
		return
	}

	apiToken, err := c.apiTokenUseCase.UpdateApiTokenByID(r.Context(), id, apiTokenToUpdate)
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

func (c *ApiTokenController) GetAllApiTokens(w http.ResponseWriter, r *http.Request) {
	apiTokens, err := c.apiTokenUseCase.GetAllApiTokens(r.Context())
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, err, nil, response.Meta{})
		return
	}

	response.Return(w, true, http.StatusOK, nil, apiTokens, response.Meta{})
}
