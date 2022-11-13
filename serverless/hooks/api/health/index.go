package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type StatusResponse struct {
	Status string `json:"status"`
}

func GetHealth(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var healthResponse StatusResponse = StatusResponse{"healthy"}
	json.NewEncoder(writer).Encode(healthResponse)
}

func Handler() {
	http.HandleFunc("/", GetHealth)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", "8080"), nil))
}
