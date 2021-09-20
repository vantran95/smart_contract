package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/keystore"
)

func main() {
	ks := keystore.NewKeyStore("/Users/van.tran/go/src/contract/keystore", keystore.StandardScryptN, keystore.StandardScryptP)
	//am := accounts.NewManager(&accounts.Config{InsecureUnlockAllowed: false}, ks)

	// Create a new account with the specified encryption passphrase.
	newAcc, _ := ks.NewAccount("123")
	fmt.Println("New Acc: ", newAcc)


	// Create an account
	//key, err := crypto.GenerateKey()
	//if err != nil {
	//	fmt.Printf("Error")
	//}
	//
	//fmt.Println("Publickey values: ", key.PublicKey)
	//
	//// Get the address
	//address := crypto.PubkeyToAddress(key.PublicKey).Hex()
	//fmt.Println("New account address: ", address)
	//
	//// Get the private key
	//privateKey := hex.EncodeToString(key.D.Bytes())
	//fmt.Println("privateKey values: ", privateKey)
	//fmt.Printf("Hello")
}
