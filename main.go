package main


import (
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"fmt"
	"context"
)

var (
    ctx         = context.Background()
    url         = "https://mainnet.infura.io/v3/407f60619ec14e538991ba8f9e0f4237"
    client, err = ethclient.DialContext(ctx, url)
)

func main()  {
	// create eth client
	client, err := ethclient.DialContext(ctx, url)

	if err != nil {
		log.Fatalf("Error to create an ether clent")
	}
	defer client.Close()

	// Get block by number
	block, err := client.BlockByNumber(ctx, nil)
	if err != nil {
		log.Fatalf("Error to get block:%v", err)
	}
	fmt.Println(block.Number())

	
	
}