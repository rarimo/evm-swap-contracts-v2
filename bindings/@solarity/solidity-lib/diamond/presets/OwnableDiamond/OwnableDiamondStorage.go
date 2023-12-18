// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ownablediamond

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

// OwnableDiamondStorageMetaData contains all meta data concerning the OwnableDiamondStorage contract.
var OwnableDiamondStorageMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"OWNABLE_DIAMOND_STORAGE_SLOT\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// OwnableDiamondStorageABI is the input ABI used to generate the binding from.
// Deprecated: Use OwnableDiamondStorageMetaData.ABI instead.
var OwnableDiamondStorageABI = OwnableDiamondStorageMetaData.ABI

// OwnableDiamondStorage is an auto generated Go binding around an Ethereum contract.
type OwnableDiamondStorage struct {
	OwnableDiamondStorageCaller     // Read-only binding to the contract
	OwnableDiamondStorageTransactor // Write-only binding to the contract
	OwnableDiamondStorageFilterer   // Log filterer for contract events
}

// OwnableDiamondStorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwnableDiamondStorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableDiamondStorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnableDiamondStorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableDiamondStorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OwnableDiamondStorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableDiamondStorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwnableDiamondStorageSession struct {
	Contract     *OwnableDiamondStorage // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// OwnableDiamondStorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwnableDiamondStorageCallerSession struct {
	Contract *OwnableDiamondStorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// OwnableDiamondStorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwnableDiamondStorageTransactorSession struct {
	Contract     *OwnableDiamondStorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// OwnableDiamondStorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwnableDiamondStorageRaw struct {
	Contract *OwnableDiamondStorage // Generic contract binding to access the raw methods on
}

// OwnableDiamondStorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwnableDiamondStorageCallerRaw struct {
	Contract *OwnableDiamondStorageCaller // Generic read-only contract binding to access the raw methods on
}

// OwnableDiamondStorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwnableDiamondStorageTransactorRaw struct {
	Contract *OwnableDiamondStorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwnableDiamondStorage creates a new instance of OwnableDiamondStorage, bound to a specific deployed contract.
func NewOwnableDiamondStorage(address common.Address, backend bind.ContractBackend) (*OwnableDiamondStorage, error) {
	contract, err := bindOwnableDiamondStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OwnableDiamondStorage{OwnableDiamondStorageCaller: OwnableDiamondStorageCaller{contract: contract}, OwnableDiamondStorageTransactor: OwnableDiamondStorageTransactor{contract: contract}, OwnableDiamondStorageFilterer: OwnableDiamondStorageFilterer{contract: contract}}, nil
}

// NewOwnableDiamondStorageCaller creates a new read-only instance of OwnableDiamondStorage, bound to a specific deployed contract.
func NewOwnableDiamondStorageCaller(address common.Address, caller bind.ContractCaller) (*OwnableDiamondStorageCaller, error) {
	contract, err := bindOwnableDiamondStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableDiamondStorageCaller{contract: contract}, nil
}

// NewOwnableDiamondStorageTransactor creates a new write-only instance of OwnableDiamondStorage, bound to a specific deployed contract.
func NewOwnableDiamondStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnableDiamondStorageTransactor, error) {
	contract, err := bindOwnableDiamondStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableDiamondStorageTransactor{contract: contract}, nil
}

// NewOwnableDiamondStorageFilterer creates a new log filterer instance of OwnableDiamondStorage, bound to a specific deployed contract.
func NewOwnableDiamondStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*OwnableDiamondStorageFilterer, error) {
	contract, err := bindOwnableDiamondStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwnableDiamondStorageFilterer{contract: contract}, nil
}

// bindOwnableDiamondStorage binds a generic wrapper to an already deployed contract.
func bindOwnableDiamondStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OwnableDiamondStorageMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OwnableDiamondStorage *OwnableDiamondStorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OwnableDiamondStorage.Contract.OwnableDiamondStorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OwnableDiamondStorage *OwnableDiamondStorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OwnableDiamondStorage.Contract.OwnableDiamondStorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OwnableDiamondStorage *OwnableDiamondStorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OwnableDiamondStorage.Contract.OwnableDiamondStorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OwnableDiamondStorage *OwnableDiamondStorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OwnableDiamondStorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OwnableDiamondStorage *OwnableDiamondStorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OwnableDiamondStorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OwnableDiamondStorage *OwnableDiamondStorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OwnableDiamondStorage.Contract.contract.Transact(opts, method, params...)
}

// OWNABLEDIAMONDSTORAGESLOT is a free data retrieval call binding the contract method 0xe3e077ad.
//
// Solidity: function OWNABLE_DIAMOND_STORAGE_SLOT() view returns(bytes32)
func (_OwnableDiamondStorage *OwnableDiamondStorageCaller) OWNABLEDIAMONDSTORAGESLOT(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _OwnableDiamondStorage.contract.Call(opts, &out, "OWNABLE_DIAMOND_STORAGE_SLOT")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OWNABLEDIAMONDSTORAGESLOT is a free data retrieval call binding the contract method 0xe3e077ad.
//
// Solidity: function OWNABLE_DIAMOND_STORAGE_SLOT() view returns(bytes32)
func (_OwnableDiamondStorage *OwnableDiamondStorageSession) OWNABLEDIAMONDSTORAGESLOT() ([32]byte, error) {
	return _OwnableDiamondStorage.Contract.OWNABLEDIAMONDSTORAGESLOT(&_OwnableDiamondStorage.CallOpts)
}

// OWNABLEDIAMONDSTORAGESLOT is a free data retrieval call binding the contract method 0xe3e077ad.
//
// Solidity: function OWNABLE_DIAMOND_STORAGE_SLOT() view returns(bytes32)
func (_OwnableDiamondStorage *OwnableDiamondStorageCallerSession) OWNABLEDIAMONDSTORAGESLOT() ([32]byte, error) {
	return _OwnableDiamondStorage.Contract.OWNABLEDIAMONDSTORAGESLOT(&_OwnableDiamondStorage.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OwnableDiamondStorage *OwnableDiamondStorageCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OwnableDiamondStorage.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OwnableDiamondStorage *OwnableDiamondStorageSession) Owner() (common.Address, error) {
	return _OwnableDiamondStorage.Contract.Owner(&_OwnableDiamondStorage.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OwnableDiamondStorage *OwnableDiamondStorageCallerSession) Owner() (common.Address, error) {
	return _OwnableDiamondStorage.Contract.Owner(&_OwnableDiamondStorage.CallOpts)
}
