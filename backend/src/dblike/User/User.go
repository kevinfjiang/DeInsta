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

// DIAccountComment is an auto generated low-level Go binding around an user-defined struct.
type DIAccountComment struct {
	Timestamp     *big.Int
	CommentNumber *big.Int
	Karma         int16
	CommentString string
	Poster        common.Address
}

// DIAccountMetaData contains all meta data concerning the DIAccount contract.
var DIAccountMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"Name_\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"PostNumber\",\"type\":\"uint256\"}],\"name\":\"PostBanner\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"PostNumber\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"CommentString\",\"type\":\"string\"}],\"name\":\"AddComment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Bio\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int8\",\"name\":\"vote\",\"type\":\"int8\"},{\"internalType\":\"uint256\",\"name\":\"PostNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"CommentNumber\",\"type\":\"uint256\"}],\"name\":\"CommentVote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CountPosts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"PostURL_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"Caption_\",\"type\":\"string\"}],\"name\":\"CreatePost\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DeleteAccount\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"PostID\",\"type\":\"uint256\"}],\"name\":\"DeletePost\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"PostID\",\"type\":\"uint256\"}],\"name\":\"GetComments\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"CommentNumber\",\"type\":\"uint256\"},{\"internalType\":\"int16\",\"name\":\"Karma\",\"type\":\"int16\"},{\"internalType\":\"string\",\"name\":\"CommentString\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"Poster\",\"type\":\"address\"}],\"internalType\":\"structDIAccount.Comment[]\",\"name\":\"CommentList\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"PostID\",\"type\":\"uint256\"}],\"name\":\"GetPost\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"CommentNumber\",\"type\":\"uint256\"},{\"internalType\":\"int16\",\"name\":\"Karma\",\"type\":\"int16\"},{\"internalType\":\"string\",\"name\":\"CommentString\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"Poster\",\"type\":\"address\"}],\"internalType\":\"structDIAccount.Comment[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Karma\",\"outputs\":[{\"internalType\":\"int16\",\"name\":\"\",\"type\":\"int16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"Posts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"PostNumber\",\"type\":\"uint256\"},{\"internalType\":\"int16\",\"name\":\"Karma\",\"type\":\"int16\"},{\"internalType\":\"string\",\"name\":\"IPFSurl\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"Caption\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"CountComments\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"Bio_\",\"type\":\"string\"}],\"name\":\"SetBio\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int8\",\"name\":\"vote\",\"type\":\"int8\"},{\"internalType\":\"uint256\",\"name\":\"PostNumber\",\"type\":\"uint256\"}],\"name\":\"Vote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"3b0cf47b": "AddComment(uint256,string)",
		"c2d2d0f9": "Bio()",
		"117d7f48": "CommentVote(int8,uint256,uint256)",
		"27677343": "CountPosts()",
		"7a8d157e": "CreatePost(string,string)",
		"9099967e": "DeleteAccount()",
		"75c16cb2": "DeletePost(uint256)",
		"d9fe8e5d": "GetComments(uint256)",
		"bf421138": "GetPost(uint256)",
		"e5163de5": "Karma()",
		"8052474d": "Name()",
		"4fafef1a": "Posts(uint256)",
		"27d0b9ec": "SetBio(string)",
		"ae81085a": "Vote(int8,uint256)",
	},
	Bin: "0x60806040523480156200001157600080fd5b506040516200167038038062001670833981016040819052620000349162000124565b80516200004990600090602084019062000068565b50506000600455600580546001600160a01b031916331790556200023d565b828054620000769062000200565b90600052602060002090601f0160209004810192826200009a5760008555620000e5565b82601f10620000b557805160ff1916838001178555620000e5565b82800160010185558215620000e5579182015b82811115620000e5578251825591602001919060010190620000c8565b50620000f3929150620000f7565b5090565b5b80821115620000f35760008155600101620000f8565b634e487b7160e01b600052604160045260246000fd5b600060208083850312156200013857600080fd5b82516001600160401b03808211156200015057600080fd5b818501915085601f8301126200016557600080fd5b8151818111156200017a576200017a6200010e565b604051601f8201601f19908116603f01168101908382118183101715620001a557620001a56200010e565b816040528281528886848701011115620001be57600080fd5b600093505b82841015620001e25784840186015181850187015292850192620001c3565b82841115620001f45760008684830101525b98975050505050505050565b600181811c908216806200021557607f821691505b602082108114156200023757634e487b7160e01b600052602260045260246000fd5b50919050565b611423806200024d6000396000f3fe6080604052600436106100dd5760003560e01c80638052474d1161007f578063bf42113811610059578063bf42113814610229578063c2d2d0f914610257578063d9fe8e5d1461026c578063e5163de51461028c57600080fd5b80638052474d146101df5780639099967e14610201578063ae81085a1461020957600080fd5b80633b0cf47b116100bb5780633b0cf47b1461014d5780634fafef1a1461016d57806375c16cb21461019f5780637a8d157e146101bf57600080fd5b8063117d7f48146100e2578063276773431461010457806327d0b9ec1461012d575b600080fd5b3480156100ee57600080fd5b506101026100fd366004610f91565b6102b9565b005b34801561011057600080fd5b5061011a60045481565b6040519081526020015b60405180910390f35b34801561013957600080fd5b50610102610148366004611067565b610528565b34801561015957600080fd5b506101026101683660046110a4565b610556565b34801561017957600080fd5b5061018d6101883660046110eb565b6106ad565b60405161012496959493929190611151565b3480156101ab57600080fd5b506101026101ba3660046110eb565b610808565b3480156101cb57600080fd5b506101026101da36600461119f565b6108a0565b3480156101eb57600080fd5b506101f46109ab565b60405161012491906111f9565b610102610a39565b34801561021557600080fd5b50610102610224366004611213565b610a5e565b34801561023557600080fd5b5061024a6102443660046110eb565b50606090565b604051610124919061123d565b34801561026357600080fd5b506101f4610cd1565b34801561027857600080fd5b5061024a6102873660046110eb565b610cde565b34801561029857600080fd5b506003546102a69060010b81565b60405160019190910b8152602001610124565b60008360000b136102de5760008360000b126102d65760006102e1565b6000196102e1565b60015b92508260000b600283815481106102fa576102fa6112e0565b9060005260206000209060090201600801828154811061031c5761031c6112e0565b6000918252602080832033845290910190526040812054900b141561037f5760405162461bcd60e51b8152602060048201526014602482015273165bdd49dd9948185b1c9958591e481d9bdd195960621b60448201526064015b60405180910390fd5b60028281548110610392576103926112e0565b906000526020600020906009020160080181815481106103b4576103b46112e0565b6000918252602080832033845290910190526040812054600280549190920b9190849081106103e5576103e56112e0565b90600052602060002090600902016007018281548110610407576104076112e0565b600091825260208220600260059092020101805490919061042c90849060010b61130c565b92506101000a81548161ffff021916908360010b61ffff1602179055508260000b60028381548110610460576104606112e0565b90600052602060002090600902016007018281548110610482576104826112e0565b60009182526020822060026005909202010180549091906104a790849060010b611352565b92506101000a81548161ffff021916908360010b61ffff16021790555082600283815481106104d8576104d86112e0565b906000526020600020906009020160080182815481106104fa576104fa6112e0565b6000918252602080832033845291909101905260409020805460ff191660ff92909216919091179055505050565b6005546001600160a01b0316331461053f57600080fd5b8051610552906001906020840190610e3a565b5050565b60328151106105a75760405162461bcd60e51b815260206004820152601c60248201527f506f7374206578636565647320636861726163746572206c696d6974000000006044820152606401610376565b6000600283815481106105bc576105bc6112e0565b9060005260206000209060090201600701600284815481106105e0576105e06112e0565b90600052602060002090600902016006015481548110610602576106026112e0565b906000526020600020906005020190504281600001819055506002838154811061062e5761062e6112e0565b6000918252602091829020600660099092020101546001830155825161065c91600384019190850190610e3a565b506004810180546001600160a01b031916331790556002805484908110610685576106856112e0565b6000918252602082206006600990920201018054916106a383611397565b9190505550505050565b600281815481106106bd57600080fd5b60009182526020909120600990910201805460018083015460028401546004850180549496509194920b926106f1906113b2565b80601f016020809104026020016040519081016040528092919081815260200182805461071d906113b2565b801561076a5780601f1061073f5761010080835404028352916020019161076a565b820191906000526020600020905b81548152906001019060200180831161074d57829003601f168201915b50505050509080600501805461077f906113b2565b80601f01602080910402602001604051908101604052809291908181526020018280546107ab906113b2565b80156107f85780601f106107cd576101008083540402835291602001916107f8565b820191906000526020600020905b8154815290600101906020018083116107db57829003601f168201915b5050505050908060060154905086565b6005546001600160a01b0316331461081f57600080fd5b60028181548110610832576108326112e0565b6000918252602082206009909102018181556001810182905560028101805461ffff19169055906108666004830182610ebe565b610874600583016000610ebe565b600682016000905560078201600061088c9190610efb565b61089b60088301600080825552565b505050565b6005546001600160a01b031633146108b757600080fd5b60328151106109085760405162461bcd60e51b815260206004820152601c60248201527f506f7374206578636565647320636861726163746572206c696d6974000000006044820152606401610376565b600060026004548154811061091f5761091f6112e0565b6000918252602091829020426009909202019081556004805460018301558551919350610953929084019190860190610e3a565b5081516109699060058301906020850190610e3a565b5060045481546040513091907f6e8001e55f8223a69373fdb10aecda8ca632280f2f689cbf7365b8c728e5187890600090a4600480549060006106a383611397565b600080546109b8906113b2565b80601f01602080910402602001604051908101604052809291908181526020018280546109e4906113b2565b8015610a315780601f10610a0657610100808354040283529160200191610a31565b820191906000526020600020905b815481529060010190602001808311610a1457829003601f168201915b505050505081565b6005546001600160a01b03163314610a5057600080fd5b6005546001600160a01b0316ff5b60008260000b13610a835760008260000b12610a7b576000610a86565b600019610a86565b60015b91508160000b60028281548110610a9f57610a9f6112e0565b6000918252602080832033845260036009909302019190910190526040812054900b1415610b065760405162461bcd60e51b8152602060048201526014602482015273165bdd49dd9948185b1c9958591e481d9bdd195960621b6044820152606401610376565b60028181548110610b1957610b196112e0565b6000918252602080832033845260036009909302019190910190526040812054600280549190920b919083908110610b5357610b536112e0565b6000918252602082206002600990920201018054909190610b7890849060010b61130c565b92506101000a81548161ffff021916908360010b61ffff16021790555060028181548110610ba857610ba86112e0565b600091825260208083203384526003600990930201820190526040822054815490830b9290610bdb90849060010b61130c565b92506101000a81548161ffff021916908360010b61ffff1602179055508160000b60028281548110610c0f57610c0f6112e0565b6000918252602082206002600990920201018054909190610c3490849060010b611352565b825461ffff9182166101009390930a92830291909202199091161790555060038054600084810b9291610c6b90849060010b611352565b92506101000a81548161ffff021916908360010b61ffff1602179055508160028281548110610c9c57610c9c6112e0565b60009182526020808320338452600992909202909101600301905260409020805460ff191660ff929092169190911790555050565b600180546109b8906113b2565b606060028281548110610cf357610cf36112e0565b9060005260206000209060090201600701805480602002602001604051908101604052809291908181526020016000905b82821015610e2f57838290600052602060002090600502016040518060a001604052908160008201548152602001600182015481526020016002820160009054906101000a900460010b60010b60010b8152602001600382018054610d88906113b2565b80601f0160208091040260200160405190810160405280929190818152602001828054610db4906113b2565b8015610e015780601f10610dd657610100808354040283529160200191610e01565b820191906000526020600020905b815481529060010190602001808311610de457829003601f168201915b5050509183525050600491909101546001600160a01b03166020918201529082526001929092019101610d24565b505050509050919050565b828054610e46906113b2565b90600052602060002090601f016020900481019282610e685760008555610eae565b82601f10610e8157805160ff1916838001178555610eae565b82800160010185558215610eae579182015b82811115610eae578251825591602001919060010190610e93565b50610eba929150610f1c565b5090565b508054610eca906113b2565b6000825580601f10610eda575050565b601f016020900490600052602060002090810190610ef89190610f1c565b50565b5080546000825560050290600052602060002090810190610ef89190610f31565b5b80821115610eba5760008155600101610f1d565b80821115610eba5760008082556001820181905560028201805461ffff19169055610f5f6003830182610ebe565b506004810180546001600160a01b0319169055600501610f31565b8035600081900b8114610f8c57600080fd5b919050565b600080600060608486031215610fa657600080fd5b610faf84610f7a565b95602085013595506040909401359392505050565b634e487b7160e01b600052604160045260246000fd5b600082601f830112610feb57600080fd5b813567ffffffffffffffff8082111561100657611006610fc4565b604051601f8301601f19908116603f0116810190828211818310171561102e5761102e610fc4565b8160405283815286602085880101111561104757600080fd5b836020870160208301376000602085830101528094505050505092915050565b60006020828403121561107957600080fd5b813567ffffffffffffffff81111561109057600080fd5b61109c84828501610fda565b949350505050565b600080604083850312156110b757600080fd5b82359150602083013567ffffffffffffffff8111156110d557600080fd5b6110e185828601610fda565b9150509250929050565b6000602082840312156110fd57600080fd5b5035919050565b6000815180845260005b8181101561112a5760208185018101518683018201520161110e565b8181111561113c576000602083870101525b50601f01601f19169290920160200192915050565b8681528560208201528460010b604082015260c06060820152600061117960c0830186611104565b828103608084015261118b8186611104565b9150508260a0830152979650505050505050565b600080604083850312156111b257600080fd5b823567ffffffffffffffff808211156111ca57600080fd5b6111d686838701610fda565b935060208501359150808211156111ec57600080fd5b506110e185828601610fda565b60208152600061120c6020830184611104565b9392505050565b6000806040838503121561122657600080fd5b61122f83610f7a565b946020939093013593505050565b60006020808301818452808551808352604092508286019150828160051b87010184880160005b838110156112d257603f19898403018552815160a08151855288820151898601528782015160010b8886015260608083015182828801526112a783880182611104565b6080948501516001600160a01b03169790940196909652505094870194925090860190600101611264565b509098975050505050505050565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b60008160010b8360010b6000811281617fff1901831281151615611332576113326112f6565b81617fff018313811615611348576113486112f6565b5090039392505050565b60008160010b8360010b6000821282617fff03821381151615611377576113776112f6565b82617fff1903821281161561138e5761138e6112f6565b50019392505050565b60006000198214156113ab576113ab6112f6565b5060010190565b600181811c908216806113c657607f821691505b602082108114156113e757634e487b7160e01b600052602260045260246000fd5b5091905056fea26469706673582212205db0081c01fd4abbfc8956f5f35bdcf2258ae05f140601932f9dc7f94d2d00ab64736f6c63430008090033",
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

// GetComments is a free data retrieval call binding the contract method 0xd9fe8e5d.
//
// Solidity: function GetComments(uint256 PostID) view returns((uint256,uint256,int16,string,address)[] CommentList)
func (_DIAccount *DIAccountCaller) GetComments(opts *bind.CallOpts, PostID *big.Int) ([]DIAccountComment, error) {
	var out []interface{}
	err := _DIAccount.contract.Call(opts, &out, "GetComments", PostID)

	if err != nil {
		return *new([]DIAccountComment), err
	}

	out0 := *abi.ConvertType(out[0], new([]DIAccountComment)).(*[]DIAccountComment)

	return out0, err

}

// GetComments is a free data retrieval call binding the contract method 0xd9fe8e5d.
//
// Solidity: function GetComments(uint256 PostID) view returns((uint256,uint256,int16,string,address)[] CommentList)
func (_DIAccount *DIAccountSession) GetComments(PostID *big.Int) ([]DIAccountComment, error) {
	return _DIAccount.Contract.GetComments(&_DIAccount.CallOpts, PostID)
}

// GetComments is a free data retrieval call binding the contract method 0xd9fe8e5d.
//
// Solidity: function GetComments(uint256 PostID) view returns((uint256,uint256,int16,string,address)[] CommentList)
func (_DIAccount *DIAccountCallerSession) GetComments(PostID *big.Int) ([]DIAccountComment, error) {
	return _DIAccount.Contract.GetComments(&_DIAccount.CallOpts, PostID)
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

// GetPost is a paid mutator transaction binding the contract method 0xbf421138.
//
// Solidity: function GetPost(uint256 PostID) returns((uint256,uint256,int16,string,address)[])
func (_DIAccount *DIAccountTransactor) GetPost(opts *bind.TransactOpts, PostID *big.Int) (*types.Transaction, error) {
	return _DIAccount.contract.Transact(opts, "GetPost", PostID)
}

// GetPost is a paid mutator transaction binding the contract method 0xbf421138.
//
// Solidity: function GetPost(uint256 PostID) returns((uint256,uint256,int16,string,address)[])
func (_DIAccount *DIAccountSession) GetPost(PostID *big.Int) (*types.Transaction, error) {
	return _DIAccount.Contract.GetPost(&_DIAccount.TransactOpts, PostID)
}

// GetPost is a paid mutator transaction binding the contract method 0xbf421138.
//
// Solidity: function GetPost(uint256 PostID) returns((uint256,uint256,int16,string,address)[])
func (_DIAccount *DIAccountTransactorSession) GetPost(PostID *big.Int) (*types.Transaction, error) {
	return _DIAccount.Contract.GetPost(&_DIAccount.TransactOpts, PostID)
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
