package front

import (
	"blog/utils/r"

	"github.com/gin-gonic/gin"
)

type BlogInfo struct{}

// 获取前台信息
func (*BlogInfo) GetFrontHomeInfo(c *gin.Context) {
	r.SuccessData(c, blogInfoService.GetFrontHomeInfo())
}
