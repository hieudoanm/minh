package committees_controller

import (
	"chatbot-functions/src/constants"
	"encoding/json"
	"net/http"

	committees_service "chatbot-functions/src/services/national-assembly/committees"

	"github.com/julienschmidt/httprouter"
)

func GetCommittees(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	writer.Header().Set("Content-Type", constants.CONTENT_TYPE_APPLICATION_JSON)
	var committees = committees_service.GetCommittees()
	json.NewEncoder(writer).Encode(committees)
}

func GetCommittee(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	writer.Header().Set("Content-Type", constants.CONTENT_TYPE_APPLICATION_JSON)
	var id string = params.ByName("id")
	var committee = committees_service.GetCommittee(id)
	json.NewEncoder(writer).Encode(committee)
}
