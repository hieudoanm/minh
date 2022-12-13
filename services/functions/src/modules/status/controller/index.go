package status_controller

import (
	"chatbot-functions/src/common/constants"
	status_client "chatbot-functions/src/modules/status/client"

	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func GetStatuses(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	writer.Header().Set("Content-Type", constants.CONTENT_TYPE_APPLICATION_JSON)
	var serviceStatuses map[string]status_client.ServiceStatus = map[string]status_client.ServiceStatus{}
	for _, service := range status_client.SERVICES {
		serviceStatuses[service] = status_client.GetStatusByService(service)
	}
	json.NewEncoder(writer).Encode(serviceStatuses)
}

func GetStatus(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	writer.Header().Set("Content-Type", constants.CONTENT_TYPE_APPLICATION_JSON)
	var service string = params.ByName("service")
	var serviceStatus status_client.ServiceStatus = status_client.GetStatusByService(service)
	json.NewEncoder(writer).Encode(serviceStatus)
}
