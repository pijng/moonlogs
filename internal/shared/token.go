package shared

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"strings"
)

const TokenOffset = 7

func GenerateRandomToken(length int) (string, error) {
	numBytes := (length * 6) / 8

	randomBytes := make([]byte, numBytes)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}

	token := base64.URLEncoding.EncodeToString(randomBytes)

	token = token[:length]

	return token, nil
}

func ExtractBearerToken(r *http.Request) string {
	authorizationHeader := r.Header.Get("Authorization")
	if strings.HasPrefix(authorizationHeader, "Bearer ") {
		return authorizationHeader[TokenOffset:]
	}

	return ""
}
