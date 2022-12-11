package health

import (
	"chatbot-functions/src/constants"
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type HealthResponse struct {
	Status string `json:"status"`
}

func GetHealth(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	writer.Header().Set("Content-Type", constants.CONTENT_TYPE_APPLICATION_JSON)
	var healthResponse HealthResponse = HealthResponse{"healthy"}
	json.NewEncoder(writer).Encode(healthResponse)
}
