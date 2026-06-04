package auth

import (
	"net/http"
	"strings"
)

func GetAPIKey(headers http.Header) string {
	apiKey := headers.Get("Authorization")

	if apiKey == "" {
		return ""
	}

	trimmed := strings.TrimPrefix(apiKey, "ApiKey ")

	return trimmed
}
