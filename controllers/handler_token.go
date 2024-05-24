package controllers

import (
	"backend/logic"
	"backend/models"
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
func GetAllTokenAddressAndNameHandler(c *gin.Context) {
	tokenAddressAndName, err := logic.GetAllTokenAddressAndName()
	if err != nil {
		zap.L().Error("GetAllTokenAddressAndNameHandler logic.GetAllTokenAddressAndName failed", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":  "服务端错误",
			"data": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  "success",
		"data": tokenAddressAndName,
	})
}

func PostTokenInfoHandler(c *gin.Context) {
	p := new(models.ContractInfoParam)
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("PostTokenInfoHandler c.ShouldBindJSON failed", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":  "请求参数错误",
			"data": false,
		})
		return
	}

	err := logic.SaveContractSocialMediaInfo(p)
	if err != nil {
		zap.L().Error("PostTokenInfoHandler logic.SaveContractInfo failed", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":  "服务端错误",
			"data": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  "success",
		"data": true,
	})
}

func GetTokenInfoHandler(c *gin.Context) {
	contractAddress, ok := c.GetQuery("contractAddress")
	if !ok {
		zap.L().Error("GetNonceHandler with invalid param")
		c.JSON(http.StatusBadRequest, gin.H{
			"msg":  "请求参数错误",
			"data": false,
		})
		return
	}

	contractInfo, err := logic.GetContractInfo(contractAddress)
	if err != nil {
		zap.L().Error("GetTokenInfoHandler logic.GetContractInfo failed", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg":  "服务端错误",
			"data": err.Error(),
		})
		return
	}
	type socialMediaRes struct {
		ContractAddress string `json:"contractAddress"`
		Homepage        string `json:"homepage"`
		XUrl            string `json:"xUrl"`
		Discord         string `json:"discord"`
		Telegram        string `json:"telegram"`
	}
	socialMedia := socialMediaRes{
		ContractAddress: contractInfo.ContractAddress,
		Homepage:        contractInfo.Homepage,
		XUrl:            contractInfo.XUrl,
		Discord:         contractInfo.Discord,
		Telegram:        contractInfo.Telegram,
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":  "success",
		"data": socialMedia,
	})
}
