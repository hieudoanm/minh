package districts_controller

import (
	"encoding/json"
	"net/http"

	"chatbot-functions/src/common/constants"
	districts_service "chatbot-functions/src/modules/vietnam/service/maps/districts"

	"github.com/julienschmidt/httprouter"
)

// swagger:route GET /districts getDistricts
//
// Consumes:
// - application/json
// Produces:
// - application/json
//
// Responses:
//
//	default: errorResponse
func GetDistricts(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	writer.Header().Set("Content-Type", constants.CONTENT_TYPE_APPLICATION_JSON)
	var districts []districts_service.District = districts_service.GetDistricts()
	json.NewEncoder(writer).Encode(districts)

}

// swagger:route GET /districts/:id getDistrict
//
// Consumes:
// - application/json
// Produces:
// - application/json
//
// Responses:
//
//	default: errorResponse
func GetDistrict(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	writer.Header().Set("Content-Type", constants.CONTENT_TYPE_APPLICATION_JSON)
	var id string = params.ByName("id")
	var district districts_service.District = districts_service.GetDistrict(id)
	json.NewEncoder(writer).Encode(district)
}
