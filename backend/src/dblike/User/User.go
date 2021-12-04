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
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"Name_\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"PostNumber\",\"type\":\"uint256\"}],\"name\":\"PostBanner\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"targ\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"PostNumber\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"CommentString\",\"type\":\"string\"}],\"name\":\"AddComment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"PostNumber\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"CommentString\",\"type\":\"string\"}],\"name\":\"AddCommentOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Bio\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"targ\",\"type\":\"address\"},{\"internalType\":\"int8\",\"name\":\"vote\",\"type\":\"int8\"},{\"internalType\":\"uint256\",\"name\":\"PostNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"CommentNumber\",\"type\":\"uint256\"}],\"name\":\"CommentVote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int8\",\"name\":\"vote\",\"type\":\"int8\"},{\"internalType\":\"uint256\",\"name\":\"PostNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"CommentNumber\",\"type\":\"uint256\"}],\"name\":\"CommentVoteOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CountPosts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"PostURL_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"Caption_\",\"type\":\"string\"}],\"name\":\"CreatePost\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DeleteAccount\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"PostID\",\"type\":\"uint256\"}],\"name\":\"DeletePost\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"targ\",\"type\":\"address\"}],\"name\":\"Follow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"Followers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"Following\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"PostID\",\"type\":\"uint256\"}],\"name\":\"GetComments\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"CommentNumber\",\"type\":\"uint256\"},{\"internalType\":\"int16\",\"name\":\"Karma\",\"type\":\"int16\"},{\"internalType\":\"string\",\"name\":\"CommentString\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"Poster\",\"type\":\"address\"}],\"internalType\":\"structDIAccount.Comment[]\",\"name\":\"CommentList\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Karma\",\"outputs\":[{\"internalType\":\"int16\",\"name\":\"\",\"type\":\"int16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NewFollowing\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"added\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"Posts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"Account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"PostNumber\",\"type\":\"uint256\"},{\"internalType\":\"int16\",\"name\":\"Karma\",\"type\":\"int16\"},{\"internalType\":\"string\",\"name\":\"IPFSurl\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"Caption\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"CountComments\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"Bio_\",\"type\":\"string\"}],\"name\":\"SetBio\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"targ\",\"type\":\"address\"}],\"name\":\"UnFollow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UnNewFollowing\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"removed\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"targ\",\"type\":\"address\"},{\"internalType\":\"int8\",\"name\":\"vote\",\"type\":\"int8\"},{\"internalType\":\"uint256\",\"name\":\"PostNumber\",\"type\":\"uint256\"}],\"name\":\"Vote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int8\",\"name\":\"vote\",\"type\":\"int8\"},{\"internalType\":\"uint256\",\"name\":\"PostNumber\",\"type\":\"uint256\"}],\"name\":\"VoteOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"b381ae75": "AddComment(address,uint256,string)",
		"25204428": "AddCommentOwner(uint256,string)",
		"c2d2d0f9": "Bio()",
		"c17f033f": "CommentVote(address,int8,uint256,uint256)",
		"ebd26c2b": "CommentVoteOwner(int8,uint256,uint256)",
		"27677343": "CountPosts()",
		"7a8d157e": "CreatePost(string,string)",
		"9099967e": "DeleteAccount()",
		"75c16cb2": "DeletePost(uint256)",
		"e06da1cb": "Follow(address)",
		"e81e97bc": "Followers(uint256)",
		"0049c1ed": "Following(uint256)",
		"d9fe8e5d": "GetComments(uint256)",
		"e5163de5": "Karma()",
		"8052474d": "Name()",
		"d6dc0b31": "NewFollowing()",
		"4fafef1a": "Posts(uint256)",
		"27d0b9ec": "SetBio(string)",
		"6e493c29": "UnFollow(address)",
		"edd69f01": "UnNewFollowing()",
		"35592d90": "Vote(address,int8,uint256)",
		"6a2397e9": "VoteOwner(int8,uint256)",
	},
	Bin: "0x60806040523480156200001157600080fd5b506040516200205238038062002052833981016040819052620000349162000198565b805162000049906000906020840190620000dc565b50506000600481905560098054336001600160a01b031991821681179092556005805460018181019092557f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db00180548316841790556007805491820181559093527fa66cc928b5edb82af9bd49922954155ab7b0942694bea4ce44661d9a8736c6889092018054909216179055620002b1565b828054620000ea9062000274565b90600052602060002090601f0160209004810192826200010e576000855562000159565b82601f106200012957805160ff191683800117855562000159565b8280016001018555821562000159579182015b82811115620001595782518255916020019190600101906200013c565b50620001679291506200016b565b5090565b5b808211156200016757600081556001016200016c565b634e487b7160e01b600052604160045260246000fd5b60006020808385031215620001ac57600080fd5b82516001600160401b0380821115620001c457600080fd5b818501915085601f830112620001d957600080fd5b815181811115620001ee57620001ee62000182565b604051601f8201601f19908116603f0116810190838211818310171562000219576200021962000182565b8160405282815288868487010111156200023257600080fd5b600093505b8284101562000256578484018601518185018701529285019262000237565b82841115620002685760008684830101525b98975050505050505050565b600181811c908216806200028957607f821691505b60208210811415620002ab57634e487b7160e01b600052602260045260246000fd5b50919050565b611d9180620002c16000396000f3fe6080604052600436106101345760003560e01c80639099967e116100ab578063d9fe8e5d1161006f578063d9fe8e5d14610353578063e06da1cb14610380578063e5163de5146103a0578063e81e97bc146103cd578063ebd26c2b146103ed578063edd69f011461040d57600080fd5b80639099967e146102d1578063b381ae75146102d9578063c17f033f146102f9578063c2d2d0f914610319578063d6dc0b311461032e57600080fd5b80634fafef1a116100fd5780634fafef1a146101fc5780636a2397e91461022f5780636e493c291461024f57806375c16cb21461026f5780637a8d157e1461028f5780638052474d146102af57600080fd5b806249c1ed146101395780632520442814610176578063276773431461019857806327d0b9ec146101bc57806335592d90146101dc575b600080fd5b34801561014557600080fd5b50610159610154366004611747565b610422565b6040516001600160a01b0390911681526020015b60405180910390f35b34801561018257600080fd5b50610196610191366004611803565b61044c565b005b3480156101a457600080fd5b506101ae60045481565b60405190815260200161016d565b3480156101c857600080fd5b506101966101d736600461184a565b6105c7565b3480156101e857600080fd5b506101966101f73660046118b3565b6105f5565b34801561020857600080fd5b5061021c610217366004611747565b61070c565b60405161016d979695949392919061193e565b34801561023b57600080fd5b5061019661024a36600461199b565b61087b565b34801561025b57600080fd5b5061019661026a3660046119c5565b610ac0565b34801561027b57600080fd5b5061019661028a366004611747565b610c2c565b34801561029b57600080fd5b506101966102aa3660046119e9565b610cd6565b3480156102bb57600080fd5b506102c4610df5565b60405161016d9190611a43565b610196610e83565b3480156102e557600080fd5b506101966102f4366004611a56565b610ea8565b34801561030557600080fd5b50610196610314366004611aaf565b610f27565b34801561032557600080fd5b506102c4611059565b34801561033a57600080fd5b50610343611066565b604051901515815260200161016d565b34801561035f57600080fd5b5061037361036e366004611747565b61115e565b60405161016d9190611af3565b34801561038c57600080fd5b5061019661039b3660046119c5565b6112ba565b3480156103ac57600080fd5b506003546103ba9060010b81565b60405160019190910b815260200161016d565b3480156103d957600080fd5b506101596103e8366004611747565b6113bd565b3480156103f957600080fd5b50610196610408366004611b96565b6113cd565b34801561041957600080fd5b50610343611576565b6007818154811061043257600080fd5b6000918252602090912001546001600160a01b0316905081565b604051633a07a5ef60e21b8152600060048201523390819063e81e97bc90602401602060405180830381865afa15801561048a573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104ae9190611bc9565b6001600160a01b0316146104c157600080fd5b6000600283815481106104d6576104d6611be6565b90600052602060002090600a0201600801600284815481106104fa576104fa611be6565b90600052602060002090600a0201600701548154811061051c5761051c611be6565b906000526020600020906005020190504281600001819055506002838154811061054857610548611be6565b60009182526020918290206007600a9092020101546001830155825161057691600384019190850190611607565b506004810180546001600160a01b03191633179055600280548490811061059f5761059f611be6565b6000918252602082206007600a90920201018054916105bd83611c12565b9190505550505050565b6009546001600160a01b031633146105de57600080fd5b80516105f1906001906020840190611607565b5050565b60008260000b1361061a5760008260000b1261061257600061061d565b60001961061d565b60015b91508160000b6002828154811061063657610636611be6565b600091825260208083203384526004600a909302019190910190526040812054900b14156106a25760405162461bcd60e51b8152602060048201526014602482015273165bdd49dd9948185b1c9958591e481d9bdd195960621b60448201526064015b60405180910390fd5b604051636a2397e960e01b8152600083900b6004820152602481018290526001600160a01b03841690636a2397e9906044015b600060405180830381600087803b1580156106ef57600080fd5b505af1158015610703573d6000803e3d6000fd5b50505050505050565b6002818154811061071c57600080fd5b60009182526020909120600a909102018054600180830154600284015460038501546005860180549597506001600160a01b03909316959194930b9290919061076490611c2d565b80601f016020809104026020016040519081016040528092919081815260200182805461079090611c2d565b80156107dd5780601f106107b2576101008083540402835291602001916107dd565b820191906000526020600020905b8154815290600101906020018083116107c057829003601f168201915b5050505050908060060180546107f290611c2d565b80601f016020809104026020016040519081016040528092919081815260200182805461081e90611c2d565b801561086b5780601f106108405761010080835404028352916020019161086b565b820191906000526020600020905b81548152906001019060200180831161084e57829003601f168201915b5050505050908060070154905087565b604051633a07a5ef60e21b8152600060048201523390819063e81e97bc90602401602060405180830381865afa1580156108b9573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108dd9190611bc9565b6001600160a01b0316146108f057600080fd5b6002818154811061090357610903611be6565b600091825260208083203384526004600a909302019190910190526040812054600280549190920b91908390811061093d5761093d611be6565b6000918252602082206003600a9092020101805490919061096290849060010b611c68565b92506101000a81548161ffff021916908360010b61ffff1602179055506002818154811061099257610992611be6565b600091825260208083203384526004600a9093020191909101905260408120546003805491830b9290916109ca90849060010b611c68565b92506101000a81548161ffff021916908360010b61ffff1602179055508160000b600282815481106109fe576109fe611be6565b6000918252602082206003600a90920201018054909190610a2390849060010b611cae565b825461ffff9182166101009390930a92830291909202199091161790555060038054600084810b9291610a5a90849060010b611cae565b92506101000a81548161ffff021916908360010b61ffff1602179055508160028281548110610a8b57610a8b611be6565b60009182526020808320338452600a92909202909101600401905260409020805460ff191660ff929092169190911790555050565b6001600160a01b03811660009081526006602052604090205415801590610b465750806001600160a01b031663edd69f016040518163ffffffff1660e01b81526004016020604051808303816000875af1158015610b22573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b469190611cf3565b610b4f57600080fd5b60058054610b5f90600190611d15565b81548110610b6f57610b6f611be6565b60009182526020808320909101546001600160a01b0384811684526006909252604090922054600580549290931692918110610bad57610bad611be6565b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b031602179055506005805480610bec57610bec611d2c565b60008281526020808220830160001990810180546001600160a01b03191690559092019092556001600160a01b0392909216815260089091526040812055565b6009546001600160a01b03163314610c4357600080fd5b60028181548110610c5657610c56611be6565b600091825260208220600a909102018181556001810180546001600160a01b03191690556002810182905560038101805461ffff1916905590610c9c600583018261168b565b610caa60068301600061168b565b6007820160009055600882016000610cc291906116c8565b610cd160098301600080825552565b505050565b6009546001600160a01b03163314610ced57600080fd5b6032815110610d3e5760405162461bcd60e51b815260206004820152601c60248201527f506f7374206578636565647320636861726163746572206c696d6974000000006044820152606401610699565b6000600260045481548110610d5557610d55611be6565b600091825260209182902042600a909202019081556001810180546001600160a01b0319163017905560045460028201558451909250610d9d91600584019190860190611607565b508151610db39060068301906020850190611607565b5060045481546040513091907f6e8001e55f8223a69373fdb10aecda8ca632280f2f689cbf7365b8c728e5187890600090a4600480549060006105bd83611c12565b60008054610e0290611c2d565b80601f0160208091040260200160405190810160405280929190818152602001828054610e2e90611c2d565b8015610e7b5780601f10610e5057610100808354040283529160200191610e7b565b820191906000526020600020905b815481529060010190602001808311610e5e57829003601f168201915b505050505081565b6009546001600160a01b03163314610e9a57600080fd5b6009546001600160a01b0316ff5b6032815110610ef95760405162461bcd60e51b815260206004820152601c60248201527f506f7374206578636565647320636861726163746572206c696d6974000000006044820152606401610699565b6040516304a4088560e31b81526001600160a01b038416906325204428906106d59085908590600401611d42565b60008360000b13610f4c5760008360000b12610f44576000610f4f565b600019610f4f565b60015b92508260000b60028381548110610f6857610f68611be6565b90600052602060002090600a02016009018281548110610f8a57610f8a611be6565b6000918252602080832033845290910190526040812054900b1415610fe85760405162461bcd60e51b8152602060048201526014602482015273165bdd49dd9948185b1c9958591e481d9bdd195960621b6044820152606401610699565b60405163ebd26c2b60e01b8152600084900b600482015260248101839052604481018290526001600160a01b0385169063ebd26c2b90606401600060405180830381600087803b15801561103b57600080fd5b505af115801561104f573d6000803e3d6000fd5b5050505050505050565b60018054610e0290611c2d565b604051633a07a5ef60e21b8152600060048201819052903390819063e81e97bc90602401602060405180830381865afa1580156110a7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110cb9190611bc9565b6001600160a01b0316146110de57600080fd5b33600090815260086020526040902054156110f857600080fd5b60078054600181810183556000929092527fa66cc928b5edb82af9bd49922954155ab7b0942694bea4ce44661d9a8736c6880180546001600160a01b031916331790556005546111489190611d15565b3360009081526008602052604090205550600190565b60606002828154811061117357611173611be6565b90600052602060002090600a0201600801805480602002602001604051908101604052809291908181526020016000905b828210156112af57838290600052602060002090600502016040518060a001604052908160008201548152602001600182015481526020016002820160009054906101000a900460010b60010b60010b815260200160038201805461120890611c2d565b80601f016020809104026020016040519081016040528092919081815260200182805461123490611c2d565b80156112815780601f1061125657610100808354040283529160200191611281565b820191906000526020600020905b81548152906001019060200180831161126457829003601f168201915b5050509183525050600491909101546001600160a01b031660209182015290825260019290920191016111a4565b505050509050919050565b6001600160a01b03811660009081526006602052604090205415801561133f5750806001600160a01b031663d6dc0b316040518163ffffffff1660e01b81526004016020604051808303816000875af115801561131b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061133f9190611cf3565b61134857600080fd5b600580546001808201835560008390527f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db090910180546001600160a01b0319166001600160a01b03851617905590546113a19190611d15565b6001600160a01b03909116600090815260066020526040902055565b6005818154811061043257600080fd5b600282815481106113e0576113e0611be6565b90600052602060002090600a0201600901818154811061140257611402611be6565b6000918252602080832033845290910190526040812054600280549190920b91908490811061143357611433611be6565b90600052602060002090600a0201600801828154811061145557611455611be6565b600091825260208220600260059092020101805490919061147a90849060010b611c68565b92506101000a81548161ffff021916908360010b61ffff1602179055508260000b600283815481106114ae576114ae611be6565b90600052602060002090600a020160080182815481106114d0576114d0611be6565b60009182526020822060026005909202010180549091906114f590849060010b611cae565b92506101000a81548161ffff021916908360010b61ffff160217905550826002838154811061152657611526611be6565b90600052602060002090600a0201600901828154811061154857611548611be6565b6000918252602080832033845291909101905260409020805460ff191660ff92909216919091179055505050565b604051633a07a5ef60e21b8152600060048201819052903390819063e81e97bc90602401602060405180830381865afa1580156115b7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906115db9190611bc9565b6001600160a01b0316146115ee57600080fd5b336000908152600860205260409020546110f857600080fd5b82805461161390611c2d565b90600052602060002090601f016020900481019282611635576000855561167b565b82601f1061164e57805160ff191683800117855561167b565b8280016001018555821561167b579182015b8281111561167b578251825591602001919060010190611660565b506116879291506116e9565b5090565b50805461169790611c2d565b6000825580601f106116a7575050565b601f0160209004906000526020600020908101906116c591906116e9565b50565b50805460008255600502906000526020600020908101906116c591906116fe565b5b8082111561168757600081556001016116ea565b808211156116875760008082556001820181905560028201805461ffff1916905561172c600383018261168b565b506004810180546001600160a01b03191690556005016116fe565b60006020828403121561175957600080fd5b5035919050565b634e487b7160e01b600052604160045260246000fd5b600082601f83011261178757600080fd5b813567ffffffffffffffff808211156117a2576117a2611760565b604051601f8301601f19908116603f011681019082821181831017156117ca576117ca611760565b816040528381528660208588010111156117e357600080fd5b836020870160208301376000602085830101528094505050505092915050565b6000806040838503121561181657600080fd5b82359150602083013567ffffffffffffffff81111561183457600080fd5b61184085828601611776565b9150509250929050565b60006020828403121561185c57600080fd5b813567ffffffffffffffff81111561187357600080fd5b61187f84828501611776565b949350505050565b6001600160a01b03811681146116c557600080fd5b8035600081900b81146118ae57600080fd5b919050565b6000806000606084860312156118c857600080fd5b83356118d381611887565b92506118e16020850161189c565b9150604084013590509250925092565b6000815180845260005b81811015611917576020818501810151868301820152016118fb565b81811115611929576000602083870101525b50601f01601f19169290920160200192915050565b87815260018060a01b03871660208201528560408201528460010b606082015260e06080820152600061197460e08301866118f1565b82810360a084015261198681866118f1565b9150508260c083015298975050505050505050565b600080604083850312156119ae57600080fd5b6119b78361189c565b946020939093013593505050565b6000602082840312156119d757600080fd5b81356119e281611887565b9392505050565b600080604083850312156119fc57600080fd5b823567ffffffffffffffff80821115611a1457600080fd5b611a2086838701611776565b93506020850135915080821115611a3657600080fd5b5061184085828601611776565b6020815260006119e260208301846118f1565b600080600060608486031215611a6b57600080fd5b8335611a7681611887565b925060208401359150604084013567ffffffffffffffff811115611a9957600080fd5b611aa586828701611776565b9150509250925092565b60008060008060808587031215611ac557600080fd5b8435611ad081611887565b9350611ade6020860161189c565b93969395505050506040820135916060013590565b60006020808301818452808551808352604092508286019150828160051b87010184880160005b83811015611b8857603f19898403018552815160a08151855288820151898601528782015160010b888601526060808301518282880152611b5d838801826118f1565b6080948501516001600160a01b03169790940196909652505094870194925090860190600101611b1a565b509098975050505050505050565b600080600060608486031215611bab57600080fd5b611bb48461189c565b95602085013595506040909401359392505050565b600060208284031215611bdb57600080fd5b81516119e281611887565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b6000600019821415611c2657611c26611bfc565b5060010190565b600181811c90821680611c4157607f821691505b60208210811415611c6257634e487b7160e01b600052602260045260246000fd5b50919050565b60008160010b8360010b6000811281617fff1901831281151615611c8e57611c8e611bfc565b81617fff018313811615611ca457611ca4611bfc565b5090039392505050565b60008160010b8360010b6000821282617fff03821381151615611cd357611cd3611bfc565b82617fff19038212811615611cea57611cea611bfc565b50019392505050565b600060208284031215611d0557600080fd5b815180151581146119e257600080fd5b600082821015611d2757611d27611bfc565b500390565b634e487b7160e01b600052603160045260246000fd5b82815260406020820152600061187f60408301846118f156fea26469706673582212203dfe5fb3810d92f13b854e8e77c5aa6cfffb1b30865ddfc3ce71a6f0e69e669f64736f6c634300080a0033",
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

// Followers is a free data retrieval call binding the contract method 0xe81e97bc.
//
// Solidity: function Followers(uint256 ) view returns(address)
func (_DIAccount *DIAccountCaller) Followers(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _DIAccount.contract.Call(opts, &out, "Followers", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Followers is a free data retrieval call binding the contract method 0xe81e97bc.
//
// Solidity: function Followers(uint256 ) view returns(address)
func (_DIAccount *DIAccountSession) Followers(arg0 *big.Int) (common.Address, error) {
	return _DIAccount.Contract.Followers(&_DIAccount.CallOpts, arg0)
}

// Followers is a free data retrieval call binding the contract method 0xe81e97bc.
//
// Solidity: function Followers(uint256 ) view returns(address)
func (_DIAccount *DIAccountCallerSession) Followers(arg0 *big.Int) (common.Address, error) {
	return _DIAccount.Contract.Followers(&_DIAccount.CallOpts, arg0)
}

// Following is a free data retrieval call binding the contract method 0x0049c1ed.
//
// Solidity: function Following(uint256 ) view returns(address)
func (_DIAccount *DIAccountCaller) Following(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _DIAccount.contract.Call(opts, &out, "Following", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Following is a free data retrieval call binding the contract method 0x0049c1ed.
//
// Solidity: function Following(uint256 ) view returns(address)
func (_DIAccount *DIAccountSession) Following(arg0 *big.Int) (common.Address, error) {
	return _DIAccount.Contract.Following(&_DIAccount.CallOpts, arg0)
}

// Following is a free data retrieval call binding the contract method 0x0049c1ed.
//
// Solidity: function Following(uint256 ) view returns(address)
func (_DIAccount *DIAccountCallerSession) Following(arg0 *big.Int) (common.Address, error) {
	return _DIAccount.Contract.Following(&_DIAccount.CallOpts, arg0)
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

// AddComment is a paid mutator transaction binding the contract method 0xb381ae75.
//
// Solidity: function AddComment(address targ, uint256 PostNumber, string CommentString) returns()
func (_DIAccount *DIAccountTransactor) AddComment(opts *bind.TransactOpts, targ common.Address, PostNumber *big.Int, CommentString string) (*types.Transaction, error) {
	return _DIAccount.contract.Transact(opts, "AddComment", targ, PostNumber, CommentString)
}

// AddComment is a paid mutator transaction binding the contract method 0xb381ae75.
//
// Solidity: function AddComment(address targ, uint256 PostNumber, string CommentString) returns()
func (_DIAccount *DIAccountSession) AddComment(targ common.Address, PostNumber *big.Int, CommentString string) (*types.Transaction, error) {
	return _DIAccount.Contract.AddComment(&_DIAccount.TransactOpts, targ, PostNumber, CommentString)
}

// AddComment is a paid mutator transaction binding the contract method 0xb381ae75.
//
// Solidity: function AddComment(address targ, uint256 PostNumber, string CommentString) returns()
func (_DIAccount *DIAccountTransactorSession) AddComment(targ common.Address, PostNumber *big.Int, CommentString string) (*types.Transaction, error) {
	return _DIAccount.Contract.AddComment(&_DIAccount.TransactOpts, targ, PostNumber, CommentString)
}

// AddCommentOwner is a paid mutator transaction binding the contract method 0x25204428.
//
// Solidity: function AddCommentOwner(uint256 PostNumber, string CommentString) returns()
func (_DIAccount *DIAccountTransactor) AddCommentOwner(opts *bind.TransactOpts, PostNumber *big.Int, CommentString string) (*types.Transaction, error) {
	return _DIAccount.contract.Transact(opts, "AddCommentOwner", PostNumber, CommentString)
}

// AddCommentOwner is a paid mutator transaction binding the contract method 0x25204428.
//
// Solidity: function AddCommentOwner(uint256 PostNumber, string CommentString) returns()
func (_DIAccount *DIAccountSession) AddCommentOwner(PostNumber *big.Int, CommentString string) (*types.Transaction, error) {
	return _DIAccount.Contract.AddCommentOwner(&_DIAccount.TransactOpts, PostNumber, CommentString)
}

// AddCommentOwner is a paid mutator transaction binding the contract method 0x25204428.
//
// Solidity: function AddCommentOwner(uint256 PostNumber, string CommentString) returns()
func (_DIAccount *DIAccountTransactorSession) AddCommentOwner(PostNumber *big.Int, CommentString string) (*types.Transaction, error) {
	return _DIAccount.Contract.AddCommentOwner(&_DIAccount.TransactOpts, PostNumber, CommentString)
}

// CommentVote is a paid mutator transaction binding the contract method 0xc17f033f.
//
// Solidity: function CommentVote(address targ, int8 vote, uint256 PostNumber, uint256 CommentNumber) returns()
func (_DIAccount *DIAccountTransactor) CommentVote(opts *bind.TransactOpts, targ common.Address, vote int8, PostNumber *big.Int, CommentNumber *big.Int) (*types.Transaction, error) {
	return _DIAccount.contract.Transact(opts, "CommentVote", targ, vote, PostNumber, CommentNumber)
}

// CommentVote is a paid mutator transaction binding the contract method 0xc17f033f.
//
// Solidity: function CommentVote(address targ, int8 vote, uint256 PostNumber, uint256 CommentNumber) returns()
func (_DIAccount *DIAccountSession) CommentVote(targ common.Address, vote int8, PostNumber *big.Int, CommentNumber *big.Int) (*types.Transaction, error) {
	return _DIAccount.Contract.CommentVote(&_DIAccount.TransactOpts, targ, vote, PostNumber, CommentNumber)
}

// CommentVote is a paid mutator transaction binding the contract method 0xc17f033f.
//
// Solidity: function CommentVote(address targ, int8 vote, uint256 PostNumber, uint256 CommentNumber) returns()
func (_DIAccount *DIAccountTransactorSession) CommentVote(targ common.Address, vote int8, PostNumber *big.Int, CommentNumber *big.Int) (*types.Transaction, error) {
	return _DIAccount.Contract.CommentVote(&_DIAccount.TransactOpts, targ, vote, PostNumber, CommentNumber)
}

// CommentVoteOwner is a paid mutator transaction binding the contract method 0xebd26c2b.
//
// Solidity: function CommentVoteOwner(int8 vote, uint256 PostNumber, uint256 CommentNumber) returns()
func (_DIAccount *DIAccountTransactor) CommentVoteOwner(opts *bind.TransactOpts, vote int8, PostNumber *big.Int, CommentNumber *big.Int) (*types.Transaction, error) {
	return _DIAccount.contract.Transact(opts, "CommentVoteOwner", vote, PostNumber, CommentNumber)
}

// CommentVoteOwner is a paid mutator transaction binding the contract method 0xebd26c2b.
//
// Solidity: function CommentVoteOwner(int8 vote, uint256 PostNumber, uint256 CommentNumber) returns()
func (_DIAccount *DIAccountSession) CommentVoteOwner(vote int8, PostNumber *big.Int, CommentNumber *big.Int) (*types.Transaction, error) {
	return _DIAccount.Contract.CommentVoteOwner(&_DIAccount.TransactOpts, vote, PostNumber, CommentNumber)
}

// CommentVoteOwner is a paid mutator transaction binding the contract method 0xebd26c2b.
//
// Solidity: function CommentVoteOwner(int8 vote, uint256 PostNumber, uint256 CommentNumber) returns()
func (_DIAccount *DIAccountTransactorSession) CommentVoteOwner(vote int8, PostNumber *big.Int, CommentNumber *big.Int) (*types.Transaction, error) {
	return _DIAccount.Contract.CommentVoteOwner(&_DIAccount.TransactOpts, vote, PostNumber, CommentNumber)
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

// Follow is a paid mutator transaction binding the contract method 0xe06da1cb.
//
// Solidity: function Follow(address targ) returns()
func (_DIAccount *DIAccountTransactor) Follow(opts *bind.TransactOpts, targ common.Address) (*types.Transaction, error) {
	return _DIAccount.contract.Transact(opts, "Follow", targ)
}

// Follow is a paid mutator transaction binding the contract method 0xe06da1cb.
//
// Solidity: function Follow(address targ) returns()
func (_DIAccount *DIAccountSession) Follow(targ common.Address) (*types.Transaction, error) {
	return _DIAccount.Contract.Follow(&_DIAccount.TransactOpts, targ)
}

// Follow is a paid mutator transaction binding the contract method 0xe06da1cb.
//
// Solidity: function Follow(address targ) returns()
func (_DIAccount *DIAccountTransactorSession) Follow(targ common.Address) (*types.Transaction, error) {
	return _DIAccount.Contract.Follow(&_DIAccount.TransactOpts, targ)
}

// NewFollowing is a paid mutator transaction binding the contract method 0xd6dc0b31.
//
// Solidity: function NewFollowing() returns(bool added)
func (_DIAccount *DIAccountTransactor) NewFollowing(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DIAccount.contract.Transact(opts, "NewFollowing")
}

// NewFollowing is a paid mutator transaction binding the contract method 0xd6dc0b31.
//
// Solidity: function NewFollowing() returns(bool added)
func (_DIAccount *DIAccountSession) NewFollowing() (*types.Transaction, error) {
	return _DIAccount.Contract.NewFollowing(&_DIAccount.TransactOpts)
}

// NewFollowing is a paid mutator transaction binding the contract method 0xd6dc0b31.
//
// Solidity: function NewFollowing() returns(bool added)
func (_DIAccount *DIAccountTransactorSession) NewFollowing() (*types.Transaction, error) {
	return _DIAccount.Contract.NewFollowing(&_DIAccount.TransactOpts)
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

// UnFollow is a paid mutator transaction binding the contract method 0x6e493c29.
//
// Solidity: function UnFollow(address targ) returns()
func (_DIAccount *DIAccountTransactor) UnFollow(opts *bind.TransactOpts, targ common.Address) (*types.Transaction, error) {
	return _DIAccount.contract.Transact(opts, "UnFollow", targ)
}

// UnFollow is a paid mutator transaction binding the contract method 0x6e493c29.
//
// Solidity: function UnFollow(address targ) returns()
func (_DIAccount *DIAccountSession) UnFollow(targ common.Address) (*types.Transaction, error) {
	return _DIAccount.Contract.UnFollow(&_DIAccount.TransactOpts, targ)
}

// UnFollow is a paid mutator transaction binding the contract method 0x6e493c29.
//
// Solidity: function UnFollow(address targ) returns()
func (_DIAccount *DIAccountTransactorSession) UnFollow(targ common.Address) (*types.Transaction, error) {
	return _DIAccount.Contract.UnFollow(&_DIAccount.TransactOpts, targ)
}

// UnNewFollowing is a paid mutator transaction binding the contract method 0xedd69f01.
//
// Solidity: function UnNewFollowing() returns(bool removed)
func (_DIAccount *DIAccountTransactor) UnNewFollowing(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DIAccount.contract.Transact(opts, "UnNewFollowing")
}

// UnNewFollowing is a paid mutator transaction binding the contract method 0xedd69f01.
//
// Solidity: function UnNewFollowing() returns(bool removed)
func (_DIAccount *DIAccountSession) UnNewFollowing() (*types.Transaction, error) {
	return _DIAccount.Contract.UnNewFollowing(&_DIAccount.TransactOpts)
}

// UnNewFollowing is a paid mutator transaction binding the contract method 0xedd69f01.
//
// Solidity: function UnNewFollowing() returns(bool removed)
func (_DIAccount *DIAccountTransactorSession) UnNewFollowing() (*types.Transaction, error) {
	return _DIAccount.Contract.UnNewFollowing(&_DIAccount.TransactOpts)
}

// Vote is a paid mutator transaction binding the contract method 0x35592d90.
//
// Solidity: function Vote(address targ, int8 vote, uint256 PostNumber) returns()
func (_DIAccount *DIAccountTransactor) Vote(opts *bind.TransactOpts, targ common.Address, vote int8, PostNumber *big.Int) (*types.Transaction, error) {
	return _DIAccount.contract.Transact(opts, "Vote", targ, vote, PostNumber)
}

// Vote is a paid mutator transaction binding the contract method 0x35592d90.
//
// Solidity: function Vote(address targ, int8 vote, uint256 PostNumber) returns()
func (_DIAccount *DIAccountSession) Vote(targ common.Address, vote int8, PostNumber *big.Int) (*types.Transaction, error) {
	return _DIAccount.Contract.Vote(&_DIAccount.TransactOpts, targ, vote, PostNumber)
}

// Vote is a paid mutator transaction binding the contract method 0x35592d90.
//
// Solidity: function Vote(address targ, int8 vote, uint256 PostNumber) returns()
func (_DIAccount *DIAccountTransactorSession) Vote(targ common.Address, vote int8, PostNumber *big.Int) (*types.Transaction, error) {
	return _DIAccount.Contract.Vote(&_DIAccount.TransactOpts, targ, vote, PostNumber)
}

// VoteOwner is a paid mutator transaction binding the contract method 0x6a2397e9.
//
// Solidity: function VoteOwner(int8 vote, uint256 PostNumber) returns()
func (_DIAccount *DIAccountTransactor) VoteOwner(opts *bind.TransactOpts, vote int8, PostNumber *big.Int) (*types.Transaction, error) {
	return _DIAccount.contract.Transact(opts, "VoteOwner", vote, PostNumber)
}

// VoteOwner is a paid mutator transaction binding the contract method 0x6a2397e9.
//
// Solidity: function VoteOwner(int8 vote, uint256 PostNumber) returns()
func (_DIAccount *DIAccountSession) VoteOwner(vote int8, PostNumber *big.Int) (*types.Transaction, error) {
	return _DIAccount.Contract.VoteOwner(&_DIAccount.TransactOpts, vote, PostNumber)
}

// VoteOwner is a paid mutator transaction binding the contract method 0x6a2397e9.
//
// Solidity: function VoteOwner(int8 vote, uint256 PostNumber) returns()
func (_DIAccount *DIAccountTransactorSession) VoteOwner(vote int8, PostNumber *big.Int) (*types.Transaction, error) {
	return _DIAccount.Contract.VoteOwner(&_DIAccount.TransactOpts, vote, PostNumber)
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
