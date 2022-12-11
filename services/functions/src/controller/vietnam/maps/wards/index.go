package wards_controller

import (
	"encoding/json"
	"net/http"

	"chatbot-functions/src/constants"
	wards_service "chatbot-functions/src/services/vietnam/maps/wards"

	"github.com/julienschmidt/httprouter"
)

// swagger:route GET /wards getWards
//
// Consumes:
// - application/json
// Produces:
// - application/json
//
// Responses:
//
//	default: errorResponse
func GetWards(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	writer.Header().Set("Content-Type", constants.CONTENT_TYPE_APPLICATION_JSON)
	var wards []wards_service.Ward = wards_service.GetWards()
	json.NewEncoder(writer).Encode(wards)
}

// swagger:route GET /wards/:id getWard
//
// Consumes:
// - application/json
// Produces:
// - application/json
//
// Responses:
//
//	default: errorResponse
func GetWard(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	writer.Header().Set("Content-Type", constants.CONTENT_TYPE_APPLICATION_JSON)
	var id string = params.ByName("id")
	var ward wards_service.Ward = wards_service.GetWard(id)
	json.NewEncoder(writer).Encode(ward)
}
