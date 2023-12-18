// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package masterfacet

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

// MasterRouterStorageMetaData contains all meta data concerning the MasterRouterStorage contract.
var MasterRouterStorageMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"MASTER_ROUTER_STORAGE_SLOT\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCallerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"caller_\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// MasterRouterStorageABI is the input ABI used to generate the binding from.
// Deprecated: Use MasterRouterStorageMetaData.ABI instead.
var MasterRouterStorageABI = MasterRouterStorageMetaData.ABI

// MasterRouterStorage is an auto generated Go binding around an Ethereum contract.
type MasterRouterStorage struct {
	MasterRouterStorageCaller     // Read-only binding to the contract
	MasterRouterStorageTransactor // Write-only binding to the contract
	MasterRouterStorageFilterer   // Log filterer for contract events
}

// MasterRouterStorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type MasterRouterStorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MasterRouterStorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MasterRouterStorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MasterRouterStorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MasterRouterStorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MasterRouterStorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MasterRouterStorageSession struct {
	Contract     *MasterRouterStorage // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// MasterRouterStorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MasterRouterStorageCallerSession struct {
	Contract *MasterRouterStorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// MasterRouterStorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MasterRouterStorageTransactorSession struct {
	Contract     *MasterRouterStorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// MasterRouterStorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type MasterRouterStorageRaw struct {
	Contract *MasterRouterStorage // Generic contract binding to access the raw methods on
}

// MasterRouterStorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MasterRouterStorageCallerRaw struct {
	Contract *MasterRouterStorageCaller // Generic read-only contract binding to access the raw methods on
}

// MasterRouterStorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MasterRouterStorageTransactorRaw struct {
	Contract *MasterRouterStorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMasterRouterStorage creates a new instance of MasterRouterStorage, bound to a specific deployed contract.
func NewMasterRouterStorage(address common.Address, backend bind.ContractBackend) (*MasterRouterStorage, error) {
	contract, err := bindMasterRouterStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MasterRouterStorage{MasterRouterStorageCaller: MasterRouterStorageCaller{contract: contract}, MasterRouterStorageTransactor: MasterRouterStorageTransactor{contract: contract}, MasterRouterStorageFilterer: MasterRouterStorageFilterer{contract: contract}}, nil
}

// NewMasterRouterStorageCaller creates a new read-only instance of MasterRouterStorage, bound to a specific deployed contract.
func NewMasterRouterStorageCaller(address common.Address, caller bind.ContractCaller) (*MasterRouterStorageCaller, error) {
	contract, err := bindMasterRouterStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MasterRouterStorageCaller{contract: contract}, nil
}

// NewMasterRouterStorageTransactor creates a new write-only instance of MasterRouterStorage, bound to a specific deployed contract.
func NewMasterRouterStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*MasterRouterStorageTransactor, error) {
	contract, err := bindMasterRouterStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MasterRouterStorageTransactor{contract: contract}, nil
}

// NewMasterRouterStorageFilterer creates a new log filterer instance of MasterRouterStorage, bound to a specific deployed contract.
func NewMasterRouterStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*MasterRouterStorageFilterer, error) {
	contract, err := bindMasterRouterStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MasterRouterStorageFilterer{contract: contract}, nil
}

// bindMasterRouterStorage binds a generic wrapper to an already deployed contract.
func bindMasterRouterStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MasterRouterStorageMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MasterRouterStorage *MasterRouterStorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MasterRouterStorage.Contract.MasterRouterStorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MasterRouterStorage *MasterRouterStorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MasterRouterStorage.Contract.MasterRouterStorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MasterRouterStorage *MasterRouterStorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MasterRouterStorage.Contract.MasterRouterStorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MasterRouterStorage *MasterRouterStorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MasterRouterStorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MasterRouterStorage *MasterRouterStorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MasterRouterStorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MasterRouterStorage *MasterRouterStorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MasterRouterStorage.Contract.contract.Transact(opts, method, params...)
}

// MASTERROUTERSTORAGESLOT is a free data retrieval call binding the contract method 0x660fd3b0.
//
// Solidity: function MASTER_ROUTER_STORAGE_SLOT() view returns(bytes32)
func (_MasterRouterStorage *MasterRouterStorageCaller) MASTERROUTERSTORAGESLOT(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MasterRouterStorage.contract.Call(opts, &out, "MASTER_ROUTER_STORAGE_SLOT")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MASTERROUTERSTORAGESLOT is a free data retrieval call binding the contract method 0x660fd3b0.
//
// Solidity: function MASTER_ROUTER_STORAGE_SLOT() view returns(bytes32)
func (_MasterRouterStorage *MasterRouterStorageSession) MASTERROUTERSTORAGESLOT() ([32]byte, error) {
	return _MasterRouterStorage.Contract.MASTERROUTERSTORAGESLOT(&_MasterRouterStorage.CallOpts)
}

// MASTERROUTERSTORAGESLOT is a free data retrieval call binding the contract method 0x660fd3b0.
//
// Solidity: function MASTER_ROUTER_STORAGE_SLOT() view returns(bytes32)
func (_MasterRouterStorage *MasterRouterStorageCallerSession) MASTERROUTERSTORAGESLOT() ([32]byte, error) {
	return _MasterRouterStorage.Contract.MASTERROUTERSTORAGESLOT(&_MasterRouterStorage.CallOpts)
}

// GetCallerAddress is a free data retrieval call binding the contract method 0x46b3353b.
//
// Solidity: function getCallerAddress() view returns(address caller_)
func (_MasterRouterStorage *MasterRouterStorageCaller) GetCallerAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MasterRouterStorage.contract.Call(opts, &out, "getCallerAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetCallerAddress is a free data retrieval call binding the contract method 0x46b3353b.
//
// Solidity: function getCallerAddress() view returns(address caller_)
func (_MasterRouterStorage *MasterRouterStorageSession) GetCallerAddress() (common.Address, error) {
	return _MasterRouterStorage.Contract.GetCallerAddress(&_MasterRouterStorage.CallOpts)
}

// GetCallerAddress is a free data retrieval call binding the contract method 0x46b3353b.
//
// Solidity: function getCallerAddress() view returns(address caller_)
func (_MasterRouterStorage *MasterRouterStorageCallerSession) GetCallerAddress() (common.Address, error) {
	return _MasterRouterStorage.Contract.GetCallerAddress(&_MasterRouterStorage.CallOpts)
}
