// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"math/big"
	"strings"

	ethereum "github.com/rupayaproject/rupaya"
	"github.com/rupayaproject/rupaya/accounts/abi"
	"github.com/rupayaproject/rupaya/accounts/abi/bind"
	"github.com/rupayaproject/rupaya/common"
	"github.com/rupayaproject/rupaya/core/types"
	"github.com/rupayaproject/rupaya/event"
)

// IRRC21ABI is the input ABI used to generate the binding from.
const IRRC21ABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"estimateFee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"issuer\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"who\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"issuer\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Fee\",\"type\":\"event\"}]"

// IRRC21Bin is the compiled bytecode used for deploying new contracts.
const IRRC21Bin = `0x`

// DeployIRRC21 deploys a new Ethereum contract, binding an instance of IRRC21 to it.
func DeployIRRC21(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *IRRC21, error) {
	parsed, err := abi.JSON(strings.NewReader(IRRC21ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(IRRC21Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &IRRC21{IRRC21Caller: IRRC21Caller{contract: contract}, IRRC21Transactor: IRRC21Transactor{contract: contract}, IRRC21Filterer: IRRC21Filterer{contract: contract}}, nil
}

// IRRC21 is an auto generated Go binding around an Ethereum contract.
type IRRC21 struct {
	IRRC21Caller     // Read-only binding to the contract
	IRRC21Transactor // Write-only binding to the contract
	IRRC21Filterer   // Log filterer for contract events
}

// IRRC21Caller is an auto generated read-only Go binding around an Ethereum contract.
type IRRC21Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRRC21Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IRRC21Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRRC21Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IRRC21Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRRC21Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IRRC21Session struct {
	Contract     *IRRC21           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IRRC21CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IRRC21CallerSession struct {
	Contract *IRRC21Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IRRC21TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IRRC21TransactorSession struct {
	Contract     *IRRC21Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IRRC21Raw is an auto generated low-level Go binding around an Ethereum contract.
type IRRC21Raw struct {
	Contract *IRRC21 // Generic contract binding to access the raw methods on
}

// IRRC21CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IRRC21CallerRaw struct {
	Contract *IRRC21Caller // Generic read-only contract binding to access the raw methods on
}

// IRRC21TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IRRC21TransactorRaw struct {
	Contract *IRRC21Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIRRC21 creates a new instance of IRRC21, bound to a specific deployed contract.
func NewIRRC21(address common.Address, backend bind.ContractBackend) (*IRRC21, error) {
	contract, err := bindIRRC21(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IRRC21{IRRC21Caller: IRRC21Caller{contract: contract}, IRRC21Transactor: IRRC21Transactor{contract: contract}, IRRC21Filterer: IRRC21Filterer{contract: contract}}, nil
}

// NewIRRC21Caller creates a new read-only instance of IRRC21, bound to a specific deployed contract.
func NewIRRC21Caller(address common.Address, caller bind.ContractCaller) (*IRRC21Caller, error) {
	contract, err := bindIRRC21(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IRRC21Caller{contract: contract}, nil
}

// NewIRRC21Transactor creates a new write-only instance of IRRC21, bound to a specific deployed contract.
func NewIRRC21Transactor(address common.Address, transactor bind.ContractTransactor) (*IRRC21Transactor, error) {
	contract, err := bindIRRC21(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IRRC21Transactor{contract: contract}, nil
}

// NewIRRC21Filterer creates a new log filterer instance of IRRC21, bound to a specific deployed contract.
func NewIRRC21Filterer(address common.Address, filterer bind.ContractFilterer) (*IRRC21Filterer, error) {
	contract, err := bindIRRC21(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IRRC21Filterer{contract: contract}, nil
}

// bindIRRC21 binds a generic wrapper to an already deployed contract.
func bindIRRC21(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IRRC21ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IRRC21 *IRRC21Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IRRC21.Contract.IRRC21Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IRRC21 *IRRC21Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IRRC21.Contract.IRRC21Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IRRC21 *IRRC21Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IRRC21.Contract.IRRC21Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IRRC21 *IRRC21CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IRRC21.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IRRC21 *IRRC21TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IRRC21.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IRRC21 *IRRC21TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IRRC21.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(uint256)
func (_IRRC21 *IRRC21Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IRRC21.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(uint256)
func (_IRRC21 *IRRC21Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IRRC21.Contract.Allowance(&_IRRC21.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(uint256)
func (_IRRC21 *IRRC21CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IRRC21.Contract.Allowance(&_IRRC21.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(uint256)
func (_IRRC21 *IRRC21Caller) BalanceOf(opts *bind.CallOpts, who common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IRRC21.contract.Call(opts, out, "balanceOf", who)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(uint256)
func (_IRRC21 *IRRC21Session) BalanceOf(who common.Address) (*big.Int, error) {
	return _IRRC21.Contract.BalanceOf(&_IRRC21.CallOpts, who)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(uint256)
func (_IRRC21 *IRRC21CallerSession) BalanceOf(who common.Address) (*big.Int, error) {
	return _IRRC21.Contract.BalanceOf(&_IRRC21.CallOpts, who)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_IRRC21 *IRRC21Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _IRRC21.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_IRRC21 *IRRC21Session) Decimals() (uint8, error) {
	return _IRRC21.Contract.Decimals(&_IRRC21.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_IRRC21 *IRRC21CallerSession) Decimals() (uint8, error) {
	return _IRRC21.Contract.Decimals(&_IRRC21.CallOpts)
}

// EstimateFee is a free data retrieval call binding the contract method 0x127e8e4d.
//
// Solidity: function estimateFee(value uint256) constant returns(uint256)
func (_IRRC21 *IRRC21Caller) EstimateFee(opts *bind.CallOpts, value *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IRRC21.contract.Call(opts, out, "estimateFee", value)
	return *ret0, err
}

// EstimateFee is a free data retrieval call binding the contract method 0x127e8e4d.
//
// Solidity: function estimateFee(value uint256) constant returns(uint256)
func (_IRRC21 *IRRC21Session) EstimateFee(value *big.Int) (*big.Int, error) {
	return _IRRC21.Contract.EstimateFee(&_IRRC21.CallOpts, value)
}

// EstimateFee is a free data retrieval call binding the contract method 0x127e8e4d.
//
// Solidity: function estimateFee(value uint256) constant returns(uint256)
func (_IRRC21 *IRRC21CallerSession) EstimateFee(value *big.Int) (*big.Int, error) {
	return _IRRC21.Contract.EstimateFee(&_IRRC21.CallOpts, value)
}

// Issuer is a free data retrieval call binding the contract method 0x1d143848.
//
// Solidity: function issuer() constant returns(address)
func (_IRRC21 *IRRC21Caller) Issuer(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IRRC21.contract.Call(opts, out, "issuer")
	return *ret0, err
}

// Issuer is a free data retrieval call binding the contract method 0x1d143848.
//
// Solidity: function issuer() constant returns(address)
func (_IRRC21 *IRRC21Session) Issuer() (common.Address, error) {
	return _IRRC21.Contract.Issuer(&_IRRC21.CallOpts)
}

// Issuer is a free data retrieval call binding the contract method 0x1d143848.
//
// Solidity: function issuer() constant returns(address)
func (_IRRC21 *IRRC21CallerSession) Issuer() (common.Address, error) {
	return _IRRC21.Contract.Issuer(&_IRRC21.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_IRRC21 *IRRC21Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IRRC21.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_IRRC21 *IRRC21Session) TotalSupply() (*big.Int, error) {
	return _IRRC21.Contract.TotalSupply(&_IRRC21.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_IRRC21 *IRRC21CallerSession) TotalSupply() (*big.Int, error) {
	return _IRRC21.Contract.TotalSupply(&_IRRC21.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(bool)
func (_IRRC21 *IRRC21Transactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IRRC21.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(bool)
func (_IRRC21 *IRRC21Session) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IRRC21.Contract.Approve(&_IRRC21.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(bool)
func (_IRRC21 *IRRC21TransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IRRC21.Contract.Approve(&_IRRC21.TransactOpts, spender, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_IRRC21 *IRRC21Transactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IRRC21.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_IRRC21 *IRRC21Session) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IRRC21.Contract.Transfer(&_IRRC21.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_IRRC21 *IRRC21TransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IRRC21.Contract.Transfer(&_IRRC21.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(bool)
func (_IRRC21 *IRRC21Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IRRC21.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(bool)
func (_IRRC21 *IRRC21Session) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IRRC21.Contract.TransferFrom(&_IRRC21.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(bool)
func (_IRRC21 *IRRC21TransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IRRC21.Contract.TransferFrom(&_IRRC21.TransactOpts, from, to, value)
}

// IRRC21ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IRRC21 contract.
type IRRC21ApprovalIterator struct {
	Event *IRRC21Approval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IRRC21ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IRRC21Approval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IRRC21Approval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IRRC21ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IRRC21ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IRRC21Approval represents a Approval event raised by the IRRC21 contract.
type IRRC21Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(owner indexed address, spender indexed address, value uint256)
func (_IRRC21 *IRRC21Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IRRC21ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IRRC21.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IRRC21ApprovalIterator{contract: _IRRC21.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(owner indexed address, spender indexed address, value uint256)
func (_IRRC21 *IRRC21Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IRRC21Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IRRC21.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IRRC21Approval)
				if err := _IRRC21.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// IRRC21FeeIterator is returned from FilterFee and is used to iterate over the raw logs and unpacked data for Fee events raised by the IRRC21 contract.
type IRRC21FeeIterator struct {
	Event *IRRC21Fee // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IRRC21FeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IRRC21Fee)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IRRC21Fee)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IRRC21FeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IRRC21FeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IRRC21Fee represents a Fee event raised by the IRRC21 contract.
type IRRC21Fee struct {
	From   common.Address
	To     common.Address
	Issuer common.Address
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFee is a free log retrieval operation binding the contract event 0xfcf5b3276434181e3c48bd3fe569b8939808f11e845d4b963693b9d7dbd6dd99.
//
// Solidity: event Fee(from indexed address, to indexed address, issuer indexed address, value uint256)
func (_IRRC21 *IRRC21Filterer) FilterFee(opts *bind.FilterOpts, from []common.Address, to []common.Address, issuer []common.Address) (*IRRC21FeeIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem)
	}

	logs, sub, err := _IRRC21.contract.FilterLogs(opts, "Fee", fromRule, toRule, issuerRule)
	if err != nil {
		return nil, err
	}
	return &IRRC21FeeIterator{contract: _IRRC21.contract, event: "Fee", logs: logs, sub: sub}, nil
}

// WatchFee is a free log subscription operation binding the contract event 0xfcf5b3276434181e3c48bd3fe569b8939808f11e845d4b963693b9d7dbd6dd99.
//
// Solidity: event Fee(from indexed address, to indexed address, issuer indexed address, value uint256)
func (_IRRC21 *IRRC21Filterer) WatchFee(opts *bind.WatchOpts, sink chan<- *IRRC21Fee, from []common.Address, to []common.Address, issuer []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem)
	}

	logs, sub, err := _IRRC21.contract.WatchLogs(opts, "Fee", fromRule, toRule, issuerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IRRC21Fee)
				if err := _IRRC21.contract.UnpackLog(event, "Fee", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// IRRC21TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IRRC21 contract.
type IRRC21TransferIterator struct {
	Event *IRRC21Transfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IRRC21TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IRRC21Transfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IRRC21Transfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IRRC21TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IRRC21TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IRRC21Transfer represents a Transfer event raised by the IRRC21 contract.
type IRRC21Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(from indexed address, to indexed address, value uint256)
func (_IRRC21 *IRRC21Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IRRC21TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IRRC21.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IRRC21TransferIterator{contract: _IRRC21.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(from indexed address, to indexed address, value uint256)
func (_IRRC21 *IRRC21Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IRRC21Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IRRC21.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IRRC21Transfer)
				if err := _IRRC21.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// MyRRC21ABI is the input ABI used to generate the binding from.
const MyRRC21ABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"owners\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"estimateFee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"removeOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"issuer\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"revokeConfirmation\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minFee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"setMinFee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"address\"}],\"name\":\"confirmations\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"depositFee\",\"type\":\"uint256\"}],\"name\":\"setDepositFee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"pending\",\"type\":\"bool\"},{\"name\":\"executed\",\"type\":\"bool\"}],\"name\":\"getTransactionCount\",\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"addOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"isConfirmed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"getConfirmationCount\",\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"transactions\",\"outputs\":[{\"name\":\"isMintingTx\",\"type\":\"bool\"},{\"name\":\"destination\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\"},{\"name\":\"executed\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"WITHDRAW_FEE\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOwners\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"from\",\"type\":\"uint256\"},{\"name\":\"to\",\"type\":\"uint256\"},{\"name\":\"pending\",\"type\":\"bool\"},{\"name\":\"executed\",\"type\":\"bool\"}],\"name\":\"getTransactionIds\",\"outputs\":[{\"name\":\"_transactionIds\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"getConfirmations\",\"outputs\":[{\"name\":\"_confirmations\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"withdrawFee\",\"type\":\"uint256\"}],\"name\":\"setWithdrawFee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"transactionCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_required\",\"type\":\"uint256\"}],\"name\":\"changeRequirement\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"confirmTransaction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferContractIssuer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"destination\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"submitTransaction\",\"outputs\":[{\"name\":\"transactionId\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_OWNER_COUNT\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"required\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"DEPOSIT_FEE\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"replaceOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"executeTransaction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"burn\",\"outputs\":[{\"name\":\"transactionId\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_owners\",\"type\":\"address[]\"},{\"name\":\"_required\",\"type\":\"uint256\"},{\"name\":\"_name\",\"type\":\"string\"},{\"name\":\"_symbol\",\"type\":\"string\"},{\"name\":\"_decimals\",\"type\":\"uint8\"},{\"name\":\"cap\",\"type\":\"uint256\"},{\"name\":\"minFee\",\"type\":\"uint256\"},{\"name\":\"depositFee\",\"type\":\"uint256\"},{\"name\":\"withdrawFee\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"Confirmation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"Revocation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"Submission\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"transactionId\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"isMintingTx\",\"type\":\"bool\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"Execution\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"transactionId\",\"type\":\"uint256\"}],\"name\":\"ExecutionFailure\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnerAddition\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnerRemoval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"required\",\"type\":\"uint256\"}],\"name\":\"RequirementChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"issuer\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Fee\",\"type\":\"event\"}]"

// MyRRC21Bin is the compiled bytecode used for deploying new contracts.
const MyRRC21Bin = `0x6080604052600060085560006009553480156200001b57600080fd5b5060405162002626380380620026268339810160409081528151602080840151928401516060850151608086015160a087015160c088015160e08901516101008a0151958a018051988b019a909895019693959294919390929160009189918991899162000090916005919086019062000356565b508151620000a690600690602085019062000356565b506007805460ff191660ff92909216919091179055505089518960328211801590620000d25750818111155b8015620000de57508015155b8015620000ea57508115155b1515620000f657600080fd5b6200010b338864010000000062000240810204565b6200011f33640100000000620002ff810204565b620001338664010000000062000337810204565b600092505b8b518310156200020b57600c60008d858151811015156200015557fe5b6020908102909101810151600160a060020a031682528101919091526040016000205460ff16158015620001ab57508b838151811015156200019357fe5b90602001906020020151600160a060020a0316600014155b1515620001b757600080fd5b6001600c60008e86815181101515620001cc57fe5b602090810291909101810151600160a060020a03168252810191909152604001600020805460ff19169115159190911790556001929092019162000138565b8b516200022090600d9060208f0190620003db565b505050600e98909855600991909155600855506200048895505050505050565b600160a060020a03821615156200025657600080fd5b60045462000273908264010000000062001da76200033c82021704565b600455600160a060020a038216600090815260208190526040902054620002a9908264010000000062001da76200033c82021704565b600160a060020a0383166000818152602081815260408083209490945583518581529351929391927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a35050565b600160a060020a03811615156200031557600080fd5b60028054600160a060020a031916600160a060020a0392909216919091179055565b600155565b6000828201838110156200034f57600080fd5b9392505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200039957805160ff1916838001178555620003c9565b82800160010185558215620003c9579182015b82811115620003c9578251825591602001919060010190620003ac565b50620003d792915062000441565b5090565b82805482825590600052602060002090810192821562000433579160200282015b82811115620004335782518254600160a060020a031916600160a060020a03909116178255602090920191600190910190620003fc565b50620003d792915062000461565b6200045e91905b80821115620003d7576000815560010162000448565b90565b6200045e91905b80821115620003d7578054600160a060020a031916815560010162000468565b61218e80620004986000396000f3006080604052600436106101ed5763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663025e7c2781146101f257806306fdde0314610226578063095ea7b3146102b0578063127e8e4d146102e8578063173825d91461031257806318160ddd146103355780631d1438481461034a57806320ea8d861461035f57806323b872dd1461037757806324ec7590146103a15780632f54bf6e146103b6578063313ce567146103d757806331ac9920146104025780633411c81c1461041a578063490ae2101461043e57806354741525146104565780637065cb481461047557806370a0823114610496578063784547a7146104b75780638b51d13f146104cf57806395d89b41146104e75780639ace38c2146104fc5780639bff5ddb146105c2578063a0e67e2b146105d7578063a8abe69a1461063c578063a9059cbb14610661578063b5dc40c314610685578063b6ac642a1461069d578063b77bf600146106b5578063ba51a6df146106ca578063c01a8c84146106e2578063c28ce6ff146106fa578063c64274741461071b578063d74f8edd14610784578063dc8452cd14610799578063dd62ed3e146107ae578063de363e65146107d5578063e20056e6146107ea578063ee22610b14610811578063fe9d930314610829575b600080fd5b3480156101fe57600080fd5b5061020a600435610887565b60408051600160a060020a039092168252519081900360200190f35b34801561023257600080fd5b5061023b6108af565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561027557818101518382015260200161025d565b50505050905090810190601f1680156102a25780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156102bc57600080fd5b506102d4600160a060020a0360043516602435610946565b604080519115158252519081900360200190f35b3480156102f457600080fd5b50610300600435610a00565b60408051918252519081900360200190f35b34801561031e57600080fd5b50610333600160a060020a0360043516610a2c565b005b34801561034157600080fd5b50610300610ba3565b34801561035657600080fd5b5061020a610ba9565b34801561036b57600080fd5b50610333600435610bb8565b34801561038357600080fd5b506102d4600160a060020a0360043581169060243516604435610c72565b3480156103ad57600080fd5b50610300610db2565b3480156103c257600080fd5b506102d4600160a060020a0360043516610db8565b3480156103e357600080fd5b506103ec610dcd565b6040805160ff9092168252519081900360200190f35b34801561040e57600080fd5b50610333600435610dd6565b34801561042657600080fd5b506102d4600435600160a060020a0360243516610df9565b34801561044a57600080fd5b50610333600435610e19565b34801561046257600080fd5b5061030060043515156024351515610e3a565b34801561048157600080fd5b50610333600160a060020a0360043516610ea6565b3480156104a257600080fd5b50610300600160a060020a0360043516610fcb565b3480156104c357600080fd5b506102d4600435610fe6565b3480156104db57600080fd5b5061030060043561106a565b3480156104f357600080fd5b5061023b6110d9565b34801561050857600080fd5b5061051460043561113a565b604051808615151515815260200185600160a060020a0316600160a060020a031681526020018481526020018060200183151515158152602001828103825284818151815260200191508051906020019080838360005b8381101561058357818101518382015260200161056b565b50505050905090810190601f1680156105b05780820380516001836020036101000a031916815260200191505b50965050505050505060405180910390f35b3480156105ce57600080fd5b50610300611202565b3480156105e357600080fd5b506105ec611208565b60408051602080825283518183015283519192839290830191858101910280838360005b83811015610628578181015183820152602001610610565b505050509050019250505060405180910390f35b34801561064857600080fd5b506105ec60043560243560443515156064351515611269565b34801561066d57600080fd5b506102d4600160a060020a03600435166024356113a2565b34801561069157600080fd5b506105ec600435611465565b3480156106a957600080fd5b506103336004356115de565b3480156106c157600080fd5b506103006115ff565b3480156106d657600080fd5b50610333600435611605565b3480156106ee57600080fd5b50610333600435611684565b34801561070657600080fd5b50610333600160a060020a0360043516611752565b34801561072757600080fd5b50604080516020600460443581810135601f8101849004840285018401909552848452610300948235600160a060020a03169460248035953695946064949201919081908401838280828437509497506117869650505050505050565b34801561079057600080fd5b506103006117a7565b3480156107a557600080fd5b506103006117ac565b3480156107ba57600080fd5b50610300600160a060020a03600435811690602435166117b2565b3480156107e157600080fd5b506103006117dd565b3480156107f657600080fd5b50610333600160a060020a03600435811690602435166117e3565b34801561081d57600080fd5b5061033360043561196d565b34801561083557600080fd5b5060408051602060046024803582810135601f8101859004850286018501909652858552610300958335953695604494919390910191908190840183828082843750949750611b509650505050505050565b600d80548290811061089557fe5b600091825260209091200154600160a060020a0316905081565b60058054604080516020601f600260001961010060018816150201909516949094049384018190048102820181019092528281526060939092909183018282801561093b5780601f106109105761010080835404028352916020019161093b565b820191906000526020600020905b81548152906001019060200180831161091e57829003601f168201915b505050505090505b90565b6000600160a060020a038316151561095d57600080fd5b60015433600090815260208190526040902054101561097b57600080fd5b336000818152600360209081526040808320600160a060020a03888116855292529091208490556002546001546109b793929190911690611c87565b604080518381529051600160a060020a0385169133917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259181900360200190a350600192915050565b600154600090610a2690610a1a848463ffffffff611d7916565b9063ffffffff611da716565b92915050565b6000333014610a3a57600080fd5b600160a060020a0382166000908152600c6020526040902054829060ff161515610a6357600080fd5b600160a060020a0383166000908152600c60205260408120805460ff1916905591505b600d5460001901821015610b3e5782600160a060020a0316600d83815481101515610aad57fe5b600091825260209091200154600160a060020a03161415610b3357600d80546000198101908110610ada57fe5b600091825260209091200154600d8054600160a060020a039092169184908110610b0057fe5b9060005260206000200160006101000a815481600160a060020a030219169083600160a060020a03160217905550610b3e565b600190910190610a86565b600d80546000190190610b5190826120a1565b50600d54600e541115610b6a57600d54610b6a90611605565b604051600160a060020a038416907f8001553a916ef2f495d26a907cc54d96ed840d7bda71e73194bf5a9df7a76b9090600090a2505050565b60045490565b600254600160a060020a031690565b336000818152600c602052604090205460ff161515610bd657600080fd5b6000828152600b602090815260408083203380855292529091205483919060ff161515610c0257600080fd5b6000848152600a6020526040902060030154849060ff1615610c2357600080fd5b6000858152600b60209081526040808320338085529252808320805460ff191690555187927ff6a317157440607f36269043eb55f1287a5a19ba2216afeab88cd46cbcfb88e991a35050505050565b600080610c8a60015484611da790919063ffffffff16565b9050600160a060020a0384161515610ca157600080fd5b80831115610cae57600080fd5b600160a060020a0385166000908152600360209081526040808320338452909152902054811115610cde57600080fd5b600160a060020a0385166000908152600360209081526040808320338452909152902054610d12908263ffffffff611db916565b600160a060020a0386166000908152600360209081526040808320338452909152902055610d41858585611c87565b600254600154610d5e918791600160a060020a0390911690611c87565b6002546001546040805191825251600160a060020a039283169287169133917ffcf5b3276434181e3c48bd3fe569b8939808f11e845d4b963693b9d7dbd6dd999181900360200190a4506001949350505050565b60015490565b600c6020526000908152604090205460ff1681565b60075460ff1690565b600254600160a060020a03163314610ded57600080fd5b610df681611dd0565b50565b600b60209081526000928352604080842090915290825290205460ff1681565b610e21610ba9565b600160a060020a03163314610e3557600080fd5b600955565b6000805b600f54811015610e9f57838015610e6757506000818152600a602052604090206003015460ff16155b80610e8b5750828015610e8b57506000818152600a602052604090206003015460ff165b15610e97576001820191505b600101610e3e565b5092915050565b333014610eb257600080fd5b600160a060020a0381166000908152600c6020526040902054819060ff1615610eda57600080fd5b81600160a060020a0381161515610ef057600080fd5b600d80549050600101600e5460328211158015610f0d5750818111155b8015610f1857508015155b8015610f2357508115155b1515610f2e57600080fd5b600160a060020a0385166000818152600c6020526040808220805460ff19166001908117909155600d8054918201815583527fd7b6990105719101dabeb77144f2a3385c8033acd3af97e9423a695e81ad1eb501805473ffffffffffffffffffffffffffffffffffffffff191684179055517ff39e6e1eb0edcf53c221607b54b00cd28f3196fed0a24994dc308b8f611b682d9190a25050505050565b600160a060020a031660009081526020819052604090205490565b600080805b600d54811015611063576000848152600b60205260408120600d80549192918490811061101457fe5b6000918252602080832090910154600160a060020a0316835282019290925260400190205460ff1615611048576001820191505b600e5482141561105b5760019250611063565b600101610feb565b5050919050565b6000805b600d548110156110d3576000838152600b60205260408120600d80549192918490811061109757fe5b6000918252602080832090910154600160a060020a0316835282019290925260400190205460ff16156110cb576001820191505b60010161106e565b50919050565b60068054604080516020601f600260001961010060018816150201909516949094049384018190048102820181019092528281526060939092909183018282801561093b5780601f106109105761010080835404028352916020019161093b565b600a6020908152600091825260409182902080546001808301546002808501805488516101009582161586026000190190911692909204601f810188900488028301880190985287825260ff85169793909404600160a060020a031695919493909290918301828280156111ef5780601f106111c4576101008083540402835291602001916111ef565b820191906000526020600020905b8154815290600101906020018083116111d257829003601f168201915b5050506003909301549192505060ff1685565b60085481565b6060600d80548060200260200160405190810160405280929190818152602001828054801561093b57602002820191906000526020600020905b8154600160a060020a03168152600190910190602001808311611242575050505050905090565b606080600080600f5460405190808252806020026020018201604052801561129b578160200160208202803883390190505b50925060009150600090505b600f54811015611322578580156112d057506000818152600a602052604090206003015460ff16155b806112f457508480156112f457506000818152600a602052604090206003015460ff165b1561131a5780838381518110151561130857fe5b60209081029091010152600191909101905b6001016112a7565b87870360405190808252806020026020018201604052801561134e578160200160208202803883390190505b5093508790505b8681101561139757828181518110151561136b57fe5b906020019060200201518489830381518110151561138557fe5b60209081029091010152600101611355565b505050949350505050565b6000806113ba60015484611da790919063ffffffff16565b9050600160a060020a03841615156113d157600080fd5b808311156113de57600080fd5b6113e9338585611c87565b6000600154111561145b57600254600154611411913391600160a060020a0390911690611c87565b6002546001546040805191825251600160a060020a039283169287169133917ffcf5b3276434181e3c48bd3fe569b8939808f11e845d4b963693b9d7dbd6dd999181900360200190a45b5060019392505050565b606080600080600d8054905060405190808252806020026020018201604052801561149a578160200160208202803883390190505b50925060009150600090505b600d54811015611557576000858152600b60205260408120600d8054919291849081106114cf57fe5b6000918252602080832090910154600160a060020a0316835282019290925260400190205460ff161561154f57600d80548290811061150a57fe5b6000918252602090912001548351600160a060020a039091169084908490811061153057fe5b600160a060020a03909216602092830290910190910152600191909101905b6001016114a6565b81604051908082528060200260200182016040528015611581578160200160208202803883390190505b509350600090505b818110156115d657828181518110151561159f57fe5b9060200190602002015184828151811015156115b757fe5b600160a060020a03909216602092830290910190910152600101611589565b505050919050565b6115e6610ba9565b600160a060020a031633146115fa57600080fd5b600855565b600f5481565b33301461161157600080fd5b600d5481603282118015906116265750818111155b801561163157508015155b801561163c57508115155b151561164757600080fd5b600e8390556040805184815290517fa3f1ee9126a074d9326c682f561767f710e927faa811f7a99829d49dc421797a9181900360200190a1505050565b336000818152600c602052604090205460ff1615156116a257600080fd5b6000828152600a602052604090205482906101009004600160a060020a031615156116cc57600080fd5b6000838152600b602090815260408083203380855292529091205484919060ff16156116f757600080fd5b6000858152600b60209081526040808320338085529252808320805460ff191660011790555187927f4a504a94899432a9846e1aa406dceb1bcfd538bb839071d49d1e5e23f5be30ef91a361174b8561196d565b5050505050565b61175a610ba9565b600160a060020a0316331461176e57600080fd5b600160a060020a03811615610df657610df633611dd5565b60006117956001858585611e19565b90506117a081611684565b9392505050565b603281565b600e5481565b600160a060020a03918216600090815260036020908152604080832093909416825291909152205490565b60095481565b60003330146117f157600080fd5b600160a060020a0383166000908152600c6020526040902054839060ff16151561181a57600080fd5b600160a060020a0383166000908152600c6020526040902054839060ff161561184257600080fd5b600092505b600d548310156118d35784600160a060020a0316600d8481548110151561186a57fe5b600091825260209091200154600160a060020a031614156118c85783600d8481548110151561189557fe5b9060005260206000200160006101000a815481600160a060020a030219169083600160a060020a031602179055506118d3565b600190920191611847565b600160a060020a038086166000818152600c6020526040808220805460ff1990811690915593881682528082208054909416600117909355915190917f8001553a916ef2f495d26a907cc54d96ed840d7bda71e73194bf5a9df7a76b9091a2604051600160a060020a038516907ff39e6e1eb0edcf53c221607b54b00cd28f3196fed0a24994dc308b8f611b682d90600090a25050505050565b336000818152600c602052604081205490919060ff16151561198e57600080fd5b6000838152600b602090815260408083203380855292529091205484919060ff1615156119ba57600080fd5b6000858152600a6020526040902060030154859060ff16156119db57600080fd5b6119e486610fe6565b15611b48576000868152600a60205260409020805490955060ff1615611b48576009546001860154611a1b9163ffffffff611db916565b600186018190558554611a3d91610100909104600160a060020a031690611f29565b60006009541115611a5b57611a5b611a53610ba9565b600954611f29565b600385018054600160ff199091168117909155855481870154604080518481526020810183905260609181018281526002808c018054610100818a1615810260001901909116929092049484018590529504600160a060020a0316958c957f7ae9762dba14e21841a58a1cd988fb97578e27069c9bcc32d65682127141e730959194919390929091608083019084908015611b375780601f10611b0c57610100808354040283529160200191611b37565b820191906000526020600020905b815481529060010190602001808311611b1a57829003601f168201915b505094505050505060405180910390a35b505050505050565b60008080808511611b6057600080fd5b600854611b7490869063ffffffff611db916565b9150611b836000338487611e19565b9250611b8f3386611fd3565b60006008541115611bad57611bad611ba5610ba9565b600854611f29565b506000828152600a6020908152604080832060038101805460ff191660011790558151848152808401869052606092810183815288519382019390935287519194339488947f7ae9762dba14e21841a58a1cd988fb97578e27069c9bcc32d65682127141e73094929389938c939192916080840191908501908083838a5b83811015611c43578181015183820152602001611c2b565b50505050905090810190601f168015611c705780820380516001836020036101000a031916815260200191505b5094505050505060405180910390a3505092915050565b600160a060020a038316600090815260208190526040902054811115611cac57600080fd5b600160a060020a0382161515611cc157600080fd5b600160a060020a038316600090815260208190526040902054611cea908263ffffffff611db916565b600160a060020a038085166000908152602081905260408082209390935590841681522054611d1f908263ffffffff611da716565b600160a060020a038084166000818152602081815260409182902094909455805185815290519193928716927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef92918290030190a3505050565b600080831515611d8c5760009150610e9f565b50828202828482811515611d9c57fe5b04146117a057600080fd5b6000828201838110156117a057600080fd5b60008083831115611dc957600080fd5b5050900390565b600155565b600160a060020a0381161515611dea57600080fd5b6002805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b600083600160a060020a0381161515611e3157600080fd5b600f546040805160a0810182528815158152600160a060020a0388811660208084019182528385018a8152606085018a8152600060808701819052888152600a84529690962085518154945160ff199095169015151774ffffffffffffffffffffffffffffffffffffffff001916610100949095169390930293909317825591516001820155925180519496509193611ed092600285019201906120ca565b50608091909101516003909101805460ff1916911515919091179055600f8054600101905560405182907fc0ba8fe4b176c1714197d43b9cc6bcf797a4a7461c5fe8d0ef6e184ae7601e5190600090a250949350505050565b600160a060020a0382161515611f3e57600080fd5b600454611f51908263ffffffff611da716565b600455600160a060020a038216600090815260208190526040902054611f7d908263ffffffff611da716565b600160a060020a0383166000818152602081815260408083209490945583518581529351929391927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a35050565b600160a060020a0382161515611fe857600080fd5b600160a060020a03821660009081526020819052604090205481111561200d57600080fd5b600454612020908263ffffffff611db916565b600455600160a060020a03821660009081526020819052604090205461204c908263ffffffff611db916565b600160a060020a038316600081815260208181526040808320949094558351858152935191937fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef929081900390910190a35050565b8154818355818111156120c5576000838152602090206120c5918101908301612148565b505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061210b57805160ff1916838001178555612138565b82800160010185558215612138579182015b8281111561213857825182559160200191906001019061211d565b50612144929150612148565b5090565b61094391905b80821115612144576000815560010161214e5600a165627a7a723058204e2f2f29b40edc5c7244c800feff0a2f4381a4b0af1c9d5b66af858037761c6f0029`

// DeployMyRRC21 deploys a new Ethereum contract, binding an instance of MyRRC21 to it.
func DeployMyRRC21(auth *bind.TransactOpts, backend bind.ContractBackend, _owners []common.Address, _required *big.Int, _name string, _symbol string, _decimals uint8, cap *big.Int, minFee *big.Int, depositFee *big.Int, withdrawFee *big.Int) (common.Address, *types.Transaction, *MyRRC21, error) {
	parsed, err := abi.JSON(strings.NewReader(MyRRC21ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MyRRC21Bin), backend, _owners, _required, _name, _symbol, _decimals, cap, minFee, depositFee, withdrawFee)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MyRRC21{MyRRC21Caller: MyRRC21Caller{contract: contract}, MyRRC21Transactor: MyRRC21Transactor{contract: contract}, MyRRC21Filterer: MyRRC21Filterer{contract: contract}}, nil
}

// MyRRC21 is an auto generated Go binding around an Ethereum contract.
type MyRRC21 struct {
	MyRRC21Caller     // Read-only binding to the contract
	MyRRC21Transactor // Write-only binding to the contract
	MyRRC21Filterer   // Log filterer for contract events
}

// MyRRC21Caller is an auto generated read-only Go binding around an Ethereum contract.
type MyRRC21Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MyRRC21Transactor is an auto generated write-only Go binding around an Ethereum contract.
type MyRRC21Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MyRRC21Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MyRRC21Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MyRRC21Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MyRRC21Session struct {
	Contract     *MyRRC21          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MyRRC21CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MyRRC21CallerSession struct {
	Contract *MyRRC21Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// MyRRC21TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MyRRC21TransactorSession struct {
	Contract     *MyRRC21Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// MyRRC21Raw is an auto generated low-level Go binding around an Ethereum contract.
type MyRRC21Raw struct {
	Contract *MyRRC21 // Generic contract binding to access the raw methods on
}

// MyRRC21CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MyRRC21CallerRaw struct {
	Contract *MyRRC21Caller // Generic read-only contract binding to access the raw methods on
}

// MyRRC21TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MyRRC21TransactorRaw struct {
	Contract *MyRRC21Transactor // Generic write-only contract binding to access the raw methods on
}

// NewMyRRC21 creates a new instance of MyRRC21, bound to a specific deployed contract.
func NewMyRRC21(address common.Address, backend bind.ContractBackend) (*MyRRC21, error) {
	contract, err := bindMyRRC21(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MyRRC21{MyRRC21Caller: MyRRC21Caller{contract: contract}, MyRRC21Transactor: MyRRC21Transactor{contract: contract}, MyRRC21Filterer: MyRRC21Filterer{contract: contract}}, nil
}

// NewMyRRC21Caller creates a new read-only instance of MyRRC21, bound to a specific deployed contract.
func NewMyRRC21Caller(address common.Address, caller bind.ContractCaller) (*MyRRC21Caller, error) {
	contract, err := bindMyRRC21(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MyRRC21Caller{contract: contract}, nil
}

// NewMyRRC21Transactor creates a new write-only instance of MyRRC21, bound to a specific deployed contract.
func NewMyRRC21Transactor(address common.Address, transactor bind.ContractTransactor) (*MyRRC21Transactor, error) {
	contract, err := bindMyRRC21(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MyRRC21Transactor{contract: contract}, nil
}

// NewMyRRC21Filterer creates a new log filterer instance of MyRRC21, bound to a specific deployed contract.
func NewMyRRC21Filterer(address common.Address, filterer bind.ContractFilterer) (*MyRRC21Filterer, error) {
	contract, err := bindMyRRC21(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MyRRC21Filterer{contract: contract}, nil
}

// bindMyRRC21 binds a generic wrapper to an already deployed contract.
func bindMyRRC21(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MyRRC21ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MyRRC21 *MyRRC21Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MyRRC21.Contract.MyRRC21Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MyRRC21 *MyRRC21Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MyRRC21.Contract.MyRRC21Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MyRRC21 *MyRRC21Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MyRRC21.Contract.MyRRC21Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MyRRC21 *MyRRC21CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MyRRC21.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MyRRC21 *MyRRC21TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MyRRC21.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MyRRC21 *MyRRC21TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MyRRC21.Contract.contract.Transact(opts, method, params...)
}

// DEPOSITFEE is a free data retrieval call binding the contract method 0xde363e65.
//
// Solidity: function DEPOSIT_FEE() constant returns(uint256)
func (_MyRRC21 *MyRRC21Caller) DEPOSITFEE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MyRRC21.contract.Call(opts, out, "DEPOSIT_FEE")
	return *ret0, err
}

// DEPOSITFEE is a free data retrieval call binding the contract method 0xde363e65.
//
// Solidity: function DEPOSIT_FEE() constant returns(uint256)
func (_MyRRC21 *MyRRC21Session) DEPOSITFEE() (*big.Int, error) {
	return _MyRRC21.Contract.DEPOSITFEE(&_MyRRC21.CallOpts)
}

// DEPOSITFEE is a free data retrieval call binding the contract method 0xde363e65.
//
// Solidity: function DEPOSIT_FEE() constant returns(uint256)
func (_MyRRC21 *MyRRC21CallerSession) DEPOSITFEE() (*big.Int, error) {
	return _MyRRC21.Contract.DEPOSITFEE(&_MyRRC21.CallOpts)
}

// MAXOWNERCOUNT is a free data retrieval call binding the contract method 0xd74f8edd.
//
// Solidity: function MAX_OWNER_COUNT() constant returns(uint256)
func (_MyRRC21 *MyRRC21Caller) MAXOWNERCOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MyRRC21.contract.Call(opts, out, "MAX_OWNER_COUNT")
	return *ret0, err
}

// MAXOWNERCOUNT is a free data retrieval call binding the contract method 0xd74f8edd.
//
// Solidity: function MAX_OWNER_COUNT() constant returns(uint256)
func (_MyRRC21 *MyRRC21Session) MAXOWNERCOUNT() (*big.Int, error) {
	return _MyRRC21.Contract.MAXOWNERCOUNT(&_MyRRC21.CallOpts)
}

// MAXOWNERCOUNT is a free data retrieval call binding the contract method 0xd74f8edd.
//
// Solidity: function MAX_OWNER_COUNT() constant returns(uint256)
func (_MyRRC21 *MyRRC21CallerSession) MAXOWNERCOUNT() (*big.Int, error) {
	return _MyRRC21.Contract.MAXOWNERCOUNT(&_MyRRC21.CallOpts)
}

// WITHDRAWFEE is a free data retrieval call binding the contract method 0x9bff5ddb.
//
// Solidity: function WITHDRAW_FEE() constant returns(uint256)
func (_MyRRC21 *MyRRC21Caller) WITHDRAWFEE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MyRRC21.contract.Call(opts, out, "WITHDRAW_FEE")
	return *ret0, err
}

// WITHDRAWFEE is a free data retrieval call binding the contract method 0x9bff5ddb.
//
// Solidity: function WITHDRAW_FEE() constant returns(uint256)
func (_MyRRC21 *MyRRC21Session) WITHDRAWFEE() (*big.Int, error) {
	return _MyRRC21.Contract.WITHDRAWFEE(&_MyRRC21.CallOpts)
}

// WITHDRAWFEE is a free data retrieval call binding the contract method 0x9bff5ddb.
//
// Solidity: function WITHDRAW_FEE() constant returns(uint256)
func (_MyRRC21 *MyRRC21CallerSession) WITHDRAWFEE() (*big.Int, error) {
	return _MyRRC21.Contract.WITHDRAWFEE(&_MyRRC21.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(uint256)
func (_MyRRC21 *MyRRC21Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MyRRC21.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(uint256)
func (_MyRRC21 *MyRRC21Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _MyRRC21.Contract.Allowance(&_MyRRC21.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(uint256)
func (_MyRRC21 *MyRRC21CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _MyRRC21.Contract.Allowance(&_MyRRC21.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(owner address) constant returns(uint256)
func (_MyRRC21 *MyRRC21Caller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MyRRC21.contract.Call(opts, out, "balanceOf", owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(owner address) constant returns(uint256)
func (_MyRRC21 *MyRRC21Session) BalanceOf(owner common.Address) (*big.Int, error) {
	return _MyRRC21.Contract.BalanceOf(&_MyRRC21.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(owner address) constant returns(uint256)
func (_MyRRC21 *MyRRC21CallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _MyRRC21.Contract.BalanceOf(&_MyRRC21.CallOpts, owner)
}

// Confirmations is a free data retrieval call binding the contract method 0x3411c81c.
//
// Solidity: function confirmations( uint256,  address) constant returns(bool)
func (_MyRRC21 *MyRRC21Caller) Confirmations(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _MyRRC21.contract.Call(opts, out, "confirmations", arg0, arg1)
	return *ret0, err
}

// Confirmations is a free data retrieval call binding the contract method 0x3411c81c.
//
// Solidity: function confirmations( uint256,  address) constant returns(bool)
func (_MyRRC21 *MyRRC21Session) Confirmations(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _MyRRC21.Contract.Confirmations(&_MyRRC21.CallOpts, arg0, arg1)
}

// Confirmations is a free data retrieval call binding the contract method 0x3411c81c.
//
// Solidity: function confirmations( uint256,  address) constant returns(bool)
func (_MyRRC21 *MyRRC21CallerSession) Confirmations(arg0 *big.Int, arg1 common.Address) (bool, error) {
	return _MyRRC21.Contract.Confirmations(&_MyRRC21.CallOpts, arg0, arg1)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_MyRRC21 *MyRRC21Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _MyRRC21.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_MyRRC21 *MyRRC21Session) Decimals() (uint8, error) {
	return _MyRRC21.Contract.Decimals(&_MyRRC21.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_MyRRC21 *MyRRC21CallerSession) Decimals() (uint8, error) {
	return _MyRRC21.Contract.Decimals(&_MyRRC21.CallOpts)
}

// EstimateFee is a free data retrieval call binding the contract method 0x127e8e4d.
//
// Solidity: function estimateFee(value uint256) constant returns(uint256)
func (_MyRRC21 *MyRRC21Caller) EstimateFee(opts *bind.CallOpts, value *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MyRRC21.contract.Call(opts, out, "estimateFee", value)
	return *ret0, err
}

// EstimateFee is a free data retrieval call binding the contract method 0x127e8e4d.
//
// Solidity: function estimateFee(value uint256) constant returns(uint256)
func (_MyRRC21 *MyRRC21Session) EstimateFee(value *big.Int) (*big.Int, error) {
	return _MyRRC21.Contract.EstimateFee(&_MyRRC21.CallOpts, value)
}

// EstimateFee is a free data retrieval call binding the contract method 0x127e8e4d.
//
// Solidity: function estimateFee(value uint256) constant returns(uint256)
func (_MyRRC21 *MyRRC21CallerSession) EstimateFee(value *big.Int) (*big.Int, error) {
	return _MyRRC21.Contract.EstimateFee(&_MyRRC21.CallOpts, value)
}

// GetConfirmationCount is a free data retrieval call binding the contract method 0x8b51d13f.
//
// Solidity: function getConfirmationCount(transactionId uint256) constant returns(count uint256)
func (_MyRRC21 *MyRRC21Caller) GetConfirmationCount(opts *bind.CallOpts, transactionId *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MyRRC21.contract.Call(opts, out, "getConfirmationCount", transactionId)
	return *ret0, err
}

// GetConfirmationCount is a free data retrieval call binding the contract method 0x8b51d13f.
//
// Solidity: function getConfirmationCount(transactionId uint256) constant returns(count uint256)
func (_MyRRC21 *MyRRC21Session) GetConfirmationCount(transactionId *big.Int) (*big.Int, error) {
	return _MyRRC21.Contract.GetConfirmationCount(&_MyRRC21.CallOpts, transactionId)
}

// GetConfirmationCount is a free data retrieval call binding the contract method 0x8b51d13f.
//
// Solidity: function getConfirmationCount(transactionId uint256) constant returns(count uint256)
func (_MyRRC21 *MyRRC21CallerSession) GetConfirmationCount(transactionId *big.Int) (*big.Int, error) {
	return _MyRRC21.Contract.GetConfirmationCount(&_MyRRC21.CallOpts, transactionId)
}

// GetConfirmations is a free data retrieval call binding the contract method 0xb5dc40c3.
//
// Solidity: function getConfirmations(transactionId uint256) constant returns(_confirmations address[])
func (_MyRRC21 *MyRRC21Caller) GetConfirmations(opts *bind.CallOpts, transactionId *big.Int) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _MyRRC21.contract.Call(opts, out, "getConfirmations", transactionId)
	return *ret0, err
}

// GetConfirmations is a free data retrieval call binding the contract method 0xb5dc40c3.
//
// Solidity: function getConfirmations(transactionId uint256) constant returns(_confirmations address[])
func (_MyRRC21 *MyRRC21Session) GetConfirmations(transactionId *big.Int) ([]common.Address, error) {
	return _MyRRC21.Contract.GetConfirmations(&_MyRRC21.CallOpts, transactionId)
}

// GetConfirmations is a free data retrieval call binding the contract method 0xb5dc40c3.
//
// Solidity: function getConfirmations(transactionId uint256) constant returns(_confirmations address[])
func (_MyRRC21 *MyRRC21CallerSession) GetConfirmations(transactionId *big.Int) ([]common.Address, error) {
	return _MyRRC21.Contract.GetConfirmations(&_MyRRC21.CallOpts, transactionId)
}

// GetOwners is a free data retrieval call binding the contract method 0xa0e67e2b.
//
// Solidity: function getOwners() constant returns(address[])
func (_MyRRC21 *MyRRC21Caller) GetOwners(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _MyRRC21.contract.Call(opts, out, "getOwners")
	return *ret0, err
}

// GetOwners is a free data retrieval call binding the contract method 0xa0e67e2b.
//
// Solidity: function getOwners() constant returns(address[])
func (_MyRRC21 *MyRRC21Session) GetOwners() ([]common.Address, error) {
	return _MyRRC21.Contract.GetOwners(&_MyRRC21.CallOpts)
}

// GetOwners is a free data retrieval call binding the contract method 0xa0e67e2b.
//
// Solidity: function getOwners() constant returns(address[])
func (_MyRRC21 *MyRRC21CallerSession) GetOwners() ([]common.Address, error) {
	return _MyRRC21.Contract.GetOwners(&_MyRRC21.CallOpts)
}

// GetTransactionCount is a free data retrieval call binding the contract method 0x54741525.
//
// Solidity: function getTransactionCount(pending bool, executed bool) constant returns(count uint256)
func (_MyRRC21 *MyRRC21Caller) GetTransactionCount(opts *bind.CallOpts, pending bool, executed bool) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MyRRC21.contract.Call(opts, out, "getTransactionCount", pending, executed)
	return *ret0, err
}

// GetTransactionCount is a free data retrieval call binding the contract method 0x54741525.
//
// Solidity: function getTransactionCount(pending bool, executed bool) constant returns(count uint256)
func (_MyRRC21 *MyRRC21Session) GetTransactionCount(pending bool, executed bool) (*big.Int, error) {
	return _MyRRC21.Contract.GetTransactionCount(&_MyRRC21.CallOpts, pending, executed)
}

// GetTransactionCount is a free data retrieval call binding the contract method 0x54741525.
//
// Solidity: function getTransactionCount(pending bool, executed bool) constant returns(count uint256)
func (_MyRRC21 *MyRRC21CallerSession) GetTransactionCount(pending bool, executed bool) (*big.Int, error) {
	return _MyRRC21.Contract.GetTransactionCount(&_MyRRC21.CallOpts, pending, executed)
}

// GetTransactionIds is a free data retrieval call binding the contract method 0xa8abe69a.
//
// Solidity: function getTransactionIds(from uint256, to uint256, pending bool, executed bool) constant returns(_transactionIds uint256[])
func (_MyRRC21 *MyRRC21Caller) GetTransactionIds(opts *bind.CallOpts, from *big.Int, to *big.Int, pending bool, executed bool) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _MyRRC21.contract.Call(opts, out, "getTransactionIds", from, to, pending, executed)
	return *ret0, err
}

// GetTransactionIds is a free data retrieval call binding the contract method 0xa8abe69a.
//
// Solidity: function getTransactionIds(from uint256, to uint256, pending bool, executed bool) constant returns(_transactionIds uint256[])
func (_MyRRC21 *MyRRC21Session) GetTransactionIds(from *big.Int, to *big.Int, pending bool, executed bool) ([]*big.Int, error) {
	return _MyRRC21.Contract.GetTransactionIds(&_MyRRC21.CallOpts, from, to, pending, executed)
}

// GetTransactionIds is a free data retrieval call binding the contract method 0xa8abe69a.
//
// Solidity: function getTransactionIds(from uint256, to uint256, pending bool, executed bool) constant returns(_transactionIds uint256[])
func (_MyRRC21 *MyRRC21CallerSession) GetTransactionIds(from *big.Int, to *big.Int, pending bool, executed bool) ([]*big.Int, error) {
	return _MyRRC21.Contract.GetTransactionIds(&_MyRRC21.CallOpts, from, to, pending, executed)
}

// IsConfirmed is a free data retrieval call binding the contract method 0x784547a7.
//
// Solidity: function isConfirmed(transactionId uint256) constant returns(bool)
func (_MyRRC21 *MyRRC21Caller) IsConfirmed(opts *bind.CallOpts, transactionId *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _MyRRC21.contract.Call(opts, out, "isConfirmed", transactionId)
	return *ret0, err
}

// IsConfirmed is a free data retrieval call binding the contract method 0x784547a7.
//
// Solidity: function isConfirmed(transactionId uint256) constant returns(bool)
func (_MyRRC21 *MyRRC21Session) IsConfirmed(transactionId *big.Int) (bool, error) {
	return _MyRRC21.Contract.IsConfirmed(&_MyRRC21.CallOpts, transactionId)
}

// IsConfirmed is a free data retrieval call binding the contract method 0x784547a7.
//
// Solidity: function isConfirmed(transactionId uint256) constant returns(bool)
func (_MyRRC21 *MyRRC21CallerSession) IsConfirmed(transactionId *big.Int) (bool, error) {
	return _MyRRC21.Contract.IsConfirmed(&_MyRRC21.CallOpts, transactionId)
}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner( address) constant returns(bool)
func (_MyRRC21 *MyRRC21Caller) IsOwner(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _MyRRC21.contract.Call(opts, out, "isOwner", arg0)
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner( address) constant returns(bool)
func (_MyRRC21 *MyRRC21Session) IsOwner(arg0 common.Address) (bool, error) {
	return _MyRRC21.Contract.IsOwner(&_MyRRC21.CallOpts, arg0)
}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner( address) constant returns(bool)
func (_MyRRC21 *MyRRC21CallerSession) IsOwner(arg0 common.Address) (bool, error) {
	return _MyRRC21.Contract.IsOwner(&_MyRRC21.CallOpts, arg0)
}

// Issuer is a free data retrieval call binding the contract method 0x1d143848.
//
// Solidity: function issuer() constant returns(address)
func (_MyRRC21 *MyRRC21Caller) Issuer(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MyRRC21.contract.Call(opts, out, "issuer")
	return *ret0, err
}

// Issuer is a free data retrieval call binding the contract method 0x1d143848.
//
// Solidity: function issuer() constant returns(address)
func (_MyRRC21 *MyRRC21Session) Issuer() (common.Address, error) {
	return _MyRRC21.Contract.Issuer(&_MyRRC21.CallOpts)
}

// Issuer is a free data retrieval call binding the contract method 0x1d143848.
//
// Solidity: function issuer() constant returns(address)
func (_MyRRC21 *MyRRC21CallerSession) Issuer() (common.Address, error) {
	return _MyRRC21.Contract.Issuer(&_MyRRC21.CallOpts)
}

// MinFee is a free data retrieval call binding the contract method 0x24ec7590.
//
// Solidity: function minFee() constant returns(uint256)
func (_MyRRC21 *MyRRC21Caller) MinFee(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MyRRC21.contract.Call(opts, out, "minFee")
	return *ret0, err
}

// MinFee is a free data retrieval call binding the contract method 0x24ec7590.
//
// Solidity: function minFee() constant returns(uint256)
func (_MyRRC21 *MyRRC21Session) MinFee() (*big.Int, error) {
	return _MyRRC21.Contract.MinFee(&_MyRRC21.CallOpts)
}

// MinFee is a free data retrieval call binding the contract method 0x24ec7590.
//
// Solidity: function minFee() constant returns(uint256)
func (_MyRRC21 *MyRRC21CallerSession) MinFee() (*big.Int, error) {
	return _MyRRC21.Contract.MinFee(&_MyRRC21.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_MyRRC21 *MyRRC21Caller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _MyRRC21.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_MyRRC21 *MyRRC21Session) Name() (string, error) {
	return _MyRRC21.Contract.Name(&_MyRRC21.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_MyRRC21 *MyRRC21CallerSession) Name() (string, error) {
	return _MyRRC21.Contract.Name(&_MyRRC21.CallOpts)
}

// Owners is a free data retrieval call binding the contract method 0x025e7c27.
//
// Solidity: function owners( uint256) constant returns(address)
func (_MyRRC21 *MyRRC21Caller) Owners(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MyRRC21.contract.Call(opts, out, "owners", arg0)
	return *ret0, err
}

// Owners is a free data retrieval call binding the contract method 0x025e7c27.
//
// Solidity: function owners( uint256) constant returns(address)
func (_MyRRC21 *MyRRC21Session) Owners(arg0 *big.Int) (common.Address, error) {
	return _MyRRC21.Contract.Owners(&_MyRRC21.CallOpts, arg0)
}

// Owners is a free data retrieval call binding the contract method 0x025e7c27.
//
// Solidity: function owners( uint256) constant returns(address)
func (_MyRRC21 *MyRRC21CallerSession) Owners(arg0 *big.Int) (common.Address, error) {
	return _MyRRC21.Contract.Owners(&_MyRRC21.CallOpts, arg0)
}

// Required is a free data retrieval call binding the contract method 0xdc8452cd.
//
// Solidity: function required() constant returns(uint256)
func (_MyRRC21 *MyRRC21Caller) Required(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MyRRC21.contract.Call(opts, out, "required")
	return *ret0, err
}

// Required is a free data retrieval call binding the contract method 0xdc8452cd.
//
// Solidity: function required() constant returns(uint256)
func (_MyRRC21 *MyRRC21Session) Required() (*big.Int, error) {
	return _MyRRC21.Contract.Required(&_MyRRC21.CallOpts)
}

// Required is a free data retrieval call binding the contract method 0xdc8452cd.
//
// Solidity: function required() constant returns(uint256)
func (_MyRRC21 *MyRRC21CallerSession) Required() (*big.Int, error) {
	return _MyRRC21.Contract.Required(&_MyRRC21.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_MyRRC21 *MyRRC21Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _MyRRC21.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_MyRRC21 *MyRRC21Session) Symbol() (string, error) {
	return _MyRRC21.Contract.Symbol(&_MyRRC21.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_MyRRC21 *MyRRC21CallerSession) Symbol() (string, error) {
	return _MyRRC21.Contract.Symbol(&_MyRRC21.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_MyRRC21 *MyRRC21Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MyRRC21.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_MyRRC21 *MyRRC21Session) TotalSupply() (*big.Int, error) {
	return _MyRRC21.Contract.TotalSupply(&_MyRRC21.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_MyRRC21 *MyRRC21CallerSession) TotalSupply() (*big.Int, error) {
	return _MyRRC21.Contract.TotalSupply(&_MyRRC21.CallOpts)
}

// TransactionCount is a free data retrieval call binding the contract method 0xb77bf600.
//
// Solidity: function transactionCount() constant returns(uint256)
func (_MyRRC21 *MyRRC21Caller) TransactionCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _MyRRC21.contract.Call(opts, out, "transactionCount")
	return *ret0, err
}

// TransactionCount is a free data retrieval call binding the contract method 0xb77bf600.
//
// Solidity: function transactionCount() constant returns(uint256)
func (_MyRRC21 *MyRRC21Session) TransactionCount() (*big.Int, error) {
	return _MyRRC21.Contract.TransactionCount(&_MyRRC21.CallOpts)
}

// TransactionCount is a free data retrieval call binding the contract method 0xb77bf600.
//
// Solidity: function transactionCount() constant returns(uint256)
func (_MyRRC21 *MyRRC21CallerSession) TransactionCount() (*big.Int, error) {
	return _MyRRC21.Contract.TransactionCount(&_MyRRC21.CallOpts)
}

// Transactions is a free data retrieval call binding the contract method 0x9ace38c2.
//
// Solidity: function transactions( uint256) constant returns(isMintingTx bool, destination address, value uint256, data bytes, executed bool)
func (_MyRRC21 *MyRRC21Caller) Transactions(opts *bind.CallOpts, arg0 *big.Int) (struct {
	IsMintingTx bool
	Destination common.Address
	Value       *big.Int
	Data        []byte
	Executed    bool
}, error) {
	ret := new(struct {
		IsMintingTx bool
		Destination common.Address
		Value       *big.Int
		Data        []byte
		Executed    bool
	})
	out := ret
	err := _MyRRC21.contract.Call(opts, out, "transactions", arg0)
	return *ret, err
}

// Transactions is a free data retrieval call binding the contract method 0x9ace38c2.
//
// Solidity: function transactions( uint256) constant returns(isMintingTx bool, destination address, value uint256, data bytes, executed bool)
func (_MyRRC21 *MyRRC21Session) Transactions(arg0 *big.Int) (struct {
	IsMintingTx bool
	Destination common.Address
	Value       *big.Int
	Data        []byte
	Executed    bool
}, error) {
	return _MyRRC21.Contract.Transactions(&_MyRRC21.CallOpts, arg0)
}

// Transactions is a free data retrieval call binding the contract method 0x9ace38c2.
//
// Solidity: function transactions( uint256) constant returns(isMintingTx bool, destination address, value uint256, data bytes, executed bool)
func (_MyRRC21 *MyRRC21CallerSession) Transactions(arg0 *big.Int) (struct {
	IsMintingTx bool
	Destination common.Address
	Value       *big.Int
	Data        []byte
	Executed    bool
}, error) {
	return _MyRRC21.Contract.Transactions(&_MyRRC21.CallOpts, arg0)
}

// AddOwner is a paid mutator transaction binding the contract method 0x7065cb48.
//
// Solidity: function addOwner(owner address) returns()
func (_MyRRC21 *MyRRC21Transactor) AddOwner(opts *bind.TransactOpts, owner common.Address) (*types.Transaction, error) {
	return _MyRRC21.contract.Transact(opts, "addOwner", owner)
}

// AddOwner is a paid mutator transaction binding the contract method 0x7065cb48.
//
// Solidity: function addOwner(owner address) returns()
func (_MyRRC21 *MyRRC21Session) AddOwner(owner common.Address) (*types.Transaction, error) {
	return _MyRRC21.Contract.AddOwner(&_MyRRC21.TransactOpts, owner)
}

// AddOwner is a paid mutator transaction binding the contract method 0x7065cb48.
//
// Solidity: function addOwner(owner address) returns()
func (_MyRRC21 *MyRRC21TransactorSession) AddOwner(owner common.Address) (*types.Transaction, error) {
	return _MyRRC21.Contract.AddOwner(&_MyRRC21.TransactOpts, owner)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(bool)
func (_MyRRC21 *MyRRC21Transactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _MyRRC21.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(bool)
func (_MyRRC21 *MyRRC21Session) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _MyRRC21.Contract.Approve(&_MyRRC21.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(bool)
func (_MyRRC21 *MyRRC21TransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _MyRRC21.Contract.Approve(&_MyRRC21.TransactOpts, spender, value)
}

// Burn is a paid mutator transaction binding the contract method 0xfe9d9303.
//
// Solidity: function burn(value uint256, data bytes) returns(transactionId uint256)
func (_MyRRC21 *MyRRC21Transactor) Burn(opts *bind.TransactOpts, value *big.Int, data []byte) (*types.Transaction, error) {
	return _MyRRC21.contract.Transact(opts, "burn", value, data)
}

// Burn is a paid mutator transaction binding the contract method 0xfe9d9303.
//
// Solidity: function burn(value uint256, data bytes) returns(transactionId uint256)
func (_MyRRC21 *MyRRC21Session) Burn(value *big.Int, data []byte) (*types.Transaction, error) {
	return _MyRRC21.Contract.Burn(&_MyRRC21.TransactOpts, value, data)
}

// Burn is a paid mutator transaction binding the contract method 0xfe9d9303.
//
// Solidity: function burn(value uint256, data bytes) returns(transactionId uint256)
func (_MyRRC21 *MyRRC21TransactorSession) Burn(value *big.Int, data []byte) (*types.Transaction, error) {
	return _MyRRC21.Contract.Burn(&_MyRRC21.TransactOpts, value, data)
}

// ChangeRequirement is a paid mutator transaction binding the contract method 0xba51a6df.
//
// Solidity: function changeRequirement(_required uint256) returns()
func (_MyRRC21 *MyRRC21Transactor) ChangeRequirement(opts *bind.TransactOpts, _required *big.Int) (*types.Transaction, error) {
	return _MyRRC21.contract.Transact(opts, "changeRequirement", _required)
}

// ChangeRequirement is a paid mutator transaction binding the contract method 0xba51a6df.
//
// Solidity: function changeRequirement(_required uint256) returns()
func (_MyRRC21 *MyRRC21Session) ChangeRequirement(_required *big.Int) (*types.Transaction, error) {
	return _MyRRC21.Contract.ChangeRequirement(&_MyRRC21.TransactOpts, _required)
}

// ChangeRequirement is a paid mutator transaction binding the contract method 0xba51a6df.
//
// Solidity: function changeRequirement(_required uint256) returns()
func (_MyRRC21 *MyRRC21TransactorSession) ChangeRequirement(_required *big.Int) (*types.Transaction, error) {
	return _MyRRC21.Contract.ChangeRequirement(&_MyRRC21.TransactOpts, _required)
}

// ConfirmTransaction is a paid mutator transaction binding the contract method 0xc01a8c84.
//
// Solidity: function confirmTransaction(transactionId uint256) returns()
func (_MyRRC21 *MyRRC21Transactor) ConfirmTransaction(opts *bind.TransactOpts, transactionId *big.Int) (*types.Transaction, error) {
	return _MyRRC21.contract.Transact(opts, "confirmTransaction", transactionId)
}

// ConfirmTransaction is a paid mutator transaction binding the contract method 0xc01a8c84.
//
// Solidity: function confirmTransaction(transactionId uint256) returns()
func (_MyRRC21 *MyRRC21Session) ConfirmTransaction(transactionId *big.Int) (*types.Transaction, error) {
	return _MyRRC21.Contract.ConfirmTransaction(&_MyRRC21.TransactOpts, transactionId)
}

// ConfirmTransaction is a paid mutator transaction binding the contract method 0xc01a8c84.
//
// Solidity: function confirmTransaction(transactionId uint256) returns()
func (_MyRRC21 *MyRRC21TransactorSession) ConfirmTransaction(transactionId *big.Int) (*types.Transaction, error) {
	return _MyRRC21.Contract.ConfirmTransaction(&_MyRRC21.TransactOpts, transactionId)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0xee22610b.
//
// Solidity: function executeTransaction(transactionId uint256) returns()
func (_MyRRC21 *MyRRC21Transactor) ExecuteTransaction(opts *bind.TransactOpts, transactionId *big.Int) (*types.Transaction, error) {
	return _MyRRC21.contract.Transact(opts, "executeTransaction", transactionId)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0xee22610b.
//
// Solidity: function executeTransaction(transactionId uint256) returns()
func (_MyRRC21 *MyRRC21Session) ExecuteTransaction(transactionId *big.Int) (*types.Transaction, error) {
	return _MyRRC21.Contract.ExecuteTransaction(&_MyRRC21.TransactOpts, transactionId)
}

// ExecuteTransaction is a paid mutator transaction binding the contract method 0xee22610b.
//
// Solidity: function executeTransaction(transactionId uint256) returns()
func (_MyRRC21 *MyRRC21TransactorSession) ExecuteTransaction(transactionId *big.Int) (*types.Transaction, error) {
	return _MyRRC21.Contract.ExecuteTransaction(&_MyRRC21.TransactOpts, transactionId)
}

// RemoveOwner is a paid mutator transaction binding the contract method 0x173825d9.
//
// Solidity: function removeOwner(owner address) returns()
func (_MyRRC21 *MyRRC21Transactor) RemoveOwner(opts *bind.TransactOpts, owner common.Address) (*types.Transaction, error) {
	return _MyRRC21.contract.Transact(opts, "removeOwner", owner)
}

// RemoveOwner is a paid mutator transaction binding the contract method 0x173825d9.
//
// Solidity: function removeOwner(owner address) returns()
func (_MyRRC21 *MyRRC21Session) RemoveOwner(owner common.Address) (*types.Transaction, error) {
	return _MyRRC21.Contract.RemoveOwner(&_MyRRC21.TransactOpts, owner)
}

// RemoveOwner is a paid mutator transaction binding the contract method 0x173825d9.
//
// Solidity: function removeOwner(owner address) returns()
func (_MyRRC21 *MyRRC21TransactorSession) RemoveOwner(owner common.Address) (*types.Transaction, error) {
	return _MyRRC21.Contract.RemoveOwner(&_MyRRC21.TransactOpts, owner)
}

// ReplaceOwner is a paid mutator transaction binding the contract method 0xe20056e6.
//
// Solidity: function replaceOwner(owner address, newOwner address) returns()
func (_MyRRC21 *MyRRC21Transactor) ReplaceOwner(opts *bind.TransactOpts, owner common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _MyRRC21.contract.Transact(opts, "replaceOwner", owner, newOwner)
}

// ReplaceOwner is a paid mutator transaction binding the contract method 0xe20056e6.
//
// Solidity: function replaceOwner(owner address, newOwner address) returns()
func (_MyRRC21 *MyRRC21Session) ReplaceOwner(owner common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _MyRRC21.Contract.ReplaceOwner(&_MyRRC21.TransactOpts, owner, newOwner)
}

// ReplaceOwner is a paid mutator transaction binding the contract method 0xe20056e6.
//
// Solidity: function replaceOwner(owner address, newOwner address) returns()
func (_MyRRC21 *MyRRC21TransactorSession) ReplaceOwner(owner common.Address, newOwner common.Address) (*types.Transaction, error) {
	return _MyRRC21.Contract.ReplaceOwner(&_MyRRC21.TransactOpts, owner, newOwner)
}

// RevokeConfirmation is a paid mutator transaction binding the contract method 0x20ea8d86.
//
// Solidity: function revokeConfirmation(transactionId uint256) returns()
func (_MyRRC21 *MyRRC21Transactor) RevokeConfirmation(opts *bind.TransactOpts, transactionId *big.Int) (*types.Transaction, error) {
	return _MyRRC21.contract.Transact(opts, "revokeConfirmation", transactionId)
}

// RevokeConfirmation is a paid mutator transaction binding the contract method 0x20ea8d86.
//
// Solidity: function revokeConfirmation(transactionId uint256) returns()
func (_MyRRC21 *MyRRC21Session) RevokeConfirmation(transactionId *big.Int) (*types.Transaction, error) {
	return _MyRRC21.Contract.RevokeConfirmation(&_MyRRC21.TransactOpts, transactionId)
}

// RevokeConfirmation is a paid mutator transaction binding the contract method 0x20ea8d86.
//
// Solidity: function revokeConfirmation(transactionId uint256) returns()
func (_MyRRC21 *MyRRC21TransactorSession) RevokeConfirmation(transactionId *big.Int) (*types.Transaction, error) {
	return _MyRRC21.Contract.RevokeConfirmation(&_MyRRC21.TransactOpts, transactionId)
}

// SetDepositFee is a paid mutator transaction binding the contract method 0x490ae210.
//
// Solidity: function setDepositFee(depositFee uint256) returns()
func (_MyRRC21 *MyRRC21Transactor) SetDepositFee(opts *bind.TransactOpts, depositFee *big.Int) (*types.Transaction, error) {
	return _MyRRC21.contract.Transact(opts, "setDepositFee", depositFee)
}

// SetDepositFee is a paid mutator transaction binding the contract method 0x490ae210.
//
// Solidity: function setDepositFee(depositFee uint256) returns()
func (_MyRRC21 *MyRRC21Session) SetDepositFee(depositFee *big.Int) (*types.Transaction, error) {
	return _MyRRC21.Contract.SetDepositFee(&_MyRRC21.TransactOpts, depositFee)
}

// SetDepositFee is a paid mutator transaction binding the contract method 0x490ae210.
//
// Solidity: function setDepositFee(depositFee uint256) returns()
func (_MyRRC21 *MyRRC21TransactorSession) SetDepositFee(depositFee *big.Int) (*types.Transaction, error) {
	return _MyRRC21.Contract.SetDepositFee(&_MyRRC21.TransactOpts, depositFee)
}

// SetMinFee is a paid mutator transaction binding the contract method 0x31ac9920.
//
// Solidity: function setMinFee(value uint256) returns()
func (_MyRRC21 *MyRRC21Transactor) SetMinFee(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _MyRRC21.contract.Transact(opts, "setMinFee", value)
}

// SetMinFee is a paid mutator transaction binding the contract method 0x31ac9920.
//
// Solidity: function setMinFee(value uint256) returns()
func (_MyRRC21 *MyRRC21Session) SetMinFee(value *big.Int) (*types.Transaction, error) {
	return _MyRRC21.Contract.SetMinFee(&_MyRRC21.TransactOpts, value)
}

// SetMinFee is a paid mutator transaction binding the contract method 0x31ac9920.
//
// Solidity: function setMinFee(value uint256) returns()
func (_MyRRC21 *MyRRC21TransactorSession) SetMinFee(value *big.Int) (*types.Transaction, error) {
	return _MyRRC21.Contract.SetMinFee(&_MyRRC21.TransactOpts, value)
}

// SetWithdrawFee is a paid mutator transaction binding the contract method 0xb6ac642a.
//
// Solidity: function setWithdrawFee(withdrawFee uint256) returns()
func (_MyRRC21 *MyRRC21Transactor) SetWithdrawFee(opts *bind.TransactOpts, withdrawFee *big.Int) (*types.Transaction, error) {
	return _MyRRC21.contract.Transact(opts, "setWithdrawFee", withdrawFee)
}

// SetWithdrawFee is a paid mutator transaction binding the contract method 0xb6ac642a.
//
// Solidity: function setWithdrawFee(withdrawFee uint256) returns()
func (_MyRRC21 *MyRRC21Session) SetWithdrawFee(withdrawFee *big.Int) (*types.Transaction, error) {
	return _MyRRC21.Contract.SetWithdrawFee(&_MyRRC21.TransactOpts, withdrawFee)
}

// SetWithdrawFee is a paid mutator transaction binding the contract method 0xb6ac642a.
//
// Solidity: function setWithdrawFee(withdrawFee uint256) returns()
func (_MyRRC21 *MyRRC21TransactorSession) SetWithdrawFee(withdrawFee *big.Int) (*types.Transaction, error) {
	return _MyRRC21.Contract.SetWithdrawFee(&_MyRRC21.TransactOpts, withdrawFee)
}

// SubmitTransaction is a paid mutator transaction binding the contract method 0xc6427474.
//
// Solidity: function submitTransaction(destination address, value uint256, data bytes) returns(transactionId uint256)
func (_MyRRC21 *MyRRC21Transactor) SubmitTransaction(opts *bind.TransactOpts, destination common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _MyRRC21.contract.Transact(opts, "submitTransaction", destination, value, data)
}

// SubmitTransaction is a paid mutator transaction binding the contract method 0xc6427474.
//
// Solidity: function submitTransaction(destination address, value uint256, data bytes) returns(transactionId uint256)
func (_MyRRC21 *MyRRC21Session) SubmitTransaction(destination common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _MyRRC21.Contract.SubmitTransaction(&_MyRRC21.TransactOpts, destination, value, data)
}

// SubmitTransaction is a paid mutator transaction binding the contract method 0xc6427474.
//
// Solidity: function submitTransaction(destination address, value uint256, data bytes) returns(transactionId uint256)
func (_MyRRC21 *MyRRC21TransactorSession) SubmitTransaction(destination common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _MyRRC21.Contract.SubmitTransaction(&_MyRRC21.TransactOpts, destination, value, data)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_MyRRC21 *MyRRC21Transactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MyRRC21.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_MyRRC21 *MyRRC21Session) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MyRRC21.Contract.Transfer(&_MyRRC21.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_MyRRC21 *MyRRC21TransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MyRRC21.Contract.Transfer(&_MyRRC21.TransactOpts, to, value)
}

// TransferContractIssuer is a paid mutator transaction binding the contract method 0xc28ce6ff.
//
// Solidity: function transferContractIssuer(newOwner address) returns()
func (_MyRRC21 *MyRRC21Transactor) TransferContractIssuer(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _MyRRC21.contract.Transact(opts, "transferContractIssuer", newOwner)
}

// TransferContractIssuer is a paid mutator transaction binding the contract method 0xc28ce6ff.
//
// Solidity: function transferContractIssuer(newOwner address) returns()
func (_MyRRC21 *MyRRC21Session) TransferContractIssuer(newOwner common.Address) (*types.Transaction, error) {
	return _MyRRC21.Contract.TransferContractIssuer(&_MyRRC21.TransactOpts, newOwner)
}

// TransferContractIssuer is a paid mutator transaction binding the contract method 0xc28ce6ff.
//
// Solidity: function transferContractIssuer(newOwner address) returns()
func (_MyRRC21 *MyRRC21TransactorSession) TransferContractIssuer(newOwner common.Address) (*types.Transaction, error) {
	return _MyRRC21.Contract.TransferContractIssuer(&_MyRRC21.TransactOpts, newOwner)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(bool)
func (_MyRRC21 *MyRRC21Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MyRRC21.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(bool)
func (_MyRRC21 *MyRRC21Session) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MyRRC21.Contract.TransferFrom(&_MyRRC21.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(bool)
func (_MyRRC21 *MyRRC21TransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _MyRRC21.Contract.TransferFrom(&_MyRRC21.TransactOpts, from, to, value)
}

// MyRRC21ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the MyRRC21 contract.
type MyRRC21ApprovalIterator struct {
	Event *MyRRC21Approval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MyRRC21ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyRRC21Approval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MyRRC21Approval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MyRRC21ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyRRC21ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyRRC21Approval represents a Approval event raised by the MyRRC21 contract.
type MyRRC21Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(owner indexed address, spender indexed address, value uint256)
func (_MyRRC21 *MyRRC21Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*MyRRC21ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _MyRRC21.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &MyRRC21ApprovalIterator{contract: _MyRRC21.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(owner indexed address, spender indexed address, value uint256)
func (_MyRRC21 *MyRRC21Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *MyRRC21Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _MyRRC21.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyRRC21Approval)
				if err := _MyRRC21.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// MyRRC21ConfirmationIterator is returned from FilterConfirmation and is used to iterate over the raw logs and unpacked data for Confirmation events raised by the MyRRC21 contract.
type MyRRC21ConfirmationIterator struct {
	Event *MyRRC21Confirmation // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MyRRC21ConfirmationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyRRC21Confirmation)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MyRRC21Confirmation)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MyRRC21ConfirmationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyRRC21ConfirmationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyRRC21Confirmation represents a Confirmation event raised by the MyRRC21 contract.
type MyRRC21Confirmation struct {
	Sender        common.Address
	TransactionId *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterConfirmation is a free log retrieval operation binding the contract event 0x4a504a94899432a9846e1aa406dceb1bcfd538bb839071d49d1e5e23f5be30ef.
//
// Solidity: event Confirmation(sender indexed address, transactionId indexed uint256)
func (_MyRRC21 *MyRRC21Filterer) FilterConfirmation(opts *bind.FilterOpts, sender []common.Address, transactionId []*big.Int) (*MyRRC21ConfirmationIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _MyRRC21.contract.FilterLogs(opts, "Confirmation", senderRule, transactionIdRule)
	if err != nil {
		return nil, err
	}
	return &MyRRC21ConfirmationIterator{contract: _MyRRC21.contract, event: "Confirmation", logs: logs, sub: sub}, nil
}

// WatchConfirmation is a free log subscription operation binding the contract event 0x4a504a94899432a9846e1aa406dceb1bcfd538bb839071d49d1e5e23f5be30ef.
//
// Solidity: event Confirmation(sender indexed address, transactionId indexed uint256)
func (_MyRRC21 *MyRRC21Filterer) WatchConfirmation(opts *bind.WatchOpts, sink chan<- *MyRRC21Confirmation, sender []common.Address, transactionId []*big.Int) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _MyRRC21.contract.WatchLogs(opts, "Confirmation", senderRule, transactionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyRRC21Confirmation)
				if err := _MyRRC21.contract.UnpackLog(event, "Confirmation", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// MyRRC21ExecutionIterator is returned from FilterExecution and is used to iterate over the raw logs and unpacked data for Execution events raised by the MyRRC21 contract.
type MyRRC21ExecutionIterator struct {
	Event *MyRRC21Execution // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MyRRC21ExecutionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyRRC21Execution)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MyRRC21Execution)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MyRRC21ExecutionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyRRC21ExecutionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyRRC21Execution represents a Execution event raised by the MyRRC21 contract.
type MyRRC21Execution struct {
	TransactionId *big.Int
	Sender        common.Address
	IsMintingTx   bool
	Value         *big.Int
	Data          []byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterExecution is a free log retrieval operation binding the contract event 0x7ae9762dba14e21841a58a1cd988fb97578e27069c9bcc32d65682127141e730.
//
// Solidity: event Execution(transactionId indexed uint256, sender indexed address, isMintingTx bool, value uint256, data bytes)
func (_MyRRC21 *MyRRC21Filterer) FilterExecution(opts *bind.FilterOpts, transactionId []*big.Int, sender []common.Address) (*MyRRC21ExecutionIterator, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _MyRRC21.contract.FilterLogs(opts, "Execution", transactionIdRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &MyRRC21ExecutionIterator{contract: _MyRRC21.contract, event: "Execution", logs: logs, sub: sub}, nil
}

// WatchExecution is a free log subscription operation binding the contract event 0x7ae9762dba14e21841a58a1cd988fb97578e27069c9bcc32d65682127141e730.
//
// Solidity: event Execution(transactionId indexed uint256, sender indexed address, isMintingTx bool, value uint256, data bytes)
func (_MyRRC21 *MyRRC21Filterer) WatchExecution(opts *bind.WatchOpts, sink chan<- *MyRRC21Execution, transactionId []*big.Int, sender []common.Address) (event.Subscription, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _MyRRC21.contract.WatchLogs(opts, "Execution", transactionIdRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyRRC21Execution)
				if err := _MyRRC21.contract.UnpackLog(event, "Execution", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// MyRRC21ExecutionFailureIterator is returned from FilterExecutionFailure and is used to iterate over the raw logs and unpacked data for ExecutionFailure events raised by the MyRRC21 contract.
type MyRRC21ExecutionFailureIterator struct {
	Event *MyRRC21ExecutionFailure // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MyRRC21ExecutionFailureIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyRRC21ExecutionFailure)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MyRRC21ExecutionFailure)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MyRRC21ExecutionFailureIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyRRC21ExecutionFailureIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyRRC21ExecutionFailure represents a ExecutionFailure event raised by the MyRRC21 contract.
type MyRRC21ExecutionFailure struct {
	TransactionId *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterExecutionFailure is a free log retrieval operation binding the contract event 0x526441bb6c1aba3c9a4a6ca1d6545da9c2333c8c48343ef398eb858d72b79236.
//
// Solidity: event ExecutionFailure(transactionId indexed uint256)
func (_MyRRC21 *MyRRC21Filterer) FilterExecutionFailure(opts *bind.FilterOpts, transactionId []*big.Int) (*MyRRC21ExecutionFailureIterator, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _MyRRC21.contract.FilterLogs(opts, "ExecutionFailure", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return &MyRRC21ExecutionFailureIterator{contract: _MyRRC21.contract, event: "ExecutionFailure", logs: logs, sub: sub}, nil
}

// WatchExecutionFailure is a free log subscription operation binding the contract event 0x526441bb6c1aba3c9a4a6ca1d6545da9c2333c8c48343ef398eb858d72b79236.
//
// Solidity: event ExecutionFailure(transactionId indexed uint256)
func (_MyRRC21 *MyRRC21Filterer) WatchExecutionFailure(opts *bind.WatchOpts, sink chan<- *MyRRC21ExecutionFailure, transactionId []*big.Int) (event.Subscription, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _MyRRC21.contract.WatchLogs(opts, "ExecutionFailure", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyRRC21ExecutionFailure)
				if err := _MyRRC21.contract.UnpackLog(event, "ExecutionFailure", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// MyRRC21FeeIterator is returned from FilterFee and is used to iterate over the raw logs and unpacked data for Fee events raised by the MyRRC21 contract.
type MyRRC21FeeIterator struct {
	Event *MyRRC21Fee // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MyRRC21FeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyRRC21Fee)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MyRRC21Fee)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MyRRC21FeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyRRC21FeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyRRC21Fee represents a Fee event raised by the MyRRC21 contract.
type MyRRC21Fee struct {
	From   common.Address
	To     common.Address
	Issuer common.Address
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFee is a free log retrieval operation binding the contract event 0xfcf5b3276434181e3c48bd3fe569b8939808f11e845d4b963693b9d7dbd6dd99.
//
// Solidity: event Fee(from indexed address, to indexed address, issuer indexed address, value uint256)
func (_MyRRC21 *MyRRC21Filterer) FilterFee(opts *bind.FilterOpts, from []common.Address, to []common.Address, issuer []common.Address) (*MyRRC21FeeIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem)
	}

	logs, sub, err := _MyRRC21.contract.FilterLogs(opts, "Fee", fromRule, toRule, issuerRule)
	if err != nil {
		return nil, err
	}
	return &MyRRC21FeeIterator{contract: _MyRRC21.contract, event: "Fee", logs: logs, sub: sub}, nil
}

// WatchFee is a free log subscription operation binding the contract event 0xfcf5b3276434181e3c48bd3fe569b8939808f11e845d4b963693b9d7dbd6dd99.
//
// Solidity: event Fee(from indexed address, to indexed address, issuer indexed address, value uint256)
func (_MyRRC21 *MyRRC21Filterer) WatchFee(opts *bind.WatchOpts, sink chan<- *MyRRC21Fee, from []common.Address, to []common.Address, issuer []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem)
	}

	logs, sub, err := _MyRRC21.contract.WatchLogs(opts, "Fee", fromRule, toRule, issuerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyRRC21Fee)
				if err := _MyRRC21.contract.UnpackLog(event, "Fee", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// MyRRC21OwnerAdditionIterator is returned from FilterOwnerAddition and is used to iterate over the raw logs and unpacked data for OwnerAddition events raised by the MyRRC21 contract.
type MyRRC21OwnerAdditionIterator struct {
	Event *MyRRC21OwnerAddition // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MyRRC21OwnerAdditionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyRRC21OwnerAddition)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MyRRC21OwnerAddition)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MyRRC21OwnerAdditionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyRRC21OwnerAdditionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyRRC21OwnerAddition represents a OwnerAddition event raised by the MyRRC21 contract.
type MyRRC21OwnerAddition struct {
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterOwnerAddition is a free log retrieval operation binding the contract event 0xf39e6e1eb0edcf53c221607b54b00cd28f3196fed0a24994dc308b8f611b682d.
//
// Solidity: event OwnerAddition(owner indexed address)
func (_MyRRC21 *MyRRC21Filterer) FilterOwnerAddition(opts *bind.FilterOpts, owner []common.Address) (*MyRRC21OwnerAdditionIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _MyRRC21.contract.FilterLogs(opts, "OwnerAddition", ownerRule)
	if err != nil {
		return nil, err
	}
	return &MyRRC21OwnerAdditionIterator{contract: _MyRRC21.contract, event: "OwnerAddition", logs: logs, sub: sub}, nil
}

// WatchOwnerAddition is a free log subscription operation binding the contract event 0xf39e6e1eb0edcf53c221607b54b00cd28f3196fed0a24994dc308b8f611b682d.
//
// Solidity: event OwnerAddition(owner indexed address)
func (_MyRRC21 *MyRRC21Filterer) WatchOwnerAddition(opts *bind.WatchOpts, sink chan<- *MyRRC21OwnerAddition, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _MyRRC21.contract.WatchLogs(opts, "OwnerAddition", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyRRC21OwnerAddition)
				if err := _MyRRC21.contract.UnpackLog(event, "OwnerAddition", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// MyRRC21OwnerRemovalIterator is returned from FilterOwnerRemoval and is used to iterate over the raw logs and unpacked data for OwnerRemoval events raised by the MyRRC21 contract.
type MyRRC21OwnerRemovalIterator struct {
	Event *MyRRC21OwnerRemoval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MyRRC21OwnerRemovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyRRC21OwnerRemoval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MyRRC21OwnerRemoval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MyRRC21OwnerRemovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyRRC21OwnerRemovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyRRC21OwnerRemoval represents a OwnerRemoval event raised by the MyRRC21 contract.
type MyRRC21OwnerRemoval struct {
	Owner common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterOwnerRemoval is a free log retrieval operation binding the contract event 0x8001553a916ef2f495d26a907cc54d96ed840d7bda71e73194bf5a9df7a76b90.
//
// Solidity: event OwnerRemoval(owner indexed address)
func (_MyRRC21 *MyRRC21Filterer) FilterOwnerRemoval(opts *bind.FilterOpts, owner []common.Address) (*MyRRC21OwnerRemovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _MyRRC21.contract.FilterLogs(opts, "OwnerRemoval", ownerRule)
	if err != nil {
		return nil, err
	}
	return &MyRRC21OwnerRemovalIterator{contract: _MyRRC21.contract, event: "OwnerRemoval", logs: logs, sub: sub}, nil
}

// WatchOwnerRemoval is a free log subscription operation binding the contract event 0x8001553a916ef2f495d26a907cc54d96ed840d7bda71e73194bf5a9df7a76b90.
//
// Solidity: event OwnerRemoval(owner indexed address)
func (_MyRRC21 *MyRRC21Filterer) WatchOwnerRemoval(opts *bind.WatchOpts, sink chan<- *MyRRC21OwnerRemoval, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _MyRRC21.contract.WatchLogs(opts, "OwnerRemoval", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyRRC21OwnerRemoval)
				if err := _MyRRC21.contract.UnpackLog(event, "OwnerRemoval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// MyRRC21RequirementChangeIterator is returned from FilterRequirementChange and is used to iterate over the raw logs and unpacked data for RequirementChange events raised by the MyRRC21 contract.
type MyRRC21RequirementChangeIterator struct {
	Event *MyRRC21RequirementChange // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MyRRC21RequirementChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyRRC21RequirementChange)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MyRRC21RequirementChange)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MyRRC21RequirementChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyRRC21RequirementChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyRRC21RequirementChange represents a RequirementChange event raised by the MyRRC21 contract.
type MyRRC21RequirementChange struct {
	Required *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRequirementChange is a free log retrieval operation binding the contract event 0xa3f1ee9126a074d9326c682f561767f710e927faa811f7a99829d49dc421797a.
//
// Solidity: event RequirementChange(required uint256)
func (_MyRRC21 *MyRRC21Filterer) FilterRequirementChange(opts *bind.FilterOpts) (*MyRRC21RequirementChangeIterator, error) {

	logs, sub, err := _MyRRC21.contract.FilterLogs(opts, "RequirementChange")
	if err != nil {
		return nil, err
	}
	return &MyRRC21RequirementChangeIterator{contract: _MyRRC21.contract, event: "RequirementChange", logs: logs, sub: sub}, nil
}

// WatchRequirementChange is a free log subscription operation binding the contract event 0xa3f1ee9126a074d9326c682f561767f710e927faa811f7a99829d49dc421797a.
//
// Solidity: event RequirementChange(required uint256)
func (_MyRRC21 *MyRRC21Filterer) WatchRequirementChange(opts *bind.WatchOpts, sink chan<- *MyRRC21RequirementChange) (event.Subscription, error) {

	logs, sub, err := _MyRRC21.contract.WatchLogs(opts, "RequirementChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyRRC21RequirementChange)
				if err := _MyRRC21.contract.UnpackLog(event, "RequirementChange", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// MyRRC21RevocationIterator is returned from FilterRevocation and is used to iterate over the raw logs and unpacked data for Revocation events raised by the MyRRC21 contract.
type MyRRC21RevocationIterator struct {
	Event *MyRRC21Revocation // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MyRRC21RevocationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyRRC21Revocation)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MyRRC21Revocation)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MyRRC21RevocationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyRRC21RevocationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyRRC21Revocation represents a Revocation event raised by the MyRRC21 contract.
type MyRRC21Revocation struct {
	Sender        common.Address
	TransactionId *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRevocation is a free log retrieval operation binding the contract event 0xf6a317157440607f36269043eb55f1287a5a19ba2216afeab88cd46cbcfb88e9.
//
// Solidity: event Revocation(sender indexed address, transactionId indexed uint256)
func (_MyRRC21 *MyRRC21Filterer) FilterRevocation(opts *bind.FilterOpts, sender []common.Address, transactionId []*big.Int) (*MyRRC21RevocationIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _MyRRC21.contract.FilterLogs(opts, "Revocation", senderRule, transactionIdRule)
	if err != nil {
		return nil, err
	}
	return &MyRRC21RevocationIterator{contract: _MyRRC21.contract, event: "Revocation", logs: logs, sub: sub}, nil
}

// WatchRevocation is a free log subscription operation binding the contract event 0xf6a317157440607f36269043eb55f1287a5a19ba2216afeab88cd46cbcfb88e9.
//
// Solidity: event Revocation(sender indexed address, transactionId indexed uint256)
func (_MyRRC21 *MyRRC21Filterer) WatchRevocation(opts *bind.WatchOpts, sink chan<- *MyRRC21Revocation, sender []common.Address, transactionId []*big.Int) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _MyRRC21.contract.WatchLogs(opts, "Revocation", senderRule, transactionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyRRC21Revocation)
				if err := _MyRRC21.contract.UnpackLog(event, "Revocation", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// MyRRC21SubmissionIterator is returned from FilterSubmission and is used to iterate over the raw logs and unpacked data for Submission events raised by the MyRRC21 contract.
type MyRRC21SubmissionIterator struct {
	Event *MyRRC21Submission // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MyRRC21SubmissionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyRRC21Submission)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MyRRC21Submission)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MyRRC21SubmissionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyRRC21SubmissionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyRRC21Submission represents a Submission event raised by the MyRRC21 contract.
type MyRRC21Submission struct {
	TransactionId *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterSubmission is a free log retrieval operation binding the contract event 0xc0ba8fe4b176c1714197d43b9cc6bcf797a4a7461c5fe8d0ef6e184ae7601e51.
//
// Solidity: event Submission(transactionId indexed uint256)
func (_MyRRC21 *MyRRC21Filterer) FilterSubmission(opts *bind.FilterOpts, transactionId []*big.Int) (*MyRRC21SubmissionIterator, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _MyRRC21.contract.FilterLogs(opts, "Submission", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return &MyRRC21SubmissionIterator{contract: _MyRRC21.contract, event: "Submission", logs: logs, sub: sub}, nil
}

// WatchSubmission is a free log subscription operation binding the contract event 0xc0ba8fe4b176c1714197d43b9cc6bcf797a4a7461c5fe8d0ef6e184ae7601e51.
//
// Solidity: event Submission(transactionId indexed uint256)
func (_MyRRC21 *MyRRC21Filterer) WatchSubmission(opts *bind.WatchOpts, sink chan<- *MyRRC21Submission, transactionId []*big.Int) (event.Subscription, error) {

	var transactionIdRule []interface{}
	for _, transactionIdItem := range transactionId {
		transactionIdRule = append(transactionIdRule, transactionIdItem)
	}

	logs, sub, err := _MyRRC21.contract.WatchLogs(opts, "Submission", transactionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyRRC21Submission)
				if err := _MyRRC21.contract.UnpackLog(event, "Submission", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// MyRRC21TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the MyRRC21 contract.
type MyRRC21TransferIterator struct {
	Event *MyRRC21Transfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MyRRC21TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MyRRC21Transfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MyRRC21Transfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MyRRC21TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MyRRC21TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MyRRC21Transfer represents a Transfer event raised by the MyRRC21 contract.
type MyRRC21Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(from indexed address, to indexed address, value uint256)
func (_MyRRC21 *MyRRC21Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*MyRRC21TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MyRRC21.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MyRRC21TransferIterator{contract: _MyRRC21.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(from indexed address, to indexed address, value uint256)
func (_MyRRC21 *MyRRC21Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *MyRRC21Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _MyRRC21.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MyRRC21Transfer)
				if err := _MyRRC21.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// RRC21ABI is the input ABI used to generate the binding from.
const RRC21ABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"estimateFee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"issuer\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"minFee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"setMinFee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"symbol\",\"type\":\"string\"},{\"name\":\"decimals\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"issuer\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Fee\",\"type\":\"event\"}]"

// RRC21Bin is the compiled bytecode used for deploying new contracts.
const RRC21Bin = `0x608060405234801561001057600080fd5b50604051610a19380380610a1983398101604090815281516020808401519284015191840180519094939093019261004e916005919086019061007f565b50815161006290600690602085019061007f565b506007805460ff191660ff929092169190911790555061011a9050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106100c057805160ff19168380011785556100ed565b828001600101855582156100ed579182015b828111156100ed5782518255916020019190600101906100d2565b506100f99291506100fd565b5090565b61011791905b808211156100f95760008155600101610103565b90565b6108f0806101296000396000f3006080604052600436106100c45763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166306fdde0381146100c9578063095ea7b314610153578063127e8e4d1461018b57806318160ddd146101b55780631d143848146101ca57806323b872dd146101fb57806324ec759014610225578063313ce5671461023a57806331ac99201461026557806370a082311461027f57806395d89b41146102a0578063a9059cbb146102b5578063dd62ed3e146102d9575b600080fd5b3480156100d557600080fd5b506100de610300565b6040805160208082528351818301528351919283929083019185019080838360005b83811015610118578181015183820152602001610100565b50505050905090810190601f1680156101455780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561015f57600080fd5b50610177600160a060020a0360043516602435610396565b604080519115158252519081900360200190f35b34801561019757600080fd5b506101a3600435610450565b60408051918252519081900360200190f35b3480156101c157600080fd5b506101a361047c565b3480156101d657600080fd5b506101df610482565b60408051600160a060020a039092168252519081900360200190f35b34801561020757600080fd5b50610177600160a060020a0360043581169060243516604435610491565b34801561023157600080fd5b506101a36105d1565b34801561024657600080fd5b5061024f6105d7565b6040805160ff9092168252519081900360200190f35b34801561027157600080fd5b5061027d6004356105e0565b005b34801561028b57600080fd5b506101a3600160a060020a0360043516610603565b3480156102ac57600080fd5b506100de61061e565b3480156102c157600080fd5b50610177600160a060020a036004351660243561067f565b3480156102e557600080fd5b506101a3600160a060020a0360043581169060243516610744565b60058054604080516020601f600260001961010060018816150201909516949094049384018190048102820181019092528281526060939092909183018282801561038c5780601f106103615761010080835404028352916020019161038c565b820191906000526020600020905b81548152906001019060200180831161036f57829003601f168201915b5050505050905090565b6000600160a060020a03831615156103ad57600080fd5b6001543360009081526020819052604090205410156103cb57600080fd5b336000818152600360209081526040808320600160a060020a03888116855292529091208490556002546001546104079392919091169061076f565b604080518381529051600160a060020a0385169133917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259181900360200190a350600192915050565b6001546000906104769061046a848463ffffffff61086116565b9063ffffffff61089616565b92915050565b60045490565b600254600160a060020a031690565b6000806104a96001548461089690919063ffffffff16565b9050600160a060020a03841615156104c057600080fd5b808311156104cd57600080fd5b600160a060020a03851660009081526003602090815260408083203384529091529020548111156104fd57600080fd5b600160a060020a0385166000908152600360209081526040808320338452909152902054610531908263ffffffff6108a816565b600160a060020a038616600090815260036020908152604080832033845290915290205561056085858561076f565b60025460015461057d918791600160a060020a039091169061076f565b6002546001546040805191825251600160a060020a039283169287169133917ffcf5b3276434181e3c48bd3fe569b8939808f11e845d4b963693b9d7dbd6dd999181900360200190a4506001949350505050565b60015490565b60075460ff1690565b600254600160a060020a031633146105f757600080fd5b610600816108bf565b50565b600160a060020a031660009081526020819052604090205490565b60068054604080516020601f600260001961010060018816150201909516949094049384018190048102820181019092528281526060939092909183018282801561038c5780601f106103615761010080835404028352916020019161038c565b6000806106976001548461089690919063ffffffff16565b9050600160a060020a03841615156106ae57600080fd5b808311156106bb57600080fd5b6106c633858561076f565b60006001541115610738576002546001546106ee913391600160a060020a039091169061076f565b6002546001546040805191825251600160a060020a039283169287169133917ffcf5b3276434181e3c48bd3fe569b8939808f11e845d4b963693b9d7dbd6dd999181900360200190a45b600191505b5092915050565b600160a060020a03918216600090815260036020908152604080832093909416825291909152205490565b600160a060020a03831660009081526020819052604090205481111561079457600080fd5b600160a060020a03821615156107a957600080fd5b600160a060020a0383166000908152602081905260409020546107d2908263ffffffff6108a816565b600160a060020a038085166000908152602081905260408082209390935590841681522054610807908263ffffffff61089616565b600160a060020a038084166000818152602081815260409182902094909455805185815290519193928716927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef92918290030190a3505050565b600080831515610874576000915061073d565b5082820282848281151561088457fe5b041461088f57600080fd5b9392505050565b60008282018381101561088f57600080fd5b600080838311156108b857600080fd5b5050900390565b6001555600a165627a7a723058202a95582cfc0b66603c8ebbf18b84c34c432c995caa84d558310570fde20cac1b0029`

// DeployRRC21 deploys a new Ethereum contract, binding an instance of RRC21 to it.
func DeployRRC21(auth *bind.TransactOpts, backend bind.ContractBackend, name string, symbol string, decimals uint8) (common.Address, *types.Transaction, *RRC21, error) {
	parsed, err := abi.JSON(strings.NewReader(RRC21ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RRC21Bin), backend, name, symbol, decimals)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RRC21{RRC21Caller: RRC21Caller{contract: contract}, RRC21Transactor: RRC21Transactor{contract: contract}, RRC21Filterer: RRC21Filterer{contract: contract}}, nil
}

// RRC21 is an auto generated Go binding around an Ethereum contract.
type RRC21 struct {
	RRC21Caller     // Read-only binding to the contract
	RRC21Transactor // Write-only binding to the contract
	RRC21Filterer   // Log filterer for contract events
}

// RRC21Caller is an auto generated read-only Go binding around an Ethereum contract.
type RRC21Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RRC21Transactor is an auto generated write-only Go binding around an Ethereum contract.
type RRC21Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RRC21Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RRC21Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RRC21Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RRC21Session struct {
	Contract     *RRC21            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RRC21CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RRC21CallerSession struct {
	Contract *RRC21Caller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// RRC21TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RRC21TransactorSession struct {
	Contract     *RRC21Transactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RRC21Raw is an auto generated low-level Go binding around an Ethereum contract.
type RRC21Raw struct {
	Contract *RRC21 // Generic contract binding to access the raw methods on
}

// RRC21CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RRC21CallerRaw struct {
	Contract *RRC21Caller // Generic read-only contract binding to access the raw methods on
}

// RRC21TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RRC21TransactorRaw struct {
	Contract *RRC21Transactor // Generic write-only contract binding to access the raw methods on
}

// NewRRC21 creates a new instance of RRC21, bound to a specific deployed contract.
func NewRRC21(address common.Address, backend bind.ContractBackend) (*RRC21, error) {
	contract, err := bindRRC21(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RRC21{RRC21Caller: RRC21Caller{contract: contract}, RRC21Transactor: RRC21Transactor{contract: contract}, RRC21Filterer: RRC21Filterer{contract: contract}}, nil
}

// NewRRC21Caller creates a new read-only instance of RRC21, bound to a specific deployed contract.
func NewRRC21Caller(address common.Address, caller bind.ContractCaller) (*RRC21Caller, error) {
	contract, err := bindRRC21(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RRC21Caller{contract: contract}, nil
}

// NewRRC21Transactor creates a new write-only instance of RRC21, bound to a specific deployed contract.
func NewRRC21Transactor(address common.Address, transactor bind.ContractTransactor) (*RRC21Transactor, error) {
	contract, err := bindRRC21(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RRC21Transactor{contract: contract}, nil
}

// NewRRC21Filterer creates a new log filterer instance of RRC21, bound to a specific deployed contract.
func NewRRC21Filterer(address common.Address, filterer bind.ContractFilterer) (*RRC21Filterer, error) {
	contract, err := bindRRC21(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RRC21Filterer{contract: contract}, nil
}

// bindRRC21 binds a generic wrapper to an already deployed contract.
func bindRRC21(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RRC21ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RRC21 *RRC21Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RRC21.Contract.RRC21Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RRC21 *RRC21Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RRC21.Contract.RRC21Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RRC21 *RRC21Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RRC21.Contract.RRC21Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RRC21 *RRC21CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _RRC21.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RRC21 *RRC21TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RRC21.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RRC21 *RRC21TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RRC21.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(uint256)
func (_RRC21 *RRC21Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RRC21.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(uint256)
func (_RRC21 *RRC21Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _RRC21.Contract.Allowance(&_RRC21.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(uint256)
func (_RRC21 *RRC21CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _RRC21.Contract.Allowance(&_RRC21.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(owner address) constant returns(uint256)
func (_RRC21 *RRC21Caller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RRC21.contract.Call(opts, out, "balanceOf", owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(owner address) constant returns(uint256)
func (_RRC21 *RRC21Session) BalanceOf(owner common.Address) (*big.Int, error) {
	return _RRC21.Contract.BalanceOf(&_RRC21.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(owner address) constant returns(uint256)
func (_RRC21 *RRC21CallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _RRC21.Contract.BalanceOf(&_RRC21.CallOpts, owner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_RRC21 *RRC21Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _RRC21.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_RRC21 *RRC21Session) Decimals() (uint8, error) {
	return _RRC21.Contract.Decimals(&_RRC21.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_RRC21 *RRC21CallerSession) Decimals() (uint8, error) {
	return _RRC21.Contract.Decimals(&_RRC21.CallOpts)
}

// EstimateFee is a free data retrieval call binding the contract method 0x127e8e4d.
//
// Solidity: function estimateFee(value uint256) constant returns(uint256)
func (_RRC21 *RRC21Caller) EstimateFee(opts *bind.CallOpts, value *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RRC21.contract.Call(opts, out, "estimateFee", value)
	return *ret0, err
}

// EstimateFee is a free data retrieval call binding the contract method 0x127e8e4d.
//
// Solidity: function estimateFee(value uint256) constant returns(uint256)
func (_RRC21 *RRC21Session) EstimateFee(value *big.Int) (*big.Int, error) {
	return _RRC21.Contract.EstimateFee(&_RRC21.CallOpts, value)
}

// EstimateFee is a free data retrieval call binding the contract method 0x127e8e4d.
//
// Solidity: function estimateFee(value uint256) constant returns(uint256)
func (_RRC21 *RRC21CallerSession) EstimateFee(value *big.Int) (*big.Int, error) {
	return _RRC21.Contract.EstimateFee(&_RRC21.CallOpts, value)
}

// Issuer is a free data retrieval call binding the contract method 0x1d143848.
//
// Solidity: function issuer() constant returns(address)
func (_RRC21 *RRC21Caller) Issuer(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _RRC21.contract.Call(opts, out, "issuer")
	return *ret0, err
}

// Issuer is a free data retrieval call binding the contract method 0x1d143848.
//
// Solidity: function issuer() constant returns(address)
func (_RRC21 *RRC21Session) Issuer() (common.Address, error) {
	return _RRC21.Contract.Issuer(&_RRC21.CallOpts)
}

// Issuer is a free data retrieval call binding the contract method 0x1d143848.
//
// Solidity: function issuer() constant returns(address)
func (_RRC21 *RRC21CallerSession) Issuer() (common.Address, error) {
	return _RRC21.Contract.Issuer(&_RRC21.CallOpts)
}

// MinFee is a free data retrieval call binding the contract method 0x24ec7590.
//
// Solidity: function minFee() constant returns(uint256)
func (_RRC21 *RRC21Caller) MinFee(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RRC21.contract.Call(opts, out, "minFee")
	return *ret0, err
}

// MinFee is a free data retrieval call binding the contract method 0x24ec7590.
//
// Solidity: function minFee() constant returns(uint256)
func (_RRC21 *RRC21Session) MinFee() (*big.Int, error) {
	return _RRC21.Contract.MinFee(&_RRC21.CallOpts)
}

// MinFee is a free data retrieval call binding the contract method 0x24ec7590.
//
// Solidity: function minFee() constant returns(uint256)
func (_RRC21 *RRC21CallerSession) MinFee() (*big.Int, error) {
	return _RRC21.Contract.MinFee(&_RRC21.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_RRC21 *RRC21Caller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _RRC21.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_RRC21 *RRC21Session) Name() (string, error) {
	return _RRC21.Contract.Name(&_RRC21.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_RRC21 *RRC21CallerSession) Name() (string, error) {
	return _RRC21.Contract.Name(&_RRC21.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_RRC21 *RRC21Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _RRC21.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_RRC21 *RRC21Session) Symbol() (string, error) {
	return _RRC21.Contract.Symbol(&_RRC21.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_RRC21 *RRC21CallerSession) Symbol() (string, error) {
	return _RRC21.Contract.Symbol(&_RRC21.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_RRC21 *RRC21Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _RRC21.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_RRC21 *RRC21Session) TotalSupply() (*big.Int, error) {
	return _RRC21.Contract.TotalSupply(&_RRC21.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_RRC21 *RRC21CallerSession) TotalSupply() (*big.Int, error) {
	return _RRC21.Contract.TotalSupply(&_RRC21.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(bool)
func (_RRC21 *RRC21Transactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _RRC21.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(bool)
func (_RRC21 *RRC21Session) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _RRC21.Contract.Approve(&_RRC21.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(bool)
func (_RRC21 *RRC21TransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _RRC21.Contract.Approve(&_RRC21.TransactOpts, spender, value)
}

// SetMinFee is a paid mutator transaction binding the contract method 0x31ac9920.
//
// Solidity: function setMinFee(value uint256) returns()
func (_RRC21 *RRC21Transactor) SetMinFee(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _RRC21.contract.Transact(opts, "setMinFee", value)
}

// SetMinFee is a paid mutator transaction binding the contract method 0x31ac9920.
//
// Solidity: function setMinFee(value uint256) returns()
func (_RRC21 *RRC21Session) SetMinFee(value *big.Int) (*types.Transaction, error) {
	return _RRC21.Contract.SetMinFee(&_RRC21.TransactOpts, value)
}

// SetMinFee is a paid mutator transaction binding the contract method 0x31ac9920.
//
// Solidity: function setMinFee(value uint256) returns()
func (_RRC21 *RRC21TransactorSession) SetMinFee(value *big.Int) (*types.Transaction, error) {
	return _RRC21.Contract.SetMinFee(&_RRC21.TransactOpts, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_RRC21 *RRC21Transactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _RRC21.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_RRC21 *RRC21Session) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _RRC21.Contract.Transfer(&_RRC21.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_RRC21 *RRC21TransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _RRC21.Contract.Transfer(&_RRC21.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(bool)
func (_RRC21 *RRC21Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _RRC21.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(bool)
func (_RRC21 *RRC21Session) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _RRC21.Contract.TransferFrom(&_RRC21.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(bool)
func (_RRC21 *RRC21TransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _RRC21.Contract.TransferFrom(&_RRC21.TransactOpts, from, to, value)
}

// RRC21ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the RRC21 contract.
type RRC21ApprovalIterator struct {
	Event *RRC21Approval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RRC21ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RRC21Approval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RRC21Approval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RRC21ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RRC21ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RRC21Approval represents a Approval event raised by the RRC21 contract.
type RRC21Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(owner indexed address, spender indexed address, value uint256)
func (_RRC21 *RRC21Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*RRC21ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _RRC21.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &RRC21ApprovalIterator{contract: _RRC21.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(owner indexed address, spender indexed address, value uint256)
func (_RRC21 *RRC21Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *RRC21Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _RRC21.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RRC21Approval)
				if err := _RRC21.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// RRC21FeeIterator is returned from FilterFee and is used to iterate over the raw logs and unpacked data for Fee events raised by the RRC21 contract.
type RRC21FeeIterator struct {
	Event *RRC21Fee // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RRC21FeeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RRC21Fee)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RRC21Fee)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RRC21FeeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RRC21FeeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RRC21Fee represents a Fee event raised by the RRC21 contract.
type RRC21Fee struct {
	From   common.Address
	To     common.Address
	Issuer common.Address
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFee is a free log retrieval operation binding the contract event 0xfcf5b3276434181e3c48bd3fe569b8939808f11e845d4b963693b9d7dbd6dd99.
//
// Solidity: event Fee(from indexed address, to indexed address, issuer indexed address, value uint256)
func (_RRC21 *RRC21Filterer) FilterFee(opts *bind.FilterOpts, from []common.Address, to []common.Address, issuer []common.Address) (*RRC21FeeIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem)
	}

	logs, sub, err := _RRC21.contract.FilterLogs(opts, "Fee", fromRule, toRule, issuerRule)
	if err != nil {
		return nil, err
	}
	return &RRC21FeeIterator{contract: _RRC21.contract, event: "Fee", logs: logs, sub: sub}, nil
}

// WatchFee is a free log subscription operation binding the contract event 0xfcf5b3276434181e3c48bd3fe569b8939808f11e845d4b963693b9d7dbd6dd99.
//
// Solidity: event Fee(from indexed address, to indexed address, issuer indexed address, value uint256)
func (_RRC21 *RRC21Filterer) WatchFee(opts *bind.WatchOpts, sink chan<- *RRC21Fee, from []common.Address, to []common.Address, issuer []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var issuerRule []interface{}
	for _, issuerItem := range issuer {
		issuerRule = append(issuerRule, issuerItem)
	}

	logs, sub, err := _RRC21.contract.WatchLogs(opts, "Fee", fromRule, toRule, issuerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RRC21Fee)
				if err := _RRC21.contract.UnpackLog(event, "Fee", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// RRC21TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the RRC21 contract.
type RRC21TransferIterator struct {
	Event *RRC21Transfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *RRC21TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RRC21Transfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(RRC21Transfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *RRC21TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RRC21TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RRC21Transfer represents a Transfer event raised by the RRC21 contract.
type RRC21Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(from indexed address, to indexed address, value uint256)
func (_RRC21 *RRC21Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*RRC21TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _RRC21.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &RRC21TransferIterator{contract: _RRC21.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(from indexed address, to indexed address, value uint256)
func (_RRC21 *RRC21Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *RRC21Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _RRC21.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RRC21Transfer)
				if err := _RRC21.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}
