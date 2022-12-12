package proxy_router

import (
	proxy_controller "chatbot-functions/src/controller/proxy"

	"github.com/julienschmidt/httprouter"
)

func ProxyRouter(router *httprouter.Router) {
	router.GET("/proxy", proxy_controller.Proxy)
}
