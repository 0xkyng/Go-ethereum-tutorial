package main

import (
	"log"

	"github.com/ethereum/go-ethereum/accounts/keystore"
)

func main() {
	// Create an encrypted wallet
	// Using keystore
	ks := keystore.NewKeyStore("./wallet", keystore.StandardScryptN, keystore.LightScryptP)
	_, err := ks.NewAccount("secret")
	if err != nil {
		log.Fatal(err)
	}
	_, err = ks.NewAccount("secret")
	if err != nil {
		log.Fatal(err)
	}

	// "00397e3b5f78291709b945d792ad959f49667c0c"
	// "f3e80494e6e77f2d19e7f76abac8912abf153a5d"

}
