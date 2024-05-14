package controllers

import (
	"backend/logic"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Test(c *gin.Context) {
	address := "0xssssssssssssssssssssssssssssssssssssss"
	contractAddresses, err := logic.GetContractByAddress(address)
	if err != nil {
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
