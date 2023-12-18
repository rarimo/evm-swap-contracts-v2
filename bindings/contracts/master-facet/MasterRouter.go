// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package masterfacet

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

// MasterRouterPayload is an auto generated low-level Go binding around an user-defined struct.
type MasterRouterPayload struct {
	Command    *big.Int
	SkipRevert bool
	Data       []byte
}

// MasterRouterMetaData contains all meta data concerning the MasterRouter contract.
var MasterRouterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"DIAMOND_STORAGE_SLOT\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MASTER_ROUTER_STORAGE_SLOT\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SWAP_DIAMOND_STORAGE_SLOT\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCallerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"caller_\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"selector_\",\"type\":\"bytes4\"}],\"name\":\"getFacetBySelector\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"facet_\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"facet_\",\"type\":\"address\"}],\"name\":\"getFacetSelectors\",\"outputs\":[{\"internalType\":\"bytes4[]\",\"name\":\"selectors_\",\"type\":\"bytes4[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFacets\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"facets_\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"selector_\",\"type\":\"bytes4\"}],\"name\":\"getSelectorType\",\"outputs\":[{\"internalType\":\"enumSwapDiamondStorage.SelectorType\",\"name\":\"selectorType_\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"command\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"skipRevert\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structMasterRouter.Payload[]\",\"name\":\"payloads_\",\"type\":\"tuple[]\"}],\"name\":\"make\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// MasterRouterABI is the input ABI used to generate the binding from.
// Deprecated: Use MasterRouterMetaData.ABI instead.
var MasterRouterABI = MasterRouterMetaData.ABI

// MasterRouter is an auto generated Go binding around an Ethereum contract.
type MasterRouter struct {
	MasterRouterCaller     // Read-only binding to the contract
	MasterRouterTransactor // Write-only binding to the contract
	MasterRouterFilterer   // Log filterer for contract events
}

// MasterRouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type MasterRouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MasterRouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MasterRouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MasterRouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MasterRouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MasterRouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MasterRouterSession struct {
	Contract     *MasterRouter     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MasterRouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MasterRouterCallerSession struct {
	Contract *MasterRouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// MasterRouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MasterRouterTransactorSession struct {
	Contract     *MasterRouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// MasterRouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type MasterRouterRaw struct {
	Contract *MasterRouter // Generic contract binding to access the raw methods on
}

// MasterRouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MasterRouterCallerRaw struct {
	Contract *MasterRouterCaller // Generic read-only contract binding to access the raw methods on
}

// MasterRouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MasterRouterTransactorRaw struct {
	Contract *MasterRouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMasterRouter creates a new instance of MasterRouter, bound to a specific deployed contract.
func NewMasterRouter(address common.Address, backend bind.ContractBackend) (*MasterRouter, error) {
	contract, err := bindMasterRouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MasterRouter{MasterRouterCaller: MasterRouterCaller{contract: contract}, MasterRouterTransactor: MasterRouterTransactor{contract: contract}, MasterRouterFilterer: MasterRouterFilterer{contract: contract}}, nil
}

// NewMasterRouterCaller creates a new read-only instance of MasterRouter, bound to a specific deployed contract.
func NewMasterRouterCaller(address common.Address, caller bind.ContractCaller) (*MasterRouterCaller, error) {
	contract, err := bindMasterRouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MasterRouterCaller{contract: contract}, nil
}

// NewMasterRouterTransactor creates a new write-only instance of MasterRouter, bound to a specific deployed contract.
func NewMasterRouterTransactor(address common.Address, transactor bind.ContractTransactor) (*MasterRouterTransactor, error) {
	contract, err := bindMasterRouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MasterRouterTransactor{contract: contract}, nil
}

// NewMasterRouterFilterer creates a new log filterer instance of MasterRouter, bound to a specific deployed contract.
func NewMasterRouterFilterer(address common.Address, filterer bind.ContractFilterer) (*MasterRouterFilterer, error) {
	contract, err := bindMasterRouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MasterRouterFilterer{contract: contract}, nil
}

// bindMasterRouter binds a generic wrapper to an already deployed contract.
func bindMasterRouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MasterRouterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MasterRouter *MasterRouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MasterRouter.Contract.MasterRouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MasterRouter *MasterRouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MasterRouter.Contract.MasterRouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MasterRouter *MasterRouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MasterRouter.Contract.MasterRouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MasterRouter *MasterRouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MasterRouter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MasterRouter *MasterRouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MasterRouter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MasterRouter *MasterRouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MasterRouter.Contract.contract.Transact(opts, method, params...)
}

// DIAMONDSTORAGESLOT is a free data retrieval call binding the contract method 0xe828c51f.
//
// Solidity: function DIAMOND_STORAGE_SLOT() view returns(bytes32)
func (_MasterRouter *MasterRouterCaller) DIAMONDSTORAGESLOT(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MasterRouter.contract.Call(opts, &out, "DIAMOND_STORAGE_SLOT")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DIAMONDSTORAGESLOT is a free data retrieval call binding the contract method 0xe828c51f.
//
// Solidity: function DIAMOND_STORAGE_SLOT() view returns(bytes32)
func (_MasterRouter *MasterRouterSession) DIAMONDSTORAGESLOT() ([32]byte, error) {
	return _MasterRouter.Contract.DIAMONDSTORAGESLOT(&_MasterRouter.CallOpts)
}

// DIAMONDSTORAGESLOT is a free data retrieval call binding the contract method 0xe828c51f.
//
// Solidity: function DIAMOND_STORAGE_SLOT() view returns(bytes32)
func (_MasterRouter *MasterRouterCallerSession) DIAMONDSTORAGESLOT() ([32]byte, error) {
	return _MasterRouter.Contract.DIAMONDSTORAGESLOT(&_MasterRouter.CallOpts)
}

// MASTERROUTERSTORAGESLOT is a free data retrieval call binding the contract method 0x660fd3b0.
//
// Solidity: function MASTER_ROUTER_STORAGE_SLOT() view returns(bytes32)
func (_MasterRouter *MasterRouterCaller) MASTERROUTERSTORAGESLOT(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MasterRouter.contract.Call(opts, &out, "MASTER_ROUTER_STORAGE_SLOT")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MASTERROUTERSTORAGESLOT is a free data retrieval call binding the contract method 0x660fd3b0.
//
// Solidity: function MASTER_ROUTER_STORAGE_SLOT() view returns(bytes32)
func (_MasterRouter *MasterRouterSession) MASTERROUTERSTORAGESLOT() ([32]byte, error) {
	return _MasterRouter.Contract.MASTERROUTERSTORAGESLOT(&_MasterRouter.CallOpts)
}

// MASTERROUTERSTORAGESLOT is a free data retrieval call binding the contract method 0x660fd3b0.
//
// Solidity: function MASTER_ROUTER_STORAGE_SLOT() view returns(bytes32)
func (_MasterRouter *MasterRouterCallerSession) MASTERROUTERSTORAGESLOT() ([32]byte, error) {
	return _MasterRouter.Contract.MASTERROUTERSTORAGESLOT(&_MasterRouter.CallOpts)
}

// SWAPDIAMONDSTORAGESLOT is a free data retrieval call binding the contract method 0xaaadd377.
//
// Solidity: function SWAP_DIAMOND_STORAGE_SLOT() view returns(bytes32)
func (_MasterRouter *MasterRouterCaller) SWAPDIAMONDSTORAGESLOT(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _MasterRouter.contract.Call(opts, &out, "SWAP_DIAMOND_STORAGE_SLOT")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SWAPDIAMONDSTORAGESLOT is a free data retrieval call binding the contract method 0xaaadd377.
//
// Solidity: function SWAP_DIAMOND_STORAGE_SLOT() view returns(bytes32)
func (_MasterRouter *MasterRouterSession) SWAPDIAMONDSTORAGESLOT() ([32]byte, error) {
	return _MasterRouter.Contract.SWAPDIAMONDSTORAGESLOT(&_MasterRouter.CallOpts)
}

// SWAPDIAMONDSTORAGESLOT is a free data retrieval call binding the contract method 0xaaadd377.
//
// Solidity: function SWAP_DIAMOND_STORAGE_SLOT() view returns(bytes32)
func (_MasterRouter *MasterRouterCallerSession) SWAPDIAMONDSTORAGESLOT() ([32]byte, error) {
	return _MasterRouter.Contract.SWAPDIAMONDSTORAGESLOT(&_MasterRouter.CallOpts)
}

// GetCallerAddress is a free data retrieval call binding the contract method 0x46b3353b.
//
// Solidity: function getCallerAddress() view returns(address caller_)
func (_MasterRouter *MasterRouterCaller) GetCallerAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MasterRouter.contract.Call(opts, &out, "getCallerAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetCallerAddress is a free data retrieval call binding the contract method 0x46b3353b.
//
// Solidity: function getCallerAddress() view returns(address caller_)
func (_MasterRouter *MasterRouterSession) GetCallerAddress() (common.Address, error) {
	return _MasterRouter.Contract.GetCallerAddress(&_MasterRouter.CallOpts)
}

// GetCallerAddress is a free data retrieval call binding the contract method 0x46b3353b.
//
// Solidity: function getCallerAddress() view returns(address caller_)
func (_MasterRouter *MasterRouterCallerSession) GetCallerAddress() (common.Address, error) {
	return _MasterRouter.Contract.GetCallerAddress(&_MasterRouter.CallOpts)
}

// GetFacetBySelector is a free data retrieval call binding the contract method 0xfe00955e.
//
// Solidity: function getFacetBySelector(bytes4 selector_) view returns(address facet_)
func (_MasterRouter *MasterRouterCaller) GetFacetBySelector(opts *bind.CallOpts, selector_ [4]byte) (common.Address, error) {
	var out []interface{}
	err := _MasterRouter.contract.Call(opts, &out, "getFacetBySelector", selector_)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetFacetBySelector is a free data retrieval call binding the contract method 0xfe00955e.
//
// Solidity: function getFacetBySelector(bytes4 selector_) view returns(address facet_)
func (_MasterRouter *MasterRouterSession) GetFacetBySelector(selector_ [4]byte) (common.Address, error) {
	return _MasterRouter.Contract.GetFacetBySelector(&_MasterRouter.CallOpts, selector_)
}

// GetFacetBySelector is a free data retrieval call binding the contract method 0xfe00955e.
//
// Solidity: function getFacetBySelector(bytes4 selector_) view returns(address facet_)
func (_MasterRouter *MasterRouterCallerSession) GetFacetBySelector(selector_ [4]byte) (common.Address, error) {
	return _MasterRouter.Contract.GetFacetBySelector(&_MasterRouter.CallOpts, selector_)
}

// GetFacetSelectors is a free data retrieval call binding the contract method 0x8ea0b248.
//
// Solidity: function getFacetSelectors(address facet_) view returns(bytes4[] selectors_)
func (_MasterRouter *MasterRouterCaller) GetFacetSelectors(opts *bind.CallOpts, facet_ common.Address) ([][4]byte, error) {
	var out []interface{}
	err := _MasterRouter.contract.Call(opts, &out, "getFacetSelectors", facet_)

	if err != nil {
		return *new([][4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][4]byte)).(*[][4]byte)

	return out0, err

}

// GetFacetSelectors is a free data retrieval call binding the contract method 0x8ea0b248.
//
// Solidity: function getFacetSelectors(address facet_) view returns(bytes4[] selectors_)
func (_MasterRouter *MasterRouterSession) GetFacetSelectors(facet_ common.Address) ([][4]byte, error) {
	return _MasterRouter.Contract.GetFacetSelectors(&_MasterRouter.CallOpts, facet_)
}

// GetFacetSelectors is a free data retrieval call binding the contract method 0x8ea0b248.
//
// Solidity: function getFacetSelectors(address facet_) view returns(bytes4[] selectors_)
func (_MasterRouter *MasterRouterCallerSession) GetFacetSelectors(facet_ common.Address) ([][4]byte, error) {
	return _MasterRouter.Contract.GetFacetSelectors(&_MasterRouter.CallOpts, facet_)
}

// GetFacets is a free data retrieval call binding the contract method 0x662ea47d.
//
// Solidity: function getFacets() view returns(address[] facets_)
func (_MasterRouter *MasterRouterCaller) GetFacets(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _MasterRouter.contract.Call(opts, &out, "getFacets")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetFacets is a free data retrieval call binding the contract method 0x662ea47d.
//
// Solidity: function getFacets() view returns(address[] facets_)
func (_MasterRouter *MasterRouterSession) GetFacets() ([]common.Address, error) {
	return _MasterRouter.Contract.GetFacets(&_MasterRouter.CallOpts)
}

// GetFacets is a free data retrieval call binding the contract method 0x662ea47d.
//
// Solidity: function getFacets() view returns(address[] facets_)
func (_MasterRouter *MasterRouterCallerSession) GetFacets() ([]common.Address, error) {
	return _MasterRouter.Contract.GetFacets(&_MasterRouter.CallOpts)
}

// GetSelectorType is a free data retrieval call binding the contract method 0xad5403eb.
//
// Solidity: function getSelectorType(bytes4 selector_) view returns(uint8 selectorType_)
func (_MasterRouter *MasterRouterCaller) GetSelectorType(opts *bind.CallOpts, selector_ [4]byte) (uint8, error) {
	var out []interface{}
	err := _MasterRouter.contract.Call(opts, &out, "getSelectorType", selector_)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetSelectorType is a free data retrieval call binding the contract method 0xad5403eb.
//
// Solidity: function getSelectorType(bytes4 selector_) view returns(uint8 selectorType_)
func (_MasterRouter *MasterRouterSession) GetSelectorType(selector_ [4]byte) (uint8, error) {
	return _MasterRouter.Contract.GetSelectorType(&_MasterRouter.CallOpts, selector_)
}

// GetSelectorType is a free data retrieval call binding the contract method 0xad5403eb.
//
// Solidity: function getSelectorType(bytes4 selector_) view returns(uint8 selectorType_)
func (_MasterRouter *MasterRouterCallerSession) GetSelectorType(selector_ [4]byte) (uint8, error) {
	return _MasterRouter.Contract.GetSelectorType(&_MasterRouter.CallOpts, selector_)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_MasterRouter *MasterRouterCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _MasterRouter.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_MasterRouter *MasterRouterSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _MasterRouter.Contract.SupportsInterface(&_MasterRouter.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_MasterRouter *MasterRouterCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _MasterRouter.Contract.SupportsInterface(&_MasterRouter.CallOpts, interfaceId)
}

// Make is a paid mutator transaction binding the contract method 0x7a344c6e.
//
// Solidity: function make((uint256,bool,bytes)[] payloads_) payable returns()
func (_MasterRouter *MasterRouterTransactor) Make(opts *bind.TransactOpts, payloads_ []MasterRouterPayload) (*types.Transaction, error) {
	return _MasterRouter.contract.Transact(opts, "make", payloads_)
}

// Make is a paid mutator transaction binding the contract method 0x7a344c6e.
//
// Solidity: function make((uint256,bool,bytes)[] payloads_) payable returns()
func (_MasterRouter *MasterRouterSession) Make(payloads_ []MasterRouterPayload) (*types.Transaction, error) {
	return _MasterRouter.Contract.Make(&_MasterRouter.TransactOpts, payloads_)
}

// Make is a paid mutator transaction binding the contract method 0x7a344c6e.
//
// Solidity: function make((uint256,bool,bytes)[] payloads_) payable returns()
func (_MasterRouter *MasterRouterTransactorSession) Make(payloads_ []MasterRouterPayload) (*types.Transaction, error) {
	return _MasterRouter.Contract.Make(&_MasterRouter.TransactOpts, payloads_)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_MasterRouter *MasterRouterTransactor) OnERC1155BatchReceived(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _MasterRouter.contract.Transact(opts, "onERC1155BatchReceived", arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_MasterRouter *MasterRouterSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _MasterRouter.Contract.OnERC1155BatchReceived(&_MasterRouter.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address , address , uint256[] , uint256[] , bytes ) returns(bytes4)
func (_MasterRouter *MasterRouterTransactorSession) OnERC1155BatchReceived(arg0 common.Address, arg1 common.Address, arg2 []*big.Int, arg3 []*big.Int, arg4 []byte) (*types.Transaction, error) {
	return _MasterRouter.Contract.OnERC1155BatchReceived(&_MasterRouter.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_MasterRouter *MasterRouterTransactor) OnERC1155Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _MasterRouter.contract.Transact(opts, "onERC1155Received", arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_MasterRouter *MasterRouterSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _MasterRouter.Contract.OnERC1155Received(&_MasterRouter.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address , address , uint256 , uint256 , bytes ) returns(bytes4)
func (_MasterRouter *MasterRouterTransactorSession) OnERC1155Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 *big.Int, arg4 []byte) (*types.Transaction, error) {
	return _MasterRouter.Contract.OnERC1155Received(&_MasterRouter.TransactOpts, arg0, arg1, arg2, arg3, arg4)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_MasterRouter *MasterRouterTransactor) OnERC721Received(opts *bind.TransactOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _MasterRouter.contract.Transact(opts, "onERC721Received", arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_MasterRouter *MasterRouterSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _MasterRouter.Contract.OnERC721Received(&_MasterRouter.TransactOpts, arg0, arg1, arg2, arg3)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address , address , uint256 , bytes ) returns(bytes4)
func (_MasterRouter *MasterRouterTransactorSession) OnERC721Received(arg0 common.Address, arg1 common.Address, arg2 *big.Int, arg3 []byte) (*types.Transaction, error) {
	return _MasterRouter.Contract.OnERC721Received(&_MasterRouter.TransactOpts, arg0, arg1, arg2, arg3)
}
