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

type AlertingRuleController struct {
	alertingRuleUseCase *usecases.AlertingRuleUseCase
}

func NewAlertingRuleController(alertingRuleUseCase *usecases.AlertingRuleUseCase) *AlertingRuleController {
	return &AlertingRuleController{
		alertingRuleUseCase: alertingRuleUseCase,
	}
}

func (arc *AlertingRuleController) CreateRule(w http.ResponseWriter, r *http.Request) {
	var newRule entities.AlertingRule

	err := serialize.NewJSONDecoder(r.Body).Decode(&newRule)
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, err, nil, response.Meta{})
		return
	}

	tag, err := arc.alertingRuleUseCase.CreateRule(r.Context(), newRule)
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, err, nil, response.Meta{})
		return
	}

	response.Return(w, true, http.StatusOK, nil, tag, response.Meta{})
}

func (arc *AlertingRuleController) DeleteRuleByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, fmt.Errorf("`id` path parameter is invalid: %w", err), nil, response.Meta{})
		return
	}

	err = arc.alertingRuleUseCase.DeleteRuleByID(r.Context(), id)
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, err, nil, response.Meta{})
		return
	}

	response.Return(w, true, http.StatusOK, nil, id, response.Meta{})
}

func (arc *AlertingRuleController) GetRuleByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, fmt.Errorf("`id` path parameter is invalid: %w", err), nil, response.Meta{})
		return
	}

	rule, err := arc.alertingRuleUseCase.GetRuleByID(r.Context(), id)
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, err, nil, response.Meta{})
		return
	}

	response.Return(w, true, http.StatusOK, nil, rule, response.Meta{})
}

func (arc *AlertingRuleController) UpdateRuleByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, fmt.Errorf("`id` path parameter is invalid: %w", err), nil, response.Meta{})
		return
	}

	var ruleToUpdate entities.AlertingRule

	err = serialize.NewJSONDecoder(r.Body).Decode(&ruleToUpdate)
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, err, nil, response.Meta{})
		return
	}

	rule, err := arc.alertingRuleUseCase.UpdateRuleByID(r.Context(), id, ruleToUpdate)
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, err, nil, response.Meta{})
		return
	}

	if rule == nil {
		response.Return(w, false, http.StatusNotFound, err, nil, response.Meta{})
		return
	}

	response.Return(w, true, http.StatusOK, nil, rule, response.Meta{})
}

func (arc *AlertingRuleController) GetAllRules(w http.ResponseWriter, r *http.Request) {
	rules, err := arc.alertingRuleUseCase.GetAllRules(r.Context())
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, err, nil, response.Meta{})
		return
	}

	response.Return(w, true, http.StatusOK, nil, rules, response.Meta{})
}
