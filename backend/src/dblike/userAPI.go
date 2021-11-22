package backendapi

import (
	"context"
	"log"
	"math/big"
	"strings"

	"github.com/kevinfjiang/DeInsta/backend/User"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

//TODO integrate these three packages to interact with one another, namely in the creation of functions.
//TODO  Essentially rewrite every desired function. This willl be a big weekend project

const RinkebyEndpoint string = "https://rinkeby.infura.io/v3/cad14f2599ad47ec8d96c4413f80b5ed"

type Wallet struct{
	Key string
	Password string
}

type AccountAPI struct{
	address common.Address
	account *User.DIAccount

	opts *bind.TransactOpts 
	client *ethclient.Client
}

func CreateConnection(endpointURL string, walletInfo Wallet)(*bind.TransactOpts, *ethclient.Client){
	blockchain, err := ethclient.Dial(endpointURL)

	if err != nil {
		log.Fatalf("Unable to connect to network:%v\n", err)
	}

	auth, err := bind.NewTransactor(strings.NewReader(walletInfo.Key), walletInfo.Password)
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}
	
	glimit, err := blockchain.SuggestGasPrice(context.Background())
	
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}
	auth.GasPrice=glimit

	return auth, blockchain
}

func CreateAccount(endpointURL string, walletInfo Wallet, Name_ string, Bio_ string)(AccountAPI){
	auth, blockchain := CreateConnection(endpointURL, walletInfo)

	addr, _, account, _ := User.DeployDIAccount(auth, blockchain, Name_)
	//TODO include a wait for the contract to deploy
	return AccountAPI{addr, account, auth, blockchain}
}

func (acc AccountAPI) SetBio(bio string){
	acc.account.SetBio(acc.opts, bio)
}

func (acc AccountAPI) DeletePost(PostID int64){
	acc.account.DeletePost(acc.opts, big.NewInt(PostID))
	// Needs some interaction with ipfs
}