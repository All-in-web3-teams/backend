package controllers

import (
	"backend/logic"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

func GetTokenListHandler(c *gin.Context) {
	address, err := getCurrentUser(c)
	if err != nil {
		zap.L().Error("GetTokenListHandler getCurrentUser failed", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":  "服务端错误",
			"data": "解析address失败",
		})
		return
	}
	contractAddresses, err := logic.GetContractByAddress(address)
	if err != nil {
		zap.L().Error("GetTokenListHandler logic.GetContractByAddress failed", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":  "服务端错误",
			"data": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": "success",
		"data": gin.H{
			"address":         address,
			"contractAddress": contractAddresses,
		},
	})
}

func PostTokenInfoHandler(c *gin.Context) {}
