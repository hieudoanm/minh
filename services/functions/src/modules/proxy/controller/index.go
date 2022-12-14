package proxy_controller

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/julienschmidt/httprouter"

	proxy_service "chatbot-functions/src/modules/proxy/service"
)

func Proxy(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	var ip string = proxy_service.GetIP()
	log.Println("IP: " + ip)
	var urlQuery string = request.URL.Query().Get("url")
	if urlQuery == "" {
		urlQuery = "https://example.com/"
	}
	url, _ := url.Parse(urlQuery)
	proxy := httputil.NewSingleHostReverseProxy(url)
	request.Host = url.Host
	request.URL.Host = url.Host
	request.URL.Scheme = url.Scheme
	var path string = request.URL.Path
	var trimPath string = strings.TrimLeft(path, "/")
	request.URL.Path = trimPath
	proxy.ServeHTTP(writer, request)
}
