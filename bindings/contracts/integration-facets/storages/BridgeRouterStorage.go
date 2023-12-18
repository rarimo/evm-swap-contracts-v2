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

// BridgeRouterStorageMetaData contains all meta data concerning the BridgeRouterStorage contract.
var BridgeRouterStorageMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"BRIDGE_ROUTER_STORAGE_SLOT\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBridgeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"bridge_\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// BridgeRouterStorageABI is the input ABI used to generate the binding from.
// Deprecated: Use BridgeRouterStorageMetaData.ABI instead.
var BridgeRouterStorageABI = BridgeRouterStorageMetaData.ABI

// BridgeRouterStorage is an auto generated Go binding around an Ethereum contract.
type BridgeRouterStorage struct {
	BridgeRouterStorageCaller     // Read-only binding to the contract
	BridgeRouterStorageTransactor // Write-only binding to the contract
	BridgeRouterStorageFilterer   // Log filterer for contract events
}

// BridgeRouterStorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type BridgeRouterStorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeRouterStorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BridgeRouterStorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeRouterStorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BridgeRouterStorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeRouterStorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BridgeRouterStorageSession struct {
	Contract     *BridgeRouterStorage // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// BridgeRouterStorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BridgeRouterStorageCallerSession struct {
	Contract *BridgeRouterStorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// BridgeRouterStorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BridgeRouterStorageTransactorSession struct {
	Contract     *BridgeRouterStorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// BridgeRouterStorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type BridgeRouterStorageRaw struct {
	Contract *BridgeRouterStorage // Generic contract binding to access the raw methods on
}

// BridgeRouterStorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BridgeRouterStorageCallerRaw struct {
	Contract *BridgeRouterStorageCaller // Generic read-only contract binding to access the raw methods on
}

// BridgeRouterStorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BridgeRouterStorageTransactorRaw struct {
	Contract *BridgeRouterStorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridgeRouterStorage creates a new instance of BridgeRouterStorage, bound to a specific deployed contract.
func NewBridgeRouterStorage(address common.Address, backend bind.ContractBackend) (*BridgeRouterStorage, error) {
	contract, err := bindBridgeRouterStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BridgeRouterStorage{BridgeRouterStorageCaller: BridgeRouterStorageCaller{contract: contract}, BridgeRouterStorageTransactor: BridgeRouterStorageTransactor{contract: contract}, BridgeRouterStorageFilterer: BridgeRouterStorageFilterer{contract: contract}}, nil
}

// NewBridgeRouterStorageCaller creates a new read-only instance of BridgeRouterStorage, bound to a specific deployed contract.
func NewBridgeRouterStorageCaller(address common.Address, caller bind.ContractCaller) (*BridgeRouterStorageCaller, error) {
	contract, err := bindBridgeRouterStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeRouterStorageCaller{contract: contract}, nil
}

// NewBridgeRouterStorageTransactor creates a new write-only instance of BridgeRouterStorage, bound to a specific deployed contract.
func NewBridgeRouterStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*BridgeRouterStorageTransactor, error) {
	contract, err := bindBridgeRouterStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeRouterStorageTransactor{contract: contract}, nil
}

// NewBridgeRouterStorageFilterer creates a new log filterer instance of BridgeRouterStorage, bound to a specific deployed contract.
func NewBridgeRouterStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*BridgeRouterStorageFilterer, error) {
	contract, err := bindBridgeRouterStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BridgeRouterStorageFilterer{contract: contract}, nil
}

// bindBridgeRouterStorage binds a generic wrapper to an already deployed contract.
func bindBridgeRouterStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BridgeRouterStorageMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeRouterStorage *BridgeRouterStorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgeRouterStorage.Contract.BridgeRouterStorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeRouterStorage *BridgeRouterStorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeRouterStorage.Contract.BridgeRouterStorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeRouterStorage *BridgeRouterStorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeRouterStorage.Contract.BridgeRouterStorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeRouterStorage *BridgeRouterStorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgeRouterStorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeRouterStorage *BridgeRouterStorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeRouterStorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeRouterStorage *BridgeRouterStorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeRouterStorage.Contract.contract.Transact(opts, method, params...)
}

// BRIDGEROUTERSTORAGESLOT is a free data retrieval call binding the contract method 0x5e080a57.
//
// Solidity: function BRIDGE_ROUTER_STORAGE_SLOT() view returns(bytes32)
func (_BridgeRouterStorage *BridgeRouterStorageCaller) BRIDGEROUTERSTORAGESLOT(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BridgeRouterStorage.contract.Call(opts, &out, "BRIDGE_ROUTER_STORAGE_SLOT")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BRIDGEROUTERSTORAGESLOT is a free data retrieval call binding the contract method 0x5e080a57.
//
// Solidity: function BRIDGE_ROUTER_STORAGE_SLOT() view returns(bytes32)
func (_BridgeRouterStorage *BridgeRouterStorageSession) BRIDGEROUTERSTORAGESLOT() ([32]byte, error) {
	return _BridgeRouterStorage.Contract.BRIDGEROUTERSTORAGESLOT(&_BridgeRouterStorage.CallOpts)
}

// BRIDGEROUTERSTORAGESLOT is a free data retrieval call binding the contract method 0x5e080a57.
//
// Solidity: function BRIDGE_ROUTER_STORAGE_SLOT() view returns(bytes32)
func (_BridgeRouterStorage *BridgeRouterStorageCallerSession) BRIDGEROUTERSTORAGESLOT() ([32]byte, error) {
	return _BridgeRouterStorage.Contract.BRIDGEROUTERSTORAGESLOT(&_BridgeRouterStorage.CallOpts)
}

// GetBridgeAddress is a free data retrieval call binding the contract method 0xfb32c508.
//
// Solidity: function getBridgeAddress() view returns(address bridge_)
func (_BridgeRouterStorage *BridgeRouterStorageCaller) GetBridgeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BridgeRouterStorage.contract.Call(opts, &out, "getBridgeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetBridgeAddress is a free data retrieval call binding the contract method 0xfb32c508.
//
// Solidity: function getBridgeAddress() view returns(address bridge_)
func (_BridgeRouterStorage *BridgeRouterStorageSession) GetBridgeAddress() (common.Address, error) {
	return _BridgeRouterStorage.Contract.GetBridgeAddress(&_BridgeRouterStorage.CallOpts)
}

// GetBridgeAddress is a free data retrieval call binding the contract method 0xfb32c508.
//
// Solidity: function getBridgeAddress() view returns(address bridge_)
func (_BridgeRouterStorage *BridgeRouterStorageCallerSession) GetBridgeAddress() (common.Address, error) {
	return _BridgeRouterStorage.Contract.GetBridgeAddress(&_BridgeRouterStorage.CallOpts)
}
