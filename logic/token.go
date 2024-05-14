package logic

import "backend/dao/mysql"

func GetContractByAddress(address string) (contractAddresses []string, err error) {
	contractAddresses, err = mysql.GetContractByAddress(address)
	return
}
