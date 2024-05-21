package mysql

import (
	"backend/models"
	"math/big"
)

func InsertContractInfo(p *models.DeployParam, address string, contractAddress string) (err error) {
	bigNum := new(big.Int)
	bigNum.SetString(p.TotalSupply, 10)                              // 将字符串转换为big.Int
	divisor := new(big.Int).Exp(big.NewInt(10), big.NewInt(18), nil) // 计算1e18的大整数表示

	totalSupply := new(big.Int)
	totalSupply.Div(bigNum, divisor) // 进行除法运算

	user := models.User{
		Address:         address,
		ContractAddress: contractAddress,
		Name:            p.Name,
		Symbol:          p.Symbol,
		TotalSupply:     totalSupply.Int64(),
		Decimals:        p.Decimals,
	}

	err = db.Table(models.User{}.TableName()).Create(&user).Error

	return
}

func GetContractByAddress(address string) (contractAddresses []string, err error) {
	err = db.Model(&models.User{}).Select("contract_address").Where("address = ?", address).Find(&contractAddresses).Error
	return
}

func GetAllTokenAddressAndName() (results []models.TokenAddressAndName, err error) {
	err = db.Table(models.User{}.TableName()).Select("contract_address, name").Scan(&results).Error
	return
}
