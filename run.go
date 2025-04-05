package traffic

import (
	"fmt"
	"github.com/zhanmengao/gateway/global"
	"github.com/zhanmengao/gateway/internal/ts"
	"net/http"
)

func Run(port int64, rt global.IRoute, urlList []*global.DBUrl) (err error) {
	ts.SetURLList(urlList)
	global.Route = rt

	hSrv := http.NewServeMux()
	hSrv.Handle("/", &ts.HttpTrafficServer{})
	if err = http.ListenAndServe(fmt.Sprintf(":%d", port), hSrv); err != nil {
		panic(err)
	}
	select {}
}
