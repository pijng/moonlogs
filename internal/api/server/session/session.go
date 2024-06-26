package session

import (
	"crypto/rand"
	"encoding/base64"
	"moonlogs/internal/entities"
	"moonlogs/internal/usecases"
	"net/http"

	"github.com/gorilla/sessions"
)

type sessionKey string

const (
	NAME            = "moonlogs"
	KEY  sessionKey = "moonlogs"
)

type SessionHelper struct {
	store       *sessions.CookieStore
	userUseCase *usecases.UserUseCase
}

var sessionHelper *SessionHelper

func RegisterSessionStore(userUseCase *usecases.UserUseCase) *SessionHelper {
	store := sessions.NewCookieStore([]byte("moonlogs"))

	store.Options = &sessions.Options{
		MaxAge:   86400 * 30,
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode, // TODO: change this MAYBE
		Secure:   true,
	}

	sessionHelper = &SessionHelper{
		store:       store,
		userUseCase: userUseCase,
	}

	return sessionHelper
}

func GetSessionStore() *sessions.CookieStore {
	return sessionHelper.store
}

func GenerateAuthToken() (string, error) {
	tokenLength := 16
	tokenBytes := make([]byte, tokenLength)

	_, err := rand.Read(tokenBytes)
	if err != nil {
		return "", err
	}

	authToken := base64.URLEncoding.EncodeToString(tokenBytes)

	return authToken, nil
}

func GetUserFromContext(r *http.Request) *entities.User {
	sessionStore, ok := r.Context().Value(KEY).(*sessions.Session)
	if !ok {
		return nil
	}

	iUserID, ok := sessionStore.Values["userID"]
	if !ok {
		return nil
	}

	userID, ok := iUserID.(int)
	if !ok {
		return nil
	}

	user, err := sessionHelper.userUseCase.GetUserByID(r.Context(), userID)
	if err != nil {
		return nil
	}

	return user
}
