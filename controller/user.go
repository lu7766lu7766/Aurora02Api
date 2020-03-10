package controller

import (
	"aurora02api/service"
	"github.com/gin-gonic/gin"
	"net/http"
)


type UserController struct {}

func (this UserController) GetList(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.UserService{}.GetList()) 
}