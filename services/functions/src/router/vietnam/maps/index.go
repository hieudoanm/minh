package vietnam_maps_router

import (
	districts_controller "chatbot-functions/src/controller/vietnam/maps/districts"
	provinces_controller "chatbot-functions/src/controller/vietnam/maps/provinces"
	wards_controller "chatbot-functions/src/controller/vietnam/maps/wards"

	"github.com/julienschmidt/httprouter"
)

func VietnamMapsRouter(router *httprouter.Router) {
	router.GET("/vietnam/maps/provinces", provinces_controller.GetProvinces)
	router.GET("/vietnam/maps/provinces/:id", provinces_controller.GetProvince)
	router.GET("/vietnam/maps/districts", districts_controller.GetDistricts)
	router.GET("/vietnam/maps/districts/:id", districts_controller.GetDistricts)
	router.GET("/vietnam/maps/wards", wards_controller.GetWards)
	router.GET("/vietnam/maps/wards/:id", wards_controller.GetWard)
}
