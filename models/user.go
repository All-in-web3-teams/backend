package models

import "gorm.io/gorm"

//数据库

type User struct {
	gorm.Model
	Address         string
	ContractAddress string
	Name            string
	Symbol          string
	TotalSupply     int64
	Decimals        int64
}

func (u User) TableName() string {
	//绑定MYSQL表名为users
	return "user"
}

// TokenAddressAndName 合约地址和名称
type TokenAddressAndName struct {
	ContractAddress string `json:"contractAddress"`
	Name            string `json:"name"`
}
