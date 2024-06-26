// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package vrfV2_5

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

// VRFProof is an auto generated low-level Go binding around an user-defined struct.
type VRFProof struct {
	Pk            [2]*big.Int
	Gamma         [2]*big.Int
	C             *big.Int
	S             *big.Int
	Seed          *big.Int
	UWitness      common.Address
	CGammaWitness [2]*big.Int
	SHashWitness  [2]*big.Int
	ZInv          *big.Int
}

// VRFTypesRequestCommitmentV2Plus is an auto generated low-level Go binding around an user-defined struct.
type VRFTypesRequestCommitmentV2Plus struct {
	BlockNum         uint64
	SubId            *big.Int
	CallbackGasLimit uint32
	NumWords         uint32
	Sender           common.Address
	ExtraArgs        []byte
}

// VRFV2PlusClientRandomWordsRequest is an auto generated low-level Go binding around an user-defined struct.
type VRFV2PlusClientRandomWordsRequest struct {
	KeyHash              [32]byte
	SubId                *big.Int
	RequestConfirmations uint16
	CallbackGasLimit     uint32
	NumWords             uint32
	ExtraArgs            []byte
}

// VrfV25MetaData contains all meta data concerning the VrfV25 contract.
var VrfV25MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"blockhashStore\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"internalBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"externalBalance\",\"type\":\"uint256\"}],\"name\":\"BalanceInvariantViolated\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNum\",\"type\":\"uint256\"}],\"name\":\"BlockhashNotInStore\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"coordinatorAddress\",\"type\":\"address\"}],\"name\":\"CoordinatorAlreadyRegistered\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"coordinatorAddress\",\"type\":\"address\"}],\"name\":\"CoordinatorNotRegistered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedToSendNative\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedToTransferLink\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"have\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"want\",\"type\":\"uint32\"}],\"name\":\"GasLimitTooBig\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxGas\",\"type\":\"uint256\"}],\"name\":\"GasPriceExceeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IncorrectCommitment\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IndexOutOfRange\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCalldata\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"consumer\",\"type\":\"address\"}],\"name\":\"InvalidConsumer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidExtraArgsTag\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"linkWei\",\"type\":\"int256\"}],\"name\":\"InvalidLinkWeiPrice\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"premiumPercentage\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"max\",\"type\":\"uint8\"}],\"name\":\"InvalidPremiumPercentage\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"have\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"min\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"max\",\"type\":\"uint16\"}],\"name\":\"InvalidRequestConfirmations\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidSubscription\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LinkAlreadySet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"flatFeeLinkDiscountPPM\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"flatFeeNativePPM\",\"type\":\"uint32\"}],\"name\":\"LinkDiscountTooHigh\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LinkNotSet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"have\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"max\",\"type\":\"uint32\"}],\"name\":\"MsgDataTooBig\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"proposedOwner\",\"type\":\"address\"}],\"name\":\"MustBeRequestedOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"MustBeSubOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoCorrespondingRequest\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"keyHash\",\"type\":\"bytes32\"}],\"name\":\"NoSuchProvingKey\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"have\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"want\",\"type\":\"uint32\"}],\"name\":\"NumWordsTooBig\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyCallableFromLink\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PaymentTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingRequestExists\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"keyHash\",\"type\":\"bytes32\"}],\"name\":\"ProvingKeyAlreadyRegistered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"Reentrant\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TooManyConsumers\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"minimumRequestConfirmations\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"maxGasLimit\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"stalenessSeconds\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"gasAfterPaymentCalculation\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"fallbackWeiPerUnitLink\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"fulfillmentFlatFeeNativePPM\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"fulfillmentFlatFeeLinkDiscountPPM\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"nativePremiumPercentage\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"linkPremiumPercentage\",\"type\":\"uint8\"}],\"name\":\"ConfigSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"coordinatorAddress\",\"type\":\"address\"}],\"name\":\"CoordinatorDeregistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"coordinatorAddress\",\"type\":\"address\"}],\"name\":\"CoordinatorRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"fallbackWeiPerUnitLink\",\"type\":\"int256\"}],\"name\":\"FallbackWeiPerUnitLinkUsed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"FundsRecovered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newCoordinator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"}],\"name\":\"MigrationCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"NativeFundsRecovered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"keyHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"maxGas\",\"type\":\"uint64\"}],\"name\":\"ProvingKeyDeregistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"keyHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"maxGas\",\"type\":\"uint64\"}],\"name\":\"ProvingKeyRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"outputSeed\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint96\",\"name\":\"payment\",\"type\":\"uint96\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"nativePayment\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"onlyPremium\",\"type\":\"bool\"}],\"name\":\"RandomWordsFulfilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"keyHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"preSeed\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"minimumRequestConfirmations\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"callbackGasLimit\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"numWords\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"extraArgs\",\"type\":\"bytes\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RandomWordsRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountLink\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountNative\",\"type\":\"uint256\"}],\"name\":\"SubscriptionCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"consumer\",\"type\":\"address\"}],\"name\":\"SubscriptionConsumerAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"consumer\",\"type\":\"address\"}],\"name\":\"SubscriptionConsumerRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"SubscriptionCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldBalance\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newBalance\",\"type\":\"uint256\"}],\"name\":\"SubscriptionFunded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldNativeBalance\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newNativeBalance\",\"type\":\"uint256\"}],\"name\":\"SubscriptionFundedWithNative\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"SubscriptionOwnerTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"SubscriptionOwnerTransferred\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BLOCKHASH_STORE\",\"outputs\":[{\"internalType\":\"contractBlockhashStoreInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LINK\",\"outputs\":[{\"internalType\":\"contractLinkTokenInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LINK_NATIVE_FEED\",\"outputs\":[{\"internalType\":\"contractAggregatorV3Interface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_CONSUMERS\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_NUM_WORDS\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_REQUEST_CONFIRMATIONS\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"}],\"name\":\"acceptSubscriptionOwnerTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"consumer\",\"type\":\"address\"}],\"name\":\"addConsumer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"cancelSubscription\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"createSubscription\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"deregisterMigratableCoordinator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"publicProvingKey\",\"type\":\"uint256[2]\"}],\"name\":\"deregisterProvingKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256[2]\",\"name\":\"pk\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"gamma\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256\",\"name\":\"c\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"s\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"seed\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"uWitness\",\"type\":\"address\"},{\"internalType\":\"uint256[2]\",\"name\":\"cGammaWitness\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"sHashWitness\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256\",\"name\":\"zInv\",\"type\":\"uint256\"}],\"internalType\":\"structVRF.Proof\",\"name\":\"proof\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"blockNum\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"callbackGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"numWords\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"extraArgs\",\"type\":\"bytes\"}],\"internalType\":\"structVRFTypes.RequestCommitmentV2Plus\",\"name\":\"rc\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"onlyPremium\",\"type\":\"bool\"}],\"name\":\"fulfillRandomWords\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"payment\",\"type\":\"uint96\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"}],\"name\":\"fundSubscriptionWithNative\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"startIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxCount\",\"type\":\"uint256\"}],\"name\":\"getActiveSubscriptionIds\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"}],\"name\":\"getSubscription\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"balance\",\"type\":\"uint96\"},{\"internalType\":\"uint96\",\"name\":\"nativeBalance\",\"type\":\"uint96\"},{\"internalType\":\"uint64\",\"name\":\"reqCount\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"subOwner\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"consumers\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"publicKey\",\"type\":\"uint256[2]\"}],\"name\":\"hashOfKey\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"newCoordinator\",\"type\":\"address\"}],\"name\":\"migrate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onTokenTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"}],\"name\":\"ownerCancelSubscription\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"}],\"name\":\"pendingRequestExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"recoverFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"recoverNativeFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"registerMigratableCoordinator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[2]\",\"name\":\"publicProvingKey\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint64\",\"name\":\"maxGas\",\"type\":\"uint64\"}],\"name\":\"registerProvingKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"consumer\",\"type\":\"address\"}],\"name\":\"removeConsumer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"keyHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"requestConfirmations\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"callbackGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"numWords\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"extraArgs\",\"type\":\"bytes\"}],\"internalType\":\"structVRFV2PlusClient.RandomWordsRequest\",\"name\":\"req\",\"type\":\"tuple\"}],\"name\":\"requestRandomWords\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"subId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"requestSubscriptionOwnerTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_config\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"minimumRequestConfirmations\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"maxGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"reentrancyLock\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"stalenessSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"gasAfterPaymentCalculation\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"fulfillmentFlatFeeNativePPM\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"fulfillmentFlatFeeLinkDiscountPPM\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"nativePremiumPercentage\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"linkPremiumPercentage\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_currentSubNonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_fallbackWeiPerUnitLink\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"s_provingKeyHashes\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"s_provingKeys\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"},{\"internalType\":\"uint64\",\"name\":\"maxGas\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"s_requestCommitments\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_totalBalance\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"\",\"type\":\"uint96\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"s_totalNativeBalance\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"\",\"type\":\"uint96\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"minimumRequestConfirmations\",\"type\":\"uint16\"},{\"internalType\":\"uint32\",\"name\":\"maxGasLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"stalenessSeconds\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"gasAfterPaymentCalculation\",\"type\":\"uint32\"},{\"internalType\":\"int256\",\"name\":\"fallbackWeiPerUnitLink\",\"type\":\"int256\"},{\"internalType\":\"uint32\",\"name\":\"fulfillmentFlatFeeNativePPM\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"fulfillmentFlatFeeLinkDiscountPPM\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"nativePremiumPercentage\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"linkPremiumPercentage\",\"type\":\"uint8\"}],\"name\":\"setConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"link\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"linkNativeFeed\",\"type\":\"address\"}],\"name\":\"setLINKAndLINKNativeFeed\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"withdrawNative\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040526004361061028c5760003560e01c80638402595e11610164578063b2a7cac5116100c6578063da2f26101161008a578063e72f6e3011610064578063e72f6e3014610904578063ee9d2d3814610924578063f2fde38b1461095157600080fd5b8063da2f261014610854578063dac83d29146108b3578063dc311dd3146108d357600080fd5b8063b2a7cac5146107b4578063bec4c08c146107d4578063caf70c4a146107f4578063cb63179714610814578063d98e620e1461083457600080fd5b80639d40a6fd11610128578063a63e0bfb11610102578063a63e0bfb14610747578063aa433aff14610767578063aefb212f1461078757600080fd5b80639d40a6fd146106da578063a21a23e414610712578063a4c0ed361461072757600080fd5b80638402595e1461064957806386fe91c7146106695780638da5cb5b1461068957806395b55cfc146106a75780639b1c385e146106ba57600080fd5b8063405b84fa1161020d57806364d51a2a116101d157806372e9d565116101ab57806372e9d565146105f457806379ba5097146106145780637a5a2aef1461062957600080fd5b806364d51a2a1461058b57806365982744146105a0578063689c4517146105c057600080fd5b8063405b84fa146104d057806340d6bb82146104f057806341af6c871461051b57806351cff8d91461054b5780635d06b4ab1461056b57600080fd5b806315c48b841161025457806315c48b84146103f157806318e3dd27146104195780631b6b6d23146104585780632f622e6b14610490578063301f42e9146104b057600080fd5b806304104edb14610291578063043bd6ae146102b3578063088070f5146102dc57806308821d58146103b15780630ae09540146103d1575b600080fd5b34801561029d57600080fd5b506102b16102ac3660046150af565b610971565b005b3480156102bf57600080fd5b506102c960105481565b6040519081526020015b60405180910390f35b3480156102e857600080fd5b50600c546103549061ffff81169063ffffffff62010000820481169160ff660100000000000082048116926701000000000000008304811692600160581b8104821692600160781b8204831692600160981b83041691600160b81b8104821691600160c01b9091041689565b6040805161ffff909a168a5263ffffffff98891660208b01529615159689019690965293861660608801529185166080870152841660a08601529290921660c084015260ff91821660e084015216610100820152610120016102d3565b3480156103bd57600080fd5b506102b16103cc3660046150dd565b610aea565b3480156103dd57600080fd5b506102b16103ec3660046150f9565b610ca7565b3480156103fd57600080fd5b5061040660c881565b60405161ffff90911681526020016102d3565b34801561042557600080fd5b50600a5461044090600160601b90046001600160601b031681565b6040516001600160601b0390911681526020016102d3565b34801561046457600080fd5b50600254610478906001600160a01b031681565b6040516001600160a01b0390911681526020016102d3565b34801561049c57600080fd5b506102b16104ab3660046150af565b610cef565b3480156104bc57600080fd5b506104406104cb36600461535b565b610e3e565b3480156104dc57600080fd5b506102b16104eb3660046150f9565b611154565b3480156104fc57600080fd5b506105066101f481565b60405163ffffffff90911681526020016102d3565b34801561052757600080fd5b5061053b610536366004615449565b611536565b60405190151581526020016102d3565b34801561055757600080fd5b506102b16105663660046150af565b6115ea565b34801561057757600080fd5b506102b16105863660046150af565b61176c565b34801561059757600080fd5b50610406606481565b3480156105ac57600080fd5b506102b16105bb366004615462565b61182a565b3480156105cc57600080fd5b506104787f0000000000000000000000001159e1889754a1f0862f8ec0e109f169aecbcd6f81565b34801561060057600080fd5b50600354610478906001600160a01b031681565b34801561062057600080fd5b506102b161188a565b34801561063557600080fd5b506102b1610644366004615490565b61193b565b34801561065557600080fd5b506102b16106643660046150af565b611a6f565b34801561067557600080fd5b50600a54610440906001600160601b031681565b34801561069557600080fd5b506000546001600160a01b0316610478565b6102b16106b5366004615449565b611b8a565b3480156106c657600080fd5b506102c96106d53660046154c4565b611cae565b3480156106e657600080fd5b506007546106fa906001600160401b031681565b6040516001600160401b0390911681526020016102d3565b34801561071e57600080fd5b506102c96120f4565b34801561073357600080fd5b506102b16107423660046154fe565b6122db565b34801561075357600080fd5b506102b16107623660046155a9565b612457565b34801561077357600080fd5b506102b1610782366004615449565b61273e565b34801561079357600080fd5b506107a76107a236600461564a565b612786565b6040516102d391906156a7565b3480156107c057600080fd5b506102b16107cf366004615449565b612888565b3480156107e057600080fd5b506102b16107ef3660046150f9565b61298c565b34801561080057600080fd5b506102c961080f3660046156ba565b612a7f565b34801561082057600080fd5b506102b161082f3660046150f9565b612aaf565b34801561084057600080fd5b506102c961084f366004615449565b612d1d565b34801561086057600080fd5b5061089461086f366004615449565b600d6020526000908152604090205460ff81169061010090046001600160401b031682565b6040805192151583526001600160401b039091166020830152016102d3565b3480156108bf57600080fd5b506102b16108ce3660046150f9565b612d3e565b3480156108df57600080fd5b506108f36108ee366004615449565b612dd8565b6040516102d395949392919061570f565b34801561091057600080fd5b506102b161091f3660046150af565b612ec6565b34801561093057600080fd5b506102c961093f366004615449565b600f6020526000908152604090205481565b34801561095d57600080fd5b506102b161096c3660046150af565b613087565b610979613098565b60115460005b81811015610abd57826001600160a01b0316601182815481106109a4576109a4615764565b6000918252602090912001546001600160a01b031603610aad5760116109cb600184615790565b815481106109db576109db615764565b600091825260209091200154601180546001600160a01b039092169183908110610a0757610a07615764565b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b031602179055506011805480610a4657610a466157a3565b6000828152602090819020600019908301810180546001600160a01b03191690559091019091556040516001600160a01b03851681527ff80a1a97fd42251f3c33cda98635e7399253033a6774fe37cd3f650b5282af3791015b60405180910390a1505050565b610ab6816157b9565b905061097f565b50604051635428d44960e01b81526001600160a01b03831660048201526024015b60405180910390fd5b50565b610af2613098565b604080518082018252600091610b21919084906002908390839080828437600092019190915250612a7f915050565b6000818152600d602090815260409182902082518084019093525460ff811615158084526101009091046001600160401b03169183019190915291925090610b7f57604051631dfd6e1360e21b815260048101839052602401610ade565b6000828152600d60205260408120805468ffffffffffffffffff19169055600e54905b81811015610c515783600e8281548110610bbe57610bbe615764565b906000526020600020015403610c4157600e610bdb600184615790565b81548110610beb57610beb615764565b9060005260206000200154600e8281548110610c0957610c09615764565b600091825260209091200155600e805480610c2657610c266157a3565b60019003818190600052602060002001600090559055610c51565b610c4a816157b9565b9050610ba2565b507f9b6868e0eb737bcd72205360baa6bfd0ba4e4819a33ade2db384e8a8025639a5838360200151604051610c999291909182526001600160401b0316602082015260400190565b60405180910390a150505050565b81610cb1816130f4565b610cb961315e565b610cc283611536565b15610ce057604051631685ecdd60e31b815260040160405180910390fd5b610cea838361318c565b505050565b610cf761315e565b610cff613098565b600b54600160601b90046001600160601b0316600003610d3257604051631e9acf1760e31b815260040160405180910390fd5b600b8054600160601b90046001600160601b0316908190600c610d5583806157d2565b92506101000a8154816001600160601b0302191690836001600160601b0316021790555080600a600c8282829054906101000a90046001600160601b0316610d9d91906157d2565b92506101000a8154816001600160601b0302191690836001600160601b031602179055506000826001600160a01b0316826001600160601b031660405160006040518083038185875af1925050503d8060008114610e17576040519150601f19603f3d011682016040523d82523d6000602084013e610e1c565b606091505b5050905080610cea5760405163950b247960e01b815260040160405180910390fd5b6000610e4861315e565b60005a9050610324361115610e7a57604051630f28961b60e01b81523660048201526103246024820152604401610ade565b6000610e868686613332565b90506000610e9c858360000151602001516135e3565b60408301516060888101519293509163ffffffff16806001600160401b03811115610ec957610ec9615129565b604051908082528060200260200182016040528015610ef2578160200160208202803683370190505b50925060005b81811015610f5a5760408051602081018590529081018290526060016040516020818303038152906040528051906020012060001c848281518110610f3f57610f3f615764565b6020908102919091010152610f53816157b9565b9050610ef8565b5050602080850180516000908152600f9092526040822082905551610f80908a8561363e565b60208a8101516000908152600690915260409020805491925090601890610fb690600160c01b90046001600160401b03166157f2565b82546101009290920a6001600160401b0381810219909316918316021790915560808a01516001600160a01b03166000908152600460209081526040808320828e0151845290915290208054909160099161101991600160481b90910416615818565b91906101000a8154816001600160401b0302191690836001600160401b0316021790555060008960a0015160018b60a00151516110569190615790565b8151811061106657611066615764565b60209101015160f81c600114905060006110828887848d6136e2565b909950905080156110cd5760208088015160105460408051928352928201527f6ca648a381f22ead7e37773d934e64885dcf861fbfbb26c40354cbf0c4662d1a910160405180910390a15b506110dd88828c6020015161371a565b6020808b015187820151604080518781526001600160601b038d16948101949094528415159084015284151560608401528b1515608084015290917faeb4b4786571e184246d39587f659abf0e26f41f6a3358692250382c0cdb47b79060a00160405180910390a3505050505050505b9392505050565b61115c61315e565b61116581613887565b61118d57604051635428d44960e01b81526001600160a01b0382166004820152602401610ade565b60008060008061119c86612dd8565b945094505093509350336001600160a01b0316826001600160a01b0316146112065760405162461bcd60e51b815260206004820152601660248201527f4e6f7420737562736372697074696f6e206f776e6572000000000000000000006044820152606401610ade565b61120f86611536565b1561125c5760405162461bcd60e51b815260206004820152601660248201527f50656e64696e67207265717565737420657869737473000000000000000000006044820152606401610ade565b6040805160c0810182526001815260208082018990526001600160a01b03851682840152606082018490526001600160601b038088166080840152861660a0830152915190916000916112b19184910161583b565b60405160208183030381529060405290506112cb886138f2565b505060405163ce3f471960e01b81526001600160a01b0388169063ce3f4719906001600160601b03881690611304908590600401615900565b6000604051808303818588803b15801561131d57600080fd5b505af1158015611331573d6000803e3d6000fd5b50506002546001600160a01b03161580159350915061135a905057506001600160601b03861615155b1561142a5760025460405163a9059cbb60e01b81526001600160a01b0389811660048301526001600160601b03891660248301529091169063a9059cbb906044016020604051808303816000875af11580156113ba573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113de9190615913565b61142a5760405162461bcd60e51b815260206004820152601260248201527f696e73756666696369656e742066756e647300000000000000000000000000006044820152606401610ade565b600c805466ff0000000000001916660100000000000017905560005b83518110156114d95783818151811061146157611461615764565b6020908102919091010151604051638ea9811760e01b81526001600160a01b038a8116600483015290911690638ea9811790602401600060405180830381600087803b1580156114b057600080fd5b505af11580156114c4573d6000803e3d6000fd5b50505050806114d2906157b9565b9050611446565b50600c805466ff00000000000019169055604080516001600160a01b0389168152602081018a90527fd63ca8cb945956747ee69bfdc3ea754c24a4caf7418db70e46052f7850be4187910160405180910390a15050505050505050565b6000818152600560205260408120600201805480830361155a575060009392505050565b60005b818110156115df5760006004600085848154811061157d5761157d615764565b60009182526020808320909101546001600160a01b0316835282810193909352604091820181208982529092529020546001600160401b03600160481b9091041611156115cf57506001949350505050565b6115d8816157b9565b905061155d565b506000949350505050565b6115f261315e565b6115fa613098565b6002546001600160a01b03166116235760405163c1f0c0a160e01b815260040160405180910390fd5b600b546001600160601b031660000361164f57604051631e9acf1760e31b815260040160405180910390fd5b600b80546001600160601b0316908190600061166b83806157d2565b92506101000a8154816001600160601b0302191690836001600160601b0316021790555080600a60008282829054906101000a90046001600160601b03166116b391906157d2565b82546101009290920a6001600160601b0381810219909316918316021790915560025460405163a9059cbb60e01b81526001600160a01b03868116600483015292851660248201529116915063a9059cbb906044016020604051808303816000875af1158015611727573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061174b9190615913565b61176857604051631e9acf1760e31b815260040160405180910390fd5b5050565b611774613098565b61177d81613887565b156117a65760405163ac8a27ef60e01b81526001600160a01b0382166004820152602401610ade565b601180546001810182556000919091527f31ecc21a745e3968a04e9570e4425bc18fa8019c68028196b546d1669c200c680180546001600160a01b0319166001600160a01b0383169081179091556040519081527fb7cabbfc11e66731fc77de0444614282023bcbd41d16781c753a431d0af016259060200160405180910390a150565b611832613098565b6002546001600160a01b03161561185c57604051631688c53760e11b815260040160405180910390fd5b600280546001600160a01b039384166001600160a01b03199182161790915560038054929093169116179055565b6001546001600160a01b031633146118e45760405162461bcd60e51b815260206004820152601660248201527f4d7573742062652070726f706f736564206f776e6572000000000000000000006044820152606401610ade565b60008054336001600160a01b0319808316821784556001805490911690556040516001600160a01b0390921692909183917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a350565b611943613098565b604080518082018252600091611972919085906002908390839080828437600092019190915250612a7f915050565b6000818152600d602052604090205490915060ff16156119a857604051634a0b8fa760e01b815260048101829052602401610ade565b60408051808201825260018082526001600160401b0385811660208085018281526000888152600d835287812096518754925168ffffffffffffffffff1990931690151568ffffffffffffffff00191617610100929095169190910293909317909455600e805493840181559091527fbb7b4a454dc3493923482f07822329ed19e8244eff582cc204f8554c3620c3fd9091018490558251848152918201527f9b911b2c240bfbef3b6a8f7ed6ee321d1258bb2a3fe6becab52ac1cd3210afd39101610aa0565b611a77613098565b600a544790600160601b90046001600160601b031681811115611ab7576040516354ced18160e11b81526004810182905260248101839052604401610ade565b81811015610cea576000611acb8284615790565b90506000846001600160a01b03168260405160006040518083038185875af1925050503d8060008114611b1a576040519150601f19603f3d011682016040523d82523d6000602084013e611b1f565b606091505b5050905080611b415760405163950b247960e01b815260040160405180910390fd5b604080516001600160a01b0387168152602081018490527f4aed7c8eed0496c8c19ea2681fcca25741c1602342e38b045d9f1e8e905d2e9c910160405180910390a15050505050565b611b9261315e565b6000818152600560205260409020546001600160a01b0316611bc757604051630fb532db60e11b815260040160405180910390fd5b60008181526006602052604090208054600160601b90046001600160601b0316903490600c611bf68385615930565b92506101000a8154816001600160601b0302191690836001600160601b0316021790555034600a600c8282829054906101000a90046001600160601b0316611c3e9190615930565b92506101000a8154816001600160601b0302191690836001600160601b03160217905550817f7603b205d03651ee812f803fccde89f1012e545a9c99f0abfea9cedd0fd8e902823484611c919190615950565b604080519283526020830191909152015b60405180910390a25050565b6000611cb861315e565b602080830135600081815260059092526040909120546001600160a01b0316611cf457604051630fb532db60e11b815260040160405180910390fd5b336000908152600460209081526040808320848452808352928190208151606081018352905460ff811615158083526001600160401b036101008304811695840195909552600160481b9091049093169181019190915290611d72576040516379bfd40160e01b815260048101849052336024820152604401610ade565b600c5461ffff16611d896060870160408801615963565b61ffff161080611dac575060c8611da66060870160408801615963565b61ffff16115b15611df257611dc16060860160408701615963565b600c5460405163539c34bb60e11b815261ffff92831660048201529116602482015260c86044820152606401610ade565b600c5462010000900463ffffffff16611e11608087016060880161597e565b63ffffffff161115611e6157611e2d608086016060870161597e565b600c54604051637aebf00f60e11b815263ffffffff9283166004820152620100009091049091166024820152604401610ade565b6101f4611e7460a087016080880161597e565b63ffffffff161115611eba57611e9060a086016080870161597e565b6040516311ce1afb60e21b815263ffffffff90911660048201526101f46024820152604401610ade565b806020018051611ec9906157f2565b6001600160401b03169052604081018051611ee3906157f2565b6001600160401b03908116909152602082810151604080518935818501819052338284015260608201899052929094166080808601919091528151808603909101815260a08501825280519084012060c085019290925260e08085018390528151808603909101815261010090940190528251929091019190912060009190955090506000611f85611f80611f7b60a08a018a615999565b613aa4565b613b25565b905085611f90613b96565b86611fa160808b0160608c0161597e565b611fb160a08c0160808d0161597e565b3386604051602001611fc997969594939291906159e6565b60405160208183030381529060405280519060200120600f600088815260200190815260200160002081905550336001600160a01b03168588600001357feb0e3652e0f44f417695e6e90f2f42c99b65cd7169074c5a654b16b9748c3a4e89868c604001602081019061203c9190615963565b8d606001602081019061204f919061597e565b8e6080016020810190612062919061597e565b8960405161207596959493929190615a3d565b60405180910390a45050600092835260209182526040928390208151815493830151929094015168ffffffffffffffffff1990931693151568ffffffffffffffff001916939093176101006001600160401b03928316021770ffffffffffffffff0000000000000000001916600160481b91909216021790555b919050565b60006120fe61315e565b6007546001600160401b031633612116600143615790565b6040516bffffffffffffffffffffffff19606093841b81166020830152914060348201523090921b1660548201526001600160c01b031960c083901b16606882015260700160408051601f1981840301815291905280516020909101209150612180816001615a7c565b6007805467ffffffffffffffff19166001600160401b03928316179055604080516000808252608082018352602080830182815283850183815260608086018581528a86526006855287862093518454935191516001600160601b039182166001600160c01b031990951694909417600160601b91909216021777ffffffffffffffffffffffffffffffffffffffffffffffff16600160c01b9290981691909102969096179055835194850184523385528481018281528585018481528884526005835294909220855181546001600160a01b03199081166001600160a01b0392831617835593516001830180549095169116179092559251805192949391926122909260028501920190614f9d565b506122a091506008905084613c17565b5060405133815283907f1d3015d7ba850fa198dc7b1a3f5d42779313a681035f77c8c03764c61005518d9060200160405180910390a2505090565b6122e361315e565b6002546001600160a01b0316331461230e576040516344b0e3c360e01b815260040160405180910390fd5b6020811461232f57604051638129bbcd60e01b815260040160405180910390fd5b600061233d82840184615449565b6000818152600560205260409020549091506001600160a01b031661237557604051630fb532db60e11b815260040160405180910390fd5b600081815260066020526040812080546001600160601b03169186919061239c8385615930565b92506101000a8154816001600160601b0302191690836001600160601b0316021790555084600a60008282829054906101000a90046001600160601b03166123e49190615930565b92506101000a8154816001600160601b0302191690836001600160601b03160217905550817f1ced9348ff549fceab2ac57cd3a9de38edaaab274b725ee82c23e8fc8c4eec7a8287846124379190615950565b6040805192835260208301919091520160405180910390a2505050505050565b61245f613098565b60c861ffff8a1611156124995760405163539c34bb60e11b815261ffff8a1660048201819052602482015260c86044820152606401610ade565b600085136124bd576040516321ea67b360e11b815260048101869052602401610ade565b8363ffffffff168363ffffffff1611156124fa576040516313c06e5960e11b815263ffffffff808516600483015285166024820152604401610ade565b609b60ff8316111561252b57604051631d66288d60e11b815260ff83166004820152609b6024820152604401610ade565b609b60ff8216111561255c57604051631d66288d60e11b815260ff82166004820152609b6024820152604401610ade565b604080516101208101825261ffff8b1680825263ffffffff808c16602084018190526000848601528b8216606085018190528b8316608086018190528a841660a08701819052938a1660c0870181905260ff808b1660e08901819052908a16610100909801889052600c8054600160c01b90990260ff60c01b19600160b81b9093029290921661ffff60b81b19600160981b90940263ffffffff60981b19600160781b9099029890981676ffffffffffffffff00000000000000000000000000000019600160581b9096026effffffff000000000000000000000019670100000000000000909802979097166effffffffffffffffff000000000000196201000090990265ffffffffffff19909c16909a179a909a1796909616979097179390931791909116959095179290921793909316929092179190911790556010869055517f2c6b6b12413678366b05b145c5f00745bdd00e739131ab5de82484a50c9d78b69061272b908b908b908b908b908b908b908b908b908b9061ffff99909916895263ffffffff97881660208a0152958716604089015293861660608801526080870192909252841660a086015290921660c084015260ff91821660e0840152166101008201526101200190565b60405180910390a1505050505050505050565b612746613098565b6000818152600560205260409020546001600160a01b03168061277c57604051630fb532db60e11b815260040160405180910390fd5b611768828261318c565b606060006127946008613c23565b90508084106127b657604051631390f2a160e01b815260040160405180910390fd5b60006127c28486615950565b9050818111806127d0575083155b6127da57806127dc565b815b905060006127ea8683615790565b9050806001600160401b0381111561280457612804615129565b60405190808252806020026020018201604052801561282d578160200160208202803683370190505b50935060005b8181101561287d576128506128488883615950565b600890613c2d565b85828151811061286257612862615764565b6020908102919091010152612876816157b9565b9050612833565b505050505b92915050565b61289061315e565b6000818152600560205260409020546001600160a01b0316806128c657604051630fb532db60e11b815260040160405180910390fd5b6000828152600560205260409020600101546001600160a01b0316331461291f576000828152600560205260409081902060010154905163d084e97560e01b81526001600160a01b039091166004820152602401610ade565b6000828152600560209081526040918290208054336001600160a01b03199182168117835560019092018054909116905582516001600160a01b03851681529182015283917fd4114ab6e9af9f597c52041f32d62dc57c5c4e4c0d4427006069635e216c93869101611ca2565b81612996816130f4565b61299e61315e565b6001600160a01b03821660009081526004602090815260408083208684529091529020805460ff16156129d15750505050565b6000848152600560205260409020600201805460631901612a05576040516305a48e0f60e01b815260040160405180910390fd5b8154600160ff199091168117835581549081018255600082815260209081902090910180546001600160a01b0319166001600160a01b03871690811790915560405190815286917f1e980d04aa7648e205713e5e8ea3808672ac163d10936d36f91b2c88ac1575e191015b60405180910390a25050505050565b600081604051602001612a929190615abf565b604051602081830303815290604052805190602001209050919050565b81612ab9816130f4565b612ac161315e565b612aca83611536565b15612ae857604051631685ecdd60e31b815260040160405180910390fd5b6001600160a01b038216600090815260046020908152604080832086845290915290205460ff16612b3e576040516379bfd40160e01b8152600481018490526001600160a01b0383166024820152604401610ade565b600083815260056020908152604080832060020180548251818502810185019093528083529192909190830182828015612ba157602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311612b83575b50505050509050600060018251612bb89190615790565b905060005b8251811015612cc157846001600160a01b0316838281518110612be257612be2615764565b60200260200101516001600160a01b031603612cb1576000838381518110612c0c57612c0c615764565b6020026020010151905080600560008981526020019081526020016000206002018381548110612c3e57612c3e615764565b600091825260208083209190910180546001600160a01b0319166001600160a01b039490941693909317909255888152600590915260409020600201805480612c8957612c896157a3565b600082815260209020810160001990810180546001600160a01b031916905501905550612cc1565b612cba816157b9565b9050612bbd565b506001600160a01b0384166000818152600460209081526040808320898452825291829020805460ff19169055905191825286917f32158c6058347c1601b2d12bc696ac6901d8a9a9aa3ba10c27ab0a983e8425a79101612a70565b600e8181548110612d2d57600080fd5b600091825260209091200154905081565b81612d48816130f4565b612d5061315e565b600083815260056020526040902060018101546001600160a01b03848116911614612dd2576001810180546001600160a01b0319166001600160a01b03851690811790915560408051338152602081019290925285917f21a4dad170a6bf476c31bbcf4a16628295b0e450672eec25d7c93308e05344a1910160405180910390a25b50505050565b600081815260056020526040812054819081906001600160a01b0316606081612e1457604051630fb532db60e11b815260040160405180910390fd5b600086815260066020908152604080832054600583529281902060020180548251818502810185019093528083526001600160601b0380861695600160601b810490911694600160c01b9091046001600160401b0316938893929091839190830182828015612eac57602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311612e8e575b505050505090509450945094509450945091939590929450565b612ece613098565b6002546001600160a01b0316612ef75760405163c1f0c0a160e01b815260040160405180910390fd5b6002546040516370a0823160e01b81523060048201526000916001600160a01b0316906370a0823190602401602060405180830381865afa158015612f40573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612f649190615acd565b600a549091506001600160601b031681811115612f9e576040516354ced18160e11b81526004810182905260248101839052604401610ade565b81811015610cea576000612fb28284615790565b60025460405163a9059cbb60e01b81526001600160a01b0387811660048301526024820184905292935091169063a9059cbb906044016020604051808303816000875af1158015613007573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061302b9190615913565b61304857604051631f01ff1360e21b815260040160405180910390fd5b604080516001600160a01b0386168152602081018390527f59bfc682b673f8cbf945f1e454df9334834abf7dfe7f92237ca29ecb9b4366009101610c99565b61308f613098565b610ae781613c39565b6000546001600160a01b031633146130f25760405162461bcd60e51b815260206004820152601660248201527f4f6e6c792063616c6c61626c65206279206f776e6572000000000000000000006044820152606401610ade565b565b6000818152600560205260409020546001600160a01b03168061312a57604051630fb532db60e11b815260040160405180910390fd5b336001600160a01b0382161461176857604051636c51fda960e11b81526001600160a01b0382166004820152602401610ade565b600c546601000000000000900460ff16156130f25760405163769dd35360e11b815260040160405180910390fd5b600080613198846138f2565b60025491935091506001600160a01b0316158015906131bf57506001600160601b03821615155b156132605760025460405163a9059cbb60e01b81526001600160a01b0385811660048301526001600160601b03851660248301529091169063a9059cbb906044016020604051808303816000875af115801561321f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906132439190615913565b61326057604051631e9acf1760e31b815260040160405180910390fd5b6000836001600160a01b0316826001600160601b031660405160006040518083038185875af1925050503d80600081146132b6576040519150601f19603f3d011682016040523d82523d6000602084013e6132bb565b606091505b50509050806132dd5760405163950b247960e01b815260040160405180910390fd5b604080516001600160a01b03861681526001600160601b03808616602083015284169181019190915285907f8c74ce8b8cf87f5eb001275c8be27eb34ea2b62bfab6814fcc62192bb63e81c490606001612a70565b6040805160a0810182526000606082018181526080830182905282526020820181905291810191909152600061336b8460000151612a7f565b6000818152600d602090815260409182902082518084019093525460ff811615158084526101009091046001600160401b031691830191909152919250906133c957604051631dfd6e1360e21b815260048101839052602401610ade565b60008286608001516040516020016133eb929190918252602082015260400190565b60408051601f1981840301815291815281516020928301206000818152600f909352908220549092509081900361343557604051631b44092560e11b815260040160405180910390fd5b85516020808801516040808a015160608b015160808c015160a08d01519351613464978a979096959101615ae6565b6040516020818303038152906040528051906020012081146134995760405163354a450b60e21b815260040160405180910390fd5b60006134a88760000151613ce2565b905080613571578651604051631d2827a760e31b81526001600160401b0390911660048201527f0000000000000000000000001159e1889754a1f0862f8ec0e109f169aecbcd6f6001600160a01b03169063e9413d3890602401602060405180830381865afa15801561351f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906135439190615acd565b90508061357157865160405163175dadad60e01b81526001600160401b039091166004820152602401610ade565b6000886080015182604051602001613593929190918252602082015260400190565b6040516020818303038152906040528051906020012060001c905060006135ba8a83613db5565b604080516060810182529788526020880196909652948601949094525092979650505050505050565b6000816001600160401b03163a111561363657821561360c57506001600160401b038116612882565b60405163435e532d60e11b81523a60048201526001600160401b0383166024820152604401610ade565b503a92915050565b6000806000631fe543e360e01b868560405160240161365e929190615b39565b60408051601f198184030181529181526020820180516001600160e01b03166001600160e01b031990941693909317909252600c805466ff000000000000191666010000000000001790559086015160808701519192506136c89163ffffffff9091169083613e20565b600c805466ff000000000000191690559695505050505050565b6000808315613701576136f6868685613e6c565b600091509150613711565b61370c868685613f7d565b915091505b94509492505050565b600081815260066020526040902082156137ee5780546001600160601b03600160601b909104811690851681101561376557604051631e9acf1760e31b815260040160405180910390fd5b61376f85826157d2565b82547fffffffffffffffff000000000000000000000000ffffffffffffffffffffffff16600160601b6001600160601b039283168102919091178455600b805488939192600c926137c4928692900416615930565b92506101000a8154816001600160601b0302191690836001600160601b0316021790555050612dd2565b80546001600160601b0390811690851681101561381e57604051631e9acf1760e31b815260040160405180910390fd5b61382885826157d2565b82546bffffffffffffffffffffffff19166001600160601b03918216178355600b8054879260009161385c91859116615930565b92506101000a8154816001600160601b0302191690836001600160601b031602179055505050505050565b601154600090815b818110156138e857836001600160a01b0316601182815481106138b4576138b4615764565b6000918252602090912001546001600160a01b0316036138d8575060019392505050565b6138e1816157b9565b905061388f565b5060009392505050565b60008181526005602090815260408083206006909252822054600290910180546001600160601b0380841694600160601b90940416925b8181101561399e576004600084838154811061394757613947615764565b60009182526020808320909101546001600160a01b0316835282810193909352604091820181208982529092529020805470ffffffffffffffffffffffffffffffffff19169055613997816157b9565b9050613929565b50600085815260056020526040812080546001600160a01b031990811682556001820180549091169055906139d66002830182615002565b50506000858152600660205260408120556139f260088661416f565b506001600160601b03841615613a4557600a8054859190600090613a209084906001600160601b03166157d2565b92506101000a8154816001600160601b0302191690836001600160601b031602179055505b6001600160601b03831615613a9d5782600a600c8282829054906101000a90046001600160601b0316613a7891906157d2565b92506101000a8154816001600160601b0302191690836001600160601b031602179055505b5050915091565b6040805160208101909152600081526000829003613ad15750604080516020810190915260008152612882565b63125fa26760e31b613ae38385615b5a565b6001600160e01b03191614613b0b57604051632923fee760e11b815260040160405180910390fd5b613b188260048186615b8a565b81019061114d9190615bb4565b60607f92fd13387c7fe7befbc38d303d6468778fb9731bc4583f17d92989c6fcfdeaaa82604051602401613b5e91511515815260200190565b60408051601f198184030181529190526020810180516001600160e01b03166001600160e01b03199093169290921790915292915050565b600046613ba28161417b565b15613c105760646001600160a01b031663a3b1b31d6040518163ffffffff1660e01b8152600401602060405180830381865afa158015613be6573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613c0a9190615acd565b91505090565b4391505090565b600061114d838361419e565b6000612882825490565b600061114d83836141ed565b336001600160a01b03821603613c915760405162461bcd60e51b815260206004820152601760248201527f43616e6e6f74207472616e7366657220746f2073656c660000000000000000006044820152606401610ade565b600180546001600160a01b0319166001600160a01b0383811691821790925560008054604051929316917fed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae12789190a350565b600046613cee8161417b565b15613da657610100836001600160401b0316613d08613b96565b613d129190615790565b1180613d2e5750613d21613b96565b836001600160401b031610155b15613d3c5750600092915050565b6040516315a03d4160e11b81526001600160401b0384166004820152606490632b407a82906024015b602060405180830381865afa158015613d82573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061114d9190615acd565b50506001600160401b03164090565b6000613de98360000151846020015185604001518660600151868860a001518960c001518a60e001518b6101000151614217565b60038360200151604051602001613e01929190615bff565b60408051601f1981840301815291905280516020909101209392505050565b60005a611388811015613e3257600080fd5b611388810390508460408204820311613e4a57600080fd5b50823b613e5657600080fd5b60008083516020850160008789f1949350505050565b600080613eaf6000368080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061444292505050565b905060005a600c54613ecf908890600160581b900463ffffffff16615950565b613ed99190615790565b613ee39086615c13565b600c54909150600090613f0890600160781b900463ffffffff1664e8d4a51000615c13565b90508415613f5457600c548190606490600160b81b900460ff16613f2c8587615950565b613f369190615c13565b613f409190615c40565b613f4a9190615950565b935050505061114d565b600c548190606490613f7090600160b81b900460ff1682615c54565b60ff16613f2c8587615950565b600080600080613f8b614522565b9150915060008213613fb3576040516321ea67b360e11b815260048101839052602401610ade565b6000613ff56000368080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061444292505050565b9050600083825a600c54614017908d90600160581b900463ffffffff16615950565b6140219190615790565b61402b908b615c13565b6140359190615950565b61404790670de0b6b3a7640000615c13565b6140519190615c40565b600c5490915060009061407a9063ffffffff600160981b8204811691600160781b900416615c6d565b61408f9063ffffffff1664e8d4a51000615c13565b90506000856140a683670de0b6b3a7640000615c13565b6140b09190615c40565b9050600089156140f157600c5482906064906140d690600160c01b900460ff1687615c13565b6140e09190615c40565b6140ea9190615950565b9050614131565b600c54829060649061410d90600160c01b900460ff1682615c54565b61411a9060ff1687615c13565b6141249190615c40565b61412e9190615950565b90505b6b033b2e3c9fd0803ce800000081111561415e5760405163e80fa38160e01b815260040160405180910390fd5b9b949a509398505050505050505050565b600061114d83836145ed565b600061a4b182148061418f575062066eed82145b8061288257505062066eee1490565b60008181526001830160205260408120546141e557508154600181810184556000848152602080822090930184905584548482528286019093526040902091909155612882565b506000612882565b600082600001828154811061420457614204615764565b9060005260206000200154905092915050565b614220896146e7565b61426c5760405162461bcd60e51b815260206004820152601a60248201527f7075626c6963206b6579206973206e6f74206f6e2063757276650000000000006044820152606401610ade565b614275886146e7565b6142c15760405162461bcd60e51b815260206004820152601560248201527f67616d6d61206973206e6f74206f6e20637572766500000000000000000000006044820152606401610ade565b6142ca836146e7565b6143165760405162461bcd60e51b815260206004820152601d60248201527f6347616d6d615769746e657373206973206e6f74206f6e2063757276650000006044820152606401610ade565b61431f826146e7565b61436b5760405162461bcd60e51b815260206004820152601c60248201527f73486173685769746e657373206973206e6f74206f6e206375727665000000006044820152606401610ade565b614377878a88876147c0565b6143c35760405162461bcd60e51b815260206004820152601960248201527f6164647228632a706b2b732a6729213d5f755769746e657373000000000000006044820152606401610ade565b60006143cf8a876148e3565b905060006143e2898b878b868989614947565b905060006143f3838d8d8a86614a73565b9050808a146144345760405162461bcd60e51b815260206004820152600d60248201526c34b73b30b634b210383937b7b360991b6044820152606401610ade565b505050505050505050505050565b60004661444e8161417b565b1561449257606c6001600160a01b031663c6f7de0e6040518163ffffffff1660e01b8152600401602060405180830381865afa158015613d82573d6000803e3d6000fd5b61449b81614ab3565b156145195773420000000000000000000000000000000000000f6001600160a01b03166349948e0e84604051806080016040528060488152602001615dca604891396040516020016144ee929190615c8a565b6040516020818303038152906040526040518263ffffffff1660e01b8152600401613d659190615900565b50600092915050565b600c5460035460408051633fabe5a360e21b81529051600093849367010000000000000090910463ffffffff169284926001600160a01b039092169163feaf968c9160048082019260a0929091908290030181865afa158015614589573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906145ad9190615cd3565b50919650909250505063ffffffff8216158015906145d957506145d08142615790565b8263ffffffff16105b925082156145e75760105493505b50509091565b600081815260018301602052604081205480156146d6576000614611600183615790565b855490915060009061462590600190615790565b905081811461468a57600086600001828154811061464557614645615764565b906000526020600020015490508087600001848154811061466857614668615764565b6000918252602080832090910192909255918252600188019052604090208390555b855486908061469b5761469b6157a3565b600190038181906000526020600020016000905590558560010160008681526020019081526020016000206000905560019350505050612882565b6000915050612882565b5092915050565b80516000906401000003d019116147405760405162461bcd60e51b815260206004820152601260248201527f696e76616c696420782d6f7264696e61746500000000000000000000000000006044820152606401610ade565b60208201516401000003d019116147995760405162461bcd60e51b815260206004820152601260248201527f696e76616c696420792d6f7264696e61746500000000000000000000000000006044820152606401610ade565b60208201516401000003d0199080096147b98360005b6020020151614aed565b1492915050565b60006001600160a01b0382166148065760405162461bcd60e51b815260206004820152600b60248201526a626164207769746e65737360a81b6044820152606401610ade565b60208401516000906001161561481d57601c614820565b601b5b9050600070014551231950b75fc4402da1732fc9bebe1985876000602002015109865170014551231950b75fc4402da1732fc9bebe19918203925060009190890987516040805160008082526020820180845287905260ff88169282019290925260608101929092526080820183905291925060019060a0016020604051602081039080840390855afa1580156148bb573d6000803e3d6000fd5b5050604051601f1901516001600160a01b039081169088161495505050505050949350505050565b6148eb615020565b6149186001848460405160200161490493929190615d23565b604051602081830303815290604052614b11565b90505b614924816146e7565b6128825780516040805160208101929092526149409101614904565b905061491b565b61494f615020565b825186516401000003d01991829006919006036149ae5760405162461bcd60e51b815260206004820152601e60248201527f706f696e747320696e2073756d206d7573742062652064697374696e637400006044820152606401610ade565b6149b9878988614b5e565b614a055760405162461bcd60e51b815260206004820152601660248201527f4669727374206d756c20636865636b206661696c6564000000000000000000006044820152606401610ade565b614a10848685614b5e565b614a5c5760405162461bcd60e51b815260206004820152601760248201527f5365636f6e64206d756c20636865636b206661696c65640000000000000000006044820152606401610ade565b614a67868484614c89565b98975050505050505050565b600060028686868587604051602001614a9196959493929190615d44565b60408051601f1981840301815291905280516020909101209695505050505050565b6000600a821480614ac557506101a482145b80614ad2575062aa37dc82145b80614ade575061210582145b8061288257505062014a331490565b6000806401000003d01980848509840990506401000003d019600782089392505050565b614b19615020565b614b2282614d50565b8152614b37614b328260006147af565b614d8b565b60208201819052600290066001036120ef576020810180516401000003d019039052919050565b600082600003614b9e5760405162461bcd60e51b815260206004820152600b60248201526a3d32b9379039b1b0b630b960a91b6044820152606401610ade565b83516020850151600090614bb490600290615da3565b15614bc057601c614bc3565b601b5b9050600070014551231950b75fc4402da1732fc9bebe198387096040805160008082526020820180845281905260ff86169282019290925260608101869052608081018390529192509060019060a0016020604051602081039080840390855afa158015614c35573d6000803e3d6000fd5b505050602060405103519050600086604051602001614c549190615db7565b60408051601f1981840301815291905280516020909101206001600160a01b0392831692169190911498975050505050505050565b614c91615020565b835160208086015185519186015160009384938493614cb293909190614dab565b919450925090506401000003d019858209600114614d125760405162461bcd60e51b815260206004820152601960248201527f696e765a206d75737420626520696e7665727365206f66207a000000000000006044820152606401610ade565b60405180604001604052806401000003d01980614d3157614d31615c2a565b87860981526020016401000003d0198785099052979650505050505050565b805160208201205b6401000003d01981106120ef57604080516020808201939093528151808203840181529082019091528051910120614d58565b6000612882826002614da46401000003d0196001615950565b901c614e8b565b60008080600180826401000003d019896401000003d019038808905060006401000003d0198b6401000003d019038a0890506000614deb83838585614f30565b9098509050614dfc88828e88614f54565b9098509050614e0d88828c87614f54565b90985090506000614e208d878b85614f54565b9098509050614e3188828686614f30565b9098509050614e4288828e89614f54565b9098509050818114614e77576401000003d019818a0998506401000003d01982890997506401000003d0198183099650614e7b565b8196505b5050505050509450945094915050565b600080614e9661503e565b6020808252818101819052604082015260608101859052608081018490526401000003d01960a0820152614ec861505c565b60208160c0846005600019fa925082600003614f265760405162461bcd60e51b815260206004820152601260248201527f6269674d6f64457870206661696c7572652100000000000000000000000000006044820152606401610ade565b5195945050505050565b6000806401000003d0198487096401000003d0198487099097909650945050505050565b600080806401000003d019878509905060006401000003d01987876401000003d019030990506401000003d0198183086401000003d01986890990999098509650505050505050565b828054828255906000526020600020908101928215614ff2579160200282015b82811115614ff257825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190614fbd565b50614ffe92915061507a565b5090565b5080546000825590600052602060002090810190610ae7919061507a565b60405180604001604052806002906020820280368337509192915050565b6040518060c001604052806006906020820280368337509192915050565b60405180602001604052806001906020820280368337509192915050565b5b80821115614ffe576000815560010161507b565b6001600160a01b0381168114610ae757600080fd5b80356120ef8161508f565b6000602082840312156150c157600080fd5b813561114d8161508f565b806040810183101561288257600080fd5b6000604082840312156150ef57600080fd5b61114d83836150cc565b6000806040838503121561510c57600080fd5b82359150602083013561511e8161508f565b809150509250929050565b634e487b7160e01b600052604160045260246000fd5b60405160c081016001600160401b038111828210171561516157615161615129565b60405290565b60405161012081016001600160401b038111828210171561516157615161615129565b604051601f8201601f191681016001600160401b03811182821017156151b2576151b2615129565b604052919050565b600082601f8301126151cb57600080fd5b604051604081018181106001600160401b03821117156151ed576151ed615129565b806040525080604084018581111561520457600080fd5b845b8181101561521e578035835260209283019201615206565b509195945050505050565b80356001600160401b03811681146120ef57600080fd5b803563ffffffff811681146120ef57600080fd5b600060c0828403121561526657600080fd5b61526e61513f565b905061527982615229565b81526020808301358183015261529160408401615240565b60408301526152a260608401615240565b606083015260808301356152b58161508f565b608083015260a08301356001600160401b03808211156152d457600080fd5b818501915085601f8301126152e857600080fd5b8135818111156152fa576152fa615129565b61530c601f8201601f1916850161518a565b9150808252868482850101111561532257600080fd5b80848401858401376000848284010152508060a085015250505092915050565b8015158114610ae757600080fd5b80356120ef81615342565b60008060008385036101e081121561537257600080fd5b6101a08082121561538257600080fd5b61538a615167565b915061539687876151ba565b82526153a587604088016151ba565b60208301526080860135604083015260a0860135606083015260c086013560808301526153d460e087016150a4565b60a08301526101006153e8888289016151ba565b60c08401526153fb8861014089016151ba565b60e0840152610180870135908301529093508401356001600160401b0381111561542457600080fd5b61543086828701615254565b9250506154406101c08501615350565b90509250925092565b60006020828403121561545b57600080fd5b5035919050565b6000806040838503121561547557600080fd5b82356154808161508f565b9150602083013561511e8161508f565b600080606083850312156154a357600080fd5b6154ad84846150cc565b91506154bb60408401615229565b90509250929050565b6000602082840312156154d657600080fd5b81356001600160401b038111156154ec57600080fd5b820160c0818503121561114d57600080fd5b6000806000806060858703121561551457600080fd5b843561551f8161508f565b93506020850135925060408501356001600160401b038082111561554257600080fd5b818701915087601f83011261555657600080fd5b81358181111561556557600080fd5b88602082850101111561557757600080fd5b95989497505060200194505050565b803561ffff811681146120ef57600080fd5b803560ff811681146120ef57600080fd5b60008060008060008060008060006101208a8c0312156155c857600080fd5b6155d18a615586565b98506155df60208b01615240565b97506155ed60408b01615240565b96506155fb60608b01615240565b955060808a0135945061561060a08b01615240565b935061561e60c08b01615240565b925061562c60e08b01615598565b915061563b6101008b01615598565b90509295985092959850929598565b6000806040838503121561565d57600080fd5b50508035926020909101359150565b600081518084526020808501945080840160005b8381101561569c57815187529582019590820190600101615680565b509495945050505050565b60208152600061114d602083018461566c565b6000604082840312156156cc57600080fd5b61114d83836151ba565b600081518084526020808501945080840160005b8381101561569c5781516001600160a01b0316875295820195908201906001016156ea565b60006001600160601b0380881683528087166020840152506001600160401b03851660408301526001600160a01b038416606083015260a0608083015261575960a08301846156d6565b979650505050505050565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b818103818111156128825761288261577a565b634e487b7160e01b600052603160045260246000fd5b6000600182016157cb576157cb61577a565b5060010190565b6001600160601b038281168282160390808211156146e0576146e061577a565b60006001600160401b0380831681810361580e5761580e61577a565b6001019392505050565b60006001600160401b038216806158315761583161577a565b6000190192915050565b6020815260ff8251166020820152602082015160408201526001600160a01b0360408301511660608201526000606083015160c0608084015261588160e08401826156d6565b905060808401516001600160601b0380821660a08601528060a08701511660c086015250508091505092915050565b60005b838110156158cb5781810151838201526020016158b3565b50506000910152565b600081518084526158ec8160208601602086016158b0565b601f01601f19169290920160200192915050565b60208152600061114d60208301846158d4565b60006020828403121561592557600080fd5b815161114d81615342565b6001600160601b038181168382160190808211156146e0576146e061577a565b808201808211156128825761288261577a565b60006020828403121561597557600080fd5b61114d82615586565b60006020828403121561599057600080fd5b61114d82615240565b6000808335601e198436030181126159b057600080fd5b8301803591506001600160401b038211156159ca57600080fd5b6020019150368190038213156159df57600080fd5b9250929050565b878152866020820152856040820152600063ffffffff80871660608401528086166080840152506001600160a01b03841660a083015260e060c0830152615a3060e08301846158d4565b9998505050505050505050565b86815285602082015261ffff85166040820152600063ffffffff808616606084015280851660808401525060c060a0830152614a6760c08301846158d4565b6001600160401b038181168382160190808211156146e0576146e061577a565b8060005b6002811015612dd2578151845260209384019390910190600101615aa0565b604081016128828284615a9c565b600060208284031215615adf57600080fd5b5051919050565b8781526001600160401b0387166020820152856040820152600063ffffffff80871660608401528086166080840152506001600160a01b03841660a083015260e060c0830152615a3060e08301846158d4565b828152604060208201526000615b52604083018461566c565b949350505050565b6001600160e01b03198135818116916004851015615b825780818660040360031b1b83161692505b505092915050565b60008085851115615b9a57600080fd5b83861115615ba757600080fd5b5050820193919092039150565b600060208284031215615bc657600080fd5b604051602081018181106001600160401b0382111715615be857615be8615129565b6040528235615bf681615342565b81529392505050565b8281526060810161114d6020830184615a9c565b80820281158282048414176128825761288261577a565b634e487b7160e01b600052601260045260246000fd5b600082615c4f57615c4f615c2a565b500490565b60ff81811683821601908111156128825761288261577a565b63ffffffff8281168282160390808211156146e0576146e061577a565b60008351615c9c8184602088016158b0565b835190830190615cb08183602088016158b0565b01949350505050565b805169ffffffffffffffffffff811681146120ef57600080fd5b600080600080600060a08688031215615ceb57600080fd5b615cf486615cb9565b9450602086015193506040860151925060608601519150615d1760808701615cb9565b90509295509295909350565b838152615d336020820184615a9c565b606081019190915260800192915050565b868152615d546020820187615a9c565b615d616060820186615a9c565b615d6e60a0820185615a9c565b615d7b60e0820184615a9c565b60609190911b6bffffffffffffffffffffffff19166101208201526101340195945050505050565b600082615db257615db2615c2a565b500690565b615dc18183615a9c565b60400191905056fe307866666666666666666666666666666666666666666666666666666666666666666666666666666666666666666666666666666666666666666666666666666666666666666666a164736f6c6343000813000a",
}

// VrfV25ABI is the input ABI used to generate the binding from.
// Deprecated: Use VrfV25MetaData.ABI instead.
var VrfV25ABI = VrfV25MetaData.ABI

// VrfV25Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use VrfV25MetaData.Bin instead.
var VrfV25Bin = VrfV25MetaData.Bin

// DeployVrfV25 deploys a new Ethereum contract, binding an instance of VrfV25 to it.
func DeployVrfV25(auth *bind.TransactOpts, backend bind.ContractBackend, blockhashStore common.Address) (common.Address, *types.Transaction, *VrfV25, error) {
	parsed, err := VrfV25MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VrfV25Bin), backend, blockhashStore)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &VrfV25{VrfV25Caller: VrfV25Caller{contract: contract}, VrfV25Transactor: VrfV25Transactor{contract: contract}, VrfV25Filterer: VrfV25Filterer{contract: contract}}, nil
}

// VrfV25 is an auto generated Go binding around an Ethereum contract.
type VrfV25 struct {
	VrfV25Caller     // Read-only binding to the contract
	VrfV25Transactor // Write-only binding to the contract
	VrfV25Filterer   // Log filterer for contract events
}

// VrfV25Caller is an auto generated read-only Go binding around an Ethereum contract.
type VrfV25Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VrfV25Transactor is an auto generated write-only Go binding around an Ethereum contract.
type VrfV25Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VrfV25Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VrfV25Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VrfV25Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VrfV25Session struct {
	Contract     *VrfV25           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VrfV25CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VrfV25CallerSession struct {
	Contract *VrfV25Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// VrfV25TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VrfV25TransactorSession struct {
	Contract     *VrfV25Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VrfV25Raw is an auto generated low-level Go binding around an Ethereum contract.
type VrfV25Raw struct {
	Contract *VrfV25 // Generic contract binding to access the raw methods on
}

// VrfV25CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VrfV25CallerRaw struct {
	Contract *VrfV25Caller // Generic read-only contract binding to access the raw methods on
}

// VrfV25TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VrfV25TransactorRaw struct {
	Contract *VrfV25Transactor // Generic write-only contract binding to access the raw methods on
}

// NewVrfV25 creates a new instance of VrfV25, bound to a specific deployed contract.
func NewVrfV25(address common.Address, backend bind.ContractBackend) (*VrfV25, error) {
	contract, err := bindVrfV25(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VrfV25{VrfV25Caller: VrfV25Caller{contract: contract}, VrfV25Transactor: VrfV25Transactor{contract: contract}, VrfV25Filterer: VrfV25Filterer{contract: contract}}, nil
}

// NewVrfV25Caller creates a new read-only instance of VrfV25, bound to a specific deployed contract.
func NewVrfV25Caller(address common.Address, caller bind.ContractCaller) (*VrfV25Caller, error) {
	contract, err := bindVrfV25(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VrfV25Caller{contract: contract}, nil
}

// NewVrfV25Transactor creates a new write-only instance of VrfV25, bound to a specific deployed contract.
func NewVrfV25Transactor(address common.Address, transactor bind.ContractTransactor) (*VrfV25Transactor, error) {
	contract, err := bindVrfV25(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VrfV25Transactor{contract: contract}, nil
}

// NewVrfV25Filterer creates a new log filterer instance of VrfV25, bound to a specific deployed contract.
func NewVrfV25Filterer(address common.Address, filterer bind.ContractFilterer) (*VrfV25Filterer, error) {
	contract, err := bindVrfV25(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VrfV25Filterer{contract: contract}, nil
}

// bindVrfV25 binds a generic wrapper to an already deployed contract.
func bindVrfV25(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VrfV25MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VrfV25 *VrfV25Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VrfV25.Contract.VrfV25Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VrfV25 *VrfV25Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VrfV25.Contract.VrfV25Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VrfV25 *VrfV25Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VrfV25.Contract.VrfV25Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VrfV25 *VrfV25CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _VrfV25.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VrfV25 *VrfV25TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VrfV25.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VrfV25 *VrfV25TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VrfV25.Contract.contract.Transact(opts, method, params...)
}

// BLOCKHASHSTORE is a free data retrieval call binding the contract method 0x689c4517.
//
// Solidity: function BLOCKHASH_STORE() view returns(address)
func (_VrfV25 *VrfV25Caller) BLOCKHASHSTORE(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VrfV25.contract.Call(opts, &out, "BLOCKHASH_STORE")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BLOCKHASHSTORE is a free data retrieval call binding the contract method 0x689c4517.
//
// Solidity: function BLOCKHASH_STORE() view returns(address)
func (_VrfV25 *VrfV25Session) BLOCKHASHSTORE() (common.Address, error) {
	return _VrfV25.Contract.BLOCKHASHSTORE(&_VrfV25.CallOpts)
}

// BLOCKHASHSTORE is a free data retrieval call binding the contract method 0x689c4517.
//
// Solidity: function BLOCKHASH_STORE() view returns(address)
func (_VrfV25 *VrfV25CallerSession) BLOCKHASHSTORE() (common.Address, error) {
	return _VrfV25.Contract.BLOCKHASHSTORE(&_VrfV25.CallOpts)
}

// LINK is a free data retrieval call binding the contract method 0x1b6b6d23.
//
// Solidity: function LINK() view returns(address)
func (_VrfV25 *VrfV25Caller) LINK(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VrfV25.contract.Call(opts, &out, "LINK")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LINK is a free data retrieval call binding the contract method 0x1b6b6d23.
//
// Solidity: function LINK() view returns(address)
func (_VrfV25 *VrfV25Session) LINK() (common.Address, error) {
	return _VrfV25.Contract.LINK(&_VrfV25.CallOpts)
}

// LINK is a free data retrieval call binding the contract method 0x1b6b6d23.
//
// Solidity: function LINK() view returns(address)
func (_VrfV25 *VrfV25CallerSession) LINK() (common.Address, error) {
	return _VrfV25.Contract.LINK(&_VrfV25.CallOpts)
}

// LINKNATIVEFEED is a free data retrieval call binding the contract method 0x72e9d565.
//
// Solidity: function LINK_NATIVE_FEED() view returns(address)
func (_VrfV25 *VrfV25Caller) LINKNATIVEFEED(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VrfV25.contract.Call(opts, &out, "LINK_NATIVE_FEED")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LINKNATIVEFEED is a free data retrieval call binding the contract method 0x72e9d565.
//
// Solidity: function LINK_NATIVE_FEED() view returns(address)
func (_VrfV25 *VrfV25Session) LINKNATIVEFEED() (common.Address, error) {
	return _VrfV25.Contract.LINKNATIVEFEED(&_VrfV25.CallOpts)
}

// LINKNATIVEFEED is a free data retrieval call binding the contract method 0x72e9d565.
//
// Solidity: function LINK_NATIVE_FEED() view returns(address)
func (_VrfV25 *VrfV25CallerSession) LINKNATIVEFEED() (common.Address, error) {
	return _VrfV25.Contract.LINKNATIVEFEED(&_VrfV25.CallOpts)
}

// MAXCONSUMERS is a free data retrieval call binding the contract method 0x64d51a2a.
//
// Solidity: function MAX_CONSUMERS() view returns(uint16)
func (_VrfV25 *VrfV25Caller) MAXCONSUMERS(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _VrfV25.contract.Call(opts, &out, "MAX_CONSUMERS")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// MAXCONSUMERS is a free data retrieval call binding the contract method 0x64d51a2a.
//
// Solidity: function MAX_CONSUMERS() view returns(uint16)
func (_VrfV25 *VrfV25Session) MAXCONSUMERS() (uint16, error) {
	return _VrfV25.Contract.MAXCONSUMERS(&_VrfV25.CallOpts)
}

// MAXCONSUMERS is a free data retrieval call binding the contract method 0x64d51a2a.
//
// Solidity: function MAX_CONSUMERS() view returns(uint16)
func (_VrfV25 *VrfV25CallerSession) MAXCONSUMERS() (uint16, error) {
	return _VrfV25.Contract.MAXCONSUMERS(&_VrfV25.CallOpts)
}

// MAXNUMWORDS is a free data retrieval call binding the contract method 0x40d6bb82.
//
// Solidity: function MAX_NUM_WORDS() view returns(uint32)
func (_VrfV25 *VrfV25Caller) MAXNUMWORDS(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _VrfV25.contract.Call(opts, &out, "MAX_NUM_WORDS")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// MAXNUMWORDS is a free data retrieval call binding the contract method 0x40d6bb82.
//
// Solidity: function MAX_NUM_WORDS() view returns(uint32)
func (_VrfV25 *VrfV25Session) MAXNUMWORDS() (uint32, error) {
	return _VrfV25.Contract.MAXNUMWORDS(&_VrfV25.CallOpts)
}

// MAXNUMWORDS is a free data retrieval call binding the contract method 0x40d6bb82.
//
// Solidity: function MAX_NUM_WORDS() view returns(uint32)
func (_VrfV25 *VrfV25CallerSession) MAXNUMWORDS() (uint32, error) {
	return _VrfV25.Contract.MAXNUMWORDS(&_VrfV25.CallOpts)
}

// MAXREQUESTCONFIRMATIONS is a free data retrieval call binding the contract method 0x15c48b84.
//
// Solidity: function MAX_REQUEST_CONFIRMATIONS() view returns(uint16)
func (_VrfV25 *VrfV25Caller) MAXREQUESTCONFIRMATIONS(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _VrfV25.contract.Call(opts, &out, "MAX_REQUEST_CONFIRMATIONS")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// MAXREQUESTCONFIRMATIONS is a free data retrieval call binding the contract method 0x15c48b84.
//
// Solidity: function MAX_REQUEST_CONFIRMATIONS() view returns(uint16)
func (_VrfV25 *VrfV25Session) MAXREQUESTCONFIRMATIONS() (uint16, error) {
	return _VrfV25.Contract.MAXREQUESTCONFIRMATIONS(&_VrfV25.CallOpts)
}

// MAXREQUESTCONFIRMATIONS is a free data retrieval call binding the contract method 0x15c48b84.
//
// Solidity: function MAX_REQUEST_CONFIRMATIONS() view returns(uint16)
func (_VrfV25 *VrfV25CallerSession) MAXREQUESTCONFIRMATIONS() (uint16, error) {
	return _VrfV25.Contract.MAXREQUESTCONFIRMATIONS(&_VrfV25.CallOpts)
}

// GetActiveSubscriptionIds is a free data retrieval call binding the contract method 0xaefb212f.
//
// Solidity: function getActiveSubscriptionIds(uint256 startIndex, uint256 maxCount) view returns(uint256[] ids)
func (_VrfV25 *VrfV25Caller) GetActiveSubscriptionIds(opts *bind.CallOpts, startIndex *big.Int, maxCount *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _VrfV25.contract.Call(opts, &out, "getActiveSubscriptionIds", startIndex, maxCount)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetActiveSubscriptionIds is a free data retrieval call binding the contract method 0xaefb212f.
//
// Solidity: function getActiveSubscriptionIds(uint256 startIndex, uint256 maxCount) view returns(uint256[] ids)
func (_VrfV25 *VrfV25Session) GetActiveSubscriptionIds(startIndex *big.Int, maxCount *big.Int) ([]*big.Int, error) {
	return _VrfV25.Contract.GetActiveSubscriptionIds(&_VrfV25.CallOpts, startIndex, maxCount)
}

// GetActiveSubscriptionIds is a free data retrieval call binding the contract method 0xaefb212f.
//
// Solidity: function getActiveSubscriptionIds(uint256 startIndex, uint256 maxCount) view returns(uint256[] ids)
func (_VrfV25 *VrfV25CallerSession) GetActiveSubscriptionIds(startIndex *big.Int, maxCount *big.Int) ([]*big.Int, error) {
	return _VrfV25.Contract.GetActiveSubscriptionIds(&_VrfV25.CallOpts, startIndex, maxCount)
}

// GetSubscription is a free data retrieval call binding the contract method 0xdc311dd3.
//
// Solidity: function getSubscription(uint256 subId) view returns(uint96 balance, uint96 nativeBalance, uint64 reqCount, address subOwner, address[] consumers)
func (_VrfV25 *VrfV25Caller) GetSubscription(opts *bind.CallOpts, subId *big.Int) (struct {
	Balance       *big.Int
	NativeBalance *big.Int
	ReqCount      uint64
	SubOwner      common.Address
	Consumers     []common.Address
}, error) {
	var out []interface{}
	err := _VrfV25.contract.Call(opts, &out, "getSubscription", subId)

	outstruct := new(struct {
		Balance       *big.Int
		NativeBalance *big.Int
		ReqCount      uint64
		SubOwner      common.Address
		Consumers     []common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Balance = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.NativeBalance = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.ReqCount = *abi.ConvertType(out[2], new(uint64)).(*uint64)
	outstruct.SubOwner = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.Consumers = *abi.ConvertType(out[4], new([]common.Address)).(*[]common.Address)

	return *outstruct, err

}

// GetSubscription is a free data retrieval call binding the contract method 0xdc311dd3.
//
// Solidity: function getSubscription(uint256 subId) view returns(uint96 balance, uint96 nativeBalance, uint64 reqCount, address subOwner, address[] consumers)
func (_VrfV25 *VrfV25Session) GetSubscription(subId *big.Int) (struct {
	Balance       *big.Int
	NativeBalance *big.Int
	ReqCount      uint64
	SubOwner      common.Address
	Consumers     []common.Address
}, error) {
	return _VrfV25.Contract.GetSubscription(&_VrfV25.CallOpts, subId)
}

// GetSubscription is a free data retrieval call binding the contract method 0xdc311dd3.
//
// Solidity: function getSubscription(uint256 subId) view returns(uint96 balance, uint96 nativeBalance, uint64 reqCount, address subOwner, address[] consumers)
func (_VrfV25 *VrfV25CallerSession) GetSubscription(subId *big.Int) (struct {
	Balance       *big.Int
	NativeBalance *big.Int
	ReqCount      uint64
	SubOwner      common.Address
	Consumers     []common.Address
}, error) {
	return _VrfV25.Contract.GetSubscription(&_VrfV25.CallOpts, subId)
}

// HashOfKey is a free data retrieval call binding the contract method 0xcaf70c4a.
//
// Solidity: function hashOfKey(uint256[2] publicKey) pure returns(bytes32)
func (_VrfV25 *VrfV25Caller) HashOfKey(opts *bind.CallOpts, publicKey [2]*big.Int) ([32]byte, error) {
	var out []interface{}
	err := _VrfV25.contract.Call(opts, &out, "hashOfKey", publicKey)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HashOfKey is a free data retrieval call binding the contract method 0xcaf70c4a.
//
// Solidity: function hashOfKey(uint256[2] publicKey) pure returns(bytes32)
func (_VrfV25 *VrfV25Session) HashOfKey(publicKey [2]*big.Int) ([32]byte, error) {
	return _VrfV25.Contract.HashOfKey(&_VrfV25.CallOpts, publicKey)
}

// HashOfKey is a free data retrieval call binding the contract method 0xcaf70c4a.
//
// Solidity: function hashOfKey(uint256[2] publicKey) pure returns(bytes32)
func (_VrfV25 *VrfV25CallerSession) HashOfKey(publicKey [2]*big.Int) ([32]byte, error) {
	return _VrfV25.Contract.HashOfKey(&_VrfV25.CallOpts, publicKey)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VrfV25 *VrfV25Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _VrfV25.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VrfV25 *VrfV25Session) Owner() (common.Address, error) {
	return _VrfV25.Contract.Owner(&_VrfV25.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_VrfV25 *VrfV25CallerSession) Owner() (common.Address, error) {
	return _VrfV25.Contract.Owner(&_VrfV25.CallOpts)
}

// PendingRequestExists is a free data retrieval call binding the contract method 0x41af6c87.
//
// Solidity: function pendingRequestExists(uint256 subId) view returns(bool)
func (_VrfV25 *VrfV25Caller) PendingRequestExists(opts *bind.CallOpts, subId *big.Int) (bool, error) {
	var out []interface{}
	err := _VrfV25.contract.Call(opts, &out, "pendingRequestExists", subId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PendingRequestExists is a free data retrieval call binding the contract method 0x41af6c87.
//
// Solidity: function pendingRequestExists(uint256 subId) view returns(bool)
func (_VrfV25 *VrfV25Session) PendingRequestExists(subId *big.Int) (bool, error) {
	return _VrfV25.Contract.PendingRequestExists(&_VrfV25.CallOpts, subId)
}

// PendingRequestExists is a free data retrieval call binding the contract method 0x41af6c87.
//
// Solidity: function pendingRequestExists(uint256 subId) view returns(bool)
func (_VrfV25 *VrfV25CallerSession) PendingRequestExists(subId *big.Int) (bool, error) {
	return _VrfV25.Contract.PendingRequestExists(&_VrfV25.CallOpts, subId)
}

// SConfig is a free data retrieval call binding the contract method 0x088070f5.
//
// Solidity: function s_config() view returns(uint16 minimumRequestConfirmations, uint32 maxGasLimit, bool reentrancyLock, uint32 stalenessSeconds, uint32 gasAfterPaymentCalculation, uint32 fulfillmentFlatFeeNativePPM, uint32 fulfillmentFlatFeeLinkDiscountPPM, uint8 nativePremiumPercentage, uint8 linkPremiumPercentage)
func (_VrfV25 *VrfV25Caller) SConfig(opts *bind.CallOpts) (struct {
	MinimumRequestConfirmations       uint16
	MaxGasLimit                       uint32
	ReentrancyLock                    bool
	StalenessSeconds                  uint32
	GasAfterPaymentCalculation        uint32
	FulfillmentFlatFeeNativePPM       uint32
	FulfillmentFlatFeeLinkDiscountPPM uint32
	NativePremiumPercentage           uint8
	LinkPremiumPercentage             uint8
}, error) {
	var out []interface{}
	err := _VrfV25.contract.Call(opts, &out, "s_config")

	outstruct := new(struct {
		MinimumRequestConfirmations       uint16
		MaxGasLimit                       uint32
		ReentrancyLock                    bool
		StalenessSeconds                  uint32
		GasAfterPaymentCalculation        uint32
		FulfillmentFlatFeeNativePPM       uint32
		FulfillmentFlatFeeLinkDiscountPPM uint32
		NativePremiumPercentage           uint8
		LinkPremiumPercentage             uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.MinimumRequestConfirmations = *abi.ConvertType(out[0], new(uint16)).(*uint16)
	outstruct.MaxGasLimit = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.ReentrancyLock = *abi.ConvertType(out[2], new(bool)).(*bool)
	outstruct.StalenessSeconds = *abi.ConvertType(out[3], new(uint32)).(*uint32)
	outstruct.GasAfterPaymentCalculation = *abi.ConvertType(out[4], new(uint32)).(*uint32)
	outstruct.FulfillmentFlatFeeNativePPM = *abi.ConvertType(out[5], new(uint32)).(*uint32)
	outstruct.FulfillmentFlatFeeLinkDiscountPPM = *abi.ConvertType(out[6], new(uint32)).(*uint32)
	outstruct.NativePremiumPercentage = *abi.ConvertType(out[7], new(uint8)).(*uint8)
	outstruct.LinkPremiumPercentage = *abi.ConvertType(out[8], new(uint8)).(*uint8)

	return *outstruct, err

}

// SConfig is a free data retrieval call binding the contract method 0x088070f5.
//
// Solidity: function s_config() view returns(uint16 minimumRequestConfirmations, uint32 maxGasLimit, bool reentrancyLock, uint32 stalenessSeconds, uint32 gasAfterPaymentCalculation, uint32 fulfillmentFlatFeeNativePPM, uint32 fulfillmentFlatFeeLinkDiscountPPM, uint8 nativePremiumPercentage, uint8 linkPremiumPercentage)
func (_VrfV25 *VrfV25Session) SConfig() (struct {
	MinimumRequestConfirmations       uint16
	MaxGasLimit                       uint32
	ReentrancyLock                    bool
	StalenessSeconds                  uint32
	GasAfterPaymentCalculation        uint32
	FulfillmentFlatFeeNativePPM       uint32
	FulfillmentFlatFeeLinkDiscountPPM uint32
	NativePremiumPercentage           uint8
	LinkPremiumPercentage             uint8
}, error) {
	return _VrfV25.Contract.SConfig(&_VrfV25.CallOpts)
}

// SConfig is a free data retrieval call binding the contract method 0x088070f5.
//
// Solidity: function s_config() view returns(uint16 minimumRequestConfirmations, uint32 maxGasLimit, bool reentrancyLock, uint32 stalenessSeconds, uint32 gasAfterPaymentCalculation, uint32 fulfillmentFlatFeeNativePPM, uint32 fulfillmentFlatFeeLinkDiscountPPM, uint8 nativePremiumPercentage, uint8 linkPremiumPercentage)
func (_VrfV25 *VrfV25CallerSession) SConfig() (struct {
	MinimumRequestConfirmations       uint16
	MaxGasLimit                       uint32
	ReentrancyLock                    bool
	StalenessSeconds                  uint32
	GasAfterPaymentCalculation        uint32
	FulfillmentFlatFeeNativePPM       uint32
	FulfillmentFlatFeeLinkDiscountPPM uint32
	NativePremiumPercentage           uint8
	LinkPremiumPercentage             uint8
}, error) {
	return _VrfV25.Contract.SConfig(&_VrfV25.CallOpts)
}

// SCurrentSubNonce is a free data retrieval call binding the contract method 0x9d40a6fd.
//
// Solidity: function s_currentSubNonce() view returns(uint64)
func (_VrfV25 *VrfV25Caller) SCurrentSubNonce(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _VrfV25.contract.Call(opts, &out, "s_currentSubNonce")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// SCurrentSubNonce is a free data retrieval call binding the contract method 0x9d40a6fd.
//
// Solidity: function s_currentSubNonce() view returns(uint64)
func (_VrfV25 *VrfV25Session) SCurrentSubNonce() (uint64, error) {
	return _VrfV25.Contract.SCurrentSubNonce(&_VrfV25.CallOpts)
}

// SCurrentSubNonce is a free data retrieval call binding the contract method 0x9d40a6fd.
//
// Solidity: function s_currentSubNonce() view returns(uint64)
func (_VrfV25 *VrfV25CallerSession) SCurrentSubNonce() (uint64, error) {
	return _VrfV25.Contract.SCurrentSubNonce(&_VrfV25.CallOpts)
}

// SFallbackWeiPerUnitLink is a free data retrieval call binding the contract method 0x043bd6ae.
//
// Solidity: function s_fallbackWeiPerUnitLink() view returns(int256)
func (_VrfV25 *VrfV25Caller) SFallbackWeiPerUnitLink(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VrfV25.contract.Call(opts, &out, "s_fallbackWeiPerUnitLink")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SFallbackWeiPerUnitLink is a free data retrieval call binding the contract method 0x043bd6ae.
//
// Solidity: function s_fallbackWeiPerUnitLink() view returns(int256)
func (_VrfV25 *VrfV25Session) SFallbackWeiPerUnitLink() (*big.Int, error) {
	return _VrfV25.Contract.SFallbackWeiPerUnitLink(&_VrfV25.CallOpts)
}

// SFallbackWeiPerUnitLink is a free data retrieval call binding the contract method 0x043bd6ae.
//
// Solidity: function s_fallbackWeiPerUnitLink() view returns(int256)
func (_VrfV25 *VrfV25CallerSession) SFallbackWeiPerUnitLink() (*big.Int, error) {
	return _VrfV25.Contract.SFallbackWeiPerUnitLink(&_VrfV25.CallOpts)
}

// SProvingKeyHashes is a free data retrieval call binding the contract method 0xd98e620e.
//
// Solidity: function s_provingKeyHashes(uint256 ) view returns(bytes32)
func (_VrfV25 *VrfV25Caller) SProvingKeyHashes(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _VrfV25.contract.Call(opts, &out, "s_provingKeyHashes", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SProvingKeyHashes is a free data retrieval call binding the contract method 0xd98e620e.
//
// Solidity: function s_provingKeyHashes(uint256 ) view returns(bytes32)
func (_VrfV25 *VrfV25Session) SProvingKeyHashes(arg0 *big.Int) ([32]byte, error) {
	return _VrfV25.Contract.SProvingKeyHashes(&_VrfV25.CallOpts, arg0)
}

// SProvingKeyHashes is a free data retrieval call binding the contract method 0xd98e620e.
//
// Solidity: function s_provingKeyHashes(uint256 ) view returns(bytes32)
func (_VrfV25 *VrfV25CallerSession) SProvingKeyHashes(arg0 *big.Int) ([32]byte, error) {
	return _VrfV25.Contract.SProvingKeyHashes(&_VrfV25.CallOpts, arg0)
}

// SProvingKeys is a free data retrieval call binding the contract method 0xda2f2610.
//
// Solidity: function s_provingKeys(bytes32 ) view returns(bool exists, uint64 maxGas)
func (_VrfV25 *VrfV25Caller) SProvingKeys(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Exists bool
	MaxGas uint64
}, error) {
	var out []interface{}
	err := _VrfV25.contract.Call(opts, &out, "s_provingKeys", arg0)

	outstruct := new(struct {
		Exists bool
		MaxGas uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Exists = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.MaxGas = *abi.ConvertType(out[1], new(uint64)).(*uint64)

	return *outstruct, err

}

// SProvingKeys is a free data retrieval call binding the contract method 0xda2f2610.
//
// Solidity: function s_provingKeys(bytes32 ) view returns(bool exists, uint64 maxGas)
func (_VrfV25 *VrfV25Session) SProvingKeys(arg0 [32]byte) (struct {
	Exists bool
	MaxGas uint64
}, error) {
	return _VrfV25.Contract.SProvingKeys(&_VrfV25.CallOpts, arg0)
}

// SProvingKeys is a free data retrieval call binding the contract method 0xda2f2610.
//
// Solidity: function s_provingKeys(bytes32 ) view returns(bool exists, uint64 maxGas)
func (_VrfV25 *VrfV25CallerSession) SProvingKeys(arg0 [32]byte) (struct {
	Exists bool
	MaxGas uint64
}, error) {
	return _VrfV25.Contract.SProvingKeys(&_VrfV25.CallOpts, arg0)
}

// SRequestCommitments is a free data retrieval call binding the contract method 0xee9d2d38.
//
// Solidity: function s_requestCommitments(uint256 ) view returns(bytes32)
func (_VrfV25 *VrfV25Caller) SRequestCommitments(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _VrfV25.contract.Call(opts, &out, "s_requestCommitments", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SRequestCommitments is a free data retrieval call binding the contract method 0xee9d2d38.
//
// Solidity: function s_requestCommitments(uint256 ) view returns(bytes32)
func (_VrfV25 *VrfV25Session) SRequestCommitments(arg0 *big.Int) ([32]byte, error) {
	return _VrfV25.Contract.SRequestCommitments(&_VrfV25.CallOpts, arg0)
}

// SRequestCommitments is a free data retrieval call binding the contract method 0xee9d2d38.
//
// Solidity: function s_requestCommitments(uint256 ) view returns(bytes32)
func (_VrfV25 *VrfV25CallerSession) SRequestCommitments(arg0 *big.Int) ([32]byte, error) {
	return _VrfV25.Contract.SRequestCommitments(&_VrfV25.CallOpts, arg0)
}

// STotalBalance is a free data retrieval call binding the contract method 0x86fe91c7.
//
// Solidity: function s_totalBalance() view returns(uint96)
func (_VrfV25 *VrfV25Caller) STotalBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VrfV25.contract.Call(opts, &out, "s_totalBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// STotalBalance is a free data retrieval call binding the contract method 0x86fe91c7.
//
// Solidity: function s_totalBalance() view returns(uint96)
func (_VrfV25 *VrfV25Session) STotalBalance() (*big.Int, error) {
	return _VrfV25.Contract.STotalBalance(&_VrfV25.CallOpts)
}

// STotalBalance is a free data retrieval call binding the contract method 0x86fe91c7.
//
// Solidity: function s_totalBalance() view returns(uint96)
func (_VrfV25 *VrfV25CallerSession) STotalBalance() (*big.Int, error) {
	return _VrfV25.Contract.STotalBalance(&_VrfV25.CallOpts)
}

// STotalNativeBalance is a free data retrieval call binding the contract method 0x18e3dd27.
//
// Solidity: function s_totalNativeBalance() view returns(uint96)
func (_VrfV25 *VrfV25Caller) STotalNativeBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _VrfV25.contract.Call(opts, &out, "s_totalNativeBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// STotalNativeBalance is a free data retrieval call binding the contract method 0x18e3dd27.
//
// Solidity: function s_totalNativeBalance() view returns(uint96)
func (_VrfV25 *VrfV25Session) STotalNativeBalance() (*big.Int, error) {
	return _VrfV25.Contract.STotalNativeBalance(&_VrfV25.CallOpts)
}

// STotalNativeBalance is a free data retrieval call binding the contract method 0x18e3dd27.
//
// Solidity: function s_totalNativeBalance() view returns(uint96)
func (_VrfV25 *VrfV25CallerSession) STotalNativeBalance() (*big.Int, error) {
	return _VrfV25.Contract.STotalNativeBalance(&_VrfV25.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_VrfV25 *VrfV25Transactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VrfV25.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_VrfV25 *VrfV25Session) AcceptOwnership() (*types.Transaction, error) {
	return _VrfV25.Contract.AcceptOwnership(&_VrfV25.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_VrfV25 *VrfV25TransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _VrfV25.Contract.AcceptOwnership(&_VrfV25.TransactOpts)
}

// AcceptSubscriptionOwnerTransfer is a paid mutator transaction binding the contract method 0xb2a7cac5.
//
// Solidity: function acceptSubscriptionOwnerTransfer(uint256 subId) returns()
func (_VrfV25 *VrfV25Transactor) AcceptSubscriptionOwnerTransfer(opts *bind.TransactOpts, subId *big.Int) (*types.Transaction, error) {
	return _VrfV25.contract.Transact(opts, "acceptSubscriptionOwnerTransfer", subId)
}

// AcceptSubscriptionOwnerTransfer is a paid mutator transaction binding the contract method 0xb2a7cac5.
//
// Solidity: function acceptSubscriptionOwnerTransfer(uint256 subId) returns()
func (_VrfV25 *VrfV25Session) AcceptSubscriptionOwnerTransfer(subId *big.Int) (*types.Transaction, error) {
	return _VrfV25.Contract.AcceptSubscriptionOwnerTransfer(&_VrfV25.TransactOpts, subId)
}

// AcceptSubscriptionOwnerTransfer is a paid mutator transaction binding the contract method 0xb2a7cac5.
//
// Solidity: function acceptSubscriptionOwnerTransfer(uint256 subId) returns()
func (_VrfV25 *VrfV25TransactorSession) AcceptSubscriptionOwnerTransfer(subId *big.Int) (*types.Transaction, error) {
	return _VrfV25.Contract.AcceptSubscriptionOwnerTransfer(&_VrfV25.TransactOpts, subId)
}

// AddConsumer is a paid mutator transaction binding the contract method 0xbec4c08c.
//
// Solidity: function addConsumer(uint256 subId, address consumer) returns()
func (_VrfV25 *VrfV25Transactor) AddConsumer(opts *bind.TransactOpts, subId *big.Int, consumer common.Address) (*types.Transaction, error) {
	return _VrfV25.contract.Transact(opts, "addConsumer", subId, consumer)
}

// AddConsumer is a paid mutator transaction binding the contract method 0xbec4c08c.
//
// Solidity: function addConsumer(uint256 subId, address consumer) returns()
func (_VrfV25 *VrfV25Session) AddConsumer(subId *big.Int, consumer common.Address) (*types.Transaction, error) {
	return _VrfV25.Contract.AddConsumer(&_VrfV25.TransactOpts, subId, consumer)
}

// AddConsumer is a paid mutator transaction binding the contract method 0xbec4c08c.
//
// Solidity: function addConsumer(uint256 subId, address consumer) returns()
func (_VrfV25 *VrfV25TransactorSession) AddConsumer(subId *big.Int, consumer common.Address) (*types.Transaction, error) {
	return _VrfV25.Contract.AddConsumer(&_VrfV25.TransactOpts, subId, consumer)
}

// CancelSubscription is a paid mutator transaction binding the contract method 0x0ae09540.
//
// Solidity: function cancelSubscription(uint256 subId, address to) returns()
func (_VrfV25 *VrfV25Transactor) CancelSubscription(opts *bind.TransactOpts, subId *big.Int, to common.Address) (*types.Transaction, error) {
	return _VrfV25.contract.Transact(opts, "cancelSubscription", subId, to)
}

// CancelSubscription is a paid mutator transaction binding the contract method 0x0ae09540.
//
// Solidity: function cancelSubscription(uint256 subId, address to) returns()
func (_VrfV25 *VrfV25Session) CancelSubscription(subId *big.Int, to common.Address) (*types.Transaction, error) {
	return _VrfV25.Contract.CancelSubscription(&_VrfV25.TransactOpts, subId, to)
}

// CancelSubscription is a paid mutator transaction binding the contract method 0x0ae09540.
//
// Solidity: function cancelSubscription(uint256 subId, address to) returns()
func (_VrfV25 *VrfV25TransactorSession) CancelSubscription(subId *big.Int, to common.Address) (*types.Transaction, error) {
	return _VrfV25.Contract.CancelSubscription(&_VrfV25.TransactOpts, subId, to)
}

// CreateSubscription is a paid mutator transaction binding the contract method 0xa21a23e4.
//
// Solidity: function createSubscription() returns(uint256 subId)
func (_VrfV25 *VrfV25Transactor) CreateSubscription(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VrfV25.contract.Transact(opts, "createSubscription")
}

// CreateSubscription is a paid mutator transaction binding the contract method 0xa21a23e4.
//
// Solidity: function createSubscription() returns(uint256 subId)
func (_VrfV25 *VrfV25Session) CreateSubscription() (*types.Transaction, error) {
	return _VrfV25.Contract.CreateSubscription(&_VrfV25.TransactOpts)
}

// CreateSubscription is a paid mutator transaction binding the contract method 0xa21a23e4.
//
// Solidity: function createSubscription() returns(uint256 subId)
func (_VrfV25 *VrfV25TransactorSession) CreateSubscription() (*types.Transaction, error) {
	return _VrfV25.Contract.CreateSubscription(&_VrfV25.TransactOpts)
}

// DeregisterMigratableCoordinator is a paid mutator transaction binding the contract method 0x04104edb.
//
// Solidity: function deregisterMigratableCoordinator(address target) returns()
func (_VrfV25 *VrfV25Transactor) DeregisterMigratableCoordinator(opts *bind.TransactOpts, target common.Address) (*types.Transaction, error) {
	return _VrfV25.contract.Transact(opts, "deregisterMigratableCoordinator", target)
}

// DeregisterMigratableCoordinator is a paid mutator transaction binding the contract method 0x04104edb.
//
// Solidity: function deregisterMigratableCoordinator(address target) returns()
func (_VrfV25 *VrfV25Session) DeregisterMigratableCoordinator(target common.Address) (*types.Transaction, error) {
	return _VrfV25.Contract.DeregisterMigratableCoordinator(&_VrfV25.TransactOpts, target)
}

// DeregisterMigratableCoordinator is a paid mutator transaction binding the contract method 0x04104edb.
//
// Solidity: function deregisterMigratableCoordinator(address target) returns()
func (_VrfV25 *VrfV25TransactorSession) DeregisterMigratableCoordinator(target common.Address) (*types.Transaction, error) {
	return _VrfV25.Contract.DeregisterMigratableCoordinator(&_VrfV25.TransactOpts, target)
}

// DeregisterProvingKey is a paid mutator transaction binding the contract method 0x08821d58.
//
// Solidity: function deregisterProvingKey(uint256[2] publicProvingKey) returns()
func (_VrfV25 *VrfV25Transactor) DeregisterProvingKey(opts *bind.TransactOpts, publicProvingKey [2]*big.Int) (*types.Transaction, error) {
	return _VrfV25.contract.Transact(opts, "deregisterProvingKey", publicProvingKey)
}

// DeregisterProvingKey is a paid mutator transaction binding the contract method 0x08821d58.
//
// Solidity: function deregisterProvingKey(uint256[2] publicProvingKey) returns()
func (_VrfV25 *VrfV25Session) DeregisterProvingKey(publicProvingKey [2]*big.Int) (*types.Transaction, error) {
	return _VrfV25.Contract.DeregisterProvingKey(&_VrfV25.TransactOpts, publicProvingKey)
}

// DeregisterProvingKey is a paid mutator transaction binding the contract method 0x08821d58.
//
// Solidity: function deregisterProvingKey(uint256[2] publicProvingKey) returns()
func (_VrfV25 *VrfV25TransactorSession) DeregisterProvingKey(publicProvingKey [2]*big.Int) (*types.Transaction, error) {
	return _VrfV25.Contract.DeregisterProvingKey(&_VrfV25.TransactOpts, publicProvingKey)
}

// FulfillRandomWords is a paid mutator transaction binding the contract method 0x301f42e9.
//
// Solidity: function fulfillRandomWords((uint256[2],uint256[2],uint256,uint256,uint256,address,uint256[2],uint256[2],uint256) proof, (uint64,uint256,uint32,uint32,address,bytes) rc, bool onlyPremium) returns(uint96 payment)
func (_VrfV25 *VrfV25Transactor) FulfillRandomWords(opts *bind.TransactOpts, proof VRFProof, rc VRFTypesRequestCommitmentV2Plus, onlyPremium bool) (*types.Transaction, error) {
	return _VrfV25.contract.Transact(opts, "fulfillRandomWords", proof, rc, onlyPremium)
}

// FulfillRandomWords is a paid mutator transaction binding the contract method 0x301f42e9.
//
// Solidity: function fulfillRandomWords((uint256[2],uint256[2],uint256,uint256,uint256,address,uint256[2],uint256[2],uint256) proof, (uint64,uint256,uint32,uint32,address,bytes) rc, bool onlyPremium) returns(uint96 payment)
func (_VrfV25 *VrfV25Session) FulfillRandomWords(proof VRFProof, rc VRFTypesRequestCommitmentV2Plus, onlyPremium bool) (*types.Transaction, error) {
	return _VrfV25.Contract.FulfillRandomWords(&_VrfV25.TransactOpts, proof, rc, onlyPremium)
}

// FulfillRandomWords is a paid mutator transaction binding the contract method 0x301f42e9.
//
// Solidity: function fulfillRandomWords((uint256[2],uint256[2],uint256,uint256,uint256,address,uint256[2],uint256[2],uint256) proof, (uint64,uint256,uint32,uint32,address,bytes) rc, bool onlyPremium) returns(uint96 payment)
func (_VrfV25 *VrfV25TransactorSession) FulfillRandomWords(proof VRFProof, rc VRFTypesRequestCommitmentV2Plus, onlyPremium bool) (*types.Transaction, error) {
	return _VrfV25.Contract.FulfillRandomWords(&_VrfV25.TransactOpts, proof, rc, onlyPremium)
}

// FundSubscriptionWithNative is a paid mutator transaction binding the contract method 0x95b55cfc.
//
// Solidity: function fundSubscriptionWithNative(uint256 subId) payable returns()
func (_VrfV25 *VrfV25Transactor) FundSubscriptionWithNative(opts *bind.TransactOpts, subId *big.Int) (*types.Transaction, error) {
	return _VrfV25.contract.Transact(opts, "fundSubscriptionWithNative", subId)
}

// FundSubscriptionWithNative is a paid mutator transaction binding the contract method 0x95b55cfc.
//
// Solidity: function fundSubscriptionWithNative(uint256 subId) payable returns()
func (_VrfV25 *VrfV25Session) FundSubscriptionWithNative(subId *big.Int) (*types.Transaction, error) {
	return _VrfV25.Contract.FundSubscriptionWithNative(&_VrfV25.TransactOpts, subId)
}

// FundSubscriptionWithNative is a paid mutator transaction binding the contract method 0x95b55cfc.
//
// Solidity: function fundSubscriptionWithNative(uint256 subId) payable returns()
func (_VrfV25 *VrfV25TransactorSession) FundSubscriptionWithNative(subId *big.Int) (*types.Transaction, error) {
	return _VrfV25.Contract.FundSubscriptionWithNative(&_VrfV25.TransactOpts, subId)
}

// Migrate is a paid mutator transaction binding the contract method 0x405b84fa.
//
// Solidity: function migrate(uint256 subId, address newCoordinator) returns()
func (_VrfV25 *VrfV25Transactor) Migrate(opts *bind.TransactOpts, subId *big.Int, newCoordinator common.Address) (*types.Transaction, error) {
	return _VrfV25.contract.Transact(opts, "migrate", subId, newCoordinator)
}

// Migrate is a paid mutator transaction binding the contract method 0x405b84fa.
//
// Solidity: function migrate(uint256 subId, address newCoordinator) returns()
func (_VrfV25 *VrfV25Session) Migrate(subId *big.Int, newCoordinator common.Address) (*types.Transaction, error) {
	return _VrfV25.Contract.Migrate(&_VrfV25.TransactOpts, subId, newCoordinator)
}

// Migrate is a paid mutator transaction binding the contract method 0x405b84fa.
//
// Solidity: function migrate(uint256 subId, address newCoordinator) returns()
func (_VrfV25 *VrfV25TransactorSession) Migrate(subId *big.Int, newCoordinator common.Address) (*types.Transaction, error) {
	return _VrfV25.Contract.Migrate(&_VrfV25.TransactOpts, subId, newCoordinator)
}

// OnTokenTransfer is a paid mutator transaction binding the contract method 0xa4c0ed36.
//
// Solidity: function onTokenTransfer(address , uint256 amount, bytes data) returns()
func (_VrfV25 *VrfV25Transactor) OnTokenTransfer(opts *bind.TransactOpts, arg0 common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _VrfV25.contract.Transact(opts, "onTokenTransfer", arg0, amount, data)
}

// OnTokenTransfer is a paid mutator transaction binding the contract method 0xa4c0ed36.
//
// Solidity: function onTokenTransfer(address , uint256 amount, bytes data) returns()
func (_VrfV25 *VrfV25Session) OnTokenTransfer(arg0 common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _VrfV25.Contract.OnTokenTransfer(&_VrfV25.TransactOpts, arg0, amount, data)
}

// OnTokenTransfer is a paid mutator transaction binding the contract method 0xa4c0ed36.
//
// Solidity: function onTokenTransfer(address , uint256 amount, bytes data) returns()
func (_VrfV25 *VrfV25TransactorSession) OnTokenTransfer(arg0 common.Address, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _VrfV25.Contract.OnTokenTransfer(&_VrfV25.TransactOpts, arg0, amount, data)
}

// OwnerCancelSubscription is a paid mutator transaction binding the contract method 0xaa433aff.
//
// Solidity: function ownerCancelSubscription(uint256 subId) returns()
func (_VrfV25 *VrfV25Transactor) OwnerCancelSubscription(opts *bind.TransactOpts, subId *big.Int) (*types.Transaction, error) {
	return _VrfV25.contract.Transact(opts, "ownerCancelSubscription", subId)
}

// OwnerCancelSubscription is a paid mutator transaction binding the contract method 0xaa433aff.
//
// Solidity: function ownerCancelSubscription(uint256 subId) returns()
func (_VrfV25 *VrfV25Session) OwnerCancelSubscription(subId *big.Int) (*types.Transaction, error) {
	return _VrfV25.Contract.OwnerCancelSubscription(&_VrfV25.TransactOpts, subId)
}

// OwnerCancelSubscription is a paid mutator transaction binding the contract method 0xaa433aff.
//
// Solidity: function ownerCancelSubscription(uint256 subId) returns()
func (_VrfV25 *VrfV25TransactorSession) OwnerCancelSubscription(subId *big.Int) (*types.Transaction, error) {
	return _VrfV25.Contract.OwnerCancelSubscription(&_VrfV25.TransactOpts, subId)
}

// RecoverFunds is a paid mutator transaction binding the contract method 0xe72f6e30.
//
// Solidity: function recoverFunds(address to) returns()
func (_VrfV25 *VrfV25Transactor) RecoverFunds(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _VrfV25.contract.Transact(opts, "recoverFunds", to)
}

// RecoverFunds is a paid mutator transaction binding the contract method 0xe72f6e30.
//
// Solidity: function recoverFunds(address to) returns()
func (_VrfV25 *VrfV25Session) RecoverFunds(to common.Address) (*types.Transaction, error) {
	return _VrfV25.Contract.RecoverFunds(&_VrfV25.TransactOpts, to)
}

// RecoverFunds is a paid mutator transaction binding the contract method 0xe72f6e30.
//
// Solidity: function recoverFunds(address to) returns()
func (_VrfV25 *VrfV25TransactorSession) RecoverFunds(to common.Address) (*types.Transaction, error) {
	return _VrfV25.Contract.RecoverFunds(&_VrfV25.TransactOpts, to)
}

// RecoverNativeFunds is a paid mutator transaction binding the contract method 0x8402595e.
//
// Solidity: function recoverNativeFunds(address to) returns()
func (_VrfV25 *VrfV25Transactor) RecoverNativeFunds(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _VrfV25.contract.Transact(opts, "recoverNativeFunds", to)
}

// RecoverNativeFunds is a paid mutator transaction binding the contract method 0x8402595e.
//
// Solidity: function recoverNativeFunds(address to) returns()
func (_VrfV25 *VrfV25Session) RecoverNativeFunds(to common.Address) (*types.Transaction, error) {
	return _VrfV25.Contract.RecoverNativeFunds(&_VrfV25.TransactOpts, to)
}

// RecoverNativeFunds is a paid mutator transaction binding the contract method 0x8402595e.
//
// Solidity: function recoverNativeFunds(address to) returns()
func (_VrfV25 *VrfV25TransactorSession) RecoverNativeFunds(to common.Address) (*types.Transaction, error) {
	return _VrfV25.Contract.RecoverNativeFunds(&_VrfV25.TransactOpts, to)
}

// RegisterMigratableCoordinator is a paid mutator transaction binding the contract method 0x5d06b4ab.
//
// Solidity: function registerMigratableCoordinator(address target) returns()
func (_VrfV25 *VrfV25Transactor) RegisterMigratableCoordinator(opts *bind.TransactOpts, target common.Address) (*types.Transaction, error) {
	return _VrfV25.contract.Transact(opts, "registerMigratableCoordinator", target)
}

// RegisterMigratableCoordinator is a paid mutator transaction binding the contract method 0x5d06b4ab.
//
// Solidity: function registerMigratableCoordinator(address target) returns()
func (_VrfV25 *VrfV25Session) RegisterMigratableCoordinator(target common.Address) (*types.Transaction, error) {
	return _VrfV25.Contract.RegisterMigratableCoordinator(&_VrfV25.TransactOpts, target)
}

// RegisterMigratableCoordinator is a paid mutator transaction binding the contract method 0x5d06b4ab.
//
// Solidity: function registerMigratableCoordinator(address target) returns()
func (_VrfV25 *VrfV25TransactorSession) RegisterMigratableCoordinator(target common.Address) (*types.Transaction, error) {
	return _VrfV25.Contract.RegisterMigratableCoordinator(&_VrfV25.TransactOpts, target)
}

// RegisterProvingKey is a paid mutator transaction binding the contract method 0x7a5a2aef.
//
// Solidity: function registerProvingKey(uint256[2] publicProvingKey, uint64 maxGas) returns()
func (_VrfV25 *VrfV25Transactor) RegisterProvingKey(opts *bind.TransactOpts, publicProvingKey [2]*big.Int, maxGas uint64) (*types.Transaction, error) {
	return _VrfV25.contract.Transact(opts, "registerProvingKey", publicProvingKey, maxGas)
}

// RegisterProvingKey is a paid mutator transaction binding the contract method 0x7a5a2aef.
//
// Solidity: function registerProvingKey(uint256[2] publicProvingKey, uint64 maxGas) returns()
func (_VrfV25 *VrfV25Session) RegisterProvingKey(publicProvingKey [2]*big.Int, maxGas uint64) (*types.Transaction, error) {
	return _VrfV25.Contract.RegisterProvingKey(&_VrfV25.TransactOpts, publicProvingKey, maxGas)
}

// RegisterProvingKey is a paid mutator transaction binding the contract method 0x7a5a2aef.
//
// Solidity: function registerProvingKey(uint256[2] publicProvingKey, uint64 maxGas) returns()
func (_VrfV25 *VrfV25TransactorSession) RegisterProvingKey(publicProvingKey [2]*big.Int, maxGas uint64) (*types.Transaction, error) {
	return _VrfV25.Contract.RegisterProvingKey(&_VrfV25.TransactOpts, publicProvingKey, maxGas)
}

// RemoveConsumer is a paid mutator transaction binding the contract method 0xcb631797.
//
// Solidity: function removeConsumer(uint256 subId, address consumer) returns()
func (_VrfV25 *VrfV25Transactor) RemoveConsumer(opts *bind.TransactOpts, subId *big.Int, consumer common.Address) (*types.Transaction, error) {
	return _VrfV25.contract.Transact(opts, "removeConsumer", subId, consumer)
}

// RemoveConsumer is a paid mutator transaction binding the contract method 0xcb631797.
//
// Solidity: function removeConsumer(uint256 subId, address consumer) returns()
func (_VrfV25 *VrfV25Session) RemoveConsumer(subId *big.Int, consumer common.Address) (*types.Transaction, error) {
	return _VrfV25.Contract.RemoveConsumer(&_VrfV25.TransactOpts, subId, consumer)
}

// RemoveConsumer is a paid mutator transaction binding the contract method 0xcb631797.
//
// Solidity: function removeConsumer(uint256 subId, address consumer) returns()
func (_VrfV25 *VrfV25TransactorSession) RemoveConsumer(subId *big.Int, consumer common.Address) (*types.Transaction, error) {
	return _VrfV25.Contract.RemoveConsumer(&_VrfV25.TransactOpts, subId, consumer)
}

// RequestRandomWords is a paid mutator transaction binding the contract method 0x9b1c385e.
//
// Solidity: function requestRandomWords((bytes32,uint256,uint16,uint32,uint32,bytes) req) returns(uint256 requestId)
func (_VrfV25 *VrfV25Transactor) RequestRandomWords(opts *bind.TransactOpts, req VRFV2PlusClientRandomWordsRequest) (*types.Transaction, error) {
	return _VrfV25.contract.Transact(opts, "requestRandomWords", req)
}

// RequestRandomWords is a paid mutator transaction binding the contract method 0x9b1c385e.
//
// Solidity: function requestRandomWords((bytes32,uint256,uint16,uint32,uint32,bytes) req) returns(uint256 requestId)
func (_VrfV25 *VrfV25Session) RequestRandomWords(req VRFV2PlusClientRandomWordsRequest) (*types.Transaction, error) {
	return _VrfV25.Contract.RequestRandomWords(&_VrfV25.TransactOpts, req)
}

// RequestRandomWords is a paid mutator transaction binding the contract method 0x9b1c385e.
//
// Solidity: function requestRandomWords((bytes32,uint256,uint16,uint32,uint32,bytes) req) returns(uint256 requestId)
func (_VrfV25 *VrfV25TransactorSession) RequestRandomWords(req VRFV2PlusClientRandomWordsRequest) (*types.Transaction, error) {
	return _VrfV25.Contract.RequestRandomWords(&_VrfV25.TransactOpts, req)
}

// RequestSubscriptionOwnerTransfer is a paid mutator transaction binding the contract method 0xdac83d29.
//
// Solidity: function requestSubscriptionOwnerTransfer(uint256 subId, address newOwner) returns()
func (_VrfV25 *VrfV25Transactor) RequestSubscriptionOwnerTransfer(opts *bind.TransactOpts, subId *big.Int, newOwner common.Address) (*types.Transaction, error) {
	return _VrfV25.contract.Transact(opts, "requestSubscriptionOwnerTransfer", subId, newOwner)
}

// RequestSubscriptionOwnerTransfer is a paid mutator transaction binding the contract method 0xdac83d29.
//
// Solidity: function requestSubscriptionOwnerTransfer(uint256 subId, address newOwner) returns()
func (_VrfV25 *VrfV25Session) RequestSubscriptionOwnerTransfer(subId *big.Int, newOwner common.Address) (*types.Transaction, error) {
	return _VrfV25.Contract.RequestSubscriptionOwnerTransfer(&_VrfV25.TransactOpts, subId, newOwner)
}

// RequestSubscriptionOwnerTransfer is a paid mutator transaction binding the contract method 0xdac83d29.
//
// Solidity: function requestSubscriptionOwnerTransfer(uint256 subId, address newOwner) returns()
func (_VrfV25 *VrfV25TransactorSession) RequestSubscriptionOwnerTransfer(subId *big.Int, newOwner common.Address) (*types.Transaction, error) {
	return _VrfV25.Contract.RequestSubscriptionOwnerTransfer(&_VrfV25.TransactOpts, subId, newOwner)
}

// SetConfig is a paid mutator transaction binding the contract method 0xa63e0bfb.
//
// Solidity: function setConfig(uint16 minimumRequestConfirmations, uint32 maxGasLimit, uint32 stalenessSeconds, uint32 gasAfterPaymentCalculation, int256 fallbackWeiPerUnitLink, uint32 fulfillmentFlatFeeNativePPM, uint32 fulfillmentFlatFeeLinkDiscountPPM, uint8 nativePremiumPercentage, uint8 linkPremiumPercentage) returns()
func (_VrfV25 *VrfV25Transactor) SetConfig(opts *bind.TransactOpts, minimumRequestConfirmations uint16, maxGasLimit uint32, stalenessSeconds uint32, gasAfterPaymentCalculation uint32, fallbackWeiPerUnitLink *big.Int, fulfillmentFlatFeeNativePPM uint32, fulfillmentFlatFeeLinkDiscountPPM uint32, nativePremiumPercentage uint8, linkPremiumPercentage uint8) (*types.Transaction, error) {
	return _VrfV25.contract.Transact(opts, "setConfig", minimumRequestConfirmations, maxGasLimit, stalenessSeconds, gasAfterPaymentCalculation, fallbackWeiPerUnitLink, fulfillmentFlatFeeNativePPM, fulfillmentFlatFeeLinkDiscountPPM, nativePremiumPercentage, linkPremiumPercentage)
}

// SetConfig is a paid mutator transaction binding the contract method 0xa63e0bfb.
//
// Solidity: function setConfig(uint16 minimumRequestConfirmations, uint32 maxGasLimit, uint32 stalenessSeconds, uint32 gasAfterPaymentCalculation, int256 fallbackWeiPerUnitLink, uint32 fulfillmentFlatFeeNativePPM, uint32 fulfillmentFlatFeeLinkDiscountPPM, uint8 nativePremiumPercentage, uint8 linkPremiumPercentage) returns()
func (_VrfV25 *VrfV25Session) SetConfig(minimumRequestConfirmations uint16, maxGasLimit uint32, stalenessSeconds uint32, gasAfterPaymentCalculation uint32, fallbackWeiPerUnitLink *big.Int, fulfillmentFlatFeeNativePPM uint32, fulfillmentFlatFeeLinkDiscountPPM uint32, nativePremiumPercentage uint8, linkPremiumPercentage uint8) (*types.Transaction, error) {
	return _VrfV25.Contract.SetConfig(&_VrfV25.TransactOpts, minimumRequestConfirmations, maxGasLimit, stalenessSeconds, gasAfterPaymentCalculation, fallbackWeiPerUnitLink, fulfillmentFlatFeeNativePPM, fulfillmentFlatFeeLinkDiscountPPM, nativePremiumPercentage, linkPremiumPercentage)
}

// SetConfig is a paid mutator transaction binding the contract method 0xa63e0bfb.
//
// Solidity: function setConfig(uint16 minimumRequestConfirmations, uint32 maxGasLimit, uint32 stalenessSeconds, uint32 gasAfterPaymentCalculation, int256 fallbackWeiPerUnitLink, uint32 fulfillmentFlatFeeNativePPM, uint32 fulfillmentFlatFeeLinkDiscountPPM, uint8 nativePremiumPercentage, uint8 linkPremiumPercentage) returns()
func (_VrfV25 *VrfV25TransactorSession) SetConfig(minimumRequestConfirmations uint16, maxGasLimit uint32, stalenessSeconds uint32, gasAfterPaymentCalculation uint32, fallbackWeiPerUnitLink *big.Int, fulfillmentFlatFeeNativePPM uint32, fulfillmentFlatFeeLinkDiscountPPM uint32, nativePremiumPercentage uint8, linkPremiumPercentage uint8) (*types.Transaction, error) {
	return _VrfV25.Contract.SetConfig(&_VrfV25.TransactOpts, minimumRequestConfirmations, maxGasLimit, stalenessSeconds, gasAfterPaymentCalculation, fallbackWeiPerUnitLink, fulfillmentFlatFeeNativePPM, fulfillmentFlatFeeLinkDiscountPPM, nativePremiumPercentage, linkPremiumPercentage)
}

// SetLINKAndLINKNativeFeed is a paid mutator transaction binding the contract method 0x65982744.
//
// Solidity: function setLINKAndLINKNativeFeed(address link, address linkNativeFeed) returns()
func (_VrfV25 *VrfV25Transactor) SetLINKAndLINKNativeFeed(opts *bind.TransactOpts, link common.Address, linkNativeFeed common.Address) (*types.Transaction, error) {
	return _VrfV25.contract.Transact(opts, "setLINKAndLINKNativeFeed", link, linkNativeFeed)
}

// SetLINKAndLINKNativeFeed is a paid mutator transaction binding the contract method 0x65982744.
//
// Solidity: function setLINKAndLINKNativeFeed(address link, address linkNativeFeed) returns()
func (_VrfV25 *VrfV25Session) SetLINKAndLINKNativeFeed(link common.Address, linkNativeFeed common.Address) (*types.Transaction, error) {
	return _VrfV25.Contract.SetLINKAndLINKNativeFeed(&_VrfV25.TransactOpts, link, linkNativeFeed)
}

// SetLINKAndLINKNativeFeed is a paid mutator transaction binding the contract method 0x65982744.
//
// Solidity: function setLINKAndLINKNativeFeed(address link, address linkNativeFeed) returns()
func (_VrfV25 *VrfV25TransactorSession) SetLINKAndLINKNativeFeed(link common.Address, linkNativeFeed common.Address) (*types.Transaction, error) {
	return _VrfV25.Contract.SetLINKAndLINKNativeFeed(&_VrfV25.TransactOpts, link, linkNativeFeed)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address to) returns()
func (_VrfV25 *VrfV25Transactor) TransferOwnership(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _VrfV25.contract.Transact(opts, "transferOwnership", to)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address to) returns()
func (_VrfV25 *VrfV25Session) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _VrfV25.Contract.TransferOwnership(&_VrfV25.TransactOpts, to)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address to) returns()
func (_VrfV25 *VrfV25TransactorSession) TransferOwnership(to common.Address) (*types.Transaction, error) {
	return _VrfV25.Contract.TransferOwnership(&_VrfV25.TransactOpts, to)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address recipient) returns()
func (_VrfV25 *VrfV25Transactor) Withdraw(opts *bind.TransactOpts, recipient common.Address) (*types.Transaction, error) {
	return _VrfV25.contract.Transact(opts, "withdraw", recipient)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address recipient) returns()
func (_VrfV25 *VrfV25Session) Withdraw(recipient common.Address) (*types.Transaction, error) {
	return _VrfV25.Contract.Withdraw(&_VrfV25.TransactOpts, recipient)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address recipient) returns()
func (_VrfV25 *VrfV25TransactorSession) Withdraw(recipient common.Address) (*types.Transaction, error) {
	return _VrfV25.Contract.Withdraw(&_VrfV25.TransactOpts, recipient)
}

// WithdrawNative is a paid mutator transaction binding the contract method 0x2f622e6b.
//
// Solidity: function withdrawNative(address recipient) returns()
func (_VrfV25 *VrfV25Transactor) WithdrawNative(opts *bind.TransactOpts, recipient common.Address) (*types.Transaction, error) {
	return _VrfV25.contract.Transact(opts, "withdrawNative", recipient)
}

// WithdrawNative is a paid mutator transaction binding the contract method 0x2f622e6b.
//
// Solidity: function withdrawNative(address recipient) returns()
func (_VrfV25 *VrfV25Session) WithdrawNative(recipient common.Address) (*types.Transaction, error) {
	return _VrfV25.Contract.WithdrawNative(&_VrfV25.TransactOpts, recipient)
}

// WithdrawNative is a paid mutator transaction binding the contract method 0x2f622e6b.
//
// Solidity: function withdrawNative(address recipient) returns()
func (_VrfV25 *VrfV25TransactorSession) WithdrawNative(recipient common.Address) (*types.Transaction, error) {
	return _VrfV25.Contract.WithdrawNative(&_VrfV25.TransactOpts, recipient)
}

// VrfV25ConfigSetIterator is returned from FilterConfigSet and is used to iterate over the raw logs and unpacked data for ConfigSet events raised by the VrfV25 contract.
type VrfV25ConfigSetIterator struct {
	Event *VrfV25ConfigSet // Event containing the contract specifics and raw log

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
func (it *VrfV25ConfigSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VrfV25ConfigSet)
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
		it.Event = new(VrfV25ConfigSet)
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
func (it *VrfV25ConfigSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VrfV25ConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VrfV25ConfigSet represents a ConfigSet event raised by the VrfV25 contract.
type VrfV25ConfigSet struct {
	MinimumRequestConfirmations       uint16
	MaxGasLimit                       uint32
	StalenessSeconds                  uint32
	GasAfterPaymentCalculation        uint32
	FallbackWeiPerUnitLink            *big.Int
	FulfillmentFlatFeeNativePPM       uint32
	FulfillmentFlatFeeLinkDiscountPPM uint32
	NativePremiumPercentage           uint8
	LinkPremiumPercentage             uint8
	Raw                               types.Log // Blockchain specific contextual infos
}

// FilterConfigSet is a free log retrieval operation binding the contract event 0x2c6b6b12413678366b05b145c5f00745bdd00e739131ab5de82484a50c9d78b6.
//
// Solidity: event ConfigSet(uint16 minimumRequestConfirmations, uint32 maxGasLimit, uint32 stalenessSeconds, uint32 gasAfterPaymentCalculation, int256 fallbackWeiPerUnitLink, uint32 fulfillmentFlatFeeNativePPM, uint32 fulfillmentFlatFeeLinkDiscountPPM, uint8 nativePremiumPercentage, uint8 linkPremiumPercentage)
func (_VrfV25 *VrfV25Filterer) FilterConfigSet(opts *bind.FilterOpts) (*VrfV25ConfigSetIterator, error) {

	logs, sub, err := _VrfV25.contract.FilterLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return &VrfV25ConfigSetIterator{contract: _VrfV25.contract, event: "ConfigSet", logs: logs, sub: sub}, nil
}

// WatchConfigSet is a free log subscription operation binding the contract event 0x2c6b6b12413678366b05b145c5f00745bdd00e739131ab5de82484a50c9d78b6.
//
// Solidity: event ConfigSet(uint16 minimumRequestConfirmations, uint32 maxGasLimit, uint32 stalenessSeconds, uint32 gasAfterPaymentCalculation, int256 fallbackWeiPerUnitLink, uint32 fulfillmentFlatFeeNativePPM, uint32 fulfillmentFlatFeeLinkDiscountPPM, uint8 nativePremiumPercentage, uint8 linkPremiumPercentage)
func (_VrfV25 *VrfV25Filterer) WatchConfigSet(opts *bind.WatchOpts, sink chan<- *VrfV25ConfigSet) (event.Subscription, error) {

	logs, sub, err := _VrfV25.contract.WatchLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VrfV25ConfigSet)
				if err := _VrfV25.contract.UnpackLog(event, "ConfigSet", log); err != nil {
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

// ParseConfigSet is a log parse operation binding the contract event 0x2c6b6b12413678366b05b145c5f00745bdd00e739131ab5de82484a50c9d78b6.
//
// Solidity: event ConfigSet(uint16 minimumRequestConfirmations, uint32 maxGasLimit, uint32 stalenessSeconds, uint32 gasAfterPaymentCalculation, int256 fallbackWeiPerUnitLink, uint32 fulfillmentFlatFeeNativePPM, uint32 fulfillmentFlatFeeLinkDiscountPPM, uint8 nativePremiumPercentage, uint8 linkPremiumPercentage)
func (_VrfV25 *VrfV25Filterer) ParseConfigSet(log types.Log) (*VrfV25ConfigSet, error) {
	event := new(VrfV25ConfigSet)
	if err := _VrfV25.contract.UnpackLog(event, "ConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VrfV25CoordinatorDeregisteredIterator is returned from FilterCoordinatorDeregistered and is used to iterate over the raw logs and unpacked data for CoordinatorDeregistered events raised by the VrfV25 contract.
type VrfV25CoordinatorDeregisteredIterator struct {
	Event *VrfV25CoordinatorDeregistered // Event containing the contract specifics and raw log

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
func (it *VrfV25CoordinatorDeregisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VrfV25CoordinatorDeregistered)
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
		it.Event = new(VrfV25CoordinatorDeregistered)
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
func (it *VrfV25CoordinatorDeregisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VrfV25CoordinatorDeregisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VrfV25CoordinatorDeregistered represents a CoordinatorDeregistered event raised by the VrfV25 contract.
type VrfV25CoordinatorDeregistered struct {
	CoordinatorAddress common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterCoordinatorDeregistered is a free log retrieval operation binding the contract event 0xf80a1a97fd42251f3c33cda98635e7399253033a6774fe37cd3f650b5282af37.
//
// Solidity: event CoordinatorDeregistered(address coordinatorAddress)
func (_VrfV25 *VrfV25Filterer) FilterCoordinatorDeregistered(opts *bind.FilterOpts) (*VrfV25CoordinatorDeregisteredIterator, error) {

	logs, sub, err := _VrfV25.contract.FilterLogs(opts, "CoordinatorDeregistered")
	if err != nil {
		return nil, err
	}
	return &VrfV25CoordinatorDeregisteredIterator{contract: _VrfV25.contract, event: "CoordinatorDeregistered", logs: logs, sub: sub}, nil
}

// WatchCoordinatorDeregistered is a free log subscription operation binding the contract event 0xf80a1a97fd42251f3c33cda98635e7399253033a6774fe37cd3f650b5282af37.
//
// Solidity: event CoordinatorDeregistered(address coordinatorAddress)
func (_VrfV25 *VrfV25Filterer) WatchCoordinatorDeregistered(opts *bind.WatchOpts, sink chan<- *VrfV25CoordinatorDeregistered) (event.Subscription, error) {

	logs, sub, err := _VrfV25.contract.WatchLogs(opts, "CoordinatorDeregistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VrfV25CoordinatorDeregistered)
				if err := _VrfV25.contract.UnpackLog(event, "CoordinatorDeregistered", log); err != nil {
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

// ParseCoordinatorDeregistered is a log parse operation binding the contract event 0xf80a1a97fd42251f3c33cda98635e7399253033a6774fe37cd3f650b5282af37.
//
// Solidity: event CoordinatorDeregistered(address coordinatorAddress)
func (_VrfV25 *VrfV25Filterer) ParseCoordinatorDeregistered(log types.Log) (*VrfV25CoordinatorDeregistered, error) {
	event := new(VrfV25CoordinatorDeregistered)
	if err := _VrfV25.contract.UnpackLog(event, "CoordinatorDeregistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VrfV25CoordinatorRegisteredIterator is returned from FilterCoordinatorRegistered and is used to iterate over the raw logs and unpacked data for CoordinatorRegistered events raised by the VrfV25 contract.
type VrfV25CoordinatorRegisteredIterator struct {
	Event *VrfV25CoordinatorRegistered // Event containing the contract specifics and raw log

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
func (it *VrfV25CoordinatorRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VrfV25CoordinatorRegistered)
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
		it.Event = new(VrfV25CoordinatorRegistered)
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
func (it *VrfV25CoordinatorRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VrfV25CoordinatorRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VrfV25CoordinatorRegistered represents a CoordinatorRegistered event raised by the VrfV25 contract.
type VrfV25CoordinatorRegistered struct {
	CoordinatorAddress common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterCoordinatorRegistered is a free log retrieval operation binding the contract event 0xb7cabbfc11e66731fc77de0444614282023bcbd41d16781c753a431d0af01625.
//
// Solidity: event CoordinatorRegistered(address coordinatorAddress)
func (_VrfV25 *VrfV25Filterer) FilterCoordinatorRegistered(opts *bind.FilterOpts) (*VrfV25CoordinatorRegisteredIterator, error) {

	logs, sub, err := _VrfV25.contract.FilterLogs(opts, "CoordinatorRegistered")
	if err != nil {
		return nil, err
	}
	return &VrfV25CoordinatorRegisteredIterator{contract: _VrfV25.contract, event: "CoordinatorRegistered", logs: logs, sub: sub}, nil
}

// WatchCoordinatorRegistered is a free log subscription operation binding the contract event 0xb7cabbfc11e66731fc77de0444614282023bcbd41d16781c753a431d0af01625.
//
// Solidity: event CoordinatorRegistered(address coordinatorAddress)
func (_VrfV25 *VrfV25Filterer) WatchCoordinatorRegistered(opts *bind.WatchOpts, sink chan<- *VrfV25CoordinatorRegistered) (event.Subscription, error) {

	logs, sub, err := _VrfV25.contract.WatchLogs(opts, "CoordinatorRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VrfV25CoordinatorRegistered)
				if err := _VrfV25.contract.UnpackLog(event, "CoordinatorRegistered", log); err != nil {
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

// ParseCoordinatorRegistered is a log parse operation binding the contract event 0xb7cabbfc11e66731fc77de0444614282023bcbd41d16781c753a431d0af01625.
//
// Solidity: event CoordinatorRegistered(address coordinatorAddress)
func (_VrfV25 *VrfV25Filterer) ParseCoordinatorRegistered(log types.Log) (*VrfV25CoordinatorRegistered, error) {
	event := new(VrfV25CoordinatorRegistered)
	if err := _VrfV25.contract.UnpackLog(event, "CoordinatorRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VrfV25FallbackWeiPerUnitLinkUsedIterator is returned from FilterFallbackWeiPerUnitLinkUsed and is used to iterate over the raw logs and unpacked data for FallbackWeiPerUnitLinkUsed events raised by the VrfV25 contract.
type VrfV25FallbackWeiPerUnitLinkUsedIterator struct {
	Event *VrfV25FallbackWeiPerUnitLinkUsed // Event containing the contract specifics and raw log

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
func (it *VrfV25FallbackWeiPerUnitLinkUsedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VrfV25FallbackWeiPerUnitLinkUsed)
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
		it.Event = new(VrfV25FallbackWeiPerUnitLinkUsed)
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
func (it *VrfV25FallbackWeiPerUnitLinkUsedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VrfV25FallbackWeiPerUnitLinkUsedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VrfV25FallbackWeiPerUnitLinkUsed represents a FallbackWeiPerUnitLinkUsed event raised by the VrfV25 contract.
type VrfV25FallbackWeiPerUnitLinkUsed struct {
	RequestId              *big.Int
	FallbackWeiPerUnitLink *big.Int
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterFallbackWeiPerUnitLinkUsed is a free log retrieval operation binding the contract event 0x6ca648a381f22ead7e37773d934e64885dcf861fbfbb26c40354cbf0c4662d1a.
//
// Solidity: event FallbackWeiPerUnitLinkUsed(uint256 requestId, int256 fallbackWeiPerUnitLink)
func (_VrfV25 *VrfV25Filterer) FilterFallbackWeiPerUnitLinkUsed(opts *bind.FilterOpts) (*VrfV25FallbackWeiPerUnitLinkUsedIterator, error) {

	logs, sub, err := _VrfV25.contract.FilterLogs(opts, "FallbackWeiPerUnitLinkUsed")
	if err != nil {
		return nil, err
	}
	return &VrfV25FallbackWeiPerUnitLinkUsedIterator{contract: _VrfV25.contract, event: "FallbackWeiPerUnitLinkUsed", logs: logs, sub: sub}, nil
}

// WatchFallbackWeiPerUnitLinkUsed is a free log subscription operation binding the contract event 0x6ca648a381f22ead7e37773d934e64885dcf861fbfbb26c40354cbf0c4662d1a.
//
// Solidity: event FallbackWeiPerUnitLinkUsed(uint256 requestId, int256 fallbackWeiPerUnitLink)
func (_VrfV25 *VrfV25Filterer) WatchFallbackWeiPerUnitLinkUsed(opts *bind.WatchOpts, sink chan<- *VrfV25FallbackWeiPerUnitLinkUsed) (event.Subscription, error) {

	logs, sub, err := _VrfV25.contract.WatchLogs(opts, "FallbackWeiPerUnitLinkUsed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VrfV25FallbackWeiPerUnitLinkUsed)
				if err := _VrfV25.contract.UnpackLog(event, "FallbackWeiPerUnitLinkUsed", log); err != nil {
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

// ParseFallbackWeiPerUnitLinkUsed is a log parse operation binding the contract event 0x6ca648a381f22ead7e37773d934e64885dcf861fbfbb26c40354cbf0c4662d1a.
//
// Solidity: event FallbackWeiPerUnitLinkUsed(uint256 requestId, int256 fallbackWeiPerUnitLink)
func (_VrfV25 *VrfV25Filterer) ParseFallbackWeiPerUnitLinkUsed(log types.Log) (*VrfV25FallbackWeiPerUnitLinkUsed, error) {
	event := new(VrfV25FallbackWeiPerUnitLinkUsed)
	if err := _VrfV25.contract.UnpackLog(event, "FallbackWeiPerUnitLinkUsed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VrfV25FundsRecoveredIterator is returned from FilterFundsRecovered and is used to iterate over the raw logs and unpacked data for FundsRecovered events raised by the VrfV25 contract.
type VrfV25FundsRecoveredIterator struct {
	Event *VrfV25FundsRecovered // Event containing the contract specifics and raw log

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
func (it *VrfV25FundsRecoveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VrfV25FundsRecovered)
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
		it.Event = new(VrfV25FundsRecovered)
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
func (it *VrfV25FundsRecoveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VrfV25FundsRecoveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VrfV25FundsRecovered represents a FundsRecovered event raised by the VrfV25 contract.
type VrfV25FundsRecovered struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFundsRecovered is a free log retrieval operation binding the contract event 0x59bfc682b673f8cbf945f1e454df9334834abf7dfe7f92237ca29ecb9b436600.
//
// Solidity: event FundsRecovered(address to, uint256 amount)
func (_VrfV25 *VrfV25Filterer) FilterFundsRecovered(opts *bind.FilterOpts) (*VrfV25FundsRecoveredIterator, error) {

	logs, sub, err := _VrfV25.contract.FilterLogs(opts, "FundsRecovered")
	if err != nil {
		return nil, err
	}
	return &VrfV25FundsRecoveredIterator{contract: _VrfV25.contract, event: "FundsRecovered", logs: logs, sub: sub}, nil
}

// WatchFundsRecovered is a free log subscription operation binding the contract event 0x59bfc682b673f8cbf945f1e454df9334834abf7dfe7f92237ca29ecb9b436600.
//
// Solidity: event FundsRecovered(address to, uint256 amount)
func (_VrfV25 *VrfV25Filterer) WatchFundsRecovered(opts *bind.WatchOpts, sink chan<- *VrfV25FundsRecovered) (event.Subscription, error) {

	logs, sub, err := _VrfV25.contract.WatchLogs(opts, "FundsRecovered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VrfV25FundsRecovered)
				if err := _VrfV25.contract.UnpackLog(event, "FundsRecovered", log); err != nil {
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

// ParseFundsRecovered is a log parse operation binding the contract event 0x59bfc682b673f8cbf945f1e454df9334834abf7dfe7f92237ca29ecb9b436600.
//
// Solidity: event FundsRecovered(address to, uint256 amount)
func (_VrfV25 *VrfV25Filterer) ParseFundsRecovered(log types.Log) (*VrfV25FundsRecovered, error) {
	event := new(VrfV25FundsRecovered)
	if err := _VrfV25.contract.UnpackLog(event, "FundsRecovered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VrfV25MigrationCompletedIterator is returned from FilterMigrationCompleted and is used to iterate over the raw logs and unpacked data for MigrationCompleted events raised by the VrfV25 contract.
type VrfV25MigrationCompletedIterator struct {
	Event *VrfV25MigrationCompleted // Event containing the contract specifics and raw log

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
func (it *VrfV25MigrationCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VrfV25MigrationCompleted)
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
		it.Event = new(VrfV25MigrationCompleted)
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
func (it *VrfV25MigrationCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VrfV25MigrationCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VrfV25MigrationCompleted represents a MigrationCompleted event raised by the VrfV25 contract.
type VrfV25MigrationCompleted struct {
	NewCoordinator common.Address
	SubId          *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterMigrationCompleted is a free log retrieval operation binding the contract event 0xd63ca8cb945956747ee69bfdc3ea754c24a4caf7418db70e46052f7850be4187.
//
// Solidity: event MigrationCompleted(address newCoordinator, uint256 subId)
func (_VrfV25 *VrfV25Filterer) FilterMigrationCompleted(opts *bind.FilterOpts) (*VrfV25MigrationCompletedIterator, error) {

	logs, sub, err := _VrfV25.contract.FilterLogs(opts, "MigrationCompleted")
	if err != nil {
		return nil, err
	}
	return &VrfV25MigrationCompletedIterator{contract: _VrfV25.contract, event: "MigrationCompleted", logs: logs, sub: sub}, nil
}

// WatchMigrationCompleted is a free log subscription operation binding the contract event 0xd63ca8cb945956747ee69bfdc3ea754c24a4caf7418db70e46052f7850be4187.
//
// Solidity: event MigrationCompleted(address newCoordinator, uint256 subId)
func (_VrfV25 *VrfV25Filterer) WatchMigrationCompleted(opts *bind.WatchOpts, sink chan<- *VrfV25MigrationCompleted) (event.Subscription, error) {

	logs, sub, err := _VrfV25.contract.WatchLogs(opts, "MigrationCompleted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VrfV25MigrationCompleted)
				if err := _VrfV25.contract.UnpackLog(event, "MigrationCompleted", log); err != nil {
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

// ParseMigrationCompleted is a log parse operation binding the contract event 0xd63ca8cb945956747ee69bfdc3ea754c24a4caf7418db70e46052f7850be4187.
//
// Solidity: event MigrationCompleted(address newCoordinator, uint256 subId)
func (_VrfV25 *VrfV25Filterer) ParseMigrationCompleted(log types.Log) (*VrfV25MigrationCompleted, error) {
	event := new(VrfV25MigrationCompleted)
	if err := _VrfV25.contract.UnpackLog(event, "MigrationCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VrfV25NativeFundsRecoveredIterator is returned from FilterNativeFundsRecovered and is used to iterate over the raw logs and unpacked data for NativeFundsRecovered events raised by the VrfV25 contract.
type VrfV25NativeFundsRecoveredIterator struct {
	Event *VrfV25NativeFundsRecovered // Event containing the contract specifics and raw log

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
func (it *VrfV25NativeFundsRecoveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VrfV25NativeFundsRecovered)
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
		it.Event = new(VrfV25NativeFundsRecovered)
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
func (it *VrfV25NativeFundsRecoveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VrfV25NativeFundsRecoveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VrfV25NativeFundsRecovered represents a NativeFundsRecovered event raised by the VrfV25 contract.
type VrfV25NativeFundsRecovered struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNativeFundsRecovered is a free log retrieval operation binding the contract event 0x4aed7c8eed0496c8c19ea2681fcca25741c1602342e38b045d9f1e8e905d2e9c.
//
// Solidity: event NativeFundsRecovered(address to, uint256 amount)
func (_VrfV25 *VrfV25Filterer) FilterNativeFundsRecovered(opts *bind.FilterOpts) (*VrfV25NativeFundsRecoveredIterator, error) {

	logs, sub, err := _VrfV25.contract.FilterLogs(opts, "NativeFundsRecovered")
	if err != nil {
		return nil, err
	}
	return &VrfV25NativeFundsRecoveredIterator{contract: _VrfV25.contract, event: "NativeFundsRecovered", logs: logs, sub: sub}, nil
}

// WatchNativeFundsRecovered is a free log subscription operation binding the contract event 0x4aed7c8eed0496c8c19ea2681fcca25741c1602342e38b045d9f1e8e905d2e9c.
//
// Solidity: event NativeFundsRecovered(address to, uint256 amount)
func (_VrfV25 *VrfV25Filterer) WatchNativeFundsRecovered(opts *bind.WatchOpts, sink chan<- *VrfV25NativeFundsRecovered) (event.Subscription, error) {

	logs, sub, err := _VrfV25.contract.WatchLogs(opts, "NativeFundsRecovered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VrfV25NativeFundsRecovered)
				if err := _VrfV25.contract.UnpackLog(event, "NativeFundsRecovered", log); err != nil {
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

// ParseNativeFundsRecovered is a log parse operation binding the contract event 0x4aed7c8eed0496c8c19ea2681fcca25741c1602342e38b045d9f1e8e905d2e9c.
//
// Solidity: event NativeFundsRecovered(address to, uint256 amount)
func (_VrfV25 *VrfV25Filterer) ParseNativeFundsRecovered(log types.Log) (*VrfV25NativeFundsRecovered, error) {
	event := new(VrfV25NativeFundsRecovered)
	if err := _VrfV25.contract.UnpackLog(event, "NativeFundsRecovered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VrfV25OwnershipTransferRequestedIterator is returned from FilterOwnershipTransferRequested and is used to iterate over the raw logs and unpacked data for OwnershipTransferRequested events raised by the VrfV25 contract.
type VrfV25OwnershipTransferRequestedIterator struct {
	Event *VrfV25OwnershipTransferRequested // Event containing the contract specifics and raw log

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
func (it *VrfV25OwnershipTransferRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VrfV25OwnershipTransferRequested)
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
		it.Event = new(VrfV25OwnershipTransferRequested)
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
func (it *VrfV25OwnershipTransferRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VrfV25OwnershipTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VrfV25OwnershipTransferRequested represents a OwnershipTransferRequested event raised by the VrfV25 contract.
type VrfV25OwnershipTransferRequested struct {
	From common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferRequested is a free log retrieval operation binding the contract event 0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278.
//
// Solidity: event OwnershipTransferRequested(address indexed from, address indexed to)
func (_VrfV25 *VrfV25Filterer) FilterOwnershipTransferRequested(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VrfV25OwnershipTransferRequestedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VrfV25.contract.FilterLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &VrfV25OwnershipTransferRequestedIterator{contract: _VrfV25.contract, event: "OwnershipTransferRequested", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferRequested is a free log subscription operation binding the contract event 0xed8889f560326eb138920d842192f0eb3dd22b4f139c87a2c57538e05bae1278.
//
// Solidity: event OwnershipTransferRequested(address indexed from, address indexed to)
func (_VrfV25 *VrfV25Filterer) WatchOwnershipTransferRequested(opts *bind.WatchOpts, sink chan<- *VrfV25OwnershipTransferRequested, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VrfV25.contract.WatchLogs(opts, "OwnershipTransferRequested", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VrfV25OwnershipTransferRequested)
				if err := _VrfV25.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
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
func (_VrfV25 *VrfV25Filterer) ParseOwnershipTransferRequested(log types.Log) (*VrfV25OwnershipTransferRequested, error) {
	event := new(VrfV25OwnershipTransferRequested)
	if err := _VrfV25.contract.UnpackLog(event, "OwnershipTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VrfV25OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the VrfV25 contract.
type VrfV25OwnershipTransferredIterator struct {
	Event *VrfV25OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *VrfV25OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VrfV25OwnershipTransferred)
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
		it.Event = new(VrfV25OwnershipTransferred)
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
func (it *VrfV25OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VrfV25OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VrfV25OwnershipTransferred represents a OwnershipTransferred event raised by the VrfV25 contract.
type VrfV25OwnershipTransferred struct {
	From common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed from, address indexed to)
func (_VrfV25 *VrfV25Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*VrfV25OwnershipTransferredIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VrfV25.contract.FilterLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &VrfV25OwnershipTransferredIterator{contract: _VrfV25.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed from, address indexed to)
func (_VrfV25 *VrfV25Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VrfV25OwnershipTransferred, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _VrfV25.contract.WatchLogs(opts, "OwnershipTransferred", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VrfV25OwnershipTransferred)
				if err := _VrfV25.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_VrfV25 *VrfV25Filterer) ParseOwnershipTransferred(log types.Log) (*VrfV25OwnershipTransferred, error) {
	event := new(VrfV25OwnershipTransferred)
	if err := _VrfV25.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VrfV25ProvingKeyDeregisteredIterator is returned from FilterProvingKeyDeregistered and is used to iterate over the raw logs and unpacked data for ProvingKeyDeregistered events raised by the VrfV25 contract.
type VrfV25ProvingKeyDeregisteredIterator struct {
	Event *VrfV25ProvingKeyDeregistered // Event containing the contract specifics and raw log

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
func (it *VrfV25ProvingKeyDeregisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VrfV25ProvingKeyDeregistered)
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
		it.Event = new(VrfV25ProvingKeyDeregistered)
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
func (it *VrfV25ProvingKeyDeregisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VrfV25ProvingKeyDeregisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VrfV25ProvingKeyDeregistered represents a ProvingKeyDeregistered event raised by the VrfV25 contract.
type VrfV25ProvingKeyDeregistered struct {
	KeyHash [32]byte
	MaxGas  uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterProvingKeyDeregistered is a free log retrieval operation binding the contract event 0x9b6868e0eb737bcd72205360baa6bfd0ba4e4819a33ade2db384e8a8025639a5.
//
// Solidity: event ProvingKeyDeregistered(bytes32 keyHash, uint64 maxGas)
func (_VrfV25 *VrfV25Filterer) FilterProvingKeyDeregistered(opts *bind.FilterOpts) (*VrfV25ProvingKeyDeregisteredIterator, error) {

	logs, sub, err := _VrfV25.contract.FilterLogs(opts, "ProvingKeyDeregistered")
	if err != nil {
		return nil, err
	}
	return &VrfV25ProvingKeyDeregisteredIterator{contract: _VrfV25.contract, event: "ProvingKeyDeregistered", logs: logs, sub: sub}, nil
}

// WatchProvingKeyDeregistered is a free log subscription operation binding the contract event 0x9b6868e0eb737bcd72205360baa6bfd0ba4e4819a33ade2db384e8a8025639a5.
//
// Solidity: event ProvingKeyDeregistered(bytes32 keyHash, uint64 maxGas)
func (_VrfV25 *VrfV25Filterer) WatchProvingKeyDeregistered(opts *bind.WatchOpts, sink chan<- *VrfV25ProvingKeyDeregistered) (event.Subscription, error) {

	logs, sub, err := _VrfV25.contract.WatchLogs(opts, "ProvingKeyDeregistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VrfV25ProvingKeyDeregistered)
				if err := _VrfV25.contract.UnpackLog(event, "ProvingKeyDeregistered", log); err != nil {
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

// ParseProvingKeyDeregistered is a log parse operation binding the contract event 0x9b6868e0eb737bcd72205360baa6bfd0ba4e4819a33ade2db384e8a8025639a5.
//
// Solidity: event ProvingKeyDeregistered(bytes32 keyHash, uint64 maxGas)
func (_VrfV25 *VrfV25Filterer) ParseProvingKeyDeregistered(log types.Log) (*VrfV25ProvingKeyDeregistered, error) {
	event := new(VrfV25ProvingKeyDeregistered)
	if err := _VrfV25.contract.UnpackLog(event, "ProvingKeyDeregistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VrfV25ProvingKeyRegisteredIterator is returned from FilterProvingKeyRegistered and is used to iterate over the raw logs and unpacked data for ProvingKeyRegistered events raised by the VrfV25 contract.
type VrfV25ProvingKeyRegisteredIterator struct {
	Event *VrfV25ProvingKeyRegistered // Event containing the contract specifics and raw log

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
func (it *VrfV25ProvingKeyRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VrfV25ProvingKeyRegistered)
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
		it.Event = new(VrfV25ProvingKeyRegistered)
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
func (it *VrfV25ProvingKeyRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VrfV25ProvingKeyRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VrfV25ProvingKeyRegistered represents a ProvingKeyRegistered event raised by the VrfV25 contract.
type VrfV25ProvingKeyRegistered struct {
	KeyHash [32]byte
	MaxGas  uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterProvingKeyRegistered is a free log retrieval operation binding the contract event 0x9b911b2c240bfbef3b6a8f7ed6ee321d1258bb2a3fe6becab52ac1cd3210afd3.
//
// Solidity: event ProvingKeyRegistered(bytes32 keyHash, uint64 maxGas)
func (_VrfV25 *VrfV25Filterer) FilterProvingKeyRegistered(opts *bind.FilterOpts) (*VrfV25ProvingKeyRegisteredIterator, error) {

	logs, sub, err := _VrfV25.contract.FilterLogs(opts, "ProvingKeyRegistered")
	if err != nil {
		return nil, err
	}
	return &VrfV25ProvingKeyRegisteredIterator{contract: _VrfV25.contract, event: "ProvingKeyRegistered", logs: logs, sub: sub}, nil
}

// WatchProvingKeyRegistered is a free log subscription operation binding the contract event 0x9b911b2c240bfbef3b6a8f7ed6ee321d1258bb2a3fe6becab52ac1cd3210afd3.
//
// Solidity: event ProvingKeyRegistered(bytes32 keyHash, uint64 maxGas)
func (_VrfV25 *VrfV25Filterer) WatchProvingKeyRegistered(opts *bind.WatchOpts, sink chan<- *VrfV25ProvingKeyRegistered) (event.Subscription, error) {

	logs, sub, err := _VrfV25.contract.WatchLogs(opts, "ProvingKeyRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VrfV25ProvingKeyRegistered)
				if err := _VrfV25.contract.UnpackLog(event, "ProvingKeyRegistered", log); err != nil {
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

// ParseProvingKeyRegistered is a log parse operation binding the contract event 0x9b911b2c240bfbef3b6a8f7ed6ee321d1258bb2a3fe6becab52ac1cd3210afd3.
//
// Solidity: event ProvingKeyRegistered(bytes32 keyHash, uint64 maxGas)
func (_VrfV25 *VrfV25Filterer) ParseProvingKeyRegistered(log types.Log) (*VrfV25ProvingKeyRegistered, error) {
	event := new(VrfV25ProvingKeyRegistered)
	if err := _VrfV25.contract.UnpackLog(event, "ProvingKeyRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VrfV25RandomWordsFulfilledIterator is returned from FilterRandomWordsFulfilled and is used to iterate over the raw logs and unpacked data for RandomWordsFulfilled events raised by the VrfV25 contract.
type VrfV25RandomWordsFulfilledIterator struct {
	Event *VrfV25RandomWordsFulfilled // Event containing the contract specifics and raw log

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
func (it *VrfV25RandomWordsFulfilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VrfV25RandomWordsFulfilled)
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
		it.Event = new(VrfV25RandomWordsFulfilled)
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
func (it *VrfV25RandomWordsFulfilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VrfV25RandomWordsFulfilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VrfV25RandomWordsFulfilled represents a RandomWordsFulfilled event raised by the VrfV25 contract.
type VrfV25RandomWordsFulfilled struct {
	RequestId     *big.Int
	OutputSeed    *big.Int
	SubId         *big.Int
	Payment       *big.Int
	NativePayment bool
	Success       bool
	OnlyPremium   bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRandomWordsFulfilled is a free log retrieval operation binding the contract event 0xaeb4b4786571e184246d39587f659abf0e26f41f6a3358692250382c0cdb47b7.
//
// Solidity: event RandomWordsFulfilled(uint256 indexed requestId, uint256 outputSeed, uint256 indexed subId, uint96 payment, bool nativePayment, bool success, bool onlyPremium)
func (_VrfV25 *VrfV25Filterer) FilterRandomWordsFulfilled(opts *bind.FilterOpts, requestId []*big.Int, subId []*big.Int) (*VrfV25RandomWordsFulfilledIterator, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VrfV25.contract.FilterLogs(opts, "RandomWordsFulfilled", requestIdRule, subIdRule)
	if err != nil {
		return nil, err
	}
	return &VrfV25RandomWordsFulfilledIterator{contract: _VrfV25.contract, event: "RandomWordsFulfilled", logs: logs, sub: sub}, nil
}

// WatchRandomWordsFulfilled is a free log subscription operation binding the contract event 0xaeb4b4786571e184246d39587f659abf0e26f41f6a3358692250382c0cdb47b7.
//
// Solidity: event RandomWordsFulfilled(uint256 indexed requestId, uint256 outputSeed, uint256 indexed subId, uint96 payment, bool nativePayment, bool success, bool onlyPremium)
func (_VrfV25 *VrfV25Filterer) WatchRandomWordsFulfilled(opts *bind.WatchOpts, sink chan<- *VrfV25RandomWordsFulfilled, requestId []*big.Int, subId []*big.Int) (event.Subscription, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VrfV25.contract.WatchLogs(opts, "RandomWordsFulfilled", requestIdRule, subIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VrfV25RandomWordsFulfilled)
				if err := _VrfV25.contract.UnpackLog(event, "RandomWordsFulfilled", log); err != nil {
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

// ParseRandomWordsFulfilled is a log parse operation binding the contract event 0xaeb4b4786571e184246d39587f659abf0e26f41f6a3358692250382c0cdb47b7.
//
// Solidity: event RandomWordsFulfilled(uint256 indexed requestId, uint256 outputSeed, uint256 indexed subId, uint96 payment, bool nativePayment, bool success, bool onlyPremium)
func (_VrfV25 *VrfV25Filterer) ParseRandomWordsFulfilled(log types.Log) (*VrfV25RandomWordsFulfilled, error) {
	event := new(VrfV25RandomWordsFulfilled)
	if err := _VrfV25.contract.UnpackLog(event, "RandomWordsFulfilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VrfV25RandomWordsRequestedIterator is returned from FilterRandomWordsRequested and is used to iterate over the raw logs and unpacked data for RandomWordsRequested events raised by the VrfV25 contract.
type VrfV25RandomWordsRequestedIterator struct {
	Event *VrfV25RandomWordsRequested // Event containing the contract specifics and raw log

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
func (it *VrfV25RandomWordsRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VrfV25RandomWordsRequested)
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
		it.Event = new(VrfV25RandomWordsRequested)
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
func (it *VrfV25RandomWordsRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VrfV25RandomWordsRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VrfV25RandomWordsRequested represents a RandomWordsRequested event raised by the VrfV25 contract.
type VrfV25RandomWordsRequested struct {
	KeyHash                     [32]byte
	RequestId                   *big.Int
	PreSeed                     *big.Int
	SubId                       *big.Int
	MinimumRequestConfirmations uint16
	CallbackGasLimit            uint32
	NumWords                    uint32
	ExtraArgs                   []byte
	Sender                      common.Address
	Raw                         types.Log // Blockchain specific contextual infos
}

// FilterRandomWordsRequested is a free log retrieval operation binding the contract event 0xeb0e3652e0f44f417695e6e90f2f42c99b65cd7169074c5a654b16b9748c3a4e.
//
// Solidity: event RandomWordsRequested(bytes32 indexed keyHash, uint256 requestId, uint256 preSeed, uint256 indexed subId, uint16 minimumRequestConfirmations, uint32 callbackGasLimit, uint32 numWords, bytes extraArgs, address indexed sender)
func (_VrfV25 *VrfV25Filterer) FilterRandomWordsRequested(opts *bind.FilterOpts, keyHash [][32]byte, subId []*big.Int, sender []common.Address) (*VrfV25RandomWordsRequestedIterator, error) {

	var keyHashRule []interface{}
	for _, keyHashItem := range keyHash {
		keyHashRule = append(keyHashRule, keyHashItem)
	}

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _VrfV25.contract.FilterLogs(opts, "RandomWordsRequested", keyHashRule, subIdRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &VrfV25RandomWordsRequestedIterator{contract: _VrfV25.contract, event: "RandomWordsRequested", logs: logs, sub: sub}, nil
}

// WatchRandomWordsRequested is a free log subscription operation binding the contract event 0xeb0e3652e0f44f417695e6e90f2f42c99b65cd7169074c5a654b16b9748c3a4e.
//
// Solidity: event RandomWordsRequested(bytes32 indexed keyHash, uint256 requestId, uint256 preSeed, uint256 indexed subId, uint16 minimumRequestConfirmations, uint32 callbackGasLimit, uint32 numWords, bytes extraArgs, address indexed sender)
func (_VrfV25 *VrfV25Filterer) WatchRandomWordsRequested(opts *bind.WatchOpts, sink chan<- *VrfV25RandomWordsRequested, keyHash [][32]byte, subId []*big.Int, sender []common.Address) (event.Subscription, error) {

	var keyHashRule []interface{}
	for _, keyHashItem := range keyHash {
		keyHashRule = append(keyHashRule, keyHashItem)
	}

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _VrfV25.contract.WatchLogs(opts, "RandomWordsRequested", keyHashRule, subIdRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VrfV25RandomWordsRequested)
				if err := _VrfV25.contract.UnpackLog(event, "RandomWordsRequested", log); err != nil {
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

// ParseRandomWordsRequested is a log parse operation binding the contract event 0xeb0e3652e0f44f417695e6e90f2f42c99b65cd7169074c5a654b16b9748c3a4e.
//
// Solidity: event RandomWordsRequested(bytes32 indexed keyHash, uint256 requestId, uint256 preSeed, uint256 indexed subId, uint16 minimumRequestConfirmations, uint32 callbackGasLimit, uint32 numWords, bytes extraArgs, address indexed sender)
func (_VrfV25 *VrfV25Filterer) ParseRandomWordsRequested(log types.Log) (*VrfV25RandomWordsRequested, error) {
	event := new(VrfV25RandomWordsRequested)
	if err := _VrfV25.contract.UnpackLog(event, "RandomWordsRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VrfV25SubscriptionCanceledIterator is returned from FilterSubscriptionCanceled and is used to iterate over the raw logs and unpacked data for SubscriptionCanceled events raised by the VrfV25 contract.
type VrfV25SubscriptionCanceledIterator struct {
	Event *VrfV25SubscriptionCanceled // Event containing the contract specifics and raw log

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
func (it *VrfV25SubscriptionCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VrfV25SubscriptionCanceled)
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
		it.Event = new(VrfV25SubscriptionCanceled)
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
func (it *VrfV25SubscriptionCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VrfV25SubscriptionCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VrfV25SubscriptionCanceled represents a SubscriptionCanceled event raised by the VrfV25 contract.
type VrfV25SubscriptionCanceled struct {
	SubId        *big.Int
	To           common.Address
	AmountLink   *big.Int
	AmountNative *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSubscriptionCanceled is a free log retrieval operation binding the contract event 0x8c74ce8b8cf87f5eb001275c8be27eb34ea2b62bfab6814fcc62192bb63e81c4.
//
// Solidity: event SubscriptionCanceled(uint256 indexed subId, address to, uint256 amountLink, uint256 amountNative)
func (_VrfV25 *VrfV25Filterer) FilterSubscriptionCanceled(opts *bind.FilterOpts, subId []*big.Int) (*VrfV25SubscriptionCanceledIterator, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VrfV25.contract.FilterLogs(opts, "SubscriptionCanceled", subIdRule)
	if err != nil {
		return nil, err
	}
	return &VrfV25SubscriptionCanceledIterator{contract: _VrfV25.contract, event: "SubscriptionCanceled", logs: logs, sub: sub}, nil
}

// WatchSubscriptionCanceled is a free log subscription operation binding the contract event 0x8c74ce8b8cf87f5eb001275c8be27eb34ea2b62bfab6814fcc62192bb63e81c4.
//
// Solidity: event SubscriptionCanceled(uint256 indexed subId, address to, uint256 amountLink, uint256 amountNative)
func (_VrfV25 *VrfV25Filterer) WatchSubscriptionCanceled(opts *bind.WatchOpts, sink chan<- *VrfV25SubscriptionCanceled, subId []*big.Int) (event.Subscription, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VrfV25.contract.WatchLogs(opts, "SubscriptionCanceled", subIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VrfV25SubscriptionCanceled)
				if err := _VrfV25.contract.UnpackLog(event, "SubscriptionCanceled", log); err != nil {
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

// ParseSubscriptionCanceled is a log parse operation binding the contract event 0x8c74ce8b8cf87f5eb001275c8be27eb34ea2b62bfab6814fcc62192bb63e81c4.
//
// Solidity: event SubscriptionCanceled(uint256 indexed subId, address to, uint256 amountLink, uint256 amountNative)
func (_VrfV25 *VrfV25Filterer) ParseSubscriptionCanceled(log types.Log) (*VrfV25SubscriptionCanceled, error) {
	event := new(VrfV25SubscriptionCanceled)
	if err := _VrfV25.contract.UnpackLog(event, "SubscriptionCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VrfV25SubscriptionConsumerAddedIterator is returned from FilterSubscriptionConsumerAdded and is used to iterate over the raw logs and unpacked data for SubscriptionConsumerAdded events raised by the VrfV25 contract.
type VrfV25SubscriptionConsumerAddedIterator struct {
	Event *VrfV25SubscriptionConsumerAdded // Event containing the contract specifics and raw log

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
func (it *VrfV25SubscriptionConsumerAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VrfV25SubscriptionConsumerAdded)
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
		it.Event = new(VrfV25SubscriptionConsumerAdded)
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
func (it *VrfV25SubscriptionConsumerAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VrfV25SubscriptionConsumerAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VrfV25SubscriptionConsumerAdded represents a SubscriptionConsumerAdded event raised by the VrfV25 contract.
type VrfV25SubscriptionConsumerAdded struct {
	SubId    *big.Int
	Consumer common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSubscriptionConsumerAdded is a free log retrieval operation binding the contract event 0x1e980d04aa7648e205713e5e8ea3808672ac163d10936d36f91b2c88ac1575e1.
//
// Solidity: event SubscriptionConsumerAdded(uint256 indexed subId, address consumer)
func (_VrfV25 *VrfV25Filterer) FilterSubscriptionConsumerAdded(opts *bind.FilterOpts, subId []*big.Int) (*VrfV25SubscriptionConsumerAddedIterator, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VrfV25.contract.FilterLogs(opts, "SubscriptionConsumerAdded", subIdRule)
	if err != nil {
		return nil, err
	}
	return &VrfV25SubscriptionConsumerAddedIterator{contract: _VrfV25.contract, event: "SubscriptionConsumerAdded", logs: logs, sub: sub}, nil
}

// WatchSubscriptionConsumerAdded is a free log subscription operation binding the contract event 0x1e980d04aa7648e205713e5e8ea3808672ac163d10936d36f91b2c88ac1575e1.
//
// Solidity: event SubscriptionConsumerAdded(uint256 indexed subId, address consumer)
func (_VrfV25 *VrfV25Filterer) WatchSubscriptionConsumerAdded(opts *bind.WatchOpts, sink chan<- *VrfV25SubscriptionConsumerAdded, subId []*big.Int) (event.Subscription, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VrfV25.contract.WatchLogs(opts, "SubscriptionConsumerAdded", subIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VrfV25SubscriptionConsumerAdded)
				if err := _VrfV25.contract.UnpackLog(event, "SubscriptionConsumerAdded", log); err != nil {
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

// ParseSubscriptionConsumerAdded is a log parse operation binding the contract event 0x1e980d04aa7648e205713e5e8ea3808672ac163d10936d36f91b2c88ac1575e1.
//
// Solidity: event SubscriptionConsumerAdded(uint256 indexed subId, address consumer)
func (_VrfV25 *VrfV25Filterer) ParseSubscriptionConsumerAdded(log types.Log) (*VrfV25SubscriptionConsumerAdded, error) {
	event := new(VrfV25SubscriptionConsumerAdded)
	if err := _VrfV25.contract.UnpackLog(event, "SubscriptionConsumerAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VrfV25SubscriptionConsumerRemovedIterator is returned from FilterSubscriptionConsumerRemoved and is used to iterate over the raw logs and unpacked data for SubscriptionConsumerRemoved events raised by the VrfV25 contract.
type VrfV25SubscriptionConsumerRemovedIterator struct {
	Event *VrfV25SubscriptionConsumerRemoved // Event containing the contract specifics and raw log

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
func (it *VrfV25SubscriptionConsumerRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VrfV25SubscriptionConsumerRemoved)
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
		it.Event = new(VrfV25SubscriptionConsumerRemoved)
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
func (it *VrfV25SubscriptionConsumerRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VrfV25SubscriptionConsumerRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VrfV25SubscriptionConsumerRemoved represents a SubscriptionConsumerRemoved event raised by the VrfV25 contract.
type VrfV25SubscriptionConsumerRemoved struct {
	SubId    *big.Int
	Consumer common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSubscriptionConsumerRemoved is a free log retrieval operation binding the contract event 0x32158c6058347c1601b2d12bc696ac6901d8a9a9aa3ba10c27ab0a983e8425a7.
//
// Solidity: event SubscriptionConsumerRemoved(uint256 indexed subId, address consumer)
func (_VrfV25 *VrfV25Filterer) FilterSubscriptionConsumerRemoved(opts *bind.FilterOpts, subId []*big.Int) (*VrfV25SubscriptionConsumerRemovedIterator, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VrfV25.contract.FilterLogs(opts, "SubscriptionConsumerRemoved", subIdRule)
	if err != nil {
		return nil, err
	}
	return &VrfV25SubscriptionConsumerRemovedIterator{contract: _VrfV25.contract, event: "SubscriptionConsumerRemoved", logs: logs, sub: sub}, nil
}

// WatchSubscriptionConsumerRemoved is a free log subscription operation binding the contract event 0x32158c6058347c1601b2d12bc696ac6901d8a9a9aa3ba10c27ab0a983e8425a7.
//
// Solidity: event SubscriptionConsumerRemoved(uint256 indexed subId, address consumer)
func (_VrfV25 *VrfV25Filterer) WatchSubscriptionConsumerRemoved(opts *bind.WatchOpts, sink chan<- *VrfV25SubscriptionConsumerRemoved, subId []*big.Int) (event.Subscription, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VrfV25.contract.WatchLogs(opts, "SubscriptionConsumerRemoved", subIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VrfV25SubscriptionConsumerRemoved)
				if err := _VrfV25.contract.UnpackLog(event, "SubscriptionConsumerRemoved", log); err != nil {
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

// ParseSubscriptionConsumerRemoved is a log parse operation binding the contract event 0x32158c6058347c1601b2d12bc696ac6901d8a9a9aa3ba10c27ab0a983e8425a7.
//
// Solidity: event SubscriptionConsumerRemoved(uint256 indexed subId, address consumer)
func (_VrfV25 *VrfV25Filterer) ParseSubscriptionConsumerRemoved(log types.Log) (*VrfV25SubscriptionConsumerRemoved, error) {
	event := new(VrfV25SubscriptionConsumerRemoved)
	if err := _VrfV25.contract.UnpackLog(event, "SubscriptionConsumerRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VrfV25SubscriptionCreatedIterator is returned from FilterSubscriptionCreated and is used to iterate over the raw logs and unpacked data for SubscriptionCreated events raised by the VrfV25 contract.
type VrfV25SubscriptionCreatedIterator struct {
	Event *VrfV25SubscriptionCreated // Event containing the contract specifics and raw log

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
func (it *VrfV25SubscriptionCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VrfV25SubscriptionCreated)
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
		it.Event = new(VrfV25SubscriptionCreated)
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
func (it *VrfV25SubscriptionCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VrfV25SubscriptionCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VrfV25SubscriptionCreated represents a SubscriptionCreated event raised by the VrfV25 contract.
type VrfV25SubscriptionCreated struct {
	SubId *big.Int
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterSubscriptionCreated is a free log retrieval operation binding the contract event 0x1d3015d7ba850fa198dc7b1a3f5d42779313a681035f77c8c03764c61005518d.
//
// Solidity: event SubscriptionCreated(uint256 indexed subId, address owner)
func (_VrfV25 *VrfV25Filterer) FilterSubscriptionCreated(opts *bind.FilterOpts, subId []*big.Int) (*VrfV25SubscriptionCreatedIterator, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VrfV25.contract.FilterLogs(opts, "SubscriptionCreated", subIdRule)
	if err != nil {
		return nil, err
	}
	return &VrfV25SubscriptionCreatedIterator{contract: _VrfV25.contract, event: "SubscriptionCreated", logs: logs, sub: sub}, nil
}

// WatchSubscriptionCreated is a free log subscription operation binding the contract event 0x1d3015d7ba850fa198dc7b1a3f5d42779313a681035f77c8c03764c61005518d.
//
// Solidity: event SubscriptionCreated(uint256 indexed subId, address owner)
func (_VrfV25 *VrfV25Filterer) WatchSubscriptionCreated(opts *bind.WatchOpts, sink chan<- *VrfV25SubscriptionCreated, subId []*big.Int) (event.Subscription, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VrfV25.contract.WatchLogs(opts, "SubscriptionCreated", subIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VrfV25SubscriptionCreated)
				if err := _VrfV25.contract.UnpackLog(event, "SubscriptionCreated", log); err != nil {
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

// ParseSubscriptionCreated is a log parse operation binding the contract event 0x1d3015d7ba850fa198dc7b1a3f5d42779313a681035f77c8c03764c61005518d.
//
// Solidity: event SubscriptionCreated(uint256 indexed subId, address owner)
func (_VrfV25 *VrfV25Filterer) ParseSubscriptionCreated(log types.Log) (*VrfV25SubscriptionCreated, error) {
	event := new(VrfV25SubscriptionCreated)
	if err := _VrfV25.contract.UnpackLog(event, "SubscriptionCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VrfV25SubscriptionFundedIterator is returned from FilterSubscriptionFunded and is used to iterate over the raw logs and unpacked data for SubscriptionFunded events raised by the VrfV25 contract.
type VrfV25SubscriptionFundedIterator struct {
	Event *VrfV25SubscriptionFunded // Event containing the contract specifics and raw log

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
func (it *VrfV25SubscriptionFundedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VrfV25SubscriptionFunded)
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
		it.Event = new(VrfV25SubscriptionFunded)
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
func (it *VrfV25SubscriptionFundedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VrfV25SubscriptionFundedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VrfV25SubscriptionFunded represents a SubscriptionFunded event raised by the VrfV25 contract.
type VrfV25SubscriptionFunded struct {
	SubId      *big.Int
	OldBalance *big.Int
	NewBalance *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSubscriptionFunded is a free log retrieval operation binding the contract event 0x1ced9348ff549fceab2ac57cd3a9de38edaaab274b725ee82c23e8fc8c4eec7a.
//
// Solidity: event SubscriptionFunded(uint256 indexed subId, uint256 oldBalance, uint256 newBalance)
func (_VrfV25 *VrfV25Filterer) FilterSubscriptionFunded(opts *bind.FilterOpts, subId []*big.Int) (*VrfV25SubscriptionFundedIterator, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VrfV25.contract.FilterLogs(opts, "SubscriptionFunded", subIdRule)
	if err != nil {
		return nil, err
	}
	return &VrfV25SubscriptionFundedIterator{contract: _VrfV25.contract, event: "SubscriptionFunded", logs: logs, sub: sub}, nil
}

// WatchSubscriptionFunded is a free log subscription operation binding the contract event 0x1ced9348ff549fceab2ac57cd3a9de38edaaab274b725ee82c23e8fc8c4eec7a.
//
// Solidity: event SubscriptionFunded(uint256 indexed subId, uint256 oldBalance, uint256 newBalance)
func (_VrfV25 *VrfV25Filterer) WatchSubscriptionFunded(opts *bind.WatchOpts, sink chan<- *VrfV25SubscriptionFunded, subId []*big.Int) (event.Subscription, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VrfV25.contract.WatchLogs(opts, "SubscriptionFunded", subIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VrfV25SubscriptionFunded)
				if err := _VrfV25.contract.UnpackLog(event, "SubscriptionFunded", log); err != nil {
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

// ParseSubscriptionFunded is a log parse operation binding the contract event 0x1ced9348ff549fceab2ac57cd3a9de38edaaab274b725ee82c23e8fc8c4eec7a.
//
// Solidity: event SubscriptionFunded(uint256 indexed subId, uint256 oldBalance, uint256 newBalance)
func (_VrfV25 *VrfV25Filterer) ParseSubscriptionFunded(log types.Log) (*VrfV25SubscriptionFunded, error) {
	event := new(VrfV25SubscriptionFunded)
	if err := _VrfV25.contract.UnpackLog(event, "SubscriptionFunded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VrfV25SubscriptionFundedWithNativeIterator is returned from FilterSubscriptionFundedWithNative and is used to iterate over the raw logs and unpacked data for SubscriptionFundedWithNative events raised by the VrfV25 contract.
type VrfV25SubscriptionFundedWithNativeIterator struct {
	Event *VrfV25SubscriptionFundedWithNative // Event containing the contract specifics and raw log

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
func (it *VrfV25SubscriptionFundedWithNativeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VrfV25SubscriptionFundedWithNative)
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
		it.Event = new(VrfV25SubscriptionFundedWithNative)
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
func (it *VrfV25SubscriptionFundedWithNativeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VrfV25SubscriptionFundedWithNativeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VrfV25SubscriptionFundedWithNative represents a SubscriptionFundedWithNative event raised by the VrfV25 contract.
type VrfV25SubscriptionFundedWithNative struct {
	SubId            *big.Int
	OldNativeBalance *big.Int
	NewNativeBalance *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterSubscriptionFundedWithNative is a free log retrieval operation binding the contract event 0x7603b205d03651ee812f803fccde89f1012e545a9c99f0abfea9cedd0fd8e902.
//
// Solidity: event SubscriptionFundedWithNative(uint256 indexed subId, uint256 oldNativeBalance, uint256 newNativeBalance)
func (_VrfV25 *VrfV25Filterer) FilterSubscriptionFundedWithNative(opts *bind.FilterOpts, subId []*big.Int) (*VrfV25SubscriptionFundedWithNativeIterator, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VrfV25.contract.FilterLogs(opts, "SubscriptionFundedWithNative", subIdRule)
	if err != nil {
		return nil, err
	}
	return &VrfV25SubscriptionFundedWithNativeIterator{contract: _VrfV25.contract, event: "SubscriptionFundedWithNative", logs: logs, sub: sub}, nil
}

// WatchSubscriptionFundedWithNative is a free log subscription operation binding the contract event 0x7603b205d03651ee812f803fccde89f1012e545a9c99f0abfea9cedd0fd8e902.
//
// Solidity: event SubscriptionFundedWithNative(uint256 indexed subId, uint256 oldNativeBalance, uint256 newNativeBalance)
func (_VrfV25 *VrfV25Filterer) WatchSubscriptionFundedWithNative(opts *bind.WatchOpts, sink chan<- *VrfV25SubscriptionFundedWithNative, subId []*big.Int) (event.Subscription, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VrfV25.contract.WatchLogs(opts, "SubscriptionFundedWithNative", subIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VrfV25SubscriptionFundedWithNative)
				if err := _VrfV25.contract.UnpackLog(event, "SubscriptionFundedWithNative", log); err != nil {
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

// ParseSubscriptionFundedWithNative is a log parse operation binding the contract event 0x7603b205d03651ee812f803fccde89f1012e545a9c99f0abfea9cedd0fd8e902.
//
// Solidity: event SubscriptionFundedWithNative(uint256 indexed subId, uint256 oldNativeBalance, uint256 newNativeBalance)
func (_VrfV25 *VrfV25Filterer) ParseSubscriptionFundedWithNative(log types.Log) (*VrfV25SubscriptionFundedWithNative, error) {
	event := new(VrfV25SubscriptionFundedWithNative)
	if err := _VrfV25.contract.UnpackLog(event, "SubscriptionFundedWithNative", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VrfV25SubscriptionOwnerTransferRequestedIterator is returned from FilterSubscriptionOwnerTransferRequested and is used to iterate over the raw logs and unpacked data for SubscriptionOwnerTransferRequested events raised by the VrfV25 contract.
type VrfV25SubscriptionOwnerTransferRequestedIterator struct {
	Event *VrfV25SubscriptionOwnerTransferRequested // Event containing the contract specifics and raw log

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
func (it *VrfV25SubscriptionOwnerTransferRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VrfV25SubscriptionOwnerTransferRequested)
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
		it.Event = new(VrfV25SubscriptionOwnerTransferRequested)
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
func (it *VrfV25SubscriptionOwnerTransferRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VrfV25SubscriptionOwnerTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VrfV25SubscriptionOwnerTransferRequested represents a SubscriptionOwnerTransferRequested event raised by the VrfV25 contract.
type VrfV25SubscriptionOwnerTransferRequested struct {
	SubId *big.Int
	From  common.Address
	To    common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterSubscriptionOwnerTransferRequested is a free log retrieval operation binding the contract event 0x21a4dad170a6bf476c31bbcf4a16628295b0e450672eec25d7c93308e05344a1.
//
// Solidity: event SubscriptionOwnerTransferRequested(uint256 indexed subId, address from, address to)
func (_VrfV25 *VrfV25Filterer) FilterSubscriptionOwnerTransferRequested(opts *bind.FilterOpts, subId []*big.Int) (*VrfV25SubscriptionOwnerTransferRequestedIterator, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VrfV25.contract.FilterLogs(opts, "SubscriptionOwnerTransferRequested", subIdRule)
	if err != nil {
		return nil, err
	}
	return &VrfV25SubscriptionOwnerTransferRequestedIterator{contract: _VrfV25.contract, event: "SubscriptionOwnerTransferRequested", logs: logs, sub: sub}, nil
}

// WatchSubscriptionOwnerTransferRequested is a free log subscription operation binding the contract event 0x21a4dad170a6bf476c31bbcf4a16628295b0e450672eec25d7c93308e05344a1.
//
// Solidity: event SubscriptionOwnerTransferRequested(uint256 indexed subId, address from, address to)
func (_VrfV25 *VrfV25Filterer) WatchSubscriptionOwnerTransferRequested(opts *bind.WatchOpts, sink chan<- *VrfV25SubscriptionOwnerTransferRequested, subId []*big.Int) (event.Subscription, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VrfV25.contract.WatchLogs(opts, "SubscriptionOwnerTransferRequested", subIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VrfV25SubscriptionOwnerTransferRequested)
				if err := _VrfV25.contract.UnpackLog(event, "SubscriptionOwnerTransferRequested", log); err != nil {
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

// ParseSubscriptionOwnerTransferRequested is a log parse operation binding the contract event 0x21a4dad170a6bf476c31bbcf4a16628295b0e450672eec25d7c93308e05344a1.
//
// Solidity: event SubscriptionOwnerTransferRequested(uint256 indexed subId, address from, address to)
func (_VrfV25 *VrfV25Filterer) ParseSubscriptionOwnerTransferRequested(log types.Log) (*VrfV25SubscriptionOwnerTransferRequested, error) {
	event := new(VrfV25SubscriptionOwnerTransferRequested)
	if err := _VrfV25.contract.UnpackLog(event, "SubscriptionOwnerTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VrfV25SubscriptionOwnerTransferredIterator is returned from FilterSubscriptionOwnerTransferred and is used to iterate over the raw logs and unpacked data for SubscriptionOwnerTransferred events raised by the VrfV25 contract.
type VrfV25SubscriptionOwnerTransferredIterator struct {
	Event *VrfV25SubscriptionOwnerTransferred // Event containing the contract specifics and raw log

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
func (it *VrfV25SubscriptionOwnerTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VrfV25SubscriptionOwnerTransferred)
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
		it.Event = new(VrfV25SubscriptionOwnerTransferred)
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
func (it *VrfV25SubscriptionOwnerTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VrfV25SubscriptionOwnerTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VrfV25SubscriptionOwnerTransferred represents a SubscriptionOwnerTransferred event raised by the VrfV25 contract.
type VrfV25SubscriptionOwnerTransferred struct {
	SubId *big.Int
	From  common.Address
	To    common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterSubscriptionOwnerTransferred is a free log retrieval operation binding the contract event 0xd4114ab6e9af9f597c52041f32d62dc57c5c4e4c0d4427006069635e216c9386.
//
// Solidity: event SubscriptionOwnerTransferred(uint256 indexed subId, address from, address to)
func (_VrfV25 *VrfV25Filterer) FilterSubscriptionOwnerTransferred(opts *bind.FilterOpts, subId []*big.Int) (*VrfV25SubscriptionOwnerTransferredIterator, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VrfV25.contract.FilterLogs(opts, "SubscriptionOwnerTransferred", subIdRule)
	if err != nil {
		return nil, err
	}
	return &VrfV25SubscriptionOwnerTransferredIterator{contract: _VrfV25.contract, event: "SubscriptionOwnerTransferred", logs: logs, sub: sub}, nil
}

// WatchSubscriptionOwnerTransferred is a free log subscription operation binding the contract event 0xd4114ab6e9af9f597c52041f32d62dc57c5c4e4c0d4427006069635e216c9386.
//
// Solidity: event SubscriptionOwnerTransferred(uint256 indexed subId, address from, address to)
func (_VrfV25 *VrfV25Filterer) WatchSubscriptionOwnerTransferred(opts *bind.WatchOpts, sink chan<- *VrfV25SubscriptionOwnerTransferred, subId []*big.Int) (event.Subscription, error) {

	var subIdRule []interface{}
	for _, subIdItem := range subId {
		subIdRule = append(subIdRule, subIdItem)
	}

	logs, sub, err := _VrfV25.contract.WatchLogs(opts, "SubscriptionOwnerTransferred", subIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VrfV25SubscriptionOwnerTransferred)
				if err := _VrfV25.contract.UnpackLog(event, "SubscriptionOwnerTransferred", log); err != nil {
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

// ParseSubscriptionOwnerTransferred is a log parse operation binding the contract event 0xd4114ab6e9af9f597c52041f32d62dc57c5c4e4c0d4427006069635e216c9386.
//
// Solidity: event SubscriptionOwnerTransferred(uint256 indexed subId, address from, address to)
func (_VrfV25 *VrfV25Filterer) ParseSubscriptionOwnerTransferred(log types.Log) (*VrfV25SubscriptionOwnerTransferred, error) {
	event := new(VrfV25SubscriptionOwnerTransferred)
	if err := _VrfV25.contract.UnpackLog(event, "SubscriptionOwnerTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
