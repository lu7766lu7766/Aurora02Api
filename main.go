package main

import (
	"aurora02api/controller"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"time"
	// "fmt"
)

var (
	host = "0.0.0.0:8099"
)

func main() {

	router := gin.Default()
	router.Use(cors.New(cors.Config{
            AllowOrigins:  []string{"*"},
            AllowMethods:  []string{"PUT", "PATCH", "GET", "POST", "DELETE"},
            AllowHeaders:  []string{"Origin"},
						ExposeHeaders:    []string{"Content-Length"},
						AllowCredentials: true,
						MaxAge: 90 * time.Hour,
        }))
	router.POST("/sysLookout/ajaxCallStatusContent2", controller.SysLookoutController{}.AjaxCallStatusContent)
	router.GET("/downloadFile/recordFile", controller.DownloadController{}.RecordFile)
	router.GET("/downloadFile/recordFilesToZip", controller.DownloadController{}.RecordFilesToZip)
	router.GET("/test", controller.UserController{}.GetList)
	// fmt.Println(controller.UserController{}.GetList())
	// if err := fasthttp.ListenAndServe(host, router.Handler); err != nil {
	// 	fmt.Println("start fasthttp fail:", err.Error())
	// }

	// log.Println("service running at " + host)
	// log.Fatal(fasthttp.ListenAndServe(host, router.Handler))
	router.Run(host)
	// close gin logger
	// gin.SetMode(gin.ReleaseMode)
	// gin.DefaultWriter = ioutil.Discard
}
