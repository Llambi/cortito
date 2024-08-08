package status

import "github.com/gin-gonic/gin"

func (h Handler) Status(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"msg": "I am UP!"})
}
