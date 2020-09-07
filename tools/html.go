package tools

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Alert(ctx *gin.Context, message string) {
	ctx.HTML(http.StatusOK, "alert.html", gin.H{
		"message": message,
	})
}
