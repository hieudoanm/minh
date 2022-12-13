package proxy_router

import (
	proxy_controller "chatbot-functions/src/modules/proxy/controller"

	"github.com/julienschmidt/httprouter"
)

func ProxyRouter(router *httprouter.Router) {
	router.GET("/proxy", proxy_controller.Proxy)
}
