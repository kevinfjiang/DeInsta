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
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"Name_\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"PostNumber\",\"type\":\"uint256\"}],\"name\":\"PostBanner\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"PostNumber\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"CommentString\",\"type\":\"string\"}],\"name\":\"AddComment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Bio\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int8\",\"name\":\"vote\",\"type\":\"int8\"},{\"internalType\":\"uint256\",\"name\":\"PostNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"CommentNumber\",\"type\":\"uint256\"}],\"name\":\"CommentVote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CountPosts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"PostURL_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"Caption_\",\"type\":\"string\"}],\"name\":\"CreatePost\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DeleteAccount\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"PostID\",\"type\":\"uint256\"}],\"name\":\"DeletePost\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"PostID\",\"type\":\"uint256\"}],\"name\":\"GetComments\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"CommentNumber\",\"type\":\"uint256\"},{\"internalType\":\"int16\",\"name\":\"Karma\",\"type\":\"int16\"},{\"internalType\":\"string\",\"name\":\"CommentString\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"Poster\",\"type\":\"address\"}],\"internalType\":\"structDIAccount.Comment[]\",\"name\":\"CommentList\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Karma\",\"outputs\":[{\"internalType\":\"int16\",\"name\":\"\",\"type\":\"int16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"Posts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"Account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"PostNumber\",\"type\":\"uint256\"},{\"internalType\":\"int16\",\"name\":\"Karma\",\"type\":\"int16\"},{\"internalType\":\"string\",\"name\":\"IPFSurl\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"Caption\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"CountComments\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"Bio_\",\"type\":\"string\"}],\"name\":\"SetBio\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int8\",\"name\":\"vote\",\"type\":\"int8\"},{\"internalType\":\"uint256\",\"name\":\"PostNumber\",\"type\":\"uint256\"}],\"name\":\"Vote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"3b0cf47b": "AddComment(uint256,string)",
		"c2d2d0f9": "Bio()",
		"117d7f48": "CommentVote(int8,uint256,uint256)",
		"27677343": "CountPosts()",
		"7a8d157e": "CreatePost(string,string)",
		"9099967e": "DeleteAccount()",
		"75c16cb2": "DeletePost(uint256)",
		"d9fe8e5d": "GetComments(uint256)",
		"e5163de5": "Karma()",
		"8052474d": "Name()",
		"4fafef1a": "Posts(uint256)",
		"27d0b9ec": "SetBio(string)",
		"ae81085a": "Vote(int8,uint256)",
	},
	Bin: "0x60806040523480156200001157600080fd5b506040516200168338038062001683833981016040819052620000349162000124565b80516200004990600090602084019062000068565b50506000600455600580546001600160a01b031916331790556200023d565b828054620000769062000200565b90600052602060002090601f0160209004810192826200009a5760008555620000e5565b82601f10620000b557805160ff1916838001178555620000e5565b82800160010185558215620000e5579182015b82811115620000e5578251825591602001919060010190620000c8565b50620000f3929150620000f7565b5090565b5b80821115620000f35760008155600101620000f8565b634e487b7160e01b600052604160045260246000fd5b600060208083850312156200013857600080fd5b82516001600160401b03808211156200015057600080fd5b818501915085601f8301126200016557600080fd5b8151818111156200017a576200017a6200010e565b604051601f8201601f19908116603f01168101908382118183101715620001a557620001a56200010e565b816040528281528886848701011115620001be57600080fd5b600093505b82841015620001e25784840186015181850187015292850192620001c3565b82841115620001f45760008684830101525b98975050505050505050565b600181811c908216806200021557607f821691505b602082108114156200023757634e487b7160e01b600052602260045260246000fd5b50919050565b611436806200024d6000396000f3fe6080604052600436106100c25760003560e01c80637a8d157e1161007f578063ae81085a11610059578063ae81085a146101ef578063c2d2d0f91461020f578063d9fe8e5d14610224578063e5163de51461025157600080fd5b80637a8d157e146101a55780638052474d146101c55780639099967e146101e757600080fd5b8063117d7f48146100c757806327677343146100e957806327d0b9ec146101125780633b0cf47b146101325780634fafef1a1461015257806375c16cb214610185575b600080fd5b3480156100d357600080fd5b506100e76100e2366004610f95565b61027e565b005b3480156100f557600080fd5b506100ff60045481565b6040519081526020015b60405180910390f35b34801561011e57600080fd5b506100e761012d36600461106b565b6104ed565b34801561013e57600080fd5b506100e761014d3660046110a8565b61051b565b34801561015e57600080fd5b5061017261016d3660046110ef565b610672565b6040516101099796959493929190611155565b34801561019157600080fd5b506100e76101a03660046110ef565b6107e1565b3480156101b157600080fd5b506100e76101c03660046111b2565b61088b565b3480156101d157600080fd5b506101da6109aa565b604051610109919061120c565b6100e7610a38565b3480156101fb57600080fd5b506100e761020a366004611226565b610a5d565b34801561021b57600080fd5b506101da610cd5565b34801561023057600080fd5b5061024461023f3660046110ef565b610ce2565b6040516101099190611250565b34801561025d57600080fd5b5060035461026b9060010b81565b60405160019190910b8152602001610109565b60008360000b136102a35760008360000b1261029b5760006102a6565b6000196102a6565b60015b92508260000b600283815481106102bf576102bf6112f3565b90600052602060002090600a020160090182815481106102e1576102e16112f3565b6000918252602080832033845290910190526040812054900b14156103445760405162461bcd60e51b8152602060048201526014602482015273165bdd49dd9948185b1c9958591e481d9bdd195960621b60448201526064015b60405180910390fd5b60028281548110610357576103576112f3565b90600052602060002090600a02016009018181548110610379576103796112f3565b6000918252602080832033845290910190526040812054600280549190920b9190849081106103aa576103aa6112f3565b90600052602060002090600a020160080182815481106103cc576103cc6112f3565b60009182526020822060026005909202010180549091906103f190849060010b61131f565b92506101000a81548161ffff021916908360010b61ffff1602179055508260000b60028381548110610425576104256112f3565b90600052602060002090600a02016008018281548110610447576104476112f3565b600091825260208220600260059092020101805490919061046c90849060010b611365565b92506101000a81548161ffff021916908360010b61ffff160217905550826002838154811061049d5761049d6112f3565b90600052602060002090600a020160090182815481106104bf576104bf6112f3565b6000918252602080832033845291909101905260409020805460ff191660ff92909216919091179055505050565b6005546001600160a01b0316331461050457600080fd5b8051610517906001906020840190610e3e565b5050565b603281511061056c5760405162461bcd60e51b815260206004820152601c60248201527f506f7374206578636565647320636861726163746572206c696d697400000000604482015260640161033b565b600060028381548110610581576105816112f3565b90600052602060002090600a0201600801600284815481106105a5576105a56112f3565b90600052602060002090600a020160070154815481106105c7576105c76112f3565b90600052602060002090600502019050428160000181905550600283815481106105f3576105f36112f3565b60009182526020918290206007600a9092020101546001830155825161062191600384019190850190610e3e565b506004810180546001600160a01b03191633179055600280548490811061064a5761064a6112f3565b6000918252602082206007600a9092020101805491610668836113aa565b9190505550505050565b6002818154811061068257600080fd5b60009182526020909120600a909102018054600180830154600284015460038501546005860180549597506001600160a01b03909316959194930b929091906106ca906113c5565b80601f01602080910402602001604051908101604052809291908181526020018280546106f6906113c5565b80156107435780601f1061071857610100808354040283529160200191610743565b820191906000526020600020905b81548152906001019060200180831161072657829003601f168201915b505050505090806006018054610758906113c5565b80601f0160208091040260200160405190810160405280929190818152602001828054610784906113c5565b80156107d15780601f106107a6576101008083540402835291602001916107d1565b820191906000526020600020905b8154815290600101906020018083116107b457829003601f168201915b5050505050908060070154905087565b6005546001600160a01b031633146107f857600080fd5b6002818154811061080b5761080b6112f3565b600091825260208220600a909102018181556001810180546001600160a01b03191690556002810182905560038101805461ffff19169055906108516005830182610ec2565b61085f600683016000610ec2565b60078201600090556008820160006108779190610eff565b61088660098301600080825552565b505050565b6005546001600160a01b031633146108a257600080fd5b60328151106108f35760405162461bcd60e51b815260206004820152601c60248201527f506f7374206578636565647320636861726163746572206c696d697400000000604482015260640161033b565b600060026004548154811061090a5761090a6112f3565b600091825260209182902042600a909202019081556001810180546001600160a01b031916301790556004546002820155845190925061095291600584019190860190610e3e565b5081516109689060068301906020850190610e3e565b5060045481546040513091907f6e8001e55f8223a69373fdb10aecda8ca632280f2f689cbf7365b8c728e5187890600090a460048054906000610668836113aa565b600080546109b7906113c5565b80601f01602080910402602001604051908101604052809291908181526020018280546109e3906113c5565b8015610a305780601f10610a0557610100808354040283529160200191610a30565b820191906000526020600020905b815481529060010190602001808311610a1357829003601f168201915b505050505081565b6005546001600160a01b03163314610a4f57600080fd5b6005546001600160a01b0316ff5b60008260000b13610a825760008260000b12610a7a576000610a85565b600019610a85565b60015b91508160000b60028281548110610a9e57610a9e6112f3565b600091825260208083203384526004600a909302019190910190526040812054900b1415610b055760405162461bcd60e51b8152602060048201526014602482015273165bdd49dd9948185b1c9958591e481d9bdd195960621b604482015260640161033b565b60028181548110610b1857610b186112f3565b600091825260208083203384526004600a909302019190910190526040812054600280549190920b919083908110610b5257610b526112f3565b6000918252602082206003600a90920201018054909190610b7790849060010b61131f565b92506101000a81548161ffff021916908360010b61ffff16021790555060028181548110610ba757610ba76112f3565b600091825260208083203384526004600a9093020191909101905260408120546003805491830b929091610bdf90849060010b61131f565b92506101000a81548161ffff021916908360010b61ffff1602179055508160000b60028281548110610c1357610c136112f3565b6000918252602082206003600a90920201018054909190610c3890849060010b611365565b825461ffff9182166101009390930a92830291909202199091161790555060038054600084810b9291610c6f90849060010b611365565b92506101000a81548161ffff021916908360010b61ffff1602179055508160028281548110610ca057610ca06112f3565b60009182526020808320338452600a92909202909101600401905260409020805460ff191660ff929092169190911790555050565b600180546109b7906113c5565b606060028281548110610cf757610cf76112f3565b90600052602060002090600a0201600801805480602002602001604051908101604052809291908181526020016000905b82821015610e3357838290600052602060002090600502016040518060a001604052908160008201548152602001600182015481526020016002820160009054906101000a900460010b60010b60010b8152602001600382018054610d8c906113c5565b80601f0160208091040260200160405190810160405280929190818152602001828054610db8906113c5565b8015610e055780601f10610dda57610100808354040283529160200191610e05565b820191906000526020600020905b815481529060010190602001808311610de857829003601f168201915b5050509183525050600491909101546001600160a01b03166020918201529082526001929092019101610d28565b505050509050919050565b828054610e4a906113c5565b90600052602060002090601f016020900481019282610e6c5760008555610eb2565b82601f10610e8557805160ff1916838001178555610eb2565b82800160010185558215610eb2579182015b82811115610eb2578251825591602001919060010190610e97565b50610ebe929150610f20565b5090565b508054610ece906113c5565b6000825580601f10610ede575050565b601f016020900490600052602060002090810190610efc9190610f20565b50565b5080546000825560050290600052602060002090810190610efc9190610f35565b5b80821115610ebe5760008155600101610f21565b80821115610ebe5760008082556001820181905560028201805461ffff19169055610f636003830182610ec2565b506004810180546001600160a01b0319169055600501610f35565b8035600081900b8114610f9057600080fd5b919050565b600080600060608486031215610faa57600080fd5b610fb384610f7e565b95602085013595506040909401359392505050565b634e487b7160e01b600052604160045260246000fd5b600082601f830112610fef57600080fd5b813567ffffffffffffffff8082111561100a5761100a610fc8565b604051601f8301601f19908116603f0116810190828211818310171561103257611032610fc8565b8160405283815286602085880101111561104b57600080fd5b836020870160208301376000602085830101528094505050505092915050565b60006020828403121561107d57600080fd5b813567ffffffffffffffff81111561109457600080fd5b6110a084828501610fde565b949350505050565b600080604083850312156110bb57600080fd5b82359150602083013567ffffffffffffffff8111156110d957600080fd5b6110e585828601610fde565b9150509250929050565b60006020828403121561110157600080fd5b5035919050565b6000815180845260005b8181101561112e57602081850181015186830182015201611112565b81811115611140576000602083870101525b50601f01601f19169290920160200192915050565b87815260018060a01b03871660208201528560408201528460010b606082015260e06080820152600061118b60e0830186611108565b82810360a084015261119d8186611108565b9150508260c083015298975050505050505050565b600080604083850312156111c557600080fd5b823567ffffffffffffffff808211156111dd57600080fd5b6111e986838701610fde565b935060208501359150808211156111ff57600080fd5b506110e585828601610fde565b60208152600061121f6020830184611108565b9392505050565b6000806040838503121561123957600080fd5b61124283610f7e565b946020939093013593505050565b60006020808301818452808551808352604092508286019150828160051b87010184880160005b838110156112e557603f19898403018552815160a08151855288820151898601528782015160010b8886015260608083015182828801526112ba83880182611108565b6080948501516001600160a01b03169790940196909652505094870194925090860190600101611277565b509098975050505050505050565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b60008160010b8360010b6000811281617fff190183128115161561134557611345611309565b81617fff01831381161561135b5761135b611309565b5090039392505050565b60008160010b8360010b6000821282617fff0382138115161561138a5761138a611309565b82617fff190382128116156113a1576113a1611309565b50019392505050565b60006000198214156113be576113be611309565b5060010190565b600181811c908216806113d957607f821691505b602082108114156113fa57634e487b7160e01b600052602260045260246000fd5b5091905056fea26469706673582212207c247fba2bf78003476ef0991c7fb47da8103addbce62ff0cbb61e9ce1c3253064736f6c634300080a0033",
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
// Solidity: function Posts(uint256 ) view returns(uint256 timestamp, address Account, uint256 PostNumber, int16 Karma, string IPFSurl, string Caption, uint256 CountComments)
func (_DIAccount *DIAccountCaller) Posts(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Timestamp     *big.Int
	Account       common.Address
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
		Account       common.Address
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
	outstruct.Account = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.PostNumber = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Karma = *abi.ConvertType(out[3], new(int16)).(*int16)
	outstruct.IPFSurl = *abi.ConvertType(out[4], new(string)).(*string)
	outstruct.Caption = *abi.ConvertType(out[5], new(string)).(*string)
	outstruct.CountComments = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Posts is a free data retrieval call binding the contract method 0x4fafef1a.
//
// Solidity: function Posts(uint256 ) view returns(uint256 timestamp, address Account, uint256 PostNumber, int16 Karma, string IPFSurl, string Caption, uint256 CountComments)
func (_DIAccount *DIAccountSession) Posts(arg0 *big.Int) (struct {
	Timestamp     *big.Int
	Account       common.Address
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
// Solidity: function Posts(uint256 ) view returns(uint256 timestamp, address Account, uint256 PostNumber, int16 Karma, string IPFSurl, string Caption, uint256 CountComments)
func (_DIAccount *DIAccountCallerSession) Posts(arg0 *big.Int) (struct {
	Timestamp     *big.Int
	Account       common.Address
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
