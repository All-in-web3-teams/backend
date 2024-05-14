package middlewares

import (
	"backend/controllers"
	"backend/pkg/jwt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

// JWTAuthMiddleware 基于JWT的认证中间件
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 1.获取cookie中的token
		cookie, err := c.Request.Cookie("token")
		if err != nil {
			zap.L().Error("JWTAuthMiddleware c.Request.Cookie failed", zap.Error(err))
			c.JSON(http.StatusUnauthorized, gin.H{
				"msg":  "需要身份验证",
				"data": false,
			})
			c.Abort()
			return
		}
		tokenValue := cookie.Value

		// 2.解析JWT
		mc, err := jwt.ParseToken(tokenValue)
		if err != nil {
			zap.L().Error("JWTAuthMiddleware jwt.ParseToken failed", zap.Error(err))
			c.JSON(http.StatusUnauthorized, gin.H{
				"msg":  "需要身份验证",
				"data": false,
			})
			c.Abort()
			return
		}
		// 3.将当前请求的username信息保存到请求的上下文c上
		c.Set(controllers.CtxAddressKey, mc.Address)
		c.Next() // 后续的处理函数可以用过c.Get(CtxAddressKey)来获取当前请求的用户信息
	}
}
