package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type RatesResponseBody struct {
	Success   bool               `json:"success"`
	Timestamp int                `json:"timestamp"`
	Base      string             `json:"base"`
	Date      string             `json:"date"`
	Rates     map[string]float64 `json:"rates"`
}

func GetForexRates(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	// HTTP Request
	var API_KEY_FIXER string = os.Getenv("API_KEY_FIXER")
	var url string = fmt.Sprintf(
		"http://data.fixer.io/api/latest?access_key=%s",
		API_KEY_FIXER,
	)
	response, httpGetError := http.Get(url)
	if httpGetError != nil {
		log.Println("Fail to GetForexRates", httpGetError)
	}
	defer response.Body.Close()
	// Get Body
	body, readBodyError := io.ReadAll(response.Body)
	if readBodyError != nil {
		log.Println("Fail to GetForexRates", readBodyError)
	}
	// Parse JSON
	var ratesResponseBody RatesResponseBody
	jsonUnmarshalError := json.Unmarshal(body, &ratesResponseBody)
	if jsonUnmarshalError != nil {
		log.Println("Fail to GetForexRates", jsonUnmarshalError)
	}

	json.NewEncoder(writer).Encode(ratesResponseBody)
}

func Handler() {
	http.HandleFunc("/", GetForexRates)
	log.Println("🚀 Server is listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
