package redirect

import (
	"github.com/gin-gonic/gin"
)

func (h Handler) Redirect(ctx *gin.Context) {
	shortUrl := ctx.Param("shortUrl")
	cortitoUrl, _ := h.RedirectService.Redirect(shortUrl) //TODO: Error
	ctx.Redirect(301, cortitoUrl.Url)
}
