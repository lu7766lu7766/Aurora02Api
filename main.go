package main

import (
	"Aurora02Api/controller"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	// "fmt" //
)

var (
	host = "0.0.0.0:8099"
)

func main() {

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
	router.POST("/sysLookout/ajaxCallStatusContent2", controller.SysLookoutController{}.AjaxCallStatusContent)
	router.GET("/downloadFile/recordFile", controller.DownloadController{}.RecordFile)
	router.GET("/downloadFile/recordFilesToZip", controller.DownloadController{}.RecordFilesToZip)
	router.Run(host)
}
