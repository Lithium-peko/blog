package v1

import (
	"blog/model/req"
	"blog/utils"
	"blog/utils/r"

	"github.com/gin-gonic/gin"
)

type Page struct{}

func (*Page) GetList(c *gin.Context) {
	r.SuccessData(c, pageService.GetList())
}

func (*Page) SaveOrUpdate(c *gin.Context) {
	r.SendCode(c, pageService.SaveOrUpdate(utils.BindJson[req.AddOrEditPage](c)))
}

func (*Page) Delete(c *gin.Context) {
	r.SendCode(c, pageService.Delete(utils.BindJson[[]int](c)))
}
