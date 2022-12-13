package health

import (
	"chatbot-functions/src/common/constants"
	health_service "chatbot-functions/src/modules/health/service"
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func GetHealth(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	writer.Header().Set("Content-Type", constants.CONTENT_TYPE_APPLICATION_JSON)
	var healthResponse health_service.HealthResponse = health_service.GetHealth()
	json.NewEncoder(writer).Encode(healthResponse)
}
