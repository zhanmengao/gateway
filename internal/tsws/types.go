package tsws

import (
	"net/http"
	"sync"
)

type TrafficWsPair struct {
	lock    sync.Mutex
	writer  http.ResponseWriter
	request *http.Request
}
