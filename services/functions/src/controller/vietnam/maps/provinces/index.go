package provinces_controller

import (
	"encoding/json"
	"net/http"

	"chatbot-functions/src/constants"
	provinces_service "chatbot-functions/src/services/maps/provinces"

	"github.com/julienschmidt/httprouter"
)

// swagger:route GET /provinces getProvinces
//
// Consumes:
// - application/json
// Produces:
// - application/json
//
// Responses:
//
//	default: errorResponse
func GetProvinces(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	writer.Header().Set("Content-Type", constants.CONTENT_TYPE_APPLICATION_JSON)
	var provinces []provinces_service.Province = provinces_service.GetProvinces()
	json.NewEncoder(writer).Encode(provinces)
}

// swagger:route GET /provinces/:id getProvince
//
// Consumes:
// - application/json
// Produces:
// - application/json
//
// Responses:
//
//	default: errorResponse
func GetProvince(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	writer.Header().Set("Content-Type", constants.CONTENT_TYPE_APPLICATION_JSON)
	var id string = params.ByName("id")
	var province provinces_service.Province = provinces_service.GetProvince(id)
	json.NewEncoder(writer).Encode(province)
}
