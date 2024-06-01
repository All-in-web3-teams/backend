// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package automation

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

// AutomationMetaData contains all meta data concerning the Automation contract.
var AutomationMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseApproval\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseApproval\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"transferAndCall\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"typeAndVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50600436106100f55760003560e01c80634000aea011610097578063a457c2d711610066578063a457c2d7146103bc578063a9059cbb146103f5578063d73dd6231461042e578063dd62ed3e14610467576100f5565b80634000aea014610280578063661884631461034857806370a082311461038157806395d89b41146103b4576100f5565b8063181f5a77116100d3578063181f5a77146101de57806323b872dd146101e6578063313ce567146102295780633950935114610247576100f5565b806306fdde03146100fa578063095ea7b31461017757806318160ddd146101c4575b600080fd5b6101026104a2565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561013c578181015183820152602001610124565b50505050905090810190601f1680156101695780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6101b06004803603604081101561018d57600080fd5b5073ffffffffffffffffffffffffffffffffffffffff8135169060200135610556565b604080519115158252519081900360200190f35b6101cc610573565b60408051918252519081900360200190f35b610102610579565b6101b0600480360360608110156101fc57600080fd5b5073ffffffffffffffffffffffffffffffffffffffff8135811691602081013590911690604001356105b0565b610231610651565b6040805160ff9092168252519081900360200190f35b6101b06004803603604081101561025d57600080fd5b5073ffffffffffffffffffffffffffffffffffffffff813516906020013561065a565b6101b06004803603606081101561029657600080fd5b73ffffffffffffffffffffffffffffffffffffffff823516916020810135918101906060810160408201356401000000008111156102d357600080fd5b8201836020820111156102e557600080fd5b8035906020019184600183028401116401000000008311171561030757600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506106b5945050505050565b6101b06004803603604081101561035e57600080fd5b5073ffffffffffffffffffffffffffffffffffffffff81351690602001356107aa565b6101cc6004803603602081101561039757600080fd5b503573ffffffffffffffffffffffffffffffffffffffff166107bd565b6101026107e5565b6101b0600480360360408110156103d257600080fd5b5073ffffffffffffffffffffffffffffffffffffffff8135169060200135610864565b6101b06004803603604081101561040b57600080fd5b5073ffffffffffffffffffffffffffffffffffffffff81351690602001356108d9565b6101b06004803603604081101561044457600080fd5b5073ffffffffffffffffffffffffffffffffffffffff81351690602001356108ed565b6101cc6004803603604081101561047d57600080fd5b5073ffffffffffffffffffffffffffffffffffffffff813581169160200135166108f9565b60038054604080516020601f60027fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff61010060018816150201909516949094049384018190048102820181019092528281526060939092909183018282801561054c5780601f106105215761010080835404028352916020019161054c565b820191906000526020600020905b81548152906001019060200180831161052f57829003601f168201915b5050505050905090565b600061056a6105636109a5565b84846109a9565b50600192915050565b60025490565b60408051808201909152600f81527f4c696e6b546f6b656e20302e302e330000000000000000000000000000000000602082015290565b60006105bd848484610a2a565b610647846105c96109a5565b610642856040518060600160405280602881526020016110186028913973ffffffffffffffffffffffffffffffffffffffff8a166000908152600160205260408120906106146109a5565b73ffffffffffffffffffffffffffffffffffffffff1681526020810191909152604001600020549190610aa5565b6109a9565b5060019392505050565b60055460ff1690565b600061056a6106676109a5565b8461064285600160006106786109a5565b73ffffffffffffffffffffffffffffffffffffffff908116825260208083019390935260409182016000908120918c168152925290205490610931565b60006106c184846108d9565b508373ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167fe19260aff97b920c7df27010903aeb9c8d2be5d310a2c67824cf3f15396e4c1685856040518083815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561075657818101518382015260200161073e565b50505050905090810190601f1680156107835780820380516001836020036101000a031916815260200191505b50935050505060405180910390a361079a84610b56565b1561064757610647848484610b5c565b60006107b68383610864565b9392505050565b73ffffffffffffffffffffffffffffffffffffffff1660009081526020819052604090205490565b60048054604080516020601f60027fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff61010060018816150201909516949094049384018190048102820181019092528281526060939092909183018282801561054c5780601f106105215761010080835404028352916020019161054c565b600061056a6108716109a5565b8461064285604051806060016040528060258152602001611089602591396001600061089b6109a5565b73ffffffffffffffffffffffffffffffffffffffff908116825260208083019390935260409182016000908120918d16815292529020549190610aa5565b600061056a6108e66109a5565b8484610a2a565b60006107b6838361065a565b73ffffffffffffffffffffffffffffffffffffffff918216600090815260016020908152604080832093909416825291909152205490565b6000828201838110156107b657604080517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b3390565b8173ffffffffffffffffffffffffffffffffffffffff8116301415610a19576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526034815260200180610fbe6034913960400191505060405180910390fd5b610a24848484610c5c565b50505050565b8173ffffffffffffffffffffffffffffffffffffffff8116301415610a9a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526034815260200180610fbe6034913960400191505060405180910390fd5b610a24848484610da3565b60008184841115610b4e576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b83811015610b13578181015183820152602001610afb565b50505050905090810190601f168015610b405780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b505050900390565b3b151590565b6040517fa4c0ed36000000000000000000000000000000000000000000000000000000008152336004820181815260248301859052606060448401908152845160648501528451879473ffffffffffffffffffffffffffffffffffffffff86169463a4c0ed369490938993899360840190602085019080838360005b83811015610bf0578181015183820152602001610bd8565b50505050905090810190601f168015610c1d5780820380516001836020036101000a031916815260200191505b50945050505050600060405180830381600087803b158015610c3e57600080fd5b505af1158015610c52573d6000803e3d6000fd5b5050505050505050565b73ffffffffffffffffffffffffffffffffffffffff8316610cc8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260248152602001806110656024913960400191505060405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8216610d34576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526022815260200180610f9c6022913960400191505060405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff808416600081815260016020908152604080832094871680845294825291829020859055815185815291517f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259281900390910190a3505050565b73ffffffffffffffffffffffffffffffffffffffff8316610e0f576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260258152602001806110406025913960400191505060405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff8216610e7b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526023815260200180610f796023913960400191505060405180910390fd5b610e86838383610f73565b610ed081604051806060016040528060268152602001610ff26026913973ffffffffffffffffffffffffffffffffffffffff86166000908152602081905260409020549190610aa5565b73ffffffffffffffffffffffffffffffffffffffff8085166000908152602081905260408082209390935590841681522054610f0c9082610931565b73ffffffffffffffffffffffffffffffffffffffff8084166000818152602081815260409182902094909455805185815290519193928716927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef92918290030190a3505050565b50505056fe45524332303a207472616e7366657220746f20746865207a65726f206164647265737345524332303a20617070726f766520746f20746865207a65726f20616464726573734c696e6b546f6b656e3a207472616e736665722f617070726f766520746f207468697320636f6e7472616374206164647265737345524332303a207472616e7366657220616d6f756e7420657863656564732062616c616e636545524332303a207472616e7366657220616d6f756e74206578636565647320616c6c6f77616e636545524332303a207472616e736665722066726f6d20746865207a65726f206164647265737345524332303a20617070726f76652066726f6d20746865207a65726f206164647265737345524332303a2064656372656173656420616c6c6f77616e63652062656c6f77207a65726fa164736f6c6343000706000a",
}

// AutomationABI is the input ABI used to generate the binding from.
// Deprecated: Use AutomationMetaData.ABI instead.
var AutomationABI = AutomationMetaData.ABI

// AutomationBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AutomationMetaData.Bin instead.
var AutomationBin = AutomationMetaData.Bin

// DeployAutomation deploys a new Ethereum contract, binding an instance of Automation to it.
func DeployAutomation(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Automation, error) {
	parsed, err := AutomationMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AutomationBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Automation{AutomationCaller: AutomationCaller{contract: contract}, AutomationTransactor: AutomationTransactor{contract: contract}, AutomationFilterer: AutomationFilterer{contract: contract}}, nil
}

// Automation is an auto generated Go binding around an Ethereum contract.
type Automation struct {
	AutomationCaller     // Read-only binding to the contract
	AutomationTransactor // Write-only binding to the contract
	AutomationFilterer   // Log filterer for contract events
}

// AutomationCaller is an auto generated read-only Go binding around an Ethereum contract.
type AutomationCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AutomationTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AutomationTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AutomationFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AutomationFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AutomationSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AutomationSession struct {
	Contract     *Automation       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AutomationCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AutomationCallerSession struct {
	Contract *AutomationCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// AutomationTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AutomationTransactorSession struct {
	Contract     *AutomationTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// AutomationRaw is an auto generated low-level Go binding around an Ethereum contract.
type AutomationRaw struct {
	Contract *Automation // Generic contract binding to access the raw methods on
}

// AutomationCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AutomationCallerRaw struct {
	Contract *AutomationCaller // Generic read-only contract binding to access the raw methods on
}

// AutomationTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AutomationTransactorRaw struct {
	Contract *AutomationTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAutomation creates a new instance of Automation, bound to a specific deployed contract.
func NewAutomation(address common.Address, backend bind.ContractBackend) (*Automation, error) {
	contract, err := bindAutomation(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Automation{AutomationCaller: AutomationCaller{contract: contract}, AutomationTransactor: AutomationTransactor{contract: contract}, AutomationFilterer: AutomationFilterer{contract: contract}}, nil
}

// NewAutomationCaller creates a new read-only instance of Automation, bound to a specific deployed contract.
func NewAutomationCaller(address common.Address, caller bind.ContractCaller) (*AutomationCaller, error) {
	contract, err := bindAutomation(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AutomationCaller{contract: contract}, nil
}

// NewAutomationTransactor creates a new write-only instance of Automation, bound to a specific deployed contract.
func NewAutomationTransactor(address common.Address, transactor bind.ContractTransactor) (*AutomationTransactor, error) {
	contract, err := bindAutomation(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AutomationTransactor{contract: contract}, nil
}

// NewAutomationFilterer creates a new log filterer instance of Automation, bound to a specific deployed contract.
func NewAutomationFilterer(address common.Address, filterer bind.ContractFilterer) (*AutomationFilterer, error) {
	contract, err := bindAutomation(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AutomationFilterer{contract: contract}, nil
}

// bindAutomation binds a generic wrapper to an already deployed contract.
func bindAutomation(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AutomationMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Automation *AutomationRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Automation.Contract.AutomationCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Automation *AutomationRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Automation.Contract.AutomationTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Automation *AutomationRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Automation.Contract.AutomationTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Automation *AutomationCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Automation.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Automation *AutomationTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Automation.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Automation *AutomationTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Automation.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Automation *AutomationCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Automation.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Automation *AutomationSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Automation.Contract.Allowance(&_Automation.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Automation *AutomationCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Automation.Contract.Allowance(&_Automation.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Automation *AutomationCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Automation.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Automation *AutomationSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Automation.Contract.BalanceOf(&_Automation.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Automation *AutomationCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Automation.Contract.BalanceOf(&_Automation.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Automation *AutomationCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Automation.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Automation *AutomationSession) Decimals() (uint8, error) {
	return _Automation.Contract.Decimals(&_Automation.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Automation *AutomationCallerSession) Decimals() (uint8, error) {
	return _Automation.Contract.Decimals(&_Automation.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Automation *AutomationCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Automation.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Automation *AutomationSession) Name() (string, error) {
	return _Automation.Contract.Name(&_Automation.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Automation *AutomationCallerSession) Name() (string, error) {
	return _Automation.Contract.Name(&_Automation.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Automation *AutomationCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Automation.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Automation *AutomationSession) Symbol() (string, error) {
	return _Automation.Contract.Symbol(&_Automation.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Automation *AutomationCallerSession) Symbol() (string, error) {
	return _Automation.Contract.Symbol(&_Automation.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Automation *AutomationCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Automation.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Automation *AutomationSession) TotalSupply() (*big.Int, error) {
	return _Automation.Contract.TotalSupply(&_Automation.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Automation *AutomationCallerSession) TotalSupply() (*big.Int, error) {
	return _Automation.Contract.TotalSupply(&_Automation.CallOpts)
}

// TypeAndVersion is a free data retrieval call binding the contract method 0x181f5a77.
//
// Solidity: function typeAndVersion() pure returns(string)
func (_Automation *AutomationCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Automation.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TypeAndVersion is a free data retrieval call binding the contract method 0x181f5a77.
//
// Solidity: function typeAndVersion() pure returns(string)
func (_Automation *AutomationSession) TypeAndVersion() (string, error) {
	return _Automation.Contract.TypeAndVersion(&_Automation.CallOpts)
}

// TypeAndVersion is a free data retrieval call binding the contract method 0x181f5a77.
//
// Solidity: function typeAndVersion() pure returns(string)
func (_Automation *AutomationCallerSession) TypeAndVersion() (string, error) {
	return _Automation.Contract.TypeAndVersion(&_Automation.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Automation *AutomationTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Automation.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Automation *AutomationSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Automation.Contract.Approve(&_Automation.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Automation *AutomationTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Automation.Contract.Approve(&_Automation.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Automation *AutomationTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Automation.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Automation *AutomationSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Automation.Contract.DecreaseAllowance(&_Automation.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Automation *AutomationTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Automation.Contract.DecreaseAllowance(&_Automation.TransactOpts, spender, subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(address spender, uint256 subtractedValue) returns(bool)
func (_Automation *AutomationTransactor) DecreaseApproval(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Automation.contract.Transact(opts, "decreaseApproval", spender, subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(address spender, uint256 subtractedValue) returns(bool)
func (_Automation *AutomationSession) DecreaseApproval(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Automation.Contract.DecreaseApproval(&_Automation.TransactOpts, spender, subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(address spender, uint256 subtractedValue) returns(bool)
func (_Automation *AutomationTransactorSession) DecreaseApproval(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Automation.Contract.DecreaseApproval(&_Automation.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Automation *AutomationTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Automation.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Automation *AutomationSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Automation.Contract.IncreaseAllowance(&_Automation.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Automation *AutomationTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Automation.Contract.IncreaseAllowance(&_Automation.TransactOpts, spender, addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(address spender, uint256 addedValue) returns(bool)
func (_Automation *AutomationTransactor) IncreaseApproval(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Automation.contract.Transact(opts, "increaseApproval", spender, addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(address spender, uint256 addedValue) returns(bool)
func (_Automation *AutomationSession) IncreaseApproval(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Automation.Contract.IncreaseApproval(&_Automation.TransactOpts, spender, addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(address spender, uint256 addedValue) returns(bool)
func (_Automation *AutomationTransactorSession) IncreaseApproval(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Automation.Contract.IncreaseApproval(&_Automation.TransactOpts, spender, addedValue)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Automation *AutomationTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Automation.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Automation *AutomationSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Automation.Contract.Transfer(&_Automation.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Automation *AutomationTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Automation.Contract.Transfer(&_Automation.TransactOpts, recipient, amount)
}

// TransferAndCall is a paid mutator transaction binding the contract method 0x4000aea0.
//
// Solidity: function transferAndCall(address to, uint256 value, bytes data) returns(bool success)
func (_Automation *AutomationTransactor) TransferAndCall(opts *bind.TransactOpts, to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _Automation.contract.Transact(opts, "transferAndCall", to, value, data)
}

// TransferAndCall is a paid mutator transaction binding the contract method 0x4000aea0.
//
// Solidity: function transferAndCall(address to, uint256 value, bytes data) returns(bool success)
func (_Automation *AutomationSession) TransferAndCall(to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _Automation.Contract.TransferAndCall(&_Automation.TransactOpts, to, value, data)
}

// TransferAndCall is a paid mutator transaction binding the contract method 0x4000aea0.
//
// Solidity: function transferAndCall(address to, uint256 value, bytes data) returns(bool success)
func (_Automation *AutomationTransactorSession) TransferAndCall(to common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _Automation.Contract.TransferAndCall(&_Automation.TransactOpts, to, value, data)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Automation *AutomationTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Automation.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Automation *AutomationSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Automation.Contract.TransferFrom(&_Automation.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Automation *AutomationTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Automation.Contract.TransferFrom(&_Automation.TransactOpts, sender, recipient, amount)
}

// AutomationApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Automation contract.
type AutomationApprovalIterator struct {
	Event *AutomationApproval // Event containing the contract specifics and raw log

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
func (it *AutomationApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AutomationApproval)
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
		it.Event = new(AutomationApproval)
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
func (it *AutomationApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AutomationApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AutomationApproval represents a Approval event raised by the Automation contract.
type AutomationApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Automation *AutomationFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*AutomationApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Automation.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &AutomationApprovalIterator{contract: _Automation.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Automation *AutomationFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *AutomationApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Automation.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AutomationApproval)
				if err := _Automation.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Automation *AutomationFilterer) ParseApproval(log types.Log) (*AutomationApproval, error) {
	event := new(AutomationApproval)
	if err := _Automation.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AutomationTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Automation contract.
type AutomationTransferIterator struct {
	Event *AutomationTransfer // Event containing the contract specifics and raw log

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
func (it *AutomationTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AutomationTransfer)
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
		it.Event = new(AutomationTransfer)
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
func (it *AutomationTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AutomationTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AutomationTransfer represents a Transfer event raised by the Automation contract.
type AutomationTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Data  []byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xe19260aff97b920c7df27010903aeb9c8d2be5d310a2c67824cf3f15396e4c16.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value, bytes data)
func (_Automation *AutomationFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*AutomationTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Automation.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &AutomationTransferIterator{contract: _Automation.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xe19260aff97b920c7df27010903aeb9c8d2be5d310a2c67824cf3f15396e4c16.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value, bytes data)
func (_Automation *AutomationFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *AutomationTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Automation.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AutomationTransfer)
				if err := _Automation.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xe19260aff97b920c7df27010903aeb9c8d2be5d310a2c67824cf3f15396e4c16.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value, bytes data)
func (_Automation *AutomationFilterer) ParseTransfer(log types.Log) (*AutomationTransfer, error) {
	event := new(AutomationTransfer)
	if err := _Automation.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AutomationTransfer0Iterator is returned from FilterTransfer0 and is used to iterate over the raw logs and unpacked data for Transfer0 events raised by the Automation contract.
type AutomationTransfer0Iterator struct {
	Event *AutomationTransfer0 // Event containing the contract specifics and raw log

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
func (it *AutomationTransfer0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AutomationTransfer0)
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
		it.Event = new(AutomationTransfer0)
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
func (it *AutomationTransfer0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AutomationTransfer0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AutomationTransfer0 represents a Transfer0 event raised by the Automation contract.
type AutomationTransfer0 struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer0 is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Automation *AutomationFilterer) FilterTransfer0(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*AutomationTransfer0Iterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Automation.contract.FilterLogs(opts, "Transfer0", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &AutomationTransfer0Iterator{contract: _Automation.contract, event: "Transfer0", logs: logs, sub: sub}, nil
}

// WatchTransfer0 is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Automation *AutomationFilterer) WatchTransfer0(opts *bind.WatchOpts, sink chan<- *AutomationTransfer0, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Automation.contract.WatchLogs(opts, "Transfer0", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AutomationTransfer0)
				if err := _Automation.contract.UnpackLog(event, "Transfer0", log); err != nil {
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

// ParseTransfer0 is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Automation *AutomationFilterer) ParseTransfer0(log types.Log) (*AutomationTransfer0, error) {
	event := new(AutomationTransfer0)
	if err := _Automation.contract.UnpackLog(event, "Transfer0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
