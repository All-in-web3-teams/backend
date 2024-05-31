package logic

import (
	"backend/controllers/vrfV2"
	"backend/settings"
	"context"
	"crypto/ecdsa"
	"errors"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
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
	address := common.HexToAddress("0x8103B0A8A00be2DDC778e6e7eaa21791Cd364625")
	instance, err := vrfV2.NewVrfV2(address, client)
	if err != nil {
		return err
	}

	var subId uint64 = 10486
	tx, err := instance.AddConsumer(auth, subId, common.HexToAddress(raffleAddress))
	if err != nil {
		return err
	}
	println(tx.Hash().Hex())
	return nil
}
