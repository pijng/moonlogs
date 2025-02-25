package controllers

import (
	"moonlogs/internal/api/server/response"
	"moonlogs/internal/lib/serialize"
	"moonlogs/internal/usecases"
	"net/http"
)

type InsightsController struct {
	insightsUseCase *usecases.InsightsUseCase
}

func NewInsightsController(insightsUseCase *usecases.InsightsUseCase) *InsightsController {
	return &InsightsController{
		insightsUseCase: insightsUseCase,
	}
}

type PromptPayload struct {
	Prompt string `json:"prompt"`
}

func (c *InsightsController) GenerateKontent(w http.ResponseWriter, r *http.Request) {
	var payload PromptPayload

	err := serialize.NewJSONDecoder(r.Body).Decode(&payload)
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, err, nil, response.Meta{})
		return
	}

	content, err := c.insightsUseCase.GenerateContent(r.Context(), payload.Prompt)
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, err, nil, response.Meta{})
		return
	}

	response.Return(w, true, http.StatusOK, nil, content, response.Meta{})
}
