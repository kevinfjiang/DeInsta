package backendapi

import(
	"log"

	"github.com/kevinfjiang/DeInsta/dblike/User"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Registry struct {
	*User.DIregistry

	opts *bind.TransactOpts 
	client *ethclient.Client
}


//For admin
func CreateRegistry(endpointURL string, walletInfo Wallet, ServerName_ string)(common.Address, Registry){
	auth, blockchain := CreateConnection(endpointURL, walletInfo)

	addr, _, registry, err := User.DeployDIregistry(auth, blockchain, ServerName_)
	if err != nil {
		log.Fatalf("Failed to deploy Registry: %v", err)
	}

	return addr, Registry{registry, auth, blockchain}
}


// For user
func ConnectRegistry(endpointURL string, walletInfo Wallet, registryAddr common.Address)(Registry){
	auth, blockchain := CreateConnection(endpointURL, walletInfo)
	registry, err := User.NewDIregistry(registryAddr, blockchain)
	if err != nil {
		log.Fatalf("Failed to deploy connect to registry: %v", err)
	}

	return Registry{registry, auth, blockchain}
}

