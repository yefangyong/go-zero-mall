// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"mall/service/order/api/internal/svc"

	"github.com/tal-tech/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/order/create",
				Handler: CreateHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/order/update",
				Handler: UpdateHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/order/remove",
				Handler: RemoveHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/order/detail",
				Handler: DetailHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/order/list",
				Handler: ListHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)
}
