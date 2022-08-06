// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package moonflip

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
)

// MoonFlipCommit is an auto generated low-level Go binding around an user-defined struct.
type MoonFlipCommit struct {
	BlockNumber *big.Int
	Nonce       *big.Int
	Value       *big.Int
	Guess       uint8
	Exists      bool
}

// MoonflipMetaData contains all meta data concerning the Moonflip contract.
var MoonflipMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumMoonFlip.Game\",\"name\":\"game\",\"type\":\"uint8\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"guess\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"}],\"indexed\":false,\"internalType\":\"structMoonFlip.Commit\",\"name\":\"commit\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"GuessSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"enumMoonFlip.Game\",\"name\":\"game\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"won\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"participant\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Outcome\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"addAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"betAmounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"gameId\",\"type\":\"uint256\"}],\"name\":\"clearPendingCommit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumMoonFlip.CoinSide\",\"name\":\"guess\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"betIndex\",\"type\":\"uint256\"}],\"name\":\"flipCoin\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"houseAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"player\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"gameId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"randomNumber\",\"type\":\"uint256\"}],\"name\":\"processGuess\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"revokeAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"ranges\",\"type\":\"uint256[]\"}],\"name\":\"setBetRange\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"houseAddress_\",\"type\":\"address\"}],\"name\":\"setHouseAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"}],\"name\":\"setLiquidityThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"state\",\"type\":\"bool\"}],\"name\":\"setPaused\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// MoonflipABI is the input ABI used to generate the binding from.
// Deprecated: Use MoonflipMetaData.ABI instead.
var MoonflipABI = MoonflipMetaData.ABI

// Moonflip is an auto generated Go binding around an Ethereum contract.
type Moonflip struct {
	MoonflipCaller     // Read-only binding to the contract
	MoonflipTransactor // Write-only binding to the contract
	MoonflipFilterer   // Log filterer for contract events
}

// MoonflipCaller is an auto generated read-only Go binding around an Ethereum contract.
type MoonflipCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MoonflipTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MoonflipTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MoonflipFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MoonflipFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MoonflipSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MoonflipSession struct {
	Contract     *Moonflip         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MoonflipCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MoonflipCallerSession struct {
	Contract *MoonflipCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// MoonflipTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MoonflipTransactorSession struct {
	Contract     *MoonflipTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// MoonflipRaw is an auto generated low-level Go binding around an Ethereum contract.
type MoonflipRaw struct {
	Contract *Moonflip // Generic contract binding to access the raw methods on
}

// MoonflipCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MoonflipCallerRaw struct {
	Contract *MoonflipCaller // Generic read-only contract binding to access the raw methods on
}

// MoonflipTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MoonflipTransactorRaw struct {
	Contract *MoonflipTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMoonflip creates a new instance of Moonflip, bound to a specific deployed contract.
func NewMoonflip(address common.Address, backend bind.ContractBackend) (*Moonflip, error) {
	contract, err := bindMoonflip(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Moonflip{MoonflipCaller: MoonflipCaller{contract: contract}, MoonflipTransactor: MoonflipTransactor{contract: contract}, MoonflipFilterer: MoonflipFilterer{contract: contract}}, nil
}

// NewMoonflipCaller creates a new read-only instance of Moonflip, bound to a specific deployed contract.
func NewMoonflipCaller(address common.Address, caller bind.ContractCaller) (*MoonflipCaller, error) {
	contract, err := bindMoonflip(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MoonflipCaller{contract: contract}, nil
}

// NewMoonflipTransactor creates a new write-only instance of Moonflip, bound to a specific deployed contract.
func NewMoonflipTransactor(address common.Address, transactor bind.ContractTransactor) (*MoonflipTransactor, error) {
	contract, err := bindMoonflip(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MoonflipTransactor{contract: contract}, nil
}

// NewMoonflipFilterer creates a new log filterer instance of Moonflip, bound to a specific deployed contract.
func NewMoonflipFilterer(address common.Address, filterer bind.ContractFilterer) (*MoonflipFilterer, error) {
	contract, err := bindMoonflip(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MoonflipFilterer{contract: contract}, nil
}

// bindMoonflip binds a generic wrapper to an already deployed contract.
func bindMoonflip(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MoonflipABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Moonflip *MoonflipRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Moonflip.Contract.MoonflipCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Moonflip *MoonflipRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Moonflip.Contract.MoonflipTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Moonflip *MoonflipRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Moonflip.Contract.MoonflipTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Moonflip *MoonflipCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Moonflip.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Moonflip *MoonflipTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Moonflip.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Moonflip *MoonflipTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Moonflip.Contract.contract.Transact(opts, method, params...)
}

// BetAmounts is a free data retrieval call binding the contract method 0x74410429.
//
// Solidity: function betAmounts(uint256 ) view returns(uint256)
func (_Moonflip *MoonflipCaller) BetAmounts(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Moonflip.contract.Call(opts, &out, "betAmounts", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BetAmounts is a free data retrieval call binding the contract method 0x74410429.
//
// Solidity: function betAmounts(uint256 ) view returns(uint256)
func (_Moonflip *MoonflipSession) BetAmounts(arg0 *big.Int) (*big.Int, error) {
	return _Moonflip.Contract.BetAmounts(&_Moonflip.CallOpts, arg0)
}

// BetAmounts is a free data retrieval call binding the contract method 0x74410429.
//
// Solidity: function betAmounts(uint256 ) view returns(uint256)
func (_Moonflip *MoonflipCallerSession) BetAmounts(arg0 *big.Int) (*big.Int, error) {
	return _Moonflip.Contract.BetAmounts(&_Moonflip.CallOpts, arg0)
}

// HouseAddress is a free data retrieval call binding the contract method 0xd7cee31e.
//
// Solidity: function houseAddress() view returns(address)
func (_Moonflip *MoonflipCaller) HouseAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Moonflip.contract.Call(opts, &out, "houseAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// HouseAddress is a free data retrieval call binding the contract method 0xd7cee31e.
//
// Solidity: function houseAddress() view returns(address)
func (_Moonflip *MoonflipSession) HouseAddress() (common.Address, error) {
	return _Moonflip.Contract.HouseAddress(&_Moonflip.CallOpts)
}

// HouseAddress is a free data retrieval call binding the contract method 0xd7cee31e.
//
// Solidity: function houseAddress() view returns(address)
func (_Moonflip *MoonflipCallerSession) HouseAddress() (common.Address, error) {
	return _Moonflip.Contract.HouseAddress(&_Moonflip.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Moonflip *MoonflipCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Moonflip.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Moonflip *MoonflipSession) Owner() (common.Address, error) {
	return _Moonflip.Contract.Owner(&_Moonflip.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Moonflip *MoonflipCallerSession) Owner() (common.Address, error) {
	return _Moonflip.Contract.Owner(&_Moonflip.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Moonflip *MoonflipCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Moonflip.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Moonflip *MoonflipSession) Paused() (bool, error) {
	return _Moonflip.Contract.Paused(&_Moonflip.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Moonflip *MoonflipCallerSession) Paused() (bool, error) {
	return _Moonflip.Contract.Paused(&_Moonflip.CallOpts)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address addr) returns()
func (_Moonflip *MoonflipTransactor) AddAdmin(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Moonflip.contract.Transact(opts, "addAdmin", addr)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address addr) returns()
func (_Moonflip *MoonflipSession) AddAdmin(addr common.Address) (*types.Transaction, error) {
	return _Moonflip.Contract.AddAdmin(&_Moonflip.TransactOpts, addr)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address addr) returns()
func (_Moonflip *MoonflipTransactorSession) AddAdmin(addr common.Address) (*types.Transaction, error) {
	return _Moonflip.Contract.AddAdmin(&_Moonflip.TransactOpts, addr)
}

// ClearPendingCommit is a paid mutator transaction binding the contract method 0x96ea025f.
//
// Solidity: function clearPendingCommit(address addr, uint256 gameId) returns()
func (_Moonflip *MoonflipTransactor) ClearPendingCommit(opts *bind.TransactOpts, addr common.Address, gameId *big.Int) (*types.Transaction, error) {
	return _Moonflip.contract.Transact(opts, "clearPendingCommit", addr, gameId)
}

// ClearPendingCommit is a paid mutator transaction binding the contract method 0x96ea025f.
//
// Solidity: function clearPendingCommit(address addr, uint256 gameId) returns()
func (_Moonflip *MoonflipSession) ClearPendingCommit(addr common.Address, gameId *big.Int) (*types.Transaction, error) {
	return _Moonflip.Contract.ClearPendingCommit(&_Moonflip.TransactOpts, addr, gameId)
}

// ClearPendingCommit is a paid mutator transaction binding the contract method 0x96ea025f.
//
// Solidity: function clearPendingCommit(address addr, uint256 gameId) returns()
func (_Moonflip *MoonflipTransactorSession) ClearPendingCommit(addr common.Address, gameId *big.Int) (*types.Transaction, error) {
	return _Moonflip.Contract.ClearPendingCommit(&_Moonflip.TransactOpts, addr, gameId)
}

// FlipCoin is a paid mutator transaction binding the contract method 0xf3b4d4cb.
//
// Solidity: function flipCoin(uint8 guess, uint256 betIndex) payable returns()
func (_Moonflip *MoonflipTransactor) FlipCoin(opts *bind.TransactOpts, guess uint8, betIndex *big.Int) (*types.Transaction, error) {
	return _Moonflip.contract.Transact(opts, "flipCoin", guess, betIndex)
}

// FlipCoin is a paid mutator transaction binding the contract method 0xf3b4d4cb.
//
// Solidity: function flipCoin(uint8 guess, uint256 betIndex) payable returns()
func (_Moonflip *MoonflipSession) FlipCoin(guess uint8, betIndex *big.Int) (*types.Transaction, error) {
	return _Moonflip.Contract.FlipCoin(&_Moonflip.TransactOpts, guess, betIndex)
}

// FlipCoin is a paid mutator transaction binding the contract method 0xf3b4d4cb.
//
// Solidity: function flipCoin(uint8 guess, uint256 betIndex) payable returns()
func (_Moonflip *MoonflipTransactorSession) FlipCoin(guess uint8, betIndex *big.Int) (*types.Transaction, error) {
	return _Moonflip.Contract.FlipCoin(&_Moonflip.TransactOpts, guess, betIndex)
}

// ProcessGuess is a paid mutator transaction binding the contract method 0xc2b411a6.
//
// Solidity: function processGuess(address player, uint256 gameId, uint256 randomNumber) returns()
func (_Moonflip *MoonflipTransactor) ProcessGuess(opts *bind.TransactOpts, player common.Address, gameId *big.Int, randomNumber *big.Int) (*types.Transaction, error) {
	return _Moonflip.contract.Transact(opts, "processGuess", player, gameId, randomNumber)
}

// ProcessGuess is a paid mutator transaction binding the contract method 0xc2b411a6.
//
// Solidity: function processGuess(address player, uint256 gameId, uint256 randomNumber) returns()
func (_Moonflip *MoonflipSession) ProcessGuess(player common.Address, gameId *big.Int, randomNumber *big.Int) (*types.Transaction, error) {
	return _Moonflip.Contract.ProcessGuess(&_Moonflip.TransactOpts, player, gameId, randomNumber)
}

// ProcessGuess is a paid mutator transaction binding the contract method 0xc2b411a6.
//
// Solidity: function processGuess(address player, uint256 gameId, uint256 randomNumber) returns()
func (_Moonflip *MoonflipTransactorSession) ProcessGuess(player common.Address, gameId *big.Int, randomNumber *big.Int) (*types.Transaction, error) {
	return _Moonflip.Contract.ProcessGuess(&_Moonflip.TransactOpts, player, gameId, randomNumber)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Moonflip *MoonflipTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Moonflip.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Moonflip *MoonflipSession) RenounceOwnership() (*types.Transaction, error) {
	return _Moonflip.Contract.RenounceOwnership(&_Moonflip.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Moonflip *MoonflipTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Moonflip.Contract.RenounceOwnership(&_Moonflip.TransactOpts)
}

// RevokeAdmin is a paid mutator transaction binding the contract method 0x2d345670.
//
// Solidity: function revokeAdmin(address addr) returns()
func (_Moonflip *MoonflipTransactor) RevokeAdmin(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Moonflip.contract.Transact(opts, "revokeAdmin", addr)
}

// RevokeAdmin is a paid mutator transaction binding the contract method 0x2d345670.
//
// Solidity: function revokeAdmin(address addr) returns()
func (_Moonflip *MoonflipSession) RevokeAdmin(addr common.Address) (*types.Transaction, error) {
	return _Moonflip.Contract.RevokeAdmin(&_Moonflip.TransactOpts, addr)
}

// RevokeAdmin is a paid mutator transaction binding the contract method 0x2d345670.
//
// Solidity: function revokeAdmin(address addr) returns()
func (_Moonflip *MoonflipTransactorSession) RevokeAdmin(addr common.Address) (*types.Transaction, error) {
	return _Moonflip.Contract.RevokeAdmin(&_Moonflip.TransactOpts, addr)
}

// SetBetRange is a paid mutator transaction binding the contract method 0x12e8ac6f.
//
// Solidity: function setBetRange(uint256[] ranges) returns()
func (_Moonflip *MoonflipTransactor) SetBetRange(opts *bind.TransactOpts, ranges []*big.Int) (*types.Transaction, error) {
	return _Moonflip.contract.Transact(opts, "setBetRange", ranges)
}

// SetBetRange is a paid mutator transaction binding the contract method 0x12e8ac6f.
//
// Solidity: function setBetRange(uint256[] ranges) returns()
func (_Moonflip *MoonflipSession) SetBetRange(ranges []*big.Int) (*types.Transaction, error) {
	return _Moonflip.Contract.SetBetRange(&_Moonflip.TransactOpts, ranges)
}

// SetBetRange is a paid mutator transaction binding the contract method 0x12e8ac6f.
//
// Solidity: function setBetRange(uint256[] ranges) returns()
func (_Moonflip *MoonflipTransactorSession) SetBetRange(ranges []*big.Int) (*types.Transaction, error) {
	return _Moonflip.Contract.SetBetRange(&_Moonflip.TransactOpts, ranges)
}

// SetHouseAddress is a paid mutator transaction binding the contract method 0xef3aaf54.
//
// Solidity: function setHouseAddress(address houseAddress_) returns()
func (_Moonflip *MoonflipTransactor) SetHouseAddress(opts *bind.TransactOpts, houseAddress_ common.Address) (*types.Transaction, error) {
	return _Moonflip.contract.Transact(opts, "setHouseAddress", houseAddress_)
}

// SetHouseAddress is a paid mutator transaction binding the contract method 0xef3aaf54.
//
// Solidity: function setHouseAddress(address houseAddress_) returns()
func (_Moonflip *MoonflipSession) SetHouseAddress(houseAddress_ common.Address) (*types.Transaction, error) {
	return _Moonflip.Contract.SetHouseAddress(&_Moonflip.TransactOpts, houseAddress_)
}

// SetHouseAddress is a paid mutator transaction binding the contract method 0xef3aaf54.
//
// Solidity: function setHouseAddress(address houseAddress_) returns()
func (_Moonflip *MoonflipTransactorSession) SetHouseAddress(houseAddress_ common.Address) (*types.Transaction, error) {
	return _Moonflip.Contract.SetHouseAddress(&_Moonflip.TransactOpts, houseAddress_)
}

// SetLiquidityThreshold is a paid mutator transaction binding the contract method 0x4baf59f3.
//
// Solidity: function setLiquidityThreshold(uint256 liquidity) returns()
func (_Moonflip *MoonflipTransactor) SetLiquidityThreshold(opts *bind.TransactOpts, liquidity *big.Int) (*types.Transaction, error) {
	return _Moonflip.contract.Transact(opts, "setLiquidityThreshold", liquidity)
}

// SetLiquidityThreshold is a paid mutator transaction binding the contract method 0x4baf59f3.
//
// Solidity: function setLiquidityThreshold(uint256 liquidity) returns()
func (_Moonflip *MoonflipSession) SetLiquidityThreshold(liquidity *big.Int) (*types.Transaction, error) {
	return _Moonflip.Contract.SetLiquidityThreshold(&_Moonflip.TransactOpts, liquidity)
}

// SetLiquidityThreshold is a paid mutator transaction binding the contract method 0x4baf59f3.
//
// Solidity: function setLiquidityThreshold(uint256 liquidity) returns()
func (_Moonflip *MoonflipTransactorSession) SetLiquidityThreshold(liquidity *big.Int) (*types.Transaction, error) {
	return _Moonflip.Contract.SetLiquidityThreshold(&_Moonflip.TransactOpts, liquidity)
}

// SetPaused is a paid mutator transaction binding the contract method 0x16c38b3c.
//
// Solidity: function setPaused(bool state) returns()
func (_Moonflip *MoonflipTransactor) SetPaused(opts *bind.TransactOpts, state bool) (*types.Transaction, error) {
	return _Moonflip.contract.Transact(opts, "setPaused", state)
}

// SetPaused is a paid mutator transaction binding the contract method 0x16c38b3c.
//
// Solidity: function setPaused(bool state) returns()
func (_Moonflip *MoonflipSession) SetPaused(state bool) (*types.Transaction, error) {
	return _Moonflip.Contract.SetPaused(&_Moonflip.TransactOpts, state)
}

// SetPaused is a paid mutator transaction binding the contract method 0x16c38b3c.
//
// Solidity: function setPaused(bool state) returns()
func (_Moonflip *MoonflipTransactorSession) SetPaused(state bool) (*types.Transaction, error) {
	return _Moonflip.Contract.SetPaused(&_Moonflip.TransactOpts, state)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Moonflip *MoonflipTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Moonflip.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Moonflip *MoonflipSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Moonflip.Contract.TransferOwnership(&_Moonflip.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Moonflip *MoonflipTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Moonflip.Contract.TransferOwnership(&_Moonflip.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Moonflip *MoonflipTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Moonflip.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Moonflip *MoonflipSession) Withdraw() (*types.Transaction, error) {
	return _Moonflip.Contract.Withdraw(&_Moonflip.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Moonflip *MoonflipTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Moonflip.Contract.Withdraw(&_Moonflip.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Moonflip *MoonflipTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Moonflip.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Moonflip *MoonflipSession) Receive() (*types.Transaction, error) {
	return _Moonflip.Contract.Receive(&_Moonflip.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Moonflip *MoonflipTransactorSession) Receive() (*types.Transaction, error) {
	return _Moonflip.Contract.Receive(&_Moonflip.TransactOpts)
}

// MoonflipGuessSubmittedIterator is returned from FilterGuessSubmitted and is used to iterate over the raw logs and unpacked data for GuessSubmitted events raised by the Moonflip contract.
type MoonflipGuessSubmittedIterator struct {
	Event *MoonflipGuessSubmitted // Event containing the contract specifics and raw log

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
func (it *MoonflipGuessSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MoonflipGuessSubmitted)
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
		it.Event = new(MoonflipGuessSubmitted)
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
func (it *MoonflipGuessSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MoonflipGuessSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MoonflipGuessSubmitted represents a GuessSubmitted event raised by the Moonflip contract.
type MoonflipGuessSubmitted struct {
	Game   uint8
	Commit MoonFlipCommit
	Sender common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterGuessSubmitted is a free log retrieval operation binding the contract event 0x2ffe001dd0156470441991399ed630f59bd6d093d0e680432ef4bc5dc9ae3f40.
//
// Solidity: event GuessSubmitted(uint8 game, (uint256,uint256,uint256,uint8,bool) commit, address sender)
func (_Moonflip *MoonflipFilterer) FilterGuessSubmitted(opts *bind.FilterOpts) (*MoonflipGuessSubmittedIterator, error) {

	logs, sub, err := _Moonflip.contract.FilterLogs(opts, "GuessSubmitted")
	if err != nil {
		return nil, err
	}
	return &MoonflipGuessSubmittedIterator{contract: _Moonflip.contract, event: "GuessSubmitted", logs: logs, sub: sub}, nil
}

// WatchGuessSubmitted is a free log subscription operation binding the contract event 0x2ffe001dd0156470441991399ed630f59bd6d093d0e680432ef4bc5dc9ae3f40.
//
// Solidity: event GuessSubmitted(uint8 game, (uint256,uint256,uint256,uint8,bool) commit, address sender)
func (_Moonflip *MoonflipFilterer) WatchGuessSubmitted(opts *bind.WatchOpts, sink chan<- *MoonflipGuessSubmitted) (event.Subscription, error) {

	logs, sub, err := _Moonflip.contract.WatchLogs(opts, "GuessSubmitted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MoonflipGuessSubmitted)
				if err := _Moonflip.contract.UnpackLog(event, "GuessSubmitted", log); err != nil {
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

// ParseGuessSubmitted is a log parse operation binding the contract event 0x2ffe001dd0156470441991399ed630f59bd6d093d0e680432ef4bc5dc9ae3f40.
//
// Solidity: event GuessSubmitted(uint8 game, (uint256,uint256,uint256,uint8,bool) commit, address sender)
func (_Moonflip *MoonflipFilterer) ParseGuessSubmitted(log types.Log) (*MoonflipGuessSubmitted, error) {
	event := new(MoonflipGuessSubmitted)
	if err := _Moonflip.contract.UnpackLog(event, "GuessSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MoonflipOutcomeIterator is returned from FilterOutcome and is used to iterate over the raw logs and unpacked data for Outcome events raised by the Moonflip contract.
type MoonflipOutcomeIterator struct {
	Event *MoonflipOutcome // Event containing the contract specifics and raw log

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
func (it *MoonflipOutcomeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MoonflipOutcome)
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
		it.Event = new(MoonflipOutcome)
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
func (it *MoonflipOutcomeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MoonflipOutcomeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MoonflipOutcome represents a Outcome event raised by the Moonflip contract.
type MoonflipOutcome struct {
	Game        uint8
	Won         bool
	Participant common.Address
	Value       *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterOutcome is a free log retrieval operation binding the contract event 0xcf9d4ce6bec133891c06506b3d12a01701296e29fdf00d7cd08b1419954d1dfd.
//
// Solidity: event Outcome(uint8 game, bool won, address participant, uint256 value)
func (_Moonflip *MoonflipFilterer) FilterOutcome(opts *bind.FilterOpts) (*MoonflipOutcomeIterator, error) {

	logs, sub, err := _Moonflip.contract.FilterLogs(opts, "Outcome")
	if err != nil {
		return nil, err
	}
	return &MoonflipOutcomeIterator{contract: _Moonflip.contract, event: "Outcome", logs: logs, sub: sub}, nil
}

// WatchOutcome is a free log subscription operation binding the contract event 0xcf9d4ce6bec133891c06506b3d12a01701296e29fdf00d7cd08b1419954d1dfd.
//
// Solidity: event Outcome(uint8 game, bool won, address participant, uint256 value)
func (_Moonflip *MoonflipFilterer) WatchOutcome(opts *bind.WatchOpts, sink chan<- *MoonflipOutcome) (event.Subscription, error) {

	logs, sub, err := _Moonflip.contract.WatchLogs(opts, "Outcome")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MoonflipOutcome)
				if err := _Moonflip.contract.UnpackLog(event, "Outcome", log); err != nil {
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

// ParseOutcome is a log parse operation binding the contract event 0xcf9d4ce6bec133891c06506b3d12a01701296e29fdf00d7cd08b1419954d1dfd.
//
// Solidity: event Outcome(uint8 game, bool won, address participant, uint256 value)
func (_Moonflip *MoonflipFilterer) ParseOutcome(log types.Log) (*MoonflipOutcome, error) {
	event := new(MoonflipOutcome)
	if err := _Moonflip.contract.UnpackLog(event, "Outcome", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MoonflipOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Moonflip contract.
type MoonflipOwnershipTransferredIterator struct {
	Event *MoonflipOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MoonflipOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MoonflipOwnershipTransferred)
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
		it.Event = new(MoonflipOwnershipTransferred)
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
func (it *MoonflipOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MoonflipOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MoonflipOwnershipTransferred represents a OwnershipTransferred event raised by the Moonflip contract.
type MoonflipOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Moonflip *MoonflipFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MoonflipOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Moonflip.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MoonflipOwnershipTransferredIterator{contract: _Moonflip.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Moonflip *MoonflipFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MoonflipOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Moonflip.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MoonflipOwnershipTransferred)
				if err := _Moonflip.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Moonflip *MoonflipFilterer) ParseOwnershipTransferred(log types.Log) (*MoonflipOwnershipTransferred, error) {
	event := new(MoonflipOwnershipTransferred)
	if err := _Moonflip.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
