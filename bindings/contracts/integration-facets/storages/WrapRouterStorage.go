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

// WrapRouterStorageMetaData contains all meta data concerning the WrapRouterStorage contract.
var WrapRouterStorageMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"WRAP_ROUTER_STORAGE_SLOT\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getWrappedNativeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"wrappedNative_\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// WrapRouterStorageABI is the input ABI used to generate the binding from.
// Deprecated: Use WrapRouterStorageMetaData.ABI instead.
var WrapRouterStorageABI = WrapRouterStorageMetaData.ABI

// WrapRouterStorage is an auto generated Go binding around an Ethereum contract.
type WrapRouterStorage struct {
	WrapRouterStorageCaller     // Read-only binding to the contract
	WrapRouterStorageTransactor // Write-only binding to the contract
	WrapRouterStorageFilterer   // Log filterer for contract events
}

// WrapRouterStorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type WrapRouterStorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WrapRouterStorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WrapRouterStorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WrapRouterStorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WrapRouterStorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WrapRouterStorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WrapRouterStorageSession struct {
	Contract     *WrapRouterStorage // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// WrapRouterStorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WrapRouterStorageCallerSession struct {
	Contract *WrapRouterStorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// WrapRouterStorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WrapRouterStorageTransactorSession struct {
	Contract     *WrapRouterStorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// WrapRouterStorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type WrapRouterStorageRaw struct {
	Contract *WrapRouterStorage // Generic contract binding to access the raw methods on
}

// WrapRouterStorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WrapRouterStorageCallerRaw struct {
	Contract *WrapRouterStorageCaller // Generic read-only contract binding to access the raw methods on
}

// WrapRouterStorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WrapRouterStorageTransactorRaw struct {
	Contract *WrapRouterStorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWrapRouterStorage creates a new instance of WrapRouterStorage, bound to a specific deployed contract.
func NewWrapRouterStorage(address common.Address, backend bind.ContractBackend) (*WrapRouterStorage, error) {
	contract, err := bindWrapRouterStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WrapRouterStorage{WrapRouterStorageCaller: WrapRouterStorageCaller{contract: contract}, WrapRouterStorageTransactor: WrapRouterStorageTransactor{contract: contract}, WrapRouterStorageFilterer: WrapRouterStorageFilterer{contract: contract}}, nil
}

// NewWrapRouterStorageCaller creates a new read-only instance of WrapRouterStorage, bound to a specific deployed contract.
func NewWrapRouterStorageCaller(address common.Address, caller bind.ContractCaller) (*WrapRouterStorageCaller, error) {
	contract, err := bindWrapRouterStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WrapRouterStorageCaller{contract: contract}, nil
}

// NewWrapRouterStorageTransactor creates a new write-only instance of WrapRouterStorage, bound to a specific deployed contract.
func NewWrapRouterStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*WrapRouterStorageTransactor, error) {
	contract, err := bindWrapRouterStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WrapRouterStorageTransactor{contract: contract}, nil
}

// NewWrapRouterStorageFilterer creates a new log filterer instance of WrapRouterStorage, bound to a specific deployed contract.
func NewWrapRouterStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*WrapRouterStorageFilterer, error) {
	contract, err := bindWrapRouterStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WrapRouterStorageFilterer{contract: contract}, nil
}

// bindWrapRouterStorage binds a generic wrapper to an already deployed contract.
func bindWrapRouterStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := WrapRouterStorageMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WrapRouterStorage *WrapRouterStorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WrapRouterStorage.Contract.WrapRouterStorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WrapRouterStorage *WrapRouterStorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WrapRouterStorage.Contract.WrapRouterStorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WrapRouterStorage *WrapRouterStorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WrapRouterStorage.Contract.WrapRouterStorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WrapRouterStorage *WrapRouterStorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WrapRouterStorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WrapRouterStorage *WrapRouterStorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WrapRouterStorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WrapRouterStorage *WrapRouterStorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WrapRouterStorage.Contract.contract.Transact(opts, method, params...)
}

// WRAPROUTERSTORAGESLOT is a free data retrieval call binding the contract method 0x53a877c0.
//
// Solidity: function WRAP_ROUTER_STORAGE_SLOT() view returns(bytes32)
func (_WrapRouterStorage *WrapRouterStorageCaller) WRAPROUTERSTORAGESLOT(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _WrapRouterStorage.contract.Call(opts, &out, "WRAP_ROUTER_STORAGE_SLOT")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// WRAPROUTERSTORAGESLOT is a free data retrieval call binding the contract method 0x53a877c0.
//
// Solidity: function WRAP_ROUTER_STORAGE_SLOT() view returns(bytes32)
func (_WrapRouterStorage *WrapRouterStorageSession) WRAPROUTERSTORAGESLOT() ([32]byte, error) {
	return _WrapRouterStorage.Contract.WRAPROUTERSTORAGESLOT(&_WrapRouterStorage.CallOpts)
}

// WRAPROUTERSTORAGESLOT is a free data retrieval call binding the contract method 0x53a877c0.
//
// Solidity: function WRAP_ROUTER_STORAGE_SLOT() view returns(bytes32)
func (_WrapRouterStorage *WrapRouterStorageCallerSession) WRAPROUTERSTORAGESLOT() ([32]byte, error) {
	return _WrapRouterStorage.Contract.WRAPROUTERSTORAGESLOT(&_WrapRouterStorage.CallOpts)
}

// GetWrappedNativeAddress is a free data retrieval call binding the contract method 0x4e07eb5e.
//
// Solidity: function getWrappedNativeAddress() view returns(address wrappedNative_)
func (_WrapRouterStorage *WrapRouterStorageCaller) GetWrappedNativeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WrapRouterStorage.contract.Call(opts, &out, "getWrappedNativeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetWrappedNativeAddress is a free data retrieval call binding the contract method 0x4e07eb5e.
//
// Solidity: function getWrappedNativeAddress() view returns(address wrappedNative_)
func (_WrapRouterStorage *WrapRouterStorageSession) GetWrappedNativeAddress() (common.Address, error) {
	return _WrapRouterStorage.Contract.GetWrappedNativeAddress(&_WrapRouterStorage.CallOpts)
}

// GetWrappedNativeAddress is a free data retrieval call binding the contract method 0x4e07eb5e.
//
// Solidity: function getWrappedNativeAddress() view returns(address wrappedNative_)
func (_WrapRouterStorage *WrapRouterStorageCallerSession) GetWrappedNativeAddress() (common.Address, error) {
	return _WrapRouterStorage.Contract.GetWrappedNativeAddress(&_WrapRouterStorage.CallOpts)
}
