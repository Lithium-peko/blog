package v1

import (
	"blog/utils"
	"blog/utils/r"
	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

type Upload struct{}

func (*Upload) UploadFile(c *gin.Context) {
	_, fileHeader, err := c.Request.FormFile("file")
	// 处理文件错误
	if err != nil {
		code := r.ERROR_FILE_RECEIVE
		utils.Logger.Error(r.GetMsg(code), zap.Error(err))
		r.SendCode(c, code)
		return
	}
	// 上传文件, 获取文件路径
	url, code := fileService.UploadFile(fileHeader)
	r.SendData(c, code, url)
}
