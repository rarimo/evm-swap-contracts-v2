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

// UniswapV3RouterStorageMetaData contains all meta data concerning the UniswapV3RouterStorage contract.
var UniswapV3RouterStorageMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"UNISWAP_V3_ROUTER_STORAGE_SLOT\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSwapV3Router\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"swapV3Router_\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// UniswapV3RouterStorageABI is the input ABI used to generate the binding from.
// Deprecated: Use UniswapV3RouterStorageMetaData.ABI instead.
var UniswapV3RouterStorageABI = UniswapV3RouterStorageMetaData.ABI

// UniswapV3RouterStorage is an auto generated Go binding around an Ethereum contract.
type UniswapV3RouterStorage struct {
	UniswapV3RouterStorageCaller     // Read-only binding to the contract
	UniswapV3RouterStorageTransactor // Write-only binding to the contract
	UniswapV3RouterStorageFilterer   // Log filterer for contract events
}

// UniswapV3RouterStorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type UniswapV3RouterStorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapV3RouterStorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UniswapV3RouterStorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapV3RouterStorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UniswapV3RouterStorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UniswapV3RouterStorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UniswapV3RouterStorageSession struct {
	Contract     *UniswapV3RouterStorage // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// UniswapV3RouterStorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UniswapV3RouterStorageCallerSession struct {
	Contract *UniswapV3RouterStorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// UniswapV3RouterStorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UniswapV3RouterStorageTransactorSession struct {
	Contract     *UniswapV3RouterStorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// UniswapV3RouterStorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type UniswapV3RouterStorageRaw struct {
	Contract *UniswapV3RouterStorage // Generic contract binding to access the raw methods on
}

// UniswapV3RouterStorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UniswapV3RouterStorageCallerRaw struct {
	Contract *UniswapV3RouterStorageCaller // Generic read-only contract binding to access the raw methods on
}

// UniswapV3RouterStorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UniswapV3RouterStorageTransactorRaw struct {
	Contract *UniswapV3RouterStorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUniswapV3RouterStorage creates a new instance of UniswapV3RouterStorage, bound to a specific deployed contract.
func NewUniswapV3RouterStorage(address common.Address, backend bind.ContractBackend) (*UniswapV3RouterStorage, error) {
	contract, err := bindUniswapV3RouterStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UniswapV3RouterStorage{UniswapV3RouterStorageCaller: UniswapV3RouterStorageCaller{contract: contract}, UniswapV3RouterStorageTransactor: UniswapV3RouterStorageTransactor{contract: contract}, UniswapV3RouterStorageFilterer: UniswapV3RouterStorageFilterer{contract: contract}}, nil
}

// NewUniswapV3RouterStorageCaller creates a new read-only instance of UniswapV3RouterStorage, bound to a specific deployed contract.
func NewUniswapV3RouterStorageCaller(address common.Address, caller bind.ContractCaller) (*UniswapV3RouterStorageCaller, error) {
	contract, err := bindUniswapV3RouterStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UniswapV3RouterStorageCaller{contract: contract}, nil
}

// NewUniswapV3RouterStorageTransactor creates a new write-only instance of UniswapV3RouterStorage, bound to a specific deployed contract.
func NewUniswapV3RouterStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*UniswapV3RouterStorageTransactor, error) {
	contract, err := bindUniswapV3RouterStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UniswapV3RouterStorageTransactor{contract: contract}, nil
}

// NewUniswapV3RouterStorageFilterer creates a new log filterer instance of UniswapV3RouterStorage, bound to a specific deployed contract.
func NewUniswapV3RouterStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*UniswapV3RouterStorageFilterer, error) {
	contract, err := bindUniswapV3RouterStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UniswapV3RouterStorageFilterer{contract: contract}, nil
}

// bindUniswapV3RouterStorage binds a generic wrapper to an already deployed contract.
func bindUniswapV3RouterStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := UniswapV3RouterStorageMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniswapV3RouterStorage *UniswapV3RouterStorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniswapV3RouterStorage.Contract.UniswapV3RouterStorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniswapV3RouterStorage *UniswapV3RouterStorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapV3RouterStorage.Contract.UniswapV3RouterStorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniswapV3RouterStorage *UniswapV3RouterStorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniswapV3RouterStorage.Contract.UniswapV3RouterStorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UniswapV3RouterStorage *UniswapV3RouterStorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UniswapV3RouterStorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UniswapV3RouterStorage *UniswapV3RouterStorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UniswapV3RouterStorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UniswapV3RouterStorage *UniswapV3RouterStorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UniswapV3RouterStorage.Contract.contract.Transact(opts, method, params...)
}

// UNISWAPV3ROUTERSTORAGESLOT is a free data retrieval call binding the contract method 0x61d2945e.
//
// Solidity: function UNISWAP_V3_ROUTER_STORAGE_SLOT() view returns(bytes32)
func (_UniswapV3RouterStorage *UniswapV3RouterStorageCaller) UNISWAPV3ROUTERSTORAGESLOT(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _UniswapV3RouterStorage.contract.Call(opts, &out, "UNISWAP_V3_ROUTER_STORAGE_SLOT")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// UNISWAPV3ROUTERSTORAGESLOT is a free data retrieval call binding the contract method 0x61d2945e.
//
// Solidity: function UNISWAP_V3_ROUTER_STORAGE_SLOT() view returns(bytes32)
func (_UniswapV3RouterStorage *UniswapV3RouterStorageSession) UNISWAPV3ROUTERSTORAGESLOT() ([32]byte, error) {
	return _UniswapV3RouterStorage.Contract.UNISWAPV3ROUTERSTORAGESLOT(&_UniswapV3RouterStorage.CallOpts)
}

// UNISWAPV3ROUTERSTORAGESLOT is a free data retrieval call binding the contract method 0x61d2945e.
//
// Solidity: function UNISWAP_V3_ROUTER_STORAGE_SLOT() view returns(bytes32)
func (_UniswapV3RouterStorage *UniswapV3RouterStorageCallerSession) UNISWAPV3ROUTERSTORAGESLOT() ([32]byte, error) {
	return _UniswapV3RouterStorage.Contract.UNISWAPV3ROUTERSTORAGESLOT(&_UniswapV3RouterStorage.CallOpts)
}

// GetSwapV3Router is a free data retrieval call binding the contract method 0xe47a7515.
//
// Solidity: function getSwapV3Router() view returns(address swapV3Router_)
func (_UniswapV3RouterStorage *UniswapV3RouterStorageCaller) GetSwapV3Router(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UniswapV3RouterStorage.contract.Call(opts, &out, "getSwapV3Router")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetSwapV3Router is a free data retrieval call binding the contract method 0xe47a7515.
//
// Solidity: function getSwapV3Router() view returns(address swapV3Router_)
func (_UniswapV3RouterStorage *UniswapV3RouterStorageSession) GetSwapV3Router() (common.Address, error) {
	return _UniswapV3RouterStorage.Contract.GetSwapV3Router(&_UniswapV3RouterStorage.CallOpts)
}

// GetSwapV3Router is a free data retrieval call binding the contract method 0xe47a7515.
//
// Solidity: function getSwapV3Router() view returns(address swapV3Router_)
func (_UniswapV3RouterStorage *UniswapV3RouterStorageCallerSession) GetSwapV3Router() (common.Address, error) {
	return _UniswapV3RouterStorage.Contract.GetSwapV3Router(&_UniswapV3RouterStorage.CallOpts)
}
