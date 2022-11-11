package forex

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"chatbot-functions/utils"

	"github.com/julienschmidt/httprouter"
)

var API_KEY_FIXER string = utils.Getenv("API_KEY_FIXER", "")

type RatesResponseBody struct {
	Success   bool               `json:"success"`
	Timestamp int                `json:"timestamp"`
	Base      string             `json:"base"`
	Date      string             `json:"date"`
	Rates     map[string]float64 `json:"rates"`
}

func GetForexRates(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json")
	// HTTP Request
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
