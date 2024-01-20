package session

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"

	"github.com/gorilla/sessions"
)

type sessionKey string

const (
	NAME            = "moonlogs"
	KEY  sessionKey = "moonlogs"
)

var store *sessions.CookieStore

func RegisterSessionStore() *sessions.CookieStore {
	store = sessions.NewCookieStore([]byte("moonlogs"))

	store.Options = &sessions.Options{
		MaxAge:   86400 * 30,
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode, // TODO: change this MAYBE
		Secure:   true,
	}

	return store
}

func GetSessionStore() *sessions.CookieStore {
	return store
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
