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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"}],\"name\":\"DataStored\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"}],\"name\":\"DataUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"getData\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getUserAtIndex\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUserCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_active\",\"type\":\"bool\"}],\"name\":\"storeData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_active\",\"type\":\"bool\"}],\"name\":\"updateData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b50610d008061001d5f395ff3fe608060405234801561000f575f80fd5b5060043610610055575f3560e01c80633bc5de30146100595780635e5bad1c1461007a578063b5cb15f714610096578063e763e87a146100b4578063ffcc7bbf146100d0575b5f80fd5b610061610100565b60405161007194939291906105e5565b60405180910390f35b610094600480360381019061008f91906107c0565b6101fb565b005b61009e61033c565b6040516100ab919061082c565b60405180910390f35b6100ce60048036038101906100c991906107c0565b610348565b005b6100ea60048036038101906100e59190610845565b61049e565b6040516100f791906108af565b60405180910390f35b60605f805f805f803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f209050805f018160010154826002015f9054906101000a900460ff16836003015483805461016e906108f5565b80601f016020809104026020016040519081016040528092919081815260200182805461019a906108f5565b80156101e55780601f106101bc576101008083540402835291602001916101e5565b820191905f5260205f20905b8154815290600101906020018083116101c857829003601f168201915b5050505050935094509450945094505090919293565b5f805f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20905083815f01908161024a9190610ac2565b5082816001018190555081816002015f6101000a81548160ff021916908315150217905550428160030181905550428160030154036102e457600133908060018154018082558091505060019003905f5260205f20015f9091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b3373ffffffffffffffffffffffffffffffffffffffff167fc11e999d9949879f09858ee0a5c407c225ded67bf476549218a4d1dd9915f37f85858560405161032e93929190610b91565b60405180910390a250505050565b5f600180549050905090565b5f805f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2060030154116103c9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103c090610c17565b60405180910390fd5b5f805f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20905083815f0190816104189190610ac2565b5082816001018190555081816002015f6101000a81548160ff0219169083151502179055504281600301819055503373ffffffffffffffffffffffffffffffffffffffff167f9b32caaae608eb0061a483eb5c20359d6fb67ee030ebcbd9536d90308b4cdae485858560405161049093929190610b91565b60405180910390a250505050565b5f60018054905082106104e6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104dd90610c7f565b60405180910390fd5b600182815481106104fa576104f9610c9d565b5b905f5260205f20015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b5f81519050919050565b5f82825260208201905092915050565b5f5b83811015610560578082015181840152602081019050610545565b5f8484015250505050565b5f601f19601f8301169050919050565b5f61058582610529565b61058f8185610533565b935061059f818560208601610543565b6105a88161056b565b840191505092915050565b5f819050919050565b6105c5816105b3565b82525050565b5f8115159050919050565b6105df816105cb565b82525050565b5f6080820190508181035f8301526105fd818761057b565b905061060c60208301866105bc565b61061960408301856105d6565b61062660608301846105bc565b95945050505050565b5f604051905090565b5f80fd5b5f80fd5b5f80fd5b5f80fd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b61067e8261056b565b810181811067ffffffffffffffff8211171561069d5761069c610648565b5b80604052505050565b5f6106af61062f565b90506106bb8282610675565b919050565b5f67ffffffffffffffff8211156106da576106d9610648565b5b6106e38261056b565b9050602081019050919050565b828183375f83830152505050565b5f61071061070b846106c0565b6106a6565b90508281526020810184848401111561072c5761072b610644565b5b6107378482856106f0565b509392505050565b5f82601f83011261075357610752610640565b5b81356107638482602086016106fe565b91505092915050565b610775816105b3565b811461077f575f80fd5b50565b5f813590506107908161076c565b92915050565b61079f816105cb565b81146107a9575f80fd5b50565b5f813590506107ba81610796565b92915050565b5f805f606084860312156107d7576107d6610638565b5b5f84013567ffffffffffffffff8111156107f4576107f361063c565b5b6108008682870161073f565b935050602061081186828701610782565b9250506040610822868287016107ac565b9150509250925092565b5f60208201905061083f5f8301846105bc565b92915050565b5f6020828403121561085a57610859610638565b5b5f61086784828501610782565b91505092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f61089982610870565b9050919050565b6108a98161088f565b82525050565b5f6020820190506108c25f8301846108a0565b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061090c57607f821691505b60208210810361091f5761091e6108c8565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026109817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610946565b61098b8683610946565b95508019841693508086168417925050509392505050565b5f819050919050565b5f6109c66109c16109bc846105b3565b6109a3565b6105b3565b9050919050565b5f819050919050565b6109df836109ac565b6109f36109eb826109cd565b848454610952565b825550505050565b5f90565b610a076109fb565b610a128184846109d6565b505050565b5b81811015610a3557610a2a5f826109ff565b600181019050610a18565b5050565b601f821115610a7a57610a4b81610925565b610a5484610937565b81016020851015610a63578190505b610a77610a6f85610937565b830182610a17565b50505b505050565b5f82821c905092915050565b5f610a9a5f1984600802610a7f565b1980831691505092915050565b5f610ab28383610a8b565b9150826002028217905092915050565b610acb82610529565b67ffffffffffffffff811115610ae457610ae3610648565b5b610aee82546108f5565b610af9828285610a39565b5f60209050601f831160018114610b2a575f8415610b18578287015190505b610b228582610aa7565b865550610b89565b601f198416610b3886610925565b5f5b82811015610b5f57848901518255600182019150602085019450602081019050610b3a565b86831015610b7c5784890151610b78601f891682610a8b565b8355505b6001600288020188555050505b505050505050565b5f6060820190508181035f830152610ba9818661057b565b9050610bb860208301856105bc565b610bc560408301846105d6565b949350505050565b7f4e6f20646174612065786973747320666f7220746869732075736572000000005f82015250565b5f610c01601c83610533565b9150610c0c82610bcd565b602082019050919050565b5f6020820190508181035f830152610c2e81610bf5565b9050919050565b7f496e646578206f7574206f6620626f756e6473000000000000000000000000005f82015250565b5f610c69601383610533565b9150610c7482610c35565b602082019050919050565b5f6020820190508181035f830152610c9681610c5d565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffdfea2646970667358221220910e4ab9e07de72c44668e5037ba56fd0e5ae227b111b7b23920423c40abe46564736f6c63430008170033",
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

// GetData is a free data retrieval call binding the contract method 0x3bc5de30.
//
// Solidity: function getData() view returns(string name, uint256 value, bool active, uint256 timestamp)
func (_Solidity *SolidityCaller) GetData(opts *bind.CallOpts) (struct {
	Name      string
	Value     *big.Int
	Active    bool
	Timestamp *big.Int
}, error) {
	var out []interface{}
	err := _Solidity.contract.Call(opts, &out, "getData")

	outstruct := new(struct {
		Name      string
		Value     *big.Int
		Active    bool
		Timestamp *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Name = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Value = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Active = *abi.ConvertType(out[2], new(bool)).(*bool)
	outstruct.Timestamp = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetData is a free data retrieval call binding the contract method 0x3bc5de30.
//
// Solidity: function getData() view returns(string name, uint256 value, bool active, uint256 timestamp)
func (_Solidity *SoliditySession) GetData() (struct {
	Name      string
	Value     *big.Int
	Active    bool
	Timestamp *big.Int
}, error) {
	return _Solidity.Contract.GetData(&_Solidity.CallOpts)
}

// GetData is a free data retrieval call binding the contract method 0x3bc5de30.
//
// Solidity: function getData() view returns(string name, uint256 value, bool active, uint256 timestamp)
func (_Solidity *SolidityCallerSession) GetData() (struct {
	Name      string
	Value     *big.Int
	Active    bool
	Timestamp *big.Int
}, error) {
	return _Solidity.Contract.GetData(&_Solidity.CallOpts)
}

// GetUserAtIndex is a free data retrieval call binding the contract method 0xffcc7bbf.
//
// Solidity: function getUserAtIndex(uint256 _index) view returns(address)
func (_Solidity *SolidityCaller) GetUserAtIndex(opts *bind.CallOpts, _index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Solidity.contract.Call(opts, &out, "getUserAtIndex", _index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetUserAtIndex is a free data retrieval call binding the contract method 0xffcc7bbf.
//
// Solidity: function getUserAtIndex(uint256 _index) view returns(address)
func (_Solidity *SoliditySession) GetUserAtIndex(_index *big.Int) (common.Address, error) {
	return _Solidity.Contract.GetUserAtIndex(&_Solidity.CallOpts, _index)
}

// GetUserAtIndex is a free data retrieval call binding the contract method 0xffcc7bbf.
//
// Solidity: function getUserAtIndex(uint256 _index) view returns(address)
func (_Solidity *SolidityCallerSession) GetUserAtIndex(_index *big.Int) (common.Address, error) {
	return _Solidity.Contract.GetUserAtIndex(&_Solidity.CallOpts, _index)
}

// GetUserCount is a free data retrieval call binding the contract method 0xb5cb15f7.
//
// Solidity: function getUserCount() view returns(uint256)
func (_Solidity *SolidityCaller) GetUserCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Solidity.contract.Call(opts, &out, "getUserCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserCount is a free data retrieval call binding the contract method 0xb5cb15f7.
//
// Solidity: function getUserCount() view returns(uint256)
func (_Solidity *SoliditySession) GetUserCount() (*big.Int, error) {
	return _Solidity.Contract.GetUserCount(&_Solidity.CallOpts)
}

// GetUserCount is a free data retrieval call binding the contract method 0xb5cb15f7.
//
// Solidity: function getUserCount() view returns(uint256)
func (_Solidity *SolidityCallerSession) GetUserCount() (*big.Int, error) {
	return _Solidity.Contract.GetUserCount(&_Solidity.CallOpts)
}

// StoreData is a paid mutator transaction binding the contract method 0x5e5bad1c.
//
// Solidity: function storeData(string _name, uint256 _value, bool _active) returns()
func (_Solidity *SolidityTransactor) StoreData(opts *bind.TransactOpts, _name string, _value *big.Int, _active bool) (*types.Transaction, error) {
	return _Solidity.contract.Transact(opts, "storeData", _name, _value, _active)
}

// StoreData is a paid mutator transaction binding the contract method 0x5e5bad1c.
//
// Solidity: function storeData(string _name, uint256 _value, bool _active) returns()
func (_Solidity *SoliditySession) StoreData(_name string, _value *big.Int, _active bool) (*types.Transaction, error) {
	return _Solidity.Contract.StoreData(&_Solidity.TransactOpts, _name, _value, _active)
}

// StoreData is a paid mutator transaction binding the contract method 0x5e5bad1c.
//
// Solidity: function storeData(string _name, uint256 _value, bool _active) returns()
func (_Solidity *SolidityTransactorSession) StoreData(_name string, _value *big.Int, _active bool) (*types.Transaction, error) {
	return _Solidity.Contract.StoreData(&_Solidity.TransactOpts, _name, _value, _active)
}

// UpdateData is a paid mutator transaction binding the contract method 0xe763e87a.
//
// Solidity: function updateData(string _name, uint256 _value, bool _active) returns()
func (_Solidity *SolidityTransactor) UpdateData(opts *bind.TransactOpts, _name string, _value *big.Int, _active bool) (*types.Transaction, error) {
	return _Solidity.contract.Transact(opts, "updateData", _name, _value, _active)
}

// UpdateData is a paid mutator transaction binding the contract method 0xe763e87a.
//
// Solidity: function updateData(string _name, uint256 _value, bool _active) returns()
func (_Solidity *SoliditySession) UpdateData(_name string, _value *big.Int, _active bool) (*types.Transaction, error) {
	return _Solidity.Contract.UpdateData(&_Solidity.TransactOpts, _name, _value, _active)
}

// UpdateData is a paid mutator transaction binding the contract method 0xe763e87a.
//
// Solidity: function updateData(string _name, uint256 _value, bool _active) returns()
func (_Solidity *SolidityTransactorSession) UpdateData(_name string, _value *big.Int, _active bool) (*types.Transaction, error) {
	return _Solidity.Contract.UpdateData(&_Solidity.TransactOpts, _name, _value, _active)
}

// SolidityDataStoredIterator is returned from FilterDataStored and is used to iterate over the raw logs and unpacked data for DataStored events raised by the Solidity contract.
type SolidityDataStoredIterator struct {
	Event *SolidityDataStored // Event containing the contract specifics and raw log

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
func (it *SolidityDataStoredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SolidityDataStored)
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
		it.Event = new(SolidityDataStored)
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
func (it *SolidityDataStoredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SolidityDataStoredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SolidityDataStored represents a DataStored event raised by the Solidity contract.
type SolidityDataStored struct {
	User   common.Address
	Name   string
	Value  *big.Int
	Active bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDataStored is a free log retrieval operation binding the contract event 0xc11e999d9949879f09858ee0a5c407c225ded67bf476549218a4d1dd9915f37f.
//
// Solidity: event DataStored(address indexed user, string name, uint256 value, bool active)
func (_Solidity *SolidityFilterer) FilterDataStored(opts *bind.FilterOpts, user []common.Address) (*SolidityDataStoredIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Solidity.contract.FilterLogs(opts, "DataStored", userRule)
	if err != nil {
		return nil, err
	}
	return &SolidityDataStoredIterator{contract: _Solidity.contract, event: "DataStored", logs: logs, sub: sub}, nil
}

// WatchDataStored is a free log subscription operation binding the contract event 0xc11e999d9949879f09858ee0a5c407c225ded67bf476549218a4d1dd9915f37f.
//
// Solidity: event DataStored(address indexed user, string name, uint256 value, bool active)
func (_Solidity *SolidityFilterer) WatchDataStored(opts *bind.WatchOpts, sink chan<- *SolidityDataStored, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Solidity.contract.WatchLogs(opts, "DataStored", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SolidityDataStored)
				if err := _Solidity.contract.UnpackLog(event, "DataStored", log); err != nil {
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

// ParseDataStored is a log parse operation binding the contract event 0xc11e999d9949879f09858ee0a5c407c225ded67bf476549218a4d1dd9915f37f.
//
// Solidity: event DataStored(address indexed user, string name, uint256 value, bool active)
func (_Solidity *SolidityFilterer) ParseDataStored(log types.Log) (*SolidityDataStored, error) {
	event := new(SolidityDataStored)
	if err := _Solidity.contract.UnpackLog(event, "DataStored", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SolidityDataUpdatedIterator is returned from FilterDataUpdated and is used to iterate over the raw logs and unpacked data for DataUpdated events raised by the Solidity contract.
type SolidityDataUpdatedIterator struct {
	Event *SolidityDataUpdated // Event containing the contract specifics and raw log

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
func (it *SolidityDataUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SolidityDataUpdated)
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
		it.Event = new(SolidityDataUpdated)
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
func (it *SolidityDataUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SolidityDataUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SolidityDataUpdated represents a DataUpdated event raised by the Solidity contract.
type SolidityDataUpdated struct {
	User   common.Address
	Name   string
	Value  *big.Int
	Active bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDataUpdated is a free log retrieval operation binding the contract event 0x9b32caaae608eb0061a483eb5c20359d6fb67ee030ebcbd9536d90308b4cdae4.
//
// Solidity: event DataUpdated(address indexed user, string name, uint256 value, bool active)
func (_Solidity *SolidityFilterer) FilterDataUpdated(opts *bind.FilterOpts, user []common.Address) (*SolidityDataUpdatedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Solidity.contract.FilterLogs(opts, "DataUpdated", userRule)
	if err != nil {
		return nil, err
	}
	return &SolidityDataUpdatedIterator{contract: _Solidity.contract, event: "DataUpdated", logs: logs, sub: sub}, nil
}

// WatchDataUpdated is a free log subscription operation binding the contract event 0x9b32caaae608eb0061a483eb5c20359d6fb67ee030ebcbd9536d90308b4cdae4.
//
// Solidity: event DataUpdated(address indexed user, string name, uint256 value, bool active)
func (_Solidity *SolidityFilterer) WatchDataUpdated(opts *bind.WatchOpts, sink chan<- *SolidityDataUpdated, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Solidity.contract.WatchLogs(opts, "DataUpdated", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SolidityDataUpdated)
				if err := _Solidity.contract.UnpackLog(event, "DataUpdated", log); err != nil {
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

// ParseDataUpdated is a log parse operation binding the contract event 0x9b32caaae608eb0061a483eb5c20359d6fb67ee030ebcbd9536d90308b4cdae4.
//
// Solidity: event DataUpdated(address indexed user, string name, uint256 value, bool active)
func (_Solidity *SolidityFilterer) ParseDataUpdated(log types.Log) (*SolidityDataUpdated, error) {
	event := new(SolidityDataUpdated)
	if err := _Solidity.contract.UnpackLog(event, "DataUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
