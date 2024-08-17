package routers

import "github.com/gin-gonic/gin"

type Router interface {
	Init(ginEngine *gin.Engine)
}

type Route struct {
	Method string
	Path   string
	Handle gin.HandlerFunc
}

type routing struct {
	Routers []Route
}

func NewRouter(routes []Route) Router {
	return &routing{
		Routers: routes,
	}
}

func (r *routing) Init(ginEngine *gin.Engine) {
	ginEngine.
		Use(gin.Logger()).
		Use(gin.Recovery()).
		Use(CORSHandler)
	ginEngine.Static("/assets", "./assets")
	ginEngine.LoadHTMLGlob("templates/*.html")

	for _, router := range r.Routers {
		ginEngine.Handle(router.Method, router.Path, router.Handle)
	}
}

func CORSHandler(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Next()
}
