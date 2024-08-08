package routing

import (
	"net/http"

	"github.com/Llambi/cortito/internal/adapters/primary/api/status"
	"github.com/Llambi/cortito/platform/routers"
)

func StatusRouting(handler status.Handler) []routers.Route {
	return []routers.Route{
		{
			Method: http.MethodGet,
			Path:   "/status",
			Handle: handler.Status,
		},
	}
}
