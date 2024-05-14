package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
)

const CtxAddressKey = "address"

var ErrorNotLogin = errors.New("地址未登陆")

// GetCurrentUser 获取token中的address信息
func getCurrentUser(c *gin.Context) (address string, err error) {

	addressAny, ok := c.Get(CtxAddressKey)
	if !ok {
		err = ErrorNotLogin
		return
	}
	address, ok = addressAny.(string)
	if !ok {
		err = ErrorNotLogin
		return
	}
	return
}
