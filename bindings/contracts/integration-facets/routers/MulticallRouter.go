// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package routers

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

// MulticallRouterMetaData contains all meta data concerning the MulticallRouter contract.
var MulticallRouterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"targets_\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values_\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"data_\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// MulticallRouterABI is the input ABI used to generate the binding from.
// Deprecated: Use MulticallRouterMetaData.ABI instead.
var MulticallRouterABI = MulticallRouterMetaData.ABI

// MulticallRouter is an auto generated Go binding around an Ethereum contract.
type MulticallRouter struct {
	MulticallRouterCaller     // Read-only binding to the contract
	MulticallRouterTransactor // Write-only binding to the contract
	MulticallRouterFilterer   // Log filterer for contract events
}

// MulticallRouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type MulticallRouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MulticallRouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MulticallRouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MulticallRouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MulticallRouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MulticallRouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MulticallRouterSession struct {
	Contract     *MulticallRouter  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MulticallRouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MulticallRouterCallerSession struct {
	Contract *MulticallRouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// MulticallRouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MulticallRouterTransactorSession struct {
	Contract     *MulticallRouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// MulticallRouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type MulticallRouterRaw struct {
	Contract *MulticallRouter // Generic contract binding to access the raw methods on
}

// MulticallRouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MulticallRouterCallerRaw struct {
	Contract *MulticallRouterCaller // Generic read-only contract binding to access the raw methods on
}

// MulticallRouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MulticallRouterTransactorRaw struct {
	Contract *MulticallRouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMulticallRouter creates a new instance of MulticallRouter, bound to a specific deployed contract.
func NewMulticallRouter(address common.Address, backend bind.ContractBackend) (*MulticallRouter, error) {
	contract, err := bindMulticallRouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MulticallRouter{MulticallRouterCaller: MulticallRouterCaller{contract: contract}, MulticallRouterTransactor: MulticallRouterTransactor{contract: contract}, MulticallRouterFilterer: MulticallRouterFilterer{contract: contract}}, nil
}

// NewMulticallRouterCaller creates a new read-only instance of MulticallRouter, bound to a specific deployed contract.
func NewMulticallRouterCaller(address common.Address, caller bind.ContractCaller) (*MulticallRouterCaller, error) {
	contract, err := bindMulticallRouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MulticallRouterCaller{contract: contract}, nil
}

// NewMulticallRouterTransactor creates a new write-only instance of MulticallRouter, bound to a specific deployed contract.
func NewMulticallRouterTransactor(address common.Address, transactor bind.ContractTransactor) (*MulticallRouterTransactor, error) {
	contract, err := bindMulticallRouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MulticallRouterTransactor{contract: contract}, nil
}

// NewMulticallRouterFilterer creates a new log filterer instance of MulticallRouter, bound to a specific deployed contract.
func NewMulticallRouterFilterer(address common.Address, filterer bind.ContractFilterer) (*MulticallRouterFilterer, error) {
	contract, err := bindMulticallRouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MulticallRouterFilterer{contract: contract}, nil
}

// bindMulticallRouter binds a generic wrapper to an already deployed contract.
func bindMulticallRouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MulticallRouterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MulticallRouter *MulticallRouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MulticallRouter.Contract.MulticallRouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MulticallRouter *MulticallRouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MulticallRouter.Contract.MulticallRouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MulticallRouter *MulticallRouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MulticallRouter.Contract.MulticallRouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MulticallRouter *MulticallRouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MulticallRouter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MulticallRouter *MulticallRouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MulticallRouter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MulticallRouter *MulticallRouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MulticallRouter.Contract.contract.Transact(opts, method, params...)
}

// Multicall is a paid mutator transaction binding the contract method 0x45f67ac7.
//
// Solidity: function multicall(address[] targets_, uint256[] values_, bytes[] data_) payable returns()
func (_MulticallRouter *MulticallRouterTransactor) Multicall(opts *bind.TransactOpts, targets_ []common.Address, values_ []*big.Int, data_ [][]byte) (*types.Transaction, error) {
	return _MulticallRouter.contract.Transact(opts, "multicall", targets_, values_, data_)
}

// Multicall is a paid mutator transaction binding the contract method 0x45f67ac7.
//
// Solidity: function multicall(address[] targets_, uint256[] values_, bytes[] data_) payable returns()
func (_MulticallRouter *MulticallRouterSession) Multicall(targets_ []common.Address, values_ []*big.Int, data_ [][]byte) (*types.Transaction, error) {
	return _MulticallRouter.Contract.Multicall(&_MulticallRouter.TransactOpts, targets_, values_, data_)
}

// Multicall is a paid mutator transaction binding the contract method 0x45f67ac7.
//
// Solidity: function multicall(address[] targets_, uint256[] values_, bytes[] data_) payable returns()
func (_MulticallRouter *MulticallRouterTransactorSession) Multicall(targets_ []common.Address, values_ []*big.Int, data_ [][]byte) (*types.Transaction, error) {
	return _MulticallRouter.Contract.Multicall(&_MulticallRouter.TransactOpts, targets_, values_, data_)
}
