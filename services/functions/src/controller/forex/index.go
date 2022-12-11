package forex_controller

import (
	forex_client "chatbot-functions/src/clients/forex"
	"chatbot-functions/src/constants"
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func GetForexRates(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	writer.Header().Set("Content-Type", constants.CONTENT_TYPE_APPLICATION_JSON)

	var ratesResponseBody forex_client.RatesResponseBody = forex_client.GetForexRates()

	json.NewEncoder(writer).Encode(ratesResponseBody)
}
