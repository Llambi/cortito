package main

import (
	"log"

	"github.com/Llambi/cortito/internal/adapters/primary/api"
	"github.com/Llambi/cortito/internal/adapters/primary/api/redirect"
	"github.com/Llambi/cortito/internal/adapters/primary/api/status"
	url "github.com/Llambi/cortito/internal/adapters/primary/api/urls"
	urlMemory "github.com/Llambi/cortito/internal/adapters/secondary/repositories/memory"
	redirectService "github.com/Llambi/cortito/internal/core/services/redirect"
	statusService "github.com/Llambi/cortito/internal/core/services/status"
	urlService "github.com/Llambi/cortito/internal/core/services/url"
	"github.com/Llambi/cortito/internal/routing"
	"github.com/Llambi/cortito/platform/routers"
)

func main() {

	store := make(map[string]string)

	urlRepo := urlMemory.Repository{Store: store}

	urlService := urlService.Service{Repo: urlRepo}
	urlHandler := url.Handler{UrlService: urlService}
	urlRouting := routing.UrlRouting(urlHandler)

	statusService := statusService.Service{}
	statusHandler := status.Handler{StatusService: statusService}
	statusRouting := routing.StatusRouting(statusHandler)

	redirectService := redirectService.Service{Repo: urlRepo}
	redirectHandler := redirect.Handler{RedirectService: redirectService}
	redirectRouting := routing.RedirectRouting(redirectHandler)

	var routes []routers.Route
	routes = append(routes, statusRouting...)
	routes = append(routes, urlRouting...)
	routes = append(routes, redirectRouting...)

	app := api.NewApp(routers.NewRouter(routes))

	log.Fatalln(app.Run())
}
