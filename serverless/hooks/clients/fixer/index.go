package fixer

import (
	"encoding/json"
	"fmt"

	"chatbot-webhook/libs/http"
)

const BASE_URL = "http://data.fixer.io/api"

type LatestResponse struct {
	Success   bool               `json:"success"`
	Timestamp int                `json:"timestamp"`
	Base      string             `json:"base"`
	Date      string             `json:"date"`
	Rates     map[string]float64 `json:"rates"`
}

func GetLatest(accessKey string) (LatestResponse, error) {
	var url string = fmt.Sprintf("%s/latest?access_key=%s", BASE_URL, accessKey)

	body, getError := http.Get(url)
	if getError != nil {
		return LatestResponse{}, getError
	}

	var latestResponse LatestResponse
	jsonUnmarshalError := json.Unmarshal(body, &latestResponse)
	if jsonUnmarshalError != nil {
		return LatestResponse{}, jsonUnmarshalError
	}

	return latestResponse, nil
}

type SymbolsResponse struct {
	Success bool              `json:"success"`
	Symbols map[string]string `json:"symbols"`
}

func GetSymbols(accessKey string) (SymbolsResponse, error) {
	var url string = fmt.Sprintf("%s/symbols?access_key=%s", BASE_URL, accessKey)

	body, getError := http.Get(url)
	if getError != nil {
		return SymbolsResponse{}, getError
	}

	var symbolsResponse SymbolsResponse
	jsonUnmarshalError := json.Unmarshal(body, &symbolsResponse)
	if jsonUnmarshalError != nil {
		return SymbolsResponse{}, jsonUnmarshalError
	}

	return symbolsResponse, nil
}
