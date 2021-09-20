// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mynft

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

// MynftABI is the input ABI used to generate the binding from.
const MynftABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// MynftBin is the compiled bytecode used for deploying new contracts.
var MynftBin = "0x60806040523480156200001157600080fd5b506040518060400160405280600581526020017f4d794e46540000000000000000000000000000000000000000000000000000008152506040518060400160405280600481526020017f4d4e465400000000000000000000000000000000000000000000000000000000815250816000908051906020019062000096929190620000b8565b508060019080519060200190620000af929190620000b8565b505050620001cd565b828054620000c69062000168565b90600052602060002090601f016020900481019282620000ea576000855562000136565b82601f106200010557805160ff191683800117855562000136565b8280016001018555821562000136579182015b828111156200013557825182559160200191906001019062000118565b5b50905062000145919062000149565b5090565b5b80821115620001645760008160009055506001016200014a565b5090565b600060028204905060018216806200018157607f821691505b602082108114156200019857620001976200019e565b5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6123cc80620001dd6000396000f3fe608060405234801561001057600080fd5b50600436106100cf5760003560e01c80636352211e1161008c578063a22cb46511610066578063a22cb46514610224578063b88d4fde14610240578063c87b56dd1461025c578063e985e9c51461028c576100cf565b80636352211e146101a657806370a08231146101d657806395d89b4114610206576100cf565b806301ffc9a7146100d457806306fdde0314610104578063081812fc14610122578063095ea7b31461015257806323b872dd1461016e57806342842e0e1461018a575b600080fd5b6100ee60048036038101906100e99190611657565b6102bc565b6040516100fb91906119dd565b60405180910390f35b61010c61039e565b60405161011991906119f8565b60405180910390f35b61013c600480360381019061013791906116b1565b610430565b6040516101499190611976565b60405180910390f35b61016c60048036038101906101679190611617565b6104b5565b005b61018860048036038101906101839190611501565b6105cd565b005b6101a4600480360381019061019f9190611501565b61062d565b005b6101c060048036038101906101bb91906116b1565b61064d565b6040516101cd9190611976565b60405180910390f35b6101f060048036038101906101eb9190611494565b6106ff565b6040516101fd9190611b9a565b60405180910390f35b61020e6107b7565b60405161021b91906119f8565b60405180910390f35b61023e600480360381019061023991906115d7565b610849565b005b61025a60048036038101906102559190611554565b6109ca565b005b610276600480360381019061027191906116b1565b610a2c565b60405161028391906119f8565b60405180910390f35b6102a660048036038101906102a191906114c1565b610ad3565b6040516102b391906119dd565b60405180910390f35b60007f80ac58cd000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916148061038757507f5b5e139f000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916145b80610397575061039682610b67565b5b9050919050565b6060600080546103ad90611dbf565b80601f01602080910402602001604051908101604052809291908181526020018280546103d990611dbf565b80156104265780601f106103fb57610100808354040283529160200191610426565b820191906000526020600020905b81548152906001019060200180831161040957829003601f168201915b5050505050905090565b600061043b82610bd1565b61047a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161047190611afa565b60405180910390fd5b6004600083815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b60006104c08261064d565b90508073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415610531576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161052890611b5a565b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff16610550610c3d565b73ffffffffffffffffffffffffffffffffffffffff16148061057f575061057e81610579610c3d565b610ad3565b5b6105be576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105b590611a9a565b60405180910390fd5b6105c88383610c45565b505050565b6105de6105d8610c3d565b82610cfe565b61061d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161061490611b7a565b60405180910390fd5b610628838383610ddc565b505050565b610648838383604051806020016040528060008152506109ca565b505050565b6000806002600084815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614156106f6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106ed90611ada565b60405180910390fd5b80915050919050565b60008073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415610770576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161076790611aba565b60405180910390fd5b600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b6060600180546107c690611dbf565b80601f01602080910402602001604051908101604052809291908181526020018280546107f290611dbf565b801561083f5780601f106108145761010080835404028352916020019161083f565b820191906000526020600020905b81548152906001019060200180831161082257829003601f168201915b5050505050905090565b610851610c3d565b73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614156108bf576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108b690611a5a565b60405180910390fd5b80600560006108cc610c3d565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055508173ffffffffffffffffffffffffffffffffffffffff16610979610c3d565b73ffffffffffffffffffffffffffffffffffffffff167f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31836040516109be91906119dd565b60405180910390a35050565b6109db6109d5610c3d565b83610cfe565b610a1a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a1190611b7a565b60405180910390fd5b610a2684848484611038565b50505050565b6060610a3782610bd1565b610a76576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a6d90611b3a565b60405180910390fd5b6000610a80611094565b90506000815111610aa05760405180602001604052806000815250610acb565b80610aaa846110ab565b604051602001610abb929190611952565b6040516020818303038152906040525b915050919050565b6000600560008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16905092915050565b60007f01ffc9a7000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916149050919050565b60008073ffffffffffffffffffffffffffffffffffffffff166002600084815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614159050919050565b600033905090565b816004600083815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550808273ffffffffffffffffffffffffffffffffffffffff16610cb88361064d565b73ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560405160405180910390a45050565b6000610d0982610bd1565b610d48576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d3f90611a7a565b60405180910390fd5b6000610d538361064d565b90508073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff161480610dc257508373ffffffffffffffffffffffffffffffffffffffff16610daa84610430565b73ffffffffffffffffffffffffffffffffffffffff16145b80610dd35750610dd28185610ad3565b5b91505092915050565b8273ffffffffffffffffffffffffffffffffffffffff16610dfc8261064d565b73ffffffffffffffffffffffffffffffffffffffff1614610e52576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e4990611b1a565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415610ec2576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610eb990611a3a565b60405180910390fd5b610ecd83838361120c565b610ed8600082610c45565b6001600360008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254610f289190611cd5565b925050819055506001600360008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000828254610f7f9190611c4e565b92505081905550816002600083815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550808273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef60405160405180910390a4505050565b611043848484610ddc565b61104f84848484611211565b61108e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161108590611a1a565b60405180910390fd5b50505050565b606060405180602001604052806000815250905090565b606060008214156110f3576040518060400160405280600181526020017f30000000000000000000000000000000000000000000000000000000000000008152509050611207565b600082905060005b6000821461112557808061110e90611e22565b915050600a8261111e9190611ca4565b91506110fb565b60008167ffffffffffffffff81111561114157611140611f58565b5b6040519080825280601f01601f1916602001820160405280156111735781602001600182028036833780820191505090505b5090505b600085146112005760018261118c9190611cd5565b9150600a8561119b9190611e6b565b60306111a79190611c4e565b60f81b8183815181106111bd576111bc611f29565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600a856111f99190611ca4565b9450611177565b8093505050505b919050565b505050565b60006112328473ffffffffffffffffffffffffffffffffffffffff166113a8565b1561139b578373ffffffffffffffffffffffffffffffffffffffff1663150b7a0261125b610c3d565b8786866040518563ffffffff1660e01b815260040161127d9493929190611991565b602060405180830381600087803b15801561129757600080fd5b505af19250505080156112c857506040513d601f19601f820116820180604052508101906112c59190611684565b60015b61134b573d80600081146112f8576040519150601f19603f3d011682016040523d82523d6000602084013e6112fd565b606091505b50600081511415611343576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161133a90611a1a565b60405180910390fd5b805181602001fd5b63150b7a0260e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916817bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916149150506113a0565b600190505b949350505050565b600080823b905060008111915050919050565b60006113ce6113c984611bda565b611bb5565b9050828152602081018484840111156113ea576113e9611f8c565b5b6113f5848285611d7d565b509392505050565b60008135905061140c8161233a565b92915050565b60008135905061142181612351565b92915050565b60008135905061143681612368565b92915050565b60008151905061144b81612368565b92915050565b600082601f83011261146657611465611f87565b5b81356114768482602086016113bb565b91505092915050565b60008135905061148e8161237f565b92915050565b6000602082840312156114aa576114a9611f96565b5b60006114b8848285016113fd565b91505092915050565b600080604083850312156114d8576114d7611f96565b5b60006114e6858286016113fd565b92505060206114f7858286016113fd565b9150509250929050565b60008060006060848603121561151a57611519611f96565b5b6000611528868287016113fd565b9350506020611539868287016113fd565b925050604061154a8682870161147f565b9150509250925092565b6000806000806080858703121561156e5761156d611f96565b5b600061157c878288016113fd565b945050602061158d878288016113fd565b935050604061159e8782880161147f565b925050606085013567ffffffffffffffff8111156115bf576115be611f91565b5b6115cb87828801611451565b91505092959194509250565b600080604083850312156115ee576115ed611f96565b5b60006115fc858286016113fd565b925050602061160d85828601611412565b9150509250929050565b6000806040838503121561162e5761162d611f96565b5b600061163c858286016113fd565b925050602061164d8582860161147f565b9150509250929050565b60006020828403121561166d5761166c611f96565b5b600061167b84828501611427565b91505092915050565b60006020828403121561169a57611699611f96565b5b60006116a88482850161143c565b91505092915050565b6000602082840312156116c7576116c6611f96565b5b60006116d58482850161147f565b91505092915050565b6116e781611d09565b82525050565b6116f681611d1b565b82525050565b600061170782611c0b565b6117118185611c21565b9350611721818560208601611d8c565b61172a81611f9b565b840191505092915050565b600061174082611c16565b61174a8185611c32565b935061175a818560208601611d8c565b61176381611f9b565b840191505092915050565b600061177982611c16565b6117838185611c43565b9350611793818560208601611d8c565b80840191505092915050565b60006117ac603283611c32565b91506117b782611fac565b604082019050919050565b60006117cf602483611c32565b91506117da82611ffb565b604082019050919050565b60006117f2601983611c32565b91506117fd8261204a565b602082019050919050565b6000611815602c83611c32565b915061182082612073565b604082019050919050565b6000611838603883611c32565b9150611843826120c2565b604082019050919050565b600061185b602a83611c32565b915061186682612111565b604082019050919050565b600061187e602983611c32565b915061188982612160565b604082019050919050565b60006118a1602c83611c32565b91506118ac826121af565b604082019050919050565b60006118c4602983611c32565b91506118cf826121fe565b604082019050919050565b60006118e7602f83611c32565b91506118f28261224d565b604082019050919050565b600061190a602183611c32565b91506119158261229c565b604082019050919050565b600061192d603183611c32565b9150611938826122eb565b604082019050919050565b61194c81611d73565b82525050565b600061195e828561176e565b915061196a828461176e565b91508190509392505050565b600060208201905061198b60008301846116de565b92915050565b60006080820190506119a660008301876116de565b6119b360208301866116de565b6119c06040830185611943565b81810360608301526119d281846116fc565b905095945050505050565b60006020820190506119f260008301846116ed565b92915050565b60006020820190508181036000830152611a128184611735565b905092915050565b60006020820190508181036000830152611a338161179f565b9050919050565b60006020820190508181036000830152611a53816117c2565b9050919050565b60006020820190508181036000830152611a73816117e5565b9050919050565b60006020820190508181036000830152611a9381611808565b9050919050565b60006020820190508181036000830152611ab38161182b565b9050919050565b60006020820190508181036000830152611ad38161184e565b9050919050565b60006020820190508181036000830152611af381611871565b9050919050565b60006020820190508181036000830152611b1381611894565b9050919050565b60006020820190508181036000830152611b33816118b7565b9050919050565b60006020820190508181036000830152611b53816118da565b9050919050565b60006020820190508181036000830152611b73816118fd565b9050919050565b60006020820190508181036000830152611b9381611920565b9050919050565b6000602082019050611baf6000830184611943565b92915050565b6000611bbf611bd0565b9050611bcb8282611df1565b919050565b6000604051905090565b600067ffffffffffffffff821115611bf557611bf4611f58565b5b611bfe82611f9b565b9050602081019050919050565b600081519050919050565b600081519050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600081905092915050565b6000611c5982611d73565b9150611c6483611d73565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03821115611c9957611c98611e9c565b5b828201905092915050565b6000611caf82611d73565b9150611cba83611d73565b925082611cca57611cc9611ecb565b5b828204905092915050565b6000611ce082611d73565b9150611ceb83611d73565b925082821015611cfe57611cfd611e9c565b5b828203905092915050565b6000611d1482611d53565b9050919050565b60008115159050919050565b60007fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b82818337600083830152505050565b60005b83811015611daa578082015181840152602081019050611d8f565b83811115611db9576000848401525b50505050565b60006002820490506001821680611dd757607f821691505b60208210811415611deb57611dea611efa565b5b50919050565b611dfa82611f9b565b810181811067ffffffffffffffff82111715611e1957611e18611f58565b5b80604052505050565b6000611e2d82611d73565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff821415611e6057611e5f611e9c565b5b600182019050919050565b6000611e7682611d73565b9150611e8183611d73565b925082611e9157611e90611ecb565b5b828206905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4552433732313a207472616e7366657220746f206e6f6e20455243373231526560008201527f63656976657220696d706c656d656e7465720000000000000000000000000000602082015250565b7f4552433732313a207472616e7366657220746f20746865207a65726f2061646460008201527f7265737300000000000000000000000000000000000000000000000000000000602082015250565b7f4552433732313a20617070726f766520746f2063616c6c657200000000000000600082015250565b7f4552433732313a206f70657261746f7220717565727920666f72206e6f6e657860008201527f697374656e7420746f6b656e0000000000000000000000000000000000000000602082015250565b7f4552433732313a20617070726f76652063616c6c6572206973206e6f74206f7760008201527f6e6572206e6f7220617070726f76656420666f7220616c6c0000000000000000602082015250565b7f4552433732313a2062616c616e636520717565727920666f7220746865207a6560008201527f726f206164647265737300000000000000000000000000000000000000000000602082015250565b7f4552433732313a206f776e657220717565727920666f72206e6f6e657869737460008201527f656e7420746f6b656e0000000000000000000000000000000000000000000000602082015250565b7f4552433732313a20617070726f76656420717565727920666f72206e6f6e657860008201527f697374656e7420746f6b656e0000000000000000000000000000000000000000602082015250565b7f4552433732313a207472616e73666572206f6620746f6b656e2074686174206960008201527f73206e6f74206f776e0000000000000000000000000000000000000000000000602082015250565b7f4552433732314d657461646174613a2055524920717565727920666f72206e6f60008201527f6e6578697374656e7420746f6b656e0000000000000000000000000000000000602082015250565b7f4552433732313a20617070726f76616c20746f2063757272656e74206f776e6560008201527f7200000000000000000000000000000000000000000000000000000000000000602082015250565b7f4552433732313a207472616e736665722063616c6c6572206973206e6f74206f60008201527f776e6572206e6f7220617070726f766564000000000000000000000000000000602082015250565b61234381611d09565b811461234e57600080fd5b50565b61235a81611d1b565b811461236557600080fd5b50565b61237181611d27565b811461237c57600080fd5b50565b61238881611d73565b811461239357600080fd5b5056fea26469706673582212201b550dc5cc3081c4f7fc82cca1b76490da2af0b80ca9e9bc136c44f5f40f007564736f6c63430008060033"

// DeployMynft deploys a new Ethereum contract, binding an instance of Mynft to it.
func DeployMynft(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Mynft, error) {
	parsed, err := abi.JSON(strings.NewReader(MynftABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MynftBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Mynft{MynftCaller: MynftCaller{contract: contract}, MynftTransactor: MynftTransactor{contract: contract}, MynftFilterer: MynftFilterer{contract: contract}}, nil
}

// Mynft is an auto generated Go binding around an Ethereum contract.
type Mynft struct {
	MynftCaller     // Read-only binding to the contract
	MynftTransactor // Write-only binding to the contract
	MynftFilterer   // Log filterer for contract events
}

// MynftCaller is an auto generated read-only Go binding around an Ethereum contract.
type MynftCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MynftTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MynftTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MynftFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MynftFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MynftSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MynftSession struct {
	Contract     *Mynft            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MynftCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MynftCallerSession struct {
	Contract *MynftCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MynftTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MynftTransactorSession struct {
	Contract     *MynftTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MynftRaw is an auto generated low-level Go binding around an Ethereum contract.
type MynftRaw struct {
	Contract *Mynft // Generic contract binding to access the raw methods on
}

// MynftCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MynftCallerRaw struct {
	Contract *MynftCaller // Generic read-only contract binding to access the raw methods on
}

// MynftTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MynftTransactorRaw struct {
	Contract *MynftTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMynft creates a new instance of Mynft, bound to a specific deployed contract.
func NewMynft(address common.Address, backend bind.ContractBackend) (*Mynft, error) {
	contract, err := bindMynft(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Mynft{MynftCaller: MynftCaller{contract: contract}, MynftTransactor: MynftTransactor{contract: contract}, MynftFilterer: MynftFilterer{contract: contract}}, nil
}

// NewMynftCaller creates a new read-only instance of Mynft, bound to a specific deployed contract.
func NewMynftCaller(address common.Address, caller bind.ContractCaller) (*MynftCaller, error) {
	contract, err := bindMynft(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MynftCaller{contract: contract}, nil
}

// NewMynftTransactor creates a new write-only instance of Mynft, bound to a specific deployed contract.
func NewMynftTransactor(address common.Address, transactor bind.ContractTransactor) (*MynftTransactor, error) {
	contract, err := bindMynft(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MynftTransactor{contract: contract}, nil
}

// NewMynftFilterer creates a new log filterer instance of Mynft, bound to a specific deployed contract.
func NewMynftFilterer(address common.Address, filterer bind.ContractFilterer) (*MynftFilterer, error) {
	contract, err := bindMynft(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MynftFilterer{contract: contract}, nil
}

// bindMynft binds a generic wrapper to an already deployed contract.
func bindMynft(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MynftABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mynft *MynftRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Mynft.Contract.MynftCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mynft *MynftRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mynft.Contract.MynftTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mynft *MynftRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mynft.Contract.MynftTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mynft *MynftCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Mynft.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mynft *MynftTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mynft.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mynft *MynftTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mynft.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Mynft *MynftCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Mynft.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Mynft *MynftSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Mynft.Contract.BalanceOf(&_Mynft.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Mynft *MynftCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Mynft.Contract.BalanceOf(&_Mynft.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Mynft *MynftCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Mynft.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Mynft *MynftSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Mynft.Contract.GetApproved(&_Mynft.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Mynft *MynftCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Mynft.Contract.GetApproved(&_Mynft.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Mynft *MynftCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Mynft.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Mynft *MynftSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Mynft.Contract.IsApprovedForAll(&_Mynft.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Mynft *MynftCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Mynft.Contract.IsApprovedForAll(&_Mynft.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Mynft *MynftCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Mynft.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Mynft *MynftSession) Name() (string, error) {
	return _Mynft.Contract.Name(&_Mynft.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Mynft *MynftCallerSession) Name() (string, error) {
	return _Mynft.Contract.Name(&_Mynft.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Mynft *MynftCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Mynft.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Mynft *MynftSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Mynft.Contract.OwnerOf(&_Mynft.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Mynft *MynftCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Mynft.Contract.OwnerOf(&_Mynft.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Mynft *MynftCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Mynft.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Mynft *MynftSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Mynft.Contract.SupportsInterface(&_Mynft.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Mynft *MynftCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Mynft.Contract.SupportsInterface(&_Mynft.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Mynft *MynftCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Mynft.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Mynft *MynftSession) Symbol() (string, error) {
	return _Mynft.Contract.Symbol(&_Mynft.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Mynft *MynftCallerSession) Symbol() (string, error) {
	return _Mynft.Contract.Symbol(&_Mynft.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Mynft *MynftCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Mynft.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Mynft *MynftSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Mynft.Contract.TokenURI(&_Mynft.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Mynft *MynftCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Mynft.Contract.TokenURI(&_Mynft.CallOpts, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Mynft *MynftTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Mynft.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Mynft *MynftSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Mynft.Contract.Approve(&_Mynft.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Mynft *MynftTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Mynft.Contract.Approve(&_Mynft.TransactOpts, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Mynft *MynftTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Mynft.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Mynft *MynftSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Mynft.Contract.SafeTransferFrom(&_Mynft.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Mynft *MynftTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Mynft.Contract.SafeTransferFrom(&_Mynft.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Mynft *MynftTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Mynft.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Mynft *MynftSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Mynft.Contract.SafeTransferFrom0(&_Mynft.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Mynft *MynftTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Mynft.Contract.SafeTransferFrom0(&_Mynft.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Mynft *MynftTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Mynft.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Mynft *MynftSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Mynft.Contract.SetApprovalForAll(&_Mynft.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Mynft *MynftTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Mynft.Contract.SetApprovalForAll(&_Mynft.TransactOpts, operator, approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Mynft *MynftTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Mynft.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Mynft *MynftSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Mynft.Contract.TransferFrom(&_Mynft.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Mynft *MynftTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Mynft.Contract.TransferFrom(&_Mynft.TransactOpts, from, to, tokenId)
}

// MynftApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Mynft contract.
type MynftApprovalIterator struct {
	Event *MynftApproval // Event containing the contract specifics and raw log

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
func (it *MynftApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MynftApproval)
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
		it.Event = new(MynftApproval)
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
func (it *MynftApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MynftApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MynftApproval represents a Approval event raised by the Mynft contract.
type MynftApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Mynft *MynftFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*MynftApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Mynft.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &MynftApprovalIterator{contract: _Mynft.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Mynft *MynftFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *MynftApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Mynft.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MynftApproval)
				if err := _Mynft.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Mynft *MynftFilterer) ParseApproval(log types.Log) (*MynftApproval, error) {
	event := new(MynftApproval)
	if err := _Mynft.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MynftApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Mynft contract.
type MynftApprovalForAllIterator struct {
	Event *MynftApprovalForAll // Event containing the contract specifics and raw log

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
func (it *MynftApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MynftApprovalForAll)
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
		it.Event = new(MynftApprovalForAll)
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
func (it *MynftApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MynftApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MynftApprovalForAll represents a ApprovalForAll event raised by the Mynft contract.
type MynftApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Mynft *MynftFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*MynftApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Mynft.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &MynftApprovalForAllIterator{contract: _Mynft.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Mynft *MynftFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *MynftApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Mynft.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MynftApprovalForAll)
				if err := _Mynft.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Mynft *MynftFilterer) ParseApprovalForAll(log types.Log) (*MynftApprovalForAll, error) {
	event := new(MynftApprovalForAll)
	if err := _Mynft.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MynftTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Mynft contract.
type MynftTransferIterator struct {
	Event *MynftTransfer // Event containing the contract specifics and raw log

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
func (it *MynftTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MynftTransfer)
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
		it.Event = new(MynftTransfer)
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
func (it *MynftTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MynftTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MynftTransfer represents a Transfer event raised by the Mynft contract.
type MynftTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Mynft *MynftFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*MynftTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Mynft.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &MynftTransferIterator{contract: _Mynft.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Mynft *MynftFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *MynftTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Mynft.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MynftTransfer)
				if err := _Mynft.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Mynft *MynftFilterer) ParseTransfer(log types.Log) (*MynftTransfer, error) {
	event := new(MynftTransfer)
	if err := _Mynft.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
