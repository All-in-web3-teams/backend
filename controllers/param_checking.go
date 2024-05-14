package controllers

import (
	"regexp"
)

func AddressCheck(address string) bool {
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	return re.MatchString(address)
}

func SigCheck(sig string) bool {
	re := regexp.MustCompile("^0x[0-9a-fA-F]{130}$")
	return re.MatchString(sig)
}

func TxHashCheck(tx string) bool {
	re := regexp.MustCompile("^0x[0-9a-fA-F]{64}$")
	return re.MatchString(tx)
}
