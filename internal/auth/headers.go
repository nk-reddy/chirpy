package auth

import (
	"fmt"
	"net/http"
	"strings"
)

func GetBearerToken(headers http.Header) (string, error) {
	auth := headers.Get("Authorization")
	if auth == "" {
		return "", fmt.Errorf("authorization doesn't exist")
	}
	rawToken, found := strings.CutPrefix(auth, "Bearer ")
	if !found || rawToken == "" {
		return "", fmt.Errorf("authorization doesn't exist")
	}

	cleanedToken := strings.TrimSpace(rawToken)
	if cleanedToken == "" {
		return "", fmt.Errorf("authorization doesn't exist")
	}
	return cleanedToken, nil
}

func GetAPIKey(headers http.Header) (string, error) {
	auth := headers.Get("Authorization")
	if auth == "" {
		return "", fmt.Errorf("authorization doesn't exist")
	}
	rawKey, found := strings.CutPrefix(auth, "ApiKey ")
	if !found || rawKey == "" {
		return "", fmt.Errorf("authorization doesn't exist")
	}

	cleanedKey := strings.TrimSpace(rawKey)
	if cleanedKey == "" {
		return "", fmt.Errorf("authorization doesn't exist")
	}
	return cleanedKey, nil
}
