package controllers

import (
	"backend/logic"
	"backend/models"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

func CheckUserLoginHandler(c *gin.Context) {
	address, err := getCurrentUser(c)
	if err != nil {
		zap.L().Error("CheckUserLoginHandler getCurrentUser failed", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":  "服务端错误",
			"data": "解析address失败",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  "success", // 成功信息
		"data": address,   // 已登录
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
