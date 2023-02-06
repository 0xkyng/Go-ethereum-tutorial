package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/accounts/keystore"
)

// Generate a keystore wallet - Enccypted private key
// Kestore is an encrypted version of the ethereum private key using
// Symmetric encryption, this allows to add an other layer of security
// Even if another person gets access to the file, he cannot use it until
// He knows the password used to encrypt it

func main() {
	key := keystore.NewKeyStore("./wallet", keystore.StandardScryptN, keystore.StandardScryptP)
	password := "password"
	account, err := key.NewAccount(password)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(account.Address)
}
