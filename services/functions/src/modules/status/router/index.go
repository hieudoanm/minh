package status_router

import (
	status_controller "chatbot-functions/src/modules/status/controller"

	"github.com/julienschmidt/httprouter"
)

func Router(router *httprouter.Router) {
	router.GET("/status", status_controller.GetStatuses)
	router.GET("/status/:service", status_controller.GetStatus)
}
