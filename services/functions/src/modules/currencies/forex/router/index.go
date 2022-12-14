package forex_router

import (
	forex_controller "chatbot-functions/src/modules/currencies/forex/controller"

	"github.com/julienschmidt/httprouter"
)

func Router(router *httprouter.Router) {
	router.GET("/forex/rates", forex_controller.GetForexRates)

}
