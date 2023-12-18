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

// ISwapRouterExactInputParams is an auto generated low-level Go binding around an user-defined struct.
type ISwapRouterExactInputParams struct {
	Path             []byte
	Recipient        common.Address
	Deadline         *big.Int
	AmountIn         *big.Int
	AmountOutMinimum *big.Int
}

// ISwapRouterExactOutputParams is an auto generated low-level Go binding around an user-defined struct.
type ISwapRouterExactOutputParams struct {
	Path            []byte
	Recipient       common.Address
	Deadline        *big.Int
	AmountOut       *big.Int
	AmountInMaximum *big.Int
}

// UniswapV3RouterMockMetaData contains all meta data concerning the UniswapV3RouterMock contract.
var UniswapV3RouterMockMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"weth9_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"WRAPPED_NATIVE\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut_\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path_\",\"type\":\"address[]\"}],\"name\":\"_getAmountsIn\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts_\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn_\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path_\",\"type\":\"address[]\"}],\"name\":\"_getAmountsOut\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amounts_\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenA_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenB_\",\"type\":\"address\"}],\"name\":\"enablePair\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMinimum\",\"type\":\"uint256\"}],\"internalType\":\"structISwapRouter.ExactInputParams\",\"name\":\"params_\",\"type\":\"tuple\"}],\"name\":\"exactInput\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut_\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"path\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMaximum\",\"type\":\"uint256\"}],\"internalType\":\"structISwapRouter.ExactOutputParams\",\"name\":\"params_\",\"type\":\"tuple\"}],\"name\":\"exactOutput\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn_\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"refundETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"}],\"name\":\"setReserve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// UniswapV3RouterMockABI is the input ABI used to generate the binding from.
// Deprecated: Use UniswapV3RouterMockMetaData.ABI instead.
var UniswapV3RouterMockABI = UniswapV3RouterMockMetaData.ABI

// UniswapV3RouterMock is an auto generated Go binding around an Ethereum contract.
type UniswapV3RouterMock struct {
	UniswapV3RouterMockCaller     // Read-only binding to the contract
	UniswapV3RouterMockTransactor // Write-only binding to the contract
	UniswapV3RouterMockFilterer   // Log filterer for contract events
}

// UniswapV3RouterMockCaller is an auto generated read-only Go binding around an Ethereum contract.
type UniswapV3RouterMockCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapV3RouterMockTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UniswapV3RouterMockTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapV3RouterMockFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UniswapV3RouterMockFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapV3RouterMockSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UniswapV3RouterMockSession struct {
	Contract     *UniswapV3RouterMock // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// UniswapV3RouterMockCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UniswapV3RouterMockCallerSession struct {
	Contract *UniswapV3RouterMockCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// UniswapV3RouterMockTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UniswapV3RouterMockTransactorSession struct {
	Contract     *UniswapV3RouterMockTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// UniswapV3RouterMockRaw is an auto generated low-level Go binding around an Ethereum contract.
type UniswapV3RouterMockRaw struct {
	Contract *UniswapV3RouterMock // Generic contract binding to access the raw methods on
}

// UniswapV3RouterMockCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UniswapV3RouterMockCallerRaw struct {
	Contract *UniswapV3RouterMockCaller // Generic read-only contract binding to access the raw methods on
}

// UniswapV3RouterMockTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UniswapV3RouterMockTransactorRaw struct {
	Contract *UniswapV3RouterMockTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUniswapV3RouterMock creates a new instance of UniswapV3RouterMock, bound to a specific deployed contract.
func NewUniswapV3RouterMock(address common.Address, backend bind.ContractBackend) (*UniswapV3RouterMock, error) {
	contract, err := bindUniswapV3RouterMock(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UniswapV3RouterMock{UniswapV3RouterMockCaller: UniswapV3RouterMockCaller{contract: contract}, UniswapV3RouterMockTransactor: UniswapV3RouterMockTransactor{contract: contract}, UniswapV3RouterMockFilterer: UniswapV3RouterMockFilterer{contract: contract}}, nil
}

// NewUniswapV3RouterMockCaller creates a new read-only instance of UniswapV3RouterMock, bound to a specific deployed contract.
func NewUniswapV3RouterMockCaller(address common.Address, caller bind.ContractCaller) (*UniswapV3RouterMockCaller, error) {
	contract, err := bindUniswapV3RouterMock(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UniswapV3RouterMockCaller{contract: contract}, nil
}

// NewUniswapV3RouterMockTransactor creates a new write-only instance of UniswapV3RouterMock, bound to a specific deployed contract.
func NewUniswapV3RouterMockTransactor(address common.Address, transactor bind.ContractTransactor) (*UniswapV3RouterMockTransactor, error) {
	contract, err := bindUniswapV3RouterMock(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UniswapV3RouterMockTransactor{contract: contract}, nil
}

// NewUniswapV3RouterMockFilterer creates a new log filterer instance of UniswapV3RouterMock, bound to a specific deployed contract.
func NewUniswapV3RouterMockFilterer(address common.Address, filterer bind.ContractFilterer) (*UniswapV3RouterMockFilterer, error) {
	contract, err := bindUniswapV3RouterMock(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UniswapV3RouterMockFilterer{contract: contract}, nil
}

// bindUniswapV3RouterMock binds a generic wrapper to an already deployed contract.
func bindUniswapV3RouterMock(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := UniswapV3RouterMockMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniswapV3RouterMock *UniswapV3RouterMockRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniswapV3RouterMock.Contract.UniswapV3RouterMockCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniswapV3RouterMock *UniswapV3RouterMockRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapV3RouterMock.Contract.UniswapV3RouterMockTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniswapV3RouterMock *UniswapV3RouterMockRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniswapV3RouterMock.Contract.UniswapV3RouterMockTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniswapV3RouterMock *UniswapV3RouterMockCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniswapV3RouterMock.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniswapV3RouterMock *UniswapV3RouterMockTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapV3RouterMock.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniswapV3RouterMock *UniswapV3RouterMockTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniswapV3RouterMock.Contract.contract.Transact(opts, method, params...)
}

// WRAPPEDNATIVE is a free data retrieval call binding the contract method 0xd999984d.
//
// Solidity: function WRAPPED_NATIVE() view returns(address)
func (_UniswapV3RouterMock *UniswapV3RouterMockCaller) WRAPPEDNATIVE(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UniswapV3RouterMock.contract.Call(opts, &out, "WRAPPED_NATIVE")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WRAPPEDNATIVE is a free data retrieval call binding the contract method 0xd999984d.
//
// Solidity: function WRAPPED_NATIVE() view returns(address)
func (_UniswapV3RouterMock *UniswapV3RouterMockSession) WRAPPEDNATIVE() (common.Address, error) {
	return _UniswapV3RouterMock.Contract.WRAPPEDNATIVE(&_UniswapV3RouterMock.CallOpts)
}

// WRAPPEDNATIVE is a free data retrieval call binding the contract method 0xd999984d.
//
// Solidity: function WRAPPED_NATIVE() view returns(address)
func (_UniswapV3RouterMock *UniswapV3RouterMockCallerSession) WRAPPEDNATIVE() (common.Address, error) {
	return _UniswapV3RouterMock.Contract.WRAPPEDNATIVE(&_UniswapV3RouterMock.CallOpts)
}

// GetAmountsIn is a free data retrieval call binding the contract method 0xbbf9a05e.
//
// Solidity: function _getAmountsIn(uint256 amountOut_, address[] path_) view returns(uint256[] amounts_)
func (_UniswapV3RouterMock *UniswapV3RouterMockCaller) GetAmountsIn(opts *bind.CallOpts, amountOut_ *big.Int, path_ []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _UniswapV3RouterMock.contract.Call(opts, &out, "_getAmountsIn", amountOut_, path_)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetAmountsIn is a free data retrieval call binding the contract method 0xbbf9a05e.
//
// Solidity: function _getAmountsIn(uint256 amountOut_, address[] path_) view returns(uint256[] amounts_)
func (_UniswapV3RouterMock *UniswapV3RouterMockSession) GetAmountsIn(amountOut_ *big.Int, path_ []common.Address) ([]*big.Int, error) {
	return _UniswapV3RouterMock.Contract.GetAmountsIn(&_UniswapV3RouterMock.CallOpts, amountOut_, path_)
}

// GetAmountsIn is a free data retrieval call binding the contract method 0xbbf9a05e.
//
// Solidity: function _getAmountsIn(uint256 amountOut_, address[] path_) view returns(uint256[] amounts_)
func (_UniswapV3RouterMock *UniswapV3RouterMockCallerSession) GetAmountsIn(amountOut_ *big.Int, path_ []common.Address) ([]*big.Int, error) {
	return _UniswapV3RouterMock.Contract.GetAmountsIn(&_UniswapV3RouterMock.CallOpts, amountOut_, path_)
}

// GetAmountsOut is a free data retrieval call binding the contract method 0x8a410396.
//
// Solidity: function _getAmountsOut(uint256 amountIn_, address[] path_) view returns(uint256[] amounts_)
func (_UniswapV3RouterMock *UniswapV3RouterMockCaller) GetAmountsOut(opts *bind.CallOpts, amountIn_ *big.Int, path_ []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _UniswapV3RouterMock.contract.Call(opts, &out, "_getAmountsOut", amountIn_, path_)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetAmountsOut is a free data retrieval call binding the contract method 0x8a410396.
//
// Solidity: function _getAmountsOut(uint256 amountIn_, address[] path_) view returns(uint256[] amounts_)
func (_UniswapV3RouterMock *UniswapV3RouterMockSession) GetAmountsOut(amountIn_ *big.Int, path_ []common.Address) ([]*big.Int, error) {
	return _UniswapV3RouterMock.Contract.GetAmountsOut(&_UniswapV3RouterMock.CallOpts, amountIn_, path_)
}

// GetAmountsOut is a free data retrieval call binding the contract method 0x8a410396.
//
// Solidity: function _getAmountsOut(uint256 amountIn_, address[] path_) view returns(uint256[] amounts_)
func (_UniswapV3RouterMock *UniswapV3RouterMockCallerSession) GetAmountsOut(amountIn_ *big.Int, path_ []common.Address) ([]*big.Int, error) {
	return _UniswapV3RouterMock.Contract.GetAmountsOut(&_UniswapV3RouterMock.CallOpts, amountIn_, path_)
}

// EnablePair is a paid mutator transaction binding the contract method 0xdfb5061c.
//
// Solidity: function enablePair(address tokenA_, address tokenB_) returns()
func (_UniswapV3RouterMock *UniswapV3RouterMockTransactor) EnablePair(opts *bind.TransactOpts, tokenA_ common.Address, tokenB_ common.Address) (*types.Transaction, error) {
	return _UniswapV3RouterMock.contract.Transact(opts, "enablePair", tokenA_, tokenB_)
}

// EnablePair is a paid mutator transaction binding the contract method 0xdfb5061c.
//
// Solidity: function enablePair(address tokenA_, address tokenB_) returns()
func (_UniswapV3RouterMock *UniswapV3RouterMockSession) EnablePair(tokenA_ common.Address, tokenB_ common.Address) (*types.Transaction, error) {
	return _UniswapV3RouterMock.Contract.EnablePair(&_UniswapV3RouterMock.TransactOpts, tokenA_, tokenB_)
}

// EnablePair is a paid mutator transaction binding the contract method 0xdfb5061c.
//
// Solidity: function enablePair(address tokenA_, address tokenB_) returns()
func (_UniswapV3RouterMock *UniswapV3RouterMockTransactorSession) EnablePair(tokenA_ common.Address, tokenB_ common.Address) (*types.Transaction, error) {
	return _UniswapV3RouterMock.Contract.EnablePair(&_UniswapV3RouterMock.TransactOpts, tokenA_, tokenB_)
}

// ExactInput is a paid mutator transaction binding the contract method 0xc04b8d59.
//
// Solidity: function exactInput((bytes,address,uint256,uint256,uint256) params_) payable returns(uint256 amountOut_)
func (_UniswapV3RouterMock *UniswapV3RouterMockTransactor) ExactInput(opts *bind.TransactOpts, params_ ISwapRouterExactInputParams) (*types.Transaction, error) {
	return _UniswapV3RouterMock.contract.Transact(opts, "exactInput", params_)
}

// ExactInput is a paid mutator transaction binding the contract method 0xc04b8d59.
//
// Solidity: function exactInput((bytes,address,uint256,uint256,uint256) params_) payable returns(uint256 amountOut_)
func (_UniswapV3RouterMock *UniswapV3RouterMockSession) ExactInput(params_ ISwapRouterExactInputParams) (*types.Transaction, error) {
	return _UniswapV3RouterMock.Contract.ExactInput(&_UniswapV3RouterMock.TransactOpts, params_)
}

// ExactInput is a paid mutator transaction binding the contract method 0xc04b8d59.
//
// Solidity: function exactInput((bytes,address,uint256,uint256,uint256) params_) payable returns(uint256 amountOut_)
func (_UniswapV3RouterMock *UniswapV3RouterMockTransactorSession) ExactInput(params_ ISwapRouterExactInputParams) (*types.Transaction, error) {
	return _UniswapV3RouterMock.Contract.ExactInput(&_UniswapV3RouterMock.TransactOpts, params_)
}

// ExactOutput is a paid mutator transaction binding the contract method 0xf28c0498.
//
// Solidity: function exactOutput((bytes,address,uint256,uint256,uint256) params_) payable returns(uint256 amountIn_)
func (_UniswapV3RouterMock *UniswapV3RouterMockTransactor) ExactOutput(opts *bind.TransactOpts, params_ ISwapRouterExactOutputParams) (*types.Transaction, error) {
	return _UniswapV3RouterMock.contract.Transact(opts, "exactOutput", params_)
}

// ExactOutput is a paid mutator transaction binding the contract method 0xf28c0498.
//
// Solidity: function exactOutput((bytes,address,uint256,uint256,uint256) params_) payable returns(uint256 amountIn_)
func (_UniswapV3RouterMock *UniswapV3RouterMockSession) ExactOutput(params_ ISwapRouterExactOutputParams) (*types.Transaction, error) {
	return _UniswapV3RouterMock.Contract.ExactOutput(&_UniswapV3RouterMock.TransactOpts, params_)
}

// ExactOutput is a paid mutator transaction binding the contract method 0xf28c0498.
//
// Solidity: function exactOutput((bytes,address,uint256,uint256,uint256) params_) payable returns(uint256 amountIn_)
func (_UniswapV3RouterMock *UniswapV3RouterMockTransactorSession) ExactOutput(params_ ISwapRouterExactOutputParams) (*types.Transaction, error) {
	return _UniswapV3RouterMock.Contract.ExactOutput(&_UniswapV3RouterMock.TransactOpts, params_)
}

// RefundETH is a paid mutator transaction binding the contract method 0x12210e8a.
//
// Solidity: function refundETH() payable returns()
func (_UniswapV3RouterMock *UniswapV3RouterMockTransactor) RefundETH(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapV3RouterMock.contract.Transact(opts, "refundETH")
}

// RefundETH is a paid mutator transaction binding the contract method 0x12210e8a.
//
// Solidity: function refundETH() payable returns()
func (_UniswapV3RouterMock *UniswapV3RouterMockSession) RefundETH() (*types.Transaction, error) {
	return _UniswapV3RouterMock.Contract.RefundETH(&_UniswapV3RouterMock.TransactOpts)
}

// RefundETH is a paid mutator transaction binding the contract method 0x12210e8a.
//
// Solidity: function refundETH() payable returns()
func (_UniswapV3RouterMock *UniswapV3RouterMockTransactorSession) RefundETH() (*types.Transaction, error) {
	return _UniswapV3RouterMock.Contract.RefundETH(&_UniswapV3RouterMock.TransactOpts)
}

// SetReserve is a paid mutator transaction binding the contract method 0x78250b63.
//
// Solidity: function setReserve(address token_, uint256 amount_) returns()
func (_UniswapV3RouterMock *UniswapV3RouterMockTransactor) SetReserve(opts *bind.TransactOpts, token_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _UniswapV3RouterMock.contract.Transact(opts, "setReserve", token_, amount_)
}

// SetReserve is a paid mutator transaction binding the contract method 0x78250b63.
//
// Solidity: function setReserve(address token_, uint256 amount_) returns()
func (_UniswapV3RouterMock *UniswapV3RouterMockSession) SetReserve(token_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _UniswapV3RouterMock.Contract.SetReserve(&_UniswapV3RouterMock.TransactOpts, token_, amount_)
}

// SetReserve is a paid mutator transaction binding the contract method 0x78250b63.
//
// Solidity: function setReserve(address token_, uint256 amount_) returns()
func (_UniswapV3RouterMock *UniswapV3RouterMockTransactorSession) SetReserve(token_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _UniswapV3RouterMock.Contract.SetReserve(&_UniswapV3RouterMock.TransactOpts, token_, amount_)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_UniswapV3RouterMock *UniswapV3RouterMockTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapV3RouterMock.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_UniswapV3RouterMock *UniswapV3RouterMockSession) Receive() (*types.Transaction, error) {
	return _UniswapV3RouterMock.Contract.Receive(&_UniswapV3RouterMock.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_UniswapV3RouterMock *UniswapV3RouterMockTransactorSession) Receive() (*types.Transaction, error) {
	return _UniswapV3RouterMock.Contract.Receive(&_UniswapV3RouterMock.TransactOpts)
}
