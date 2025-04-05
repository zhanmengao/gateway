package tshttp

import (
	"net/http"
	"time"
)

var (
	transPort = &http.Transport{
		Proxy:                  nil,
		DialContext:            nil,
		DialTLSContext:         nil,
		TLSClientConfig:        nil,
		TLSHandshakeTimeout:    0,
		DisableKeepAlives:      false,
		DisableCompression:     false,
		MaxIdleConns:           800,
		MaxIdleConnsPerHost:    800,
		MaxConnsPerHost:        0,
		IdleConnTimeout:        10 * time.Minute,
		ResponseHeaderTimeout:  0,
		ExpectContinueTimeout:  0,
		TLSNextProto:           nil,
		ProxyConnectHeader:     nil,
		GetProxyConnectHeader:  nil,
		MaxResponseHeaderBytes: 0,
		WriteBufferSize:        0,
		ReadBufferSize:         0,
		ForceAttemptHTTP2:      false,
	}
)
