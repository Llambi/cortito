package main

import (
	"log"
	"time"

	"github.com/Llambi/cortito/internal/adapters/primary/api"
	"github.com/Llambi/cortito/internal/adapters/primary/api/redirect"
	"github.com/Llambi/cortito/internal/adapters/primary/api/status"
	url "github.com/Llambi/cortito/internal/adapters/primary/api/urls"
	"github.com/Llambi/cortito/internal/adapters/primary/web/html"
	urlMemory "github.com/Llambi/cortito/internal/adapters/secondary/repositories/url/memory"
	userMemory "github.com/Llambi/cortito/internal/adapters/secondary/repositories/user/memory"
	redirectService "github.com/Llambi/cortito/internal/core/services/redirect"
	statusService "github.com/Llambi/cortito/internal/core/services/status"
	urlService "github.com/Llambi/cortito/internal/core/services/url"
	userService "github.com/Llambi/cortito/internal/core/services/user"
	"github.com/Llambi/cortito/internal/routing"
	"github.com/Llambi/cortito/platform/routers"
	"github.com/scalalang2/golang-fifo/sieve"
)

func main() {
	urlStore := sieve.New[string, string](30, 30*time.Minute)
	userStore := sieve.New[string, string](30, 30*time.Minute)
	urlRepo := urlMemory.MemoryRepository{Store: urlStore}
	userRepo := userMemory.MemoryRepository{Store: userStore}

	urlService := urlService.Service{Repo: urlRepo}
	userService := userService.Service{Repo: userRepo}
	urlHandler := url.Handler{UrlService: urlService, UserService: userService}
	urlRouting := routing.UrlRouting(urlHandler)

	statusService := statusService.Service{}
	statusHandler := status.Handler{StatusService: statusService}
	statusRouting := routing.StatusRouting(statusHandler)

	redirectService := redirectService.Service{Repo: urlRepo}
	redirectHandler := redirect.Handler{RedirectService: redirectService}
	redirectRouting := routing.RedirectRouting(redirectHandler)

	htmlHandler := html.Handler{}
	htmlRouting := routing.HtmlRouting(htmlHandler)

	var routes []routers.Route
	routes = append(routes, statusRouting...)
	routes = append(routes, urlRouting...)
	routes = append(routes, redirectRouting...)
	routes = append(routes, htmlRouting...)

	app := api.NewApp(routers.NewRouter(routes))

	log.Fatalln(app.Run())
}
