package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

type HealthResponse struct {
	Status string `json:"status"`
}

func GetHealth(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var healthResponse HealthResponse = HealthResponse{"healthy"}
	json.NewEncoder(writer).Encode(healthResponse)
}

func Handler() {
	http.HandleFunc("/", GetHealth)
	log.Println("🚀 Server is listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
