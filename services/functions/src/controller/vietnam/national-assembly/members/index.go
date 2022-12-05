package members_controller

import (
	"chatbot-functions/src/constants"
	"encoding/json"
	"net/http"

	members_service "chatbot-functions/src/services/national-assembly/members"

	"github.com/julienschmidt/httprouter"
)

func GetMembers(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	writer.Header().Set("Content-Type", constants.CONTENT_TYPE_APPLICATION_JSON)
	var members = members_service.GetMembers()
	json.NewEncoder(writer).Encode(members)
}

func GetMember(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	writer.Header().Set("Content-Type", constants.CONTENT_TYPE_APPLICATION_JSON)
	var id string = params.ByName("id")
	var member = members_service.GetMember(id)
	json.NewEncoder(writer).Encode(member)
}
