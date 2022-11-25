package forex_controller

import (
	forex_client "chatbot-functions/src/clients/forex"
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func GetForexRates(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json")

	var ratesResponseBody forex_client.RatesResponseBody = forex_client.GetForexRates()

	json.NewEncoder(writer).Encode(ratesResponseBody)
}
