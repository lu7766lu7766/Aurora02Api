package UserController

import (
	"Aurora02Api/service/UserService"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetList(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, UserService.GetList())
}
