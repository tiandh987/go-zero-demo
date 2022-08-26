package main

import (
	"flag"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/rest"
)

var (
	port    = flag.Int("port", 3333, "the port to listen")
	timeout = flag.Int64("timeout", 0, "timeout of milliseconds")
)

type Request struct {
	User string `form:"user,options=a|b"`
}

func first(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("X-Middleware", "first")
		next(w, r)
	}
}

func second(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("X-Middleware", "second")
		next(w, r)
	}
}

func handle(w http.ResponseWriter, r *http.Request) {
	var req Request
	err := httpx.Parse(r, &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	httpx.OkJson(w, "hello, "+req.User)
}

func main() {
	flag.Parse()

	engine := rest.MustNewServer(rest.RestConf{
		ServiceConf: service.ServiceConf{
			Log: logx.LogConf{
				Mode: "console",
			},
		},
		Port:     *port,
		Timeout:  *timeout,
		MaxConns: 500,
	})
	defer engine.Stop()

	engine.Use(first)
	engine.Use(second)
	engine.AddRoute(rest.Route{
		Method:  http.MethodGet,
		Path:    "/demo",
		Handler: handle,
	})
	engine.Start()
}
