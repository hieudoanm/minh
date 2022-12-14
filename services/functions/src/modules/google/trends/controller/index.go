package trends_controller

import (
	"chatbot-functions/src/common/constants"
	trends_client "chatbot-functions/src/modules/google/trends/client"

	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func GetGoogleTrends(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	writer.Header().Set("Content-Type", constants.CONTENT_TYPE_APPLICATION_JSON)

	var trendsResponseBody = trends_client.GetTrends()

	json.NewEncoder(writer).Encode(trendsResponseBody)
}
