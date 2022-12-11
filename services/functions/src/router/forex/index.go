package forex_router

import (
	forex_controller "chatbot-functions/src/controller/forex"

	"github.com/julienschmidt/httprouter"
)

func ForexRouter(router *httprouter.Router) {
	router.GET("/forex/rates", forex_controller.GetForexRates)

}
