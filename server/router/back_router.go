package router

import (
	"blog/config"
	"blog/router/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 后台管理页面的接口路由
func BackRouter() http.Handler {
	gin.SetMode(config.Cfg.Server.AppMode)

	r := gin.New()
	r.SetTrustedProxies([]string{"*"})

	// 使用本地文件上传, 需要静态文件服务, 使用七牛云不需要
	if config.Cfg.Upload.OssType == "local" {
		r.Static("/public", "./public")
		r.StaticFS("/dir", http.Dir("./public")) // 将 public 目录内的文件列举展示
	}

	r.Use(middleware.Logger())             // 自定义的 zap 日志中间件
	r.Use(middleware.ErrorRecovery(false)) // 自定义错误处理中间件
	r.Use(middleware.Cors())               // 跨域中间件
}
