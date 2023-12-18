// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package arrays

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

// ArrayHelperMetaData contains all meta data concerning the ArrayHelper contract.
var ArrayHelperMetaData = &bind.MetaData{
	ABI: "[]",
}

// ArrayHelperABI is the input ABI used to generate the binding from.
// Deprecated: Use ArrayHelperMetaData.ABI instead.
var ArrayHelperABI = ArrayHelperMetaData.ABI

// ArrayHelper is an auto generated Go binding around an Ethereum contract.
type ArrayHelper struct {
	ArrayHelperCaller     // Read-only binding to the contract
	ArrayHelperTransactor // Write-only binding to the contract
	ArrayHelperFilterer   // Log filterer for contract events
}

// ArrayHelperCaller is an auto generated read-only Go binding around an Ethereum contract.
type ArrayHelperCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArrayHelperTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ArrayHelperTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArrayHelperFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ArrayHelperFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ArrayHelperSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ArrayHelperSession struct {
	Contract     *ArrayHelper      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ArrayHelperCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ArrayHelperCallerSession struct {
	Contract *ArrayHelperCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ArrayHelperTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ArrayHelperTransactorSession struct {
	Contract     *ArrayHelperTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ArrayHelperRaw is an auto generated low-level Go binding around an Ethereum contract.
type ArrayHelperRaw struct {
	Contract *ArrayHelper // Generic contract binding to access the raw methods on
}

// ArrayHelperCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ArrayHelperCallerRaw struct {
	Contract *ArrayHelperCaller // Generic read-only contract binding to access the raw methods on
}

// ArrayHelperTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ArrayHelperTransactorRaw struct {
	Contract *ArrayHelperTransactor // Generic write-only contract binding to access the raw methods on
}

// NewArrayHelper creates a new instance of ArrayHelper, bound to a specific deployed contract.
func NewArrayHelper(address common.Address, backend bind.ContractBackend) (*ArrayHelper, error) {
	contract, err := bindArrayHelper(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ArrayHelper{ArrayHelperCaller: ArrayHelperCaller{contract: contract}, ArrayHelperTransactor: ArrayHelperTransactor{contract: contract}, ArrayHelperFilterer: ArrayHelperFilterer{contract: contract}}, nil
}

// NewArrayHelperCaller creates a new read-only instance of ArrayHelper, bound to a specific deployed contract.
func NewArrayHelperCaller(address common.Address, caller bind.ContractCaller) (*ArrayHelperCaller, error) {
	contract, err := bindArrayHelper(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ArrayHelperCaller{contract: contract}, nil
}

// NewArrayHelperTransactor creates a new write-only instance of ArrayHelper, bound to a specific deployed contract.
func NewArrayHelperTransactor(address common.Address, transactor bind.ContractTransactor) (*ArrayHelperTransactor, error) {
	contract, err := bindArrayHelper(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ArrayHelperTransactor{contract: contract}, nil
}

// NewArrayHelperFilterer creates a new log filterer instance of ArrayHelper, bound to a specific deployed contract.
func NewArrayHelperFilterer(address common.Address, filterer bind.ContractFilterer) (*ArrayHelperFilterer, error) {
	contract, err := bindArrayHelper(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ArrayHelperFilterer{contract: contract}, nil
}

// bindArrayHelper binds a generic wrapper to an already deployed contract.
func bindArrayHelper(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ArrayHelperMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ArrayHelper *ArrayHelperRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ArrayHelper.Contract.ArrayHelperCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ArrayHelper *ArrayHelperRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArrayHelper.Contract.ArrayHelperTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ArrayHelper *ArrayHelperRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArrayHelper.Contract.ArrayHelperTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ArrayHelper *ArrayHelperCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ArrayHelper.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ArrayHelper *ArrayHelperTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ArrayHelper.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ArrayHelper *ArrayHelperTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ArrayHelper.Contract.contract.Transact(opts, method, params...)
}
