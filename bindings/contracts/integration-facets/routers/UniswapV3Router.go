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

// UniswapV3RouterMetaData contains all meta data concerning the UniswapV3Router contract.
var UniswapV3RouterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"OWNABLE_DIAMOND_STORAGE_SLOT\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UNISWAP_V3_ROUTER_STORAGE_SLOT\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"isNative_\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"receiver_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutMinimum_\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"path_\",\"type\":\"bytes\"}],\"name\":\"exactInput\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"isNative_\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"receiver_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountOut_\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInMaximum_\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"path_\",\"type\":\"bytes\"}],\"name\":\"exactOutput\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSwapV3Router\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"swapV3Router_\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"swapV3Router_\",\"type\":\"address\"}],\"name\":\"setUniswapV3RouterAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// UniswapV3RouterABI is the input ABI used to generate the binding from.
// Deprecated: Use UniswapV3RouterMetaData.ABI instead.
var UniswapV3RouterABI = UniswapV3RouterMetaData.ABI

// UniswapV3Router is an auto generated Go binding around an Ethereum contract.
type UniswapV3Router struct {
	UniswapV3RouterCaller     // Read-only binding to the contract
	UniswapV3RouterTransactor // Write-only binding to the contract
	UniswapV3RouterFilterer   // Log filterer for contract events
}

// UniswapV3RouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type UniswapV3RouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapV3RouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UniswapV3RouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapV3RouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UniswapV3RouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapV3RouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UniswapV3RouterSession struct {
	Contract     *UniswapV3Router  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UniswapV3RouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UniswapV3RouterCallerSession struct {
	Contract *UniswapV3RouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// UniswapV3RouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UniswapV3RouterTransactorSession struct {
	Contract     *UniswapV3RouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// UniswapV3RouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type UniswapV3RouterRaw struct {
	Contract *UniswapV3Router // Generic contract binding to access the raw methods on
}

// UniswapV3RouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UniswapV3RouterCallerRaw struct {
	Contract *UniswapV3RouterCaller // Generic read-only contract binding to access the raw methods on
}

// UniswapV3RouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UniswapV3RouterTransactorRaw struct {
	Contract *UniswapV3RouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUniswapV3Router creates a new instance of UniswapV3Router, bound to a specific deployed contract.
func NewUniswapV3Router(address common.Address, backend bind.ContractBackend) (*UniswapV3Router, error) {
	contract, err := bindUniswapV3Router(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UniswapV3Router{UniswapV3RouterCaller: UniswapV3RouterCaller{contract: contract}, UniswapV3RouterTransactor: UniswapV3RouterTransactor{contract: contract}, UniswapV3RouterFilterer: UniswapV3RouterFilterer{contract: contract}}, nil
}

// NewUniswapV3RouterCaller creates a new read-only instance of UniswapV3Router, bound to a specific deployed contract.
func NewUniswapV3RouterCaller(address common.Address, caller bind.ContractCaller) (*UniswapV3RouterCaller, error) {
	contract, err := bindUniswapV3Router(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UniswapV3RouterCaller{contract: contract}, nil
}

// NewUniswapV3RouterTransactor creates a new write-only instance of UniswapV3Router, bound to a specific deployed contract.
func NewUniswapV3RouterTransactor(address common.Address, transactor bind.ContractTransactor) (*UniswapV3RouterTransactor, error) {
	contract, err := bindUniswapV3Router(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UniswapV3RouterTransactor{contract: contract}, nil
}

// NewUniswapV3RouterFilterer creates a new log filterer instance of UniswapV3Router, bound to a specific deployed contract.
func NewUniswapV3RouterFilterer(address common.Address, filterer bind.ContractFilterer) (*UniswapV3RouterFilterer, error) {
	contract, err := bindUniswapV3Router(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UniswapV3RouterFilterer{contract: contract}, nil
}

// bindUniswapV3Router binds a generic wrapper to an already deployed contract.
func bindUniswapV3Router(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := UniswapV3RouterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniswapV3Router *UniswapV3RouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniswapV3Router.Contract.UniswapV3RouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniswapV3Router *UniswapV3RouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapV3Router.Contract.UniswapV3RouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniswapV3Router *UniswapV3RouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniswapV3Router.Contract.UniswapV3RouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniswapV3Router *UniswapV3RouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniswapV3Router.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniswapV3Router *UniswapV3RouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapV3Router.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniswapV3Router *UniswapV3RouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniswapV3Router.Contract.contract.Transact(opts, method, params...)
}

// OWNABLEDIAMONDSTORAGESLOT is a free data retrieval call binding the contract method 0xe3e077ad.
//
// Solidity: function OWNABLE_DIAMOND_STORAGE_SLOT() view returns(bytes32)
func (_UniswapV3Router *UniswapV3RouterCaller) OWNABLEDIAMONDSTORAGESLOT(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _UniswapV3Router.contract.Call(opts, &out, "OWNABLE_DIAMOND_STORAGE_SLOT")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OWNABLEDIAMONDSTORAGESLOT is a free data retrieval call binding the contract method 0xe3e077ad.
//
// Solidity: function OWNABLE_DIAMOND_STORAGE_SLOT() view returns(bytes32)
func (_UniswapV3Router *UniswapV3RouterSession) OWNABLEDIAMONDSTORAGESLOT() ([32]byte, error) {
	return _UniswapV3Router.Contract.OWNABLEDIAMONDSTORAGESLOT(&_UniswapV3Router.CallOpts)
}

// OWNABLEDIAMONDSTORAGESLOT is a free data retrieval call binding the contract method 0xe3e077ad.
//
// Solidity: function OWNABLE_DIAMOND_STORAGE_SLOT() view returns(bytes32)
func (_UniswapV3Router *UniswapV3RouterCallerSession) OWNABLEDIAMONDSTORAGESLOT() ([32]byte, error) {
	return _UniswapV3Router.Contract.OWNABLEDIAMONDSTORAGESLOT(&_UniswapV3Router.CallOpts)
}

// UNISWAPV3ROUTERSTORAGESLOT is a free data retrieval call binding the contract method 0x61d2945e.
//
// Solidity: function UNISWAP_V3_ROUTER_STORAGE_SLOT() view returns(bytes32)
func (_UniswapV3Router *UniswapV3RouterCaller) UNISWAPV3ROUTERSTORAGESLOT(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _UniswapV3Router.contract.Call(opts, &out, "UNISWAP_V3_ROUTER_STORAGE_SLOT")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// UNISWAPV3ROUTERSTORAGESLOT is a free data retrieval call binding the contract method 0x61d2945e.
//
// Solidity: function UNISWAP_V3_ROUTER_STORAGE_SLOT() view returns(bytes32)
func (_UniswapV3Router *UniswapV3RouterSession) UNISWAPV3ROUTERSTORAGESLOT() ([32]byte, error) {
	return _UniswapV3Router.Contract.UNISWAPV3ROUTERSTORAGESLOT(&_UniswapV3Router.CallOpts)
}

// UNISWAPV3ROUTERSTORAGESLOT is a free data retrieval call binding the contract method 0x61d2945e.
//
// Solidity: function UNISWAP_V3_ROUTER_STORAGE_SLOT() view returns(bytes32)
func (_UniswapV3Router *UniswapV3RouterCallerSession) UNISWAPV3ROUTERSTORAGESLOT() ([32]byte, error) {
	return _UniswapV3Router.Contract.UNISWAPV3ROUTERSTORAGESLOT(&_UniswapV3Router.CallOpts)
}

// GetSwapV3Router is a free data retrieval call binding the contract method 0xe47a7515.
//
// Solidity: function getSwapV3Router() view returns(address swapV3Router_)
func (_UniswapV3Router *UniswapV3RouterCaller) GetSwapV3Router(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UniswapV3Router.contract.Call(opts, &out, "getSwapV3Router")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetSwapV3Router is a free data retrieval call binding the contract method 0xe47a7515.
//
// Solidity: function getSwapV3Router() view returns(address swapV3Router_)
func (_UniswapV3Router *UniswapV3RouterSession) GetSwapV3Router() (common.Address, error) {
	return _UniswapV3Router.Contract.GetSwapV3Router(&_UniswapV3Router.CallOpts)
}

// GetSwapV3Router is a free data retrieval call binding the contract method 0xe47a7515.
//
// Solidity: function getSwapV3Router() view returns(address swapV3Router_)
func (_UniswapV3Router *UniswapV3RouterCallerSession) GetSwapV3Router() (common.Address, error) {
	return _UniswapV3Router.Contract.GetSwapV3Router(&_UniswapV3Router.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_UniswapV3Router *UniswapV3RouterCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UniswapV3Router.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_UniswapV3Router *UniswapV3RouterSession) Owner() (common.Address, error) {
	return _UniswapV3Router.Contract.Owner(&_UniswapV3Router.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_UniswapV3Router *UniswapV3RouterCallerSession) Owner() (common.Address, error) {
	return _UniswapV3Router.Contract.Owner(&_UniswapV3Router.CallOpts)
}

// ExactInput is a paid mutator transaction binding the contract method 0x272bf559.
//
// Solidity: function exactInput(bool isNative_, address receiver_, uint256 amountIn_, uint256 amountOutMinimum_, bytes path_) payable returns()
func (_UniswapV3Router *UniswapV3RouterTransactor) ExactInput(opts *bind.TransactOpts, isNative_ bool, receiver_ common.Address, amountIn_ *big.Int, amountOutMinimum_ *big.Int, path_ []byte) (*types.Transaction, error) {
	return _UniswapV3Router.contract.Transact(opts, "exactInput", isNative_, receiver_, amountIn_, amountOutMinimum_, path_)
}

// ExactInput is a paid mutator transaction binding the contract method 0x272bf559.
//
// Solidity: function exactInput(bool isNative_, address receiver_, uint256 amountIn_, uint256 amountOutMinimum_, bytes path_) payable returns()
func (_UniswapV3Router *UniswapV3RouterSession) ExactInput(isNative_ bool, receiver_ common.Address, amountIn_ *big.Int, amountOutMinimum_ *big.Int, path_ []byte) (*types.Transaction, error) {
	return _UniswapV3Router.Contract.ExactInput(&_UniswapV3Router.TransactOpts, isNative_, receiver_, amountIn_, amountOutMinimum_, path_)
}

// ExactInput is a paid mutator transaction binding the contract method 0x272bf559.
//
// Solidity: function exactInput(bool isNative_, address receiver_, uint256 amountIn_, uint256 amountOutMinimum_, bytes path_) payable returns()
func (_UniswapV3Router *UniswapV3RouterTransactorSession) ExactInput(isNative_ bool, receiver_ common.Address, amountIn_ *big.Int, amountOutMinimum_ *big.Int, path_ []byte) (*types.Transaction, error) {
	return _UniswapV3Router.Contract.ExactInput(&_UniswapV3Router.TransactOpts, isNative_, receiver_, amountIn_, amountOutMinimum_, path_)
}

// ExactOutput is a paid mutator transaction binding the contract method 0x654833e7.
//
// Solidity: function exactOutput(bool isNative_, address receiver_, uint256 amountOut_, uint256 amountInMaximum_, bytes path_) payable returns()
func (_UniswapV3Router *UniswapV3RouterTransactor) ExactOutput(opts *bind.TransactOpts, isNative_ bool, receiver_ common.Address, amountOut_ *big.Int, amountInMaximum_ *big.Int, path_ []byte) (*types.Transaction, error) {
	return _UniswapV3Router.contract.Transact(opts, "exactOutput", isNative_, receiver_, amountOut_, amountInMaximum_, path_)
}

// ExactOutput is a paid mutator transaction binding the contract method 0x654833e7.
//
// Solidity: function exactOutput(bool isNative_, address receiver_, uint256 amountOut_, uint256 amountInMaximum_, bytes path_) payable returns()
func (_UniswapV3Router *UniswapV3RouterSession) ExactOutput(isNative_ bool, receiver_ common.Address, amountOut_ *big.Int, amountInMaximum_ *big.Int, path_ []byte) (*types.Transaction, error) {
	return _UniswapV3Router.Contract.ExactOutput(&_UniswapV3Router.TransactOpts, isNative_, receiver_, amountOut_, amountInMaximum_, path_)
}

// ExactOutput is a paid mutator transaction binding the contract method 0x654833e7.
//
// Solidity: function exactOutput(bool isNative_, address receiver_, uint256 amountOut_, uint256 amountInMaximum_, bytes path_) payable returns()
func (_UniswapV3Router *UniswapV3RouterTransactorSession) ExactOutput(isNative_ bool, receiver_ common.Address, amountOut_ *big.Int, amountInMaximum_ *big.Int, path_ []byte) (*types.Transaction, error) {
	return _UniswapV3Router.Contract.ExactOutput(&_UniswapV3Router.TransactOpts, isNative_, receiver_, amountOut_, amountInMaximum_, path_)
}

// SetUniswapV3RouterAddress is a paid mutator transaction binding the contract method 0xa22e091c.
//
// Solidity: function setUniswapV3RouterAddress(address swapV3Router_) returns()
func (_UniswapV3Router *UniswapV3RouterTransactor) SetUniswapV3RouterAddress(opts *bind.TransactOpts, swapV3Router_ common.Address) (*types.Transaction, error) {
	return _UniswapV3Router.contract.Transact(opts, "setUniswapV3RouterAddress", swapV3Router_)
}

// SetUniswapV3RouterAddress is a paid mutator transaction binding the contract method 0xa22e091c.
//
// Solidity: function setUniswapV3RouterAddress(address swapV3Router_) returns()
func (_UniswapV3Router *UniswapV3RouterSession) SetUniswapV3RouterAddress(swapV3Router_ common.Address) (*types.Transaction, error) {
	return _UniswapV3Router.Contract.SetUniswapV3RouterAddress(&_UniswapV3Router.TransactOpts, swapV3Router_)
}

// SetUniswapV3RouterAddress is a paid mutator transaction binding the contract method 0xa22e091c.
//
// Solidity: function setUniswapV3RouterAddress(address swapV3Router_) returns()
func (_UniswapV3Router *UniswapV3RouterTransactorSession) SetUniswapV3RouterAddress(swapV3Router_ common.Address) (*types.Transaction, error) {
	return _UniswapV3Router.Contract.SetUniswapV3RouterAddress(&_UniswapV3Router.TransactOpts, swapV3Router_)
}
