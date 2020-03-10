package controller

import (
	"aurora02api/service"
	"github.com/gin-gonic/gin"
	"net/http"
	// "fmt"
)

type SysLookoutController struct {}

func (this SysLookoutController) AjaxCallStatusContent(ctx *gin.Context) {
	res := service.SysLookoutService{}.GetCallStatusContent(ctx.DefaultPostForm("userId", ""))
	ctx.JSON(http.StatusOK, res)
}