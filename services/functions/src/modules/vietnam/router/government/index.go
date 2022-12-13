package vietnam_government_router

import (
	ministers_controller "chatbot-functions/src/modules/vietnam/controller/government/ministers"
	ministries_controller "chatbot-functions/src/modules/vietnam/controller/government/ministries"
	presidents_controller "chatbot-functions/src/modules/vietnam/controller/government/presidents"

	"github.com/julienschmidt/httprouter"
)

func VietnamGovernmentRouter(router *httprouter.Router) {
	router.GET("/vietnam/government/ministers", ministers_controller.GetMinisters)
	router.GET("/vietnam/government/ministers/:id", ministers_controller.GetMinister)
	router.GET("/vietnam/government/ministries", ministries_controller.GetMinistries)
	router.GET("/vietnam/government/ministries/:id", ministries_controller.GetMinistry)
	router.GET("/vietnam/government/presidents", presidents_controller.GetPresidents)
	router.GET("/vietnam/government/presidents/:id", presidents_controller.GetPresident)
}
