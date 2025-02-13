package controllers

import (
	"moonlogs/internal/api/server/response"
	"moonlogs/internal/entities"
	"moonlogs/internal/lib/serialize"
	"moonlogs/internal/usecases"
	"net/http"
)

type IncidentController struct {
	incidentUseCase *usecases.IncidentUseCase
}

func NewIncidentController(incidentUseCase *usecases.IncidentUseCase) *IncidentController {
	return &IncidentController{
		incidentUseCase: incidentUseCase,
	}
}

func (ic *IncidentController) GetIncidentsByKeys(w http.ResponseWriter, r *http.Request) {
	var incidentToGet entities.Incident

	err := serialize.NewJSONDecoder(r.Body).Decode(&incidentToGet)
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, err, nil, response.Meta{})
		return
	}

	incidents, err := ic.incidentUseCase.GetIncidentsByKeys(r.Context(), incidentToGet.Keys)
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, err, nil, response.Meta{})
		return
	}

	response.Return(w, true, http.StatusOK, nil, incidents, response.Meta{})
}

func (ic *IncidentController) GetAllIncidents(w http.ResponseWriter, r *http.Request) {
	incidents, err := ic.incidentUseCase.GetAllIncidents(r.Context())
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, err, nil, response.Meta{})
		return
	}

	response.Return(w, true, http.StatusOK, nil, incidents, response.Meta{})
}
