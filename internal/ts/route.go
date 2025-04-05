package ts

import (
	"context"
	"errors"
	"fmt"
	"github.com/zhanmengao/gateway/global"
	"net/http"
	"strings"
)

func route(ctx context.Context, request *http.Request) (target *TRouteTarget, err error) {
	defer func() {
		if err != nil {
			global.Log.Errorf(ctx, "Route %s error = %s ", request.URL.Path, err)
		}
	}()
	routeLock.RLock()
	defer routeLock.RUnlock()
	var matchUrl *global.DBUrl
	//先匹配服务
	for _, u := range urlRouteSlice {
		//前缀是否一致
		if strings.HasPrefix(request.URL.Path, u.URI) {
			matchUrl = u
			break
		}
	}
	//如果没匹配上，返回error
	if matchUrl == nil {
		err = errors.New(fmt.Sprintf("url %s not match", request.URL.Path))
		return
	}
	target = &TRouteTarget{
		DBUrl: matchUrl,
	}

	//从该服务中路由一个节点
	key := ""
	node, err := global.Route.Route(ctx, matchUrl.Service, key)
	if err != nil {
		return
	}

	switch target.Type {
	case global.NETWORK_HTTP:
		target.Addr = node.GetHttpAddr()
	case global.NETWORK_WS:
		target.Addr = node.GetWsAddr()
	}
	return
}
