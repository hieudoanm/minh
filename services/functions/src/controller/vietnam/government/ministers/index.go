package ministers_controller

import (
	"chatbot-functions/src/constants"
	"encoding/json"
	"net/http"

	ministers_service "chatbot-functions/src/services/government/ministers"

	"github.com/julienschmidt/httprouter"
)

func GetMinisters(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	writer.Header().Set("Content-Type", constants.CONTENT_TYPE_APPLICATION_JSON)
	var ministers = ministers_service.GetMinisters()
	json.NewEncoder(writer).Encode(ministers)
}

func GetMinister(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	writer.Header().Set("Content-Type", constants.CONTENT_TYPE_APPLICATION_JSON)
	var id string = params.ByName("id")
	var minister = ministers_service.GetMinister(id)
	json.NewEncoder(writer).Encode(minister)
}
