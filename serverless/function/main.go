package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func GetCryptoCoins(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var healthResponse HealthResponse = HealthResponse{"healthy"}
	json.NewEncoder(writer).Encode(healthResponse)
}

func GetForexRates(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var healthResponse HealthResponse = HealthResponse{"healthy"}
	json.NewEncoder(writer).Encode(healthResponse)
}

type HealthResponse struct {
	Status string `json:"status"`
}

func GetHeath(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var healthResponse HealthResponse = HealthResponse{"healthy"}
	json.NewEncoder(writer).Encode(healthResponse)
}

func GetVnindexCompanies(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var healthResponse HealthResponse = HealthResponse{"healthy"}
	json.NewEncoder(writer).Encode(healthResponse)
}

func GetVnindexHistory(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var healthResponse HealthResponse = HealthResponse{"healthy"}
	json.NewEncoder(writer).Encode(healthResponse)
}

func main() {
	// Router
	http.HandleFunc("/crypto/coins", GetCryptoCoins)
	http.HandleFunc("/forex/rates", GetForexRates)
	http.HandleFunc("/health", GetHeath)
	http.HandleFunc("/vnindex/companies", GetVnindexCompanies)
	http.HandleFunc("/vnindex/history", GetVnindexHistory)
	// Start
	log.Println("Server is listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
