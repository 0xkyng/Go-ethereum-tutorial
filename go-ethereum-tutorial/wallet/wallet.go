package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	// Generate wallet address

	// Generate a private key
	pvKey, err := crypto.GenerateKey()

	if err != nil {
		log.Fatal(err)
	}

	// Covert the private key to a string
	pvKeyData := crypto.FromECDSA(pvKey)
	// Encode the privateKeyData
	fmt.Println(hexutil.Encode(pvKeyData))

	// Create a public key from the private key
	pubKeyData := crypto.FromECDSAPub(&pvKey.PublicKey)
	fmt.Println(hexutil.Encode(pubKeyData))

	// Create public address from the public key
	fmt.Println(crypto.PubkeyToAddress(pvKey.PublicKey).Hex())

}
