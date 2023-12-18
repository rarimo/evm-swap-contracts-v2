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

// TraderJoeRouterStorageMetaData contains all meta data concerning the TraderJoeRouterStorage contract.
var TraderJoeRouterStorageMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"TRADER_JOE_ROUTER_STORAGE_SLOT\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTraderJoeRouter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"traderJoeRouter_\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// TraderJoeRouterStorageABI is the input ABI used to generate the binding from.
// Deprecated: Use TraderJoeRouterStorageMetaData.ABI instead.
var TraderJoeRouterStorageABI = TraderJoeRouterStorageMetaData.ABI

// TraderJoeRouterStorage is an auto generated Go binding around an Ethereum contract.
type TraderJoeRouterStorage struct {
	TraderJoeRouterStorageCaller     // Read-only binding to the contract
	TraderJoeRouterStorageTransactor // Write-only binding to the contract
	TraderJoeRouterStorageFilterer   // Log filterer for contract events
}

// TraderJoeRouterStorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type TraderJoeRouterStorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TraderJoeRouterStorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TraderJoeRouterStorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TraderJoeRouterStorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TraderJoeRouterStorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TraderJoeRouterStorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TraderJoeRouterStorageSession struct {
	Contract     *TraderJoeRouterStorage // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// TraderJoeRouterStorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TraderJoeRouterStorageCallerSession struct {
	Contract *TraderJoeRouterStorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// TraderJoeRouterStorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TraderJoeRouterStorageTransactorSession struct {
	Contract     *TraderJoeRouterStorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// TraderJoeRouterStorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type TraderJoeRouterStorageRaw struct {
	Contract *TraderJoeRouterStorage // Generic contract binding to access the raw methods on
}

// TraderJoeRouterStorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TraderJoeRouterStorageCallerRaw struct {
	Contract *TraderJoeRouterStorageCaller // Generic read-only contract binding to access the raw methods on
}

// TraderJoeRouterStorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TraderJoeRouterStorageTransactorRaw struct {
	Contract *TraderJoeRouterStorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTraderJoeRouterStorage creates a new instance of TraderJoeRouterStorage, bound to a specific deployed contract.
func NewTraderJoeRouterStorage(address common.Address, backend bind.ContractBackend) (*TraderJoeRouterStorage, error) {
	contract, err := bindTraderJoeRouterStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TraderJoeRouterStorage{TraderJoeRouterStorageCaller: TraderJoeRouterStorageCaller{contract: contract}, TraderJoeRouterStorageTransactor: TraderJoeRouterStorageTransactor{contract: contract}, TraderJoeRouterStorageFilterer: TraderJoeRouterStorageFilterer{contract: contract}}, nil
}

// NewTraderJoeRouterStorageCaller creates a new read-only instance of TraderJoeRouterStorage, bound to a specific deployed contract.
func NewTraderJoeRouterStorageCaller(address common.Address, caller bind.ContractCaller) (*TraderJoeRouterStorageCaller, error) {
	contract, err := bindTraderJoeRouterStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TraderJoeRouterStorageCaller{contract: contract}, nil
}

// NewTraderJoeRouterStorageTransactor creates a new write-only instance of TraderJoeRouterStorage, bound to a specific deployed contract.
func NewTraderJoeRouterStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*TraderJoeRouterStorageTransactor, error) {
	contract, err := bindTraderJoeRouterStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TraderJoeRouterStorageTransactor{contract: contract}, nil
}

// NewTraderJoeRouterStorageFilterer creates a new log filterer instance of TraderJoeRouterStorage, bound to a specific deployed contract.
func NewTraderJoeRouterStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*TraderJoeRouterStorageFilterer, error) {
	contract, err := bindTraderJoeRouterStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TraderJoeRouterStorageFilterer{contract: contract}, nil
}

// bindTraderJoeRouterStorage binds a generic wrapper to an already deployed contract.
func bindTraderJoeRouterStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TraderJoeRouterStorageMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TraderJoeRouterStorage *TraderJoeRouterStorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TraderJoeRouterStorage.Contract.TraderJoeRouterStorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TraderJoeRouterStorage *TraderJoeRouterStorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TraderJoeRouterStorage.Contract.TraderJoeRouterStorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TraderJoeRouterStorage *TraderJoeRouterStorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TraderJoeRouterStorage.Contract.TraderJoeRouterStorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TraderJoeRouterStorage *TraderJoeRouterStorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TraderJoeRouterStorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TraderJoeRouterStorage *TraderJoeRouterStorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TraderJoeRouterStorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TraderJoeRouterStorage *TraderJoeRouterStorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TraderJoeRouterStorage.Contract.contract.Transact(opts, method, params...)
}

// TRADERJOEROUTERSTORAGESLOT is a free data retrieval call binding the contract method 0x7d35e208.
//
// Solidity: function TRADER_JOE_ROUTER_STORAGE_SLOT() view returns(bytes32)
func (_TraderJoeRouterStorage *TraderJoeRouterStorageCaller) TRADERJOEROUTERSTORAGESLOT(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TraderJoeRouterStorage.contract.Call(opts, &out, "TRADER_JOE_ROUTER_STORAGE_SLOT")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TRADERJOEROUTERSTORAGESLOT is a free data retrieval call binding the contract method 0x7d35e208.
//
// Solidity: function TRADER_JOE_ROUTER_STORAGE_SLOT() view returns(bytes32)
func (_TraderJoeRouterStorage *TraderJoeRouterStorageSession) TRADERJOEROUTERSTORAGESLOT() ([32]byte, error) {
	return _TraderJoeRouterStorage.Contract.TRADERJOEROUTERSTORAGESLOT(&_TraderJoeRouterStorage.CallOpts)
}

// TRADERJOEROUTERSTORAGESLOT is a free data retrieval call binding the contract method 0x7d35e208.
//
// Solidity: function TRADER_JOE_ROUTER_STORAGE_SLOT() view returns(bytes32)
func (_TraderJoeRouterStorage *TraderJoeRouterStorageCallerSession) TRADERJOEROUTERSTORAGESLOT() ([32]byte, error) {
	return _TraderJoeRouterStorage.Contract.TRADERJOEROUTERSTORAGESLOT(&_TraderJoeRouterStorage.CallOpts)
}

// GetTraderJoeRouter is a free data retrieval call binding the contract method 0x3a144e2d.
//
// Solidity: function getTraderJoeRouter() view returns(address traderJoeRouter_)
func (_TraderJoeRouterStorage *TraderJoeRouterStorageCaller) GetTraderJoeRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TraderJoeRouterStorage.contract.Call(opts, &out, "getTraderJoeRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetTraderJoeRouter is a free data retrieval call binding the contract method 0x3a144e2d.
//
// Solidity: function getTraderJoeRouter() view returns(address traderJoeRouter_)
func (_TraderJoeRouterStorage *TraderJoeRouterStorageSession) GetTraderJoeRouter() (common.Address, error) {
	return _TraderJoeRouterStorage.Contract.GetTraderJoeRouter(&_TraderJoeRouterStorage.CallOpts)
}

// GetTraderJoeRouter is a free data retrieval call binding the contract method 0x3a144e2d.
//
// Solidity: function getTraderJoeRouter() view returns(address traderJoeRouter_)
func (_TraderJoeRouterStorage *TraderJoeRouterStorageCallerSession) GetTraderJoeRouter() (common.Address, error) {
	return _TraderJoeRouterStorage.Contract.GetTraderJoeRouter(&_TraderJoeRouterStorage.CallOpts)
}
