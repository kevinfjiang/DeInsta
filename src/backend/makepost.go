// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package backend

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

// UserMetaData contains all meta data concerning the User contract.
var UserMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"postString\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"CountPosts\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"CommentString\",\"type\":\"string\"}],\"name\":\"AddComment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Balance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"Comments\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"CommentNumber\",\"type\":\"uint256\"},{\"internalType\":\"int16\",\"name\":\"Karma\",\"type\":\"int16\"},{\"internalType\":\"string\",\"name\":\"CommentString\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"Poster\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CountComments\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DeletePost\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Karma\",\"outputs\":[{\"internalType\":\"int16\",\"name\":\"\",\"type\":\"int16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PostNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PostString\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Transfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int8\",\"name\":\"vote\",\"type\":\"int8\"}],\"name\":\"Vote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int8\",\"name\":\"vote\",\"type\":\"int8\"},{\"internalType\":\"uint256\",\"name\":\"CommentNumber\",\"type\":\"uint256\"}],\"name\":\"Vote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// UserABI is the input ABI used to generate the binding from.
// Deprecated: Use UserMetaData.ABI instead.
var UserABI = UserMetaData.ABI

// User is an auto generated Go binding around an Ethereum contract.
type User struct {
	UserCaller     // Read-only binding to the contract
	UserTransactor // Write-only binding to the contract
	UserFilterer   // Log filterer for contract events
}

// UserCaller is an auto generated read-only Go binding around an Ethereum contract.
type UserCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UserTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UserTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UserFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UserFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UserSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UserSession struct {
	Contract     *User             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UserCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UserCallerSession struct {
	Contract *UserCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// UserTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UserTransactorSession struct {
	Contract     *UserTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UserRaw is an auto generated low-level Go binding around an Ethereum contract.
type UserRaw struct {
	Contract *User // Generic contract binding to access the raw methods on
}

// UserCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UserCallerRaw struct {
	Contract *UserCaller // Generic read-only contract binding to access the raw methods on
}

// UserTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UserTransactorRaw struct {
	Contract *UserTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUser creates a new instance of User, bound to a specific deployed contract.
func NewUser(address common.Address, backend bind.ContractBackend) (*User, error) {
	contract, err := bindUser(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &User{UserCaller: UserCaller{contract: contract}, UserTransactor: UserTransactor{contract: contract}, UserFilterer: UserFilterer{contract: contract}}, nil
}

// NewUserCaller creates a new read-only instance of User, bound to a specific deployed contract.
func NewUserCaller(address common.Address, caller bind.ContractCaller) (*UserCaller, error) {
	contract, err := bindUser(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UserCaller{contract: contract}, nil
}

// NewUserTransactor creates a new write-only instance of User, bound to a specific deployed contract.
func NewUserTransactor(address common.Address, transactor bind.ContractTransactor) (*UserTransactor, error) {
	contract, err := bindUser(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UserTransactor{contract: contract}, nil
}

// NewUserFilterer creates a new log filterer instance of User, bound to a specific deployed contract.
func NewUserFilterer(address common.Address, filterer bind.ContractFilterer) (*UserFilterer, error) {
	contract, err := bindUser(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UserFilterer{contract: contract}, nil
}

// bindUser binds a generic wrapper to an already deployed contract.
func bindUser(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UserABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_User *UserRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _User.Contract.UserCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_User *UserRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _User.Contract.UserTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_User *UserRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _User.Contract.UserTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_User *UserCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _User.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_User *UserTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _User.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_User *UserTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _User.Contract.contract.Transact(opts, method, params...)
}

// Balance is a free data retrieval call binding the contract method 0x0ef67887.
//
// Solidity: function Balance() view returns(uint256)
func (_User *UserCaller) Balance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _User.contract.Call(opts, &out, "Balance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balance is a free data retrieval call binding the contract method 0x0ef67887.
//
// Solidity: function Balance() view returns(uint256)
func (_User *UserSession) Balance() (*big.Int, error) {
	return _User.Contract.Balance(&_User.CallOpts)
}

// Balance is a free data retrieval call binding the contract method 0x0ef67887.
//
// Solidity: function Balance() view returns(uint256)
func (_User *UserCallerSession) Balance() (*big.Int, error) {
	return _User.Contract.Balance(&_User.CallOpts)
}

// Comments is a free data retrieval call binding the contract method 0x8ea20ddc.
//
// Solidity: function Comments(uint256 ) view returns(uint256 timestamp, uint256 CommentNumber, int16 Karma, string CommentString, address Poster)
func (_User *UserCaller) Comments(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Timestamp     *big.Int
	CommentNumber *big.Int
	Karma         int16
	CommentString string
	Poster        common.Address
}, error) {
	var out []interface{}
	err := _User.contract.Call(opts, &out, "Comments", arg0)

	outstruct := new(struct {
		Timestamp     *big.Int
		CommentNumber *big.Int
		Karma         int16
		CommentString string
		Poster        common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Timestamp = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.CommentNumber = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Karma = *abi.ConvertType(out[2], new(int16)).(*int16)
	outstruct.CommentString = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.Poster = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// Comments is a free data retrieval call binding the contract method 0x8ea20ddc.
//
// Solidity: function Comments(uint256 ) view returns(uint256 timestamp, uint256 CommentNumber, int16 Karma, string CommentString, address Poster)
func (_User *UserSession) Comments(arg0 *big.Int) (struct {
	Timestamp     *big.Int
	CommentNumber *big.Int
	Karma         int16
	CommentString string
	Poster        common.Address
}, error) {
	return _User.Contract.Comments(&_User.CallOpts, arg0)
}

// Comments is a free data retrieval call binding the contract method 0x8ea20ddc.
//
// Solidity: function Comments(uint256 ) view returns(uint256 timestamp, uint256 CommentNumber, int16 Karma, string CommentString, address Poster)
func (_User *UserCallerSession) Comments(arg0 *big.Int) (struct {
	Timestamp     *big.Int
	CommentNumber *big.Int
	Karma         int16
	CommentString string
	Poster        common.Address
}, error) {
	return _User.Contract.Comments(&_User.CallOpts, arg0)
}

// CountComments is a free data retrieval call binding the contract method 0xdc28c100.
//
// Solidity: function CountComments() view returns(uint256)
func (_User *UserCaller) CountComments(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _User.contract.Call(opts, &out, "CountComments")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CountComments is a free data retrieval call binding the contract method 0xdc28c100.
//
// Solidity: function CountComments() view returns(uint256)
func (_User *UserSession) CountComments() (*big.Int, error) {
	return _User.Contract.CountComments(&_User.CallOpts)
}

// CountComments is a free data retrieval call binding the contract method 0xdc28c100.
//
// Solidity: function CountComments() view returns(uint256)
func (_User *UserCallerSession) CountComments() (*big.Int, error) {
	return _User.Contract.CountComments(&_User.CallOpts)
}

// Karma is a free data retrieval call binding the contract method 0xe5163de5.
//
// Solidity: function Karma() view returns(int16)
func (_User *UserCaller) Karma(opts *bind.CallOpts) (int16, error) {
	var out []interface{}
	err := _User.contract.Call(opts, &out, "Karma")

	if err != nil {
		return *new(int16), err
	}

	out0 := *abi.ConvertType(out[0], new(int16)).(*int16)

	return out0, err

}

// Karma is a free data retrieval call binding the contract method 0xe5163de5.
//
// Solidity: function Karma() view returns(int16)
func (_User *UserSession) Karma() (int16, error) {
	return _User.Contract.Karma(&_User.CallOpts)
}

// Karma is a free data retrieval call binding the contract method 0xe5163de5.
//
// Solidity: function Karma() view returns(int16)
func (_User *UserCallerSession) Karma() (int16, error) {
	return _User.Contract.Karma(&_User.CallOpts)
}

// PostNumber is a free data retrieval call binding the contract method 0x98c024bc.
//
// Solidity: function PostNumber() view returns(uint256)
func (_User *UserCaller) PostNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _User.contract.Call(opts, &out, "PostNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PostNumber is a free data retrieval call binding the contract method 0x98c024bc.
//
// Solidity: function PostNumber() view returns(uint256)
func (_User *UserSession) PostNumber() (*big.Int, error) {
	return _User.Contract.PostNumber(&_User.CallOpts)
}

// PostNumber is a free data retrieval call binding the contract method 0x98c024bc.
//
// Solidity: function PostNumber() view returns(uint256)
func (_User *UserCallerSession) PostNumber() (*big.Int, error) {
	return _User.Contract.PostNumber(&_User.CallOpts)
}

// PostString is a free data retrieval call binding the contract method 0xfe6c9752.
//
// Solidity: function PostString() view returns(string)
func (_User *UserCaller) PostString(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _User.contract.Call(opts, &out, "PostString")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// PostString is a free data retrieval call binding the contract method 0xfe6c9752.
//
// Solidity: function PostString() view returns(string)
func (_User *UserSession) PostString() (string, error) {
	return _User.Contract.PostString(&_User.CallOpts)
}

// PostString is a free data retrieval call binding the contract method 0xfe6c9752.
//
// Solidity: function PostString() view returns(string)
func (_User *UserCallerSession) PostString() (string, error) {
	return _User.Contract.PostString(&_User.CallOpts)
}

// AddComment is a paid mutator transaction binding the contract method 0x6bde24b6.
//
// Solidity: function AddComment(string CommentString) returns()
func (_User *UserTransactor) AddComment(opts *bind.TransactOpts, CommentString string) (*types.Transaction, error) {
	return _User.contract.Transact(opts, "AddComment", CommentString)
}

// AddComment is a paid mutator transaction binding the contract method 0x6bde24b6.
//
// Solidity: function AddComment(string CommentString) returns()
func (_User *UserSession) AddComment(CommentString string) (*types.Transaction, error) {
	return _User.Contract.AddComment(&_User.TransactOpts, CommentString)
}

// AddComment is a paid mutator transaction binding the contract method 0x6bde24b6.
//
// Solidity: function AddComment(string CommentString) returns()
func (_User *UserTransactorSession) AddComment(CommentString string) (*types.Transaction, error) {
	return _User.Contract.AddComment(&_User.TransactOpts, CommentString)
}

// DeletePost is a paid mutator transaction binding the contract method 0x85141ee3.
//
// Solidity: function DeletePost() returns()
func (_User *UserTransactor) DeletePost(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _User.contract.Transact(opts, "DeletePost")
}

// DeletePost is a paid mutator transaction binding the contract method 0x85141ee3.
//
// Solidity: function DeletePost() returns()
func (_User *UserSession) DeletePost() (*types.Transaction, error) {
	return _User.Contract.DeletePost(&_User.TransactOpts)
}

// DeletePost is a paid mutator transaction binding the contract method 0x85141ee3.
//
// Solidity: function DeletePost() returns()
func (_User *UserTransactorSession) DeletePost() (*types.Transaction, error) {
	return _User.Contract.DeletePost(&_User.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0x406dade3.
//
// Solidity: function Transfer() returns()
func (_User *UserTransactor) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _User.contract.Transact(opts, "Transfer")
}

// Transfer is a paid mutator transaction binding the contract method 0x406dade3.
//
// Solidity: function Transfer() returns()
func (_User *UserSession) Transfer() (*types.Transaction, error) {
	return _User.Contract.Transfer(&_User.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0x406dade3.
//
// Solidity: function Transfer() returns()
func (_User *UserTransactorSession) Transfer() (*types.Transaction, error) {
	return _User.Contract.Transfer(&_User.TransactOpts)
}

// Vote is a paid mutator transaction binding the contract method 0x4e49d2da.
//
// Solidity: function Vote(int8 vote) returns()
func (_User *UserTransactor) Vote(opts *bind.TransactOpts, vote int8) (*types.Transaction, error) {
	return _User.contract.Transact(opts, "Vote", vote)
}

// Vote is a paid mutator transaction binding the contract method 0x4e49d2da.
//
// Solidity: function Vote(int8 vote) returns()
func (_User *UserSession) Vote(vote int8) (*types.Transaction, error) {
	return _User.Contract.Vote(&_User.TransactOpts, vote)
}

// Vote is a paid mutator transaction binding the contract method 0x4e49d2da.
//
// Solidity: function Vote(int8 vote) returns()
func (_User *UserTransactorSession) Vote(vote int8) (*types.Transaction, error) {
	return _User.Contract.Vote(&_User.TransactOpts, vote)
}

// Vote0 is a paid mutator transaction binding the contract method 0xae81085a.
//
// Solidity: function Vote(int8 vote, uint256 CommentNumber) returns()
func (_User *UserTransactor) Vote0(opts *bind.TransactOpts, vote int8, CommentNumber *big.Int) (*types.Transaction, error) {
	return _User.contract.Transact(opts, "Vote0", vote, CommentNumber)
}

// Vote0 is a paid mutator transaction binding the contract method 0xae81085a.
//
// Solidity: function Vote(int8 vote, uint256 CommentNumber) returns()
func (_User *UserSession) Vote0(vote int8, CommentNumber *big.Int) (*types.Transaction, error) {
	return _User.Contract.Vote0(&_User.TransactOpts, vote, CommentNumber)
}

// Vote0 is a paid mutator transaction binding the contract method 0xae81085a.
//
// Solidity: function Vote(int8 vote, uint256 CommentNumber) returns()
func (_User *UserTransactorSession) Vote0(vote int8, CommentNumber *big.Int) (*types.Transaction, error) {
	return _User.Contract.Vote0(&_User.TransactOpts, vote, CommentNumber)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_User *UserTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _User.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_User *UserSession) Receive() (*types.Transaction, error) {
	return _User.Contract.Receive(&_User.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_User *UserTransactorSession) Receive() (*types.Transaction, error) {
	return _User.Contract.Receive(&_User.TransactOpts)
}
