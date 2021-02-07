package UserController

import (
	"Aurora02Api/service/UserService"
	"Aurora02Api/tools/ReturnMessage"

	"github.com/gin-gonic/gin"
)

func GetList(ctx *gin.Context) {
	ReturnMessage.Success(ctx, UserService.GetList())
}
