package emit_items

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
	"log"
	"math/big"
	"strings"
	"testing"
)

func TestEmitItemsDeployContract(t *testing.T) {
	//const address = "http://127.0.0.1:8545"
	//const key = `{"address":"c292eb3833a2506713d315e587d1521f04efbb16","crypto":{"cipher":"aes-128-ctr","ciphertext":"4be208667540e735aa0b4d6b42251db557520d207eb274b804df481745fb499b","cipherparams":{"iv":"4c5d2021efd58f68e55f663f6c62fa87"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"1e98822fed6cf3c9f80fca3e02bf587e2ac457381d49d55759d5737c05ddf013"},"mac":"44782c35cabcccb32ba4b4751372b5316decce96f7265cafd6fe95f5444dddb0"},"id":"ccaaaa04-9d99-409f-ac67-1ae17be009d8","version":3}`
	//const key = `{"address":"9f0b91204fb9daa8c4cf0b3289c267b3c33fae4d","crypto":{"cipher":"aes-128-ctr","ciphertext":"8cf7ac663144d1caa405bd70314d0e397ab7ba3b5f160dc7b716788e980bb9b7","cipherparams":{"iv":"61bbf8f0fb01cda61af708c9fa5157b3"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"391caaecfffed1895c5e05d58cec2e176f5464ee799e9e31dea067095efb30c9"},"mac":"8e0cb285ced87b376b34db5c112dbb8bf0700f3c3a4d030e0a192a4902fb8829"},"id":"9b1ed4d8-21a7-4584-95c1-0414f2297073","version":3}`


	const address = "https://volta-rpc.energyweb.org"
	const key = `{"address":"22810405bc0e12ae9b82f2ff2f03b56a2610a4b4","crypto":{"cipher":"aes-128-ctr","ciphertext":"73fccd82aaa2ea95b6373ea8e6393c2e33268170801248a4b1fdea31c1beb6cc","cipherparams":{"iv":"b38730c3d75e60d1d16dd0d19b14e03b"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"bdf9508510559ad262ef6c1b822184713487553a7a8a34c777ad8f2c54ac0ea8"},"mac":"004173990df959112aed1371a655d34bf1eab18512951ca52af452407f3427d6"},"id":"aa8db14e-8bb4-4550-9faa-411c41776d8e","version":3}`
	const pwd = "" //password from Geth
	conn, err := ethclient.Dial(address)
	assert.NoError(t, err)
	chainID := big.NewInt(73799)
	auth, err := bind.NewTransactorWithChainID(strings.NewReader(key), pwd, chainID)
	assert.NoError(t, err)

	rs, _, _, _ := DeployEmitItems(auth, conn)
	strAddress := fmt.Sprintf("Contract pending deploy: 0x%x\n", rs)
	fmt.Println(strAddress)
	assert.True(t, len(strAddress) > 0)
}


func TestSetItem(t *testing.T) {
	const address = "https://volta-rpc.energyweb.org"
	//const address = "http://127.0.0.1:8545"
	const key = `{"address":"c292eb3833a2506713d315e587d1521f04efbb16","crypto":{"cipher":"aes-128-ctr","ciphertext":"4be208667540e735aa0b4d6b42251db557520d207eb274b804df481745fb499b","cipherparams":{"iv":"4c5d2021efd58f68e55f663f6c62fa87"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"1e98822fed6cf3c9f80fca3e02bf587e2ac457381d49d55759d5737c05ddf013"},"mac":"44782c35cabcccb32ba4b4751372b5316decce96f7265cafd6fe95f5444dddb0"},"id":"ccaaaa04-9d99-409f-ac67-1ae17be009d8","version":3}`
	const pwd = "1234" //password from Geth
	conn, err := ethclient.Dial(address)
	assert.NoError(t, err)
	chainID := big.NewInt(73799)
	//contractAddress := common.HexToAddress("0x0d20570A52B45eb5Ccc1503C639eEFC2d4D3c159")


	strAddress := common.HexToAddress("0x0d20570A52B45eb5Ccc1503C639eEFC2d4D3c159")
	emit, err := NewEmitItems(strAddress, conn)
	assert.NoError(t, err)

	auth, err := bind.NewTransactorWithChainID(strings.NewReader(key), pwd, chainID)
	assert.NoError(t, err)
	//k := [32]byte{'J', 'O', 'H', 'N'}
	//value := [32]byte{'A', 'B', 'C', 'D'}
	k := [32]byte{}
	value := [32]byte{}
	b := []byte("foo")
	v := []byte("bar")
	copy(k[:], b)
	copy(value[:], v)

	rs, err := emit.SetItem(auth, k, value)
	assert.NoError(t, err)
	fmt.Println("tx sent: ", rs.Hash().Hex())

	getKey, err := emit.Items(nil, k)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("key: ", getKey)

	getValue, err := emit.Items(nil, value)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("value : ", string(getValue[:]))

	printCosts(rs)
}

func TestReadingLogItem(t *testing.T) {
	const address = "https://volta-rpc.energyweb.org"
	//const address = "http://127.0.0.1:8545"
	//const key = `{"address":"c292eb3833a2506713d315e587d1521f04efbb16","crypto":{"cipher":"aes-128-ctr","ciphertext":"4be208667540e735aa0b4d6b42251db557520d207eb274b804df481745fb499b","cipherparams":{"iv":"4c5d2021efd58f68e55f663f6c62fa87"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"1e98822fed6cf3c9f80fca3e02bf587e2ac457381d49d55759d5737c05ddf013"},"mac":"44782c35cabcccb32ba4b4751372b5316decce96f7265cafd6fe95f5444dddb0"},"id":"ccaaaa04-9d99-409f-ac67-1ae17be009d8","version":3}`
	//const pwd = "1234" //password from Geth
	conn, err := ethclient.Dial(address)
	assert.NoError(t, err)
	//chainID := big.NewInt(73799)
	//contractAddress := common.HexToAddress("0x0d20570A52B45eb5Ccc1503C639eEFC2d4D3c159")


	contractAddress := common.HexToAddress("0x0d20570A52B45eb5Ccc1503C639eEFC2d4D3c159")
	bytecode, err := conn.CodeAt(context.Background(), contractAddress, nil) // nil is latest block
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("block id: ", hex.EncodeToString(bytecode))
	//blockId := hex.EncodeToString(bytecode)

	// Get the latest block number
	header, err := conn.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(header.Number.String()) // 5671744
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(13438328),
		ToBlock:   big.NewInt(13438328),
		Addresses: []common.Address{
			contractAddress,
		},
	}

	logs, err := conn.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	contractAbi, err := abi.JSON(strings.NewReader(EmitItemsABI))
	if err != nil {
		log.Fatal(err)
	}

	for _, vLog := range logs {
		fmt.Println("log: ", vLog)
		fmt.Println(vLog.BlockHash.Hex()) // 0x3404b8c050aa0aacd0223e91b5c32fee6400f357764771d0684fa7b3f448f1a8
		fmt.Println(vLog.BlockNumber)     // 2394201
		fmt.Println(vLog.TxHash.Hex())    // 0x280201eda63c9ff6f305fcee51d5eb86167fab40ca3108ec784e8652a0e2b1a6
		fmt.Println("data: ",vLog.Data)

		event := struct {
			Key   [32]byte
			Value [32]byte
		}{}

		//_, err := contractAbi.Unpack("ItemSet", vLog.Data)
		// The parsed ABI interface to decode the raw log data.
		err := contractAbi.UnpackIntoInterface(&event, "ItemSet", vLog.Data)
		if err != nil {
			log.Fatal(err)
		}

		//fmt.Println(event.Key)   // foo
		//fmt.Println(event.Value) // bar

		fmt.Println(string(event.Key[:]))   // foo
		fmt.Println(string(event.Value[:])) // bar

		var topics [4]string
		for i := range vLog.Topics {
			topics[i] = vLog.Topics[i].Hex()
		}

		fmt.Println("topic 0: ", topics[0]) // 0xe79e73da417710ae99aa2088575580a60415d359acfad9cdd3382d59c80281d4
		fmt.Println("topic 1: ", topics[1]) // 0xe79e73da417710ae99aa2088575580a60415d359acfad9cdd3382d59c80281d4

		eventSignature := []byte("ItemSet(bytes32,bytes32)")
		hash := crypto.Keccak256Hash(eventSignature)
		fmt.Println(hash.Hex())
	}
}

func printCosts(tnx *types.Transaction) {
	tnx.Hash()
	wei := new(big.Int)
	wei.Set(CalcGasCost(tnx.Gas(), tnx.GasPrice()))
	cost := ToDecimal(wei, 18)
	fmt.Println(fmt.Sprintf("Gas Price: %d (in wei) X Gas Limit: %d (units) = Tx Cost: %s (in wei)", tnx.GasPrice(), tnx.Gas(), cost.String()))
}