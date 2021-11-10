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

// EIaccountMetaData contains all meta data concerning the EIaccount contract.
var EIaccountMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"Name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"Bio_\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"User\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"Hash\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"PostDescription\",\"type\":\"string\"}],\"name\":\"PostBanner\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"CommentString\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"PostNum\",\"type\":\"uint256\"}],\"name\":\"AddComment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddFollower\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"Hash_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"Description_\",\"type\":\"string\"}],\"name\":\"AddPost\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Bio\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"CheckFollowPost\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"FollowPost\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CountFollowers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CountFollowing\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CountPosts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"PostID\",\"type\":\"uint256\"}],\"name\":\"DeletePost\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Follow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"Followers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"Following\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetOwnerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"adminAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Karma\",\"outputs\":[{\"internalType\":\"int16\",\"name\":\"\",\"type\":\"int16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"Posts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"PostNumber\",\"type\":\"uint256\"},{\"internalType\":\"int16\",\"name\":\"Karma\",\"type\":\"int16\"},{\"internalType\":\"string\",\"name\":\"Hash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"PostDescription\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"Owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"CountComments\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"RemoveFollower\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Transfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"UnFollow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int8\",\"name\":\"vote\",\"type\":\"int8\"},{\"internalType\":\"uint256\",\"name\":\"PostNum\",\"type\":\"uint256\"}],\"name\":\"Vote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"adminDeleteAccount\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"cc55e360": "AddComment(string,uint256)",
		"f23014cb": "AddFollower(address)",
		"198f3d11": "AddPost(string,string)",
		"c2d2d0f9": "Bio()",
		"63a917f2": "CheckFollowPost(address)",
		"ae8a2430": "CountFollowers()",
		"34555be0": "CountFollowing()",
		"27677343": "CountPosts()",
		"75c16cb2": "DeletePost(uint256)",
		"e06da1cb": "Follow(address)",
		"3df67974": "Followers(address)",
		"dbf13de5": "Following(address)",
		"e6f7bf89": "GetOwnerAddress()",
		"e5163de5": "Karma()",
		"8052474d": "Name()",
		"4fafef1a": "Posts(uint256)",
		"300638f9": "RemoveFollower(address)",
		"406dade3": "Transfer()",
		"6e493c29": "UnFollow(address)",
		"ae81085a": "Vote(int8,uint256)",
		"3e450fff": "adminDeleteAccount()",
	},
	Bin: "0x60806040523480156200001157600080fd5b50604051620015183803806200151883398101604081905262000034916200029c565b600f8251106200008b5760405162461bcd60e51b815260206004820152601c60248201527f4e616d65206578636565647320636861726163746572206c696d69740000000060448201526064015b60405180910390fd5b601e815110620000de5760405162461bcd60e51b815260206004820152601b60248201527f42696f206578636565647320636861726163746572206c696d69740000000000604482015260640162000082565b8151620000f390600090602085019062000129565b5080516200010990600190602084019062000129565b5050600060045550600580546001600160a01b0319163317905562000343565b828054620001379062000306565b90600052602060002090601f0160209004810192826200015b5760008555620001a6565b82601f106200017657805160ff1916838001178555620001a6565b82800160010185558215620001a6579182015b82811115620001a657825182559160200191906001019062000189565b50620001b4929150620001b8565b5090565b5b80821115620001b45760008155600101620001b9565b634e487b7160e01b600052604160045260246000fd5b600082601f830112620001f757600080fd5b81516001600160401b0380821115620002145762000214620001cf565b604051601f8301601f19908116603f011681019082821181831017156200023f576200023f620001cf565b816040528381526020925086838588010111156200025c57600080fd5b600091505b8382101562000280578582018301518183018401529082019062000261565b83821115620002925760008385830101525b9695505050505050565b60008060408385031215620002b057600080fd5b82516001600160401b0380821115620002c857600080fd5b620002d686838701620001e5565b93506020850151915080821115620002ed57600080fd5b50620002fc85828601620001e5565b9150509250929050565b600181811c908216806200031b57607f821691505b602082108114156200033d57634e487b7160e01b600052602260045260246000fd5b50919050565b6111c580620003536000396000f3fe60806040526004361061012a5760003560e01c806375c16cb2116100ab578063cc55e3601161006f578063cc55e36014610326578063dbf13de514610346578063e06da1cb14610376578063e5163de514610396578063e6f7bf89146103c3578063f23014cb1461017a57600080fd5b806375c16cb2146102995780638052474d146102b9578063ae81085a146102db578063ae8a2430146102fb578063c2d2d0f91461031157600080fd5b80633e450fff116100f25780633e450fff146101f0578063406dade3146101f85780634fafef1a1461020d57806363a917f2146102405780636e493c291461027957600080fd5b8063198f3d111461012f5780632767734314610151578063300638f91461017a57806334555be01461019a5780633df67974146101b0575b600080fd5b34801561013b57600080fd5b5061014f61014a366004610e49565b6103eb565b005b34801561015d57600080fd5b5061016760045481565b6040519081526020015b60405180910390f35b34801561018657600080fd5b5061014f610195366004610ead565b61053f565b3480156101a657600080fd5b5061016760095481565b3480156101bc57600080fd5b506101e06101cb366004610ead565b60066020526000908152604090205460ff1681565b6040519015158152602001610171565b61014f61058f565b34801561020457600080fd5b5061014f6105b4565b34801561021957600080fd5b5061022d610228366004610edd565b610607565b6040516101719796959493929190610f43565b34801561024c57600080fd5b506101e061025b366004610ead565b6001600160a01b031660009081526006602052604090205460ff1690565b34801561028557600080fd5b5061014f610294366004610ead565b610772565b3480156102a557600080fd5b5061014f6102b4366004610edd565b6107b7565b3480156102c557600080fd5b506102ce610841565b6040516101719190610f9d565b3480156102e757600080fd5b5061014f6102f6366004610fb0565b6108cf565b34801561030757600080fd5b5061016760075481565b34801561031d57600080fd5b506102ce610b43565b34801561033257600080fd5b5061014f610341366004610fe2565b610b50565b34801561035257600080fd5b506101e0610361366004610ead565b60086020526000908152604090205460ff1681565b34801561038257600080fd5b5061014f610391366004610ead565b610c8f565b3480156103a257600080fd5b506003546103b09060010b81565b60405160019190910b8152602001610171565b3480156103cf57600080fd5b506005546040516001600160a01b039091168152602001610171565b6005546001600160a01b0316331461040257600080fd5b600082511161041057600080fd5b60648151106104665760405162461bcd60e51b815260206004820152601c60248201527f506f7374206578636565647320636861726163746572206c696d69740000000060448201526064015b60405180910390fd5b3361047057600080fd5b600060026004548154811061048757610487611027565b60009182526020918290204260099092020190815584519092506104b391600484019190860190610cd7565b5081516104c99060058301906020850190610cd7565b506004805460018301556006820180546001600160a01b03191633179055546040517f303dbb34abb7589def7f6abbd68426bbf4f4a3f6062603673caedc6edc1441849161051d913091908790879061103d565b60405180910390a16004805490600061053583611097565b9190505550505050565b6005546001600160a01b0316331461055657600080fd5b6001600160a01b0381166000908152600660205260408120805460ff19166001179055600980549161058783611097565b919050555050565b6005546001600160a01b031633146105a657600080fd5b6005546001600160a01b0316ff5b6005546001600160a01b031633146105cb57600080fd5b6005546040516001600160a01b03909116904780156108fc02916000818181858888f19350505050158015610604573d6000803e3d6000fd5b50565b6002818154811061061757600080fd5b60009182526020909120600990910201805460018083015460028401546004850180549496509194920b9261064b906110b2565b80601f0160208091040260200160405190810160405280929190818152602001828054610677906110b2565b80156106c45780601f10610699576101008083540402835291602001916106c4565b820191906000526020600020905b8154815290600101906020018083116106a757829003601f168201915b5050505050908060050180546106d9906110b2565b80601f0160208091040260200160405190810160405280929190818152602001828054610705906110b2565b80156107525780601f1061072757610100808354040283529160200191610752565b820191906000526020600020905b81548152906001019060200180831161073557829003601f168201915b50505050600683015460079093015491926001600160a01b031691905087565b6005546001600160a01b0316331461078957600080fd5b6001600160a01b0381166000908152600860205260408120805460ff191690556007805491610587836110ed565b6005546001600160a01b031633146107ce57600080fd5b600281815481106107e1576107e1611027565b6000918252602082206009909102018181556001810182905560028101805461ffff19169055906108156004830182610d5b565b610823600583016000610d5b565b506006810180546001600160a01b0319169055600060079091015550565b6000805461084e906110b2565b80601f016020809104026020016040519081016040528092919081815260200182805461087a906110b2565b80156108c75780601f1061089c576101008083540402835291602001916108c7565b820191906000526020600020905b8154815290600101906020018083116108aa57829003601f168201915b505050505081565b60008260000b136108f45760008260000b126108ec5760006108f7565b6000196108f7565b60015b91508160000b6002828154811061091057610910611027565b6000918252602080832033845260036009909302019190910190526040812054900b14156109785760405162461bcd60e51b81526020600482015260156024820152744d757374206368616e6765207072657620766f746560581b604482015260640161045d565b6002818154811061098b5761098b611027565b600091825260208083203384526003600990930201820190526040822054815490830b92906109be90849060010b611104565b92506101000a81548161ffff021916908360010b61ffff160217905550600281815481106109ee576109ee611027565b6000918252602080832033845260036009909302019190910190526040812054600280549190920b919083908110610a2857610a28611027565b6000918252602082206002600990920201018054909190610a4d90849060010b611104565b825461ffff9182166101009390930a92830291909202199091161790555060038054600084810b9291610a8490849060010b61114a565b92506101000a81548161ffff021916908360010b61ffff1602179055508160000b60028281548110610ab857610ab8611027565b6000918252602082206002600990920201018054909190610add90849060010b61114a565b92506101000a81548161ffff021916908360010b61ffff1602179055508160028281548110610b0e57610b0e611027565b60009182526020808320338452600992909202909101600301905260409020805460ff191660ff929092169190911790555050565b6001805461084e906110b2565b6064825110610ba15760405162461bcd60e51b815260206004820152601c60248201527f506f7374206578636565647320636861726163746572206c696d697400000000604482015260640161045d565b600060028281548110610bb657610bb6611027565b9060005260206000209060090201600801600060028481548110610bdc57610bdc611027565b9060005260206000209060090201600701548152602001908152602001600020905042816000018190555060028281548110610c1a57610c1a611027565b60009182526020918290206007600990920201015460018301558351610c4891600484019190860190610cd7565b506005810180546001600160a01b031916331790556002805483908110610c7157610c71611027565b60009182526020822060076009909202010180549161053583611097565b6005546001600160a01b03163314610ca657600080fd5b6001600160a01b0381166000908152600860205260408120805460ff19166001179055600780549161058783611097565b828054610ce3906110b2565b90600052602060002090601f016020900481019282610d055760008555610d4b565b82601f10610d1e57805160ff1916838001178555610d4b565b82800160010185558215610d4b579182015b82811115610d4b578251825591602001919060010190610d30565b50610d57929150610d91565b5090565b508054610d67906110b2565b6000825580601f10610d77575050565b601f01602090049060005260206000209081019061060491905b5b80821115610d575760008155600101610d92565b634e487b7160e01b600052604160045260246000fd5b600082601f830112610dcd57600080fd5b813567ffffffffffffffff80821115610de857610de8610da6565b604051601f8301601f19908116603f01168101908282118183101715610e1057610e10610da6565b81604052838152866020858801011115610e2957600080fd5b836020870160208301376000602085830101528094505050505092915050565b60008060408385031215610e5c57600080fd5b823567ffffffffffffffff80821115610e7457600080fd5b610e8086838701610dbc565b93506020850135915080821115610e9657600080fd5b50610ea385828601610dbc565b9150509250929050565b600060208284031215610ebf57600080fd5b81356001600160a01b0381168114610ed657600080fd5b9392505050565b600060208284031215610eef57600080fd5b5035919050565b6000815180845260005b81811015610f1c57602081850181015186830182015201610f00565b81811115610f2e576000602083870101525b50601f01601f19169290920160200192915050565b8781528660208201528560010b604082015260e060608201526000610f6b60e0830187610ef6565b8281036080840152610f7d8187610ef6565b6001600160a01b039590951660a0840152505060c0015295945050505050565b602081526000610ed66020830184610ef6565b60008060408385031215610fc357600080fd5b82358060000b8114610fd457600080fd5b946020939093013593505050565b60008060408385031215610ff557600080fd5b823567ffffffffffffffff81111561100c57600080fd5b61101885828601610dbc565b95602094909401359450505050565b634e487b7160e01b600052603260045260246000fd5b60018060a01b03851681528360208201526080604082015260006110646080830185610ef6565b82810360608401526110768185610ef6565b979650505050505050565b634e487b7160e01b600052601160045260246000fd5b60006000198214156110ab576110ab611081565b5060010190565b600181811c908216806110c657607f821691505b602082108114156110e757634e487b7160e01b600052602260045260246000fd5b50919050565b6000816110fc576110fc611081565b506000190190565b60008160010b8360010b6000811281617fff190183128115161561112a5761112a611081565b81617fff01831381161561114057611140611081565b5090039392505050565b60008160010b8360010b6000821282617fff0382138115161561116f5761116f611081565b82617fff1903821281161561118657611186611081565b5001939250505056fea26469706673582212204ba47bb2c13c6ecc92aec01d7fb2559a788012bc708cf749cf90c4f9d96e13a364736f6c63430008090033",
}

// EIaccountABI is the input ABI used to generate the binding from.
// Deprecated: Use EIaccountMetaData.ABI instead.
var EIaccountABI = EIaccountMetaData.ABI

// Deprecated: Use EIaccountMetaData.Sigs instead.
// EIaccountFuncSigs maps the 4-byte function signature to its string representation.
var EIaccountFuncSigs = EIaccountMetaData.Sigs

// EIaccountBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EIaccountMetaData.Bin instead.
var EIaccountBin = EIaccountMetaData.Bin

// DeployEIaccount deploys a new Ethereum contract, binding an instance of EIaccount to it.
func DeployEIaccount(auth *bind.TransactOpts, backend bind.ContractBackend, Name_ string, Bio_ string) (common.Address, *types.Transaction, *EIaccount, error) {
	parsed, err := EIaccountMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EIaccountBin), backend, Name_, Bio_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EIaccount{EIaccountCaller: EIaccountCaller{contract: contract}, EIaccountTransactor: EIaccountTransactor{contract: contract}, EIaccountFilterer: EIaccountFilterer{contract: contract}}, nil
}

// EIaccount is an auto generated Go binding around an Ethereum contract.
type EIaccount struct {
	EIaccountCaller     // Read-only binding to the contract
	EIaccountTransactor // Write-only binding to the contract
	EIaccountFilterer   // Log filterer for contract events
}

// EIaccountCaller is an auto generated read-only Go binding around an Ethereum contract.
type EIaccountCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EIaccountTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EIaccountTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EIaccountFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EIaccountFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EIaccountSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EIaccountSession struct {
	Contract     *EIaccount        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EIaccountCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EIaccountCallerSession struct {
	Contract *EIaccountCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// EIaccountTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EIaccountTransactorSession struct {
	Contract     *EIaccountTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// EIaccountRaw is an auto generated low-level Go binding around an Ethereum contract.
type EIaccountRaw struct {
	Contract *EIaccount // Generic contract binding to access the raw methods on
}

// EIaccountCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EIaccountCallerRaw struct {
	Contract *EIaccountCaller // Generic read-only contract binding to access the raw methods on
}

// EIaccountTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EIaccountTransactorRaw struct {
	Contract *EIaccountTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEIaccount creates a new instance of EIaccount, bound to a specific deployed contract.
func NewEIaccount(address common.Address, backend bind.ContractBackend) (*EIaccount, error) {
	contract, err := bindEIaccount(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EIaccount{EIaccountCaller: EIaccountCaller{contract: contract}, EIaccountTransactor: EIaccountTransactor{contract: contract}, EIaccountFilterer: EIaccountFilterer{contract: contract}}, nil
}

// NewEIaccountCaller creates a new read-only instance of EIaccount, bound to a specific deployed contract.
func NewEIaccountCaller(address common.Address, caller bind.ContractCaller) (*EIaccountCaller, error) {
	contract, err := bindEIaccount(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EIaccountCaller{contract: contract}, nil
}

// NewEIaccountTransactor creates a new write-only instance of EIaccount, bound to a specific deployed contract.
func NewEIaccountTransactor(address common.Address, transactor bind.ContractTransactor) (*EIaccountTransactor, error) {
	contract, err := bindEIaccount(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EIaccountTransactor{contract: contract}, nil
}

// NewEIaccountFilterer creates a new log filterer instance of EIaccount, bound to a specific deployed contract.
func NewEIaccountFilterer(address common.Address, filterer bind.ContractFilterer) (*EIaccountFilterer, error) {
	contract, err := bindEIaccount(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EIaccountFilterer{contract: contract}, nil
}

// bindEIaccount binds a generic wrapper to an already deployed contract.
func bindEIaccount(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EIaccountABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EIaccount *EIaccountRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EIaccount.Contract.EIaccountCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EIaccount *EIaccountRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EIaccount.Contract.EIaccountTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EIaccount *EIaccountRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EIaccount.Contract.EIaccountTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EIaccount *EIaccountCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EIaccount.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EIaccount *EIaccountTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EIaccount.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EIaccount *EIaccountTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EIaccount.Contract.contract.Transact(opts, method, params...)
}

// Bio is a free data retrieval call binding the contract method 0xc2d2d0f9.
//
// Solidity: function Bio() view returns(string)
func (_EIaccount *EIaccountCaller) Bio(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _EIaccount.contract.Call(opts, &out, "Bio")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Bio is a free data retrieval call binding the contract method 0xc2d2d0f9.
//
// Solidity: function Bio() view returns(string)
func (_EIaccount *EIaccountSession) Bio() (string, error) {
	return _EIaccount.Contract.Bio(&_EIaccount.CallOpts)
}

// Bio is a free data retrieval call binding the contract method 0xc2d2d0f9.
//
// Solidity: function Bio() view returns(string)
func (_EIaccount *EIaccountCallerSession) Bio() (string, error) {
	return _EIaccount.Contract.Bio(&_EIaccount.CallOpts)
}

// CheckFollowPost is a free data retrieval call binding the contract method 0x63a917f2.
//
// Solidity: function CheckFollowPost(address account) view returns(bool FollowPost)
func (_EIaccount *EIaccountCaller) CheckFollowPost(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _EIaccount.contract.Call(opts, &out, "CheckFollowPost", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckFollowPost is a free data retrieval call binding the contract method 0x63a917f2.
//
// Solidity: function CheckFollowPost(address account) view returns(bool FollowPost)
func (_EIaccount *EIaccountSession) CheckFollowPost(account common.Address) (bool, error) {
	return _EIaccount.Contract.CheckFollowPost(&_EIaccount.CallOpts, account)
}

// CheckFollowPost is a free data retrieval call binding the contract method 0x63a917f2.
//
// Solidity: function CheckFollowPost(address account) view returns(bool FollowPost)
func (_EIaccount *EIaccountCallerSession) CheckFollowPost(account common.Address) (bool, error) {
	return _EIaccount.Contract.CheckFollowPost(&_EIaccount.CallOpts, account)
}

// CountFollowers is a free data retrieval call binding the contract method 0xae8a2430.
//
// Solidity: function CountFollowers() view returns(uint256)
func (_EIaccount *EIaccountCaller) CountFollowers(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EIaccount.contract.Call(opts, &out, "CountFollowers")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CountFollowers is a free data retrieval call binding the contract method 0xae8a2430.
//
// Solidity: function CountFollowers() view returns(uint256)
func (_EIaccount *EIaccountSession) CountFollowers() (*big.Int, error) {
	return _EIaccount.Contract.CountFollowers(&_EIaccount.CallOpts)
}

// CountFollowers is a free data retrieval call binding the contract method 0xae8a2430.
//
// Solidity: function CountFollowers() view returns(uint256)
func (_EIaccount *EIaccountCallerSession) CountFollowers() (*big.Int, error) {
	return _EIaccount.Contract.CountFollowers(&_EIaccount.CallOpts)
}

// CountFollowing is a free data retrieval call binding the contract method 0x34555be0.
//
// Solidity: function CountFollowing() view returns(uint256)
func (_EIaccount *EIaccountCaller) CountFollowing(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EIaccount.contract.Call(opts, &out, "CountFollowing")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CountFollowing is a free data retrieval call binding the contract method 0x34555be0.
//
// Solidity: function CountFollowing() view returns(uint256)
func (_EIaccount *EIaccountSession) CountFollowing() (*big.Int, error) {
	return _EIaccount.Contract.CountFollowing(&_EIaccount.CallOpts)
}

// CountFollowing is a free data retrieval call binding the contract method 0x34555be0.
//
// Solidity: function CountFollowing() view returns(uint256)
func (_EIaccount *EIaccountCallerSession) CountFollowing() (*big.Int, error) {
	return _EIaccount.Contract.CountFollowing(&_EIaccount.CallOpts)
}

// CountPosts is a free data retrieval call binding the contract method 0x27677343.
//
// Solidity: function CountPosts() view returns(uint256)
func (_EIaccount *EIaccountCaller) CountPosts(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EIaccount.contract.Call(opts, &out, "CountPosts")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CountPosts is a free data retrieval call binding the contract method 0x27677343.
//
// Solidity: function CountPosts() view returns(uint256)
func (_EIaccount *EIaccountSession) CountPosts() (*big.Int, error) {
	return _EIaccount.Contract.CountPosts(&_EIaccount.CallOpts)
}

// CountPosts is a free data retrieval call binding the contract method 0x27677343.
//
// Solidity: function CountPosts() view returns(uint256)
func (_EIaccount *EIaccountCallerSession) CountPosts() (*big.Int, error) {
	return _EIaccount.Contract.CountPosts(&_EIaccount.CallOpts)
}

// Followers is a free data retrieval call binding the contract method 0x3df67974.
//
// Solidity: function Followers(address ) view returns(bool)
func (_EIaccount *EIaccountCaller) Followers(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _EIaccount.contract.Call(opts, &out, "Followers", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Followers is a free data retrieval call binding the contract method 0x3df67974.
//
// Solidity: function Followers(address ) view returns(bool)
func (_EIaccount *EIaccountSession) Followers(arg0 common.Address) (bool, error) {
	return _EIaccount.Contract.Followers(&_EIaccount.CallOpts, arg0)
}

// Followers is a free data retrieval call binding the contract method 0x3df67974.
//
// Solidity: function Followers(address ) view returns(bool)
func (_EIaccount *EIaccountCallerSession) Followers(arg0 common.Address) (bool, error) {
	return _EIaccount.Contract.Followers(&_EIaccount.CallOpts, arg0)
}

// Following is a free data retrieval call binding the contract method 0xdbf13de5.
//
// Solidity: function Following(address ) view returns(bool)
func (_EIaccount *EIaccountCaller) Following(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _EIaccount.contract.Call(opts, &out, "Following", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Following is a free data retrieval call binding the contract method 0xdbf13de5.
//
// Solidity: function Following(address ) view returns(bool)
func (_EIaccount *EIaccountSession) Following(arg0 common.Address) (bool, error) {
	return _EIaccount.Contract.Following(&_EIaccount.CallOpts, arg0)
}

// Following is a free data retrieval call binding the contract method 0xdbf13de5.
//
// Solidity: function Following(address ) view returns(bool)
func (_EIaccount *EIaccountCallerSession) Following(arg0 common.Address) (bool, error) {
	return _EIaccount.Contract.Following(&_EIaccount.CallOpts, arg0)
}

// GetOwnerAddress is a free data retrieval call binding the contract method 0xe6f7bf89.
//
// Solidity: function GetOwnerAddress() view returns(address adminAddress)
func (_EIaccount *EIaccountCaller) GetOwnerAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EIaccount.contract.Call(opts, &out, "GetOwnerAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOwnerAddress is a free data retrieval call binding the contract method 0xe6f7bf89.
//
// Solidity: function GetOwnerAddress() view returns(address adminAddress)
func (_EIaccount *EIaccountSession) GetOwnerAddress() (common.Address, error) {
	return _EIaccount.Contract.GetOwnerAddress(&_EIaccount.CallOpts)
}

// GetOwnerAddress is a free data retrieval call binding the contract method 0xe6f7bf89.
//
// Solidity: function GetOwnerAddress() view returns(address adminAddress)
func (_EIaccount *EIaccountCallerSession) GetOwnerAddress() (common.Address, error) {
	return _EIaccount.Contract.GetOwnerAddress(&_EIaccount.CallOpts)
}

// Karma is a free data retrieval call binding the contract method 0xe5163de5.
//
// Solidity: function Karma() view returns(int16)
func (_EIaccount *EIaccountCaller) Karma(opts *bind.CallOpts) (int16, error) {
	var out []interface{}
	err := _EIaccount.contract.Call(opts, &out, "Karma")

	if err != nil {
		return *new(int16), err
	}

	out0 := *abi.ConvertType(out[0], new(int16)).(*int16)

	return out0, err

}

// Karma is a free data retrieval call binding the contract method 0xe5163de5.
//
// Solidity: function Karma() view returns(int16)
func (_EIaccount *EIaccountSession) Karma() (int16, error) {
	return _EIaccount.Contract.Karma(&_EIaccount.CallOpts)
}

// Karma is a free data retrieval call binding the contract method 0xe5163de5.
//
// Solidity: function Karma() view returns(int16)
func (_EIaccount *EIaccountCallerSession) Karma() (int16, error) {
	return _EIaccount.Contract.Karma(&_EIaccount.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x8052474d.
//
// Solidity: function Name() view returns(string)
func (_EIaccount *EIaccountCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _EIaccount.contract.Call(opts, &out, "Name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x8052474d.
//
// Solidity: function Name() view returns(string)
func (_EIaccount *EIaccountSession) Name() (string, error) {
	return _EIaccount.Contract.Name(&_EIaccount.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x8052474d.
//
// Solidity: function Name() view returns(string)
func (_EIaccount *EIaccountCallerSession) Name() (string, error) {
	return _EIaccount.Contract.Name(&_EIaccount.CallOpts)
}

// Posts is a free data retrieval call binding the contract method 0x4fafef1a.
//
// Solidity: function Posts(uint256 ) view returns(uint256 timestamp, uint256 PostNumber, int16 Karma, string Hash, string PostDescription, address Owner, uint256 CountComments)
func (_EIaccount *EIaccountCaller) Posts(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Timestamp       *big.Int
	PostNumber      *big.Int
	Karma           int16
	Hash            string
	PostDescription string
	Owner           common.Address
	CountComments   *big.Int
}, error) {
	var out []interface{}
	err := _EIaccount.contract.Call(opts, &out, "Posts", arg0)

	outstruct := new(struct {
		Timestamp       *big.Int
		PostNumber      *big.Int
		Karma           int16
		Hash            string
		PostDescription string
		Owner           common.Address
		CountComments   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Timestamp = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.PostNumber = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Karma = *abi.ConvertType(out[2], new(int16)).(*int16)
	outstruct.Hash = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.PostDescription = *abi.ConvertType(out[4], new(string)).(*string)
	outstruct.Owner = *abi.ConvertType(out[5], new(common.Address)).(*common.Address)
	outstruct.CountComments = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Posts is a free data retrieval call binding the contract method 0x4fafef1a.
//
// Solidity: function Posts(uint256 ) view returns(uint256 timestamp, uint256 PostNumber, int16 Karma, string Hash, string PostDescription, address Owner, uint256 CountComments)
func (_EIaccount *EIaccountSession) Posts(arg0 *big.Int) (struct {
	Timestamp       *big.Int
	PostNumber      *big.Int
	Karma           int16
	Hash            string
	PostDescription string
	Owner           common.Address
	CountComments   *big.Int
}, error) {
	return _EIaccount.Contract.Posts(&_EIaccount.CallOpts, arg0)
}

// Posts is a free data retrieval call binding the contract method 0x4fafef1a.
//
// Solidity: function Posts(uint256 ) view returns(uint256 timestamp, uint256 PostNumber, int16 Karma, string Hash, string PostDescription, address Owner, uint256 CountComments)
func (_EIaccount *EIaccountCallerSession) Posts(arg0 *big.Int) (struct {
	Timestamp       *big.Int
	PostNumber      *big.Int
	Karma           int16
	Hash            string
	PostDescription string
	Owner           common.Address
	CountComments   *big.Int
}, error) {
	return _EIaccount.Contract.Posts(&_EIaccount.CallOpts, arg0)
}

// AddComment is a paid mutator transaction binding the contract method 0xcc55e360.
//
// Solidity: function AddComment(string CommentString, uint256 PostNum) returns()
func (_EIaccount *EIaccountTransactor) AddComment(opts *bind.TransactOpts, CommentString string, PostNum *big.Int) (*types.Transaction, error) {
	return _EIaccount.contract.Transact(opts, "AddComment", CommentString, PostNum)
}

// AddComment is a paid mutator transaction binding the contract method 0xcc55e360.
//
// Solidity: function AddComment(string CommentString, uint256 PostNum) returns()
func (_EIaccount *EIaccountSession) AddComment(CommentString string, PostNum *big.Int) (*types.Transaction, error) {
	return _EIaccount.Contract.AddComment(&_EIaccount.TransactOpts, CommentString, PostNum)
}

// AddComment is a paid mutator transaction binding the contract method 0xcc55e360.
//
// Solidity: function AddComment(string CommentString, uint256 PostNum) returns()
func (_EIaccount *EIaccountTransactorSession) AddComment(CommentString string, PostNum *big.Int) (*types.Transaction, error) {
	return _EIaccount.Contract.AddComment(&_EIaccount.TransactOpts, CommentString, PostNum)
}

// AddFollower is a paid mutator transaction binding the contract method 0xf23014cb.
//
// Solidity: function AddFollower(address account) returns()
func (_EIaccount *EIaccountTransactor) AddFollower(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _EIaccount.contract.Transact(opts, "AddFollower", account)
}

// AddFollower is a paid mutator transaction binding the contract method 0xf23014cb.
//
// Solidity: function AddFollower(address account) returns()
func (_EIaccount *EIaccountSession) AddFollower(account common.Address) (*types.Transaction, error) {
	return _EIaccount.Contract.AddFollower(&_EIaccount.TransactOpts, account)
}

// AddFollower is a paid mutator transaction binding the contract method 0xf23014cb.
//
// Solidity: function AddFollower(address account) returns()
func (_EIaccount *EIaccountTransactorSession) AddFollower(account common.Address) (*types.Transaction, error) {
	return _EIaccount.Contract.AddFollower(&_EIaccount.TransactOpts, account)
}

// AddPost is a paid mutator transaction binding the contract method 0x198f3d11.
//
// Solidity: function AddPost(string Hash_, string Description_) returns()
func (_EIaccount *EIaccountTransactor) AddPost(opts *bind.TransactOpts, Hash_ string, Description_ string) (*types.Transaction, error) {
	return _EIaccount.contract.Transact(opts, "AddPost", Hash_, Description_)
}

// AddPost is a paid mutator transaction binding the contract method 0x198f3d11.
//
// Solidity: function AddPost(string Hash_, string Description_) returns()
func (_EIaccount *EIaccountSession) AddPost(Hash_ string, Description_ string) (*types.Transaction, error) {
	return _EIaccount.Contract.AddPost(&_EIaccount.TransactOpts, Hash_, Description_)
}

// AddPost is a paid mutator transaction binding the contract method 0x198f3d11.
//
// Solidity: function AddPost(string Hash_, string Description_) returns()
func (_EIaccount *EIaccountTransactorSession) AddPost(Hash_ string, Description_ string) (*types.Transaction, error) {
	return _EIaccount.Contract.AddPost(&_EIaccount.TransactOpts, Hash_, Description_)
}

// DeletePost is a paid mutator transaction binding the contract method 0x75c16cb2.
//
// Solidity: function DeletePost(uint256 PostID) returns()
func (_EIaccount *EIaccountTransactor) DeletePost(opts *bind.TransactOpts, PostID *big.Int) (*types.Transaction, error) {
	return _EIaccount.contract.Transact(opts, "DeletePost", PostID)
}

// DeletePost is a paid mutator transaction binding the contract method 0x75c16cb2.
//
// Solidity: function DeletePost(uint256 PostID) returns()
func (_EIaccount *EIaccountSession) DeletePost(PostID *big.Int) (*types.Transaction, error) {
	return _EIaccount.Contract.DeletePost(&_EIaccount.TransactOpts, PostID)
}

// DeletePost is a paid mutator transaction binding the contract method 0x75c16cb2.
//
// Solidity: function DeletePost(uint256 PostID) returns()
func (_EIaccount *EIaccountTransactorSession) DeletePost(PostID *big.Int) (*types.Transaction, error) {
	return _EIaccount.Contract.DeletePost(&_EIaccount.TransactOpts, PostID)
}

// Follow is a paid mutator transaction binding the contract method 0xe06da1cb.
//
// Solidity: function Follow(address account) returns()
func (_EIaccount *EIaccountTransactor) Follow(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _EIaccount.contract.Transact(opts, "Follow", account)
}

// Follow is a paid mutator transaction binding the contract method 0xe06da1cb.
//
// Solidity: function Follow(address account) returns()
func (_EIaccount *EIaccountSession) Follow(account common.Address) (*types.Transaction, error) {
	return _EIaccount.Contract.Follow(&_EIaccount.TransactOpts, account)
}

// Follow is a paid mutator transaction binding the contract method 0xe06da1cb.
//
// Solidity: function Follow(address account) returns()
func (_EIaccount *EIaccountTransactorSession) Follow(account common.Address) (*types.Transaction, error) {
	return _EIaccount.Contract.Follow(&_EIaccount.TransactOpts, account)
}

// RemoveFollower is a paid mutator transaction binding the contract method 0x300638f9.
//
// Solidity: function RemoveFollower(address account) returns()
func (_EIaccount *EIaccountTransactor) RemoveFollower(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _EIaccount.contract.Transact(opts, "RemoveFollower", account)
}

// RemoveFollower is a paid mutator transaction binding the contract method 0x300638f9.
//
// Solidity: function RemoveFollower(address account) returns()
func (_EIaccount *EIaccountSession) RemoveFollower(account common.Address) (*types.Transaction, error) {
	return _EIaccount.Contract.RemoveFollower(&_EIaccount.TransactOpts, account)
}

// RemoveFollower is a paid mutator transaction binding the contract method 0x300638f9.
//
// Solidity: function RemoveFollower(address account) returns()
func (_EIaccount *EIaccountTransactorSession) RemoveFollower(account common.Address) (*types.Transaction, error) {
	return _EIaccount.Contract.RemoveFollower(&_EIaccount.TransactOpts, account)
}

// Transfer is a paid mutator transaction binding the contract method 0x406dade3.
//
// Solidity: function Transfer() returns()
func (_EIaccount *EIaccountTransactor) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EIaccount.contract.Transact(opts, "Transfer")
}

// Transfer is a paid mutator transaction binding the contract method 0x406dade3.
//
// Solidity: function Transfer() returns()
func (_EIaccount *EIaccountSession) Transfer() (*types.Transaction, error) {
	return _EIaccount.Contract.Transfer(&_EIaccount.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0x406dade3.
//
// Solidity: function Transfer() returns()
func (_EIaccount *EIaccountTransactorSession) Transfer() (*types.Transaction, error) {
	return _EIaccount.Contract.Transfer(&_EIaccount.TransactOpts)
}

// UnFollow is a paid mutator transaction binding the contract method 0x6e493c29.
//
// Solidity: function UnFollow(address account) returns()
func (_EIaccount *EIaccountTransactor) UnFollow(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _EIaccount.contract.Transact(opts, "UnFollow", account)
}

// UnFollow is a paid mutator transaction binding the contract method 0x6e493c29.
//
// Solidity: function UnFollow(address account) returns()
func (_EIaccount *EIaccountSession) UnFollow(account common.Address) (*types.Transaction, error) {
	return _EIaccount.Contract.UnFollow(&_EIaccount.TransactOpts, account)
}

// UnFollow is a paid mutator transaction binding the contract method 0x6e493c29.
//
// Solidity: function UnFollow(address account) returns()
func (_EIaccount *EIaccountTransactorSession) UnFollow(account common.Address) (*types.Transaction, error) {
	return _EIaccount.Contract.UnFollow(&_EIaccount.TransactOpts, account)
}

// Vote is a paid mutator transaction binding the contract method 0xae81085a.
//
// Solidity: function Vote(int8 vote, uint256 PostNum) returns()
func (_EIaccount *EIaccountTransactor) Vote(opts *bind.TransactOpts, vote int8, PostNum *big.Int) (*types.Transaction, error) {
	return _EIaccount.contract.Transact(opts, "Vote", vote, PostNum)
}

// Vote is a paid mutator transaction binding the contract method 0xae81085a.
//
// Solidity: function Vote(int8 vote, uint256 PostNum) returns()
func (_EIaccount *EIaccountSession) Vote(vote int8, PostNum *big.Int) (*types.Transaction, error) {
	return _EIaccount.Contract.Vote(&_EIaccount.TransactOpts, vote, PostNum)
}

// Vote is a paid mutator transaction binding the contract method 0xae81085a.
//
// Solidity: function Vote(int8 vote, uint256 PostNum) returns()
func (_EIaccount *EIaccountTransactorSession) Vote(vote int8, PostNum *big.Int) (*types.Transaction, error) {
	return _EIaccount.Contract.Vote(&_EIaccount.TransactOpts, vote, PostNum)
}

// AdminDeleteAccount is a paid mutator transaction binding the contract method 0x3e450fff.
//
// Solidity: function adminDeleteAccount() payable returns()
func (_EIaccount *EIaccountTransactor) AdminDeleteAccount(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EIaccount.contract.Transact(opts, "adminDeleteAccount")
}

// AdminDeleteAccount is a paid mutator transaction binding the contract method 0x3e450fff.
//
// Solidity: function adminDeleteAccount() payable returns()
func (_EIaccount *EIaccountSession) AdminDeleteAccount() (*types.Transaction, error) {
	return _EIaccount.Contract.AdminDeleteAccount(&_EIaccount.TransactOpts)
}

// AdminDeleteAccount is a paid mutator transaction binding the contract method 0x3e450fff.
//
// Solidity: function adminDeleteAccount() payable returns()
func (_EIaccount *EIaccountTransactorSession) AdminDeleteAccount() (*types.Transaction, error) {
	return _EIaccount.Contract.AdminDeleteAccount(&_EIaccount.TransactOpts)
}

// EIaccountPostBannerIterator is returned from FilterPostBanner and is used to iterate over the raw logs and unpacked data for PostBanner events raised by the EIaccount contract.
type EIaccountPostBannerIterator struct {
	Event *EIaccountPostBanner // Event containing the contract specifics and raw log

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
func (it *EIaccountPostBannerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EIaccountPostBanner)
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
		it.Event = new(EIaccountPostBanner)
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
func (it *EIaccountPostBannerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EIaccountPostBannerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EIaccountPostBanner represents a PostBanner event raised by the EIaccount contract.
type EIaccountPostBanner struct {
	User            common.Address
	Id              *big.Int
	Hash            string
	PostDescription string
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPostBanner is a free log retrieval operation binding the contract event 0x303dbb34abb7589def7f6abbd68426bbf4f4a3f6062603673caedc6edc144184.
//
// Solidity: event PostBanner(address User, uint256 id, string Hash, string PostDescription)
func (_EIaccount *EIaccountFilterer) FilterPostBanner(opts *bind.FilterOpts) (*EIaccountPostBannerIterator, error) {

	logs, sub, err := _EIaccount.contract.FilterLogs(opts, "PostBanner")
	if err != nil {
		return nil, err
	}
	return &EIaccountPostBannerIterator{contract: _EIaccount.contract, event: "PostBanner", logs: logs, sub: sub}, nil
}

// WatchPostBanner is a free log subscription operation binding the contract event 0x303dbb34abb7589def7f6abbd68426bbf4f4a3f6062603673caedc6edc144184.
//
// Solidity: event PostBanner(address User, uint256 id, string Hash, string PostDescription)
func (_EIaccount *EIaccountFilterer) WatchPostBanner(opts *bind.WatchOpts, sink chan<- *EIaccountPostBanner) (event.Subscription, error) {

	logs, sub, err := _EIaccount.contract.WatchLogs(opts, "PostBanner")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EIaccountPostBanner)
				if err := _EIaccount.contract.UnpackLog(event, "PostBanner", log); err != nil {
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

// ParsePostBanner is a log parse operation binding the contract event 0x303dbb34abb7589def7f6abbd68426bbf4f4a3f6062603673caedc6edc144184.
//
// Solidity: event PostBanner(address User, uint256 id, string Hash, string PostDescription)
func (_EIaccount *EIaccountFilterer) ParsePostBanner(log types.Log) (*EIaccountPostBanner, error) {
	event := new(EIaccountPostBanner)
	if err := _EIaccount.contract.UnpackLog(event, "PostBanner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
