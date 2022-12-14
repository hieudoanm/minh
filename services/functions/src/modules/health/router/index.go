package health_router

import (
	health_controller "chatbot-functions/src/modules/health/controller"

	"github.com/julienschmidt/httprouter"
)

func Router(router *httprouter.Router) {
	router.GET("/health", health_controller.GetHealth)
}
