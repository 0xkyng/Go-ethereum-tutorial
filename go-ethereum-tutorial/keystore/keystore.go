package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

// Generate a keystore wallet - Enccypted private key
// Kestore is an encrypted version of the ethereum private key using
// Symmetric encryption, this allows to add an other layer of security
// Even if another person gets access to the file, he cannot use it until
// He knows the password used to encrypt it

func main() {
	// key := keystore.NewKeyStore("./wallet", keystore.StandardScryptN, keystore.StandardScryptP)
	password := "password"
	// account, err := key.NewAccount(password)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(account.Address)

	// Decrypt the file and get the private key

	file, err := ioutil.ReadFile("/wallet/UTC--2023-02-06T19-31-00.372104270Z--ad272b84ce085aba470b1b1f46c8f30e31ec80d5")
	if err != nil {
		log.Fatal(err)
	}

	key, err := keystore.DecryptKey(file, password)
	if err != nil {
		log.Fatal(err)
	}

	pData := crypto.FromECDSA(key.PrivateKey)
	hexutil.Encode(pData)

}
