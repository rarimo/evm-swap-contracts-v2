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

// UniswapV2RouterMockMetaData contains all meta data concerning the UniswapV2RouterMock contract.
var UniswapV2RouterMockMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"weth9_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"WRAPPED_NATIVE\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut_\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path_\",\"type\":\"address[]\"}],\"name\":\"_getAmountsIn\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts_\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn_\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path_\",\"type\":\"address[]\"}],\"name\":\"_getAmountsOut\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts_\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB_\",\"type\":\"address\"}],\"name\":\"enablePair\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"}],\"name\":\"setReserve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut_\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path_\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"receiver_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"swapETHForExactTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts_\",\"type\":\"uint256[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOutMin_\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path_\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"receiver_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"swapExactETHForTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts_\",\"type\":\"uint256[]\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin_\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path_\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"receiver_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForETH\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts_\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin_\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path_\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"receiver_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"swapExactTokensForTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts_\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMax_\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path_\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"receiver_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"swapTokensForExactETH\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts_\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMax_\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path_\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"receiver_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"swapTokensForExactTokens\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts_\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// UniswapV2RouterMockABI is the input ABI used to generate the binding from.
// Deprecated: Use UniswapV2RouterMockMetaData.ABI instead.
var UniswapV2RouterMockABI = UniswapV2RouterMockMetaData.ABI

// UniswapV2RouterMock is an auto generated Go binding around an Ethereum contract.
type UniswapV2RouterMock struct {
	UniswapV2RouterMockCaller     // Read-only binding to the contract
	UniswapV2RouterMockTransactor // Write-only binding to the contract
	UniswapV2RouterMockFilterer   // Log filterer for contract events
}

// UniswapV2RouterMockCaller is an auto generated read-only Go binding around an Ethereum contract.
type UniswapV2RouterMockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapV2RouterMockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UniswapV2RouterMockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapV2RouterMockFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UniswapV2RouterMockFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapV2RouterMockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UniswapV2RouterMockSession struct {
	Contract     *UniswapV2RouterMock // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// UniswapV2RouterMockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UniswapV2RouterMockCallerSession struct {
	Contract *UniswapV2RouterMockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// UniswapV2RouterMockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UniswapV2RouterMockTransactorSession struct {
	Contract     *UniswapV2RouterMockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// UniswapV2RouterMockRaw is an auto generated low-level Go binding around an Ethereum contract.
type UniswapV2RouterMockRaw struct {
	Contract *UniswapV2RouterMock // Generic contract binding to access the raw methods on
}

// UniswapV2RouterMockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UniswapV2RouterMockCallerRaw struct {
	Contract *UniswapV2RouterMockCaller // Generic read-only contract binding to access the raw methods on
}

// UniswapV2RouterMockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UniswapV2RouterMockTransactorRaw struct {
	Contract *UniswapV2RouterMockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUniswapV2RouterMock creates a new instance of UniswapV2RouterMock, bound to a specific deployed contract.
func NewUniswapV2RouterMock(address common.Address, backend bind.ContractBackend) (*UniswapV2RouterMock, error) {
	contract, err := bindUniswapV2RouterMock(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UniswapV2RouterMock{UniswapV2RouterMockCaller: UniswapV2RouterMockCaller{contract: contract}, UniswapV2RouterMockTransactor: UniswapV2RouterMockTransactor{contract: contract}, UniswapV2RouterMockFilterer: UniswapV2RouterMockFilterer{contract: contract}}, nil
}

// NewUniswapV2RouterMockCaller creates a new read-only instance of UniswapV2RouterMock, bound to a specific deployed contract.
func NewUniswapV2RouterMockCaller(address common.Address, caller bind.ContractCaller) (*UniswapV2RouterMockCaller, error) {
	contract, err := bindUniswapV2RouterMock(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UniswapV2RouterMockCaller{contract: contract}, nil
}

// NewUniswapV2RouterMockTransactor creates a new write-only instance of UniswapV2RouterMock, bound to a specific deployed contract.
func NewUniswapV2RouterMockTransactor(address common.Address, transactor bind.ContractTransactor) (*UniswapV2RouterMockTransactor, error) {
	contract, err := bindUniswapV2RouterMock(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UniswapV2RouterMockTransactor{contract: contract}, nil
}

// NewUniswapV2RouterMockFilterer creates a new log filterer instance of UniswapV2RouterMock, bound to a specific deployed contract.
func NewUniswapV2RouterMockFilterer(address common.Address, filterer bind.ContractFilterer) (*UniswapV2RouterMockFilterer, error) {
	contract, err := bindUniswapV2RouterMock(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UniswapV2RouterMockFilterer{contract: contract}, nil
}

// bindUniswapV2RouterMock binds a generic wrapper to an already deployed contract.
func bindUniswapV2RouterMock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := UniswapV2RouterMockMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniswapV2RouterMock *UniswapV2RouterMockRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniswapV2RouterMock.Contract.UniswapV2RouterMockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniswapV2RouterMock *UniswapV2RouterMockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapV2RouterMock.Contract.UniswapV2RouterMockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniswapV2RouterMock *UniswapV2RouterMockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniswapV2RouterMock.Contract.UniswapV2RouterMockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniswapV2RouterMock *UniswapV2RouterMockCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniswapV2RouterMock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniswapV2RouterMock *UniswapV2RouterMockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapV2RouterMock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniswapV2RouterMock *UniswapV2RouterMockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniswapV2RouterMock.Contract.contract.Transact(opts, method, params...)
}

// WRAPPEDNATIVE is a free data retrieval call binding the contract method 0xd999984d.
//
// Solidity: function WRAPPED_NATIVE() view returns(address)
func (_UniswapV2RouterMock *UniswapV2RouterMockCaller) WRAPPEDNATIVE(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UniswapV2RouterMock.contract.Call(opts, &out, "WRAPPED_NATIVE")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WRAPPEDNATIVE is a free data retrieval call binding the contract method 0xd999984d.
//
// Solidity: function WRAPPED_NATIVE() view returns(address)
func (_UniswapV2RouterMock *UniswapV2RouterMockSession) WRAPPEDNATIVE() (common.Address, error) {
	return _UniswapV2RouterMock.Contract.WRAPPEDNATIVE(&_UniswapV2RouterMock.CallOpts)
}

// WRAPPEDNATIVE is a free data retrieval call binding the contract method 0xd999984d.
//
// Solidity: function WRAPPED_NATIVE() view returns(address)
func (_UniswapV2RouterMock *UniswapV2RouterMockCallerSession) WRAPPEDNATIVE() (common.Address, error) {
	return _UniswapV2RouterMock.Contract.WRAPPEDNATIVE(&_UniswapV2RouterMock.CallOpts)
}

// GetAmountsIn is a free data retrieval call binding the contract method 0xbbf9a05e.
//
// Solidity: function _getAmountsIn(uint256 amountOut_, address[] path_) view returns(uint256[] amounts_)
func (_UniswapV2RouterMock *UniswapV2RouterMockCaller) GetAmountsIn(opts *bind.CallOpts, amountOut_ *big.Int, path_ []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _UniswapV2RouterMock.contract.Call(opts, &out, "_getAmountsIn", amountOut_, path_)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetAmountsIn is a free data retrieval call binding the contract method 0xbbf9a05e.
//
// Solidity: function _getAmountsIn(uint256 amountOut_, address[] path_) view returns(uint256[] amounts_)
func (_UniswapV2RouterMock *UniswapV2RouterMockSession) GetAmountsIn(amountOut_ *big.Int, path_ []common.Address) ([]*big.Int, error) {
	return _UniswapV2RouterMock.Contract.GetAmountsIn(&_UniswapV2RouterMock.CallOpts, amountOut_, path_)
}

// GetAmountsIn is a free data retrieval call binding the contract method 0xbbf9a05e.
//
// Solidity: function _getAmountsIn(uint256 amountOut_, address[] path_) view returns(uint256[] amounts_)
func (_UniswapV2RouterMock *UniswapV2RouterMockCallerSession) GetAmountsIn(amountOut_ *big.Int, path_ []common.Address) ([]*big.Int, error) {
	return _UniswapV2RouterMock.Contract.GetAmountsIn(&_UniswapV2RouterMock.CallOpts, amountOut_, path_)
}

// GetAmountsOut is a free data retrieval call binding the contract method 0x8a410396.
//
// Solidity: function _getAmountsOut(uint256 amountIn_, address[] path_) view returns(uint256[] amounts_)
func (_UniswapV2RouterMock *UniswapV2RouterMockCaller) GetAmountsOut(opts *bind.CallOpts, amountIn_ *big.Int, path_ []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _UniswapV2RouterMock.contract.Call(opts, &out, "_getAmountsOut", amountIn_, path_)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetAmountsOut is a free data retrieval call binding the contract method 0x8a410396.
//
// Solidity: function _getAmountsOut(uint256 amountIn_, address[] path_) view returns(uint256[] amounts_)
func (_UniswapV2RouterMock *UniswapV2RouterMockSession) GetAmountsOut(amountIn_ *big.Int, path_ []common.Address) ([]*big.Int, error) {
	return _UniswapV2RouterMock.Contract.GetAmountsOut(&_UniswapV2RouterMock.CallOpts, amountIn_, path_)
}

// GetAmountsOut is a free data retrieval call binding the contract method 0x8a410396.
//
// Solidity: function _getAmountsOut(uint256 amountIn_, address[] path_) view returns(uint256[] amounts_)
func (_UniswapV2RouterMock *UniswapV2RouterMockCallerSession) GetAmountsOut(amountIn_ *big.Int, path_ []common.Address) ([]*big.Int, error) {
	return _UniswapV2RouterMock.Contract.GetAmountsOut(&_UniswapV2RouterMock.CallOpts, amountIn_, path_)
}

// EnablePair is a paid mutator transaction binding the contract method 0xdfb5061c.
//
// Solidity: function enablePair(address tokenA_, address tokenB_) returns()
func (_UniswapV2RouterMock *UniswapV2RouterMockTransactor) EnablePair(opts *bind.TransactOpts, tokenA_ common.Address, tokenB_ common.Address) (*types.Transaction, error) {
	return _UniswapV2RouterMock.contract.Transact(opts, "enablePair", tokenA_, tokenB_)
}

// EnablePair is a paid mutator transaction binding the contract method 0xdfb5061c.
//
// Solidity: function enablePair(address tokenA_, address tokenB_) returns()
func (_UniswapV2RouterMock *UniswapV2RouterMockSession) EnablePair(tokenA_ common.Address, tokenB_ common.Address) (*types.Transaction, error) {
	return _UniswapV2RouterMock.Contract.EnablePair(&_UniswapV2RouterMock.TransactOpts, tokenA_, tokenB_)
}

// EnablePair is a paid mutator transaction binding the contract method 0xdfb5061c.
//
// Solidity: function enablePair(address tokenA_, address tokenB_) returns()
func (_UniswapV2RouterMock *UniswapV2RouterMockTransactorSession) EnablePair(tokenA_ common.Address, tokenB_ common.Address) (*types.Transaction, error) {
	return _UniswapV2RouterMock.Contract.EnablePair(&_UniswapV2RouterMock.TransactOpts, tokenA_, tokenB_)
}

// SetReserve is a paid mutator transaction binding the contract method 0x78250b63.
//
// Solidity: function setReserve(address token_, uint256 amount_) returns()
func (_UniswapV2RouterMock *UniswapV2RouterMockTransactor) SetReserve(opts *bind.TransactOpts, token_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _UniswapV2RouterMock.contract.Transact(opts, "setReserve", token_, amount_)
}

// SetReserve is a paid mutator transaction binding the contract method 0x78250b63.
//
// Solidity: function setReserve(address token_, uint256 amount_) returns()
func (_UniswapV2RouterMock *UniswapV2RouterMockSession) SetReserve(token_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _UniswapV2RouterMock.Contract.SetReserve(&_UniswapV2RouterMock.TransactOpts, token_, amount_)
}

// SetReserve is a paid mutator transaction binding the contract method 0x78250b63.
//
// Solidity: function setReserve(address token_, uint256 amount_) returns()
func (_UniswapV2RouterMock *UniswapV2RouterMockTransactorSession) SetReserve(token_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _UniswapV2RouterMock.Contract.SetReserve(&_UniswapV2RouterMock.TransactOpts, token_, amount_)
}

// SwapETHForExactTokens is a paid mutator transaction binding the contract method 0xfb3bdb41.
//
// Solidity: function swapETHForExactTokens(uint256 amountOut_, address[] path_, address receiver_, uint256 ) payable returns(uint256[] amounts_)
func (_UniswapV2RouterMock *UniswapV2RouterMockTransactor) SwapETHForExactTokens(opts *bind.TransactOpts, amountOut_ *big.Int, path_ []common.Address, receiver_ common.Address, arg3 *big.Int) (*types.Transaction, error) {
	return _UniswapV2RouterMock.contract.Transact(opts, "swapETHForExactTokens", amountOut_, path_, receiver_, arg3)
}

// SwapETHForExactTokens is a paid mutator transaction binding the contract method 0xfb3bdb41.
//
// Solidity: function swapETHForExactTokens(uint256 amountOut_, address[] path_, address receiver_, uint256 ) payable returns(uint256[] amounts_)
func (_UniswapV2RouterMock *UniswapV2RouterMockSession) SwapETHForExactTokens(amountOut_ *big.Int, path_ []common.Address, receiver_ common.Address, arg3 *big.Int) (*types.Transaction, error) {
	return _UniswapV2RouterMock.Contract.SwapETHForExactTokens(&_UniswapV2RouterMock.TransactOpts, amountOut_, path_, receiver_, arg3)
}

// SwapETHForExactTokens is a paid mutator transaction binding the contract method 0xfb3bdb41.
//
// Solidity: function swapETHForExactTokens(uint256 amountOut_, address[] path_, address receiver_, uint256 ) payable returns(uint256[] amounts_)
func (_UniswapV2RouterMock *UniswapV2RouterMockTransactorSession) SwapETHForExactTokens(amountOut_ *big.Int, path_ []common.Address, receiver_ common.Address, arg3 *big.Int) (*types.Transaction, error) {
	return _UniswapV2RouterMock.Contract.SwapETHForExactTokens(&_UniswapV2RouterMock.TransactOpts, amountOut_, path_, receiver_, arg3)
}

// SwapExactETHForTokens is a paid mutator transaction binding the contract method 0x7ff36ab5.
//
// Solidity: function swapExactETHForTokens(uint256 amountOutMin_, address[] path_, address receiver_, uint256 ) payable returns(uint256[] amounts_)
func (_UniswapV2RouterMock *UniswapV2RouterMockTransactor) SwapExactETHForTokens(opts *bind.TransactOpts, amountOutMin_ *big.Int, path_ []common.Address, receiver_ common.Address, arg3 *big.Int) (*types.Transaction, error) {
	return _UniswapV2RouterMock.contract.Transact(opts, "swapExactETHForTokens", amountOutMin_, path_, receiver_, arg3)
}

// SwapExactETHForTokens is a paid mutator transaction binding the contract method 0x7ff36ab5.
//
// Solidity: function swapExactETHForTokens(uint256 amountOutMin_, address[] path_, address receiver_, uint256 ) payable returns(uint256[] amounts_)
func (_UniswapV2RouterMock *UniswapV2RouterMockSession) SwapExactETHForTokens(amountOutMin_ *big.Int, path_ []common.Address, receiver_ common.Address, arg3 *big.Int) (*types.Transaction, error) {
	return _UniswapV2RouterMock.Contract.SwapExactETHForTokens(&_UniswapV2RouterMock.TransactOpts, amountOutMin_, path_, receiver_, arg3)
}

// SwapExactETHForTokens is a paid mutator transaction binding the contract method 0x7ff36ab5.
//
// Solidity: function swapExactETHForTokens(uint256 amountOutMin_, address[] path_, address receiver_, uint256 ) payable returns(uint256[] amounts_)
func (_UniswapV2RouterMock *UniswapV2RouterMockTransactorSession) SwapExactETHForTokens(amountOutMin_ *big.Int, path_ []common.Address, receiver_ common.Address, arg3 *big.Int) (*types.Transaction, error) {
	return _UniswapV2RouterMock.Contract.SwapExactETHForTokens(&_UniswapV2RouterMock.TransactOpts, amountOutMin_, path_, receiver_, arg3)
}

// SwapExactTokensForETH is a paid mutator transaction binding the contract method 0x18cbafe5.
//
// Solidity: function swapExactTokensForETH(uint256 amountIn_, uint256 amountOutMin_, address[] path_, address receiver_, uint256 ) returns(uint256[] amounts_)
func (_UniswapV2RouterMock *UniswapV2RouterMockTransactor) SwapExactTokensForETH(opts *bind.TransactOpts, amountIn_ *big.Int, amountOutMin_ *big.Int, path_ []common.Address, receiver_ common.Address, arg4 *big.Int) (*types.Transaction, error) {
	return _UniswapV2RouterMock.contract.Transact(opts, "swapExactTokensForETH", amountIn_, amountOutMin_, path_, receiver_, arg4)
}

// SwapExactTokensForETH is a paid mutator transaction binding the contract method 0x18cbafe5.
//
// Solidity: function swapExactTokensForETH(uint256 amountIn_, uint256 amountOutMin_, address[] path_, address receiver_, uint256 ) returns(uint256[] amounts_)
func (_UniswapV2RouterMock *UniswapV2RouterMockSession) SwapExactTokensForETH(amountIn_ *big.Int, amountOutMin_ *big.Int, path_ []common.Address, receiver_ common.Address, arg4 *big.Int) (*types.Transaction, error) {
	return _UniswapV2RouterMock.Contract.SwapExactTokensForETH(&_UniswapV2RouterMock.TransactOpts, amountIn_, amountOutMin_, path_, receiver_, arg4)
}

// SwapExactTokensForETH is a paid mutator transaction binding the contract method 0x18cbafe5.
//
// Solidity: function swapExactTokensForETH(uint256 amountIn_, uint256 amountOutMin_, address[] path_, address receiver_, uint256 ) returns(uint256[] amounts_)
func (_UniswapV2RouterMock *UniswapV2RouterMockTransactorSession) SwapExactTokensForETH(amountIn_ *big.Int, amountOutMin_ *big.Int, path_ []common.Address, receiver_ common.Address, arg4 *big.Int) (*types.Transaction, error) {
	return _UniswapV2RouterMock.Contract.SwapExactTokensForETH(&_UniswapV2RouterMock.TransactOpts, amountIn_, amountOutMin_, path_, receiver_, arg4)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x38ed1739.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn_, uint256 amountOutMin_, address[] path_, address receiver_, uint256 ) returns(uint256[] amounts_)
func (_UniswapV2RouterMock *UniswapV2RouterMockTransactor) SwapExactTokensForTokens(opts *bind.TransactOpts, amountIn_ *big.Int, amountOutMin_ *big.Int, path_ []common.Address, receiver_ common.Address, arg4 *big.Int) (*types.Transaction, error) {
	return _UniswapV2RouterMock.contract.Transact(opts, "swapExactTokensForTokens", amountIn_, amountOutMin_, path_, receiver_, arg4)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x38ed1739.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn_, uint256 amountOutMin_, address[] path_, address receiver_, uint256 ) returns(uint256[] amounts_)
func (_UniswapV2RouterMock *UniswapV2RouterMockSession) SwapExactTokensForTokens(amountIn_ *big.Int, amountOutMin_ *big.Int, path_ []common.Address, receiver_ common.Address, arg4 *big.Int) (*types.Transaction, error) {
	return _UniswapV2RouterMock.Contract.SwapExactTokensForTokens(&_UniswapV2RouterMock.TransactOpts, amountIn_, amountOutMin_, path_, receiver_, arg4)
}

// SwapExactTokensForTokens is a paid mutator transaction binding the contract method 0x38ed1739.
//
// Solidity: function swapExactTokensForTokens(uint256 amountIn_, uint256 amountOutMin_, address[] path_, address receiver_, uint256 ) returns(uint256[] amounts_)
func (_UniswapV2RouterMock *UniswapV2RouterMockTransactorSession) SwapExactTokensForTokens(amountIn_ *big.Int, amountOutMin_ *big.Int, path_ []common.Address, receiver_ common.Address, arg4 *big.Int) (*types.Transaction, error) {
	return _UniswapV2RouterMock.Contract.SwapExactTokensForTokens(&_UniswapV2RouterMock.TransactOpts, amountIn_, amountOutMin_, path_, receiver_, arg4)
}

// SwapTokensForExactETH is a paid mutator transaction binding the contract method 0x4a25d94a.
//
// Solidity: function swapTokensForExactETH(uint256 amountOut_, uint256 amountInMax_, address[] path_, address receiver_, uint256 ) returns(uint256[] amounts_)
func (_UniswapV2RouterMock *UniswapV2RouterMockTransactor) SwapTokensForExactETH(opts *bind.TransactOpts, amountOut_ *big.Int, amountInMax_ *big.Int, path_ []common.Address, receiver_ common.Address, arg4 *big.Int) (*types.Transaction, error) {
	return _UniswapV2RouterMock.contract.Transact(opts, "swapTokensForExactETH", amountOut_, amountInMax_, path_, receiver_, arg4)
}

// SwapTokensForExactETH is a paid mutator transaction binding the contract method 0x4a25d94a.
//
// Solidity: function swapTokensForExactETH(uint256 amountOut_, uint256 amountInMax_, address[] path_, address receiver_, uint256 ) returns(uint256[] amounts_)
func (_UniswapV2RouterMock *UniswapV2RouterMockSession) SwapTokensForExactETH(amountOut_ *big.Int, amountInMax_ *big.Int, path_ []common.Address, receiver_ common.Address, arg4 *big.Int) (*types.Transaction, error) {
	return _UniswapV2RouterMock.Contract.SwapTokensForExactETH(&_UniswapV2RouterMock.TransactOpts, amountOut_, amountInMax_, path_, receiver_, arg4)
}

// SwapTokensForExactETH is a paid mutator transaction binding the contract method 0x4a25d94a.
//
// Solidity: function swapTokensForExactETH(uint256 amountOut_, uint256 amountInMax_, address[] path_, address receiver_, uint256 ) returns(uint256[] amounts_)
func (_UniswapV2RouterMock *UniswapV2RouterMockTransactorSession) SwapTokensForExactETH(amountOut_ *big.Int, amountInMax_ *big.Int, path_ []common.Address, receiver_ common.Address, arg4 *big.Int) (*types.Transaction, error) {
	return _UniswapV2RouterMock.Contract.SwapTokensForExactETH(&_UniswapV2RouterMock.TransactOpts, amountOut_, amountInMax_, path_, receiver_, arg4)
}

// SwapTokensForExactTokens is a paid mutator transaction binding the contract method 0x8803dbee.
//
// Solidity: function swapTokensForExactTokens(uint256 amountOut_, uint256 amountInMax_, address[] path_, address receiver_, uint256 ) returns(uint256[] amounts_)
func (_UniswapV2RouterMock *UniswapV2RouterMockTransactor) SwapTokensForExactTokens(opts *bind.TransactOpts, amountOut_ *big.Int, amountInMax_ *big.Int, path_ []common.Address, receiver_ common.Address, arg4 *big.Int) (*types.Transaction, error) {
	return _UniswapV2RouterMock.contract.Transact(opts, "swapTokensForExactTokens", amountOut_, amountInMax_, path_, receiver_, arg4)
}

// SwapTokensForExactTokens is a paid mutator transaction binding the contract method 0x8803dbee.
//
// Solidity: function swapTokensForExactTokens(uint256 amountOut_, uint256 amountInMax_, address[] path_, address receiver_, uint256 ) returns(uint256[] amounts_)
func (_UniswapV2RouterMock *UniswapV2RouterMockSession) SwapTokensForExactTokens(amountOut_ *big.Int, amountInMax_ *big.Int, path_ []common.Address, receiver_ common.Address, arg4 *big.Int) (*types.Transaction, error) {
	return _UniswapV2RouterMock.Contract.SwapTokensForExactTokens(&_UniswapV2RouterMock.TransactOpts, amountOut_, amountInMax_, path_, receiver_, arg4)
}

// SwapTokensForExactTokens is a paid mutator transaction binding the contract method 0x8803dbee.
//
// Solidity: function swapTokensForExactTokens(uint256 amountOut_, uint256 amountInMax_, address[] path_, address receiver_, uint256 ) returns(uint256[] amounts_)
func (_UniswapV2RouterMock *UniswapV2RouterMockTransactorSession) SwapTokensForExactTokens(amountOut_ *big.Int, amountInMax_ *big.Int, path_ []common.Address, receiver_ common.Address, arg4 *big.Int) (*types.Transaction, error) {
	return _UniswapV2RouterMock.Contract.SwapTokensForExactTokens(&_UniswapV2RouterMock.TransactOpts, amountOut_, amountInMax_, path_, receiver_, arg4)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_UniswapV2RouterMock *UniswapV2RouterMockTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapV2RouterMock.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_UniswapV2RouterMock *UniswapV2RouterMockSession) Receive() (*types.Transaction, error) {
	return _UniswapV2RouterMock.Contract.Receive(&_UniswapV2RouterMock.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_UniswapV2RouterMock *UniswapV2RouterMockTransactorSession) Receive() (*types.Transaction, error) {
	return _UniswapV2RouterMock.Contract.Receive(&_UniswapV2RouterMock.TransactOpts)
}
