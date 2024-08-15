// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	health "github.com/ryantokmanmokmtm/chat-app-server/app/chat/cmd/api/internal/handler/health"
	ws "github.com/ryantokmanmokmtm/chat-app-server/app/chat/cmd/api/internal/handler/ws"
	"github.com/ryantokmanmokmtm/chat-app-server/app/chat/cmd/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/chat/ping",
				Handler: health.HealthCheckHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/ws",
				Handler: ws.ConnectWsHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/api/v1"),
	)
}
