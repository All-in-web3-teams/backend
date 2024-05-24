package models

//请求参数

type LoginParam struct {
	Signature string `json:"signature" binding:"required"`
	Address   string `json:"address" binding:"required"`
}

type DeployParam struct {
	//Address     string `json:"address" binding:"required"`
	TxHash      string `json:"txHash" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Symbol      string `json:"symbol" binding:"required"`
	TotalSupply string `json:"totalSupply" binding:"required"`
	Decimals    int64  `json:"decimals" binding:"required"`
}

type ContractInfoParam struct {
	ContractAddress string `json:"contractAddress" binding:"required"`
	Homepage        string `json:"homepage"`
	XUrl            string `json:"xUrl"`
	Discord         string `json:"discord"`
	Telegram        string `json:"telegram"`
}
