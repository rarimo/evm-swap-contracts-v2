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

// UniswapV2RouterMetaData contains all meta data concerning the UniswapV2Router contract.
var UniswapV2RouterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"OWNABLE_DIAMOND_STORAGE_SLOT\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UNISWAP_V2_ROUTER_STORAGE_SLOT\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSwapV2Router\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"swapV2Router_\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"swapV2Router_\",\"type\":\"address\"}],\"name\":\"setUniswapV2RouterAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountOut_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMax_\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path_\",\"type\":\"address[]\"}],\"name\":\"swapETHForExactTokens\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin_\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path_\",\"type\":\"address[]\"}],\"name\":\"swapExactETHForTokens\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin_\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path_\",\"type\":\"address[]\"}],\"name\":\"swapExactTokensForETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMin_\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path_\",\"type\":\"address[]\"}],\"name\":\"swapExactTokensForTokensV2\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountOut_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMax_\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path_\",\"type\":\"address[]\"}],\"name\":\"swapTokensForExactETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountOut_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMax_\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"path_\",\"type\":\"address[]\"}],\"name\":\"swapTokensForExactTokensV2\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// UniswapV2RouterABI is the input ABI used to generate the binding from.
// Deprecated: Use UniswapV2RouterMetaData.ABI instead.
var UniswapV2RouterABI = UniswapV2RouterMetaData.ABI

// UniswapV2Router is an auto generated Go binding around an Ethereum contract.
type UniswapV2Router struct {
	UniswapV2RouterCaller     // Read-only binding to the contract
	UniswapV2RouterTransactor // Write-only binding to the contract
	UniswapV2RouterFilterer   // Log filterer for contract events
}

// UniswapV2RouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type UniswapV2RouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapV2RouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UniswapV2RouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapV2RouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UniswapV2RouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapV2RouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UniswapV2RouterSession struct {
	Contract     *UniswapV2Router  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UniswapV2RouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UniswapV2RouterCallerSession struct {
	Contract *UniswapV2RouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// UniswapV2RouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UniswapV2RouterTransactorSession struct {
	Contract     *UniswapV2RouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// UniswapV2RouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type UniswapV2RouterRaw struct {
	Contract *UniswapV2Router // Generic contract binding to access the raw methods on
}

// UniswapV2RouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UniswapV2RouterCallerRaw struct {
	Contract *UniswapV2RouterCaller // Generic read-only contract binding to access the raw methods on
}

// UniswapV2RouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UniswapV2RouterTransactorRaw struct {
	Contract *UniswapV2RouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUniswapV2Router creates a new instance of UniswapV2Router, bound to a specific deployed contract.
func NewUniswapV2Router(address common.Address, backend bind.ContractBackend) (*UniswapV2Router, error) {
	contract, err := bindUniswapV2Router(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UniswapV2Router{UniswapV2RouterCaller: UniswapV2RouterCaller{contract: contract}, UniswapV2RouterTransactor: UniswapV2RouterTransactor{contract: contract}, UniswapV2RouterFilterer: UniswapV2RouterFilterer{contract: contract}}, nil
}

// NewUniswapV2RouterCaller creates a new read-only instance of UniswapV2Router, bound to a specific deployed contract.
func NewUniswapV2RouterCaller(address common.Address, caller bind.ContractCaller) (*UniswapV2RouterCaller, error) {
	contract, err := bindUniswapV2Router(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UniswapV2RouterCaller{contract: contract}, nil
}

// NewUniswapV2RouterTransactor creates a new write-only instance of UniswapV2Router, bound to a specific deployed contract.
func NewUniswapV2RouterTransactor(address common.Address, transactor bind.ContractTransactor) (*UniswapV2RouterTransactor, error) {
	contract, err := bindUniswapV2Router(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UniswapV2RouterTransactor{contract: contract}, nil
}

// NewUniswapV2RouterFilterer creates a new log filterer instance of UniswapV2Router, bound to a specific deployed contract.
func NewUniswapV2RouterFilterer(address common.Address, filterer bind.ContractFilterer) (*UniswapV2RouterFilterer, error) {
	contract, err := bindUniswapV2Router(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UniswapV2RouterFilterer{contract: contract}, nil
}

// bindUniswapV2Router binds a generic wrapper to an already deployed contract.
func bindUniswapV2Router(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := UniswapV2RouterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniswapV2Router *UniswapV2RouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniswapV2Router.Contract.UniswapV2RouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniswapV2Router *UniswapV2RouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapV2Router.Contract.UniswapV2RouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniswapV2Router *UniswapV2RouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniswapV2Router.Contract.UniswapV2RouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniswapV2Router *UniswapV2RouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniswapV2Router.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniswapV2Router *UniswapV2RouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapV2Router.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniswapV2Router *UniswapV2RouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniswapV2Router.Contract.contract.Transact(opts, method, params...)
}

// OWNABLEDIAMONDSTORAGESLOT is a free data retrieval call binding the contract method 0xe3e077ad.
//
// Solidity: function OWNABLE_DIAMOND_STORAGE_SLOT() view returns(bytes32)
func (_UniswapV2Router *UniswapV2RouterCaller) OWNABLEDIAMONDSTORAGESLOT(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _UniswapV2Router.contract.Call(opts, &out, "OWNABLE_DIAMOND_STORAGE_SLOT")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OWNABLEDIAMONDSTORAGESLOT is a free data retrieval call binding the contract method 0xe3e077ad.
//
// Solidity: function OWNABLE_DIAMOND_STORAGE_SLOT() view returns(bytes32)
func (_UniswapV2Router *UniswapV2RouterSession) OWNABLEDIAMONDSTORAGESLOT() ([32]byte, error) {
	return _UniswapV2Router.Contract.OWNABLEDIAMONDSTORAGESLOT(&_UniswapV2Router.CallOpts)
}

// OWNABLEDIAMONDSTORAGESLOT is a free data retrieval call binding the contract method 0xe3e077ad.
//
// Solidity: function OWNABLE_DIAMOND_STORAGE_SLOT() view returns(bytes32)
func (_UniswapV2Router *UniswapV2RouterCallerSession) OWNABLEDIAMONDSTORAGESLOT() ([32]byte, error) {
	return _UniswapV2Router.Contract.OWNABLEDIAMONDSTORAGESLOT(&_UniswapV2Router.CallOpts)
}

// UNISWAPV2ROUTERSTORAGESLOT is a free data retrieval call binding the contract method 0x3e5aae47.
//
// Solidity: function UNISWAP_V2_ROUTER_STORAGE_SLOT() view returns(bytes32)
func (_UniswapV2Router *UniswapV2RouterCaller) UNISWAPV2ROUTERSTORAGESLOT(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _UniswapV2Router.contract.Call(opts, &out, "UNISWAP_V2_ROUTER_STORAGE_SLOT")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// UNISWAPV2ROUTERSTORAGESLOT is a free data retrieval call binding the contract method 0x3e5aae47.
//
// Solidity: function UNISWAP_V2_ROUTER_STORAGE_SLOT() view returns(bytes32)
func (_UniswapV2Router *UniswapV2RouterSession) UNISWAPV2ROUTERSTORAGESLOT() ([32]byte, error) {
	return _UniswapV2Router.Contract.UNISWAPV2ROUTERSTORAGESLOT(&_UniswapV2Router.CallOpts)
}

// UNISWAPV2ROUTERSTORAGESLOT is a free data retrieval call binding the contract method 0x3e5aae47.
//
// Solidity: function UNISWAP_V2_ROUTER_STORAGE_SLOT() view returns(bytes32)
func (_UniswapV2Router *UniswapV2RouterCallerSession) UNISWAPV2ROUTERSTORAGESLOT() ([32]byte, error) {
	return _UniswapV2Router.Contract.UNISWAPV2ROUTERSTORAGESLOT(&_UniswapV2Router.CallOpts)
}

// GetSwapV2Router is a free data retrieval call binding the contract method 0x918e44d1.
//
// Solidity: function getSwapV2Router() view returns(address swapV2Router_)
func (_UniswapV2Router *UniswapV2RouterCaller) GetSwapV2Router(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UniswapV2Router.contract.Call(opts, &out, "getSwapV2Router")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetSwapV2Router is a free data retrieval call binding the contract method 0x918e44d1.
//
// Solidity: function getSwapV2Router() view returns(address swapV2Router_)
func (_UniswapV2Router *UniswapV2RouterSession) GetSwapV2Router() (common.Address, error) {
	return _UniswapV2Router.Contract.GetSwapV2Router(&_UniswapV2Router.CallOpts)
}

// GetSwapV2Router is a free data retrieval call binding the contract method 0x918e44d1.
//
// Solidity: function getSwapV2Router() view returns(address swapV2Router_)
func (_UniswapV2Router *UniswapV2RouterCallerSession) GetSwapV2Router() (common.Address, error) {
	return _UniswapV2Router.Contract.GetSwapV2Router(&_UniswapV2Router.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_UniswapV2Router *UniswapV2RouterCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UniswapV2Router.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_UniswapV2Router *UniswapV2RouterSession) Owner() (common.Address, error) {
	return _UniswapV2Router.Contract.Owner(&_UniswapV2Router.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_UniswapV2Router *UniswapV2RouterCallerSession) Owner() (common.Address, error) {
	return _UniswapV2Router.Contract.Owner(&_UniswapV2Router.CallOpts)
}

// SetUniswapV2RouterAddress is a paid mutator transaction binding the contract method 0xbf7fb44b.
//
// Solidity: function setUniswapV2RouterAddress(address swapV2Router_) returns()
func (_UniswapV2Router *UniswapV2RouterTransactor) SetUniswapV2RouterAddress(opts *bind.TransactOpts, swapV2Router_ common.Address) (*types.Transaction, error) {
	return _UniswapV2Router.contract.Transact(opts, "setUniswapV2RouterAddress", swapV2Router_)
}

// SetUniswapV2RouterAddress is a paid mutator transaction binding the contract method 0xbf7fb44b.
//
// Solidity: function setUniswapV2RouterAddress(address swapV2Router_) returns()
func (_UniswapV2Router *UniswapV2RouterSession) SetUniswapV2RouterAddress(swapV2Router_ common.Address) (*types.Transaction, error) {
	return _UniswapV2Router.Contract.SetUniswapV2RouterAddress(&_UniswapV2Router.TransactOpts, swapV2Router_)
}

// SetUniswapV2RouterAddress is a paid mutator transaction binding the contract method 0xbf7fb44b.
//
// Solidity: function setUniswapV2RouterAddress(address swapV2Router_) returns()
func (_UniswapV2Router *UniswapV2RouterTransactorSession) SetUniswapV2RouterAddress(swapV2Router_ common.Address) (*types.Transaction, error) {
	return _UniswapV2Router.Contract.SetUniswapV2RouterAddress(&_UniswapV2Router.TransactOpts, swapV2Router_)
}

// SwapETHForExactTokens is a paid mutator transaction binding the contract method 0x9858a066.
//
// Solidity: function swapETHForExactTokens(address receiver_, uint256 amountOut_, uint256 amountInMax_, address[] path_) payable returns()
func (_UniswapV2Router *UniswapV2RouterTransactor) SwapETHForExactTokens(opts *bind.TransactOpts, receiver_ common.Address, amountOut_ *big.Int, amountInMax_ *big.Int, path_ []common.Address) (*types.Transaction, error) {
	return _UniswapV2Router.contract.Transact(opts, "swapETHForExactTokens", receiver_, amountOut_, amountInMax_, path_)
}

// SwapETHForExactTokens is a paid mutator transaction binding the contract method 0x9858a066.
//
// Solidity: function swapETHForExactTokens(address receiver_, uint256 amountOut_, uint256 amountInMax_, address[] path_) payable returns()
func (_UniswapV2Router *UniswapV2RouterSession) SwapETHForExactTokens(receiver_ common.Address, amountOut_ *big.Int, amountInMax_ *big.Int, path_ []common.Address) (*types.Transaction, error) {
	return _UniswapV2Router.Contract.SwapETHForExactTokens(&_UniswapV2Router.TransactOpts, receiver_, amountOut_, amountInMax_, path_)
}

// SwapETHForExactTokens is a paid mutator transaction binding the contract method 0x9858a066.
//
// Solidity: function swapETHForExactTokens(address receiver_, uint256 amountOut_, uint256 amountInMax_, address[] path_) payable returns()
func (_UniswapV2Router *UniswapV2RouterTransactorSession) SwapETHForExactTokens(receiver_ common.Address, amountOut_ *big.Int, amountInMax_ *big.Int, path_ []common.Address) (*types.Transaction, error) {
	return _UniswapV2Router.Contract.SwapETHForExactTokens(&_UniswapV2Router.TransactOpts, receiver_, amountOut_, amountInMax_, path_)
}

// SwapExactETHForTokens is a paid mutator transaction binding the contract method 0x04a84c1d.
//
// Solidity: function swapExactETHForTokens(address receiver_, uint256 amountIn_, uint256 amountOutMin_, address[] path_) payable returns()
func (_UniswapV2Router *UniswapV2RouterTransactor) SwapExactETHForTokens(opts *bind.TransactOpts, receiver_ common.Address, amountIn_ *big.Int, amountOutMin_ *big.Int, path_ []common.Address) (*types.Transaction, error) {
	return _UniswapV2Router.contract.Transact(opts, "swapExactETHForTokens", receiver_, amountIn_, amountOutMin_, path_)
}

// SwapExactETHForTokens is a paid mutator transaction binding the contract method 0x04a84c1d.
//
// Solidity: function swapExactETHForTokens(address receiver_, uint256 amountIn_, uint256 amountOutMin_, address[] path_) payable returns()
func (_UniswapV2Router *UniswapV2RouterSession) SwapExactETHForTokens(receiver_ common.Address, amountIn_ *big.Int, amountOutMin_ *big.Int, path_ []common.Address) (*types.Transaction, error) {
	return _UniswapV2Router.Contract.SwapExactETHForTokens(&_UniswapV2Router.TransactOpts, receiver_, amountIn_, amountOutMin_, path_)
}

// SwapExactETHForTokens is a paid mutator transaction binding the contract method 0x04a84c1d.
//
// Solidity: function swapExactETHForTokens(address receiver_, uint256 amountIn_, uint256 amountOutMin_, address[] path_) payable returns()
func (_UniswapV2Router *UniswapV2RouterTransactorSession) SwapExactETHForTokens(receiver_ common.Address, amountIn_ *big.Int, amountOutMin_ *big.Int, path_ []common.Address) (*types.Transaction, error) {
	return _UniswapV2Router.Contract.SwapExactETHForTokens(&_UniswapV2Router.TransactOpts, receiver_, amountIn_, amountOutMin_, path_)
}

// SwapExactTokensForETH is a paid mutator transaction binding the contract method 0x2f45a771.
//
// Solidity: function swapExactTokensForETH(address receiver_, uint256 amountIn_, uint256 amountOutMin_, address[] path_) payable returns()
func (_UniswapV2Router *UniswapV2RouterTransactor) SwapExactTokensForETH(opts *bind.TransactOpts, receiver_ common.Address, amountIn_ *big.Int, amountOutMin_ *big.Int, path_ []common.Address) (*types.Transaction, error) {
	return _UniswapV2Router.contract.Transact(opts, "swapExactTokensForETH", receiver_, amountIn_, amountOutMin_, path_)
}

// SwapExactTokensForETH is a paid mutator transaction binding the contract method 0x2f45a771.
//
// Solidity: function swapExactTokensForETH(address receiver_, uint256 amountIn_, uint256 amountOutMin_, address[] path_) payable returns()
func (_UniswapV2Router *UniswapV2RouterSession) SwapExactTokensForETH(receiver_ common.Address, amountIn_ *big.Int, amountOutMin_ *big.Int, path_ []common.Address) (*types.Transaction, error) {
	return _UniswapV2Router.Contract.SwapExactTokensForETH(&_UniswapV2Router.TransactOpts, receiver_, amountIn_, amountOutMin_, path_)
}

// SwapExactTokensForETH is a paid mutator transaction binding the contract method 0x2f45a771.
//
// Solidity: function swapExactTokensForETH(address receiver_, uint256 amountIn_, uint256 amountOutMin_, address[] path_) payable returns()
func (_UniswapV2Router *UniswapV2RouterTransactorSession) SwapExactTokensForETH(receiver_ common.Address, amountIn_ *big.Int, amountOutMin_ *big.Int, path_ []common.Address) (*types.Transaction, error) {
	return _UniswapV2Router.Contract.SwapExactTokensForETH(&_UniswapV2Router.TransactOpts, receiver_, amountIn_, amountOutMin_, path_)
}

// SwapExactTokensForTokensV2 is a paid mutator transaction binding the contract method 0x84a98cf8.
//
// Solidity: function swapExactTokensForTokensV2(address receiver_, uint256 amountIn_, uint256 amountOutMin_, address[] path_) payable returns()
func (_UniswapV2Router *UniswapV2RouterTransactor) SwapExactTokensForTokensV2(opts *bind.TransactOpts, receiver_ common.Address, amountIn_ *big.Int, amountOutMin_ *big.Int, path_ []common.Address) (*types.Transaction, error) {
	return _UniswapV2Router.contract.Transact(opts, "swapExactTokensForTokensV2", receiver_, amountIn_, amountOutMin_, path_)
}

// SwapExactTokensForTokensV2 is a paid mutator transaction binding the contract method 0x84a98cf8.
//
// Solidity: function swapExactTokensForTokensV2(address receiver_, uint256 amountIn_, uint256 amountOutMin_, address[] path_) payable returns()
func (_UniswapV2Router *UniswapV2RouterSession) SwapExactTokensForTokensV2(receiver_ common.Address, amountIn_ *big.Int, amountOutMin_ *big.Int, path_ []common.Address) (*types.Transaction, error) {
	return _UniswapV2Router.Contract.SwapExactTokensForTokensV2(&_UniswapV2Router.TransactOpts, receiver_, amountIn_, amountOutMin_, path_)
}

// SwapExactTokensForTokensV2 is a paid mutator transaction binding the contract method 0x84a98cf8.
//
// Solidity: function swapExactTokensForTokensV2(address receiver_, uint256 amountIn_, uint256 amountOutMin_, address[] path_) payable returns()
func (_UniswapV2Router *UniswapV2RouterTransactorSession) SwapExactTokensForTokensV2(receiver_ common.Address, amountIn_ *big.Int, amountOutMin_ *big.Int, path_ []common.Address) (*types.Transaction, error) {
	return _UniswapV2Router.Contract.SwapExactTokensForTokensV2(&_UniswapV2Router.TransactOpts, receiver_, amountIn_, amountOutMin_, path_)
}

// SwapTokensForExactETH is a paid mutator transaction binding the contract method 0x52f6f214.
//
// Solidity: function swapTokensForExactETH(address receiver_, uint256 amountOut_, uint256 amountInMax_, address[] path_) payable returns()
func (_UniswapV2Router *UniswapV2RouterTransactor) SwapTokensForExactETH(opts *bind.TransactOpts, receiver_ common.Address, amountOut_ *big.Int, amountInMax_ *big.Int, path_ []common.Address) (*types.Transaction, error) {
	return _UniswapV2Router.contract.Transact(opts, "swapTokensForExactETH", receiver_, amountOut_, amountInMax_, path_)
}

// SwapTokensForExactETH is a paid mutator transaction binding the contract method 0x52f6f214.
//
// Solidity: function swapTokensForExactETH(address receiver_, uint256 amountOut_, uint256 amountInMax_, address[] path_) payable returns()
func (_UniswapV2Router *UniswapV2RouterSession) SwapTokensForExactETH(receiver_ common.Address, amountOut_ *big.Int, amountInMax_ *big.Int, path_ []common.Address) (*types.Transaction, error) {
	return _UniswapV2Router.Contract.SwapTokensForExactETH(&_UniswapV2Router.TransactOpts, receiver_, amountOut_, amountInMax_, path_)
}

// SwapTokensForExactETH is a paid mutator transaction binding the contract method 0x52f6f214.
//
// Solidity: function swapTokensForExactETH(address receiver_, uint256 amountOut_, uint256 amountInMax_, address[] path_) payable returns()
func (_UniswapV2Router *UniswapV2RouterTransactorSession) SwapTokensForExactETH(receiver_ common.Address, amountOut_ *big.Int, amountInMax_ *big.Int, path_ []common.Address) (*types.Transaction, error) {
	return _UniswapV2Router.Contract.SwapTokensForExactETH(&_UniswapV2Router.TransactOpts, receiver_, amountOut_, amountInMax_, path_)
}

// SwapTokensForExactTokensV2 is a paid mutator transaction binding the contract method 0xcbb9b3dd.
//
// Solidity: function swapTokensForExactTokensV2(address receiver_, uint256 amountOut_, uint256 amountInMax_, address[] path_) payable returns()
func (_UniswapV2Router *UniswapV2RouterTransactor) SwapTokensForExactTokensV2(opts *bind.TransactOpts, receiver_ common.Address, amountOut_ *big.Int, amountInMax_ *big.Int, path_ []common.Address) (*types.Transaction, error) {
	return _UniswapV2Router.contract.Transact(opts, "swapTokensForExactTokensV2", receiver_, amountOut_, amountInMax_, path_)
}

// SwapTokensForExactTokensV2 is a paid mutator transaction binding the contract method 0xcbb9b3dd.
//
// Solidity: function swapTokensForExactTokensV2(address receiver_, uint256 amountOut_, uint256 amountInMax_, address[] path_) payable returns()
func (_UniswapV2Router *UniswapV2RouterSession) SwapTokensForExactTokensV2(receiver_ common.Address, amountOut_ *big.Int, amountInMax_ *big.Int, path_ []common.Address) (*types.Transaction, error) {
	return _UniswapV2Router.Contract.SwapTokensForExactTokensV2(&_UniswapV2Router.TransactOpts, receiver_, amountOut_, amountInMax_, path_)
}

// SwapTokensForExactTokensV2 is a paid mutator transaction binding the contract method 0xcbb9b3dd.
//
// Solidity: function swapTokensForExactTokensV2(address receiver_, uint256 amountOut_, uint256 amountInMax_, address[] path_) payable returns()
func (_UniswapV2Router *UniswapV2RouterTransactorSession) SwapTokensForExactTokensV2(receiver_ common.Address, amountOut_ *big.Int, amountInMax_ *big.Int, path_ []common.Address) (*types.Transaction, error) {
	return _UniswapV2Router.Contract.SwapTokensForExactTokensV2(&_UniswapV2Router.TransactOpts, receiver_, amountOut_, amountInMax_, path_)
}
