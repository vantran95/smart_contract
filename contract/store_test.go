package contract

import (
	"fmt"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
)

func TestDeployToRPC(t *testing.T) {
	//const address = "https://volta-rpc.energyweb.org"
	const address = "http://127.0.0.1:8545"
	const key = `{"address":"c292eb3833a2506713d315e587d1521f04efbb16","crypto":{"cipher":"aes-128-ctr","ciphertext":"4be208667540e735aa0b4d6b42251db557520d207eb274b804df481745fb499b","cipherparams":{"iv":"4c5d2021efd58f68e55f663f6c62fa87"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"1e98822fed6cf3c9f80fca3e02bf587e2ac457381d49d55759d5737c05ddf013"},"mac":"44782c35cabcccb32ba4b4751372b5316decce96f7265cafd6fe95f5444dddb0"},"id":"ccaaaa04-9d99-409f-ac67-1ae17be009d8","version":3}`
	const pwd = "1234" //password from Geth
	conn, err := ethclient.Dial(address)
	assert.NoError(t, err)

	auth, err := bind.NewTransactor(strings.NewReader(key), pwd)
	assert.NoError(t, err)
	rs, _, _, _ := DeployStore(auth, conn)
	strAddress := fmt.Sprintf("Contract pending deploy: 0x%x\n", rs)
	fmt.Println(strAddress)
	assert.True(t, len(strAddress) > 0)
}

func TestRetrieveFromRPC(t *testing.T) {
	//const address = "http://127.0.0.1:8545"
	//const key = `{"address":"f3b02b166a375a00930d5fb63fac2c41cff0134f","crypto":{"cipher":"aes-128-ctr","ciphertext":"05dbc4131d2cbc2c0ae1cd240262b89c424e9934087bdd071116749bc725bd73","cipherparams":{"iv":"09af21099717cde83fc740e493acc257"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"8d48cbd5c73a1b2a1d42aaed30d2038205476853527bfd52d2a939e88f4269eb"},"mac":"c8b5bd996690a896f1a2f8fed6903a24a9ca043e64aed506d65892bd36266fc8"},"id":"47faab3d-2247-4f1e-94e7-b042ca290850","version":3}`
	//const pwd = "van200295" //password from MetaMask
	const address = "http://127.0.0.1:8545"
	const key = `{"address":"c292eb3833a2506713d315e587d1521f04efbb16","crypto":{"cipher":"aes-128-ctr","ciphertext":"4be208667540e735aa0b4d6b42251db557520d207eb274b804df481745fb499b","cipherparams":{"iv":"4c5d2021efd58f68e55f663f6c62fa87"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"1e98822fed6cf3c9f80fca3e02bf587e2ac457381d49d55759d5737c05ddf013"},"mac":"44782c35cabcccb32ba4b4751372b5316decce96f7265cafd6fe95f5444dddb0"},"id":"ccaaaa04-9d99-409f-ac67-1ae17be009d8","version":3}`
	const pwd = "1234" //password from MetaMask
	conn, err := ethclient.Dial(address)
	assert.NoError(t, err)
	// this address has been stored to psql once deploy success
	strAddress := common.HexToAddress("0xc292eB3833a2506713D315E587d1521f04efBb16")
	store, err := NewStore(strAddress, conn)
	assert.NoError(t, err)
	auth, err := bind.NewTransactor(strings.NewReader(key), pwd)
	assert.NoError(t, err)
	rs, err := store.Retrieve(auth)
	assert.NoError(t, err)
	fmt.Println("rs.ChainId: ", *rs.ChainId())
	fmt.Println("rs.Value: ", *rs.Value())
	fmt.Println("rs.GasPrice: ", *rs.GasPrice())
	fmt.Println("rs.Gas: ", rs.Gas())
}