package health

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type HealthResponse struct {
	Status string `json:"status"`
}

func GetHealth(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json")
	var healthResponse HealthResponse = HealthResponse{"healthy"}
	json.NewEncoder(writer).Encode(healthResponse)
}
