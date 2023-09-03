package controllers

import (
	"encoding/json"
	"fmt"
	"moonlogs/api/server/util"
	"moonlogs/ent"
	"moonlogs/ent/schema"
	"moonlogs/internal/repository"
	"net/http"
	"slices"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

var roles = []string{
	string(schema.RoleMember),
	string(schema.RoleAdmin),
}

func UserCreate(w http.ResponseWriter, r *http.Request) {
	var newUser ent.User

	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		util.Return(w, false, http.StatusBadRequest, err, nil, util.Meta{})
		return
	}

	if len(newUser.Role) > 0 {
		isValidRole := slices.Contains(roles, newUser.Role)
		if !isValidRole {
			appropriateRoles := strings.Join(roles, ", ")
			error := fmt.Errorf("`role` field should be one of: %v", appropriateRoles)

			util.Return(w, false, http.StatusBadRequest, error, nil, util.Meta{})
			return
		}
	}

	createdUser, err := repository.NewUserRepository(r.Context()).Create(newUser)
	if err != nil {
		util.Return(w, false, http.StatusInternalServerError, err, nil, util.Meta{})
		return
	}

	util.Return(w, true, http.StatusOK, nil, createdUser, util.Meta{})
}

func UserDestroyById(w http.ResponseWriter, r *http.Request) {
	var userToDestroy ent.User

	err := json.NewDecoder(r.Body).Decode(&userToDestroy)
	if err != nil {
		util.Return(w, false, http.StatusBadRequest, err, nil, util.Meta{})
		return
	}

	err = repository.NewUserRepository(r.Context()).DestroyById(userToDestroy.ID)
	if err != nil {
		util.Return(w, false, http.StatusInternalServerError, err, nil, util.Meta{})
		return
	}

	util.Return(w, true, http.StatusOK, nil, userToDestroy.ID, util.Meta{})
}

func UserGetById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		error := fmt.Errorf("`id` path parameter is invalid")
		util.Return(w, false, http.StatusBadRequest, error, nil, util.Meta{})
		return
	}

	user, err := repository.NewUserRepository(r.Context()).GetById(id)
	if err != nil {
		util.Return(w, false, http.StatusNotFound, err, nil, util.Meta{})
		return
	}

	util.Return(w, true, http.StatusOK, nil, user, util.Meta{})
}

func UserUpdateById(w http.ResponseWriter, r *http.Request) {
	var userToUpdate ent.User

	err := json.NewDecoder(r.Body).Decode(&userToUpdate)
	if err != nil {
		util.Return(w, false, http.StatusBadRequest, err, nil, util.Meta{})
		return
	}

	u, err := repository.NewUserRepository(r.Context()).UpdateById(userToUpdate)
	if err != nil {
		util.Return(w, false, http.StatusInternalServerError, err, nil, util.Meta{})
		return
	}

	util.Return(w, true, http.StatusOK, nil, u, util.Meta{})
}

func UserGetAll(w http.ResponseWriter, r *http.Request) {
	users, err := repository.NewUserRepository(r.Context()).GetAll()
	if err != nil {
		util.Return(w, false, http.StatusNotFound, err, nil, util.Meta{})
		return
	}

	util.Return(w, true, http.StatusOK, nil, users, util.Meta{})
}
