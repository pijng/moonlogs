package controllers

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"moonlogs/api/server/session"
	"moonlogs/api/server/util"
	"moonlogs/internal/repository"
	"net/http"
	"strings"
)

type Credentials struct {
	Email    string
	Password string
}

type Session struct {
	Token string `json:"token"`
}

var SHA256Hasher = sha256.New()

func Login(w http.ResponseWriter, r *http.Request) {
	var token string
	var credentials Credentials

	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		util.Return(w, false, http.StatusBadRequest, err, nil, util.Meta{})
		return
	}

	userRepository := repository.NewUserRepository(r.Context())

	user, err := userRepository.GetByEmail(credentials.Email)
	if err != nil {
		util.Return(w, false, http.StatusNotFound, err, nil, util.Meta{})
		return
	}

	_, err = SHA256Hasher.Write([]byte(credentials.Password))
	if err != nil {
		util.Return(w, false, http.StatusInternalServerError, err, nil, util.Meta{})
		return
	}

	hashBytes := SHA256Hasher.Sum(nil)
	hashString := hex.EncodeToString(hashBytes)
	SHA256Hasher.Reset()

	if hashString != user.PasswordDigest {
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
		err = userRepository.UpdateTokenById(user.ID, token)
		if err != nil {
			util.Return(w, false, http.StatusInternalServerError, err, nil, util.Meta{})
			return
		}
	}

	util.Return(w, true, http.StatusOK, nil, Session{Token: token}, util.Meta{})
}

func GetSession(w http.ResponseWriter, r *http.Request) {
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

	user, _ := repository.NewUserRepository(r.Context()).GetByToken(bearerToken)

	token, ok := session.Values["token"].(string)
	if !ok && user == nil {
		util.Return(w, false, http.StatusUnauthorized, nil, nil, util.Meta{})
		return
	}

	if token == "" && user != nil {
		token = user.Token
		session.Values["token"] = token
		session.Values["userID"] = user.ID
	}

	err = session.Save(r, w)
	if err != nil {
		util.Return(w, false, http.StatusInternalServerError, err, nil, util.Meta{})
		return
	}

	util.Return(w, true, http.StatusOK, nil, Session{Token: token}, util.Meta{})
}
