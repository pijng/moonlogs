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

type UserDTO struct {
	Email string        `json:"email"`
	Id    int           `json:"id"`
	Name  string        `json:"name"`
	Role  entities.Role `json:"role"`
	Tags  entities.Tags `json:"tags"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser entities.User

	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		util.Return(w, false, http.StatusBadRequest, err, nil, util.Meta{})
		return
	}

	userRepository := repositories.NewUserRepository(r.Context())
	user, err := usecases.NewUserUseCase(userRepository).CreateUser(newUser)
	if err != nil {
		util.Return(w, false, http.StatusBadRequest, err, nil, util.Meta{})
		return
	}

	util.Return(w, true, http.StatusOK, nil, UserToDTO(user), util.Meta{})
}

func DestroyUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		util.Return(w, false, http.StatusBadRequest, fmt.Errorf("`id` path parameter is invalid: %w", err), nil, util.Meta{})
		return
	}

	userRepository := repositories.NewUserRepository(r.Context())
	err = usecases.NewUserUseCase(userRepository).DestroyUserByID(id)
	if err != nil {
		util.Return(w, false, http.StatusInternalServerError, err, nil, util.Meta{})
		return
	}

	util.Return(w, true, http.StatusOK, nil, id, util.Meta{})
}

func GetUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		util.Return(w, false, http.StatusBadRequest, fmt.Errorf("`id` path parameter is invalid: %w", err), nil, util.Meta{})
		return
	}

	userRepository := repositories.NewUserRepository(r.Context())
	user, err := usecases.NewUserUseCase(userRepository).GetUserByID(id)
	if err != nil {
		util.Return(w, false, http.StatusBadRequest, err, nil, util.Meta{})
		return
	}

	if user.ID == 0 {
		util.Return(w, false, http.StatusNotFound, err, nil, util.Meta{})
		return
	}

	util.Return(w, true, http.StatusOK, nil, UserToDTO(user), util.Meta{})
}

func UpdateUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		util.Return(w, false, http.StatusBadRequest, fmt.Errorf("`id` path parameter is invalid: %w", err), nil, util.Meta{})
		return
	}

	var userToUpdate entities.User

	err = json.NewDecoder(r.Body).Decode(&userToUpdate)
	if err != nil {
		util.Return(w, false, http.StatusBadRequest, err, nil, util.Meta{})
		return
	}

	userRepository := repositories.NewUserRepository(r.Context())
	user, err := usecases.NewUserUseCase(userRepository).UpdateUserByID(id, userToUpdate)
	if err != nil {
		util.Return(w, false, http.StatusBadRequest, err, nil, util.Meta{})
		return
	}

	if user == nil {
		util.Return(w, false, http.StatusNotFound, err, nil, util.Meta{})
		return
	}

	util.Return(w, true, http.StatusOK, nil, UserToDTO(user), util.Meta{})
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	userRepository := repositories.NewUserRepository(r.Context())
	users, err := usecases.NewUserUseCase(userRepository).GetAllUsers()
	if err != nil {
		util.Return(w, false, http.StatusBadRequest, err, nil, util.Meta{})
		return
	}

	util.Return(w, true, http.StatusOK, nil, UsersToDTO(users), util.Meta{})
}

func CreateInitialAdmin(w http.ResponseWriter, r *http.Request) {
	userUserCase := usecases.NewUserUseCase(repositories.NewUserRepository(r.Context()))
	shouldCreateInitialAdmin, err := userUserCase.ShouldCreateInitialAdmin()

	if err != nil {
		util.Return(w, false, http.StatusBadRequest, err, nil, util.Meta{})
		return
	}

	if !shouldCreateInitialAdmin {
		util.Return(w, false, http.StatusBadRequest, err, nil, util.Meta{})
		return
	}

	var newAdmin entities.User

	err = json.NewDecoder(r.Body).Decode(&newAdmin)
	if err != nil {
		util.Return(w, false, http.StatusBadRequest, err, nil, util.Meta{})
		return
	}

	admin, err := userUserCase.CreateInitialAdmin(newAdmin)
	if err != nil {
		util.Return(w, false, http.StatusBadRequest, fmt.Errorf("failed creating initial admin: %w", err), nil, util.Meta{})
		return
	}

	util.Return(w, true, http.StatusOK, nil, Session{Token: admin.Token}, util.Meta{})
}

func UsersToDTO(users []*entities.User) []UserDTO {
	usersDTO := make([]UserDTO, 0)
	for _, user := range users {
		usersDTO = append(usersDTO, UserToDTO(user))
	}

	return usersDTO
}

func UserToDTO(user *entities.User) UserDTO {
	return UserDTO{
		Id:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Role:  user.Role,
		Tags:  user.Tags,
	}
}
