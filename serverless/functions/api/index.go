package handler

import (
	"encoding/json"
	"net/http"
)

type HealthResponse struct {
	Status string `json:"status"`
}

func Handler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var healthResponse HealthResponse = HealthResponse{"healthy"}
	json.NewEncoder(writer).Encode(healthResponse)
}
