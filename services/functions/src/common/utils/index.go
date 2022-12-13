package utils

import (
	"net/http"
	"os"
)

func Getenv(key string, defaultValue string) string {
	var value string = os.Getenv(key)
	if defaultValue == "" {
		return value
	}
	if defaultValue != "" && value == "" {
		return defaultValue
	}
	return value
}

func GetQueryParameter(request *http.Request, key string, defaultValue string) string {
	var value string = request.URL.Query().Get(key)
	if defaultValue == "" {
		return value
	}
	if defaultValue != "" && value == "" {
		return defaultValue
	}
	return value
}
