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

type ContractInfo struct {
	gorm.Model
	ContractAddress string
	Homepage        string
	XUrl            string
	Discord         string
	Telegram        string
	RaffleAddress   string
}

func (c ContractInfo) TableName() string {
	//绑定MYSQL表名为users
	return "contract_info"
}

// TokenAddressAndName 合约地址和名称
type TokenAddressAndName struct {
	ContractAddress string `json:"contractAddress"`
	Name            string `json:"name"`
}
