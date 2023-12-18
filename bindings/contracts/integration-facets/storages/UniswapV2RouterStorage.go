// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package storages

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

// UniswapV2RouterStorageMetaData contains all meta data concerning the UniswapV2RouterStorage contract.
var UniswapV2RouterStorageMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"UNISWAP_V2_ROUTER_STORAGE_SLOT\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSwapV2Router\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"swapV2Router_\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// UniswapV2RouterStorageABI is the input ABI used to generate the binding from.
// Deprecated: Use UniswapV2RouterStorageMetaData.ABI instead.
var UniswapV2RouterStorageABI = UniswapV2RouterStorageMetaData.ABI

// UniswapV2RouterStorage is an auto generated Go binding around an Ethereum contract.
type UniswapV2RouterStorage struct {
	UniswapV2RouterStorageCaller     // Read-only binding to the contract
	UniswapV2RouterStorageTransactor // Write-only binding to the contract
	UniswapV2RouterStorageFilterer   // Log filterer for contract events
}

// UniswapV2RouterStorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type UniswapV2RouterStorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapV2RouterStorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UniswapV2RouterStorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapV2RouterStorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UniswapV2RouterStorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapV2RouterStorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UniswapV2RouterStorageSession struct {
	Contract     *UniswapV2RouterStorage // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// UniswapV2RouterStorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UniswapV2RouterStorageCallerSession struct {
	Contract *UniswapV2RouterStorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// UniswapV2RouterStorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UniswapV2RouterStorageTransactorSession struct {
	Contract     *UniswapV2RouterStorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// UniswapV2RouterStorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type UniswapV2RouterStorageRaw struct {
	Contract *UniswapV2RouterStorage // Generic contract binding to access the raw methods on
}

// UniswapV2RouterStorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UniswapV2RouterStorageCallerRaw struct {
	Contract *UniswapV2RouterStorageCaller // Generic read-only contract binding to access the raw methods on
}

// UniswapV2RouterStorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UniswapV2RouterStorageTransactorRaw struct {
	Contract *UniswapV2RouterStorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUniswapV2RouterStorage creates a new instance of UniswapV2RouterStorage, bound to a specific deployed contract.
func NewUniswapV2RouterStorage(address common.Address, backend bind.ContractBackend) (*UniswapV2RouterStorage, error) {
	contract, err := bindUniswapV2RouterStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UniswapV2RouterStorage{UniswapV2RouterStorageCaller: UniswapV2RouterStorageCaller{contract: contract}, UniswapV2RouterStorageTransactor: UniswapV2RouterStorageTransactor{contract: contract}, UniswapV2RouterStorageFilterer: UniswapV2RouterStorageFilterer{contract: contract}}, nil
}

// NewUniswapV2RouterStorageCaller creates a new read-only instance of UniswapV2RouterStorage, bound to a specific deployed contract.
func NewUniswapV2RouterStorageCaller(address common.Address, caller bind.ContractCaller) (*UniswapV2RouterStorageCaller, error) {
	contract, err := bindUniswapV2RouterStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UniswapV2RouterStorageCaller{contract: contract}, nil
}

// NewUniswapV2RouterStorageTransactor creates a new write-only instance of UniswapV2RouterStorage, bound to a specific deployed contract.
func NewUniswapV2RouterStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*UniswapV2RouterStorageTransactor, error) {
	contract, err := bindUniswapV2RouterStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UniswapV2RouterStorageTransactor{contract: contract}, nil
}

// NewUniswapV2RouterStorageFilterer creates a new log filterer instance of UniswapV2RouterStorage, bound to a specific deployed contract.
func NewUniswapV2RouterStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*UniswapV2RouterStorageFilterer, error) {
	contract, err := bindUniswapV2RouterStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UniswapV2RouterStorageFilterer{contract: contract}, nil
}

// bindUniswapV2RouterStorage binds a generic wrapper to an already deployed contract.
func bindUniswapV2RouterStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := UniswapV2RouterStorageMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniswapV2RouterStorage *UniswapV2RouterStorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniswapV2RouterStorage.Contract.UniswapV2RouterStorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniswapV2RouterStorage *UniswapV2RouterStorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapV2RouterStorage.Contract.UniswapV2RouterStorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniswapV2RouterStorage *UniswapV2RouterStorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniswapV2RouterStorage.Contract.UniswapV2RouterStorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniswapV2RouterStorage *UniswapV2RouterStorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniswapV2RouterStorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniswapV2RouterStorage *UniswapV2RouterStorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapV2RouterStorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniswapV2RouterStorage *UniswapV2RouterStorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniswapV2RouterStorage.Contract.contract.Transact(opts, method, params...)
}

// UNISWAPV2ROUTERSTORAGESLOT is a free data retrieval call binding the contract method 0x3e5aae47.
//
// Solidity: function UNISWAP_V2_ROUTER_STORAGE_SLOT() view returns(bytes32)
func (_UniswapV2RouterStorage *UniswapV2RouterStorageCaller) UNISWAPV2ROUTERSTORAGESLOT(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _UniswapV2RouterStorage.contract.Call(opts, &out, "UNISWAP_V2_ROUTER_STORAGE_SLOT")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// UNISWAPV2ROUTERSTORAGESLOT is a free data retrieval call binding the contract method 0x3e5aae47.
//
// Solidity: function UNISWAP_V2_ROUTER_STORAGE_SLOT() view returns(bytes32)
func (_UniswapV2RouterStorage *UniswapV2RouterStorageSession) UNISWAPV2ROUTERSTORAGESLOT() ([32]byte, error) {
	return _UniswapV2RouterStorage.Contract.UNISWAPV2ROUTERSTORAGESLOT(&_UniswapV2RouterStorage.CallOpts)
}

// UNISWAPV2ROUTERSTORAGESLOT is a free data retrieval call binding the contract method 0x3e5aae47.
//
// Solidity: function UNISWAP_V2_ROUTER_STORAGE_SLOT() view returns(bytes32)
func (_UniswapV2RouterStorage *UniswapV2RouterStorageCallerSession) UNISWAPV2ROUTERSTORAGESLOT() ([32]byte, error) {
	return _UniswapV2RouterStorage.Contract.UNISWAPV2ROUTERSTORAGESLOT(&_UniswapV2RouterStorage.CallOpts)
}

// GetSwapV2Router is a free data retrieval call binding the contract method 0x918e44d1.
//
// Solidity: function getSwapV2Router() view returns(address swapV2Router_)
func (_UniswapV2RouterStorage *UniswapV2RouterStorageCaller) GetSwapV2Router(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UniswapV2RouterStorage.contract.Call(opts, &out, "getSwapV2Router")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetSwapV2Router is a free data retrieval call binding the contract method 0x918e44d1.
//
// Solidity: function getSwapV2Router() view returns(address swapV2Router_)
func (_UniswapV2RouterStorage *UniswapV2RouterStorageSession) GetSwapV2Router() (common.Address, error) {
	return _UniswapV2RouterStorage.Contract.GetSwapV2Router(&_UniswapV2RouterStorage.CallOpts)
}

// GetSwapV2Router is a free data retrieval call binding the contract method 0x918e44d1.
//
// Solidity: function getSwapV2Router() view returns(address swapV2Router_)
func (_UniswapV2RouterStorage *UniswapV2RouterStorageCallerSession) GetSwapV2Router() (common.Address, error) {
	return _UniswapV2RouterStorage.Contract.GetSwapV2Router(&_UniswapV2RouterStorage.CallOpts)
}
