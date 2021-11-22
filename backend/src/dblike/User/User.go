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
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"Name_\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"Bio\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CountPosts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"PostID\",\"type\":\"uint256\"}],\"name\":\"DeletePost\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetOwnerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"adminAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"PostID\",\"type\":\"uint256\"}],\"name\":\"GetPost\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"postaddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Karma\",\"outputs\":[{\"internalType\":\"int16\",\"name\":\"\",\"type\":\"int16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"PostName_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"PostString_\",\"type\":\"string\"}],\"name\":\"PostMessage\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"postaddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"Posts\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"Bio_\",\"type\":\"string\"}],\"name\":\"SetBio\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"adminDeleteAccount\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"c2d2d0f9": "Bio()",
		"27677343": "CountPosts()",
		"75c16cb2": "DeletePost(uint256)",
		"e6f7bf89": "GetOwnerAddress()",
		"bf421138": "GetPost(uint256)",
		"e5163de5": "Karma()",
		"8052474d": "Name()",
		"01465d36": "PostMessage(string,string)",
		"4fafef1a": "Posts(uint256)",
		"27d0b9ec": "SetBio(string)",
		"3e450fff": "adminDeleteAccount()",
	},
	Bin: "0x60806040523480156200001157600080fd5b506040516200194138038062001941833981016040819052620000349162000124565b80516200004990600090602084019062000068565b50506000600455600580546001600160a01b031916321790556200023d565b828054620000769062000200565b90600052602060002090601f0160209004810192826200009a5760008555620000e5565b82601f10620000b557805160ff1916838001178555620000e5565b82800160010185558215620000e5579182015b82811115620000e5578251825591602001919060010190620000c8565b50620000f3929150620000f7565b5090565b5b80821115620000f35760008155600101620000f8565b634e487b7160e01b600052604160045260246000fd5b600060208083850312156200013857600080fd5b82516001600160401b03808211156200015057600080fd5b818501915085601f8301126200016557600080fd5b8151818111156200017a576200017a6200010e565b604051601f8201601f19908116603f01168101908382118183101715620001a557620001a56200010e565b816040528281528886848701011115620001be57600080fd5b600093505b82841015620001e25784840186015181850187015292850192620001c3565b82841115620001f45760008684830101525b98975050505050505050565b600181811c908216806200021557607f821691505b602082108114156200023757634e487b7160e01b600052602260045260246000fd5b50919050565b6116f4806200024d6000396000f3fe608060405260043610620000a95760003560e01c806375c16cb2116200006c57806375c16cb2146200016d5780638052474d1462000192578063bf42113814620001b9578063c2d2d0f914620001de578063e5163de514620001f6578063e6f7bf89146200022657600080fd5b806301465d3614620000ae5780632767734314620000f057806327d0b9ec14620001175780633e450fff146200013e5780634fafef1a1462000148575b600080fd5b348015620000bb57600080fd5b50620000d3620000cd366004620006c7565b62000246565b6040516001600160a01b0390911681526020015b60405180910390f35b348015620000fd57600080fd5b506200010860045481565b604051908152602001620000e7565b3480156200012457600080fd5b506200013c6200013636600462000732565b6200034c565b005b6200013c6200037d565b3480156200015557600080fd5b50620000d36200016736600462000773565b620003a3565b3480156200017a57600080fd5b506200013c6200018c36600462000773565b620003ce565b3480156200019f57600080fd5b50620001aa62000490565b604051620000e79190620007dd565b348015620001c657600080fd5b50620000d3620001d836600462000773565b62000526565b348015620001eb57600080fd5b50620001aa62000559565b3480156200020357600080fd5b50600354620002129060010b81565b60405160019190910b8152602001620000e7565b3480156200023357600080fd5b506005546001600160a01b0316620000d3565b6005546000906001600160a01b031633146200026157600080fd5b6101f4825110620002b85760405162461bcd60e51b815260206004820152601c60248201527f506f7374206578636565647320636861726163746572206c696d697400000000604482015260640160405180910390fd5b60008383600454604051620002cd9062000568565b620002db93929190620007f9565b604051809103906000f080158015620002f8573d6000803e3d6000fd5b50600280546001810182556000919091527f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace0180546001600160a01b0319166001600160a01b038316179055949350505050565b6005546001600160a01b031633146200036457600080fd5b80516200037990600190602084019062000576565b5050565b6005546001600160a01b031633146200039557600080fd5b6005546001600160a01b0316ff5b60028181548110620003b457600080fd5b6000918252602090912001546001600160a01b0316905081565b6005546001600160a01b03163314620003e657600080fd5b60028181548110620003fc57620003fc62000833565b6000918252602082200154604080516385141ee360e01b815290516001600160a01b03909216926385141ee39260048084019382900301818387803b1580156200044557600080fd5b505af11580156200045a573d6000803e3d6000fd5b505050506002818154811062000474576200047462000833565b600091825260209091200180546001600160a01b031916905550565b600080546200049f9062000849565b80601f0160208091040260200160405190810160405280929190818152602001828054620004cd9062000849565b80156200051e5780601f10620004f2576101008083540402835291602001916200051e565b820191906000526020600020905b8154815290600101906020018083116200050057829003601f168201915b505050505081565b6000600282815481106200053e576200053e62000833565b6000918252602090912001546001600160a01b031692915050565b600180546200049f9062000849565b610e38806200088783390190565b828054620005849062000849565b90600052602060002090601f016020900481019282620005a85760008555620005f3565b82601f10620005c357805160ff1916838001178555620005f3565b82800160010185558215620005f3579182015b82811115620005f3578251825591602001919060010190620005d6565b506200060192915062000605565b5090565b5b8082111562000601576000815560010162000606565b634e487b7160e01b600052604160045260246000fd5b600082601f8301126200064457600080fd5b813567ffffffffffffffff808211156200066257620006626200061c565b604051601f8301601f19908116603f011681019082821181831017156200068d576200068d6200061c565b81604052838152866020858801011115620006a757600080fd5b836020870160208301376000602085830101528094505050505092915050565b60008060408385031215620006db57600080fd5b823567ffffffffffffffff80821115620006f457600080fd5b620007028683870162000632565b935060208501359150808211156200071957600080fd5b50620007288582860162000632565b9150509250929050565b6000602082840312156200074557600080fd5b813567ffffffffffffffff8111156200075d57600080fd5b6200076b8482850162000632565b949350505050565b6000602082840312156200078657600080fd5b5035919050565b6000815180845260005b81811015620007b55760208185018101518683018201520162000797565b81811115620007c8576000602083870101525b50601f01601f19169290920160200192915050565b602081526000620007f260208301846200078d565b9392505050565b6060815260006200080e60608301866200078d565b82810360208401526200082281866200078d565b915050826040830152949350505050565b634e487b7160e01b600052603260045260246000fd5b600181811c908216806200085e57607f821691505b602082108114156200088057634e487b7160e01b600052602260045260246000fd5b5091905056fe60806040523480156200001157600080fd5b5060405162000e3838038062000e3883398101604081905262000034916200023a565b4260005582516200004d906004906020860190620000c7565b50815162000063906005906020850190620000c7565b506001819055600680546001600160a01b031916321790556000546040517f7307c220a36817a866099ba844095153e39ee120706b0822e482241b2091ac5e91620000b691309060049060059062000392565b60405180910390a1505050620003dd565b828054620000d590620002ad565b90600052602060002090601f016020900481019282620000f9576000855562000144565b82601f106200011457805160ff191683800117855562000144565b8280016001018555821562000144579182015b828111156200014457825182559160200191906001019062000127565b506200015292915062000156565b5090565b5b8082111562000152576000815560010162000157565b634e487b7160e01b600052604160045260246000fd5b600082601f8301126200019557600080fd5b81516001600160401b0380821115620001b257620001b26200016d565b604051601f8301601f19908116603f01168101908282118183101715620001dd57620001dd6200016d565b81604052838152602092508683858801011115620001fa57600080fd5b600091505b838210156200021e5785820183015181830184015290820190620001ff565b83821115620002305760008385830101525b9695505050505050565b6000806000606084860312156200025057600080fd5b83516001600160401b03808211156200026857600080fd5b620002768783880162000183565b945060208601519150808211156200028d57600080fd5b506200029c8682870162000183565b925050604084015190509250925092565b600181811c90821680620002c257607f821691505b60208210811415620002e457634e487b7160e01b600052602260045260246000fd5b50919050565b8054600090600181811c90808316806200030557607f831692505b60208084108214156200032857634e487b7160e01b600052602260045260246000fd5b83885260208801828015620003465760018114620003585762000385565b60ff1987168252828201975062000385565b60008981526020902060005b878110156200037f5781548482015290860190840162000364565b83019850505b5050505050505092915050565b8481526001600160a01b0384166020820152608060408201819052600090620003be90830185620002ea565b8281036060840152620003d28185620002ea565b979650505050505050565b610a4b80620003ed6000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c806385141ee31161007157806385141ee3146101045780638ea20ddc1461010c57806398c024bc14610130578063ae81085a14610147578063dc28c1001461015a578063e5163de51461016357600080fd5b80631b529115146100ae578063406dade3146100cc5780634e49d2da146100d6578063680e8ebc146100e95780636bde24b6146100f1575b600080fd5b6100b6610183565b6040516100c39190610786565b60405180910390f35b6100d4610211565b005b6100d46100e43660046107b7565b610264565b6100b6610391565b6100d46100ff3660046107e8565b61039e565b6100d4610450565b61011f61011a366004610899565b610475565b6040516100c39594939291906108b2565b61013960015481565b6040519081526020016100c3565b6100d46101553660046108f4565b61053b565b61013960075481565b6002546101709060010b81565b60405160019190910b81526020016100c3565b600580546101909061091e565b80601f01602080910402602001604051908101604052809291908181526020018280546101bc9061091e565b80156102095780601f106101de57610100808354040283529160200191610209565b820191906000526020600020905b8154815290600101906020018083116101ec57829003601f168201915b505050505081565b6006546001600160a01b0316331461022857600080fd5b6006546040516001600160a01b03909116904780156108fc02916000818181858888f19350505050158015610261573d6000803e3d6000fd5b50565b60008160000b136102895760008160000b1261028157600061028c565b60001961028c565b60015b3360009081526003602052604081205491925082810b91900b14156102f05760405162461bcd60e51b81526020600482015260156024820152744d757374206368616e6765207072657620766f746560581b60448201526064015b60405180910390fd5b336000908152600360205260408120546002805491830b92909161031890849060010b61096f565b825461ffff9182166101009390930a92830291909202199091161790555060028054600083810b929161034f90849060010b6109b5565b825461ffff9182166101009390930a928302919092021990911617905550336000908152600360205260409020805460ff90921660ff19909216919091179055565b600480546101909061091e565b60fa8151106103ef5760405162461bcd60e51b815260206004820152601c60248201527f506f7374206578636565647320636861726163746572206c696d69740000000060448201526064016102e7565b600780546000908152600860209081526040909120428155915460018301558251610422916004840191908501906106a0565b506005810180546001600160a01b0319163317905560078054906000610447836109fa565b91905055505050565b6006546001600160a01b0316331461046757600080fd5b6006546001600160a01b0316ff5b600860205260009081526040902080546001808301546002840154600485018054949592949190930b92906104a99061091e565b80601f01602080910402602001604051908101604052809291908181526020018280546104d59061091e565b80156105225780601f106104f757610100808354040283529160200191610522565b820191906000526020600020905b81548152906001019060200180831161050557829003601f168201915b505050600590930154919250506001600160a01b031685565b60008260000b136105605760008260000b12610558576000610563565b600019610563565b60015b600082815260086020908152604080832033845260030190915281205491935083810b91900b14156105cf5760405162461bcd60e51b81526020600482015260156024820152744d757374206368616e6765207072657620766f746560581b60448201526064016102e7565b60008181526008602081815260408084203385526003810183529084205485855292909152600201805491830b92909161060d90849060010b61096f565b825461ffff9182166101009390930a9283029190920219909116179055506000818152600860205260408120600201805484830b929061065190849060010b6109b5565b825461ffff9182166101009390930a92830291909202199091161790555060009081526008602090815260408083203384526003019091529020805460ff90921660ff19909216919091179055565b8280546106ac9061091e565b90600052602060002090601f0160209004810192826106ce5760008555610714565b82601f106106e757805160ff1916838001178555610714565b82800160010185558215610714579182015b828111156107145782518255916020019190600101906106f9565b50610720929150610724565b5090565b5b808211156107205760008155600101610725565b6000815180845260005b8181101561075f57602081850181015186830182015201610743565b81811115610771576000602083870101525b50601f01601f19169290920160200192915050565b6020815260006107996020830184610739565b9392505050565b8035600081900b81146107b257600080fd5b919050565b6000602082840312156107c957600080fd5b610799826107a0565b634e487b7160e01b600052604160045260246000fd5b6000602082840312156107fa57600080fd5b813567ffffffffffffffff8082111561081257600080fd5b818401915084601f83011261082657600080fd5b813581811115610838576108386107d2565b604051601f8201601f19908116603f01168101908382118183101715610860576108606107d2565b8160405282815287602084870101111561087957600080fd5b826020860160208301376000928101602001929092525095945050505050565b6000602082840312156108ab57600080fd5b5035919050565b8581528460208201528360010b604082015260a0606082015260006108da60a0830185610739565b905060018060a01b03831660808301529695505050505050565b6000806040838503121561090757600080fd5b610910836107a0565b946020939093013593505050565b600181811c9082168061093257607f821691505b6020821081141561095357634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052601160045260246000fd5b60008160010b8360010b6000811281617fff190183128115161561099557610995610959565b81617fff0183138116156109ab576109ab610959565b5090039392505050565b60008160010b8360010b6000821282617fff038213811516156109da576109da610959565b82617fff190382128116156109f1576109f1610959565b50019392505050565b6000600019821415610a0e57610a0e610959565b506001019056fea2646970667358221220446a900890e11eb3781b7d80076fd04370fa9b0e6b9b37cb78851a0baba1b73764736f6c63430008090033a2646970667358221220c13b09117799ecf7a18e75d01107490fc15b7fe6d72dbfca53dbd213a5906fa364736f6c63430008090033",
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

// GetOwnerAddress is a free data retrieval call binding the contract method 0xe6f7bf89.
//
// Solidity: function GetOwnerAddress() view returns(address adminAddress)
func (_DIAccount *DIAccountCaller) GetOwnerAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DIAccount.contract.Call(opts, &out, "GetOwnerAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOwnerAddress is a free data retrieval call binding the contract method 0xe6f7bf89.
//
// Solidity: function GetOwnerAddress() view returns(address adminAddress)
func (_DIAccount *DIAccountSession) GetOwnerAddress() (common.Address, error) {
	return _DIAccount.Contract.GetOwnerAddress(&_DIAccount.CallOpts)
}

// GetOwnerAddress is a free data retrieval call binding the contract method 0xe6f7bf89.
//
// Solidity: function GetOwnerAddress() view returns(address adminAddress)
func (_DIAccount *DIAccountCallerSession) GetOwnerAddress() (common.Address, error) {
	return _DIAccount.Contract.GetOwnerAddress(&_DIAccount.CallOpts)
}

// GetPost is a free data retrieval call binding the contract method 0xbf421138.
//
// Solidity: function GetPost(uint256 PostID) view returns(address postaddress)
func (_DIAccount *DIAccountCaller) GetPost(opts *bind.CallOpts, PostID *big.Int) (common.Address, error) {
	var out []interface{}
	err := _DIAccount.contract.Call(opts, &out, "GetPost", PostID)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPost is a free data retrieval call binding the contract method 0xbf421138.
//
// Solidity: function GetPost(uint256 PostID) view returns(address postaddress)
func (_DIAccount *DIAccountSession) GetPost(PostID *big.Int) (common.Address, error) {
	return _DIAccount.Contract.GetPost(&_DIAccount.CallOpts, PostID)
}

// GetPost is a free data retrieval call binding the contract method 0xbf421138.
//
// Solidity: function GetPost(uint256 PostID) view returns(address postaddress)
func (_DIAccount *DIAccountCallerSession) GetPost(PostID *big.Int) (common.Address, error) {
	return _DIAccount.Contract.GetPost(&_DIAccount.CallOpts, PostID)
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
// Solidity: function Posts(uint256 ) view returns(address)
func (_DIAccount *DIAccountCaller) Posts(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _DIAccount.contract.Call(opts, &out, "Posts", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Posts is a free data retrieval call binding the contract method 0x4fafef1a.
//
// Solidity: function Posts(uint256 ) view returns(address)
func (_DIAccount *DIAccountSession) Posts(arg0 *big.Int) (common.Address, error) {
	return _DIAccount.Contract.Posts(&_DIAccount.CallOpts, arg0)
}

// Posts is a free data retrieval call binding the contract method 0x4fafef1a.
//
// Solidity: function Posts(uint256 ) view returns(address)
func (_DIAccount *DIAccountCallerSession) Posts(arg0 *big.Int) (common.Address, error) {
	return _DIAccount.Contract.Posts(&_DIAccount.CallOpts, arg0)
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

// PostMessage is a paid mutator transaction binding the contract method 0x01465d36.
//
// Solidity: function PostMessage(string PostName_, string PostString_) returns(address postaddress)
func (_DIAccount *DIAccountTransactor) PostMessage(opts *bind.TransactOpts, PostName_ string, PostString_ string) (*types.Transaction, error) {
	return _DIAccount.contract.Transact(opts, "PostMessage", PostName_, PostString_)
}

// PostMessage is a paid mutator transaction binding the contract method 0x01465d36.
//
// Solidity: function PostMessage(string PostName_, string PostString_) returns(address postaddress)
func (_DIAccount *DIAccountSession) PostMessage(PostName_ string, PostString_ string) (*types.Transaction, error) {
	return _DIAccount.Contract.PostMessage(&_DIAccount.TransactOpts, PostName_, PostString_)
}

// PostMessage is a paid mutator transaction binding the contract method 0x01465d36.
//
// Solidity: function PostMessage(string PostName_, string PostString_) returns(address postaddress)
func (_DIAccount *DIAccountTransactorSession) PostMessage(PostName_ string, PostString_ string) (*types.Transaction, error) {
	return _DIAccount.Contract.PostMessage(&_DIAccount.TransactOpts, PostName_, PostString_)
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

// AdminDeleteAccount is a paid mutator transaction binding the contract method 0x3e450fff.
//
// Solidity: function adminDeleteAccount() payable returns()
func (_DIAccount *DIAccountTransactor) AdminDeleteAccount(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DIAccount.contract.Transact(opts, "adminDeleteAccount")
}

// AdminDeleteAccount is a paid mutator transaction binding the contract method 0x3e450fff.
//
// Solidity: function adminDeleteAccount() payable returns()
func (_DIAccount *DIAccountSession) AdminDeleteAccount() (*types.Transaction, error) {
	return _DIAccount.Contract.AdminDeleteAccount(&_DIAccount.TransactOpts)
}

// AdminDeleteAccount is a paid mutator transaction binding the contract method 0x3e450fff.
//
// Solidity: function adminDeleteAccount() payable returns()
func (_DIAccount *DIAccountTransactorSession) AdminDeleteAccount() (*types.Transaction, error) {
	return _DIAccount.Contract.AdminDeleteAccount(&_DIAccount.TransactOpts)
}

// PostMetaData contains all meta data concerning the Post contract.
var PostMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"PostName_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"Caption_\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"CountPosts\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"Post\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"PostName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"PostString\",\"type\":\"string\"}],\"name\":\"PostBanner\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"CommentString\",\"type\":\"string\"}],\"name\":\"AddComment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Caption\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"Comments\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"CommentNumber\",\"type\":\"uint256\"},{\"internalType\":\"int16\",\"name\":\"Karma\",\"type\":\"int16\"},{\"internalType\":\"string\",\"name\":\"CommentString\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"Poster\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CountComments\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DeletePost\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Karma\",\"outputs\":[{\"internalType\":\"int16\",\"name\":\"\",\"type\":\"int16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PostName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PostNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Transfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int8\",\"name\":\"vote\",\"type\":\"int8\"}],\"name\":\"Vote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int8\",\"name\":\"vote\",\"type\":\"int8\"},{\"internalType\":\"uint256\",\"name\":\"CommentNumber\",\"type\":\"uint256\"}],\"name\":\"Vote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"6bde24b6": "AddComment(string)",
		"1b529115": "Caption()",
		"8ea20ddc": "Comments(uint256)",
		"dc28c100": "CountComments()",
		"85141ee3": "DeletePost()",
		"e5163de5": "Karma()",
		"680e8ebc": "PostName()",
		"98c024bc": "PostNumber()",
		"406dade3": "Transfer()",
		"4e49d2da": "Vote(int8)",
		"ae81085a": "Vote(int8,uint256)",
	},
	Bin: "0x60806040523480156200001157600080fd5b5060405162000e3838038062000e3883398101604081905262000034916200023a565b4260005582516200004d906004906020860190620000c7565b50815162000063906005906020850190620000c7565b506001819055600680546001600160a01b031916321790556000546040517f7307c220a36817a866099ba844095153e39ee120706b0822e482241b2091ac5e91620000b691309060049060059062000392565b60405180910390a1505050620003dd565b828054620000d590620002ad565b90600052602060002090601f016020900481019282620000f9576000855562000144565b82601f106200011457805160ff191683800117855562000144565b8280016001018555821562000144579182015b828111156200014457825182559160200191906001019062000127565b506200015292915062000156565b5090565b5b8082111562000152576000815560010162000157565b634e487b7160e01b600052604160045260246000fd5b600082601f8301126200019557600080fd5b81516001600160401b0380821115620001b257620001b26200016d565b604051601f8301601f19908116603f01168101908282118183101715620001dd57620001dd6200016d565b81604052838152602092508683858801011115620001fa57600080fd5b600091505b838210156200021e5785820183015181830184015290820190620001ff565b83821115620002305760008385830101525b9695505050505050565b6000806000606084860312156200025057600080fd5b83516001600160401b03808211156200026857600080fd5b620002768783880162000183565b945060208601519150808211156200028d57600080fd5b506200029c8682870162000183565b925050604084015190509250925092565b600181811c90821680620002c257607f821691505b60208210811415620002e457634e487b7160e01b600052602260045260246000fd5b50919050565b8054600090600181811c90808316806200030557607f831692505b60208084108214156200032857634e487b7160e01b600052602260045260246000fd5b83885260208801828015620003465760018114620003585762000385565b60ff1987168252828201975062000385565b60008981526020902060005b878110156200037f5781548482015290860190840162000364565b83019850505b5050505050505092915050565b8481526001600160a01b0384166020820152608060408201819052600090620003be90830185620002ea565b8281036060840152620003d28185620002ea565b979650505050505050565b610a4b80620003ed6000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c806385141ee31161007157806385141ee3146101045780638ea20ddc1461010c57806398c024bc14610130578063ae81085a14610147578063dc28c1001461015a578063e5163de51461016357600080fd5b80631b529115146100ae578063406dade3146100cc5780634e49d2da146100d6578063680e8ebc146100e95780636bde24b6146100f1575b600080fd5b6100b6610183565b6040516100c39190610786565b60405180910390f35b6100d4610211565b005b6100d46100e43660046107b7565b610264565b6100b6610391565b6100d46100ff3660046107e8565b61039e565b6100d4610450565b61011f61011a366004610899565b610475565b6040516100c39594939291906108b2565b61013960015481565b6040519081526020016100c3565b6100d46101553660046108f4565b61053b565b61013960075481565b6002546101709060010b81565b60405160019190910b81526020016100c3565b600580546101909061091e565b80601f01602080910402602001604051908101604052809291908181526020018280546101bc9061091e565b80156102095780601f106101de57610100808354040283529160200191610209565b820191906000526020600020905b8154815290600101906020018083116101ec57829003601f168201915b505050505081565b6006546001600160a01b0316331461022857600080fd5b6006546040516001600160a01b03909116904780156108fc02916000818181858888f19350505050158015610261573d6000803e3d6000fd5b50565b60008160000b136102895760008160000b1261028157600061028c565b60001961028c565b60015b3360009081526003602052604081205491925082810b91900b14156102f05760405162461bcd60e51b81526020600482015260156024820152744d757374206368616e6765207072657620766f746560581b60448201526064015b60405180910390fd5b336000908152600360205260408120546002805491830b92909161031890849060010b61096f565b825461ffff9182166101009390930a92830291909202199091161790555060028054600083810b929161034f90849060010b6109b5565b825461ffff9182166101009390930a928302919092021990911617905550336000908152600360205260409020805460ff90921660ff19909216919091179055565b600480546101909061091e565b60fa8151106103ef5760405162461bcd60e51b815260206004820152601c60248201527f506f7374206578636565647320636861726163746572206c696d69740000000060448201526064016102e7565b600780546000908152600860209081526040909120428155915460018301558251610422916004840191908501906106a0565b506005810180546001600160a01b0319163317905560078054906000610447836109fa565b91905055505050565b6006546001600160a01b0316331461046757600080fd5b6006546001600160a01b0316ff5b600860205260009081526040902080546001808301546002840154600485018054949592949190930b92906104a99061091e565b80601f01602080910402602001604051908101604052809291908181526020018280546104d59061091e565b80156105225780601f106104f757610100808354040283529160200191610522565b820191906000526020600020905b81548152906001019060200180831161050557829003601f168201915b505050600590930154919250506001600160a01b031685565b60008260000b136105605760008260000b12610558576000610563565b600019610563565b60015b600082815260086020908152604080832033845260030190915281205491935083810b91900b14156105cf5760405162461bcd60e51b81526020600482015260156024820152744d757374206368616e6765207072657620766f746560581b60448201526064016102e7565b60008181526008602081815260408084203385526003810183529084205485855292909152600201805491830b92909161060d90849060010b61096f565b825461ffff9182166101009390930a9283029190920219909116179055506000818152600860205260408120600201805484830b929061065190849060010b6109b5565b825461ffff9182166101009390930a92830291909202199091161790555060009081526008602090815260408083203384526003019091529020805460ff90921660ff19909216919091179055565b8280546106ac9061091e565b90600052602060002090601f0160209004810192826106ce5760008555610714565b82601f106106e757805160ff1916838001178555610714565b82800160010185558215610714579182015b828111156107145782518255916020019190600101906106f9565b50610720929150610724565b5090565b5b808211156107205760008155600101610725565b6000815180845260005b8181101561075f57602081850181015186830182015201610743565b81811115610771576000602083870101525b50601f01601f19169290920160200192915050565b6020815260006107996020830184610739565b9392505050565b8035600081900b81146107b257600080fd5b919050565b6000602082840312156107c957600080fd5b610799826107a0565b634e487b7160e01b600052604160045260246000fd5b6000602082840312156107fa57600080fd5b813567ffffffffffffffff8082111561081257600080fd5b818401915084601f83011261082657600080fd5b813581811115610838576108386107d2565b604051601f8201601f19908116603f01168101908382118183101715610860576108606107d2565b8160405282815287602084870101111561087957600080fd5b826020860160208301376000928101602001929092525095945050505050565b6000602082840312156108ab57600080fd5b5035919050565b8581528460208201528360010b604082015260a0606082015260006108da60a0830185610739565b905060018060a01b03831660808301529695505050505050565b6000806040838503121561090757600080fd5b610910836107a0565b946020939093013593505050565b600181811c9082168061093257607f821691505b6020821081141561095357634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052601160045260246000fd5b60008160010b8360010b6000811281617fff190183128115161561099557610995610959565b81617fff0183138116156109ab576109ab610959565b5090039392505050565b60008160010b8360010b6000821282617fff038213811516156109da576109da610959565b82617fff190382128116156109f1576109f1610959565b50019392505050565b6000600019821415610a0e57610a0e610959565b506001019056fea2646970667358221220446a900890e11eb3781b7d80076fd04370fa9b0e6b9b37cb78851a0baba1b73764736f6c63430008090033",
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
func DeployPost(auth *bind.TransactOpts, backend bind.ContractBackend, PostName_ string, Caption_ string, CountPosts *big.Int) (common.Address, *types.Transaction, *Post, error) {
	parsed, err := PostMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PostBin), backend, PostName_, Caption_, CountPosts)
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

// Caption is a free data retrieval call binding the contract method 0x1b529115.
//
// Solidity: function Caption() view returns(string)
func (_Post *PostCaller) Caption(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Post.contract.Call(opts, &out, "Caption")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Caption is a free data retrieval call binding the contract method 0x1b529115.
//
// Solidity: function Caption() view returns(string)
func (_Post *PostSession) Caption() (string, error) {
	return _Post.Contract.Caption(&_Post.CallOpts)
}

// Caption is a free data retrieval call binding the contract method 0x1b529115.
//
// Solidity: function Caption() view returns(string)
func (_Post *PostCallerSession) Caption() (string, error) {
	return _Post.Contract.Caption(&_Post.CallOpts)
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
	Post       common.Address
	PostName   string
	PostString string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPostBanner is a free log retrieval operation binding the contract event 0x7307c220a36817a866099ba844095153e39ee120706b0822e482241b2091ac5e.
//
// Solidity: event PostBanner(uint256 timestamp, address Post, string PostName, string PostString)
func (_Post *PostFilterer) FilterPostBanner(opts *bind.FilterOpts) (*PostPostBannerIterator, error) {

	logs, sub, err := _Post.contract.FilterLogs(opts, "PostBanner")
	if err != nil {
		return nil, err
	}
	return &PostPostBannerIterator{contract: _Post.contract, event: "PostBanner", logs: logs, sub: sub}, nil
}

// WatchPostBanner is a free log subscription operation binding the contract event 0x7307c220a36817a866099ba844095153e39ee120706b0822e482241b2091ac5e.
//
// Solidity: event PostBanner(uint256 timestamp, address Post, string PostName, string PostString)
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

// ParsePostBanner is a log parse operation binding the contract event 0x7307c220a36817a866099ba844095153e39ee120706b0822e482241b2091ac5e.
//
// Solidity: event PostBanner(uint256 timestamp, address Post, string PostName, string PostString)
func (_Post *PostFilterer) ParsePostBanner(log types.Log) (*PostPostBanner, error) {
	event := new(PostPostBanner)
	if err := _Post.contract.UnpackLog(event, "PostBanner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
