package routing

import (
	"net/http"

	url "github.com/Llambi/cortito/internal/adapters/primary/api/urls"
	"github.com/Llambi/cortito/platform/routers"
)

func UrlRouting(handler url.Handler) []routers.Route {
	return []routers.Route {
		{
			Method: http.MethodGet,
			Path: "/url",
			Handle: handler.GetByKey,
		},
		{
			Method: http.MethodPost,
			Path: "/url",
			Handle: handler.Create,
		},
	}
}