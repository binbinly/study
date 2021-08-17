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

// LotteryABI is the input ABI used to generate the binding from.
const LotteryABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"nums\",\"type\":\"string\"}],\"name\":\"BetFinish\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"nums\",\"type\":\"uint256[]\"}],\"name\":\"DrawFinish\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"accountList\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"nums\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"accounts\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"nums\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_nums\",\"type\":\"string\"}],\"name\":\"bet\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"draw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"period\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"retNums\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"retNumsList\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Lottery is an auto generated Go binding around an Ethereum contract.
type Lottery struct {
	LotteryCaller     // Read-only binding to the contract
	LotteryTransactor // Write-only binding to the contract
	LotteryFilterer   // Log filterer for contract events
}

// LotteryCaller is an auto generated read-only Go binding around an Ethereum contract.
type LotteryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LotteryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LotteryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LotteryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LotteryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LotterySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LotterySession struct {
	Contract     *Lottery          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LotteryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LotteryCallerSession struct {
	Contract *LotteryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// LotteryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LotteryTransactorSession struct {
	Contract     *LotteryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// LotteryRaw is an auto generated low-level Go binding around an Ethereum contract.
type LotteryRaw struct {
	Contract *Lottery // Generic contract binding to access the raw methods on
}

// LotteryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LotteryCallerRaw struct {
	Contract *LotteryCaller // Generic read-only contract binding to access the raw methods on
}

// LotteryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LotteryTransactorRaw struct {
	Contract *LotteryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLottery creates a new instance of Lottery, bound to a specific deployed contract.
func NewLottery(address common.Address, backend bind.ContractBackend) (*Lottery, error) {
	contract, err := bindLottery(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Lottery{LotteryCaller: LotteryCaller{contract: contract}, LotteryTransactor: LotteryTransactor{contract: contract}, LotteryFilterer: LotteryFilterer{contract: contract}}, nil
}

// NewLotteryCaller creates a new read-only instance of Lottery, bound to a specific deployed contract.
func NewLotteryCaller(address common.Address, caller bind.ContractCaller) (*LotteryCaller, error) {
	contract, err := bindLottery(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LotteryCaller{contract: contract}, nil
}

// NewLotteryTransactor creates a new write-only instance of Lottery, bound to a specific deployed contract.
func NewLotteryTransactor(address common.Address, transactor bind.ContractTransactor) (*LotteryTransactor, error) {
	contract, err := bindLottery(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LotteryTransactor{contract: contract}, nil
}

// NewLotteryFilterer creates a new log filterer instance of Lottery, bound to a specific deployed contract.
func NewLotteryFilterer(address common.Address, filterer bind.ContractFilterer) (*LotteryFilterer, error) {
	contract, err := bindLottery(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LotteryFilterer{contract: contract}, nil
}

// bindLottery binds a generic wrapper to an already deployed contract.
func bindLottery(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LotteryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Lottery *LotteryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Lottery.Contract.LotteryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Lottery *LotteryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lottery.Contract.LotteryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Lottery *LotteryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Lottery.Contract.LotteryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Lottery *LotteryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Lottery.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Lottery *LotteryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lottery.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Lottery *LotteryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Lottery.Contract.contract.Transact(opts, method, params...)
}

// AccountList is a free data retrieval call binding the contract method 0x4ea98d16.
//
// Solidity: function accountList(uint256 ) view returns(address addr, string nums)
func (_Lottery *LotteryCaller) AccountList(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Addr common.Address
	Nums string
}, error) {
	var out []interface{}
	err := _Lottery.contract.Call(opts, &out, "accountList", arg0)

	outstruct := new(struct {
		Addr common.Address
		Nums string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Addr = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Nums = *abi.ConvertType(out[1], new(string)).(*string)

	return *outstruct, err

}

// AccountList is a free data retrieval call binding the contract method 0x4ea98d16.
//
// Solidity: function accountList(uint256 ) view returns(address addr, string nums)
func (_Lottery *LotterySession) AccountList(arg0 *big.Int) (struct {
	Addr common.Address
	Nums string
}, error) {
	return _Lottery.Contract.AccountList(&_Lottery.CallOpts, arg0)
}

// AccountList is a free data retrieval call binding the contract method 0x4ea98d16.
//
// Solidity: function accountList(uint256 ) view returns(address addr, string nums)
func (_Lottery *LotteryCallerSession) AccountList(arg0 *big.Int) (struct {
	Addr common.Address
	Nums string
}, error) {
	return _Lottery.Contract.AccountList(&_Lottery.CallOpts, arg0)
}

// Accounts is a free data retrieval call binding the contract method 0x2069a73a.
//
// Solidity: function accounts(uint256 , uint256 ) view returns(address addr, string nums)
func (_Lottery *LotteryCaller) Accounts(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (struct {
	Addr common.Address
	Nums string
}, error) {
	var out []interface{}
	err := _Lottery.contract.Call(opts, &out, "accounts", arg0, arg1)

	outstruct := new(struct {
		Addr common.Address
		Nums string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Addr = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Nums = *abi.ConvertType(out[1], new(string)).(*string)

	return *outstruct, err

}

// Accounts is a free data retrieval call binding the contract method 0x2069a73a.
//
// Solidity: function accounts(uint256 , uint256 ) view returns(address addr, string nums)
func (_Lottery *LotterySession) Accounts(arg0 *big.Int, arg1 *big.Int) (struct {
	Addr common.Address
	Nums string
}, error) {
	return _Lottery.Contract.Accounts(&_Lottery.CallOpts, arg0, arg1)
}

// Accounts is a free data retrieval call binding the contract method 0x2069a73a.
//
// Solidity: function accounts(uint256 , uint256 ) view returns(address addr, string nums)
func (_Lottery *LotteryCallerSession) Accounts(arg0 *big.Int, arg1 *big.Int) (struct {
	Addr common.Address
	Nums string
}, error) {
	return _Lottery.Contract.Accounts(&_Lottery.CallOpts, arg0, arg1)
}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address addr) view returns(uint256)
func (_Lottery *LotteryCaller) GetBalance(opts *bind.CallOpts, addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Lottery.contract.Call(opts, &out, "getBalance", addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address addr) view returns(uint256)
func (_Lottery *LotterySession) GetBalance(addr common.Address) (*big.Int, error) {
	return _Lottery.Contract.GetBalance(&_Lottery.CallOpts, addr)
}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address addr) view returns(uint256)
func (_Lottery *LotteryCallerSession) GetBalance(addr common.Address) (*big.Int, error) {
	return _Lottery.Contract.GetBalance(&_Lottery.CallOpts, addr)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Lottery *LotteryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Lottery.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Lottery *LotterySession) Owner() (common.Address, error) {
	return _Lottery.Contract.Owner(&_Lottery.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Lottery *LotteryCallerSession) Owner() (common.Address, error) {
	return _Lottery.Contract.Owner(&_Lottery.CallOpts)
}

// Period is a free data retrieval call binding the contract method 0xef78d4fd.
//
// Solidity: function period() view returns(uint256)
func (_Lottery *LotteryCaller) Period(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lottery.contract.Call(opts, &out, "period")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Period is a free data retrieval call binding the contract method 0xef78d4fd.
//
// Solidity: function period() view returns(uint256)
func (_Lottery *LotterySession) Period() (*big.Int, error) {
	return _Lottery.Contract.Period(&_Lottery.CallOpts)
}

// Period is a free data retrieval call binding the contract method 0xef78d4fd.
//
// Solidity: function period() view returns(uint256)
func (_Lottery *LotteryCallerSession) Period() (*big.Int, error) {
	return _Lottery.Contract.Period(&_Lottery.CallOpts)
}

// RetNums is a free data retrieval call binding the contract method 0x59d64215.
//
// Solidity: function retNums(uint256 , uint256 ) view returns(uint256)
func (_Lottery *LotteryCaller) RetNums(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Lottery.contract.Call(opts, &out, "retNums", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RetNums is a free data retrieval call binding the contract method 0x59d64215.
//
// Solidity: function retNums(uint256 , uint256 ) view returns(uint256)
func (_Lottery *LotterySession) RetNums(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _Lottery.Contract.RetNums(&_Lottery.CallOpts, arg0, arg1)
}

// RetNums is a free data retrieval call binding the contract method 0x59d64215.
//
// Solidity: function retNums(uint256 , uint256 ) view returns(uint256)
func (_Lottery *LotteryCallerSession) RetNums(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _Lottery.Contract.RetNums(&_Lottery.CallOpts, arg0, arg1)
}

// RetNumsList is a free data retrieval call binding the contract method 0xc2c26b62.
//
// Solidity: function retNumsList(uint256 ) view returns(uint256)
func (_Lottery *LotteryCaller) RetNumsList(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Lottery.contract.Call(opts, &out, "retNumsList", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RetNumsList is a free data retrieval call binding the contract method 0xc2c26b62.
//
// Solidity: function retNumsList(uint256 ) view returns(uint256)
func (_Lottery *LotterySession) RetNumsList(arg0 *big.Int) (*big.Int, error) {
	return _Lottery.Contract.RetNumsList(&_Lottery.CallOpts, arg0)
}

// RetNumsList is a free data retrieval call binding the contract method 0xc2c26b62.
//
// Solidity: function retNumsList(uint256 ) view returns(uint256)
func (_Lottery *LotteryCallerSession) RetNumsList(arg0 *big.Int) (*big.Int, error) {
	return _Lottery.Contract.RetNumsList(&_Lottery.CallOpts, arg0)
}

// Bet is a paid mutator transaction binding the contract method 0x74030531.
//
// Solidity: function bet(string _nums) payable returns()
func (_Lottery *LotteryTransactor) Bet(opts *bind.TransactOpts, _nums string) (*types.Transaction, error) {
	return _Lottery.contract.Transact(opts, "bet", _nums)
}

// Bet is a paid mutator transaction binding the contract method 0x74030531.
//
// Solidity: function bet(string _nums) payable returns()
func (_Lottery *LotterySession) Bet(_nums string) (*types.Transaction, error) {
	return _Lottery.Contract.Bet(&_Lottery.TransactOpts, _nums)
}

// Bet is a paid mutator transaction binding the contract method 0x74030531.
//
// Solidity: function bet(string _nums) payable returns()
func (_Lottery *LotteryTransactorSession) Bet(_nums string) (*types.Transaction, error) {
	return _Lottery.Contract.Bet(&_Lottery.TransactOpts, _nums)
}

// Draw is a paid mutator transaction binding the contract method 0x0eecae21.
//
// Solidity: function draw() returns()
func (_Lottery *LotteryTransactor) Draw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lottery.contract.Transact(opts, "draw")
}

// Draw is a paid mutator transaction binding the contract method 0x0eecae21.
//
// Solidity: function draw() returns()
func (_Lottery *LotterySession) Draw() (*types.Transaction, error) {
	return _Lottery.Contract.Draw(&_Lottery.TransactOpts)
}

// Draw is a paid mutator transaction binding the contract method 0x0eecae21.
//
// Solidity: function draw() returns()
func (_Lottery *LotteryTransactorSession) Draw() (*types.Transaction, error) {
	return _Lottery.Contract.Draw(&_Lottery.TransactOpts)
}

// LotteryBetFinishIterator is returned from FilterBetFinish and is used to iterate over the raw logs and unpacked data for BetFinish events raised by the Lottery contract.
type LotteryBetFinishIterator struct {
	Event *LotteryBetFinish // Event containing the contract specifics and raw log

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
func (it *LotteryBetFinishIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LotteryBetFinish)
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
		it.Event = new(LotteryBetFinish)
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
func (it *LotteryBetFinishIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LotteryBetFinishIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LotteryBetFinish represents a BetFinish event raised by the Lottery contract.
type LotteryBetFinish struct {
	Addr common.Address
	Nums string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterBetFinish is a free log retrieval operation binding the contract event 0xe6a709e620781c0796d8e03377cdbdcdb7297e37e63fa9541832bd0538e58b14.
//
// Solidity: event BetFinish(address addr, string nums)
func (_Lottery *LotteryFilterer) FilterBetFinish(opts *bind.FilterOpts) (*LotteryBetFinishIterator, error) {

	logs, sub, err := _Lottery.contract.FilterLogs(opts, "BetFinish")
	if err != nil {
		return nil, err
	}
	return &LotteryBetFinishIterator{contract: _Lottery.contract, event: "BetFinish", logs: logs, sub: sub}, nil
}

// WatchBetFinish is a free log subscription operation binding the contract event 0xe6a709e620781c0796d8e03377cdbdcdb7297e37e63fa9541832bd0538e58b14.
//
// Solidity: event BetFinish(address addr, string nums)
func (_Lottery *LotteryFilterer) WatchBetFinish(opts *bind.WatchOpts, sink chan<- *LotteryBetFinish) (event.Subscription, error) {

	logs, sub, err := _Lottery.contract.WatchLogs(opts, "BetFinish")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LotteryBetFinish)
				if err := _Lottery.contract.UnpackLog(event, "BetFinish", log); err != nil {
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

// ParseBetFinish is a log parse operation binding the contract event 0xe6a709e620781c0796d8e03377cdbdcdb7297e37e63fa9541832bd0538e58b14.
//
// Solidity: event BetFinish(address addr, string nums)
func (_Lottery *LotteryFilterer) ParseBetFinish(log types.Log) (*LotteryBetFinish, error) {
	event := new(LotteryBetFinish)
	if err := _Lottery.contract.UnpackLog(event, "BetFinish", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LotteryDrawFinishIterator is returned from FilterDrawFinish and is used to iterate over the raw logs and unpacked data for DrawFinish events raised by the Lottery contract.
type LotteryDrawFinishIterator struct {
	Event *LotteryDrawFinish // Event containing the contract specifics and raw log

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
func (it *LotteryDrawFinishIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LotteryDrawFinish)
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
		it.Event = new(LotteryDrawFinish)
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
func (it *LotteryDrawFinishIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LotteryDrawFinishIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LotteryDrawFinish represents a DrawFinish event raised by the Lottery contract.
type LotteryDrawFinish struct {
	Nums []*big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterDrawFinish is a free log retrieval operation binding the contract event 0xad7be3a9d1eb2a0bb5c9b39014e9d6519886349942fab5459bd584dcb251e6c4.
//
// Solidity: event DrawFinish(uint256[] nums)
func (_Lottery *LotteryFilterer) FilterDrawFinish(opts *bind.FilterOpts) (*LotteryDrawFinishIterator, error) {

	logs, sub, err := _Lottery.contract.FilterLogs(opts, "DrawFinish")
	if err != nil {
		return nil, err
	}
	return &LotteryDrawFinishIterator{contract: _Lottery.contract, event: "DrawFinish", logs: logs, sub: sub}, nil
}

// WatchDrawFinish is a free log subscription operation binding the contract event 0xad7be3a9d1eb2a0bb5c9b39014e9d6519886349942fab5459bd584dcb251e6c4.
//
// Solidity: event DrawFinish(uint256[] nums)
func (_Lottery *LotteryFilterer) WatchDrawFinish(opts *bind.WatchOpts, sink chan<- *LotteryDrawFinish) (event.Subscription, error) {

	logs, sub, err := _Lottery.contract.WatchLogs(opts, "DrawFinish")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LotteryDrawFinish)
				if err := _Lottery.contract.UnpackLog(event, "DrawFinish", log); err != nil {
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

// ParseDrawFinish is a log parse operation binding the contract event 0xad7be3a9d1eb2a0bb5c9b39014e9d6519886349942fab5459bd584dcb251e6c4.
//
// Solidity: event DrawFinish(uint256[] nums)
func (_Lottery *LotteryFilterer) ParseDrawFinish(log types.Log) (*LotteryDrawFinish, error) {
	event := new(LotteryDrawFinish)
	if err := _Lottery.contract.UnpackLog(event, "DrawFinish", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
