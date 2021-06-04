// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// CharityABI is the input ABI used to generate the binding from.
const CharityABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"getBalance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"donate\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"}]"

// CharityFuncSigs maps the 4-byte function signature to its string representation.
var CharityFuncSigs = map[string]string{
	"ed88c68e": "donate()",
	"12065fe0": "getBalance()",
	"3ccfd60b": "withdraw()",
}

// CharityBin is the compiled bytecode used for deploying new contracts.
var CharityBin = "0x608060405234801561001057600080fd5b50610151806100206000396000f3fe6080604052600436106100565763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166312065fe081146100585780633ccfd60b1461007f578063ed88c68e14610087575b005b34801561006457600080fd5b5061006d61008f565b60408051918252519081900360200190f35b6100566100a2565b61005661010e565b3360009081526020819052604090205490565b3360008181526020819052604080822054905190929083908381818185875af1925050503d80600081146100f2576040519150601f19603f3d011682016040523d82523d6000602084013e6100f7565b606091505b505033600090815260208190526040812055505050565b33600090815260208190526040902080543401905556fea165627a7a72305820f1640cff381b229efedec5c706dae270a102d76e53caa2121ad5c0cb4049815b0029"

// DeployCharity deploys a new Ethereum contract, binding an instance of Charity to it.
func DeployCharity(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Charity, error) {
	parsed, err := abi.JSON(strings.NewReader(CharityABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CharityBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Charity{CharityCaller: CharityCaller{contract: contract}, CharityTransactor: CharityTransactor{contract: contract}, CharityFilterer: CharityFilterer{contract: contract}}, nil
}

// Charity is an auto generated Go binding around an Ethereum contract.
type Charity struct {
	CharityCaller     // Read-only binding to the contract
	CharityTransactor // Write-only binding to the contract
	CharityFilterer   // Log filterer for contract events
}

// CharityCaller is an auto generated read-only Go binding around an Ethereum contract.
type CharityCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CharityTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CharityTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CharityFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CharityFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CharitySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CharitySession struct {
	Contract     *Charity          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CharityCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CharityCallerSession struct {
	Contract *CharityCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// CharityTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CharityTransactorSession struct {
	Contract     *CharityTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// CharityRaw is an auto generated low-level Go binding around an Ethereum contract.
type CharityRaw struct {
	Contract *Charity // Generic contract binding to access the raw methods on
}

// CharityCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CharityCallerRaw struct {
	Contract *CharityCaller // Generic read-only contract binding to access the raw methods on
}

// CharityTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CharityTransactorRaw struct {
	Contract *CharityTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCharity creates a new instance of Charity, bound to a specific deployed contract.
func NewCharity(address common.Address, backend bind.ContractBackend) (*Charity, error) {
	contract, err := bindCharity(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Charity{CharityCaller: CharityCaller{contract: contract}, CharityTransactor: CharityTransactor{contract: contract}, CharityFilterer: CharityFilterer{contract: contract}}, nil
}

// NewCharityCaller creates a new read-only instance of Charity, bound to a specific deployed contract.
func NewCharityCaller(address common.Address, caller bind.ContractCaller) (*CharityCaller, error) {
	contract, err := bindCharity(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CharityCaller{contract: contract}, nil
}

// NewCharityTransactor creates a new write-only instance of Charity, bound to a specific deployed contract.
func NewCharityTransactor(address common.Address, transactor bind.ContractTransactor) (*CharityTransactor, error) {
	contract, err := bindCharity(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CharityTransactor{contract: contract}, nil
}

// NewCharityFilterer creates a new log filterer instance of Charity, bound to a specific deployed contract.
func NewCharityFilterer(address common.Address, filterer bind.ContractFilterer) (*CharityFilterer, error) {
	contract, err := bindCharity(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CharityFilterer{contract: contract}, nil
}

// bindCharity binds a generic wrapper to an already deployed contract.
func bindCharity(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CharityABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Charity *CharityRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Charity.Contract.CharityCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Charity *CharityRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Charity.Contract.CharityTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Charity *CharityRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Charity.Contract.CharityTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Charity *CharityCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Charity.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Charity *CharityTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Charity.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Charity *CharityTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Charity.Contract.contract.Transact(opts, method, params...)
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_Charity *CharityCaller) GetBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Charity.contract.Call(opts, &out, "getBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_Charity *CharitySession) GetBalance() (*big.Int, error) {
	return _Charity.Contract.GetBalance(&_Charity.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_Charity *CharityCallerSession) GetBalance() (*big.Int, error) {
	return _Charity.Contract.GetBalance(&_Charity.CallOpts)
}

// Donate is a paid mutator transaction binding the contract method 0xed88c68e.
//
// Solidity: function donate() payable returns()
func (_Charity *CharityTransactor) Donate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Charity.contract.Transact(opts, "donate")
}

// Donate is a paid mutator transaction binding the contract method 0xed88c68e.
//
// Solidity: function donate() payable returns()
func (_Charity *CharitySession) Donate() (*types.Transaction, error) {
	return _Charity.Contract.Donate(&_Charity.TransactOpts)
}

// Donate is a paid mutator transaction binding the contract method 0xed88c68e.
//
// Solidity: function donate() payable returns()
func (_Charity *CharityTransactorSession) Donate() (*types.Transaction, error) {
	return _Charity.Contract.Donate(&_Charity.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() payable returns()
func (_Charity *CharityTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Charity.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() payable returns()
func (_Charity *CharitySession) Withdraw() (*types.Transaction, error) {
	return _Charity.Contract.Withdraw(&_Charity.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() payable returns()
func (_Charity *CharityTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Charity.Contract.Withdraw(&_Charity.TransactOpts)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Charity *CharityTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Charity.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Charity *CharitySession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Charity.Contract.Fallback(&_Charity.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Charity *CharityTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Charity.Contract.Fallback(&_Charity.TransactOpts, calldata)
}
