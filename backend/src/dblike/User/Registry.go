// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package User

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

// DIregistryMetaData contains all meta data concerning the DIregistry contract.
var DIregistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"ServerName\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"accountNumber\",\"type\":\"uint256\"}],\"name\":\"NewUser\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"Name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_registrationDisabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"adminRetrieveDonations\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"accountAdmin\",\"type\":\"address\"}],\"name\":\"adminSetAccountAdministrator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"registrationDisabled\",\"type\":\"bool\"}],\"name\":\"adminSetRegistrationDisabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"adminUnregister\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getAddressOfId\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"getAddressOfName\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getNameOfAddress\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numberOfAccounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"accountAddress\",\"type\":\"address\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unregister\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"8052474d": "Name()",
		"4a51a936": "_registrationDisabled()",
		"345e3416": "adminRetrieveDonations()",
		"49f0c90d": "adminSetAccountAdministrator(address)",
		"9b6d86d6": "adminSetRegistrationDisabled(bool)",
		"0cc04b55": "adminUnregister(string)",
		"ec43eeb6": "getAddressOfId(uint256)",
		"38ec6ba8": "getAddressOfName(string)",
		"b2dad25d": "getNameOfAddress(address)",
		"0f03e4c3": "numberOfAccounts()",
		"1e59c529": "register(string,address)",
		"e79a198f": "unregister()",
	},
	Bin: "0x60806040523480156200001157600080fd5b5060405162000e0638038062000e0683398101604081905262000034916200013a565b60148151106200004357600080fd5b8051620000589060009060208401906200007e565b50506005805460006004556001600160a81b0319163360ff60a01b191617905562000253565b8280546200008c9062000216565b90600052602060002090601f016020900481019282620000b05760008555620000fb565b82601f10620000cb57805160ff1916838001178555620000fb565b82800160010185558215620000fb579182015b82811115620000fb578251825591602001919060010190620000de565b50620001099291506200010d565b5090565b5b808211156200010957600081556001016200010e565b634e487b7160e01b600052604160045260246000fd5b600060208083850312156200014e57600080fd5b82516001600160401b03808211156200016657600080fd5b818501915085601f8301126200017b57600080fd5b81518181111562000190576200019062000124565b604051601f8201601f19908116603f01168101908382118183101715620001bb57620001bb62000124565b816040528281528886848701011115620001d457600080fd5b600093505b82841015620001f85784840186015181850187015292850192620001d9565b828411156200020a5760008684830101525b98975050505050505050565b600181811c908216806200022b57607f821691505b602082108114156200024d57634e487b7160e01b600052602260045260246000fd5b50919050565b610ba380620002636000396000f3fe608060405234801561001057600080fd5b50600436106100b45760003560e01c80634a51a936116100715780634a51a936146101435780638052474d146101675780639b6d86d61461017c578063b2dad25d1461018f578063e79a198f146101a2578063ec43eeb6146101aa57600080fd5b80630cc04b55146100b95780630f03e4c3146100ce5780631e59c529146100ea578063345e3416146100fd57806338ec6ba81461010557806349f0c90d14610130575b600080fd5b6100cc6100c7366004610986565b6101bd565b005b6100d760045481565b6040519081526020015b60405180910390f35b6100cc6100f83660046109df565b610270565b6100cc6104dc565b610118610113366004610986565b61052f565b6040516001600160a01b0390911681526020016100e1565b6100cc61013e366004610a2d565b610560565b60055461015790600160a01b900460ff1681565b60405190151581526020016100e1565b61016f610599565b6040516100e19190610a7f565b6100cc61018a366004610ab2565b610627565b61016f61019d366004610a2d565b61065c565b6100cc610708565b6101186101b8366004610ad4565b610814565b6005546001600160a01b031633146101d457600080fd5b60006003826040516101e69190610aed565b9081526040805191829003602090810183205483820180845260008086526001600160a01b039092168083526001909352929020925190935061022a92919061084a565b50600060038360405161023d9190610aed565b90815260405190819003602001902080546001600160a01b03929092166001600160a01b03199092169190911790555050565b60006001600160a01b031660038360405161028b9190610aed565b908152604051908190036020019020546001600160a01b031614156102e45760405162461bcd60e51b815260206004820152600a6024820152692730b6b2902a30b5b2b760b11b60448201526064015b60405180910390fd5b6001600160a01b0381166000908152600160205260409020805461030790610b09565b151590506103575760405162461bcd60e51b815260206004820152601a60248201527f4164647265737320616c7265616479207265676973746572656400000000000060448201526064016102db565b6040825110156103995760405162461bcd60e51b815260206004820152600d60248201526c4e616d6520746f6f206c6f6e6760981b60448201526064016102db565b600554600160a01b900460ff166103f25760405162461bcd60e51b815260206004820181905260248201527f52656769737472792064697361626c656420666f72206e65772075706461746560448201526064016102db565b6001600160a01b0381166000908152600160209081526040909120835161041b9285019061084a565b508060038360405161042d9190610aed565b908152604080516020928190038301902080546001600160a01b03199081166001600160a01b039586161790915560048054600090815260029094529183208054909116938516939093179092558154919061048883610b44565b9091555050600454604080514281526001600160a01b038416602082015280820192909252517f46385ef64f03cddac3514117a2aac0e2528b47d904c8ad4417ba7cc0e76746349181900360600190a15050565b6005546001600160a01b031633146104f357600080fd5b6005546040516001600160a01b03909116904780156108fc02916000818181858888f1935050505015801561052c573d6000803e3d6000fd5b50565b60006003826040516105419190610aed565b908152604051908190036020019020546001600160a01b031692915050565b6005546001600160a01b0316331461057757600080fd5b600580546001600160a01b0319166001600160a01b0392909216919091179055565b600080546105a690610b09565b80601f01602080910402602001604051908101604052809291908181526020018280546105d290610b09565b801561061f5780601f106105f45761010080835404028352916020019161061f565b820191906000526020600020905b81548152906001019060200180831161060257829003601f168201915b505050505081565b6005546001600160a01b0316331461063e57600080fd5b60058054911515600160a01b0260ff60a01b19909216919091179055565b6001600160a01b038116600090815260016020526040902080546060919061068390610b09565b80601f01602080910402602001604051908101604052809291908181526020018280546106af90610b09565b80156106fc5780601f106106d1576101008083540402835291602001916106fc565b820191906000526020600020905b8154815290600101906020018083116106df57829003601f168201915b50505050509050919050565b336000908152600160205260408120805461072290610b09565b80601f016020809104026020016040519081016040528092919081815260200182805461074e90610b09565b801561079b5780601f106107705761010080835404028352916020019161079b565b820191906000526020600020905b81548152906001019060200180831161077e57829003601f168201915b50506040805160208082018084526000808452338152600190925292902090519596506107cf95909450909250905061084a565b5060006003826040516107e29190610aed565b90815260405190819003602001902080546001600160a01b03929092166001600160a01b031990921691909117905550565b6005546000906001600160a01b0316331461082e57600080fd5b506000908152600260205260409020546001600160a01b031690565b82805461085690610b09565b90600052602060002090601f01602090048101928261087857600085556108be565b82601f1061089157805160ff19168380011785556108be565b828001600101855582156108be579182015b828111156108be5782518255916020019190600101906108a3565b506108ca9291506108ce565b5090565b5b808211156108ca57600081556001016108cf565b634e487b7160e01b600052604160045260246000fd5b600082601f83011261090a57600080fd5b813567ffffffffffffffff80821115610925576109256108e3565b604051601f8301601f19908116603f0116810190828211818310171561094d5761094d6108e3565b8160405283815286602085880101111561096657600080fd5b836020870160208301376000602085830101528094505050505092915050565b60006020828403121561099857600080fd5b813567ffffffffffffffff8111156109af57600080fd5b6109bb848285016108f9565b949350505050565b80356001600160a01b03811681146109da57600080fd5b919050565b600080604083850312156109f257600080fd5b823567ffffffffffffffff811115610a0957600080fd5b610a15858286016108f9565b925050610a24602084016109c3565b90509250929050565b600060208284031215610a3f57600080fd5b610a48826109c3565b9392505050565b60005b83811015610a6a578181015183820152602001610a52565b83811115610a79576000848401525b50505050565b6020815260008251806020840152610a9e816040850160208701610a4f565b601f01601f19169190910160400192915050565b600060208284031215610ac457600080fd5b81358015158114610a4857600080fd5b600060208284031215610ae657600080fd5b5035919050565b60008251610aff818460208701610a4f565b9190910192915050565b600181811c90821680610b1d57607f821691505b60208210811415610b3e57634e487b7160e01b600052602260045260246000fd5b50919050565b6000600019821415610b6657634e487b7160e01b600052601160045260246000fd5b506001019056fea26469706673582212203c9fca9fbd7e941176caa0a01c3306da5ec28773eed561e75c096c7c65a53ecf64736f6c63430008090033",
}

// DIregistryABI is the input ABI used to generate the binding from.
// Deprecated: Use DIregistryMetaData.ABI instead.
var DIregistryABI = DIregistryMetaData.ABI

// Deprecated: Use DIregistryMetaData.Sigs instead.
// DIregistryFuncSigs maps the 4-byte function signature to its string representation.
var DIregistryFuncSigs = DIregistryMetaData.Sigs

// DIregistryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DIregistryMetaData.Bin instead.
var DIregistryBin = DIregistryMetaData.Bin

// DeployDIregistry deploys a new Ethereum contract, binding an instance of DIregistry to it.
func DeployDIregistry(auth *bind.TransactOpts, backend bind.ContractBackend, ServerName string) (common.Address, *types.Transaction, *DIregistry, error) {
	parsed, err := DIregistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DIregistryBin), backend, ServerName)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DIregistry{DIregistryCaller: DIregistryCaller{contract: contract}, DIregistryTransactor: DIregistryTransactor{contract: contract}, DIregistryFilterer: DIregistryFilterer{contract: contract}}, nil
}

// DIregistry is an auto generated Go binding around an Ethereum contract.
type DIregistry struct {
	DIregistryCaller     // Read-only binding to the contract
	DIregistryTransactor // Write-only binding to the contract
	DIregistryFilterer   // Log filterer for contract events
}

// DIregistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type DIregistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DIregistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DIregistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DIregistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DIregistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DIregistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DIregistrySession struct {
	Contract     *DIregistry       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DIregistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DIregistryCallerSession struct {
	Contract *DIregistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// DIregistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DIregistryTransactorSession struct {
	Contract     *DIregistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// DIregistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type DIregistryRaw struct {
	Contract *DIregistry // Generic contract binding to access the raw methods on
}

// DIregistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DIregistryCallerRaw struct {
	Contract *DIregistryCaller // Generic read-only contract binding to access the raw methods on
}

// DIregistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DIregistryTransactorRaw struct {
	Contract *DIregistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDIregistry creates a new instance of DIregistry, bound to a specific deployed contract.
func NewDIregistry(address common.Address, backend bind.ContractBackend) (*DIregistry, error) {
	contract, err := bindDIregistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DIregistry{DIregistryCaller: DIregistryCaller{contract: contract}, DIregistryTransactor: DIregistryTransactor{contract: contract}, DIregistryFilterer: DIregistryFilterer{contract: contract}}, nil
}

// NewDIregistryCaller creates a new read-only instance of DIregistry, bound to a specific deployed contract.
func NewDIregistryCaller(address common.Address, caller bind.ContractCaller) (*DIregistryCaller, error) {
	contract, err := bindDIregistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DIregistryCaller{contract: contract}, nil
}

// NewDIregistryTransactor creates a new write-only instance of DIregistry, bound to a specific deployed contract.
func NewDIregistryTransactor(address common.Address, transactor bind.ContractTransactor) (*DIregistryTransactor, error) {
	contract, err := bindDIregistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DIregistryTransactor{contract: contract}, nil
}

// NewDIregistryFilterer creates a new log filterer instance of DIregistry, bound to a specific deployed contract.
func NewDIregistryFilterer(address common.Address, filterer bind.ContractFilterer) (*DIregistryFilterer, error) {
	contract, err := bindDIregistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DIregistryFilterer{contract: contract}, nil
}

// bindDIregistry binds a generic wrapper to an already deployed contract.
func bindDIregistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DIregistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DIregistry *DIregistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DIregistry.Contract.DIregistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DIregistry *DIregistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DIregistry.Contract.DIregistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DIregistry *DIregistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DIregistry.Contract.DIregistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DIregistry *DIregistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DIregistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DIregistry *DIregistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DIregistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DIregistry *DIregistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DIregistry.Contract.contract.Transact(opts, method, params...)
}

// Name is a free data retrieval call binding the contract method 0x8052474d.
//
// Solidity: function Name() view returns(string)
func (_DIregistry *DIregistryCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _DIregistry.contract.Call(opts, &out, "Name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x8052474d.
//
// Solidity: function Name() view returns(string)
func (_DIregistry *DIregistrySession) Name() (string, error) {
	return _DIregistry.Contract.Name(&_DIregistry.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x8052474d.
//
// Solidity: function Name() view returns(string)
func (_DIregistry *DIregistryCallerSession) Name() (string, error) {
	return _DIregistry.Contract.Name(&_DIregistry.CallOpts)
}

// RegistrationDisabled is a free data retrieval call binding the contract method 0x4a51a936.
//
// Solidity: function _registrationDisabled() view returns(bool)
func (_DIregistry *DIregistryCaller) RegistrationDisabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _DIregistry.contract.Call(opts, &out, "_registrationDisabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RegistrationDisabled is a free data retrieval call binding the contract method 0x4a51a936.
//
// Solidity: function _registrationDisabled() view returns(bool)
func (_DIregistry *DIregistrySession) RegistrationDisabled() (bool, error) {
	return _DIregistry.Contract.RegistrationDisabled(&_DIregistry.CallOpts)
}

// RegistrationDisabled is a free data retrieval call binding the contract method 0x4a51a936.
//
// Solidity: function _registrationDisabled() view returns(bool)
func (_DIregistry *DIregistryCallerSession) RegistrationDisabled() (bool, error) {
	return _DIregistry.Contract.RegistrationDisabled(&_DIregistry.CallOpts)
}

// GetAddressOfId is a free data retrieval call binding the contract method 0xec43eeb6.
//
// Solidity: function getAddressOfId(uint256 id) view returns(address addr)
func (_DIregistry *DIregistryCaller) GetAddressOfId(opts *bind.CallOpts, id *big.Int) (common.Address, error) {
	var out []interface{}
	err := _DIregistry.contract.Call(opts, &out, "getAddressOfId", id)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddressOfId is a free data retrieval call binding the contract method 0xec43eeb6.
//
// Solidity: function getAddressOfId(uint256 id) view returns(address addr)
func (_DIregistry *DIregistrySession) GetAddressOfId(id *big.Int) (common.Address, error) {
	return _DIregistry.Contract.GetAddressOfId(&_DIregistry.CallOpts, id)
}

// GetAddressOfId is a free data retrieval call binding the contract method 0xec43eeb6.
//
// Solidity: function getAddressOfId(uint256 id) view returns(address addr)
func (_DIregistry *DIregistryCallerSession) GetAddressOfId(id *big.Int) (common.Address, error) {
	return _DIregistry.Contract.GetAddressOfId(&_DIregistry.CallOpts, id)
}

// GetAddressOfName is a free data retrieval call binding the contract method 0x38ec6ba8.
//
// Solidity: function getAddressOfName(string name) view returns(address addr)
func (_DIregistry *DIregistryCaller) GetAddressOfName(opts *bind.CallOpts, name string) (common.Address, error) {
	var out []interface{}
	err := _DIregistry.contract.Call(opts, &out, "getAddressOfName", name)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddressOfName is a free data retrieval call binding the contract method 0x38ec6ba8.
//
// Solidity: function getAddressOfName(string name) view returns(address addr)
func (_DIregistry *DIregistrySession) GetAddressOfName(name string) (common.Address, error) {
	return _DIregistry.Contract.GetAddressOfName(&_DIregistry.CallOpts, name)
}

// GetAddressOfName is a free data retrieval call binding the contract method 0x38ec6ba8.
//
// Solidity: function getAddressOfName(string name) view returns(address addr)
func (_DIregistry *DIregistryCallerSession) GetAddressOfName(name string) (common.Address, error) {
	return _DIregistry.Contract.GetAddressOfName(&_DIregistry.CallOpts, name)
}

// GetNameOfAddress is a free data retrieval call binding the contract method 0xb2dad25d.
//
// Solidity: function getNameOfAddress(address addr) view returns(string name)
func (_DIregistry *DIregistryCaller) GetNameOfAddress(opts *bind.CallOpts, addr common.Address) (string, error) {
	var out []interface{}
	err := _DIregistry.contract.Call(opts, &out, "getNameOfAddress", addr)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetNameOfAddress is a free data retrieval call binding the contract method 0xb2dad25d.
//
// Solidity: function getNameOfAddress(address addr) view returns(string name)
func (_DIregistry *DIregistrySession) GetNameOfAddress(addr common.Address) (string, error) {
	return _DIregistry.Contract.GetNameOfAddress(&_DIregistry.CallOpts, addr)
}

// GetNameOfAddress is a free data retrieval call binding the contract method 0xb2dad25d.
//
// Solidity: function getNameOfAddress(address addr) view returns(string name)
func (_DIregistry *DIregistryCallerSession) GetNameOfAddress(addr common.Address) (string, error) {
	return _DIregistry.Contract.GetNameOfAddress(&_DIregistry.CallOpts, addr)
}

// NumberOfAccounts is a free data retrieval call binding the contract method 0x0f03e4c3.
//
// Solidity: function numberOfAccounts() view returns(uint256)
func (_DIregistry *DIregistryCaller) NumberOfAccounts(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DIregistry.contract.Call(opts, &out, "numberOfAccounts")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumberOfAccounts is a free data retrieval call binding the contract method 0x0f03e4c3.
//
// Solidity: function numberOfAccounts() view returns(uint256)
func (_DIregistry *DIregistrySession) NumberOfAccounts() (*big.Int, error) {
	return _DIregistry.Contract.NumberOfAccounts(&_DIregistry.CallOpts)
}

// NumberOfAccounts is a free data retrieval call binding the contract method 0x0f03e4c3.
//
// Solidity: function numberOfAccounts() view returns(uint256)
func (_DIregistry *DIregistryCallerSession) NumberOfAccounts() (*big.Int, error) {
	return _DIregistry.Contract.NumberOfAccounts(&_DIregistry.CallOpts)
}

// AdminRetrieveDonations is a paid mutator transaction binding the contract method 0x345e3416.
//
// Solidity: function adminRetrieveDonations() returns()
func (_DIregistry *DIregistryTransactor) AdminRetrieveDonations(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DIregistry.contract.Transact(opts, "adminRetrieveDonations")
}

// AdminRetrieveDonations is a paid mutator transaction binding the contract method 0x345e3416.
//
// Solidity: function adminRetrieveDonations() returns()
func (_DIregistry *DIregistrySession) AdminRetrieveDonations() (*types.Transaction, error) {
	return _DIregistry.Contract.AdminRetrieveDonations(&_DIregistry.TransactOpts)
}

// AdminRetrieveDonations is a paid mutator transaction binding the contract method 0x345e3416.
//
// Solidity: function adminRetrieveDonations() returns()
func (_DIregistry *DIregistryTransactorSession) AdminRetrieveDonations() (*types.Transaction, error) {
	return _DIregistry.Contract.AdminRetrieveDonations(&_DIregistry.TransactOpts)
}

// AdminSetAccountAdministrator is a paid mutator transaction binding the contract method 0x49f0c90d.
//
// Solidity: function adminSetAccountAdministrator(address accountAdmin) returns()
func (_DIregistry *DIregistryTransactor) AdminSetAccountAdministrator(opts *bind.TransactOpts, accountAdmin common.Address) (*types.Transaction, error) {
	return _DIregistry.contract.Transact(opts, "adminSetAccountAdministrator", accountAdmin)
}

// AdminSetAccountAdministrator is a paid mutator transaction binding the contract method 0x49f0c90d.
//
// Solidity: function adminSetAccountAdministrator(address accountAdmin) returns()
func (_DIregistry *DIregistrySession) AdminSetAccountAdministrator(accountAdmin common.Address) (*types.Transaction, error) {
	return _DIregistry.Contract.AdminSetAccountAdministrator(&_DIregistry.TransactOpts, accountAdmin)
}

// AdminSetAccountAdministrator is a paid mutator transaction binding the contract method 0x49f0c90d.
//
// Solidity: function adminSetAccountAdministrator(address accountAdmin) returns()
func (_DIregistry *DIregistryTransactorSession) AdminSetAccountAdministrator(accountAdmin common.Address) (*types.Transaction, error) {
	return _DIregistry.Contract.AdminSetAccountAdministrator(&_DIregistry.TransactOpts, accountAdmin)
}

// AdminSetRegistrationDisabled is a paid mutator transaction binding the contract method 0x9b6d86d6.
//
// Solidity: function adminSetRegistrationDisabled(bool registrationDisabled) returns()
func (_DIregistry *DIregistryTransactor) AdminSetRegistrationDisabled(opts *bind.TransactOpts, registrationDisabled bool) (*types.Transaction, error) {
	return _DIregistry.contract.Transact(opts, "adminSetRegistrationDisabled", registrationDisabled)
}

// AdminSetRegistrationDisabled is a paid mutator transaction binding the contract method 0x9b6d86d6.
//
// Solidity: function adminSetRegistrationDisabled(bool registrationDisabled) returns()
func (_DIregistry *DIregistrySession) AdminSetRegistrationDisabled(registrationDisabled bool) (*types.Transaction, error) {
	return _DIregistry.Contract.AdminSetRegistrationDisabled(&_DIregistry.TransactOpts, registrationDisabled)
}

// AdminSetRegistrationDisabled is a paid mutator transaction binding the contract method 0x9b6d86d6.
//
// Solidity: function adminSetRegistrationDisabled(bool registrationDisabled) returns()
func (_DIregistry *DIregistryTransactorSession) AdminSetRegistrationDisabled(registrationDisabled bool) (*types.Transaction, error) {
	return _DIregistry.Contract.AdminSetRegistrationDisabled(&_DIregistry.TransactOpts, registrationDisabled)
}

// AdminUnregister is a paid mutator transaction binding the contract method 0x0cc04b55.
//
// Solidity: function adminUnregister(string name) returns()
func (_DIregistry *DIregistryTransactor) AdminUnregister(opts *bind.TransactOpts, name string) (*types.Transaction, error) {
	return _DIregistry.contract.Transact(opts, "adminUnregister", name)
}

// AdminUnregister is a paid mutator transaction binding the contract method 0x0cc04b55.
//
// Solidity: function adminUnregister(string name) returns()
func (_DIregistry *DIregistrySession) AdminUnregister(name string) (*types.Transaction, error) {
	return _DIregistry.Contract.AdminUnregister(&_DIregistry.TransactOpts, name)
}

// AdminUnregister is a paid mutator transaction binding the contract method 0x0cc04b55.
//
// Solidity: function adminUnregister(string name) returns()
func (_DIregistry *DIregistryTransactorSession) AdminUnregister(name string) (*types.Transaction, error) {
	return _DIregistry.Contract.AdminUnregister(&_DIregistry.TransactOpts, name)
}

// Register is a paid mutator transaction binding the contract method 0x1e59c529.
//
// Solidity: function register(string name, address accountAddress) returns()
func (_DIregistry *DIregistryTransactor) Register(opts *bind.TransactOpts, name string, accountAddress common.Address) (*types.Transaction, error) {
	return _DIregistry.contract.Transact(opts, "register", name, accountAddress)
}

// Register is a paid mutator transaction binding the contract method 0x1e59c529.
//
// Solidity: function register(string name, address accountAddress) returns()
func (_DIregistry *DIregistrySession) Register(name string, accountAddress common.Address) (*types.Transaction, error) {
	return _DIregistry.Contract.Register(&_DIregistry.TransactOpts, name, accountAddress)
}

// Register is a paid mutator transaction binding the contract method 0x1e59c529.
//
// Solidity: function register(string name, address accountAddress) returns()
func (_DIregistry *DIregistryTransactorSession) Register(name string, accountAddress common.Address) (*types.Transaction, error) {
	return _DIregistry.Contract.Register(&_DIregistry.TransactOpts, name, accountAddress)
}

// Unregister is a paid mutator transaction binding the contract method 0xe79a198f.
//
// Solidity: function unregister() returns()
func (_DIregistry *DIregistryTransactor) Unregister(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DIregistry.contract.Transact(opts, "unregister")
}

// Unregister is a paid mutator transaction binding the contract method 0xe79a198f.
//
// Solidity: function unregister() returns()
func (_DIregistry *DIregistrySession) Unregister() (*types.Transaction, error) {
	return _DIregistry.Contract.Unregister(&_DIregistry.TransactOpts)
}

// Unregister is a paid mutator transaction binding the contract method 0xe79a198f.
//
// Solidity: function unregister() returns()
func (_DIregistry *DIregistryTransactorSession) Unregister() (*types.Transaction, error) {
	return _DIregistry.Contract.Unregister(&_DIregistry.TransactOpts)
}

// DIregistryNewUserIterator is returned from FilterNewUser and is used to iterate over the raw logs and unpacked data for NewUser events raised by the DIregistry contract.
type DIregistryNewUserIterator struct {
	Event *DIregistryNewUser // Event containing the contract specifics and raw log

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
func (it *DIregistryNewUserIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DIregistryNewUser)
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
		it.Event = new(DIregistryNewUser)
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
func (it *DIregistryNewUserIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DIregistryNewUserIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DIregistryNewUser represents a NewUser event raised by the DIregistry contract.
type DIregistryNewUser struct {
	Timestamp     *big.Int
	Addr          common.Address
	AccountNumber *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterNewUser is a free log retrieval operation binding the contract event 0x46385ef64f03cddac3514117a2aac0e2528b47d904c8ad4417ba7cc0e7674634.
//
// Solidity: event NewUser(uint256 timestamp, address addr, uint256 accountNumber)
func (_DIregistry *DIregistryFilterer) FilterNewUser(opts *bind.FilterOpts) (*DIregistryNewUserIterator, error) {

	logs, sub, err := _DIregistry.contract.FilterLogs(opts, "NewUser")
	if err != nil {
		return nil, err
	}
	return &DIregistryNewUserIterator{contract: _DIregistry.contract, event: "NewUser", logs: logs, sub: sub}, nil
}

// WatchNewUser is a free log subscription operation binding the contract event 0x46385ef64f03cddac3514117a2aac0e2528b47d904c8ad4417ba7cc0e7674634.
//
// Solidity: event NewUser(uint256 timestamp, address addr, uint256 accountNumber)
func (_DIregistry *DIregistryFilterer) WatchNewUser(opts *bind.WatchOpts, sink chan<- *DIregistryNewUser) (event.Subscription, error) {

	logs, sub, err := _DIregistry.contract.WatchLogs(opts, "NewUser")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DIregistryNewUser)
				if err := _DIregistry.contract.UnpackLog(event, "NewUser", log); err != nil {
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

// ParseNewUser is a log parse operation binding the contract event 0x46385ef64f03cddac3514117a2aac0e2528b47d904c8ad4417ba7cc0e7674634.
//
// Solidity: event NewUser(uint256 timestamp, address addr, uint256 accountNumber)
func (_DIregistry *DIregistryFilterer) ParseNewUser(log types.Log) (*DIregistryNewUser, error) {
	event := new(DIregistryNewUser)
	if err := _DIregistry.contract.UnpackLog(event, "NewUser", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
