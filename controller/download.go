package controller

import (
	"Aurora02Api/service"
	"Aurora02Api/tools"

	"github.com/gin-gonic/gin"
)

type DownloadController struct{}

func (this DownloadController) RecordFile(ctx *gin.Context) {
	filePath := service.DownloadService{}.GetRecordFile(
		ctx.DefaultQuery("userId", ""),
		ctx.DefaultQuery("connectDate", ""),
		ctx.DefaultQuery("fileName", ""))
	if tools.FileExists(filePath) {
		tools.Attachment(ctx, filePath)
	} else {
		tools.Alert(ctx, "找不到檔案")
	}
}

func (this DownloadController) RecordFilesToZip(ctx *gin.Context) {

	filePath := service.DownloadService{}.GetRecordFilesToZip(
		ctx.DefaultQuery("choice", ""),
		ctx.DefaultQuery("callStartBillingDate", ""),
		ctx.DefaultQuery("callStopBillingDate", ""),
		ctx.DefaultQuery("extensionNo", ""),
		ctx.DefaultQuery("orgCalledID", ""),
		ctx.DefaultQuery("durationCondition", ""),
		ctx.DefaultQuery("callDuration", ""))

	tools.Attachment(ctx, filePath)
	// tools.DownloadFile(ctx, UserID + "RecordFile.zip", filePath)
	// ctx.JSON(http.StatusOK, res)
}
