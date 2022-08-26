package handler

import (
	"github.com/zeromicro/go-zero/rest/httpx"
	"go-zero-demo/http/download/internal/svc"
	"go-zero-demo/http/download/internal/types"
	"io/ioutil"
	"net/http"
)

func DownloadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request

		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		body, err := ioutil.ReadFile(req.File)
		if err != nil {
			httpx.Error(w, err)
			return
		}

		w.Write(body)
	}
}
