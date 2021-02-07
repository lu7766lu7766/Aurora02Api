package router

import (
	"Aurora02Api/controller/DownloadController"
	"Aurora02Api/controller/SysLookoutController"
	"Aurora02Api/controller/UserController"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	// "fmt" //
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           90 * time.Hour,
	}))
	router.LoadHTMLGlob("view/*")
	router.GET("/user/list", UserController.GetList)
	router.POST("/sysLookout/ajaxCallStatusContent2", SysLookoutController.AjaxCallStatusContent)
	router.GET("/downloadFile/recordFile", DownloadController.RecordFile)
	router.GET("/downloadFile/recordFilesToZip", DownloadController.RecordFilesToZip)
	return router
}
