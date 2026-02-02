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

// OutcomeTokenMetaData contains all meta data concerning the OutcomeToken contract.
var OutcomeTokenMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"name_\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"symbol_\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"market\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"burn\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"mint\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ERC20InsufficientAllowance\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InsufficientBalance\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidApprover\",\"inputs\":[{\"name\":\"approver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidReceiver\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidSender\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidSpender\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
	Bin: "0x608060405234801562000010575f80fd5b506040516200198438038062001984833981810160405281019062000036919062000396565b80838381600390816200004a919062000664565b5080600490816200005c919062000664565b5050505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603620000d2575f6040517f1e4fbdf7000000000000000000000000000000000000000000000000000000008152600401620000c9919062000759565b60405180910390fd5b620000e381620000ed60201b60201c565b5050505062000774565b5f60055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508160055f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f604051905090565b5f80fd5b5f80fd5b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6200021182620001c9565b810181811067ffffffffffffffff82111715620002335762000232620001d9565b5b80604052505050565b5f62000247620001b0565b905062000255828262000206565b919050565b5f67ffffffffffffffff821115620002775762000276620001d9565b5b6200028282620001c9565b9050602081019050919050565b5f5b83811015620002ae57808201518184015260208101905062000291565b5f8484015250505050565b5f620002cf620002c9846200025a565b6200023c565b905082815260208101848484011115620002ee57620002ed620001c5565b5b620002fb8482856200028f565b509392505050565b5f82601f8301126200031a5762000319620001c1565b5b81516200032c848260208601620002b9565b91505092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f620003608262000335565b9050919050565b620003728162000354565b81146200037d575f80fd5b50565b5f81519050620003908162000367565b92915050565b5f805f60608486031215620003b057620003af620001b9565b5b5f84015167ffffffffffffffff811115620003d057620003cf620001bd565b5b620003de8682870162000303565b935050602084015167ffffffffffffffff811115620004025762000401620001bd565b5b620004108682870162000303565b9250506040620004238682870162000380565b9150509250925092565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f60028204905060018216806200047c57607f821691505b60208210810362000492576200049162000437565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f60088302620004f67fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82620004b9565b620005028683620004b9565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f6200054c6200054662000540846200051a565b62000523565b6200051a565b9050919050565b5f819050919050565b62000567836200052c565b6200057f620005768262000553565b848454620004c5565b825550505050565b5f90565b6200059562000587565b620005a28184846200055c565b505050565b5b81811015620005c957620005bd5f826200058b565b600181019050620005a8565b5050565b601f8211156200061857620005e28162000498565b620005ed84620004aa565b81016020851015620005fd578190505b620006156200060c85620004aa565b830182620005a7565b50505b505050565b5f82821c905092915050565b5f6200063a5f19846008026200061d565b1980831691505092915050565b5f62000654838362000629565b9150826002028217905092915050565b6200066f826200042d565b67ffffffffffffffff8111156200068b576200068a620001d9565b5b62000697825462000464565b620006a4828285620005cd565b5f60209050601f831160018114620006da575f8415620006c5578287015190505b620006d1858262000647565b86555062000740565b601f198416620006ea8662000498565b5f5b828110156200071357848901518255600182019150602085019450602081019050620006ec565b868310156200073357848901516200072f601f89168262000629565b8355505b6001600288020188555050505b505050505050565b620007538162000354565b82525050565b5f6020820190506200076e5f83018462000748565b92915050565b61120280620007825f395ff3fe608060405234801561000f575f80fd5b50600436106100e8575f3560e01c8063715018a61161008a5780639dc29fac116100645780639dc29fac14610238578063a9059cbb14610254578063dd62ed3e14610284578063f2fde38b146102b4576100e8565b8063715018a6146101f25780638da5cb5b146101fc57806395d89b411461021a576100e8565b806323b872dd116100c657806323b872dd14610158578063313ce5671461018857806340c10f19146101a657806370a08231146101c2576100e8565b806306fdde03146100ec578063095ea7b31461010a57806318160ddd1461013a575b5f80fd5b6100f46102d0565b6040516101019190610e7b565b60405180910390f35b610124600480360381019061011f9190610f2c565b610360565b6040516101319190610f84565b60405180910390f35b610142610382565b60405161014f9190610fac565b60405180910390f35b610172600480360381019061016d9190610fc5565b61038b565b60405161017f9190610f84565b60405180910390f35b6101906103b9565b60405161019d9190611030565b60405180910390f35b6101c060048036038101906101bb9190610f2c565b6103c1565b005b6101dc60048036038101906101d79190611049565b6103d7565b6040516101e99190610fac565b60405180910390f35b6101fa61041c565b005b61020461042f565b6040516102119190611083565b60405180910390f35b610222610457565b60405161022f9190610e7b565b60405180910390f35b610252600480360381019061024d9190610f2c565b6104e7565b005b61026e60048036038101906102699190610f2c565b6104fd565b60405161027b9190610f84565b60405180910390f35b61029e6004803603810190610299919061109c565b61051f565b6040516102ab9190610fac565b60405180910390f35b6102ce60048036038101906102c99190611049565b6105a1565b005b6060600380546102df90611107565b80601f016020809104026020016040519081016040528092919081815260200182805461030b90611107565b80156103565780601f1061032d57610100808354040283529160200191610356565b820191905f5260205f20905b81548152906001019060200180831161033957829003601f168201915b5050505050905090565b5f8061036a610625565b905061037781858561062c565b600191505092915050565b5f600254905090565b5f80610395610625565b90506103a285828561063e565b6103ad8585856106d1565b60019150509392505050565b5f6012905090565b6103c96107c1565b6103d38282610848565b5050565b5f805f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20549050919050565b6104246107c1565b61042d5f6108c7565b565b5f60055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60606004805461046690611107565b80601f016020809104026020016040519081016040528092919081815260200182805461049290611107565b80156104dd5780601f106104b4576101008083540402835291602001916104dd565b820191905f5260205f20905b8154815290600101906020018083116104c057829003601f168201915b5050505050905090565b6104ef6107c1565b6104f9828261098a565b5050565b5f80610507610625565b90506105148185856106d1565b600191505092915050565b5f60015f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054905092915050565b6105a96107c1565b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610619575f6040517f1e4fbdf70000000000000000000000000000000000000000000000000000000081526004016106109190611083565b60405180910390fd5b610622816108c7565b50565b5f33905090565b6106398383836001610a09565b505050565b5f610649848461051f565b90507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8110156106cb57818110156106bc578281836040517ffb8f41b20000000000000000000000000000000000000000000000000000000081526004016106b393929190611137565b60405180910390fd5b6106ca84848484035f610a09565b5b50505050565b5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610741575f6040517f96c6fd1e0000000000000000000000000000000000000000000000000000000081526004016107389190611083565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036107b1575f6040517fec442f050000000000000000000000000000000000000000000000000000000081526004016107a89190611083565b60405180910390fd5b6107bc838383610bd8565b505050565b6107c9610625565b73ffffffffffffffffffffffffffffffffffffffff166107e761042f565b73ffffffffffffffffffffffffffffffffffffffff16146108465761080a610625565b6040517f118cdaa700000000000000000000000000000000000000000000000000000000815260040161083d9190611083565b60405180910390fd5b565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036108b8575f6040517fec442f050000000000000000000000000000000000000000000000000000000081526004016108af9190611083565b60405180910390fd5b6108c35f8383610bd8565b5050565b5f60055f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690508160055f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508173ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a35050565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036109fa575f6040517f96c6fd1e0000000000000000000000000000000000000000000000000000000081526004016109f19190611083565b60405180910390fd5b610a05825f83610bd8565b5050565b5f73ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1603610a79575f6040517fe602df05000000000000000000000000000000000000000000000000000000008152600401610a709190611083565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610ae9575f6040517f94280d62000000000000000000000000000000000000000000000000000000008152600401610ae09190611083565b60405180910390fd5b8160015f8673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055508015610bd2578273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92584604051610bc99190610fac565b60405180910390a35b50505050565b5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610c28578060025f828254610c1c9190611199565b92505081905550610cf6565b5f805f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2054905081811015610cb1578381836040517fe450d38c000000000000000000000000000000000000000000000000000000008152600401610ca893929190611137565b60405180910390fd5b8181035f808673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2081905550505b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610d3d578060025f8282540392505081905550610d87565b805f808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825401925050819055505b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef83604051610de49190610fac565b60405180910390a3505050565b5f81519050919050565b5f82825260208201905092915050565b5f5b83811015610e28578082015181840152602081019050610e0d565b5f8484015250505050565b5f601f19601f8301169050919050565b5f610e4d82610df1565b610e578185610dfb565b9350610e67818560208601610e0b565b610e7081610e33565b840191505092915050565b5f6020820190508181035f830152610e938184610e43565b905092915050565b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610ec882610e9f565b9050919050565b610ed881610ebe565b8114610ee2575f80fd5b50565b5f81359050610ef381610ecf565b92915050565b5f819050919050565b610f0b81610ef9565b8114610f15575f80fd5b50565b5f81359050610f2681610f02565b92915050565b5f8060408385031215610f4257610f41610e9b565b5b5f610f4f85828601610ee5565b9250506020610f6085828601610f18565b9150509250929050565b5f8115159050919050565b610f7e81610f6a565b82525050565b5f602082019050610f975f830184610f75565b92915050565b610fa681610ef9565b82525050565b5f602082019050610fbf5f830184610f9d565b92915050565b5f805f60608486031215610fdc57610fdb610e9b565b5b5f610fe986828701610ee5565b9350506020610ffa86828701610ee5565b925050604061100b86828701610f18565b9150509250925092565b5f60ff82169050919050565b61102a81611015565b82525050565b5f6020820190506110435f830184611021565b92915050565b5f6020828403121561105e5761105d610e9b565b5b5f61106b84828501610ee5565b91505092915050565b61107d81610ebe565b82525050565b5f6020820190506110965f830184611074565b92915050565b5f80604083850312156110b2576110b1610e9b565b5b5f6110bf85828601610ee5565b92505060206110d085828601610ee5565b9150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061111e57607f821691505b602082108103611131576111306110da565b5b50919050565b5f60608201905061114a5f830186611074565b6111576020830185610f9d565b6111646040830184610f9d565b949350505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f6111a382610ef9565b91506111ae83610ef9565b92508282019050808211156111c6576111c561116c565b5b9291505056fea26469706673582212201bcf72ca6e2fec90c156b43c964a6dcd666873fd40e60e4eaff253aae59d065564736f6c63430008140033",
}

// OutcomeTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use OutcomeTokenMetaData.ABI instead.
var OutcomeTokenABI = OutcomeTokenMetaData.ABI

// OutcomeTokenBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OutcomeTokenMetaData.Bin instead.
var OutcomeTokenBin = OutcomeTokenMetaData.Bin

// DeployOutcomeToken deploys a new Ethereum contract, binding an instance of OutcomeToken to it.
func DeployOutcomeToken(auth *bind.TransactOpts, backend bind.ContractBackend, name_ string, symbol_ string, market common.Address) (common.Address, *types.Transaction, *OutcomeToken, error) {
	parsed, err := OutcomeTokenMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OutcomeTokenBin), backend, name_, symbol_, market)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OutcomeToken{OutcomeTokenCaller: OutcomeTokenCaller{contract: contract}, OutcomeTokenTransactor: OutcomeTokenTransactor{contract: contract}, OutcomeTokenFilterer: OutcomeTokenFilterer{contract: contract}}, nil
}

// OutcomeToken is an auto generated Go binding around an Ethereum contract.
type OutcomeToken struct {
	OutcomeTokenCaller     // Read-only binding to the contract
	OutcomeTokenTransactor // Write-only binding to the contract
	OutcomeTokenFilterer   // Log filterer for contract events
}

// OutcomeTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type OutcomeTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OutcomeTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OutcomeTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OutcomeTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OutcomeTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OutcomeTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OutcomeTokenSession struct {
	Contract     *OutcomeToken     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OutcomeTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OutcomeTokenCallerSession struct {
	Contract *OutcomeTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// OutcomeTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OutcomeTokenTransactorSession struct {
	Contract     *OutcomeTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// OutcomeTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type OutcomeTokenRaw struct {
	Contract *OutcomeToken // Generic contract binding to access the raw methods on
}

// OutcomeTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OutcomeTokenCallerRaw struct {
	Contract *OutcomeTokenCaller // Generic read-only contract binding to access the raw methods on
}

// OutcomeTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OutcomeTokenTransactorRaw struct {
	Contract *OutcomeTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOutcomeToken creates a new instance of OutcomeToken, bound to a specific deployed contract.
func NewOutcomeToken(address common.Address, backend bind.ContractBackend) (*OutcomeToken, error) {
	contract, err := bindOutcomeToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OutcomeToken{OutcomeTokenCaller: OutcomeTokenCaller{contract: contract}, OutcomeTokenTransactor: OutcomeTokenTransactor{contract: contract}, OutcomeTokenFilterer: OutcomeTokenFilterer{contract: contract}}, nil
}

// NewOutcomeTokenCaller creates a new read-only instance of OutcomeToken, bound to a specific deployed contract.
func NewOutcomeTokenCaller(address common.Address, caller bind.ContractCaller) (*OutcomeTokenCaller, error) {
	contract, err := bindOutcomeToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OutcomeTokenCaller{contract: contract}, nil
}

// NewOutcomeTokenTransactor creates a new write-only instance of OutcomeToken, bound to a specific deployed contract.
func NewOutcomeTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*OutcomeTokenTransactor, error) {
	contract, err := bindOutcomeToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OutcomeTokenTransactor{contract: contract}, nil
}

// NewOutcomeTokenFilterer creates a new log filterer instance of OutcomeToken, bound to a specific deployed contract.
func NewOutcomeTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*OutcomeTokenFilterer, error) {
	contract, err := bindOutcomeToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OutcomeTokenFilterer{contract: contract}, nil
}

// bindOutcomeToken binds a generic wrapper to an already deployed contract.
func bindOutcomeToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OutcomeTokenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OutcomeToken *OutcomeTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OutcomeToken.Contract.OutcomeTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OutcomeToken *OutcomeTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OutcomeToken.Contract.OutcomeTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OutcomeToken *OutcomeTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OutcomeToken.Contract.OutcomeTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OutcomeToken *OutcomeTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OutcomeToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OutcomeToken *OutcomeTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OutcomeToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OutcomeToken *OutcomeTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OutcomeToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_OutcomeToken *OutcomeTokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _OutcomeToken.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_OutcomeToken *OutcomeTokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _OutcomeToken.Contract.Allowance(&_OutcomeToken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_OutcomeToken *OutcomeTokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _OutcomeToken.Contract.Allowance(&_OutcomeToken.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_OutcomeToken *OutcomeTokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _OutcomeToken.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_OutcomeToken *OutcomeTokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _OutcomeToken.Contract.BalanceOf(&_OutcomeToken.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_OutcomeToken *OutcomeTokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _OutcomeToken.Contract.BalanceOf(&_OutcomeToken.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_OutcomeToken *OutcomeTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _OutcomeToken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_OutcomeToken *OutcomeTokenSession) Decimals() (uint8, error) {
	return _OutcomeToken.Contract.Decimals(&_OutcomeToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_OutcomeToken *OutcomeTokenCallerSession) Decimals() (uint8, error) {
	return _OutcomeToken.Contract.Decimals(&_OutcomeToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_OutcomeToken *OutcomeTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _OutcomeToken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_OutcomeToken *OutcomeTokenSession) Name() (string, error) {
	return _OutcomeToken.Contract.Name(&_OutcomeToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_OutcomeToken *OutcomeTokenCallerSession) Name() (string, error) {
	return _OutcomeToken.Contract.Name(&_OutcomeToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OutcomeToken *OutcomeTokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OutcomeToken.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OutcomeToken *OutcomeTokenSession) Owner() (common.Address, error) {
	return _OutcomeToken.Contract.Owner(&_OutcomeToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OutcomeToken *OutcomeTokenCallerSession) Owner() (common.Address, error) {
	return _OutcomeToken.Contract.Owner(&_OutcomeToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_OutcomeToken *OutcomeTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _OutcomeToken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_OutcomeToken *OutcomeTokenSession) Symbol() (string, error) {
	return _OutcomeToken.Contract.Symbol(&_OutcomeToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_OutcomeToken *OutcomeTokenCallerSession) Symbol() (string, error) {
	return _OutcomeToken.Contract.Symbol(&_OutcomeToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_OutcomeToken *OutcomeTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OutcomeToken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_OutcomeToken *OutcomeTokenSession) TotalSupply() (*big.Int, error) {
	return _OutcomeToken.Contract.TotalSupply(&_OutcomeToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_OutcomeToken *OutcomeTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _OutcomeToken.Contract.TotalSupply(&_OutcomeToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_OutcomeToken *OutcomeTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _OutcomeToken.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_OutcomeToken *OutcomeTokenSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _OutcomeToken.Contract.Approve(&_OutcomeToken.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_OutcomeToken *OutcomeTokenTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _OutcomeToken.Contract.Approve(&_OutcomeToken.TransactOpts, spender, value)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address from, uint256 amount) returns()
func (_OutcomeToken *OutcomeTokenTransactor) Burn(opts *bind.TransactOpts, from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _OutcomeToken.contract.Transact(opts, "burn", from, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address from, uint256 amount) returns()
func (_OutcomeToken *OutcomeTokenSession) Burn(from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _OutcomeToken.Contract.Burn(&_OutcomeToken.TransactOpts, from, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address from, uint256 amount) returns()
func (_OutcomeToken *OutcomeTokenTransactorSession) Burn(from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _OutcomeToken.Contract.Burn(&_OutcomeToken.TransactOpts, from, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_OutcomeToken *OutcomeTokenTransactor) Mint(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _OutcomeToken.contract.Transact(opts, "mint", to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_OutcomeToken *OutcomeTokenSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _OutcomeToken.Contract.Mint(&_OutcomeToken.TransactOpts, to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_OutcomeToken *OutcomeTokenTransactorSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _OutcomeToken.Contract.Mint(&_OutcomeToken.TransactOpts, to, amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OutcomeToken *OutcomeTokenTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OutcomeToken.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OutcomeToken *OutcomeTokenSession) RenounceOwnership() (*types.Transaction, error) {
	return _OutcomeToken.Contract.RenounceOwnership(&_OutcomeToken.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OutcomeToken *OutcomeTokenTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _OutcomeToken.Contract.RenounceOwnership(&_OutcomeToken.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_OutcomeToken *OutcomeTokenTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _OutcomeToken.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_OutcomeToken *OutcomeTokenSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _OutcomeToken.Contract.Transfer(&_OutcomeToken.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_OutcomeToken *OutcomeTokenTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _OutcomeToken.Contract.Transfer(&_OutcomeToken.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_OutcomeToken *OutcomeTokenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _OutcomeToken.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_OutcomeToken *OutcomeTokenSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _OutcomeToken.Contract.TransferFrom(&_OutcomeToken.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_OutcomeToken *OutcomeTokenTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _OutcomeToken.Contract.TransferFrom(&_OutcomeToken.TransactOpts, from, to, value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OutcomeToken *OutcomeTokenTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _OutcomeToken.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OutcomeToken *OutcomeTokenSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _OutcomeToken.Contract.TransferOwnership(&_OutcomeToken.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OutcomeToken *OutcomeTokenTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _OutcomeToken.Contract.TransferOwnership(&_OutcomeToken.TransactOpts, newOwner)
}

// OutcomeTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the OutcomeToken contract.
type OutcomeTokenApprovalIterator struct {
	Event *OutcomeTokenApproval // Event containing the contract specifics and raw log

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
func (it *OutcomeTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OutcomeTokenApproval)
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
		it.Event = new(OutcomeTokenApproval)
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
func (it *OutcomeTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OutcomeTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OutcomeTokenApproval represents a Approval event raised by the OutcomeToken contract.
type OutcomeTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_OutcomeToken *OutcomeTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*OutcomeTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _OutcomeToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &OutcomeTokenApprovalIterator{contract: _OutcomeToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_OutcomeToken *OutcomeTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *OutcomeTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _OutcomeToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OutcomeTokenApproval)
				if err := _OutcomeToken.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_OutcomeToken *OutcomeTokenFilterer) ParseApproval(log types.Log) (*OutcomeTokenApproval, error) {
	event := new(OutcomeTokenApproval)
	if err := _OutcomeToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OutcomeTokenOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the OutcomeToken contract.
type OutcomeTokenOwnershipTransferredIterator struct {
	Event *OutcomeTokenOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OutcomeTokenOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OutcomeTokenOwnershipTransferred)
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
		it.Event = new(OutcomeTokenOwnershipTransferred)
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
func (it *OutcomeTokenOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OutcomeTokenOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OutcomeTokenOwnershipTransferred represents a OwnershipTransferred event raised by the OutcomeToken contract.
type OutcomeTokenOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OutcomeToken *OutcomeTokenFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OutcomeTokenOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _OutcomeToken.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OutcomeTokenOwnershipTransferredIterator{contract: _OutcomeToken.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OutcomeToken *OutcomeTokenFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OutcomeTokenOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _OutcomeToken.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OutcomeTokenOwnershipTransferred)
				if err := _OutcomeToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_OutcomeToken *OutcomeTokenFilterer) ParseOwnershipTransferred(log types.Log) (*OutcomeTokenOwnershipTransferred, error) {
	event := new(OutcomeTokenOwnershipTransferred)
	if err := _OutcomeToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OutcomeTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the OutcomeToken contract.
type OutcomeTokenTransferIterator struct {
	Event *OutcomeTokenTransfer // Event containing the contract specifics and raw log

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
func (it *OutcomeTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OutcomeTokenTransfer)
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
		it.Event = new(OutcomeTokenTransfer)
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
func (it *OutcomeTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OutcomeTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OutcomeTokenTransfer represents a Transfer event raised by the OutcomeToken contract.
type OutcomeTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_OutcomeToken *OutcomeTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OutcomeTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OutcomeToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &OutcomeTokenTransferIterator{contract: _OutcomeToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_OutcomeToken *OutcomeTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *OutcomeTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OutcomeToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OutcomeTokenTransfer)
				if err := _OutcomeToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_OutcomeToken *OutcomeTokenFilterer) ParseTransfer(log types.Log) (*OutcomeTokenTransfer, error) {
	event := new(OutcomeTokenTransfer)
	if err := _OutcomeToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
