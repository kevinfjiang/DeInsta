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
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"Name_\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"PostNumber\",\"type\":\"uint256\"}],\"name\":\"PostBanner\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"targ\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"PostNumber\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"CommentString\",\"type\":\"string\"}],\"name\":\"AddComment\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"PostNumber\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"CommentString\",\"type\":\"string\"}],\"name\":\"AddCommentOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Bio\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"targ\",\"type\":\"address\"},{\"internalType\":\"int8\",\"name\":\"vote\",\"type\":\"int8\"},{\"internalType\":\"uint256\",\"name\":\"PostNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"CommentNumber\",\"type\":\"uint256\"}],\"name\":\"CommentVote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int8\",\"name\":\"vote\",\"type\":\"int8\"},{\"internalType\":\"uint256\",\"name\":\"PostNumber\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"CommentNumber\",\"type\":\"uint256\"}],\"name\":\"CommentVoteOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CountPosts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"PostURL_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"Caption_\",\"type\":\"string\"}],\"name\":\"CreatePost\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DeleteAccount\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"PostID\",\"type\":\"uint256\"}],\"name\":\"DeletePost\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"targ\",\"type\":\"address\"}],\"name\":\"Follow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"Followers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"Following\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"PostID\",\"type\":\"uint256\"}],\"name\":\"GetComments\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"CommentNumber\",\"type\":\"uint256\"},{\"internalType\":\"int16\",\"name\":\"Karma\",\"type\":\"int16\"},{\"internalType\":\"string\",\"name\":\"CommentString\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"Poster\",\"type\":\"address\"}],\"internalType\":\"structDIAccount.Comment[]\",\"name\":\"CommentList\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Karma\",\"outputs\":[{\"internalType\":\"int16\",\"name\":\"\",\"type\":\"int16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"Name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NewFollowing\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"added\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"Posts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"Account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"PostNumber\",\"type\":\"uint256\"},{\"internalType\":\"int16\",\"name\":\"Karma\",\"type\":\"int16\"},{\"internalType\":\"string\",\"name\":\"IPFSurl\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"Caption\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"CountComments\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"Bio_\",\"type\":\"string\"}],\"name\":\"SetBio\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"targ\",\"type\":\"address\"}],\"name\":\"UnFollow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UnNewFollowing\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"removed\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"targ\",\"type\":\"address\"},{\"internalType\":\"int8\",\"name\":\"vote\",\"type\":\"int8\"},{\"internalType\":\"uint256\",\"name\":\"PostNumber\",\"type\":\"uint256\"}],\"name\":\"Vote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int8\",\"name\":\"vote\",\"type\":\"int8\"},{\"internalType\":\"uint256\",\"name\":\"PostNumber\",\"type\":\"uint256\"}],\"name\":\"VoteOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
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
	Bin: "0x60806040523480156200001157600080fd5b50604051620023363803806200233683398101604081905262000034916200019a565b805162000049906000906020840190620000de565b505060006004819055600980546001600160a01b031990811633179091556005805460018181019092557f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db00180548316309081179091556007805492830181559093527fa66cc928b5edb82af9bd49922954155ab7b0942694bea4ce44661d9a8736c6880180549091169091179055620002b3565b828054620000ec9062000276565b90600052602060002090601f0160209004810192826200011057600085556200015b565b82601f106200012b57805160ff19168380011785556200015b565b828001600101855582156200015b579182015b828111156200015b5782518255916020019190600101906200013e565b50620001699291506200016d565b5090565b5b808211156200016957600081556001016200016e565b634e487b7160e01b600052604160045260246000fd5b60006020808385031215620001ae57600080fd5b82516001600160401b0380821115620001c657600080fd5b818501915085601f830112620001db57600080fd5b815181811115620001f057620001f062000184565b604051601f8201601f19908116603f011681019083821181831017156200021b576200021b62000184565b8160405282815288868487010111156200023457600080fd5b600093505b8284101562000258578484018601518185018701529285019262000239565b828411156200026a5760008684830101525b98975050505050505050565b600181811c908216806200028b57607f821691505b60208210811415620002ad57634e487b7160e01b600052602260045260246000fd5b50919050565b61207380620002c36000396000f3fe6080604052600436106101385760003560e01c80639099967e116100ab578063d9fe8e5d1161006f578063d9fe8e5d146103ed578063e06da1cb1461041a578063e5163de51461043a578063e81e97bc14610467578063ebd26c2b14610487578063edd69f01146104a757600080fd5b80639099967e1461036b578063b381ae7514610373578063c17f033f14610393578063c2d2d0f9146103b3578063d6dc0b31146103c857600080fd5b80634fafef1a116100fd5780634fafef1a146102965780636a2397e9146102c95780636e493c29146102e957806375c16cb2146103095780637a8d157e146103295780638052474d1461034957600080fd5b806249c1ed146101d35780632520442814610210578063276773431461023257806327d0b9ec1461025657806335592d901461027657600080fd5b366101ce57600034116101925760405162461bcd60e51b815260206004820152601a60248201527f53756d206d7573742062652067726561746572207468616e203000000000000060448201526064015b60405180910390fd5b6009546040516001600160a01b03909116903480156108fc02916000818181858888f193505050501580156101cb573d6000803e3d6000fd5b50005b600080fd5b3480156101df57600080fd5b506101f36101ee366004611a29565b6104bc565b6040516001600160a01b0390911681526020015b60405180910390f35b34801561021c57600080fd5b5061023061022b366004611ae5565b6104e6565b005b34801561023e57600080fd5b5061024860045481565b604051908152602001610207565b34801561026257600080fd5b50610230610271366004611b2c565b61068e565b34801561028257600080fd5b50610230610291366004611b95565b6106bc565b3480156102a257600080fd5b506102b66102b1366004611a29565b610754565b6040516102079796959493929190611c20565b3480156102d557600080fd5b506102306102e4366004611c7d565b6108c3565b3480156102f557600080fd5b50610230610304366004611ca7565b610b94565b34801561031557600080fd5b50610230610324366004611a29565b610d4e565b34801561033557600080fd5b50610230610344366004611ccb565b610df8565b34801561035557600080fd5b5061035e610f23565b6040516102079190611d25565b610230610fb1565b34801561037f57600080fd5b5061023061038e366004611d38565b610fd6565b34801561039f57600080fd5b506102306103ae366004611d91565b611055565b3480156103bf57600080fd5b5061035e6110f4565b3480156103d457600080fd5b506103dd611101565b6040519015158152602001610207565b3480156103f957600080fd5b5061040d610408366004611a29565b6111f9565b6040516102079190611dd5565b34801561042657600080fd5b50610230610435366004611ca7565b611355565b34801561044657600080fd5b506003546104549060010b81565b60405160019190910b8152602001610207565b34801561047357600080fd5b506101f3610482366004611a29565b611458565b34801561049357600080fd5b506102306104a2366004611e78565b611468565b3480156104b357600080fd5b506103dd61172b565b600781815481106104cc57600080fd5b6000918252602090912001546001600160a01b0316905081565b604051633a07a5ef60e21b8152600060048201523390819063e81e97bc90602401602060405180830381865afa158015610524573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105489190611eab565b6001600160a01b03161461055b57600080fd5b6002828154811061056e5761056e611ec8565b6000918252602082206008600a9092020101805460010181558152600280548490811061059d5761059d611ec8565b90600052602060002090600a0201600801600284815481106105c1576105c1611ec8565b90600052602060002090600a020160070154815481106105e3576105e3611ec8565b906000526020600020906005020190504281600001819055506002838154811061060f5761060f611ec8565b60009182526020918290206007600a9092020101546001830155825161063d916003840191908501906118e9565b506004810180546001600160a01b03191633179055600280548490811061066657610666611ec8565b6000918252602082206007600a909202010180549161068483611ef4565b9190505550505050565b6009546001600160a01b031633146106a557600080fd5b80516106b89060019060208401906118e9565b5050565b826001600160a01b0316636a2397e960008460000b136106f05760008460000b126106e85760006106f3565b6000196106f3565b60015b6040516001600160e01b031960e084901b16815260009190910b6004820152602481018490526044015b600060405180830381600087803b15801561073757600080fd5b505af115801561074b573d6000803e3d6000fd5b50505050505050565b6002818154811061076457600080fd5b60009182526020909120600a909102018054600180830154600284015460038501546005860180549597506001600160a01b03909316959194930b929091906107ac90611f0f565b80601f01602080910402602001604051908101604052809291908181526020018280546107d890611f0f565b80156108255780601f106107fa57610100808354040283529160200191610825565b820191906000526020600020905b81548152906001019060200180831161080857829003601f168201915b50505050509080600601805461083a90611f0f565b80601f016020809104026020016040519081016040528092919081815260200182805461086690611f0f565b80156108b35780601f10610888576101008083540402835291602001916108b3565b820191906000526020600020905b81548152906001019060200180831161089657829003601f168201915b5050505050908060070154905087565b604051633a07a5ef60e21b8152600060048201523390819063e81e97bc90602401602060405180830381865afa158015610901573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109259190611eab565b6001600160a01b03161461093857600080fd5b6002818154811061094b5761094b611ec8565b600091825260208083203384526004600a909302019190910190526040812054810b90839081900b8214156109c25760405162461bcd60e51b815260206004820152601a60248201527f596f75206d757374206368616e676520796f757220766f7465210000000000006044820152606401610189565b600283815481106109d5576109d5611ec8565b600091825260208083203384526004600a909302019190910190526040812054600280549190920b919085908110610a0f57610a0f611ec8565b6000918252602082206003600a90920201018054909190610a3490849060010b611f4a565b92506101000a81548161ffff021916908360010b61ffff16021790555060028381548110610a6457610a64611ec8565b600091825260208083203384526004600a9093020191909101905260408120546003805491830b929091610a9c90849060010b611f4a565b92506101000a81548161ffff021916908360010b61ffff1602179055508360000b60028481548110610ad057610ad0611ec8565b6000918252602082206003600a90920201018054909190610af590849060010b611f90565b825461ffff9182166101009390930a92830291909202199091161790555060038054600086810b9291610b2c90849060010b611f90565b92506101000a81548161ffff021916908360010b61ffff1602179055508360028481548110610b5d57610b5d611ec8565b60009182526020808320338452600a92909202909101600401905260409020805460ff191660ff9290921691909117905550505050565b6001600160a01b03811660009081526006602052604090205415801590610c1a5750806001600160a01b031663edd69f016040518163ffffffff1660e01b81526004016020604051808303816000875af1158015610bf6573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c1a9190611fd5565b610c2357600080fd5b60058054610c3390600190611ff7565b81548110610c4357610c43611ec8565b60009182526020808320909101546001600160a01b0384811684526006909252604090922054600580549290931692918110610c8157610c81611ec8565b600091825260208083209190910180546001600160a01b0319166001600160a01b03948516179055918316815260069182905260408120546005805491939291610ccd90600190611ff7565b81548110610cdd57610cdd611ec8565b60009182526020808320909101546001600160a01b03908116845283820194909452604092830182209490945591841682526006909252908120556005805480610d2957610d2961200e565b600082815260209020810160001990810180546001600160a01b031916905501905550565b6009546001600160a01b03163314610d6557600080fd5b60028181548110610d7857610d78611ec8565b600091825260208220600a909102018181556001810180546001600160a01b03191690556002810182905560038101805461ffff1916905590610dbe600583018261196d565b610dcc60068301600061196d565b6007820160009055600882016000610de491906119aa565b610df360098301600080825552565b505050565b6009546001600160a01b03163314610e0f57600080fd5b6032815110610e605760405162461bcd60e51b815260206004820152601c60248201527f506f7374206578636565647320636861726163746572206c696d6974000000006044820152606401610189565b6002805460010180825560008281526004549092918110610e8357610e83611ec8565b600091825260209182902042600a909202019081556001810180546001600160a01b0319163017905560045460028201558451909250610ecb916005840191908601906118e9565b508151610ee190600683019060208501906118e9565b5060045481546040513091907f6e8001e55f8223a69373fdb10aecda8ca632280f2f689cbf7365b8c728e5187890600090a46004805490600061068483611ef4565b60008054610f3090611f0f565b80601f0160208091040260200160405190810160405280929190818152602001828054610f5c90611f0f565b8015610fa95780601f10610f7e57610100808354040283529160200191610fa9565b820191906000526020600020905b815481529060010190602001808311610f8c57829003601f168201915b505050505081565b6009546001600160a01b03163314610fc857600080fd5b6009546001600160a01b0316ff5b60328151106110275760405162461bcd60e51b815260206004820152601c60248201527f506f7374206578636565647320636861726163746572206c696d6974000000006044820152606401610189565b6040516304a4088560e31b81526001600160a01b0384169063252044289061071d9085908590600401612024565b836001600160a01b031663ebd26c2b60008560000b136110895760008560000b1261108157600061108c565b60001961108c565b60015b6040516001600160e01b031960e084901b16815260009190910b60048201526024810185905260448101849052606401600060405180830381600087803b1580156110d657600080fd5b505af11580156110ea573d6000803e3d6000fd5b5050505050505050565b60018054610f3090611f0f565b604051633a07a5ef60e21b8152600060048201819052903390819063e81e97bc90602401602060405180830381865afa158015611142573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111669190611eab565b6001600160a01b03161461117957600080fd5b336000908152600860205260409020541561119357600080fd5b60078054600181810183556000929092527fa66cc928b5edb82af9bd49922954155ab7b0942694bea4ce44661d9a8736c6880180546001600160a01b031916331790556005546111e39190611ff7565b3360009081526008602052604090205550600190565b60606002828154811061120e5761120e611ec8565b90600052602060002090600a0201600801805480602002602001604051908101604052809291908181526020016000905b8282101561134a57838290600052602060002090600502016040518060a001604052908160008201548152602001600182015481526020016002820160009054906101000a900460010b60010b60010b81526020016003820180546112a390611f0f565b80601f01602080910402602001604051908101604052809291908181526020018280546112cf90611f0f565b801561131c5780601f106112f15761010080835404028352916020019161131c565b820191906000526020600020905b8154815290600101906020018083116112ff57829003601f168201915b5050509183525050600491909101546001600160a01b0316602091820152908252600192909201910161123f565b505050509050919050565b6001600160a01b0381166000908152600660205260409020541580156113da5750806001600160a01b031663d6dc0b316040518163ffffffff1660e01b81526004016020604051808303816000875af11580156113b6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113da9190611fd5565b6113e357600080fd5b600580546001808201835560008390527f036b6384b5eca791c62761152d0c79bb0604c104a5fb6f4eb0703f3154bb3db090910180546001600160a01b0319166001600160a01b038516179055905461143c9190611ff7565b6001600160a01b03909116600090815260066020526040902055565b600581815481106104cc57600080fd5b604051633a07a5ef60e21b8152600060048201523390819063e81e97bc90602401602060405180830381865afa1580156114a6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114ca9190611eab565b6001600160a01b0316146114dd57600080fd5b600282815481106114f0576114f0611ec8565b90600052602060002090600a0201600901818154811061151257611512611ec8565b6000918252602080832033845290910190526040812054810b90849081900b8214156115805760405162461bcd60e51b815260206004820152601a60248201527f596f75206d757374206368616e676520796f757220766f7465210000000000006044820152606401610189565b6002848154811061159357611593611ec8565b90600052602060002090600a020160090183815481106115b5576115b5611ec8565b6000918252602080832033845290910190526040812054600280549190920b9190869081106115e6576115e6611ec8565b90600052602060002090600a0201600801848154811061160857611608611ec8565b600091825260208220600260059092020101805490919061162d90849060010b611f4a565b92506101000a81548161ffff021916908360010b61ffff1602179055508460000b6002858154811061166157611661611ec8565b90600052602060002090600a0201600801848154811061168357611683611ec8565b60009182526020822060026005909202010180549091906116a890849060010b611f90565b92506101000a81548161ffff021916908360010b61ffff16021790555084600285815481106116d9576116d9611ec8565b90600052602060002090600a020160090184815481106116fb576116fb611ec8565b6000918252602080832033845291909101905260409020805460ff191660ff929092169190911790555050505050565b604051633a07a5ef60e21b8152600060048201819052903390819063e81e97bc90602401602060405180830381865afa15801561176c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906117909190611eab565b6001600160a01b0316146117a357600080fd5b336000908152600860205260409020546117bc57600080fd5b6005546007906117ce90600190611ff7565b815481106117de576117de611ec8565b60009182526020808320909101543383526008909152604090912054600780546001600160a01b0390931692909190811061181b5761181b611ec8565b600091825260208083209190910180546001600160a01b0319166001600160a01b0394909416939093179092553381526008918290526040812054600780549193929161186a90600190611ff7565b8154811061187a5761187a611ec8565b60009182526020808320909101546001600160a01b03168352828101939093526040918201812093909355338352600890915281205560078054806118c1576118c161200e565b600082815260209020810160001990810180546001600160a01b031916905501905550600190565b8280546118f590611f0f565b90600052602060002090601f016020900481019282611917576000855561195d565b82601f1061193057805160ff191683800117855561195d565b8280016001018555821561195d579182015b8281111561195d578251825591602001919060010190611942565b506119699291506119cb565b5090565b50805461197990611f0f565b6000825580601f10611989575050565b601f0160209004906000526020600020908101906119a791906119cb565b50565b50805460008255600502906000526020600020908101906119a791906119e0565b5b8082111561196957600081556001016119cc565b808211156119695760008082556001820181905560028201805461ffff19169055611a0e600383018261196d565b506004810180546001600160a01b03191690556005016119e0565b600060208284031215611a3b57600080fd5b5035919050565b634e487b7160e01b600052604160045260246000fd5b600082601f830112611a6957600080fd5b813567ffffffffffffffff80821115611a8457611a84611a42565b604051601f8301601f19908116603f01168101908282118183101715611aac57611aac611a42565b81604052838152866020858801011115611ac557600080fd5b836020870160208301376000602085830101528094505050505092915050565b60008060408385031215611af857600080fd5b82359150602083013567ffffffffffffffff811115611b1657600080fd5b611b2285828601611a58565b9150509250929050565b600060208284031215611b3e57600080fd5b813567ffffffffffffffff811115611b5557600080fd5b611b6184828501611a58565b949350505050565b6001600160a01b03811681146119a757600080fd5b8035600081900b8114611b9057600080fd5b919050565b600080600060608486031215611baa57600080fd5b8335611bb581611b69565b9250611bc360208501611b7e565b9150604084013590509250925092565b6000815180845260005b81811015611bf957602081850181015186830182015201611bdd565b81811115611c0b576000602083870101525b50601f01601f19169290920160200192915050565b87815260018060a01b03871660208201528560408201528460010b606082015260e060808201526000611c5660e0830186611bd3565b82810360a0840152611c688186611bd3565b9150508260c083015298975050505050505050565b60008060408385031215611c9057600080fd5b611c9983611b7e565b946020939093013593505050565b600060208284031215611cb957600080fd5b8135611cc481611b69565b9392505050565b60008060408385031215611cde57600080fd5b823567ffffffffffffffff80821115611cf657600080fd5b611d0286838701611a58565b93506020850135915080821115611d1857600080fd5b50611b2285828601611a58565b602081526000611cc46020830184611bd3565b600080600060608486031215611d4d57600080fd5b8335611d5881611b69565b925060208401359150604084013567ffffffffffffffff811115611d7b57600080fd5b611d8786828701611a58565b9150509250925092565b60008060008060808587031215611da757600080fd5b8435611db281611b69565b9350611dc060208601611b7e565b93969395505050506040820135916060013590565b60006020808301818452808551808352604092508286019150828160051b87010184880160005b83811015611e6a57603f19898403018552815160a08151855288820151898601528782015160010b888601526060808301518282880152611e3f83880182611bd3565b6080948501516001600160a01b03169790940196909652505094870194925090860190600101611dfc565b509098975050505050505050565b600080600060608486031215611e8d57600080fd5b611e9684611b7e565b95602085013595506040909401359392505050565b600060208284031215611ebd57600080fd5b8151611cc481611b69565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b6000600019821415611f0857611f08611ede565b5060010190565b600181811c90821680611f2357607f821691505b60208210811415611f4457634e487b7160e01b600052602260045260246000fd5b50919050565b60008160010b8360010b6000811281617fff1901831281151615611f7057611f70611ede565b81617fff018313811615611f8657611f86611ede565b5090039392505050565b60008160010b8360010b6000821282617fff03821381151615611fb557611fb5611ede565b82617fff19038212811615611fcc57611fcc611ede565b50019392505050565b600060208284031215611fe757600080fd5b81518015158114611cc457600080fd5b60008282101561200957612009611ede565b500390565b634e487b7160e01b600052603160045260246000fd5b828152604060208201526000611b616040830184611bd356fea26469706673582212200bad65042f6b8c8b7a10b33db891f4ad0b5d1526331ec076686c75209b7a9f9864736f6c634300080a0033",
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

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_DIAccount *DIAccountTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DIAccount.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_DIAccount *DIAccountSession) Receive() (*types.Transaction, error) {
	return _DIAccount.Contract.Receive(&_DIAccount.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_DIAccount *DIAccountTransactorSession) Receive() (*types.Transaction, error) {
	return _DIAccount.Contract.Receive(&_DIAccount.TransactOpts)
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
