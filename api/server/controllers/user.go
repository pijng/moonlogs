package controllers

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"moonlogs/api/server/session"
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

type UserDTO struct {
	Email string      `json:"email"`
	Id    int         `json:"id"`
	Name  string      `json:"name"`
	Role  schema.Role `json:"role"`
}

var roles = []string{
	string(schema.RoleMember),
	string(schema.RoleAdmin),
}

var passwordHasher = sha256.New()

func UserCreate(w http.ResponseWriter, r *http.Request) {
	var newUser ent.User

	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		util.Return(w, false, http.StatusBadRequest, err, nil, util.Meta{})
		return
	}

	userRepository := repository.NewUserRepository(r.Context())

	userWithIdenticalEmail, err := userRepository.GetByEmail(newUser.Email)
	if userWithIdenticalEmail != nil && err == nil {
		util.Return(w, false, http.StatusBadRequest, fmt.Errorf("User with email %s already exists", newUser.Email), nil, util.Meta{})
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

	_, err = passwordHasher.Write([]byte(newUser.Password))
	if err != nil {
		util.Return(w, false, http.StatusInternalServerError, err, nil, util.Meta{})
		return
	}

	hashBytes := passwordHasher.Sum(nil)
	hashString := hex.EncodeToString(hashBytes)
	passwordHasher.Reset()

	newUser.PasswordDigest = hashString

	createdUser, err := userRepository.Create(newUser)
	if err != nil {
		util.Return(w, false, http.StatusInternalServerError, err, nil, util.Meta{})
		return
	}

	util.Return(w, true, http.StatusOK, nil, UserToDTO(createdUser), util.Meta{})
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

	util.Return(w, true, http.StatusOK, nil, UserToDTO(user), util.Meta{})
}

func UserUpdateById(w http.ResponseWriter, r *http.Request) {
	var userToUpdate ent.User

	err := json.NewDecoder(r.Body).Decode(&userToUpdate)
	if err != nil {
		util.Return(w, false, http.StatusBadRequest, err, nil, util.Meta{})
		return
	}

	if len(userToUpdate.Password) > 0 {
		_, err = passwordHasher.Write([]byte(userToUpdate.Password))
		if err != nil {
			util.Return(w, false, http.StatusInternalServerError, err, nil, util.Meta{})
			return
		}

		hashBytes := passwordHasher.Sum(nil)
		hashString := hex.EncodeToString(hashBytes)
		passwordHasher.Reset()

		userToUpdate.PasswordDigest = hashString
	}

	if len(userToUpdate.PasswordDigest) > 0 {
		token, err := session.GenerateAuthToken()
		if err != nil {
			util.Return(w, false, http.StatusInternalServerError, err, nil, util.Meta{})
			return
		}

		userToUpdate.Token = token
	}

	user, err := repository.NewUserRepository(r.Context()).UpdateById(userToUpdate)
	if err != nil {
		util.Return(w, false, http.StatusInternalServerError, err, nil, util.Meta{})
		return
	}

	util.Return(w, true, http.StatusOK, nil, UserToDTO(user), util.Meta{})
}

func UserGetAll(w http.ResponseWriter, r *http.Request) {
	users, err := repository.NewUserRepository(r.Context()).GetAll()
	if err != nil {
		util.Return(w, false, http.StatusNotFound, err, nil, util.Meta{})
		return
	}

	util.Return(w, true, http.StatusOK, nil, UsersToDTO(users), util.Meta{})
}

func UsersToDTO(users []*ent.User) []UserDTO {
	usersDTO := make([]UserDTO, 0)
	for _, user := range users {
		usersDTO = append(usersDTO, UserToDTO(user))
	}

	return usersDTO
}

func UserToDTO(user *ent.User) UserDTO {
	return UserDTO{
		Id:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Role:  schema.Role(user.Role),
	}
}
