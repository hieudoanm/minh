package presidents_controller

import (
	"chatbot-functions/src/common/constants"
	presidents_service "chatbot-functions/src/modules/vietnam/service/government/presidents"

	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func GetPresidents(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	writer.Header().Set("Content-Type", constants.CONTENT_TYPE_APPLICATION_JSON)
	var presidents = presidents_service.GetPresidents()
	json.NewEncoder(writer).Encode(presidents)
}

func GetPresident(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	writer.Header().Set("Content-Type", constants.CONTENT_TYPE_APPLICATION_JSON)
	var id string = params.ByName("id")
	var president = presidents_service.GetPresident(id)
	json.NewEncoder(writer).Encode(president)
}
