package ministries_controller

import (
	"chatbot-functions/src/common/constants"
	ministries_service "chatbot-functions/src/modules/vietnam/service/government/ministries"

	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func GetMinistries(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	writer.Header().Set("Content-Type", constants.CONTENT_TYPE_APPLICATION_JSON)
	var ministries = ministries_service.GetMinistries()
	json.NewEncoder(writer).Encode(ministries)
}

func GetMinistry(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	writer.Header().Set("Content-Type", constants.CONTENT_TYPE_APPLICATION_JSON)
	var id string = params.ByName("id")
	var ministry = ministries_service.GetMinistry(id)
	json.NewEncoder(writer).Encode(ministry)
}
