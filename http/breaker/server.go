package main

import (
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/core/stat"
	"github.com/zeromicro/go-zero/core/syncx"
	"github.com/zeromicro/go-zero/rest"
	"net/http"
	"runtime"
	"time"
)

func main() {
	logx.Disable()

	stat.SetReporter(nil)

	server := rest.MustNewServer(rest.RestConf{
		ServiceConf: service.ServiceConf{
			Name: "breaker",
			Log: logx.LogConf{
				Mode: "console",
			},
		},
		Host:     "0.0.0.0",
		Port:     8080,
		MaxConns: 1000,
		Timeout:  3000,
	})

	latch := syncx.NewLimit(10)

	server.AddRoute(rest.Route{
		Method: http.MethodGet,
		Path:   "/heavy",
		Handler: func(w http.ResponseWriter, r *http.Request) {
			if latch.TryBorrow() {
				defer latch.Return()

				runtime.LockOSThread()
				defer runtime.UnlockOSThread()

				begin := time.Now()

				for {
					if time.Now().Sub(begin) > time.Millisecond*50 {
						break
					}
				}
			} else {
				w.WriteHeader(http.StatusInternalServerError)
			}
		},
	})

	server.AddRoute(rest.Route{
		Method: http.MethodGet,
		Path:   "/good",
		Handler: func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		},
	})

	defer server.Stop()
	server.Start()
}
