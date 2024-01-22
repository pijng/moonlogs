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

type UserDTO struct {
	Email string        `json:"email"`
	Id    int           `json:"id"`
	Name  string        `json:"name"`
	Role  entities.Role `json:"role"`
	Tags  entities.Tags `json:"tag_ids"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser entities.User

	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, err, nil, response.Meta{})
		return
	}

	userRepository := repositories.NewUserRepository(r.Context())
	user, err := usecases.NewUserUseCase(userRepository).CreateUser(newUser)
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, err, nil, response.Meta{})
		return
	}

	response.Return(w, true, http.StatusOK, nil, UserToDTO(user), response.Meta{})
}

func DestroyUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, fmt.Errorf("`id` path parameter is invalid: %w", err), nil, response.Meta{})
		return
	}

	userRepository := repositories.NewUserRepository(r.Context())
	err = usecases.NewUserUseCase(userRepository).DestroyUserByID(id)
	if err != nil {
		response.Return(w, false, http.StatusInternalServerError, err, nil, response.Meta{})
		return
	}

	response.Return(w, true, http.StatusOK, nil, id, response.Meta{})
}

func GetUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, fmt.Errorf("`id` path parameter is invalid: %w", err), nil, response.Meta{})
		return
	}

	userRepository := repositories.NewUserRepository(r.Context())
	user, err := usecases.NewUserUseCase(userRepository).GetUserByID(id)
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, err, nil, response.Meta{})
		return
	}

	if user.ID == 0 {
		response.Return(w, false, http.StatusNotFound, err, nil, response.Meta{})
		return
	}

	response.Return(w, true, http.StatusOK, nil, UserToDTO(user), response.Meta{})
}

func UpdateUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, fmt.Errorf("`id` path parameter is invalid: %w", err), nil, response.Meta{})
		return
	}

	var userToUpdate entities.User

	err = json.NewDecoder(r.Body).Decode(&userToUpdate)
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, err, nil, response.Meta{})
		return
	}

	userRepository := repositories.NewUserRepository(r.Context())
	user, err := usecases.NewUserUseCase(userRepository).UpdateUserByID(id, userToUpdate)
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, err, nil, response.Meta{})
		return
	}

	if user == nil {
		response.Return(w, false, http.StatusNotFound, err, nil, response.Meta{})
		return
	}

	response.Return(w, true, http.StatusOK, nil, UserToDTO(user), response.Meta{})
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	userRepository := repositories.NewUserRepository(r.Context())
	users, err := usecases.NewUserUseCase(userRepository).GetAllUsers()
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, err, nil, response.Meta{})
		return
	}

	response.Return(w, true, http.StatusOK, nil, UsersToDTO(users), response.Meta{})
}

func CreateInitialAdmin(w http.ResponseWriter, r *http.Request) {
	userUserCase := usecases.NewUserUseCase(repositories.NewUserRepository(r.Context()))
	shouldCreateInitialAdmin, err := userUserCase.ShouldCreateInitialAdmin()

	if err != nil {
		response.Return(w, false, http.StatusBadRequest, err, nil, response.Meta{})
		return
	}

	if !shouldCreateInitialAdmin {
		response.Return(w, false, http.StatusBadRequest, err, nil, response.Meta{})
		return
	}

	var newAdmin entities.User

	err = json.NewDecoder(r.Body).Decode(&newAdmin)
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, err, nil, response.Meta{})
		return
	}

	admin, err := userUserCase.CreateInitialAdmin(newAdmin)
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, fmt.Errorf("failed creating initial admin: %w", err), nil, response.Meta{})
		return
	}

	response.Return(w, true, http.StatusOK, nil, Session{Token: admin.Token}, response.Meta{})
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
