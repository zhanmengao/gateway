package tshttp

import (
	"context"
	"fmt"
	"github.com/zhanmengao/gateway/global"
	"github.com/zhanmengao/gateway/internal/ts"
	"net/http"
	"net/http/httputil"
	"net/url"
)

// TrafficHttpRequest 收到CLB请求，转发给游戏服务
func TrafficHttpRequest(ctx context.Context, writer http.ResponseWriter, request *http.Request, target *ts.TRouteTarget) {
	proxyReverse(ctx, writer, request, target)
}

func proxyReverse(ctx context.Context, writer http.ResponseWriter, request *http.Request, target *ts.TRouteTarget) {
	u, err := url.Parse(fmt.Sprintf("http://%s%s", target.Addr, request.RequestURI))
	if err != nil {
		global.Log.Errorf(ctx, "Parse %s error = %s ", target.Addr, err)
		return
	}
	proxy := &httputil.ReverseProxy{
		Director: func(request *http.Request) {
			request.URL = u
		},
		Transport: transPort,
	}
	proxy.ErrorHandler = func(writer http.ResponseWriter, request *http.Request, err error) {
		global.Log.Errorf(ctx, "write to %s.%s error = %s ", u.String(), target, err)
	}
	proxy.ModifyResponse = func(response *http.Response) error {
		if !global.UseDebug {
			return nil
		}
		//打印一下回包
		body, err := httputil.DumpResponse(response, true)
		//打印一下回包
		if err != nil {
			msg := fmt.Sprintf("path : %s ReadAll error = %s ", request.URL.Path, err.Error())
			global.Log.Errorf(ctx, msg)
		} else {
			global.Log.Infof(ctx, "rcv %s ", string(body))
		}
		return nil
	}
	proxy.ServeHTTP(writer, request)
}
