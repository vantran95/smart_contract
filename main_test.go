package main

import (
	"fmt"
	"strings"
	"testing"

	"contract/contract"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
)

func TestDeployToRPC(t *testing.T) {
	const address = "https://volta-rpc.energyweb.org"
	const key = `{"address":"f3b02b166a375a00930d5fb63fac2c41cff0134f","crypto":{"cipher":"aes-128-ctr","ciphertext":"05dbc4131d2cbc2c0ae1cd240262b89c424e9934087bdd071116749bc725bd73","cipherparams":{"iv":"09af21099717cde83fc740e493acc257"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"8d48cbd5c73a1b2a1d42aaed30d2038205476853527bfd52d2a939e88f4269eb"},"mac":"c8b5bd996690a896f1a2f8fed6903a24a9ca043e64aed506d65892bd36266fc8"},"id":"47faab3d-2247-4f1e-94e7-b042ca290850","version":3}`
	const pwd = "van200295" //password from MetaMask
	conn, err := ethclient.Dial(address)
	assert.NoError(t, err)
	auth, err := bind.NewTransactor(strings.NewReader(key), pwd)
	assert.NoError(t, err)
	rs, _, _, _ := contract.DeployStore(auth, conn)
	strAddress := fmt.Sprintf("Contract pending deploy: 0x%x\n", rs)
	fmt.Println(strAddress)
	assert.True(t, len(strAddress) > 0)
}

func TestContractSuccess(t *testing.T) {
	const address = "https://volta-rpc.energyweb.org"
	const key = `{"address":"f3b02b166a375a00930d5fb63fac2c41cff0134f","crypto":{"cipher":"aes-128-ctr","ciphertext":"05dbc4131d2cbc2c0ae1cd240262b89c424e9934087bdd071116749bc725bd73","cipherparams":{"iv":"09af21099717cde83fc740e493acc257"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"8d48cbd5c73a1b2a1d42aaed30d2038205476853527bfd52d2a939e88f4269eb"},"mac":"c8b5bd996690a896f1a2f8fed6903a24a9ca043e64aed506d65892bd36266fc8"},"id":"47faab3d-2247-4f1e-94e7-b042ca290850","version":3}`
	const pwd = "van200295" //password from MetaMask
	conn, err := ethclient.Dial(address)
	assert.NoError(t, err)
	auth, err := bind.ContractCaller()
}