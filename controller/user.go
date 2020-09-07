package controller

import (
	"Aurora02Api/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (this UserController) GetList(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.UserService{}.GetList())
}
