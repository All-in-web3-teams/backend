package logic

import (
	"backend/controllers/automation"
	"backend/controllers/automationRegistrar2_1"
	"backend/controllers/vrfV2_5"
	"backend/settings"
	"context"
	"crypto/ecdsa"
	"errors"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"time"
)

func AddConsumer(raffleAddress string) error {
	client, err := ethclient.Dial(settings.Conf.EthClientConfig.Url)
	if err != nil {
		return err
	}
	privateKey, err := crypto.HexToECDSA(settings.Conf.PrivateKey)
	if err != nil {
		return err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return errors.New("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return err
	}
	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(500000) // in units
	auth.GasPrice = gasPrice
	address := common.HexToAddress("0x9DdfaCa8183c41ad55329BdeeD9F6A8d53168B1B")
	instance, err := vrfV2_5.NewVrfV25(address, client)
	if err != nil {
		return err
	}
	numStr := "19709626778135997837660097477283008945859889359416389923405312488169579505131"
	subId, _ := new(big.Int).SetString(numStr, 10)
	tx, err := instance.AddConsumer(auth, subId, common.HexToAddress(raffleAddress))
	if err != nil {
		return err
	}
	println(tx.Hash().Hex())
	return nil
}
func AddAutomation(raffleAddress string) error {
	client, err := ethclient.Dial(settings.Conf.EthClientConfig.Url)
	if err != nil {
		return err
	}
	privateKey, err := crypto.HexToECDSA(settings.Conf.PrivateKey)
	if err != nil {
		return err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return errors.New("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return err
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return err
	}
	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(500000) // in units
	auth.GasPrice = gasPrice

	addressLink := common.HexToAddress("0x779877A7B0D9E8603169DdbD7836e478b4624789")
	instanceLink, err := automation.NewAutomation(addressLink, client)
	txLink, err := instanceLink.Approve(auth, common.HexToAddress("0xb0E49c5D0d05cbc241d68c05BC5BA1d1B7B72976"), big.NewInt(5000000000000000000))
	if err != nil {
		return err
	}
	println(txLink.Hash().Hex())

	time.Sleep(60 * time.Second)

	nonce, err = client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return err
	}

	gasPrice, err = client.SuggestGasPrice(context.Background())
	if err != nil {
		return err
	}
	auth = bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(500000) // in units
	auth.GasPrice = gasPrice

	addressRegistrar := common.HexToAddress("0xb0E49c5D0d05cbc241d68c05BC5BA1d1B7B72976")
	instanceRegister, err := automationRegistrar2_1.NewAutomationRegistrar21(addressRegistrar, client)
	if err != nil {
		return err
	}
	now := time.Now()
	// 自定义的时间格式布局
	layout := "20060102-15-04"
	// 根据布局格式化时间
	formattedTime := now.Format(layout)
	requstParams := automationRegistrar2_1.AutomationRegistrar21RegistrationParams{
		Name:           "raffle" + formattedTime,
		EncryptedEmail: []byte{},
		UpkeepContract: common.HexToAddress(raffleAddress),
		GasLimit:       uint32(500000),
		AdminAddress:   common.HexToAddress("0x85dd2e97Db4416159B128f833C99A9134b939475"),
		TriggerType:    uint8(0),
		CheckData:      common.Hex2Bytes("0x"),
		TriggerConfig:  common.Hex2Bytes("0x"),
		OffchainConfig: common.Hex2Bytes("0x"),
		Amount:         big.NewInt(5000000000000000000),
	}
	txRegistrar, err := instanceRegister.RegisterUpkeep(auth, requstParams)
	if err != nil {
		return err
	}
	println(txRegistrar.Hash().Hex())
	return nil
}
