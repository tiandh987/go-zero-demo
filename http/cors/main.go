package main

import (
	"flag"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

var configFile = flag.String("f", "config.yaml", "the config file")

type Request struct {
	User string `form:"user"`
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

func handleHello(w http.ResponseWriter, r *http.Request) {
	var req Request
	err := httpx.Parse(r, &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	httpx.OkJson(w, "welcome, "+req.User)
}

func main() {
	flag.Parse()

	logx.DisableStat()

	var c rest.RestConf
	conf.MustLoad(*configFile, &c)
	srv := rest.MustNewServer(c, rest.WithCors())
	defer srv.Stop()

	// 使用中间件
	srv.Use(first)
	srv.Use(second)

	srv.AddRoute(rest.Route{
		Method:  http.MethodGet,
		Path:    "/hello",
		Handler: handleHello,
	})

	srv.Start()
}
