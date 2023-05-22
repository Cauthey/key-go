package middleware

import (
	"github.com/gin-gonic/gin"
	"key-go/internal/app/config"
	"key-go/internal/app/contextx"
	"strings"
)

func AuthExpire(skippers ...SkipperFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		if SkipHandler(c, skippers...) {
			c.Next()
			return
		}
		uri := c.Request.URL.String()
		if uri == "/" || strings.HasPrefix(uri, "/#") {
			c.Next()
			return
		}
		ctx := c.Request.Context()
		userID := contextx.FromUserID(ctx)
		// 获取配置过期时间
		sessionExpire := config.C.HTTP.SessionExpire
		if sessionExpire == 0 {
			c.Next()
			return
		} else {
			if userID != 0 {

			}
		}
		c.Next()
		return
	}
}
