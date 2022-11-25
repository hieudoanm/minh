package forex_controller

import (
	forex_client "chatbot-functions/src/clients/forex"
	"chatbot-functions/src/utils"
	"encoding/json"
	"net/http"

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

	var ratesResponseBody forex_client.RatesResponseBody = forex_client.GetForexRates()

	json.NewEncoder(writer).Encode(ratesResponseBody)
}
