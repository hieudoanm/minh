package health_router

import (
	health_controller "chatbot-functions/src/controller/health"

	"github.com/julienschmidt/httprouter"
)

func HealthRouter(router *httprouter.Router) {
	router.GET("/health", health_controller.GetHealth)
}
