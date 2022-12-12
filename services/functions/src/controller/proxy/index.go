package proxy_controller

import (
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"

	"github.com/julienschmidt/httprouter"
)

func GetLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

func Proxy(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	var localIP string = GetLocalIP()
	log.Println("localIP: " + localIP)
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
