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

type UserDTO struct {
	Email     string        `json:"email"`
	Id        int           `json:"id"`
	Name      string        `json:"name"`
	Role      entities.Role `json:"role"`
	Tags      entities.Tags `json:"tag_ids"`
	IsRevoked bool          `json:"is_revoked"`
}

type UserController struct {
	userUseCase *usecases.UserUseCase
}

func NewUserController(userUseCase *usecases.UserUseCase) *UserController {
	return &UserController{
		userUseCase: userUseCase,
	}
}

func (c *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser entities.User

	err := serialize.NewJSONDecoder(r.Body).Decode(&newUser)
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, err, nil, response.Meta{})
		return
	}

	user, err := c.userUseCase.CreateUser(r.Context(), newUser)
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, err, nil, response.Meta{})
		return
	}

	response.Return(w, true, http.StatusOK, nil, UserToDTO(user), response.Meta{})
}

func (c *UserController) DeleteUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, fmt.Errorf("`id` path parameter is invalid: %w", err), nil, response.Meta{})
		return
	}

	err = c.userUseCase.DeleteUserByID(r.Context(), id)
	if err != nil {
		response.Return(w, false, http.StatusInternalServerError, err, nil, response.Meta{})
		return
	}

	response.Return(w, true, http.StatusOK, nil, id, response.Meta{})
}

func (c *UserController) GetUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, fmt.Errorf("`id` path parameter is invalid: %w", err), nil, response.Meta{})
		return
	}

	user, err := c.userUseCase.GetUserByID(r.Context(), id)
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

func (c *UserController) UpdateUserByID(w http.ResponseWriter, r *http.Request) {
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

	user, err := c.userUseCase.UpdateUserByID(r.Context(), id, userToUpdate)
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

func (c *UserController) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := c.userUseCase.GetAllUsers(r.Context())
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, err, nil, response.Meta{})
		return
	}

	response.Return(w, true, http.StatusOK, nil, UsersToDTO(users), response.Meta{})
}

func (c *UserController) CreateInitialAdmin(w http.ResponseWriter, r *http.Request) {
	shouldCreateInitialAdmin, err := c.userUseCase.ShouldCreateInitialAdmin(r.Context())

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

	admin, err := c.userUseCase.CreateInitialAdmin(r.Context(), newAdmin)
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
