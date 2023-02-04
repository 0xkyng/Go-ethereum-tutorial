package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	ctx         = context.Background()
	url         = "https://mainnet.infura.io/v3/407f60619ec14e538991ba8f9e0f4237"
	client, err = ethclient.DialContext(ctx, url)
)

func main() {
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
	fmt.Println("The block number:", block.Number())

	// Check the balance of an ether wallet
	addr := "0x2412784d09cd692Acc14142b2C8b0DEC022d52c2"
	address := common.HexToAddress(addr)

	balance, err := client.BalanceAt(ctx, address, nil)
	if err != nil {
		log.Fatalf("Error to get the balance:%v", err)
	}
	fmt.Println("The balance:", balance)
	// Covert the bigint to bigfloat
	// 1 ether = 10^18 wei
	fBalance := new(big.Float)
	fBalance.SetString(balance.String())
	fmt.Println(fBalance)

}
