package url

import (
	"github.com/Llambi/cortito/internal/core/domain"
	"github.com/gin-gonic/gin"
)

func (h Handler) Create(ctx *gin.Context) {
	var createCortitoUrlRequest domain.CreateCortitoUrlRequest
	createCortitoUserRequest := domain.CreateCortitoUserRequest{Ip: ctx.ClientIP()}
	if err := ctx.BindJSON(&createCortitoUrlRequest); err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
		return
	}

	cortitoUserResponse, err := h.UserService.Upsert(createCortitoUserRequest)
	if err != nil {
		ctx.AbortWithStatusJSON(429, gin.H{"error": err.Error()})
		return
	}
	cortitoUrlResponse, _ := h.UrlService.Create(createCortitoUrlRequest) //TODO: Error

	ctx.JSON(200, gin.H{"data": cortitoUrlResponse, "user": cortitoUserResponse})
}
