package status_router

import (
	status_controller "chatbot-functions/src/controller/status"

	"github.com/julienschmidt/httprouter"
)

func StatusRouter(router *httprouter.Router) {
	router.GET("/status", status_controller.GetStatuses)
	router.GET("/status/:service", status_controller.GetStatus)
}
