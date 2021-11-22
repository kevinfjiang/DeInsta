package backendapi

import(
	"context"
	"log"
	"strings"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	// "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Wallet struct{
	key string
	password string
}

func CreateConnection(endpointURL string, walletInfo Wallet)(*bind.TransactOpts, *ethclient.Client){
	blockchain, err := ethclient.Dial(endpointURL)

	if err != nil {
		log.Fatalf("Unable to connect to network:%v\n", err)
	}

	auth, err := bind.NewTransactor(strings.NewReader(walletInfo.key), walletInfo.password)
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