// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package login

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

// LoginMetaData contains all meta data concerning the Login contract.
var LoginMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"GetData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_data\",\"type\":\"uint256\"}],\"name\":\"SetData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"data\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b506101718061001d5f395ff3fe608060405234801561000f575f80fd5b506004361061003f575f3560e01c8063213621171461004357806373d4a13a1461005f57806376b8e5281461007d575b5f80fd5b61005d600480360381019061005891906100e8565b61009b565b005b6100676100a4565b6040516100749190610122565b60405180910390f35b6100856100a9565b6040516100929190610122565b60405180910390f35b805f8190555050565b5f5481565b5f8054905090565b5f80fd5b5f819050919050565b6100c7816100b5565b81146100d1575f80fd5b50565b5f813590506100e2816100be565b92915050565b5f602082840312156100fd576100fc6100b1565b5b5f61010a848285016100d4565b91505092915050565b61011c816100b5565b82525050565b5f6020820190506101355f830184610113565b9291505056fea2646970667358221220b3d2c376b68cf40e73216d54fe725bb724eb36674180dc829238a793f5d9d04964736f6c63430008180033",
}

// LoginABI is the input ABI used to generate the binding from.
// Deprecated: Use LoginMetaData.ABI instead.
var LoginABI = LoginMetaData.ABI

// LoginBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use LoginMetaData.Bin instead.
var LoginBin = LoginMetaData.Bin

// DeployLogin deploys a new Ethereum contract, binding an instance of Login to it.
func DeployLogin(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Login, error) {
	parsed, err := LoginMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(LoginBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Login{LoginCaller: LoginCaller{contract: contract}, LoginTransactor: LoginTransactor{contract: contract}, LoginFilterer: LoginFilterer{contract: contract}}, nil
}

// Login is an auto generated Go binding around an Ethereum contract.
type Login struct {
	LoginCaller     // Read-only binding to the contract
	LoginTransactor // Write-only binding to the contract
	LoginFilterer   // Log filterer for contract events
}

// LoginCaller is an auto generated read-only Go binding around an Ethereum contract.
type LoginCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LoginTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LoginTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LoginFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LoginFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LoginSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LoginSession struct {
	Contract     *Login            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LoginCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LoginCallerSession struct {
	Contract *LoginCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// LoginTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LoginTransactorSession struct {
	Contract     *LoginTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LoginRaw is an auto generated low-level Go binding around an Ethereum contract.
type LoginRaw struct {
	Contract *Login // Generic contract binding to access the raw methods on
}

// LoginCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LoginCallerRaw struct {
	Contract *LoginCaller // Generic read-only contract binding to access the raw methods on
}

// LoginTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LoginTransactorRaw struct {
	Contract *LoginTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLogin creates a new instance of Login, bound to a specific deployed contract.
func NewLogin(address common.Address, backend bind.ContractBackend) (*Login, error) {
	contract, err := bindLogin(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Login{LoginCaller: LoginCaller{contract: contract}, LoginTransactor: LoginTransactor{contract: contract}, LoginFilterer: LoginFilterer{contract: contract}}, nil
}

// NewLoginCaller creates a new read-only instance of Login, bound to a specific deployed contract.
func NewLoginCaller(address common.Address, caller bind.ContractCaller) (*LoginCaller, error) {
	contract, err := bindLogin(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LoginCaller{contract: contract}, nil
}

// NewLoginTransactor creates a new write-only instance of Login, bound to a specific deployed contract.
func NewLoginTransactor(address common.Address, transactor bind.ContractTransactor) (*LoginTransactor, error) {
	contract, err := bindLogin(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LoginTransactor{contract: contract}, nil
}

// NewLoginFilterer creates a new log filterer instance of Login, bound to a specific deployed contract.
func NewLoginFilterer(address common.Address, filterer bind.ContractFilterer) (*LoginFilterer, error) {
	contract, err := bindLogin(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LoginFilterer{contract: contract}, nil
}

// bindLogin binds a generic wrapper to an already deployed contract.
func bindLogin(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LoginMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Login *LoginRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Login.Contract.LoginCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Login *LoginRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Login.Contract.LoginTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Login *LoginRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Login.Contract.LoginTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Login *LoginCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Login.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Login *LoginTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Login.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Login *LoginTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Login.Contract.contract.Transact(opts, method, params...)
}

// GetData is a free data retrieval call binding the contract method 0x76b8e528.
//
// Solidity: function GetData() view returns(uint256)
func (_Login *LoginCaller) GetData(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Login.contract.Call(opts, &out, "GetData")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetData is a free data retrieval call binding the contract method 0x76b8e528.
//
// Solidity: function GetData() view returns(uint256)
func (_Login *LoginSession) GetData() (*big.Int, error) {
	return _Login.Contract.GetData(&_Login.CallOpts)
}

// GetData is a free data retrieval call binding the contract method 0x76b8e528.
//
// Solidity: function GetData() view returns(uint256)
func (_Login *LoginCallerSession) GetData() (*big.Int, error) {
	return _Login.Contract.GetData(&_Login.CallOpts)
}

// Data is a free data retrieval call binding the contract method 0x73d4a13a.
//
// Solidity: function data() view returns(uint256)
func (_Login *LoginCaller) Data(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Login.contract.Call(opts, &out, "data")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Data is a free data retrieval call binding the contract method 0x73d4a13a.
//
// Solidity: function data() view returns(uint256)
func (_Login *LoginSession) Data() (*big.Int, error) {
	return _Login.Contract.Data(&_Login.CallOpts)
}

// Data is a free data retrieval call binding the contract method 0x73d4a13a.
//
// Solidity: function data() view returns(uint256)
func (_Login *LoginCallerSession) Data() (*big.Int, error) {
	return _Login.Contract.Data(&_Login.CallOpts)
}

// SetData is a paid mutator transaction binding the contract method 0x21362117.
//
// Solidity: function SetData(uint256 _data) returns()
func (_Login *LoginTransactor) SetData(opts *bind.TransactOpts, _data *big.Int) (*types.Transaction, error) {
	return _Login.contract.Transact(opts, "SetData", _data)
}

// SetData is a paid mutator transaction binding the contract method 0x21362117.
//
// Solidity: function SetData(uint256 _data) returns()
func (_Login *LoginSession) SetData(_data *big.Int) (*types.Transaction, error) {
	return _Login.Contract.SetData(&_Login.TransactOpts, _data)
}

// SetData is a paid mutator transaction binding the contract method 0x21362117.
//
// Solidity: function SetData(uint256 _data) returns()
func (_Login *LoginTransactorSession) SetData(_data *big.Int) (*types.Transaction, error) {
	return _Login.Contract.SetData(&_Login.TransactOpts, _data)
}
