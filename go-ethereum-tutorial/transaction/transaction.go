package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"

	// "github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/eth/gasprice"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	url = "https://goerli.infura.io/v3/f7cdc50f468747ec96cdd15ca75141ae"
)

func main() {
	// Create an encrypted wallet
	// Using keystore
	// ks := keystore.NewKeyStore("./wallet", keystore.StandardScryptN, keystore.LightScryptP)
	// _, err := ks.NewAccount("secret")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// _, err = ks.NewAccount("secret")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// "00397e3b5f78291709b945d792ad959f49667c0c"
	// "f3e80494e6e77f2d19e7f76abac8912abf153a5d"

	client, err := ethclient.Dial(url)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	account1 := common.HexToAddress("f3e80494e6e77f2d19e7f76abac8912abf153a5d")
	account2 := common.HexToAddress("00397e3b5f78291709b945d792ad959f49667c0c")

	account1Balance, err := client.BalanceAt(context.Background(), account1, nil)
	if err != nil {
		log.Fatal(err)
	}

	account2Balance, err := client.BalanceAt(context.Background(), account2, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("The balance for account1 is", account1Balance)
	fmt.Println("The balance for account2 is", account2Balance)

	// Make a transction
	nonce, err := client.PendingNonceAt(context.Background(), account1)
	if err != nil {
		log.Fatal(err)
	}
	amount := big.NewInt(1000000000000000)
	gasprice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	tx := types.NewTransaction(nonce, account2, amount, 2100, gasprice, nil)
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	b, err := ioutil.ReadFile()
	types.SignTx(tx, types.NewEIP155Signer(chainID))

}
