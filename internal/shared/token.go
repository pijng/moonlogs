package shared

import (
	"crypto/rand"
	"encoding/base64"
)

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
