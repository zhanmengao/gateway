package tsws

import (
	"context"
	"github.com/zhanmengao/gateway/internal/ts"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
	"time"
)

func TrafficWsHijack(ctx context.Context, writer http.ResponseWriter, request *http.Request, target *ts.TRouteTarget) {
	var err error
	var srvConn net.Conn
	var cliConn net.Conn
	cliConn, _, err = writer.(http.Hijacker).Hijack()
	if err != nil {
		return
	}
	reqBuf, err := httputil.DumpRequest(request, false)
	if err != nil {
		return
	}
	srvConn, err = net.DialTimeout("tcp", target.Addr, 1*time.Second)
	if err != nil {
		return
	}
	_, err = srvConn.Write(reqBuf)
	if err != nil {
		return
	}
	{
		defer func() {
			cliConn.Close()
			srvConn.Close()
		}()
		if _, err2 := io.Copy(cliConn, srvConn); err2 != nil {
			return
		}
	}
	{
		defer func() {
			cliConn.Close()
			srvConn.Close()
		}()
		_, _ = io.Copy(srvConn, cliConn)
	}
	return
}
