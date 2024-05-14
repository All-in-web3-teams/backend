package redis

import (
	"context"
	"time"
)

func SaveAddressAndNonce(address string, nonce string) (err error) {
	err = rdb.Set(context.Background(), address, nonce, time.Minute).Err()
	return
}

func GetNonceByAddress(address string) (nonce string, err error) {
	nonce, err = rdb.Get(context.Background(), address).Result()
	if err != nil {
		return "", err
	} else {
		return
	}

}
