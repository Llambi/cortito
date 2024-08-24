package html

import "github.com/gin-gonic/gin"

func (h Handler) Home(ctx *gin.Context) {
	ctx.HTML(200, "index.html", gin.H{
		"content": "This is an index page...",
	})
}
