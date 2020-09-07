package tools

import (
	"os"
	"path"

	"github.com/gin-gonic/gin"
)

func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func Attachment(ctx *gin.Context, filePath string) {
	if _, err := os.Stat(filePath); err == nil {
		ctx.Header("Content-Description", "File Transfer")
		ctx.Header("Content-Transfer-Encoding", "binary")
		ctx.Header("Content-Disposition", "attachment; filename="+path.Base(filePath))
		ctx.Header("Content-Type", "application/octet-stream")
		ctx.File(filePath)
	}
}
