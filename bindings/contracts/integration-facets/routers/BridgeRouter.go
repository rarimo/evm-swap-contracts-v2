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

// IBridgeFacadeDepositFeeERC1155Parameters is an auto generated low-level Go binding around an user-defined struct.
type IBridgeFacadeDepositFeeERC1155Parameters struct {
	FeeToken common.Address
}

// IBridgeFacadeDepositFeeERC20Parameters is an auto generated low-level Go binding around an user-defined struct.
type IBridgeFacadeDepositFeeERC20Parameters struct {
	FeeToken common.Address
}

// IBridgeFacadeDepositFeeERC721Parameters is an auto generated low-level Go binding around an user-defined struct.
type IBridgeFacadeDepositFeeERC721Parameters struct {
	FeeToken common.Address
}

// IBridgeFacadeDepositFeeNativeParameters is an auto generated low-level Go binding around an user-defined struct.
type IBridgeFacadeDepositFeeNativeParameters struct {
	FeeToken common.Address
}

// IBundlerBundle is an auto generated low-level Go binding around an user-defined struct.
type IBundlerBundle struct {
	Salt   [32]byte
	Bundle []byte
}

// IERC1155HandlerDepositERC1155Parameters is an auto generated low-level Go binding around an user-defined struct.
type IERC1155HandlerDepositERC1155Parameters struct {
	Token     common.Address
	TokenId   *big.Int
	Amount    *big.Int
	Bundle    IBundlerBundle
	Network   string
	Receiver  string
	IsWrapped bool
}

// IERC20HandlerDepositERC20Parameters is an auto generated low-level Go binding around an user-defined struct.
type IERC20HandlerDepositERC20Parameters struct {
	Token     common.Address
	Amount    *big.Int
	Bundle    IBundlerBundle
	Network   string
	Receiver  string
	IsWrapped bool
}

// IERC721HandlerDepositERC721Parameters is an auto generated low-level Go binding around an user-defined struct.
type IERC721HandlerDepositERC721Parameters struct {
	Token     common.Address
	TokenId   *big.Int
	Bundle    IBundlerBundle
	Network   string
	Receiver  string
	IsWrapped bool
}

// INativeHandlerDepositNativeParameters is an auto generated low-level Go binding around an user-defined struct.
type INativeHandlerDepositNativeParameters struct {
	Amount   *big.Int
	Bundle   IBundlerBundle
	Network  string
	Receiver string
}

// BridgeRouterMetaData contains all meta data concerning the BridgeRouter contract.
var BridgeRouterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"BRIDGE_ROUTER_STORAGE_SLOT\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OWNABLE_DIAMOND_STORAGE_SLOT\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"}],\"internalType\":\"structIBridgeFacade.DepositFeeERC1155Parameters\",\"name\":\"feeParams_\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"}],\"internalType\":\"structIBundler.Bundle\",\"name\":\"bundle\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"network\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"receiver\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isWrapped\",\"type\":\"bool\"}],\"internalType\":\"structIERC1155Handler.DepositERC1155Parameters\",\"name\":\"depositParams_\",\"type\":\"tuple\"}],\"name\":\"bridgeERC1155\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"}],\"internalType\":\"structIBridgeFacade.DepositFeeERC20Parameters\",\"name\":\"feeParams_\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"}],\"internalType\":\"structIBundler.Bundle\",\"name\":\"bundle\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"network\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"receiver\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isWrapped\",\"type\":\"bool\"}],\"internalType\":\"structIERC20Handler.DepositERC20Parameters\",\"name\":\"depositParams_\",\"type\":\"tuple\"}],\"name\":\"bridgeERC20\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"}],\"internalType\":\"structIBridgeFacade.DepositFeeERC721Parameters\",\"name\":\"feeParams_\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"}],\"internalType\":\"structIBundler.Bundle\",\"name\":\"bundle\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"network\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"receiver\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isWrapped\",\"type\":\"bool\"}],\"internalType\":\"structIERC721Handler.DepositERC721Parameters\",\"name\":\"depositParams_\",\"type\":\"tuple\"}],\"name\":\"bridgeERC721\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"feeToken\",\"type\":\"address\"}],\"internalType\":\"structIBridgeFacade.DepositFeeNativeParameters\",\"name\":\"feeParams_\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"bundle\",\"type\":\"bytes\"}],\"internalType\":\"structIBundler.Bundle\",\"name\":\"bundle\",\"type\":\"tuple\"},{\"internalType\":\"string\",\"name\":\"network\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"receiver\",\"type\":\"string\"}],\"internalType\":\"structINativeHandler.DepositNativeParameters\",\"name\":\"depositParams_\",\"type\":\"tuple\"}],\"name\":\"bridgeNative\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBridgeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"bridge_\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bridge_\",\"type\":\"address\"}],\"name\":\"setBridgeAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// BridgeRouterABI is the input ABI used to generate the binding from.
// Deprecated: Use BridgeRouterMetaData.ABI instead.
var BridgeRouterABI = BridgeRouterMetaData.ABI

// BridgeRouter is an auto generated Go binding around an Ethereum contract.
type BridgeRouter struct {
	BridgeRouterCaller     // Read-only binding to the contract
	BridgeRouterTransactor // Write-only binding to the contract
	BridgeRouterFilterer   // Log filterer for contract events
}

// BridgeRouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type BridgeRouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeRouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BridgeRouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeRouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BridgeRouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeRouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BridgeRouterSession struct {
	Contract     *BridgeRouter     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeRouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BridgeRouterCallerSession struct {
	Contract *BridgeRouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// BridgeRouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BridgeRouterTransactorSession struct {
	Contract     *BridgeRouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// BridgeRouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type BridgeRouterRaw struct {
	Contract *BridgeRouter // Generic contract binding to access the raw methods on
}

// BridgeRouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BridgeRouterCallerRaw struct {
	Contract *BridgeRouterCaller // Generic read-only contract binding to access the raw methods on
}

// BridgeRouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BridgeRouterTransactorRaw struct {
	Contract *BridgeRouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridgeRouter creates a new instance of BridgeRouter, bound to a specific deployed contract.
func NewBridgeRouter(address common.Address, backend bind.ContractBackend) (*BridgeRouter, error) {
	contract, err := bindBridgeRouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BridgeRouter{BridgeRouterCaller: BridgeRouterCaller{contract: contract}, BridgeRouterTransactor: BridgeRouterTransactor{contract: contract}, BridgeRouterFilterer: BridgeRouterFilterer{contract: contract}}, nil
}

// NewBridgeRouterCaller creates a new read-only instance of BridgeRouter, bound to a specific deployed contract.
func NewBridgeRouterCaller(address common.Address, caller bind.ContractCaller) (*BridgeRouterCaller, error) {
	contract, err := bindBridgeRouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeRouterCaller{contract: contract}, nil
}

// NewBridgeRouterTransactor creates a new write-only instance of BridgeRouter, bound to a specific deployed contract.
func NewBridgeRouterTransactor(address common.Address, transactor bind.ContractTransactor) (*BridgeRouterTransactor, error) {
	contract, err := bindBridgeRouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeRouterTransactor{contract: contract}, nil
}

// NewBridgeRouterFilterer creates a new log filterer instance of BridgeRouter, bound to a specific deployed contract.
func NewBridgeRouterFilterer(address common.Address, filterer bind.ContractFilterer) (*BridgeRouterFilterer, error) {
	contract, err := bindBridgeRouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BridgeRouterFilterer{contract: contract}, nil
}

// bindBridgeRouter binds a generic wrapper to an already deployed contract.
func bindBridgeRouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BridgeRouterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeRouter *BridgeRouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgeRouter.Contract.BridgeRouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeRouter *BridgeRouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeRouter.Contract.BridgeRouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeRouter *BridgeRouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeRouter.Contract.BridgeRouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeRouter *BridgeRouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BridgeRouter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeRouter *BridgeRouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeRouter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeRouter *BridgeRouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeRouter.Contract.contract.Transact(opts, method, params...)
}

// BRIDGEROUTERSTORAGESLOT is a free data retrieval call binding the contract method 0x5e080a57.
//
// Solidity: function BRIDGE_ROUTER_STORAGE_SLOT() view returns(bytes32)
func (_BridgeRouter *BridgeRouterCaller) BRIDGEROUTERSTORAGESLOT(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BridgeRouter.contract.Call(opts, &out, "BRIDGE_ROUTER_STORAGE_SLOT")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BRIDGEROUTERSTORAGESLOT is a free data retrieval call binding the contract method 0x5e080a57.
//
// Solidity: function BRIDGE_ROUTER_STORAGE_SLOT() view returns(bytes32)
func (_BridgeRouter *BridgeRouterSession) BRIDGEROUTERSTORAGESLOT() ([32]byte, error) {
	return _BridgeRouter.Contract.BRIDGEROUTERSTORAGESLOT(&_BridgeRouter.CallOpts)
}

// BRIDGEROUTERSTORAGESLOT is a free data retrieval call binding the contract method 0x5e080a57.
//
// Solidity: function BRIDGE_ROUTER_STORAGE_SLOT() view returns(bytes32)
func (_BridgeRouter *BridgeRouterCallerSession) BRIDGEROUTERSTORAGESLOT() ([32]byte, error) {
	return _BridgeRouter.Contract.BRIDGEROUTERSTORAGESLOT(&_BridgeRouter.CallOpts)
}

// OWNABLEDIAMONDSTORAGESLOT is a free data retrieval call binding the contract method 0xe3e077ad.
//
// Solidity: function OWNABLE_DIAMOND_STORAGE_SLOT() view returns(bytes32)
func (_BridgeRouter *BridgeRouterCaller) OWNABLEDIAMONDSTORAGESLOT(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _BridgeRouter.contract.Call(opts, &out, "OWNABLE_DIAMOND_STORAGE_SLOT")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OWNABLEDIAMONDSTORAGESLOT is a free data retrieval call binding the contract method 0xe3e077ad.
//
// Solidity: function OWNABLE_DIAMOND_STORAGE_SLOT() view returns(bytes32)
func (_BridgeRouter *BridgeRouterSession) OWNABLEDIAMONDSTORAGESLOT() ([32]byte, error) {
	return _BridgeRouter.Contract.OWNABLEDIAMONDSTORAGESLOT(&_BridgeRouter.CallOpts)
}

// OWNABLEDIAMONDSTORAGESLOT is a free data retrieval call binding the contract method 0xe3e077ad.
//
// Solidity: function OWNABLE_DIAMOND_STORAGE_SLOT() view returns(bytes32)
func (_BridgeRouter *BridgeRouterCallerSession) OWNABLEDIAMONDSTORAGESLOT() ([32]byte, error) {
	return _BridgeRouter.Contract.OWNABLEDIAMONDSTORAGESLOT(&_BridgeRouter.CallOpts)
}

// GetBridgeAddress is a free data retrieval call binding the contract method 0xfb32c508.
//
// Solidity: function getBridgeAddress() view returns(address bridge_)
func (_BridgeRouter *BridgeRouterCaller) GetBridgeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BridgeRouter.contract.Call(opts, &out, "getBridgeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetBridgeAddress is a free data retrieval call binding the contract method 0xfb32c508.
//
// Solidity: function getBridgeAddress() view returns(address bridge_)
func (_BridgeRouter *BridgeRouterSession) GetBridgeAddress() (common.Address, error) {
	return _BridgeRouter.Contract.GetBridgeAddress(&_BridgeRouter.CallOpts)
}

// GetBridgeAddress is a free data retrieval call binding the contract method 0xfb32c508.
//
// Solidity: function getBridgeAddress() view returns(address bridge_)
func (_BridgeRouter *BridgeRouterCallerSession) GetBridgeAddress() (common.Address, error) {
	return _BridgeRouter.Contract.GetBridgeAddress(&_BridgeRouter.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BridgeRouter *BridgeRouterCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BridgeRouter.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BridgeRouter *BridgeRouterSession) Owner() (common.Address, error) {
	return _BridgeRouter.Contract.Owner(&_BridgeRouter.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BridgeRouter *BridgeRouterCallerSession) Owner() (common.Address, error) {
	return _BridgeRouter.Contract.Owner(&_BridgeRouter.CallOpts)
}

// BridgeERC1155 is a paid mutator transaction binding the contract method 0x4f7e12b5.
//
// Solidity: function bridgeERC1155((address) feeParams_, (address,uint256,uint256,(bytes32,bytes),string,string,bool) depositParams_) payable returns()
func (_BridgeRouter *BridgeRouterTransactor) BridgeERC1155(opts *bind.TransactOpts, feeParams_ IBridgeFacadeDepositFeeERC1155Parameters, depositParams_ IERC1155HandlerDepositERC1155Parameters) (*types.Transaction, error) {
	return _BridgeRouter.contract.Transact(opts, "bridgeERC1155", feeParams_, depositParams_)
}

// BridgeERC1155 is a paid mutator transaction binding the contract method 0x4f7e12b5.
//
// Solidity: function bridgeERC1155((address) feeParams_, (address,uint256,uint256,(bytes32,bytes),string,string,bool) depositParams_) payable returns()
func (_BridgeRouter *BridgeRouterSession) BridgeERC1155(feeParams_ IBridgeFacadeDepositFeeERC1155Parameters, depositParams_ IERC1155HandlerDepositERC1155Parameters) (*types.Transaction, error) {
	return _BridgeRouter.Contract.BridgeERC1155(&_BridgeRouter.TransactOpts, feeParams_, depositParams_)
}

// BridgeERC1155 is a paid mutator transaction binding the contract method 0x4f7e12b5.
//
// Solidity: function bridgeERC1155((address) feeParams_, (address,uint256,uint256,(bytes32,bytes),string,string,bool) depositParams_) payable returns()
func (_BridgeRouter *BridgeRouterTransactorSession) BridgeERC1155(feeParams_ IBridgeFacadeDepositFeeERC1155Parameters, depositParams_ IERC1155HandlerDepositERC1155Parameters) (*types.Transaction, error) {
	return _BridgeRouter.Contract.BridgeERC1155(&_BridgeRouter.TransactOpts, feeParams_, depositParams_)
}

// BridgeERC20 is a paid mutator transaction binding the contract method 0xf3fe433a.
//
// Solidity: function bridgeERC20((address) feeParams_, (address,uint256,(bytes32,bytes),string,string,bool) depositParams_) payable returns()
func (_BridgeRouter *BridgeRouterTransactor) BridgeERC20(opts *bind.TransactOpts, feeParams_ IBridgeFacadeDepositFeeERC20Parameters, depositParams_ IERC20HandlerDepositERC20Parameters) (*types.Transaction, error) {
	return _BridgeRouter.contract.Transact(opts, "bridgeERC20", feeParams_, depositParams_)
}

// BridgeERC20 is a paid mutator transaction binding the contract method 0xf3fe433a.
//
// Solidity: function bridgeERC20((address) feeParams_, (address,uint256,(bytes32,bytes),string,string,bool) depositParams_) payable returns()
func (_BridgeRouter *BridgeRouterSession) BridgeERC20(feeParams_ IBridgeFacadeDepositFeeERC20Parameters, depositParams_ IERC20HandlerDepositERC20Parameters) (*types.Transaction, error) {
	return _BridgeRouter.Contract.BridgeERC20(&_BridgeRouter.TransactOpts, feeParams_, depositParams_)
}

// BridgeERC20 is a paid mutator transaction binding the contract method 0xf3fe433a.
//
// Solidity: function bridgeERC20((address) feeParams_, (address,uint256,(bytes32,bytes),string,string,bool) depositParams_) payable returns()
func (_BridgeRouter *BridgeRouterTransactorSession) BridgeERC20(feeParams_ IBridgeFacadeDepositFeeERC20Parameters, depositParams_ IERC20HandlerDepositERC20Parameters) (*types.Transaction, error) {
	return _BridgeRouter.Contract.BridgeERC20(&_BridgeRouter.TransactOpts, feeParams_, depositParams_)
}

// BridgeERC721 is a paid mutator transaction binding the contract method 0x6ffad695.
//
// Solidity: function bridgeERC721((address) feeParams_, (address,uint256,(bytes32,bytes),string,string,bool) depositParams_) payable returns()
func (_BridgeRouter *BridgeRouterTransactor) BridgeERC721(opts *bind.TransactOpts, feeParams_ IBridgeFacadeDepositFeeERC721Parameters, depositParams_ IERC721HandlerDepositERC721Parameters) (*types.Transaction, error) {
	return _BridgeRouter.contract.Transact(opts, "bridgeERC721", feeParams_, depositParams_)
}

// BridgeERC721 is a paid mutator transaction binding the contract method 0x6ffad695.
//
// Solidity: function bridgeERC721((address) feeParams_, (address,uint256,(bytes32,bytes),string,string,bool) depositParams_) payable returns()
func (_BridgeRouter *BridgeRouterSession) BridgeERC721(feeParams_ IBridgeFacadeDepositFeeERC721Parameters, depositParams_ IERC721HandlerDepositERC721Parameters) (*types.Transaction, error) {
	return _BridgeRouter.Contract.BridgeERC721(&_BridgeRouter.TransactOpts, feeParams_, depositParams_)
}

// BridgeERC721 is a paid mutator transaction binding the contract method 0x6ffad695.
//
// Solidity: function bridgeERC721((address) feeParams_, (address,uint256,(bytes32,bytes),string,string,bool) depositParams_) payable returns()
func (_BridgeRouter *BridgeRouterTransactorSession) BridgeERC721(feeParams_ IBridgeFacadeDepositFeeERC721Parameters, depositParams_ IERC721HandlerDepositERC721Parameters) (*types.Transaction, error) {
	return _BridgeRouter.Contract.BridgeERC721(&_BridgeRouter.TransactOpts, feeParams_, depositParams_)
}

// BridgeNative is a paid mutator transaction binding the contract method 0x4fd4d774.
//
// Solidity: function bridgeNative((address) feeParams_, (uint256,(bytes32,bytes),string,string) depositParams_) payable returns()
func (_BridgeRouter *BridgeRouterTransactor) BridgeNative(opts *bind.TransactOpts, feeParams_ IBridgeFacadeDepositFeeNativeParameters, depositParams_ INativeHandlerDepositNativeParameters) (*types.Transaction, error) {
	return _BridgeRouter.contract.Transact(opts, "bridgeNative", feeParams_, depositParams_)
}

// BridgeNative is a paid mutator transaction binding the contract method 0x4fd4d774.
//
// Solidity: function bridgeNative((address) feeParams_, (uint256,(bytes32,bytes),string,string) depositParams_) payable returns()
func (_BridgeRouter *BridgeRouterSession) BridgeNative(feeParams_ IBridgeFacadeDepositFeeNativeParameters, depositParams_ INativeHandlerDepositNativeParameters) (*types.Transaction, error) {
	return _BridgeRouter.Contract.BridgeNative(&_BridgeRouter.TransactOpts, feeParams_, depositParams_)
}

// BridgeNative is a paid mutator transaction binding the contract method 0x4fd4d774.
//
// Solidity: function bridgeNative((address) feeParams_, (uint256,(bytes32,bytes),string,string) depositParams_) payable returns()
func (_BridgeRouter *BridgeRouterTransactorSession) BridgeNative(feeParams_ IBridgeFacadeDepositFeeNativeParameters, depositParams_ INativeHandlerDepositNativeParameters) (*types.Transaction, error) {
	return _BridgeRouter.Contract.BridgeNative(&_BridgeRouter.TransactOpts, feeParams_, depositParams_)
}

// SetBridgeAddress is a paid mutator transaction binding the contract method 0x7f5a22f9.
//
// Solidity: function setBridgeAddress(address bridge_) returns()
func (_BridgeRouter *BridgeRouterTransactor) SetBridgeAddress(opts *bind.TransactOpts, bridge_ common.Address) (*types.Transaction, error) {
	return _BridgeRouter.contract.Transact(opts, "setBridgeAddress", bridge_)
}

// SetBridgeAddress is a paid mutator transaction binding the contract method 0x7f5a22f9.
//
// Solidity: function setBridgeAddress(address bridge_) returns()
func (_BridgeRouter *BridgeRouterSession) SetBridgeAddress(bridge_ common.Address) (*types.Transaction, error) {
	return _BridgeRouter.Contract.SetBridgeAddress(&_BridgeRouter.TransactOpts, bridge_)
}

// SetBridgeAddress is a paid mutator transaction binding the contract method 0x7f5a22f9.
//
// Solidity: function setBridgeAddress(address bridge_) returns()
func (_BridgeRouter *BridgeRouterTransactorSession) SetBridgeAddress(bridge_ common.Address) (*types.Transaction, error) {
	return _BridgeRouter.Contract.SetBridgeAddress(&_BridgeRouter.TransactOpts, bridge_)
}
