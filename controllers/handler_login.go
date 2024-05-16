package controllers

import (
	"backend/logic"
	"backend/models"
	"backend/pkg/jwt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

func CheckUserLoginHandler(c *gin.Context) {
	// 1.获取cookie中的token
	cookie, err := c.Request.Cookie("token")
	if err != nil {
		zap.L().Error("CheckUserLoginHandler c.Request.Cookie failed", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"msg":  "未登陆",
			"data": nil,
		})
		return
	}
	tokenValue := cookie.Value

	// 2.解析JWT
	mc, err := jwt.ParseToken(tokenValue)
	if err != nil {
		zap.L().Error("CheckUserLoginHandler jwt.ParseToken failed", zap.Error(err))
		c.JSON(http.StatusOK, gin.H{
			"msg":  "未登陆",
			"data": nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  "success",
		"data": mc.Address,
	})
}

// GetNonceHandler 获取nonce
func GetNonceHandler(c *gin.Context) {
	//1.参数校验
	address, ok := c.GetQuery("address")
	if !ok {
		zap.L().Error("GetNonceHandler with invalid param")
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":  "请求参数错误",
			"data": false,
		})
		return
	}

	if !AddressCheck(address) {
		zap.L().Error("GetNonceHandler with invalid param")
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":  "请求参数错误",
			"data": false,
		})
		return
	}

	//2.生成nonce并存储
	nonce, err := logic.GetAndSaveNonce(address)
	if err != nil {
		zap.L().Error("GetNonceHandler logic.GetAndSaveNonce failed", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":  "服务端错误",
			"data": err.Error(),
		})
		return
	}

	//3.返回响应
	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
		"data": gin.H{
			"address": address,
			"nonce":   nonce,
		},
	})
}

// LoginHandler 使用签名登陆
func LoginHandler(c *gin.Context) {
	var (
		err   error
		token string
	)
	//1.参数校验
	p := new(models.LoginParam)
	if err = c.ShouldBindJSON(p); err != nil {
		zap.L().Error("LoginHandler c.ShouldBindJSON failed", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":  "请求参数错误",
			"data": false,
		})
		return
	}
	if !AddressCheck(p.Address) || !SigCheck(p.Signature) {
		zap.L().Error("LoginHandler with invalid param")
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":  "请求参数错误",
			"data": false,
		})
		return
	}

	//2.验证签名并返回token
	token, err = logic.Login(p)
	if err != nil {
		zap.L().Error("LoginHandler logic.Login failed", zap.Error(err))
		c.JSON(http.StatusUnauthorized, gin.H{
			"msg":  "需要身份验证",
			"data": false,
		})
		return
	}

	//3.返回响应
	//将token存入cookie中
	cookie := &http.Cookie{
		Name:     "token",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
	}
	http.SetCookie(c.Writer, cookie)

	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
		"data": gin.H{
			"address": p.Address,
		},
	})
}
