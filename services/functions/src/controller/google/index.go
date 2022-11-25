package google_controller

import (
	google_client "chatbot-functions/src/clients/google"

	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func GetGoogleTrends(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json")

	var trendsResponseBody = google_client.GetTrends()

	json.NewEncoder(writer).Encode(trendsResponseBody)
}
