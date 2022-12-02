package status_controller

import (
	status_client "chatbot-functions/src/clients/status"

	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func GetStatuses(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json")
	var serviceStatuses map[string]status_client.ServiceStatus = map[string]status_client.ServiceStatus{}
	for _, service := range status_client.SERVICES {
		serviceStatuses[service] = status_client.GetStatusByService(service)
	}
	json.NewEncoder(writer).Encode(serviceStatuses)
}

func GetStatus(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	writer.Header().Set("Content-Type", "application/json")
	var service string = params.ByName("service")
	var serviceStatus status_client.ServiceStatus = status_client.GetStatusByService(service)
	json.NewEncoder(writer).Encode(serviceStatus)
}