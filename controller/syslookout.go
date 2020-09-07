package controller

import (
	"Aurora02Api/service"
	"net/http"

	"github.com/gin-gonic/gin"
	// "fmt"
)

type SysLookoutController struct{}

func (this SysLookoutController) AjaxCallStatusContent(ctx *gin.Context) {
	res := service.SysLookoutService{}.GetCallStatusContent(ctx.DefaultPostForm("userId", ""))
	ctx.JSON(http.StatusOK, res)
}
