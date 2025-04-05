package traffic

import (
	"github.com/zhanmengao/gateway/global"
	"github.com/zhanmengao/gateway/internal/ts"
)

func Run(port int64, rt global.IRoute, urlList []*global.DBUrl) (err error) {
	ts.SetURLList(urlList)
	select {}
}
