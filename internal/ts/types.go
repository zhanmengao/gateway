package ts

import "github.com/zhanmengao/gateway/global"

type HttpTrafficServer struct {
}

type TRouteTarget struct {
	*global.DBUrl
	Addr string
}
