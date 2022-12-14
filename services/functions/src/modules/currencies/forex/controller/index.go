package forex_controller

import (
	"chatbot-functions/src/common/constants"
	forex_client "chatbot-functions/src/modules/currencies/forex/client"
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func GetForexRates(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	writer.Header().Set("Content-Type", constants.CONTENT_TYPE_APPLICATION_JSON)

	var ratesResponseBody forex_client.RatesResponseBody = forex_client.GetForexRates()

	json.NewEncoder(writer).Encode(ratesResponseBody)
}
