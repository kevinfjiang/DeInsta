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

// ERRegistryMetaData contains all meta data concerning the ERRegistry contract.
var ERRegistryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"SubName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"SubDescription\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"Description\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_registrationDisabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"adminRetrieveDonations\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"accountAdmin\",\"type\":\"address\"}],\"name\":\"adminSetAccountAdministrator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"registrationDisabled\",\"type\":\"bool\"}],\"name\":\"adminSetRegistrationDisabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"adminUnregister\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getAddressOfId\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"getAddressOfName\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getNameOfAddress\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numberOfAccounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"accountAddress\",\"type\":\"address\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unregister\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"8191745f": "Description()",
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
	Bin: "0x60806040523480156200001157600080fd5b5060405162000e4738038062000e47833981016040819052620000349162000208565b60148251106200004357600080fd5b81516200005890600090602085019062000095565b5080516200006e90600190602084019062000095565b50506006805460006005556001600160a81b0319163360ff60a01b191617905550620002af565b828054620000a39062000272565b90600052602060002090601f016020900481019282620000c7576000855562000112565b82601f10620000e257805160ff191683800117855562000112565b8280016001018555821562000112579182015b8281111562000112578251825591602001919060010190620000f5565b506200012092915062000124565b5090565b5b8082111562000120576000815560010162000125565b634e487b7160e01b600052604160045260246000fd5b600082601f8301126200016357600080fd5b81516001600160401b03808211156200018057620001806200013b565b604051601f8301601f19908116603f01168101908282118183101715620001ab57620001ab6200013b565b81604052838152602092508683858801011115620001c857600080fd5b600091505b83821015620001ec5785820183015181830184015290820190620001cd565b83821115620001fe5760008385830101525b9695505050505050565b600080604083850312156200021c57600080fd5b82516001600160401b03808211156200023457600080fd5b620002428683870162000151565b935060208501519150808211156200025957600080fd5b50620002688582860162000151565b9150509250929050565b600181811c908216806200028757607f821691505b60208210811415620002a957634e487b7160e01b600052602260045260246000fd5b50919050565b610b8880620002bf6000396000f3fe608060405234801561001057600080fd5b50600436106100cf5760003560e01c80634a51a9361161008c5780639b6d86d6116100665780639b6d86d61461019f578063b2dad25d146101b2578063e79a198f146101c5578063ec43eeb6146101cd57600080fd5b80634a51a9361461015e5780638052474d146101825780638191745f1461019757600080fd5b80630cc04b55146100d45780630f03e4c3146100e95780631e59c52914610105578063345e34161461011857806338ec6ba81461012057806349f0c90d1461014b575b600080fd5b6100e76100e236600461096b565b6101e0565b005b6100f260055481565b6040519081526020015b60405180910390f35b6100e76101133660046109c4565b610293565b6100e76104b4565b61013361012e36600461096b565b610507565b6040516001600160a01b0390911681526020016100fc565b6100e7610159366004610a12565b610538565b60065461017290600160a01b900460ff1681565b60405190151581526020016100fc565b61018a610571565b6040516100fc9190610a64565b61018a6105ff565b6100e76101ad366004610a97565b61060c565b61018a6101c0366004610a12565b610641565b6100e76106ed565b6101336101db366004610ab9565b6107f9565b6006546001600160a01b031633146101f757600080fd5b60006004826040516102099190610ad2565b9081526040805191829003602090810183205483820180845260008086526001600160a01b039092168083526002909352929020925190935061024d92919061082f565b5060006004836040516102609190610ad2565b90815260405190819003602001902080546001600160a01b03929092166001600160a01b03199092169190911790555050565b60006001600160a01b03166004836040516102ae9190610ad2565b908152604051908190036020019020546001600160a01b031614156103075760405162461bcd60e51b815260206004820152600a6024820152692730b6b2902a30b5b2b760b11b60448201526064015b60405180910390fd5b6001600160a01b0381166000908152600260205260409020805461032a90610aee565b1515905061037a5760405162461bcd60e51b815260206004820152601a60248201527f4164647265737320616c7265616479207265676973746572656400000000000060448201526064016102fe565b6040825110156103bc5760405162461bcd60e51b815260206004820152600d60248201526c4e616d6520746f6f206c6f6e6760981b60448201526064016102fe565b600654600160a01b900460ff166104155760405162461bcd60e51b815260206004820181905260248201527f52656769737472792064697361626c656420666f72206e65772075706461746560448201526064016102fe565b6001600160a01b0381166000908152600260209081526040909120835161043e9285019061082f565b50806004836040516104509190610ad2565b908152604080516020928190038301902080546001600160a01b03199081166001600160a01b03958616179091556005805460009081526003909452918320805490911693851693909317909255815491906104ab83610b29565b91905055505050565b6006546001600160a01b031633146104cb57600080fd5b6006546040516001600160a01b03909116904780156108fc02916000818181858888f19350505050158015610504573d6000803e3d6000fd5b50565b60006004826040516105199190610ad2565b908152604051908190036020019020546001600160a01b031692915050565b6006546001600160a01b0316331461054f57600080fd5b600680546001600160a01b0319166001600160a01b0392909216919091179055565b6000805461057e90610aee565b80601f01602080910402602001604051908101604052809291908181526020018280546105aa90610aee565b80156105f75780601f106105cc576101008083540402835291602001916105f7565b820191906000526020600020905b8154815290600101906020018083116105da57829003601f168201915b505050505081565b6001805461057e90610aee565b6006546001600160a01b0316331461062357600080fd5b60068054911515600160a01b0260ff60a01b19909216919091179055565b6001600160a01b038116600090815260026020526040902080546060919061066890610aee565b80601f016020809104026020016040519081016040528092919081815260200182805461069490610aee565b80156106e15780601f106106b6576101008083540402835291602001916106e1565b820191906000526020600020905b8154815290600101906020018083116106c457829003601f168201915b50505050509050919050565b336000908152600260205260408120805461070790610aee565b80601f016020809104026020016040519081016040528092919081815260200182805461073390610aee565b80156107805780601f1061075557610100808354040283529160200191610780565b820191906000526020600020905b81548152906001019060200180831161076357829003601f168201915b50506040805160208082018084526000808452338152600290925292902090519596506107b495909450909250905061082f565b5060006004826040516107c79190610ad2565b90815260405190819003602001902080546001600160a01b03929092166001600160a01b031990921691909117905550565b6006546000906001600160a01b0316331461081357600080fd5b506000908152600360205260409020546001600160a01b031690565b82805461083b90610aee565b90600052602060002090601f01602090048101928261085d57600085556108a3565b82601f1061087657805160ff19168380011785556108a3565b828001600101855582156108a3579182015b828111156108a3578251825591602001919060010190610888565b506108af9291506108b3565b5090565b5b808211156108af57600081556001016108b4565b634e487b7160e01b600052604160045260246000fd5b600082601f8301126108ef57600080fd5b813567ffffffffffffffff8082111561090a5761090a6108c8565b604051601f8301601f19908116603f01168101908282118183101715610932576109326108c8565b8160405283815286602085880101111561094b57600080fd5b836020870160208301376000602085830101528094505050505092915050565b60006020828403121561097d57600080fd5b813567ffffffffffffffff81111561099457600080fd5b6109a0848285016108de565b949350505050565b80356001600160a01b03811681146109bf57600080fd5b919050565b600080604083850312156109d757600080fd5b823567ffffffffffffffff8111156109ee57600080fd5b6109fa858286016108de565b925050610a09602084016109a8565b90509250929050565b600060208284031215610a2457600080fd5b610a2d826109a8565b9392505050565b60005b83811015610a4f578181015183820152602001610a37565b83811115610a5e576000848401525b50505050565b6020815260008251806020840152610a83816040850160208701610a34565b601f01601f19169190910160400192915050565b600060208284031215610aa957600080fd5b81358015158114610a2d57600080fd5b600060208284031215610acb57600080fd5b5035919050565b60008251610ae4818460208701610a34565b9190910192915050565b600181811c90821680610b0257607f821691505b60208210811415610b2357634e487b7160e01b600052602260045260246000fd5b50919050565b6000600019821415610b4b57634e487b7160e01b600052601160045260246000fd5b506001019056fea26469706673582212209724f9901a7a3ba74bf953429e7bc9885f86a90371853a66ddf514570886825a64736f6c63430008090033",
}

// ERRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use ERRegistryMetaData.ABI instead.
var ERRegistryABI = ERRegistryMetaData.ABI

// Deprecated: Use ERRegistryMetaData.Sigs instead.
// ERRegistryFuncSigs maps the 4-byte function signature to its string representation.
var ERRegistryFuncSigs = ERRegistryMetaData.Sigs

// ERRegistryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ERRegistryMetaData.Bin instead.
var ERRegistryBin = ERRegistryMetaData.Bin

// DeployERRegistry deploys a new Ethereum contract, binding an instance of ERRegistry to it.
func DeployERRegistry(auth *bind.TransactOpts, backend bind.ContractBackend, SubName string, SubDescription string) (common.Address, *types.Transaction, *ERRegistry, error) {
	parsed, err := ERRegistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ERRegistryBin), backend, SubName, SubDescription)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERRegistry{ERRegistryCaller: ERRegistryCaller{contract: contract}, ERRegistryTransactor: ERRegistryTransactor{contract: contract}, ERRegistryFilterer: ERRegistryFilterer{contract: contract}}, nil
}

// ERRegistry is an auto generated Go binding around an Ethereum contract.
type ERRegistry struct {
	ERRegistryCaller     // Read-only binding to the contract
	ERRegistryTransactor // Write-only binding to the contract
	ERRegistryFilterer   // Log filterer for contract events
}

// ERRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERRegistrySession struct {
	Contract     *ERRegistry       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERRegistryCallerSession struct {
	Contract *ERRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ERRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERRegistryTransactorSession struct {
	Contract     *ERRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ERRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERRegistryRaw struct {
	Contract *ERRegistry // Generic contract binding to access the raw methods on
}

// ERRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERRegistryCallerRaw struct {
	Contract *ERRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// ERRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERRegistryTransactorRaw struct {
	Contract *ERRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERRegistry creates a new instance of ERRegistry, bound to a specific deployed contract.
func NewERRegistry(address common.Address, backend bind.ContractBackend) (*ERRegistry, error) {
	contract, err := bindERRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERRegistry{ERRegistryCaller: ERRegistryCaller{contract: contract}, ERRegistryTransactor: ERRegistryTransactor{contract: contract}, ERRegistryFilterer: ERRegistryFilterer{contract: contract}}, nil
}

// NewERRegistryCaller creates a new read-only instance of ERRegistry, bound to a specific deployed contract.
func NewERRegistryCaller(address common.Address, caller bind.ContractCaller) (*ERRegistryCaller, error) {
	contract, err := bindERRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERRegistryCaller{contract: contract}, nil
}

// NewERRegistryTransactor creates a new write-only instance of ERRegistry, bound to a specific deployed contract.
func NewERRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*ERRegistryTransactor, error) {
	contract, err := bindERRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERRegistryTransactor{contract: contract}, nil
}

// NewERRegistryFilterer creates a new log filterer instance of ERRegistry, bound to a specific deployed contract.
func NewERRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*ERRegistryFilterer, error) {
	contract, err := bindERRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERRegistryFilterer{contract: contract}, nil
}

// bindERRegistry binds a generic wrapper to an already deployed contract.
func bindERRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERRegistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERRegistry *ERRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERRegistry.Contract.ERRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERRegistry *ERRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERRegistry.Contract.ERRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERRegistry *ERRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERRegistry.Contract.ERRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERRegistry *ERRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERRegistry *ERRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERRegistry *ERRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERRegistry.Contract.contract.Transact(opts, method, params...)
}

// Description is a free data retrieval call binding the contract method 0x8191745f.
//
// Solidity: function Description() view returns(string)
func (_ERRegistry *ERRegistryCaller) Description(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERRegistry.contract.Call(opts, &out, "Description")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Description is a free data retrieval call binding the contract method 0x8191745f.
//
// Solidity: function Description() view returns(string)
func (_ERRegistry *ERRegistrySession) Description() (string, error) {
	return _ERRegistry.Contract.Description(&_ERRegistry.CallOpts)
}

// Description is a free data retrieval call binding the contract method 0x8191745f.
//
// Solidity: function Description() view returns(string)
func (_ERRegistry *ERRegistryCallerSession) Description() (string, error) {
	return _ERRegistry.Contract.Description(&_ERRegistry.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x8052474d.
//
// Solidity: function Name() view returns(string)
func (_ERRegistry *ERRegistryCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERRegistry.contract.Call(opts, &out, "Name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x8052474d.
//
// Solidity: function Name() view returns(string)
func (_ERRegistry *ERRegistrySession) Name() (string, error) {
	return _ERRegistry.Contract.Name(&_ERRegistry.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x8052474d.
//
// Solidity: function Name() view returns(string)
func (_ERRegistry *ERRegistryCallerSession) Name() (string, error) {
	return _ERRegistry.Contract.Name(&_ERRegistry.CallOpts)
}

// RegistrationDisabled is a free data retrieval call binding the contract method 0x4a51a936.
//
// Solidity: function _registrationDisabled() view returns(bool)
func (_ERRegistry *ERRegistryCaller) RegistrationDisabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ERRegistry.contract.Call(opts, &out, "_registrationDisabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RegistrationDisabled is a free data retrieval call binding the contract method 0x4a51a936.
//
// Solidity: function _registrationDisabled() view returns(bool)
func (_ERRegistry *ERRegistrySession) RegistrationDisabled() (bool, error) {
	return _ERRegistry.Contract.RegistrationDisabled(&_ERRegistry.CallOpts)
}

// RegistrationDisabled is a free data retrieval call binding the contract method 0x4a51a936.
//
// Solidity: function _registrationDisabled() view returns(bool)
func (_ERRegistry *ERRegistryCallerSession) RegistrationDisabled() (bool, error) {
	return _ERRegistry.Contract.RegistrationDisabled(&_ERRegistry.CallOpts)
}

// GetAddressOfId is a free data retrieval call binding the contract method 0xec43eeb6.
//
// Solidity: function getAddressOfId(uint256 id) view returns(address addr)
func (_ERRegistry *ERRegistryCaller) GetAddressOfId(opts *bind.CallOpts, id *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ERRegistry.contract.Call(opts, &out, "getAddressOfId", id)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddressOfId is a free data retrieval call binding the contract method 0xec43eeb6.
//
// Solidity: function getAddressOfId(uint256 id) view returns(address addr)
func (_ERRegistry *ERRegistrySession) GetAddressOfId(id *big.Int) (common.Address, error) {
	return _ERRegistry.Contract.GetAddressOfId(&_ERRegistry.CallOpts, id)
}

// GetAddressOfId is a free data retrieval call binding the contract method 0xec43eeb6.
//
// Solidity: function getAddressOfId(uint256 id) view returns(address addr)
func (_ERRegistry *ERRegistryCallerSession) GetAddressOfId(id *big.Int) (common.Address, error) {
	return _ERRegistry.Contract.GetAddressOfId(&_ERRegistry.CallOpts, id)
}

// GetAddressOfName is a free data retrieval call binding the contract method 0x38ec6ba8.
//
// Solidity: function getAddressOfName(string name) view returns(address addr)
func (_ERRegistry *ERRegistryCaller) GetAddressOfName(opts *bind.CallOpts, name string) (common.Address, error) {
	var out []interface{}
	err := _ERRegistry.contract.Call(opts, &out, "getAddressOfName", name)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddressOfName is a free data retrieval call binding the contract method 0x38ec6ba8.
//
// Solidity: function getAddressOfName(string name) view returns(address addr)
func (_ERRegistry *ERRegistrySession) GetAddressOfName(name string) (common.Address, error) {
	return _ERRegistry.Contract.GetAddressOfName(&_ERRegistry.CallOpts, name)
}

// GetAddressOfName is a free data retrieval call binding the contract method 0x38ec6ba8.
//
// Solidity: function getAddressOfName(string name) view returns(address addr)
func (_ERRegistry *ERRegistryCallerSession) GetAddressOfName(name string) (common.Address, error) {
	return _ERRegistry.Contract.GetAddressOfName(&_ERRegistry.CallOpts, name)
}

// GetNameOfAddress is a free data retrieval call binding the contract method 0xb2dad25d.
//
// Solidity: function getNameOfAddress(address addr) view returns(string name)
func (_ERRegistry *ERRegistryCaller) GetNameOfAddress(opts *bind.CallOpts, addr common.Address) (string, error) {
	var out []interface{}
	err := _ERRegistry.contract.Call(opts, &out, "getNameOfAddress", addr)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetNameOfAddress is a free data retrieval call binding the contract method 0xb2dad25d.
//
// Solidity: function getNameOfAddress(address addr) view returns(string name)
func (_ERRegistry *ERRegistrySession) GetNameOfAddress(addr common.Address) (string, error) {
	return _ERRegistry.Contract.GetNameOfAddress(&_ERRegistry.CallOpts, addr)
}

// GetNameOfAddress is a free data retrieval call binding the contract method 0xb2dad25d.
//
// Solidity: function getNameOfAddress(address addr) view returns(string name)
func (_ERRegistry *ERRegistryCallerSession) GetNameOfAddress(addr common.Address) (string, error) {
	return _ERRegistry.Contract.GetNameOfAddress(&_ERRegistry.CallOpts, addr)
}

// NumberOfAccounts is a free data retrieval call binding the contract method 0x0f03e4c3.
//
// Solidity: function numberOfAccounts() view returns(uint256)
func (_ERRegistry *ERRegistryCaller) NumberOfAccounts(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERRegistry.contract.Call(opts, &out, "numberOfAccounts")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumberOfAccounts is a free data retrieval call binding the contract method 0x0f03e4c3.
//
// Solidity: function numberOfAccounts() view returns(uint256)
func (_ERRegistry *ERRegistrySession) NumberOfAccounts() (*big.Int, error) {
	return _ERRegistry.Contract.NumberOfAccounts(&_ERRegistry.CallOpts)
}

// NumberOfAccounts is a free data retrieval call binding the contract method 0x0f03e4c3.
//
// Solidity: function numberOfAccounts() view returns(uint256)
func (_ERRegistry *ERRegistryCallerSession) NumberOfAccounts() (*big.Int, error) {
	return _ERRegistry.Contract.NumberOfAccounts(&_ERRegistry.CallOpts)
}

// AdminRetrieveDonations is a paid mutator transaction binding the contract method 0x345e3416.
//
// Solidity: function adminRetrieveDonations() returns()
func (_ERRegistry *ERRegistryTransactor) AdminRetrieveDonations(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERRegistry.contract.Transact(opts, "adminRetrieveDonations")
}

// AdminRetrieveDonations is a paid mutator transaction binding the contract method 0x345e3416.
//
// Solidity: function adminRetrieveDonations() returns()
func (_ERRegistry *ERRegistrySession) AdminRetrieveDonations() (*types.Transaction, error) {
	return _ERRegistry.Contract.AdminRetrieveDonations(&_ERRegistry.TransactOpts)
}

// AdminRetrieveDonations is a paid mutator transaction binding the contract method 0x345e3416.
//
// Solidity: function adminRetrieveDonations() returns()
func (_ERRegistry *ERRegistryTransactorSession) AdminRetrieveDonations() (*types.Transaction, error) {
	return _ERRegistry.Contract.AdminRetrieveDonations(&_ERRegistry.TransactOpts)
}

// AdminSetAccountAdministrator is a paid mutator transaction binding the contract method 0x49f0c90d.
//
// Solidity: function adminSetAccountAdministrator(address accountAdmin) returns()
func (_ERRegistry *ERRegistryTransactor) AdminSetAccountAdministrator(opts *bind.TransactOpts, accountAdmin common.Address) (*types.Transaction, error) {
	return _ERRegistry.contract.Transact(opts, "adminSetAccountAdministrator", accountAdmin)
}

// AdminSetAccountAdministrator is a paid mutator transaction binding the contract method 0x49f0c90d.
//
// Solidity: function adminSetAccountAdministrator(address accountAdmin) returns()
func (_ERRegistry *ERRegistrySession) AdminSetAccountAdministrator(accountAdmin common.Address) (*types.Transaction, error) {
	return _ERRegistry.Contract.AdminSetAccountAdministrator(&_ERRegistry.TransactOpts, accountAdmin)
}

// AdminSetAccountAdministrator is a paid mutator transaction binding the contract method 0x49f0c90d.
//
// Solidity: function adminSetAccountAdministrator(address accountAdmin) returns()
func (_ERRegistry *ERRegistryTransactorSession) AdminSetAccountAdministrator(accountAdmin common.Address) (*types.Transaction, error) {
	return _ERRegistry.Contract.AdminSetAccountAdministrator(&_ERRegistry.TransactOpts, accountAdmin)
}

// AdminSetRegistrationDisabled is a paid mutator transaction binding the contract method 0x9b6d86d6.
//
// Solidity: function adminSetRegistrationDisabled(bool registrationDisabled) returns()
func (_ERRegistry *ERRegistryTransactor) AdminSetRegistrationDisabled(opts *bind.TransactOpts, registrationDisabled bool) (*types.Transaction, error) {
	return _ERRegistry.contract.Transact(opts, "adminSetRegistrationDisabled", registrationDisabled)
}

// AdminSetRegistrationDisabled is a paid mutator transaction binding the contract method 0x9b6d86d6.
//
// Solidity: function adminSetRegistrationDisabled(bool registrationDisabled) returns()
func (_ERRegistry *ERRegistrySession) AdminSetRegistrationDisabled(registrationDisabled bool) (*types.Transaction, error) {
	return _ERRegistry.Contract.AdminSetRegistrationDisabled(&_ERRegistry.TransactOpts, registrationDisabled)
}

// AdminSetRegistrationDisabled is a paid mutator transaction binding the contract method 0x9b6d86d6.
//
// Solidity: function adminSetRegistrationDisabled(bool registrationDisabled) returns()
func (_ERRegistry *ERRegistryTransactorSession) AdminSetRegistrationDisabled(registrationDisabled bool) (*types.Transaction, error) {
	return _ERRegistry.Contract.AdminSetRegistrationDisabled(&_ERRegistry.TransactOpts, registrationDisabled)
}

// AdminUnregister is a paid mutator transaction binding the contract method 0x0cc04b55.
//
// Solidity: function adminUnregister(string name) returns()
func (_ERRegistry *ERRegistryTransactor) AdminUnregister(opts *bind.TransactOpts, name string) (*types.Transaction, error) {
	return _ERRegistry.contract.Transact(opts, "adminUnregister", name)
}

// AdminUnregister is a paid mutator transaction binding the contract method 0x0cc04b55.
//
// Solidity: function adminUnregister(string name) returns()
func (_ERRegistry *ERRegistrySession) AdminUnregister(name string) (*types.Transaction, error) {
	return _ERRegistry.Contract.AdminUnregister(&_ERRegistry.TransactOpts, name)
}

// AdminUnregister is a paid mutator transaction binding the contract method 0x0cc04b55.
//
// Solidity: function adminUnregister(string name) returns()
func (_ERRegistry *ERRegistryTransactorSession) AdminUnregister(name string) (*types.Transaction, error) {
	return _ERRegistry.Contract.AdminUnregister(&_ERRegistry.TransactOpts, name)
}

// Register is a paid mutator transaction binding the contract method 0x1e59c529.
//
// Solidity: function register(string name, address accountAddress) returns()
func (_ERRegistry *ERRegistryTransactor) Register(opts *bind.TransactOpts, name string, accountAddress common.Address) (*types.Transaction, error) {
	return _ERRegistry.contract.Transact(opts, "register", name, accountAddress)
}

// Register is a paid mutator transaction binding the contract method 0x1e59c529.
//
// Solidity: function register(string name, address accountAddress) returns()
func (_ERRegistry *ERRegistrySession) Register(name string, accountAddress common.Address) (*types.Transaction, error) {
	return _ERRegistry.Contract.Register(&_ERRegistry.TransactOpts, name, accountAddress)
}

// Register is a paid mutator transaction binding the contract method 0x1e59c529.
//
// Solidity: function register(string name, address accountAddress) returns()
func (_ERRegistry *ERRegistryTransactorSession) Register(name string, accountAddress common.Address) (*types.Transaction, error) {
	return _ERRegistry.Contract.Register(&_ERRegistry.TransactOpts, name, accountAddress)
}

// Unregister is a paid mutator transaction binding the contract method 0xe79a198f.
//
// Solidity: function unregister() returns()
func (_ERRegistry *ERRegistryTransactor) Unregister(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERRegistry.contract.Transact(opts, "unregister")
}

// Unregister is a paid mutator transaction binding the contract method 0xe79a198f.
//
// Solidity: function unregister() returns()
func (_ERRegistry *ERRegistrySession) Unregister() (*types.Transaction, error) {
	return _ERRegistry.Contract.Unregister(&_ERRegistry.TransactOpts)
}

// Unregister is a paid mutator transaction binding the contract method 0xe79a198f.
//
// Solidity: function unregister() returns()
func (_ERRegistry *ERRegistryTransactorSession) Unregister() (*types.Transaction, error) {
	return _ERRegistry.Contract.Unregister(&_ERRegistry.TransactOpts)
}
