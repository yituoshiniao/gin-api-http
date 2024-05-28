package v1

import (
	"fmt"
	"net"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yituoshiniao/kit/xlog"
	"github.com/yituoshiniao/kit/xtrace"
)

// IPWhiteListMiddleware 是一个中间件函数，用于实现IP白名单功能 //[]string{"127.0.0.1", "::1"} 运行本地
func IPWhiteListMiddleware(whitelist []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		clientIP := c.ClientIP()
		// msg := "Forbidden: Your IP is not allowed. " + "[clientIP:]" + clientIP
		msg := fmt.Sprintf("Forbidden: Your IP is not allowed. [clientIP:] %s", clientIP)
		// 检查客户端IP是否在白名单中
		if !isIPAllowed(clientIP, whitelist) {
			xlog.S(c.Request.Context()).Warnw("clientIP-信息", "clientIP", clientIP)
			c.JSON(http.StatusForbidden, gin.H{
				"code":    1,
				"msg":     msg,
				"data":    whitelist,
				"traceId": xtrace.TraceIdFromContext(c.Request.Context()),
			})
			c.Abort()
			return
		}

		// 继续处理下一个中间件或路由处理器
		c.Next()
	}
}

// isIPAllowed 检查IP是否在白名单中
func isIPAllowed(ip string, whitelist []string) bool {
	clientAddr := net.ParseIP(ip)
	for _, allowedIP := range whitelist {
		allowedAddr := net.ParseIP(allowedIP)
		if clientAddr.Equal(allowedAddr) {
			return true
		}
	}
	return false
}
