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

// ApproverMetaData contains all meta data concerning the Approver contract.
var ApproverMetaData = &bind.MetaData{
	ABI: "[]",
}

// ApproverABI is the input ABI used to generate the binding from.
// Deprecated: Use ApproverMetaData.ABI instead.
var ApproverABI = ApproverMetaData.ABI

// Approver is an auto generated Go binding around an Ethereum contract.
type Approver struct {
	ApproverCaller     // Read-only binding to the contract
	ApproverTransactor // Write-only binding to the contract
	ApproverFilterer   // Log filterer for contract events
}

// ApproverCaller is an auto generated read-only Go binding around an Ethereum contract.
type ApproverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApproverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ApproverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApproverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ApproverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApproverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ApproverSession struct {
	Contract     *Approver         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ApproverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ApproverCallerSession struct {
	Contract *ApproverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ApproverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ApproverTransactorSession struct {
	Contract     *ApproverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ApproverRaw is an auto generated low-level Go binding around an Ethereum contract.
type ApproverRaw struct {
	Contract *Approver // Generic contract binding to access the raw methods on
}

// ApproverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ApproverCallerRaw struct {
	Contract *ApproverCaller // Generic read-only contract binding to access the raw methods on
}

// ApproverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ApproverTransactorRaw struct {
	Contract *ApproverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewApprover creates a new instance of Approver, bound to a specific deployed contract.
func NewApprover(address common.Address, backend bind.ContractBackend) (*Approver, error) {
	contract, err := bindApprover(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Approver{ApproverCaller: ApproverCaller{contract: contract}, ApproverTransactor: ApproverTransactor{contract: contract}, ApproverFilterer: ApproverFilterer{contract: contract}}, nil
}

// NewApproverCaller creates a new read-only instance of Approver, bound to a specific deployed contract.
func NewApproverCaller(address common.Address, caller bind.ContractCaller) (*ApproverCaller, error) {
	contract, err := bindApprover(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ApproverCaller{contract: contract}, nil
}

// NewApproverTransactor creates a new write-only instance of Approver, bound to a specific deployed contract.
func NewApproverTransactor(address common.Address, transactor bind.ContractTransactor) (*ApproverTransactor, error) {
	contract, err := bindApprover(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ApproverTransactor{contract: contract}, nil
}

// NewApproverFilterer creates a new log filterer instance of Approver, bound to a specific deployed contract.
func NewApproverFilterer(address common.Address, filterer bind.ContractFilterer) (*ApproverFilterer, error) {
	contract, err := bindApprover(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ApproverFilterer{contract: contract}, nil
}

// bindApprover binds a generic wrapper to an already deployed contract.
func bindApprover(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ApproverMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Approver *ApproverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Approver.Contract.ApproverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Approver *ApproverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Approver.Contract.ApproverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Approver *ApproverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Approver.Contract.ApproverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Approver *ApproverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Approver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Approver *ApproverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Approver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Approver *ApproverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Approver.Contract.contract.Transact(opts, method, params...)
}
