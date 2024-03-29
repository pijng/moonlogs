package controllers

import (
	"fmt"
	"moonlogs/internal/api/server/response"
	"moonlogs/internal/config"
	"moonlogs/internal/entities"
	"moonlogs/internal/lib/serialize"
	"moonlogs/internal/storage"
	"moonlogs/internal/usecases"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type UserDTO struct {
	Email     string        `json:"email"`
	Id        int           `json:"id"`
	Name      string        `json:"name"`
	Role      entities.Role `json:"role"`
	Tags      entities.Tags `json:"tag_ids"`
	IsRevoked bool          `json:"is_revoked"`
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser entities.User

	err := serialize.NewJSONDecoder(r.Body).Decode(&newUser)
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, err, nil, response.Meta{})
		return
	}

	userStorage := storage.NewUserStorage(r.Context(), config.Get().DBAdapter)
	user, err := usecases.NewUserUseCase(userStorage).CreateUser(newUser)
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, err, nil, response.Meta{})
		return
	}

	response.Return(w, true, http.StatusOK, nil, UserToDTO(user), response.Meta{})
}

func DeleteUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, fmt.Errorf("`id` path parameter is invalid: %w", err), nil, response.Meta{})
		return
	}

	userStorage := storage.NewUserStorage(r.Context(), config.Get().DBAdapter)
	err = usecases.NewUserUseCase(userStorage).DeleteUserByID(id)
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

	userStorage := storage.NewUserStorage(r.Context(), config.Get().DBAdapter)
	user, err := usecases.NewUserUseCase(userStorage).GetUserByID(id)
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

	err = serialize.NewJSONDecoder(r.Body).Decode(&userToUpdate)
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, err, nil, response.Meta{})
		return
	}

	userStorage := storage.NewUserStorage(r.Context(), config.Get().DBAdapter)
	user, err := usecases.NewUserUseCase(userStorage).UpdateUserByID(id, userToUpdate)
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
	userStorage := storage.NewUserStorage(r.Context(), config.Get().DBAdapter)
	users, err := usecases.NewUserUseCase(userStorage).GetAllUsers()
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, err, nil, response.Meta{})
		return
	}

	response.Return(w, true, http.StatusOK, nil, UsersToDTO(users), response.Meta{})
}

func CreateInitialAdmin(w http.ResponseWriter, r *http.Request) {
	userStorage := storage.NewUserStorage(r.Context(), config.Get().DBAdapter)
	userUserCase := usecases.NewUserUseCase(userStorage)
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

	err = serialize.NewJSONDecoder(r.Body).Decode(&newAdmin)
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
		Id:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Role:      user.Role,
		Tags:      user.Tags,
		IsRevoked: bool(user.IsRevoked),
	}
}
