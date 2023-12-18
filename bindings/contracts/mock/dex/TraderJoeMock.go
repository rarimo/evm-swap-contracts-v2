// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dex

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

// TraderJoeMockMetaData contains all meta data concerning the TraderJoeMock contract.
var TraderJoeMockMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"wavax_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"WRAPPED_NATIVE\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut_\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path_\",\"type\":\"address[]\"}],\"name\":\"_getAmountsIn\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts_\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn_\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path_\",\"type\":\"address[]\"}],\"name\":\"_getAmountsOut\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts_\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB_\",\"type\":\"address\"}],\"name\":\"enablePair\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"}],\"name\":\"setReserve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut_\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path_\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"receiver_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"swapAVAXForExactTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts_\",\"type\":\"uint256[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOutMin_\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path_\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"receiver_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"swapExactAVAXForTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts_\",\"type\":\"uint256[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin_\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path_\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"receiver_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForAVAX\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts_\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin_\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path_\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"receiver_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts_\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMax_\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path_\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"receiver_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"swapTokensForExactAVAX\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts_\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMax_\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path_\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"receiver_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"swapTokensForExactTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts_\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// TraderJoeMockABI is the input ABI used to generate the binding from.
// Deprecated: Use TraderJoeMockMetaData.ABI instead.
var TraderJoeMockABI = TraderJoeMockMetaData.ABI

// TraderJoeMock is an auto generated Go binding around an Ethereum contract.
type TraderJoeMock struct {
	TraderJoeMockCaller     // Read-only binding to the contract
	TraderJoeMockTransactor // Write-only binding to the contract
	TraderJoeMockFilterer   // Log filterer for contract events
}

// TraderJoeMockCaller is an auto generated read-only Go binding around an Ethereum contract.
type TraderJoeMockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TraderJoeMockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TraderJoeMockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TraderJoeMockFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TraderJoeMockFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TraderJoeMockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TraderJoeMockSession struct {
	Contract     *TraderJoeMock    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TraderJoeMockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TraderJoeMockCallerSession struct {
	Contract *TraderJoeMockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// TraderJoeMockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TraderJoeMockTransactorSession struct {
	Contract     *TraderJoeMockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// TraderJoeMockRaw is an auto generated low-level Go binding around an Ethereum contract.
type TraderJoeMockRaw struct {
	Contract *TraderJoeMock // Generic contract binding to access the raw methods on
}

// TraderJoeMockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TraderJoeMockCallerRaw struct {
	Contract *TraderJoeMockCaller // Generic read-only contract binding to access the raw methods on
}

// TraderJoeMockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TraderJoeMockTransactorRaw struct {
	Contract *TraderJoeMockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTraderJoeMock creates a new instance of TraderJoeMock, bound to a specific deployed contract.
func NewTraderJoeMock(address common.Address, backend bind.ContractBackend) (*TraderJoeMock, error) {
	contract, err := bindTraderJoeMock(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TraderJoeMock{TraderJoeMockCaller: TraderJoeMockCaller{contract: contract}, TraderJoeMockTransactor: TraderJoeMockTransactor{contract: contract}, TraderJoeMockFilterer: TraderJoeMockFilterer{contract: contract}}, nil
}

// NewTraderJoeMockCaller creates a new read-only instance of TraderJoeMock, bound to a specific deployed contract.
func NewTraderJoeMockCaller(address common.Address, caller bind.ContractCaller) (*TraderJoeMockCaller, error) {
	contract, err := bindTraderJoeMock(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TraderJoeMockCaller{contract: contract}, nil
}

// NewTraderJoeMockTransactor creates a new write-only instance of TraderJoeMock, bound to a specific deployed contract.
func NewTraderJoeMockTransactor(address common.Address, transactor bind.ContractTransactor) (*TraderJoeMockTransactor, error) {
	contract, err := bindTraderJoeMock(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TraderJoeMockTransactor{contract: contract}, nil
}

// NewTraderJoeMockFilterer creates a new log filterer instance of TraderJoeMock, bound to a specific deployed contract.
func NewTraderJoeMockFilterer(address common.Address, filterer bind.ContractFilterer) (*TraderJoeMockFilterer, error) {
	contract, err := bindTraderJoeMock(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TraderJoeMockFilterer{contract: contract}, nil
}

// bindTraderJoeMock binds a generic wrapper to an already deployed contract.
func bindTraderJoeMock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TraderJoeMockMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TraderJoeMock *TraderJoeMockRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TraderJoeMock.Contract.TraderJoeMockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TraderJoeMock *TraderJoeMockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TraderJoeMock.Contract.TraderJoeMockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TraderJoeMock *TraderJoeMockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TraderJoeMock.Contract.TraderJoeMockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TraderJoeMock *TraderJoeMockCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TraderJoeMock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TraderJoeMock *TraderJoeMockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TraderJoeMock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TraderJoeMock *TraderJoeMockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TraderJoeMock.Contract.contract.Transact(opts, method, params...)
}

// WRAPPEDNATIVE is a free data retrieval call binding the contract method 0xd999984d.
//
// Solidity: function WRAPPED_NATIVE() view returns(address)
func (_TraderJoeMock *TraderJoeMockCaller) WRAPPEDNATIVE(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TraderJoeMock.contract.Call(opts, &out, "WRAPPED_NATIVE")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WRAPPEDNATIVE is a free data retrieval call binding the contract method 0xd999984d.
//
// Solidity: function WRAPPED_NATIVE() view returns(address)
func (_TraderJoeMock *TraderJoeMockSession) WRAPPEDNATIVE() (common.Address, error) {
	return _TraderJoeMock.Contract.WRAPPEDNATIVE(&_TraderJoeMock.CallOpts)
}

// WRAPPEDNATIVE is a free data retrieval call binding the contract method 0xd999984d.
//
// Solidity: function WRAPPED_NATIVE() view returns(address)
func (_TraderJoeMock *TraderJoeMockCallerSession) WRAPPEDNATIVE() (common.Address, error) {
	return _TraderJoeMock.Contract.WRAPPEDNATIVE(&_TraderJoeMock.CallOpts)
}

// GetAmountsIn is a free data retrieval call binding the contract method 0xbbf9a05e.
//
// Solidity: function _getAmountsIn(uint256 amountOut_, address[] path_) view returns(uint256[] amounts_)
func (_TraderJoeMock *TraderJoeMockCaller) GetAmountsIn(opts *bind.CallOpts, amountOut_ *big.Int, path_ []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _TraderJoeMock.contract.Call(opts, &out, "_getAmountsIn", amountOut_, path_)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetAmountsIn is a free data retrieval call binding the contract method 0xbbf9a05e.
//
// Solidity: function _getAmountsIn(uint256 amountOut_, address[] path_) view returns(uint256[] amounts_)
func (_TraderJoeMock *TraderJoeMockSession) GetAmountsIn(amountOut_ *big.Int, path_ []common.Address) ([]*big.Int, error) {
	return _TraderJoeMock.Contract.GetAmountsIn(&_TraderJoeMock.CallOpts, amountOut_, path_)
}

// GetAmountsIn is a free data retrieval call binding the contract method 0xbbf9a05e.
//
// Solidity: function _getAmountsIn(uint256 amountOut_, address[] path_) view returns(uint256[] amounts_)
func (_TraderJoeMock *TraderJoeMockCallerSession) GetAmountsIn(amountOut_ *big.Int, path_ []common.Address) ([]*big.Int, error) {
	return _TraderJoeMock.Contract.GetAmountsIn(&_TraderJoeMock.CallOpts, amountOut_, path_)
}

// GetAmountsOut is a free data retrieval call binding the contract method 0x8a410396.
//
// Solidity: function _getAmountsOut(uint256 amountIn_, address[] path_) view returns(uint256[] amounts_)
func (_TraderJoeMock *TraderJoeMockCaller) GetAmountsOut(opts *bind.CallOpts, amountIn_ *big.Int, path_ []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _TraderJoeMock.contract.Call(opts, &out, "_getAmountsOut", amountIn_, path_)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetAmountsOut is a free data retrieval call binding the contract method 0x8a410396.
//
// Solidity: function _getAmountsOut(uint256 amountIn_, address[] path_) view returns(uint256[] amounts_)
func (_TraderJoeMock *TraderJoeMockSession) GetAmountsOut(amountIn_ *big.Int, path_ []common.Address) ([]*big.Int, error) {
	return _TraderJoeMock.Contract.GetAmountsOut(&_TraderJoeMock.CallOpts, amountIn_, path_)
}

// GetAmountsOut is a free data retrieval call binding the contract method 0x8a410396.
//
// Solidity: function _getAmountsOut(uint256 amountIn_, address[] path_) view returns(uint256[] amounts_)
func (_TraderJoeMock *TraderJoeMockCallerSession) GetAmountsOut(amountIn_ *big.Int, path_ []common.Address) ([]*big.Int, error) {
	return _TraderJoeMock.Contract.GetAmountsOut(&_TraderJoeMock.CallOpts, amountIn_, path_)
}

// EnablePair is a paid mutator transaction binding the contract method 0xdfb5061c.
//
// Solidity: function enablePair(address tokenA_, address tokenB_) returns()
func (_TraderJoeMock *TraderJoeMockTransactor) EnablePair(opts *bind.TransactOpts, tokenA_ common.Address, tokenB_ common.Address) (*types.Transaction, error) {
	return _TraderJoeMock.contract.Transact(opts, "enablePair", tokenA_, tokenB_)
}

// EnablePair is a paid mutator transaction binding the contract method 0xdfb5061c.
//
// Solidity: function enablePair(address tokenA_, address tokenB_) returns()
func (_TraderJoeMock *TraderJoeMockSession) EnablePair(tokenA_ common.Address, tokenB_ common.Address) (*types.Transaction, error) {
	return _TraderJoeMock.Contract.EnablePair(&_TraderJoeMock.TransactOpts, tokenA_, tokenB_)
}

// EnablePair is a paid mutator transaction binding the contract method 0xdfb5061c.
//
// Solidity: function enablePair(address tokenA_, address tokenB_) returns()
func (_TraderJoeMock *TraderJoeMockTransactorSession) EnablePair(tokenA_ common.Address, tokenB_ common.Address) (*types.Transaction, error) {
	return _TraderJoeMock.Contract.EnablePair(&_TraderJoeMock.TransactOpts, tokenA_, tokenB_)
}

// SetReserve is a paid mutator transaction binding the contract method 0x78250b63.
//
// Solidity: function setReserve(address token_, uint256 amount_) returns()
func (_TraderJoeMock *TraderJoeMockTransactor) SetReserve(opts *bind.TransactOpts, token_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _TraderJoeMock.contract.Transact(opts, "setReserve", token_, amount_)
}

// SetReserve is a paid mutator transaction binding the contract method 0x78250b63.
//
// Solidity: function setReserve(address token_, uint256 amount_) returns()
func (_TraderJoeMock *TraderJoeMockSession) SetReserve(token_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _TraderJoeMock.Contract.SetReserve(&_TraderJoeMock.TransactOpts, token_, amount_)
}

// SetReserve is a paid mutator transaction binding the contract method 0x78250b63.
//
// Solidity: function setReserve(address token_, uint256 amount_) returns()
func (_TraderJoeMock *TraderJoeMockTransactorSession) SetReserve(token_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _TraderJoeMock.Contract.SetReserve(&_TraderJoeMock.TransactOpts, token_, amount_)
}

// SwapAVAXForExactTokens is a paid mutator transaction binding the contract method 0x8a657e67.
//
// Solidity: function swapAVAXForExactTokens(uint256 amountOut_, address[] path_, address receiver_, uint256 ) payable returns(uint256[] amounts_)
func (_TraderJoeMock *TraderJoeMockTransactor) SwapAVAXForExactTokens(opts *bind.TransactOpts, amountOut_ *big.Int, path_ []common.Address, receiver_ common.Address, arg3 *big.Int) (*types.Transaction, error) {
	return _TraderJoeMock.contract.Transact(opts, "swapAVAXForExactTokens", amountOut_, path_, receiver_, arg3)
}

// SwapAVAXForExactTokens is a paid mutator transaction binding the contract method 0x8a657e67.
//
// Solidity: function swapAVAXForExactTokens(uint256 amountOut_, address[] path_, address receiver_, uint256 ) payable returns(uint256[] amounts_)
func (_TraderJoeMock *TraderJoeMockSession) SwapAVAXForExactTokens(amountOut_ *big.Int, path_ []common.Address, receiver_ common.Address, arg3 *big.Int) (*types.Transaction, error) {
	return _TraderJoeMock.Contract.SwapAVAXForExactTokens(&_TraderJoeMock.TransactOpts, amountOut_, path_, receiver_, arg3)
}

// SwapAVAXForExactTokens is a paid mutator transaction binding the contract method 0x8a657e67.
//
// Solidity: function swapAVAXForExactTokens(uint256 amountOut_, address[] path_, address receiver_, uint256 ) payable returns(uint256[] amounts_)
func (_TraderJoeMock *TraderJoeMockTransactorSession) SwapAVAXForExactTokens(amountOut_ *big.Int, path_ []common.Address, receiver_ common.Address, arg3 *big.Int) (*types.Transaction, error) {
	return _TraderJoeMock.Contract.SwapAVAXForExactTokens(&_TraderJoeMock.TransactOpts, amountOut_, path_, receiver_, arg3)
}

// SwapExactAVAXForTokens is a paid mutator transaction binding the contract method 0xa2a1623d.
//
// Solidity: function swapExactAVAXForTokens(uint256 amountOutMin_, address[] path_, address receiver_, uint256 ) payable returns(uint256[] amounts_)
func (_TraderJoeMock *TraderJoeMockTransactor) SwapExactAVAXForTokens(opts *bind.TransactOpts, amountOutMin_ *big.Int, path_ []common.Address, receiver_ common.Address, arg3 *big.Int) (*types.Transaction, error) {
	return _TraderJoeMock.contract.Transact(opts, "swapExactAVAXForTokens", amountOutMin_, path_, receiver_, arg3)
}

// SwapExactAVAXForTokens is a paid mutator transaction binding the contract method 0xa2a1623d.
//
// Solidity: function swapExactAVAXForTokens(uint256 amountOutMin_, address[] path_, address receiver_, uint256 ) payable returns(uint256[] amounts_)
func (_TraderJoeMock *TraderJoeMockSession) SwapExactAVAXForTokens(amountOutMin_ *big.Int, path_ []common.Address, receiver_ common.Address, arg3 *big.Int) (*types.Transaction, error) {
	return _TraderJoeMock.Contract.SwapExactAVAXForTokens(&_TraderJoeMock.TransactOpts, amountOutMin_, path_, receiver_, arg3)
}

// SwapExactAVAXForTokens is a paid mutator transaction binding the contract method 0xa2a1623d.
//
// Solidity: function swapExactAVAXForTokens(uint256 amountOutMin_, address[] path_, address receiver_, uint256 ) payable returns(uint256[] amounts_)
func (_TraderJoeMock *TraderJoeMockTransactorSession) SwapExactAVAXForTokens(amountOutMin_ *big.Int, path_ []common.Address, receiver_ common.Address, arg3 *big.Int) (*types.Transaction, error) {
	return _TraderJoeMock.Contract.SwapExactAVAXForTokens(&_TraderJoeMock.TransactOpts, amountOutMin_, path_, receiver_, arg3)
}

// SwapExactTokensForAVAX is a paid mutator transaction binding the contract method 0x676528d1.
//
// Solidity: function swapExactTokensForAVAX(uint256 amountIn_, uint256 amountOutMin_, address[] path_, address receiver_, uint256 ) returns(uint256[] amounts_)
func (_TraderJoeMock *TraderJoeMockTransactor) SwapExactTokensForAVAX(opts *bind.TransactOpts, amountIn_ *big.Int, amountOutMin_ *big.Int, path_ []common.Address, receiver_ common.Address, arg4 *big.Int) (*types.Transaction, error) {
	return _TraderJoeMock.contract.Transact(opts, "swapExactTokensForAVAX", amountIn_, amountOutMin_, path_, receiver_, arg4)
}

// SwapExactTokensForAVAX is a paid mutator transaction binding the contract method 0x676528d1.
//
// Solidity: function swapExactTokensForAVAX(uint256 amountIn_, uint256 amountOutMin_, address[] path_, address receiver_, uint256 ) returns(uint256[] amounts_)
func (_TraderJoeMock *TraderJoeMockSession) SwapExactTokensForAVAX(amountIn_ *big.Int, amountOutMin_ *big.Int, path_ []common.Address, receiver_ common.Address, arg4 *big.Int) (*types.Transaction, error) {
	return _TraderJoeMock.Contract.SwapExactTokensForAVAX(&_TraderJoeMock.TransactOpts, amountIn_, amountOutMin_, path_, receiver_, arg4)
}

// SwapExactTokensForAVAX is a paid mutator transaction binding the contract method 0x676528d1.
//
// Solidity: function swapExactTokensForAVAX(uint256 amountIn_, uint256 amountOutMin_, address[] path_, address receiver_, uint256 ) returns(uint256[] amounts_)
func (_TraderJoeMock *TraderJoeMockTransactorSession) SwapExactTokensForAVAX(amountIn_ *big.Int, amountOutMin_ *big.Int, path_ []common.Address, receiver_ common.Address, arg4 *big.Int) (*types.Transaction, error) {
	return _TraderJoeMock.Contract.SwapExactTokensForAVAX(&_TraderJoeMock.TransactOpts, amountIn_, amountOutMin_, path_, receiver_, arg4)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x38ed1739.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn_, uint256 amountOutMin_, address[] path_, address receiver_, uint256 ) returns(uint256[] amounts_)
func (_TraderJoeMock *TraderJoeMockTransactor) SwapExactTokensForTokens(opts *bind.TransactOpts, amountIn_ *big.Int, amountOutMin_ *big.Int, path_ []common.Address, receiver_ common.Address, arg4 *big.Int) (*types.Transaction, error) {
	return _TraderJoeMock.contract.Transact(opts, "swapExactTokensForTokens", amountIn_, amountOutMin_, path_, receiver_, arg4)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x38ed1739.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn_, uint256 amountOutMin_, address[] path_, address receiver_, uint256 ) returns(uint256[] amounts_)
func (_TraderJoeMock *TraderJoeMockSession) SwapExactTokensForTokens(amountIn_ *big.Int, amountOutMin_ *big.Int, path_ []common.Address, receiver_ common.Address, arg4 *big.Int) (*types.Transaction, error) {
	return _TraderJoeMock.Contract.SwapExactTokensForTokens(&_TraderJoeMock.TransactOpts, amountIn_, amountOutMin_, path_, receiver_, arg4)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x38ed1739.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn_, uint256 amountOutMin_, address[] path_, address receiver_, uint256 ) returns(uint256[] amounts_)
func (_TraderJoeMock *TraderJoeMockTransactorSession) SwapExactTokensForTokens(amountIn_ *big.Int, amountOutMin_ *big.Int, path_ []common.Address, receiver_ common.Address, arg4 *big.Int) (*types.Transaction, error) {
	return _TraderJoeMock.Contract.SwapExactTokensForTokens(&_TraderJoeMock.TransactOpts, amountIn_, amountOutMin_, path_, receiver_, arg4)
}

// SwapTokensForExactAVAX is a paid mutator transaction binding the contract method 0x7a42416a.
//
// Solidity: function swapTokensForExactAVAX(uint256 amountOut_, uint256 amountInMax_, address[] path_, address receiver_, uint256 ) returns(uint256[] amounts_)
func (_TraderJoeMock *TraderJoeMockTransactor) SwapTokensForExactAVAX(opts *bind.TransactOpts, amountOut_ *big.Int, amountInMax_ *big.Int, path_ []common.Address, receiver_ common.Address, arg4 *big.Int) (*types.Transaction, error) {
	return _TraderJoeMock.contract.Transact(opts, "swapTokensForExactAVAX", amountOut_, amountInMax_, path_, receiver_, arg4)
}

// SwapTokensForExactAVAX is a paid mutator transaction binding the contract method 0x7a42416a.
//
// Solidity: function swapTokensForExactAVAX(uint256 amountOut_, uint256 amountInMax_, address[] path_, address receiver_, uint256 ) returns(uint256[] amounts_)
func (_TraderJoeMock *TraderJoeMockSession) SwapTokensForExactAVAX(amountOut_ *big.Int, amountInMax_ *big.Int, path_ []common.Address, receiver_ common.Address, arg4 *big.Int) (*types.Transaction, error) {
	return _TraderJoeMock.Contract.SwapTokensForExactAVAX(&_TraderJoeMock.TransactOpts, amountOut_, amountInMax_, path_, receiver_, arg4)
}

// SwapTokensForExactAVAX is a paid mutator transaction binding the contract method 0x7a42416a.
//
// Solidity: function swapTokensForExactAVAX(uint256 amountOut_, uint256 amountInMax_, address[] path_, address receiver_, uint256 ) returns(uint256[] amounts_)
func (_TraderJoeMock *TraderJoeMockTransactorSession) SwapTokensForExactAVAX(amountOut_ *big.Int, amountInMax_ *big.Int, path_ []common.Address, receiver_ common.Address, arg4 *big.Int) (*types.Transaction, error) {
	return _TraderJoeMock.Contract.SwapTokensForExactAVAX(&_TraderJoeMock.TransactOpts, amountOut_, amountInMax_, path_, receiver_, arg4)
}

// SwapTokensForExactTokens is a paid mutator transaction binding the contract method 0x8803dbee.
//
// Solidity: function swapTokensForExactTokens(uint256 amountOut_, uint256 amountInMax_, address[] path_, address receiver_, uint256 ) returns(uint256[] amounts_)
func (_TraderJoeMock *TraderJoeMockTransactor) SwapTokensForExactTokens(opts *bind.TransactOpts, amountOut_ *big.Int, amountInMax_ *big.Int, path_ []common.Address, receiver_ common.Address, arg4 *big.Int) (*types.Transaction, error) {
	return _TraderJoeMock.contract.Transact(opts, "swapTokensForExactTokens", amountOut_, amountInMax_, path_, receiver_, arg4)
}

// SwapTokensForExactTokens is a paid mutator transaction binding the contract method 0x8803dbee.
//
// Solidity: function swapTokensForExactTokens(uint256 amountOut_, uint256 amountInMax_, address[] path_, address receiver_, uint256 ) returns(uint256[] amounts_)
func (_TraderJoeMock *TraderJoeMockSession) SwapTokensForExactTokens(amountOut_ *big.Int, amountInMax_ *big.Int, path_ []common.Address, receiver_ common.Address, arg4 *big.Int) (*types.Transaction, error) {
	return _TraderJoeMock.Contract.SwapTokensForExactTokens(&_TraderJoeMock.TransactOpts, amountOut_, amountInMax_, path_, receiver_, arg4)
}

// SwapTokensForExactTokens is a paid mutator transaction binding the contract method 0x8803dbee.
//
// Solidity: function swapTokensForExactTokens(uint256 amountOut_, uint256 amountInMax_, address[] path_, address receiver_, uint256 ) returns(uint256[] amounts_)
func (_TraderJoeMock *TraderJoeMockTransactorSession) SwapTokensForExactTokens(amountOut_ *big.Int, amountInMax_ *big.Int, path_ []common.Address, receiver_ common.Address, arg4 *big.Int) (*types.Transaction, error) {
	return _TraderJoeMock.Contract.SwapTokensForExactTokens(&_TraderJoeMock.TransactOpts, amountOut_, amountInMax_, path_, receiver_, arg4)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TraderJoeMock *TraderJoeMockTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TraderJoeMock.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TraderJoeMock *TraderJoeMockSession) Receive() (*types.Transaction, error) {
	return _TraderJoeMock.Contract.Receive(&_TraderJoeMock.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TraderJoeMock *TraderJoeMockTransactorSession) Receive() (*types.Transaction, error) {
	return _TraderJoeMock.Contract.Receive(&_TraderJoeMock.TransactOpts)
}
