package logic

import (
	"backend/dao/mysql"
	"backend/models"
	"backend/settings"
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func CheckDeploySuccess(txHash string) (string, string, error) {
	client, err := ethclient.Dial(settings.Conf.EthClientConfig.Url)
	if err != nil {
		return "", "", err
	}

	// 获取交易收据
	receipt, err := client.TransactionReceipt(context.Background(), common.HexToHash(txHash))
	//tx, isPending, err := client.TransactionByHash(context.Background(), common.HexToHash(txHash))
	if err != nil {
		return "", "", err
	}

	// 判断合约部署状态
	if receipt == nil {
		return "Pending", "", nil
	} else if receipt.Status == 1 {
		return "Success", receipt.ContractAddress.Hex(), nil
	} else {
		return "Failed", "", nil
	}
}

func SaveContractInfo(p *models.DeployParam, address string, contractAddress string) (err error) {
	err = mysql.InsertContractInfo(p, address, contractAddress)
	return
}
