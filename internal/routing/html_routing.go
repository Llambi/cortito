package routing

import (
	"net/http"

	"github.com/Llambi/cortito/internal/adapters/primary/web/html"
	"github.com/Llambi/cortito/platform/routers"
)

func HtmlRouting(handler html.Handler) []routers.Route {
	return []routers.Route{
		{
			Method: http.MethodGet,
			Path:   "/",
			Handle: handler.Home,
		},
	}
}
