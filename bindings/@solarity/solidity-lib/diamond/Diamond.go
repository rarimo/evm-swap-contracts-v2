// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package diamond

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

// DiamondMetaData contains all meta data concerning the Diamond contract.
var DiamondMetaData = &bind.MetaData{
	ABI: "[{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"DIAMOND_STORAGE_SLOT\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"selector_\",\"type\":\"bytes4\"}],\"name\":\"getFacetBySelector\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"facet_\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"facet_\",\"type\":\"address\"}],\"name\":\"getFacetSelectors\",\"outputs\":[{\"internalType\":\"bytes4[]\",\"name\":\"selectors_\",\"type\":\"bytes4[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFacets\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"facets_\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// DiamondABI is the input ABI used to generate the binding from.
// Deprecated: Use DiamondMetaData.ABI instead.
var DiamondABI = DiamondMetaData.ABI

// Diamond is an auto generated Go binding around an Ethereum contract.
type Diamond struct {
	DiamondCaller     // Read-only binding to the contract
	DiamondTransactor // Write-only binding to the contract
	DiamondFilterer   // Log filterer for contract events
}

// DiamondCaller is an auto generated read-only Go binding around an Ethereum contract.
type DiamondCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DiamondTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DiamondTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DiamondFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DiamondFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DiamondSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DiamondSession struct {
	Contract     *Diamond          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DiamondCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DiamondCallerSession struct {
	Contract *DiamondCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// DiamondTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DiamondTransactorSession struct {
	Contract     *DiamondTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// DiamondRaw is an auto generated low-level Go binding around an Ethereum contract.
type DiamondRaw struct {
	Contract *Diamond // Generic contract binding to access the raw methods on
}

// DiamondCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DiamondCallerRaw struct {
	Contract *DiamondCaller // Generic read-only contract binding to access the raw methods on
}

// DiamondTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DiamondTransactorRaw struct {
	Contract *DiamondTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDiamond creates a new instance of Diamond, bound to a specific deployed contract.
func NewDiamond(address common.Address, backend bind.ContractBackend) (*Diamond, error) {
	contract, err := bindDiamond(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Diamond{DiamondCaller: DiamondCaller{contract: contract}, DiamondTransactor: DiamondTransactor{contract: contract}, DiamondFilterer: DiamondFilterer{contract: contract}}, nil
}

// NewDiamondCaller creates a new read-only instance of Diamond, bound to a specific deployed contract.
func NewDiamondCaller(address common.Address, caller bind.ContractCaller) (*DiamondCaller, error) {
	contract, err := bindDiamond(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DiamondCaller{contract: contract}, nil
}

// NewDiamondTransactor creates a new write-only instance of Diamond, bound to a specific deployed contract.
func NewDiamondTransactor(address common.Address, transactor bind.ContractTransactor) (*DiamondTransactor, error) {
	contract, err := bindDiamond(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DiamondTransactor{contract: contract}, nil
}

// NewDiamondFilterer creates a new log filterer instance of Diamond, bound to a specific deployed contract.
func NewDiamondFilterer(address common.Address, filterer bind.ContractFilterer) (*DiamondFilterer, error) {
	contract, err := bindDiamond(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DiamondFilterer{contract: contract}, nil
}

// bindDiamond binds a generic wrapper to an already deployed contract.
func bindDiamond(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DiamondMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Diamond *DiamondRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Diamond.Contract.DiamondCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Diamond *DiamondRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Diamond.Contract.DiamondTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Diamond *DiamondRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Diamond.Contract.DiamondTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Diamond *DiamondCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Diamond.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Diamond *DiamondTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Diamond.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Diamond *DiamondTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Diamond.Contract.contract.Transact(opts, method, params...)
}

// DIAMONDSTORAGESLOT is a free data retrieval call binding the contract method 0xe828c51f.
//
// Solidity: function DIAMOND_STORAGE_SLOT() view returns(bytes32)
func (_Diamond *DiamondCaller) DIAMONDSTORAGESLOT(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Diamond.contract.Call(opts, &out, "DIAMOND_STORAGE_SLOT")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DIAMONDSTORAGESLOT is a free data retrieval call binding the contract method 0xe828c51f.
//
// Solidity: function DIAMOND_STORAGE_SLOT() view returns(bytes32)
func (_Diamond *DiamondSession) DIAMONDSTORAGESLOT() ([32]byte, error) {
	return _Diamond.Contract.DIAMONDSTORAGESLOT(&_Diamond.CallOpts)
}

// DIAMONDSTORAGESLOT is a free data retrieval call binding the contract method 0xe828c51f.
//
// Solidity: function DIAMOND_STORAGE_SLOT() view returns(bytes32)
func (_Diamond *DiamondCallerSession) DIAMONDSTORAGESLOT() ([32]byte, error) {
	return _Diamond.Contract.DIAMONDSTORAGESLOT(&_Diamond.CallOpts)
}

// GetFacetBySelector is a free data retrieval call binding the contract method 0xfe00955e.
//
// Solidity: function getFacetBySelector(bytes4 selector_) view returns(address facet_)
func (_Diamond *DiamondCaller) GetFacetBySelector(opts *bind.CallOpts, selector_ [4]byte) (common.Address, error) {
	var out []interface{}
	err := _Diamond.contract.Call(opts, &out, "getFacetBySelector", selector_)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetFacetBySelector is a free data retrieval call binding the contract method 0xfe00955e.
//
// Solidity: function getFacetBySelector(bytes4 selector_) view returns(address facet_)
func (_Diamond *DiamondSession) GetFacetBySelector(selector_ [4]byte) (common.Address, error) {
	return _Diamond.Contract.GetFacetBySelector(&_Diamond.CallOpts, selector_)
}

// GetFacetBySelector is a free data retrieval call binding the contract method 0xfe00955e.
//
// Solidity: function getFacetBySelector(bytes4 selector_) view returns(address facet_)
func (_Diamond *DiamondCallerSession) GetFacetBySelector(selector_ [4]byte) (common.Address, error) {
	return _Diamond.Contract.GetFacetBySelector(&_Diamond.CallOpts, selector_)
}

// GetFacetSelectors is a free data retrieval call binding the contract method 0x8ea0b248.
//
// Solidity: function getFacetSelectors(address facet_) view returns(bytes4[] selectors_)
func (_Diamond *DiamondCaller) GetFacetSelectors(opts *bind.CallOpts, facet_ common.Address) ([][4]byte, error) {
	var out []interface{}
	err := _Diamond.contract.Call(opts, &out, "getFacetSelectors", facet_)

	if err != nil {
		return *new([][4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][4]byte)).(*[][4]byte)

	return out0, err

}

// GetFacetSelectors is a free data retrieval call binding the contract method 0x8ea0b248.
//
// Solidity: function getFacetSelectors(address facet_) view returns(bytes4[] selectors_)
func (_Diamond *DiamondSession) GetFacetSelectors(facet_ common.Address) ([][4]byte, error) {
	return _Diamond.Contract.GetFacetSelectors(&_Diamond.CallOpts, facet_)
}

// GetFacetSelectors is a free data retrieval call binding the contract method 0x8ea0b248.
//
// Solidity: function getFacetSelectors(address facet_) view returns(bytes4[] selectors_)
func (_Diamond *DiamondCallerSession) GetFacetSelectors(facet_ common.Address) ([][4]byte, error) {
	return _Diamond.Contract.GetFacetSelectors(&_Diamond.CallOpts, facet_)
}

// GetFacets is a free data retrieval call binding the contract method 0x662ea47d.
//
// Solidity: function getFacets() view returns(address[] facets_)
func (_Diamond *DiamondCaller) GetFacets(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Diamond.contract.Call(opts, &out, "getFacets")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetFacets is a free data retrieval call binding the contract method 0x662ea47d.
//
// Solidity: function getFacets() view returns(address[] facets_)
func (_Diamond *DiamondSession) GetFacets() ([]common.Address, error) {
	return _Diamond.Contract.GetFacets(&_Diamond.CallOpts)
}

// GetFacets is a free data retrieval call binding the contract method 0x662ea47d.
//
// Solidity: function getFacets() view returns(address[] facets_)
func (_Diamond *DiamondCallerSession) GetFacets() ([]common.Address, error) {
	return _Diamond.Contract.GetFacets(&_Diamond.CallOpts)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Diamond *DiamondTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Diamond.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Diamond *DiamondSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Diamond.Contract.Fallback(&_Diamond.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Diamond *DiamondTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Diamond.Contract.Fallback(&_Diamond.TransactOpts, calldata)
}
