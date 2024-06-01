package mysql

import (
	"backend/models"
	"errors"
	"gorm.io/gorm"
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

func InsertContractSocialMediaInfo(p *models.ContractInfoParam) (err error) {
	var contractInfo = models.ContractInfo{
		ContractAddress: p.ContractAddress,
		Homepage:        p.Homepage,
		XUrl:            p.XUrl,
		Discord:         p.Discord,
		Telegram:        p.Telegram,
	}
	var exists models.ContractInfo
	if err = db.Table(models.ContractInfo{}.TableName()).Where("contract_address = ?", p.ContractAddress).First(&exists).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) { // 如果记录不存在，插入新记录
			return db.Create(&contractInfo).Error
		} else {
			return err // 返回其他错误
		}
	}
	var fieldsToUpdate = map[string]interface{}{
		"ContractAddress": p.ContractAddress,
		"Homepage":        p.Homepage,
		"XUrl":            p.XUrl,
		"Discord":         p.Discord,
		"Telegram":        p.Telegram,
	}
	err = db.Model(&exists).Updates(&fieldsToUpdate).Error
	return
}

func GetContractInfoByAddress(contractAddress string) (contractInfo models.ContractInfo, err error) {
	err = db.Table(models.ContractInfo{}.TableName()).Where("contract_address = ?", contractAddress).First(&contractInfo).Error
	return
}

func SaveRaffleAddress(contractAddress string, raffleAddress string) (err error) {
	var contractInfo = models.ContractInfo{
		ContractAddress: contractAddress,
		Homepage:        "",
		XUrl:            "",
		Discord:         "",
		Telegram:        "",
		RaffleAddress:   raffleAddress,
	}
	var exists models.ContractInfo
	if err = db.Table(models.ContractInfo{}.TableName()).Where("contract_address = ?", contractAddress).First(&exists).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) { // 如果记录不存在，插入新记录
			return db.Create(&contractInfo).Error
		} else {
			return err // 返回其他错误
		}
	}
	err = db.Table(models.ContractInfo{}.TableName()).Where("contract_address = ?", contractAddress).Update("raffle_address", raffleAddress).Error
	return
}
