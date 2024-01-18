package controllers

import (
	"encoding/json"
	"fmt"
	"moonlogs/api/server/session"
	"moonlogs/api/server/util"
	"moonlogs/internal/entities"
	"moonlogs/internal/repositories"
	"moonlogs/internal/usecases"
	"net/http"
	"strings"

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
}

func Login(w http.ResponseWriter, r *http.Request) {
	var token string
	var credentials Credentials

	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		util.Return(w, false, http.StatusBadRequest, err, nil, util.Meta{})
		return
	}

	userRepository := repositories.NewUserRepository(r.Context())
	userUseCase := usecases.NewUserUseCase(userRepository)

	user, err := userUseCase.GetUserByEmail(credentials.Email)
	if err != nil || user.ID == 0 {
		util.Return(w, false, http.StatusUnauthorized, err, nil, util.Meta{})
		return
	}

	err = checkPassword(user.PasswordDigest, credentials.Password)
	if err != nil {
		util.Return(w, false, http.StatusUnauthorized, err, nil, util.Meta{})
		return
	}

	token = user.Token

	if token == "" {
		token, err = session.GenerateAuthToken()
		if err != nil {
			util.Return(w, false, http.StatusInternalServerError, err, nil, util.Meta{})
			return
		}
	}

	store := session.GetSessionStore()
	session, err := store.Get(r, session.NAME)
	if err != nil {
		util.Return(w, false, http.StatusInternalServerError, err, nil, util.Meta{})
		return
	}

	session.Values["token"] = token
	session.Values["userID"] = user.ID

	err = session.Save(r, w)
	if err != nil {
		util.Return(w, false, http.StatusInternalServerError, err, nil, util.Meta{})
		return
	}

	if token != user.Token {
		err = userUseCase.UpdateUserTokenByID(user.ID, token)
		if err != nil {
			util.Return(w, false, http.StatusInternalServerError, err, nil, util.Meta{})
			return
		}
	}

	sessionPayload := Session{
		Token: token,
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Role:  user.Role,
	}

	util.Return(w, true, http.StatusOK, nil, sessionPayload, util.Meta{})
}

func GetSession(w http.ResponseWriter, r *http.Request) {
	userUserCase := usecases.NewUserUseCase(repositories.NewUserRepository(r.Context()))

	shouldCreateInitialAdmin, err := userUserCase.ShouldCreateInitialAdmin()
	if err != nil {
		util.Return(w, false, http.StatusInternalServerError, fmt.Errorf("failed checking if initial admin is required: %w", err), nil, util.Meta{})
		return
	}

	if shouldCreateInitialAdmin {
		util.Return(w, true, http.StatusOK, nil, Session{ShouldCreateInitialAdmin: true}, util.Meta{})
		return
	}

	store := session.GetSessionStore()
	session, err := store.Get(r, session.NAME)
	if err != nil {
		util.Return(w, false, http.StatusInternalServerError, err, nil, util.Meta{})
		return
	}

	var bearerToken string

	reqAuth := r.Header.Get("Authorization")
	splitToken := strings.Split(reqAuth, "Bearer ")
	if len(splitToken) == 2 {
		bearerToken = splitToken[1]
	}

	if bearerToken == "" {
		bearerToken, _ = session.Values["token"].(string)
	}

	user, err := userUserCase.GetUserByToken(bearerToken)
	if err != nil {
		util.Return(w, false, http.StatusInternalServerError, err, nil, util.Meta{})
		return
	}

	if user.ID == 0 {
		util.Return(w, false, http.StatusUnauthorized, nil, nil, util.Meta{})
		return
	}

	if bearerToken == "" {
		bearerToken = user.Token
		session.Values["token"] = bearerToken
		session.Values["userID"] = user.ID
	}

	err = session.Save(r, w)
	if err != nil {
		util.Return(w, false, http.StatusInternalServerError, err, nil, util.Meta{})
		return
	}

	sessionPayload := Session{
		Token: bearerToken,
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Role:  user.Role,
	}

	util.Return(w, true, http.StatusOK, nil, sessionPayload, util.Meta{})
}

func checkPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
