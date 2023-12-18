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

// TransferRouterMetaData contains all meta data concerning the TransferRouter contract.
var TransferRouterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"MASTER_ROUTER_STORAGE_SLOT\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCallerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"caller_\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver_\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenIds_\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts_\",\"type\":\"uint256[]\"}],\"name\":\"transferERC1155\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"}],\"name\":\"transferERC20\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver_\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"nftIds_\",\"type\":\"uint256[]\"}],\"name\":\"transferERC721\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token_\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenIds_\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts_\",\"type\":\"uint256[]\"}],\"name\":\"transferFromERC1155\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"}],\"name\":\"transferFromERC20\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token_\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"nftIds_\",\"type\":\"uint256[]\"}],\"name\":\"transferFromERC721\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"}],\"name\":\"transferNative\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// TransferRouterABI is the input ABI used to generate the binding from.
// Deprecated: Use TransferRouterMetaData.ABI instead.
var TransferRouterABI = TransferRouterMetaData.ABI

// TransferRouter is an auto generated Go binding around an Ethereum contract.
type TransferRouter struct {
	TransferRouterCaller     // Read-only binding to the contract
	TransferRouterTransactor // Write-only binding to the contract
	TransferRouterFilterer   // Log filterer for contract events
}

// TransferRouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type TransferRouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransferRouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TransferRouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransferRouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TransferRouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransferRouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TransferRouterSession struct {
	Contract     *TransferRouter   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TransferRouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TransferRouterCallerSession struct {
	Contract *TransferRouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// TransferRouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TransferRouterTransactorSession struct {
	Contract     *TransferRouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// TransferRouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type TransferRouterRaw struct {
	Contract *TransferRouter // Generic contract binding to access the raw methods on
}

// TransferRouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TransferRouterCallerRaw struct {
	Contract *TransferRouterCaller // Generic read-only contract binding to access the raw methods on
}

// TransferRouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TransferRouterTransactorRaw struct {
	Contract *TransferRouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTransferRouter creates a new instance of TransferRouter, bound to a specific deployed contract.
func NewTransferRouter(address common.Address, backend bind.ContractBackend) (*TransferRouter, error) {
	contract, err := bindTransferRouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TransferRouter{TransferRouterCaller: TransferRouterCaller{contract: contract}, TransferRouterTransactor: TransferRouterTransactor{contract: contract}, TransferRouterFilterer: TransferRouterFilterer{contract: contract}}, nil
}

// NewTransferRouterCaller creates a new read-only instance of TransferRouter, bound to a specific deployed contract.
func NewTransferRouterCaller(address common.Address, caller bind.ContractCaller) (*TransferRouterCaller, error) {
	contract, err := bindTransferRouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TransferRouterCaller{contract: contract}, nil
}

// NewTransferRouterTransactor creates a new write-only instance of TransferRouter, bound to a specific deployed contract.
func NewTransferRouterTransactor(address common.Address, transactor bind.ContractTransactor) (*TransferRouterTransactor, error) {
	contract, err := bindTransferRouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TransferRouterTransactor{contract: contract}, nil
}

// NewTransferRouterFilterer creates a new log filterer instance of TransferRouter, bound to a specific deployed contract.
func NewTransferRouterFilterer(address common.Address, filterer bind.ContractFilterer) (*TransferRouterFilterer, error) {
	contract, err := bindTransferRouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TransferRouterFilterer{contract: contract}, nil
}

// bindTransferRouter binds a generic wrapper to an already deployed contract.
func bindTransferRouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TransferRouterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TransferRouter *TransferRouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TransferRouter.Contract.TransferRouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TransferRouter *TransferRouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TransferRouter.Contract.TransferRouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TransferRouter *TransferRouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TransferRouter.Contract.TransferRouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TransferRouter *TransferRouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TransferRouter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TransferRouter *TransferRouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TransferRouter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TransferRouter *TransferRouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TransferRouter.Contract.contract.Transact(opts, method, params...)
}

// MASTERROUTERSTORAGESLOT is a free data retrieval call binding the contract method 0x660fd3b0.
//
// Solidity: function MASTER_ROUTER_STORAGE_SLOT() view returns(bytes32)
func (_TransferRouter *TransferRouterCaller) MASTERROUTERSTORAGESLOT(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TransferRouter.contract.Call(opts, &out, "MASTER_ROUTER_STORAGE_SLOT")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MASTERROUTERSTORAGESLOT is a free data retrieval call binding the contract method 0x660fd3b0.
//
// Solidity: function MASTER_ROUTER_STORAGE_SLOT() view returns(bytes32)
func (_TransferRouter *TransferRouterSession) MASTERROUTERSTORAGESLOT() ([32]byte, error) {
	return _TransferRouter.Contract.MASTERROUTERSTORAGESLOT(&_TransferRouter.CallOpts)
}

// MASTERROUTERSTORAGESLOT is a free data retrieval call binding the contract method 0x660fd3b0.
//
// Solidity: function MASTER_ROUTER_STORAGE_SLOT() view returns(bytes32)
func (_TransferRouter *TransferRouterCallerSession) MASTERROUTERSTORAGESLOT() ([32]byte, error) {
	return _TransferRouter.Contract.MASTERROUTERSTORAGESLOT(&_TransferRouter.CallOpts)
}

// GetCallerAddress is a free data retrieval call binding the contract method 0x46b3353b.
//
// Solidity: function getCallerAddress() view returns(address caller_)
func (_TransferRouter *TransferRouterCaller) GetCallerAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TransferRouter.contract.Call(opts, &out, "getCallerAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetCallerAddress is a free data retrieval call binding the contract method 0x46b3353b.
//
// Solidity: function getCallerAddress() view returns(address caller_)
func (_TransferRouter *TransferRouterSession) GetCallerAddress() (common.Address, error) {
	return _TransferRouter.Contract.GetCallerAddress(&_TransferRouter.CallOpts)
}

// GetCallerAddress is a free data retrieval call binding the contract method 0x46b3353b.
//
// Solidity: function getCallerAddress() view returns(address caller_)
func (_TransferRouter *TransferRouterCallerSession) GetCallerAddress() (common.Address, error) {
	return _TransferRouter.Contract.GetCallerAddress(&_TransferRouter.CallOpts)
}

// TransferERC1155 is a paid mutator transaction binding the contract method 0x1cb56859.
//
// Solidity: function transferERC1155(address token_, address receiver_, uint256[] tokenIds_, uint256[] amounts_) payable returns()
func (_TransferRouter *TransferRouterTransactor) TransferERC1155(opts *bind.TransactOpts, token_ common.Address, receiver_ common.Address, tokenIds_ []*big.Int, amounts_ []*big.Int) (*types.Transaction, error) {
	return _TransferRouter.contract.Transact(opts, "transferERC1155", token_, receiver_, tokenIds_, amounts_)
}

// TransferERC1155 is a paid mutator transaction binding the contract method 0x1cb56859.
//
// Solidity: function transferERC1155(address token_, address receiver_, uint256[] tokenIds_, uint256[] amounts_) payable returns()
func (_TransferRouter *TransferRouterSession) TransferERC1155(token_ common.Address, receiver_ common.Address, tokenIds_ []*big.Int, amounts_ []*big.Int) (*types.Transaction, error) {
	return _TransferRouter.Contract.TransferERC1155(&_TransferRouter.TransactOpts, token_, receiver_, tokenIds_, amounts_)
}

// TransferERC1155 is a paid mutator transaction binding the contract method 0x1cb56859.
//
// Solidity: function transferERC1155(address token_, address receiver_, uint256[] tokenIds_, uint256[] amounts_) payable returns()
func (_TransferRouter *TransferRouterTransactorSession) TransferERC1155(token_ common.Address, receiver_ common.Address, tokenIds_ []*big.Int, amounts_ []*big.Int) (*types.Transaction, error) {
	return _TransferRouter.Contract.TransferERC1155(&_TransferRouter.TransactOpts, token_, receiver_, tokenIds_, amounts_)
}

// TransferERC20 is a paid mutator transaction binding the contract method 0x9db5dbe4.
//
// Solidity: function transferERC20(address token_, address receiver_, uint256 amount_) payable returns()
func (_TransferRouter *TransferRouterTransactor) TransferERC20(opts *bind.TransactOpts, token_ common.Address, receiver_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _TransferRouter.contract.Transact(opts, "transferERC20", token_, receiver_, amount_)
}

// TransferERC20 is a paid mutator transaction binding the contract method 0x9db5dbe4.
//
// Solidity: function transferERC20(address token_, address receiver_, uint256 amount_) payable returns()
func (_TransferRouter *TransferRouterSession) TransferERC20(token_ common.Address, receiver_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _TransferRouter.Contract.TransferERC20(&_TransferRouter.TransactOpts, token_, receiver_, amount_)
}

// TransferERC20 is a paid mutator transaction binding the contract method 0x9db5dbe4.
//
// Solidity: function transferERC20(address token_, address receiver_, uint256 amount_) payable returns()
func (_TransferRouter *TransferRouterTransactorSession) TransferERC20(token_ common.Address, receiver_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _TransferRouter.Contract.TransferERC20(&_TransferRouter.TransactOpts, token_, receiver_, amount_)
}

// TransferERC721 is a paid mutator transaction binding the contract method 0x8ca5e396.
//
// Solidity: function transferERC721(address token_, address receiver_, uint256[] nftIds_) payable returns()
func (_TransferRouter *TransferRouterTransactor) TransferERC721(opts *bind.TransactOpts, token_ common.Address, receiver_ common.Address, nftIds_ []*big.Int) (*types.Transaction, error) {
	return _TransferRouter.contract.Transact(opts, "transferERC721", token_, receiver_, nftIds_)
}

// TransferERC721 is a paid mutator transaction binding the contract method 0x8ca5e396.
//
// Solidity: function transferERC721(address token_, address receiver_, uint256[] nftIds_) payable returns()
func (_TransferRouter *TransferRouterSession) TransferERC721(token_ common.Address, receiver_ common.Address, nftIds_ []*big.Int) (*types.Transaction, error) {
	return _TransferRouter.Contract.TransferERC721(&_TransferRouter.TransactOpts, token_, receiver_, nftIds_)
}

// TransferERC721 is a paid mutator transaction binding the contract method 0x8ca5e396.
//
// Solidity: function transferERC721(address token_, address receiver_, uint256[] nftIds_) payable returns()
func (_TransferRouter *TransferRouterTransactorSession) TransferERC721(token_ common.Address, receiver_ common.Address, nftIds_ []*big.Int) (*types.Transaction, error) {
	return _TransferRouter.Contract.TransferERC721(&_TransferRouter.TransactOpts, token_, receiver_, nftIds_)
}

// TransferFromERC1155 is a paid mutator transaction binding the contract method 0x57e2c45f.
//
// Solidity: function transferFromERC1155(address token_, uint256[] tokenIds_, uint256[] amounts_) payable returns()
func (_TransferRouter *TransferRouterTransactor) TransferFromERC1155(opts *bind.TransactOpts, token_ common.Address, tokenIds_ []*big.Int, amounts_ []*big.Int) (*types.Transaction, error) {
	return _TransferRouter.contract.Transact(opts, "transferFromERC1155", token_, tokenIds_, amounts_)
}

// TransferFromERC1155 is a paid mutator transaction binding the contract method 0x57e2c45f.
//
// Solidity: function transferFromERC1155(address token_, uint256[] tokenIds_, uint256[] amounts_) payable returns()
func (_TransferRouter *TransferRouterSession) TransferFromERC1155(token_ common.Address, tokenIds_ []*big.Int, amounts_ []*big.Int) (*types.Transaction, error) {
	return _TransferRouter.Contract.TransferFromERC1155(&_TransferRouter.TransactOpts, token_, tokenIds_, amounts_)
}

// TransferFromERC1155 is a paid mutator transaction binding the contract method 0x57e2c45f.
//
// Solidity: function transferFromERC1155(address token_, uint256[] tokenIds_, uint256[] amounts_) payable returns()
func (_TransferRouter *TransferRouterTransactorSession) TransferFromERC1155(token_ common.Address, tokenIds_ []*big.Int, amounts_ []*big.Int) (*types.Transaction, error) {
	return _TransferRouter.Contract.TransferFromERC1155(&_TransferRouter.TransactOpts, token_, tokenIds_, amounts_)
}

// TransferFromERC20 is a paid mutator transaction binding the contract method 0xbd31ed1f.
//
// Solidity: function transferFromERC20(address token_, uint256 amount_) payable returns()
func (_TransferRouter *TransferRouterTransactor) TransferFromERC20(opts *bind.TransactOpts, token_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _TransferRouter.contract.Transact(opts, "transferFromERC20", token_, amount_)
}

// TransferFromERC20 is a paid mutator transaction binding the contract method 0xbd31ed1f.
//
// Solidity: function transferFromERC20(address token_, uint256 amount_) payable returns()
func (_TransferRouter *TransferRouterSession) TransferFromERC20(token_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _TransferRouter.Contract.TransferFromERC20(&_TransferRouter.TransactOpts, token_, amount_)
}

// TransferFromERC20 is a paid mutator transaction binding the contract method 0xbd31ed1f.
//
// Solidity: function transferFromERC20(address token_, uint256 amount_) payable returns()
func (_TransferRouter *TransferRouterTransactorSession) TransferFromERC20(token_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _TransferRouter.Contract.TransferFromERC20(&_TransferRouter.TransactOpts, token_, amount_)
}

// TransferFromERC721 is a paid mutator transaction binding the contract method 0x6ef8622b.
//
// Solidity: function transferFromERC721(address token_, uint256[] nftIds_) payable returns()
func (_TransferRouter *TransferRouterTransactor) TransferFromERC721(opts *bind.TransactOpts, token_ common.Address, nftIds_ []*big.Int) (*types.Transaction, error) {
	return _TransferRouter.contract.Transact(opts, "transferFromERC721", token_, nftIds_)
}

// TransferFromERC721 is a paid mutator transaction binding the contract method 0x6ef8622b.
//
// Solidity: function transferFromERC721(address token_, uint256[] nftIds_) payable returns()
func (_TransferRouter *TransferRouterSession) TransferFromERC721(token_ common.Address, nftIds_ []*big.Int) (*types.Transaction, error) {
	return _TransferRouter.Contract.TransferFromERC721(&_TransferRouter.TransactOpts, token_, nftIds_)
}

// TransferFromERC721 is a paid mutator transaction binding the contract method 0x6ef8622b.
//
// Solidity: function transferFromERC721(address token_, uint256[] nftIds_) payable returns()
func (_TransferRouter *TransferRouterTransactorSession) TransferFromERC721(token_ common.Address, nftIds_ []*big.Int) (*types.Transaction, error) {
	return _TransferRouter.Contract.TransferFromERC721(&_TransferRouter.TransactOpts, token_, nftIds_)
}

// TransferNative is a paid mutator transaction binding the contract method 0x7d2e90c2.
//
// Solidity: function transferNative(address receiver_, uint256 amount_) payable returns()
func (_TransferRouter *TransferRouterTransactor) TransferNative(opts *bind.TransactOpts, receiver_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _TransferRouter.contract.Transact(opts, "transferNative", receiver_, amount_)
}

// TransferNative is a paid mutator transaction binding the contract method 0x7d2e90c2.
//
// Solidity: function transferNative(address receiver_, uint256 amount_) payable returns()
func (_TransferRouter *TransferRouterSession) TransferNative(receiver_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _TransferRouter.Contract.TransferNative(&_TransferRouter.TransactOpts, receiver_, amount_)
}

// TransferNative is a paid mutator transaction binding the contract method 0x7d2e90c2.
//
// Solidity: function transferNative(address receiver_, uint256 amount_) payable returns()
func (_TransferRouter *TransferRouterTransactorSession) TransferNative(receiver_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _TransferRouter.Contract.TransferNative(&_TransferRouter.TransactOpts, receiver_, amount_)
}
