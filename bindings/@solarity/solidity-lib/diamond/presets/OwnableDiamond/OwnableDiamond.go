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

// OwnableDiamondMetaData contains all meta data concerning the OwnableDiamond contract.
var OwnableDiamondMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"DIAMOND_STORAGE_SLOT\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OWNABLE_DIAMOND_STORAGE_SLOT\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"facet_\",\"type\":\"address\"},{\"internalType\":\"bytes4[]\",\"name\":\"selectors_\",\"type\":\"bytes4[]\"}],\"name\":\"addFacet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"selector_\",\"type\":\"bytes4\"}],\"name\":\"getFacetBySelector\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"facet_\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"facet_\",\"type\":\"address\"}],\"name\":\"getFacetSelectors\",\"outputs\":[{\"internalType\":\"bytes4[]\",\"name\":\"selectors_\",\"type\":\"bytes4[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFacets\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"facets_\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"facet_\",\"type\":\"address\"},{\"internalType\":\"bytes4[]\",\"name\":\"selectors_\",\"type\":\"bytes4[]\"}],\"name\":\"removeFacet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner_\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"facet_\",\"type\":\"address\"},{\"internalType\":\"bytes4[]\",\"name\":\"fromSelectors_\",\"type\":\"bytes4[]\"},{\"internalType\":\"bytes4[]\",\"name\":\"toSelectors_\",\"type\":\"bytes4[]\"}],\"name\":\"updateFacet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// OwnableDiamondABI is the input ABI used to generate the binding from.
// Deprecated: Use OwnableDiamondMetaData.ABI instead.
var OwnableDiamondABI = OwnableDiamondMetaData.ABI

// OwnableDiamond is an auto generated Go binding around an Ethereum contract.
type OwnableDiamond struct {
	OwnableDiamondCaller     // Read-only binding to the contract
	OwnableDiamondTransactor // Write-only binding to the contract
	OwnableDiamondFilterer   // Log filterer for contract events
}

// OwnableDiamondCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwnableDiamondCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableDiamondTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnableDiamondTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableDiamondFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OwnableDiamondFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableDiamondSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwnableDiamondSession struct {
	Contract     *OwnableDiamond   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnableDiamondCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwnableDiamondCallerSession struct {
	Contract *OwnableDiamondCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// OwnableDiamondTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwnableDiamondTransactorSession struct {
	Contract     *OwnableDiamondTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// OwnableDiamondRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwnableDiamondRaw struct {
	Contract *OwnableDiamond // Generic contract binding to access the raw methods on
}

// OwnableDiamondCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwnableDiamondCallerRaw struct {
	Contract *OwnableDiamondCaller // Generic read-only contract binding to access the raw methods on
}

// OwnableDiamondTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwnableDiamondTransactorRaw struct {
	Contract *OwnableDiamondTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwnableDiamond creates a new instance of OwnableDiamond, bound to a specific deployed contract.
func NewOwnableDiamond(address common.Address, backend bind.ContractBackend) (*OwnableDiamond, error) {
	contract, err := bindOwnableDiamond(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OwnableDiamond{OwnableDiamondCaller: OwnableDiamondCaller{contract: contract}, OwnableDiamondTransactor: OwnableDiamondTransactor{contract: contract}, OwnableDiamondFilterer: OwnableDiamondFilterer{contract: contract}}, nil
}

// NewOwnableDiamondCaller creates a new read-only instance of OwnableDiamond, bound to a specific deployed contract.
func NewOwnableDiamondCaller(address common.Address, caller bind.ContractCaller) (*OwnableDiamondCaller, error) {
	contract, err := bindOwnableDiamond(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableDiamondCaller{contract: contract}, nil
}

// NewOwnableDiamondTransactor creates a new write-only instance of OwnableDiamond, bound to a specific deployed contract.
func NewOwnableDiamondTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnableDiamondTransactor, error) {
	contract, err := bindOwnableDiamond(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableDiamondTransactor{contract: contract}, nil
}

// NewOwnableDiamondFilterer creates a new log filterer instance of OwnableDiamond, bound to a specific deployed contract.
func NewOwnableDiamondFilterer(address common.Address, filterer bind.ContractFilterer) (*OwnableDiamondFilterer, error) {
	contract, err := bindOwnableDiamond(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwnableDiamondFilterer{contract: contract}, nil
}

// bindOwnableDiamond binds a generic wrapper to an already deployed contract.
func bindOwnableDiamond(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OwnableDiamondMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OwnableDiamond *OwnableDiamondRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OwnableDiamond.Contract.OwnableDiamondCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OwnableDiamond *OwnableDiamondRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OwnableDiamond.Contract.OwnableDiamondTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OwnableDiamond *OwnableDiamondRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OwnableDiamond.Contract.OwnableDiamondTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OwnableDiamond *OwnableDiamondCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OwnableDiamond.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OwnableDiamond *OwnableDiamondTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OwnableDiamond.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OwnableDiamond *OwnableDiamondTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OwnableDiamond.Contract.contract.Transact(opts, method, params...)
}

// DIAMONDSTORAGESLOT is a free data retrieval call binding the contract method 0xe828c51f.
//
// Solidity: function DIAMOND_STORAGE_SLOT() view returns(bytes32)
func (_OwnableDiamond *OwnableDiamondCaller) DIAMONDSTORAGESLOT(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _OwnableDiamond.contract.Call(opts, &out, "DIAMOND_STORAGE_SLOT")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DIAMONDSTORAGESLOT is a free data retrieval call binding the contract method 0xe828c51f.
//
// Solidity: function DIAMOND_STORAGE_SLOT() view returns(bytes32)
func (_OwnableDiamond *OwnableDiamondSession) DIAMONDSTORAGESLOT() ([32]byte, error) {
	return _OwnableDiamond.Contract.DIAMONDSTORAGESLOT(&_OwnableDiamond.CallOpts)
}

// DIAMONDSTORAGESLOT is a free data retrieval call binding the contract method 0xe828c51f.
//
// Solidity: function DIAMOND_STORAGE_SLOT() view returns(bytes32)
func (_OwnableDiamond *OwnableDiamondCallerSession) DIAMONDSTORAGESLOT() ([32]byte, error) {
	return _OwnableDiamond.Contract.DIAMONDSTORAGESLOT(&_OwnableDiamond.CallOpts)
}

// OWNABLEDIAMONDSTORAGESLOT is a free data retrieval call binding the contract method 0xe3e077ad.
//
// Solidity: function OWNABLE_DIAMOND_STORAGE_SLOT() view returns(bytes32)
func (_OwnableDiamond *OwnableDiamondCaller) OWNABLEDIAMONDSTORAGESLOT(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _OwnableDiamond.contract.Call(opts, &out, "OWNABLE_DIAMOND_STORAGE_SLOT")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OWNABLEDIAMONDSTORAGESLOT is a free data retrieval call binding the contract method 0xe3e077ad.
//
// Solidity: function OWNABLE_DIAMOND_STORAGE_SLOT() view returns(bytes32)
func (_OwnableDiamond *OwnableDiamondSession) OWNABLEDIAMONDSTORAGESLOT() ([32]byte, error) {
	return _OwnableDiamond.Contract.OWNABLEDIAMONDSTORAGESLOT(&_OwnableDiamond.CallOpts)
}

// OWNABLEDIAMONDSTORAGESLOT is a free data retrieval call binding the contract method 0xe3e077ad.
//
// Solidity: function OWNABLE_DIAMOND_STORAGE_SLOT() view returns(bytes32)
func (_OwnableDiamond *OwnableDiamondCallerSession) OWNABLEDIAMONDSTORAGESLOT() ([32]byte, error) {
	return _OwnableDiamond.Contract.OWNABLEDIAMONDSTORAGESLOT(&_OwnableDiamond.CallOpts)
}

// GetFacetBySelector is a free data retrieval call binding the contract method 0xfe00955e.
//
// Solidity: function getFacetBySelector(bytes4 selector_) view returns(address facet_)
func (_OwnableDiamond *OwnableDiamondCaller) GetFacetBySelector(opts *bind.CallOpts, selector_ [4]byte) (common.Address, error) {
	var out []interface{}
	err := _OwnableDiamond.contract.Call(opts, &out, "getFacetBySelector", selector_)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetFacetBySelector is a free data retrieval call binding the contract method 0xfe00955e.
//
// Solidity: function getFacetBySelector(bytes4 selector_) view returns(address facet_)
func (_OwnableDiamond *OwnableDiamondSession) GetFacetBySelector(selector_ [4]byte) (common.Address, error) {
	return _OwnableDiamond.Contract.GetFacetBySelector(&_OwnableDiamond.CallOpts, selector_)
}

// GetFacetBySelector is a free data retrieval call binding the contract method 0xfe00955e.
//
// Solidity: function getFacetBySelector(bytes4 selector_) view returns(address facet_)
func (_OwnableDiamond *OwnableDiamondCallerSession) GetFacetBySelector(selector_ [4]byte) (common.Address, error) {
	return _OwnableDiamond.Contract.GetFacetBySelector(&_OwnableDiamond.CallOpts, selector_)
}

// GetFacetSelectors is a free data retrieval call binding the contract method 0x8ea0b248.
//
// Solidity: function getFacetSelectors(address facet_) view returns(bytes4[] selectors_)
func (_OwnableDiamond *OwnableDiamondCaller) GetFacetSelectors(opts *bind.CallOpts, facet_ common.Address) ([][4]byte, error) {
	var out []interface{}
	err := _OwnableDiamond.contract.Call(opts, &out, "getFacetSelectors", facet_)

	if err != nil {
		return *new([][4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][4]byte)).(*[][4]byte)

	return out0, err

}

// GetFacetSelectors is a free data retrieval call binding the contract method 0x8ea0b248.
//
// Solidity: function getFacetSelectors(address facet_) view returns(bytes4[] selectors_)
func (_OwnableDiamond *OwnableDiamondSession) GetFacetSelectors(facet_ common.Address) ([][4]byte, error) {
	return _OwnableDiamond.Contract.GetFacetSelectors(&_OwnableDiamond.CallOpts, facet_)
}

// GetFacetSelectors is a free data retrieval call binding the contract method 0x8ea0b248.
//
// Solidity: function getFacetSelectors(address facet_) view returns(bytes4[] selectors_)
func (_OwnableDiamond *OwnableDiamondCallerSession) GetFacetSelectors(facet_ common.Address) ([][4]byte, error) {
	return _OwnableDiamond.Contract.GetFacetSelectors(&_OwnableDiamond.CallOpts, facet_)
}

// GetFacets is a free data retrieval call binding the contract method 0x662ea47d.
//
// Solidity: function getFacets() view returns(address[] facets_)
func (_OwnableDiamond *OwnableDiamondCaller) GetFacets(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _OwnableDiamond.contract.Call(opts, &out, "getFacets")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetFacets is a free data retrieval call binding the contract method 0x662ea47d.
//
// Solidity: function getFacets() view returns(address[] facets_)
func (_OwnableDiamond *OwnableDiamondSession) GetFacets() ([]common.Address, error) {
	return _OwnableDiamond.Contract.GetFacets(&_OwnableDiamond.CallOpts)
}

// GetFacets is a free data retrieval call binding the contract method 0x662ea47d.
//
// Solidity: function getFacets() view returns(address[] facets_)
func (_OwnableDiamond *OwnableDiamondCallerSession) GetFacets() ([]common.Address, error) {
	return _OwnableDiamond.Contract.GetFacets(&_OwnableDiamond.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OwnableDiamond *OwnableDiamondCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OwnableDiamond.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OwnableDiamond *OwnableDiamondSession) Owner() (common.Address, error) {
	return _OwnableDiamond.Contract.Owner(&_OwnableDiamond.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OwnableDiamond *OwnableDiamondCallerSession) Owner() (common.Address, error) {
	return _OwnableDiamond.Contract.Owner(&_OwnableDiamond.CallOpts)
}

// AddFacet is a paid mutator transaction binding the contract method 0x5547dad6.
//
// Solidity: function addFacet(address facet_, bytes4[] selectors_) returns()
func (_OwnableDiamond *OwnableDiamondTransactor) AddFacet(opts *bind.TransactOpts, facet_ common.Address, selectors_ [][4]byte) (*types.Transaction, error) {
	return _OwnableDiamond.contract.Transact(opts, "addFacet", facet_, selectors_)
}

// AddFacet is a paid mutator transaction binding the contract method 0x5547dad6.
//
// Solidity: function addFacet(address facet_, bytes4[] selectors_) returns()
func (_OwnableDiamond *OwnableDiamondSession) AddFacet(facet_ common.Address, selectors_ [][4]byte) (*types.Transaction, error) {
	return _OwnableDiamond.Contract.AddFacet(&_OwnableDiamond.TransactOpts, facet_, selectors_)
}

// AddFacet is a paid mutator transaction binding the contract method 0x5547dad6.
//
// Solidity: function addFacet(address facet_, bytes4[] selectors_) returns()
func (_OwnableDiamond *OwnableDiamondTransactorSession) AddFacet(facet_ common.Address, selectors_ [][4]byte) (*types.Transaction, error) {
	return _OwnableDiamond.Contract.AddFacet(&_OwnableDiamond.TransactOpts, facet_, selectors_)
}

// RemoveFacet is a paid mutator transaction binding the contract method 0xcf380c86.
//
// Solidity: function removeFacet(address facet_, bytes4[] selectors_) returns()
func (_OwnableDiamond *OwnableDiamondTransactor) RemoveFacet(opts *bind.TransactOpts, facet_ common.Address, selectors_ [][4]byte) (*types.Transaction, error) {
	return _OwnableDiamond.contract.Transact(opts, "removeFacet", facet_, selectors_)
}

// RemoveFacet is a paid mutator transaction binding the contract method 0xcf380c86.
//
// Solidity: function removeFacet(address facet_, bytes4[] selectors_) returns()
func (_OwnableDiamond *OwnableDiamondSession) RemoveFacet(facet_ common.Address, selectors_ [][4]byte) (*types.Transaction, error) {
	return _OwnableDiamond.Contract.RemoveFacet(&_OwnableDiamond.TransactOpts, facet_, selectors_)
}

// RemoveFacet is a paid mutator transaction binding the contract method 0xcf380c86.
//
// Solidity: function removeFacet(address facet_, bytes4[] selectors_) returns()
func (_OwnableDiamond *OwnableDiamondTransactorSession) RemoveFacet(facet_ common.Address, selectors_ [][4]byte) (*types.Transaction, error) {
	return _OwnableDiamond.Contract.RemoveFacet(&_OwnableDiamond.TransactOpts, facet_, selectors_)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner_) returns()
func (_OwnableDiamond *OwnableDiamondTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner_ common.Address) (*types.Transaction, error) {
	return _OwnableDiamond.contract.Transact(opts, "transferOwnership", newOwner_)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner_) returns()
func (_OwnableDiamond *OwnableDiamondSession) TransferOwnership(newOwner_ common.Address) (*types.Transaction, error) {
	return _OwnableDiamond.Contract.TransferOwnership(&_OwnableDiamond.TransactOpts, newOwner_)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner_) returns()
func (_OwnableDiamond *OwnableDiamondTransactorSession) TransferOwnership(newOwner_ common.Address) (*types.Transaction, error) {
	return _OwnableDiamond.Contract.TransferOwnership(&_OwnableDiamond.TransactOpts, newOwner_)
}

// UpdateFacet is a paid mutator transaction binding the contract method 0x218bc10a.
//
// Solidity: function updateFacet(address facet_, bytes4[] fromSelectors_, bytes4[] toSelectors_) returns()
func (_OwnableDiamond *OwnableDiamondTransactor) UpdateFacet(opts *bind.TransactOpts, facet_ common.Address, fromSelectors_ [][4]byte, toSelectors_ [][4]byte) (*types.Transaction, error) {
	return _OwnableDiamond.contract.Transact(opts, "updateFacet", facet_, fromSelectors_, toSelectors_)
}

// UpdateFacet is a paid mutator transaction binding the contract method 0x218bc10a.
//
// Solidity: function updateFacet(address facet_, bytes4[] fromSelectors_, bytes4[] toSelectors_) returns()
func (_OwnableDiamond *OwnableDiamondSession) UpdateFacet(facet_ common.Address, fromSelectors_ [][4]byte, toSelectors_ [][4]byte) (*types.Transaction, error) {
	return _OwnableDiamond.Contract.UpdateFacet(&_OwnableDiamond.TransactOpts, facet_, fromSelectors_, toSelectors_)
}

// UpdateFacet is a paid mutator transaction binding the contract method 0x218bc10a.
//
// Solidity: function updateFacet(address facet_, bytes4[] fromSelectors_, bytes4[] toSelectors_) returns()
func (_OwnableDiamond *OwnableDiamondTransactorSession) UpdateFacet(facet_ common.Address, fromSelectors_ [][4]byte, toSelectors_ [][4]byte) (*types.Transaction, error) {
	return _OwnableDiamond.Contract.UpdateFacet(&_OwnableDiamond.TransactOpts, facet_, fromSelectors_, toSelectors_)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_OwnableDiamond *OwnableDiamondTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _OwnableDiamond.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_OwnableDiamond *OwnableDiamondSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _OwnableDiamond.Contract.Fallback(&_OwnableDiamond.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_OwnableDiamond *OwnableDiamondTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _OwnableDiamond.Contract.Fallback(&_OwnableDiamond.TransactOpts, calldata)
}
