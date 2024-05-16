package logic

import (
	"backend/dao/mysql"
	"backend/models"
)

func GetContractByAddress(address string) (contractAddresses []string, err error) {
	contractAddresses, err = mysql.GetContractByAddress(address)
	return
}

func GetAllTokenAddressAndName() (results []models.TokenAddressAndName, err error) {
	results, err = mysql.GetAllTokenAddressAndName()
	if err == nil {
		results = append(results, models.TokenAddressAndName{
			ContractAddress: "0x16EFdA168bDe70E05CA6D349A690749d622F95e0",
			Name:            "WBTC",
		})
		results = append(results, models.TokenAddressAndName{
			ContractAddress: "0x6F6bB5dADDB05718382A0192B65603492C939f8F",
			Name:            "USDC",
		})
		results = append(results, models.TokenAddressAndName{
			ContractAddress: "0xe83EbF27702244dD5997526692f03415f741b98F",
			Name:            "DAI",
		})
		results = append(results, models.TokenAddressAndName{
			ContractAddress: "0x271B34781c76fB06bfc54eD9cfE7c817d89f7759",
			Name:            "USDT",
		})
		results = append(results, models.TokenAddressAndName{
			ContractAddress: "0x097D90c9d3E0B50Ca60e1ae45F6A81010f9FB534",
			Name:            "WETH",
		})
	}
	return
}
