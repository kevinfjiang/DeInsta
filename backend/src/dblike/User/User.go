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

// DIAccountMetaData contains all meta data concerning the DIAccount contract.
var DIAccountMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"Name_\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"PostNumber\",\"type\":\"uint256\"}],\"name\":\"PostBanner\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"PostNumber\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"CommentString\",\"type\":\"string\"}],\"name\":\"AddComment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Bio\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int8\",\"name\":\"vote\",\"type\":\"int8\"},{\"internalType\":\"uint256\",\"name\":\"PostNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"CommentNumber\",\"type\":\"uint256\"}],\"name\":\"CommentVote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CountPosts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"PostURL_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"Caption_\",\"type\":\"string\"}],\"name\":\"CreatePost\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DeleteAccount\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"PostID\",\"type\":\"uint256\"}],\"name\":\"DeletePost\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Karma\",\"outputs\":[{\"internalType\":\"int16\",\"name\":\"\",\"type\":\"int16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"Posts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"PostNumber\",\"type\":\"uint256\"},{\"internalType\":\"int16\",\"name\":\"Karma\",\"type\":\"int16\"},{\"internalType\":\"string\",\"name\":\"IPFSurl\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"Caption\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"CountComments\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"Bio_\",\"type\":\"string\"}],\"name\":\"SetBio\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int8\",\"name\":\"vote\",\"type\":\"int8\"},{\"internalType\":\"uint256\",\"name\":\"PostNumber\",\"type\":\"uint256\"}],\"name\":\"Vote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"3b0cf47b": "AddComment(uint256,string)",
		"c2d2d0f9": "Bio()",
		"117d7f48": "CommentVote(int8,uint256,uint256)",
		"27677343": "CountPosts()",
		"7a8d157e": "CreatePost(string,string)",
		"9099967e": "DeleteAccount()",
		"75c16cb2": "DeletePost(uint256)",
		"e5163de5": "Karma()",
		"8052474d": "Name()",
		"4fafef1a": "Posts(uint256)",
		"27d0b9ec": "SetBio(string)",
		"ae81085a": "Vote(int8,uint256)",
	},
	Bin: "0x60806040523480156200001157600080fd5b50604051620013f8380380620013f8833981016040819052620000349162000124565b80516200004990600090602084019062000068565b50506000600455600580546001600160a01b031916331790556200023d565b828054620000769062000200565b90600052602060002090601f0160209004810192826200009a5760008555620000e5565b82601f10620000b557805160ff1916838001178555620000e5565b82800160010185558215620000e5579182015b82811115620000e5578251825591602001919060010190620000c8565b50620000f3929150620000f7565b5090565b5b80821115620000f35760008155600101620000f8565b634e487b7160e01b600052604160045260246000fd5b600060208083850312156200013857600080fd5b82516001600160401b03808211156200015057600080fd5b818501915085601f8301126200016557600080fd5b8151818111156200017a576200017a6200010e565b604051601f8201601f19908116603f01168101908382118183101715620001a557620001a56200010e565b816040528281528886848701011115620001be57600080fd5b600093505b82841015620001e25784840186015181850187015292850192620001c3565b82841115620001f45760008684830101525b98975050505050505050565b600181811c908216806200021557607f821691505b602082108114156200023757634e487b7160e01b600052602260045260246000fd5b50919050565b6111ab806200024d6000396000f3fe6080604052600436106100a75760003560e01c80637a8d157e116100645780637a8d157e146101895780638052474d146101a95780639099967e146101cb578063ae81085a146101d3578063c2d2d0f9146101f3578063e5163de51461020857600080fd5b8063117d7f48146100ac57806327677343146100ce57806327d0b9ec146100f75780633b0cf47b146101175780634fafef1a1461013757806375c16cb214610169575b600080fd5b3480156100b857600080fd5b506100cc6100c7366004610dbc565b610235565b005b3480156100da57600080fd5b506100e460045481565b6040519081526020015b60405180910390f35b34801561010357600080fd5b506100cc610112366004610e92565b6104be565b34801561012357600080fd5b506100cc610132366004610ecf565b6104ec565b34801561014357600080fd5b50610157610152366004610f16565b610643565b6040516100ee96959493929190610f7c565b34801561017557600080fd5b506100cc610184366004610f16565b61079e565b34801561019557600080fd5b506100cc6101a4366004610fca565b610827565b3480156101b557600080fd5b506101be610932565b6040516100ee9190611024565b6100cc6109c0565b3480156101df57600080fd5b506100cc6101ee36600461103e565b6109e5565b3480156101ff57600080fd5b506101be610c58565b34801561021457600080fd5b506003546102229060010b81565b60405160019190910b81526020016100ee565b60008360000b1361025a5760008360000b1261025257600061025d565b60001961025d565b60015b92508260000b6002838154811061027657610276611068565b9060005260206000209060080201600701828154811061029857610298611068565b6000918252602080832033845260036006909302019190910190526040812054900b14156103045760405162461bcd60e51b8152602060048201526014602482015273165bdd49dd9948185b1c9958591e481d9bdd195960621b60448201526064015b60405180910390fd5b6002828154811061031757610317611068565b9060005260206000209060080201600701818154811061033957610339611068565b6000918252602080832033845260036006909302019190910190526040812054600280549190920b91908490811061037357610373611068565b9060005260206000209060080201600701828154811061039557610395611068565b60009182526020822060026006909202010180549091906103ba90849060010b611094565b92506101000a81548161ffff021916908360010b61ffff1602179055508260000b600283815481106103ee576103ee611068565b9060005260206000209060080201600701828154811061041057610410611068565b600091825260208220600260069092020101805490919061043590849060010b6110da565b92506101000a81548161ffff021916908360010b61ffff160217905550826002838154811061046657610466611068565b9060005260206000209060080201600701828154811061048857610488611068565b60009182526020808320338452600692909202909101600301905260409020805460ff191660ff92909216919091179055505050565b6005546001600160a01b031633146104d557600080fd5b80516104e8906001906020840190610c65565b5050565b603281511061053d5760405162461bcd60e51b815260206004820152601c60248201527f506f7374206578636565647320636861726163746572206c696d69740000000060448201526064016102fb565b60006002838154811061055257610552611068565b90600052602060002090600802016007016002848154811061057657610576611068565b9060005260206000209060080201600601548154811061059857610598611068565b90600052602060002090600602019050428160000181905550600283815481106105c4576105c4611068565b600091825260209182902060066008909202010154600183015582516105f291600484019190850190610c65565b506005810180546001600160a01b03191633179055600280548490811061061b5761061b611068565b6000918252602082206006600890920201018054916106398361111f565b9190505550505050565b6002818154811061065357600080fd5b60009182526020909120600890910201805460018083015460028401546004850180549496509194920b926106879061113a565b80601f01602080910402602001604051908101604052809291908181526020018280546106b39061113a565b80156107005780601f106106d557610100808354040283529160200191610700565b820191906000526020600020905b8154815290600101906020018083116106e357829003601f168201915b5050505050908060050180546107159061113a565b80601f01602080910402602001604051908101604052809291908181526020018280546107419061113a565b801561078e5780601f106107635761010080835404028352916020019161078e565b820191906000526020600020905b81548152906001019060200180831161077157829003601f168201915b5050505050908060060154905086565b6005546001600160a01b031633146107b557600080fd5b600281815481106107c8576107c8611068565b6000918252602082206008909102018181556001810182905560028101805461ffff19169055906107fc6004830182610ce9565b61080a600583016000610ce9565b60068201600090556007820160006108229190610d26565b505050565b6005546001600160a01b0316331461083e57600080fd5b603281511061088f5760405162461bcd60e51b815260206004820152601c60248201527f506f7374206578636565647320636861726163746572206c696d69740000000060448201526064016102fb565b60006002600454815481106108a6576108a6611068565b60009182526020918290204260089092020190815560048054600183015585519193506108da929084019190860190610c65565b5081516108f09060058301906020850190610c65565b5060045481546040513091907f6e8001e55f8223a69373fdb10aecda8ca632280f2f689cbf7365b8c728e5187890600090a4600480549060006106398361111f565b6000805461093f9061113a565b80601f016020809104026020016040519081016040528092919081815260200182805461096b9061113a565b80156109b85780601f1061098d576101008083540402835291602001916109b8565b820191906000526020600020905b81548152906001019060200180831161099b57829003601f168201915b505050505081565b6005546001600160a01b031633146109d757600080fd5b6005546001600160a01b0316ff5b60008260000b13610a0a5760008260000b12610a02576000610a0d565b600019610a0d565b60015b91508160000b60028281548110610a2657610a26611068565b6000918252602080832033845260036008909302019190910190526040812054900b1415610a8d5760405162461bcd60e51b8152602060048201526014602482015273165bdd49dd9948185b1c9958591e481d9bdd195960621b60448201526064016102fb565b60028181548110610aa057610aa0611068565b6000918252602080832033845260036008909302019190910190526040812054600280549190920b919083908110610ada57610ada611068565b6000918252602082206002600890920201018054909190610aff90849060010b611094565b92506101000a81548161ffff021916908360010b61ffff16021790555060028181548110610b2f57610b2f611068565b600091825260208083203384526003600890930201820190526040822054815490830b9290610b6290849060010b611094565b92506101000a81548161ffff021916908360010b61ffff1602179055508160000b60028281548110610b9657610b96611068565b6000918252602082206002600890920201018054909190610bbb90849060010b6110da565b825461ffff9182166101009390930a92830291909202199091161790555060038054600084810b9291610bf290849060010b6110da565b92506101000a81548161ffff021916908360010b61ffff1602179055508160028281548110610c2357610c23611068565b60009182526020808320338452600892909202909101600301905260409020805460ff191660ff929092169190911790555050565b6001805461093f9061113a565b828054610c719061113a565b90600052602060002090601f016020900481019282610c935760008555610cd9565b82601f10610cac57805160ff1916838001178555610cd9565b82800160010185558215610cd9579182015b82811115610cd9578251825591602001919060010190610cbe565b50610ce5929150610d47565b5090565b508054610cf59061113a565b6000825580601f10610d05575050565b601f016020900490600052602060002090810190610d239190610d47565b50565b5080546000825560060290600052602060002090810190610d239190610d5c565b5b80821115610ce55760008155600101610d48565b80821115610ce55760008082556001820181905560028201805461ffff19169055610d8a6004830182610ce9565b506005810180546001600160a01b0319169055600601610d5c565b8035600081900b8114610db757600080fd5b919050565b600080600060608486031215610dd157600080fd5b610dda84610da5565b95602085013595506040909401359392505050565b634e487b7160e01b600052604160045260246000fd5b600082601f830112610e1657600080fd5b813567ffffffffffffffff80821115610e3157610e31610def565b604051601f8301601f19908116603f01168101908282118183101715610e5957610e59610def565b81604052838152866020858801011115610e7257600080fd5b836020870160208301376000602085830101528094505050505092915050565b600060208284031215610ea457600080fd5b813567ffffffffffffffff811115610ebb57600080fd5b610ec784828501610e05565b949350505050565b60008060408385031215610ee257600080fd5b82359150602083013567ffffffffffffffff811115610f0057600080fd5b610f0c85828601610e05565b9150509250929050565b600060208284031215610f2857600080fd5b5035919050565b6000815180845260005b81811015610f5557602081850181015186830182015201610f39565b81811115610f67576000602083870101525b50601f01601f19169290920160200192915050565b8681528560208201528460010b604082015260c060608201526000610fa460c0830186610f2f565b8281036080840152610fb68186610f2f565b9150508260a0830152979650505050505050565b60008060408385031215610fdd57600080fd5b823567ffffffffffffffff80821115610ff557600080fd5b61100186838701610e05565b9350602085013591508082111561101757600080fd5b50610f0c85828601610e05565b6020815260006110376020830184610f2f565b9392505050565b6000806040838503121561105157600080fd5b61105a83610da5565b946020939093013593505050565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b60008160010b8360010b6000811281617fff19018312811516156110ba576110ba61107e565b81617fff0183138116156110d0576110d061107e565b5090039392505050565b60008160010b8360010b6000821282617fff038213811516156110ff576110ff61107e565b82617fff190382128116156111165761111661107e565b50019392505050565b60006000198214156111335761113361107e565b5060010190565b600181811c9082168061114e57607f821691505b6020821081141561116f57634e487b7160e01b600052602260045260246000fd5b5091905056fea2646970667358221220fc8af8ebb0b4ccaa0e7f8c72e47e508648d9a93e69360c0fb91ed19c7573c42b64736f6c63430008090033",
}

// DIAccountABI is the input ABI used to generate the binding from.
// Deprecated: Use DIAccountMetaData.ABI instead.
var DIAccountABI = DIAccountMetaData.ABI

// Deprecated: Use DIAccountMetaData.Sigs instead.
// DIAccountFuncSigs maps the 4-byte function signature to its string representation.
var DIAccountFuncSigs = DIAccountMetaData.Sigs

// DIAccountBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DIAccountMetaData.Bin instead.
var DIAccountBin = DIAccountMetaData.Bin

// DeployDIAccount deploys a new Ethereum contract, binding an instance of DIAccount to it.
func DeployDIAccount(auth *bind.TransactOpts, backend bind.ContractBackend, Name_ string) (common.Address, *types.Transaction, *DIAccount, error) {
	parsed, err := DIAccountMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DIAccountBin), backend, Name_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DIAccount{DIAccountCaller: DIAccountCaller{contract: contract}, DIAccountTransactor: DIAccountTransactor{contract: contract}, DIAccountFilterer: DIAccountFilterer{contract: contract}}, nil
}

// DIAccount is an auto generated Go binding around an Ethereum contract.
type DIAccount struct {
	DIAccountCaller     // Read-only binding to the contract
	DIAccountTransactor // Write-only binding to the contract
	DIAccountFilterer   // Log filterer for contract events
}

// DIAccountCaller is an auto generated read-only Go binding around an Ethereum contract.
type DIAccountCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DIAccountTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DIAccountTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DIAccountFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DIAccountFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DIAccountSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DIAccountSession struct {
	Contract     *DIAccount        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DIAccountCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DIAccountCallerSession struct {
	Contract *DIAccountCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// DIAccountTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DIAccountTransactorSession struct {
	Contract     *DIAccountTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// DIAccountRaw is an auto generated low-level Go binding around an Ethereum contract.
type DIAccountRaw struct {
	Contract *DIAccount // Generic contract binding to access the raw methods on
}

// DIAccountCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DIAccountCallerRaw struct {
	Contract *DIAccountCaller // Generic read-only contract binding to access the raw methods on
}

// DIAccountTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DIAccountTransactorRaw struct {
	Contract *DIAccountTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDIAccount creates a new instance of DIAccount, bound to a specific deployed contract.
func NewDIAccount(address common.Address, backend bind.ContractBackend) (*DIAccount, error) {
	contract, err := bindDIAccount(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DIAccount{DIAccountCaller: DIAccountCaller{contract: contract}, DIAccountTransactor: DIAccountTransactor{contract: contract}, DIAccountFilterer: DIAccountFilterer{contract: contract}}, nil
}

// NewDIAccountCaller creates a new read-only instance of DIAccount, bound to a specific deployed contract.
func NewDIAccountCaller(address common.Address, caller bind.ContractCaller) (*DIAccountCaller, error) {
	contract, err := bindDIAccount(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DIAccountCaller{contract: contract}, nil
}

// NewDIAccountTransactor creates a new write-only instance of DIAccount, bound to a specific deployed contract.
func NewDIAccountTransactor(address common.Address, transactor bind.ContractTransactor) (*DIAccountTransactor, error) {
	contract, err := bindDIAccount(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DIAccountTransactor{contract: contract}, nil
}

// NewDIAccountFilterer creates a new log filterer instance of DIAccount, bound to a specific deployed contract.
func NewDIAccountFilterer(address common.Address, filterer bind.ContractFilterer) (*DIAccountFilterer, error) {
	contract, err := bindDIAccount(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DIAccountFilterer{contract: contract}, nil
}

// bindDIAccount binds a generic wrapper to an already deployed contract.
func bindDIAccount(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DIAccountABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DIAccount *DIAccountRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DIAccount.Contract.DIAccountCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DIAccount *DIAccountRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DIAccount.Contract.DIAccountTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DIAccount *DIAccountRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DIAccount.Contract.DIAccountTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DIAccount *DIAccountCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DIAccount.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DIAccount *DIAccountTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DIAccount.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DIAccount *DIAccountTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DIAccount.Contract.contract.Transact(opts, method, params...)
}

// Bio is a free data retrieval call binding the contract method 0xc2d2d0f9.
//
// Solidity: function Bio() view returns(string)
func (_DIAccount *DIAccountCaller) Bio(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _DIAccount.contract.Call(opts, &out, "Bio")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Bio is a free data retrieval call binding the contract method 0xc2d2d0f9.
//
// Solidity: function Bio() view returns(string)
func (_DIAccount *DIAccountSession) Bio() (string, error) {
	return _DIAccount.Contract.Bio(&_DIAccount.CallOpts)
}

// Bio is a free data retrieval call binding the contract method 0xc2d2d0f9.
//
// Solidity: function Bio() view returns(string)
func (_DIAccount *DIAccountCallerSession) Bio() (string, error) {
	return _DIAccount.Contract.Bio(&_DIAccount.CallOpts)
}

// CountPosts is a free data retrieval call binding the contract method 0x27677343.
//
// Solidity: function CountPosts() view returns(uint256)
func (_DIAccount *DIAccountCaller) CountPosts(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _DIAccount.contract.Call(opts, &out, "CountPosts")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CountPosts is a free data retrieval call binding the contract method 0x27677343.
//
// Solidity: function CountPosts() view returns(uint256)
func (_DIAccount *DIAccountSession) CountPosts() (*big.Int, error) {
	return _DIAccount.Contract.CountPosts(&_DIAccount.CallOpts)
}

// CountPosts is a free data retrieval call binding the contract method 0x27677343.
//
// Solidity: function CountPosts() view returns(uint256)
func (_DIAccount *DIAccountCallerSession) CountPosts() (*big.Int, error) {
	return _DIAccount.Contract.CountPosts(&_DIAccount.CallOpts)
}

// Karma is a free data retrieval call binding the contract method 0xe5163de5.
//
// Solidity: function Karma() view returns(int16)
func (_DIAccount *DIAccountCaller) Karma(opts *bind.CallOpts) (int16, error) {
	var out []interface{}
	err := _DIAccount.contract.Call(opts, &out, "Karma")

	if err != nil {
		return *new(int16), err
	}

	out0 := *abi.ConvertType(out[0], new(int16)).(*int16)

	return out0, err

}

// Karma is a free data retrieval call binding the contract method 0xe5163de5.
//
// Solidity: function Karma() view returns(int16)
func (_DIAccount *DIAccountSession) Karma() (int16, error) {
	return _DIAccount.Contract.Karma(&_DIAccount.CallOpts)
}

// Karma is a free data retrieval call binding the contract method 0xe5163de5.
//
// Solidity: function Karma() view returns(int16)
func (_DIAccount *DIAccountCallerSession) Karma() (int16, error) {
	return _DIAccount.Contract.Karma(&_DIAccount.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x8052474d.
//
// Solidity: function Name() view returns(string)
func (_DIAccount *DIAccountCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _DIAccount.contract.Call(opts, &out, "Name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x8052474d.
//
// Solidity: function Name() view returns(string)
func (_DIAccount *DIAccountSession) Name() (string, error) {
	return _DIAccount.Contract.Name(&_DIAccount.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x8052474d.
//
// Solidity: function Name() view returns(string)
func (_DIAccount *DIAccountCallerSession) Name() (string, error) {
	return _DIAccount.Contract.Name(&_DIAccount.CallOpts)
}

// Posts is a free data retrieval call binding the contract method 0x4fafef1a.
//
// Solidity: function Posts(uint256 ) view returns(uint256 timestamp, uint256 PostNumber, int16 Karma, string IPFSurl, string Caption, uint256 CountComments)
func (_DIAccount *DIAccountCaller) Posts(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Timestamp     *big.Int
	PostNumber    *big.Int
	Karma         int16
	IPFSurl       string
	Caption       string
	CountComments *big.Int
}, error) {
	var out []interface{}
	err := _DIAccount.contract.Call(opts, &out, "Posts", arg0)

	outstruct := new(struct {
		Timestamp     *big.Int
		PostNumber    *big.Int
		Karma         int16
		IPFSurl       string
		Caption       string
		CountComments *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Timestamp = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.PostNumber = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Karma = *abi.ConvertType(out[2], new(int16)).(*int16)
	outstruct.IPFSurl = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.Caption = *abi.ConvertType(out[4], new(string)).(*string)
	outstruct.CountComments = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Posts is a free data retrieval call binding the contract method 0x4fafef1a.
//
// Solidity: function Posts(uint256 ) view returns(uint256 timestamp, uint256 PostNumber, int16 Karma, string IPFSurl, string Caption, uint256 CountComments)
func (_DIAccount *DIAccountSession) Posts(arg0 *big.Int) (struct {
	Timestamp     *big.Int
	PostNumber    *big.Int
	Karma         int16
	IPFSurl       string
	Caption       string
	CountComments *big.Int
}, error) {
	return _DIAccount.Contract.Posts(&_DIAccount.CallOpts, arg0)
}

// Posts is a free data retrieval call binding the contract method 0x4fafef1a.
//
// Solidity: function Posts(uint256 ) view returns(uint256 timestamp, uint256 PostNumber, int16 Karma, string IPFSurl, string Caption, uint256 CountComments)
func (_DIAccount *DIAccountCallerSession) Posts(arg0 *big.Int) (struct {
	Timestamp     *big.Int
	PostNumber    *big.Int
	Karma         int16
	IPFSurl       string
	Caption       string
	CountComments *big.Int
}, error) {
	return _DIAccount.Contract.Posts(&_DIAccount.CallOpts, arg0)
}

// AddComment is a paid mutator transaction binding the contract method 0x3b0cf47b.
//
// Solidity: function AddComment(uint256 PostNumber, string CommentString) returns()
func (_DIAccount *DIAccountTransactor) AddComment(opts *bind.TransactOpts, PostNumber *big.Int, CommentString string) (*types.Transaction, error) {
	return _DIAccount.contract.Transact(opts, "AddComment", PostNumber, CommentString)
}

// AddComment is a paid mutator transaction binding the contract method 0x3b0cf47b.
//
// Solidity: function AddComment(uint256 PostNumber, string CommentString) returns()
func (_DIAccount *DIAccountSession) AddComment(PostNumber *big.Int, CommentString string) (*types.Transaction, error) {
	return _DIAccount.Contract.AddComment(&_DIAccount.TransactOpts, PostNumber, CommentString)
}

// AddComment is a paid mutator transaction binding the contract method 0x3b0cf47b.
//
// Solidity: function AddComment(uint256 PostNumber, string CommentString) returns()
func (_DIAccount *DIAccountTransactorSession) AddComment(PostNumber *big.Int, CommentString string) (*types.Transaction, error) {
	return _DIAccount.Contract.AddComment(&_DIAccount.TransactOpts, PostNumber, CommentString)
}

// CommentVote is a paid mutator transaction binding the contract method 0x117d7f48.
//
// Solidity: function CommentVote(int8 vote, uint256 PostNumber, uint256 CommentNumber) returns()
func (_DIAccount *DIAccountTransactor) CommentVote(opts *bind.TransactOpts, vote int8, PostNumber *big.Int, CommentNumber *big.Int) (*types.Transaction, error) {
	return _DIAccount.contract.Transact(opts, "CommentVote", vote, PostNumber, CommentNumber)
}

// CommentVote is a paid mutator transaction binding the contract method 0x117d7f48.
//
// Solidity: function CommentVote(int8 vote, uint256 PostNumber, uint256 CommentNumber) returns()
func (_DIAccount *DIAccountSession) CommentVote(vote int8, PostNumber *big.Int, CommentNumber *big.Int) (*types.Transaction, error) {
	return _DIAccount.Contract.CommentVote(&_DIAccount.TransactOpts, vote, PostNumber, CommentNumber)
}

// CommentVote is a paid mutator transaction binding the contract method 0x117d7f48.
//
// Solidity: function CommentVote(int8 vote, uint256 PostNumber, uint256 CommentNumber) returns()
func (_DIAccount *DIAccountTransactorSession) CommentVote(vote int8, PostNumber *big.Int, CommentNumber *big.Int) (*types.Transaction, error) {
	return _DIAccount.Contract.CommentVote(&_DIAccount.TransactOpts, vote, PostNumber, CommentNumber)
}

// CreatePost is a paid mutator transaction binding the contract method 0x7a8d157e.
//
// Solidity: function CreatePost(string PostURL_, string Caption_) returns()
func (_DIAccount *DIAccountTransactor) CreatePost(opts *bind.TransactOpts, PostURL_ string, Caption_ string) (*types.Transaction, error) {
	return _DIAccount.contract.Transact(opts, "CreatePost", PostURL_, Caption_)
}

// CreatePost is a paid mutator transaction binding the contract method 0x7a8d157e.
//
// Solidity: function CreatePost(string PostURL_, string Caption_) returns()
func (_DIAccount *DIAccountSession) CreatePost(PostURL_ string, Caption_ string) (*types.Transaction, error) {
	return _DIAccount.Contract.CreatePost(&_DIAccount.TransactOpts, PostURL_, Caption_)
}

// CreatePost is a paid mutator transaction binding the contract method 0x7a8d157e.
//
// Solidity: function CreatePost(string PostURL_, string Caption_) returns()
func (_DIAccount *DIAccountTransactorSession) CreatePost(PostURL_ string, Caption_ string) (*types.Transaction, error) {
	return _DIAccount.Contract.CreatePost(&_DIAccount.TransactOpts, PostURL_, Caption_)
}

// DeleteAccount is a paid mutator transaction binding the contract method 0x9099967e.
//
// Solidity: function DeleteAccount() payable returns()
func (_DIAccount *DIAccountTransactor) DeleteAccount(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DIAccount.contract.Transact(opts, "DeleteAccount")
}

// DeleteAccount is a paid mutator transaction binding the contract method 0x9099967e.
//
// Solidity: function DeleteAccount() payable returns()
func (_DIAccount *DIAccountSession) DeleteAccount() (*types.Transaction, error) {
	return _DIAccount.Contract.DeleteAccount(&_DIAccount.TransactOpts)
}

// DeleteAccount is a paid mutator transaction binding the contract method 0x9099967e.
//
// Solidity: function DeleteAccount() payable returns()
func (_DIAccount *DIAccountTransactorSession) DeleteAccount() (*types.Transaction, error) {
	return _DIAccount.Contract.DeleteAccount(&_DIAccount.TransactOpts)
}

// DeletePost is a paid mutator transaction binding the contract method 0x75c16cb2.
//
// Solidity: function DeletePost(uint256 PostID) returns()
func (_DIAccount *DIAccountTransactor) DeletePost(opts *bind.TransactOpts, PostID *big.Int) (*types.Transaction, error) {
	return _DIAccount.contract.Transact(opts, "DeletePost", PostID)
}

// DeletePost is a paid mutator transaction binding the contract method 0x75c16cb2.
//
// Solidity: function DeletePost(uint256 PostID) returns()
func (_DIAccount *DIAccountSession) DeletePost(PostID *big.Int) (*types.Transaction, error) {
	return _DIAccount.Contract.DeletePost(&_DIAccount.TransactOpts, PostID)
}

// DeletePost is a paid mutator transaction binding the contract method 0x75c16cb2.
//
// Solidity: function DeletePost(uint256 PostID) returns()
func (_DIAccount *DIAccountTransactorSession) DeletePost(PostID *big.Int) (*types.Transaction, error) {
	return _DIAccount.Contract.DeletePost(&_DIAccount.TransactOpts, PostID)
}

// SetBio is a paid mutator transaction binding the contract method 0x27d0b9ec.
//
// Solidity: function SetBio(string Bio_) returns()
func (_DIAccount *DIAccountTransactor) SetBio(opts *bind.TransactOpts, Bio_ string) (*types.Transaction, error) {
	return _DIAccount.contract.Transact(opts, "SetBio", Bio_)
}

// SetBio is a paid mutator transaction binding the contract method 0x27d0b9ec.
//
// Solidity: function SetBio(string Bio_) returns()
func (_DIAccount *DIAccountSession) SetBio(Bio_ string) (*types.Transaction, error) {
	return _DIAccount.Contract.SetBio(&_DIAccount.TransactOpts, Bio_)
}

// SetBio is a paid mutator transaction binding the contract method 0x27d0b9ec.
//
// Solidity: function SetBio(string Bio_) returns()
func (_DIAccount *DIAccountTransactorSession) SetBio(Bio_ string) (*types.Transaction, error) {
	return _DIAccount.Contract.SetBio(&_DIAccount.TransactOpts, Bio_)
}

// Vote is a paid mutator transaction binding the contract method 0xae81085a.
//
// Solidity: function Vote(int8 vote, uint256 PostNumber) returns()
func (_DIAccount *DIAccountTransactor) Vote(opts *bind.TransactOpts, vote int8, PostNumber *big.Int) (*types.Transaction, error) {
	return _DIAccount.contract.Transact(opts, "Vote", vote, PostNumber)
}

// Vote is a paid mutator transaction binding the contract method 0xae81085a.
//
// Solidity: function Vote(int8 vote, uint256 PostNumber) returns()
func (_DIAccount *DIAccountSession) Vote(vote int8, PostNumber *big.Int) (*types.Transaction, error) {
	return _DIAccount.Contract.Vote(&_DIAccount.TransactOpts, vote, PostNumber)
}

// Vote is a paid mutator transaction binding the contract method 0xae81085a.
//
// Solidity: function Vote(int8 vote, uint256 PostNumber) returns()
func (_DIAccount *DIAccountTransactorSession) Vote(vote int8, PostNumber *big.Int) (*types.Transaction, error) {
	return _DIAccount.Contract.Vote(&_DIAccount.TransactOpts, vote, PostNumber)
}

// DIAccountPostBannerIterator is returned from FilterPostBanner and is used to iterate over the raw logs and unpacked data for PostBanner events raised by the DIAccount contract.
type DIAccountPostBannerIterator struct {
	Event *DIAccountPostBanner // Event containing the contract specifics and raw log

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
func (it *DIAccountPostBannerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DIAccountPostBanner)
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
		it.Event = new(DIAccountPostBanner)
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
func (it *DIAccountPostBannerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DIAccountPostBannerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DIAccountPostBanner represents a PostBanner event raised by the DIAccount contract.
type DIAccountPostBanner struct {
	Timestamp  *big.Int
	Owner      common.Address
	PostNumber *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPostBanner is a free log retrieval operation binding the contract event 0x6e8001e55f8223a69373fdb10aecda8ca632280f2f689cbf7365b8c728e51878.
//
// Solidity: event PostBanner(uint256 indexed timestamp, address indexed owner, uint256 indexed PostNumber)
func (_DIAccount *DIAccountFilterer) FilterPostBanner(opts *bind.FilterOpts, timestamp []*big.Int, owner []common.Address, PostNumber []*big.Int) (*DIAccountPostBannerIterator, error) {

	var timestampRule []interface{}
	for _, timestampItem := range timestamp {
		timestampRule = append(timestampRule, timestampItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var PostNumberRule []interface{}
	for _, PostNumberItem := range PostNumber {
		PostNumberRule = append(PostNumberRule, PostNumberItem)
	}

	logs, sub, err := _DIAccount.contract.FilterLogs(opts, "PostBanner", timestampRule, ownerRule, PostNumberRule)
	if err != nil {
		return nil, err
	}
	return &DIAccountPostBannerIterator{contract: _DIAccount.contract, event: "PostBanner", logs: logs, sub: sub}, nil
}

// WatchPostBanner is a free log subscription operation binding the contract event 0x6e8001e55f8223a69373fdb10aecda8ca632280f2f689cbf7365b8c728e51878.
//
// Solidity: event PostBanner(uint256 indexed timestamp, address indexed owner, uint256 indexed PostNumber)
func (_DIAccount *DIAccountFilterer) WatchPostBanner(opts *bind.WatchOpts, sink chan<- *DIAccountPostBanner, timestamp []*big.Int, owner []common.Address, PostNumber []*big.Int) (event.Subscription, error) {

	var timestampRule []interface{}
	for _, timestampItem := range timestamp {
		timestampRule = append(timestampRule, timestampItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var PostNumberRule []interface{}
	for _, PostNumberItem := range PostNumber {
		PostNumberRule = append(PostNumberRule, PostNumberItem)
	}

	logs, sub, err := _DIAccount.contract.WatchLogs(opts, "PostBanner", timestampRule, ownerRule, PostNumberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DIAccountPostBanner)
				if err := _DIAccount.contract.UnpackLog(event, "PostBanner", log); err != nil {
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

// ParsePostBanner is a log parse operation binding the contract event 0x6e8001e55f8223a69373fdb10aecda8ca632280f2f689cbf7365b8c728e51878.
//
// Solidity: event PostBanner(uint256 indexed timestamp, address indexed owner, uint256 indexed PostNumber)
func (_DIAccount *DIAccountFilterer) ParsePostBanner(log types.Log) (*DIAccountPostBanner, error) {
	event := new(DIAccountPostBanner)
	if err := _DIAccount.contract.UnpackLog(event, "PostBanner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
