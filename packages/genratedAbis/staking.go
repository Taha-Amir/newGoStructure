// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package staking

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

// AccessPassV4userDetails is an auto generated low-level Go binding around an user-defined struct.
type AccessPassV4userDetails struct {
	UserAddress            common.Address
	TokensDeposited        *big.Int
	AccessPassedTokenPrice *big.Int
	LatestTrxPrice         *big.Int
	TotalPriceBUSD         *big.Int
	TrxTime                *big.Int
	WorthIncrease          *big.Int
	HasDeposited           bool
	Status                 bool
	AutoWithdraw           bool
}

// StakingMetaData contains all meta data concerning the Staking contract.
var StakingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_curToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_busd\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BalanceClear\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"bnbSwapped\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"TokenAddressCleared\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"Amount\",\"type\":\"uint256\"}],\"name\":\"clearToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"tokensBought\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BUSD\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CheckAutoWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"Users\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOutMin\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path\",\"type\":\"address[]\"}],\"name\":\"accessPassTokens\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"addresses\",\"type\":\"address[]\"}],\"name\":\"autoWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentDepositBusdLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentTokenDepositPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentTokenPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"depositTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feePercentage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amountIn\",\"type\":\"uint256\"}],\"name\":\"getAmountOutMin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"inActiveUserTransferHistory\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"TokensDeposited\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accessPassedTokenPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"latestTrxPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalPriceBUSD\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"trxTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"worthIncrease\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"hasDeposited\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"autoWithdraw\",\"type\":\"bool\"}],\"internalType\":\"structAccessPassV4.userDetails[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"isEligibleToWithdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"current\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"withdraw_possible\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountPercentage\",\"type\":\"uint256\"}],\"name\":\"recoverStuckBnb\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokens\",\"type\":\"uint256\"}],\"name\":\"recoverStuckToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"busdAmount\",\"type\":\"uint256\"}],\"name\":\"setDepositLimitInBusd\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"percent\",\"type\":\"uint256\"}],\"name\":\"subtractPercent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenWorthIncrease\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"inactiveUsers\",\"type\":\"address[]\"}],\"name\":\"transferUsersBalanceToOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"uniswapV2Router\",\"outputs\":[{\"internalType\":\"contractIUniswapV2Router02\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"users\",\"type\":\"address[]\"},{\"internalType\":\"bool\",\"name\":\"autoWithdrawStatus\",\"type\":\"bool\"}],\"name\":\"updateAutoWithdrawStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newToken\",\"type\":\"address\"}],\"name\":\"updateCurrentToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"autoWithdrawStatus\",\"type\":\"bool\"}],\"name\":\"updateOwnAutoWithdrawStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"percentage\",\"type\":\"uint256\"}],\"name\":\"updateTokenWorthIncrease\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"TokensDeposited\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accessPassedTokenPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"latestTrxPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"trxTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"worthIncrease\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"hasDeposited\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"autoWithdrawStatus\",\"type\":\"bool\"}],\"name\":\"updateUserDetails\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"users\",\"type\":\"address[]\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"}],\"name\":\"updateUserStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_isEnabled\",\"type\":\"bool\"}],\"name\":\"updatefeeEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_feepercent\",\"type\":\"uint256\"}],\"name\":\"updatefeePercentage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"userDepositBalane\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"userHistory\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"TokensDeposited\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accessPassedTokenPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"latestTrxPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalPriceBUSD\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"trxTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"worthIncrease\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"hasDeposited\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"autoWithdraw\",\"type\":\"bool\"}],\"internalType\":\"structAccessPassV4.userDetails[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userProfile\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"TokensDeposited\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accessPassedTokenPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"latestTrxPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalPriceBUSD\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"trxTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"worthIncrease\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"hasDeposited\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"autoWithdraw\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"userStatus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"userWithdrawProfit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"usersHistory\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"userAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"TokensDeposited\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accessPassedTokenPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"latestTrxPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalPriceBUSD\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"trxTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"worthIncrease\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"hasDeposited\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"autoWithdraw\",\"type\":\"bool\"}],\"internalType\":\"structAccessPassV4.userDetails[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// StakingABI is the input ABI used to generate the binding from.
// Deprecated: Use StakingMetaData.ABI instead.
var StakingABI = StakingMetaData.ABI

// Staking is an auto generated Go binding around an Ethereum contract.
type Staking struct {
	StakingCaller     // Read-only binding to the contract
	StakingTransactor // Write-only binding to the contract
	StakingFilterer   // Log filterer for contract events
}

// StakingCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakingSession struct {
	Contract     *Staking          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakingCallerSession struct {
	Contract *StakingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// StakingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakingTransactorSession struct {
	Contract     *StakingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// StakingRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakingRaw struct {
	Contract *Staking // Generic contract binding to access the raw methods on
}

// StakingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakingCallerRaw struct {
	Contract *StakingCaller // Generic read-only contract binding to access the raw methods on
}

// StakingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakingTransactorRaw struct {
	Contract *StakingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStaking creates a new instance of Staking, bound to a specific deployed contract.
func NewStaking(address common.Address, backend bind.ContractBackend) (*Staking, error) {
	contract, err := bindStaking(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Staking{StakingCaller: StakingCaller{contract: contract}, StakingTransactor: StakingTransactor{contract: contract}, StakingFilterer: StakingFilterer{contract: contract}}, nil
}

// NewStakingCaller creates a new read-only instance of Staking, bound to a specific deployed contract.
func NewStakingCaller(address common.Address, caller bind.ContractCaller) (*StakingCaller, error) {
	contract, err := bindStaking(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakingCaller{contract: contract}, nil
}

// NewStakingTransactor creates a new write-only instance of Staking, bound to a specific deployed contract.
func NewStakingTransactor(address common.Address, transactor bind.ContractTransactor) (*StakingTransactor, error) {
	contract, err := bindStaking(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakingTransactor{contract: contract}, nil
}

// NewStakingFilterer creates a new log filterer instance of Staking, bound to a specific deployed contract.
func NewStakingFilterer(address common.Address, filterer bind.ContractFilterer) (*StakingFilterer, error) {
	contract, err := bindStaking(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakingFilterer{contract: contract}, nil
}

// bindStaking binds a generic wrapper to an already deployed contract.
func bindStaking(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StakingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Staking *StakingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Staking.Contract.StakingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Staking *StakingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.Contract.StakingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Staking *StakingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Staking.Contract.StakingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Staking *StakingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Staking.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Staking *StakingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Staking *StakingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Staking.Contract.contract.Transact(opts, method, params...)
}

// BUSD is a free data retrieval call binding the contract method 0x484f4ea9.
//
// Solidity: function BUSD() view returns(address)
func (_Staking *StakingCaller) BUSD(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "BUSD")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BUSD is a free data retrieval call binding the contract method 0x484f4ea9.
//
// Solidity: function BUSD() view returns(address)
func (_Staking *StakingSession) BUSD() (common.Address, error) {
	return _Staking.Contract.BUSD(&_Staking.CallOpts)
}

// BUSD is a free data retrieval call binding the contract method 0x484f4ea9.
//
// Solidity: function BUSD() view returns(address)
func (_Staking *StakingCallerSession) BUSD() (common.Address, error) {
	return _Staking.Contract.BUSD(&_Staking.CallOpts)
}

// CheckAutoWithdraw is a free data retrieval call binding the contract method 0x3d9102a3.
//
// Solidity: function CheckAutoWithdraw() view returns(uint256 count, address[] Users)
func (_Staking *StakingCaller) CheckAutoWithdraw(opts *bind.CallOpts) (struct {
	Count *big.Int
	Users []common.Address
}, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "CheckAutoWithdraw")

	outstruct := new(struct {
		Count *big.Int
		Users []common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Count = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Users = *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)

	return *outstruct, err

}

// CheckAutoWithdraw is a free data retrieval call binding the contract method 0x3d9102a3.
//
// Solidity: function CheckAutoWithdraw() view returns(uint256 count, address[] Users)
func (_Staking *StakingSession) CheckAutoWithdraw() (struct {
	Count *big.Int
	Users []common.Address
}, error) {
	return _Staking.Contract.CheckAutoWithdraw(&_Staking.CallOpts)
}

// CheckAutoWithdraw is a free data retrieval call binding the contract method 0x3d9102a3.
//
// Solidity: function CheckAutoWithdraw() view returns(uint256 count, address[] Users)
func (_Staking *StakingCallerSession) CheckAutoWithdraw() (struct {
	Count *big.Int
	Users []common.Address
}, error) {
	return _Staking.Contract.CheckAutoWithdraw(&_Staking.CallOpts)
}

// CurrentDepositBusdLimit is a free data retrieval call binding the contract method 0x5960150b.
//
// Solidity: function currentDepositBusdLimit() view returns(uint256)
func (_Staking *StakingCaller) CurrentDepositBusdLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "currentDepositBusdLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentDepositBusdLimit is a free data retrieval call binding the contract method 0x5960150b.
//
// Solidity: function currentDepositBusdLimit() view returns(uint256)
func (_Staking *StakingSession) CurrentDepositBusdLimit() (*big.Int, error) {
	return _Staking.Contract.CurrentDepositBusdLimit(&_Staking.CallOpts)
}

// CurrentDepositBusdLimit is a free data retrieval call binding the contract method 0x5960150b.
//
// Solidity: function currentDepositBusdLimit() view returns(uint256)
func (_Staking *StakingCallerSession) CurrentDepositBusdLimit() (*big.Int, error) {
	return _Staking.Contract.CurrentDepositBusdLimit(&_Staking.CallOpts)
}

// CurrentToken is a free data retrieval call binding the contract method 0x836c081d.
//
// Solidity: function currentToken() view returns(address)
func (_Staking *StakingCaller) CurrentToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "currentToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CurrentToken is a free data retrieval call binding the contract method 0x836c081d.
//
// Solidity: function currentToken() view returns(address)
func (_Staking *StakingSession) CurrentToken() (common.Address, error) {
	return _Staking.Contract.CurrentToken(&_Staking.CallOpts)
}

// CurrentToken is a free data retrieval call binding the contract method 0x836c081d.
//
// Solidity: function currentToken() view returns(address)
func (_Staking *StakingCallerSession) CurrentToken() (common.Address, error) {
	return _Staking.Contract.CurrentToken(&_Staking.CallOpts)
}

// CurrentTokenDepositPrice is a free data retrieval call binding the contract method 0x6449cb68.
//
// Solidity: function currentTokenDepositPrice() view returns(uint256)
func (_Staking *StakingCaller) CurrentTokenDepositPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "currentTokenDepositPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentTokenDepositPrice is a free data retrieval call binding the contract method 0x6449cb68.
//
// Solidity: function currentTokenDepositPrice() view returns(uint256)
func (_Staking *StakingSession) CurrentTokenDepositPrice() (*big.Int, error) {
	return _Staking.Contract.CurrentTokenDepositPrice(&_Staking.CallOpts)
}

// CurrentTokenDepositPrice is a free data retrieval call binding the contract method 0x6449cb68.
//
// Solidity: function currentTokenDepositPrice() view returns(uint256)
func (_Staking *StakingCallerSession) CurrentTokenDepositPrice() (*big.Int, error) {
	return _Staking.Contract.CurrentTokenDepositPrice(&_Staking.CallOpts)
}

// CurrentTokenPrice is a free data retrieval call binding the contract method 0x71b3659e.
//
// Solidity: function currentTokenPrice() view returns(uint256)
func (_Staking *StakingCaller) CurrentTokenPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "currentTokenPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentTokenPrice is a free data retrieval call binding the contract method 0x71b3659e.
//
// Solidity: function currentTokenPrice() view returns(uint256)
func (_Staking *StakingSession) CurrentTokenPrice() (*big.Int, error) {
	return _Staking.Contract.CurrentTokenPrice(&_Staking.CallOpts)
}

// CurrentTokenPrice is a free data retrieval call binding the contract method 0x71b3659e.
//
// Solidity: function currentTokenPrice() view returns(uint256)
func (_Staking *StakingCallerSession) CurrentTokenPrice() (*big.Int, error) {
	return _Staking.Contract.CurrentTokenPrice(&_Staking.CallOpts)
}

// FeeEnabled is a free data retrieval call binding the contract method 0xa771ebc7.
//
// Solidity: function feeEnabled() view returns(bool)
func (_Staking *StakingCaller) FeeEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "feeEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// FeeEnabled is a free data retrieval call binding the contract method 0xa771ebc7.
//
// Solidity: function feeEnabled() view returns(bool)
func (_Staking *StakingSession) FeeEnabled() (bool, error) {
	return _Staking.Contract.FeeEnabled(&_Staking.CallOpts)
}

// FeeEnabled is a free data retrieval call binding the contract method 0xa771ebc7.
//
// Solidity: function feeEnabled() view returns(bool)
func (_Staking *StakingCallerSession) FeeEnabled() (bool, error) {
	return _Staking.Contract.FeeEnabled(&_Staking.CallOpts)
}

// FeePercentage is a free data retrieval call binding the contract method 0xa001ecdd.
//
// Solidity: function feePercentage() view returns(uint256)
func (_Staking *StakingCaller) FeePercentage(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "feePercentage")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeePercentage is a free data retrieval call binding the contract method 0xa001ecdd.
//
// Solidity: function feePercentage() view returns(uint256)
func (_Staking *StakingSession) FeePercentage() (*big.Int, error) {
	return _Staking.Contract.FeePercentage(&_Staking.CallOpts)
}

// FeePercentage is a free data retrieval call binding the contract method 0xa001ecdd.
//
// Solidity: function feePercentage() view returns(uint256)
func (_Staking *StakingCallerSession) FeePercentage() (*big.Int, error) {
	return _Staking.Contract.FeePercentage(&_Staking.CallOpts)
}

// GetAmountOutMin is a free data retrieval call binding the contract method 0x3c50eec1.
//
// Solidity: function getAmountOutMin(address _tokenIn, address _tokenOut, uint256 _amountIn) view returns(uint256)
func (_Staking *StakingCaller) GetAmountOutMin(opts *bind.CallOpts, _tokenIn common.Address, _tokenOut common.Address, _amountIn *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "getAmountOutMin", _tokenIn, _tokenOut, _amountIn)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAmountOutMin is a free data retrieval call binding the contract method 0x3c50eec1.
//
// Solidity: function getAmountOutMin(address _tokenIn, address _tokenOut, uint256 _amountIn) view returns(uint256)
func (_Staking *StakingSession) GetAmountOutMin(_tokenIn common.Address, _tokenOut common.Address, _amountIn *big.Int) (*big.Int, error) {
	return _Staking.Contract.GetAmountOutMin(&_Staking.CallOpts, _tokenIn, _tokenOut, _amountIn)
}

// GetAmountOutMin is a free data retrieval call binding the contract method 0x3c50eec1.
//
// Solidity: function getAmountOutMin(address _tokenIn, address _tokenOut, uint256 _amountIn) view returns(uint256)
func (_Staking *StakingCallerSession) GetAmountOutMin(_tokenIn common.Address, _tokenOut common.Address, _amountIn *big.Int) (*big.Int, error) {
	return _Staking.Contract.GetAmountOutMin(&_Staking.CallOpts, _tokenIn, _tokenOut, _amountIn)
}

// InActiveUserTransferHistory is a free data retrieval call binding the contract method 0x72a3815f.
//
// Solidity: function inActiveUserTransferHistory() view returns((address,uint256,uint256,uint256,uint256,uint256,uint256,bool,bool,bool)[])
func (_Staking *StakingCaller) InActiveUserTransferHistory(opts *bind.CallOpts) ([]AccessPassV4userDetails, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "inActiveUserTransferHistory")

	if err != nil {
		return *new([]AccessPassV4userDetails), err
	}

	out0 := *abi.ConvertType(out[0], new([]AccessPassV4userDetails)).(*[]AccessPassV4userDetails)

	return out0, err

}

// InActiveUserTransferHistory is a free data retrieval call binding the contract method 0x72a3815f.
//
// Solidity: function inActiveUserTransferHistory() view returns((address,uint256,uint256,uint256,uint256,uint256,uint256,bool,bool,bool)[])
func (_Staking *StakingSession) InActiveUserTransferHistory() ([]AccessPassV4userDetails, error) {
	return _Staking.Contract.InActiveUserTransferHistory(&_Staking.CallOpts)
}

// InActiveUserTransferHistory is a free data retrieval call binding the contract method 0x72a3815f.
//
// Solidity: function inActiveUserTransferHistory() view returns((address,uint256,uint256,uint256,uint256,uint256,uint256,bool,bool,bool)[])
func (_Staking *StakingCallerSession) InActiveUserTransferHistory() ([]AccessPassV4userDetails, error) {
	return _Staking.Contract.InActiveUserTransferHistory(&_Staking.CallOpts)
}

// IsEligibleToWithdraw is a free data retrieval call binding the contract method 0xf769cddc.
//
// Solidity: function isEligibleToWithdraw(address user) view returns(uint256 current, uint256 expected, bool withdraw_possible)
func (_Staking *StakingCaller) IsEligibleToWithdraw(opts *bind.CallOpts, user common.Address) (struct {
	Current          *big.Int
	Expected         *big.Int
	WithdrawPossible bool
}, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "isEligibleToWithdraw", user)

	outstruct := new(struct {
		Current          *big.Int
		Expected         *big.Int
		WithdrawPossible bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Current = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Expected = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.WithdrawPossible = *abi.ConvertType(out[2], new(bool)).(*bool)

	return *outstruct, err

}

// IsEligibleToWithdraw is a free data retrieval call binding the contract method 0xf769cddc.
//
// Solidity: function isEligibleToWithdraw(address user) view returns(uint256 current, uint256 expected, bool withdraw_possible)
func (_Staking *StakingSession) IsEligibleToWithdraw(user common.Address) (struct {
	Current          *big.Int
	Expected         *big.Int
	WithdrawPossible bool
}, error) {
	return _Staking.Contract.IsEligibleToWithdraw(&_Staking.CallOpts, user)
}

// IsEligibleToWithdraw is a free data retrieval call binding the contract method 0xf769cddc.
//
// Solidity: function isEligibleToWithdraw(address user) view returns(uint256 current, uint256 expected, bool withdraw_possible)
func (_Staking *StakingCallerSession) IsEligibleToWithdraw(user common.Address) (struct {
	Current          *big.Int
	Expected         *big.Int
	WithdrawPossible bool
}, error) {
	return _Staking.Contract.IsEligibleToWithdraw(&_Staking.CallOpts, user)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Staking *StakingCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Staking *StakingSession) Owner() (common.Address, error) {
	return _Staking.Contract.Owner(&_Staking.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Staking *StakingCallerSession) Owner() (common.Address, error) {
	return _Staking.Contract.Owner(&_Staking.CallOpts)
}

// SubtractPercent is a free data retrieval call binding the contract method 0x47309d1c.
//
// Solidity: function subtractPercent(uint256 amount, uint256 percent) pure returns(uint256)
func (_Staking *StakingCaller) SubtractPercent(opts *bind.CallOpts, amount *big.Int, percent *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "subtractPercent", amount, percent)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SubtractPercent is a free data retrieval call binding the contract method 0x47309d1c.
//
// Solidity: function subtractPercent(uint256 amount, uint256 percent) pure returns(uint256)
func (_Staking *StakingSession) SubtractPercent(amount *big.Int, percent *big.Int) (*big.Int, error) {
	return _Staking.Contract.SubtractPercent(&_Staking.CallOpts, amount, percent)
}

// SubtractPercent is a free data retrieval call binding the contract method 0x47309d1c.
//
// Solidity: function subtractPercent(uint256 amount, uint256 percent) pure returns(uint256)
func (_Staking *StakingCallerSession) SubtractPercent(amount *big.Int, percent *big.Int) (*big.Int, error) {
	return _Staking.Contract.SubtractPercent(&_Staking.CallOpts, amount, percent)
}

// TokenWorthIncrease is a free data retrieval call binding the contract method 0x876cc61c.
//
// Solidity: function tokenWorthIncrease() view returns(uint256)
func (_Staking *StakingCaller) TokenWorthIncrease(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "tokenWorthIncrease")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenWorthIncrease is a free data retrieval call binding the contract method 0x876cc61c.
//
// Solidity: function tokenWorthIncrease() view returns(uint256)
func (_Staking *StakingSession) TokenWorthIncrease() (*big.Int, error) {
	return _Staking.Contract.TokenWorthIncrease(&_Staking.CallOpts)
}

// TokenWorthIncrease is a free data retrieval call binding the contract method 0x876cc61c.
//
// Solidity: function tokenWorthIncrease() view returns(uint256)
func (_Staking *StakingCallerSession) TokenWorthIncrease() (*big.Int, error) {
	return _Staking.Contract.TokenWorthIncrease(&_Staking.CallOpts)
}

// UniswapV2Router is a free data retrieval call binding the contract method 0x1694505e.
//
// Solidity: function uniswapV2Router() view returns(address)
func (_Staking *StakingCaller) UniswapV2Router(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "uniswapV2Router")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UniswapV2Router is a free data retrieval call binding the contract method 0x1694505e.
//
// Solidity: function uniswapV2Router() view returns(address)
func (_Staking *StakingSession) UniswapV2Router() (common.Address, error) {
	return _Staking.Contract.UniswapV2Router(&_Staking.CallOpts)
}

// UniswapV2Router is a free data retrieval call binding the contract method 0x1694505e.
//
// Solidity: function uniswapV2Router() view returns(address)
func (_Staking *StakingCallerSession) UniswapV2Router() (common.Address, error) {
	return _Staking.Contract.UniswapV2Router(&_Staking.CallOpts)
}

// UserDepositBalane is a free data retrieval call binding the contract method 0x98c21bc1.
//
// Solidity: function userDepositBalane(address user) view returns(uint256 amount)
func (_Staking *StakingCaller) UserDepositBalane(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "userDepositBalane", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserDepositBalane is a free data retrieval call binding the contract method 0x98c21bc1.
//
// Solidity: function userDepositBalane(address user) view returns(uint256 amount)
func (_Staking *StakingSession) UserDepositBalane(user common.Address) (*big.Int, error) {
	return _Staking.Contract.UserDepositBalane(&_Staking.CallOpts, user)
}

// UserDepositBalane is a free data retrieval call binding the contract method 0x98c21bc1.
//
// Solidity: function userDepositBalane(address user) view returns(uint256 amount)
func (_Staking *StakingCallerSession) UserDepositBalane(user common.Address) (*big.Int, error) {
	return _Staking.Contract.UserDepositBalane(&_Staking.CallOpts, user)
}

// UserHistory is a free data retrieval call binding the contract method 0x100cd71d.
//
// Solidity: function userHistory(address user) view returns((address,uint256,uint256,uint256,uint256,uint256,uint256,bool,bool,bool)[])
func (_Staking *StakingCaller) UserHistory(opts *bind.CallOpts, user common.Address) ([]AccessPassV4userDetails, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "userHistory", user)

	if err != nil {
		return *new([]AccessPassV4userDetails), err
	}

	out0 := *abi.ConvertType(out[0], new([]AccessPassV4userDetails)).(*[]AccessPassV4userDetails)

	return out0, err

}

// UserHistory is a free data retrieval call binding the contract method 0x100cd71d.
//
// Solidity: function userHistory(address user) view returns((address,uint256,uint256,uint256,uint256,uint256,uint256,bool,bool,bool)[])
func (_Staking *StakingSession) UserHistory(user common.Address) ([]AccessPassV4userDetails, error) {
	return _Staking.Contract.UserHistory(&_Staking.CallOpts, user)
}

// UserHistory is a free data retrieval call binding the contract method 0x100cd71d.
//
// Solidity: function userHistory(address user) view returns((address,uint256,uint256,uint256,uint256,uint256,uint256,bool,bool,bool)[])
func (_Staking *StakingCallerSession) UserHistory(user common.Address) ([]AccessPassV4userDetails, error) {
	return _Staking.Contract.UserHistory(&_Staking.CallOpts, user)
}

// UserProfile is a free data retrieval call binding the contract method 0x1c8e8962.
//
// Solidity: function userProfile(address ) view returns(address userAddress, uint256 TokensDeposited, uint256 accessPassedTokenPrice, uint256 latestTrxPrice, uint256 totalPriceBUSD, uint256 trxTime, uint256 worthIncrease, bool hasDeposited, bool status, bool autoWithdraw)
func (_Staking *StakingCaller) UserProfile(opts *bind.CallOpts, arg0 common.Address) (struct {
	UserAddress            common.Address
	TokensDeposited        *big.Int
	AccessPassedTokenPrice *big.Int
	LatestTrxPrice         *big.Int
	TotalPriceBUSD         *big.Int
	TrxTime                *big.Int
	WorthIncrease          *big.Int
	HasDeposited           bool
	Status                 bool
	AutoWithdraw           bool
}, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "userProfile", arg0)

	outstruct := new(struct {
		UserAddress            common.Address
		TokensDeposited        *big.Int
		AccessPassedTokenPrice *big.Int
		LatestTrxPrice         *big.Int
		TotalPriceBUSD         *big.Int
		TrxTime                *big.Int
		WorthIncrease          *big.Int
		HasDeposited           bool
		Status                 bool
		AutoWithdraw           bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.UserAddress = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.TokensDeposited = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.AccessPassedTokenPrice = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.LatestTrxPrice = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.TotalPriceBUSD = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.TrxTime = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.WorthIncrease = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.HasDeposited = *abi.ConvertType(out[7], new(bool)).(*bool)
	outstruct.Status = *abi.ConvertType(out[8], new(bool)).(*bool)
	outstruct.AutoWithdraw = *abi.ConvertType(out[9], new(bool)).(*bool)

	return *outstruct, err

}

// UserProfile is a free data retrieval call binding the contract method 0x1c8e8962.
//
// Solidity: function userProfile(address ) view returns(address userAddress, uint256 TokensDeposited, uint256 accessPassedTokenPrice, uint256 latestTrxPrice, uint256 totalPriceBUSD, uint256 trxTime, uint256 worthIncrease, bool hasDeposited, bool status, bool autoWithdraw)
func (_Staking *StakingSession) UserProfile(arg0 common.Address) (struct {
	UserAddress            common.Address
	TokensDeposited        *big.Int
	AccessPassedTokenPrice *big.Int
	LatestTrxPrice         *big.Int
	TotalPriceBUSD         *big.Int
	TrxTime                *big.Int
	WorthIncrease          *big.Int
	HasDeposited           bool
	Status                 bool
	AutoWithdraw           bool
}, error) {
	return _Staking.Contract.UserProfile(&_Staking.CallOpts, arg0)
}

// UserProfile is a free data retrieval call binding the contract method 0x1c8e8962.
//
// Solidity: function userProfile(address ) view returns(address userAddress, uint256 TokensDeposited, uint256 accessPassedTokenPrice, uint256 latestTrxPrice, uint256 totalPriceBUSD, uint256 trxTime, uint256 worthIncrease, bool hasDeposited, bool status, bool autoWithdraw)
func (_Staking *StakingCallerSession) UserProfile(arg0 common.Address) (struct {
	UserAddress            common.Address
	TokensDeposited        *big.Int
	AccessPassedTokenPrice *big.Int
	LatestTrxPrice         *big.Int
	TotalPriceBUSD         *big.Int
	TrxTime                *big.Int
	WorthIncrease          *big.Int
	HasDeposited           bool
	Status                 bool
	AutoWithdraw           bool
}, error) {
	return _Staking.Contract.UserProfile(&_Staking.CallOpts, arg0)
}

// UserStatus is a free data retrieval call binding the contract method 0x225d29a1.
//
// Solidity: function userStatus(address user) view returns(bool active)
func (_Staking *StakingCaller) UserStatus(opts *bind.CallOpts, user common.Address) (bool, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "userStatus", user)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// UserStatus is a free data retrieval call binding the contract method 0x225d29a1.
//
// Solidity: function userStatus(address user) view returns(bool active)
func (_Staking *StakingSession) UserStatus(user common.Address) (bool, error) {
	return _Staking.Contract.UserStatus(&_Staking.CallOpts, user)
}

// UserStatus is a free data retrieval call binding the contract method 0x225d29a1.
//
// Solidity: function userStatus(address user) view returns(bool active)
func (_Staking *StakingCallerSession) UserStatus(user common.Address) (bool, error) {
	return _Staking.Contract.UserStatus(&_Staking.CallOpts, user)
}

// UsersHistory is a free data retrieval call binding the contract method 0x8ab2e9e4.
//
// Solidity: function usersHistory() view returns((address,uint256,uint256,uint256,uint256,uint256,uint256,bool,bool,bool)[])
func (_Staking *StakingCaller) UsersHistory(opts *bind.CallOpts) ([]AccessPassV4userDetails, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "usersHistory")

	if err != nil {
		return *new([]AccessPassV4userDetails), err
	}

	out0 := *abi.ConvertType(out[0], new([]AccessPassV4userDetails)).(*[]AccessPassV4userDetails)

	return out0, err

}

// UsersHistory is a free data retrieval call binding the contract method 0x8ab2e9e4.
//
// Solidity: function usersHistory() view returns((address,uint256,uint256,uint256,uint256,uint256,uint256,bool,bool,bool)[])
func (_Staking *StakingSession) UsersHistory() ([]AccessPassV4userDetails, error) {
	return _Staking.Contract.UsersHistory(&_Staking.CallOpts)
}

// UsersHistory is a free data retrieval call binding the contract method 0x8ab2e9e4.
//
// Solidity: function usersHistory() view returns((address,uint256,uint256,uint256,uint256,uint256,uint256,bool,bool,bool)[])
func (_Staking *StakingCallerSession) UsersHistory() ([]AccessPassV4userDetails, error) {
	return _Staking.Contract.UsersHistory(&_Staking.CallOpts)
}

// AccessPassTokens is a paid mutator transaction binding the contract method 0x0621437e.
//
// Solidity: function accessPassTokens(uint256 amountOutMin, address[] path) payable returns()
func (_Staking *StakingTransactor) AccessPassTokens(opts *bind.TransactOpts, amountOutMin *big.Int, path []common.Address) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "accessPassTokens", amountOutMin, path)
}

// AccessPassTokens is a paid mutator transaction binding the contract method 0x0621437e.
//
// Solidity: function accessPassTokens(uint256 amountOutMin, address[] path) payable returns()
func (_Staking *StakingSession) AccessPassTokens(amountOutMin *big.Int, path []common.Address) (*types.Transaction, error) {
	return _Staking.Contract.AccessPassTokens(&_Staking.TransactOpts, amountOutMin, path)
}

// AccessPassTokens is a paid mutator transaction binding the contract method 0x0621437e.
//
// Solidity: function accessPassTokens(uint256 amountOutMin, address[] path) payable returns()
func (_Staking *StakingTransactorSession) AccessPassTokens(amountOutMin *big.Int, path []common.Address) (*types.Transaction, error) {
	return _Staking.Contract.AccessPassTokens(&_Staking.TransactOpts, amountOutMin, path)
}

// AutoWithdraw is a paid mutator transaction binding the contract method 0xd336121c.
//
// Solidity: function autoWithdraw(address[] addresses) returns()
func (_Staking *StakingTransactor) AutoWithdraw(opts *bind.TransactOpts, addresses []common.Address) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "autoWithdraw", addresses)
}

// AutoWithdraw is a paid mutator transaction binding the contract method 0xd336121c.
//
// Solidity: function autoWithdraw(address[] addresses) returns()
func (_Staking *StakingSession) AutoWithdraw(addresses []common.Address) (*types.Transaction, error) {
	return _Staking.Contract.AutoWithdraw(&_Staking.TransactOpts, addresses)
}

// AutoWithdraw is a paid mutator transaction binding the contract method 0xd336121c.
//
// Solidity: function autoWithdraw(address[] addresses) returns()
func (_Staking *StakingTransactorSession) AutoWithdraw(addresses []common.Address) (*types.Transaction, error) {
	return _Staking.Contract.AutoWithdraw(&_Staking.TransactOpts, addresses)
}

// DepositTokens is a paid mutator transaction binding the contract method 0xdd49756e.
//
// Solidity: function depositTokens(uint256 _amount) returns()
func (_Staking *StakingTransactor) DepositTokens(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "depositTokens", _amount)
}

// DepositTokens is a paid mutator transaction binding the contract method 0xdd49756e.
//
// Solidity: function depositTokens(uint256 _amount) returns()
func (_Staking *StakingSession) DepositTokens(_amount *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.DepositTokens(&_Staking.TransactOpts, _amount)
}

// DepositTokens is a paid mutator transaction binding the contract method 0xdd49756e.
//
// Solidity: function depositTokens(uint256 _amount) returns()
func (_Staking *StakingTransactorSession) DepositTokens(_amount *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.DepositTokens(&_Staking.TransactOpts, _amount)
}

// RecoverStuckBnb is a paid mutator transaction binding the contract method 0x5d896f49.
//
// Solidity: function recoverStuckBnb(uint256 amountPercentage) returns()
func (_Staking *StakingTransactor) RecoverStuckBnb(opts *bind.TransactOpts, amountPercentage *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "recoverStuckBnb", amountPercentage)
}

// RecoverStuckBnb is a paid mutator transaction binding the contract method 0x5d896f49.
//
// Solidity: function recoverStuckBnb(uint256 amountPercentage) returns()
func (_Staking *StakingSession) RecoverStuckBnb(amountPercentage *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.RecoverStuckBnb(&_Staking.TransactOpts, amountPercentage)
}

// RecoverStuckBnb is a paid mutator transaction binding the contract method 0x5d896f49.
//
// Solidity: function recoverStuckBnb(uint256 amountPercentage) returns()
func (_Staking *StakingTransactorSession) RecoverStuckBnb(amountPercentage *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.RecoverStuckBnb(&_Staking.TransactOpts, amountPercentage)
}

// RecoverStuckToken is a paid mutator transaction binding the contract method 0xfd3e76c4.
//
// Solidity: function recoverStuckToken(address tokenAddress, uint256 tokens) returns(bool success)
func (_Staking *StakingTransactor) RecoverStuckToken(opts *bind.TransactOpts, tokenAddress common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "recoverStuckToken", tokenAddress, tokens)
}

// RecoverStuckToken is a paid mutator transaction binding the contract method 0xfd3e76c4.
//
// Solidity: function recoverStuckToken(address tokenAddress, uint256 tokens) returns(bool success)
func (_Staking *StakingSession) RecoverStuckToken(tokenAddress common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.RecoverStuckToken(&_Staking.TransactOpts, tokenAddress, tokens)
}

// RecoverStuckToken is a paid mutator transaction binding the contract method 0xfd3e76c4.
//
// Solidity: function recoverStuckToken(address tokenAddress, uint256 tokens) returns(bool success)
func (_Staking *StakingTransactorSession) RecoverStuckToken(tokenAddress common.Address, tokens *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.RecoverStuckToken(&_Staking.TransactOpts, tokenAddress, tokens)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Staking *StakingTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Staking *StakingSession) RenounceOwnership() (*types.Transaction, error) {
	return _Staking.Contract.RenounceOwnership(&_Staking.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Staking *StakingTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Staking.Contract.RenounceOwnership(&_Staking.TransactOpts)
}

// SetDepositLimitInBusd is a paid mutator transaction binding the contract method 0x72b0a693.
//
// Solidity: function setDepositLimitInBusd(uint256 busdAmount) returns()
func (_Staking *StakingTransactor) SetDepositLimitInBusd(opts *bind.TransactOpts, busdAmount *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "setDepositLimitInBusd", busdAmount)
}

// SetDepositLimitInBusd is a paid mutator transaction binding the contract method 0x72b0a693.
//
// Solidity: function setDepositLimitInBusd(uint256 busdAmount) returns()
func (_Staking *StakingSession) SetDepositLimitInBusd(busdAmount *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.SetDepositLimitInBusd(&_Staking.TransactOpts, busdAmount)
}

// SetDepositLimitInBusd is a paid mutator transaction binding the contract method 0x72b0a693.
//
// Solidity: function setDepositLimitInBusd(uint256 busdAmount) returns()
func (_Staking *StakingTransactorSession) SetDepositLimitInBusd(busdAmount *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.SetDepositLimitInBusd(&_Staking.TransactOpts, busdAmount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Staking *StakingTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Staking *StakingSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Staking.Contract.TransferOwnership(&_Staking.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Staking *StakingTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Staking.Contract.TransferOwnership(&_Staking.TransactOpts, newOwner)
}

// TransferUsersBalanceToOwner is a paid mutator transaction binding the contract method 0x46cdcf44.
//
// Solidity: function transferUsersBalanceToOwner(address[] inactiveUsers) returns()
func (_Staking *StakingTransactor) TransferUsersBalanceToOwner(opts *bind.TransactOpts, inactiveUsers []common.Address) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "transferUsersBalanceToOwner", inactiveUsers)
}

// TransferUsersBalanceToOwner is a paid mutator transaction binding the contract method 0x46cdcf44.
//
// Solidity: function transferUsersBalanceToOwner(address[] inactiveUsers) returns()
func (_Staking *StakingSession) TransferUsersBalanceToOwner(inactiveUsers []common.Address) (*types.Transaction, error) {
	return _Staking.Contract.TransferUsersBalanceToOwner(&_Staking.TransactOpts, inactiveUsers)
}

// TransferUsersBalanceToOwner is a paid mutator transaction binding the contract method 0x46cdcf44.
//
// Solidity: function transferUsersBalanceToOwner(address[] inactiveUsers) returns()
func (_Staking *StakingTransactorSession) TransferUsersBalanceToOwner(inactiveUsers []common.Address) (*types.Transaction, error) {
	return _Staking.Contract.TransferUsersBalanceToOwner(&_Staking.TransactOpts, inactiveUsers)
}

// UpdateAutoWithdrawStatus is a paid mutator transaction binding the contract method 0x826ae09b.
//
// Solidity: function updateAutoWithdrawStatus(address[] users, bool autoWithdrawStatus) returns()
func (_Staking *StakingTransactor) UpdateAutoWithdrawStatus(opts *bind.TransactOpts, users []common.Address, autoWithdrawStatus bool) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "updateAutoWithdrawStatus", users, autoWithdrawStatus)
}

// UpdateAutoWithdrawStatus is a paid mutator transaction binding the contract method 0x826ae09b.
//
// Solidity: function updateAutoWithdrawStatus(address[] users, bool autoWithdrawStatus) returns()
func (_Staking *StakingSession) UpdateAutoWithdrawStatus(users []common.Address, autoWithdrawStatus bool) (*types.Transaction, error) {
	return _Staking.Contract.UpdateAutoWithdrawStatus(&_Staking.TransactOpts, users, autoWithdrawStatus)
}

// UpdateAutoWithdrawStatus is a paid mutator transaction binding the contract method 0x826ae09b.
//
// Solidity: function updateAutoWithdrawStatus(address[] users, bool autoWithdrawStatus) returns()
func (_Staking *StakingTransactorSession) UpdateAutoWithdrawStatus(users []common.Address, autoWithdrawStatus bool) (*types.Transaction, error) {
	return _Staking.Contract.UpdateAutoWithdrawStatus(&_Staking.TransactOpts, users, autoWithdrawStatus)
}

// UpdateCurrentToken is a paid mutator transaction binding the contract method 0xfae9c2a4.
//
// Solidity: function updateCurrentToken(address newToken) returns()
func (_Staking *StakingTransactor) UpdateCurrentToken(opts *bind.TransactOpts, newToken common.Address) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "updateCurrentToken", newToken)
}

// UpdateCurrentToken is a paid mutator transaction binding the contract method 0xfae9c2a4.
//
// Solidity: function updateCurrentToken(address newToken) returns()
func (_Staking *StakingSession) UpdateCurrentToken(newToken common.Address) (*types.Transaction, error) {
	return _Staking.Contract.UpdateCurrentToken(&_Staking.TransactOpts, newToken)
}

// UpdateCurrentToken is a paid mutator transaction binding the contract method 0xfae9c2a4.
//
// Solidity: function updateCurrentToken(address newToken) returns()
func (_Staking *StakingTransactorSession) UpdateCurrentToken(newToken common.Address) (*types.Transaction, error) {
	return _Staking.Contract.UpdateCurrentToken(&_Staking.TransactOpts, newToken)
}

// UpdateOwnAutoWithdrawStatus is a paid mutator transaction binding the contract method 0x058a84cd.
//
// Solidity: function updateOwnAutoWithdrawStatus(bool autoWithdrawStatus) returns()
func (_Staking *StakingTransactor) UpdateOwnAutoWithdrawStatus(opts *bind.TransactOpts, autoWithdrawStatus bool) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "updateOwnAutoWithdrawStatus", autoWithdrawStatus)
}

// UpdateOwnAutoWithdrawStatus is a paid mutator transaction binding the contract method 0x058a84cd.
//
// Solidity: function updateOwnAutoWithdrawStatus(bool autoWithdrawStatus) returns()
func (_Staking *StakingSession) UpdateOwnAutoWithdrawStatus(autoWithdrawStatus bool) (*types.Transaction, error) {
	return _Staking.Contract.UpdateOwnAutoWithdrawStatus(&_Staking.TransactOpts, autoWithdrawStatus)
}

// UpdateOwnAutoWithdrawStatus is a paid mutator transaction binding the contract method 0x058a84cd.
//
// Solidity: function updateOwnAutoWithdrawStatus(bool autoWithdrawStatus) returns()
func (_Staking *StakingTransactorSession) UpdateOwnAutoWithdrawStatus(autoWithdrawStatus bool) (*types.Transaction, error) {
	return _Staking.Contract.UpdateOwnAutoWithdrawStatus(&_Staking.TransactOpts, autoWithdrawStatus)
}

// UpdateTokenWorthIncrease is a paid mutator transaction binding the contract method 0x094e8268.
//
// Solidity: function updateTokenWorthIncrease(uint256 percentage) returns()
func (_Staking *StakingTransactor) UpdateTokenWorthIncrease(opts *bind.TransactOpts, percentage *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "updateTokenWorthIncrease", percentage)
}

// UpdateTokenWorthIncrease is a paid mutator transaction binding the contract method 0x094e8268.
//
// Solidity: function updateTokenWorthIncrease(uint256 percentage) returns()
func (_Staking *StakingSession) UpdateTokenWorthIncrease(percentage *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.UpdateTokenWorthIncrease(&_Staking.TransactOpts, percentage)
}

// UpdateTokenWorthIncrease is a paid mutator transaction binding the contract method 0x094e8268.
//
// Solidity: function updateTokenWorthIncrease(uint256 percentage) returns()
func (_Staking *StakingTransactorSession) UpdateTokenWorthIncrease(percentage *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.UpdateTokenWorthIncrease(&_Staking.TransactOpts, percentage)
}

// UpdateUserDetails is a paid mutator transaction binding the contract method 0x59fab1de.
//
// Solidity: function updateUserDetails(address user, uint256 TokensDeposited, uint256 accessPassedTokenPrice, uint256 latestTrxPrice, uint256 totalPrice, uint256 trxTime, uint256 worthIncrease, bool hasDeposited, bool status, bool autoWithdrawStatus) returns()
func (_Staking *StakingTransactor) UpdateUserDetails(opts *bind.TransactOpts, user common.Address, TokensDeposited *big.Int, accessPassedTokenPrice *big.Int, latestTrxPrice *big.Int, totalPrice *big.Int, trxTime *big.Int, worthIncrease *big.Int, hasDeposited bool, status bool, autoWithdrawStatus bool) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "updateUserDetails", user, TokensDeposited, accessPassedTokenPrice, latestTrxPrice, totalPrice, trxTime, worthIncrease, hasDeposited, status, autoWithdrawStatus)
}

// UpdateUserDetails is a paid mutator transaction binding the contract method 0x59fab1de.
//
// Solidity: function updateUserDetails(address user, uint256 TokensDeposited, uint256 accessPassedTokenPrice, uint256 latestTrxPrice, uint256 totalPrice, uint256 trxTime, uint256 worthIncrease, bool hasDeposited, bool status, bool autoWithdrawStatus) returns()
func (_Staking *StakingSession) UpdateUserDetails(user common.Address, TokensDeposited *big.Int, accessPassedTokenPrice *big.Int, latestTrxPrice *big.Int, totalPrice *big.Int, trxTime *big.Int, worthIncrease *big.Int, hasDeposited bool, status bool, autoWithdrawStatus bool) (*types.Transaction, error) {
	return _Staking.Contract.UpdateUserDetails(&_Staking.TransactOpts, user, TokensDeposited, accessPassedTokenPrice, latestTrxPrice, totalPrice, trxTime, worthIncrease, hasDeposited, status, autoWithdrawStatus)
}

// UpdateUserDetails is a paid mutator transaction binding the contract method 0x59fab1de.
//
// Solidity: function updateUserDetails(address user, uint256 TokensDeposited, uint256 accessPassedTokenPrice, uint256 latestTrxPrice, uint256 totalPrice, uint256 trxTime, uint256 worthIncrease, bool hasDeposited, bool status, bool autoWithdrawStatus) returns()
func (_Staking *StakingTransactorSession) UpdateUserDetails(user common.Address, TokensDeposited *big.Int, accessPassedTokenPrice *big.Int, latestTrxPrice *big.Int, totalPrice *big.Int, trxTime *big.Int, worthIncrease *big.Int, hasDeposited bool, status bool, autoWithdrawStatus bool) (*types.Transaction, error) {
	return _Staking.Contract.UpdateUserDetails(&_Staking.TransactOpts, user, TokensDeposited, accessPassedTokenPrice, latestTrxPrice, totalPrice, trxTime, worthIncrease, hasDeposited, status, autoWithdrawStatus)
}

// UpdateUserStatus is a paid mutator transaction binding the contract method 0x51452a62.
//
// Solidity: function updateUserStatus(address[] users, bool active) returns()
func (_Staking *StakingTransactor) UpdateUserStatus(opts *bind.TransactOpts, users []common.Address, active bool) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "updateUserStatus", users, active)
}

// UpdateUserStatus is a paid mutator transaction binding the contract method 0x51452a62.
//
// Solidity: function updateUserStatus(address[] users, bool active) returns()
func (_Staking *StakingSession) UpdateUserStatus(users []common.Address, active bool) (*types.Transaction, error) {
	return _Staking.Contract.UpdateUserStatus(&_Staking.TransactOpts, users, active)
}

// UpdateUserStatus is a paid mutator transaction binding the contract method 0x51452a62.
//
// Solidity: function updateUserStatus(address[] users, bool active) returns()
func (_Staking *StakingTransactorSession) UpdateUserStatus(users []common.Address, active bool) (*types.Transaction, error) {
	return _Staking.Contract.UpdateUserStatus(&_Staking.TransactOpts, users, active)
}

// UpdatefeeEnabled is a paid mutator transaction binding the contract method 0x703893f1.
//
// Solidity: function updatefeeEnabled(bool _isEnabled) returns()
func (_Staking *StakingTransactor) UpdatefeeEnabled(opts *bind.TransactOpts, _isEnabled bool) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "updatefeeEnabled", _isEnabled)
}

// UpdatefeeEnabled is a paid mutator transaction binding the contract method 0x703893f1.
//
// Solidity: function updatefeeEnabled(bool _isEnabled) returns()
func (_Staking *StakingSession) UpdatefeeEnabled(_isEnabled bool) (*types.Transaction, error) {
	return _Staking.Contract.UpdatefeeEnabled(&_Staking.TransactOpts, _isEnabled)
}

// UpdatefeeEnabled is a paid mutator transaction binding the contract method 0x703893f1.
//
// Solidity: function updatefeeEnabled(bool _isEnabled) returns()
func (_Staking *StakingTransactorSession) UpdatefeeEnabled(_isEnabled bool) (*types.Transaction, error) {
	return _Staking.Contract.UpdatefeeEnabled(&_Staking.TransactOpts, _isEnabled)
}

// UpdatefeePercentage is a paid mutator transaction binding the contract method 0xc46447a8.
//
// Solidity: function updatefeePercentage(uint256 _feepercent) returns()
func (_Staking *StakingTransactor) UpdatefeePercentage(opts *bind.TransactOpts, _feepercent *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "updatefeePercentage", _feepercent)
}

// UpdatefeePercentage is a paid mutator transaction binding the contract method 0xc46447a8.
//
// Solidity: function updatefeePercentage(uint256 _feepercent) returns()
func (_Staking *StakingSession) UpdatefeePercentage(_feepercent *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.UpdatefeePercentage(&_Staking.TransactOpts, _feepercent)
}

// UpdatefeePercentage is a paid mutator transaction binding the contract method 0xc46447a8.
//
// Solidity: function updatefeePercentage(uint256 _feepercent) returns()
func (_Staking *StakingTransactorSession) UpdatefeePercentage(_feepercent *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.UpdatefeePercentage(&_Staking.TransactOpts, _feepercent)
}

// UserWithdrawProfit is a paid mutator transaction binding the contract method 0x02b73811.
//
// Solidity: function userWithdrawProfit() returns()
func (_Staking *StakingTransactor) UserWithdrawProfit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "userWithdrawProfit")
}

// UserWithdrawProfit is a paid mutator transaction binding the contract method 0x02b73811.
//
// Solidity: function userWithdrawProfit() returns()
func (_Staking *StakingSession) UserWithdrawProfit() (*types.Transaction, error) {
	return _Staking.Contract.UserWithdrawProfit(&_Staking.TransactOpts)
}

// UserWithdrawProfit is a paid mutator transaction binding the contract method 0x02b73811.
//
// Solidity: function userWithdrawProfit() returns()
func (_Staking *StakingTransactorSession) UserWithdrawProfit() (*types.Transaction, error) {
	return _Staking.Contract.UserWithdrawProfit(&_Staking.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns()
func (_Staking *StakingTransactor) Withdraw(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "withdraw", _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns()
func (_Staking *StakingSession) Withdraw(_amount *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.Withdraw(&_Staking.TransactOpts, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _amount) returns()
func (_Staking *StakingTransactorSession) Withdraw(_amount *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.Withdraw(&_Staking.TransactOpts, _amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Staking *StakingTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Staking *StakingSession) Receive() (*types.Transaction, error) {
	return _Staking.Contract.Receive(&_Staking.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Staking *StakingTransactorSession) Receive() (*types.Transaction, error) {
	return _Staking.Contract.Receive(&_Staking.TransactOpts)
}

// StakingBalanceClearIterator is returned from FilterBalanceClear and is used to iterate over the raw logs and unpacked data for BalanceClear events raised by the Staking contract.
type StakingBalanceClearIterator struct {
	Event *StakingBalanceClear // Event containing the contract specifics and raw log

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
func (it *StakingBalanceClearIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingBalanceClear)
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
		it.Event = new(StakingBalanceClear)
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
func (it *StakingBalanceClearIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingBalanceClearIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingBalanceClear represents a BalanceClear event raised by the Staking contract.
type StakingBalanceClear struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBalanceClear is a free log retrieval operation binding the contract event 0x377153983e64f72f80af47182763316780b9133be808b7082a16710202b8fbda.
//
// Solidity: event BalanceClear(uint256 amount)
func (_Staking *StakingFilterer) FilterBalanceClear(opts *bind.FilterOpts) (*StakingBalanceClearIterator, error) {

	logs, sub, err := _Staking.contract.FilterLogs(opts, "BalanceClear")
	if err != nil {
		return nil, err
	}
	return &StakingBalanceClearIterator{contract: _Staking.contract, event: "BalanceClear", logs: logs, sub: sub}, nil
}

// WatchBalanceClear is a free log subscription operation binding the contract event 0x377153983e64f72f80af47182763316780b9133be808b7082a16710202b8fbda.
//
// Solidity: event BalanceClear(uint256 amount)
func (_Staking *StakingFilterer) WatchBalanceClear(opts *bind.WatchOpts, sink chan<- *StakingBalanceClear) (event.Subscription, error) {

	logs, sub, err := _Staking.contract.WatchLogs(opts, "BalanceClear")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingBalanceClear)
				if err := _Staking.contract.UnpackLog(event, "BalanceClear", log); err != nil {
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

// ParseBalanceClear is a log parse operation binding the contract event 0x377153983e64f72f80af47182763316780b9133be808b7082a16710202b8fbda.
//
// Solidity: event BalanceClear(uint256 amount)
func (_Staking *StakingFilterer) ParseBalanceClear(log types.Log) (*StakingBalanceClear, error) {
	event := new(StakingBalanceClear)
	if err := _Staking.contract.UnpackLog(event, "BalanceClear", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Staking contract.
type StakingDepositIterator struct {
	Event *StakingDeposit // Event containing the contract specifics and raw log

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
func (it *StakingDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingDeposit)
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
		it.Event = new(StakingDeposit)
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
func (it *StakingDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingDeposit represents a Deposit event raised by the Staking contract.
type StakingDeposit struct {
	User   common.Address
	Amount *big.Int
	Token  common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xe31c7b8d08ee7db0afa68782e1028ef92305caeea8626633ad44d413e30f6b2f.
//
// Solidity: event Deposit(address indexed user, uint256 amount, address token)
func (_Staking *StakingFilterer) FilterDeposit(opts *bind.FilterOpts, user []common.Address) (*StakingDepositIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "Deposit", userRule)
	if err != nil {
		return nil, err
	}
	return &StakingDepositIterator{contract: _Staking.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xe31c7b8d08ee7db0afa68782e1028ef92305caeea8626633ad44d413e30f6b2f.
//
// Solidity: event Deposit(address indexed user, uint256 amount, address token)
func (_Staking *StakingFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *StakingDeposit, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "Deposit", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingDeposit)
				if err := _Staking.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xe31c7b8d08ee7db0afa68782e1028ef92305caeea8626633ad44d413e30f6b2f.
//
// Solidity: event Deposit(address indexed user, uint256 amount, address token)
func (_Staking *StakingFilterer) ParseDeposit(log types.Log) (*StakingDeposit, error) {
	event := new(StakingDeposit)
	if err := _Staking.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Staking contract.
type StakingOwnershipTransferredIterator struct {
	Event *StakingOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *StakingOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingOwnershipTransferred)
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
		it.Event = new(StakingOwnershipTransferred)
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
func (it *StakingOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingOwnershipTransferred represents a OwnershipTransferred event raised by the Staking contract.
type StakingOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Staking *StakingFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*StakingOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &StakingOwnershipTransferredIterator{contract: _Staking.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Staking *StakingFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *StakingOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingOwnershipTransferred)
				if err := _Staking.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Staking *StakingFilterer) ParseOwnershipTransferred(log types.Log) (*StakingOwnershipTransferred, error) {
	event := new(StakingOwnershipTransferred)
	if err := _Staking.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Staking contract.
type StakingTransferIterator struct {
	Event *StakingTransfer // Event containing the contract specifics and raw log

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
func (it *StakingTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingTransfer)
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
		it.Event = new(StakingTransfer)
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
func (it *StakingTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingTransfer represents a Transfer event raised by the Staking contract.
type StakingTransfer struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_Staking *StakingFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*StakingTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &StakingTransferIterator{contract: _Staking.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_Staking *StakingFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *StakingTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingTransfer)
				if err := _Staking.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_Staking *StakingFilterer) ParseTransfer(log types.Log) (*StakingTransfer, error) {
	event := new(StakingTransfer)
	if err := _Staking.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Staking contract.
type StakingWithdrawIterator struct {
	Event *StakingWithdraw // Event containing the contract specifics and raw log

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
func (it *StakingWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingWithdraw)
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
		it.Event = new(StakingWithdraw)
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
func (it *StakingWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingWithdraw represents a Withdraw event raised by the Staking contract.
type StakingWithdraw struct {
	User   common.Address
	Amount *big.Int
	Token  common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x56c54ba9bd38d8fd62012e42c7ee564519b09763c426d331b3661b537ead19b2.
//
// Solidity: event Withdraw(address indexed user, uint256 amount, address token)
func (_Staking *StakingFilterer) FilterWithdraw(opts *bind.FilterOpts, user []common.Address) (*StakingWithdrawIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "Withdraw", userRule)
	if err != nil {
		return nil, err
	}
	return &StakingWithdrawIterator{contract: _Staking.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x56c54ba9bd38d8fd62012e42c7ee564519b09763c426d331b3661b537ead19b2.
//
// Solidity: event Withdraw(address indexed user, uint256 amount, address token)
func (_Staking *StakingFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *StakingWithdraw, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "Withdraw", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingWithdraw)
				if err := _Staking.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x56c54ba9bd38d8fd62012e42c7ee564519b09763c426d331b3661b537ead19b2.
//
// Solidity: event Withdraw(address indexed user, uint256 amount, address token)
func (_Staking *StakingFilterer) ParseWithdraw(log types.Log) (*StakingWithdraw, error) {
	event := new(StakingWithdraw)
	if err := _Staking.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingBnbSwappedIterator is returned from FilterBnbSwapped and is used to iterate over the raw logs and unpacked data for BnbSwapped events raised by the Staking contract.
type StakingBnbSwappedIterator struct {
	Event *StakingBnbSwapped // Event containing the contract specifics and raw log

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
func (it *StakingBnbSwappedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingBnbSwapped)
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
		it.Event = new(StakingBnbSwapped)
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
func (it *StakingBnbSwappedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingBnbSwappedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingBnbSwapped represents a BnbSwapped event raised by the Staking contract.
type StakingBnbSwapped struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBnbSwapped is a free log retrieval operation binding the contract event 0x8cd9932bee04826679d294b4759416111d2f1dbd4602012633ebf7290c266bcc.
//
// Solidity: event bnbSwapped(address user, uint256 amount)
func (_Staking *StakingFilterer) FilterBnbSwapped(opts *bind.FilterOpts) (*StakingBnbSwappedIterator, error) {

	logs, sub, err := _Staking.contract.FilterLogs(opts, "bnbSwapped")
	if err != nil {
		return nil, err
	}
	return &StakingBnbSwappedIterator{contract: _Staking.contract, event: "bnbSwapped", logs: logs, sub: sub}, nil
}

// WatchBnbSwapped is a free log subscription operation binding the contract event 0x8cd9932bee04826679d294b4759416111d2f1dbd4602012633ebf7290c266bcc.
//
// Solidity: event bnbSwapped(address user, uint256 amount)
func (_Staking *StakingFilterer) WatchBnbSwapped(opts *bind.WatchOpts, sink chan<- *StakingBnbSwapped) (event.Subscription, error) {

	logs, sub, err := _Staking.contract.WatchLogs(opts, "bnbSwapped")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingBnbSwapped)
				if err := _Staking.contract.UnpackLog(event, "bnbSwapped", log); err != nil {
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

// ParseBnbSwapped is a log parse operation binding the contract event 0x8cd9932bee04826679d294b4759416111d2f1dbd4602012633ebf7290c266bcc.
//
// Solidity: event bnbSwapped(address user, uint256 amount)
func (_Staking *StakingFilterer) ParseBnbSwapped(log types.Log) (*StakingBnbSwapped, error) {
	event := new(StakingBnbSwapped)
	if err := _Staking.contract.UnpackLog(event, "bnbSwapped", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingClearTokenIterator is returned from FilterClearToken and is used to iterate over the raw logs and unpacked data for ClearToken events raised by the Staking contract.
type StakingClearTokenIterator struct {
	Event *StakingClearToken // Event containing the contract specifics and raw log

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
func (it *StakingClearTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingClearToken)
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
		it.Event = new(StakingClearToken)
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
func (it *StakingClearTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingClearTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingClearToken represents a ClearToken event raised by the Staking contract.
type StakingClearToken struct {
	TokenAddressCleared common.Address
	Amount              *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterClearToken is a free log retrieval operation binding the contract event 0x960427a4fde284dcc8da2c51a43d210cfd7ed6f4328a24ba33eee49910db564b.
//
// Solidity: event clearToken(address TokenAddressCleared, uint256 Amount)
func (_Staking *StakingFilterer) FilterClearToken(opts *bind.FilterOpts) (*StakingClearTokenIterator, error) {

	logs, sub, err := _Staking.contract.FilterLogs(opts, "clearToken")
	if err != nil {
		return nil, err
	}
	return &StakingClearTokenIterator{contract: _Staking.contract, event: "clearToken", logs: logs, sub: sub}, nil
}

// WatchClearToken is a free log subscription operation binding the contract event 0x960427a4fde284dcc8da2c51a43d210cfd7ed6f4328a24ba33eee49910db564b.
//
// Solidity: event clearToken(address TokenAddressCleared, uint256 Amount)
func (_Staking *StakingFilterer) WatchClearToken(opts *bind.WatchOpts, sink chan<- *StakingClearToken) (event.Subscription, error) {

	logs, sub, err := _Staking.contract.WatchLogs(opts, "clearToken")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingClearToken)
				if err := _Staking.contract.UnpackLog(event, "clearToken", log); err != nil {
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

// ParseClearToken is a log parse operation binding the contract event 0x960427a4fde284dcc8da2c51a43d210cfd7ed6f4328a24ba33eee49910db564b.
//
// Solidity: event clearToken(address TokenAddressCleared, uint256 Amount)
func (_Staking *StakingFilterer) ParseClearToken(log types.Log) (*StakingClearToken, error) {
	event := new(StakingClearToken)
	if err := _Staking.contract.UnpackLog(event, "clearToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingTokensBoughtIterator is returned from FilterTokensBought and is used to iterate over the raw logs and unpacked data for TokensBought events raised by the Staking contract.
type StakingTokensBoughtIterator struct {
	Event *StakingTokensBought // Event containing the contract specifics and raw log

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
func (it *StakingTokensBoughtIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingTokensBought)
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
		it.Event = new(StakingTokensBought)
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
func (it *StakingTokensBoughtIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingTokensBoughtIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingTokensBought represents a TokensBought event raised by the Staking contract.
type StakingTokensBought struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTokensBought is a free log retrieval operation binding the contract event 0x6ebf55d0cee4245affe84127b666e1c5c6614cce8e0de4290cc3210d2b83b630.
//
// Solidity: event tokensBought(address user, uint256 amount)
func (_Staking *StakingFilterer) FilterTokensBought(opts *bind.FilterOpts) (*StakingTokensBoughtIterator, error) {

	logs, sub, err := _Staking.contract.FilterLogs(opts, "tokensBought")
	if err != nil {
		return nil, err
	}
	return &StakingTokensBoughtIterator{contract: _Staking.contract, event: "tokensBought", logs: logs, sub: sub}, nil
}

// WatchTokensBought is a free log subscription operation binding the contract event 0x6ebf55d0cee4245affe84127b666e1c5c6614cce8e0de4290cc3210d2b83b630.
//
// Solidity: event tokensBought(address user, uint256 amount)
func (_Staking *StakingFilterer) WatchTokensBought(opts *bind.WatchOpts, sink chan<- *StakingTokensBought) (event.Subscription, error) {

	logs, sub, err := _Staking.contract.WatchLogs(opts, "tokensBought")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingTokensBought)
				if err := _Staking.contract.UnpackLog(event, "tokensBought", log); err != nil {
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

// ParseTokensBought is a log parse operation binding the contract event 0x6ebf55d0cee4245affe84127b666e1c5c6614cce8e0de4290cc3210d2b83b630.
//
// Solidity: event tokensBought(address user, uint256 amount)
func (_Staking *StakingFilterer) ParseTokensBought(log types.Log) (*StakingTokensBought, error) {
	event := new(StakingTokensBought)
	if err := _Staking.contract.UnpackLog(event, "tokensBought", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
