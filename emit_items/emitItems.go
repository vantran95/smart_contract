// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package emit_items

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// EmitItemsABI is the input ABI used to generate the binding from.
const EmitItemsABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"value\",\"type\":\"bytes32\"}],\"name\":\"ItemSet\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"items\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_version\",\"type\":\"string\"}],\"name\":\"set\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"value\",\"type\":\"bytes32\"}],\"name\":\"setItem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// EmitItemsBin is the compiled bytecode used for deploying new contracts.
var EmitItemsBin = "0x608060405234801561001057600080fd5b50610669806100206000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c806348f343f3146100515780634ed3885e1461008157806354fd4d501461009d578063f56256c7146100bb575b600080fd5b61006b60048036038101906100669190610314565b6100d7565b6040516100789190610412565b60405180910390f35b61009b60048036038101906100969190610381565b6100ef565b005b6100a5610109565b6040516100b29190610456565b60405180910390f35b6100d560048036038101906100d09190610341565b610197565b005b60016020528060005260406000206000915090505481565b80600090805190602001906101059291906101ec565b5050565b6000805461011690610536565b80601f016020809104026020016040519081016040528092919081815260200182805461014290610536565b801561018f5780601f106101645761010080835404028352916020019161018f565b820191906000526020600020905b81548152906001019060200180831161017257829003601f168201915b505050505081565b8060016000848152602001908152602001600020819055507fe79e73da417710ae99aa2088575580a60415d359acfad9cdd3382d59c80281d482826040516101e092919061042d565b60405180910390a15050565b8280546101f890610536565b90600052602060002090601f01602090048101928261021a5760008555610261565b82601f1061023357805160ff1916838001178555610261565b82800160010185558215610261579182015b82811115610260578251825591602001919060010190610245565b5b50905061026e9190610272565b5090565b5b8082111561028b576000816000905550600101610273565b5090565b60006102a261029d8461049d565b610478565b9050828152602081018484840111156102be576102bd6105fc565b5b6102c98482856104f4565b509392505050565b6000813590506102e08161061c565b92915050565b600082601f8301126102fb576102fa6105f7565b5b813561030b84826020860161028f565b91505092915050565b60006020828403121561032a57610329610606565b5b6000610338848285016102d1565b91505092915050565b6000806040838503121561035857610357610606565b5b6000610366858286016102d1565b9250506020610377858286016102d1565b9150509250929050565b60006020828403121561039757610396610606565b5b600082013567ffffffffffffffff8111156103b5576103b4610601565b5b6103c1848285016102e6565b91505092915050565b6103d3816104ea565b82525050565b60006103e4826104ce565b6103ee81856104d9565b93506103fe818560208601610503565b6104078161060b565b840191505092915050565b600060208201905061042760008301846103ca565b92915050565b600060408201905061044260008301856103ca565b61044f60208301846103ca565b9392505050565b6000602082019050818103600083015261047081846103d9565b905092915050565b6000610482610493565b905061048e8282610568565b919050565b6000604051905090565b600067ffffffffffffffff8211156104b8576104b76105c8565b5b6104c18261060b565b9050602081019050919050565b600081519050919050565b600082825260208201905092915050565b6000819050919050565b82818337600083830152505050565b60005b83811015610521578082015181840152602081019050610506565b83811115610530576000848401525b50505050565b6000600282049050600182168061054e57607f821691505b6020821081141561056257610561610599565b5b50919050565b6105718261060b565b810181811067ffffffffffffffff821117156105905761058f6105c8565b5b80604052505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b610625816104ea565b811461063057600080fd5b5056fea2646970667358221220366b0e1a578f836dd84d79f98bed793dca6c245966111a505f16385ea783941264736f6c63430008060033"

// DeployEmitItems deploys a new Ethereum contract, binding an instance of EmitItems to it.
func DeployEmitItems(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *EmitItems, error) {
	parsed, err := abi.JSON(strings.NewReader(EmitItemsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(EmitItemsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EmitItems{EmitItemsCaller: EmitItemsCaller{contract: contract}, EmitItemsTransactor: EmitItemsTransactor{contract: contract}, EmitItemsFilterer: EmitItemsFilterer{contract: contract}}, nil
}

// EmitItems is an auto generated Go binding around an Ethereum contract.
type EmitItems struct {
	EmitItemsCaller     // Read-only binding to the contract
	EmitItemsTransactor // Write-only binding to the contract
	EmitItemsFilterer   // Log filterer for contract events
}

// EmitItemsCaller is an auto generated read-only Go binding around an Ethereum contract.
type EmitItemsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EmitItemsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EmitItemsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EmitItemsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EmitItemsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EmitItemsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EmitItemsSession struct {
	Contract     *EmitItems        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EmitItemsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EmitItemsCallerSession struct {
	Contract *EmitItemsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// EmitItemsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EmitItemsTransactorSession struct {
	Contract     *EmitItemsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// EmitItemsRaw is an auto generated low-level Go binding around an Ethereum contract.
type EmitItemsRaw struct {
	Contract *EmitItems // Generic contract binding to access the raw methods on
}

// EmitItemsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EmitItemsCallerRaw struct {
	Contract *EmitItemsCaller // Generic read-only contract binding to access the raw methods on
}

// EmitItemsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EmitItemsTransactorRaw struct {
	Contract *EmitItemsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEmitItems creates a new instance of EmitItems, bound to a specific deployed contract.
func NewEmitItems(address common.Address, backend bind.ContractBackend) (*EmitItems, error) {
	contract, err := bindEmitItems(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EmitItems{EmitItemsCaller: EmitItemsCaller{contract: contract}, EmitItemsTransactor: EmitItemsTransactor{contract: contract}, EmitItemsFilterer: EmitItemsFilterer{contract: contract}}, nil
}

// NewEmitItemsCaller creates a new read-only instance of EmitItems, bound to a specific deployed contract.
func NewEmitItemsCaller(address common.Address, caller bind.ContractCaller) (*EmitItemsCaller, error) {
	contract, err := bindEmitItems(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EmitItemsCaller{contract: contract}, nil
}

// NewEmitItemsTransactor creates a new write-only instance of EmitItems, bound to a specific deployed contract.
func NewEmitItemsTransactor(address common.Address, transactor bind.ContractTransactor) (*EmitItemsTransactor, error) {
	contract, err := bindEmitItems(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EmitItemsTransactor{contract: contract}, nil
}

// NewEmitItemsFilterer creates a new log filterer instance of EmitItems, bound to a specific deployed contract.
func NewEmitItemsFilterer(address common.Address, filterer bind.ContractFilterer) (*EmitItemsFilterer, error) {
	contract, err := bindEmitItems(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EmitItemsFilterer{contract: contract}, nil
}

// bindEmitItems binds a generic wrapper to an already deployed contract.
func bindEmitItems(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EmitItemsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EmitItems *EmitItemsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EmitItems.Contract.EmitItemsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EmitItems *EmitItemsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EmitItems.Contract.EmitItemsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EmitItems *EmitItemsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EmitItems.Contract.EmitItemsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EmitItems *EmitItemsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EmitItems.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EmitItems *EmitItemsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EmitItems.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EmitItems *EmitItemsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EmitItems.Contract.contract.Transact(opts, method, params...)
}

// Items is a free data retrieval call binding the contract method 0x48f343f3.
//
// Solidity: function items(bytes32 ) view returns(bytes32)
func (_EmitItems *EmitItemsCaller) Items(opts *bind.CallOpts, arg0 [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _EmitItems.contract.Call(opts, &out, "items", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Items is a free data retrieval call binding the contract method 0x48f343f3.
//
// Solidity: function items(bytes32 ) view returns(bytes32)
func (_EmitItems *EmitItemsSession) Items(arg0 [32]byte) ([32]byte, error) {
	return _EmitItems.Contract.Items(&_EmitItems.CallOpts, arg0)
}

// Items is a free data retrieval call binding the contract method 0x48f343f3.
//
// Solidity: function items(bytes32 ) view returns(bytes32)
func (_EmitItems *EmitItemsCallerSession) Items(arg0 [32]byte) ([32]byte, error) {
	return _EmitItems.Contract.Items(&_EmitItems.CallOpts, arg0)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_EmitItems *EmitItemsCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _EmitItems.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_EmitItems *EmitItemsSession) Version() (string, error) {
	return _EmitItems.Contract.Version(&_EmitItems.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_EmitItems *EmitItemsCallerSession) Version() (string, error) {
	return _EmitItems.Contract.Version(&_EmitItems.CallOpts)
}

// Set is a paid mutator transaction binding the contract method 0x4ed3885e.
//
// Solidity: function set(string _version) returns()
func (_EmitItems *EmitItemsTransactor) Set(opts *bind.TransactOpts, _version string) (*types.Transaction, error) {
	return _EmitItems.contract.Transact(opts, "set", _version)
}

// Set is a paid mutator transaction binding the contract method 0x4ed3885e.
//
// Solidity: function set(string _version) returns()
func (_EmitItems *EmitItemsSession) Set(_version string) (*types.Transaction, error) {
	return _EmitItems.Contract.Set(&_EmitItems.TransactOpts, _version)
}

// Set is a paid mutator transaction binding the contract method 0x4ed3885e.
//
// Solidity: function set(string _version) returns()
func (_EmitItems *EmitItemsTransactorSession) Set(_version string) (*types.Transaction, error) {
	return _EmitItems.Contract.Set(&_EmitItems.TransactOpts, _version)
}

// SetItem is a paid mutator transaction binding the contract method 0xf56256c7.
//
// Solidity: function setItem(bytes32 key, bytes32 value) returns()
func (_EmitItems *EmitItemsTransactor) SetItem(opts *bind.TransactOpts, key [32]byte, value [32]byte) (*types.Transaction, error) {
	return _EmitItems.contract.Transact(opts, "setItem", key, value)
}

// SetItem is a paid mutator transaction binding the contract method 0xf56256c7.
//
// Solidity: function setItem(bytes32 key, bytes32 value) returns()
func (_EmitItems *EmitItemsSession) SetItem(key [32]byte, value [32]byte) (*types.Transaction, error) {
	return _EmitItems.Contract.SetItem(&_EmitItems.TransactOpts, key, value)
}

// SetItem is a paid mutator transaction binding the contract method 0xf56256c7.
//
// Solidity: function setItem(bytes32 key, bytes32 value) returns()
func (_EmitItems *EmitItemsTransactorSession) SetItem(key [32]byte, value [32]byte) (*types.Transaction, error) {
	return _EmitItems.Contract.SetItem(&_EmitItems.TransactOpts, key, value)
}

// EmitItemsItemSetIterator is returned from FilterItemSet and is used to iterate over the raw logs and unpacked data for ItemSet events raised by the EmitItems contract.
type EmitItemsItemSetIterator struct {
	Event *EmitItemsItemSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *EmitItemsItemSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EmitItemsItemSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(EmitItemsItemSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *EmitItemsItemSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EmitItemsItemSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EmitItemsItemSet represents a ItemSet event raised by the EmitItems contract.
type EmitItemsItemSet struct {
	Key   [32]byte
	Value [32]byte
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterItemSet is a free log retrieval operation binding the contract event 0xe79e73da417710ae99aa2088575580a60415d359acfad9cdd3382d59c80281d4.
//
// Solidity: event ItemSet(bytes32 key, bytes32 value)
func (_EmitItems *EmitItemsFilterer) FilterItemSet(opts *bind.FilterOpts) (*EmitItemsItemSetIterator, error) {

	logs, sub, err := _EmitItems.contract.FilterLogs(opts, "ItemSet")
	if err != nil {
		return nil, err
	}
	return &EmitItemsItemSetIterator{contract: _EmitItems.contract, event: "ItemSet", logs: logs, sub: sub}, nil
}

// WatchItemSet is a free log subscription operation binding the contract event 0xe79e73da417710ae99aa2088575580a60415d359acfad9cdd3382d59c80281d4.
//
// Solidity: event ItemSet(bytes32 key, bytes32 value)
func (_EmitItems *EmitItemsFilterer) WatchItemSet(opts *bind.WatchOpts, sink chan<- *EmitItemsItemSet) (event.Subscription, error) {

	logs, sub, err := _EmitItems.contract.WatchLogs(opts, "ItemSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EmitItemsItemSet)
				if err := _EmitItems.contract.UnpackLog(event, "ItemSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseItemSet is a log parse operation binding the contract event 0xe79e73da417710ae99aa2088575580a60415d359acfad9cdd3382d59c80281d4.
//
// Solidity: event ItemSet(bytes32 key, bytes32 value)
func (_EmitItems *EmitItemsFilterer) ParseItemSet(log types.Log) (*EmitItemsItemSet, error) {
	event := new(EmitItemsItemSet)
	if err := _EmitItems.contract.UnpackLog(event, "ItemSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
