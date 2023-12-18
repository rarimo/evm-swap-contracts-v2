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

// SwapDiamondMetaData contains all meta data concerning the SwapDiamond contract.
var SwapDiamondMetaData = &bind.MetaData{
	ABI: "[{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[],\"name\":\"DIAMOND_STORAGE_SLOT\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OWNABLE_DIAMOND_STORAGE_SLOT\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SWAP_DIAMOND_STORAGE_SLOT\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"facet_\",\"type\":\"address\"},{\"internalType\":\"bytes4[]\",\"name\":\"selectors_\",\"type\":\"bytes4[]\"}],\"name\":\"addFacet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"facet_\",\"type\":\"address\"},{\"internalType\":\"bytes4[]\",\"name\":\"selectors_\",\"type\":\"bytes4[]\"},{\"internalType\":\"enumSwapDiamondStorage.SelectorType[]\",\"name\":\"types_\",\"type\":\"uint8[]\"}],\"name\":\"addFacet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"selector_\",\"type\":\"bytes4\"}],\"name\":\"getFacetBySelector\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"facet_\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"facet_\",\"type\":\"address\"}],\"name\":\"getFacetSelectors\",\"outputs\":[{\"internalType\":\"bytes4[]\",\"name\":\"selectors_\",\"type\":\"bytes4[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFacets\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"facets_\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"selector_\",\"type\":\"bytes4\"}],\"name\":\"getSelectorType\",\"outputs\":[{\"internalType\":\"enumSwapDiamondStorage.SelectorType\",\"name\":\"selectorType_\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"facet_\",\"type\":\"address\"},{\"internalType\":\"bytes4[]\",\"name\":\"selectors_\",\"type\":\"bytes4[]\"}],\"name\":\"removeFacet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner_\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"facet_\",\"type\":\"address\"},{\"internalType\":\"bytes4[]\",\"name\":\"fromSelectors_\",\"type\":\"bytes4[]\"},{\"internalType\":\"bytes4[]\",\"name\":\"toSelectors_\",\"type\":\"bytes4[]\"},{\"internalType\":\"enumSwapDiamondStorage.SelectorType[]\",\"name\":\"toTypes_\",\"type\":\"uint8[]\"}],\"name\":\"updateFacet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"facet_\",\"type\":\"address\"},{\"internalType\":\"bytes4[]\",\"name\":\"fromSelectors_\",\"type\":\"bytes4[]\"},{\"internalType\":\"bytes4[]\",\"name\":\"toSelectors_\",\"type\":\"bytes4[]\"}],\"name\":\"updateFacet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// SwapDiamondABI is the input ABI used to generate the binding from.
// Deprecated: Use SwapDiamondMetaData.ABI instead.
var SwapDiamondABI = SwapDiamondMetaData.ABI

// SwapDiamond is an auto generated Go binding around an Ethereum contract.
type SwapDiamond struct {
	SwapDiamondCaller     // Read-only binding to the contract
	SwapDiamondTransactor // Write-only binding to the contract
	SwapDiamondFilterer   // Log filterer for contract events
}

// SwapDiamondCaller is an auto generated read-only Go binding around an Ethereum contract.
type SwapDiamondCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapDiamondTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SwapDiamondTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapDiamondFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SwapDiamondFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SwapDiamondSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SwapDiamondSession struct {
	Contract     *SwapDiamond      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SwapDiamondCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SwapDiamondCallerSession struct {
	Contract *SwapDiamondCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// SwapDiamondTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SwapDiamondTransactorSession struct {
	Contract     *SwapDiamondTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// SwapDiamondRaw is an auto generated low-level Go binding around an Ethereum contract.
type SwapDiamondRaw struct {
	Contract *SwapDiamond // Generic contract binding to access the raw methods on
}

// SwapDiamondCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SwapDiamondCallerRaw struct {
	Contract *SwapDiamondCaller // Generic read-only contract binding to access the raw methods on
}

// SwapDiamondTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SwapDiamondTransactorRaw struct {
	Contract *SwapDiamondTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSwapDiamond creates a new instance of SwapDiamond, bound to a specific deployed contract.
func NewSwapDiamond(address common.Address, backend bind.ContractBackend) (*SwapDiamond, error) {
	contract, err := bindSwapDiamond(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SwapDiamond{SwapDiamondCaller: SwapDiamondCaller{contract: contract}, SwapDiamondTransactor: SwapDiamondTransactor{contract: contract}, SwapDiamondFilterer: SwapDiamondFilterer{contract: contract}}, nil
}

// NewSwapDiamondCaller creates a new read-only instance of SwapDiamond, bound to a specific deployed contract.
func NewSwapDiamondCaller(address common.Address, caller bind.ContractCaller) (*SwapDiamondCaller, error) {
	contract, err := bindSwapDiamond(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SwapDiamondCaller{contract: contract}, nil
}

// NewSwapDiamondTransactor creates a new write-only instance of SwapDiamond, bound to a specific deployed contract.
func NewSwapDiamondTransactor(address common.Address, transactor bind.ContractTransactor) (*SwapDiamondTransactor, error) {
	contract, err := bindSwapDiamond(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SwapDiamondTransactor{contract: contract}, nil
}

// NewSwapDiamondFilterer creates a new log filterer instance of SwapDiamond, bound to a specific deployed contract.
func NewSwapDiamondFilterer(address common.Address, filterer bind.ContractFilterer) (*SwapDiamondFilterer, error) {
	contract, err := bindSwapDiamond(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SwapDiamondFilterer{contract: contract}, nil
}

// bindSwapDiamond binds a generic wrapper to an already deployed contract.
func bindSwapDiamond(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SwapDiamondMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SwapDiamond *SwapDiamondRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SwapDiamond.Contract.SwapDiamondCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SwapDiamond *SwapDiamondRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SwapDiamond.Contract.SwapDiamondTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SwapDiamond *SwapDiamondRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SwapDiamond.Contract.SwapDiamondTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SwapDiamond *SwapDiamondCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SwapDiamond.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SwapDiamond *SwapDiamondTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SwapDiamond.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SwapDiamond *SwapDiamondTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SwapDiamond.Contract.contract.Transact(opts, method, params...)
}

// DIAMONDSTORAGESLOT is a free data retrieval call binding the contract method 0xe828c51f.
//
// Solidity: function DIAMOND_STORAGE_SLOT() view returns(bytes32)
func (_SwapDiamond *SwapDiamondCaller) DIAMONDSTORAGESLOT(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _SwapDiamond.contract.Call(opts, &out, "DIAMOND_STORAGE_SLOT")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DIAMONDSTORAGESLOT is a free data retrieval call binding the contract method 0xe828c51f.
//
// Solidity: function DIAMOND_STORAGE_SLOT() view returns(bytes32)
func (_SwapDiamond *SwapDiamondSession) DIAMONDSTORAGESLOT() ([32]byte, error) {
	return _SwapDiamond.Contract.DIAMONDSTORAGESLOT(&_SwapDiamond.CallOpts)
}

// DIAMONDSTORAGESLOT is a free data retrieval call binding the contract method 0xe828c51f.
//
// Solidity: function DIAMOND_STORAGE_SLOT() view returns(bytes32)
func (_SwapDiamond *SwapDiamondCallerSession) DIAMONDSTORAGESLOT() ([32]byte, error) {
	return _SwapDiamond.Contract.DIAMONDSTORAGESLOT(&_SwapDiamond.CallOpts)
}

// OWNABLEDIAMONDSTORAGESLOT is a free data retrieval call binding the contract method 0xe3e077ad.
//
// Solidity: function OWNABLE_DIAMOND_STORAGE_SLOT() view returns(bytes32)
func (_SwapDiamond *SwapDiamondCaller) OWNABLEDIAMONDSTORAGESLOT(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _SwapDiamond.contract.Call(opts, &out, "OWNABLE_DIAMOND_STORAGE_SLOT")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OWNABLEDIAMONDSTORAGESLOT is a free data retrieval call binding the contract method 0xe3e077ad.
//
// Solidity: function OWNABLE_DIAMOND_STORAGE_SLOT() view returns(bytes32)
func (_SwapDiamond *SwapDiamondSession) OWNABLEDIAMONDSTORAGESLOT() ([32]byte, error) {
	return _SwapDiamond.Contract.OWNABLEDIAMONDSTORAGESLOT(&_SwapDiamond.CallOpts)
}

// OWNABLEDIAMONDSTORAGESLOT is a free data retrieval call binding the contract method 0xe3e077ad.
//
// Solidity: function OWNABLE_DIAMOND_STORAGE_SLOT() view returns(bytes32)
func (_SwapDiamond *SwapDiamondCallerSession) OWNABLEDIAMONDSTORAGESLOT() ([32]byte, error) {
	return _SwapDiamond.Contract.OWNABLEDIAMONDSTORAGESLOT(&_SwapDiamond.CallOpts)
}

// SWAPDIAMONDSTORAGESLOT is a free data retrieval call binding the contract method 0xaaadd377.
//
// Solidity: function SWAP_DIAMOND_STORAGE_SLOT() view returns(bytes32)
func (_SwapDiamond *SwapDiamondCaller) SWAPDIAMONDSTORAGESLOT(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _SwapDiamond.contract.Call(opts, &out, "SWAP_DIAMOND_STORAGE_SLOT")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SWAPDIAMONDSTORAGESLOT is a free data retrieval call binding the contract method 0xaaadd377.
//
// Solidity: function SWAP_DIAMOND_STORAGE_SLOT() view returns(bytes32)
func (_SwapDiamond *SwapDiamondSession) SWAPDIAMONDSTORAGESLOT() ([32]byte, error) {
	return _SwapDiamond.Contract.SWAPDIAMONDSTORAGESLOT(&_SwapDiamond.CallOpts)
}

// SWAPDIAMONDSTORAGESLOT is a free data retrieval call binding the contract method 0xaaadd377.
//
// Solidity: function SWAP_DIAMOND_STORAGE_SLOT() view returns(bytes32)
func (_SwapDiamond *SwapDiamondCallerSession) SWAPDIAMONDSTORAGESLOT() ([32]byte, error) {
	return _SwapDiamond.Contract.SWAPDIAMONDSTORAGESLOT(&_SwapDiamond.CallOpts)
}

// GetFacetBySelector is a free data retrieval call binding the contract method 0xfe00955e.
//
// Solidity: function getFacetBySelector(bytes4 selector_) view returns(address facet_)
func (_SwapDiamond *SwapDiamondCaller) GetFacetBySelector(opts *bind.CallOpts, selector_ [4]byte) (common.Address, error) {
	var out []interface{}
	err := _SwapDiamond.contract.Call(opts, &out, "getFacetBySelector", selector_)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetFacetBySelector is a free data retrieval call binding the contract method 0xfe00955e.
//
// Solidity: function getFacetBySelector(bytes4 selector_) view returns(address facet_)
func (_SwapDiamond *SwapDiamondSession) GetFacetBySelector(selector_ [4]byte) (common.Address, error) {
	return _SwapDiamond.Contract.GetFacetBySelector(&_SwapDiamond.CallOpts, selector_)
}

// GetFacetBySelector is a free data retrieval call binding the contract method 0xfe00955e.
//
// Solidity: function getFacetBySelector(bytes4 selector_) view returns(address facet_)
func (_SwapDiamond *SwapDiamondCallerSession) GetFacetBySelector(selector_ [4]byte) (common.Address, error) {
	return _SwapDiamond.Contract.GetFacetBySelector(&_SwapDiamond.CallOpts, selector_)
}

// GetFacetSelectors is a free data retrieval call binding the contract method 0x8ea0b248.
//
// Solidity: function getFacetSelectors(address facet_) view returns(bytes4[] selectors_)
func (_SwapDiamond *SwapDiamondCaller) GetFacetSelectors(opts *bind.CallOpts, facet_ common.Address) ([][4]byte, error) {
	var out []interface{}
	err := _SwapDiamond.contract.Call(opts, &out, "getFacetSelectors", facet_)

	if err != nil {
		return *new([][4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][4]byte)).(*[][4]byte)

	return out0, err

}

// GetFacetSelectors is a free data retrieval call binding the contract method 0x8ea0b248.
//
// Solidity: function getFacetSelectors(address facet_) view returns(bytes4[] selectors_)
func (_SwapDiamond *SwapDiamondSession) GetFacetSelectors(facet_ common.Address) ([][4]byte, error) {
	return _SwapDiamond.Contract.GetFacetSelectors(&_SwapDiamond.CallOpts, facet_)
}

// GetFacetSelectors is a free data retrieval call binding the contract method 0x8ea0b248.
//
// Solidity: function getFacetSelectors(address facet_) view returns(bytes4[] selectors_)
func (_SwapDiamond *SwapDiamondCallerSession) GetFacetSelectors(facet_ common.Address) ([][4]byte, error) {
	return _SwapDiamond.Contract.GetFacetSelectors(&_SwapDiamond.CallOpts, facet_)
}

// GetFacets is a free data retrieval call binding the contract method 0x662ea47d.
//
// Solidity: function getFacets() view returns(address[] facets_)
func (_SwapDiamond *SwapDiamondCaller) GetFacets(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _SwapDiamond.contract.Call(opts, &out, "getFacets")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetFacets is a free data retrieval call binding the contract method 0x662ea47d.
//
// Solidity: function getFacets() view returns(address[] facets_)
func (_SwapDiamond *SwapDiamondSession) GetFacets() ([]common.Address, error) {
	return _SwapDiamond.Contract.GetFacets(&_SwapDiamond.CallOpts)
}

// GetFacets is a free data retrieval call binding the contract method 0x662ea47d.
//
// Solidity: function getFacets() view returns(address[] facets_)
func (_SwapDiamond *SwapDiamondCallerSession) GetFacets() ([]common.Address, error) {
	return _SwapDiamond.Contract.GetFacets(&_SwapDiamond.CallOpts)
}

// GetSelectorType is a free data retrieval call binding the contract method 0xad5403eb.
//
// Solidity: function getSelectorType(bytes4 selector_) view returns(uint8 selectorType_)
func (_SwapDiamond *SwapDiamondCaller) GetSelectorType(opts *bind.CallOpts, selector_ [4]byte) (uint8, error) {
	var out []interface{}
	err := _SwapDiamond.contract.Call(opts, &out, "getSelectorType", selector_)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetSelectorType is a free data retrieval call binding the contract method 0xad5403eb.
//
// Solidity: function getSelectorType(bytes4 selector_) view returns(uint8 selectorType_)
func (_SwapDiamond *SwapDiamondSession) GetSelectorType(selector_ [4]byte) (uint8, error) {
	return _SwapDiamond.Contract.GetSelectorType(&_SwapDiamond.CallOpts, selector_)
}

// GetSelectorType is a free data retrieval call binding the contract method 0xad5403eb.
//
// Solidity: function getSelectorType(bytes4 selector_) view returns(uint8 selectorType_)
func (_SwapDiamond *SwapDiamondCallerSession) GetSelectorType(selector_ [4]byte) (uint8, error) {
	return _SwapDiamond.Contract.GetSelectorType(&_SwapDiamond.CallOpts, selector_)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SwapDiamond *SwapDiamondCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SwapDiamond.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SwapDiamond *SwapDiamondSession) Owner() (common.Address, error) {
	return _SwapDiamond.Contract.Owner(&_SwapDiamond.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SwapDiamond *SwapDiamondCallerSession) Owner() (common.Address, error) {
	return _SwapDiamond.Contract.Owner(&_SwapDiamond.CallOpts)
}

// AddFacet is a paid mutator transaction binding the contract method 0x5547dad6.
//
// Solidity: function addFacet(address facet_, bytes4[] selectors_) returns()
func (_SwapDiamond *SwapDiamondTransactor) AddFacet(opts *bind.TransactOpts, facet_ common.Address, selectors_ [][4]byte) (*types.Transaction, error) {
	return _SwapDiamond.contract.Transact(opts, "addFacet", facet_, selectors_)
}

// AddFacet is a paid mutator transaction binding the contract method 0x5547dad6.
//
// Solidity: function addFacet(address facet_, bytes4[] selectors_) returns()
func (_SwapDiamond *SwapDiamondSession) AddFacet(facet_ common.Address, selectors_ [][4]byte) (*types.Transaction, error) {
	return _SwapDiamond.Contract.AddFacet(&_SwapDiamond.TransactOpts, facet_, selectors_)
}

// AddFacet is a paid mutator transaction binding the contract method 0x5547dad6.
//
// Solidity: function addFacet(address facet_, bytes4[] selectors_) returns()
func (_SwapDiamond *SwapDiamondTransactorSession) AddFacet(facet_ common.Address, selectors_ [][4]byte) (*types.Transaction, error) {
	return _SwapDiamond.Contract.AddFacet(&_SwapDiamond.TransactOpts, facet_, selectors_)
}

// AddFacet0 is a paid mutator transaction binding the contract method 0xc8ae30fd.
//
// Solidity: function addFacet(address facet_, bytes4[] selectors_, uint8[] types_) returns()
func (_SwapDiamond *SwapDiamondTransactor) AddFacet0(opts *bind.TransactOpts, facet_ common.Address, selectors_ [][4]byte, types_ []uint8) (*types.Transaction, error) {
	return _SwapDiamond.contract.Transact(opts, "addFacet0", facet_, selectors_, types_)
}

// AddFacet0 is a paid mutator transaction binding the contract method 0xc8ae30fd.
//
// Solidity: function addFacet(address facet_, bytes4[] selectors_, uint8[] types_) returns()
func (_SwapDiamond *SwapDiamondSession) AddFacet0(facet_ common.Address, selectors_ [][4]byte, types_ []uint8) (*types.Transaction, error) {
	return _SwapDiamond.Contract.AddFacet0(&_SwapDiamond.TransactOpts, facet_, selectors_, types_)
}

// AddFacet0 is a paid mutator transaction binding the contract method 0xc8ae30fd.
//
// Solidity: function addFacet(address facet_, bytes4[] selectors_, uint8[] types_) returns()
func (_SwapDiamond *SwapDiamondTransactorSession) AddFacet0(facet_ common.Address, selectors_ [][4]byte, types_ []uint8) (*types.Transaction, error) {
	return _SwapDiamond.Contract.AddFacet0(&_SwapDiamond.TransactOpts, facet_, selectors_, types_)
}

// RemoveFacet is a paid mutator transaction binding the contract method 0xcf380c86.
//
// Solidity: function removeFacet(address facet_, bytes4[] selectors_) returns()
func (_SwapDiamond *SwapDiamondTransactor) RemoveFacet(opts *bind.TransactOpts, facet_ common.Address, selectors_ [][4]byte) (*types.Transaction, error) {
	return _SwapDiamond.contract.Transact(opts, "removeFacet", facet_, selectors_)
}

// RemoveFacet is a paid mutator transaction binding the contract method 0xcf380c86.
//
// Solidity: function removeFacet(address facet_, bytes4[] selectors_) returns()
func (_SwapDiamond *SwapDiamondSession) RemoveFacet(facet_ common.Address, selectors_ [][4]byte) (*types.Transaction, error) {
	return _SwapDiamond.Contract.RemoveFacet(&_SwapDiamond.TransactOpts, facet_, selectors_)
}

// RemoveFacet is a paid mutator transaction binding the contract method 0xcf380c86.
//
// Solidity: function removeFacet(address facet_, bytes4[] selectors_) returns()
func (_SwapDiamond *SwapDiamondTransactorSession) RemoveFacet(facet_ common.Address, selectors_ [][4]byte) (*types.Transaction, error) {
	return _SwapDiamond.Contract.RemoveFacet(&_SwapDiamond.TransactOpts, facet_, selectors_)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner_) returns()
func (_SwapDiamond *SwapDiamondTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner_ common.Address) (*types.Transaction, error) {
	return _SwapDiamond.contract.Transact(opts, "transferOwnership", newOwner_)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner_) returns()
func (_SwapDiamond *SwapDiamondSession) TransferOwnership(newOwner_ common.Address) (*types.Transaction, error) {
	return _SwapDiamond.Contract.TransferOwnership(&_SwapDiamond.TransactOpts, newOwner_)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner_) returns()
func (_SwapDiamond *SwapDiamondTransactorSession) TransferOwnership(newOwner_ common.Address) (*types.Transaction, error) {
	return _SwapDiamond.Contract.TransferOwnership(&_SwapDiamond.TransactOpts, newOwner_)
}

// UpdateFacet is a paid mutator transaction binding the contract method 0x01aa2450.
//
// Solidity: function updateFacet(address facet_, bytes4[] fromSelectors_, bytes4[] toSelectors_, uint8[] toTypes_) returns()
func (_SwapDiamond *SwapDiamondTransactor) UpdateFacet(opts *bind.TransactOpts, facet_ common.Address, fromSelectors_ [][4]byte, toSelectors_ [][4]byte, toTypes_ []uint8) (*types.Transaction, error) {
	return _SwapDiamond.contract.Transact(opts, "updateFacet", facet_, fromSelectors_, toSelectors_, toTypes_)
}

// UpdateFacet is a paid mutator transaction binding the contract method 0x01aa2450.
//
// Solidity: function updateFacet(address facet_, bytes4[] fromSelectors_, bytes4[] toSelectors_, uint8[] toTypes_) returns()
func (_SwapDiamond *SwapDiamondSession) UpdateFacet(facet_ common.Address, fromSelectors_ [][4]byte, toSelectors_ [][4]byte, toTypes_ []uint8) (*types.Transaction, error) {
	return _SwapDiamond.Contract.UpdateFacet(&_SwapDiamond.TransactOpts, facet_, fromSelectors_, toSelectors_, toTypes_)
}

// UpdateFacet is a paid mutator transaction binding the contract method 0x01aa2450.
//
// Solidity: function updateFacet(address facet_, bytes4[] fromSelectors_, bytes4[] toSelectors_, uint8[] toTypes_) returns()
func (_SwapDiamond *SwapDiamondTransactorSession) UpdateFacet(facet_ common.Address, fromSelectors_ [][4]byte, toSelectors_ [][4]byte, toTypes_ []uint8) (*types.Transaction, error) {
	return _SwapDiamond.Contract.UpdateFacet(&_SwapDiamond.TransactOpts, facet_, fromSelectors_, toSelectors_, toTypes_)
}

// UpdateFacet0 is a paid mutator transaction binding the contract method 0x218bc10a.
//
// Solidity: function updateFacet(address facet_, bytes4[] fromSelectors_, bytes4[] toSelectors_) returns()
func (_SwapDiamond *SwapDiamondTransactor) UpdateFacet0(opts *bind.TransactOpts, facet_ common.Address, fromSelectors_ [][4]byte, toSelectors_ [][4]byte) (*types.Transaction, error) {
	return _SwapDiamond.contract.Transact(opts, "updateFacet0", facet_, fromSelectors_, toSelectors_)
}

// UpdateFacet0 is a paid mutator transaction binding the contract method 0x218bc10a.
//
// Solidity: function updateFacet(address facet_, bytes4[] fromSelectors_, bytes4[] toSelectors_) returns()
func (_SwapDiamond *SwapDiamondSession) UpdateFacet0(facet_ common.Address, fromSelectors_ [][4]byte, toSelectors_ [][4]byte) (*types.Transaction, error) {
	return _SwapDiamond.Contract.UpdateFacet0(&_SwapDiamond.TransactOpts, facet_, fromSelectors_, toSelectors_)
}

// UpdateFacet0 is a paid mutator transaction binding the contract method 0x218bc10a.
//
// Solidity: function updateFacet(address facet_, bytes4[] fromSelectors_, bytes4[] toSelectors_) returns()
func (_SwapDiamond *SwapDiamondTransactorSession) UpdateFacet0(facet_ common.Address, fromSelectors_ [][4]byte, toSelectors_ [][4]byte) (*types.Transaction, error) {
	return _SwapDiamond.Contract.UpdateFacet0(&_SwapDiamond.TransactOpts, facet_, fromSelectors_, toSelectors_)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_SwapDiamond *SwapDiamondTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _SwapDiamond.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_SwapDiamond *SwapDiamondSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _SwapDiamond.Contract.Fallback(&_SwapDiamond.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_SwapDiamond *SwapDiamondTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _SwapDiamond.Contract.Fallback(&_SwapDiamond.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SwapDiamond *SwapDiamondTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SwapDiamond.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SwapDiamond *SwapDiamondSession) Receive() (*types.Transaction, error) {
	return _SwapDiamond.Contract.Receive(&_SwapDiamond.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_SwapDiamond *SwapDiamondTransactorSession) Receive() (*types.Transaction, error) {
	return _SwapDiamond.Contract.Receive(&_SwapDiamond.TransactOpts)
}
