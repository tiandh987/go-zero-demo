package handler

import (
	"github.com/zeromicro/go-zero/rest"
	"go-zero-demo/http/download/internal/svc"
	"net/http"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/static/:file",
				Handler: DownloadHandler(serverCtx),
			},
		},
	)
}
