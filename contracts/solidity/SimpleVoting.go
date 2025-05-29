// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package solidity

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

// SolidityMetaData contains all meta data concerning the Solidity contract.
var SolidityMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"name\":\"ProposalCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"ProposalExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"VoteCast\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"chairperson\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_description\",\"type\":\"string\"}],\"name\":\"createProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_proposalId\",\"type\":\"uint256\"}],\"name\":\"executeProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_proposalId\",\"type\":\"uint256\"}],\"name\":\"getProposal\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"voteCount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"executed\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getProposalCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"proposals\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"voteCount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"executed\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_proposalId\",\"type\":\"uint256\"}],\"name\":\"vote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"voters\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"hasVoted\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"votedProposal\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b50335f806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506112bd8061005c5f395ff3fe608060405234801561000f575f80fd5b5060043610610086575f3560e01c806349c2a1a61161005957806349c2a1a614610112578063a3ec138d1461012e578063c08cc02d1461015f578063c7f758a81461017d57610086565b80630121b93f1461008a578063013cf08b146100a65780630d61b519146100d85780632e4176cf146100f4575b5f80fd5b6100a4600480360381019061009f9190610969565b6101af565b005b6100c060048036038101906100bb9190610969565b6103cd565b6040516100cf93929190610a47565b60405180910390f35b6100f260048036038101906100ed9190610969565b610494565b005b6100fc61064d565b6040516101099190610ac2565b60405180910390f35b61012c60048036038101906101279190610c07565b610670565b005b61014860048036038101906101439190610c78565b6107ce565b604051610156929190610ca3565b60405180910390f35b6101676107fa565b6040516101749190610cca565b60405180910390f35b61019760048036038101906101929190610969565b610806565b6040516101a693929190610a47565b60405180910390f35b60015f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f015f9054906101000a900460ff161561023b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161023290610d2d565b60405180910390fd5b6002805490508110610282576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161027990610d95565b60405180910390fd5b6002818154811061029657610295610db3565b5b905f5260205f2090600302016002015f9054906101000a900460ff16156102f2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102e990610e2a565b60405180910390fd5b5f60015f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2090506001815f015f6101000a81548160ff02191690831515021790555081816001018190555060016002838154811061036c5761036b610db3565b5b905f5260205f2090600302016001015f8282546103899190610e75565b925050819055507fa36cc2bebb74db33e9f88110a07ef56e1b31b24b4c4f51b54b1664266e29f45b33836040516103c1929190610ea8565b60405180910390a15050565b600281815481106103dc575f80fd5b905f5260205f2090600302015f91509050805f0180546103fb90610efc565b80601f016020809104026020016040519081016040528092919081815260200182805461042790610efc565b80156104725780601f1061044957610100808354040283529160200191610472565b820191905f5260205f20905b81548152906001019060200180831161045557829003601f168201915b505050505090806001015490806002015f9054906101000a900460ff16905083565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610521576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161051890610f9c565b60405180910390fd5b6002805490508110610568576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161055f90610d95565b60405180910390fd5b6002818154811061057c5761057b610db3565b5b905f5260205f2090600302016002015f9054906101000a900460ff16156105d8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105cf90610e2a565b60405180910390fd5b6001600282815481106105ee576105ed610db3565b5b905f5260205f2090600302016002015f6101000a81548160ff0219169083151502179055507f712ae1383f79ac853f8d882153778e0260ef8f03b504e2866e0593e04d2b291f816040516106429190610cca565b60405180910390a150565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f8054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146106fd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106f490610f9c565b60405180910390fd5b600260405180606001604052808381526020015f81526020015f1515815250908060018154018082558091505060019003905f5260205f2090600302015f909190919091505f820151815f0190816107559190611157565b50602082015181600101556040820151816002015f6101000a81548160ff02191690831515021790555050507f9c770c289ab5bf7e57cb1d23c8ceae993aea46eb64847072fd3d78ca60d3e43260016002805490506107b49190611226565b826040516107c3929190611259565b60405180910390a150565b6001602052805f5260405f205f91509050805f015f9054906101000a900460ff16908060010154905082565b5f600280549050905090565b60605f806002805490508410610851576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161084890610d95565b60405180910390fd5b5f6002858154811061086657610865610db3565b5b905f5260205f2090600302019050805f018160010154826002015f9054906101000a900460ff1682805461089990610efc565b80601f01602080910402602001604051908101604052809291908181526020018280546108c590610efc565b80156109105780601f106108e757610100808354040283529160200191610910565b820191905f5260205f20905b8154815290600101906020018083116108f357829003601f168201915b50505050509250935093509350509193909250565b5f604051905090565b5f80fd5b5f80fd5b5f819050919050565b61094881610936565b8114610952575f80fd5b50565b5f813590506109638161093f565b92915050565b5f6020828403121561097e5761097d61092e565b5b5f61098b84828501610955565b91505092915050565b5f81519050919050565b5f82825260208201905092915050565b5f5b838110156109cb5780820151818401526020810190506109b0565b5f8484015250505050565b5f601f19601f8301169050919050565b5f6109f082610994565b6109fa818561099e565b9350610a0a8185602086016109ae565b610a13816109d6565b840191505092915050565b610a2781610936565b82525050565b5f8115159050919050565b610a4181610a2d565b82525050565b5f6060820190508181035f830152610a5f81866109e6565b9050610a6e6020830185610a1e565b610a7b6040830184610a38565b949350505050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610aac82610a83565b9050919050565b610abc81610aa2565b82525050565b5f602082019050610ad55f830184610ab3565b92915050565b5f80fd5b5f80fd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b610b19826109d6565b810181811067ffffffffffffffff82111715610b3857610b37610ae3565b5b80604052505050565b5f610b4a610925565b9050610b568282610b10565b919050565b5f67ffffffffffffffff821115610b7557610b74610ae3565b5b610b7e826109d6565b9050602081019050919050565b828183375f83830152505050565b5f610bab610ba684610b5b565b610b41565b905082815260208101848484011115610bc757610bc6610adf565b5b610bd2848285610b8b565b509392505050565b5f82601f830112610bee57610bed610adb565b5b8135610bfe848260208601610b99565b91505092915050565b5f60208284031215610c1c57610c1b61092e565b5b5f82013567ffffffffffffffff811115610c3957610c38610932565b5b610c4584828501610bda565b91505092915050565b610c5781610aa2565b8114610c61575f80fd5b50565b5f81359050610c7281610c4e565b92915050565b5f60208284031215610c8d57610c8c61092e565b5b5f610c9a84828501610c64565b91505092915050565b5f604082019050610cb65f830185610a38565b610cc36020830184610a1e565b9392505050565b5f602082019050610cdd5f830184610a1e565b92915050565b7f416c726561647920766f746564000000000000000000000000000000000000005f82015250565b5f610d17600d8361099e565b9150610d2282610ce3565b602082019050919050565b5f6020820190508181035f830152610d4481610d0b565b9050919050565b7f496e76616c69642070726f706f73616c000000000000000000000000000000005f82015250565b5f610d7f60108361099e565b9150610d8a82610d4b565b602082019050919050565b5f6020820190508181035f830152610dac81610d73565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f50726f706f73616c20616c7265616479206578656375746564000000000000005f82015250565b5f610e1460198361099e565b9150610e1f82610de0565b602082019050919050565b5f6020820190508181035f830152610e4181610e08565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f610e7f82610936565b9150610e8a83610936565b9250828201905080821115610ea257610ea1610e48565b5b92915050565b5f604082019050610ebb5f830185610ab3565b610ec86020830184610a1e565b9392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f6002820490506001821680610f1357607f821691505b602082108103610f2657610f25610ecf565b5b50919050565b7f4f6e6c79206368616972706572736f6e2063616e2063616c6c207468697320665f8201527f756e6374696f6e00000000000000000000000000000000000000000000000000602082015250565b5f610f8660278361099e565b9150610f9182610f2c565b604082019050919050565b5f6020820190508181035f830152610fb381610f7a565b9050919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026110167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610fdb565b6110208683610fdb565b95508019841693508086168417925050509392505050565b5f819050919050565b5f61105b61105661105184610936565b611038565b610936565b9050919050565b5f819050919050565b61107483611041565b61108861108082611062565b848454610fe7565b825550505050565b5f90565b61109c611090565b6110a781848461106b565b505050565b5b818110156110ca576110bf5f82611094565b6001810190506110ad565b5050565b601f82111561110f576110e081610fba565b6110e984610fcc565b810160208510156110f8578190505b61110c61110485610fcc565b8301826110ac565b50505b505050565b5f82821c905092915050565b5f61112f5f1984600802611114565b1980831691505092915050565b5f6111478383611120565b9150826002028217905092915050565b61116082610994565b67ffffffffffffffff81111561117957611178610ae3565b5b6111838254610efc565b61118e8282856110ce565b5f60209050601f8311600181146111bf575f84156111ad578287015190505b6111b7858261113c565b86555061121e565b601f1984166111cd86610fba565b5f5b828110156111f4578489015182556001820191506020850194506020810190506111cf565b86831015611211578489015161120d601f891682611120565b8355505b6001600288020188555050505b505050505050565b5f61123082610936565b915061123b83610936565b925082820390508181111561125357611252610e48565b5b92915050565b5f60408201905061126c5f830185610a1e565b818103602083015261127e81846109e6565b9050939250505056fea2646970667358221220d54df8543eba0fce5f0948a786abc9cdf422e47577a2570eb14597b1be6197a064736f6c63430008170033",
}

// SolidityABI is the input ABI used to generate the binding from.
// Deprecated: Use SolidityMetaData.ABI instead.
var SolidityABI = SolidityMetaData.ABI

// SolidityBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SolidityMetaData.Bin instead.
var SolidityBin = SolidityMetaData.Bin

// DeploySolidity deploys a new Ethereum contract, binding an instance of Solidity to it.
func DeploySolidity(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Solidity, error) {
	parsed, err := SolidityMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SolidityBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Solidity{SolidityCaller: SolidityCaller{contract: contract}, SolidityTransactor: SolidityTransactor{contract: contract}, SolidityFilterer: SolidityFilterer{contract: contract}}, nil
}

// Solidity is an auto generated Go binding around an Ethereum contract.
type Solidity struct {
	SolidityCaller     // Read-only binding to the contract
	SolidityTransactor // Write-only binding to the contract
	SolidityFilterer   // Log filterer for contract events
}

// SolidityCaller is an auto generated read-only Go binding around an Ethereum contract.
type SolidityCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SolidityTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SolidityTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SolidityFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SolidityFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SoliditySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SoliditySession struct {
	Contract     *Solidity         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SolidityCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SolidityCallerSession struct {
	Contract *SolidityCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SolidityTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SolidityTransactorSession struct {
	Contract     *SolidityTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SolidityRaw is an auto generated low-level Go binding around an Ethereum contract.
type SolidityRaw struct {
	Contract *Solidity // Generic contract binding to access the raw methods on
}

// SolidityCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SolidityCallerRaw struct {
	Contract *SolidityCaller // Generic read-only contract binding to access the raw methods on
}

// SolidityTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SolidityTransactorRaw struct {
	Contract *SolidityTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSolidity creates a new instance of Solidity, bound to a specific deployed contract.
func NewSolidity(address common.Address, backend bind.ContractBackend) (*Solidity, error) {
	contract, err := bindSolidity(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Solidity{SolidityCaller: SolidityCaller{contract: contract}, SolidityTransactor: SolidityTransactor{contract: contract}, SolidityFilterer: SolidityFilterer{contract: contract}}, nil
}

// NewSolidityCaller creates a new read-only instance of Solidity, bound to a specific deployed contract.
func NewSolidityCaller(address common.Address, caller bind.ContractCaller) (*SolidityCaller, error) {
	contract, err := bindSolidity(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SolidityCaller{contract: contract}, nil
}

// NewSolidityTransactor creates a new write-only instance of Solidity, bound to a specific deployed contract.
func NewSolidityTransactor(address common.Address, transactor bind.ContractTransactor) (*SolidityTransactor, error) {
	contract, err := bindSolidity(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SolidityTransactor{contract: contract}, nil
}

// NewSolidityFilterer creates a new log filterer instance of Solidity, bound to a specific deployed contract.
func NewSolidityFilterer(address common.Address, filterer bind.ContractFilterer) (*SolidityFilterer, error) {
	contract, err := bindSolidity(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SolidityFilterer{contract: contract}, nil
}

// bindSolidity binds a generic wrapper to an already deployed contract.
func bindSolidity(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SolidityMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Solidity *SolidityRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Solidity.Contract.SolidityCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Solidity *SolidityRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Solidity.Contract.SolidityTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Solidity *SolidityRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Solidity.Contract.SolidityTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Solidity *SolidityCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Solidity.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Solidity *SolidityTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Solidity.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Solidity *SolidityTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Solidity.Contract.contract.Transact(opts, method, params...)
}

// Chairperson is a free data retrieval call binding the contract method 0x2e4176cf.
//
// Solidity: function chairperson() view returns(address)
func (_Solidity *SolidityCaller) Chairperson(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Solidity.contract.Call(opts, &out, "chairperson")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Chairperson is a free data retrieval call binding the contract method 0x2e4176cf.
//
// Solidity: function chairperson() view returns(address)
func (_Solidity *SoliditySession) Chairperson() (common.Address, error) {
	return _Solidity.Contract.Chairperson(&_Solidity.CallOpts)
}

// Chairperson is a free data retrieval call binding the contract method 0x2e4176cf.
//
// Solidity: function chairperson() view returns(address)
func (_Solidity *SolidityCallerSession) Chairperson() (common.Address, error) {
	return _Solidity.Contract.Chairperson(&_Solidity.CallOpts)
}

// GetProposal is a free data retrieval call binding the contract method 0xc7f758a8.
//
// Solidity: function getProposal(uint256 _proposalId) view returns(string description, uint256 voteCount, bool executed)
func (_Solidity *SolidityCaller) GetProposal(opts *bind.CallOpts, _proposalId *big.Int) (struct {
	Description string
	VoteCount   *big.Int
	Executed    bool
}, error) {
	var out []interface{}
	err := _Solidity.contract.Call(opts, &out, "getProposal", _proposalId)

	outstruct := new(struct {
		Description string
		VoteCount   *big.Int
		Executed    bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Description = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.VoteCount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Executed = *abi.ConvertType(out[2], new(bool)).(*bool)

	return *outstruct, err

}

// GetProposal is a free data retrieval call binding the contract method 0xc7f758a8.
//
// Solidity: function getProposal(uint256 _proposalId) view returns(string description, uint256 voteCount, bool executed)
func (_Solidity *SoliditySession) GetProposal(_proposalId *big.Int) (struct {
	Description string
	VoteCount   *big.Int
	Executed    bool
}, error) {
	return _Solidity.Contract.GetProposal(&_Solidity.CallOpts, _proposalId)
}

// GetProposal is a free data retrieval call binding the contract method 0xc7f758a8.
//
// Solidity: function getProposal(uint256 _proposalId) view returns(string description, uint256 voteCount, bool executed)
func (_Solidity *SolidityCallerSession) GetProposal(_proposalId *big.Int) (struct {
	Description string
	VoteCount   *big.Int
	Executed    bool
}, error) {
	return _Solidity.Contract.GetProposal(&_Solidity.CallOpts, _proposalId)
}

// GetProposalCount is a free data retrieval call binding the contract method 0xc08cc02d.
//
// Solidity: function getProposalCount() view returns(uint256)
func (_Solidity *SolidityCaller) GetProposalCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Solidity.contract.Call(opts, &out, "getProposalCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetProposalCount is a free data retrieval call binding the contract method 0xc08cc02d.
//
// Solidity: function getProposalCount() view returns(uint256)
func (_Solidity *SoliditySession) GetProposalCount() (*big.Int, error) {
	return _Solidity.Contract.GetProposalCount(&_Solidity.CallOpts)
}

// GetProposalCount is a free data retrieval call binding the contract method 0xc08cc02d.
//
// Solidity: function getProposalCount() view returns(uint256)
func (_Solidity *SolidityCallerSession) GetProposalCount() (*big.Int, error) {
	return _Solidity.Contract.GetProposalCount(&_Solidity.CallOpts)
}

// Proposals is a free data retrieval call binding the contract method 0x013cf08b.
//
// Solidity: function proposals(uint256 ) view returns(string description, uint256 voteCount, bool executed)
func (_Solidity *SolidityCaller) Proposals(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Description string
	VoteCount   *big.Int
	Executed    bool
}, error) {
	var out []interface{}
	err := _Solidity.contract.Call(opts, &out, "proposals", arg0)

	outstruct := new(struct {
		Description string
		VoteCount   *big.Int
		Executed    bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Description = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.VoteCount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Executed = *abi.ConvertType(out[2], new(bool)).(*bool)

	return *outstruct, err

}

// Proposals is a free data retrieval call binding the contract method 0x013cf08b.
//
// Solidity: function proposals(uint256 ) view returns(string description, uint256 voteCount, bool executed)
func (_Solidity *SoliditySession) Proposals(arg0 *big.Int) (struct {
	Description string
	VoteCount   *big.Int
	Executed    bool
}, error) {
	return _Solidity.Contract.Proposals(&_Solidity.CallOpts, arg0)
}

// Proposals is a free data retrieval call binding the contract method 0x013cf08b.
//
// Solidity: function proposals(uint256 ) view returns(string description, uint256 voteCount, bool executed)
func (_Solidity *SolidityCallerSession) Proposals(arg0 *big.Int) (struct {
	Description string
	VoteCount   *big.Int
	Executed    bool
}, error) {
	return _Solidity.Contract.Proposals(&_Solidity.CallOpts, arg0)
}

// Voters is a free data retrieval call binding the contract method 0xa3ec138d.
//
// Solidity: function voters(address ) view returns(bool hasVoted, uint256 votedProposal)
func (_Solidity *SolidityCaller) Voters(opts *bind.CallOpts, arg0 common.Address) (struct {
	HasVoted      bool
	VotedProposal *big.Int
}, error) {
	var out []interface{}
	err := _Solidity.contract.Call(opts, &out, "voters", arg0)

	outstruct := new(struct {
		HasVoted      bool
		VotedProposal *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.HasVoted = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.VotedProposal = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Voters is a free data retrieval call binding the contract method 0xa3ec138d.
//
// Solidity: function voters(address ) view returns(bool hasVoted, uint256 votedProposal)
func (_Solidity *SoliditySession) Voters(arg0 common.Address) (struct {
	HasVoted      bool
	VotedProposal *big.Int
}, error) {
	return _Solidity.Contract.Voters(&_Solidity.CallOpts, arg0)
}

// Voters is a free data retrieval call binding the contract method 0xa3ec138d.
//
// Solidity: function voters(address ) view returns(bool hasVoted, uint256 votedProposal)
func (_Solidity *SolidityCallerSession) Voters(arg0 common.Address) (struct {
	HasVoted      bool
	VotedProposal *big.Int
}, error) {
	return _Solidity.Contract.Voters(&_Solidity.CallOpts, arg0)
}

// CreateProposal is a paid mutator transaction binding the contract method 0x49c2a1a6.
//
// Solidity: function createProposal(string _description) returns()
func (_Solidity *SolidityTransactor) CreateProposal(opts *bind.TransactOpts, _description string) (*types.Transaction, error) {
	return _Solidity.contract.Transact(opts, "createProposal", _description)
}

// CreateProposal is a paid mutator transaction binding the contract method 0x49c2a1a6.
//
// Solidity: function createProposal(string _description) returns()
func (_Solidity *SoliditySession) CreateProposal(_description string) (*types.Transaction, error) {
	return _Solidity.Contract.CreateProposal(&_Solidity.TransactOpts, _description)
}

// CreateProposal is a paid mutator transaction binding the contract method 0x49c2a1a6.
//
// Solidity: function createProposal(string _description) returns()
func (_Solidity *SolidityTransactorSession) CreateProposal(_description string) (*types.Transaction, error) {
	return _Solidity.Contract.CreateProposal(&_Solidity.TransactOpts, _description)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0x0d61b519.
//
// Solidity: function executeProposal(uint256 _proposalId) returns()
func (_Solidity *SolidityTransactor) ExecuteProposal(opts *bind.TransactOpts, _proposalId *big.Int) (*types.Transaction, error) {
	return _Solidity.contract.Transact(opts, "executeProposal", _proposalId)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0x0d61b519.
//
// Solidity: function executeProposal(uint256 _proposalId) returns()
func (_Solidity *SoliditySession) ExecuteProposal(_proposalId *big.Int) (*types.Transaction, error) {
	return _Solidity.Contract.ExecuteProposal(&_Solidity.TransactOpts, _proposalId)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0x0d61b519.
//
// Solidity: function executeProposal(uint256 _proposalId) returns()
func (_Solidity *SolidityTransactorSession) ExecuteProposal(_proposalId *big.Int) (*types.Transaction, error) {
	return _Solidity.Contract.ExecuteProposal(&_Solidity.TransactOpts, _proposalId)
}

// Vote is a paid mutator transaction binding the contract method 0x0121b93f.
//
// Solidity: function vote(uint256 _proposalId) returns()
func (_Solidity *SolidityTransactor) Vote(opts *bind.TransactOpts, _proposalId *big.Int) (*types.Transaction, error) {
	return _Solidity.contract.Transact(opts, "vote", _proposalId)
}

// Vote is a paid mutator transaction binding the contract method 0x0121b93f.
//
// Solidity: function vote(uint256 _proposalId) returns()
func (_Solidity *SoliditySession) Vote(_proposalId *big.Int) (*types.Transaction, error) {
	return _Solidity.Contract.Vote(&_Solidity.TransactOpts, _proposalId)
}

// Vote is a paid mutator transaction binding the contract method 0x0121b93f.
//
// Solidity: function vote(uint256 _proposalId) returns()
func (_Solidity *SolidityTransactorSession) Vote(_proposalId *big.Int) (*types.Transaction, error) {
	return _Solidity.Contract.Vote(&_Solidity.TransactOpts, _proposalId)
}

// SolidityProposalCreatedIterator is returned from FilterProposalCreated and is used to iterate over the raw logs and unpacked data for ProposalCreated events raised by the Solidity contract.
type SolidityProposalCreatedIterator struct {
	Event *SolidityProposalCreated // Event containing the contract specifics and raw log

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
func (it *SolidityProposalCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SolidityProposalCreated)
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
		it.Event = new(SolidityProposalCreated)
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
func (it *SolidityProposalCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SolidityProposalCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SolidityProposalCreated represents a ProposalCreated event raised by the Solidity contract.
type SolidityProposalCreated struct {
	ProposalId  *big.Int
	Description string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterProposalCreated is a free log retrieval operation binding the contract event 0x9c770c289ab5bf7e57cb1d23c8ceae993aea46eb64847072fd3d78ca60d3e432.
//
// Solidity: event ProposalCreated(uint256 proposalId, string description)
func (_Solidity *SolidityFilterer) FilterProposalCreated(opts *bind.FilterOpts) (*SolidityProposalCreatedIterator, error) {

	logs, sub, err := _Solidity.contract.FilterLogs(opts, "ProposalCreated")
	if err != nil {
		return nil, err
	}
	return &SolidityProposalCreatedIterator{contract: _Solidity.contract, event: "ProposalCreated", logs: logs, sub: sub}, nil
}

// WatchProposalCreated is a free log subscription operation binding the contract event 0x9c770c289ab5bf7e57cb1d23c8ceae993aea46eb64847072fd3d78ca60d3e432.
//
// Solidity: event ProposalCreated(uint256 proposalId, string description)
func (_Solidity *SolidityFilterer) WatchProposalCreated(opts *bind.WatchOpts, sink chan<- *SolidityProposalCreated) (event.Subscription, error) {

	logs, sub, err := _Solidity.contract.WatchLogs(opts, "ProposalCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SolidityProposalCreated)
				if err := _Solidity.contract.UnpackLog(event, "ProposalCreated", log); err != nil {
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

// ParseProposalCreated is a log parse operation binding the contract event 0x9c770c289ab5bf7e57cb1d23c8ceae993aea46eb64847072fd3d78ca60d3e432.
//
// Solidity: event ProposalCreated(uint256 proposalId, string description)
func (_Solidity *SolidityFilterer) ParseProposalCreated(log types.Log) (*SolidityProposalCreated, error) {
	event := new(SolidityProposalCreated)
	if err := _Solidity.contract.UnpackLog(event, "ProposalCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SolidityProposalExecutedIterator is returned from FilterProposalExecuted and is used to iterate over the raw logs and unpacked data for ProposalExecuted events raised by the Solidity contract.
type SolidityProposalExecutedIterator struct {
	Event *SolidityProposalExecuted // Event containing the contract specifics and raw log

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
func (it *SolidityProposalExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SolidityProposalExecuted)
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
		it.Event = new(SolidityProposalExecuted)
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
func (it *SolidityProposalExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SolidityProposalExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SolidityProposalExecuted represents a ProposalExecuted event raised by the Solidity contract.
type SolidityProposalExecuted struct {
	ProposalId *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterProposalExecuted is a free log retrieval operation binding the contract event 0x712ae1383f79ac853f8d882153778e0260ef8f03b504e2866e0593e04d2b291f.
//
// Solidity: event ProposalExecuted(uint256 proposalId)
func (_Solidity *SolidityFilterer) FilterProposalExecuted(opts *bind.FilterOpts) (*SolidityProposalExecutedIterator, error) {

	logs, sub, err := _Solidity.contract.FilterLogs(opts, "ProposalExecuted")
	if err != nil {
		return nil, err
	}
	return &SolidityProposalExecutedIterator{contract: _Solidity.contract, event: "ProposalExecuted", logs: logs, sub: sub}, nil
}

// WatchProposalExecuted is a free log subscription operation binding the contract event 0x712ae1383f79ac853f8d882153778e0260ef8f03b504e2866e0593e04d2b291f.
//
// Solidity: event ProposalExecuted(uint256 proposalId)
func (_Solidity *SolidityFilterer) WatchProposalExecuted(opts *bind.WatchOpts, sink chan<- *SolidityProposalExecuted) (event.Subscription, error) {

	logs, sub, err := _Solidity.contract.WatchLogs(opts, "ProposalExecuted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SolidityProposalExecuted)
				if err := _Solidity.contract.UnpackLog(event, "ProposalExecuted", log); err != nil {
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

// ParseProposalExecuted is a log parse operation binding the contract event 0x712ae1383f79ac853f8d882153778e0260ef8f03b504e2866e0593e04d2b291f.
//
// Solidity: event ProposalExecuted(uint256 proposalId)
func (_Solidity *SolidityFilterer) ParseProposalExecuted(log types.Log) (*SolidityProposalExecuted, error) {
	event := new(SolidityProposalExecuted)
	if err := _Solidity.contract.UnpackLog(event, "ProposalExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SolidityVoteCastIterator is returned from FilterVoteCast and is used to iterate over the raw logs and unpacked data for VoteCast events raised by the Solidity contract.
type SolidityVoteCastIterator struct {
	Event *SolidityVoteCast // Event containing the contract specifics and raw log

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
func (it *SolidityVoteCastIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SolidityVoteCast)
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
		it.Event = new(SolidityVoteCast)
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
func (it *SolidityVoteCastIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SolidityVoteCastIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SolidityVoteCast represents a VoteCast event raised by the Solidity contract.
type SolidityVoteCast struct {
	Voter      common.Address
	ProposalId *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterVoteCast is a free log retrieval operation binding the contract event 0xa36cc2bebb74db33e9f88110a07ef56e1b31b24b4c4f51b54b1664266e29f45b.
//
// Solidity: event VoteCast(address voter, uint256 proposalId)
func (_Solidity *SolidityFilterer) FilterVoteCast(opts *bind.FilterOpts) (*SolidityVoteCastIterator, error) {

	logs, sub, err := _Solidity.contract.FilterLogs(opts, "VoteCast")
	if err != nil {
		return nil, err
	}
	return &SolidityVoteCastIterator{contract: _Solidity.contract, event: "VoteCast", logs: logs, sub: sub}, nil
}

// WatchVoteCast is a free log subscription operation binding the contract event 0xa36cc2bebb74db33e9f88110a07ef56e1b31b24b4c4f51b54b1664266e29f45b.
//
// Solidity: event VoteCast(address voter, uint256 proposalId)
func (_Solidity *SolidityFilterer) WatchVoteCast(opts *bind.WatchOpts, sink chan<- *SolidityVoteCast) (event.Subscription, error) {

	logs, sub, err := _Solidity.contract.WatchLogs(opts, "VoteCast")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SolidityVoteCast)
				if err := _Solidity.contract.UnpackLog(event, "VoteCast", log); err != nil {
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

// ParseVoteCast is a log parse operation binding the contract event 0xa36cc2bebb74db33e9f88110a07ef56e1b31b24b4c4f51b54b1664266e29f45b.
//
// Solidity: event VoteCast(address voter, uint256 proposalId)
func (_Solidity *SolidityFilterer) ParseVoteCast(log types.Log) (*SolidityVoteCast, error) {
	event := new(SolidityVoteCast)
	if err := _Solidity.contract.UnpackLog(event, "VoteCast", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
