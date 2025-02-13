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

type NotificationProfileController struct {
	notificationProfileUseCase *usecases.NotificationProfileUseCase
}

func NewNotificationProfileController(notificationProfileUseCase *usecases.NotificationProfileUseCase) *NotificationProfileController {
	return &NotificationProfileController{
		notificationProfileUseCase: notificationProfileUseCase,
	}
}

func (c *NotificationProfileController) CreateNotificationProfile(w http.ResponseWriter, r *http.Request) {
	var newNotificationProfile entities.NotificationProfile

	err := serialize.NewJSONDecoder(r.Body).Decode(&newNotificationProfile)
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, err, nil, response.Meta{})
		return
	}

	profile, err := c.notificationProfileUseCase.CreateNotificationProfile(r.Context(), newNotificationProfile)
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, err, nil, response.Meta{})
		return
	}

	response.Return(w, true, http.StatusOK, nil, profile, response.Meta{})
}

func (c *NotificationProfileController) DeleteNotificationProfileByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, fmt.Errorf("`id` path parameter is invalid: %w", err), nil, response.Meta{})
		return
	}

	err = c.notificationProfileUseCase.DeleteNotificationProfileByID(r.Context(), id)
	if err != nil {
		response.Return(w, false, http.StatusInternalServerError, err, nil, response.Meta{})
		return
	}

	response.Return(w, true, http.StatusOK, nil, id, response.Meta{})
}

func (c *NotificationProfileController) GetNotificationProfileByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, fmt.Errorf("`id` path parameter is invalid: %w", err), nil, response.Meta{})
		return
	}

	profile, err := c.notificationProfileUseCase.GetNotificationProfileByID(r.Context(), id)
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, err, nil, response.Meta{})
		return
	}

	response.Return(w, true, http.StatusOK, nil, profile, response.Meta{})
}

func (c *NotificationProfileController) UpdateNotificationProfileByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, fmt.Errorf("`id` path parameter is invalid: %w", err), nil, response.Meta{})
		return
	}

	var profileToUpdate entities.NotificationProfile

	err = serialize.NewJSONDecoder(r.Body).Decode(&profileToUpdate)
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, err, nil, response.Meta{})
		return
	}

	profile, err := c.notificationProfileUseCase.UpdateNotificationProfileByID(r.Context(), id, profileToUpdate)
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, err, nil, response.Meta{})
		return
	}

	if profile == nil {
		response.Return(w, false, http.StatusNotFound, err, nil, response.Meta{})
		return
	}

	response.Return(w, true, http.StatusOK, nil, profile, response.Meta{})
}

func (c *NotificationProfileController) GetAllNotificationProfiles(w http.ResponseWriter, r *http.Request) {
	profiles, err := c.notificationProfileUseCase.GetAllNotificationProfiles(r.Context())
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, err, nil, response.Meta{})
		return
	}

	response.Return(w, true, http.StatusOK, nil, profiles, response.Meta{})
}
