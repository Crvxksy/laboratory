// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	verify "looklook/app/identity/cmd/api/internal/handler/verify"
	"looklook/app/identity/cmd/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/verify/token",
				Handler: verify.TokenHandler(serverCtx),
			},
		},
		rest.WithPrefix("/identity/v1"),
	)
}
