// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// SwapDiamondStorageMetaData contains all meta data concerning the SwapDiamondStorage contract.
var SwapDiamondStorageMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"SWAP_DIAMOND_STORAGE_SLOT\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"selector_\",\"type\":\"bytes4\"}],\"name\":\"getSelectorType\",\"outputs\":[{\"internalType\":\"enumSwapDiamondStorage.SelectorType\",\"name\":\"selectorType_\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// SwapDiamondStorageABI is the input ABI used to generate the binding from.
// Deprecated: Use SwapDiamondStorageMetaData.ABI instead.
var SwapDiamondStorageABI = SwapDiamondStorageMetaData.ABI

// SwapDiamondStorage is an auto generated Go binding around an Ethereum contract.
type SwapDiamondStorage struct {
	SwapDiamondStorageCaller     // Read-only binding to the contract
	SwapDiamondStorageTransactor // Write-only binding to the contract
	SwapDiamondStorageFilterer   // Log filterer for contract events
}

// SwapDiamondStorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type SwapDiamondStorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapDiamondStorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SwapDiamondStorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapDiamondStorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SwapDiamondStorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapDiamondStorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SwapDiamondStorageSession struct {
	Contract     *SwapDiamondStorage // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SwapDiamondStorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SwapDiamondStorageCallerSession struct {
	Contract *SwapDiamondStorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// SwapDiamondStorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SwapDiamondStorageTransactorSession struct {
	Contract     *SwapDiamondStorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// SwapDiamondStorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type SwapDiamondStorageRaw struct {
	Contract *SwapDiamondStorage // Generic contract binding to access the raw methods on
}

// SwapDiamondStorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SwapDiamondStorageCallerRaw struct {
	Contract *SwapDiamondStorageCaller // Generic read-only contract binding to access the raw methods on
}

// SwapDiamondStorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SwapDiamondStorageTransactorRaw struct {
	Contract *SwapDiamondStorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSwapDiamondStorage creates a new instance of SwapDiamondStorage, bound to a specific deployed contract.
func NewSwapDiamondStorage(address common.Address, backend bind.ContractBackend) (*SwapDiamondStorage, error) {
	contract, err := bindSwapDiamondStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SwapDiamondStorage{SwapDiamondStorageCaller: SwapDiamondStorageCaller{contract: contract}, SwapDiamondStorageTransactor: SwapDiamondStorageTransactor{contract: contract}, SwapDiamondStorageFilterer: SwapDiamondStorageFilterer{contract: contract}}, nil
}

// NewSwapDiamondStorageCaller creates a new read-only instance of SwapDiamondStorage, bound to a specific deployed contract.
func NewSwapDiamondStorageCaller(address common.Address, caller bind.ContractCaller) (*SwapDiamondStorageCaller, error) {
	contract, err := bindSwapDiamondStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SwapDiamondStorageCaller{contract: contract}, nil
}

// NewSwapDiamondStorageTransactor creates a new write-only instance of SwapDiamondStorage, bound to a specific deployed contract.
func NewSwapDiamondStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*SwapDiamondStorageTransactor, error) {
	contract, err := bindSwapDiamondStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SwapDiamondStorageTransactor{contract: contract}, nil
}

// NewSwapDiamondStorageFilterer creates a new log filterer instance of SwapDiamondStorage, bound to a specific deployed contract.
func NewSwapDiamondStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*SwapDiamondStorageFilterer, error) {
	contract, err := bindSwapDiamondStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SwapDiamondStorageFilterer{contract: contract}, nil
}

// bindSwapDiamondStorage binds a generic wrapper to an already deployed contract.
func bindSwapDiamondStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SwapDiamondStorageMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SwapDiamondStorage *SwapDiamondStorageRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SwapDiamondStorage.Contract.SwapDiamondStorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SwapDiamondStorage *SwapDiamondStorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SwapDiamondStorage.Contract.SwapDiamondStorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SwapDiamondStorage *SwapDiamondStorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SwapDiamondStorage.Contract.SwapDiamondStorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SwapDiamondStorage *SwapDiamondStorageCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SwapDiamondStorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SwapDiamondStorage *SwapDiamondStorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SwapDiamondStorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SwapDiamondStorage *SwapDiamondStorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SwapDiamondStorage.Contract.contract.Transact(opts, method, params...)
}

// SWAPDIAMONDSTORAGESLOT is a free data retrieval call binding the contract method 0xaaadd377.
//
// Solidity: function SWAP_DIAMOND_STORAGE_SLOT() view returns(bytes32)
func (_SwapDiamondStorage *SwapDiamondStorageCaller) SWAPDIAMONDSTORAGESLOT(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _SwapDiamondStorage.contract.Call(opts, &out, "SWAP_DIAMOND_STORAGE_SLOT")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SWAPDIAMONDSTORAGESLOT is a free data retrieval call binding the contract method 0xaaadd377.
//
// Solidity: function SWAP_DIAMOND_STORAGE_SLOT() view returns(bytes32)
func (_SwapDiamondStorage *SwapDiamondStorageSession) SWAPDIAMONDSTORAGESLOT() ([32]byte, error) {
	return _SwapDiamondStorage.Contract.SWAPDIAMONDSTORAGESLOT(&_SwapDiamondStorage.CallOpts)
}

// SWAPDIAMONDSTORAGESLOT is a free data retrieval call binding the contract method 0xaaadd377.
//
// Solidity: function SWAP_DIAMOND_STORAGE_SLOT() view returns(bytes32)
func (_SwapDiamondStorage *SwapDiamondStorageCallerSession) SWAPDIAMONDSTORAGESLOT() ([32]byte, error) {
	return _SwapDiamondStorage.Contract.SWAPDIAMONDSTORAGESLOT(&_SwapDiamondStorage.CallOpts)
}

// GetSelectorType is a free data retrieval call binding the contract method 0xad5403eb.
//
// Solidity: function getSelectorType(bytes4 selector_) view returns(uint8 selectorType_)
func (_SwapDiamondStorage *SwapDiamondStorageCaller) GetSelectorType(opts *bind.CallOpts, selector_ [4]byte) (uint8, error) {
	var out []interface{}
	err := _SwapDiamondStorage.contract.Call(opts, &out, "getSelectorType", selector_)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetSelectorType is a free data retrieval call binding the contract method 0xad5403eb.
//
// Solidity: function getSelectorType(bytes4 selector_) view returns(uint8 selectorType_)
func (_SwapDiamondStorage *SwapDiamondStorageSession) GetSelectorType(selector_ [4]byte) (uint8, error) {
	return _SwapDiamondStorage.Contract.GetSelectorType(&_SwapDiamondStorage.CallOpts, selector_)
}

// GetSelectorType is a free data retrieval call binding the contract method 0xad5403eb.
//
// Solidity: function getSelectorType(bytes4 selector_) view returns(uint8 selectorType_)
func (_SwapDiamondStorage *SwapDiamondStorageCallerSession) GetSelectorType(selector_ [4]byte) (uint8, error) {
	return _SwapDiamondStorage.Contract.GetSelectorType(&_SwapDiamondStorage.CallOpts, selector_)
}
