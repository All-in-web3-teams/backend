// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package automationRegistrar2_1

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// AutomationRegistrar21InitialTriggerConfig is an auto generated low-level Go binding around an user-defined struct.
type AutomationRegistrar21InitialTriggerConfig struct {
	TriggerType           uint8
	AutoApproveType       uint8
	AutoApproveMaxAllowed uint32
}

// AutomationRegistrar21RegistrationParams is an auto generated low-level Go binding around an user-defined struct.
type AutomationRegistrar21RegistrationParams struct {
	Name           string
	EncryptedEmail []byte
	UpkeepContract common.Address
	GasLimit       uint32
	AdminAddress   common.Address
	TriggerType    uint8
	CheckData      []byte
	TriggerConfig  []byte
	OffchainConfig []byte
	Amount         *big.Int
}

// AutomationRegistrar21TriggerRegistrationStorage is an auto generated low-level Go binding around an user-defined struct.
type AutomationRegistrar21TriggerRegistrationStorage struct {
	AutoApproveType       uint8
	AutoApproveMaxAllowed uint32
	ApprovedCount         uint32
}

// AutomationRegistrar21MetaData contains all meta data concerning the AutomationRegistrar21 contract.
var AutomationRegistrar21MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"LINKAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"keeperRegistry\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"minLINKJuels\",\"type\":\"uint96\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"triggerType\",\"type\":\"uint8\"},{\"internalType\":\"enumAutomationRegistrar2_1.AutoApproveType\",\"name\":\"autoApproveType\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"autoApproveMaxAllowed\",\"type\":\"uint32\"}],\"internalType\":\"structAutomationRegistrar2_1.InitialTriggerConfig[]\",\"name\":\"triggerConfigs\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AmountMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FunctionNotPermitted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HashMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientPayment\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAdminAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidDataLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"LinkTransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyAdminOrOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyLink\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RegistrationRequestFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RequestNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SenderMismatch\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"AutoApproveAllowedSenderSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"keeperRegistry\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint96\",\"name\":\"minLINKJuels\",\"type\":\"uint96\"}],\"name\":\"ConfigChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"displayName\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"upkeepId\",\"type\":\"uint256\"}],\"name\":\"RegistrationApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"RegistrationRejected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"encryptedEmail\",\"type\":\"bytes\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"upkeepContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"gasLimit\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"adminAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"triggerType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"triggerConfig\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"checkData\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint96\",\"name\":\"amount\",\"type\":\"uint96\"}],\"name\":\"RegistrationRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"triggerType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"enumAutomationRegistrar2_1.AutoApproveType\",\"name\":\"autoApproveType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"autoApproveMaxAllowed\",\"type\":\"uint32\"}],\"name\":\"TriggerConfigSet\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"LINK\",\"outputs\":[{\"internalType\":\"contractLinkTokenInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"upkeepContract\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"gasLimit\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"adminAddress\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"triggerType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"checkData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"triggerConfig\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"cancel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"}],\"name\":\"getAutoApproveAllowedSender\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getConfig\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"keeperRegistry\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minLINKJuels\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"getPendingRequest\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"\",\"type\":\"uint96\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"triggerType\",\"type\":\"uint8\"}],\"name\":\"getTriggerRegistrationDetails\",\"outputs\":[{\"components\":[{\"internalType\":\"enumAutomationRegistrar2_1.AutoApproveType\",\"name\":\"autoApproveType\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"autoApproveMaxAllowed\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"approvedCount\",\"type\":\"uint32\"}],\"internalType\":\"structAutomationRegistrar2_1.TriggerRegistrationStorage\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onTokenTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"encryptedEmail\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"upkeepContract\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"gasLimit\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"adminAddress\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"triggerType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"checkData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"triggerConfig\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"},{\"internalType\":\"uint96\",\"name\":\"amount\",\"type\":\"uint96\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"encryptedEmail\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"upkeepContract\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"gasLimit\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"adminAddress\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"triggerType\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"checkData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"triggerConfig\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"offchainConfig\",\"type\":\"bytes\"},{\"internalType\":\"uint96\",\"name\":\"amount\",\"type\":\"uint96\"}],\"internalType\":\"structAutomationRegistrar2_1.RegistrationParams\",\"name\":\"requestParams\",\"type\":\"tuple\"}],\"name\":\"registerUpkeep\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"senderAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"allowed\",\"type\":\"bool\"}],\"name\":\"setAutoApproveAllowedSender\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"keeperRegistry\",\"type\":\"address\"},{\"internalType\":\"uint96\",\"name\":\"minLINKJuels\",\"type\":\"uint96\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"triggerType\",\"type\":\"uint8\"},{\"internalType\":\"enumAutomationRegistrar2_1.AutoApproveType\",\"name\":\"autoApproveType\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"autoApproveMaxAllowed\",\"type\":\"uint32\"}],\"name\":\"setTriggerConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506004361061011b5760003560e01c8063856853e6116100b2578063b5ff5b4111610081578063c4d252f511610066578063c4d252f5146103e3578063e8d4070d146103f6578063f2fde38b1461040957600080fd5b8063b5ff5b4114610369578063c3f909d41461037c57600080fd5b8063856853e61461027857806388b12d551461028b5780638da5cb5b14610338578063a4c0ed361461035657600080fd5b80633f678e11116100ee5780633f678e11146101f35780636c4cdfc31461021457806379ba5097146102275780637e776f7f1461022f57600080fd5b8063181f5a77146101205780631b6b6d2314610172578063212d0884146101be578063367b9b4f146101de575b600080fd5b61015c6040518060400160405280601981526020017f4175746f6d6174696f6e52656769737472617220322e312e300000000000000081525081565b6040516101699190611a74565b60405180910390f35b6101997f000000000000000000000000779877a7b0d9e8603169ddbd7836e478b462478981565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610169565b6101d16101cc366004611aa4565b61041c565b6040516101699190611b29565b6101f16101ec366004611b9d565b6104a9565b005b610206610201366004611bd6565b61053b565b604051908152602001610169565b6101f1610222366004611c2e565b6106d3565b6101f161076d565b61026861023d366004611c63565b73ffffffffffffffffffffffffffffffffffffffff1660009081526005602052604090205460ff1690565b6040519015158152602001610169565b6101f1610286366004611de1565b61086f565b6102ff610299366004611f40565b60009081526002602090815260409182902082518084019093525473ffffffffffffffffffffffffffffffffffffffff8116808452740100000000000000000000000000000000000000009091046bffffffffffffffffffffffff169290910182905291565b6040805173ffffffffffffffffffffffffffffffffffffffff90931683526bffffffffffffffffffffffff909116602083015201610169565b60005473ffffffffffffffffffffffffffffffffffffffff16610199565b6101f1610364366004611f59565b6109a5565b6101f1610377366004611fb5565b610ce3565b60408051808201825260045473ffffffffffffffffffffffffffffffffffffffff8116808352740100000000000000000000000000000000000000009091046bffffffffffffffffffffffff16602092830181905283519182529181019190915201610169565b6101f16103f1366004611f40565b610dc2565b6101f1610404366004611ffe565b61104c565b6101f1610417366004611c63565b6112d9565b60408051606080820183526000808352602080840182905283850182905260ff86811683526003909152908490208451928301909452835492939192839116600281111561046c5761046c611abf565b600281111561047d5761047d611abf565b8152905463ffffffff610100820481166020840152650100000000009091041660409091015292915050565b6104b16112ed565b73ffffffffffffffffffffffffffffffffffffffff821660008181526005602090815260409182902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001685151590811790915591519182527f20c6237dac83526a849285a9f79d08a483291bdd3a056a0ef9ae94ecee1ad356910160405180910390a25050565b6004546000907401000000000000000000000000000000000000000090046bffffffffffffffffffffffff1661057961014084016101208501612109565b6bffffffffffffffffffffffff1610156105bf576040517fcd1c886700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000779877a7b0d9e8603169ddbd7836e478b4624789166323b872dd333061060f61014087016101208801612109565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e086901b16815273ffffffffffffffffffffffffffffffffffffffff93841660048201529290911660248301526bffffffffffffffffffffffff1660448201526064016020604051808303816000875af1158015610696573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106ba9190612124565b506106cd6106c783612141565b33611370565b92915050565b6106db6112ed565b60408051808201825273ffffffffffffffffffffffffffffffffffffffff84168082526bffffffffffffffffffffffff8416602092830181905274010000000000000000000000000000000000000000810282176004558351918252918101919091527f39ce5d867555f0b0183e358fce5b158e7ca4fecd7c01cb7e0e19f1e23285838a910160405180910390a15050565b60015473ffffffffffffffffffffffffffffffffffffffff1633146107f3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e65720000000000000000000060448201526064015b60405180910390fd5b60008054337fffffffffffffffffffffffff00000000000000000000000000000000000000008083168217845560018054909116905560405173ffffffffffffffffffffffffffffffffffffffff90921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b3373ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000779877a7b0d9e8603169ddbd7836e478b462478916146108de576040517f018d10be00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6109966040518061014001604052808e81526020018d8d8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050509082525073ffffffffffffffffffffffffffffffffffffffff808d16602083015263ffffffff8c1660408301528a16606082015260ff8916608082015260a0810188905260c0810187905260e081018690526bffffffffffffffffffffffff85166101009091015282611370565b50505050505050505050505050565b3373ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000779877a7b0d9e8603169ddbd7836e478b46247891614610a14576040517f018d10be00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b81818080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050505060208101517fffffffff0000000000000000000000000000000000000000000000000000000081167f856853e60000000000000000000000000000000000000000000000000000000014610aca576040517fe3d6792100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8484846000610adc8260048186612276565b810190610ae991906122a0565b509950505050505050505050806bffffffffffffffffffffffff168414610b3c576040517f55e97b0d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8988886000610b4e8260048186612276565b810190610b5b91906122a0565b9a50505050505050505050508073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1614610bcc576040517ff8c5638e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6004547401000000000000000000000000000000000000000090046bffffffffffffffffffffffff168d1015610c2e576040517fcd1c886700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60003073ffffffffffffffffffffffffffffffffffffffff168d8d604051610c579291906123dd565b600060405180830381855af49150503d8060008114610c92576040519150601f19603f3d011682016040523d82523d6000602084013e610c97565b606091505b5050905080610cd2576040517f649bf81000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b505050505050505050505050505050565b610ceb6112ed565b60ff8316600090815260036020526040902080548391907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001836002811115610d3857610d38611abf565b021790555060ff83166000908152600360205260409081902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffff00000000ff1661010063ffffffff851602179055517f830a6d06a4e2caac67eba04323de22bdb04f032dd8b3d6a0c52b503d9a7036a390610db5908590859085906123ed565b60405180910390a1505050565b60008181526002602090815260409182902082518084019093525473ffffffffffffffffffffffffffffffffffffffff8116808452740100000000000000000000000000000000000000009091046bffffffffffffffffffffffff1691830191909152331480610e49575060005473ffffffffffffffffffffffffffffffffffffffff1633145b610e7f576040517f61685c2b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b805173ffffffffffffffffffffffffffffffffffffffff16610ecd576040517f4b13b31e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600082815260026020908152604080832083905583519184015190517fa9059cbb0000000000000000000000000000000000000000000000000000000081527f000000000000000000000000779877a7b0d9e8603169ddbd7836e478b462478973ffffffffffffffffffffffffffffffffffffffff169263a9059cbb92610f859260040173ffffffffffffffffffffffffffffffffffffffff9290921682526bffffffffffffffffffffffff16602082015260400190565b6020604051808303816000875af1158015610fa4573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610fc89190612124565b90508061101c5781516040517fc2e4dce800000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff90911660048201526024016107ea565b60405183907f3663fb28ebc87645eb972c9dad8521bf665c623f287e79f1c56f1eb374b82a2290600090a2505050565b6110546112ed565b60008181526002602090815260409182902082518084019093525473ffffffffffffffffffffffffffffffffffffffff8116808452740100000000000000000000000000000000000000009091046bffffffffffffffffffffffff16918301919091526110ed576040517f4b13b31e00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008b8b8b8b8b8b8b8b8b60405160200161111099989796959493929190612461565b604051602081830303815290604052805190602001209050808314611161576040517f3f4d605300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60026000848152602001908152602001600020600080820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff02191690556000820160146101000a8154906bffffffffffffffffffffffff021916905550506112c96040518061014001604052808f81526020016040518060200160405280600081525081526020018e73ffffffffffffffffffffffffffffffffffffffff1681526020018d63ffffffff1681526020018c73ffffffffffffffffffffffffffffffffffffffff1681526020018b60ff1681526020018a8a8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050509082525060208082018a905260408051601f8a0183900483028101830182528981529201919089908990819084018382808284376000920191909152505050908252506020858101516bffffffffffffffffffffffff1691015282611647565b5050505050505050505050505050565b6112e16112ed565b6112ea81611876565b50565b60005473ffffffffffffffffffffffffffffffffffffffff16331461136e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e65720000000000000000000060448201526064016107ea565b565b608082015160009073ffffffffffffffffffffffffffffffffffffffff166113c4576040517f05bb467c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008360400151846060015185608001518660a001518760c001518860e0015189610100015160405160200161140097969594939291906124e7565b604051602081830303815290604052805190602001209050836040015173ffffffffffffffffffffffffffffffffffffffff16817f7684390ebb103102f7f48c71439c2408713f8d437782a6fab2756acc0e42c1b786600001518760200151886060015189608001518a60a001518b60e001518c61010001518d60c001518e610120015160405161149999989796959493929190612569565b60405180910390a360a084015160ff9081166000908152600360205260408082208151606081019092528054929361151c9383911660028111156114df576114df611abf565b60028111156114f0576114f0611abf565b8152905463ffffffff61010082048116602084015265010000000000909104166040909101528561196b565b156115845760a085015160ff166000908152600360205260409020805465010000000000900463ffffffff1690600561155483612653565b91906101000a81548163ffffffff021916908363ffffffff1602179055505061157d8583611647565b905061163f565b61012085015160008381526002602052604081205490916115ca917401000000000000000000000000000000000000000090046bffffffffffffffffffffffff16612676565b604080518082018252608089015173ffffffffffffffffffffffffffffffffffffffff90811682526bffffffffffffffffffffffff9384166020808401918252600089815260029091529390932091519251909316740100000000000000000000000000000000000000000291909216179055505b949350505050565b600480546040808501516060860151608087015160a088015160c089015160e08a01516101008b015196517f28f32f3800000000000000000000000000000000000000000000000000000000815260009973ffffffffffffffffffffffffffffffffffffffff909916988a988a986328f32f38986116d29891979096919590949193909291016124e7565b6020604051808303816000875af11580156116f1573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061171591906126a2565b905060007f000000000000000000000000779877a7b0d9e8603169ddbd7836e478b462478973ffffffffffffffffffffffffffffffffffffffff16634000aea0848861012001518560405160200161176f91815260200190565b6040516020818303038152906040526040518463ffffffff1660e01b815260040161179c939291906126bb565b6020604051808303816000875af11580156117bb573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906117df9190612124565b905080611830576040517fc2e4dce800000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff841660048201526024016107ea565b81857fb9a292fb7e3edd920cd2d2829a3615a640c43fd7de0a0820aa0668feb4c37d4b88600001516040516118659190611a74565b60405180910390a350949350505050565b3373ffffffffffffffffffffffffffffffffffffffff8216036118f5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c6600000000000000000060448201526064016107ea565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b6000808351600281111561198157611981611abf565b0361198e575060006106cd565b6001835160028111156119a3576119a3611abf565b1480156119d6575073ffffffffffffffffffffffffffffffffffffffff821660009081526005602052604090205460ff16155b156119e3575060006106cd565b826020015163ffffffff16836040015163ffffffff161015611a07575060016106cd565b50600092915050565b6000815180845260005b81811015611a3657602081850181015186830182015201611a1a565b5060006020828601015260207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f83011685010191505092915050565b602081526000611a876020830184611a10565b9392505050565b803560ff81168114611a9f57600080fd5b919050565b600060208284031215611ab657600080fd5b611a8782611a8e565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60038110611b25577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b9052565b6000606082019050611b3c828451611aee565b602083015163ffffffff8082166020850152806040860151166040850152505092915050565b73ffffffffffffffffffffffffffffffffffffffff811681146112ea57600080fd5b8035611a9f81611b62565b80151581146112ea57600080fd5b60008060408385031215611bb057600080fd5b8235611bbb81611b62565b91506020830135611bcb81611b8f565b809150509250929050565b600060208284031215611be857600080fd5b813567ffffffffffffffff811115611bff57600080fd5b82016101408185031215611a8757600080fd5b80356bffffffffffffffffffffffff81168114611a9f57600080fd5b60008060408385031215611c4157600080fd5b8235611c4c81611b62565b9150611c5a60208401611c12565b90509250929050565b600060208284031215611c7557600080fd5b8135611a8781611b62565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051610140810167ffffffffffffffff81118282101715611cd357611cd3611c80565b60405290565b600082601f830112611cea57600080fd5b813567ffffffffffffffff80821115611d0557611d05611c80565b604051601f83017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f01168101908282118183101715611d4b57611d4b611c80565b81604052838152866020858801011115611d6457600080fd5b836020870160208301376000602085830101528094505050505092915050565b60008083601f840112611d9657600080fd5b50813567ffffffffffffffff811115611dae57600080fd5b602083019150836020828501011115611dc657600080fd5b9250929050565b803563ffffffff81168114611a9f57600080fd5b6000806000806000806000806000806000806101608d8f031215611e0457600080fd5b67ffffffffffffffff8d351115611e1a57600080fd5b611e278e8e358f01611cd9565b9b5067ffffffffffffffff60208e01351115611e4257600080fd5b611e528e60208f01358f01611d84565b909b509950611e6360408e01611b84565b9850611e7160608e01611dcd565b9750611e7f60808e01611b84565b9650611e8d60a08e01611a8e565b955067ffffffffffffffff60c08e01351115611ea857600080fd5b611eb88e60c08f01358f01611cd9565b945067ffffffffffffffff60e08e01351115611ed357600080fd5b611ee38e60e08f01358f01611cd9565b935067ffffffffffffffff6101008e01351115611eff57600080fd5b611f108e6101008f01358f01611cd9565b9250611f1f6101208e01611c12565b9150611f2e6101408e01611b84565b90509295989b509295989b509295989b565b600060208284031215611f5257600080fd5b5035919050565b60008060008060608587031215611f6f57600080fd5b8435611f7a81611b62565b935060208501359250604085013567ffffffffffffffff811115611f9d57600080fd5b611fa987828801611d84565b95989497509550505050565b600080600060608486031215611fca57600080fd5b611fd384611a8e565b9250602084013560038110611fe757600080fd5b9150611ff560408501611dcd565b90509250925092565b60008060008060008060008060008060006101208c8e03121561202057600080fd5b67ffffffffffffffff808d35111561203757600080fd5b6120448e8e358f01611cd9565b9b5061205260208e01611b84565b9a5061206060408e01611dcd565b995061206e60608e01611b84565b985061207c60808e01611a8e565b97508060a08e0135111561208f57600080fd5b61209f8e60a08f01358f01611d84565b909750955060c08d01358110156120b557600080fd5b6120c58e60c08f01358f01611cd9565b94508060e08e013511156120d857600080fd5b506120e98d60e08e01358e01611d84565b81945080935050506101008c013590509295989b509295989b9093969950565b60006020828403121561211b57600080fd5b611a8782611c12565b60006020828403121561213657600080fd5b8151611a8781611b8f565b6000610140823603121561215457600080fd5b61215c611caf565b823567ffffffffffffffff8082111561217457600080fd5b61218036838701611cd9565b8352602085013591508082111561219657600080fd5b6121a236838701611cd9565b60208401526121b360408601611b84565b60408401526121c460608601611dcd565b60608401526121d560808601611b84565b60808401526121e660a08601611a8e565b60a084015260c08501359150808211156121ff57600080fd5b61220b36838701611cd9565b60c084015260e085013591508082111561222457600080fd5b61223036838701611cd9565b60e08401526101009150818501358181111561224b57600080fd5b61225736828801611cd9565b8385015250505061012061226c818501611c12565b9082015292915050565b6000808585111561228657600080fd5b8386111561229357600080fd5b5050820193919092039150565b60008060008060008060008060008060006101608c8e0312156122c257600080fd5b67ffffffffffffffff808d3511156122d957600080fd5b6122e68e8e358f01611cd9565b9b508060208e013511156122f957600080fd5b6123098e60208f01358f01611cd9565b9a5061231760408e01611b84565b995061232560608e01611dcd565b985061233360808e01611b84565b975061234160a08e01611a8e565b96508060c08e0135111561235457600080fd5b6123648e60c08f01358f01611cd9565b95508060e08e0135111561237757600080fd5b6123878e60e08f01358f01611cd9565b9450806101008e0135111561239b57600080fd5b506123ad8d6101008e01358e01611cd9565b92506123bc6101208d01611c12565b91506123cb6101408d01611b84565b90509295989b509295989b9093969950565b8183823760009101908152919050565b60ff84168152606081016124046020830185611aee565b63ffffffff83166040830152949350505050565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b600073ffffffffffffffffffffffffffffffffffffffff808c16835263ffffffff8b166020840152808a1660408401525060ff8816606083015260e060808301526124b060e083018789612418565b82810360a08401526124c28187611a10565b905082810360c08401526124d7818587612418565b9c9b505050505050505050505050565b600073ffffffffffffffffffffffffffffffffffffffff808a16835263ffffffff8916602084015280881660408401525060ff8616606083015260e0608083015261253560e0830186611a10565b82810360a08401526125478186611a10565b905082810360c084015261255b8185611a10565b9a9950505050505050505050565b600061012080835261257d8184018d611a10565b90508281036020840152612591818c611a10565b905063ffffffff8a16604084015273ffffffffffffffffffffffffffffffffffffffff8916606084015260ff8816608084015282810360a08401526125d68188611a10565b905082810360c08401526125ea8187611a10565b905082810360e08401526125fe8186611a10565b9150506bffffffffffffffffffffffff83166101008301529a9950505050505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600063ffffffff80831681810361266c5761266c612624565b6001019392505050565b6bffffffffffffffffffffffff81811683821601908082111561269b5761269b612624565b5092915050565b6000602082840312156126b457600080fd5b5051919050565b73ffffffffffffffffffffffffffffffffffffffff841681526bffffffffffffffffffffffff831660208201526060604082015260006126fe6060830184611a10565b9594505050505056fea164736f6c6343000810000a",
}

// AutomationRegistrar21ABI is the input ABI used to generate the binding from.
// Deprecated: Use AutomationRegistrar21MetaData.ABI instead.
var AutomationRegistrar21ABI = AutomationRegistrar21MetaData.ABI

// AutomationRegistrar21Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AutomationRegistrar21MetaData.Bin instead.
var AutomationRegistrar21Bin = AutomationRegistrar21MetaData.Bin

// DeployAutomationRegistrar21 deploys a new Ethereum contract, binding an instance of AutomationRegistrar21 to it.
func DeployAutomationRegistrar21(auth *bind.TransactOpts, backend bind.ContractBackend, LINKAddress common.Address, keeperRegistry common.Address, minLINKJuels *big.Int, triggerConfigs []AutomationRegistrar21InitialTriggerConfig) (common.Address, *types.Transaction, *AutomationRegistrar21, error) {
	parsed, err := AutomationRegistrar21MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AutomationRegistrar21Bin), backend, LINKAddress, keeperRegistry, minLINKJuels, triggerConfigs)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AutomationRegistrar21{AutomationRegistrar21Caller: AutomationRegistrar21Caller{contract: contract}, AutomationRegistrar21Transactor: AutomationRegistrar21Transactor{contract: contract}, AutomationRegistrar21Filterer: AutomationRegistrar21Filterer{contract: contract}}, nil
}

// AutomationRegistrar21 is an auto generated Go binding around an Ethereum contract.
type AutomationRegistrar21 struct {
	AutomationRegistrar21Caller     // Read-only binding to the contract
	AutomationRegistrar21Transactor // Write-only binding to the contract
	AutomationRegistrar21Filterer   // Log filterer for contract events
}

// AutomationRegistrar21Caller is an auto generated read-only Go binding around an Ethereum contract.
type AutomationRegistrar21Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AutomationRegistrar21Transactor is an auto generated write-only Go binding around an Ethereum contract.
type AutomationRegistrar21Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AutomationRegistrar21Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AutomationRegistrar21Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AutomationRegistrar21Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AutomationRegistrar21Session struct {
	Contract     *AutomationRegistrar21 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// AutomationRegistrar21CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AutomationRegistrar21CallerSession struct {
	Contract *AutomationRegistrar21Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// AutomationRegistrar21TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AutomationRegistrar21TransactorSession struct {
	Contract     *AutomationRegistrar21Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// AutomationRegistrar21Raw is an auto generated low-level Go binding around an Ethereum contract.
type AutomationRegistrar21Raw struct {
	Contract *AutomationRegistrar21 // Generic contract binding to access the raw methods on
}

// AutomationRegistrar21CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AutomationRegistrar21CallerRaw struct {
	Contract *AutomationRegistrar21Caller // Generic read-only contract binding to access the raw methods on
}

// AutomationRegistrar21TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AutomationRegistrar21TransactorRaw struct {
	Contract *AutomationRegistrar21Transactor // Generic write-only contract binding to access the raw methods on
}

// NewAutomationRegistrar21 creates a new instance of AutomationRegistrar21, bound to a specific deployed contract.
func NewAutomationRegistrar21(address common.Address, backend bind.ContractBackend) (*AutomationRegistrar21, error) {
	contract, err := bindAutomationRegistrar21(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AutomationRegistrar21{AutomationRegistrar21Caller: AutomationRegistrar21Caller{contract: contract}, AutomationRegistrar21Transactor: AutomationRegistrar21Transactor{contract: contract}, AutomationRegistrar21Filterer: AutomationRegistrar21Filterer{contract: contract}}, nil
}

// NewAutomationRegistrar21Caller creates a new read-only instance of AutomationRegistrar21, bound to a specific deployed contract.
func NewAutomationRegistrar21Caller(address common.Address, caller bind.ContractCaller) (*AutomationRegistrar21Caller, error) {
	contract, err := bindAutomationRegistrar21(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AutomationRegistrar21Caller{contract: contract}, nil
}

// NewAutomationRegistrar21Transactor creates a new write-only instance of AutomationRegistrar21, bound to a specific deployed contract.
func NewAutomationRegistrar21Transactor(address common.Address, transactor bind.ContractTransactor) (*AutomationRegistrar21Transactor, error) {
	contract, err := bindAutomationRegistrar21(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AutomationRegistrar21Transactor{contract: contract}, nil
}

// NewAutomationRegistrar21Filterer creates a new log filterer instance of AutomationRegistrar21, bound to a specific deployed contract.
func NewAutomationRegistrar21Filterer(address common.Address, filterer bind.ContractFilterer) (*AutomationRegistrar21Filterer, error) {
	contract, err := bindAutomationRegistrar21(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AutomationRegistrar21Filterer{contract: contract}, nil
}

// bindAutomationRegistrar21 binds a generic wrapper to an already deployed contract.
func bindAutomationRegistrar21(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AutomationRegistrar21MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AutomationRegistrar21 *AutomationRegistrar21Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AutomationRegistrar21.Contract.AutomationRegistrar21Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AutomationRegistrar21 *AutomationRegistrar21Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AutomationRegistrar21.Contract.AutomationRegistrar21Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AutomationRegistrar21 *AutomationRegistrar21Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AutomationRegistrar21.Contract.AutomationRegistrar21Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AutomationRegistrar21 *AutomationRegistrar21CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AutomationRegistrar21.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AutomationRegistrar21 *AutomationRegistrar21TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AutomationRegistrar21.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AutomationRegistrar21 *AutomationRegistrar21TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AutomationRegistrar21.Contract.contract.Transact(opts, method, params...)
}

// LINK is a free data retrieval call binding the contract method 0x1b6b6d23.
//
// Solidity: function LINK() view returns(address)
func (_AutomationRegistrar21 *AutomationRegistrar21Caller) LINK(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AutomationRegistrar21.contract.Call(opts, &out, "LINK")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LINK is a free data retrieval call binding the contract method 0x1b6b6d23.
//
// Solidity: function LINK() view returns(address)
func (_AutomationRegistrar21 *AutomationRegistrar21Session) LINK() (common.Address, error) {
	return _AutomationRegistrar21.Contract.LINK(&_AutomationRegistrar21.CallOpts)
}

// LINK is a free data retrieval call binding the contract method 0x1b6b6d23.
//
// Solidity: function LINK() view returns(address)
func (_AutomationRegistrar21 *AutomationRegistrar21CallerSession) LINK() (common.Address, error) {
	return _AutomationRegistrar21.Contract.LINK(&_AutomationRegistrar21.CallOpts)
}

// GetAutoApproveAllowedSender is a free data retrieval call binding the contract method 0x7e776f7f.
//
// Solidity: function getAutoApproveAllowedSender(address senderAddress) view returns(bool)
func (_AutomationRegistrar21 *AutomationRegistrar21Caller) GetAutoApproveAllowedSender(opts *bind.CallOpts, senderAddress common.Address) (bool, error) {
	var out []interface{}
	err := _AutomationRegistrar21.contract.Call(opts, &out, "getAutoApproveAllowedSender", senderAddress)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetAutoApproveAllowedSender is a free data retrieval call binding the contract method 0x7e776f7f.
//
// Solidity: function getAutoApproveAllowedSender(address senderAddress) view returns(bool)
func (_AutomationRegistrar21 *AutomationRegistrar21Session) GetAutoApproveAllowedSender(senderAddress common.Address) (bool, error) {
	return _AutomationRegistrar21.Contract.GetAutoApproveAllowedSender(&_AutomationRegistrar21.CallOpts, senderAddress)
}

// GetAutoApproveAllowedSender is a free data retrieval call binding the contract method 0x7e776f7f.
//
// Solidity: function getAutoApproveAllowedSender(address senderAddress) view returns(bool)
func (_AutomationRegistrar21 *AutomationRegistrar21CallerSession) GetAutoApproveAllowedSender(senderAddress common.Address) (bool, error) {
	return _AutomationRegistrar21.Contract.GetAutoApproveAllowedSender(&_AutomationRegistrar21.CallOpts, senderAddress)
}

// GetConfig is a free data retrieval call binding the contract method 0xc3f909d4.
//
// Solidity: function getConfig() view returns(address keeperRegistry, uint256 minLINKJuels)
func (_AutomationRegistrar21 *AutomationRegistrar21Caller) GetConfig(opts *bind.CallOpts) (struct {
	KeeperRegistry common.Address
	MinLINKJuels   *big.Int
}, error) {
	var out []interface{}
	err := _AutomationRegistrar21.contract.Call(opts, &out, "getConfig")

	outstruct := new(struct {
		KeeperRegistry common.Address
		MinLINKJuels   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.KeeperRegistry = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.MinLINKJuels = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetConfig is a free data retrieval call binding the contract method 0xc3f909d4.
//
// Solidity: function getConfig() view returns(address keeperRegistry, uint256 minLINKJuels)
func (_AutomationRegistrar21 *AutomationRegistrar21Session) GetConfig() (struct {
	KeeperRegistry common.Address
	MinLINKJuels   *big.Int
}, error) {
	return _AutomationRegistrar21.Contract.GetConfig(&_AutomationRegistrar21.CallOpts)
}

// GetConfig is a free data retrieval call binding the contract method 0xc3f909d4.
//
// Solidity: function getConfig() view returns(address keeperRegistry, uint256 minLINKJuels)
func (_AutomationRegistrar21 *AutomationRegistrar21CallerSession) GetConfig() (struct {
	KeeperRegistry common.Address
	MinLINKJuels   *big.Int
}, error) {
	return _AutomationRegistrar21.Contract.GetConfig(&_AutomationRegistrar21.CallOpts)
}

// GetPendingRequest is a free data retrieval call binding the contract method 0x88b12d55.
//
// Solidity: function getPendingRequest(bytes32 hash) view returns(address, uint96)
func (_AutomationRegistrar21 *AutomationRegistrar21Caller) GetPendingRequest(opts *bind.CallOpts, hash [32]byte) (common.Address, *big.Int, error) {
	var out []interface{}
	err := _AutomationRegistrar21.contract.Call(opts, &out, "getPendingRequest", hash)

	if err != nil {
		return *new(common.Address), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetPendingRequest is a free data retrieval call binding the contract method 0x88b12d55.
//
// Solidity: function getPendingRequest(bytes32 hash) view returns(address, uint96)
func (_AutomationRegistrar21 *AutomationRegistrar21Session) GetPendingRequest(hash [32]byte) (common.Address, *big.Int, error) {
	return _AutomationRegistrar21.Contract.GetPendingRequest(&_AutomationRegistrar21.CallOpts, hash)
}

// GetPendingRequest is a free data retrieval call binding the contract method 0x88b12d55.
//
// Solidity: function getPendingRequest(bytes32 hash) view returns(address, uint96)
func (_AutomationRegistrar21 *AutomationRegistrar21CallerSession) GetPendingRequest(hash [32]byte) (common.Address, *big.Int, error) {
	return _AutomationRegistrar21.Contract.GetPendingRequest(&_AutomationRegistrar21.CallOpts, hash)
}

// GetTriggerRegistrationDetails is a free data retrieval call binding the contract method 0x212d0884.
//
// Solidity: function getTriggerRegistrationDetails(uint8 triggerType) view returns((uint8,uint32,uint32))
func (_AutomationRegistrar21 *AutomationRegistrar21Caller) GetTriggerRegistrationDetails(opts *bind.CallOpts, triggerType uint8) (AutomationRegistrar21TriggerRegistrationStorage, error) {
	var out []interface{}
	err := _AutomationRegistrar21.contract.Call(opts, &out, "getTriggerRegistrationDetails", triggerType)

	if err != nil {
		return *new(AutomationRegistrar21TriggerRegistrationStorage), err
	}

	out0 := *abi.ConvertType(out[0], new(AutomationRegistrar21TriggerRegistrationStorage)).(*AutomationRegistrar21TriggerRegistrationStorage)

	return out0, err

}

// GetTriggerRegistrationDetails is a free data retrieval call binding the contract method 0x212d0884.
//
// Solidity: function getTriggerRegistrationDetails(uint8 triggerType) view returns((uint8,uint32,uint32))
func (_AutomationRegistrar21 *AutomationRegistrar21Session) GetTriggerRegistrationDetails(triggerType uint8) (AutomationRegistrar21TriggerRegistrationStorage, error) {
	return _AutomationRegistrar21.Contract.GetTriggerRegistrationDetails(&_AutomationRegistrar21.CallOpts, triggerType)
}

// GetTriggerRegistrationDetails is a free data retrieval call binding the contract method 0x212d0884.
//
// Solidity: function getTriggerRegistrationDetails(uint8 triggerType) view returns((uint8,uint32,uint32))
func (_AutomationRegistrar21 *AutomationRegistrar21CallerSession) GetTriggerRegistrationDetails(triggerType uint8) (AutomationRegistrar21TriggerRegistrationStorage, error) {
	return _AutomationRegistrar21.Contract.GetTriggerRegistrationDetails(&_AutomationRegistrar21.CallOpts, triggerType)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AutomationRegistrar21 *AutomationRegistrar21Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AutomationRegistrar21.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AutomationRegistrar21 *AutomationRegistrar21Session) Owner() (common.Address, error) {
	return _AutomationRegistrar21.Contract.Owner(&_AutomationRegistrar21.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AutomationRegistrar21 *AutomationRegistrar21CallerSession) Owner() (common.Address, error) {
	return _AutomationRegistrar21.Contract.Owner(&_AutomationRegistrar21.CallOpts)
}

// TypeAndVersion is a free data retrieval call binding the contract method 0x181f5a77.
//
// Solidity: function typeAndVersion() view returns(string)
func (_AutomationRegistrar21 *AutomationRegistrar21Caller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _AutomationRegistrar21.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TypeAndVersion is a free data retrieval call binding the contract method 0x181f5a77.
//
// Solidity: function typeAndVersion() view returns(string)
func (_AutomationRegistrar21 *AutomationRegistrar21Session) TypeAndVersion() (string, error) {
	return _AutomationRegistrar21.Contract.TypeAndVersion(&_AutomationRegistrar21.CallOpts)
}

// TypeAndVersion is a free data retrieval call binding the contract method 0x181f5a77.
//
// Solidity: function typeAndVersion() view returns(string)
func (_AutomationRegistrar21 *AutomationRegistrar21CallerSession) TypeAndVersion() (string, error) {
	return _AutomationRegistrar21.Contract.TypeAndVersion(&_AutomationRegistrar21.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_AutomationRegistrar21 *AutomationRegistrar21Transactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AutomationRegistrar21.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_AutomationRegistrar21 *AutomationRegistrar21Session) AcceptOwnership() (*types.Transaction, error) {
	return _AutomationRegistrar21.Contract.AcceptOwnership(&_AutomationRegistrar21.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_AutomationRegistrar21 *AutomationRegistrar21TransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _AutomationRegistrar21.Contract.AcceptOwnership(&_AutomationRegistrar21.TransactOpts)
}

// Approve is a paid mutator transaction binding the contract method 0xe8d4070d.
//
// Solidity: function approve(string name, address upkeepContract, uint32 gasLimit, address adminAddress, uint8 triggerType, bytes checkData, bytes triggerConfig, bytes offchainConfig, bytes32 hash) returns()
func (_AutomationRegistrar21 *AutomationRegistrar21Transactor) Approve(opts *bind.TransactOpts, name string, upkeepContract common.Address, gasLimit uint32, adminAddress common.Address, triggerType uint8, checkData []byte, triggerConfig []byte, offchainConfig []byte, hash [32]byte) (*types.Transaction, error) {
	return _AutomationRegistrar21.contract.Transact(opts, "approve", name, upkeepContract, gasLimit, adminAddress, triggerType, checkData, triggerConfig, offchainConfig, hash)
}

// Approve is a paid mutator transaction binding the contract method 0xe8d4070d.
//
// Solidity: function approve(string name, address upkeepContract, uint32 gasLimit, address adminAddress, uint8 triggerType, bytes checkData, bytes triggerConfig, bytes offchainConfig, bytes32 hash) returns()
func (_AutomationRegistrar21 *AutomationRegistrar21Session) Approve(name string, upkeepContract common.Address, gasLimit uint32, adminAddress common.Address, triggerType uint8, checkData []byte, triggerConfig []byte, offchainConfig []byte, hash [32]byte) (*types.Transaction, error) {
	return _AutomationRegistrar21.Contract.Approve(&_AutomationRegistrar21.TransactOpts, name, upkeepContract, gasLimit, adminAddress, triggerType, checkData, triggerConfig, offchainConfig, hash)
}

// Approve is a paid mutator transaction binding the contract method 0xe8d4070d.
//
// Solidity: function approve(string name, address upkeepContract, uint32 gasLimit, address adminAddress, uint8 triggerType, bytes checkData, bytes triggerConfig, bytes offchainConfig, bytes32 hash) returns()
func (_AutomationRegistrar21 *AutomationRegistrar21TransactorSession) Approve(name string, upkeepContract common.Address, gasLimit uint32, adminAddress common.Address, triggerType uint8, checkData []byte, triggerConfig []byte, offchainConfig []byte, hash [32]byte) (*types.Transaction, error) {
	return _AutomationRegistrar21.Contract.Approve(&_AutomationRegistrar21.TransactOpts, name, upkeepContract, gasLimit, adminAddress, triggerType, checkData, triggerConfig, offchainConfig, hash)
}

// Cancel is a paid mutator transaction binding the contract method 0xc4d252f5.
//
// Solidity: function cancel(bytes32 hash) returns()
func (_AutomationRegistrar21 *AutomationRegistrar21Transactor) Cancel(opts *bind.TransactOpts, hash [32]byte) (*types.Transaction, error) {
	return _AutomationRegistrar21.contract.Transact(opts, "cancel", hash)
}

// Cancel is a paid mutator transaction binding the contract method 0xc4d252f5.
//
// Solidity: function cancel(bytes32 hash) returns()
func (_AutomationRegistrar21 *AutomationRegistrar21Session) Cancel(hash [32]byte) (*types.Transaction, error) {
	return _AutomationRegistrar21.Contract.Cancel(&_AutomationRegistrar21.TransactOpts, hash)
}

// Cancel is a paid mutator transaction binding the contract method 0xc4d252f5.
//
// Solidity: function cancel(bytes32 hash) returns()
func (_AutomationRegistrar21 *AutomationRegistrar21TransactorSession) Cancel(hash [32]byte) (*types.Transaction, error) {
	return _AutomationRegistrar21.Contract.Cancel(&_AutomationRegistrar21.TransactOpts, hash)
}

// OnTokenTransfer is a paid mutator transaction binding the contract method 0xa4c0ed36.
//
// Solidity: function onTokenTransfer(address sender, uint256 amount, bytes data) returns()
func (_AutomationRegistrar21 *AutomationRegistrar21Transactor) OnTokenTransfer(opts *bind.TransactOpts, sender common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _AutomationRegistrar21.contract.Transact(opts, "onTokenTransfer", sender, amount, data)
}

// OnTokenTransfer is a paid mutator transaction binding the contract method 0xa4c0ed36.
//
// Solidity: function onTokenTransfer(address sender, uint256 amount, bytes data) returns()
func (_AutomationRegistrar21 *AutomationRegistrar21Session) OnTokenTransfer(sender common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _AutomationRegistrar21.Contract.OnTokenTransfer(&_AutomationRegistrar21.TransactOpts, sender, amount, data)
}

// OnTokenTransfer is a paid mutator transaction binding the contract method 0xa4c0ed36.
//
// Solidity: function onTokenTransfer(address sender, uint256 amount, bytes data) returns()
func (_AutomationRegistrar21 *AutomationRegistrar21TransactorSession) OnTokenTransfer(sender common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _AutomationRegistrar21.Contract.OnTokenTransfer(&_AutomationRegistrar21.TransactOpts, sender, amount, data)
}

// Register is a paid mutator transaction binding the contract method 0x856853e6.
//
// Solidity: function register(string name, bytes encryptedEmail, address upkeepContract, uint32 gasLimit, address adminAddress, uint8 triggerType, bytes checkData, bytes triggerConfig, bytes offchainConfig, uint96 amount, address sender) returns()
func (_AutomationRegistrar21 *AutomationRegistrar21Transactor) Register(opts *bind.TransactOpts, name string, encryptedEmail []byte, upkeepContract common.Address, gasLimit uint32, adminAddress common.Address, triggerType uint8, checkData []byte, triggerConfig []byte, offchainConfig []byte, amount *big.Int, sender common.Address) (*types.Transaction, error) {
	return _AutomationRegistrar21.contract.Transact(opts, "register", name, encryptedEmail, upkeepContract, gasLimit, adminAddress, triggerType, checkData, triggerConfig, offchainConfig, amount, sender)
}

// Register is a paid mutator transaction binding the contract method 0x856853e6.
//
// Solidity: function register(string name, bytes encryptedEmail, address upkeepContract, uint32 gasLimit, address adminAddress, uint8 triggerType, bytes checkData, bytes triggerConfig, bytes offchainConfig, uint96 amount, address sender) returns()
func (_AutomationRegistrar21 *AutomationRegistrar21Session) Register(name string, encryptedEmail []byte, upkeepContract common.Address, gasLimit uint32, adminAddress common.Address, triggerType uint8, checkData []byte, triggerConfig []byte, offchainConfig []byte, amount *big.Int, sender common.Address) (*types.Transaction, error) {
	return _AutomationRegistrar21.Contract.Register(&_AutomationRegistrar21.TransactOpts, name, encryptedEmail, upkeepContract, gasLimit, adminAddress, triggerType, checkData, triggerConfig, offchainConfig, amount, sender)
}

// Register is a paid mutator transaction binding the contract method 0x856853e6.
//
// Solidity: function register(string name, bytes encryptedEmail, address upkeepContract, uint32 gasLimit, address adminAddress, uint8 triggerType, bytes checkData, bytes triggerConfig, bytes offchainConfig, uint96 amount, address sender) returns()
func (_AutomationRegistrar21 *AutomationRegistrar21TransactorSession) Register(name string, encryptedEmail []byte, upkeepContract common.Address, gasLimit uint32, adminAddress common.Address, triggerType uint8, checkData []byte, triggerConfig []byte, offchainConfig []byte, amount *big.Int, sender common.Address) (*types.Transaction, error) {
	return _AutomationRegistrar21.Contract.Register(&_AutomationRegistrar21.TransactOpts, name, encryptedEmail, upkeepContract, gasLimit, adminAddress, triggerType, checkData, triggerConfig, offchainConfig, amount, sender)
}

// RegisterUpkeep is a paid mutator transaction binding the contract method 0x3f678e11.
//
// Solidity: function registerUpkeep((string,bytes,address,uint32,address,uint8,bytes,bytes,bytes,uint96) requestParams) returns(uint256)
func (_AutomationRegistrar21 *AutomationRegistrar21Transactor) RegisterUpkeep(opts *bind.TransactOpts, requestParams AutomationRegistrar21RegistrationParams) (*types.Transaction, error) {
	return _AutomationRegistrar21.contract.Transact(opts, "registerUpkeep", requestParams)
}

// RegisterUpkeep is a paid mutator transaction binding the contract method 0x3f678e11.
//
// Solidity: function registerUpkeep((string,bytes,address,uint32,address,uint8,bytes,bytes,bytes,uint96) requestParams) returns(uint256)
func (_AutomationRegistrar21 *AutomationRegistrar21Session) RegisterUpkeep(requestParams AutomationRegistrar21RegistrationParams) (*types.Transaction, error) {
	return _AutomationRegistrar21.Contract.RegisterUpkeep(&_AutomationRegistrar21.TransactOpts, requestParams)
}

// RegisterUpkeep is a paid mutator transaction binding the contract method 0x3f678e11.
//
// Solidity: function registerUpkeep((string,bytes,address,uint32,address,uint8,bytes,bytes,bytes,uint96) requestParams) returns(uint256)
func (_AutomationRegistrar21 *AutomationRegistrar21TransactorSession) RegisterUpkeep(requestParams AutomationRegistrar21RegistrationParams) (*types.Transaction, error) {
	return _AutomationRegistrar21.Contract.RegisterUpkeep(&_AutomationRegistrar21.TransactOpts, requestParams)
}

// SetAutoApproveAllowedSender is a paid mutator transaction binding the contract method 0x367b9b4f.
//
// Solidity: function setAutoApproveAllowedSender(address senderAddress, bool allowed) returns()
func (_AutomationRegistrar21 *AutomationRegistrar21Transactor) SetAutoApproveAllowedSender(opts *bind.TransactOpts, senderAddress common.Address, allowed bool) (*types.Transaction, error) {
	return _AutomationRegistrar21.contract.Transact(opts, "setAutoApproveAllowedSender", senderAddress, allowed)
}

// SetAutoApproveAllowedSender is a paid mutator transaction binding the contract method 0x367b9b4f.
//
// Solidity: function setAutoApproveAllowedSender(address senderAddress, bool allowed) returns()
func (_AutomationRegistrar21 *AutomationRegistrar21Session) SetAutoApproveAllowedSender(senderAddress common.Address, allowed bool) (*types.Transaction, error) {
	return _AutomationRegistrar21.Contract.SetAutoApproveAllowedSender(&_AutomationRegistrar21.TransactOpts, senderAddress, allowed)
}

// SetAutoApproveAllowedSender is a paid mutator transaction binding the contract method 0x367b9b4f.
//
// Solidity: function setAutoApproveAllowedSender(address senderAddress, bool allowed) returns()
func (_AutomationRegistrar21 *AutomationRegistrar21TransactorSession) SetAutoApproveAllowedSender(senderAddress common.Address, allowed bool) (*types.Transaction, error) {
	return _AutomationRegistrar21.Contract.SetAutoApproveAllowedSender(&_AutomationRegistrar21.TransactOpts, senderAddress, allowed)
}

// SetConfig is a paid mutator transaction binding the contract method 0x6c4cdfc3.
//
// Solidity: function setConfig(address keeperRegistry, uint96 minLINKJuels) returns()
func (_AutomationRegistrar21 *AutomationRegistrar21Transactor) SetConfig(opts *bind.TransactOpts, keeperRegistry common.Address, minLINKJuels *big.Int) (*types.Transaction, error) {
	return _AutomationRegistrar21.contract.Transact(opts, "setConfig", keeperRegistry, minLINKJuels)
}

// SetConfig is a paid mutator transaction binding the contract method 0x6c4cdfc3.
//
// Solidity: function setConfig(address keeperRegistry, uint96 minLINKJuels) returns()
func (_AutomationRegistrar21 *AutomationRegistrar21Session) SetConfig(keeperRegistry common.Address, minLINKJuels *big.Int) (*types.Transaction, error) {
	return _AutomationRegistrar21.Contract.SetConfig(&_AutomationRegistrar21.TransactOpts, keeperRegistry, minLINKJuels)
}

// SetConfig is a paid mutator transaction binding the contract method 0x6c4cdfc3.
//
// Solidity: function setConfig(address keeperRegistry, uint96 minLINKJuels) returns()
func (_AutomationRegistrar21 *AutomationRegistrar21TransactorSession) SetConfig(keeperRegistry common.Address, minLINKJuels *big.Int) (*types.Transaction, error) {
	return _AutomationRegistrar21.Contract.SetConfig(&_AutomationRegistrar21.TransactOpts, keeperRegistry, minLINKJuels)
}

// SetTriggerConfig is a paid mutator transaction binding the contract method 0xb5ff5b41.
//
// Solidity: function setTriggerConfig(uint8 triggerType, uint8 autoApproveType, uint32 autoApproveMaxAllowed) returns()
func (_AutomationRegistrar21 *AutomationRegistrar21Transactor) SetTriggerConfig(opts *bind.TransactOpts, triggerType uint8, autoApproveType uint8, autoApproveMaxAllowed uint32) (*types.Transaction, error) {
	return _AutomationRegistrar21.contract.Transact(opts, "setTriggerConfig", triggerType, autoApproveType, autoApproveMaxAllowed)
}

// SetTriggerConfig is a paid mutator transaction binding the contract method 0xb5ff5b41.
//
// Solidity: function setTriggerConfig(uint8 triggerType, uint8 autoApproveType, uint32 autoApproveMaxAllowed) returns()
func (_AutomationRegistrar21 *AutomationRegistrar21Session) SetTriggerConfig(triggerType uint8, autoApproveType uint8, autoApproveMaxAllowed uint32) (*types.Transaction, error) {
	return _AutomationRegistrar21.Contract.SetTriggerConfig(&_AutomationRegistrar21.TransactOpts, triggerType, autoApproveType, autoApproveMaxAllowed)
}

// SetTriggerConfig is a paid mutator transaction binding the contract method 0xb5ff5b41.
//
// Solidity: function setTriggerConfig(uint8 triggerType, uint8 autoApproveType, uint32 autoApproveMaxAllowed) returns()
func (_AutomationRegistrar21 *AutomationRegistrar21TransactorSession) SetTriggerConfig(triggerType uint8, autoApproveType uint8, autoApproveMaxAllowed uint32) (*types.Transaction, error) {
	return _AutomationRegistrar21.Contract.SetTriggerConfig(&_AutomationRegistrar21.TransactOpts, triggerType, autoApproveType, autoApproveMaxAllowed)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address to) returns()
func (_AutomationRegistrar21 *AutomationRegistrar21Transactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _AutomationRegistrar21.contract.Transact(opts, "transferOwnership", to)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address to) returns()
func (_AutomationRegistrar21 *AutomationRegistrar21Session) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _AutomationRegistrar21.Contract.TransferOwnership(&_AutomationRegistrar21.TransactOpts, to)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address to) returns()
func (_AutomationRegistrar21 *AutomationRegistrar21TransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _AutomationRegistrar21.Contract.TransferOwnership(&_AutomationRegistrar21.TransactOpts, to)
}

// AutomationRegistrar21AutoApproveAllowedSenderSetIterator is returned from FilterAutoApproveAllowedSenderSet and is used to iterate over the raw logs and unpacked data for AutoApproveAllowedSenderSet events raised by the AutomationRegistrar21 contract.
type AutomationRegistrar21AutoApproveAllowedSenderSetIterator struct {
	Event *AutomationRegistrar21AutoApproveAllowedSenderSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AutomationRegistrar21AutoApproveAllowedSenderSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AutomationRegistrar21AutoApproveAllowedSenderSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AutomationRegistrar21AutoApproveAllowedSenderSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AutomationRegistrar21AutoApproveAllowedSenderSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AutomationRegistrar21AutoApproveAllowedSenderSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AutomationRegistrar21AutoApproveAllowedSenderSet represents a AutoApproveAllowedSenderSet event raised by the AutomationRegistrar21 contract.
type AutomationRegistrar21AutoApproveAllowedSenderSet struct {
	SenderAddress common.Address
	Allowed       bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAutoApproveAllowedSenderSet is a free log retrieval operation binding the contract event 0x20c6237dac83526a849285a9f79d08a483291bdd3a056a0ef9ae94ecee1ad356.
//
// Solidity: event AutoApproveAllowedSenderSet(address indexed senderAddress, bool allowed)
func (_AutomationRegistrar21 *AutomationRegistrar21Filterer) FilterAutoApproveAllowedSenderSet(opts *bind.FilterOpts, senderAddress []common.Address) (*AutomationRegistrar21AutoApproveAllowedSenderSetIterator, error) {

	var senderAddressRule []interface{}
	for _, senderAddressItem := range senderAddress {
		senderAddressRule = append(senderAddressRule, senderAddressItem)
	}

	logs, sub, err := _AutomationRegistrar21.contract.FilterLogs(opts, "AutoApproveAllowedSenderSet", senderAddressRule)
	if err != nil {
		return nil, err
	}
	return &AutomationRegistrar21AutoApproveAllowedSenderSetIterator{contract: _AutomationRegistrar21.contract, event: "AutoApproveAllowedSenderSet", logs: logs, sub: sub}, nil
}

// WatchAutoApproveAllowedSenderSet is a free log subscription operation binding the contract event 0x20c6237dac83526a849285a9f79d08a483291bdd3a056a0ef9ae94ecee1ad356.
//
// Solidity: event AutoApproveAllowedSenderSet(address indexed senderAddress, bool allowed)
func (_AutomationRegistrar21 *AutomationRegistrar21Filterer) WatchAutoApproveAllowedSenderSet(opts *bind.WatchOpts, sink chan<- *AutomationRegistrar21AutoApproveAllowedSenderSet, senderAddress []common.Address) (event.Subscription, error) {

	var senderAddressRule []interface{}
	for _, senderAddressItem := range senderAddress {
		senderAddressRule = append(senderAddressRule, senderAddressItem)
	}

	logs, sub, err := _AutomationRegistrar21.contract.WatchLogs(opts, "AutoApproveAllowedSenderSet", senderAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AutomationRegistrar21AutoApproveAllowedSenderSet)
				if err := _AutomationRegistrar21.contract.UnpackLog(event, "AutoApproveAllowedSenderSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAutoApproveAllowedSenderSet is a log parse operation binding the contract event 0x20c6237dac83526a849285a9f79d08a483291bdd3a056a0ef9ae94ecee1ad356.
//
// Solidity: event AutoApproveAllowedSenderSet(address indexed senderAddress, bool allowed)
func (_AutomationRegistrar21 *AutomationRegistrar21Filterer) ParseAutoApproveAllowedSenderSet(log types.Log) (*AutomationRegistrar21AutoApproveAllowedSenderSet, error) {
	event := new(AutomationRegistrar21AutoApproveAllowedSenderSet)
	if err := _AutomationRegistrar21.contract.UnpackLog(event, "AutoApproveAllowedSenderSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AutomationRegistrar21ConfigChangedIterator is returned from FilterConfigChanged and is used to iterate over the raw logs and unpacked data for ConfigChanged events raised by the AutomationRegistrar21 contract.
type AutomationRegistrar21ConfigChangedIterator struct {
	Event *AutomationRegistrar21ConfigChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AutomationRegistrar21ConfigChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AutomationRegistrar21ConfigChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AutomationRegistrar21ConfigChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AutomationRegistrar21ConfigChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AutomationRegistrar21ConfigChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AutomationRegistrar21ConfigChanged represents a ConfigChanged event raised by the AutomationRegistrar21 contract.
type AutomationRegistrar21ConfigChanged struct {
	KeeperRegistry common.Address
	MinLINKJuels   *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterConfigChanged is a free log retrieval operation binding the contract event 0x39ce5d867555f0b0183e358fce5b158e7ca4fecd7c01cb7e0e19f1e23285838a.
//
// Solidity: event ConfigChanged(address keeperRegistry, uint96 minLINKJuels)
func (_AutomationRegistrar21 *AutomationRegistrar21Filterer) FilterConfigChanged(opts *bind.FilterOpts) (*AutomationRegistrar21ConfigChangedIterator, error) {

	logs, sub, err := _AutomationRegistrar21.contract.FilterLogs(opts, "ConfigChanged")
	if err != nil {
		return nil, err
	}
	return &AutomationRegistrar21ConfigChangedIterator{contract: _AutomationRegistrar21.contract, event: "ConfigChanged", logs: logs, sub: sub}, nil
}

// WatchConfigChanged is a free log subscription operation binding the contract event 0x39ce5d867555f0b0183e358fce5b158e7ca4fecd7c01cb7e0e19f1e23285838a.
//
// Solidity: event ConfigChanged(address keeperRegistry, uint96 minLINKJuels)
func (_AutomationRegistrar21 *AutomationRegistrar21Filterer) WatchConfigChanged(opts *bind.WatchOpts, sink chan<- *AutomationRegistrar21ConfigChanged) (event.Subscription, error) {

	logs, sub, err := _AutomationRegistrar21.contract.WatchLogs(opts, "ConfigChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AutomationRegistrar21ConfigChanged)
				if err := _AutomationRegistrar21.contract.UnpackLog(event, "ConfigChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseConfigChanged is a log parse operation binding the contract event 0x39ce5d867555f0b0183e358fce5b158e7ca4fecd7c01cb7e0e19f1e23285838a.
//
// Solidity: event ConfigChanged(address keeperRegistry, uint96 minLINKJuels)
func (_AutomationRegistrar21 *AutomationRegistrar21Filterer) ParseConfigChanged(log types.Log) (*AutomationRegistrar21ConfigChanged, error) {
	event := new(AutomationRegistrar21ConfigChanged)
	if err := _AutomationRegistrar21.contract.UnpackLog(event, "ConfigChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AutomationRegistrar21OwnershipTransferRequestedIterator is returned from FilterOwnershipTransferRequested and is used to iterate over the raw logs and unpacked data for OwnershipTransferRequested events raised by the AutomationRegistrar21 contract.
type AutomationRegistrar21OwnershipTransferRequestedIterator struct {
	Event *AutomationRegistrar21OwnershipTransferRequested // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AutomationRegistrar21OwnershipTransferRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AutomationRegistrar21OwnershipTransferRequested)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AutomationRegistrar21OwnershipTransferRequested)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AutomationRegistrar21OwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AutomationRegistrar21OwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AutomationRegistrar21OwnershipTransferRequested represents a OwnershipTransferRequested event raised by the AutomationRegistrar21 contract.
type AutomationRegistrar21OwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferRequested is a free log retrieval operation binding the contract event 0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278.
//
// Solidity: event OwnershipTransferRequested(address indexed from, address indexed to)
func (_AutomationRegistrar21 *AutomationRegistrar21Filterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*AutomationRegistrar21OwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AutomationRegistrar21.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &AutomationRegistrar21OwnershipTransferRequestedIterator{contract: _AutomationRegistrar21.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferRequested is a free log subscription operation binding the contract event 0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278.
//
// Solidity: event OwnershipTransferRequested(address indexed from, address indexed to)
func (_AutomationRegistrar21 *AutomationRegistrar21Filterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *AutomationRegistrar21OwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AutomationRegistrar21.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AutomationRegistrar21OwnershipTransferRequested)
				if err := _AutomationRegistrar21.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferRequested is a log parse operation binding the contract event 0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278.
//
// Solidity: event OwnershipTransferRequested(address indexed from, address indexed to)
func (_AutomationRegistrar21 *AutomationRegistrar21Filterer) ParseOwnershipTransferRequested(log types.Log) (*AutomationRegistrar21OwnershipTransferRequested, error) {
	event := new(AutomationRegistrar21OwnershipTransferRequested)
	if err := _AutomationRegistrar21.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AutomationRegistrar21OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the AutomationRegistrar21 contract.
type AutomationRegistrar21OwnershipTransferredIterator struct {
	Event *AutomationRegistrar21OwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AutomationRegistrar21OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AutomationRegistrar21OwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AutomationRegistrar21OwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AutomationRegistrar21OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AutomationRegistrar21OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AutomationRegistrar21OwnershipTransferred represents a OwnershipTransferred event raised by the AutomationRegistrar21 contract.
type AutomationRegistrar21OwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed from, address indexed to)
func (_AutomationRegistrar21 *AutomationRegistrar21Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*AutomationRegistrar21OwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AutomationRegistrar21.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &AutomationRegistrar21OwnershipTransferredIterator{contract: _AutomationRegistrar21.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed from, address indexed to)
func (_AutomationRegistrar21 *AutomationRegistrar21Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AutomationRegistrar21OwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AutomationRegistrar21.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AutomationRegistrar21OwnershipTransferred)
				if err := _AutomationRegistrar21.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed from, address indexed to)
func (_AutomationRegistrar21 *AutomationRegistrar21Filterer) ParseOwnershipTransferred(log types.Log) (*AutomationRegistrar21OwnershipTransferred, error) {
	event := new(AutomationRegistrar21OwnershipTransferred)
	if err := _AutomationRegistrar21.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AutomationRegistrar21RegistrationApprovedIterator is returned from FilterRegistrationApproved and is used to iterate over the raw logs and unpacked data for RegistrationApproved events raised by the AutomationRegistrar21 contract.
type AutomationRegistrar21RegistrationApprovedIterator struct {
	Event *AutomationRegistrar21RegistrationApproved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AutomationRegistrar21RegistrationApprovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AutomationRegistrar21RegistrationApproved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AutomationRegistrar21RegistrationApproved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AutomationRegistrar21RegistrationApprovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AutomationRegistrar21RegistrationApprovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AutomationRegistrar21RegistrationApproved represents a RegistrationApproved event raised by the AutomationRegistrar21 contract.
type AutomationRegistrar21RegistrationApproved struct {
	Hash        [32]byte
	DisplayName string
	UpkeepId    *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRegistrationApproved is a free log retrieval operation binding the contract event 0xb9a292fb7e3edd920cd2d2829a3615a640c43fd7de0a0820aa0668feb4c37d4b.
//
// Solidity: event RegistrationApproved(bytes32 indexed hash, string displayName, uint256 indexed upkeepId)
func (_AutomationRegistrar21 *AutomationRegistrar21Filterer) FilterRegistrationApproved(opts *bind.FilterOpts, hash [][32]byte, upkeepId []*big.Int) (*AutomationRegistrar21RegistrationApprovedIterator, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}

	var upkeepIdRule []interface{}
	for _, upkeepIdItem := range upkeepId {
		upkeepIdRule = append(upkeepIdRule, upkeepIdItem)
	}

	logs, sub, err := _AutomationRegistrar21.contract.FilterLogs(opts, "RegistrationApproved", hashRule, upkeepIdRule)
	if err != nil {
		return nil, err
	}
	return &AutomationRegistrar21RegistrationApprovedIterator{contract: _AutomationRegistrar21.contract, event: "RegistrationApproved", logs: logs, sub: sub}, nil
}

// WatchRegistrationApproved is a free log subscription operation binding the contract event 0xb9a292fb7e3edd920cd2d2829a3615a640c43fd7de0a0820aa0668feb4c37d4b.
//
// Solidity: event RegistrationApproved(bytes32 indexed hash, string displayName, uint256 indexed upkeepId)
func (_AutomationRegistrar21 *AutomationRegistrar21Filterer) WatchRegistrationApproved(opts *bind.WatchOpts, sink chan<- *AutomationRegistrar21RegistrationApproved, hash [][32]byte, upkeepId []*big.Int) (event.Subscription, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}

	var upkeepIdRule []interface{}
	for _, upkeepIdItem := range upkeepId {
		upkeepIdRule = append(upkeepIdRule, upkeepIdItem)
	}

	logs, sub, err := _AutomationRegistrar21.contract.WatchLogs(opts, "RegistrationApproved", hashRule, upkeepIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AutomationRegistrar21RegistrationApproved)
				if err := _AutomationRegistrar21.contract.UnpackLog(event, "RegistrationApproved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRegistrationApproved is a log parse operation binding the contract event 0xb9a292fb7e3edd920cd2d2829a3615a640c43fd7de0a0820aa0668feb4c37d4b.
//
// Solidity: event RegistrationApproved(bytes32 indexed hash, string displayName, uint256 indexed upkeepId)
func (_AutomationRegistrar21 *AutomationRegistrar21Filterer) ParseRegistrationApproved(log types.Log) (*AutomationRegistrar21RegistrationApproved, error) {
	event := new(AutomationRegistrar21RegistrationApproved)
	if err := _AutomationRegistrar21.contract.UnpackLog(event, "RegistrationApproved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AutomationRegistrar21RegistrationRejectedIterator is returned from FilterRegistrationRejected and is used to iterate over the raw logs and unpacked data for RegistrationRejected events raised by the AutomationRegistrar21 contract.
type AutomationRegistrar21RegistrationRejectedIterator struct {
	Event *AutomationRegistrar21RegistrationRejected // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AutomationRegistrar21RegistrationRejectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AutomationRegistrar21RegistrationRejected)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AutomationRegistrar21RegistrationRejected)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AutomationRegistrar21RegistrationRejectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AutomationRegistrar21RegistrationRejectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AutomationRegistrar21RegistrationRejected represents a RegistrationRejected event raised by the AutomationRegistrar21 contract.
type AutomationRegistrar21RegistrationRejected struct {
	Hash [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterRegistrationRejected is a free log retrieval operation binding the contract event 0x3663fb28ebc87645eb972c9dad8521bf665c623f287e79f1c56f1eb374b82a22.
//
// Solidity: event RegistrationRejected(bytes32 indexed hash)
func (_AutomationRegistrar21 *AutomationRegistrar21Filterer) FilterRegistrationRejected(opts *bind.FilterOpts, hash [][32]byte) (*AutomationRegistrar21RegistrationRejectedIterator, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}

	logs, sub, err := _AutomationRegistrar21.contract.FilterLogs(opts, "RegistrationRejected", hashRule)
	if err != nil {
		return nil, err
	}
	return &AutomationRegistrar21RegistrationRejectedIterator{contract: _AutomationRegistrar21.contract, event: "RegistrationRejected", logs: logs, sub: sub}, nil
}

// WatchRegistrationRejected is a free log subscription operation binding the contract event 0x3663fb28ebc87645eb972c9dad8521bf665c623f287e79f1c56f1eb374b82a22.
//
// Solidity: event RegistrationRejected(bytes32 indexed hash)
func (_AutomationRegistrar21 *AutomationRegistrar21Filterer) WatchRegistrationRejected(opts *bind.WatchOpts, sink chan<- *AutomationRegistrar21RegistrationRejected, hash [][32]byte) (event.Subscription, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}

	logs, sub, err := _AutomationRegistrar21.contract.WatchLogs(opts, "RegistrationRejected", hashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AutomationRegistrar21RegistrationRejected)
				if err := _AutomationRegistrar21.contract.UnpackLog(event, "RegistrationRejected", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRegistrationRejected is a log parse operation binding the contract event 0x3663fb28ebc87645eb972c9dad8521bf665c623f287e79f1c56f1eb374b82a22.
//
// Solidity: event RegistrationRejected(bytes32 indexed hash)
func (_AutomationRegistrar21 *AutomationRegistrar21Filterer) ParseRegistrationRejected(log types.Log) (*AutomationRegistrar21RegistrationRejected, error) {
	event := new(AutomationRegistrar21RegistrationRejected)
	if err := _AutomationRegistrar21.contract.UnpackLog(event, "RegistrationRejected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AutomationRegistrar21RegistrationRequestedIterator is returned from FilterRegistrationRequested and is used to iterate over the raw logs and unpacked data for RegistrationRequested events raised by the AutomationRegistrar21 contract.
type AutomationRegistrar21RegistrationRequestedIterator struct {
	Event *AutomationRegistrar21RegistrationRequested // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AutomationRegistrar21RegistrationRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AutomationRegistrar21RegistrationRequested)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AutomationRegistrar21RegistrationRequested)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AutomationRegistrar21RegistrationRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AutomationRegistrar21RegistrationRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AutomationRegistrar21RegistrationRequested represents a RegistrationRequested event raised by the AutomationRegistrar21 contract.
type AutomationRegistrar21RegistrationRequested struct {
	Hash           [32]byte
	Name           string
	EncryptedEmail []byte
	UpkeepContract common.Address
	GasLimit       uint32
	AdminAddress   common.Address
	TriggerType    uint8
	TriggerConfig  []byte
	OffchainConfig []byte
	CheckData      []byte
	Amount         *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterRegistrationRequested is a free log retrieval operation binding the contract event 0x7684390ebb103102f7f48c71439c2408713f8d437782a6fab2756acc0e42c1b7.
//
// Solidity: event RegistrationRequested(bytes32 indexed hash, string name, bytes encryptedEmail, address indexed upkeepContract, uint32 gasLimit, address adminAddress, uint8 triggerType, bytes triggerConfig, bytes offchainConfig, bytes checkData, uint96 amount)
func (_AutomationRegistrar21 *AutomationRegistrar21Filterer) FilterRegistrationRequested(opts *bind.FilterOpts, hash [][32]byte, upkeepContract []common.Address) (*AutomationRegistrar21RegistrationRequestedIterator, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}

	var upkeepContractRule []interface{}
	for _, upkeepContractItem := range upkeepContract {
		upkeepContractRule = append(upkeepContractRule, upkeepContractItem)
	}

	logs, sub, err := _AutomationRegistrar21.contract.FilterLogs(opts, "RegistrationRequested", hashRule, upkeepContractRule)
	if err != nil {
		return nil, err
	}
	return &AutomationRegistrar21RegistrationRequestedIterator{contract: _AutomationRegistrar21.contract, event: "RegistrationRequested", logs: logs, sub: sub}, nil
}

// WatchRegistrationRequested is a free log subscription operation binding the contract event 0x7684390ebb103102f7f48c71439c2408713f8d437782a6fab2756acc0e42c1b7.
//
// Solidity: event RegistrationRequested(bytes32 indexed hash, string name, bytes encryptedEmail, address indexed upkeepContract, uint32 gasLimit, address adminAddress, uint8 triggerType, bytes triggerConfig, bytes offchainConfig, bytes checkData, uint96 amount)
func (_AutomationRegistrar21 *AutomationRegistrar21Filterer) WatchRegistrationRequested(opts *bind.WatchOpts, sink chan<- *AutomationRegistrar21RegistrationRequested, hash [][32]byte, upkeepContract []common.Address) (event.Subscription, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}

	var upkeepContractRule []interface{}
	for _, upkeepContractItem := range upkeepContract {
		upkeepContractRule = append(upkeepContractRule, upkeepContractItem)
	}

	logs, sub, err := _AutomationRegistrar21.contract.WatchLogs(opts, "RegistrationRequested", hashRule, upkeepContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AutomationRegistrar21RegistrationRequested)
				if err := _AutomationRegistrar21.contract.UnpackLog(event, "RegistrationRequested", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRegistrationRequested is a log parse operation binding the contract event 0x7684390ebb103102f7f48c71439c2408713f8d437782a6fab2756acc0e42c1b7.
//
// Solidity: event RegistrationRequested(bytes32 indexed hash, string name, bytes encryptedEmail, address indexed upkeepContract, uint32 gasLimit, address adminAddress, uint8 triggerType, bytes triggerConfig, bytes offchainConfig, bytes checkData, uint96 amount)
func (_AutomationRegistrar21 *AutomationRegistrar21Filterer) ParseRegistrationRequested(log types.Log) (*AutomationRegistrar21RegistrationRequested, error) {
	event := new(AutomationRegistrar21RegistrationRequested)
	if err := _AutomationRegistrar21.contract.UnpackLog(event, "RegistrationRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AutomationRegistrar21TriggerConfigSetIterator is returned from FilterTriggerConfigSet and is used to iterate over the raw logs and unpacked data for TriggerConfigSet events raised by the AutomationRegistrar21 contract.
type AutomationRegistrar21TriggerConfigSetIterator struct {
	Event *AutomationRegistrar21TriggerConfigSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *AutomationRegistrar21TriggerConfigSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AutomationRegistrar21TriggerConfigSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(AutomationRegistrar21TriggerConfigSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *AutomationRegistrar21TriggerConfigSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AutomationRegistrar21TriggerConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AutomationRegistrar21TriggerConfigSet represents a TriggerConfigSet event raised by the AutomationRegistrar21 contract.
type AutomationRegistrar21TriggerConfigSet struct {
	TriggerType           uint8
	AutoApproveType       uint8
	AutoApproveMaxAllowed uint32
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterTriggerConfigSet is a free log retrieval operation binding the contract event 0x830a6d06a4e2caac67eba04323de22bdb04f032dd8b3d6a0c52b503d9a7036a3.
//
// Solidity: event TriggerConfigSet(uint8 triggerType, uint8 autoApproveType, uint32 autoApproveMaxAllowed)
func (_AutomationRegistrar21 *AutomationRegistrar21Filterer) FilterTriggerConfigSet(opts *bind.FilterOpts) (*AutomationRegistrar21TriggerConfigSetIterator, error) {

	logs, sub, err := _AutomationRegistrar21.contract.FilterLogs(opts, "TriggerConfigSet")
	if err != nil {
		return nil, err
	}
	return &AutomationRegistrar21TriggerConfigSetIterator{contract: _AutomationRegistrar21.contract, event: "TriggerConfigSet", logs: logs, sub: sub}, nil
}

// WatchTriggerConfigSet is a free log subscription operation binding the contract event 0x830a6d06a4e2caac67eba04323de22bdb04f032dd8b3d6a0c52b503d9a7036a3.
//
// Solidity: event TriggerConfigSet(uint8 triggerType, uint8 autoApproveType, uint32 autoApproveMaxAllowed)
func (_AutomationRegistrar21 *AutomationRegistrar21Filterer) WatchTriggerConfigSet(opts *bind.WatchOpts, sink chan<- *AutomationRegistrar21TriggerConfigSet) (event.Subscription, error) {

	logs, sub, err := _AutomationRegistrar21.contract.WatchLogs(opts, "TriggerConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AutomationRegistrar21TriggerConfigSet)
				if err := _AutomationRegistrar21.contract.UnpackLog(event, "TriggerConfigSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTriggerConfigSet is a log parse operation binding the contract event 0x830a6d06a4e2caac67eba04323de22bdb04f032dd8b3d6a0c52b503d9a7036a3.
//
// Solidity: event TriggerConfigSet(uint8 triggerType, uint8 autoApproveType, uint32 autoApproveMaxAllowed)
func (_AutomationRegistrar21 *AutomationRegistrar21Filterer) ParseTriggerConfigSet(log types.Log) (*AutomationRegistrar21TriggerConfigSet, error) {
	event := new(AutomationRegistrar21TriggerConfigSet)
	if err := _AutomationRegistrar21.contract.UnpackLog(event, "TriggerConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
