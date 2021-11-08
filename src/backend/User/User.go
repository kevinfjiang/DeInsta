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

// ERAccountMetaData contains all meta data concerning the ERAccount contract.
var ERAccountMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"Name_\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"CountPosts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"PostID\",\"type\":\"uint256\"}],\"name\":\"DeletePost\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetOwnerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"adminAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"PostID\",\"type\":\"uint256\"}],\"name\":\"GetPost\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"postaddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Karma\",\"outputs\":[{\"internalType\":\"int16\",\"name\":\"\",\"type\":\"int16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"PostName_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"PostString_\",\"type\":\"string\"}],\"name\":\"PostMessage\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"postaddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"Posts\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"adminDeleteAccount\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"27677343": "CountPosts()",
		"75c16cb2": "DeletePost(uint256)",
		"e6f7bf89": "GetOwnerAddress()",
		"bf421138": "GetPost(uint256)",
		"e5163de5": "Karma()",
		"8052474d": "Name()",
		"01465d36": "PostMessage(string,string)",
		"4fafef1a": "Posts(uint256)",
		"3e450fff": "adminDeleteAccount()",
	},
	Bin: "0x608060405234801561001057600080fd5b50604051620018e5380380620018e583398101604081905261003191610111565b8051610044906000906020840190610062565b50506000600355600480546001600160a01b0319163217905561021b565b82805461006e906101e0565b90600052602060002090601f01602090048101928261009057600085556100d6565b82601f106100a957805160ff19168380011785556100d6565b828001600101855582156100d6579182015b828111156100d65782518255916020019190600101906100bb565b506100e29291506100e6565b5090565b5b808211156100e257600081556001016100e7565b634e487b7160e01b600052604160045260246000fd5b6000602080838503121561012457600080fd5b82516001600160401b038082111561013b57600080fd5b818501915085601f83011261014f57600080fd5b815181811115610161576101616100fb565b604051601f8201601f19908116603f01168101908382118183101715610189576101896100fb565b8160405282815288868487010111156101a157600080fd5b600093505b828410156101c357848401860151818501870152928501926101a6565b828411156101d45760008684830101525b98975050505050505050565b600181811c908216806101f457607f821691505b6020821081141561021557634e487b7160e01b600052602260045260246000fd5b50919050565b6116ba806200022b6000396000f3fe608060405260043610620000915760003560e01c806375c16cb2116200006057806375c16cb214620001305780638052474d1462000155578063bf421138146200017c578063e5163de514620001a1578063e6f7bf8914620001d157600080fd5b806301465d3614620000965780632767734314620000d85780633e450fff14620000ff5780634fafef1a146200010b575b600080fd5b348015620000a357600080fd5b50620000bb620000b53660046200058b565b620001f1565b6040516001600160a01b0390911681526020015b60405180910390f35b348015620000e557600080fd5b50620000f060035481565b604051908152602001620000cf565b62000109620002f6565b005b3480156200011857600080fd5b50620000bb6200012a366004620005f6565b6200031c565b3480156200013d57600080fd5b50620001096200014f366004620005f6565b62000347565b3480156200016257600080fd5b506200016d62000409565b604051620000cf919062000660565b3480156200018957600080fd5b50620000bb6200019b366004620005f6565b6200049f565b348015620001ae57600080fd5b50600254620001bd9060010b81565b60405160019190910b8152602001620000cf565b348015620001de57600080fd5b506004546001600160a01b0316620000bb565b6004546000906001600160a01b031633146200020c57600080fd5b6101f4825110620002635760405162461bcd60e51b815260206004820152601c60248201527f506f7374206578636565647320636861726163746572206c696d697400000000604482015260640160405180910390fd5b600083836003546040516200027890620004d2565b62000286939291906200067c565b604051809103906000f080158015620002a3573d6000803e3d6000fd5b506001805480820182556000919091527fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf60180546001600160a01b0319166001600160a01b038316179055949350505050565b6004546001600160a01b031633146200030e57600080fd5b6004546001600160a01b0316ff5b600181815481106200032d57600080fd5b6000918252602090912001546001600160a01b0316905081565b6004546001600160a01b031633146200035f57600080fd5b60018181548110620003755762000375620006b6565b6000918252602082200154604080516385141ee360e01b815290516001600160a01b03909216926385141ee39260048084019382900301818387803b158015620003be57600080fd5b505af1158015620003d3573d6000803e3d6000fd5b5050505060018181548110620003ed57620003ed620006b6565b600091825260209091200180546001600160a01b031916905550565b600080546200041890620006cc565b80601f01602080910402602001604051908101604052809291908181526020018280546200044690620006cc565b8015620004975780601f106200046b5761010080835404028352916020019162000497565b820191906000526020600020905b8154815290600101906020018083116200047957829003601f168201915b505050505081565b600060018281548110620004b757620004b7620006b6565b6000918252602090912001546001600160a01b031692915050565b610f7b806200070a83390190565b634e487b7160e01b600052604160045260246000fd5b600082601f8301126200050857600080fd5b813567ffffffffffffffff80821115620005265762000526620004e0565b604051601f8301601f19908116603f01168101908282118183101715620005515762000551620004e0565b816040528381528660208588010111156200056b57600080fd5b836020870160208301376000602085830101528094505050505092915050565b600080604083850312156200059f57600080fd5b823567ffffffffffffffff80821115620005b857600080fd5b620005c686838701620004f6565b93506020850135915080821115620005dd57600080fd5b50620005ec85828601620004f6565b9150509250929050565b6000602082840312156200060957600080fd5b5035919050565b6000815180845260005b8181101562000638576020818501810151868301820152016200061a565b818111156200064b576000602083870101525b50601f01601f19169290920160200192915050565b60208152600062000675602083018462000610565b9392505050565b60608152600062000691606083018662000610565b8281036020840152620006a5818662000610565b915050826040830152949350505050565b634e487b7160e01b600052603260045260246000fd5b600181811c90821680620006e157607f821691505b602082108114156200070357634e487b7160e01b600052602260045260246000fd5b5091905056fe60806040523480156200001157600080fd5b5060405162000f7b38038062000f7b833981016040819052620000349162000241565b4260005582516200004d906004906020860190620000ce565b50815162000063906005906020850190620000ce565b506001818155600680546001600160a01b031916321790556000546002546040517fe7af46f298185143ea7dc9c524c80748fc11ca4248529d56a451f895aa5a7c8093620000bd9392900b90309060049060059062000399565b60405180910390a1505050620003eb565b828054620000dc90620002b4565b90600052602060002090601f0160209004810192826200010057600085556200014b565b82601f106200011b57805160ff19168380011785556200014b565b828001600101855582156200014b579182015b828111156200014b5782518255916020019190600101906200012e565b50620001599291506200015d565b5090565b5b808211156200015957600081556001016200015e565b634e487b7160e01b600052604160045260246000fd5b600082601f8301126200019c57600080fd5b81516001600160401b0380821115620001b957620001b962000174565b604051601f8301601f19908116603f01168101908282118183101715620001e457620001e462000174565b816040528381526020925086838588010111156200020157600080fd5b600091505b8382101562000225578582018301518183018401529082019062000206565b83821115620002375760008385830101525b9695505050505050565b6000806000606084860312156200025757600080fd5b83516001600160401b03808211156200026f57600080fd5b6200027d878388016200018a565b945060208601519150808211156200029457600080fd5b50620002a3868287016200018a565b925050604084015190509250925092565b600181811c90821680620002c957607f821691505b60208210811415620002eb57634e487b7160e01b600052602260045260246000fd5b50919050565b8054600090600181811c90808316806200030c57607f831692505b60208084108214156200032f57634e487b7160e01b600052602260045260246000fd5b838852602088018280156200034d57600181146200035f576200038c565b60ff198716825282820197506200038c565b60008981526020902060005b8781101562000386578154848201529086019084016200036b565b83019850505b5050505050505092915050565b8581528460010b602082015260018060a01b038416604082015260a060608201526000620003cb60a0830185620002f1565b8281036080840152620003df8185620002f1565b98975050505050505050565b610b8080620003fb6000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c80638ea20ddc116100715780638ea20ddc1461010457806398c024bc14610128578063ae81085a1461013f578063dc28c10014610152578063e5163de51461015b578063fe6c97521461017b57600080fd5b8063406dade3146100ae5780634e49d2da146100b8578063680e8ebc146100cb5780636bde24b6146100e957806385141ee3146100fc575b600080fd5b6100b6610183565b005b6100b66100c6366004610797565b6101d6565b6100d361034a565b6040516100e09190610806565b60405180910390f35b6100b66100f736600461082f565b6103d8565b6100b661048a565b6101176101123660046108e0565b6104af565b6040516100e09594939291906108f9565b61013160015481565b6040519081526020016100e0565b6100b661014d36600461093b565b610575565b61013160075481565b6002546101689060010b81565b60405160019190910b81526020016100e0565b6100d36106da565b6006546001600160a01b0316331461019a57600080fd5b6006546040516001600160a01b03909116904780156108fc02916000818181858888f193505050501580156101d3573d6000803e3d6000fd5b50565b60008160000b136101fb5760008160000b126101f35760006101fe565b6000196101fe565b60015b3360009081526003602052604081205491925082810b91900b14156102625760405162461bcd60e51b81526020600482015260156024820152744d757374206368616e6765207072657620766f746560581b60448201526064015b60405180910390fd5b336000908152600360205260408120546002805491830b92909161028a90849060010b61097b565b825461ffff9182166101009390930a92830291909202199091161790555060028054600083810b92916102c190849060010b6109c1565b82546101009290920a61ffff8181021990931691909216919091021790555033600090815260036020526040808220805460ff191660ff8516179055905460025491517fe7af46f298185143ea7dc9c524c80748fc11ca4248529d56a451f895aa5a7c809261033f929160019190910b903090600490600590610ae1565b60405180910390a150565b6004805461035790610a06565b80601f016020809104026020016040519081016040528092919081815260200182805461038390610a06565b80156103d05780601f106103a5576101008083540402835291602001916103d0565b820191906000526020600020905b8154815290600101906020018083116103b357829003601f168201915b505050505081565b60fa8151106104295760405162461bcd60e51b815260206004820152601c60248201527f506f7374206578636565647320636861726163746572206c696d6974000000006044820152606401610259565b60078054600090815260086020908152604090912042815591546001830155825161045c916004840191908501906106e7565b506005810180546001600160a01b031916331790556007805490600061048183610b2f565b91905055505050565b6006546001600160a01b031633146104a157600080fd5b6006546001600160a01b0316ff5b600860205260009081526040902080546001808301546002840154600485018054949592949190930b92906104e390610a06565b80601f016020809104026020016040519081016040528092919081815260200182805461050f90610a06565b801561055c5780601f106105315761010080835404028352916020019161055c565b820191906000526020600020905b81548152906001019060200180831161053f57829003601f168201915b505050600590930154919250506001600160a01b031685565b60008260000b1361059a5760008260000b1261059257600061059d565b60001961059d565b60015b600082815260086020908152604080832033845260030190915281205491935083810b91900b14156106095760405162461bcd60e51b81526020600482015260156024820152744d757374206368616e6765207072657620766f746560581b6044820152606401610259565b60008181526008602081815260408084203385526003810183529084205485855292909152600201805491830b92909161064790849060010b61097b565b825461ffff9182166101009390930a9283029190920219909116179055506000818152600860205260408120600201805484830b929061068b90849060010b6109c1565b825461ffff9182166101009390930a92830291909202199091161790555060009081526008602090815260408083203384526003019091529020805460ff90921660ff19909216919091179055565b6005805461035790610a06565b8280546106f390610a06565b90600052602060002090601f016020900481019282610715576000855561075b565b82601f1061072e57805160ff191683800117855561075b565b8280016001018555821561075b579182015b8281111561075b578251825591602001919060010190610740565b5061076792915061076b565b5090565b5b80821115610767576000815560010161076c565b8035600081900b811461079257600080fd5b919050565b6000602082840312156107a957600080fd5b6107b282610780565b9392505050565b6000815180845260005b818110156107df576020818501810151868301820152016107c3565b818111156107f1576000602083870101525b50601f01601f19169290920160200192915050565b6020815260006107b260208301846107b9565b634e487b7160e01b600052604160045260246000fd5b60006020828403121561084157600080fd5b813567ffffffffffffffff8082111561085957600080fd5b818401915084601f83011261086d57600080fd5b81358181111561087f5761087f610819565b604051601f8201601f19908116603f011681019083821181831017156108a7576108a7610819565b816040528281528760208487010111156108c057600080fd5b826020860160208301376000928101602001929092525095945050505050565b6000602082840312156108f257600080fd5b5035919050565b8581528460208201528360010b604082015260a06060820152600061092160a08301856107b9565b905060018060a01b03831660808301529695505050505050565b6000806040838503121561094e57600080fd5b61095783610780565b946020939093013593505050565b634e487b7160e01b600052601160045260246000fd5b60008160010b8360010b6000811281617fff19018312811516156109a1576109a1610965565b81617fff0183138116156109b7576109b7610965565b5090039392505050565b60008160010b8360010b6000821282617fff038213811516156109e6576109e6610965565b82617fff190382128116156109fd576109fd610965565b50019392505050565b600181811c90821680610a1a57607f821691505b60208210811415610a3b57634e487b7160e01b600052602260045260246000fd5b50919050565b8054600090600181811c9080831680610a5b57607f831692505b6020808410821415610a7d57634e487b7160e01b600052602260045260246000fd5b83885260208801828015610a985760018114610aa957610ad4565b60ff19871682528282019750610ad4565b60008981526020902060005b87811015610ace57815484820152908601908401610ab5565b83019850505b5050505050505092915050565b8581528460010b602082015260018060a01b038416604082015260a060608201526000610b1160a0830185610a41565b8281036080840152610b238185610a41565b98975050505050505050565b6000600019821415610b4357610b43610965565b506001019056fea2646970667358221220644f09dd28e689894945015501237899179602563292f57f2ed9f36ae94dafd464736f6c63430008090033a26469706673582212209109b6114c7cac0d873ad0ec83bad5d21bd11ce72c91fc0c8979d65ff2013c2b64736f6c63430008090033",
}

// ERAccountABI is the input ABI used to generate the binding from.
// Deprecated: Use ERAccountMetaData.ABI instead.
var ERAccountABI = ERAccountMetaData.ABI

// Deprecated: Use ERAccountMetaData.Sigs instead.
// ERAccountFuncSigs maps the 4-byte function signature to its string representation.
var ERAccountFuncSigs = ERAccountMetaData.Sigs

// ERAccountBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ERAccountMetaData.Bin instead.
var ERAccountBin = ERAccountMetaData.Bin

// DeployERAccount deploys a new Ethereum contract, binding an instance of ERAccount to it.
func DeployERAccount(auth *bind.TransactOpts, backend bind.ContractBackend, Name_ string) (common.Address, *types.Transaction, *ERAccount, error) {
	parsed, err := ERAccountMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ERAccountBin), backend, Name_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERAccount{ERAccountCaller: ERAccountCaller{contract: contract}, ERAccountTransactor: ERAccountTransactor{contract: contract}, ERAccountFilterer: ERAccountFilterer{contract: contract}}, nil
}

// ERAccount is an auto generated Go binding around an Ethereum contract.
type ERAccount struct {
	ERAccountCaller     // Read-only binding to the contract
	ERAccountTransactor // Write-only binding to the contract
	ERAccountFilterer   // Log filterer for contract events
}

// ERAccountCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERAccountCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERAccountTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERAccountTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERAccountFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERAccountFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERAccountSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERAccountSession struct {
	Contract     *ERAccount        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERAccountCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERAccountCallerSession struct {
	Contract *ERAccountCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ERAccountTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERAccountTransactorSession struct {
	Contract     *ERAccountTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ERAccountRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERAccountRaw struct {
	Contract *ERAccount // Generic contract binding to access the raw methods on
}

// ERAccountCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERAccountCallerRaw struct {
	Contract *ERAccountCaller // Generic read-only contract binding to access the raw methods on
}

// ERAccountTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERAccountTransactorRaw struct {
	Contract *ERAccountTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERAccount creates a new instance of ERAccount, bound to a specific deployed contract.
func NewERAccount(address common.Address, backend bind.ContractBackend) (*ERAccount, error) {
	contract, err := bindERAccount(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERAccount{ERAccountCaller: ERAccountCaller{contract: contract}, ERAccountTransactor: ERAccountTransactor{contract: contract}, ERAccountFilterer: ERAccountFilterer{contract: contract}}, nil
}

// NewERAccountCaller creates a new read-only instance of ERAccount, bound to a specific deployed contract.
func NewERAccountCaller(address common.Address, caller bind.ContractCaller) (*ERAccountCaller, error) {
	contract, err := bindERAccount(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERAccountCaller{contract: contract}, nil
}

// NewERAccountTransactor creates a new write-only instance of ERAccount, bound to a specific deployed contract.
func NewERAccountTransactor(address common.Address, transactor bind.ContractTransactor) (*ERAccountTransactor, error) {
	contract, err := bindERAccount(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERAccountTransactor{contract: contract}, nil
}

// NewERAccountFilterer creates a new log filterer instance of ERAccount, bound to a specific deployed contract.
func NewERAccountFilterer(address common.Address, filterer bind.ContractFilterer) (*ERAccountFilterer, error) {
	contract, err := bindERAccount(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERAccountFilterer{contract: contract}, nil
}

// bindERAccount binds a generic wrapper to an already deployed contract.
func bindERAccount(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERAccountABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERAccount *ERAccountRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERAccount.Contract.ERAccountCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERAccount *ERAccountRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERAccount.Contract.ERAccountTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERAccount *ERAccountRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERAccount.Contract.ERAccountTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERAccount *ERAccountCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERAccount.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERAccount *ERAccountTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERAccount.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERAccount *ERAccountTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERAccount.Contract.contract.Transact(opts, method, params...)
}

// CountPosts is a free data retrieval call binding the contract method 0x27677343.
//
// Solidity: function CountPosts() view returns(uint256)
func (_ERAccount *ERAccountCaller) CountPosts(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERAccount.contract.Call(opts, &out, "CountPosts")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CountPosts is a free data retrieval call binding the contract method 0x27677343.
//
// Solidity: function CountPosts() view returns(uint256)
func (_ERAccount *ERAccountSession) CountPosts() (*big.Int, error) {
	return _ERAccount.Contract.CountPosts(&_ERAccount.CallOpts)
}

// CountPosts is a free data retrieval call binding the contract method 0x27677343.
//
// Solidity: function CountPosts() view returns(uint256)
func (_ERAccount *ERAccountCallerSession) CountPosts() (*big.Int, error) {
	return _ERAccount.Contract.CountPosts(&_ERAccount.CallOpts)
}

// GetOwnerAddress is a free data retrieval call binding the contract method 0xe6f7bf89.
//
// Solidity: function GetOwnerAddress() view returns(address adminAddress)
func (_ERAccount *ERAccountCaller) GetOwnerAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ERAccount.contract.Call(opts, &out, "GetOwnerAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOwnerAddress is a free data retrieval call binding the contract method 0xe6f7bf89.
//
// Solidity: function GetOwnerAddress() view returns(address adminAddress)
func (_ERAccount *ERAccountSession) GetOwnerAddress() (common.Address, error) {
	return _ERAccount.Contract.GetOwnerAddress(&_ERAccount.CallOpts)
}

// GetOwnerAddress is a free data retrieval call binding the contract method 0xe6f7bf89.
//
// Solidity: function GetOwnerAddress() view returns(address adminAddress)
func (_ERAccount *ERAccountCallerSession) GetOwnerAddress() (common.Address, error) {
	return _ERAccount.Contract.GetOwnerAddress(&_ERAccount.CallOpts)
}

// GetPost is a free data retrieval call binding the contract method 0xbf421138.
//
// Solidity: function GetPost(uint256 PostID) view returns(address postaddress)
func (_ERAccount *ERAccountCaller) GetPost(opts *bind.CallOpts, PostID *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ERAccount.contract.Call(opts, &out, "GetPost", PostID)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPost is a free data retrieval call binding the contract method 0xbf421138.
//
// Solidity: function GetPost(uint256 PostID) view returns(address postaddress)
func (_ERAccount *ERAccountSession) GetPost(PostID *big.Int) (common.Address, error) {
	return _ERAccount.Contract.GetPost(&_ERAccount.CallOpts, PostID)
}

// GetPost is a free data retrieval call binding the contract method 0xbf421138.
//
// Solidity: function GetPost(uint256 PostID) view returns(address postaddress)
func (_ERAccount *ERAccountCallerSession) GetPost(PostID *big.Int) (common.Address, error) {
	return _ERAccount.Contract.GetPost(&_ERAccount.CallOpts, PostID)
}

// Karma is a free data retrieval call binding the contract method 0xe5163de5.
//
// Solidity: function Karma() view returns(int16)
func (_ERAccount *ERAccountCaller) Karma(opts *bind.CallOpts) (int16, error) {
	var out []interface{}
	err := _ERAccount.contract.Call(opts, &out, "Karma")

	if err != nil {
		return *new(int16), err
	}

	out0 := *abi.ConvertType(out[0], new(int16)).(*int16)

	return out0, err

}

// Karma is a free data retrieval call binding the contract method 0xe5163de5.
//
// Solidity: function Karma() view returns(int16)
func (_ERAccount *ERAccountSession) Karma() (int16, error) {
	return _ERAccount.Contract.Karma(&_ERAccount.CallOpts)
}

// Karma is a free data retrieval call binding the contract method 0xe5163de5.
//
// Solidity: function Karma() view returns(int16)
func (_ERAccount *ERAccountCallerSession) Karma() (int16, error) {
	return _ERAccount.Contract.Karma(&_ERAccount.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x8052474d.
//
// Solidity: function Name() view returns(string)
func (_ERAccount *ERAccountCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERAccount.contract.Call(opts, &out, "Name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x8052474d.
//
// Solidity: function Name() view returns(string)
func (_ERAccount *ERAccountSession) Name() (string, error) {
	return _ERAccount.Contract.Name(&_ERAccount.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x8052474d.
//
// Solidity: function Name() view returns(string)
func (_ERAccount *ERAccountCallerSession) Name() (string, error) {
	return _ERAccount.Contract.Name(&_ERAccount.CallOpts)
}

// Posts is a free data retrieval call binding the contract method 0x4fafef1a.
//
// Solidity: function Posts(uint256 ) view returns(address)
func (_ERAccount *ERAccountCaller) Posts(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ERAccount.contract.Call(opts, &out, "Posts", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Posts is a free data retrieval call binding the contract method 0x4fafef1a.
//
// Solidity: function Posts(uint256 ) view returns(address)
func (_ERAccount *ERAccountSession) Posts(arg0 *big.Int) (common.Address, error) {
	return _ERAccount.Contract.Posts(&_ERAccount.CallOpts, arg0)
}

// Posts is a free data retrieval call binding the contract method 0x4fafef1a.
//
// Solidity: function Posts(uint256 ) view returns(address)
func (_ERAccount *ERAccountCallerSession) Posts(arg0 *big.Int) (common.Address, error) {
	return _ERAccount.Contract.Posts(&_ERAccount.CallOpts, arg0)
}

// DeletePost is a paid mutator transaction binding the contract method 0x75c16cb2.
//
// Solidity: function DeletePost(uint256 PostID) returns()
func (_ERAccount *ERAccountTransactor) DeletePost(opts *bind.TransactOpts, PostID *big.Int) (*types.Transaction, error) {
	return _ERAccount.contract.Transact(opts, "DeletePost", PostID)
}

// DeletePost is a paid mutator transaction binding the contract method 0x75c16cb2.
//
// Solidity: function DeletePost(uint256 PostID) returns()
func (_ERAccount *ERAccountSession) DeletePost(PostID *big.Int) (*types.Transaction, error) {
	return _ERAccount.Contract.DeletePost(&_ERAccount.TransactOpts, PostID)
}

// DeletePost is a paid mutator transaction binding the contract method 0x75c16cb2.
//
// Solidity: function DeletePost(uint256 PostID) returns()
func (_ERAccount *ERAccountTransactorSession) DeletePost(PostID *big.Int) (*types.Transaction, error) {
	return _ERAccount.Contract.DeletePost(&_ERAccount.TransactOpts, PostID)
}

// PostMessage is a paid mutator transaction binding the contract method 0x01465d36.
//
// Solidity: function PostMessage(string PostName_, string PostString_) returns(address postaddress)
func (_ERAccount *ERAccountTransactor) PostMessage(opts *bind.TransactOpts, PostName_ string, PostString_ string) (*types.Transaction, error) {
	return _ERAccount.contract.Transact(opts, "PostMessage", PostName_, PostString_)
}

// PostMessage is a paid mutator transaction binding the contract method 0x01465d36.
//
// Solidity: function PostMessage(string PostName_, string PostString_) returns(address postaddress)
func (_ERAccount *ERAccountSession) PostMessage(PostName_ string, PostString_ string) (*types.Transaction, error) {
	return _ERAccount.Contract.PostMessage(&_ERAccount.TransactOpts, PostName_, PostString_)
}

// PostMessage is a paid mutator transaction binding the contract method 0x01465d36.
//
// Solidity: function PostMessage(string PostName_, string PostString_) returns(address postaddress)
func (_ERAccount *ERAccountTransactorSession) PostMessage(PostName_ string, PostString_ string) (*types.Transaction, error) {
	return _ERAccount.Contract.PostMessage(&_ERAccount.TransactOpts, PostName_, PostString_)
}

// AdminDeleteAccount is a paid mutator transaction binding the contract method 0x3e450fff.
//
// Solidity: function adminDeleteAccount() payable returns()
func (_ERAccount *ERAccountTransactor) AdminDeleteAccount(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERAccount.contract.Transact(opts, "adminDeleteAccount")
}

// AdminDeleteAccount is a paid mutator transaction binding the contract method 0x3e450fff.
//
// Solidity: function adminDeleteAccount() payable returns()
func (_ERAccount *ERAccountSession) AdminDeleteAccount() (*types.Transaction, error) {
	return _ERAccount.Contract.AdminDeleteAccount(&_ERAccount.TransactOpts)
}

// AdminDeleteAccount is a paid mutator transaction binding the contract method 0x3e450fff.
//
// Solidity: function adminDeleteAccount() payable returns()
func (_ERAccount *ERAccountTransactorSession) AdminDeleteAccount() (*types.Transaction, error) {
	return _ERAccount.Contract.AdminDeleteAccount(&_ERAccount.TransactOpts)
}

// PostMetaData contains all meta data concerning the Post contract.
var PostMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"PostName_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"PostString_\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"CountPosts\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int16\",\"name\":\"Karma\",\"type\":\"int16\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"Post\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"PostName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"PostString\",\"type\":\"string\"}],\"name\":\"PostBanner\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"CommentString\",\"type\":\"string\"}],\"name\":\"AddComment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"Comments\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"CommentNumber\",\"type\":\"uint256\"},{\"internalType\":\"int16\",\"name\":\"Karma\",\"type\":\"int16\"},{\"internalType\":\"string\",\"name\":\"CommentString\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"Poster\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CountComments\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DeletePost\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Karma\",\"outputs\":[{\"internalType\":\"int16\",\"name\":\"\",\"type\":\"int16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PostName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PostNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PostString\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Transfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int8\",\"name\":\"vote\",\"type\":\"int8\"}],\"name\":\"Vote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int8\",\"name\":\"vote\",\"type\":\"int8\"},{\"internalType\":\"uint256\",\"name\":\"CommentNumber\",\"type\":\"uint256\"}],\"name\":\"Vote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"6bde24b6": "AddComment(string)",
		"8ea20ddc": "Comments(uint256)",
		"dc28c100": "CountComments()",
		"85141ee3": "DeletePost()",
		"e5163de5": "Karma()",
		"680e8ebc": "PostName()",
		"98c024bc": "PostNumber()",
		"fe6c9752": "PostString()",
		"406dade3": "Transfer()",
		"4e49d2da": "Vote(int8)",
		"ae81085a": "Vote(int8,uint256)",
	},
	Bin: "0x60806040523480156200001157600080fd5b5060405162000f7b38038062000f7b833981016040819052620000349162000241565b4260005582516200004d906004906020860190620000ce565b50815162000063906005906020850190620000ce565b506001818155600680546001600160a01b031916321790556000546002546040517fe7af46f298185143ea7dc9c524c80748fc11ca4248529d56a451f895aa5a7c8093620000bd9392900b90309060049060059062000399565b60405180910390a1505050620003eb565b828054620000dc90620002b4565b90600052602060002090601f0160209004810192826200010057600085556200014b565b82601f106200011b57805160ff19168380011785556200014b565b828001600101855582156200014b579182015b828111156200014b5782518255916020019190600101906200012e565b50620001599291506200015d565b5090565b5b808211156200015957600081556001016200015e565b634e487b7160e01b600052604160045260246000fd5b600082601f8301126200019c57600080fd5b81516001600160401b0380821115620001b957620001b962000174565b604051601f8301601f19908116603f01168101908282118183101715620001e457620001e462000174565b816040528381526020925086838588010111156200020157600080fd5b600091505b8382101562000225578582018301518183018401529082019062000206565b83821115620002375760008385830101525b9695505050505050565b6000806000606084860312156200025757600080fd5b83516001600160401b03808211156200026f57600080fd5b6200027d878388016200018a565b945060208601519150808211156200029457600080fd5b50620002a3868287016200018a565b925050604084015190509250925092565b600181811c90821680620002c957607f821691505b60208210811415620002eb57634e487b7160e01b600052602260045260246000fd5b50919050565b8054600090600181811c90808316806200030c57607f831692505b60208084108214156200032f57634e487b7160e01b600052602260045260246000fd5b838852602088018280156200034d57600181146200035f576200038c565b60ff198716825282820197506200038c565b60008981526020902060005b8781101562000386578154848201529086019084016200036b565b83019850505b5050505050505092915050565b8581528460010b602082015260018060a01b038416604082015260a060608201526000620003cb60a0830185620002f1565b8281036080840152620003df8185620002f1565b98975050505050505050565b610b8080620003fb6000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c80638ea20ddc116100715780638ea20ddc1461010457806398c024bc14610128578063ae81085a1461013f578063dc28c10014610152578063e5163de51461015b578063fe6c97521461017b57600080fd5b8063406dade3146100ae5780634e49d2da146100b8578063680e8ebc146100cb5780636bde24b6146100e957806385141ee3146100fc575b600080fd5b6100b6610183565b005b6100b66100c6366004610797565b6101d6565b6100d361034a565b6040516100e09190610806565b60405180910390f35b6100b66100f736600461082f565b6103d8565b6100b661048a565b6101176101123660046108e0565b6104af565b6040516100e09594939291906108f9565b61013160015481565b6040519081526020016100e0565b6100b661014d36600461093b565b610575565b61013160075481565b6002546101689060010b81565b60405160019190910b81526020016100e0565b6100d36106da565b6006546001600160a01b0316331461019a57600080fd5b6006546040516001600160a01b03909116904780156108fc02916000818181858888f193505050501580156101d3573d6000803e3d6000fd5b50565b60008160000b136101fb5760008160000b126101f35760006101fe565b6000196101fe565b60015b3360009081526003602052604081205491925082810b91900b14156102625760405162461bcd60e51b81526020600482015260156024820152744d757374206368616e6765207072657620766f746560581b60448201526064015b60405180910390fd5b336000908152600360205260408120546002805491830b92909161028a90849060010b61097b565b825461ffff9182166101009390930a92830291909202199091161790555060028054600083810b92916102c190849060010b6109c1565b82546101009290920a61ffff8181021990931691909216919091021790555033600090815260036020526040808220805460ff191660ff8516179055905460025491517fe7af46f298185143ea7dc9c524c80748fc11ca4248529d56a451f895aa5a7c809261033f929160019190910b903090600490600590610ae1565b60405180910390a150565b6004805461035790610a06565b80601f016020809104026020016040519081016040528092919081815260200182805461038390610a06565b80156103d05780601f106103a5576101008083540402835291602001916103d0565b820191906000526020600020905b8154815290600101906020018083116103b357829003601f168201915b505050505081565b60fa8151106104295760405162461bcd60e51b815260206004820152601c60248201527f506f7374206578636565647320636861726163746572206c696d6974000000006044820152606401610259565b60078054600090815260086020908152604090912042815591546001830155825161045c916004840191908501906106e7565b506005810180546001600160a01b031916331790556007805490600061048183610b2f565b91905055505050565b6006546001600160a01b031633146104a157600080fd5b6006546001600160a01b0316ff5b600860205260009081526040902080546001808301546002840154600485018054949592949190930b92906104e390610a06565b80601f016020809104026020016040519081016040528092919081815260200182805461050f90610a06565b801561055c5780601f106105315761010080835404028352916020019161055c565b820191906000526020600020905b81548152906001019060200180831161053f57829003601f168201915b505050600590930154919250506001600160a01b031685565b60008260000b1361059a5760008260000b1261059257600061059d565b60001961059d565b60015b600082815260086020908152604080832033845260030190915281205491935083810b91900b14156106095760405162461bcd60e51b81526020600482015260156024820152744d757374206368616e6765207072657620766f746560581b6044820152606401610259565b60008181526008602081815260408084203385526003810183529084205485855292909152600201805491830b92909161064790849060010b61097b565b825461ffff9182166101009390930a9283029190920219909116179055506000818152600860205260408120600201805484830b929061068b90849060010b6109c1565b825461ffff9182166101009390930a92830291909202199091161790555060009081526008602090815260408083203384526003019091529020805460ff90921660ff19909216919091179055565b6005805461035790610a06565b8280546106f390610a06565b90600052602060002090601f016020900481019282610715576000855561075b565b82601f1061072e57805160ff191683800117855561075b565b8280016001018555821561075b579182015b8281111561075b578251825591602001919060010190610740565b5061076792915061076b565b5090565b5b80821115610767576000815560010161076c565b8035600081900b811461079257600080fd5b919050565b6000602082840312156107a957600080fd5b6107b282610780565b9392505050565b6000815180845260005b818110156107df576020818501810151868301820152016107c3565b818111156107f1576000602083870101525b50601f01601f19169290920160200192915050565b6020815260006107b260208301846107b9565b634e487b7160e01b600052604160045260246000fd5b60006020828403121561084157600080fd5b813567ffffffffffffffff8082111561085957600080fd5b818401915084601f83011261086d57600080fd5b81358181111561087f5761087f610819565b604051601f8201601f19908116603f011681019083821181831017156108a7576108a7610819565b816040528281528760208487010111156108c057600080fd5b826020860160208301376000928101602001929092525095945050505050565b6000602082840312156108f257600080fd5b5035919050565b8581528460208201528360010b604082015260a06060820152600061092160a08301856107b9565b905060018060a01b03831660808301529695505050505050565b6000806040838503121561094e57600080fd5b61095783610780565b946020939093013593505050565b634e487b7160e01b600052601160045260246000fd5b60008160010b8360010b6000811281617fff19018312811516156109a1576109a1610965565b81617fff0183138116156109b7576109b7610965565b5090039392505050565b60008160010b8360010b6000821282617fff038213811516156109e6576109e6610965565b82617fff190382128116156109fd576109fd610965565b50019392505050565b600181811c90821680610a1a57607f821691505b60208210811415610a3b57634e487b7160e01b600052602260045260246000fd5b50919050565b8054600090600181811c9080831680610a5b57607f831692505b6020808410821415610a7d57634e487b7160e01b600052602260045260246000fd5b83885260208801828015610a985760018114610aa957610ad4565b60ff19871682528282019750610ad4565b60008981526020902060005b87811015610ace57815484820152908601908401610ab5565b83019850505b5050505050505092915050565b8581528460010b602082015260018060a01b038416604082015260a060608201526000610b1160a0830185610a41565b8281036080840152610b238185610a41565b98975050505050505050565b6000600019821415610b4357610b43610965565b506001019056fea2646970667358221220644f09dd28e689894945015501237899179602563292f57f2ed9f36ae94dafd464736f6c63430008090033",
}

// PostABI is the input ABI used to generate the binding from.
// Deprecated: Use PostMetaData.ABI instead.
var PostABI = PostMetaData.ABI

// Deprecated: Use PostMetaData.Sigs instead.
// PostFuncSigs maps the 4-byte function signature to its string representation.
var PostFuncSigs = PostMetaData.Sigs

// PostBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PostMetaData.Bin instead.
var PostBin = PostMetaData.Bin

// DeployPost deploys a new Ethereum contract, binding an instance of Post to it.
func DeployPost(auth *bind.TransactOpts, backend bind.ContractBackend, PostName_ string, PostString_ string, CountPosts *big.Int) (common.Address, *types.Transaction, *Post, error) {
	parsed, err := PostMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PostBin), backend, PostName_, PostString_, CountPosts)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Post{PostCaller: PostCaller{contract: contract}, PostTransactor: PostTransactor{contract: contract}, PostFilterer: PostFilterer{contract: contract}}, nil
}

// Post is an auto generated Go binding around an Ethereum contract.
type Post struct {
	PostCaller     // Read-only binding to the contract
	PostTransactor // Write-only binding to the contract
	PostFilterer   // Log filterer for contract events
}

// PostCaller is an auto generated read-only Go binding around an Ethereum contract.
type PostCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PostTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PostTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PostFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PostFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PostSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PostSession struct {
	Contract     *Post             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PostCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PostCallerSession struct {
	Contract *PostCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// PostTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PostTransactorSession struct {
	Contract     *PostTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PostRaw is an auto generated low-level Go binding around an Ethereum contract.
type PostRaw struct {
	Contract *Post // Generic contract binding to access the raw methods on
}

// PostCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PostCallerRaw struct {
	Contract *PostCaller // Generic read-only contract binding to access the raw methods on
}

// PostTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PostTransactorRaw struct {
	Contract *PostTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPost creates a new instance of Post, bound to a specific deployed contract.
func NewPost(address common.Address, backend bind.ContractBackend) (*Post, error) {
	contract, err := bindPost(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Post{PostCaller: PostCaller{contract: contract}, PostTransactor: PostTransactor{contract: contract}, PostFilterer: PostFilterer{contract: contract}}, nil
}

// NewPostCaller creates a new read-only instance of Post, bound to a specific deployed contract.
func NewPostCaller(address common.Address, caller bind.ContractCaller) (*PostCaller, error) {
	contract, err := bindPost(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PostCaller{contract: contract}, nil
}

// NewPostTransactor creates a new write-only instance of Post, bound to a specific deployed contract.
func NewPostTransactor(address common.Address, transactor bind.ContractTransactor) (*PostTransactor, error) {
	contract, err := bindPost(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PostTransactor{contract: contract}, nil
}

// NewPostFilterer creates a new log filterer instance of Post, bound to a specific deployed contract.
func NewPostFilterer(address common.Address, filterer bind.ContractFilterer) (*PostFilterer, error) {
	contract, err := bindPost(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PostFilterer{contract: contract}, nil
}

// bindPost binds a generic wrapper to an already deployed contract.
func bindPost(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PostABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Post *PostRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Post.Contract.PostCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Post *PostRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Post.Contract.PostTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Post *PostRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Post.Contract.PostTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Post *PostCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Post.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Post *PostTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Post.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Post *PostTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Post.Contract.contract.Transact(opts, method, params...)
}

// Comments is a free data retrieval call binding the contract method 0x8ea20ddc.
//
// Solidity: function Comments(uint256 ) view returns(uint256 timestamp, uint256 CommentNumber, int16 Karma, string CommentString, address Poster)
func (_Post *PostCaller) Comments(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Timestamp     *big.Int
	CommentNumber *big.Int
	Karma         int16
	CommentString string
	Poster        common.Address
}, error) {
	var out []interface{}
	err := _Post.contract.Call(opts, &out, "Comments", arg0)

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
func (_Post *PostSession) Comments(arg0 *big.Int) (struct {
	Timestamp     *big.Int
	CommentNumber *big.Int
	Karma         int16
	CommentString string
	Poster        common.Address
}, error) {
	return _Post.Contract.Comments(&_Post.CallOpts, arg0)
}

// Comments is a free data retrieval call binding the contract method 0x8ea20ddc.
//
// Solidity: function Comments(uint256 ) view returns(uint256 timestamp, uint256 CommentNumber, int16 Karma, string CommentString, address Poster)
func (_Post *PostCallerSession) Comments(arg0 *big.Int) (struct {
	Timestamp     *big.Int
	CommentNumber *big.Int
	Karma         int16
	CommentString string
	Poster        common.Address
}, error) {
	return _Post.Contract.Comments(&_Post.CallOpts, arg0)
}

// CountComments is a free data retrieval call binding the contract method 0xdc28c100.
//
// Solidity: function CountComments() view returns(uint256)
func (_Post *PostCaller) CountComments(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Post.contract.Call(opts, &out, "CountComments")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CountComments is a free data retrieval call binding the contract method 0xdc28c100.
//
// Solidity: function CountComments() view returns(uint256)
func (_Post *PostSession) CountComments() (*big.Int, error) {
	return _Post.Contract.CountComments(&_Post.CallOpts)
}

// CountComments is a free data retrieval call binding the contract method 0xdc28c100.
//
// Solidity: function CountComments() view returns(uint256)
func (_Post *PostCallerSession) CountComments() (*big.Int, error) {
	return _Post.Contract.CountComments(&_Post.CallOpts)
}

// Karma is a free data retrieval call binding the contract method 0xe5163de5.
//
// Solidity: function Karma() view returns(int16)
func (_Post *PostCaller) Karma(opts *bind.CallOpts) (int16, error) {
	var out []interface{}
	err := _Post.contract.Call(opts, &out, "Karma")

	if err != nil {
		return *new(int16), err
	}

	out0 := *abi.ConvertType(out[0], new(int16)).(*int16)

	return out0, err

}

// Karma is a free data retrieval call binding the contract method 0xe5163de5.
//
// Solidity: function Karma() view returns(int16)
func (_Post *PostSession) Karma() (int16, error) {
	return _Post.Contract.Karma(&_Post.CallOpts)
}

// Karma is a free data retrieval call binding the contract method 0xe5163de5.
//
// Solidity: function Karma() view returns(int16)
func (_Post *PostCallerSession) Karma() (int16, error) {
	return _Post.Contract.Karma(&_Post.CallOpts)
}

// PostName is a free data retrieval call binding the contract method 0x680e8ebc.
//
// Solidity: function PostName() view returns(string)
func (_Post *PostCaller) PostName(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Post.contract.Call(opts, &out, "PostName")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// PostName is a free data retrieval call binding the contract method 0x680e8ebc.
//
// Solidity: function PostName() view returns(string)
func (_Post *PostSession) PostName() (string, error) {
	return _Post.Contract.PostName(&_Post.CallOpts)
}

// PostName is a free data retrieval call binding the contract method 0x680e8ebc.
//
// Solidity: function PostName() view returns(string)
func (_Post *PostCallerSession) PostName() (string, error) {
	return _Post.Contract.PostName(&_Post.CallOpts)
}

// PostNumber is a free data retrieval call binding the contract method 0x98c024bc.
//
// Solidity: function PostNumber() view returns(uint256)
func (_Post *PostCaller) PostNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Post.contract.Call(opts, &out, "PostNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PostNumber is a free data retrieval call binding the contract method 0x98c024bc.
//
// Solidity: function PostNumber() view returns(uint256)
func (_Post *PostSession) PostNumber() (*big.Int, error) {
	return _Post.Contract.PostNumber(&_Post.CallOpts)
}

// PostNumber is a free data retrieval call binding the contract method 0x98c024bc.
//
// Solidity: function PostNumber() view returns(uint256)
func (_Post *PostCallerSession) PostNumber() (*big.Int, error) {
	return _Post.Contract.PostNumber(&_Post.CallOpts)
}

// PostString is a free data retrieval call binding the contract method 0xfe6c9752.
//
// Solidity: function PostString() view returns(string)
func (_Post *PostCaller) PostString(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Post.contract.Call(opts, &out, "PostString")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// PostString is a free data retrieval call binding the contract method 0xfe6c9752.
//
// Solidity: function PostString() view returns(string)
func (_Post *PostSession) PostString() (string, error) {
	return _Post.Contract.PostString(&_Post.CallOpts)
}

// PostString is a free data retrieval call binding the contract method 0xfe6c9752.
//
// Solidity: function PostString() view returns(string)
func (_Post *PostCallerSession) PostString() (string, error) {
	return _Post.Contract.PostString(&_Post.CallOpts)
}

// AddComment is a paid mutator transaction binding the contract method 0x6bde24b6.
//
// Solidity: function AddComment(string CommentString) returns()
func (_Post *PostTransactor) AddComment(opts *bind.TransactOpts, CommentString string) (*types.Transaction, error) {
	return _Post.contract.Transact(opts, "AddComment", CommentString)
}

// AddComment is a paid mutator transaction binding the contract method 0x6bde24b6.
//
// Solidity: function AddComment(string CommentString) returns()
func (_Post *PostSession) AddComment(CommentString string) (*types.Transaction, error) {
	return _Post.Contract.AddComment(&_Post.TransactOpts, CommentString)
}

// AddComment is a paid mutator transaction binding the contract method 0x6bde24b6.
//
// Solidity: function AddComment(string CommentString) returns()
func (_Post *PostTransactorSession) AddComment(CommentString string) (*types.Transaction, error) {
	return _Post.Contract.AddComment(&_Post.TransactOpts, CommentString)
}

// DeletePost is a paid mutator transaction binding the contract method 0x85141ee3.
//
// Solidity: function DeletePost() returns()
func (_Post *PostTransactor) DeletePost(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Post.contract.Transact(opts, "DeletePost")
}

// DeletePost is a paid mutator transaction binding the contract method 0x85141ee3.
//
// Solidity: function DeletePost() returns()
func (_Post *PostSession) DeletePost() (*types.Transaction, error) {
	return _Post.Contract.DeletePost(&_Post.TransactOpts)
}

// DeletePost is a paid mutator transaction binding the contract method 0x85141ee3.
//
// Solidity: function DeletePost() returns()
func (_Post *PostTransactorSession) DeletePost() (*types.Transaction, error) {
	return _Post.Contract.DeletePost(&_Post.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0x406dade3.
//
// Solidity: function Transfer() returns()
func (_Post *PostTransactor) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Post.contract.Transact(opts, "Transfer")
}

// Transfer is a paid mutator transaction binding the contract method 0x406dade3.
//
// Solidity: function Transfer() returns()
func (_Post *PostSession) Transfer() (*types.Transaction, error) {
	return _Post.Contract.Transfer(&_Post.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0x406dade3.
//
// Solidity: function Transfer() returns()
func (_Post *PostTransactorSession) Transfer() (*types.Transaction, error) {
	return _Post.Contract.Transfer(&_Post.TransactOpts)
}

// Vote is a paid mutator transaction binding the contract method 0x4e49d2da.
//
// Solidity: function Vote(int8 vote) returns()
func (_Post *PostTransactor) Vote(opts *bind.TransactOpts, vote int8) (*types.Transaction, error) {
	return _Post.contract.Transact(opts, "Vote", vote)
}

// Vote is a paid mutator transaction binding the contract method 0x4e49d2da.
//
// Solidity: function Vote(int8 vote) returns()
func (_Post *PostSession) Vote(vote int8) (*types.Transaction, error) {
	return _Post.Contract.Vote(&_Post.TransactOpts, vote)
}

// Vote is a paid mutator transaction binding the contract method 0x4e49d2da.
//
// Solidity: function Vote(int8 vote) returns()
func (_Post *PostTransactorSession) Vote(vote int8) (*types.Transaction, error) {
	return _Post.Contract.Vote(&_Post.TransactOpts, vote)
}

// Vote0 is a paid mutator transaction binding the contract method 0xae81085a.
//
// Solidity: function Vote(int8 vote, uint256 CommentNumber) returns()
func (_Post *PostTransactor) Vote0(opts *bind.TransactOpts, vote int8, CommentNumber *big.Int) (*types.Transaction, error) {
	return _Post.contract.Transact(opts, "Vote0", vote, CommentNumber)
}

// Vote0 is a paid mutator transaction binding the contract method 0xae81085a.
//
// Solidity: function Vote(int8 vote, uint256 CommentNumber) returns()
func (_Post *PostSession) Vote0(vote int8, CommentNumber *big.Int) (*types.Transaction, error) {
	return _Post.Contract.Vote0(&_Post.TransactOpts, vote, CommentNumber)
}

// Vote0 is a paid mutator transaction binding the contract method 0xae81085a.
//
// Solidity: function Vote(int8 vote, uint256 CommentNumber) returns()
func (_Post *PostTransactorSession) Vote0(vote int8, CommentNumber *big.Int) (*types.Transaction, error) {
	return _Post.Contract.Vote0(&_Post.TransactOpts, vote, CommentNumber)
}

// PostPostBannerIterator is returned from FilterPostBanner and is used to iterate over the raw logs and unpacked data for PostBanner events raised by the Post contract.
type PostPostBannerIterator struct {
	Event *PostPostBanner // Event containing the contract specifics and raw log

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
func (it *PostPostBannerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PostPostBanner)
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
		it.Event = new(PostPostBanner)
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
func (it *PostPostBannerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PostPostBannerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PostPostBanner represents a PostBanner event raised by the Post contract.
type PostPostBanner struct {
	Timestamp  *big.Int
	Karma      int16
	Post       common.Address
	PostName   string
	PostString string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPostBanner is a free log retrieval operation binding the contract event 0xe7af46f298185143ea7dc9c524c80748fc11ca4248529d56a451f895aa5a7c80.
//
// Solidity: event PostBanner(uint256 timestamp, int16 Karma, address Post, string PostName, string PostString)
func (_Post *PostFilterer) FilterPostBanner(opts *bind.FilterOpts) (*PostPostBannerIterator, error) {

	logs, sub, err := _Post.contract.FilterLogs(opts, "PostBanner")
	if err != nil {
		return nil, err
	}
	return &PostPostBannerIterator{contract: _Post.contract, event: "PostBanner", logs: logs, sub: sub}, nil
}

// WatchPostBanner is a free log subscription operation binding the contract event 0xe7af46f298185143ea7dc9c524c80748fc11ca4248529d56a451f895aa5a7c80.
//
// Solidity: event PostBanner(uint256 timestamp, int16 Karma, address Post, string PostName, string PostString)
func (_Post *PostFilterer) WatchPostBanner(opts *bind.WatchOpts, sink chan<- *PostPostBanner) (event.Subscription, error) {

	logs, sub, err := _Post.contract.WatchLogs(opts, "PostBanner")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PostPostBanner)
				if err := _Post.contract.UnpackLog(event, "PostBanner", log); err != nil {
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

// ParsePostBanner is a log parse operation binding the contract event 0xe7af46f298185143ea7dc9c524c80748fc11ca4248529d56a451f895aa5a7c80.
//
// Solidity: event PostBanner(uint256 timestamp, int16 Karma, address Post, string PostName, string PostString)
func (_Post *PostFilterer) ParsePostBanner(log types.Log) (*PostPostBanner, error) {
	event := new(PostPostBanner)
	if err := _Post.contract.UnpackLog(event, "PostBanner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
