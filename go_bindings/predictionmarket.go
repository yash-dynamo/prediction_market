// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// PredictionMarketMetaData contains all meta data concerning the PredictionMarket contract.
var PredictionMarketMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_usdc\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"buyA\",\"inputs\":[{\"name\":\"usdcIn\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getProbabilities\",\"inputs\":[],\"outputs\":[{\"name\":\"pA\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"pB\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"pC\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"invariantK\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"outcomeA\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractOutcomeToken\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"outcomeB\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractOutcomeToken\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"outcomeC\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractOutcomeToken\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"usdc\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractMockUSDC\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"xA\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"xB\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"xC\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"}]",
	Bin: "0x608060405234801562000010575f80fd5b5060405162002a7538038062002a758339818101604052810190620000369190620002f9565b8060035f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555030604051620000859062000286565b620000919190620003e6565b604051809103905ff080158015620000ab573d5f803e3d5ffd5b505f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555030604051620000f99062000286565b620001059190620004c7565b604051809103905ff0801580156200011f573d5f803e3d5ffd5b5060015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550306040516200016e9062000286565b6200017a9190620005a8565b604051809103905ff08015801562000194573d5f803e3d5ffd5b5060025f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550670de0b6b3a76400006064620001eb919062000623565b600481905550670de0b6b3a7640000606462000208919062000623565b600581905550670de0b6b3a7640000606462000225919062000623565b6006819055506006546006546200023d919062000623565b6005546005546200024f919062000623565b60055460045462000261919062000623565b6200026d91906200066d565b6200027991906200066d565b60078190555050620006a7565b61198480620010f183390190565b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f620002c38262000298565b9050919050565b620002d581620002b7565b8114620002e0575f80fd5b50565b5f81519050620002f381620002ca565b92915050565b5f6020828403121562000311576200031062000294565b5b5f6200032084828501620002e3565b91505092915050565b5f82825260208201905092915050565b7f4f7574636f6d65204100000000000000000000000000000000000000000000005f82015250565b5f6200036f60098362000329565b91506200037c8262000339565b602082019050919050565b7f41000000000000000000000000000000000000000000000000000000000000005f82015250565b5f620003bd60018362000329565b9150620003ca8262000387565b602082019050919050565b620003e081620002b7565b82525050565b5f6060820190508181035f830152620003ff8162000361565b905081810360208301526200041481620003af565b9050620004256040830184620003d5565b92915050565b7f4f7574636f6d65204200000000000000000000000000000000000000000000005f82015250565b5f6200046160098362000329565b91506200046e826200042b565b602082019050919050565b7f42000000000000000000000000000000000000000000000000000000000000005f82015250565b5f620004af60018362000329565b9150620004bc8262000479565b602082019050919050565b5f6060820190508181035f830152620004e08162000453565b90508181036020830152620004f581620004a1565b9050620005066040830184620003d5565b92915050565b7f4f7574636f6d65204300000000000000000000000000000000000000000000005f82015250565b5f6200054260098362000329565b91506200054f826200050c565b602082019050919050565b7f63000000000000000000000000000000000000000000000000000000000000005f82015250565b5f6200059060018362000329565b91506200059d826200055a565b602082019050919050565b5f6060820190508181035f830152620005c18162000534565b90508181036020830152620005d68162000582565b9050620005e76040830184620003d5565b92915050565b5f819050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f6200062f82620005ed565b91506200063c83620005ed565b92508282026200064c81620005ed565b91508282048414831517620006665762000665620005f6565b5b5092915050565b5f6200067982620005ed565b91506200068683620005ed565b9250828201905080821115620006a157620006a0620005f6565b5b92915050565b610a3c80620006b55f395ff3fe608060405234801561000f575f80fd5b506004361061009c575f3560e01c80635f1b0e73116100645780635f1b0e73146101385780637c28ed7914610156578063a06d363c14610174578063de161d1614610190578063ee23c1a8146101ae5761009c565b8063062f7b8c146100a0578063342ae6b6146100c057806335cff1ca146100de5780633e413bee146100fc5780635b66f21e1461011a575b5f80fd5b6100a86101cc565b6040516100b79392919061060e565b60405180910390f35b6100c861025e565b6040516100d59190610643565b60405180910390f35b6100e6610264565b6040516100f391906106d6565b60405180910390f35b610104610289565b604051610111919061070f565b60405180910390f35b6101226102ae565b60405161012f9190610643565b60405180910390f35b6101406102b4565b60405161014d91906106d6565b60405180910390f35b61015e6102d7565b60405161016b9190610643565b60405180910390f35b61018e60048036038101906101899190610756565b6102dd565b005b610198610554565b6040516101a591906106d6565b60405180910390f35b6101b6610579565b6040516101c39190610643565b60405180910390f35b5f805f806006546005546004546101e391906107ae565b6101ed91906107ae565b905080670de0b6b3a764000060045461020691906107e1565b610210919061084f565b935080670de0b6b3a764000060055461022991906107e1565b610233919061084f565b925080670de0b6b3a764000060065461024c91906107e1565b610256919061084f565b915050909192565b60055481565b60025f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60075481565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60045481565b5f811161031f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610316906108d9565b60405180910390fd5b60035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166323b872dd3330846040518463ffffffff1660e01b815260040161037d93929190610917565b6020604051808303815f875af1158015610399573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906103bd9190610981565b505f6004549050670de0b6b3a7640000826103d891906107e1565b60045f8282546103e891906107ae565b925050819055505f60045460045461040091906107e1565b60075461040d91906109ac565b90505f60065460065461042091906107e1565b60055460055461043091906107e1565b61043a91906107ae565b90505f61046482670de0b6b3a76400008561045591906107e1565b61045f919061084f565b61057f565b9050670de0b6b3a76400008160055461047d91906107e1565b610487919061084f565b600581905550670de0b6b3a7640000816006546104a491906107e1565b6104ae919061084f565b6006819055505f846004546104c391906109ac565b90505f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166340c10f1933836040518363ffffffff1660e01b815260040161051f9291906109df565b5f604051808303815f87803b158015610536575f80fd5b505af1158015610548573d5f803e3d5ffd5b50505050505050505050565b60015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60065481565b5f60038211156105e4578190505f600160028461059c919061084f565b6105a691906107ae565b90505b818110156105de5780915060028182856105c3919061084f565b6105cd91906107ae565b6105d7919061084f565b90506105a9565b506105f1565b5f82146105f057600190505b5b919050565b5f819050919050565b610608816105f6565b82525050565b5f6060820190506106215f8301866105ff565b61062e60208301856105ff565b61063b60408301846105ff565b949350505050565b5f6020820190506106565f8301846105ff565b92915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f819050919050565b5f61069e6106996106948461065c565b61067b565b61065c565b9050919050565b5f6106af82610684565b9050919050565b5f6106c0826106a5565b9050919050565b6106d0816106b6565b82525050565b5f6020820190506106e95f8301846106c7565b92915050565b5f6106f9826106a5565b9050919050565b610709816106ef565b82525050565b5f6020820190506107225f830184610700565b92915050565b5f80fd5b610735816105f6565b811461073f575f80fd5b50565b5f813590506107508161072c565b92915050565b5f6020828403121561076b5761076a610728565b5b5f61077884828501610742565b91505092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f6107b8826105f6565b91506107c3836105f6565b92508282019050808211156107db576107da610781565b5b92915050565b5f6107eb826105f6565b91506107f6836105f6565b9250828202610804816105f6565b9150828204841483151761081b5761081a610781565b5b5092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b5f610859826105f6565b9150610864836105f6565b92508261087457610873610822565b5b828204905092915050565b5f82825260208201905092915050565b7f5a65726f207472616465000000000000000000000000000000000000000000005f82015250565b5f6108c3600a8361087f565b91506108ce8261088f565b602082019050919050565b5f6020820190508181035f8301526108f0816108b7565b9050919050565b5f6109018261065c565b9050919050565b610911816108f7565b82525050565b5f60608201905061092a5f830186610908565b6109376020830185610908565b61094460408301846105ff565b949350505050565b5f8115159050919050565b6109608161094c565b811461096a575f80fd5b50565b5f8151905061097b81610957565b92915050565b5f6020828403121561099657610995610728565b5b5f6109a38482850161096d565b91505092915050565b5f6109b6826105f6565b91506109c1836105f6565b92508282039050818111156109d9576109d8610781565b5b92915050565b5f6040820190506109f25f830185610908565b6109ff60208301846105ff565b939250505056fea2646970667358221220689b4b38949793bd52f5051de895b3f5c8fa00616ed5b9d0ec385146f3f5834f64736f6c63430008140033608060405234801562000010575f80fd5b506040516200198438038062001984833981810160405281019062000036919062000396565b80838381600390816200004a919062000664565b5080600490816200005c919062000664565b5050505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603620000d2575f6040517f1e4fbdf7000000000000000000000000000000000000000000000000000000008152600401620000c9919062000759565b60405180910390fd5b620000e381620000ed60201b60201c565b5050505062000774565b5f60055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508160055f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f604051905090565b5f80fd5b5f80fd5b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6200021182620001c9565b810181811067ffffffffffffffff82111715620002335762000232620001d9565b5b80604052505050565b5f62000247620001b0565b905062000255828262000206565b919050565b5f67ffffffffffffffff821115620002775762000276620001d9565b5b6200028282620001c9565b9050602081019050919050565b5f5b83811015620002ae57808201518184015260208101905062000291565b5f8484015250505050565b5f620002cf620002c9846200025a565b6200023c565b905082815260208101848484011115620002ee57620002ed620001c5565b5b620002fb8482856200028f565b509392505050565b5f82601f8301126200031a5762000319620001c1565b5b81516200032c848260208601620002b9565b91505092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f620003608262000335565b9050919050565b620003728162000354565b81146200037d575f80fd5b50565b5f81519050620003908162000367565b92915050565b5f805f60608486031215620003b057620003af620001b9565b5b5f84015167ffffffffffffffff811115620003d057620003cf620001bd565b5b620003de8682870162000303565b935050602084015167ffffffffffffffff811115620004025762000401620001bd565b5b620004108682870162000303565b9250506040620004238682870162000380565b9150509250925092565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f60028204905060018216806200047c57607f821691505b60208210810362000492576200049162000437565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f60088302620004f67fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82620004b9565b620005028683620004b9565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f6200054c6200054662000540846200051a565b62000523565b6200051a565b9050919050565b5f819050919050565b62000567836200052c565b6200057f620005768262000553565b848454620004c5565b825550505050565b5f90565b6200059562000587565b620005a28184846200055c565b505050565b5b81811015620005c957620005bd5f826200058b565b600181019050620005a8565b5050565b601f8211156200061857620005e28162000498565b620005ed84620004aa565b81016020851015620005fd578190505b620006156200060c85620004aa565b830182620005a7565b50505b505050565b5f82821c905092915050565b5f6200063a5f19846008026200061d565b1980831691505092915050565b5f62000654838362000629565b9150826002028217905092915050565b6200066f826200042d565b67ffffffffffffffff8111156200068b576200068a620001d9565b5b62000697825462000464565b620006a4828285620005cd565b5f60209050601f831160018114620006da575f8415620006c5578287015190505b620006d1858262000647565b86555062000740565b601f198416620006ea8662000498565b5f5b828110156200071357848901518255600182019150602085019450602081019050620006ec565b868310156200073357848901516200072f601f89168262000629565b8355505b6001600288020188555050505b505050505050565b620007538162000354565b82525050565b5f6020820190506200076e5f83018462000748565b92915050565b61120280620007825f395ff3fe608060405234801561000f575f80fd5b50600436106100e8575f3560e01c8063715018a61161008a5780639dc29fac116100645780639dc29fac14610238578063a9059cbb14610254578063dd62ed3e14610284578063f2fde38b146102b4576100e8565b8063715018a6146101f25780638da5cb5b146101fc57806395d89b411461021a576100e8565b806323b872dd116100c657806323b872dd14610158578063313ce5671461018857806340c10f19146101a657806370a08231146101c2576100e8565b806306fdde03146100ec578063095ea7b31461010a57806318160ddd1461013a575b5f80fd5b6100f46102d0565b6040516101019190610e7b565b60405180910390f35b610124600480360381019061011f9190610f2c565b610360565b6040516101319190610f84565b60405180910390f35b610142610382565b60405161014f9190610fac565b60405180910390f35b610172600480360381019061016d9190610fc5565b61038b565b60405161017f9190610f84565b60405180910390f35b6101906103b9565b60405161019d9190611030565b60405180910390f35b6101c060048036038101906101bb9190610f2c565b6103c1565b005b6101dc60048036038101906101d79190611049565b6103d7565b6040516101e99190610fac565b60405180910390f35b6101fa61041c565b005b61020461042f565b6040516102119190611083565b60405180910390f35b610222610457565b60405161022f9190610e7b565b60405180910390f35b610252600480360381019061024d9190610f2c565b6104e7565b005b61026e60048036038101906102699190610f2c565b6104fd565b60405161027b9190610f84565b60405180910390f35b61029e6004803603810190610299919061109c565b61051f565b6040516102ab9190610fac565b60405180910390f35b6102ce60048036038101906102c99190611049565b6105a1565b005b6060600380546102df90611107565b80601f016020809104026020016040519081016040528092919081815260200182805461030b90611107565b80156103565780601f1061032d57610100808354040283529160200191610356565b820191905f5260205f20905b81548152906001019060200180831161033957829003601f168201915b5050505050905090565b5f8061036a610625565b905061037781858561062c565b600191505092915050565b5f600254905090565b5f80610395610625565b90506103a285828561063e565b6103ad8585856106d1565b60019150509392505050565b5f6012905090565b6103c96107c1565b6103d38282610848565b5050565b5f805f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20549050919050565b6104246107c1565b61042d5f6108c7565b565b5f60055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60606004805461046690611107565b80601f016020809104026020016040519081016040528092919081815260200182805461049290611107565b80156104dd5780601f106104b4576101008083540402835291602001916104dd565b820191905f5260205f20905b8154815290600101906020018083116104c057829003601f168201915b5050505050905090565b6104ef6107c1565b6104f9828261098a565b5050565b5f80610507610625565b90506105148185856106d1565b600191505092915050565b5f60015f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054905092915050565b6105a96107c1565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610619575f6040517f1e4fbdf70000000000000000000000000000000000000000000000000000000081526004016106109190611083565b60405180910390fd5b610622816108c7565b50565b5f33905090565b6106398383836001610a09565b505050565b5f610649848461051f565b90507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8110156106cb57818110156106bc578281836040517ffb8f41b20000000000000000000000000000000000000000000000000000000081526004016106b393929190611137565b60405180910390fd5b6106ca84848484035f610a09565b5b50505050565b5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610741575f6040517f96c6fd1e0000000000000000000000000000000000000000000000000000000081526004016107389190611083565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036107b1575f6040517fec442f050000000000000000000000000000000000000000000000000000000081526004016107a89190611083565b60405180910390fd5b6107bc838383610bd8565b505050565b6107c9610625565b73ffffffffffffffffffffffffffffffffffffffff166107e761042f565b73ffffffffffffffffffffffffffffffffffffffff16146108465761080a610625565b6040517f118cdaa700000000000000000000000000000000000000000000000000000000815260040161083d9190611083565b60405180910390fd5b565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036108b8575f6040517fec442f050000000000000000000000000000000000000000000000000000000081526004016108af9190611083565b60405180910390fd5b6108c35f8383610bd8565b5050565b5f60055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508160055f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036109fa575f6040517f96c6fd1e0000000000000000000000000000000000000000000000000000000081526004016109f19190611083565b60405180910390fd5b610a05825f83610bd8565b5050565b5f73ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1603610a79575f6040517fe602df05000000000000000000000000000000000000000000000000000000008152600401610a709190611083565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610ae9575f6040517f94280d62000000000000000000000000000000000000000000000000000000008152600401610ae09190611083565b60405180910390fd5b8160015f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055508015610bd2578273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92584604051610bc99190610fac565b60405180910390a35b50505050565b5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610c28578060025f828254610c1c9190611199565b92505081905550610cf6565b5f805f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054905081811015610cb1578381836040517fe450d38c000000000000000000000000000000000000000000000000000000008152600401610ca893929190611137565b60405180910390fd5b8181035f808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2081905550505b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610d3d578060025f8282540392505081905550610d87565b805f808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825401925050819055505b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef83604051610de49190610fac565b60405180910390a3505050565b5f81519050919050565b5f82825260208201905092915050565b5f5b83811015610e28578082015181840152602081019050610e0d565b5f8484015250505050565b5f601f19601f8301169050919050565b5f610e4d82610df1565b610e578185610dfb565b9350610e67818560208601610e0b565b610e7081610e33565b840191505092915050565b5f6020820190508181035f830152610e938184610e43565b905092915050565b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610ec882610e9f565b9050919050565b610ed881610ebe565b8114610ee2575f80fd5b50565b5f81359050610ef381610ecf565b92915050565b5f819050919050565b610f0b81610ef9565b8114610f15575f80fd5b50565b5f81359050610f2681610f02565b92915050565b5f8060408385031215610f4257610f41610e9b565b5b5f610f4f85828601610ee5565b9250506020610f6085828601610f18565b9150509250929050565b5f8115159050919050565b610f7e81610f6a565b82525050565b5f602082019050610f975f830184610f75565b92915050565b610fa681610ef9565b82525050565b5f602082019050610fbf5f830184610f9d565b92915050565b5f805f60608486031215610fdc57610fdb610e9b565b5b5f610fe986828701610ee5565b9350506020610ffa86828701610ee5565b925050604061100b86828701610f18565b9150509250925092565b5f60ff82169050919050565b61102a81611015565b82525050565b5f6020820190506110435f830184611021565b92915050565b5f6020828403121561105e5761105d610e9b565b5b5f61106b84828501610ee5565b91505092915050565b61107d81610ebe565b82525050565b5f6020820190506110965f830184611074565b92915050565b5f80604083850312156110b2576110b1610e9b565b5b5f6110bf85828601610ee5565b92505060206110d085828601610ee5565b9150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061111e57607f821691505b602082108103611131576111306110da565b5b50919050565b5f60608201905061114a5f830186611074565b6111576020830185610f9d565b6111646040830184610f9d565b949350505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f6111a382610ef9565b91506111ae83610ef9565b92508282019050808211156111c6576111c561116c565b5b9291505056fea26469706673582212201bcf72ca6e2fec90c156b43c964a6dcd666873fd40e60e4eaff253aae59d065564736f6c63430008140033",
}

// PredictionMarketABI is the input ABI used to generate the binding from.
// Deprecated: Use PredictionMarketMetaData.ABI instead.
var PredictionMarketABI = PredictionMarketMetaData.ABI

// PredictionMarketBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PredictionMarketMetaData.Bin instead.
var PredictionMarketBin = PredictionMarketMetaData.Bin

// DeployPredictionMarket deploys a new Ethereum contract, binding an instance of PredictionMarket to it.
func DeployPredictionMarket(auth *bind.TransactOpts, backend bind.ContractBackend, _usdc common.Address) (common.Address, *types.Transaction, *PredictionMarket, error) {
	parsed, err := PredictionMarketMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PredictionMarketBin), backend, _usdc)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PredictionMarket{PredictionMarketCaller: PredictionMarketCaller{contract: contract}, PredictionMarketTransactor: PredictionMarketTransactor{contract: contract}, PredictionMarketFilterer: PredictionMarketFilterer{contract: contract}}, nil
}

// PredictionMarket is an auto generated Go binding around an Ethereum contract.
type PredictionMarket struct {
	PredictionMarketCaller     // Read-only binding to the contract
	PredictionMarketTransactor // Write-only binding to the contract
	PredictionMarketFilterer   // Log filterer for contract events
}

// PredictionMarketCaller is an auto generated read-only Go binding around an Ethereum contract.
type PredictionMarketCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PredictionMarketTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PredictionMarketTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PredictionMarketFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PredictionMarketFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PredictionMarketSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PredictionMarketSession struct {
	Contract     *PredictionMarket // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PredictionMarketCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PredictionMarketCallerSession struct {
	Contract *PredictionMarketCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// PredictionMarketTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PredictionMarketTransactorSession struct {
	Contract     *PredictionMarketTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// PredictionMarketRaw is an auto generated low-level Go binding around an Ethereum contract.
type PredictionMarketRaw struct {
	Contract *PredictionMarket // Generic contract binding to access the raw methods on
}

// PredictionMarketCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PredictionMarketCallerRaw struct {
	Contract *PredictionMarketCaller // Generic read-only contract binding to access the raw methods on
}

// PredictionMarketTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PredictionMarketTransactorRaw struct {
	Contract *PredictionMarketTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPredictionMarket creates a new instance of PredictionMarket, bound to a specific deployed contract.
func NewPredictionMarket(address common.Address, backend bind.ContractBackend) (*PredictionMarket, error) {
	contract, err := bindPredictionMarket(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PredictionMarket{PredictionMarketCaller: PredictionMarketCaller{contract: contract}, PredictionMarketTransactor: PredictionMarketTransactor{contract: contract}, PredictionMarketFilterer: PredictionMarketFilterer{contract: contract}}, nil
}

// NewPredictionMarketCaller creates a new read-only instance of PredictionMarket, bound to a specific deployed contract.
func NewPredictionMarketCaller(address common.Address, caller bind.ContractCaller) (*PredictionMarketCaller, error) {
	contract, err := bindPredictionMarket(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PredictionMarketCaller{contract: contract}, nil
}

// NewPredictionMarketTransactor creates a new write-only instance of PredictionMarket, bound to a specific deployed contract.
func NewPredictionMarketTransactor(address common.Address, transactor bind.ContractTransactor) (*PredictionMarketTransactor, error) {
	contract, err := bindPredictionMarket(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PredictionMarketTransactor{contract: contract}, nil
}

// NewPredictionMarketFilterer creates a new log filterer instance of PredictionMarket, bound to a specific deployed contract.
func NewPredictionMarketFilterer(address common.Address, filterer bind.ContractFilterer) (*PredictionMarketFilterer, error) {
	contract, err := bindPredictionMarket(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PredictionMarketFilterer{contract: contract}, nil
}

// bindPredictionMarket binds a generic wrapper to an already deployed contract.
func bindPredictionMarket(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PredictionMarketMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PredictionMarket *PredictionMarketRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PredictionMarket.Contract.PredictionMarketCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PredictionMarket *PredictionMarketRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PredictionMarket.Contract.PredictionMarketTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PredictionMarket *PredictionMarketRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PredictionMarket.Contract.PredictionMarketTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PredictionMarket *PredictionMarketCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PredictionMarket.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PredictionMarket *PredictionMarketTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PredictionMarket.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PredictionMarket *PredictionMarketTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PredictionMarket.Contract.contract.Transact(opts, method, params...)
}

// GetProbabilities is a free data retrieval call binding the contract method 0x062f7b8c.
//
// Solidity: function getProbabilities() view returns(uint256 pA, uint256 pB, uint256 pC)
func (_PredictionMarket *PredictionMarketCaller) GetProbabilities(opts *bind.CallOpts) (struct {
	PA *big.Int
	PB *big.Int
	PC *big.Int
}, error) {
	var out []interface{}
	err := _PredictionMarket.contract.Call(opts, &out, "getProbabilities")

	outstruct := new(struct {
		PA *big.Int
		PB *big.Int
		PC *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PA = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.PB = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.PC = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetProbabilities is a free data retrieval call binding the contract method 0x062f7b8c.
//
// Solidity: function getProbabilities() view returns(uint256 pA, uint256 pB, uint256 pC)
func (_PredictionMarket *PredictionMarketSession) GetProbabilities() (struct {
	PA *big.Int
	PB *big.Int
	PC *big.Int
}, error) {
	return _PredictionMarket.Contract.GetProbabilities(&_PredictionMarket.CallOpts)
}

// GetProbabilities is a free data retrieval call binding the contract method 0x062f7b8c.
//
// Solidity: function getProbabilities() view returns(uint256 pA, uint256 pB, uint256 pC)
func (_PredictionMarket *PredictionMarketCallerSession) GetProbabilities() (struct {
	PA *big.Int
	PB *big.Int
	PC *big.Int
}, error) {
	return _PredictionMarket.Contract.GetProbabilities(&_PredictionMarket.CallOpts)
}

// InvariantK is a free data retrieval call binding the contract method 0x5b66f21e.
//
// Solidity: function invariantK() view returns(uint256)
func (_PredictionMarket *PredictionMarketCaller) InvariantK(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PredictionMarket.contract.Call(opts, &out, "invariantK")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InvariantK is a free data retrieval call binding the contract method 0x5b66f21e.
//
// Solidity: function invariantK() view returns(uint256)
func (_PredictionMarket *PredictionMarketSession) InvariantK() (*big.Int, error) {
	return _PredictionMarket.Contract.InvariantK(&_PredictionMarket.CallOpts)
}

// InvariantK is a free data retrieval call binding the contract method 0x5b66f21e.
//
// Solidity: function invariantK() view returns(uint256)
func (_PredictionMarket *PredictionMarketCallerSession) InvariantK() (*big.Int, error) {
	return _PredictionMarket.Contract.InvariantK(&_PredictionMarket.CallOpts)
}

// OutcomeA is a free data retrieval call binding the contract method 0x5f1b0e73.
//
// Solidity: function outcomeA() view returns(address)
func (_PredictionMarket *PredictionMarketCaller) OutcomeA(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PredictionMarket.contract.Call(opts, &out, "outcomeA")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OutcomeA is a free data retrieval call binding the contract method 0x5f1b0e73.
//
// Solidity: function outcomeA() view returns(address)
func (_PredictionMarket *PredictionMarketSession) OutcomeA() (common.Address, error) {
	return _PredictionMarket.Contract.OutcomeA(&_PredictionMarket.CallOpts)
}

// OutcomeA is a free data retrieval call binding the contract method 0x5f1b0e73.
//
// Solidity: function outcomeA() view returns(address)
func (_PredictionMarket *PredictionMarketCallerSession) OutcomeA() (common.Address, error) {
	return _PredictionMarket.Contract.OutcomeA(&_PredictionMarket.CallOpts)
}

// OutcomeB is a free data retrieval call binding the contract method 0xde161d16.
//
// Solidity: function outcomeB() view returns(address)
func (_PredictionMarket *PredictionMarketCaller) OutcomeB(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PredictionMarket.contract.Call(opts, &out, "outcomeB")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OutcomeB is a free data retrieval call binding the contract method 0xde161d16.
//
// Solidity: function outcomeB() view returns(address)
func (_PredictionMarket *PredictionMarketSession) OutcomeB() (common.Address, error) {
	return _PredictionMarket.Contract.OutcomeB(&_PredictionMarket.CallOpts)
}

// OutcomeB is a free data retrieval call binding the contract method 0xde161d16.
//
// Solidity: function outcomeB() view returns(address)
func (_PredictionMarket *PredictionMarketCallerSession) OutcomeB() (common.Address, error) {
	return _PredictionMarket.Contract.OutcomeB(&_PredictionMarket.CallOpts)
}

// OutcomeC is a free data retrieval call binding the contract method 0x35cff1ca.
//
// Solidity: function outcomeC() view returns(address)
func (_PredictionMarket *PredictionMarketCaller) OutcomeC(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PredictionMarket.contract.Call(opts, &out, "outcomeC")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OutcomeC is a free data retrieval call binding the contract method 0x35cff1ca.
//
// Solidity: function outcomeC() view returns(address)
func (_PredictionMarket *PredictionMarketSession) OutcomeC() (common.Address, error) {
	return _PredictionMarket.Contract.OutcomeC(&_PredictionMarket.CallOpts)
}

// OutcomeC is a free data retrieval call binding the contract method 0x35cff1ca.
//
// Solidity: function outcomeC() view returns(address)
func (_PredictionMarket *PredictionMarketCallerSession) OutcomeC() (common.Address, error) {
	return _PredictionMarket.Contract.OutcomeC(&_PredictionMarket.CallOpts)
}

// Usdc is a free data retrieval call binding the contract method 0x3e413bee.
//
// Solidity: function usdc() view returns(address)
func (_PredictionMarket *PredictionMarketCaller) Usdc(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PredictionMarket.contract.Call(opts, &out, "usdc")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Usdc is a free data retrieval call binding the contract method 0x3e413bee.
//
// Solidity: function usdc() view returns(address)
func (_PredictionMarket *PredictionMarketSession) Usdc() (common.Address, error) {
	return _PredictionMarket.Contract.Usdc(&_PredictionMarket.CallOpts)
}

// Usdc is a free data retrieval call binding the contract method 0x3e413bee.
//
// Solidity: function usdc() view returns(address)
func (_PredictionMarket *PredictionMarketCallerSession) Usdc() (common.Address, error) {
	return _PredictionMarket.Contract.Usdc(&_PredictionMarket.CallOpts)
}

// XA is a free data retrieval call binding the contract method 0x7c28ed79.
//
// Solidity: function xA() view returns(uint256)
func (_PredictionMarket *PredictionMarketCaller) XA(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PredictionMarket.contract.Call(opts, &out, "xA")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// XA is a free data retrieval call binding the contract method 0x7c28ed79.
//
// Solidity: function xA() view returns(uint256)
func (_PredictionMarket *PredictionMarketSession) XA() (*big.Int, error) {
	return _PredictionMarket.Contract.XA(&_PredictionMarket.CallOpts)
}

// XA is a free data retrieval call binding the contract method 0x7c28ed79.
//
// Solidity: function xA() view returns(uint256)
func (_PredictionMarket *PredictionMarketCallerSession) XA() (*big.Int, error) {
	return _PredictionMarket.Contract.XA(&_PredictionMarket.CallOpts)
}

// XB is a free data retrieval call binding the contract method 0x342ae6b6.
//
// Solidity: function xB() view returns(uint256)
func (_PredictionMarket *PredictionMarketCaller) XB(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PredictionMarket.contract.Call(opts, &out, "xB")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// XB is a free data retrieval call binding the contract method 0x342ae6b6.
//
// Solidity: function xB() view returns(uint256)
func (_PredictionMarket *PredictionMarketSession) XB() (*big.Int, error) {
	return _PredictionMarket.Contract.XB(&_PredictionMarket.CallOpts)
}

// XB is a free data retrieval call binding the contract method 0x342ae6b6.
//
// Solidity: function xB() view returns(uint256)
func (_PredictionMarket *PredictionMarketCallerSession) XB() (*big.Int, error) {
	return _PredictionMarket.Contract.XB(&_PredictionMarket.CallOpts)
}

// XC is a free data retrieval call binding the contract method 0xee23c1a8.
//
// Solidity: function xC() view returns(uint256)
func (_PredictionMarket *PredictionMarketCaller) XC(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PredictionMarket.contract.Call(opts, &out, "xC")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// XC is a free data retrieval call binding the contract method 0xee23c1a8.
//
// Solidity: function xC() view returns(uint256)
func (_PredictionMarket *PredictionMarketSession) XC() (*big.Int, error) {
	return _PredictionMarket.Contract.XC(&_PredictionMarket.CallOpts)
}

// XC is a free data retrieval call binding the contract method 0xee23c1a8.
//
// Solidity: function xC() view returns(uint256)
func (_PredictionMarket *PredictionMarketCallerSession) XC() (*big.Int, error) {
	return _PredictionMarket.Contract.XC(&_PredictionMarket.CallOpts)
}

// BuyA is a paid mutator transaction binding the contract method 0xa06d363c.
//
// Solidity: function buyA(uint256 usdcIn) returns()
func (_PredictionMarket *PredictionMarketTransactor) BuyA(opts *bind.TransactOpts, usdcIn *big.Int) (*types.Transaction, error) {
	return _PredictionMarket.contract.Transact(opts, "buyA", usdcIn)
}

// BuyA is a paid mutator transaction binding the contract method 0xa06d363c.
//
// Solidity: function buyA(uint256 usdcIn) returns()
func (_PredictionMarket *PredictionMarketSession) BuyA(usdcIn *big.Int) (*types.Transaction, error) {
	return _PredictionMarket.Contract.BuyA(&_PredictionMarket.TransactOpts, usdcIn)
}

// BuyA is a paid mutator transaction binding the contract method 0xa06d363c.
//
// Solidity: function buyA(uint256 usdcIn) returns()
func (_PredictionMarket *PredictionMarketTransactorSession) BuyA(usdcIn *big.Int) (*types.Transaction, error) {
	return _PredictionMarket.Contract.BuyA(&_PredictionMarket.TransactOpts, usdcIn)
}
