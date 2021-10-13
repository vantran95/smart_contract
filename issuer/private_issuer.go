// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package issuer

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

// PrivateIssuerPrivateTransferRequest is an auto generated low-level Go binding around an user-defined struct.
type PrivateIssuerPrivateTransferRequest struct {
	Owner common.Address
	Hash  [32]byte
}

// PrivateIssuerProof is an auto generated low-level Go binding around an user-defined struct.
type PrivateIssuerProof struct {
	Left bool
	Hash [32]byte
}

// PrivateIssuerRequestStateChange is an auto generated low-level Go binding around an user-defined struct.
type PrivateIssuerRequestStateChange struct {
	Owner         common.Address
	CertificateId *big.Int
	Hash          [32]byte
	Approved      bool
}

// PrivateIssuerABI is the input ABI used to generate the binding from.
const PrivateIssuerABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_certificateId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"CertificateMigratedToPublic\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_commitment\",\"type\":\"bytes32\"}],\"name\":\"CommitmentUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"MigrateToPublicRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_certificateId\",\"type\":\"uint256\"}],\"name\":\"PrivateCertificationRequestApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_certificateId\",\"type\":\"uint256\"}],\"name\":\"PrivateTransferRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"issuer\",\"outputs\":[{\"internalType\":\"contractIssuer\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"registry\",\"outputs\":[{\"internalType\":\"contractRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\",\"payable\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_issuer\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"certificateId\",\"type\":\"uint256\"}],\"name\":\"getCertificateCommitment\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_commitment\",\"type\":\"bytes32\"}],\"name\":\"approveCertificationRequestPrivate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_commitment\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"issuePrivate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_certificateId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_ownerAddressLeafHash\",\"type\":\"bytes32\"}],\"name\":\"requestPrivateTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_certificateId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"left\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"internalType\":\"structPrivateIssuer.Proof[]\",\"name\":\"_proof\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes32\",\"name\":\"_previousCommitment\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_commitment\",\"type\":\"bytes32\"}],\"name\":\"approvePrivateTransfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_certificateId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_ownerAddressLeafHash\",\"type\":\"bytes32\"}],\"name\":\"requestMigrateToPublic\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_migrationRequestId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_certificateId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"_ownerAddressLeafHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"_forAddress\",\"type\":\"address\"}],\"name\":\"requestMigrateToPublicFor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_migrationRequestId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_certificateId\",\"type\":\"uint256\"}],\"name\":\"getPrivateTransferRequest\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"internalType\":\"structPrivateIssuer.PrivateTransferRequest\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"}],\"name\":\"getMigrationRequest\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"certificateId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"internalType\":\"structPrivateIssuer.RequestStateChange\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_certificateId\",\"type\":\"uint256\"}],\"name\":\"getMigrationRequestId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_migrationRequestId\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_volume\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_salt\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"left\",\"type\":\"bool\"},{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"internalType\":\"structPrivateIssuer.Proof[]\",\"name\":\"_proof\",\"type\":\"tuple[]\"}],\"name\":\"migrateToPublic\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getIssuerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\",\"constant\":true}]"

// PrivateIssuerBin is the compiled bytecode used for deploying new contracts.
var PrivateIssuerBin = "0x608060405234801561001057600080fd5b506123fd806100206000396000f3fe608060405260043610610147576000357c0100000000000000000000000000000000000000000000000000000000900480637b103999116100c8578063abba27b51161008c578063abba27b5146103ec578063b747c9111461040c578063c13cbb631461042c578063c45029221461044c578063c4d66de81461046c578063f2fde38b1461048c57600080fd5b80637b1039991461031a5780637dbca0061461033a5780638da5cb5b1461035a5780639317058614610378578063a8551fd3146103a857600080fd5b80633659cfe61161010f5780633659cfe6146102575780634f1ef2861461027957806354fd4d501461028c578063715018a6146102d8578063750e247b146102ed57600080fd5b80631927b6a51461014c5780631d143848146101b357806326eb2fd1146101eb5780633163e28a1461021957806332e71ead14610239575b600080fd5b34801561015857600080fd5b5061016c610167366004611f36565b6104ac565b6040516101aa91908151600160a060020a03168152602080830151908201526040808301519082015260609182015115159181019190915260800190565b60405180910390f35b3480156101bf57600080fd5b5060c9546101d390600160a060020a031681565b604051600160a060020a0390911681526020016101aa565b3480156101f757600080fd5b5061020b610206366004611f36565b61055c565b6040519081526020016101aa565b34801561022557600080fd5b5061020b610234366004611d55565b610625565b34801561024557600080fd5b5060c954600160a060020a03166101d3565b34801561026357600080fd5b50610277610272366004611d1d565b610763565b005b610277610287366004611dac565b61078a565b34801561029857600080fd5b50604080518082018252600481527f76302e3100000000000000000000000000000000000000000000000000000000602082015290516101aa91906121b7565b3480156102e457600080fd5b506102776107a3565b3480156102f957600080fd5b5061020b610308366004611f36565b600090815260ce602052604090205490565b34801561032657600080fd5b5060ca546101d390600160a060020a031681565b34801561034657600080fd5b5061020b610355366004611fe3565b6107dc565b34801561036657600080fd5b50603354600160a060020a03166101d3565b34801561038457600080fd5b50610398610393366004611f66565b61081f565b60405190151581526020016101aa565b3480156103b457600080fd5b506103c86103c3366004611f36565b6109c4565b604080518251600160a060020a0316815260209283015192810192909252016101aa565b3480156103f857600080fd5b5061020b610407366004611fc2565b610a38565b34801561041857600080fd5b50610277610427366004611fc2565b610a45565b34801561043857600080fd5b5061027761044736600461201b565b610b3d565b34801561045857600080fd5b5061020b610467366004611fc2565b610edb565b34801561047857600080fd5b50610277610487366004611d1d565b611096565b34801561049857600080fd5b506102776104a7366004611d1d565b611233565b604080516080810182526000808252602082018190529181018290526060810191909152603354600160a060020a031633146105065760405160e560020a62461bcd0281526004016104fd90612227565b60405180910390fd5b50600081815260cb602090815260409182902082516080810184528154600160a060020a0316815260018201549281019290925260028101549282019290925260039091015460ff16151560608201525b919050565b603354600090600160a060020a0316331461058c5760405160e560020a62461bcd0281526004016104fd90612227565b600060015b60cf5481116105ce57600081815260cb60205260409020600101548414156105bc5791506105579050565b806105c68161230e565b915050610591565b508061061f5760405160e560020a62461bcd02815260206004820181905260248201527f756e61626c6520746f2066696e64206d6967726174696f6e207265717565737460448201526064016104fd565b50919050565b603354600090600160a060020a031633146106555760405160e560020a62461bcd0281526004016104fd90612227565b600160a060020a0384166106ae5760405160e560020a62461bcd02815260206004820152601660248201527f43616e6e6f74207573652061646472657373203078300000000000000000000060448201526064016104fd565b60c9546040517f20f2637d000000000000000000000000000000000000000000000000000000008152600091600160a060020a0316906320f2637d906106fa908690899060040161218c565b602060405180830381600087803b15801561071457600080fd5b505af1158015610728573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061074c9190611f4e565b90506107588185610edb565b9150505b9392505050565b61076c816112e8565b61078781604051806020016040528060008152506000611315565b50565b610793826112e8565b61079f82826001611315565b5050565b603354600160a060020a031633146107d05760405160e560020a62461bcd0281526004016104fd90612227565b6107da60006114dc565b565b603354600090600160a060020a0316331461080c5760405160e560020a62461bcd0281526004016104fd90612227565b61081784848461152e565b949350505050565b603354600090600160a060020a0316331461084f5760405160e560020a62461bcd0281526004016104fd90612227565b600086815260cc602052604090208054600160a060020a03166108b75760405160e560020a62461bcd02815260206004820152601d60248201527f4e6f6e2d6578697374696e672070726976617465207472616e7366657200000060448201526064016104fd565b6109188160010154848888808060200260200160405190810160405280939291908181526020016000905b8282101561090e576108ff60408302860136819003810190611edf565b815260200190600101906108e2565b505050505061164b565b6109675760405160e560020a62461bcd02815260206004820152601360248201527f496e76616c6964206d65726b6c6520747265650000000000000000000000000060448201526064016104fd565b604080518082018252600080825260208083018281528b835260cc90915292902090518154600160a060020a031916600160a060020a0390911617815590516001909101556109b7878585611730565b5060019695505050505050565b6040805180820190915260008082526020820152603354600160a060020a03163314610a055760405160e560020a62461bcd0281526004016104fd90612227565b50600090815260cc602090815260409182902082518084019093528054600160a060020a03168352600101549082015290565b600061075c83833361152e565b600082815260cc602052604090208054600160a060020a031615610aae5760405160e560020a62461bcd02815260206004820152601e60248201527f4f6e652070726976617465207472616e7366657220617420612074696d65000060448201526064016104fd565b6040518060400160405280610ac03390565b600160a060020a0390811682526020918201859052600086815260cc83526040902083518154600160a060020a03191692169190911781559101516001909101558233600160a060020a03167fde442453a4fabb4e6fe9cd90c427b54511e2e7712b8dfe60307a8de6209db0f160405160405180910390a3505050565b603354600160a060020a03163314610b6a5760405160e560020a62461bcd0281526004016104fd90612227565b600086815260cb60205260409020600381015460ff1615610bd05760405160e560020a62461bcd02815260206004820152601860248201527f5265717565737420616c726561647920617070726f766564000000000000000060448201526064016104fd565b6001810154600090815260cd602052604090205460ff1615610c375760405160e560020a62461bcd02815260206004820152601c60248201527f436572746966696361746520616c7265616479206d696772617465640000000060448201526064016104fd565b8054604051610c5891600160a060020a0316908890889088906020016120f2565b60405160208183030381529060405280519060200120816002015414610cc35760405160e560020a62461bcd02815260206004820152601d60248201527f526571756573746564206861736820646f6573206e6f74206d6174636800000060448201526064016104fd565b8054604080516020601f8801819004810282018101909252868152610d6c92600160a060020a031691899190899089908190840183828082843760009201829052506001890154815260ce602090815260408083205481518d8402810184019092528c8252955093508b92508a91829185015b82821015610d6257610d5360408302860136819003810190611edf565b81526020019060010190610d36565b50505050506117ea565b610dbb5760405160e560020a62461bcd02815260206004820152600d60248201527f496e76616c69642070726f6f660000000000000000000000000000000000000060448201526064016104fd565b60038101805460ff1916600190811790915560ca549082015482546040517f836a10400000000000000000000000000000000000000000000000000000000081526004810192909252600160a060020a039081166024830152604482018990529091169063836a104090606401600060405180830381600087803b158015610e4257600080fd5b505af1158015610e56573d6000803e3d6000fd5b5050505060018181018054600090815260cd60209081526040808320805460ff1916909517909455915480825260ce90925291822054610e9592611730565b805460018201546040518892600160a060020a031691907f5d348ef3cec3f8493cd139d1ee84e4c84f5b987f49646d7dfa226dbe98b770c990600090a450505050505050565b603354600090600160a060020a03163314610f0b5760405160e560020a62461bcd0281526004016104fd90612227565b60c9546040517fe3e47cb60000000000000000000000000000000000000000000000000000000081526004810185905260006024820181905291600160a060020a03169063e3e47cb690604401602060405180830381600087803b158015610f7257600080fd5b505af1158015610f86573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610faa9190611f4e565b9050610fb881600085611730565b60c9546040517f3db5717000000000000000000000000000000000000000000000000000000000815260048101869052600091600160a060020a031690633db571709060240160006040518083038186803b15801561101657600080fd5b505afa15801561102a573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526110529190810190611dfa565b805160405191925083918791600160a060020a0316907f382c89c3e15b040e673bdec044eba907f4e9af06293e0058163ca396cecb39ff90600090a4509392505050565b600054610100900460ff16806110af575060005460ff16155b6110ce5760405160e560020a62461bcd0281526004016104fd906121ca565b600054610100900460ff161580156110f0576000805461ffff19166101011790555b600160a060020a0382166111495760405160e560020a62461bcd02815260206004820152601660248201527f43616e6e6f74207573652061646472657373203078300000000000000000000060448201526064016104fd565b60c98054600160a060020a031916600160a060020a038416908117909155604080517ff21de1e8000000000000000000000000000000000000000000000000000000008152905163f21de1e891600480820192602092909190829003018186803b1580156111b657600080fd5b505afa1580156111ca573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111ee9190611d39565b60ca8054600160a060020a031916600160a060020a0392909216919091179055611216611830565b61121e6118ae565b801561079f576000805461ff00191690555050565b603354600160a060020a031633146112605760405160e560020a62461bcd0281526004016104fd90612227565b600160a060020a0381166112df5760405160e560020a62461bcd02815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f646472657373000000000000000000000000000000000000000000000000000060648201526084016104fd565b610787816114dc565b603354600160a060020a031633146107875760405160e560020a62461bcd0281526004016104fd90612227565b60006113487f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc54600160a060020a031690565b905061135384611918565b6000835111806113605750815b156113715761136f84846119d0565b505b7f4910fdfa16fed3260ed0e7147f7cc6da11a60208b5b9406d12a635614ffd9143805460ff166114d557805460ff19166001178155604051600160a060020a038316602482015261141e90869060440160408051601f198184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f3659cfe6000000000000000000000000000000000000000000000000000000001790526119d0565b50805460ff191681557f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc54600160a060020a038381169116146114cc5760405160e560020a62461bcd02815260206004820152602f60248201527f45524331393637557067726164653a207570677261646520627265616b73206660448201527f757274686572207570677261646573000000000000000000000000000000000060648201526084016104fd565b6114d585611ad5565b5050505050565b60338054600160a060020a03838116600160a060020a0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b60008061153a85611b15565b9050801561158d5760405160e560020a62461bcd02815260206004820181905260248201527f4d6967726174696f6e207265717565737420616c72656164792065786973747360448201526064016104fd565b600060cf6000815461159e9061230e565b918290555060408051608081018252600160a060020a0387811680835260208084018c81528486018c815260006060870181815289825260cb90945287812096518754600160a060020a031916961695909517865590516001860155516002850155516003909301805460ff1916931515939093179092559151929350839290917fd18a53554a32a70376b6c223bfb0d84fc24caa3c704ef04eaf1b868e74ee8f6a91a395945050505050565b600083815b8351811015611725576000848281518110611694577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001015190508060000151156116e3576020808201516040516116c692869101918252602082015260400190565b604051602081830303815290604052805190602001209250611712565b602080820151604080519283018690528201526060016040516020818303038152906040528051906020012092505b508061171d8161230e565b915050611650565b509092149392505050565b600083815260ce602052604090205482146117905760405160e560020a62461bcd02815260206004820152601b60248201527f50726576696f757320636f6d6d69746d656e7420696e76616c6964000000000060448201526064016104fd565b600083815260ce602052604090208190558233600160a060020a03167f20c9608d29cbc79a883dcda48c0f0a9f1babeccea0680f65ef20ce17b06dad15836040516117dd91815260200190565b60405180910390a3505050565b6000808686866040516020016118029392919061212b565b60405160208183030381529060405280519060200120905061182581858561164b565b979650505050505050565b600054610100900460ff1680611849575060005460ff16155b6118685760405160e560020a62461bcd0281526004016104fd906121ca565b600054610100900460ff1615801561188a576000805461ffff19166101011790555b611892611b60565b61189a611bcd565b8015610787576000805461ff001916905550565b600054610100900460ff16806118c7575060005460ff16155b6118e65760405160e560020a62461bcd0281526004016104fd906121ca565b600054610100900460ff16158015611908576000805461ffff19166101011790555b611910611b60565b61189a611b60565b803b61198f5760405160e560020a62461bcd02815260206004820152602d60248201527f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60448201527f6f74206120636f6e74726163740000000000000000000000000000000000000060648201526084016104fd565b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc8054600160a060020a031916600160a060020a0392909216919091179055565b6060823b611a495760405160e560020a62461bcd02815260206004820152602660248201527f416464726573733a2064656c65676174652063616c6c20746f206e6f6e2d636f60448201527f6e7472616374000000000000000000000000000000000000000000000000000060648201526084016104fd565b60008084600160a060020a031684604051611a649190612170565b600060405180830381855af49150503d8060008114611a9f576040519150601f19603f3d011682016040523d82523d6000602084013e611aa4565b606091505b5091509150611acc82826040518060600160405280602781526020016123a160279139611c30565b95945050505050565b611ade81611918565b604051600160a060020a038216907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a250565b60008060015b60cf548111611b5957600081815260cb6020526040902060010154841415611b47575060019392505050565b80611b518161230e565b915050611b1b565b5092915050565b600054610100900460ff1680611b79575060005460ff16155b611b985760405160e560020a62461bcd0281526004016104fd906121ca565b600054610100900460ff1615801561189a576000805461ffff19166101011790558015610787576000805461ff001916905550565b600054610100900460ff1680611be6575060005460ff16155b611c055760405160e560020a62461bcd0281526004016104fd906121ca565b600054610100900460ff16158015611c27576000805461ffff19166101011790555b61189a336114dc565b60608315611c3f57508161075c565b825115611c4f5782518084602001fd5b8160405160e560020a62461bcd0281526004016104fd91906121b7565b80516105578161237d565b60008083601f840112611c88578182fd5b50813567ffffffffffffffff811115611c9f578182fd5b602083019150836020604083028501011115611cba57600080fd5b9250929050565b805161055781612392565b600082601f830112611cdc578081fd5b8135611cef611cea826122b6565b612285565b818152846020838601011115611d03578283fd5b816020850160208301379081016020019190915292915050565b600060208284031215611d2e578081fd5b813561075c8161237d565b600060208284031215611d4a578081fd5b815161075c8161237d565b600080600060608486031215611d69578182fd5b8335611d748161237d565b925060208401359150604084013567ffffffffffffffff811115611d96578182fd5b611da286828701611ccc565b9150509250925092565b60008060408385031215611dbe578182fd5b8235611dc98161237d565b9150602083013567ffffffffffffffff811115611de4578182fd5b611df085828601611ccc565b9150509250929050565b60006020808385031215611e0c578182fd5b825167ffffffffffffffff80821115611e23578384fd5b9084019060a08287031215611e36578384fd5b611e3e61225c565b8251611e498161237d565b81528284015182811115611e5b578586fd5b83019150601f82018713611e6d578485fd5b8151611e7b611cea826122b6565b8181528886838601011115611e8e578687fd5b611e9d828783018887016122de565b8286015250611eae60408401611cc1565b6040820152611ebf60608401611cc1565b6060820152611ed060808401611c6c565b60808201529695505050505050565b600060408284031215611ef0578081fd5b6040516040810181811067ffffffffffffffff82111715611f1357611f1361234e565b6040528235611f2181612392565b81526020928301359281019290925250919050565b600060208284031215611f47578081fd5b5035919050565b600060208284031215611f5f578081fd5b5051919050565b600080600080600060808688031215611f7d578081fd5b85359450602086013567ffffffffffffffff811115611f9a578182fd5b611fa688828901611c77565b9699909850959660408101359660609091013595509350505050565b60008060408385031215611fd4578182fd5b50508035926020909101359150565b600080600060608486031215611ff7578081fd5b833592506020840135915060408401356120108161237d565b809150509250925092565b60008060008060008060808789031215612033578384fd5b8635955060208701359450604087013567ffffffffffffffff80821115612058578586fd5b818901915089601f83011261206b578586fd5b813581811115612079578687fd5b8a602082850101111561208a578687fd5b6020830196508095505060608901359150808211156120a7578283fd5b506120b489828a01611c77565b979a9699509497509295939492505050565b600081518084526120de8160208601602086016122de565b601f01601f19169290920160200192915050565b6c01000000000000000000000000600160a060020a03861602815283601482015281836034830137600091016034019081529392505050565b6c01000000000000000000000000600160a060020a038516028152826014820152600082516121618160348501602087016122de565b91909101603401949350505050565b600082516121828184602087016122de565b9190910192915050565b60408152600061219f60408301856120c6565b9050600160a060020a03831660208301529392505050565b60208152600061075c60208301846120c6565b6020808252602e908201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160408201527f647920696e697469616c697a6564000000000000000000000000000000000000606082015260800190565b6020808252818101527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604082015260600190565b60405160a0810167ffffffffffffffff8111828210171561227f5761227f61234e565b60405290565b604051601f8201601f1916810167ffffffffffffffff811182821017156122ae576122ae61234e565b604052919050565b600067ffffffffffffffff8211156122d0576122d061234e565b50601f01601f191660200190565b60005b838110156122f95781810151838201526020016122e1565b83811115612308576000848401525b50505050565b6000600019821415612347577f4e487b710000000000000000000000000000000000000000000000000000000081526011600452602481fd5b5060010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600160a060020a038116811461078757600080fd5b801515811461078757600080fdfe416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a264697066735822122069b4ca6c9148fc61b26a468fdeffac68e06f7cf7b7ab9fea0038bfcbb86dd0c964736f6c63430008040033"

// DeployPrivateIssuer deploys a new Ethereum contract, binding an instance of PrivateIssuer to it.
func DeployPrivateIssuer(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PrivateIssuer, error) {
	parsed, err := abi.JSON(strings.NewReader(PrivateIssuerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PrivateIssuerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PrivateIssuer{PrivateIssuerCaller: PrivateIssuerCaller{contract: contract}, PrivateIssuerTransactor: PrivateIssuerTransactor{contract: contract}, PrivateIssuerFilterer: PrivateIssuerFilterer{contract: contract}}, nil
}

// PrivateIssuer is an auto generated Go binding around an Ethereum contract.
type PrivateIssuer struct {
	PrivateIssuerCaller     // Read-only binding to the contract
	PrivateIssuerTransactor // Write-only binding to the contract
	PrivateIssuerFilterer   // Log filterer for contract events
}

// PrivateIssuerCaller is an auto generated read-only Go binding around an Ethereum contract.
type PrivateIssuerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PrivateIssuerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PrivateIssuerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PrivateIssuerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PrivateIssuerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PrivateIssuerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PrivateIssuerSession struct {
	Contract     *PrivateIssuer    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PrivateIssuerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PrivateIssuerCallerSession struct {
	Contract *PrivateIssuerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// PrivateIssuerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PrivateIssuerTransactorSession struct {
	Contract     *PrivateIssuerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// PrivateIssuerRaw is an auto generated low-level Go binding around an Ethereum contract.
type PrivateIssuerRaw struct {
	Contract *PrivateIssuer // Generic contract binding to access the raw methods on
}

// PrivateIssuerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PrivateIssuerCallerRaw struct {
	Contract *PrivateIssuerCaller // Generic read-only contract binding to access the raw methods on
}

// PrivateIssuerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PrivateIssuerTransactorRaw struct {
	Contract *PrivateIssuerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPrivateIssuer creates a new instance of PrivateIssuer, bound to a specific deployed contract.
func NewPrivateIssuer(address common.Address, backend bind.ContractBackend) (*PrivateIssuer, error) {
	contract, err := bindPrivateIssuer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PrivateIssuer{PrivateIssuerCaller: PrivateIssuerCaller{contract: contract}, PrivateIssuerTransactor: PrivateIssuerTransactor{contract: contract}, PrivateIssuerFilterer: PrivateIssuerFilterer{contract: contract}}, nil
}

// NewPrivateIssuerCaller creates a new read-only instance of PrivateIssuer, bound to a specific deployed contract.
func NewPrivateIssuerCaller(address common.Address, caller bind.ContractCaller) (*PrivateIssuerCaller, error) {
	contract, err := bindPrivateIssuer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PrivateIssuerCaller{contract: contract}, nil
}

// NewPrivateIssuerTransactor creates a new write-only instance of PrivateIssuer, bound to a specific deployed contract.
func NewPrivateIssuerTransactor(address common.Address, transactor bind.ContractTransactor) (*PrivateIssuerTransactor, error) {
	contract, err := bindPrivateIssuer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PrivateIssuerTransactor{contract: contract}, nil
}

// NewPrivateIssuerFilterer creates a new log filterer instance of PrivateIssuer, bound to a specific deployed contract.
func NewPrivateIssuerFilterer(address common.Address, filterer bind.ContractFilterer) (*PrivateIssuerFilterer, error) {
	contract, err := bindPrivateIssuer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PrivateIssuerFilterer{contract: contract}, nil
}

// bindPrivateIssuer binds a generic wrapper to an already deployed contract.
func bindPrivateIssuer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PrivateIssuerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PrivateIssuer *PrivateIssuerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PrivateIssuer.Contract.PrivateIssuerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PrivateIssuer *PrivateIssuerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PrivateIssuer.Contract.PrivateIssuerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PrivateIssuer *PrivateIssuerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PrivateIssuer.Contract.PrivateIssuerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PrivateIssuer *PrivateIssuerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PrivateIssuer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PrivateIssuer *PrivateIssuerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PrivateIssuer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PrivateIssuer *PrivateIssuerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PrivateIssuer.Contract.contract.Transact(opts, method, params...)
}

// GetCertificateCommitment is a free data retrieval call binding the contract method 0x750e247b.
//
// Solidity: function getCertificateCommitment(uint256 certificateId) view returns(bytes32)
func (_PrivateIssuer *PrivateIssuerCaller) GetCertificateCommitment(opts *bind.CallOpts, certificateId *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _PrivateIssuer.contract.Call(opts, &out, "getCertificateCommitment", certificateId)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetCertificateCommitment is a free data retrieval call binding the contract method 0x750e247b.
//
// Solidity: function getCertificateCommitment(uint256 certificateId) view returns(bytes32)
func (_PrivateIssuer *PrivateIssuerSession) GetCertificateCommitment(certificateId *big.Int) ([32]byte, error) {
	return _PrivateIssuer.Contract.GetCertificateCommitment(&_PrivateIssuer.CallOpts, certificateId)
}

// GetCertificateCommitment is a free data retrieval call binding the contract method 0x750e247b.
//
// Solidity: function getCertificateCommitment(uint256 certificateId) view returns(bytes32)
func (_PrivateIssuer *PrivateIssuerCallerSession) GetCertificateCommitment(certificateId *big.Int) ([32]byte, error) {
	return _PrivateIssuer.Contract.GetCertificateCommitment(&_PrivateIssuer.CallOpts, certificateId)
}

// GetIssuerAddress is a free data retrieval call binding the contract method 0x32e71ead.
//
// Solidity: function getIssuerAddress() view returns(address)
func (_PrivateIssuer *PrivateIssuerCaller) GetIssuerAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PrivateIssuer.contract.Call(opts, &out, "getIssuerAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetIssuerAddress is a free data retrieval call binding the contract method 0x32e71ead.
//
// Solidity: function getIssuerAddress() view returns(address)
func (_PrivateIssuer *PrivateIssuerSession) GetIssuerAddress() (common.Address, error) {
	return _PrivateIssuer.Contract.GetIssuerAddress(&_PrivateIssuer.CallOpts)
}

// GetIssuerAddress is a free data retrieval call binding the contract method 0x32e71ead.
//
// Solidity: function getIssuerAddress() view returns(address)
func (_PrivateIssuer *PrivateIssuerCallerSession) GetIssuerAddress() (common.Address, error) {
	return _PrivateIssuer.Contract.GetIssuerAddress(&_PrivateIssuer.CallOpts)
}

// GetMigrationRequest is a free data retrieval call binding the contract method 0x1927b6a5.
//
// Solidity: function getMigrationRequest(uint256 _requestId) view returns((address,uint256,bytes32,bool))
func (_PrivateIssuer *PrivateIssuerCaller) GetMigrationRequest(opts *bind.CallOpts, _requestId *big.Int) (PrivateIssuerRequestStateChange, error) {
	var out []interface{}
	err := _PrivateIssuer.contract.Call(opts, &out, "getMigrationRequest", _requestId)

	if err != nil {
		return *new(PrivateIssuerRequestStateChange), err
	}

	out0 := *abi.ConvertType(out[0], new(PrivateIssuerRequestStateChange)).(*PrivateIssuerRequestStateChange)

	return out0, err

}

// GetMigrationRequest is a free data retrieval call binding the contract method 0x1927b6a5.
//
// Solidity: function getMigrationRequest(uint256 _requestId) view returns((address,uint256,bytes32,bool))
func (_PrivateIssuer *PrivateIssuerSession) GetMigrationRequest(_requestId *big.Int) (PrivateIssuerRequestStateChange, error) {
	return _PrivateIssuer.Contract.GetMigrationRequest(&_PrivateIssuer.CallOpts, _requestId)
}

// GetMigrationRequest is a free data retrieval call binding the contract method 0x1927b6a5.
//
// Solidity: function getMigrationRequest(uint256 _requestId) view returns((address,uint256,bytes32,bool))
func (_PrivateIssuer *PrivateIssuerCallerSession) GetMigrationRequest(_requestId *big.Int) (PrivateIssuerRequestStateChange, error) {
	return _PrivateIssuer.Contract.GetMigrationRequest(&_PrivateIssuer.CallOpts, _requestId)
}

// GetMigrationRequestId is a free data retrieval call binding the contract method 0x26eb2fd1.
//
// Solidity: function getMigrationRequestId(uint256 _certificateId) view returns(uint256 _migrationRequestId)
func (_PrivateIssuer *PrivateIssuerCaller) GetMigrationRequestId(opts *bind.CallOpts, _certificateId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PrivateIssuer.contract.Call(opts, &out, "getMigrationRequestId", _certificateId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMigrationRequestId is a free data retrieval call binding the contract method 0x26eb2fd1.
//
// Solidity: function getMigrationRequestId(uint256 _certificateId) view returns(uint256 _migrationRequestId)
func (_PrivateIssuer *PrivateIssuerSession) GetMigrationRequestId(_certificateId *big.Int) (*big.Int, error) {
	return _PrivateIssuer.Contract.GetMigrationRequestId(&_PrivateIssuer.CallOpts, _certificateId)
}

// GetMigrationRequestId is a free data retrieval call binding the contract method 0x26eb2fd1.
//
// Solidity: function getMigrationRequestId(uint256 _certificateId) view returns(uint256 _migrationRequestId)
func (_PrivateIssuer *PrivateIssuerCallerSession) GetMigrationRequestId(_certificateId *big.Int) (*big.Int, error) {
	return _PrivateIssuer.Contract.GetMigrationRequestId(&_PrivateIssuer.CallOpts, _certificateId)
}

// GetPrivateTransferRequest is a free data retrieval call binding the contract method 0xa8551fd3.
//
// Solidity: function getPrivateTransferRequest(uint256 _certificateId) view returns((address,bytes32))
func (_PrivateIssuer *PrivateIssuerCaller) GetPrivateTransferRequest(opts *bind.CallOpts, _certificateId *big.Int) (PrivateIssuerPrivateTransferRequest, error) {
	var out []interface{}
	err := _PrivateIssuer.contract.Call(opts, &out, "getPrivateTransferRequest", _certificateId)

	if err != nil {
		return *new(PrivateIssuerPrivateTransferRequest), err
	}

	out0 := *abi.ConvertType(out[0], new(PrivateIssuerPrivateTransferRequest)).(*PrivateIssuerPrivateTransferRequest)

	return out0, err

}

// GetPrivateTransferRequest is a free data retrieval call binding the contract method 0xa8551fd3.
//
// Solidity: function getPrivateTransferRequest(uint256 _certificateId) view returns((address,bytes32))
func (_PrivateIssuer *PrivateIssuerSession) GetPrivateTransferRequest(_certificateId *big.Int) (PrivateIssuerPrivateTransferRequest, error) {
	return _PrivateIssuer.Contract.GetPrivateTransferRequest(&_PrivateIssuer.CallOpts, _certificateId)
}

// GetPrivateTransferRequest is a free data retrieval call binding the contract method 0xa8551fd3.
//
// Solidity: function getPrivateTransferRequest(uint256 _certificateId) view returns((address,bytes32))
func (_PrivateIssuer *PrivateIssuerCallerSession) GetPrivateTransferRequest(_certificateId *big.Int) (PrivateIssuerPrivateTransferRequest, error) {
	return _PrivateIssuer.Contract.GetPrivateTransferRequest(&_PrivateIssuer.CallOpts, _certificateId)
}

// Issuer is a free data retrieval call binding the contract method 0x1d143848.
//
// Solidity: function issuer() view returns(address)
func (_PrivateIssuer *PrivateIssuerCaller) Issuer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PrivateIssuer.contract.Call(opts, &out, "issuer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Issuer is a free data retrieval call binding the contract method 0x1d143848.
//
// Solidity: function issuer() view returns(address)
func (_PrivateIssuer *PrivateIssuerSession) Issuer() (common.Address, error) {
	return _PrivateIssuer.Contract.Issuer(&_PrivateIssuer.CallOpts)
}

// Issuer is a free data retrieval call binding the contract method 0x1d143848.
//
// Solidity: function issuer() view returns(address)
func (_PrivateIssuer *PrivateIssuerCallerSession) Issuer() (common.Address, error) {
	return _PrivateIssuer.Contract.Issuer(&_PrivateIssuer.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PrivateIssuer *PrivateIssuerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PrivateIssuer.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PrivateIssuer *PrivateIssuerSession) Owner() (common.Address, error) {
	return _PrivateIssuer.Contract.Owner(&_PrivateIssuer.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PrivateIssuer *PrivateIssuerCallerSession) Owner() (common.Address, error) {
	return _PrivateIssuer.Contract.Owner(&_PrivateIssuer.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_PrivateIssuer *PrivateIssuerCaller) Registry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PrivateIssuer.contract.Call(opts, &out, "registry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_PrivateIssuer *PrivateIssuerSession) Registry() (common.Address, error) {
	return _PrivateIssuer.Contract.Registry(&_PrivateIssuer.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_PrivateIssuer *PrivateIssuerCallerSession) Registry() (common.Address, error) {
	return _PrivateIssuer.Contract.Registry(&_PrivateIssuer.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_PrivateIssuer *PrivateIssuerCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _PrivateIssuer.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_PrivateIssuer *PrivateIssuerSession) Version() (string, error) {
	return _PrivateIssuer.Contract.Version(&_PrivateIssuer.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_PrivateIssuer *PrivateIssuerCallerSession) Version() (string, error) {
	return _PrivateIssuer.Contract.Version(&_PrivateIssuer.CallOpts)
}

// ApproveCertificationRequestPrivate is a paid mutator transaction binding the contract method 0xc4502922.
//
// Solidity: function approveCertificationRequestPrivate(uint256 _requestId, bytes32 _commitment) returns(uint256)
func (_PrivateIssuer *PrivateIssuerTransactor) ApproveCertificationRequestPrivate(opts *bind.TransactOpts, _requestId *big.Int, _commitment [32]byte) (*types.Transaction, error) {
	return _PrivateIssuer.contract.Transact(opts, "approveCertificationRequestPrivate", _requestId, _commitment)
}

// ApproveCertificationRequestPrivate is a paid mutator transaction binding the contract method 0xc4502922.
//
// Solidity: function approveCertificationRequestPrivate(uint256 _requestId, bytes32 _commitment) returns(uint256)
func (_PrivateIssuer *PrivateIssuerSession) ApproveCertificationRequestPrivate(_requestId *big.Int, _commitment [32]byte) (*types.Transaction, error) {
	return _PrivateIssuer.Contract.ApproveCertificationRequestPrivate(&_PrivateIssuer.TransactOpts, _requestId, _commitment)
}

// ApproveCertificationRequestPrivate is a paid mutator transaction binding the contract method 0xc4502922.
//
// Solidity: function approveCertificationRequestPrivate(uint256 _requestId, bytes32 _commitment) returns(uint256)
func (_PrivateIssuer *PrivateIssuerTransactorSession) ApproveCertificationRequestPrivate(_requestId *big.Int, _commitment [32]byte) (*types.Transaction, error) {
	return _PrivateIssuer.Contract.ApproveCertificationRequestPrivate(&_PrivateIssuer.TransactOpts, _requestId, _commitment)
}

// ApprovePrivateTransfer is a paid mutator transaction binding the contract method 0x93170586.
//
// Solidity: function approvePrivateTransfer(uint256 _certificateId, (bool,bytes32)[] _proof, bytes32 _previousCommitment, bytes32 _commitment) returns(bool)
func (_PrivateIssuer *PrivateIssuerTransactor) ApprovePrivateTransfer(opts *bind.TransactOpts, _certificateId *big.Int, _proof []PrivateIssuerProof, _previousCommitment [32]byte, _commitment [32]byte) (*types.Transaction, error) {
	return _PrivateIssuer.contract.Transact(opts, "approvePrivateTransfer", _certificateId, _proof, _previousCommitment, _commitment)
}

// ApprovePrivateTransfer is a paid mutator transaction binding the contract method 0x93170586.
//
// Solidity: function approvePrivateTransfer(uint256 _certificateId, (bool,bytes32)[] _proof, bytes32 _previousCommitment, bytes32 _commitment) returns(bool)
func (_PrivateIssuer *PrivateIssuerSession) ApprovePrivateTransfer(_certificateId *big.Int, _proof []PrivateIssuerProof, _previousCommitment [32]byte, _commitment [32]byte) (*types.Transaction, error) {
	return _PrivateIssuer.Contract.ApprovePrivateTransfer(&_PrivateIssuer.TransactOpts, _certificateId, _proof, _previousCommitment, _commitment)
}

// ApprovePrivateTransfer is a paid mutator transaction binding the contract method 0x93170586.
//
// Solidity: function approvePrivateTransfer(uint256 _certificateId, (bool,bytes32)[] _proof, bytes32 _previousCommitment, bytes32 _commitment) returns(bool)
func (_PrivateIssuer *PrivateIssuerTransactorSession) ApprovePrivateTransfer(_certificateId *big.Int, _proof []PrivateIssuerProof, _previousCommitment [32]byte, _commitment [32]byte) (*types.Transaction, error) {
	return _PrivateIssuer.Contract.ApprovePrivateTransfer(&_PrivateIssuer.TransactOpts, _certificateId, _proof, _previousCommitment, _commitment)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _issuer) returns()
func (_PrivateIssuer *PrivateIssuerTransactor) Initialize(opts *bind.TransactOpts, _issuer common.Address) (*types.Transaction, error) {
	return _PrivateIssuer.contract.Transact(opts, "initialize", _issuer)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _issuer) returns()
func (_PrivateIssuer *PrivateIssuerSession) Initialize(_issuer common.Address) (*types.Transaction, error) {
	return _PrivateIssuer.Contract.Initialize(&_PrivateIssuer.TransactOpts, _issuer)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _issuer) returns()
func (_PrivateIssuer *PrivateIssuerTransactorSession) Initialize(_issuer common.Address) (*types.Transaction, error) {
	return _PrivateIssuer.Contract.Initialize(&_PrivateIssuer.TransactOpts, _issuer)
}

// IssuePrivate is a paid mutator transaction binding the contract method 0x3163e28a.
//
// Solidity: function issuePrivate(address _to, bytes32 _commitment, bytes _data) returns(uint256)
func (_PrivateIssuer *PrivateIssuerTransactor) IssuePrivate(opts *bind.TransactOpts, _to common.Address, _commitment [32]byte, _data []byte) (*types.Transaction, error) {
	return _PrivateIssuer.contract.Transact(opts, "issuePrivate", _to, _commitment, _data)
}

// IssuePrivate is a paid mutator transaction binding the contract method 0x3163e28a.
//
// Solidity: function issuePrivate(address _to, bytes32 _commitment, bytes _data) returns(uint256)
func (_PrivateIssuer *PrivateIssuerSession) IssuePrivate(_to common.Address, _commitment [32]byte, _data []byte) (*types.Transaction, error) {
	return _PrivateIssuer.Contract.IssuePrivate(&_PrivateIssuer.TransactOpts, _to, _commitment, _data)
}

// IssuePrivate is a paid mutator transaction binding the contract method 0x3163e28a.
//
// Solidity: function issuePrivate(address _to, bytes32 _commitment, bytes _data) returns(uint256)
func (_PrivateIssuer *PrivateIssuerTransactorSession) IssuePrivate(_to common.Address, _commitment [32]byte, _data []byte) (*types.Transaction, error) {
	return _PrivateIssuer.Contract.IssuePrivate(&_PrivateIssuer.TransactOpts, _to, _commitment, _data)
}

// MigrateToPublic is a paid mutator transaction binding the contract method 0xc13cbb63.
//
// Solidity: function migrateToPublic(uint256 _requestId, uint256 _volume, string _salt, (bool,bytes32)[] _proof) returns()
func (_PrivateIssuer *PrivateIssuerTransactor) MigrateToPublic(opts *bind.TransactOpts, _requestId *big.Int, _volume *big.Int, _salt string, _proof []PrivateIssuerProof) (*types.Transaction, error) {
	return _PrivateIssuer.contract.Transact(opts, "migrateToPublic", _requestId, _volume, _salt, _proof)
}

// MigrateToPublic is a paid mutator transaction binding the contract method 0xc13cbb63.
//
// Solidity: function migrateToPublic(uint256 _requestId, uint256 _volume, string _salt, (bool,bytes32)[] _proof) returns()
func (_PrivateIssuer *PrivateIssuerSession) MigrateToPublic(_requestId *big.Int, _volume *big.Int, _salt string, _proof []PrivateIssuerProof) (*types.Transaction, error) {
	return _PrivateIssuer.Contract.MigrateToPublic(&_PrivateIssuer.TransactOpts, _requestId, _volume, _salt, _proof)
}

// MigrateToPublic is a paid mutator transaction binding the contract method 0xc13cbb63.
//
// Solidity: function migrateToPublic(uint256 _requestId, uint256 _volume, string _salt, (bool,bytes32)[] _proof) returns()
func (_PrivateIssuer *PrivateIssuerTransactorSession) MigrateToPublic(_requestId *big.Int, _volume *big.Int, _salt string, _proof []PrivateIssuerProof) (*types.Transaction, error) {
	return _PrivateIssuer.Contract.MigrateToPublic(&_PrivateIssuer.TransactOpts, _requestId, _volume, _salt, _proof)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PrivateIssuer *PrivateIssuerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PrivateIssuer.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PrivateIssuer *PrivateIssuerSession) RenounceOwnership() (*types.Transaction, error) {
	return _PrivateIssuer.Contract.RenounceOwnership(&_PrivateIssuer.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PrivateIssuer *PrivateIssuerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _PrivateIssuer.Contract.RenounceOwnership(&_PrivateIssuer.TransactOpts)
}

// RequestMigrateToPublic is a paid mutator transaction binding the contract method 0xabba27b5.
//
// Solidity: function requestMigrateToPublic(uint256 _certificateId, bytes32 _ownerAddressLeafHash) returns(uint256 _migrationRequestId)
func (_PrivateIssuer *PrivateIssuerTransactor) RequestMigrateToPublic(opts *bind.TransactOpts, _certificateId *big.Int, _ownerAddressLeafHash [32]byte) (*types.Transaction, error) {
	return _PrivateIssuer.contract.Transact(opts, "requestMigrateToPublic", _certificateId, _ownerAddressLeafHash)
}

// RequestMigrateToPublic is a paid mutator transaction binding the contract method 0xabba27b5.
//
// Solidity: function requestMigrateToPublic(uint256 _certificateId, bytes32 _ownerAddressLeafHash) returns(uint256 _migrationRequestId)
func (_PrivateIssuer *PrivateIssuerSession) RequestMigrateToPublic(_certificateId *big.Int, _ownerAddressLeafHash [32]byte) (*types.Transaction, error) {
	return _PrivateIssuer.Contract.RequestMigrateToPublic(&_PrivateIssuer.TransactOpts, _certificateId, _ownerAddressLeafHash)
}

// RequestMigrateToPublic is a paid mutator transaction binding the contract method 0xabba27b5.
//
// Solidity: function requestMigrateToPublic(uint256 _certificateId, bytes32 _ownerAddressLeafHash) returns(uint256 _migrationRequestId)
func (_PrivateIssuer *PrivateIssuerTransactorSession) RequestMigrateToPublic(_certificateId *big.Int, _ownerAddressLeafHash [32]byte) (*types.Transaction, error) {
	return _PrivateIssuer.Contract.RequestMigrateToPublic(&_PrivateIssuer.TransactOpts, _certificateId, _ownerAddressLeafHash)
}

// RequestMigrateToPublicFor is a paid mutator transaction binding the contract method 0x7dbca006.
//
// Solidity: function requestMigrateToPublicFor(uint256 _certificateId, bytes32 _ownerAddressLeafHash, address _forAddress) returns(uint256 _migrationRequestId)
func (_PrivateIssuer *PrivateIssuerTransactor) RequestMigrateToPublicFor(opts *bind.TransactOpts, _certificateId *big.Int, _ownerAddressLeafHash [32]byte, _forAddress common.Address) (*types.Transaction, error) {
	return _PrivateIssuer.contract.Transact(opts, "requestMigrateToPublicFor", _certificateId, _ownerAddressLeafHash, _forAddress)
}

// RequestMigrateToPublicFor is a paid mutator transaction binding the contract method 0x7dbca006.
//
// Solidity: function requestMigrateToPublicFor(uint256 _certificateId, bytes32 _ownerAddressLeafHash, address _forAddress) returns(uint256 _migrationRequestId)
func (_PrivateIssuer *PrivateIssuerSession) RequestMigrateToPublicFor(_certificateId *big.Int, _ownerAddressLeafHash [32]byte, _forAddress common.Address) (*types.Transaction, error) {
	return _PrivateIssuer.Contract.RequestMigrateToPublicFor(&_PrivateIssuer.TransactOpts, _certificateId, _ownerAddressLeafHash, _forAddress)
}

// RequestMigrateToPublicFor is a paid mutator transaction binding the contract method 0x7dbca006.
//
// Solidity: function requestMigrateToPublicFor(uint256 _certificateId, bytes32 _ownerAddressLeafHash, address _forAddress) returns(uint256 _migrationRequestId)
func (_PrivateIssuer *PrivateIssuerTransactorSession) RequestMigrateToPublicFor(_certificateId *big.Int, _ownerAddressLeafHash [32]byte, _forAddress common.Address) (*types.Transaction, error) {
	return _PrivateIssuer.Contract.RequestMigrateToPublicFor(&_PrivateIssuer.TransactOpts, _certificateId, _ownerAddressLeafHash, _forAddress)
}

// RequestPrivateTransfer is a paid mutator transaction binding the contract method 0xb747c911.
//
// Solidity: function requestPrivateTransfer(uint256 _certificateId, bytes32 _ownerAddressLeafHash) returns()
func (_PrivateIssuer *PrivateIssuerTransactor) RequestPrivateTransfer(opts *bind.TransactOpts, _certificateId *big.Int, _ownerAddressLeafHash [32]byte) (*types.Transaction, error) {
	return _PrivateIssuer.contract.Transact(opts, "requestPrivateTransfer", _certificateId, _ownerAddressLeafHash)
}

// RequestPrivateTransfer is a paid mutator transaction binding the contract method 0xb747c911.
//
// Solidity: function requestPrivateTransfer(uint256 _certificateId, bytes32 _ownerAddressLeafHash) returns()
func (_PrivateIssuer *PrivateIssuerSession) RequestPrivateTransfer(_certificateId *big.Int, _ownerAddressLeafHash [32]byte) (*types.Transaction, error) {
	return _PrivateIssuer.Contract.RequestPrivateTransfer(&_PrivateIssuer.TransactOpts, _certificateId, _ownerAddressLeafHash)
}

// RequestPrivateTransfer is a paid mutator transaction binding the contract method 0xb747c911.
//
// Solidity: function requestPrivateTransfer(uint256 _certificateId, bytes32 _ownerAddressLeafHash) returns()
func (_PrivateIssuer *PrivateIssuerTransactorSession) RequestPrivateTransfer(_certificateId *big.Int, _ownerAddressLeafHash [32]byte) (*types.Transaction, error) {
	return _PrivateIssuer.Contract.RequestPrivateTransfer(&_PrivateIssuer.TransactOpts, _certificateId, _ownerAddressLeafHash)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PrivateIssuer *PrivateIssuerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _PrivateIssuer.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PrivateIssuer *PrivateIssuerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PrivateIssuer.Contract.TransferOwnership(&_PrivateIssuer.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PrivateIssuer *PrivateIssuerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PrivateIssuer.Contract.TransferOwnership(&_PrivateIssuer.TransactOpts, newOwner)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_PrivateIssuer *PrivateIssuerTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _PrivateIssuer.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_PrivateIssuer *PrivateIssuerSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _PrivateIssuer.Contract.UpgradeTo(&_PrivateIssuer.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_PrivateIssuer *PrivateIssuerTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _PrivateIssuer.Contract.UpgradeTo(&_PrivateIssuer.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_PrivateIssuer *PrivateIssuerTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _PrivateIssuer.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_PrivateIssuer *PrivateIssuerSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _PrivateIssuer.Contract.UpgradeToAndCall(&_PrivateIssuer.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_PrivateIssuer *PrivateIssuerTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _PrivateIssuer.Contract.UpgradeToAndCall(&_PrivateIssuer.TransactOpts, newImplementation, data)
}

// PrivateIssuerAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the PrivateIssuer contract.
type PrivateIssuerAdminChangedIterator struct {
	Event *PrivateIssuerAdminChanged // Event containing the contract specifics and raw log

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
func (it *PrivateIssuerAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PrivateIssuerAdminChanged)
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
		it.Event = new(PrivateIssuerAdminChanged)
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
func (it *PrivateIssuerAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PrivateIssuerAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PrivateIssuerAdminChanged represents a AdminChanged event raised by the PrivateIssuer contract.
type PrivateIssuerAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_PrivateIssuer *PrivateIssuerFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*PrivateIssuerAdminChangedIterator, error) {

	logs, sub, err := _PrivateIssuer.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &PrivateIssuerAdminChangedIterator{contract: _PrivateIssuer.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_PrivateIssuer *PrivateIssuerFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *PrivateIssuerAdminChanged) (event.Subscription, error) {

	logs, sub, err := _PrivateIssuer.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PrivateIssuerAdminChanged)
				if err := _PrivateIssuer.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_PrivateIssuer *PrivateIssuerFilterer) ParseAdminChanged(log types.Log) (*PrivateIssuerAdminChanged, error) {
	event := new(PrivateIssuerAdminChanged)
	if err := _PrivateIssuer.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PrivateIssuerBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the PrivateIssuer contract.
type PrivateIssuerBeaconUpgradedIterator struct {
	Event *PrivateIssuerBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *PrivateIssuerBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PrivateIssuerBeaconUpgraded)
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
		it.Event = new(PrivateIssuerBeaconUpgraded)
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
func (it *PrivateIssuerBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PrivateIssuerBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PrivateIssuerBeaconUpgraded represents a BeaconUpgraded event raised by the PrivateIssuer contract.
type PrivateIssuerBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_PrivateIssuer *PrivateIssuerFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*PrivateIssuerBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _PrivateIssuer.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &PrivateIssuerBeaconUpgradedIterator{contract: _PrivateIssuer.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_PrivateIssuer *PrivateIssuerFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *PrivateIssuerBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _PrivateIssuer.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PrivateIssuerBeaconUpgraded)
				if err := _PrivateIssuer.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_PrivateIssuer *PrivateIssuerFilterer) ParseBeaconUpgraded(log types.Log) (*PrivateIssuerBeaconUpgraded, error) {
	event := new(PrivateIssuerBeaconUpgraded)
	if err := _PrivateIssuer.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PrivateIssuerCertificateMigratedToPublicIterator is returned from FilterCertificateMigratedToPublic and is used to iterate over the raw logs and unpacked data for CertificateMigratedToPublic events raised by the PrivateIssuer contract.
type PrivateIssuerCertificateMigratedToPublicIterator struct {
	Event *PrivateIssuerCertificateMigratedToPublic // Event containing the contract specifics and raw log

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
func (it *PrivateIssuerCertificateMigratedToPublicIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PrivateIssuerCertificateMigratedToPublic)
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
		it.Event = new(PrivateIssuerCertificateMigratedToPublic)
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
func (it *PrivateIssuerCertificateMigratedToPublicIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PrivateIssuerCertificateMigratedToPublicIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PrivateIssuerCertificateMigratedToPublic represents a CertificateMigratedToPublic event raised by the PrivateIssuer contract.
type PrivateIssuerCertificateMigratedToPublic struct {
	CertificateId *big.Int
	Owner         common.Address
	Amount        *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCertificateMigratedToPublic is a free log retrieval operation binding the contract event 0x5d348ef3cec3f8493cd139d1ee84e4c84f5b987f49646d7dfa226dbe98b770c9.
//
// Solidity: event CertificateMigratedToPublic(uint256 indexed _certificateId, address indexed _owner, uint256 indexed _amount)
func (_PrivateIssuer *PrivateIssuerFilterer) FilterCertificateMigratedToPublic(opts *bind.FilterOpts, _certificateId []*big.Int, _owner []common.Address, _amount []*big.Int) (*PrivateIssuerCertificateMigratedToPublicIterator, error) {

	var _certificateIdRule []interface{}
	for _, _certificateIdItem := range _certificateId {
		_certificateIdRule = append(_certificateIdRule, _certificateIdItem)
	}
	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _amountRule []interface{}
	for _, _amountItem := range _amount {
		_amountRule = append(_amountRule, _amountItem)
	}

	logs, sub, err := _PrivateIssuer.contract.FilterLogs(opts, "CertificateMigratedToPublic", _certificateIdRule, _ownerRule, _amountRule)
	if err != nil {
		return nil, err
	}
	return &PrivateIssuerCertificateMigratedToPublicIterator{contract: _PrivateIssuer.contract, event: "CertificateMigratedToPublic", logs: logs, sub: sub}, nil
}

// WatchCertificateMigratedToPublic is a free log subscription operation binding the contract event 0x5d348ef3cec3f8493cd139d1ee84e4c84f5b987f49646d7dfa226dbe98b770c9.
//
// Solidity: event CertificateMigratedToPublic(uint256 indexed _certificateId, address indexed _owner, uint256 indexed _amount)
func (_PrivateIssuer *PrivateIssuerFilterer) WatchCertificateMigratedToPublic(opts *bind.WatchOpts, sink chan<- *PrivateIssuerCertificateMigratedToPublic, _certificateId []*big.Int, _owner []common.Address, _amount []*big.Int) (event.Subscription, error) {

	var _certificateIdRule []interface{}
	for _, _certificateIdItem := range _certificateId {
		_certificateIdRule = append(_certificateIdRule, _certificateIdItem)
	}
	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _amountRule []interface{}
	for _, _amountItem := range _amount {
		_amountRule = append(_amountRule, _amountItem)
	}

	logs, sub, err := _PrivateIssuer.contract.WatchLogs(opts, "CertificateMigratedToPublic", _certificateIdRule, _ownerRule, _amountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PrivateIssuerCertificateMigratedToPublic)
				if err := _PrivateIssuer.contract.UnpackLog(event, "CertificateMigratedToPublic", log); err != nil {
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

// ParseCertificateMigratedToPublic is a log parse operation binding the contract event 0x5d348ef3cec3f8493cd139d1ee84e4c84f5b987f49646d7dfa226dbe98b770c9.
//
// Solidity: event CertificateMigratedToPublic(uint256 indexed _certificateId, address indexed _owner, uint256 indexed _amount)
func (_PrivateIssuer *PrivateIssuerFilterer) ParseCertificateMigratedToPublic(log types.Log) (*PrivateIssuerCertificateMigratedToPublic, error) {
	event := new(PrivateIssuerCertificateMigratedToPublic)
	if err := _PrivateIssuer.contract.UnpackLog(event, "CertificateMigratedToPublic", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PrivateIssuerCommitmentUpdatedIterator is returned from FilterCommitmentUpdated and is used to iterate over the raw logs and unpacked data for CommitmentUpdated events raised by the PrivateIssuer contract.
type PrivateIssuerCommitmentUpdatedIterator struct {
	Event *PrivateIssuerCommitmentUpdated // Event containing the contract specifics and raw log

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
func (it *PrivateIssuerCommitmentUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PrivateIssuerCommitmentUpdated)
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
		it.Event = new(PrivateIssuerCommitmentUpdated)
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
func (it *PrivateIssuerCommitmentUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PrivateIssuerCommitmentUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PrivateIssuerCommitmentUpdated represents a CommitmentUpdated event raised by the PrivateIssuer contract.
type PrivateIssuerCommitmentUpdated struct {
	Owner      common.Address
	Id         *big.Int
	Commitment [32]byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCommitmentUpdated is a free log retrieval operation binding the contract event 0x20c9608d29cbc79a883dcda48c0f0a9f1babeccea0680f65ef20ce17b06dad15.
//
// Solidity: event CommitmentUpdated(address indexed _owner, uint256 indexed _id, bytes32 _commitment)
func (_PrivateIssuer *PrivateIssuerFilterer) FilterCommitmentUpdated(opts *bind.FilterOpts, _owner []common.Address, _id []*big.Int) (*PrivateIssuerCommitmentUpdatedIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _PrivateIssuer.contract.FilterLogs(opts, "CommitmentUpdated", _ownerRule, _idRule)
	if err != nil {
		return nil, err
	}
	return &PrivateIssuerCommitmentUpdatedIterator{contract: _PrivateIssuer.contract, event: "CommitmentUpdated", logs: logs, sub: sub}, nil
}

// WatchCommitmentUpdated is a free log subscription operation binding the contract event 0x20c9608d29cbc79a883dcda48c0f0a9f1babeccea0680f65ef20ce17b06dad15.
//
// Solidity: event CommitmentUpdated(address indexed _owner, uint256 indexed _id, bytes32 _commitment)
func (_PrivateIssuer *PrivateIssuerFilterer) WatchCommitmentUpdated(opts *bind.WatchOpts, sink chan<- *PrivateIssuerCommitmentUpdated, _owner []common.Address, _id []*big.Int) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _PrivateIssuer.contract.WatchLogs(opts, "CommitmentUpdated", _ownerRule, _idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PrivateIssuerCommitmentUpdated)
				if err := _PrivateIssuer.contract.UnpackLog(event, "CommitmentUpdated", log); err != nil {
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

// ParseCommitmentUpdated is a log parse operation binding the contract event 0x20c9608d29cbc79a883dcda48c0f0a9f1babeccea0680f65ef20ce17b06dad15.
//
// Solidity: event CommitmentUpdated(address indexed _owner, uint256 indexed _id, bytes32 _commitment)
func (_PrivateIssuer *PrivateIssuerFilterer) ParseCommitmentUpdated(log types.Log) (*PrivateIssuerCommitmentUpdated, error) {
	event := new(PrivateIssuerCommitmentUpdated)
	if err := _PrivateIssuer.contract.UnpackLog(event, "CommitmentUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PrivateIssuerMigrateToPublicRequestedIterator is returned from FilterMigrateToPublicRequested and is used to iterate over the raw logs and unpacked data for MigrateToPublicRequested events raised by the PrivateIssuer contract.
type PrivateIssuerMigrateToPublicRequestedIterator struct {
	Event *PrivateIssuerMigrateToPublicRequested // Event containing the contract specifics and raw log

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
func (it *PrivateIssuerMigrateToPublicRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PrivateIssuerMigrateToPublicRequested)
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
		it.Event = new(PrivateIssuerMigrateToPublicRequested)
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
func (it *PrivateIssuerMigrateToPublicRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PrivateIssuerMigrateToPublicRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PrivateIssuerMigrateToPublicRequested represents a MigrateToPublicRequested event raised by the PrivateIssuer contract.
type PrivateIssuerMigrateToPublicRequested struct {
	Owner common.Address
	Id    *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterMigrateToPublicRequested is a free log retrieval operation binding the contract event 0xd18a53554a32a70376b6c223bfb0d84fc24caa3c704ef04eaf1b868e74ee8f6a.
//
// Solidity: event MigrateToPublicRequested(address indexed _owner, uint256 indexed _id)
func (_PrivateIssuer *PrivateIssuerFilterer) FilterMigrateToPublicRequested(opts *bind.FilterOpts, _owner []common.Address, _id []*big.Int) (*PrivateIssuerMigrateToPublicRequestedIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _PrivateIssuer.contract.FilterLogs(opts, "MigrateToPublicRequested", _ownerRule, _idRule)
	if err != nil {
		return nil, err
	}
	return &PrivateIssuerMigrateToPublicRequestedIterator{contract: _PrivateIssuer.contract, event: "MigrateToPublicRequested", logs: logs, sub: sub}, nil
}

// WatchMigrateToPublicRequested is a free log subscription operation binding the contract event 0xd18a53554a32a70376b6c223bfb0d84fc24caa3c704ef04eaf1b868e74ee8f6a.
//
// Solidity: event MigrateToPublicRequested(address indexed _owner, uint256 indexed _id)
func (_PrivateIssuer *PrivateIssuerFilterer) WatchMigrateToPublicRequested(opts *bind.WatchOpts, sink chan<- *PrivateIssuerMigrateToPublicRequested, _owner []common.Address, _id []*big.Int) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}

	logs, sub, err := _PrivateIssuer.contract.WatchLogs(opts, "MigrateToPublicRequested", _ownerRule, _idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PrivateIssuerMigrateToPublicRequested)
				if err := _PrivateIssuer.contract.UnpackLog(event, "MigrateToPublicRequested", log); err != nil {
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

// ParseMigrateToPublicRequested is a log parse operation binding the contract event 0xd18a53554a32a70376b6c223bfb0d84fc24caa3c704ef04eaf1b868e74ee8f6a.
//
// Solidity: event MigrateToPublicRequested(address indexed _owner, uint256 indexed _id)
func (_PrivateIssuer *PrivateIssuerFilterer) ParseMigrateToPublicRequested(log types.Log) (*PrivateIssuerMigrateToPublicRequested, error) {
	event := new(PrivateIssuerMigrateToPublicRequested)
	if err := _PrivateIssuer.contract.UnpackLog(event, "MigrateToPublicRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PrivateIssuerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the PrivateIssuer contract.
type PrivateIssuerOwnershipTransferredIterator struct {
	Event *PrivateIssuerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PrivateIssuerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PrivateIssuerOwnershipTransferred)
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
		it.Event = new(PrivateIssuerOwnershipTransferred)
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
func (it *PrivateIssuerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PrivateIssuerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PrivateIssuerOwnershipTransferred represents a OwnershipTransferred event raised by the PrivateIssuer contract.
type PrivateIssuerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PrivateIssuer *PrivateIssuerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PrivateIssuerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PrivateIssuer.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PrivateIssuerOwnershipTransferredIterator{contract: _PrivateIssuer.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PrivateIssuer *PrivateIssuerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PrivateIssuerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PrivateIssuer.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PrivateIssuerOwnershipTransferred)
				if err := _PrivateIssuer.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PrivateIssuer *PrivateIssuerFilterer) ParseOwnershipTransferred(log types.Log) (*PrivateIssuerOwnershipTransferred, error) {
	event := new(PrivateIssuerOwnershipTransferred)
	if err := _PrivateIssuer.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PrivateIssuerPrivateCertificationRequestApprovedIterator is returned from FilterPrivateCertificationRequestApproved and is used to iterate over the raw logs and unpacked data for PrivateCertificationRequestApproved events raised by the PrivateIssuer contract.
type PrivateIssuerPrivateCertificationRequestApprovedIterator struct {
	Event *PrivateIssuerPrivateCertificationRequestApproved // Event containing the contract specifics and raw log

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
func (it *PrivateIssuerPrivateCertificationRequestApprovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PrivateIssuerPrivateCertificationRequestApproved)
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
		it.Event = new(PrivateIssuerPrivateCertificationRequestApproved)
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
func (it *PrivateIssuerPrivateCertificationRequestApprovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PrivateIssuerPrivateCertificationRequestApprovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PrivateIssuerPrivateCertificationRequestApproved represents a PrivateCertificationRequestApproved event raised by the PrivateIssuer contract.
type PrivateIssuerPrivateCertificationRequestApproved struct {
	Owner         common.Address
	Id            *big.Int
	CertificateId *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterPrivateCertificationRequestApproved is a free log retrieval operation binding the contract event 0x382c89c3e15b040e673bdec044eba907f4e9af06293e0058163ca396cecb39ff.
//
// Solidity: event PrivateCertificationRequestApproved(address indexed _owner, uint256 indexed _id, uint256 indexed _certificateId)
func (_PrivateIssuer *PrivateIssuerFilterer) FilterPrivateCertificationRequestApproved(opts *bind.FilterOpts, _owner []common.Address, _id []*big.Int, _certificateId []*big.Int) (*PrivateIssuerPrivateCertificationRequestApprovedIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}
	var _certificateIdRule []interface{}
	for _, _certificateIdItem := range _certificateId {
		_certificateIdRule = append(_certificateIdRule, _certificateIdItem)
	}

	logs, sub, err := _PrivateIssuer.contract.FilterLogs(opts, "PrivateCertificationRequestApproved", _ownerRule, _idRule, _certificateIdRule)
	if err != nil {
		return nil, err
	}
	return &PrivateIssuerPrivateCertificationRequestApprovedIterator{contract: _PrivateIssuer.contract, event: "PrivateCertificationRequestApproved", logs: logs, sub: sub}, nil
}

// WatchPrivateCertificationRequestApproved is a free log subscription operation binding the contract event 0x382c89c3e15b040e673bdec044eba907f4e9af06293e0058163ca396cecb39ff.
//
// Solidity: event PrivateCertificationRequestApproved(address indexed _owner, uint256 indexed _id, uint256 indexed _certificateId)
func (_PrivateIssuer *PrivateIssuerFilterer) WatchPrivateCertificationRequestApproved(opts *bind.WatchOpts, sink chan<- *PrivateIssuerPrivateCertificationRequestApproved, _owner []common.Address, _id []*big.Int, _certificateId []*big.Int) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _idRule []interface{}
	for _, _idItem := range _id {
		_idRule = append(_idRule, _idItem)
	}
	var _certificateIdRule []interface{}
	for _, _certificateIdItem := range _certificateId {
		_certificateIdRule = append(_certificateIdRule, _certificateIdItem)
	}

	logs, sub, err := _PrivateIssuer.contract.WatchLogs(opts, "PrivateCertificationRequestApproved", _ownerRule, _idRule, _certificateIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PrivateIssuerPrivateCertificationRequestApproved)
				if err := _PrivateIssuer.contract.UnpackLog(event, "PrivateCertificationRequestApproved", log); err != nil {
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

// ParsePrivateCertificationRequestApproved is a log parse operation binding the contract event 0x382c89c3e15b040e673bdec044eba907f4e9af06293e0058163ca396cecb39ff.
//
// Solidity: event PrivateCertificationRequestApproved(address indexed _owner, uint256 indexed _id, uint256 indexed _certificateId)
func (_PrivateIssuer *PrivateIssuerFilterer) ParsePrivateCertificationRequestApproved(log types.Log) (*PrivateIssuerPrivateCertificationRequestApproved, error) {
	event := new(PrivateIssuerPrivateCertificationRequestApproved)
	if err := _PrivateIssuer.contract.UnpackLog(event, "PrivateCertificationRequestApproved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PrivateIssuerPrivateTransferRequestedIterator is returned from FilterPrivateTransferRequested and is used to iterate over the raw logs and unpacked data for PrivateTransferRequested events raised by the PrivateIssuer contract.
type PrivateIssuerPrivateTransferRequestedIterator struct {
	Event *PrivateIssuerPrivateTransferRequested // Event containing the contract specifics and raw log

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
func (it *PrivateIssuerPrivateTransferRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PrivateIssuerPrivateTransferRequested)
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
		it.Event = new(PrivateIssuerPrivateTransferRequested)
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
func (it *PrivateIssuerPrivateTransferRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PrivateIssuerPrivateTransferRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PrivateIssuerPrivateTransferRequested represents a PrivateTransferRequested event raised by the PrivateIssuer contract.
type PrivateIssuerPrivateTransferRequested struct {
	Owner         common.Address
	CertificateId *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterPrivateTransferRequested is a free log retrieval operation binding the contract event 0xde442453a4fabb4e6fe9cd90c427b54511e2e7712b8dfe60307a8de6209db0f1.
//
// Solidity: event PrivateTransferRequested(address indexed _owner, uint256 indexed _certificateId)
func (_PrivateIssuer *PrivateIssuerFilterer) FilterPrivateTransferRequested(opts *bind.FilterOpts, _owner []common.Address, _certificateId []*big.Int) (*PrivateIssuerPrivateTransferRequestedIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _certificateIdRule []interface{}
	for _, _certificateIdItem := range _certificateId {
		_certificateIdRule = append(_certificateIdRule, _certificateIdItem)
	}

	logs, sub, err := _PrivateIssuer.contract.FilterLogs(opts, "PrivateTransferRequested", _ownerRule, _certificateIdRule)
	if err != nil {
		return nil, err
	}
	return &PrivateIssuerPrivateTransferRequestedIterator{contract: _PrivateIssuer.contract, event: "PrivateTransferRequested", logs: logs, sub: sub}, nil
}

// WatchPrivateTransferRequested is a free log subscription operation binding the contract event 0xde442453a4fabb4e6fe9cd90c427b54511e2e7712b8dfe60307a8de6209db0f1.
//
// Solidity: event PrivateTransferRequested(address indexed _owner, uint256 indexed _certificateId)
func (_PrivateIssuer *PrivateIssuerFilterer) WatchPrivateTransferRequested(opts *bind.WatchOpts, sink chan<- *PrivateIssuerPrivateTransferRequested, _owner []common.Address, _certificateId []*big.Int) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _certificateIdRule []interface{}
	for _, _certificateIdItem := range _certificateId {
		_certificateIdRule = append(_certificateIdRule, _certificateIdItem)
	}

	logs, sub, err := _PrivateIssuer.contract.WatchLogs(opts, "PrivateTransferRequested", _ownerRule, _certificateIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PrivateIssuerPrivateTransferRequested)
				if err := _PrivateIssuer.contract.UnpackLog(event, "PrivateTransferRequested", log); err != nil {
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

// ParsePrivateTransferRequested is a log parse operation binding the contract event 0xde442453a4fabb4e6fe9cd90c427b54511e2e7712b8dfe60307a8de6209db0f1.
//
// Solidity: event PrivateTransferRequested(address indexed _owner, uint256 indexed _certificateId)
func (_PrivateIssuer *PrivateIssuerFilterer) ParsePrivateTransferRequested(log types.Log) (*PrivateIssuerPrivateTransferRequested, error) {
	event := new(PrivateIssuerPrivateTransferRequested)
	if err := _PrivateIssuer.contract.UnpackLog(event, "PrivateTransferRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PrivateIssuerUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the PrivateIssuer contract.
type PrivateIssuerUpgradedIterator struct {
	Event *PrivateIssuerUpgraded // Event containing the contract specifics and raw log

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
func (it *PrivateIssuerUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PrivateIssuerUpgraded)
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
		it.Event = new(PrivateIssuerUpgraded)
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
func (it *PrivateIssuerUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PrivateIssuerUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PrivateIssuerUpgraded represents a Upgraded event raised by the PrivateIssuer contract.
type PrivateIssuerUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_PrivateIssuer *PrivateIssuerFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*PrivateIssuerUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _PrivateIssuer.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &PrivateIssuerUpgradedIterator{contract: _PrivateIssuer.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_PrivateIssuer *PrivateIssuerFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *PrivateIssuerUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _PrivateIssuer.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PrivateIssuerUpgraded)
				if err := _PrivateIssuer.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_PrivateIssuer *PrivateIssuerFilterer) ParseUpgraded(log types.Log) (*PrivateIssuerUpgraded, error) {
	event := new(PrivateIssuerUpgraded)
	if err := _PrivateIssuer.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
