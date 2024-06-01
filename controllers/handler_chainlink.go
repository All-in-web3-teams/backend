package controllers

import (
	"backend/logic"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

func AddConsumerHandler(c *gin.Context) {
	raffleAddress, ok := c.GetQuery("raffleAddress")
	if !ok {
		zap.L().Error("AddConsumerHandler with invalid param")
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":  "请求参数错误",
			"data": false,
		})
		return
	}

	if !AddressCheck(raffleAddress) {
		zap.L().Error("AddConsumerHandler with invalid param")
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":  "请求参数错误",
			"data": false,
		})
		return
	}

	err := logic.AddConsumer(raffleAddress)
	if err != nil {
		zap.L().Error("AddConsumerHandler logic.AddConsumer failed", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":  "服务端错误",
			"data": false,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":  "success",
		"data": true,
	})
}
func AddAutomationHandler(c *gin.Context) {
	raffleAddress, ok := c.GetQuery("raffleAddress")
	if !ok {
		zap.L().Error("AddConsumerHandler with invalid param")
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":  "请求参数错误",
			"data": false,
		})
		return
	}

	if !AddressCheck(raffleAddress) {
		zap.L().Error("AddConsumerHandler with invalid param")
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":  "请求参数错误",
			"data": false,
		})
		return
	}

	err := logic.AddAutomation(raffleAddress)
	if err != nil {
		zap.L().Error("AddConsumerHandler logic.AddConsumer failed", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":  "服务端错误",
			"data": false,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":  "success",
		"data": true,
	})
}
