package ts

import (
	"github.com/zhanmengao/gateway/global"
	"sync"
)

type TRouteList []*global.DBUrl

var (
	routeLock     sync.RWMutex
	urlRouteSlice TRouteList
)

func SetURLList(ul []*global.DBUrl) {
	urlRouteSlice = ul
}

func (r TRouteList) Len() int {
	return len(r)
}
func (r TRouteList) Swap(i, j int) {
	tmp := r[i]
	r[i] = r[j]
	r[j] = tmp
}
func (r TRouteList) Less(i, j int) bool {
	//长前缀放后边
	if len(r[i].URI) > len(r[j].URI) {
		return true
	} else if len(r[i].URI) < len(r[j].URI) {
		return false
	} else if r[i].Service < r[j].Service {
		return true
	} else if r[i].Service < r[j].Service {
		return false
	} else {
		return r[i].Type < r[j].Type
	}
}

func UrlList2Proto(srvName string, httpURL []string, wsURL []string) []*global.DBUrl {
	urlList := make([]*global.DBUrl, 0, len(httpURL)+len(wsURL))
	for _, u := range httpURL {
		urlList = append(urlList, &global.DBUrl{
			Service: srvName,
			URI:     u,
			Type:    global.NETWORK_HTTP,
		})
	}
	for _, u := range wsURL {
		urlList = append(urlList, &global.DBUrl{
			Service: srvName,
			URI:     u,
			Type:    global.NETWORK_WS,
		})
	}
	return urlList
}
