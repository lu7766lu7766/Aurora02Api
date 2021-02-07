package SysLookoutController

import (
	"Aurora02Api/service/SysLookoutService"
	"net/http"

	"github.com/gin-gonic/gin"
	// "fmt"
)

func AjaxCallStatusContent(ctx *gin.Context) {
	res := SysLookoutService.GetCallStatusContent(ctx.DefaultPostForm("userId", ""))
	ctx.JSON(http.StatusOK, res)
}
