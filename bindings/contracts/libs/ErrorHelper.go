// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package libs

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

// ErrorHelperMetaData contains all meta data concerning the ErrorHelper contract.
var ErrorHelperMetaData = &bind.MetaData{
	ABI: "[]",
}

// ErrorHelperABI is the input ABI used to generate the binding from.
// Deprecated: Use ErrorHelperMetaData.ABI instead.
var ErrorHelperABI = ErrorHelperMetaData.ABI

// ErrorHelper is an auto generated Go binding around an Ethereum contract.
type ErrorHelper struct {
	ErrorHelperCaller     // Read-only binding to the contract
	ErrorHelperTransactor // Write-only binding to the contract
	ErrorHelperFilterer   // Log filterer for contract events
}

// ErrorHelperCaller is an auto generated read-only Go binding around an Ethereum contract.
type ErrorHelperCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ErrorHelperTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ErrorHelperTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ErrorHelperFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ErrorHelperFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ErrorHelperSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ErrorHelperSession struct {
	Contract     *ErrorHelper      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ErrorHelperCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ErrorHelperCallerSession struct {
	Contract *ErrorHelperCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ErrorHelperTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ErrorHelperTransactorSession struct {
	Contract     *ErrorHelperTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ErrorHelperRaw is an auto generated low-level Go binding around an Ethereum contract.
type ErrorHelperRaw struct {
	Contract *ErrorHelper // Generic contract binding to access the raw methods on
}

// ErrorHelperCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ErrorHelperCallerRaw struct {
	Contract *ErrorHelperCaller // Generic read-only contract binding to access the raw methods on
}

// ErrorHelperTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ErrorHelperTransactorRaw struct {
	Contract *ErrorHelperTransactor // Generic write-only contract binding to access the raw methods on
}

// NewErrorHelper creates a new instance of ErrorHelper, bound to a specific deployed contract.
func NewErrorHelper(address common.Address, backend bind.ContractBackend) (*ErrorHelper, error) {
	contract, err := bindErrorHelper(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ErrorHelper{ErrorHelperCaller: ErrorHelperCaller{contract: contract}, ErrorHelperTransactor: ErrorHelperTransactor{contract: contract}, ErrorHelperFilterer: ErrorHelperFilterer{contract: contract}}, nil
}

// NewErrorHelperCaller creates a new read-only instance of ErrorHelper, bound to a specific deployed contract.
func NewErrorHelperCaller(address common.Address, caller bind.ContractCaller) (*ErrorHelperCaller, error) {
	contract, err := bindErrorHelper(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ErrorHelperCaller{contract: contract}, nil
}

// NewErrorHelperTransactor creates a new write-only instance of ErrorHelper, bound to a specific deployed contract.
func NewErrorHelperTransactor(address common.Address, transactor bind.ContractTransactor) (*ErrorHelperTransactor, error) {
	contract, err := bindErrorHelper(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ErrorHelperTransactor{contract: contract}, nil
}

// NewErrorHelperFilterer creates a new log filterer instance of ErrorHelper, bound to a specific deployed contract.
func NewErrorHelperFilterer(address common.Address, filterer bind.ContractFilterer) (*ErrorHelperFilterer, error) {
	contract, err := bindErrorHelper(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ErrorHelperFilterer{contract: contract}, nil
}

// bindErrorHelper binds a generic wrapper to an already deployed contract.
func bindErrorHelper(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ErrorHelperMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ErrorHelper *ErrorHelperRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ErrorHelper.Contract.ErrorHelperCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ErrorHelper *ErrorHelperRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ErrorHelper.Contract.ErrorHelperTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ErrorHelper *ErrorHelperRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ErrorHelper.Contract.ErrorHelperTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ErrorHelper *ErrorHelperCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ErrorHelper.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ErrorHelper *ErrorHelperTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ErrorHelper.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ErrorHelper *ErrorHelperTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ErrorHelper.Contract.contract.Transact(opts, method, params...)
}
