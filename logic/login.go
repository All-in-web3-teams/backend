package logic

import (
	"backend/dao/redis"
	"backend/models"
	"backend/pkg/jwt"
	"backend/pkg/nonce"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"strings"
)

var ErrorInvalidSig = errors.New("无效的签名")

const NonceDigit = 10

func GetAndSaveNonce(address string) (nonceStr string, err error) {
	nonceStr = nonce.GetNonce(NonceDigit)
	err = redis.SaveAddressAndNonce(address, nonceStr)
	if err != nil {
		return "", err
	}

	return nonceStr, nil
}

func Login(p *models.LoginParam) (token string, err error) {
	var nonce string
	nonce, err = redis.GetNonceByAddress(p.Address)
	if err != nil {
		return "", err
	}
	//nonce = "309766"
	message := []byte(nonce)
	prefix := fmt.Sprintf("\x19Ethereum Signed Message:\n%d", len(message))
	data := []byte(prefix + string(message))
	hash := crypto.Keccak256Hash(data)
	signatureBytes, err := hexutil.Decode(p.Signature)
	if err != nil {
		return "", err
	}
	if signatureBytes[64] >= 27 {
		signatureBytes[64] -= 27 // 对于旧版以太坊签名需要做的调整
	}

	sigPublicKeyECDSA, err := crypto.SigToPub(hash.Bytes(), signatureBytes)
	if err != nil {
		return "", err
	}
	address := crypto.PubkeyToAddress(*sigPublicKeyECDSA).Hex()
	address = strings.ToLower(address)
	//println(address)
	if address == p.Address {
		token, err = jwt.GenToken(p.Address)
		return
	} else {
		return "", ErrorInvalidSig
	}
}
