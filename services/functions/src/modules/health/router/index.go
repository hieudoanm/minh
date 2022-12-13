package health_router

import (
	health_controller "chatbot-functions/src/modules/health/controller"

	"github.com/julienschmidt/httprouter"
)

func HealthRouter(router *httprouter.Router) {
	router.GET("/health", health_controller.GetHealth)
}
