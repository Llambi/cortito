package url

import (
	"github.com/Llambi/cortito/internal/core/domain"
	"github.com/gin-gonic/gin"
)

func (h Handler) GetByKey(ctx *gin.Context) {
	var cortitoUrl domain.FindCortitoUrlRequest
	if err := ctx.ShouldBindQuery(&cortitoUrl); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
	}

	cortitoUrlResponse, _ := h.UrlService.GetByKey(cortitoUrl.ShortUrl) //TODO: Fix this to use a DTO

	ctx.JSON(200, cortitoUrlResponse)
}
