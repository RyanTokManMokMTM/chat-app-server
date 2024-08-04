// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	assets "api/app/assets/cmd/api/internal/handler/assets"
	health "api/app/assets/cmd/api/internal/handler/health"
	"api/app/assets/cmd/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				// Upload only image
				Method:  http.MethodPost,
				Path:    "/file/image/upload",
				Handler: assets.UploadImageHandler(serverCtx),
			},
			{
				// Upload any file
				Method:  http.MethodPost,
				Path:    "/file/upload",
				Handler: assets.UploadFileHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/ping",
				Handler: health.HealthCheckHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/v1"),
	)
}