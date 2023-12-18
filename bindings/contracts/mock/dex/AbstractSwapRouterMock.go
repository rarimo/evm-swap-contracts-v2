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

// AbstractSwapRouterMockMetaData contains all meta data concerning the AbstractSwapRouterMock contract.
var AbstractSwapRouterMockMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"WRAPPED_NATIVE\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut_\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path_\",\"type\":\"address[]\"}],\"name\":\"_getAmountsIn\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts_\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn_\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path_\",\"type\":\"address[]\"}],\"name\":\"_getAmountsOut\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts_\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB_\",\"type\":\"address\"}],\"name\":\"enablePair\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"}],\"name\":\"setReserve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// AbstractSwapRouterMockABI is the input ABI used to generate the binding from.
// Deprecated: Use AbstractSwapRouterMockMetaData.ABI instead.
var AbstractSwapRouterMockABI = AbstractSwapRouterMockMetaData.ABI

// AbstractSwapRouterMock is an auto generated Go binding around an Ethereum contract.
type AbstractSwapRouterMock struct {
	AbstractSwapRouterMockCaller     // Read-only binding to the contract
	AbstractSwapRouterMockTransactor // Write-only binding to the contract
	AbstractSwapRouterMockFilterer   // Log filterer for contract events
}

// AbstractSwapRouterMockCaller is an auto generated read-only Go binding around an Ethereum contract.
type AbstractSwapRouterMockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbstractSwapRouterMockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AbstractSwapRouterMockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbstractSwapRouterMockFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AbstractSwapRouterMockFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AbstractSwapRouterMockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AbstractSwapRouterMockSession struct {
	Contract     *AbstractSwapRouterMock // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// AbstractSwapRouterMockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AbstractSwapRouterMockCallerSession struct {
	Contract *AbstractSwapRouterMockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// AbstractSwapRouterMockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AbstractSwapRouterMockTransactorSession struct {
	Contract     *AbstractSwapRouterMockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// AbstractSwapRouterMockRaw is an auto generated low-level Go binding around an Ethereum contract.
type AbstractSwapRouterMockRaw struct {
	Contract *AbstractSwapRouterMock // Generic contract binding to access the raw methods on
}

// AbstractSwapRouterMockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AbstractSwapRouterMockCallerRaw struct {
	Contract *AbstractSwapRouterMockCaller // Generic read-only contract binding to access the raw methods on
}

// AbstractSwapRouterMockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AbstractSwapRouterMockTransactorRaw struct {
	Contract *AbstractSwapRouterMockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAbstractSwapRouterMock creates a new instance of AbstractSwapRouterMock, bound to a specific deployed contract.
func NewAbstractSwapRouterMock(address common.Address, backend bind.ContractBackend) (*AbstractSwapRouterMock, error) {
	contract, err := bindAbstractSwapRouterMock(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AbstractSwapRouterMock{AbstractSwapRouterMockCaller: AbstractSwapRouterMockCaller{contract: contract}, AbstractSwapRouterMockTransactor: AbstractSwapRouterMockTransactor{contract: contract}, AbstractSwapRouterMockFilterer: AbstractSwapRouterMockFilterer{contract: contract}}, nil
}

// NewAbstractSwapRouterMockCaller creates a new read-only instance of AbstractSwapRouterMock, bound to a specific deployed contract.
func NewAbstractSwapRouterMockCaller(address common.Address, caller bind.ContractCaller) (*AbstractSwapRouterMockCaller, error) {
	contract, err := bindAbstractSwapRouterMock(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AbstractSwapRouterMockCaller{contract: contract}, nil
}

// NewAbstractSwapRouterMockTransactor creates a new write-only instance of AbstractSwapRouterMock, bound to a specific deployed contract.
func NewAbstractSwapRouterMockTransactor(address common.Address, transactor bind.ContractTransactor) (*AbstractSwapRouterMockTransactor, error) {
	contract, err := bindAbstractSwapRouterMock(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AbstractSwapRouterMockTransactor{contract: contract}, nil
}

// NewAbstractSwapRouterMockFilterer creates a new log filterer instance of AbstractSwapRouterMock, bound to a specific deployed contract.
func NewAbstractSwapRouterMockFilterer(address common.Address, filterer bind.ContractFilterer) (*AbstractSwapRouterMockFilterer, error) {
	contract, err := bindAbstractSwapRouterMock(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AbstractSwapRouterMockFilterer{contract: contract}, nil
}

// bindAbstractSwapRouterMock binds a generic wrapper to an already deployed contract.
func bindAbstractSwapRouterMock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AbstractSwapRouterMockMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AbstractSwapRouterMock *AbstractSwapRouterMockRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AbstractSwapRouterMock.Contract.AbstractSwapRouterMockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AbstractSwapRouterMock *AbstractSwapRouterMockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AbstractSwapRouterMock.Contract.AbstractSwapRouterMockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AbstractSwapRouterMock *AbstractSwapRouterMockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AbstractSwapRouterMock.Contract.AbstractSwapRouterMockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AbstractSwapRouterMock *AbstractSwapRouterMockCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AbstractSwapRouterMock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AbstractSwapRouterMock *AbstractSwapRouterMockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AbstractSwapRouterMock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AbstractSwapRouterMock *AbstractSwapRouterMockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AbstractSwapRouterMock.Contract.contract.Transact(opts, method, params...)
}

// WRAPPEDNATIVE is a free data retrieval call binding the contract method 0xd999984d.
//
// Solidity: function WRAPPED_NATIVE() view returns(address)
func (_AbstractSwapRouterMock *AbstractSwapRouterMockCaller) WRAPPEDNATIVE(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AbstractSwapRouterMock.contract.Call(opts, &out, "WRAPPED_NATIVE")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WRAPPEDNATIVE is a free data retrieval call binding the contract method 0xd999984d.
//
// Solidity: function WRAPPED_NATIVE() view returns(address)
func (_AbstractSwapRouterMock *AbstractSwapRouterMockSession) WRAPPEDNATIVE() (common.Address, error) {
	return _AbstractSwapRouterMock.Contract.WRAPPEDNATIVE(&_AbstractSwapRouterMock.CallOpts)
}

// WRAPPEDNATIVE is a free data retrieval call binding the contract method 0xd999984d.
//
// Solidity: function WRAPPED_NATIVE() view returns(address)
func (_AbstractSwapRouterMock *AbstractSwapRouterMockCallerSession) WRAPPEDNATIVE() (common.Address, error) {
	return _AbstractSwapRouterMock.Contract.WRAPPEDNATIVE(&_AbstractSwapRouterMock.CallOpts)
}

// GetAmountsIn is a free data retrieval call binding the contract method 0xbbf9a05e.
//
// Solidity: function _getAmountsIn(uint256 amountOut_, address[] path_) view returns(uint256[] amounts_)
func (_AbstractSwapRouterMock *AbstractSwapRouterMockCaller) GetAmountsIn(opts *bind.CallOpts, amountOut_ *big.Int, path_ []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _AbstractSwapRouterMock.contract.Call(opts, &out, "_getAmountsIn", amountOut_, path_)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetAmountsIn is a free data retrieval call binding the contract method 0xbbf9a05e.
//
// Solidity: function _getAmountsIn(uint256 amountOut_, address[] path_) view returns(uint256[] amounts_)
func (_AbstractSwapRouterMock *AbstractSwapRouterMockSession) GetAmountsIn(amountOut_ *big.Int, path_ []common.Address) ([]*big.Int, error) {
	return _AbstractSwapRouterMock.Contract.GetAmountsIn(&_AbstractSwapRouterMock.CallOpts, amountOut_, path_)
}

// GetAmountsIn is a free data retrieval call binding the contract method 0xbbf9a05e.
//
// Solidity: function _getAmountsIn(uint256 amountOut_, address[] path_) view returns(uint256[] amounts_)
func (_AbstractSwapRouterMock *AbstractSwapRouterMockCallerSession) GetAmountsIn(amountOut_ *big.Int, path_ []common.Address) ([]*big.Int, error) {
	return _AbstractSwapRouterMock.Contract.GetAmountsIn(&_AbstractSwapRouterMock.CallOpts, amountOut_, path_)
}

// GetAmountsOut is a free data retrieval call binding the contract method 0x8a410396.
//
// Solidity: function _getAmountsOut(uint256 amountIn_, address[] path_) view returns(uint256[] amounts_)
func (_AbstractSwapRouterMock *AbstractSwapRouterMockCaller) GetAmountsOut(opts *bind.CallOpts, amountIn_ *big.Int, path_ []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _AbstractSwapRouterMock.contract.Call(opts, &out, "_getAmountsOut", amountIn_, path_)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetAmountsOut is a free data retrieval call binding the contract method 0x8a410396.
//
// Solidity: function _getAmountsOut(uint256 amountIn_, address[] path_) view returns(uint256[] amounts_)
func (_AbstractSwapRouterMock *AbstractSwapRouterMockSession) GetAmountsOut(amountIn_ *big.Int, path_ []common.Address) ([]*big.Int, error) {
	return _AbstractSwapRouterMock.Contract.GetAmountsOut(&_AbstractSwapRouterMock.CallOpts, amountIn_, path_)
}

// GetAmountsOut is a free data retrieval call binding the contract method 0x8a410396.
//
// Solidity: function _getAmountsOut(uint256 amountIn_, address[] path_) view returns(uint256[] amounts_)
func (_AbstractSwapRouterMock *AbstractSwapRouterMockCallerSession) GetAmountsOut(amountIn_ *big.Int, path_ []common.Address) ([]*big.Int, error) {
	return _AbstractSwapRouterMock.Contract.GetAmountsOut(&_AbstractSwapRouterMock.CallOpts, amountIn_, path_)
}

// EnablePair is a paid mutator transaction binding the contract method 0xdfb5061c.
//
// Solidity: function enablePair(address tokenA_, address tokenB_) returns()
func (_AbstractSwapRouterMock *AbstractSwapRouterMockTransactor) EnablePair(opts *bind.TransactOpts, tokenA_ common.Address, tokenB_ common.Address) (*types.Transaction, error) {
	return _AbstractSwapRouterMock.contract.Transact(opts, "enablePair", tokenA_, tokenB_)
}

// EnablePair is a paid mutator transaction binding the contract method 0xdfb5061c.
//
// Solidity: function enablePair(address tokenA_, address tokenB_) returns()
func (_AbstractSwapRouterMock *AbstractSwapRouterMockSession) EnablePair(tokenA_ common.Address, tokenB_ common.Address) (*types.Transaction, error) {
	return _AbstractSwapRouterMock.Contract.EnablePair(&_AbstractSwapRouterMock.TransactOpts, tokenA_, tokenB_)
}

// EnablePair is a paid mutator transaction binding the contract method 0xdfb5061c.
//
// Solidity: function enablePair(address tokenA_, address tokenB_) returns()
func (_AbstractSwapRouterMock *AbstractSwapRouterMockTransactorSession) EnablePair(tokenA_ common.Address, tokenB_ common.Address) (*types.Transaction, error) {
	return _AbstractSwapRouterMock.Contract.EnablePair(&_AbstractSwapRouterMock.TransactOpts, tokenA_, tokenB_)
}

// SetReserve is a paid mutator transaction binding the contract method 0x78250b63.
//
// Solidity: function setReserve(address token_, uint256 amount_) returns()
func (_AbstractSwapRouterMock *AbstractSwapRouterMockTransactor) SetReserve(opts *bind.TransactOpts, token_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _AbstractSwapRouterMock.contract.Transact(opts, "setReserve", token_, amount_)
}

// SetReserve is a paid mutator transaction binding the contract method 0x78250b63.
//
// Solidity: function setReserve(address token_, uint256 amount_) returns()
func (_AbstractSwapRouterMock *AbstractSwapRouterMockSession) SetReserve(token_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _AbstractSwapRouterMock.Contract.SetReserve(&_AbstractSwapRouterMock.TransactOpts, token_, amount_)
}

// SetReserve is a paid mutator transaction binding the contract method 0x78250b63.
//
// Solidity: function setReserve(address token_, uint256 amount_) returns()
func (_AbstractSwapRouterMock *AbstractSwapRouterMockTransactorSession) SetReserve(token_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _AbstractSwapRouterMock.Contract.SetReserve(&_AbstractSwapRouterMock.TransactOpts, token_, amount_)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_AbstractSwapRouterMock *AbstractSwapRouterMockTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AbstractSwapRouterMock.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_AbstractSwapRouterMock *AbstractSwapRouterMockSession) Receive() (*types.Transaction, error) {
	return _AbstractSwapRouterMock.Contract.Receive(&_AbstractSwapRouterMock.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_AbstractSwapRouterMock *AbstractSwapRouterMockTransactorSession) Receive() (*types.Transaction, error) {
	return _AbstractSwapRouterMock.Contract.Receive(&_AbstractSwapRouterMock.TransactOpts)
}
