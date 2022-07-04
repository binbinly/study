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

// PokemonMarketOrder is an auto generated low-level Go binding around an user-defined struct.
type PokemonMarketOrder struct {
	Id       *big.Int
	Owner    common.Address
	Nft      common.Address
	NftId    *big.Int
	Token    common.Address
	Price    *big.Int
	Create   *big.Int
	Buyer    common.Address
	DealTime *big.Int
}

// MarketABI is the input ABI used to generate the binding from.
const MarketABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"Trade\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"FEE_DENOMINATOR\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"addManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"buy\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"nftIds\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"cancelOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"detail\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"create\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"dealTime\",\"type\":\"uint256\"}],\"internalType\":\"structPokemonMarket.Order\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getMyOrdersLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"startIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultLength\",\"type\":\"uint256\"}],\"name\":\"getMyPurchasedOrders\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"create\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"dealTime\",\"type\":\"uint256\"}],\"internalType\":\"structPokemonMarket.Order[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"startIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endIndex\",\"type\":\"uint256\"}],\"name\":\"getMySoldOrders\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"create\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"dealTime\",\"type\":\"uint256\"}],\"internalType\":\"structPokemonMarket.Order[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"startIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"resultLength\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getOrders\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"create\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"dealTime\",\"type\":\"uint256\"}],\"internalType\":\"structPokemonMarket.Order[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"number\",\"type\":\"uint256\"}],\"name\":\"getRecentlySold\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"create\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"dealTime\",\"type\":\"uint256\"}],\"internalType\":\"structPokemonMarket.Order[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"}],\"name\":\"history\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"create\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"dealTime\",\"type\":\"uint256\"}],\"internalType\":\"structPokemonMarket.Order[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"idCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"incomeOf\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"incomes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"manager\",\"outputs\":[{\"internalType\":\"contractManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"myNftSoldOrders\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"create\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"dealTime\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"myOrderNums\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"myPurchasedOrders\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"create\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"dealTime\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"myPurchasedOrdersNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"mySoldNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nftAddress\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"nftIndexes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"nftOrders\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"create\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"dealTime\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"orderIndexes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"orderNums\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"orders\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"create\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"dealTime\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"ordersNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"recentlySold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"create\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"dealTime\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nftId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"sell\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"nftIds\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"prices\",\"type\":\"uint256[]\"}],\"name\":\"sells\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeAddr\",\"type\":\"address\"}],\"name\":\"setFeeAddr\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"enable\",\"type\":\"bool\"}],\"name\":\"setNftAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"enable\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"priceMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeRatio\",\"type\":\"uint256\"}],\"name\":\"setTokenWhite\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokenWhites\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"enabled\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"priceMin\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amounts\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"counter\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userIdCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Market is an auto generated Go binding around an Ethereum contract.
type Market struct {
	MarketCaller     // Read-only binding to the contract
	MarketTransactor // Write-only binding to the contract
	MarketFilterer   // Log filterer for contract events
}

// MarketCaller is an auto generated read-only Go binding around an Ethereum contract.
type MarketCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MarketTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MarketFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MarketSession struct {
	Contract     *Market           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MarketCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MarketCallerSession struct {
	Contract *MarketCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MarketTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MarketTransactorSession struct {
	Contract     *MarketTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MarketRaw is an auto generated low-level Go binding around an Ethereum contract.
type MarketRaw struct {
	Contract *Market // Generic contract binding to access the raw methods on
}

// MarketCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MarketCallerRaw struct {
	Contract *MarketCaller // Generic read-only contract binding to access the raw methods on
}

// MarketTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MarketTransactorRaw struct {
	Contract *MarketTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMarket creates a new instance of Market, bound to a specific deployed contract.
func NewMarket(address common.Address, backend bind.ContractBackend) (*Market, error) {
	contract, err := bindMarket(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Market{MarketCaller: MarketCaller{contract: contract}, MarketTransactor: MarketTransactor{contract: contract}, MarketFilterer: MarketFilterer{contract: contract}}, nil
}

// NewMarketCaller creates a new read-only instance of Market, bound to a specific deployed contract.
func NewMarketCaller(address common.Address, caller bind.ContractCaller) (*MarketCaller, error) {
	contract, err := bindMarket(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MarketCaller{contract: contract}, nil
}

// NewMarketTransactor creates a new write-only instance of Market, bound to a specific deployed contract.
func NewMarketTransactor(address common.Address, transactor bind.ContractTransactor) (*MarketTransactor, error) {
	contract, err := bindMarket(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MarketTransactor{contract: contract}, nil
}

// NewMarketFilterer creates a new log filterer instance of Market, bound to a specific deployed contract.
func NewMarketFilterer(address common.Address, filterer bind.ContractFilterer) (*MarketFilterer, error) {
	contract, err := bindMarket(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MarketFilterer{contract: contract}, nil
}

// bindMarket binds a generic wrapper to an already deployed contract.
func bindMarket(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MarketABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Market *MarketRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Market.Contract.MarketCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Market *MarketRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Market.Contract.MarketTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Market *MarketRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Market.Contract.MarketTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Market *MarketCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Market.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Market *MarketTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Market.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Market *MarketTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Market.Contract.contract.Transact(opts, method, params...)
}

// FEEDENOMINATOR is a free data retrieval call binding the contract method 0xd73792a9.
//
// Solidity: function FEE_DENOMINATOR() view returns(uint256)
func (_Market *MarketCaller) FEEDENOMINATOR(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "FEE_DENOMINATOR")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FEEDENOMINATOR is a free data retrieval call binding the contract method 0xd73792a9.
//
// Solidity: function FEE_DENOMINATOR() view returns(uint256)
func (_Market *MarketSession) FEEDENOMINATOR() (*big.Int, error) {
	return _Market.Contract.FEEDENOMINATOR(&_Market.CallOpts)
}

// FEEDENOMINATOR is a free data retrieval call binding the contract method 0xd73792a9.
//
// Solidity: function FEE_DENOMINATOR() view returns(uint256)
func (_Market *MarketCallerSession) FEEDENOMINATOR() (*big.Int, error) {
	return _Market.Contract.FEEDENOMINATOR(&_Market.CallOpts)
}

// Detail is a free data retrieval call binding the contract method 0x3cff1b3f.
//
// Solidity: function detail(address nft, uint256 nftId, address token) view returns((uint256,address,address,uint256,address,uint256,uint256,address,uint256))
func (_Market *MarketCaller) Detail(opts *bind.CallOpts, nft common.Address, nftId *big.Int, token common.Address) (PokemonMarketOrder, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "detail", nft, nftId, token)

	if err != nil {
		return *new(PokemonMarketOrder), err
	}

	out0 := *abi.ConvertType(out[0], new(PokemonMarketOrder)).(*PokemonMarketOrder)

	return out0, err

}

// Detail is a free data retrieval call binding the contract method 0x3cff1b3f.
//
// Solidity: function detail(address nft, uint256 nftId, address token) view returns((uint256,address,address,uint256,address,uint256,uint256,address,uint256))
func (_Market *MarketSession) Detail(nft common.Address, nftId *big.Int, token common.Address) (PokemonMarketOrder, error) {
	return _Market.Contract.Detail(&_Market.CallOpts, nft, nftId, token)
}

// Detail is a free data retrieval call binding the contract method 0x3cff1b3f.
//
// Solidity: function detail(address nft, uint256 nftId, address token) view returns((uint256,address,address,uint256,address,uint256,uint256,address,uint256))
func (_Market *MarketCallerSession) Detail(nft common.Address, nftId *big.Int, token common.Address) (PokemonMarketOrder, error) {
	return _Market.Contract.Detail(&_Market.CallOpts, nft, nftId, token)
}

// FeeAddr is a free data retrieval call binding the contract method 0x39e7fddc.
//
// Solidity: function feeAddr() view returns(address)
func (_Market *MarketCaller) FeeAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "feeAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeAddr is a free data retrieval call binding the contract method 0x39e7fddc.
//
// Solidity: function feeAddr() view returns(address)
func (_Market *MarketSession) FeeAddr() (common.Address, error) {
	return _Market.Contract.FeeAddr(&_Market.CallOpts)
}

// FeeAddr is a free data retrieval call binding the contract method 0x39e7fddc.
//
// Solidity: function feeAddr() view returns(address)
func (_Market *MarketCallerSession) FeeAddr() (common.Address, error) {
	return _Market.Contract.FeeAddr(&_Market.CallOpts)
}

// GetMyOrdersLength is a free data retrieval call binding the contract method 0x5ac24eb0.
//
// Solidity: function getMyOrdersLength(address owner, address nft, address token) view returns(uint256)
func (_Market *MarketCaller) GetMyOrdersLength(opts *bind.CallOpts, owner common.Address, nft common.Address, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "getMyOrdersLength", owner, nft, token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMyOrdersLength is a free data retrieval call binding the contract method 0x5ac24eb0.
//
// Solidity: function getMyOrdersLength(address owner, address nft, address token) view returns(uint256)
func (_Market *MarketSession) GetMyOrdersLength(owner common.Address, nft common.Address, token common.Address) (*big.Int, error) {
	return _Market.Contract.GetMyOrdersLength(&_Market.CallOpts, owner, nft, token)
}

// GetMyOrdersLength is a free data retrieval call binding the contract method 0x5ac24eb0.
//
// Solidity: function getMyOrdersLength(address owner, address nft, address token) view returns(uint256)
func (_Market *MarketCallerSession) GetMyOrdersLength(owner common.Address, nft common.Address, token common.Address) (*big.Int, error) {
	return _Market.Contract.GetMyOrdersLength(&_Market.CallOpts, owner, nft, token)
}

// GetMyPurchasedOrders is a free data retrieval call binding the contract method 0xc841f6dd.
//
// Solidity: function getMyPurchasedOrders(address owner, uint256 startIndex, uint256 endIndex, uint256 resultLength) view returns((uint256,address,address,uint256,address,uint256,uint256,address,uint256)[])
func (_Market *MarketCaller) GetMyPurchasedOrders(opts *bind.CallOpts, owner common.Address, startIndex *big.Int, endIndex *big.Int, resultLength *big.Int) ([]PokemonMarketOrder, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "getMyPurchasedOrders", owner, startIndex, endIndex, resultLength)

	if err != nil {
		return *new([]PokemonMarketOrder), err
	}

	out0 := *abi.ConvertType(out[0], new([]PokemonMarketOrder)).(*[]PokemonMarketOrder)

	return out0, err

}

// GetMyPurchasedOrders is a free data retrieval call binding the contract method 0xc841f6dd.
//
// Solidity: function getMyPurchasedOrders(address owner, uint256 startIndex, uint256 endIndex, uint256 resultLength) view returns((uint256,address,address,uint256,address,uint256,uint256,address,uint256)[])
func (_Market *MarketSession) GetMyPurchasedOrders(owner common.Address, startIndex *big.Int, endIndex *big.Int, resultLength *big.Int) ([]PokemonMarketOrder, error) {
	return _Market.Contract.GetMyPurchasedOrders(&_Market.CallOpts, owner, startIndex, endIndex, resultLength)
}

// GetMyPurchasedOrders is a free data retrieval call binding the contract method 0xc841f6dd.
//
// Solidity: function getMyPurchasedOrders(address owner, uint256 startIndex, uint256 endIndex, uint256 resultLength) view returns((uint256,address,address,uint256,address,uint256,uint256,address,uint256)[])
func (_Market *MarketCallerSession) GetMyPurchasedOrders(owner common.Address, startIndex *big.Int, endIndex *big.Int, resultLength *big.Int) ([]PokemonMarketOrder, error) {
	return _Market.Contract.GetMyPurchasedOrders(&_Market.CallOpts, owner, startIndex, endIndex, resultLength)
}

// GetMySoldOrders is a free data retrieval call binding the contract method 0xedc7e315.
//
// Solidity: function getMySoldOrders(address owner, uint256 startIndex, uint256 endIndex) view returns((uint256,address,address,uint256,address,uint256,uint256,address,uint256)[])
func (_Market *MarketCaller) GetMySoldOrders(opts *bind.CallOpts, owner common.Address, startIndex *big.Int, endIndex *big.Int) ([]PokemonMarketOrder, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "getMySoldOrders", owner, startIndex, endIndex)

	if err != nil {
		return *new([]PokemonMarketOrder), err
	}

	out0 := *abi.ConvertType(out[0], new([]PokemonMarketOrder)).(*[]PokemonMarketOrder)

	return out0, err

}

// GetMySoldOrders is a free data retrieval call binding the contract method 0xedc7e315.
//
// Solidity: function getMySoldOrders(address owner, uint256 startIndex, uint256 endIndex) view returns((uint256,address,address,uint256,address,uint256,uint256,address,uint256)[])
func (_Market *MarketSession) GetMySoldOrders(owner common.Address, startIndex *big.Int, endIndex *big.Int) ([]PokemonMarketOrder, error) {
	return _Market.Contract.GetMySoldOrders(&_Market.CallOpts, owner, startIndex, endIndex)
}

// GetMySoldOrders is a free data retrieval call binding the contract method 0xedc7e315.
//
// Solidity: function getMySoldOrders(address owner, uint256 startIndex, uint256 endIndex) view returns((uint256,address,address,uint256,address,uint256,uint256,address,uint256)[])
func (_Market *MarketCallerSession) GetMySoldOrders(owner common.Address, startIndex *big.Int, endIndex *big.Int) ([]PokemonMarketOrder, error) {
	return _Market.Contract.GetMySoldOrders(&_Market.CallOpts, owner, startIndex, endIndex)
}

// GetOrders is a free data retrieval call binding the contract method 0x3e827b6f.
//
// Solidity: function getOrders(uint256 startIndex, uint256 endIndex, uint256 resultLength, address owner, address nft, address token) view returns((uint256,address,address,uint256,address,uint256,uint256,address,uint256)[])
func (_Market *MarketCaller) GetOrders(opts *bind.CallOpts, startIndex *big.Int, endIndex *big.Int, resultLength *big.Int, owner common.Address, nft common.Address, token common.Address) ([]PokemonMarketOrder, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "getOrders", startIndex, endIndex, resultLength, owner, nft, token)

	if err != nil {
		return *new([]PokemonMarketOrder), err
	}

	out0 := *abi.ConvertType(out[0], new([]PokemonMarketOrder)).(*[]PokemonMarketOrder)

	return out0, err

}

// GetOrders is a free data retrieval call binding the contract method 0x3e827b6f.
//
// Solidity: function getOrders(uint256 startIndex, uint256 endIndex, uint256 resultLength, address owner, address nft, address token) view returns((uint256,address,address,uint256,address,uint256,uint256,address,uint256)[])
func (_Market *MarketSession) GetOrders(startIndex *big.Int, endIndex *big.Int, resultLength *big.Int, owner common.Address, nft common.Address, token common.Address) ([]PokemonMarketOrder, error) {
	return _Market.Contract.GetOrders(&_Market.CallOpts, startIndex, endIndex, resultLength, owner, nft, token)
}

// GetOrders is a free data retrieval call binding the contract method 0x3e827b6f.
//
// Solidity: function getOrders(uint256 startIndex, uint256 endIndex, uint256 resultLength, address owner, address nft, address token) view returns((uint256,address,address,uint256,address,uint256,uint256,address,uint256)[])
func (_Market *MarketCallerSession) GetOrders(startIndex *big.Int, endIndex *big.Int, resultLength *big.Int, owner common.Address, nft common.Address, token common.Address) ([]PokemonMarketOrder, error) {
	return _Market.Contract.GetOrders(&_Market.CallOpts, startIndex, endIndex, resultLength, owner, nft, token)
}

// GetRecentlySold is a free data retrieval call binding the contract method 0x133cb7fc.
//
// Solidity: function getRecentlySold(address nft, uint256 number) view returns((uint256,address,address,uint256,address,uint256,uint256,address,uint256)[])
func (_Market *MarketCaller) GetRecentlySold(opts *bind.CallOpts, nft common.Address, number *big.Int) ([]PokemonMarketOrder, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "getRecentlySold", nft, number)

	if err != nil {
		return *new([]PokemonMarketOrder), err
	}

	out0 := *abi.ConvertType(out[0], new([]PokemonMarketOrder)).(*[]PokemonMarketOrder)

	return out0, err

}

// GetRecentlySold is a free data retrieval call binding the contract method 0x133cb7fc.
//
// Solidity: function getRecentlySold(address nft, uint256 number) view returns((uint256,address,address,uint256,address,uint256,uint256,address,uint256)[])
func (_Market *MarketSession) GetRecentlySold(nft common.Address, number *big.Int) ([]PokemonMarketOrder, error) {
	return _Market.Contract.GetRecentlySold(&_Market.CallOpts, nft, number)
}

// GetRecentlySold is a free data retrieval call binding the contract method 0x133cb7fc.
//
// Solidity: function getRecentlySold(address nft, uint256 number) view returns((uint256,address,address,uint256,address,uint256,uint256,address,uint256)[])
func (_Market *MarketCallerSession) GetRecentlySold(nft common.Address, number *big.Int) ([]PokemonMarketOrder, error) {
	return _Market.Contract.GetRecentlySold(&_Market.CallOpts, nft, number)
}

// History is a free data retrieval call binding the contract method 0x7718f4ec.
//
// Solidity: function history(address nft, uint256 nftId) view returns((uint256,address,address,uint256,address,uint256,uint256,address,uint256)[])
func (_Market *MarketCaller) History(opts *bind.CallOpts, nft common.Address, nftId *big.Int) ([]PokemonMarketOrder, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "history", nft, nftId)

	if err != nil {
		return *new([]PokemonMarketOrder), err
	}

	out0 := *abi.ConvertType(out[0], new([]PokemonMarketOrder)).(*[]PokemonMarketOrder)

	return out0, err

}

// History is a free data retrieval call binding the contract method 0x7718f4ec.
//
// Solidity: function history(address nft, uint256 nftId) view returns((uint256,address,address,uint256,address,uint256,uint256,address,uint256)[])
func (_Market *MarketSession) History(nft common.Address, nftId *big.Int) ([]PokemonMarketOrder, error) {
	return _Market.Contract.History(&_Market.CallOpts, nft, nftId)
}

// History is a free data retrieval call binding the contract method 0x7718f4ec.
//
// Solidity: function history(address nft, uint256 nftId) view returns((uint256,address,address,uint256,address,uint256,uint256,address,uint256)[])
func (_Market *MarketCallerSession) History(nft common.Address, nftId *big.Int) ([]PokemonMarketOrder, error) {
	return _Market.Contract.History(&_Market.CallOpts, nft, nftId)
}

// IdCount is a free data retrieval call binding the contract method 0xe6e46238.
//
// Solidity: function idCount() view returns(uint256)
func (_Market *MarketCaller) IdCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "idCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IdCount is a free data retrieval call binding the contract method 0xe6e46238.
//
// Solidity: function idCount() view returns(uint256)
func (_Market *MarketSession) IdCount() (*big.Int, error) {
	return _Market.Contract.IdCount(&_Market.CallOpts)
}

// IdCount is a free data retrieval call binding the contract method 0xe6e46238.
//
// Solidity: function idCount() view returns(uint256)
func (_Market *MarketCallerSession) IdCount() (*big.Int, error) {
	return _Market.Contract.IdCount(&_Market.CallOpts)
}

// IncomeOf is a free data retrieval call binding the contract method 0x3b3c340d.
//
// Solidity: function incomeOf(address user, address[] tokens) view returns(uint256[])
func (_Market *MarketCaller) IncomeOf(opts *bind.CallOpts, user common.Address, tokens []common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "incomeOf", user, tokens)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// IncomeOf is a free data retrieval call binding the contract method 0x3b3c340d.
//
// Solidity: function incomeOf(address user, address[] tokens) view returns(uint256[])
func (_Market *MarketSession) IncomeOf(user common.Address, tokens []common.Address) ([]*big.Int, error) {
	return _Market.Contract.IncomeOf(&_Market.CallOpts, user, tokens)
}

// IncomeOf is a free data retrieval call binding the contract method 0x3b3c340d.
//
// Solidity: function incomeOf(address user, address[] tokens) view returns(uint256[])
func (_Market *MarketCallerSession) IncomeOf(user common.Address, tokens []common.Address) ([]*big.Int, error) {
	return _Market.Contract.IncomeOf(&_Market.CallOpts, user, tokens)
}

// Incomes is a free data retrieval call binding the contract method 0xa7a78282.
//
// Solidity: function incomes(address , address ) view returns(uint256)
func (_Market *MarketCaller) Incomes(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "incomes", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Incomes is a free data retrieval call binding the contract method 0xa7a78282.
//
// Solidity: function incomes(address , address ) view returns(uint256)
func (_Market *MarketSession) Incomes(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Market.Contract.Incomes(&_Market.CallOpts, arg0, arg1)
}

// Incomes is a free data retrieval call binding the contract method 0xa7a78282.
//
// Solidity: function incomes(address , address ) view returns(uint256)
func (_Market *MarketCallerSession) Incomes(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Market.Contract.Incomes(&_Market.CallOpts, arg0, arg1)
}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() view returns(address)
func (_Market *MarketCaller) Manager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "manager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() view returns(address)
func (_Market *MarketSession) Manager() (common.Address, error) {
	return _Market.Contract.Manager(&_Market.CallOpts)
}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() view returns(address)
func (_Market *MarketCallerSession) Manager() (common.Address, error) {
	return _Market.Contract.Manager(&_Market.CallOpts)
}

// MyNftSoldOrders is a free data retrieval call binding the contract method 0xe27ada1a.
//
// Solidity: function myNftSoldOrders(address , uint256 ) view returns(uint256 id, address owner, address nft, uint256 nftId, address token, uint256 price, uint256 create, address buyer, uint256 dealTime)
func (_Market *MarketCaller) MyNftSoldOrders(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	Id       *big.Int
	Owner    common.Address
	Nft      common.Address
	NftId    *big.Int
	Token    common.Address
	Price    *big.Int
	Create   *big.Int
	Buyer    common.Address
	DealTime *big.Int
}, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "myNftSoldOrders", arg0, arg1)

	outstruct := new(struct {
		Id       *big.Int
		Owner    common.Address
		Nft      common.Address
		NftId    *big.Int
		Token    common.Address
		Price    *big.Int
		Create   *big.Int
		Buyer    common.Address
		DealTime *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Owner = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Nft = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.NftId = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Token = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Price = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.Create = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.Buyer = *abi.ConvertType(out[7], new(common.Address)).(*common.Address)
	outstruct.DealTime = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// MyNftSoldOrders is a free data retrieval call binding the contract method 0xe27ada1a.
//
// Solidity: function myNftSoldOrders(address , uint256 ) view returns(uint256 id, address owner, address nft, uint256 nftId, address token, uint256 price, uint256 create, address buyer, uint256 dealTime)
func (_Market *MarketSession) MyNftSoldOrders(arg0 common.Address, arg1 *big.Int) (struct {
	Id       *big.Int
	Owner    common.Address
	Nft      common.Address
	NftId    *big.Int
	Token    common.Address
	Price    *big.Int
	Create   *big.Int
	Buyer    common.Address
	DealTime *big.Int
}, error) {
	return _Market.Contract.MyNftSoldOrders(&_Market.CallOpts, arg0, arg1)
}

// MyNftSoldOrders is a free data retrieval call binding the contract method 0xe27ada1a.
//
// Solidity: function myNftSoldOrders(address , uint256 ) view returns(uint256 id, address owner, address nft, uint256 nftId, address token, uint256 price, uint256 create, address buyer, uint256 dealTime)
func (_Market *MarketCallerSession) MyNftSoldOrders(arg0 common.Address, arg1 *big.Int) (struct {
	Id       *big.Int
	Owner    common.Address
	Nft      common.Address
	NftId    *big.Int
	Token    common.Address
	Price    *big.Int
	Create   *big.Int
	Buyer    common.Address
	DealTime *big.Int
}, error) {
	return _Market.Contract.MyNftSoldOrders(&_Market.CallOpts, arg0, arg1)
}

// MyOrderNums is a free data retrieval call binding the contract method 0x52ab1acc.
//
// Solidity: function myOrderNums(address , address , address ) view returns(uint256)
func (_Market *MarketCaller) MyOrderNums(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "myOrderNums", arg0, arg1, arg2)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MyOrderNums is a free data retrieval call binding the contract method 0x52ab1acc.
//
// Solidity: function myOrderNums(address , address , address ) view returns(uint256)
func (_Market *MarketSession) MyOrderNums(arg0 common.Address, arg1 common.Address, arg2 common.Address) (*big.Int, error) {
	return _Market.Contract.MyOrderNums(&_Market.CallOpts, arg0, arg1, arg2)
}

// MyOrderNums is a free data retrieval call binding the contract method 0x52ab1acc.
//
// Solidity: function myOrderNums(address , address , address ) view returns(uint256)
func (_Market *MarketCallerSession) MyOrderNums(arg0 common.Address, arg1 common.Address, arg2 common.Address) (*big.Int, error) {
	return _Market.Contract.MyOrderNums(&_Market.CallOpts, arg0, arg1, arg2)
}

// MyPurchasedOrders is a free data retrieval call binding the contract method 0x82b92a2a.
//
// Solidity: function myPurchasedOrders(address , uint256 ) view returns(uint256 id, address owner, address nft, uint256 nftId, address token, uint256 price, uint256 create, address buyer, uint256 dealTime)
func (_Market *MarketCaller) MyPurchasedOrders(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	Id       *big.Int
	Owner    common.Address
	Nft      common.Address
	NftId    *big.Int
	Token    common.Address
	Price    *big.Int
	Create   *big.Int
	Buyer    common.Address
	DealTime *big.Int
}, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "myPurchasedOrders", arg0, arg1)

	outstruct := new(struct {
		Id       *big.Int
		Owner    common.Address
		Nft      common.Address
		NftId    *big.Int
		Token    common.Address
		Price    *big.Int
		Create   *big.Int
		Buyer    common.Address
		DealTime *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Owner = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Nft = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.NftId = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Token = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Price = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.Create = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.Buyer = *abi.ConvertType(out[7], new(common.Address)).(*common.Address)
	outstruct.DealTime = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// MyPurchasedOrders is a free data retrieval call binding the contract method 0x82b92a2a.
//
// Solidity: function myPurchasedOrders(address , uint256 ) view returns(uint256 id, address owner, address nft, uint256 nftId, address token, uint256 price, uint256 create, address buyer, uint256 dealTime)
func (_Market *MarketSession) MyPurchasedOrders(arg0 common.Address, arg1 *big.Int) (struct {
	Id       *big.Int
	Owner    common.Address
	Nft      common.Address
	NftId    *big.Int
	Token    common.Address
	Price    *big.Int
	Create   *big.Int
	Buyer    common.Address
	DealTime *big.Int
}, error) {
	return _Market.Contract.MyPurchasedOrders(&_Market.CallOpts, arg0, arg1)
}

// MyPurchasedOrders is a free data retrieval call binding the contract method 0x82b92a2a.
//
// Solidity: function myPurchasedOrders(address , uint256 ) view returns(uint256 id, address owner, address nft, uint256 nftId, address token, uint256 price, uint256 create, address buyer, uint256 dealTime)
func (_Market *MarketCallerSession) MyPurchasedOrders(arg0 common.Address, arg1 *big.Int) (struct {
	Id       *big.Int
	Owner    common.Address
	Nft      common.Address
	NftId    *big.Int
	Token    common.Address
	Price    *big.Int
	Create   *big.Int
	Buyer    common.Address
	DealTime *big.Int
}, error) {
	return _Market.Contract.MyPurchasedOrders(&_Market.CallOpts, arg0, arg1)
}

// MyPurchasedOrdersNum is a free data retrieval call binding the contract method 0x3968160b.
//
// Solidity: function myPurchasedOrdersNum(address owner) view returns(uint256)
func (_Market *MarketCaller) MyPurchasedOrdersNum(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "myPurchasedOrdersNum", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MyPurchasedOrdersNum is a free data retrieval call binding the contract method 0x3968160b.
//
// Solidity: function myPurchasedOrdersNum(address owner) view returns(uint256)
func (_Market *MarketSession) MyPurchasedOrdersNum(owner common.Address) (*big.Int, error) {
	return _Market.Contract.MyPurchasedOrdersNum(&_Market.CallOpts, owner)
}

// MyPurchasedOrdersNum is a free data retrieval call binding the contract method 0x3968160b.
//
// Solidity: function myPurchasedOrdersNum(address owner) view returns(uint256)
func (_Market *MarketCallerSession) MyPurchasedOrdersNum(owner common.Address) (*big.Int, error) {
	return _Market.Contract.MyPurchasedOrdersNum(&_Market.CallOpts, owner)
}

// MySoldNum is a free data retrieval call binding the contract method 0xdd3b0de1.
//
// Solidity: function mySoldNum(address owner) view returns(uint256)
func (_Market *MarketCaller) MySoldNum(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "mySoldNum", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MySoldNum is a free data retrieval call binding the contract method 0xdd3b0de1.
//
// Solidity: function mySoldNum(address owner) view returns(uint256)
func (_Market *MarketSession) MySoldNum(owner common.Address) (*big.Int, error) {
	return _Market.Contract.MySoldNum(&_Market.CallOpts, owner)
}

// MySoldNum is a free data retrieval call binding the contract method 0xdd3b0de1.
//
// Solidity: function mySoldNum(address owner) view returns(uint256)
func (_Market *MarketCallerSession) MySoldNum(owner common.Address) (*big.Int, error) {
	return _Market.Contract.MySoldNum(&_Market.CallOpts, owner)
}

// NftAddress is a free data retrieval call binding the contract method 0x99d2041f.
//
// Solidity: function nftAddress(address ) view returns(bool)
func (_Market *MarketCaller) NftAddress(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "nftAddress", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// NftAddress is a free data retrieval call binding the contract method 0x99d2041f.
//
// Solidity: function nftAddress(address ) view returns(bool)
func (_Market *MarketSession) NftAddress(arg0 common.Address) (bool, error) {
	return _Market.Contract.NftAddress(&_Market.CallOpts, arg0)
}

// NftAddress is a free data retrieval call binding the contract method 0x99d2041f.
//
// Solidity: function nftAddress(address ) view returns(bool)
func (_Market *MarketCallerSession) NftAddress(arg0 common.Address) (bool, error) {
	return _Market.Contract.NftAddress(&_Market.CallOpts, arg0)
}

// NftIndexes is a free data retrieval call binding the contract method 0xcb0beeed.
//
// Solidity: function nftIndexes(address , uint256 ) view returns(uint256)
func (_Market *MarketCaller) NftIndexes(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "nftIndexes", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NftIndexes is a free data retrieval call binding the contract method 0xcb0beeed.
//
// Solidity: function nftIndexes(address , uint256 ) view returns(uint256)
func (_Market *MarketSession) NftIndexes(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Market.Contract.NftIndexes(&_Market.CallOpts, arg0, arg1)
}

// NftIndexes is a free data retrieval call binding the contract method 0xcb0beeed.
//
// Solidity: function nftIndexes(address , uint256 ) view returns(uint256)
func (_Market *MarketCallerSession) NftIndexes(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Market.Contract.NftIndexes(&_Market.CallOpts, arg0, arg1)
}

// NftOrders is a free data retrieval call binding the contract method 0xd7addda5.
//
// Solidity: function nftOrders(address , uint256 , uint256 ) view returns(uint256 id, address owner, address nft, uint256 nftId, address token, uint256 price, uint256 create, address buyer, uint256 dealTime)
func (_Market *MarketCaller) NftOrders(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int, arg2 *big.Int) (struct {
	Id       *big.Int
	Owner    common.Address
	Nft      common.Address
	NftId    *big.Int
	Token    common.Address
	Price    *big.Int
	Create   *big.Int
	Buyer    common.Address
	DealTime *big.Int
}, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "nftOrders", arg0, arg1, arg2)

	outstruct := new(struct {
		Id       *big.Int
		Owner    common.Address
		Nft      common.Address
		NftId    *big.Int
		Token    common.Address
		Price    *big.Int
		Create   *big.Int
		Buyer    common.Address
		DealTime *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Owner = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Nft = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.NftId = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Token = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Price = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.Create = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.Buyer = *abi.ConvertType(out[7], new(common.Address)).(*common.Address)
	outstruct.DealTime = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// NftOrders is a free data retrieval call binding the contract method 0xd7addda5.
//
// Solidity: function nftOrders(address , uint256 , uint256 ) view returns(uint256 id, address owner, address nft, uint256 nftId, address token, uint256 price, uint256 create, address buyer, uint256 dealTime)
func (_Market *MarketSession) NftOrders(arg0 common.Address, arg1 *big.Int, arg2 *big.Int) (struct {
	Id       *big.Int
	Owner    common.Address
	Nft      common.Address
	NftId    *big.Int
	Token    common.Address
	Price    *big.Int
	Create   *big.Int
	Buyer    common.Address
	DealTime *big.Int
}, error) {
	return _Market.Contract.NftOrders(&_Market.CallOpts, arg0, arg1, arg2)
}

// NftOrders is a free data retrieval call binding the contract method 0xd7addda5.
//
// Solidity: function nftOrders(address , uint256 , uint256 ) view returns(uint256 id, address owner, address nft, uint256 nftId, address token, uint256 price, uint256 create, address buyer, uint256 dealTime)
func (_Market *MarketCallerSession) NftOrders(arg0 common.Address, arg1 *big.Int, arg2 *big.Int) (struct {
	Id       *big.Int
	Owner    common.Address
	Nft      common.Address
	NftId    *big.Int
	Token    common.Address
	Price    *big.Int
	Create   *big.Int
	Buyer    common.Address
	DealTime *big.Int
}, error) {
	return _Market.Contract.NftOrders(&_Market.CallOpts, arg0, arg1, arg2)
}

// OrderIndexes is a free data retrieval call binding the contract method 0x1a8640de.
//
// Solidity: function orderIndexes(address , uint256 ) view returns(uint256)
func (_Market *MarketCaller) OrderIndexes(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "orderIndexes", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OrderIndexes is a free data retrieval call binding the contract method 0x1a8640de.
//
// Solidity: function orderIndexes(address , uint256 ) view returns(uint256)
func (_Market *MarketSession) OrderIndexes(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Market.Contract.OrderIndexes(&_Market.CallOpts, arg0, arg1)
}

// OrderIndexes is a free data retrieval call binding the contract method 0x1a8640de.
//
// Solidity: function orderIndexes(address , uint256 ) view returns(uint256)
func (_Market *MarketCallerSession) OrderIndexes(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Market.Contract.OrderIndexes(&_Market.CallOpts, arg0, arg1)
}

// OrderNums is a free data retrieval call binding the contract method 0x5114379c.
//
// Solidity: function orderNums(address , address ) view returns(uint256)
func (_Market *MarketCaller) OrderNums(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "orderNums", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OrderNums is a free data retrieval call binding the contract method 0x5114379c.
//
// Solidity: function orderNums(address , address ) view returns(uint256)
func (_Market *MarketSession) OrderNums(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Market.Contract.OrderNums(&_Market.CallOpts, arg0, arg1)
}

// OrderNums is a free data retrieval call binding the contract method 0x5114379c.
//
// Solidity: function orderNums(address , address ) view returns(uint256)
func (_Market *MarketCallerSession) OrderNums(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Market.Contract.OrderNums(&_Market.CallOpts, arg0, arg1)
}

// Orders is a free data retrieval call binding the contract method 0xdd02df16.
//
// Solidity: function orders(address , address , uint256 ) view returns(uint256 id, address owner, address nft, uint256 nftId, address token, uint256 price, uint256 create, address buyer, uint256 dealTime)
func (_Market *MarketCaller) Orders(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address, arg2 *big.Int) (struct {
	Id       *big.Int
	Owner    common.Address
	Nft      common.Address
	NftId    *big.Int
	Token    common.Address
	Price    *big.Int
	Create   *big.Int
	Buyer    common.Address
	DealTime *big.Int
}, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "orders", arg0, arg1, arg2)

	outstruct := new(struct {
		Id       *big.Int
		Owner    common.Address
		Nft      common.Address
		NftId    *big.Int
		Token    common.Address
		Price    *big.Int
		Create   *big.Int
		Buyer    common.Address
		DealTime *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Owner = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Nft = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.NftId = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Token = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Price = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.Create = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.Buyer = *abi.ConvertType(out[7], new(common.Address)).(*common.Address)
	outstruct.DealTime = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Orders is a free data retrieval call binding the contract method 0xdd02df16.
//
// Solidity: function orders(address , address , uint256 ) view returns(uint256 id, address owner, address nft, uint256 nftId, address token, uint256 price, uint256 create, address buyer, uint256 dealTime)
func (_Market *MarketSession) Orders(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (struct {
	Id       *big.Int
	Owner    common.Address
	Nft      common.Address
	NftId    *big.Int
	Token    common.Address
	Price    *big.Int
	Create   *big.Int
	Buyer    common.Address
	DealTime *big.Int
}, error) {
	return _Market.Contract.Orders(&_Market.CallOpts, arg0, arg1, arg2)
}

// Orders is a free data retrieval call binding the contract method 0xdd02df16.
//
// Solidity: function orders(address , address , uint256 ) view returns(uint256 id, address owner, address nft, uint256 nftId, address token, uint256 price, uint256 create, address buyer, uint256 dealTime)
func (_Market *MarketCallerSession) Orders(arg0 common.Address, arg1 common.Address, arg2 *big.Int) (struct {
	Id       *big.Int
	Owner    common.Address
	Nft      common.Address
	NftId    *big.Int
	Token    common.Address
	Price    *big.Int
	Create   *big.Int
	Buyer    common.Address
	DealTime *big.Int
}, error) {
	return _Market.Contract.Orders(&_Market.CallOpts, arg0, arg1, arg2)
}

// OrdersNum is a free data retrieval call binding the contract method 0xd9e39329.
//
// Solidity: function ordersNum(address nft, address token) view returns(uint256)
func (_Market *MarketCaller) OrdersNum(opts *bind.CallOpts, nft common.Address, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "ordersNum", nft, token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OrdersNum is a free data retrieval call binding the contract method 0xd9e39329.
//
// Solidity: function ordersNum(address nft, address token) view returns(uint256)
func (_Market *MarketSession) OrdersNum(nft common.Address, token common.Address) (*big.Int, error) {
	return _Market.Contract.OrdersNum(&_Market.CallOpts, nft, token)
}

// OrdersNum is a free data retrieval call binding the contract method 0xd9e39329.
//
// Solidity: function ordersNum(address nft, address token) view returns(uint256)
func (_Market *MarketCallerSession) OrdersNum(nft common.Address, token common.Address) (*big.Int, error) {
	return _Market.Contract.OrdersNum(&_Market.CallOpts, nft, token)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Market *MarketCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Market *MarketSession) Owner() (common.Address, error) {
	return _Market.Contract.Owner(&_Market.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Market *MarketCallerSession) Owner() (common.Address, error) {
	return _Market.Contract.Owner(&_Market.CallOpts)
}

// RecentlySold is a free data retrieval call binding the contract method 0x6a14ba48.
//
// Solidity: function recentlySold(address , uint256 ) view returns(uint256 id, address owner, address nft, uint256 nftId, address token, uint256 price, uint256 create, address buyer, uint256 dealTime)
func (_Market *MarketCaller) RecentlySold(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	Id       *big.Int
	Owner    common.Address
	Nft      common.Address
	NftId    *big.Int
	Token    common.Address
	Price    *big.Int
	Create   *big.Int
	Buyer    common.Address
	DealTime *big.Int
}, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "recentlySold", arg0, arg1)

	outstruct := new(struct {
		Id       *big.Int
		Owner    common.Address
		Nft      common.Address
		NftId    *big.Int
		Token    common.Address
		Price    *big.Int
		Create   *big.Int
		Buyer    common.Address
		DealTime *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Owner = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Nft = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.NftId = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Token = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Price = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.Create = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.Buyer = *abi.ConvertType(out[7], new(common.Address)).(*common.Address)
	outstruct.DealTime = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// RecentlySold is a free data retrieval call binding the contract method 0x6a14ba48.
//
// Solidity: function recentlySold(address , uint256 ) view returns(uint256 id, address owner, address nft, uint256 nftId, address token, uint256 price, uint256 create, address buyer, uint256 dealTime)
func (_Market *MarketSession) RecentlySold(arg0 common.Address, arg1 *big.Int) (struct {
	Id       *big.Int
	Owner    common.Address
	Nft      common.Address
	NftId    *big.Int
	Token    common.Address
	Price    *big.Int
	Create   *big.Int
	Buyer    common.Address
	DealTime *big.Int
}, error) {
	return _Market.Contract.RecentlySold(&_Market.CallOpts, arg0, arg1)
}

// RecentlySold is a free data retrieval call binding the contract method 0x6a14ba48.
//
// Solidity: function recentlySold(address , uint256 ) view returns(uint256 id, address owner, address nft, uint256 nftId, address token, uint256 price, uint256 create, address buyer, uint256 dealTime)
func (_Market *MarketCallerSession) RecentlySold(arg0 common.Address, arg1 *big.Int) (struct {
	Id       *big.Int
	Owner    common.Address
	Nft      common.Address
	NftId    *big.Int
	Token    common.Address
	Price    *big.Int
	Create   *big.Int
	Buyer    common.Address
	DealTime *big.Int
}, error) {
	return _Market.Contract.RecentlySold(&_Market.CallOpts, arg0, arg1)
}

// TokenWhites is a free data retrieval call binding the contract method 0x128b6796.
//
// Solidity: function tokenWhites(address , address ) view returns(bool enabled, uint256 priceMin, uint256 amounts, uint256 counter, uint256 feeRatio, uint256 fee)
func (_Market *MarketCaller) TokenWhites(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (struct {
	Enabled  bool
	PriceMin *big.Int
	Amounts  *big.Int
	Counter  *big.Int
	FeeRatio *big.Int
	Fee      *big.Int
}, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "tokenWhites", arg0, arg1)

	outstruct := new(struct {
		Enabled  bool
		PriceMin *big.Int
		Amounts  *big.Int
		Counter  *big.Int
		FeeRatio *big.Int
		Fee      *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Enabled = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.PriceMin = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Amounts = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Counter = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.FeeRatio = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Fee = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// TokenWhites is a free data retrieval call binding the contract method 0x128b6796.
//
// Solidity: function tokenWhites(address , address ) view returns(bool enabled, uint256 priceMin, uint256 amounts, uint256 counter, uint256 feeRatio, uint256 fee)
func (_Market *MarketSession) TokenWhites(arg0 common.Address, arg1 common.Address) (struct {
	Enabled  bool
	PriceMin *big.Int
	Amounts  *big.Int
	Counter  *big.Int
	FeeRatio *big.Int
	Fee      *big.Int
}, error) {
	return _Market.Contract.TokenWhites(&_Market.CallOpts, arg0, arg1)
}

// TokenWhites is a free data retrieval call binding the contract method 0x128b6796.
//
// Solidity: function tokenWhites(address , address ) view returns(bool enabled, uint256 priceMin, uint256 amounts, uint256 counter, uint256 feeRatio, uint256 fee)
func (_Market *MarketCallerSession) TokenWhites(arg0 common.Address, arg1 common.Address) (struct {
	Enabled  bool
	PriceMin *big.Int
	Amounts  *big.Int
	Counter  *big.Int
	FeeRatio *big.Int
	Fee      *big.Int
}, error) {
	return _Market.Contract.TokenWhites(&_Market.CallOpts, arg0, arg1)
}

// UserIdCount is a free data retrieval call binding the contract method 0x76a7f8b7.
//
// Solidity: function userIdCount(address ) view returns(uint256)
func (_Market *MarketCaller) UserIdCount(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Market.contract.Call(opts, &out, "userIdCount", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserIdCount is a free data retrieval call binding the contract method 0x76a7f8b7.
//
// Solidity: function userIdCount(address ) view returns(uint256)
func (_Market *MarketSession) UserIdCount(arg0 common.Address) (*big.Int, error) {
	return _Market.Contract.UserIdCount(&_Market.CallOpts, arg0)
}

// UserIdCount is a free data retrieval call binding the contract method 0x76a7f8b7.
//
// Solidity: function userIdCount(address ) view returns(uint256)
func (_Market *MarketCallerSession) UserIdCount(arg0 common.Address) (*big.Int, error) {
	return _Market.Contract.UserIdCount(&_Market.CallOpts, arg0)
}

// AddManager is a paid mutator transaction binding the contract method 0x2d06177a.
//
// Solidity: function addManager(address addr) returns()
func (_Market *MarketTransactor) AddManager(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "addManager", addr)
}

// AddManager is a paid mutator transaction binding the contract method 0x2d06177a.
//
// Solidity: function addManager(address addr) returns()
func (_Market *MarketSession) AddManager(addr common.Address) (*types.Transaction, error) {
	return _Market.Contract.AddManager(&_Market.TransactOpts, addr)
}

// AddManager is a paid mutator transaction binding the contract method 0x2d06177a.
//
// Solidity: function addManager(address addr) returns()
func (_Market *MarketTransactorSession) AddManager(addr common.Address) (*types.Transaction, error) {
	return _Market.Contract.AddManager(&_Market.TransactOpts, addr)
}

// Buy is a paid mutator transaction binding the contract method 0xdb61c76e.
//
// Solidity: function buy(address nft, uint256 nftId, address token) payable returns()
func (_Market *MarketTransactor) Buy(opts *bind.TransactOpts, nft common.Address, nftId *big.Int, token common.Address) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "buy", nft, nftId, token)
}

// Buy is a paid mutator transaction binding the contract method 0xdb61c76e.
//
// Solidity: function buy(address nft, uint256 nftId, address token) payable returns()
func (_Market *MarketSession) Buy(nft common.Address, nftId *big.Int, token common.Address) (*types.Transaction, error) {
	return _Market.Contract.Buy(&_Market.TransactOpts, nft, nftId, token)
}

// Buy is a paid mutator transaction binding the contract method 0xdb61c76e.
//
// Solidity: function buy(address nft, uint256 nftId, address token) payable returns()
func (_Market *MarketTransactorSession) Buy(nft common.Address, nftId *big.Int, token common.Address) (*types.Transaction, error) {
	return _Market.Contract.Buy(&_Market.TransactOpts, nft, nftId, token)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x8a7b919a.
//
// Solidity: function cancelOrder(address nft, uint256[] nftIds, address token) returns()
func (_Market *MarketTransactor) CancelOrder(opts *bind.TransactOpts, nft common.Address, nftIds []*big.Int, token common.Address) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "cancelOrder", nft, nftIds, token)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x8a7b919a.
//
// Solidity: function cancelOrder(address nft, uint256[] nftIds, address token) returns()
func (_Market *MarketSession) CancelOrder(nft common.Address, nftIds []*big.Int, token common.Address) (*types.Transaction, error) {
	return _Market.Contract.CancelOrder(&_Market.TransactOpts, nft, nftIds, token)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x8a7b919a.
//
// Solidity: function cancelOrder(address nft, uint256[] nftIds, address token) returns()
func (_Market *MarketTransactorSession) CancelOrder(nft common.Address, nftIds []*big.Int, token common.Address) (*types.Transaction, error) {
	return _Market.Contract.CancelOrder(&_Market.TransactOpts, nft, nftIds, token)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Market *MarketTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Market *MarketSession) RenounceOwnership() (*types.Transaction, error) {
	return _Market.Contract.RenounceOwnership(&_Market.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Market *MarketTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Market.Contract.RenounceOwnership(&_Market.TransactOpts)
}

// Sell is a paid mutator transaction binding the contract method 0x627eb0d4.
//
// Solidity: function sell(address nft, uint256 nftId, address token, uint256 price) returns()
func (_Market *MarketTransactor) Sell(opts *bind.TransactOpts, nft common.Address, nftId *big.Int, token common.Address, price *big.Int) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "sell", nft, nftId, token, price)
}

// Sell is a paid mutator transaction binding the contract method 0x627eb0d4.
//
// Solidity: function sell(address nft, uint256 nftId, address token, uint256 price) returns()
func (_Market *MarketSession) Sell(nft common.Address, nftId *big.Int, token common.Address, price *big.Int) (*types.Transaction, error) {
	return _Market.Contract.Sell(&_Market.TransactOpts, nft, nftId, token, price)
}

// Sell is a paid mutator transaction binding the contract method 0x627eb0d4.
//
// Solidity: function sell(address nft, uint256 nftId, address token, uint256 price) returns()
func (_Market *MarketTransactorSession) Sell(nft common.Address, nftId *big.Int, token common.Address, price *big.Int) (*types.Transaction, error) {
	return _Market.Contract.Sell(&_Market.TransactOpts, nft, nftId, token, price)
}

// Sells is a paid mutator transaction binding the contract method 0x60c0fe84.
//
// Solidity: function sells(address nft, uint256[] nftIds, address token, uint256[] prices) returns()
func (_Market *MarketTransactor) Sells(opts *bind.TransactOpts, nft common.Address, nftIds []*big.Int, token common.Address, prices []*big.Int) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "sells", nft, nftIds, token, prices)
}

// Sells is a paid mutator transaction binding the contract method 0x60c0fe84.
//
// Solidity: function sells(address nft, uint256[] nftIds, address token, uint256[] prices) returns()
func (_Market *MarketSession) Sells(nft common.Address, nftIds []*big.Int, token common.Address, prices []*big.Int) (*types.Transaction, error) {
	return _Market.Contract.Sells(&_Market.TransactOpts, nft, nftIds, token, prices)
}

// Sells is a paid mutator transaction binding the contract method 0x60c0fe84.
//
// Solidity: function sells(address nft, uint256[] nftIds, address token, uint256[] prices) returns()
func (_Market *MarketTransactorSession) Sells(nft common.Address, nftIds []*big.Int, token common.Address, prices []*big.Int) (*types.Transaction, error) {
	return _Market.Contract.Sells(&_Market.TransactOpts, nft, nftIds, token, prices)
}

// SetFeeAddr is a paid mutator transaction binding the contract method 0xb2855b4f.
//
// Solidity: function setFeeAddr(address _feeAddr) returns()
func (_Market *MarketTransactor) SetFeeAddr(opts *bind.TransactOpts, _feeAddr common.Address) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "setFeeAddr", _feeAddr)
}

// SetFeeAddr is a paid mutator transaction binding the contract method 0xb2855b4f.
//
// Solidity: function setFeeAddr(address _feeAddr) returns()
func (_Market *MarketSession) SetFeeAddr(_feeAddr common.Address) (*types.Transaction, error) {
	return _Market.Contract.SetFeeAddr(&_Market.TransactOpts, _feeAddr)
}

// SetFeeAddr is a paid mutator transaction binding the contract method 0xb2855b4f.
//
// Solidity: function setFeeAddr(address _feeAddr) returns()
func (_Market *MarketTransactorSession) SetFeeAddr(_feeAddr common.Address) (*types.Transaction, error) {
	return _Market.Contract.SetFeeAddr(&_Market.TransactOpts, _feeAddr)
}

// SetNftAddress is a paid mutator transaction binding the contract method 0x813133c4.
//
// Solidity: function setNftAddress(address addr, bool enable) returns()
func (_Market *MarketTransactor) SetNftAddress(opts *bind.TransactOpts, addr common.Address, enable bool) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "setNftAddress", addr, enable)
}

// SetNftAddress is a paid mutator transaction binding the contract method 0x813133c4.
//
// Solidity: function setNftAddress(address addr, bool enable) returns()
func (_Market *MarketSession) SetNftAddress(addr common.Address, enable bool) (*types.Transaction, error) {
	return _Market.Contract.SetNftAddress(&_Market.TransactOpts, addr, enable)
}

// SetNftAddress is a paid mutator transaction binding the contract method 0x813133c4.
//
// Solidity: function setNftAddress(address addr, bool enable) returns()
func (_Market *MarketTransactorSession) SetNftAddress(addr common.Address, enable bool) (*types.Transaction, error) {
	return _Market.Contract.SetNftAddress(&_Market.TransactOpts, addr, enable)
}

// SetTokenWhite is a paid mutator transaction binding the contract method 0x5426831d.
//
// Solidity: function setTokenWhite(address nft, address token, bool enable, uint256 priceMin, uint256 feeRatio) returns()
func (_Market *MarketTransactor) SetTokenWhite(opts *bind.TransactOpts, nft common.Address, token common.Address, enable bool, priceMin *big.Int, feeRatio *big.Int) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "setTokenWhite", nft, token, enable, priceMin, feeRatio)
}

// SetTokenWhite is a paid mutator transaction binding the contract method 0x5426831d.
//
// Solidity: function setTokenWhite(address nft, address token, bool enable, uint256 priceMin, uint256 feeRatio) returns()
func (_Market *MarketSession) SetTokenWhite(nft common.Address, token common.Address, enable bool, priceMin *big.Int, feeRatio *big.Int) (*types.Transaction, error) {
	return _Market.Contract.SetTokenWhite(&_Market.TransactOpts, nft, token, enable, priceMin, feeRatio)
}

// SetTokenWhite is a paid mutator transaction binding the contract method 0x5426831d.
//
// Solidity: function setTokenWhite(address nft, address token, bool enable, uint256 priceMin, uint256 feeRatio) returns()
func (_Market *MarketTransactorSession) SetTokenWhite(nft common.Address, token common.Address, enable bool, priceMin *big.Int, feeRatio *big.Int) (*types.Transaction, error) {
	return _Market.Contract.SetTokenWhite(&_Market.TransactOpts, nft, token, enable, priceMin, feeRatio)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Market *MarketTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Market.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Market *MarketSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Market.Contract.TransferOwnership(&_Market.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Market *MarketTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Market.Contract.TransferOwnership(&_Market.TransactOpts, newOwner)
}

// MarketOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Market contract.
type MarketOwnershipTransferredIterator struct {
	Event *MarketOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MarketOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketOwnershipTransferred)
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
		it.Event = new(MarketOwnershipTransferred)
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
func (it *MarketOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketOwnershipTransferred represents a OwnershipTransferred event raised by the Market contract.
type MarketOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Market *MarketFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MarketOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Market.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MarketOwnershipTransferredIterator{contract: _Market.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Market *MarketFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MarketOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Market.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketOwnershipTransferred)
				if err := _Market.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Market *MarketFilterer) ParseOwnershipTransferred(log types.Log) (*MarketOwnershipTransferred, error) {
	event := new(MarketOwnershipTransferred)
	if err := _Market.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketTradeIterator is returned from FilterTrade and is used to iterate over the raw logs and unpacked data for Trade events raised by the Market contract.
type MarketTradeIterator struct {
	Event *MarketTrade // Event containing the contract specifics and raw log

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
func (it *MarketTradeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketTrade)
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
		it.Event = new(MarketTrade)
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
func (it *MarketTradeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketTradeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketTrade represents a Trade event raised by the Market contract.
type MarketTrade struct {
	Id    *big.Int
	From  common.Address
	To    common.Address
	Nft   common.Address
	NftId *big.Int
	Token common.Address
	Price *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTrade is a free log retrieval operation binding the contract event 0xf49141886563be203826fd184c8487cc09ff268ebfc5508583ac909f60924aef.
//
// Solidity: event Trade(uint256 indexed id, address indexed from, address indexed to, address nft, uint256 nftId, address token, uint256 price)
func (_Market *MarketFilterer) FilterTrade(opts *bind.FilterOpts, id []*big.Int, from []common.Address, to []common.Address) (*MarketTradeIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Market.contract.FilterLogs(opts, "Trade", idRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &MarketTradeIterator{contract: _Market.contract, event: "Trade", logs: logs, sub: sub}, nil
}

// WatchTrade is a free log subscription operation binding the contract event 0xf49141886563be203826fd184c8487cc09ff268ebfc5508583ac909f60924aef.
//
// Solidity: event Trade(uint256 indexed id, address indexed from, address indexed to, address nft, uint256 nftId, address token, uint256 price)
func (_Market *MarketFilterer) WatchTrade(opts *bind.WatchOpts, sink chan<- *MarketTrade, id []*big.Int, from []common.Address, to []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Market.contract.WatchLogs(opts, "Trade", idRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketTrade)
				if err := _Market.contract.UnpackLog(event, "Trade", log); err != nil {
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

// ParseTrade is a log parse operation binding the contract event 0xf49141886563be203826fd184c8487cc09ff268ebfc5508583ac909f60924aef.
//
// Solidity: event Trade(uint256 indexed id, address indexed from, address indexed to, address nft, uint256 nftId, address token, uint256 price)
func (_Market *MarketFilterer) ParseTrade(log types.Log) (*MarketTrade, error) {
	event := new(MarketTrade)
	if err := _Market.contract.UnpackLog(event, "Trade", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
