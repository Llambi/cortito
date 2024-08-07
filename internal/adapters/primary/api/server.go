package api

import (
	"fmt"

	"github.com/Llambi/cortito/platform/routers"
	"github.com/gin-gonic/gin"
)


type App struct{
	ginEngine *gin.Engine
	router routers.Router
	port int
}

func NewApp(router routers.Router) *App {
	app := &App {
		ginEngine: gin.Default(),
		router: router,
		port: 8000,
	}
	
	router.Init(app.ginEngine)

	return app
}

func (app *App) Run() error {
	return app.ginEngine.Run(fmt.Sprintf(":%d",app.port))
}