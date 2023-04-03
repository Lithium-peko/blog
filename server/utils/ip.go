package utils

import "github.com/gin-gonic/gin"

var IP = new(ipUtil)

type ipUtil struct{}

// 获取用户发送请求的 IP 地址
// 如果服务器不经过代理, 直接把自己的 IP 暴露出去, c.Request.RemoteAddr 就可以直接获取 IP
// 目前流行的架构中, 请求经过服务器前基本会经过代理 (Nginx 最常见), 此时直接获取 IP 拿到的是代理服务器的 IP
func (*ipUtil) GetIpAddress(c *gin.Context) (ipAddress string) {
	// c.ClientIP() 获取的是代理服务器的 IP (Nginx)

	// X-Real-IP: Nginx 服务器代理, 本项目明确使用 Nginx 作代理, 因此优先获取这个
	ipAddress = c.Request.Header.Get("X-Real-IP")
}
