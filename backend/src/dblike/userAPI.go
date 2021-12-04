package backendapi

import (
	"math/big"
	"errors"
	"log"
	
	"github.com/kevinfjiang/DeInsta/dblike/User"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

//TODO integrate these three packages to interact with one another, namely in the creation of functions.
//TODO  Essentially rewrite every desired function. This willl be a big weekend project

const RinkebyEndpoint string = "https://rinkeby.infura.io/v3/cad14f2599ad47ec8d96c4413f80b5ed"
type AccountAPI struct{
	account *User.DIAccount

	address common.Address
	opts *bind.TransactOpts 
	client *ethclient.Client
}

func (AccountAPI) CreateAccount(endpointURL string, walletInfo Wallet, Name_ string, Bio_ string)(AccountAPI){
	auth, blockchain := CreateConnection(endpointURL, walletInfo)

	addr, _, account, err := User.DeployDIAccount(auth, blockchain, Name_)
	if err != nil {
		log.Fatalf("Failed to creating account: %v", err)
	}
	//TODO include a wait for the contract to deploy
	return AccountAPI{account, addr, auth, blockchain}
}

func (AccountAPI) LogIn(accountName_ string, r Registry)(AccountAPI){
	addr, err := r.GetAddressOfName(&bind.CallOpts{Pending: true}, accountName_)
	if err != nil {
		log.Fatalf("Failed to creating account: %v", err)
	}

	account, err := User.NewDIAccount(addr, r.client)
	if err != nil {
		log.Fatalf("Failed to creating logging in: %v", err)
	}
	//TODO include a wait for the contract to deploy
	return AccountAPI{account, addr, r.opts, r.client}
}

func (acc AccountAPI) SetBio(bio string){
	acc.account.SetBio(acc.opts, bio)
}

func (acc AccountAPI) CreatePost(PostID int64, PostURL_ string, Caption_ string){
	acc.account.CreatePost(acc.opts, PostURL_, Caption_)
	// Needs some interaction with ipfs
}

func (acc AccountAPI) GetPost(PostID int64)(*DisplayPost, error){
	post_, err := acc.account.Posts(&bind.CallOpts{Pending: true}, big.NewInt(PostID))
	if err != nil {
		return nil, err
	}
	
	comments_, err := acc.GetComments(PostID)
	if err != nil {
		return nil, err
	}
	
	return &(DisplayPost{post_, comments_}), nil
}

func (acc AccountAPI) GetComments(PostID int64)([]User.DIAccountComment, error){
	comments_, err := acc.account.GetComments(&bind.CallOpts{Pending: true}, big.NewInt(PostID))
	if err != nil {
		return nil, err
	}
	return comments_, nil
}

func (acc AccountAPI) PostVote(targ common.Address, P DisplayPost, vote int8) (error){
	if _, present := map[int8]bool{-1: true, 0: true, 1: true}[vote]; !present{
		return errors.New("Invalid vote, make sure the vote casted is either -1, 0, or 1")
	}

	acct, err := User.NewDIAccount(P.Post.Account, acc.client)
	if err != nil {
		return errors.New("Failed to find account: %v")
	}

	_, err = acct.Vote(acc.opts, targ, vote, P.Post.PostNumber)
	if err != nil {
		return err
	}
	return nil
}

func (acc AccountAPI) VoteComment(targ common.Address, P DisplayPost, CommentID int64, vote int8)(error){
	if _, present := map[int8]bool{-1: true, 0: true, 1: true}[vote]; !present{
		return errors.New("Invalid vote, make sure the vote casted is either -1, 0, or 1")
	}
	acct, err := User.NewDIAccount(P.Post.Account, acc.client)
	if err != nil {
		return errors.New("Failed to find account: %v")
	}

	_, err = acct.CommentVote(acc.opts, targ, vote, P.Post.PostNumber, big.NewInt(CommentID))
	if err != nil {
		return err
	}
	return nil
}

func (acc AccountAPI) DeletePost(PostID int64){
	acc.account.DeletePost(acc.opts, big.NewInt(PostID))
	// Needs some interaction with ipfs
}

func (acc AccountAPI) Follow(targ common.Address,){
	acc.account.Follow(acc.opts, targ)
}

func (acc AccountAPI) UnFollow(targ common.Address,){
	acc.account.UnFollow(acc.opts, targ)
}