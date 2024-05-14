package mysql

import (
	"backend/models"
)

func InsertContractInfo(p *models.DeployParam, address string, contractAddress string) (err error) {
	user := models.User{
		Address:         address,
		ContractAddress: contractAddress,
		Name:            p.Name,
		Symbol:          p.Symbol,
		TotalSupply:     p.TotalSupply,
		Decimals:        p.Decimals,
	}

	err = db.Create(&user).Error

	return
}

func GetContractByAddress(address string) (contractAddresses []string, err error) {
	err = db.Model(&models.User{}).Select("contract_address").Where("address = ?", address).Find(&contractAddresses).Error
	return
}
