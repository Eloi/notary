// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package notarydapp

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// NotarydappABI is the input ABI used to generate the binding from.
const NotarydappABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"getDocumentByHash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"owner\",\"type\":\"string\"}],\"name\":\"getDocumentHashesByOwner\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"documentHash\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"documentOwner\",\"type\":\"string\"}],\"name\":\"registerDocument\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Notarydapp is an auto generated Go binding around an Ethereum contract.
type Notarydapp struct {
	NotarydappCaller     // Read-only binding to the contract
	NotarydappTransactor // Write-only binding to the contract
	NotarydappFilterer   // Log filterer for contract events
}

// NotarydappCaller is an auto generated read-only Go binding around an Ethereum contract.
type NotarydappCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NotarydappTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NotarydappTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NotarydappFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NotarydappFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NotarydappSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NotarydappSession struct {
	Contract     *Notarydapp       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NotarydappCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NotarydappCallerSession struct {
	Contract *NotarydappCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// NotarydappTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NotarydappTransactorSession struct {
	Contract     *NotarydappTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// NotarydappRaw is an auto generated low-level Go binding around an Ethereum contract.
type NotarydappRaw struct {
	Contract *Notarydapp // Generic contract binding to access the raw methods on
}

// NotarydappCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NotarydappCallerRaw struct {
	Contract *NotarydappCaller // Generic read-only contract binding to access the raw methods on
}

// NotarydappTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NotarydappTransactorRaw struct {
	Contract *NotarydappTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNotarydapp creates a new instance of Notarydapp, bound to a specific deployed contract.
func NewNotarydapp(address common.Address, backend bind.ContractBackend) (*Notarydapp, error) {
	contract, err := bindNotarydapp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Notarydapp{NotarydappCaller: NotarydappCaller{contract: contract}, NotarydappTransactor: NotarydappTransactor{contract: contract}, NotarydappFilterer: NotarydappFilterer{contract: contract}}, nil
}

// NewNotarydappCaller creates a new read-only instance of Notarydapp, bound to a specific deployed contract.
func NewNotarydappCaller(address common.Address, caller bind.ContractCaller) (*NotarydappCaller, error) {
	contract, err := bindNotarydapp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NotarydappCaller{contract: contract}, nil
}

// NewNotarydappTransactor creates a new write-only instance of Notarydapp, bound to a specific deployed contract.
func NewNotarydappTransactor(address common.Address, transactor bind.ContractTransactor) (*NotarydappTransactor, error) {
	contract, err := bindNotarydapp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NotarydappTransactor{contract: contract}, nil
}

// NewNotarydappFilterer creates a new log filterer instance of Notarydapp, bound to a specific deployed contract.
func NewNotarydappFilterer(address common.Address, filterer bind.ContractFilterer) (*NotarydappFilterer, error) {
	contract, err := bindNotarydapp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NotarydappFilterer{contract: contract}, nil
}

// bindNotarydapp binds a generic wrapper to an already deployed contract.
func bindNotarydapp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NotarydappABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Notarydapp *NotarydappRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Notarydapp.Contract.NotarydappCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Notarydapp *NotarydappRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Notarydapp.Contract.NotarydappTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Notarydapp *NotarydappRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Notarydapp.Contract.NotarydappTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Notarydapp *NotarydappCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Notarydapp.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Notarydapp *NotarydappTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Notarydapp.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Notarydapp *NotarydappTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Notarydapp.Contract.contract.Transact(opts, method, params...)
}

// GetDocumentByHash is a free data retrieval call binding the contract method 0x07fb5bc2.
//
// Solidity: function getDocumentByHash(bytes32 hash) constant returns(uint256, string)
func (_Notarydapp *NotarydappCaller) GetDocumentByHash(opts *bind.CallOpts, hash [32]byte) (*big.Int, string, error) {
	var (
		ret0 = new(*big.Int)
		ret1 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Notarydapp.contract.Call(opts, out, "getDocumentByHash", hash)
	return *ret0, *ret1, err
}

// GetDocumentByHash is a free data retrieval call binding the contract method 0x07fb5bc2.
//
// Solidity: function getDocumentByHash(bytes32 hash) constant returns(uint256, string)
func (_Notarydapp *NotarydappSession) GetDocumentByHash(hash [32]byte) (*big.Int, string, error) {
	return _Notarydapp.Contract.GetDocumentByHash(&_Notarydapp.CallOpts, hash)
}

// GetDocumentByHash is a free data retrieval call binding the contract method 0x07fb5bc2.
//
// Solidity: function getDocumentByHash(bytes32 hash) constant returns(uint256, string)
func (_Notarydapp *NotarydappCallerSession) GetDocumentByHash(hash [32]byte) (*big.Int, string, error) {
	return _Notarydapp.Contract.GetDocumentByHash(&_Notarydapp.CallOpts, hash)
}

// GetDocumentHashesByOwner is a free data retrieval call binding the contract method 0xd6cc7ac8.
//
// Solidity: function getDocumentHashesByOwner(string owner) constant returns(bytes32[])
func (_Notarydapp *NotarydappCaller) GetDocumentHashesByOwner(opts *bind.CallOpts, owner string) ([][32]byte, error) {
	var (
		ret0 = new([][32]byte)
	)
	out := ret0
	err := _Notarydapp.contract.Call(opts, out, "getDocumentHashesByOwner", owner)
	return *ret0, err
}

// GetDocumentHashesByOwner is a free data retrieval call binding the contract method 0xd6cc7ac8.
//
// Solidity: function getDocumentHashesByOwner(string owner) constant returns(bytes32[])
func (_Notarydapp *NotarydappSession) GetDocumentHashesByOwner(owner string) ([][32]byte, error) {
	return _Notarydapp.Contract.GetDocumentHashesByOwner(&_Notarydapp.CallOpts, owner)
}

// GetDocumentHashesByOwner is a free data retrieval call binding the contract method 0xd6cc7ac8.
//
// Solidity: function getDocumentHashesByOwner(string owner) constant returns(bytes32[])
func (_Notarydapp *NotarydappCallerSession) GetDocumentHashesByOwner(owner string) ([][32]byte, error) {
	return _Notarydapp.Contract.GetDocumentHashesByOwner(&_Notarydapp.CallOpts, owner)
}

// RegisterDocument is a paid mutator transaction binding the contract method 0xfa17185a.
//
// Solidity: function registerDocument(bytes32 documentHash, string documentOwner) returns(bool)
func (_Notarydapp *NotarydappTransactor) RegisterDocument(opts *bind.TransactOpts, documentHash [32]byte, documentOwner string) (*types.Transaction, error) {
	return _Notarydapp.contract.Transact(opts, "registerDocument", documentHash, documentOwner)
}

// RegisterDocument is a paid mutator transaction binding the contract method 0xfa17185a.
//
// Solidity: function registerDocument(bytes32 documentHash, string documentOwner) returns(bool)
func (_Notarydapp *NotarydappSession) RegisterDocument(documentHash [32]byte, documentOwner string) (*types.Transaction, error) {
	return _Notarydapp.Contract.RegisterDocument(&_Notarydapp.TransactOpts, documentHash, documentOwner)
}

// RegisterDocument is a paid mutator transaction binding the contract method 0xfa17185a.
//
// Solidity: function registerDocument(bytes32 documentHash, string documentOwner) returns(bool)
func (_Notarydapp *NotarydappTransactorSession) RegisterDocument(documentHash [32]byte, documentOwner string) (*types.Transaction, error) {
	return _Notarydapp.Contract.RegisterDocument(&_Notarydapp.TransactOpts, documentHash, documentOwner)
}
