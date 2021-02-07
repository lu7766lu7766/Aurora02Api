package ReturnMessage

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Success(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": data,
	})
}

func Error(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": -1,
		"data": data,
	})
}
