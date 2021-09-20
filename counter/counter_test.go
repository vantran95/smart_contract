package counter

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
)

func TestCounterDeployContract(t *testing.T) {
	const address = "https://volta-rpc.energyweb.org"
	//const address = "http://127.0.0.1:8545"
	const key = `{"address":"c292eb3833a2506713d315e587d1521f04efbb16","crypto":{"cipher":"aes-128-ctr","ciphertext":"4be208667540e735aa0b4d6b42251db557520d207eb274b804df481745fb499b","cipherparams":{"iv":"4c5d2021efd58f68e55f663f6c62fa87"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"1e98822fed6cf3c9f80fca3e02bf587e2ac457381d49d55759d5737c05ddf013"},"mac":"44782c35cabcccb32ba4b4751372b5316decce96f7265cafd6fe95f5444dddb0"},"id":"ccaaaa04-9d99-409f-ac67-1ae17be009d8","version":3}`
	const pwd = "1234" //password from Geth
	client, err := ethclient.Dial(address)
	assert.NoError(t, err)
	assert.NoError(t, err)

	//estGas, err := client.SuggestGasPrice(context.Background())
	//assert.NoError(t, err)
	//fmt.Println("Suggest Gas: ", estGas)


	fromAddress := common.HexToAddress("0xc292eb3833a2506713d315e587d1521f04efbb16")
	estimatedGas, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
		From: fromAddress,
		Data: common.FromHex(CounterBin),
	})
	fmt.Println("Data: ", common.FromHex(CounterBin))

	chainID := big.NewInt(73799)

	auth, err := bind.NewTransactorWithChainID(strings.NewReader(key), pwd, chainID)
	assert.NoError(t, err)
	auth.NoSend = true
	rs, tnx, _, _ := DeployCounter(auth, client)
	strAddress := fmt.Sprintf("Contract pending deploy: 0x%x\n", rs)
	fmt.Println("Input: ", tnx.Data())
	fmt.Println(strAddress)
	assert.True(t, len(strAddress) > 0)
	gasLimit := tnx.Gas()
	fmt.Println("Gas Limit: ", gasLimit)
	gasPrice := tnx.GasPrice()
	fmt.Println("Gas Price: ", gasPrice)
	fmt.Println("Est Gas: ", estimatedGas)
	gasTotal := int64(gasLimit) * gasPrice.Int64()
	fmt.Printf("Gas Total: Gas price %d * Gas Limit %d = Total %d", gasPrice.Uint64(), gasLimit, gasTotal)
}

func TestIncrement(t *testing.T) {
	const address = "https://volta-rpc.energyweb.org"
	//const address = "http://127.0.0.1:8545"
	const key = `{"address":"c292eb3833a2506713d315e587d1521f04efbb16","crypto":{"cipher":"aes-128-ctr","ciphertext":"4be208667540e735aa0b4d6b42251db557520d207eb274b804df481745fb499b","cipherparams":{"iv":"4c5d2021efd58f68e55f663f6c62fa87"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"1e98822fed6cf3c9f80fca3e02bf587e2ac457381d49d55759d5737c05ddf013"},"mac":"44782c35cabcccb32ba4b4751372b5316decce96f7265cafd6fe95f5444dddb0"},"id":"ccaaaa04-9d99-409f-ac67-1ae17be009d8","version":3}`
	const pwd = "1234" //password from Geth
	conn, err := ethclient.Dial(address)
	assert.NoError(t, err)

	// this address has been stored to psql once deploy success
	strAddress := common.HexToAddress("0x9640B432044C04ac2656e96213594D8da8c33490")
	counter, err := NewCounter(strAddress, conn)
	assert.NoError(t, err)
	chainID := big.NewInt(73799)
	auth, err := bind.NewTransactorWithChainID(strings.NewReader(key), pwd, chainID)
	assert.NoError(t, err)

	inc, err := counter.Increment(auth)
	assert.NoError(t, err)
	count, err := counter.GetCount(&bind.CallOpts{})
	fmt.Println("counter: ", count)
	fmt.Println("rs.Weight: ", *inc.ChainId())
	fmt.Println("rs.Value: ", *inc.Value())
	fmt.Println("rs.GasPrice: ", *inc.GasPrice())
	fmt.Println("rs.Gas: ", inc.Gas())
}

func TestDecrement(t *testing.T) {

	const address = "https://volta-rpc.energyweb.org"
	//const address = "http://127.0.0.1:8545"
	const key = `{"address":"c292eb3833a2506713d315e587d1521f04efbb16","crypto":{"cipher":"aes-128-ctr","ciphertext":"4be208667540e735aa0b4d6b42251db557520d207eb274b804df481745fb499b","cipherparams":{"iv":"4c5d2021efd58f68e55f663f6c62fa87"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"1e98822fed6cf3c9f80fca3e02bf587e2ac457381d49d55759d5737c05ddf013"},"mac":"44782c35cabcccb32ba4b4751372b5316decce96f7265cafd6fe95f5444dddb0"},"id":"ccaaaa04-9d99-409f-ac67-1ae17be009d8","version":3}`
	const pwd = "1234" //password from Geth
	conn, err := ethclient.Dial(address)
	assert.NoError(t, err)

	// this address has been stored to psql once deploy success
	strAddress := common.HexToAddress("0x9640B432044C04ac2656e96213594D8da8c33490")
	counter, err := NewCounter(strAddress, conn)
	assert.NoError(t, err)
	chainID := big.NewInt(73799)
	auth, err := bind.NewTransactorWithChainID(strings.NewReader(key), pwd, chainID)
	assert.NoError(t, err)

	decrement, err := counter.Decrement(auth)
	assert.NoError(t, err)
	count, err := counter.GetCount(&bind.CallOpts{})
	fmt.Println("counter: ", count)
	fmt.Println("rs.Weight: ", *decrement.ChainId())
	fmt.Println("rs.Value: ", *decrement.Value())
	fmt.Println("rs.GasPrice: ", *decrement.GasPrice())
	fmt.Println("rs.Gas: ", decrement.Gas())
}