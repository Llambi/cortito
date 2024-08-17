package routing

import (
	"net/http"

	"github.com/Llambi/cortito/internal/adapters/primary/api/redirect"
	"github.com/Llambi/cortito/platform/routers"
)

func RedirectRouting(handler redirect.Handler) []routers.Route {
	return []routers.Route{
		{
			Method: http.MethodGet,
			Path:   "/r/:shortUrl",
			Handle: handler.Redirect,
		},
	}
}
