package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

const TABLEBASE_URL = "https://raw.githubusercontent.com/hieudoanm/tablebase/master/json"

type HistoryResponseBody struct {
	Date      string `json:"date"`
	Symbol    string `json:"symbol"`
	Open      string `json:"open"`
	High      string `json:"high"`
	Low       string `json:"low"`
	Close     string `json:"close"`
	Volume    string `json:"volume"`
	Timestamp string `json:"timestamp"`
}

type ErrorResponse struct {
	Message string `json:"message"`
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

func GetVnindexHistory(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var symbol string = GetQueryParameter(request, "symbol", "")
	if symbol == "" {
		var errorResponse ErrorResponse = ErrorResponse{Message: "Missing Symbol"}
		json.NewEncoder(writer).Encode(errorResponse)
		return
	}
	// HTTP Request
	var url string = fmt.Sprintf(
		"%s/vietnam/stock/history/%s.json",
		TABLEBASE_URL,
		strings.ToUpper(symbol),
	)
	response, httpGetError := http.Get(url)
	if httpGetError != nil {
		log.Println("Fail to GetVnindexHistory", httpGetError)
	}
	defer response.Body.Close()
	// Get Body
	body, readBodyError := io.ReadAll(response.Body)
	if readBodyError != nil {
		log.Println("Fail to GetVnindexHistory", readBodyError)
	}
	// Parse JSON
	var companiesResponseBody []HistoryResponseBody
	jsonUnmarshalError := json.Unmarshal(body, &companiesResponseBody)
	if jsonUnmarshalError != nil {
		log.Println("Fail to GetVnindexHistory", jsonUnmarshalError)
	}

	json.NewEncoder(writer).Encode(companiesResponseBody)
}

func Handler() {
	http.HandleFunc("/", GetVnindexHistory)
	log.Println("🚀 Server is listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
