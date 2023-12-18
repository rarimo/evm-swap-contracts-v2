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

// WrapRouterMetaData contains all meta data concerning the WrapRouter contract.
var WrapRouterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"MASTER_ROUTER_STORAGE_SLOT\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OWNABLE_DIAMOND_STORAGE_SLOT\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WRAP_ROUTER_STORAGE_SLOT\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCallerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"caller_\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getWrappedNativeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"wrappedNative_\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"wrappedNative_\",\"type\":\"address\"}],\"name\":\"setWrappedNativeAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver_\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenIds_\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts_\",\"type\":\"uint256[]\"}],\"name\":\"transferERC1155\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"}],\"name\":\"transferERC20\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver_\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"nftIds_\",\"type\":\"uint256[]\"}],\"name\":\"transferERC721\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token_\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenIds_\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts_\",\"type\":\"uint256[]\"}],\"name\":\"transferFromERC1155\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"}],\"name\":\"transferFromERC20\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token_\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"nftIds_\",\"type\":\"uint256[]\"}],\"name\":\"transferFromERC721\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"}],\"name\":\"transferNative\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"}],\"name\":\"unwrap\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount_\",\"type\":\"uint256\"}],\"name\":\"wrap\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// WrapRouterABI is the input ABI used to generate the binding from.
// Deprecated: Use WrapRouterMetaData.ABI instead.
var WrapRouterABI = WrapRouterMetaData.ABI

// WrapRouter is an auto generated Go binding around an Ethereum contract.
type WrapRouter struct {
	WrapRouterCaller     // Read-only binding to the contract
	WrapRouterTransactor // Write-only binding to the contract
	WrapRouterFilterer   // Log filterer for contract events
}

// WrapRouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type WrapRouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WrapRouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WrapRouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WrapRouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WrapRouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WrapRouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WrapRouterSession struct {
	Contract     *WrapRouter       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WrapRouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WrapRouterCallerSession struct {
	Contract *WrapRouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// WrapRouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WrapRouterTransactorSession struct {
	Contract     *WrapRouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// WrapRouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type WrapRouterRaw struct {
	Contract *WrapRouter // Generic contract binding to access the raw methods on
}

// WrapRouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WrapRouterCallerRaw struct {
	Contract *WrapRouterCaller // Generic read-only contract binding to access the raw methods on
}

// WrapRouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WrapRouterTransactorRaw struct {
	Contract *WrapRouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWrapRouter creates a new instance of WrapRouter, bound to a specific deployed contract.
func NewWrapRouter(address common.Address, backend bind.ContractBackend) (*WrapRouter, error) {
	contract, err := bindWrapRouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WrapRouter{WrapRouterCaller: WrapRouterCaller{contract: contract}, WrapRouterTransactor: WrapRouterTransactor{contract: contract}, WrapRouterFilterer: WrapRouterFilterer{contract: contract}}, nil
}

// NewWrapRouterCaller creates a new read-only instance of WrapRouter, bound to a specific deployed contract.
func NewWrapRouterCaller(address common.Address, caller bind.ContractCaller) (*WrapRouterCaller, error) {
	contract, err := bindWrapRouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WrapRouterCaller{contract: contract}, nil
}

// NewWrapRouterTransactor creates a new write-only instance of WrapRouter, bound to a specific deployed contract.
func NewWrapRouterTransactor(address common.Address, transactor bind.ContractTransactor) (*WrapRouterTransactor, error) {
	contract, err := bindWrapRouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WrapRouterTransactor{contract: contract}, nil
}

// NewWrapRouterFilterer creates a new log filterer instance of WrapRouter, bound to a specific deployed contract.
func NewWrapRouterFilterer(address common.Address, filterer bind.ContractFilterer) (*WrapRouterFilterer, error) {
	contract, err := bindWrapRouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WrapRouterFilterer{contract: contract}, nil
}

// bindWrapRouter binds a generic wrapper to an already deployed contract.
func bindWrapRouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := WrapRouterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WrapRouter *WrapRouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WrapRouter.Contract.WrapRouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WrapRouter *WrapRouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WrapRouter.Contract.WrapRouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WrapRouter *WrapRouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WrapRouter.Contract.WrapRouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WrapRouter *WrapRouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WrapRouter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WrapRouter *WrapRouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WrapRouter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WrapRouter *WrapRouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WrapRouter.Contract.contract.Transact(opts, method, params...)
}

// MASTERROUTERSTORAGESLOT is a free data retrieval call binding the contract method 0x660fd3b0.
//
// Solidity: function MASTER_ROUTER_STORAGE_SLOT() view returns(bytes32)
func (_WrapRouter *WrapRouterCaller) MASTERROUTERSTORAGESLOT(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _WrapRouter.contract.Call(opts, &out, "MASTER_ROUTER_STORAGE_SLOT")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MASTERROUTERSTORAGESLOT is a free data retrieval call binding the contract method 0x660fd3b0.
//
// Solidity: function MASTER_ROUTER_STORAGE_SLOT() view returns(bytes32)
func (_WrapRouter *WrapRouterSession) MASTERROUTERSTORAGESLOT() ([32]byte, error) {
	return _WrapRouter.Contract.MASTERROUTERSTORAGESLOT(&_WrapRouter.CallOpts)
}

// MASTERROUTERSTORAGESLOT is a free data retrieval call binding the contract method 0x660fd3b0.
//
// Solidity: function MASTER_ROUTER_STORAGE_SLOT() view returns(bytes32)
func (_WrapRouter *WrapRouterCallerSession) MASTERROUTERSTORAGESLOT() ([32]byte, error) {
	return _WrapRouter.Contract.MASTERROUTERSTORAGESLOT(&_WrapRouter.CallOpts)
}

// OWNABLEDIAMONDSTORAGESLOT is a free data retrieval call binding the contract method 0xe3e077ad.
//
// Solidity: function OWNABLE_DIAMOND_STORAGE_SLOT() view returns(bytes32)
func (_WrapRouter *WrapRouterCaller) OWNABLEDIAMONDSTORAGESLOT(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _WrapRouter.contract.Call(opts, &out, "OWNABLE_DIAMOND_STORAGE_SLOT")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OWNABLEDIAMONDSTORAGESLOT is a free data retrieval call binding the contract method 0xe3e077ad.
//
// Solidity: function OWNABLE_DIAMOND_STORAGE_SLOT() view returns(bytes32)
func (_WrapRouter *WrapRouterSession) OWNABLEDIAMONDSTORAGESLOT() ([32]byte, error) {
	return _WrapRouter.Contract.OWNABLEDIAMONDSTORAGESLOT(&_WrapRouter.CallOpts)
}

// OWNABLEDIAMONDSTORAGESLOT is a free data retrieval call binding the contract method 0xe3e077ad.
//
// Solidity: function OWNABLE_DIAMOND_STORAGE_SLOT() view returns(bytes32)
func (_WrapRouter *WrapRouterCallerSession) OWNABLEDIAMONDSTORAGESLOT() ([32]byte, error) {
	return _WrapRouter.Contract.OWNABLEDIAMONDSTORAGESLOT(&_WrapRouter.CallOpts)
}

// WRAPROUTERSTORAGESLOT is a free data retrieval call binding the contract method 0x53a877c0.
//
// Solidity: function WRAP_ROUTER_STORAGE_SLOT() view returns(bytes32)
func (_WrapRouter *WrapRouterCaller) WRAPROUTERSTORAGESLOT(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _WrapRouter.contract.Call(opts, &out, "WRAP_ROUTER_STORAGE_SLOT")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// WRAPROUTERSTORAGESLOT is a free data retrieval call binding the contract method 0x53a877c0.
//
// Solidity: function WRAP_ROUTER_STORAGE_SLOT() view returns(bytes32)
func (_WrapRouter *WrapRouterSession) WRAPROUTERSTORAGESLOT() ([32]byte, error) {
	return _WrapRouter.Contract.WRAPROUTERSTORAGESLOT(&_WrapRouter.CallOpts)
}

// WRAPROUTERSTORAGESLOT is a free data retrieval call binding the contract method 0x53a877c0.
//
// Solidity: function WRAP_ROUTER_STORAGE_SLOT() view returns(bytes32)
func (_WrapRouter *WrapRouterCallerSession) WRAPROUTERSTORAGESLOT() ([32]byte, error) {
	return _WrapRouter.Contract.WRAPROUTERSTORAGESLOT(&_WrapRouter.CallOpts)
}

// GetCallerAddress is a free data retrieval call binding the contract method 0x46b3353b.
//
// Solidity: function getCallerAddress() view returns(address caller_)
func (_WrapRouter *WrapRouterCaller) GetCallerAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WrapRouter.contract.Call(opts, &out, "getCallerAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetCallerAddress is a free data retrieval call binding the contract method 0x46b3353b.
//
// Solidity: function getCallerAddress() view returns(address caller_)
func (_WrapRouter *WrapRouterSession) GetCallerAddress() (common.Address, error) {
	return _WrapRouter.Contract.GetCallerAddress(&_WrapRouter.CallOpts)
}

// GetCallerAddress is a free data retrieval call binding the contract method 0x46b3353b.
//
// Solidity: function getCallerAddress() view returns(address caller_)
func (_WrapRouter *WrapRouterCallerSession) GetCallerAddress() (common.Address, error) {
	return _WrapRouter.Contract.GetCallerAddress(&_WrapRouter.CallOpts)
}

// GetWrappedNativeAddress is a free data retrieval call binding the contract method 0x4e07eb5e.
//
// Solidity: function getWrappedNativeAddress() view returns(address wrappedNative_)
func (_WrapRouter *WrapRouterCaller) GetWrappedNativeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WrapRouter.contract.Call(opts, &out, "getWrappedNativeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetWrappedNativeAddress is a free data retrieval call binding the contract method 0x4e07eb5e.
//
// Solidity: function getWrappedNativeAddress() view returns(address wrappedNative_)
func (_WrapRouter *WrapRouterSession) GetWrappedNativeAddress() (common.Address, error) {
	return _WrapRouter.Contract.GetWrappedNativeAddress(&_WrapRouter.CallOpts)
}

// GetWrappedNativeAddress is a free data retrieval call binding the contract method 0x4e07eb5e.
//
// Solidity: function getWrappedNativeAddress() view returns(address wrappedNative_)
func (_WrapRouter *WrapRouterCallerSession) GetWrappedNativeAddress() (common.Address, error) {
	return _WrapRouter.Contract.GetWrappedNativeAddress(&_WrapRouter.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WrapRouter *WrapRouterCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WrapRouter.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WrapRouter *WrapRouterSession) Owner() (common.Address, error) {
	return _WrapRouter.Contract.Owner(&_WrapRouter.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WrapRouter *WrapRouterCallerSession) Owner() (common.Address, error) {
	return _WrapRouter.Contract.Owner(&_WrapRouter.CallOpts)
}

// SetWrappedNativeAddress is a paid mutator transaction binding the contract method 0x9ad21e01.
//
// Solidity: function setWrappedNativeAddress(address wrappedNative_) returns()
func (_WrapRouter *WrapRouterTransactor) SetWrappedNativeAddress(opts *bind.TransactOpts, wrappedNative_ common.Address) (*types.Transaction, error) {
	return _WrapRouter.contract.Transact(opts, "setWrappedNativeAddress", wrappedNative_)
}

// SetWrappedNativeAddress is a paid mutator transaction binding the contract method 0x9ad21e01.
//
// Solidity: function setWrappedNativeAddress(address wrappedNative_) returns()
func (_WrapRouter *WrapRouterSession) SetWrappedNativeAddress(wrappedNative_ common.Address) (*types.Transaction, error) {
	return _WrapRouter.Contract.SetWrappedNativeAddress(&_WrapRouter.TransactOpts, wrappedNative_)
}

// SetWrappedNativeAddress is a paid mutator transaction binding the contract method 0x9ad21e01.
//
// Solidity: function setWrappedNativeAddress(address wrappedNative_) returns()
func (_WrapRouter *WrapRouterTransactorSession) SetWrappedNativeAddress(wrappedNative_ common.Address) (*types.Transaction, error) {
	return _WrapRouter.Contract.SetWrappedNativeAddress(&_WrapRouter.TransactOpts, wrappedNative_)
}

// TransferERC1155 is a paid mutator transaction binding the contract method 0x1cb56859.
//
// Solidity: function transferERC1155(address token_, address receiver_, uint256[] tokenIds_, uint256[] amounts_) payable returns()
func (_WrapRouter *WrapRouterTransactor) TransferERC1155(opts *bind.TransactOpts, token_ common.Address, receiver_ common.Address, tokenIds_ []*big.Int, amounts_ []*big.Int) (*types.Transaction, error) {
	return _WrapRouter.contract.Transact(opts, "transferERC1155", token_, receiver_, tokenIds_, amounts_)
}

// TransferERC1155 is a paid mutator transaction binding the contract method 0x1cb56859.
//
// Solidity: function transferERC1155(address token_, address receiver_, uint256[] tokenIds_, uint256[] amounts_) payable returns()
func (_WrapRouter *WrapRouterSession) TransferERC1155(token_ common.Address, receiver_ common.Address, tokenIds_ []*big.Int, amounts_ []*big.Int) (*types.Transaction, error) {
	return _WrapRouter.Contract.TransferERC1155(&_WrapRouter.TransactOpts, token_, receiver_, tokenIds_, amounts_)
}

// TransferERC1155 is a paid mutator transaction binding the contract method 0x1cb56859.
//
// Solidity: function transferERC1155(address token_, address receiver_, uint256[] tokenIds_, uint256[] amounts_) payable returns()
func (_WrapRouter *WrapRouterTransactorSession) TransferERC1155(token_ common.Address, receiver_ common.Address, tokenIds_ []*big.Int, amounts_ []*big.Int) (*types.Transaction, error) {
	return _WrapRouter.Contract.TransferERC1155(&_WrapRouter.TransactOpts, token_, receiver_, tokenIds_, amounts_)
}

// TransferERC20 is a paid mutator transaction binding the contract method 0x9db5dbe4.
//
// Solidity: function transferERC20(address token_, address receiver_, uint256 amount_) payable returns()
func (_WrapRouter *WrapRouterTransactor) TransferERC20(opts *bind.TransactOpts, token_ common.Address, receiver_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _WrapRouter.contract.Transact(opts, "transferERC20", token_, receiver_, amount_)
}

// TransferERC20 is a paid mutator transaction binding the contract method 0x9db5dbe4.
//
// Solidity: function transferERC20(address token_, address receiver_, uint256 amount_) payable returns()
func (_WrapRouter *WrapRouterSession) TransferERC20(token_ common.Address, receiver_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _WrapRouter.Contract.TransferERC20(&_WrapRouter.TransactOpts, token_, receiver_, amount_)
}

// TransferERC20 is a paid mutator transaction binding the contract method 0x9db5dbe4.
//
// Solidity: function transferERC20(address token_, address receiver_, uint256 amount_) payable returns()
func (_WrapRouter *WrapRouterTransactorSession) TransferERC20(token_ common.Address, receiver_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _WrapRouter.Contract.TransferERC20(&_WrapRouter.TransactOpts, token_, receiver_, amount_)
}

// TransferERC721 is a paid mutator transaction binding the contract method 0x8ca5e396.
//
// Solidity: function transferERC721(address token_, address receiver_, uint256[] nftIds_) payable returns()
func (_WrapRouter *WrapRouterTransactor) TransferERC721(opts *bind.TransactOpts, token_ common.Address, receiver_ common.Address, nftIds_ []*big.Int) (*types.Transaction, error) {
	return _WrapRouter.contract.Transact(opts, "transferERC721", token_, receiver_, nftIds_)
}

// TransferERC721 is a paid mutator transaction binding the contract method 0x8ca5e396.
//
// Solidity: function transferERC721(address token_, address receiver_, uint256[] nftIds_) payable returns()
func (_WrapRouter *WrapRouterSession) TransferERC721(token_ common.Address, receiver_ common.Address, nftIds_ []*big.Int) (*types.Transaction, error) {
	return _WrapRouter.Contract.TransferERC721(&_WrapRouter.TransactOpts, token_, receiver_, nftIds_)
}

// TransferERC721 is a paid mutator transaction binding the contract method 0x8ca5e396.
//
// Solidity: function transferERC721(address token_, address receiver_, uint256[] nftIds_) payable returns()
func (_WrapRouter *WrapRouterTransactorSession) TransferERC721(token_ common.Address, receiver_ common.Address, nftIds_ []*big.Int) (*types.Transaction, error) {
	return _WrapRouter.Contract.TransferERC721(&_WrapRouter.TransactOpts, token_, receiver_, nftIds_)
}

// TransferFromERC1155 is a paid mutator transaction binding the contract method 0x57e2c45f.
//
// Solidity: function transferFromERC1155(address token_, uint256[] tokenIds_, uint256[] amounts_) payable returns()
func (_WrapRouter *WrapRouterTransactor) TransferFromERC1155(opts *bind.TransactOpts, token_ common.Address, tokenIds_ []*big.Int, amounts_ []*big.Int) (*types.Transaction, error) {
	return _WrapRouter.contract.Transact(opts, "transferFromERC1155", token_, tokenIds_, amounts_)
}

// TransferFromERC1155 is a paid mutator transaction binding the contract method 0x57e2c45f.
//
// Solidity: function transferFromERC1155(address token_, uint256[] tokenIds_, uint256[] amounts_) payable returns()
func (_WrapRouter *WrapRouterSession) TransferFromERC1155(token_ common.Address, tokenIds_ []*big.Int, amounts_ []*big.Int) (*types.Transaction, error) {
	return _WrapRouter.Contract.TransferFromERC1155(&_WrapRouter.TransactOpts, token_, tokenIds_, amounts_)
}

// TransferFromERC1155 is a paid mutator transaction binding the contract method 0x57e2c45f.
//
// Solidity: function transferFromERC1155(address token_, uint256[] tokenIds_, uint256[] amounts_) payable returns()
func (_WrapRouter *WrapRouterTransactorSession) TransferFromERC1155(token_ common.Address, tokenIds_ []*big.Int, amounts_ []*big.Int) (*types.Transaction, error) {
	return _WrapRouter.Contract.TransferFromERC1155(&_WrapRouter.TransactOpts, token_, tokenIds_, amounts_)
}

// TransferFromERC20 is a paid mutator transaction binding the contract method 0xbd31ed1f.
//
// Solidity: function transferFromERC20(address token_, uint256 amount_) payable returns()
func (_WrapRouter *WrapRouterTransactor) TransferFromERC20(opts *bind.TransactOpts, token_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _WrapRouter.contract.Transact(opts, "transferFromERC20", token_, amount_)
}

// TransferFromERC20 is a paid mutator transaction binding the contract method 0xbd31ed1f.
//
// Solidity: function transferFromERC20(address token_, uint256 amount_) payable returns()
func (_WrapRouter *WrapRouterSession) TransferFromERC20(token_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _WrapRouter.Contract.TransferFromERC20(&_WrapRouter.TransactOpts, token_, amount_)
}

// TransferFromERC20 is a paid mutator transaction binding the contract method 0xbd31ed1f.
//
// Solidity: function transferFromERC20(address token_, uint256 amount_) payable returns()
func (_WrapRouter *WrapRouterTransactorSession) TransferFromERC20(token_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _WrapRouter.Contract.TransferFromERC20(&_WrapRouter.TransactOpts, token_, amount_)
}

// TransferFromERC721 is a paid mutator transaction binding the contract method 0x6ef8622b.
//
// Solidity: function transferFromERC721(address token_, uint256[] nftIds_) payable returns()
func (_WrapRouter *WrapRouterTransactor) TransferFromERC721(opts *bind.TransactOpts, token_ common.Address, nftIds_ []*big.Int) (*types.Transaction, error) {
	return _WrapRouter.contract.Transact(opts, "transferFromERC721", token_, nftIds_)
}

// TransferFromERC721 is a paid mutator transaction binding the contract method 0x6ef8622b.
//
// Solidity: function transferFromERC721(address token_, uint256[] nftIds_) payable returns()
func (_WrapRouter *WrapRouterSession) TransferFromERC721(token_ common.Address, nftIds_ []*big.Int) (*types.Transaction, error) {
	return _WrapRouter.Contract.TransferFromERC721(&_WrapRouter.TransactOpts, token_, nftIds_)
}

// TransferFromERC721 is a paid mutator transaction binding the contract method 0x6ef8622b.
//
// Solidity: function transferFromERC721(address token_, uint256[] nftIds_) payable returns()
func (_WrapRouter *WrapRouterTransactorSession) TransferFromERC721(token_ common.Address, nftIds_ []*big.Int) (*types.Transaction, error) {
	return _WrapRouter.Contract.TransferFromERC721(&_WrapRouter.TransactOpts, token_, nftIds_)
}

// TransferNative is a paid mutator transaction binding the contract method 0x7d2e90c2.
//
// Solidity: function transferNative(address receiver_, uint256 amount_) payable returns()
func (_WrapRouter *WrapRouterTransactor) TransferNative(opts *bind.TransactOpts, receiver_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _WrapRouter.contract.Transact(opts, "transferNative", receiver_, amount_)
}

// TransferNative is a paid mutator transaction binding the contract method 0x7d2e90c2.
//
// Solidity: function transferNative(address receiver_, uint256 amount_) payable returns()
func (_WrapRouter *WrapRouterSession) TransferNative(receiver_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _WrapRouter.Contract.TransferNative(&_WrapRouter.TransactOpts, receiver_, amount_)
}

// TransferNative is a paid mutator transaction binding the contract method 0x7d2e90c2.
//
// Solidity: function transferNative(address receiver_, uint256 amount_) payable returns()
func (_WrapRouter *WrapRouterTransactorSession) TransferNative(receiver_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _WrapRouter.Contract.TransferNative(&_WrapRouter.TransactOpts, receiver_, amount_)
}

// Unwrap is a paid mutator transaction binding the contract method 0x39f47693.
//
// Solidity: function unwrap(address receiver_, uint256 amount_) payable returns()
func (_WrapRouter *WrapRouterTransactor) Unwrap(opts *bind.TransactOpts, receiver_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _WrapRouter.contract.Transact(opts, "unwrap", receiver_, amount_)
}

// Unwrap is a paid mutator transaction binding the contract method 0x39f47693.
//
// Solidity: function unwrap(address receiver_, uint256 amount_) payable returns()
func (_WrapRouter *WrapRouterSession) Unwrap(receiver_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _WrapRouter.Contract.Unwrap(&_WrapRouter.TransactOpts, receiver_, amount_)
}

// Unwrap is a paid mutator transaction binding the contract method 0x39f47693.
//
// Solidity: function unwrap(address receiver_, uint256 amount_) payable returns()
func (_WrapRouter *WrapRouterTransactorSession) Unwrap(receiver_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _WrapRouter.Contract.Unwrap(&_WrapRouter.TransactOpts, receiver_, amount_)
}

// Wrap is a paid mutator transaction binding the contract method 0xbf376c7a.
//
// Solidity: function wrap(address receiver_, uint256 amount_) payable returns()
func (_WrapRouter *WrapRouterTransactor) Wrap(opts *bind.TransactOpts, receiver_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _WrapRouter.contract.Transact(opts, "wrap", receiver_, amount_)
}

// Wrap is a paid mutator transaction binding the contract method 0xbf376c7a.
//
// Solidity: function wrap(address receiver_, uint256 amount_) payable returns()
func (_WrapRouter *WrapRouterSession) Wrap(receiver_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _WrapRouter.Contract.Wrap(&_WrapRouter.TransactOpts, receiver_, amount_)
}

// Wrap is a paid mutator transaction binding the contract method 0xbf376c7a.
//
// Solidity: function wrap(address receiver_, uint256 amount_) payable returns()
func (_WrapRouter *WrapRouterTransactorSession) Wrap(receiver_ common.Address, amount_ *big.Int) (*types.Transaction, error) {
	return _WrapRouter.Contract.Wrap(&_WrapRouter.TransactOpts, receiver_, amount_)
}
