// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package routers

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

// TraderJoeRouterMetaData contains all meta data concerning the TraderJoeRouter contract.
var TraderJoeRouterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"OWNABLE_DIAMOND_STORAGE_SLOT\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TRADER_JOE_ROUTER_STORAGE_SLOT\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTraderJoeRouter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"traderJoeRouter_\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"traderJoeRouter_\",\"type\":\"address\"}],\"name\":\"setTraderJoeRouterAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountOut_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMax_\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path_\",\"type\":\"address[]\"}],\"name\":\"swapAVAXForExactTokens\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin_\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path_\",\"type\":\"address[]\"}],\"name\":\"swapExactAVAXForTokens\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin_\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path_\",\"type\":\"address[]\"}],\"name\":\"swapExactTokensForAVAX\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin_\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path_\",\"type\":\"address[]\"}],\"name\":\"swapExactTokensForTokensTJ\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountOut_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMax_\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path_\",\"type\":\"address[]\"}],\"name\":\"swapTokensForExactAVAX\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountOut_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMax_\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path_\",\"type\":\"address[]\"}],\"name\":\"swapTokensForExactTokensTJ\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// TraderJoeRouterABI is the input ABI used to generate the binding from.
// Deprecated: Use TraderJoeRouterMetaData.ABI instead.
var TraderJoeRouterABI = TraderJoeRouterMetaData.ABI

// TraderJoeRouter is an auto generated Go binding around an Ethereum contract.
type TraderJoeRouter struct {
	TraderJoeRouterCaller     // Read-only binding to the contract
	TraderJoeRouterTransactor // Write-only binding to the contract
	TraderJoeRouterFilterer   // Log filterer for contract events
}

// TraderJoeRouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type TraderJoeRouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TraderJoeRouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TraderJoeRouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TraderJoeRouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TraderJoeRouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TraderJoeRouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TraderJoeRouterSession struct {
	Contract     *TraderJoeRouter  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TraderJoeRouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TraderJoeRouterCallerSession struct {
	Contract *TraderJoeRouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// TraderJoeRouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TraderJoeRouterTransactorSession struct {
	Contract     *TraderJoeRouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// TraderJoeRouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type TraderJoeRouterRaw struct {
	Contract *TraderJoeRouter // Generic contract binding to access the raw methods on
}

// TraderJoeRouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TraderJoeRouterCallerRaw struct {
	Contract *TraderJoeRouterCaller // Generic read-only contract binding to access the raw methods on
}

// TraderJoeRouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TraderJoeRouterTransactorRaw struct {
	Contract *TraderJoeRouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTraderJoeRouter creates a new instance of TraderJoeRouter, bound to a specific deployed contract.
func NewTraderJoeRouter(address common.Address, backend bind.ContractBackend) (*TraderJoeRouter, error) {
	contract, err := bindTraderJoeRouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TraderJoeRouter{TraderJoeRouterCaller: TraderJoeRouterCaller{contract: contract}, TraderJoeRouterTransactor: TraderJoeRouterTransactor{contract: contract}, TraderJoeRouterFilterer: TraderJoeRouterFilterer{contract: contract}}, nil
}

// NewTraderJoeRouterCaller creates a new read-only instance of TraderJoeRouter, bound to a specific deployed contract.
func NewTraderJoeRouterCaller(address common.Address, caller bind.ContractCaller) (*TraderJoeRouterCaller, error) {
	contract, err := bindTraderJoeRouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TraderJoeRouterCaller{contract: contract}, nil
}

// NewTraderJoeRouterTransactor creates a new write-only instance of TraderJoeRouter, bound to a specific deployed contract.
func NewTraderJoeRouterTransactor(address common.Address, transactor bind.ContractTransactor) (*TraderJoeRouterTransactor, error) {
	contract, err := bindTraderJoeRouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TraderJoeRouterTransactor{contract: contract}, nil
}

// NewTraderJoeRouterFilterer creates a new log filterer instance of TraderJoeRouter, bound to a specific deployed contract.
func NewTraderJoeRouterFilterer(address common.Address, filterer bind.ContractFilterer) (*TraderJoeRouterFilterer, error) {
	contract, err := bindTraderJoeRouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TraderJoeRouterFilterer{contract: contract}, nil
}

// bindTraderJoeRouter binds a generic wrapper to an already deployed contract.
func bindTraderJoeRouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TraderJoeRouterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TraderJoeRouter *TraderJoeRouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TraderJoeRouter.Contract.TraderJoeRouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TraderJoeRouter *TraderJoeRouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TraderJoeRouter.Contract.TraderJoeRouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TraderJoeRouter *TraderJoeRouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TraderJoeRouter.Contract.TraderJoeRouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TraderJoeRouter *TraderJoeRouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TraderJoeRouter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TraderJoeRouter *TraderJoeRouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TraderJoeRouter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TraderJoeRouter *TraderJoeRouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TraderJoeRouter.Contract.contract.Transact(opts, method, params...)
}

// OWNABLEDIAMONDSTORAGESLOT is a free data retrieval call binding the contract method 0xe3e077ad.
//
// Solidity: function OWNABLE_DIAMOND_STORAGE_SLOT() view returns(bytes32)
func (_TraderJoeRouter *TraderJoeRouterCaller) OWNABLEDIAMONDSTORAGESLOT(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TraderJoeRouter.contract.Call(opts, &out, "OWNABLE_DIAMOND_STORAGE_SLOT")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OWNABLEDIAMONDSTORAGESLOT is a free data retrieval call binding the contract method 0xe3e077ad.
//
// Solidity: function OWNABLE_DIAMOND_STORAGE_SLOT() view returns(bytes32)
func (_TraderJoeRouter *TraderJoeRouterSession) OWNABLEDIAMONDSTORAGESLOT() ([32]byte, error) {
	return _TraderJoeRouter.Contract.OWNABLEDIAMONDSTORAGESLOT(&_TraderJoeRouter.CallOpts)
}

// OWNABLEDIAMONDSTORAGESLOT is a free data retrieval call binding the contract method 0xe3e077ad.
//
// Solidity: function OWNABLE_DIAMOND_STORAGE_SLOT() view returns(bytes32)
func (_TraderJoeRouter *TraderJoeRouterCallerSession) OWNABLEDIAMONDSTORAGESLOT() ([32]byte, error) {
	return _TraderJoeRouter.Contract.OWNABLEDIAMONDSTORAGESLOT(&_TraderJoeRouter.CallOpts)
}

// TRADERJOEROUTERSTORAGESLOT is a free data retrieval call binding the contract method 0x7d35e208.
//
// Solidity: function TRADER_JOE_ROUTER_STORAGE_SLOT() view returns(bytes32)
func (_TraderJoeRouter *TraderJoeRouterCaller) TRADERJOEROUTERSTORAGESLOT(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TraderJoeRouter.contract.Call(opts, &out, "TRADER_JOE_ROUTER_STORAGE_SLOT")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TRADERJOEROUTERSTORAGESLOT is a free data retrieval call binding the contract method 0x7d35e208.
//
// Solidity: function TRADER_JOE_ROUTER_STORAGE_SLOT() view returns(bytes32)
func (_TraderJoeRouter *TraderJoeRouterSession) TRADERJOEROUTERSTORAGESLOT() ([32]byte, error) {
	return _TraderJoeRouter.Contract.TRADERJOEROUTERSTORAGESLOT(&_TraderJoeRouter.CallOpts)
}

// TRADERJOEROUTERSTORAGESLOT is a free data retrieval call binding the contract method 0x7d35e208.
//
// Solidity: function TRADER_JOE_ROUTER_STORAGE_SLOT() view returns(bytes32)
func (_TraderJoeRouter *TraderJoeRouterCallerSession) TRADERJOEROUTERSTORAGESLOT() ([32]byte, error) {
	return _TraderJoeRouter.Contract.TRADERJOEROUTERSTORAGESLOT(&_TraderJoeRouter.CallOpts)
}

// GetTraderJoeRouter is a free data retrieval call binding the contract method 0x3a144e2d.
//
// Solidity: function getTraderJoeRouter() view returns(address traderJoeRouter_)
func (_TraderJoeRouter *TraderJoeRouterCaller) GetTraderJoeRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TraderJoeRouter.contract.Call(opts, &out, "getTraderJoeRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetTraderJoeRouter is a free data retrieval call binding the contract method 0x3a144e2d.
//
// Solidity: function getTraderJoeRouter() view returns(address traderJoeRouter_)
func (_TraderJoeRouter *TraderJoeRouterSession) GetTraderJoeRouter() (common.Address, error) {
	return _TraderJoeRouter.Contract.GetTraderJoeRouter(&_TraderJoeRouter.CallOpts)
}

// GetTraderJoeRouter is a free data retrieval call binding the contract method 0x3a144e2d.
//
// Solidity: function getTraderJoeRouter() view returns(address traderJoeRouter_)
func (_TraderJoeRouter *TraderJoeRouterCallerSession) GetTraderJoeRouter() (common.Address, error) {
	return _TraderJoeRouter.Contract.GetTraderJoeRouter(&_TraderJoeRouter.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TraderJoeRouter *TraderJoeRouterCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TraderJoeRouter.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TraderJoeRouter *TraderJoeRouterSession) Owner() (common.Address, error) {
	return _TraderJoeRouter.Contract.Owner(&_TraderJoeRouter.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TraderJoeRouter *TraderJoeRouterCallerSession) Owner() (common.Address, error) {
	return _TraderJoeRouter.Contract.Owner(&_TraderJoeRouter.CallOpts)
}

// SetTraderJoeRouterAddress is a paid mutator transaction binding the contract method 0x4cc64fc4.
//
// Solidity: function setTraderJoeRouterAddress(address traderJoeRouter_) returns()
func (_TraderJoeRouter *TraderJoeRouterTransactor) SetTraderJoeRouterAddress(opts *bind.TransactOpts, traderJoeRouter_ common.Address) (*types.Transaction, error) {
	return _TraderJoeRouter.contract.Transact(opts, "setTraderJoeRouterAddress", traderJoeRouter_)
}

// SetTraderJoeRouterAddress is a paid mutator transaction binding the contract method 0x4cc64fc4.
//
// Solidity: function setTraderJoeRouterAddress(address traderJoeRouter_) returns()
func (_TraderJoeRouter *TraderJoeRouterSession) SetTraderJoeRouterAddress(traderJoeRouter_ common.Address) (*types.Transaction, error) {
	return _TraderJoeRouter.Contract.SetTraderJoeRouterAddress(&_TraderJoeRouter.TransactOpts, traderJoeRouter_)
}

// SetTraderJoeRouterAddress is a paid mutator transaction binding the contract method 0x4cc64fc4.
//
// Solidity: function setTraderJoeRouterAddress(address traderJoeRouter_) returns()
func (_TraderJoeRouter *TraderJoeRouterTransactorSession) SetTraderJoeRouterAddress(traderJoeRouter_ common.Address) (*types.Transaction, error) {
	return _TraderJoeRouter.Contract.SetTraderJoeRouterAddress(&_TraderJoeRouter.TransactOpts, traderJoeRouter_)
}

// SwapAVAXForExactTokens is a paid mutator transaction binding the contract method 0x23a012f0.
//
// Solidity: function swapAVAXForExactTokens(address receiver_, uint256 amountOut_, uint256 amountInMax_, address[] path_) payable returns()
func (_TraderJoeRouter *TraderJoeRouterTransactor) SwapAVAXForExactTokens(opts *bind.TransactOpts, receiver_ common.Address, amountOut_ *big.Int, amountInMax_ *big.Int, path_ []common.Address) (*types.Transaction, error) {
	return _TraderJoeRouter.contract.Transact(opts, "swapAVAXForExactTokens", receiver_, amountOut_, amountInMax_, path_)
}

// SwapAVAXForExactTokens is a paid mutator transaction binding the contract method 0x23a012f0.
//
// Solidity: function swapAVAXForExactTokens(address receiver_, uint256 amountOut_, uint256 amountInMax_, address[] path_) payable returns()
func (_TraderJoeRouter *TraderJoeRouterSession) SwapAVAXForExactTokens(receiver_ common.Address, amountOut_ *big.Int, amountInMax_ *big.Int, path_ []common.Address) (*types.Transaction, error) {
	return _TraderJoeRouter.Contract.SwapAVAXForExactTokens(&_TraderJoeRouter.TransactOpts, receiver_, amountOut_, amountInMax_, path_)
}

// SwapAVAXForExactTokens is a paid mutator transaction binding the contract method 0x23a012f0.
//
// Solidity: function swapAVAXForExactTokens(address receiver_, uint256 amountOut_, uint256 amountInMax_, address[] path_) payable returns()
func (_TraderJoeRouter *TraderJoeRouterTransactorSession) SwapAVAXForExactTokens(receiver_ common.Address, amountOut_ *big.Int, amountInMax_ *big.Int, path_ []common.Address) (*types.Transaction, error) {
	return _TraderJoeRouter.Contract.SwapAVAXForExactTokens(&_TraderJoeRouter.TransactOpts, receiver_, amountOut_, amountInMax_, path_)
}

// SwapExactAVAXForTokens is a paid mutator transaction binding the contract method 0x519c1647.
//
// Solidity: function swapExactAVAXForTokens(address receiver_, uint256 amountIn_, uint256 amountOutMin_, address[] path_) payable returns()
func (_TraderJoeRouter *TraderJoeRouterTransactor) SwapExactAVAXForTokens(opts *bind.TransactOpts, receiver_ common.Address, amountIn_ *big.Int, amountOutMin_ *big.Int, path_ []common.Address) (*types.Transaction, error) {
	return _TraderJoeRouter.contract.Transact(opts, "swapExactAVAXForTokens", receiver_, amountIn_, amountOutMin_, path_)
}

// SwapExactAVAXForTokens is a paid mutator transaction binding the contract method 0x519c1647.
//
// Solidity: function swapExactAVAXForTokens(address receiver_, uint256 amountIn_, uint256 amountOutMin_, address[] path_) payable returns()
func (_TraderJoeRouter *TraderJoeRouterSession) SwapExactAVAXForTokens(receiver_ common.Address, amountIn_ *big.Int, amountOutMin_ *big.Int, path_ []common.Address) (*types.Transaction, error) {
	return _TraderJoeRouter.Contract.SwapExactAVAXForTokens(&_TraderJoeRouter.TransactOpts, receiver_, amountIn_, amountOutMin_, path_)
}

// SwapExactAVAXForTokens is a paid mutator transaction binding the contract method 0x519c1647.
//
// Solidity: function swapExactAVAXForTokens(address receiver_, uint256 amountIn_, uint256 amountOutMin_, address[] path_) payable returns()
func (_TraderJoeRouter *TraderJoeRouterTransactorSession) SwapExactAVAXForTokens(receiver_ common.Address, amountIn_ *big.Int, amountOutMin_ *big.Int, path_ []common.Address) (*types.Transaction, error) {
	return _TraderJoeRouter.Contract.SwapExactAVAXForTokens(&_TraderJoeRouter.TransactOpts, receiver_, amountIn_, amountOutMin_, path_)
}

// SwapExactTokensForAVAX is a paid mutator transaction binding the contract method 0xd00dac7c.
//
// Solidity: function swapExactTokensForAVAX(address receiver_, uint256 amountIn_, uint256 amountOutMin_, address[] path_) payable returns()
func (_TraderJoeRouter *TraderJoeRouterTransactor) SwapExactTokensForAVAX(opts *bind.TransactOpts, receiver_ common.Address, amountIn_ *big.Int, amountOutMin_ *big.Int, path_ []common.Address) (*types.Transaction, error) {
	return _TraderJoeRouter.contract.Transact(opts, "swapExactTokensForAVAX", receiver_, amountIn_, amountOutMin_, path_)
}

// SwapExactTokensForAVAX is a paid mutator transaction binding the contract method 0xd00dac7c.
//
// Solidity: function swapExactTokensForAVAX(address receiver_, uint256 amountIn_, uint256 amountOutMin_, address[] path_) payable returns()
func (_TraderJoeRouter *TraderJoeRouterSession) SwapExactTokensForAVAX(receiver_ common.Address, amountIn_ *big.Int, amountOutMin_ *big.Int, path_ []common.Address) (*types.Transaction, error) {
	return _TraderJoeRouter.Contract.SwapExactTokensForAVAX(&_TraderJoeRouter.TransactOpts, receiver_, amountIn_, amountOutMin_, path_)
}

// SwapExactTokensForAVAX is a paid mutator transaction binding the contract method 0xd00dac7c.
//
// Solidity: function swapExactTokensForAVAX(address receiver_, uint256 amountIn_, uint256 amountOutMin_, address[] path_) payable returns()
func (_TraderJoeRouter *TraderJoeRouterTransactorSession) SwapExactTokensForAVAX(receiver_ common.Address, amountIn_ *big.Int, amountOutMin_ *big.Int, path_ []common.Address) (*types.Transaction, error) {
	return _TraderJoeRouter.Contract.SwapExactTokensForAVAX(&_TraderJoeRouter.TransactOpts, receiver_, amountIn_, amountOutMin_, path_)
}

// SwapExactTokensForTokensTJ is a paid mutator transaction binding the contract method 0xe691b80a.
//
// Solidity: function swapExactTokensForTokensTJ(address receiver_, uint256 amountIn_, uint256 amountOutMin_, address[] path_) payable returns()
func (_TraderJoeRouter *TraderJoeRouterTransactor) SwapExactTokensForTokensTJ(opts *bind.TransactOpts, receiver_ common.Address, amountIn_ *big.Int, amountOutMin_ *big.Int, path_ []common.Address) (*types.Transaction, error) {
	return _TraderJoeRouter.contract.Transact(opts, "swapExactTokensForTokensTJ", receiver_, amountIn_, amountOutMin_, path_)
}

// SwapExactTokensForTokensTJ is a paid mutator transaction binding the contract method 0xe691b80a.
//
// Solidity: function swapExactTokensForTokensTJ(address receiver_, uint256 amountIn_, uint256 amountOutMin_, address[] path_) payable returns()
func (_TraderJoeRouter *TraderJoeRouterSession) SwapExactTokensForTokensTJ(receiver_ common.Address, amountIn_ *big.Int, amountOutMin_ *big.Int, path_ []common.Address) (*types.Transaction, error) {
	return _TraderJoeRouter.Contract.SwapExactTokensForTokensTJ(&_TraderJoeRouter.TransactOpts, receiver_, amountIn_, amountOutMin_, path_)
}

// SwapExactTokensForTokensTJ is a paid mutator transaction binding the contract method 0xe691b80a.
//
// Solidity: function swapExactTokensForTokensTJ(address receiver_, uint256 amountIn_, uint256 amountOutMin_, address[] path_) payable returns()
func (_TraderJoeRouter *TraderJoeRouterTransactorSession) SwapExactTokensForTokensTJ(receiver_ common.Address, amountIn_ *big.Int, amountOutMin_ *big.Int, path_ []common.Address) (*types.Transaction, error) {
	return _TraderJoeRouter.Contract.SwapExactTokensForTokensTJ(&_TraderJoeRouter.TransactOpts, receiver_, amountIn_, amountOutMin_, path_)
}

// SwapTokensForExactAVAX is a paid mutator transaction binding the contract method 0xe8a30798.
//
// Solidity: function swapTokensForExactAVAX(address receiver_, uint256 amountOut_, uint256 amountInMax_, address[] path_) payable returns()
func (_TraderJoeRouter *TraderJoeRouterTransactor) SwapTokensForExactAVAX(opts *bind.TransactOpts, receiver_ common.Address, amountOut_ *big.Int, amountInMax_ *big.Int, path_ []common.Address) (*types.Transaction, error) {
	return _TraderJoeRouter.contract.Transact(opts, "swapTokensForExactAVAX", receiver_, amountOut_, amountInMax_, path_)
}

// SwapTokensForExactAVAX is a paid mutator transaction binding the contract method 0xe8a30798.
//
// Solidity: function swapTokensForExactAVAX(address receiver_, uint256 amountOut_, uint256 amountInMax_, address[] path_) payable returns()
func (_TraderJoeRouter *TraderJoeRouterSession) SwapTokensForExactAVAX(receiver_ common.Address, amountOut_ *big.Int, amountInMax_ *big.Int, path_ []common.Address) (*types.Transaction, error) {
	return _TraderJoeRouter.Contract.SwapTokensForExactAVAX(&_TraderJoeRouter.TransactOpts, receiver_, amountOut_, amountInMax_, path_)
}

// SwapTokensForExactAVAX is a paid mutator transaction binding the contract method 0xe8a30798.
//
// Solidity: function swapTokensForExactAVAX(address receiver_, uint256 amountOut_, uint256 amountInMax_, address[] path_) payable returns()
func (_TraderJoeRouter *TraderJoeRouterTransactorSession) SwapTokensForExactAVAX(receiver_ common.Address, amountOut_ *big.Int, amountInMax_ *big.Int, path_ []common.Address) (*types.Transaction, error) {
	return _TraderJoeRouter.Contract.SwapTokensForExactAVAX(&_TraderJoeRouter.TransactOpts, receiver_, amountOut_, amountInMax_, path_)
}

// SwapTokensForExactTokensTJ is a paid mutator transaction binding the contract method 0x9b1d014e.
//
// Solidity: function swapTokensForExactTokensTJ(address receiver_, uint256 amountOut_, uint256 amountInMax_, address[] path_) payable returns()
func (_TraderJoeRouter *TraderJoeRouterTransactor) SwapTokensForExactTokensTJ(opts *bind.TransactOpts, receiver_ common.Address, amountOut_ *big.Int, amountInMax_ *big.Int, path_ []common.Address) (*types.Transaction, error) {
	return _TraderJoeRouter.contract.Transact(opts, "swapTokensForExactTokensTJ", receiver_, amountOut_, amountInMax_, path_)
}

// SwapTokensForExactTokensTJ is a paid mutator transaction binding the contract method 0x9b1d014e.
//
// Solidity: function swapTokensForExactTokensTJ(address receiver_, uint256 amountOut_, uint256 amountInMax_, address[] path_) payable returns()
func (_TraderJoeRouter *TraderJoeRouterSession) SwapTokensForExactTokensTJ(receiver_ common.Address, amountOut_ *big.Int, amountInMax_ *big.Int, path_ []common.Address) (*types.Transaction, error) {
	return _TraderJoeRouter.Contract.SwapTokensForExactTokensTJ(&_TraderJoeRouter.TransactOpts, receiver_, amountOut_, amountInMax_, path_)
}

// SwapTokensForExactTokensTJ is a paid mutator transaction binding the contract method 0x9b1d014e.
//
// Solidity: function swapTokensForExactTokensTJ(address receiver_, uint256 amountOut_, uint256 amountInMax_, address[] path_) payable returns()
func (_TraderJoeRouter *TraderJoeRouterTransactorSession) SwapTokensForExactTokensTJ(receiver_ common.Address, amountOut_ *big.Int, amountInMax_ *big.Int, path_ []common.Address) (*types.Transaction, error) {
	return _TraderJoeRouter.Contract.SwapTokensForExactTokensTJ(&_TraderJoeRouter.TransactOpts, receiver_, amountOut_, amountInMax_, path_)
}
