package global

var (
	Log      ILog
	UseDebug bool
	Route    IRoute
)

func init() {
	Log = &consoleLog{}
}
