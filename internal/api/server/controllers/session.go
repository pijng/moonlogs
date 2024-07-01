package controllers

import (
	"errors"
	"fmt"
	"moonlogs/internal/api/server/response"
	"moonlogs/internal/api/server/session"
	"moonlogs/internal/entities"
	"moonlogs/internal/lib/serialize"
	"moonlogs/internal/shared"
	"moonlogs/internal/storage"
	"moonlogs/internal/usecases"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

type Credentials struct {
	Email    string
	Password string
}

type Session struct {
	Token                    string        `json:"token"`
	ID                       int           `json:"id"`
	Name                     string        `json:"name"`
	Email                    string        `json:"email"`
	Role                     entities.Role `json:"role"`
	ShouldCreateInitialAdmin bool          `json:"should_create_initial_admin"`
	IsRevoked                bool          `json:"is_revoked"`
	Tags                     entities.Tags `json:"tag_ids"`
}

type SessionController struct {
	userUseCase *usecases.UserUseCase
}

func NewSessionController(userUseCase *usecases.UserUseCase) *SessionController {
	return &SessionController{
		userUseCase: userUseCase,
	}
}

func (c *SessionController) Login(w http.ResponseWriter, r *http.Request) {
	var token string
	var credentials Credentials

	err := serialize.NewJSONDecoder(r.Body).Decode(&credentials)
	if err != nil {
		response.Return(w, false, http.StatusBadRequest, err, nil, response.Meta{})
		return
	}

	user, err := c.userUseCase.GetUserByEmail(r.Context(), credentials.Email)
	if err != nil {
		response.Return(w, false, http.StatusUnauthorized, err, nil, response.Meta{})
		return
	}

	err = checkPassword(user.PasswordDigest, credentials.Password)
	if err != nil {
		response.Return(w, false, http.StatusUnauthorized, err, nil, response.Meta{})
		return
	}

	token = user.Token

	if token == "" {
		token, err = shared.GenerateRandomToken(16)
		if err != nil {
			response.Return(w, false, http.StatusInternalServerError, err, nil, response.Meta{})
			return
		}
	}

	store := session.GetSessionStore()
	session, err := store.Get(r, session.NAME)
	if err != nil {
		response.Return(w, false, http.StatusInternalServerError, err, nil, response.Meta{})
		return
	}

	session.Values["token"] = token
	session.Values["userID"] = user.ID
	session.Values["role"] = string(user.Role)

	err = session.Save(r, w)
	if err != nil {
		response.Return(w, false, http.StatusInternalServerError, err, nil, response.Meta{})
		return
	}

	if token != user.Token {
		err = c.userUseCase.UpdateUserTokenByID(r.Context(), user.ID, token)
		if err != nil {
			response.Return(w, false, http.StatusInternalServerError, err, nil, response.Meta{})
			return
		}
	}

	sessionPayload := Session{
		Token:     token,
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Role:      user.Role,
		IsRevoked: bool(user.IsRevoked),
		Tags:      user.Tags,
	}

	response.Return(w, true, http.StatusOK, nil, sessionPayload, response.Meta{})
}

func (c *SessionController) GetSession(w http.ResponseWriter, r *http.Request) {

	shouldCreateInitialAdmin, err := c.userUseCase.ShouldCreateInitialAdmin(r.Context())
	if err != nil {
		response.Return(w, false, http.StatusInternalServerError, fmt.Errorf("failed checking if initial admin is required: %w", err), nil, response.Meta{})
		return
	}

	if shouldCreateInitialAdmin {
		response.Return(w, true, http.StatusOK, nil, Session{ShouldCreateInitialAdmin: true}, response.Meta{})
		return
	}

	store := session.GetSessionStore()
	session, err := store.Get(r, session.NAME)
	if err != nil {
		response.Return(w, false, http.StatusInternalServerError, err, nil, response.Meta{})
		return
	}

	bearerToken := shared.ExtractBearerToken(r)

	if bearerToken == "" {
		bearerToken, _ = session.Values["token"].(string)
	}

	user, err := c.userUseCase.GetUserByToken(r.Context(), bearerToken)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			response.Return(w, false, http.StatusUnauthorized, err, nil, response.Meta{})
			return
		}

		response.Return(w, false, http.StatusInternalServerError, err, nil, response.Meta{})
		return
	}

	if user == nil {
		response.Return(w, false, http.StatusUnauthorized, nil, nil, response.Meta{})
		return
	}

	if bearerToken == "" {
		bearerToken = user.Token
		session.Values["token"] = bearerToken
		session.Values["userID"] = user.ID
		session.Values["role"] = string(user.Role)
	}

	err = session.Save(r, w)
	if err != nil {
		response.Return(w, false, http.StatusInternalServerError, err, nil, response.Meta{})
		return
	}

	sessionPayload := Session{
		Token:     bearerToken,
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Role:      user.Role,
		IsRevoked: bool(user.IsRevoked),
		Tags:      user.Tags,
	}

	response.Return(w, true, http.StatusOK, nil, sessionPayload, response.Meta{})
}

func checkPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
