package ts

import (
	"github.com/zhanmengao/gateway/global"
	"github.com/zhanmengao/gateway/internal/tshttp"
	"github.com/zhanmengao/gateway/internal/tsws"
	"net/http"
)

func (s *HttpTrafficServer) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	/*
		Access-Control-Allow-Credentials: true
		Access-Control-Allow-Headers: Content-Type, Content-Length
		Access-Control-Allow-Origin: *
		Access-Control-Allow-Methods: GET,POST,PUT,HEAD,OPTIONS
	*/
	ctx := request.Context()
	if request.Method == http.MethodOptions {
		writer.Header().Set("Access-Control-Allow-Credentials", "true")
		writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length")
		writer.Header().Set("Access-Control-Allow-Origin", "*")
		writer.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,HEAD,OPTIONS")
		writer.WriteHeader(http.StatusNoContent)
		return
	}
	//如果访问根路径，返回200
	if request.URL.Path == "/" {
		writer.Header().Set("Server", "traffic")
		writer.WriteHeader(http.StatusOK)
		return
	}
	target, err := route(ctx, request)
	if err != nil {
		//路由失败，报错
		global.Log.Errorf(ctx, "Route %s error = %s ", request.URL.RawPath, err)
		_, _ = writer.Write([]byte("404 page not found"))
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	switch target.Type {
	case global.NETWORK_HTTP:
		tshttp.TrafficHttpRequest(ctx, writer, request, target)
	case global.NETWORK_WS:
		tsws.TrafficWsHijack(ctx, writer, request, target)
	}
}
