package main


import (
	"github.com/ethereum/go-ethereum"
	"log"
)

var infuralUrl = "https://mainnet.infura.io/v3/407f60619ec14e538991ba8f9e0f4237"

func main()  {
	// create eth client
	client, err := ethclient.DialContext(context.Background(), infuralUrl)

	if err != nil {
		log.Fatalf("Error to create an ether clent")
	}
	defer client.Close()

	
	
}