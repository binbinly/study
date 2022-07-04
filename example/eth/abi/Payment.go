// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package eth

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

// PaymentABI is the input ABI used to generate the binding from.
const PaymentABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"order_no\",\"type\":\"string\"}],\"name\":\"OnFinish\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"minPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"receiveAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"setMinPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setReceiveAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"setToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"trades\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"date\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"order_no\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"pay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"order_no\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"query\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Payment is an auto generated Go binding around an Ethereum contract.
type Payment struct {
	PaymentCaller     // Read-only binding to the contract
	PaymentTransactor // Write-only binding to the contract
	PaymentFilterer   // Log filterer for contract events
}

// PaymentCaller is an auto generated read-only Go binding around an Ethereum contract.
type PaymentCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PaymentTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PaymentTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PaymentFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PaymentFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PaymentSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PaymentSession struct {
	Contract     *Payment          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PaymentCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PaymentCallerSession struct {
	Contract *PaymentCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// PaymentTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PaymentTransactorSession struct {
	Contract     *PaymentTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// PaymentRaw is an auto generated low-level Go binding around an Ethereum contract.
type PaymentRaw struct {
	Contract *Payment // Generic contract binding to access the raw methods on
}

// PaymentCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PaymentCallerRaw struct {
	Contract *PaymentCaller // Generic read-only contract binding to access the raw methods on
}

// PaymentTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PaymentTransactorRaw struct {
	Contract *PaymentTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPayment creates a new instance of Payment, bound to a specific deployed contract.
func NewPayment(address common.Address, backend bind.ContractBackend) (*Payment, error) {
	contract, err := bindPayment(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Payment{PaymentCaller: PaymentCaller{contract: contract}, PaymentTransactor: PaymentTransactor{contract: contract}, PaymentFilterer: PaymentFilterer{contract: contract}}, nil
}

// NewPaymentCaller creates a new read-only instance of Payment, bound to a specific deployed contract.
func NewPaymentCaller(address common.Address, caller bind.ContractCaller) (*PaymentCaller, error) {
	contract, err := bindPayment(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PaymentCaller{contract: contract}, nil
}

// NewPaymentTransactor creates a new write-only instance of Payment, bound to a specific deployed contract.
func NewPaymentTransactor(address common.Address, transactor bind.ContractTransactor) (*PaymentTransactor, error) {
	contract, err := bindPayment(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PaymentTransactor{contract: contract}, nil
}

// NewPaymentFilterer creates a new log filterer instance of Payment, bound to a specific deployed contract.
func NewPaymentFilterer(address common.Address, filterer bind.ContractFilterer) (*PaymentFilterer, error) {
	contract, err := bindPayment(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PaymentFilterer{contract: contract}, nil
}

// bindPayment binds a generic wrapper to an already deployed contract.
func bindPayment(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PaymentABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Payment *PaymentRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Payment.Contract.PaymentCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Payment *PaymentRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Payment.Contract.PaymentTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Payment *PaymentRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Payment.Contract.PaymentTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Payment *PaymentCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Payment.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Payment *PaymentTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Payment.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Payment *PaymentTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Payment.Contract.contract.Transact(opts, method, params...)
}

// MinPrice is a free data retrieval call binding the contract method 0xe45be8eb.
//
// Solidity: function minPrice() view returns(uint256)
func (_Payment *PaymentCaller) MinPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Payment.contract.Call(opts, &out, "minPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinPrice is a free data retrieval call binding the contract method 0xe45be8eb.
//
// Solidity: function minPrice() view returns(uint256)
func (_Payment *PaymentSession) MinPrice() (*big.Int, error) {
	return _Payment.Contract.MinPrice(&_Payment.CallOpts)
}

// MinPrice is a free data retrieval call binding the contract method 0xe45be8eb.
//
// Solidity: function minPrice() view returns(uint256)
func (_Payment *PaymentCallerSession) MinPrice() (*big.Int, error) {
	return _Payment.Contract.MinPrice(&_Payment.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Payment *PaymentCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Payment.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Payment *PaymentSession) Owner() (common.Address, error) {
	return _Payment.Contract.Owner(&_Payment.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Payment *PaymentCallerSession) Owner() (common.Address, error) {
	return _Payment.Contract.Owner(&_Payment.CallOpts)
}

// Query is a free data retrieval call binding the contract method 0x7521357d.
//
// Solidity: function query(string order_no, uint256 value) view returns(bool)
func (_Payment *PaymentCaller) Query(opts *bind.CallOpts, order_no string, value *big.Int) (bool, error) {
	var out []interface{}
	err := _Payment.contract.Call(opts, &out, "query", order_no, value)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Query is a free data retrieval call binding the contract method 0x7521357d.
//
// Solidity: function query(string order_no, uint256 value) view returns(bool)
func (_Payment *PaymentSession) Query(order_no string, value *big.Int) (bool, error) {
	return _Payment.Contract.Query(&_Payment.CallOpts, order_no, value)
}

// Query is a free data retrieval call binding the contract method 0x7521357d.
//
// Solidity: function query(string order_no, uint256 value) view returns(bool)
func (_Payment *PaymentCallerSession) Query(order_no string, value *big.Int) (bool, error) {
	return _Payment.Contract.Query(&_Payment.CallOpts, order_no, value)
}

// ReceiveAddress is a free data retrieval call binding the contract method 0xfffe42e9.
//
// Solidity: function receiveAddress() view returns(address)
func (_Payment *PaymentCaller) ReceiveAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Payment.contract.Call(opts, &out, "receiveAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ReceiveAddress is a free data retrieval call binding the contract method 0xfffe42e9.
//
// Solidity: function receiveAddress() view returns(address)
func (_Payment *PaymentSession) ReceiveAddress() (common.Address, error) {
	return _Payment.Contract.ReceiveAddress(&_Payment.CallOpts)
}

// ReceiveAddress is a free data retrieval call binding the contract method 0xfffe42e9.
//
// Solidity: function receiveAddress() view returns(address)
func (_Payment *PaymentCallerSession) ReceiveAddress() (common.Address, error) {
	return _Payment.Contract.ReceiveAddress(&_Payment.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Payment *PaymentCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Payment.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Payment *PaymentSession) Token() (common.Address, error) {
	return _Payment.Contract.Token(&_Payment.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Payment *PaymentCallerSession) Token() (common.Address, error) {
	return _Payment.Contract.Token(&_Payment.CallOpts)
}

// Trades is a free data retrieval call binding the contract method 0x47a04100.
//
// Solidity: function trades(string ) view returns(address addr, uint256 amount, uint256 date)
func (_Payment *PaymentCaller) Trades(opts *bind.CallOpts, arg0 string) (struct {
	Addr   common.Address
	Amount *big.Int
	Date   *big.Int
}, error) {
	var out []interface{}
	err := _Payment.contract.Call(opts, &out, "trades", arg0)

	outstruct := new(struct {
		Addr   common.Address
		Amount *big.Int
		Date   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Addr = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Amount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Date = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Trades is a free data retrieval call binding the contract method 0x47a04100.
//
// Solidity: function trades(string ) view returns(address addr, uint256 amount, uint256 date)
func (_Payment *PaymentSession) Trades(arg0 string) (struct {
	Addr   common.Address
	Amount *big.Int
	Date   *big.Int
}, error) {
	return _Payment.Contract.Trades(&_Payment.CallOpts, arg0)
}

// Trades is a free data retrieval call binding the contract method 0x47a04100.
//
// Solidity: function trades(string ) view returns(address addr, uint256 amount, uint256 date)
func (_Payment *PaymentCallerSession) Trades(arg0 string) (struct {
	Addr   common.Address
	Amount *big.Int
	Date   *big.Int
}, error) {
	return _Payment.Contract.Trades(&_Payment.CallOpts, arg0)
}

// Pay is a paid mutator transaction binding the contract method 0x901c5953.
//
// Solidity: function pay(string order_no, uint256 value) returns()
func (_Payment *PaymentTransactor) Pay(opts *bind.TransactOpts, order_no string, value *big.Int) (*types.Transaction, error) {
	return _Payment.contract.Transact(opts, "pay", order_no, value)
}

// Pay is a paid mutator transaction binding the contract method 0x901c5953.
//
// Solidity: function pay(string order_no, uint256 value) returns()
func (_Payment *PaymentSession) Pay(order_no string, value *big.Int) (*types.Transaction, error) {
	return _Payment.Contract.Pay(&_Payment.TransactOpts, order_no, value)
}

// Pay is a paid mutator transaction binding the contract method 0x901c5953.
//
// Solidity: function pay(string order_no, uint256 value) returns()
func (_Payment *PaymentTransactorSession) Pay(order_no string, value *big.Int) (*types.Transaction, error) {
	return _Payment.Contract.Pay(&_Payment.TransactOpts, order_no, value)
}

// SetMinPrice is a paid mutator transaction binding the contract method 0x5ea8cd12.
//
// Solidity: function setMinPrice(uint256 _price) returns()
func (_Payment *PaymentTransactor) SetMinPrice(opts *bind.TransactOpts, _price *big.Int) (*types.Transaction, error) {
	return _Payment.contract.Transact(opts, "setMinPrice", _price)
}

// SetMinPrice is a paid mutator transaction binding the contract method 0x5ea8cd12.
//
// Solidity: function setMinPrice(uint256 _price) returns()
func (_Payment *PaymentSession) SetMinPrice(_price *big.Int) (*types.Transaction, error) {
	return _Payment.Contract.SetMinPrice(&_Payment.TransactOpts, _price)
}

// SetMinPrice is a paid mutator transaction binding the contract method 0x5ea8cd12.
//
// Solidity: function setMinPrice(uint256 _price) returns()
func (_Payment *PaymentTransactorSession) SetMinPrice(_price *big.Int) (*types.Transaction, error) {
	return _Payment.Contract.SetMinPrice(&_Payment.TransactOpts, _price)
}

// SetReceiveAddress is a paid mutator transaction binding the contract method 0x5ec4b7a8.
//
// Solidity: function setReceiveAddress(address addr) returns()
func (_Payment *PaymentTransactor) SetReceiveAddress(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Payment.contract.Transact(opts, "setReceiveAddress", addr)
}

// SetReceiveAddress is a paid mutator transaction binding the contract method 0x5ec4b7a8.
//
// Solidity: function setReceiveAddress(address addr) returns()
func (_Payment *PaymentSession) SetReceiveAddress(addr common.Address) (*types.Transaction, error) {
	return _Payment.Contract.SetReceiveAddress(&_Payment.TransactOpts, addr)
}

// SetReceiveAddress is a paid mutator transaction binding the contract method 0x5ec4b7a8.
//
// Solidity: function setReceiveAddress(address addr) returns()
func (_Payment *PaymentTransactorSession) SetReceiveAddress(addr common.Address) (*types.Transaction, error) {
	return _Payment.Contract.SetReceiveAddress(&_Payment.TransactOpts, addr)
}

// SetToken is a paid mutator transaction binding the contract method 0x144fa6d7.
//
// Solidity: function setToken(address _token) returns()
func (_Payment *PaymentTransactor) SetToken(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _Payment.contract.Transact(opts, "setToken", _token)
}

// SetToken is a paid mutator transaction binding the contract method 0x144fa6d7.
//
// Solidity: function setToken(address _token) returns()
func (_Payment *PaymentSession) SetToken(_token common.Address) (*types.Transaction, error) {
	return _Payment.Contract.SetToken(&_Payment.TransactOpts, _token)
}

// SetToken is a paid mutator transaction binding the contract method 0x144fa6d7.
//
// Solidity: function setToken(address _token) returns()
func (_Payment *PaymentTransactorSession) SetToken(_token common.Address) (*types.Transaction, error) {
	return _Payment.Contract.SetToken(&_Payment.TransactOpts, _token)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Payment *PaymentTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Payment.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Payment *PaymentSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Payment.Contract.TransferOwnership(&_Payment.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Payment *PaymentTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Payment.Contract.TransferOwnership(&_Payment.TransactOpts, newOwner)
}

// PaymentOnFinishIterator is returned from FilterOnFinish and is used to iterate over the raw logs and unpacked data for OnFinish events raised by the Payment contract.
type PaymentOnFinishIterator struct {
	Event *PaymentOnFinish // Event containing the contract specifics and raw log

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
func (it *PaymentOnFinishIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PaymentOnFinish)
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
		it.Event = new(PaymentOnFinish)
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
func (it *PaymentOnFinishIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PaymentOnFinishIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PaymentOnFinish represents a OnFinish event raised by the Payment contract.
type PaymentOnFinish struct {
	Addr    common.Address
	OrderNo string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterOnFinish is a free log retrieval operation binding the contract event 0x135b7b32f6d6347a78767b3c41000f81ce759b6d67750431d9df82eb9d80593a.
//
// Solidity: event OnFinish(address addr, string order_no)
func (_Payment *PaymentFilterer) FilterOnFinish(opts *bind.FilterOpts) (*PaymentOnFinishIterator, error) {

	logs, sub, err := _Payment.contract.FilterLogs(opts, "OnFinish")
	if err != nil {
		return nil, err
	}
	return &PaymentOnFinishIterator{contract: _Payment.contract, event: "OnFinish", logs: logs, sub: sub}, nil
}

// WatchOnFinish is a free log subscription operation binding the contract event 0x135b7b32f6d6347a78767b3c41000f81ce759b6d67750431d9df82eb9d80593a.
//
// Solidity: event OnFinish(address addr, string order_no)
func (_Payment *PaymentFilterer) WatchOnFinish(opts *bind.WatchOpts, sink chan<- *PaymentOnFinish) (event.Subscription, error) {

	logs, sub, err := _Payment.contract.WatchLogs(opts, "OnFinish")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PaymentOnFinish)
				if err := _Payment.contract.UnpackLog(event, "OnFinish", log); err != nil {
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

// ParseOnFinish is a log parse operation binding the contract event 0x135b7b32f6d6347a78767b3c41000f81ce759b6d67750431d9df82eb9d80593a.
//
// Solidity: event OnFinish(address addr, string order_no)
func (_Payment *PaymentFilterer) ParseOnFinish(log types.Log) (*PaymentOnFinish, error) {
	event := new(PaymentOnFinish)
	if err := _Payment.contract.UnpackLog(event, "OnFinish", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
