package controller

import (
	"aurora02api/service"
	"github.com/gin-gonic/gin"
	"os"
	"path"
)


type DownloadController struct {}

func (this DownloadController) RecordFile(ctx *gin.Context) {
	filePath := service.DownloadService{}.GetRecordFile(
		ctx.DefaultQuery("userId", ""),
		ctx.DefaultQuery("connectDate", ""),
		ctx.DefaultQuery("fileName", ""))

	this.Attachment(ctx, filePath)
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

	this.Attachment(ctx, filePath)
	// tools.DownloadFile(ctx, UserID + "RecordFile.zip", filePath)
	// ctx.JSON(http.StatusOK, res)
}

func (this DownloadController) Attachment(ctx *gin.Context, filePath string) {
	if _, err := os.Stat(filePath); err == nil {
		ctx.Header("Content-Description", "File Transfer")
		ctx.Header("Content-Transfer-Encoding", "binary")
		ctx.Header("Content-Disposition", "attachment; filename="+path.Base(filePath))
		ctx.Header("Content-Type", "application/octet-stream")
		ctx.File(filePath)
	}
}
