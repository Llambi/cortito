package url

import (
	"github.com/Llambi/cortito/internal/core/domain"
	"github.com/gin-gonic/gin"
)

func (h Handler) Create(ctx *gin.Context) {
	var cortitoUrl domain.CreateCortitoUrlRequest
	if err := ctx.BindJSON(&cortitoUrl); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
	}

	cortitoUrlResponse, _ := h.UrlService.Create(cortitoUrl) //TODO: Error

	ctx.JSON(200, cortitoUrlResponse)
}
